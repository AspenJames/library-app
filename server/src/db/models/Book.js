import Base from './Base.js';

import { boolType, stringType } from './_schemaTypes.js';
import Author from './Author.js';

class Book extends Base {
  static get tableName() {
    return 'books'; // 20221212010204_create-books.mjs
  }

  static get idColumn() {
    return 'id';
  }

  static get jsonSchema() {
    return {
      type: 'object',
      required: ['title'],
      properties: {
        isbn: stringType(),
        title: stringType({ minLength: 1 }),
        edition: stringType(),
        readStatus: boolType()
      }
    }
  }

  static get relationMappings() {
    return {
      authors: {
        relation: Base.ManyToManyRelation,
        modelClass: Author,
        join: {
          from: 'books.id',
          through: {
            from: 'book_authors.book_id',
            to: 'book_authors.author_id',
          },
          to: 'authors.id'
        }
      }
    }
  }
}

export default Book;
