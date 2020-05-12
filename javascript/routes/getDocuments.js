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
  return rp(documents)
    .then((body) => {
      fs.writeFileSync("example.zip", body);
      console.log("Documents Saved");
      return "Successful";
    })
    .catch((err) => {
      console.log("Documents not Saved");
      console.log(err.error);
      return "Unsuccessful";
    });
};

router.get("/", async function (req, res, next) {
  const { envelopeid } = req.query;
  try {
    let results = await getDocuments(envelopeid);
    console.log(results);
    res.render("getDocuments", { results, envelopeid });
  } catch (error) {
    console.log(error);
  }
});

module.exports = router;
