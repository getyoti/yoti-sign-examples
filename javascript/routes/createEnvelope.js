require("dotenv").config();
const express = require("express");
const router = express.Router();
const rp = require("request-promise");
const fs = require("fs");

const createEnvelope = () => {
  const options = {
    name: "envelope name",
    emails: {
      invitation: {
        body: {
          message: "Please sign this document",
        },
      },
    },
    recipients: [
      {
        name: "User 1",
        email: "example@email.com",
        role: "Signee",
        auth_type: "no-auth",
        sign_group: 1,
        tags: [
          {
            page_number: 1,
            x: 0.1,
            y: 0.1,
            type: "signature",
            optional: false,
            file_name: "example.pdf",
          },
        ],
      },
    ],
    notifications: {
      destination: "https://mysite.com/events",
      subscriptions: ["envelope_completion", "upload_errors"],
    },
  };

  const envelope = {
    method: "POST",
    uri: `${process.env.BASE_URL}/v2/envelopes`,
    formData: {
      file: fs.createReadStream("documents/example.pdf"), // file system location of PDF
      options: JSON.stringify(options),
    },
    headers: {
      authorization: `Bearer ${process.env.API_KEY}`, // API key
    },
  };
  return rp(envelope)
    .then((body) => body)
    .catch((err) => {
      err;
      console.log(err.message);
    });
};

router.get("/", async function (req, res, next) {
  let result = await createEnvelope();
  console.log(result);
  let envelopeid = JSON.parse(result).envelope_id;
  res.render("index", { envelopeid });
});

module.exports = router;
