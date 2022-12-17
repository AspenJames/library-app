// npm run knex migrate:make create-books
export const up = async (knex) =>
  knex.schema.createTable("books", (table) => {
    table.uuid("id", { primaryKey: true })
      .defaultTo(knex.raw('gen_random_uuid()'));
    table.timestamps({
      useTimestamps: true,
      defaultToNow: true,
      useCamelCase: false,
    });
    table.string("isbn", 255);
    table.string("title", 255).notNullable();
    table.string("edition", 255);
    table.boolean("read_status").defaultTo(false);
  });

export const down = async (knex) => knex.schema.dropTableIfExists("books");
