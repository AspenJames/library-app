import Base from './Base.js';

import { stringType } from './_schemaTypes.js';
import Book from './Book.js';

class Author extends Base {
  static get tableName() {
    return 'authors'; // 20221214044901_create-authors.mjs
  }

  static get idColumn() {
    return 'id';
  }

  static get jsonSchema() {
    return {
      type: 'object',
      required: ['name'],
      properties: {
        name: stringType({ minLength: 1 }),
      }
    }
  }

  static get relationMappings() {
    return {
      books: {
        relation: Base.ManyToManyRelation,
        modelClass: Book,
        join: {
          from: 'authors.id',
          through: {
            from: 'book_authors.author_id',
            to: 'book_authors.book_id',
          },
          to: 'books.id'
        }
      }
    }
  }
}

export default Author;
