# Yoti Sign Embedded Signing Go Example

## Setup

1) Register for Yoti Sign at <https://www.yotisign.com/app/contact-us/>
1) Request an API key (Bearer Authentication Token)

## Running the example

1) Clone this repo
2) Rename `.env.example` to `.env`
3) Add your Bearer Authentication Token to the `YOTI_AUTHENTICATION_TOKEN` variable in the `.env` file
4) Update `options.json` to have your email address
5) Create the binary executable file: `go build`
6) Run the executable: `./embedded-signing` (UNIX) or `embedded-signing.exe` (Windows)
7) visit <https://localhost:8080>
