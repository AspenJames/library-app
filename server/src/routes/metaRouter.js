import { Router } from 'express';

import { description, license, gitUrl, version } from '../constants.js';

const router = Router();

router.get('/healthcheck', (_, res) => {
  res.sendStatus(200);
});

router.get('/info', (_, res) => {
  res.json({
    'apiVersion': version,
    description,
    gitUrl,
    license,
  });
});


export default router;