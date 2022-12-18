import { Router } from 'express';
import { NotFoundError } from 'objection';

import Book from '../db/models/Book.js';

const router = Router();

// Get books
router.get('/', async (req, res) => {
  const books = await Book.query();
  res.json({ data: books });
});

// Create book
router.post('/', async (req, res, next) => {
  const bookParams = req.body;
  const book = await Book.createBook(bookParams).catch(next);
  res.json({ data: { book } });
});

// Get book by id
router.get('/:id', async (req, res, next) => {
  const { id } = req.params;
  const book = await Book.query().findById(id).catch(next);
  if (book) {
    res.json({ data: book });
  } else {
    next(new NotFoundError({
      message: `Book with id ${id} not found`,
      modelClass: Book
    }));
  }
})

export default router;
