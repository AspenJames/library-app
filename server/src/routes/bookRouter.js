import { Router } from 'express';

import Book from '../db/models/Book.js';

const router = Router();

// Get books
router.get('/', async (req, res) => {
  const books = await Book.query();
  res.json({ data: books });
});

// Create book
router.post('/', async (req, res, next) => {
  const book = await Book.create(req.body).catch(next);
  res.json({ data: { book } });
});

// Get book by id
router.get('/:id', async (req, res, next) => {
  const { id } = req.params;
  const book = await Book.query().findById(id).catch(next);
  if (book) {
    res.json({ data: book });
  } else {
    next(Book.createNotFoundError(null, {
      message: `Book with id ${id} not found`,
      modelClass: Book
    }));
  }
});

// Edit book by id
router.post('/:id', async (req, res, next) => {
  const { id } = req.params;
  const book = await Book.query().findById(id).catch(next);
  if (book) {
    const updated = await book.update(req.body).catch(next);
    res.json({ data: updated });
  } else {
    next(Book.createNotFoundError(null, {
      message: `Book with id ${id} not found`,
      modelClass: Book
    }));
  }
});

// Delete book by id
router.delete('/:id', async (req, res, next) => {
  const { id } = req.params;
  await Book.query().deleteById(id).catch(next);
  res.sendStatus(200);
});

export default router;
