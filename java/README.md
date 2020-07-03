# Java Yoti Sign Example

## Setup

1) Register for Yoti Sign at <https://www.yotisign.com/app/contact-us/>
1) Request an API key (Bearer Authentication Token)

## Running the example

1) Clone this repo
1) Run `cd java`
1) Rename `application.yml.example` to `application.yml`
1) Add the Yoti Sign Base URL and your Bearer Authentication Token to the `baseUrl` and `authenticationToken` variables in the `application.yml` file`
1) Customise the `options.json` file as you wish
1) Run `mvn clean package`
1) Run `java -jar target/sign-demo-1.0.0.jar'
1) visit <https://localhost:8443>
