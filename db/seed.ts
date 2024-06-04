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
import type { Subsession as SubsessionType } from "$lib/types";
import users from "./users.json";
import cars from "./cars.json";
import carClasses from "./car-classes.json";
import seasons from "./seasons.json";
import pastSeasons from "./past-seasons-output.json";
import jackStandings from "./jack-standings-output.json";
import jakeStandings from "./jake-standings-output.json";
import seasonIdArray from "./distinct-season-ids.json";
import subsessions from "./subsessions-output.json";

function batchInserts(array: Array<any>): Array<Array<any>> {
  const size = 100;
  const count = array.length;
  const batches = Math.ceil(count / size);
  return Array(batches)
    .fill({})
    .map((_, index) =>
      array.slice(index * size, (index + 1) * size).filter((v) => v)
    );
}

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
        first_sale: new Date(first_sale),
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
  const standings = [...jackStandings, ...jakeStandings];
  console.log("Seeding Standings...");
  await db.insert(Standing).values(
    standings.map((standing: any) => {
      const { car_class_id, season_driver_data, season_id, ...rest } = standing;
      return {
        id: `${season_id}_${car_class_id}_${season_driver_data.cust_id}`,
        car_class_id: Number(car_class_id),
        season_id: Number(season_id),
        ...season_driver_data,
        ...rest,
      };
    })
  );
  console.log("Standings Seeded!");
  console.log("Seeding Subsessions...");
  const subsessionsTyped = subsessions as Record<string, SubsessionType>;
  const {
    allPracticeResults,
    allQualifyingResults,
    allRaceResults,
    allSubsessions,
  } = Object.keys(subsessionsTyped).reduce(
    // @todo Provide an accurate type for the unknown here
    (object: Record<string, Array<any>>, key) => {
      const { session_results, end_time, start_time, ...rest } =
        subsessionsTyped[key];
      session_results.forEach((sessionResult) => {
        const { results, ...rest } = sessionResult;
        results.forEach((result) => {
          const resultWithSession = {
            subsession_id: key,
            ...result,
            ...rest,
            id: `${key}_${rest.simsession_name}_${result.cust_id}`,
          };
          if (rest.simsession_type_name.match("Practice")) {
            object.allPracticeResults.push(resultWithSession);
          } else if (rest.simsession_type_name.match("Qualifying")) {
            object.allQualifyingResults.push(resultWithSession);
          } else if (rest.simsession_type_name.match("Race")) {
            object.allRaceResults.push(resultWithSession);
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
      batchInserts(allSubsessions).map((batch) =>
        db.insert(Subsession).values(batch)
      )
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
      batchInserts(allPracticeResults).map((batch) =>
        db.insert(SubsessionPracticeResults).values(batch)
      )
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
      batchInserts(allQualifyingResults).map((batch) =>
        db.insert(SubsessionQualifyingResults).values(batch)
      )
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
      batchInserts(allRaceResults).map((batch) =>
        db.insert(SubsessionRaceResults).values(batch)
      )
    );
  }
  console.log("Subsession Race Results Seeded!");
  console.log("Subsessions Seeded!");
}
