# IBM Cloud DNS Services Go SDK
[![Build Status](https://travis.ibm.com/ibmcloud/dns-svcs-go-sdk.svg?token=5N69GxPxKisDHsait9qz&branch=master)](https://travis.ibm.com/ibmcloud/dns-svcs-go-sdk)

Go client library to use the [IBM Cloud DNS Services API](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-dns-svcs-api).

<details>
<summary>Table of Contents</summary>

- [IBM Cloud DNS Services Go SDK](#ibm-cloud-dns-services-go-sdk)
  - [Overview](#overview)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
    - [1. `go get` command](#1-go-get-command)
    - [2. Go modules](#2-go-modules)
    - [3. `dep` dependency manager](#3-dep-dependency-manager)
  - [Using the SDK](#using-the-sdk)
    - [Constructing service clients](#constructing-service-clients)
    - [Authentication](#authentication)
      - [Examples:](#examples)
    - [Passing operation parameters via an options struct](#passing-operation-parameters-via-an-options-struct)
    - [Receiving operation responses](#receiving-operation-responses)
      - [Example:](#example)
    - [Error Handling](#error-handling)
      - [Example:](#example-1)
    - [Default headers](#default-headers)
      - [Example:](#example-2)
    - [Sending request headers](#sending-request-headers)
      - [Example:](#example-3)
    - [Transaction IDs](#transaction-ids)
  - [License](#license)

</details>

## Overview

The `dns-svcs-go-sdk` allows developers to programmatically interact with the
IBM Cloud DNS Services API.

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration?target=%2Fdeveloper%2Fwatson&

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* An installation of Go (version 1.12 or above) on your local machine.

## Installation
There are a few different ways to download and install the `dns-svcs-go-sdk` project for use by your
Go application:
### 1. `go get` command
Use this command to download and install the `dns-svcs-go-sdk` project to allow your Go application to
use it:
```
go get -u github.com/IBM/dns-svcs-go-sdk
```
### 2. Go modules
If your application is using Go modules, you can add a suitable import to your
Go application, like this:
```go
import (
    "github.com/IBM/dns-svcs-go-sdk/dnssvcsv1"
)
```
then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.
### 3. `dep` dependency manager
If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file.  Here is an example:
```
[[constraint]]
  name = "github.com/IBM/dns-svcs-go/dnssvcsv1"
  version = "0.0.1"

```
then run `dep ensure`.

## Using the SDK
This section provides general information on how to use the services contained in this SDK.

### Constructing service clients
Each service is implemented in its own package (e.g. `dnssvcsv1`).
The package will contain a "service client"
struct (a client-side representation of the service), as well as an "options" struct that is used to
construct instances of the service client.  
Here's an example of how to construct an instance of "My Service":
```go
import (
    "github.com/IBM/go-sdk-core/core"
    "github.com/IBM/dns-svcs-go-sdk/dnssvcsv1"
)

// Create an authenticator.
authenticator := /* create an authenticator - see examples below */

// Create an instance of the "MyServiceV1Options"  struct.
myserviceURL := "https://api.dns-svcs.cloud.ibm.com/v1"
options := &dnssvcsv1.MyServiceV1Options{
    Authenticator: authenticator,
    URL: myserviceURL,
}

// Create an instance of the "MyServiceV1" service client.
service, err := NewMyServiceV1(options)
if err != nil {
    // handle error
}

// Service operations can now be called using the "service" variable.

```

### Authentication
DNS Services use token-based Identity and Access Management (IAM) authentication.

IAM authentication uses an API key to obtain an access token, which is then used to authenticate
each API request.  Access tokens are valid for a limited amount of time and must be regenerated.

To provide credentials to the SDK, you supply either an IAM service **API key** or an **access token**:

- Specify the IAM API key to have the SDK manage the lifecycle of the access token.
The SDK requests an access token, ensures that the access token is valid, and refreshes it when
necessary.
- Specify the access token if you want to manage the lifecycle yourself.
For details, see [Managing IAM and IBM Cloud DNS Services](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-iam).

#### Examples:
* Supplying the IAM API key and letting the SDK manage the access token for you:

```go
// letting the SDK manage the IAM access token
import (
    "github.com/IBM/go-sdk-core/core"
    "github.com/IBM/dns-svcs-go-sdk/dnssvcsv1"
)
...
// Create the IAM authenticator.
authenticator := &core.IamAuthenticator{
    ApiKey: "myapikey",
}

// Create the service options struct.
options := &dnssvcsv1.MyServiceV1Options{
    Authenticator: authenticator,
}

// Construct the service instance.
service, err := dnssvcsv1.NewMyServiceV1(options)

```

* Supplying the access token (a bearer token) and managing it yourself:

```go
import {
    "github.com/IBM/go-sdk-core/core"
    "github.com/IBM/dns-svcs-go-sdk/dnssvcsv1"
}
...
// Create the BearerToken authenticator.
authenticator := &core.BearerTokenAuthenticator{
    BearerToken: "my IAM access token",
}

// Create the service options struct.
options := &dnssvcsv1.MyServiceV1Options{
    Authenticator: authenticator,
}

// Construct the service instance.
service, err := dnssvcsv1.NewMyServiceV1(options)

...
// Later when the access token expires, the application must refresh the access token,
// then set the new access token on the authenticator.
// Subsequent request invocations will include the new access token.
authenticator.BearerToken = /* new access token */
```
For more information on authentication, including the full set of authentication schemes supported by
the underlying Go Core library, see
[this page](https://github.com/IBM/go-sdk-core/blob/master/Authentication.md)

### Passing operation parameters via an options struct
For each operation belonging to a service, an "options" struct is defined as a container for
the parameters associated with the operation.
The name of the struct will be `<operation-name>Options` and it will contain a field for each
operation parameter.  
Here's an example of an options struct for the `GetDnszone` operation:
```go
// GetDnszoneOptions : The GetDnszone options.
type GetDnszoneOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}
```
In this example, the `GetDnszone` operation has two required parameters - `InstanceID` and `DnszoneID`.
When invoking this operation, the application first creates an instance of the `GetDnszoneOptions`
struct and then sets the parameter values within it.  Along with the "options" struct, a constructor
function is also provided.  
Here's an example:
```go
options := service.NewGetDnszoneOptions("instance-id-1", "dnszone-id-1")
```
Then the operation can be called like this:
```go
result, detailedResponse, err := service.GetDnszone(options)
```
This use of the "options" struct pattern (instead of listing each operation parameter within the
argument list of the service method) allows for future expansion of the API (within certain
guidelines) without impacting applications.

### Receiving operation responses

Each service method (operation) will return the following values:
1. `result` - An operation-specific result (if the operation is defined as returning a result).
2. `detailedResponse` - An instance of the `core.DetailedResponse` struct.
This will contain the following fields:
* `StatusCode` - the HTTP status code returned in the response message
* `Headers` - the HTTP headers returned in the response message
* `Result` - the operation result (if available). This is the same value returned in the `result` return value
mentioned above.
3. `err` - An error object.  This return value will be nil if the operation was successful, or non-nil
if unsuccessful.

#### Example:
1. Here's an example of calling the `GetDnszone` operation which returns resource of the `Dnszone`
struct as its result:
```go
// Construct the service instance.
service, err := dnssvcsv1.NewDnsSvcsV1(
    &dnssvcsv1.DnsSvcsV1Options{
        Authenticator: authenticator,
    })

// Call the GetDnszone operation and receive the returned Resource.
options := service.NewGetDnszoneOptions("instance-id-1", "dnszone-id-1")
result, detailedResponse, err := service.GetDnszone(options)

// Now use 'result' which should be a dns zone of 'Dnszone'.
```
2. Here's an example of calling the `DeleteDnszone` operation which does not return a response object:
```go
// Construct the service instance.
service, err := dnssvcsv1.NewDnsSvcsV1(
    &dnssvcsv1.DnsSvcsV1Options{
        Authenticator: authenticator,
    })

// Call the DeleteDnszone operation and receive the returned Resource.
options := service.NewDeleteDnszoneOptions("instance-id-1", "dnszone-id-1")
detailedResponse, err := service.DeleteDnszone(options)
```

### Error Handling

In the case of an error response from the server endpoint, the `dns-svcs-go-sdk` will do the following:
1. The service method (operation) will return a non-nil `error` object.  This `error` object will
contain the error message retrieved from the HTTP response if possible, or a generic error message
otherwise.
2. The `detailedResponse.Result` field will contain the unmarshalled response (in the form of a
`map[string]interface{}`) if the operation returned a JSON response.  
This allows the application to examine all of the error information returned in the HTTP
response message.
3. The `detailedResponse.RawResult` field will contain the raw response body as a `[]byte` if the
operation returned a non-JSON response.

#### Example:
Here's an example of checking the `error` object after invoking the `GetDnszone` operation:
```go
// Call the GetDnszone operation and receive the returned Resource.
options := service.NewGetDnszoneOptions("bad-instance-id-1", "bad-dnszone-id-1")
result, detailedResponse, err := service.GetDnszone(options)
if err != nil {
    fmt.Println("Error retrieving the resource: ", err.Error())
    fmt.Println("   full error response: ", detailedResponse.Result)
}
```

### Default headers
Default HTTP headers can be specified by using the `SetDefaultHeaders(http.Header)`
method of the client instance.  Once set on the service client, default headers are sent with
every outbound request.  

#### Example:
The example below sets the header `Custom-Header` with the value "custom_value" as a default
header:
```go
// Construct the service instance.
service, err := dnssvcsv1.NewDnsSvcsV1(
    &dnssvcsv1.DnsSvcsV1Options{
        Authenticator: authenticator,
    })

customHeaders := http.Header{}
customHeaders.Add("Custom-Header", "custom_value")
service.Service.SetDefaultHeaders(customHeaders)

// "Custom-Header" will now be included with all subsequent requests invoked from "service".
```

### Sending request headers
Custom HTTP headers can also be passed with any individual request.
To do so, add the headers to the "options" struct passed to the service method.

#### Example:
Here's an example that sets "Custom-Header" on the `GetDnszoneOptions` instance and then
invokes the `GetDnszone` operation:
```go

// Call the GetDnszone operation, passing our Custom-Header.
options := service.NewGetDnszoneOptions("instance-id-1", "dnszone-id-1")
customHeaders := make(map[string]interface{})
customHeaders["Custom-Header"] = "custom_value"
options.SetHeaders(customHeaders)
result, detailedResponse, err := service.GetDnszone(options)
// "Custom-Header" will be sent along with the "GetDnszone" request.
```

### Transaction IDs

Every call from the SDK will receive a response which will contain a transaction ID, accessible via the `x-correlation-id` header. This transaction ID is useful for troubleshooting and accessing relevant logs from your service instance.

## License

The IBM Cloud DNS Services Go SDK is released under the Apache 2.0 license. The license's full text can be found in [LICENSE](LICENSE).
