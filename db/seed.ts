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
import { connectToDatabase } from "$lib/mongodb";
import { Collection } from "mongodb";
import type { Document } from "mongodb";

async function paginatedQuery(
  collection: Collection<Document>
): Promise<Array<Document>> {
  const size = 100;
  const count = await collection.countDocuments({});
  const batches = Math.ceil(count / size);
  const results = await Promise.all(
    Array(batches)
      .fill({})
      .map((_, index) => {
        console.log(
          `Processing batch ${index + 1} of ${batches} for ${
            collection.collectionName
          }...`
        );
        return collection
          .find({})
          .skip(index * size)
          .limit(size)
          .toArray();
      })
  );
  return results.flat();
}
function batchInserts(array: Array<any>): Array<Array<any>> {
  const size = 2000;
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
  console.log("Attempting database connection...");
  const dbConnection = await connectToDatabase();
  const mongoDb = dbConnection.db;
  console.log("Connection made!");
  const userCollection = mongoDb.collection("users");
  const users = await userCollection.find({}).toArray();
  console.log("Seeding Users...");
  await db.delete(User);
  await db.insert(User).values(
    users.map((user: any) => {
      const { _id, last_login, read_comp_rules, read_pp, read_tc, ...rest } =
        user;
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
  const carCollection = mongoDb.collection("cars");
  const cars = await carCollection.find({}).toArray();
  console.log("Seeding Cars...");
  await db.delete(Car);
  await db.insert(Car).values(
    cars.map((car: any) => {
      const { _id, created, first_sale, ...rest } = car;
      return {
        ...rest,
        created: new Date(created),
        first_sale: new Date(first_sale),
      };
    })
  );
  console.log("Cars Seeded!");
  const carClassCollection = mongoDb.collection("carclasses");
  const carClasses = await paginatedQuery(carClassCollection);
  console.log("Seeding Car Classes...");
  await db.delete(CarClass);
  await db.insert(CarClass).values(
    carClasses.map((carClass: any) => {
      const { _id, ...rest } = carClass;
      return rest;
    })
  );
  console.log("Car Classes Seeded!");
  const seasonsCollection = mongoDb.collection("seasons");
  const seasons = await paginatedQuery(seasonsCollection);
  console.log("Seeding Seasons...");
  await db.delete(Season);
  await db.insert(Season).values(
    seasons.map((season: any) => {
      const { _id, start_date, ...rest } = season;
      return {
        ...rest,
        start_date: new Date(start_date),
      };
    })
  );
  console.log("Seeded Seasons!");
  const pastSeasonsCollection = mongoDb.collection("pastseasons");
  const pastSeasons = await paginatedQuery(pastSeasonsCollection);
  console.log("Seeding Past Seasons...");
  await db.delete(PastSeason);
  await db.insert(PastSeason).values(
    pastSeasons.map((pastSeason: any) => {
      const { _id, ...rest } = pastSeason;
      return rest;
    })
  );
  console.log("Past Seasons Seeded!");
  const standingsCollection = mongoDb.collection("standings");
  const standings = await paginatedQuery(standingsCollection);
  console.log("Seeding Standings...");
  await db.delete(Standing);
  await db.insert(Standing).values(
    standings.map((standing: any) => {
      const {
        _id: id,
        car_class_id,
        season_driver_data,
        season_id,
        ...rest
      } = standing;
      return {
        id,
        car_class_id: Number(car_class_id),
        season_id: Number(season_id),
        ...season_driver_data,
        ...rest,
      };
    })
  );
  console.log("Standings Seeded!");
  const subsessionsCollection = mongoDb.collection("subsessions");
  // @ts-ignore
  const subsessions: Array<SubsessionType> = await paginatedQuery(
    subsessionsCollection
  );
  console.log("Seeding Subsessions...");
  const {
    allPracticeResults,
    allQualifyingResults,
    allRaceResults,
    allSubsessions,
  } = subsessions.reduce(
    // @todo Provide an accurate type for the unknown here
    (object: Record<string, Array<any>>, subsession) => {
      const { _id, session_results, end_time, start_time, ...rest } =
        subsession;
      session_results.forEach((sessionResult) => {
        const { results, ...rest } = sessionResult;
        results.forEach((result) => {
          const resultWithSession = {
            subsession_id: _id,
            ...result,
            ...rest,
            id: `${_id}_${rest.simsession_name}_${result.cust_id}`,
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
    const practiceResultInserts: Array<any> = [];
    allPracticeResults.forEach((result) => {
      practiceResultInserts.push(
        db.insert(SubsessionPracticeResults).values(result)
      );
    });
    await Promise.all(
      batchInserts(practiceResultInserts).map((batch, index) => {
        console.log(
          `Processing batch ${index + 1} of ${Math.ceil(
            practiceResultInserts.length / 2000
          )} for practiceResults`
        );
        // @ts-expect-error
        return db.batch(batch);
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
    const qualifyingResultInserts: Array<any> = [];
    allQualifyingResults.forEach((result) => {
      qualifyingResultInserts.push(
        db.insert(SubsessionQualifyingResults).values(result)
      );
    });
    await Promise.all(
      batchInserts(qualifyingResultInserts).map((batch, index) => {
        console.log(
          `Processing batch ${index + 1} of ${Math.ceil(
            qualifyingResultInserts.length / 2000
          )} for qualifyingResults`
        );
        // @ts-expect-error
        return db.batch(batch);
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
    const raceResultInserts: Array<any> = [];
    allRaceResults.forEach((result) => {
      raceResultInserts.push(db.insert(SubsessionRaceResults).values(result));
    });
    await Promise.all(
      batchInserts(raceResultInserts).map((batch, index) => {
        console.log(
          `Processing batch ${index + 1} of ${Math.ceil(
            raceResultInserts.length / 2000
          )} for raceResults`
        );
        // @ts-expect-error
        return db.batch(batch);
      })
    );
  }
  console.log("Subsession Race Results Seeded!");
  console.log("Subsessions Seeded!");
}
