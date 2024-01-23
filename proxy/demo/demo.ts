import dotenv from "dotenv";

dotenv.config({
  path: "../../.env"
});

const { FLY_API_TOKEN, REGION } = process.env;

const authString = "Bearer " + FLY_API_TOKEN;
const baseName = "my-little-proxy";
const machinesApiUrl = "https://api.machines.dev";

const main = async () => {
  const appName = baseName + "-" + Math.random().toString(36).substring(2, 15);
  const region = REGION ?? "sea";

  console.log("Creating app...", appName);

  // Create the app.
  await post(machinesApiUrl + "/v1/apps", {
    app_name: appName,
    org_slug: "personal"
  })
}

const post = async (url: string, body: any) => {
  const response = await fetch(url, {
    method: "POST",
    body: JSON.stringify(body),
    headers: {
      "Content-Type": "application/json",
      "Authorization": authString
    }
  });

  if (!response.ok) {
    throw new Error("Request failed: " + response.statusText);
  }

  return await response.json();
}

main();