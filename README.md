# Test app for Mattermost

## Install

1. Running `make` will build the executable and start the server.
  - A base URL can be added so links are sent based on that url (e.g. `make BASE=http://myurl.com`). Defaults to `http://localhost:3000`.
  - An address can be added for the "ListenAndServe" function (e.g. `make ADDR=:3000`). Defaults to `:3000`.
3. Run the following command in Mattermost `/apps install --url http://localhost:3000/manifest`.
  - If a base URL has been set on step 2, run the install command with that URL. (e.g. `/app install --url http://myurl.com/manifest`)
4. As secret key, use `1234`.

## Provision

To provision this PR to AWS run `make dist` to generate the App bundle and then follow the steps [here](https://github.com/mattermost/mattermost-plugin-apps#provisioning).
