import { Client, uploadPresignedFile } from "singlebase-js";

const client = new Client({
  apiKey: "my-api-key",
  endpointKey: "vector-db",
});

const payload = { op: "ping" };

client.dispatch(payload).then((res) => {
  if (res.ok) {
    console.log("✅ Success:", res.data);
  } else {
    console.error("❌ Error:", res.error);
  }
});
