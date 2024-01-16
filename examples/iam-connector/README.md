# IAM Connector Example

This example application shows how to use the IAM Connector to perform a series of actions on a user, and then query the identity service's audit log to see the logs of those actions.

This example uses Okta; however, any supported identity provider can be used. The configuration steps will need to be adjusted, but the main application logic will remain the same.

The application writes a file `tenant_store.yaml` that contains the data you need in persistent storage to use an integration across multiple runs. If not present, this file is created and populated with the data from the first run. If present, the file is read and the data is used to connect to Synqly.

## Setup

1. You need the URL, Client ID, and API token for your Okta tenant.
2. You need an API token for your Synqly Organization.
3. You need a user to perform the actions against. This should not be a real user as the actions will force a password reset and disable the user.

All of this information should be placed in a `config.yaml` file in this directory.

To use environment variables instead of a config file, set the following environment variables:

* `SYNQLY_ORG_TOKEN`
* `SYNQLY_OKTA_URL`
* `SYNQLY_OKTA_TOKEN`
* `SYNQLY_OKTA_CLIENTID`

## Running the Example

To run the example, run the following command:

```bash
go run main.go
```

## Cleanup

To clean up the example, delete the `tenant_store.yaml` file. You will also need to manually delete the `Secure Identities Co` account created in your Synqly Organization.
