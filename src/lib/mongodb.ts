import { MongoClient } from "mongodb";

if (!import.meta.env.MONGODB_URI) {
  throw new Error(
    "Please define the MONGODB_URI environment variable inside a root .env file"
  );
}

if (!import.meta.env.MONGODB_DB) {
  throw new Error(
    "Please define the MONGODB_DB environment variable inside a root .env file"
  );
}

/**
 * Global is used here to maintain a cached connection across hot reloads
 * in development. This prevents connections growing exponentially
 * during API Route usage.
 */
let cached = global.mongo;

if (!cached) {
  cached = global.mongo = { conn: null, promise: null };
}

export async function connectToDatabase() {
  if (cached.conn) {
    return cached.conn;
  }

  if (!cached.promise) {
    cached.promise = MongoClient.connect(import.meta.env.MONGODB_URI).then(
      (client) => {
        return {
          client,
          db: client.db(import.meta.env.MONGODB_DB),
        };
      }
    );
  }
  cached.conn = await cached.promise;
  return cached.conn;
}
