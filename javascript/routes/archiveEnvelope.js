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
  return rp(envelope)
    .then((body) => {
      body;
      return "Successful";
    })
    .catch((err) => {
      err;
      return `Unsuccessful: ${err.message}`;
    });
};

router.get("/", async function (req, res, next) {
  const { envelopeid } = req.query;
  let result = await archiveEnvelope(envelopeid);
  console.log(result);
  res.render("archiveEnvelope", { result, envelopeid });
});

module.exports = router;
