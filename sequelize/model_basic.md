```js
const { Sequelize, Op, Model, DataTypes } = require("sequelize");

const sequelize = new Sequelize("<DatabaseName>", "<UserName>", "<Password>", {
  host: "localhost",
  dialect: "postgres",
  logging: (...msg) => console.log(msg),
});

try {
  sequelize.authenticate().then((res) => {
    console.log("Connection has been established successfully");
    console.log(res);
  });
} catch (error) {
  console.error("Unable to connect to database: ", error);
}

module.exports = sequelize;
```
