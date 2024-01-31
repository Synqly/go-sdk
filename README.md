\*\*\*\*\*\***CONFIDENTIAL**\*\*\*\*\*\*

# Synqly Go SDK

This repository contains Go packages for integrating with Synqly APIs.

# SDK Access

Access to this repository is currently restricted to authenticated users, necessitating a couple extra steps in order to use the packages locally. The steps in this section will cover how to configure access to `github.com/Synqly/go-sdk` for use in your application.

### Step 1: Set GOPRIVATE

The first step to using a private Go package is to add it to the `GOPRIVATE` environment variable. This will prevent `go mod` from checking for the package on  public proxies, causing it to instead download the package directly from Github.

```
export GOPRIVATE=github.com/synqly/go-sdk
```

### Step 2: Configure Github Authentication to Synqly/go-sdk

In order for `go mod` to successfully pull `github.com/Synqly/go-sdk`, your local machine must be able to authenticate with Github. The two most common ways to authenticate are through either a Github Personal Access Token or an SSH key linked to a Github account. The following sections describe the `go mod` configuration for both methods.

If your local machine is already configured to authenticate as the Github user which has been granted access to `github.com/Synqly/go-sdk`, you can skip to Step 3.


#### Github Personal Access Token (PAT)

This section describes how to use a Github Personal Access Token (PAT) to pull `github.com/Synqly/go-sdk`. The method below relies on Git's built-in [gitcredentials](https://git-scm.com/docs/gitcredentials) to cache a PAT for use with `go mod`.

In order to configure access, first navigate to a temporary directory. This directory can be anywhere on the local system.

```bash
mkdir ~/synqly-tmp
cd ~/synqly-tmp
```

Next, attempt to clone Synqly/go-sdk:

```bash
git clone https://github.com/Synqly/go-sdk.git
```

If the system is not already authenticated to pull from Github, it should prompt for a Username, then a Password. The Password field expects a Personal Access Token.

```bash
# git clone output
Cloning into 'go-sdk'...
Username for 'https://github.com': <MY_USER_NAME>
Password for 'https://<MY_USER_NAME>@github.com': <PAT_VALUE>
```

Enter a Username and Token value for the Github user that has been granted access to Synqly's go-sdk. This will cache the provided values using git's credential helper, making them available for `go mod` later on.

If the authentication values provided are valid, the repository should pull successfully.

```bash
# git clone output
remote: Enumerating objects: 465, done.
remote: Counting objects: 100% (465/465), done.
remote: Compressing objects: 100% (278/278), done.
remote: Total 465 (delta 180), reused 451 (delta 169), pack-reused 0
Receiving objects: 100% (465/465), 257.65 KiB | 3.53 MiB/s, done.
Resolving deltas: 100% (180/180), done.
```

#### SSH Key

This section describes how to set `go mod` to use an SSH key when pulling `github.com/Synqly/go-sdk`. This section assumes you have already followed Github's [Add a New SSH Key](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account) guide. 

Open your local `~/.gitconfig` file in your editor of choice. Add the following block if it does not already exist. This configuration will instruct tools that pull from github.com, such as `go mod`, to use SSH authentication. 

```
[url "ssh://git@github.com/"]
	insteadOf = https://github.com/
```

### Step 3: (Optional) Test Synqly/go-sdk Module Import

Once the local system is authenticated to clone `github.com/Synyqly/go-sdk`, `go mod` should be able to use the same credentials to import the SDK as a Go module. You can now add packages from `github.com/Synqly/go-sdk` to a Go program to test whether they are importable. 

The following example uses a minimal sample application to test package importability.

First, create a sample application directory:

```bash
mkdir ~/sample-synqly-app
cd ~/sample-synqly-app
```

Next, create a new go program using `go mod`:

```bash
go mod init synqly-go
```

Create a `main.go` file to print one of Synqly's OCSF Event Enum values.

```go
cat << EOF > main.go
    package main

    import (
        "fmt"

        engine "github.com/synqly/go-sdk/client/engine"
    )

    func main() {
        exampleStatus := engine.NotificationStatusOpen
        fmt.Printf("Example Synqly OCSF Notification status: %v\n", exampleStatus)
    }
EOF
```

Test that `go mod` can import `github.com/synqly/go-sdk`. If this step succeeds without error, then the import was successful.

```bash
go mod tidy
```

As a final check, run the `main.go` program to test the module is behaving as expected.

```bash
go run main.go
# Go run output
Example Synqly OCSF Notification status: OPEN
```

# SIEM Connector Example

The `examples` directory of this repository contains example implementations that demonstrate how to incorporate Synqly SDKs into a Go application.

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
3. The example depends on environment variables to connect to Synqly and Splunk. To allow the sample program to connect to Synqly, set the following environment variables to your Organization ID and API Token values.

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
