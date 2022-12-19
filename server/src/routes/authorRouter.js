import { Router } from 'express';

import Author from '../db/models/Author.js';

const router = Router();

// Get authors
router.get('/', async (req, res) => {
  const authors = await Author.query();
  res.json({ data: authors });
});

// Create author
router.post('/', async (req, res, next) => {
  const author = await Author.create(req.body).catch(next);
  res.json({ data: { author } });
});

// Get author by id
router.get('/:id', async (req, res, next) => {
  const { id } = req.params;
  const author = await Author.query().findById(id).catch(next);
  if (author) {
    res.json({ data: author });
  } else {
    next(Author.createNotFoundError(null, {
      message: `Author with id ${id} not found`,
      modelClass: Author
    }));
  }
});

// Edit author by id
router.post('/:id', async (req, res, next) => {
  const { id } = req.params;
  const author = await Author.query().findById(id).catch(next);
  if (author) {
    const updated = await author.update(req.body).catch(next);
    res.json({ data: updated });
  } else {
    next(Author.createNotFoundError(null, {
      message: `Author with id ${id} not found`,
      modelClass: Author
    }));
  }
});

// Delete author by id
router.delete('/:id', async (req, res, next) => {
  const { id } = req.params;
  await Author.query().deleteById(id).catch(next);
  res.sendStatus(200);
});

export default router;
