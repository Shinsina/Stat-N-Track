
import type { APIRoute } from "astro";
import { connectToDatabase } from '$lib/mongodb';

export const get: APIRoute = async () => {
  try {
    const dbConnection = await connectToDatabase();
    const db = dbConnection.db;
    const collection = db.collection('subsessions');
    const data = await collection.find();
    new Response(JSON.stringify(data), {
      status: 200,
      headers: {
        "Content-Type": "application/json"
      }
    });
  } catch {
    return new Response(JSON.stringify({}), {
      status: 404,
      statusText: 'Not found'
    });
  }
  return new Response(JSON.stringify({}), {
    status: 500,
    statusText: 'Request Unable to be Processed',
  });
}