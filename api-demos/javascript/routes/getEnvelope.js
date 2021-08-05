require("dotenv").config();
const express = require("express");
const router = express.Router();
const rp = require("request-promise");

const getEnvelope = (envelopeid) => {
  const envelope = {
    method: "GET",
    uri: `${process.env.BASE_URL}/v2/envelopes/${envelopeid}`,
    headers: {
      authorization: `Bearer ${process.env.API_KEY}`, // API key
    },
  };
  return rp(envelope);
};

router.get("/", async function (req, res, next) {
  try {
    let result = await getEnvelope(envelopeid);
    console.log(result);
    res.render("getEnvelope", { result, envelopeid });
  } catch (error) {
    console.log(error.message);
    res.render("error", error);
  }
});

module.exports = router;
