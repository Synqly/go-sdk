# Vulnerabilities Connector Example

This example application demonstrates how to use the Synqly Vulnerability Connectors to query assets, findings and scans from vulnerability management systems. It demonstrates how to use and configure many different providers and can be extended to support other operations, such as creating assets or findings.

## Setup
1. A [Synqly](https://synqly.com) `Organization` API token
2. Go 1.20 or later
3. An account and credentials for one of the supported vulnerabilities providers:
    - [CroudStrike](https://docs.synqly.com/docs/crowdstrike-vulns-setup)
        - Client ID
        - Client Secret
        - CroudStrike API URL
    - [Nucleus](https://docs.synqly.com/docs/nucleus-vulns-setup)
        - Project ID
        - A valid Token
        - Nucleus API URL
    - [Qualys](https://docs.synqly.com/docs/qualys-vulns-setup)
        - Secret
        - Qualys URL
        - Valid user
    - [Rapid7](https://docs.synqly.com/docs/rapid7-vulns-setup)
        - A valid token
        - Rapid7 API URL
    - [Tanium](https://docs.synqly.com/docs/tanium-setup)
        - A valid token
        - Tanium API URL
    - [Tenable](https://docs.synqly.com/docs/tenable-vulns-setup)
        - Valid token
        - Tenable URL

For any provider you wish to test, its configuration details should be placed in the `config.env` file. See the included config.env file for the precise details.

The application writes a file `tenant_store_PROVIDER_NAME.yaml` that contains the data you need in persistent storage to use an integration across multiple runs. If not present, this file is created and populated with the data from the first run. If present, the file is read and the data is used to connect to Synqly.

## Running the Example
```bash
go run main.go
```

## Cleanup
The program will automatically cleanup your Synqly organization and configuration files, if you want to prevent this please use:
```bash
go run main.go --cleanup=false
```

Note:
- If you manually delete one of the test accounts or the configuration file for the account, please delete the counter part:
    - If you delete the test account, delete also the file.
    - If you delete the file, delete also the account on your Synqly organization.