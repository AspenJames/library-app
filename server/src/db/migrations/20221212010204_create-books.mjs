// npm run knex migrate:make create-books
import { defaultCols } from "../_tools.js";

export const up = async (knex) =>
  knex.schema.createTable("books", (table) => {
    defaultCols(table, knex);
    table.string("isbn", 255);
    table.string("title", 255).notNullable();
    table.string("edition", 255);
    table.boolean("read_status").defaultTo(false);
  });

export const down = async (knex) => knex.schema.dropTableIfExists("books");
