# Sentinel Connector Example

This section acts as supporting documentation for the sample program in `examples/sentinel-connector`. The `sentinel-connector` example demonstrates how to use the Synqly's Sentinel connector to send events to a Sentinel Provider from within a multi-tenant application. 

The `sentinel-connector` example will:

- Define multiple tenants in a sample application
- Define an `Integration` for each tenant (using Azure as the target Sentinel Provider)
- Run a background job to simulate load for each tenant
- Send events from the background job to Synqly
- Demonstrate that events are sent in OCSF format and transformed by Synqly before
  being forwarded to the target SIEM Provider.

## Prerequisites

- A [Synqly](https://synqly.com) `Organization`
- Your Synqly Organization ID and API Token
- Go 1.20 or later
- An Azure account -- [sign up for a free trial](https://azure.microsoft.com/en-us/pricing/purchase-options/azure-account)
- Create data collection rules (DCRs) in Azure Monitor -- [create a new DCR](https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/data-collection-rule-create-edit?tabs=cli#create-or-edit-a-dcr-using-the-azure-portal)

## Setup and run the example

1. Clone this repository and navigate to the `examples/sentinel-example` directory.
    ```bash
    cd examples/sentinel-example
   
    ```
2. Initialize go modules (this will also verify access that `go` can pull from
   this repository). 
   ```bash
   go mod tidy
   ```
3. The example could be setup it with environment variables or completing the config.env file to connect to Synqly and Splunk. To allow the sample program to connect to Synqly, set the following environment variables to your Organization ID and API Token values.

    ```bash
    export SYNQLY_ORG_TOKEN=your-api-token
    ```

4. Access to Splunk is also configured via environment variable. Set the following variables in your terminal:

    ```bash
    export SPLUNK_URL=https://my-org.splunkcloud.com:8088/services/collector/event
    export SPLUNK_HEC_TOKEN=my-splunk-token
    ```

   Alternatively, you can create a `config.env` file in the `examples/siem-connector` directory with the following content:

    ```bash
   SYNQLY_ORG_TOKEN=<synqly_org_token>
   SPLUNK_HEC_URL=<splunk_hec_url>
   SPLUNK_HEC_TOKEN=<splunk_hec_token>
   SPLUNK_SEARCH_URL=<splunk_search_url>
   SPLUNK_SEARCH_TOKEN=<splunk_search_token>
   DURATION_SECONDS=600
   ```


5. Run the example:

    ```bash
    go run . 
    ```

Exit the example by pressing `Ctrl+C`.

6. View the events in Splunk. In the Splunk UI, navigate to **Search** > **Search & Reporting** and run a search for `sourcetype=httpevent`.