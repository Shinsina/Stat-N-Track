import {
  db,
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

// https://astro.build/db/seed
export default async function seed() {
  console.log("Attempting database connection...");
  const dbConnection = await connectToDatabase();
  const mongoDb = dbConnection.db;
  console.log("Connection made!");
  const userCollection = mongoDb.collection("users");
  const users = await userCollection.find({}).toArray();
  console.log("Seeding Users...");
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
  const standingsCollection = mongoDb.collection("standings");
  const standings = await paginatedQuery(standingsCollection);
  console.log("Seeding Standings...");
  await db.insert(Standing).values(
    standings.map((standing: any) => {
      const { _id: id, season_driver_data, ...rest } = standing;
      return {
        id,
        ...season_driver_data,
        ...rest,
      };
    })
  );
  console.log("Standings Seeded!");
  const pastSeasonsCollection = mongoDb.collection("pastseasons");
  const pastSeasons = await paginatedQuery(pastSeasonsCollection);
  console.log("Seeding Past Seasons...");
  await db.insert(PastSeason).values(
    pastSeasons.map((pastSeason: any) => {
      const { _id, ...rest } = pastSeason;
      return rest;
    })
  );
  console.log("Past Seasons Seeded!");
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
  await db.insert(Subsession).values(allSubsessions);
  console.log("Seeding Subsession Practice Results...");
  await Promise.all(
    allPracticeResults.map((result) =>
      db.insert(SubsessionPracticeResults).values(result)
    )
  );
  console.log("Subsession Practice Results Seeded!");
  console.log("Seeding Subsession Qualifying Results...");
  await Promise.all(
    allQualifyingResults.map((result) =>
      db.insert(SubsessionQualifyingResults).values(result)
    )
  );
  console.log("Subsession Qualifying Results Seeded!");
  console.log("Seeding Subsession Race Results...");
  await Promise.all(
    allRaceResults.map((result) =>
      db.insert(SubsessionRaceResults).values(result)
    )
  );
  console.log("Subsession Race Results Seeded!");
  console.log("Subsessions Seeded!");
}
