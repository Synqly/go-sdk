# SIEM Connector Example

This section acts as supporting documentation for the sample program in `examples/siem-connector`. The `siem-connector` example demonstrates how to use the Synqly's SIEM Connector to send events to a SIEM Provider from within a multi-tenant application. 

The `siem-connector` example will:

- Define multiple tenants in a sample application
- Define an `Integration` for each tenant (using Splunk as the target SIEM Provider)
- Run a background job to simulate load for each tenant
- Send events from the background job to Synqly
- Demonstrate that events are sent in OCSF format and transformed by Synqly before
  being forwarded to the target SIEM Provider.

## Prerequisites

- A [Synqly](https://synqly.com) `Organization`
- Your Synqly Organization ID and API Token
- Go 1.20 or later
- A Splunk account -- [sign up for a free trial](https://www.splunk.com/en_us/download.html)
- A Splunk HTTP Event Collector (HEC) endpoint and API token -- [create a new HEC token](https://docs.splunk.com/Documentation/Splunk/8.1.3/Data/UsetheHTTPEventCollector#Create_an_Event_Collector_token)

## Setup and run the example

1. Clone this repository and navigate to the `examples/siem-connector` directory.
    ```bash
    cd examples/siem-connector
    ```
2. Initialize go modules (this will also verify access that `go` can pull from
   this repository). 
   ```bash
   go mod tidy
   ```
3. The example depends on environment variables to connect to Synqly Splunk and Sumo Logic. To allow the sample program to connect to Synqly, set the following environment variables to your Organization ID and API Token values.

    ```bash
    export SYNQLY_ORG_TOKEN=your-api-token
    ```

4. Access to Splunk is also configured via environment variable. Set the following variables in your terminal:

    ```bash
    export SPLUNK_URL=https://my-org.splunkcloud.com:8088/services/collector/event
    export SPLUNK_HEC_TOKEN=my-splunk-token
    ```

5. Run the example:

    ```bash
    go run main.go
    ```

Exit the example by pressing `Ctrl+C`.

6. View the events in Splunk. In the Splunk UI, navigate to **Search** > **Search & Reporting** and run a search for `sourcetype=httpevent`.