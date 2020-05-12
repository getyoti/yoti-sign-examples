# nodejs Yoti Sign Example

- Clone this project
- Run `npm install`
- Register for Yoti Sign at https://www.yotisign.com/app/contact-us/
- Optain an API key once you have registered for Yoti Sign
- Create a .env file using the example `.env.example` file
- Fill in the missing environment variables using your API key and the API base URL
- Run the project with `npm start` or `nodemon`
- Add a PDF to /documents and configure the options object in /routes/createEnvelope.js
- Go to http://localhost:3000 - this will create an envelope and send a Yoti Sign email to your recipients.
- To see notifications you need to have your destination publicly accessible
