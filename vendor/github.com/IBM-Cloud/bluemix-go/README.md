# Bluemix SDK for Go

[![Build Status](https://travis-ci.org/IBM-Cloud/bluemix-go.svg?branch=master)](https://travis-ci.org/IBM-Cloud/bluemix-go) [![GoDoc](https://godoc.org/github.com/IBM-Cloud/bluemix-go?status.svg)](https://godoc.org/github.com/IBM-Cloud/bluemix-go)

bluemix-go provides the Go implementation for operating the IBM Bluemix platform, which is based on the [Cloud Foundry API][cloudfoundry_api].

## Installing

1. Install the SDK using the following command

```bash
go get github.com/IBM-Cloud/bluemix-go
```

2. Update the SDK to the latest version using the following command

```bash
go get -u github.com/IBM-Cloud/bluemix-go
```


## Using the SDK

You must have a working Bluemix account to use the APIs. [Sign up][bluemix_signup] if you don't have one.

The SDK has ```examples``` folder which cites few examples on how to use the SDK.
First you need to create a session.

```go
import "github.com/IBM-Cloud/bluemix-go/session"

func main(){

    s := session.New()
    .....
}
```

Creating session in this way creates a default configuration which reads the value from the environment variables.
You must export the following environment variables.
* IBMID - This is the IBM ID
* IBMID_PASSWORD - This is the password for the above ID

OR

* BM_API_KEY/BLUEMIX_API_KEY - This is the Bluemix API Key. Login to [Bluemix][bluemix_login] to create one if you don't already have one. Follow Manage -> Account -> Users. Click on _Bluemix API Keys_

The default region is _us_south_. You can override it in the [Config struct][bluemix_go_config]. You can also provide the value via environment variables; either via _BM_REGION_ or _BLUEMIX_REGION_. Valid regions are -
* us-south
* eu-gb
* eu-de
* au-syd

[bluemix_signup]: https://console.ng.bluemix.net/registration/?target=%2Fdashboard%2Fapps
[bluemix_login]: https://console.ng.bluemix.net
[bluemix_go_config]: https://godoc.org/github.com/IBM-Cloud/bluemix-go#Config
[cloudfoundry_api]: https://apidocs.cloudfoundry.org/264/
