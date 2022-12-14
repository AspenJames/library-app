// npm run knex migrate:make create-authors
export const up = async (knex) =>
  knex.schema.createTable("authors", (table) => {
    table.uuid("id", { primaryKey: true });
    table.timestamps({
      useTimestamps: true,
      defaultToNow: true,
      useCamelCase: false,
    });
    table.string("name", 255);
  });

export const down = async (knex) => knex.schema.dropTableIfExists("authors");
