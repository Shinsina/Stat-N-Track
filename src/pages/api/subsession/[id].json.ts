import type { APIRoute } from "astro";
import example from "../../../data/example";

export const get: APIRoute = ({ params }) => {
  const id = params.id;

  if (example.subsession_id !== Number(id)) {
    return new Response(null, {
      status: 404,
      statusText: 'Not found'
    });
  }

  return new Response(JSON.stringify(example), {
    status: 200,
    headers: {
      "Content-Type": "application/json"
    }
  });
}
