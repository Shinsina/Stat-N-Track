import type { APIRoute } from "astro";
import { connectToDatabase } from '$lib/mongodb';

export const get: APIRoute = async ({ params }) => {
  const id = Number(params.id);
  const dbConnection = await connectToDatabase();
  const db = dbConnection.db;
  const collection = db.collection('subsessions');
  const data = await collection.findOne({ _id: id });
  if (data) {
    return new Response(JSON.stringify(data), {
      status: 200,
      headers: {
        "Content-Type": "application/json"
      }
    });
  }
  return new Response(JSON.stringify({}), {
    status: 404,
    statusText: 'Not found'
  });
}
