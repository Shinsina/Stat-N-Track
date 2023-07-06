import { MongoClient, Db } from "mongodb";

declare global {
  namespace globalThis {
      var mongo: {
        conn: { client: MongoClient; db: Db; } | null;
        promise: Promise<{ client: MongoClient; db: Db; }> | null;
      };
  }
}
