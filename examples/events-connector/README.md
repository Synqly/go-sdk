# Events Connector Example

This tutorial demonstrates how to use the Synqly Events Connector to send events to a SIEM in a multi-tenant application.

It simulates a multi-tenant application that runs a background job for each tenant. The background job sends events to Synqly, which are then forwarded to a SIEM.

## Prerequisites

- A [Synqly](https://synqly.com) Organization
- Your Synqly Organization ID and API Token
- Go 1.20 or later
- A Splunk account -- [sign up for a free trial](https://www.splunk.com/en_us/download.html)
- A Splunk HTTP Event Collector (HEC) endpoint and API token -- [create a new HEC token](https://docs.splunk.com/Documentation/Splunk/8.1.3/Data/UsetheHTTPEventCollector#Create_an_Event_Collector_token)

## Setup and run the example

1. Clone this repository and navigate to the `examples/events-connector` directory.
2. The example uses environment variables to read your Synqly Organization ID and API Token. Set these variables in your terminal:

    ```bash
    export SYNQLY_ORG_ID=your-org-id
    export SYNQLY_API_TOKEN=your-api-token
    ```

3. Splunk is also configured using environment variables. Set these variables in your terminal:

    ```bash
    export SPLUNK_URL=https://my-org.splunkcloud.com:8088/services/collector/event
    export SPLUNK_TOKEN=my-splunk-token
    ```

4. Run the example:

    ```bash
    go run main.go
    ```

Exit the example by pressing `Ctrl+C`.

5. View the events in Splunk. In the Splunk UI, navigate to **Search** > **Search & Reporting** and run a search for recent events.
