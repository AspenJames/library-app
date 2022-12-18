import objection from 'objection';
const {
  ValidationError,
  NotFoundError,
  DBError,
  UniqueViolationError,
  NotNullViolationError,
  ForeignKeyViolationError,
  CheckViolationError,
  DataError
} = objection;

// errResp sets default vaules and returns formatted JSON
const errResp = (errJson = {
  msg: '',
  type: 'UnknownError',
  data: {},
}) => {
  const respJson = {
    msg: errJson.msg || '',
    type: errJson.type || 'UnknownError',
    data: errJson.data || {},
    ...errJson
  };
  return { error: respJson };
};

// TODO: need much nicer logging
export const errorLogger = (err, req, res, next) => {
  console.error({
    error: err,
    url: req.originalUrl,
    route: req.route,
  });
  next(err);
};


// Adapted from https://vincit.github.io/objection.js/recipes/error-handling.html
export const errorHandler = (err, req, res, next) => {
  if (res.headersSent) return next(err);

  if (err instanceof ValidationError) {
    res.status(400).send(errResp({
      message: err.message,
      type: err.type,
      data: err.data,
    }));
  } else if (err instanceof NotFoundError) {
    res.status(404).send(errResp({
      message: err.message,
      type: 'NotFound',
    }));
  } else if (err instanceof UniqueViolationError) {
    res.status(409).send(errResp({
      message: `unique constraint for columns '${err.columns}' on table '${err.table}' violated`,
      type: 'UniqueViolation',
      data: {
        columns: err.columns,
        table: err.table,
      },
    }));
  } else if (err instanceof NotNullViolationError) {
    res.status(400).send(errResp({
      message: `'${err.table}'.'${err.column}' must not be null`,
      type: 'NotNullViolation',
      data: {
        columns: err.columns,
        table: err.table,
      },
    }));
  } else if (err instanceof ForeignKeyViolationError) {
    res.status(409).send(errResp({
      message: `foreign key constraint violated for ${err.table}`,
      type: 'ForeignKeyViolation',
      data: {
        table: err.table,
      },
    }));
  } else if (err instanceof CheckViolationError) {
    res.status(400).send(errResp({
      message: `check violation for ${err.table}`,
      type: 'CheckViolation',
      data: {
        table: err.table,
      },
    }));
  } else if (err instanceof DataError) {
    res.status(400).send(errResp({
      message: 'invalid data received',
      type: 'InvalidData',
    }));
  } else if (err instanceof DBError) {
    res.status(500).send(errResp({
      message: 'unknown database error occurred',
      type: 'UnknownDatabaseError',
    }));
  } else {
    res.status(500).send(errResp({
      message: 'internal server error',
      type: 'UnknownError',
    }));
  }
}

