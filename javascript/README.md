# nodejs Yoti Sign Example

- Clone this project
- Run `npm install`
- Register for Yoti Sign at https://www.yotisign.com/app/contact-us/
- Optain an API key once you have registered for Yoti Sign
- Rename `.env.example` to `.env`
- Fill in the missing environment variables using your API key and the API base URL
- Run the project with `npm start`
- Add a PDF to /documents and specify the name of the PDF to be example.pdf
- Go to http://localhost:3000 - this will create an envelope and send a Yoti Sign email to your recipients.
- To see notifications you need to have your destination publicly accessible
