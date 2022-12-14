// npm run knex migrate:make create-book-authors
export const up = async (knex) =>
  knex.schema.createTable("book_authors", (table) => {
    table.timestamps({
      useTimestamps: true,
      defaultToNow: true,
      useCamelCase: false,
    });
    table.uuid("book_id").references('books.id')
    table.uuid("author_id").references('authors.id')
    // Enforce unique constraint
    table.unique(["book_id", "author_id"], { indexName: 'book_author_unique' })
  });

export const down = async (knex) => knex.schema.dropTableIfExists("book_authors");
