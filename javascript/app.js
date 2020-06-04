const createError = require("http-errors");
const express = require("express");
const path = require("path");
const cookieParser = require("cookie-parser");

//routes
const createEnvelopeRouter = require("./routes/createEnvelope");
const getEnvelopeRouter = require("./routes/getEnvelope");
const archiveEnvelopeRouter = require("./routes/archiveEnvelope");
const getDocumentsRouter = require("./routes/getDocuments");

const app = express();

// view engine setup
app.set("views", path.join(__dirname, "views"));
app.set("view engine", "hbs");

app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, "public")));

app.use("/", createEnvelopeRouter);
app.use("/getEnvelope", getEnvelopeRouter);
app.use("/archiveEnvelope", archiveEnvelopeRouter);
app.use("/getDocuments", getDocumentsRouter);

// catch 404 and forward to error handler
app.use(function (req, res, next) {
  next(createError(404));
});

// error handler
app.use(function (err, req, res, next) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get("env") === "development" ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.render("error");
});

module.exports = app;
