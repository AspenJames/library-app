import path from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

export default {
  client: "postgresql",
  connection: process.env.DB_URL,
  migrations: {
    directory: path.join(__dirname, "./migrations"),
    extension: "mjs",
    loadExtensions: [".mjs"],
  },
};
