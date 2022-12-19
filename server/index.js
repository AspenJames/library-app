import app from "./src/app.js";
import setupDb from "./src/db/db.js";
import logger from "./src/logger.js";


setupDb();

const port = process.env.PORT || 3000;

app.listen(port, () => {
  logger.log('info', `Example app listening on port ${port}`);
});
