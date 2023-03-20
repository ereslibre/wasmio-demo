/**
 * Builds a reply to the given request
 */
const reply = (request) => {
  if (request.method != "GET") {
    return new Response("Method not allowed", {
      status: 405
    });
  }

  // Body response
  const body = JSON.stringify({ reason: "Definitely a teapot" });

  // Build a new response
  return new Response(body, { status: 418, headers: { "x-generated-by": "wasm-workers-server" } });
}

// Subscribe to the Fetch event
addEventListener("fetch", event => {
  return event.respondWith(reply(event.request));
});
