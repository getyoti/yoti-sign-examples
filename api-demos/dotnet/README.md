# .NET Yoti Sign Example

## Setup

1) Register for Yoti Sign at <https://www.yotisign.com/app/contact-us/>
1) Request an API key (Bearer Authentication Token)

## Running the example

1) Clone this repo
1) Run `cd dotnet/YotiSignDemo`
1) Rename `.env.example` to `.env`
1) Add your Bearer Authentication Token to the `YOTI_AUTHENTICATION_TOKEN` variable in the `.env` file
1) Set the API URL to match the production or demo environment: `/Controllers/SignController.cs L85` ;. Demo keys will not work in production and vice versa
1) Customise the `options.json` file as you wish
1) Run `dotnet run`
1) visit <https://localhost:5001/sign>
