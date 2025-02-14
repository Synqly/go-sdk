## Asset Management Example

This example showcases how to use Synqly's Asset Management Connector to interact with a asset management system. We'll be using Service Now, Armis Centrix and Nozomi Vantage, which are currently supported by Synqly. Assets Connector Overview (https://docs.synqly.com/docs/assets-connector-overview)


## Setup

You need the following to run this example:

1. An API token for your Synqly Organization. [Auth Doc] (https://docs.synqly.com/reference/api-authentication) [API Keys] (https://app.synqly.com/prod/settings/secrets)
2. A ServiceNow account and API token. Follow [ServiceNow configuration guide](https://docs.synqly.com/docs/servicenow-webhook-setup).
3. A Armis Centrix account and API token. Follow [Armix Centrix configuration guide](https://docs.synqly.com/docs/armis-centrix-setup)
4. A Nozomi Vantage account and API token. Follow [Nozomi Vantage configuration guide](https://docs.synqly.com/docs/nozomi-vantage-setup)

All of this information should be placed in a `config.env` file in this directory or set through environment variables.

The example will only execute for providers with authentication details specified in the `config.env` file. If authentication details are missing, that provider will be skipped.

## Running the Example

To run the example:

```bash
go run main.go
```

The logs for each action performed will be displayed in the console.