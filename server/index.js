import app from "./src/app.js";
import setupDb from "./src/db/db.js";

setupDb();

const port = process.env.PORT || 3000;

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
