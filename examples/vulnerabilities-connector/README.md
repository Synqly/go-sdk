# Vulnerabilities Connector Example

This example application demonstrates how to use Synqly vulnerabilities connectors to query assets, findings and scans, using different providers, this can be extended to make other supported calls.

## Setup
1. A [Synqly](https://synqly.com) `Organization` API token
2. Go 1.20 or later
3. An account and credentials for one of the supported vulnerabilities providers:
    - CroudStrike
        - Client ID
        - Client Secret
        - CroudStrike API URL
    - Nucleus
        - Project ID
        - A valid Token
        - Nucleus API URL
    - Qualys
        - Secret
        - Qualys URL
        - Valid user
    - Rapid7
        - A valid token
        - Rapid7 API URL
    - Tanium
        - A valid token
        - Tanium API URL
    - Tenable
        - Valid token
        - Tenable URL

All this information (Or the one required for the provider you want to test), should be place in `config.env` file that is on this directory.

The application writes a file `tenant_store_PROVIDER_NAME.yaml` that contains the data you need in persistent storage to use an integration across multiple runs. If not present, this file is created and populated with the data from the first run. If present, the file is read and the data is used to connect to Synqly.

## Running the Example

To run the example, you can use one of the following commands.

Run for all providers:
```bash
go run main.go
```

Run for a specific provider:
```bash
go run main.go --provider PROVIDER_NAME
```

The available providers are:
- CroudStrike
- Nucleus
- Qualys
- Rapid7
- Tanium
- Tenable

## Cleanup

To clean up the example, delete the `tenant_store_PROVIDER_NAME.yaml` file. You will also need to manually delete the `Zenith Systems - PROVIDER_NAME` account created in your Synqly Organization.
