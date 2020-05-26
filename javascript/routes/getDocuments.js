require("dotenv").config();
const express = require("express");
const router = express.Router();
const rp = require("request-promise");
const fs = require("fs");

const getDocuments = (envelopeid) => {
  const documents = {
    method: "GET",
    uri: `${process.env.BASE_URL}/v2/envelopes/${envelopeid}/completed-documents`,
    headers: {
      authorization: `Bearer ${process.env.API_KEY}`, // API key
    },
    encoding: null,
  };
  return rp(documents);
};

router.get("/", async function (req, res, next) {
  const { envelopeid } = req.query;
  try {
    let body = await getDocuments(envelopeid);
    fs.writeFileSync("example.zip", body);
    console.log("Documents Saved");
    let results = "Successful";
    res.render("getDocuments", { results, envelopeid });
  } catch (error) {
    console.log(error.message);
    res.render("error", error);
  }
});

module.exports = router;
