// npm run knex migrate:make create-authors
import { defaultCols } from "../_tools.js";

export const up = async (knex) =>
  knex.schema.createTable("authors", (table) => {
    defaultCols(table, knex);
    table.string("name", 255).notNullable();
  });

export const down = async (knex) => knex.schema.dropTableIfExists("authors");
