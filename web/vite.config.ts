import { defineConfig } from "vite";
import { copyFile } from "fs/promises";
import { resolve } from "path";

const root = process.cwd();
const sourceCityJson = resolve(root, "..", "data", "city.json");
const publicDirectory = resolve(root, "public");
const publicCityJson = resolve(publicDirectory, "city.json");

async function copyCityJson() {
  await copyFile(sourceCityJson, publicCityJson);
}

function cityJsonCopyPlugin() {
  return {
    name: "city-json-copy",
    buildStart: copyCityJson,
    configureServer() {
      copyCityJson().catch(() => {});
    },
  };
}

export default defineConfig({
  server: {
    port: 5173,
    open: true,
  },
  plugins: [cityJsonCopyPlugin()],
});
