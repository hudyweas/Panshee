# Panshee
The tool is based on microservice architecture. It lets check basic 
information about wallets and contained tokens from the BNB chain. 
The whole tool is segmented into 5 microservices: API, auth, wallet, 
email and front-end.

## Caution!
To work environmental variables are needed to set (.env in the main folder of every service). 
Env var files schemas are set in config.go files in every service (Config structure).

Api and Auth service need postgres databases (ORM will auto-create tables).

Because of using free API and it's requests limit, the tool can only handle two wallets on an account.
