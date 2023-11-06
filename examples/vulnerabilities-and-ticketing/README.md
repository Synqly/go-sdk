# Vulnerabilities and Ticketing Example

This example application shows how Synqly Connectors can be combined to build a workflow that automatically creates a ticket in a ticketing system when a vulnerability is detected. The application queries vulnerabilities and then creates a ticket for each vulnerability found. If a ticket has already been opened for a vulnerability, the application logs that a ticket already exists and does not create a new ticket.

This example uses Tenable and Jira as the vulnerability scanner and ticketing system, respectively. However, it could be adapted to use any supported vulnerability scanner and ticketing system.

The application writes a file `tenant_store.yaml` that contains the data you need in persistent storage to use an integration across multiple runs. If not present, this file is created and populated with the data from the first run. If present, the file is read and the data is used to connect to Synqly.

## Setup

You need the following to run this example:

1. An API token for your Synqly Organization.
2. A Jira account and nn API token. Follow [this guide to create a token](https://support.atlassian.com/atlassian-account/docs/manage-api-tokens-for-your-atlassian-account/).
3. A Tenable Cloud account and an API token. Follow [this guide to create a token](https://docs.tenable.com/vulnerability-management/Content/Settings/my-account/GenerateAPIKey.htm).

All of this information should be placed in a `config.yaml` file in this directory or set through environment variables.

For your Tenable Cloud token, you need to use the format `accessKey=<value>;secretKey=<value>`.

## Running the Example

To run the example, run the following command:

```bash
go run main.go
```

## Cleanup

To clean up the example, delete the `tenant_store.yaml` file. You will also need to manually delete the `Secure Identities Co` account created in your Synqly Organization.
