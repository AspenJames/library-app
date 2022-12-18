import express from 'express';

import { version } from './constants.js';
import { errorHandler, errorLogger } from './error.js';

import bookRouter from './routes/bookRouter.js';
import metaRouter from './routes/metaRouter.js';

const app = express();

// Mount middlewares
app.use(express.json());
app.disable('x-powered-by');
app.use((_, res, next) => {
  res.set('library-app-api-version', version);
  next();
})

// Mount routers
app.use('/', metaRouter);
app.use('/books', bookRouter);

app.use(errorLogger);
app.use(errorHandler);

export default app;
