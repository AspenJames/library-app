import winston from 'winston';

const level = process.env.LOG_LEVEL || 'info';

const logger = winston.createLogger({
  transports: [
    new winston.transports.Console({ level }),
  ],
  format: winston.format.json(),
});

export default logger;