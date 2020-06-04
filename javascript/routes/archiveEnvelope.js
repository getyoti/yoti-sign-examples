require("dotenv").config();
const express = require("express");
const router = express.Router();
const rp = require("request-promise");

const archiveEnvelope = (envelopeid) => {
  const envelope = {
    method: "PATCH",
    uri: `${process.env.BASE_URL}/v2/envelopes/${envelopeid}`,
    headers: {
      authorization: `Bearer ${process.env.API_KEY}`, // API key
    },
  };
  return rp(envelope);
};

router.get("/", async function (req, res, next) {
  const { envelopeid } = req.query;
  try {
    let result = await archiveEnvelope(envelopeid);
    console.log("Successful", result);
    res.render("archiveEnvelope", { envelopeid });
  } catch (error) {
    console.log(error.message);
    res.render("error", error);
  }
});

module.exports = router;
