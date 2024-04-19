import { MongoClient } from "mongodb";

if (!process.env.MONGODB_URI && !import.meta.env.MONGODB_URI) {
  throw new Error(
    "Please define the MONGODB_URI environment variable inside a root .env file"
  );
}

if (!process.env.MONGODB_DB && !import.meta.env.MONGODB_DB) {
  throw new Error(
    "Please define the MONGODB_DB environment variable inside a root .env file"
  );
}

const mongodbURI = process.env.MONGODB_URI || import.meta.env.MONGODB_URI;
const mongodbDb = process.env.MONGODB_DB || import.meta.env.MONGODB_DB

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
    cached.promise = MongoClient.connect(mongodbURI).then(
      (client) => {
        return {
          client,
          db: client.db(mongodbDb),
        };
      }
    );
  }
  cached.conn = await cached.promise;
  return cached.conn;
}
