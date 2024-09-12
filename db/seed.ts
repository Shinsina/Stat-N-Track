import {
  db,
  count,
  Car,
  CarClass,
  PastSeason,
  Season,
  Standing,
  Subsession,
  SubsessionPracticeResults,
  SubsessionQualifyingResults,
  SubsessionRaceResults,
  User,
} from "astro:db";
import fs from "fs";
import type { Subsession as SubsessionType } from "$lib/types";
import users from "./users.json";
import cars from "./cars.json";
import carClasses from "./car-classes.json";
import seasons from "./seasons.json";
import pastSeasons from "./past-seasons-output.json";
import standings from "./standings-output.json";
import seasonIdArray from "./distinct-season-ids.json";

const batchSize = 100;
function batchInserts(array: Array<any>): Array<Array<any>> {
  const count = array.length;
  const batches = Math.ceil(count / batchSize);
  return Array(batches)
    .fill({})
    .map((_, index) =>
      array.slice(index * batchSize, (index + 1) * batchSize).filter((v) => v)
    );
}

function delay(delayInMilliseconds: number): Promise<unknown> {
  return new Promise((resolution) =>
    setTimeout(resolution, delayInMilliseconds)
  );
}

// @todo Re-write this so it is additive or "upsert-like" rather than starting from ground 0 every time
// https://astro.build/db/seed
export default async function seed() {
  await db.delete(SubsessionRaceResults);
  await db.delete(SubsessionQualifyingResults);
  await db.delete(SubsessionPracticeResults);
  await db.delete(Subsession);
  await db.delete(Standing);
  await db.delete(PastSeason);
  await db.delete(Season);
  await db.delete(CarClass);
  await db.delete(Car);
  await db.delete(User);
  console.log("Seeding Users...");
  await db.insert(User).values(
    users.map((user: any) => {
      const { last_login, read_comp_rules, read_pp, read_tc, ...rest } = user;
      return {
        ...rest,
        last_login: new Date(last_login),
        read_comp_rules: new Date(read_comp_rules),
        read_pp: new Date(read_pp),
        read_tc: new Date(read_tc),
      };
    })
  );
  console.log("Users Seeded!");
  console.log("Seeding Cars...");
  await db.insert(Car).values(
    cars.map((car: any) => {
      const { created, first_sale, ...rest } = car;
      return {
        ...rest,
        created: new Date(created),
        // Appears the Mini Stock doesn't have this field so just assume "now" in these cases
        first_sale: first_sale ? new Date(first_sale) : new Date(),
      };
    })
  );
  console.log("Cars Seeded!");
  console.log("Seeding Car Classes...");
  await db.insert(CarClass).values(carClasses);
  console.log("Car Classes Seeded!");
  console.log("Seeding Seasons...");
  const seasonsAsArray = Array.isArray(seasons)
    ? seasons.map((season: any) => {
        const { start_date, ...rest } = season;
        return {
          ...rest,
          start_date: new Date(start_date),
        };
      })
    : [];
  await db.insert(Season).values(seasonsAsArray);
  console.log("Seeded Seasons!");
  console.log("Seeding Past Seasons...");
  const pastSeasonsAsObject = pastSeasons as Record<string, any>;
  await db.insert(PastSeason).values(
    Object.keys(pastSeasonsAsObject)
      .map((key) =>
        pastSeasonsAsObject[key].series.seasons.filter(
          (season: typeof PastSeason) =>
            seasonIdArray.includes(Number(season.season_id))
        )
      )
      .flat()
  );
  console.log("Past Seasons Seeded!");
  console.log("Seeding Standings...");
  const allStandings = standings.map((standing: any) => {
    const { car_class_id, season_driver_data, season_id, ...rest } = standing;
    return {
      id: `${season_id}_${car_class_id}_${season_driver_data.cust_id}`,
      car_class_id: Number(car_class_id),
      season_id: Number(season_id),
      ...season_driver_data,
      ...rest,
    };
  });
  await Promise.all(
    batchInserts(allStandings).map(async (batch, index) => {
      const trueIndex = index + 1;
      await delay(2000 * trueIndex);
      console.log(
        `Processing up to the ${trueIndex * batchSize}th standings result`
      );
      // The following seasons appear to have had their car classes removed:
      // Ring Meister Ricmotech Series - Fixed - 2024 Season 3
      // Weekly Race Challenge - 2024 Season 1
      return db.insert(Standing).values(batch.filter((v) => carClasses.map((v) => v.car_class_id).includes(v.car_class_id)));
    })
  );
  console.log("Standings Seeded!");
  console.log("Seeding Subsessions...");
  const subsessions: Array<SubsessionType> = [];
  for (let i = 1; i <= 10; i += 1) {
    const raw = fs.readFileSync(`./db/${i}-subsessions-output.json`);
    const batchSubsessions = JSON.parse(String(raw));
    subsessions.push(...batchSubsessions);
  }
  const {
    allPracticeResults,
    allQualifyingResults,
    allRaceResults,
    allSubsessions,
  } = subsessions.reduce(
    // @todo Provide an accurate type for the unknown here
    (object: Record<string, Array<any>>, subsession) => {
      const { session_results, end_time, start_time, ...rest } =
        subsession;
      session_results.forEach((sessionResult) => {
        const { results, ...rest } = sessionResult;
        results.forEach((result) => {
          if (!result.driver_results) {
            const resultWithSession = {
              subsession_id: subsession.subsession_id,
              ...result,
              ...rest,
              id: `${subsession.subsession_id}_${rest.simsession_name}_${result.cust_id}`,
            };
            if (rest.simsession_type_name.match("Practice")) {
              object.allPracticeResults.push(resultWithSession);
            } else if (rest.simsession_type_name.match("Qualifying")) {
              object.allQualifyingResults.push(resultWithSession);
            } else if (rest.simsession_type_name.match("Race")) {
              object.allRaceResults.push(resultWithSession);
            }
          } else if (result.driver_results && Array.isArray(result.driver_results)) {
            result.driver_results.forEach((driverResult) => {
              const resultWithSession = {
                subsession_id: subsession.subsession_id,
                ...driverResult,
                ...rest,
                // Apparently you can be entered into the EXACT SAME SPLIT on multiple teams, this accounts for that.
                // @ts-expect-error
                id: `${subsession.subsession_id}_${rest.simsession_name}_${driverResult.cust_id}_${driverResult.team_id}`,
              };
              if (rest.simsession_type_name.match("Practice")) {
                object.allPracticeResults.push(resultWithSession);
              } else if (rest.simsession_type_name.match("Qualifying")) {
                object.allQualifyingResults.push(resultWithSession);
              } else if (rest.simsession_type_name.match("Race")) {
                object.allRaceResults.push(resultWithSession);
              }
            });
          }
        });
      });
      object.allSubsessions.push({
        ...rest,
        end_time: new Date(end_time),
        start_time: new Date(start_time),
      });
      return object;
    },
    {
      allRaceResults: [],
      allQualifyingResults: [],
      allPracticeResults: [],
      allSubsessions: [],
    }
  );
  const currentSubsessions = await db
    .select({ count: count(Subsession.subsession_id) })
    .from(Subsession);
  const { count: subsessionCount } = currentSubsessions[0];
  console.log(
    `Currently have ${subsessionCount} subsessions, new pending count ${allSubsessions.length}`
  );
  if (subsessionCount !== allSubsessions.length) {
    await db.delete(Subsession);
    await Promise.all(
      batchInserts(allSubsessions).map(async (batch, index) => {
        const trueIndex = index + 1;
        await delay(2000 * trueIndex);
        console.log(
          `Processing up to the ${trueIndex * batchSize}th subsession result`
        );
        return db.insert(Subsession).values(batch);
      })
    );
    await db.insert(Subsession).values(allSubsessions);
  }
  console.log("Seeding Subsession Practice Results...");
  const currentPracticeResults = await db
    .select({ count: count(SubsessionPracticeResults.cust_id) })
    .from(SubsessionPracticeResults);
  const { count: practiceResultCount } = currentPracticeResults[0];
  console.log(
    `Currently have ${practiceResultCount} practice results, new pending count ${allPracticeResults.length}`
  );
  if (practiceResultCount !== allPracticeResults.length) {
    await db.delete(SubsessionPracticeResults);
    await Promise.all(
      batchInserts(allPracticeResults).map(async (batch, index) => {
        const trueIndex = index + 1;
        await delay(2000 * trueIndex);
        console.log(
          `Processing up to the ${trueIndex * batchSize}th practice result`
        );
        return db.insert(SubsessionPracticeResults).values(batch);
      })
    );
  }
  console.log("Subsession Practice Results Seeded!");
  console.log("Seeding Subsession Qualifying Results...");
  const currentQualifyingResults = await db
    .select({ count: count(SubsessionQualifyingResults.cust_id) })
    .from(SubsessionQualifyingResults);
  const { count: qualifyingResultCount } = currentQualifyingResults[0];
  console.log(
    `Currently have ${qualifyingResultCount} qualifying results, new pending count ${allQualifyingResults.length}`
  );
  if (qualifyingResultCount !== allQualifyingResults.length) {
    await db.delete(SubsessionQualifyingResults);
    await Promise.all(
      batchInserts(allQualifyingResults).map(async (batch, index) => {
        const trueIndex = index + 1;
        await delay(2000 * trueIndex);
        console.log(
          `Processing up to the ${trueIndex * batchSize}th qualifying result`
        );
        return db.insert(SubsessionQualifyingResults).values(batch);
      })
    );
  }
  console.log("Subsession Qualifying Results Seeded!");
  console.log("Seeding Subsession Race Results...");
  const currentRaceResults = await db
    .select({ count: count(SubsessionRaceResults.cust_id) })
    .from(SubsessionRaceResults);
  const { count: raceResultCount } = currentRaceResults[0];
  console.log(
    `Currently have ${raceResultCount} race results, new pending count ${allRaceResults.length}`
  );
  if (raceResultCount !== allRaceResults.length) {
    await db.delete(SubsessionRaceResults);
    await Promise.all(
      batchInserts(allRaceResults).map(async (batch, index) => {
        const trueIndex = index + 1;
        await delay(2000 * trueIndex);
        console.log(
          `Processing up to the ${trueIndex * batchSize}th race result`
        );
        return db.insert(SubsessionRaceResults).values(batch);
      })
    );
  }
  console.log("Subsession Race Results Seeded!");
  console.log("Subsessions Seeded!");
}
