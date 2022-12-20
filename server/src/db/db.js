import Knex from "knex";
import { Model, knexSnakeCaseMappers } from "objection";

import config from "./knexfile.js";

// Setup Knex connection, run migrations,
// register knex with Objeciton models.
function setupDb() {
  const knex = Knex({
    ...config,
    ...knexSnakeCaseMappers(),
  });

  knex.migrate.latest();

  Model.knex(knex);
}

export default setupDb;
