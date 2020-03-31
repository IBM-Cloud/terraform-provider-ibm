# Authentication
The go-sdk-core project supports the following types of authentication:
- Basic Authentication
- Bearer Token 
- Identity and Access Management (IAM)
- Cloud Pak for Data
- No Authentication

The SDK user configures the appropriate type of authentication for use with service instances.  
The authentication types that are appropriate for a particular service may vary from service to service, so it is important for the SDK user to consult with the appropriate service documentation to understand which authenticators are supported for that service.

The go-sdk-core allows an authenticator to be specified in one of two ways:
1. programmatically - the SDK user invokes the appropriate function(s) to create an instance of the desired authenticator and then passes the authenticator instance when constructing an instance of the service.
2. configuration - the SDK user provides external configuration information (in the form of environment variables or a credentials file) to indicate the type of authenticator along with the configuration of the necessary properties for that authenticator.  The SDK user then invokes the configuration-based authenticator factory to construct an instance of the authenticator that is described in the external configuration information.

The sections below will provide detailed information for each authenticator
which will include the following:
- A description of the authenticator
- The properties associated with the authenticator
- An example of how to construct the authenticator programmatically
- An example of how to configure the authenticator through the use of external
configuration information.  The configuration examples below will use
environment variables, although the same properties could be specified in a
credentials file instead.

## Basic Authentication
The `BasicAuthenticator` is used to add Basic Authentication information to
each outbound request in the `Authorization` header in the form:
```
   Authorization: Basic <encoded username and password>
```
### Properties
- Username: (required) the basic auth username
- Password: (required) the basic auth password
### Programming example
```go
import {
    "github.com/IBM/go-sdk-core/core"
    "<appropriate-git-repo-url>/exampleservicev1"
}
...
// Create the authenticator.
authenticator := &core.BasicAuthenticator{
    Username: "myuser",
    Password: "mypassword",
}

// Create the service options struct.
options := &exampleservicev1.ExampleServiceV1Options{
    Authenticator: authenticator,
}

// Construct the service instance.
service := exampleservicev1.NewExampleServiceV1(options)

// 'service' can now be used to invoke operations.
```
### Configuration example
External configuration:
```
export EXAMPLE_SERVICE_AUTH_TYPE=basic
export EXAMPLE_SERVICE_USERNAME=myuser
export EXAMPLE_SERVICE_PASSWORD=mypassword
```
Application code:
```go
import {
    "github.com/IBM/go-sdk-core/core"
    "<appropriate-git-repo-url>/exampleservicev1"
}
...
// Construct the authenticator from external configuration information for service "example_service".
authenticator := &core.GetAuthenticatorFromEnvironment("example_service")

// Create the service options struct.
options := &exampleservicev1.ExampleServiceV1Options{
    Authenticator: authenticator,
}

// Construct the service instance.
service := exampleservicev1.NewExampleServiceV1(options)

// 'service' can now be used to invoke operations.
```

## Bearer Token Authentication
The `BearerTokenAuthenticator` will add a user-supplied bearer token to
each outbound request in the `Authorization` header in the form:
```
    Authorization: Bearer <bearer-token>
```
### Properties
- BearerToken: (required) the bearer token value
### Programming example
```go
import {
    "github.com/IBM/go-sdk-core/core"
    "<appropriate-git-repo-url>/exampleservicev1"
}
...
// Create the authenticator.
bearerToken := // ... obtain bearer token value ...
authenticator := &core.BearerTokenAuthenticator{
    BearerToken: bearerToken,
}

// Create the service options struct.
options := &exampleservicev1.ExampleServiceV1Options{
    Authenticator: authenticator,
}

// Construct the service instance.
service := exampleservicev1.NewExampleServiceV1(options)

// 'service' can now be used to invoke operations.
...
// Later, if your bearer token value expires, you can set a new one like this:
newToken := // ... obtain new bearer token value
authenticator.BearerToken = newToken
```
### Configuration example
External configuration:
```
export EXAMPLE_SERVICE_AUTH_TYPE=bearertoken
export EXAMPLE_SERVICE_BEARER_TOKEN=<the bearer token value>
```
Application code:
```go
import {
    "github.com/IBM/go-sdk-core/core"
    "<appropriate-git-repo-url>/exampleservicev1"
}
...
// Construct the authenticator from external configuration information for service "example_service".
authenticator := &core.GetAuthenticatorFromEnvironment("example_service")

// Create the service options struct.
options := &exampleservicev1.ExampleServiceV1Options{
    Authenticator: authenticator,
}

// Construct the service instance.
service := exampleservicev1.NewExampleServiceV1(options)

// 'service' can now be used to invoke operations.
...
// Later, if your bearer token value expires, you can set a new one like this:
newToken := // ... obtain new bearer token value
authenticator.BearerToken = newToken
```
Note that the use of external configuration is not as useful with the `BearerTokenAuthenticator` as it
is for other authenticator types because bearer tokens typically need to be obtained and refreshed
programmatically since they normally have a relatively short lifespan before they expire.  This
authenticator type is intended for situations in which the application will be managing the bearer 
token itself in terms of initial acquisition and refreshing as needed.

## Identity and Access Management Authentication (IAM)
The `IamAuthenticator` will accept a user-supplied api key and will perform
the necessary interactions with the IAM token service to obtain a suitable
bearer token for the specified api key.  The authenticator will also obtain 
a new bearer token when the current token expires.  The bearer token is 
then added to each outbound request in the `Authorization` header in the
form:
```
   Authorization: Bearer <bearer-token>
```
### Properties
- ApiKey: (required) the IAM api key
- URL: (optional) The URL representing the IAM token service endpoint.  If not specified, a suitable
default value is used.
- ClientId/ClientSecret: (optional) The `ClientId` and `ClientSecret` fields are used to form a 
"basic auth" Authorization header for interactions with the IAM token server. If neither field 
is specified, then no Authorization header will be sent with token server requests.  These fields 
are optional, but must be specified together.
- DisableSSLVerification: (optional) A flag that indicates whether verificaton of the server's SSL 
certificate should be disabled or not. The default value is `false`.
- Headers: (optional) A set of key/value pairs that will be sent as HTTP headers in requests
made to the IAM token service.
- Client: (Optional) The `http.Client` object used to invoke token servive requests. If not specified
by the user, a suitable default Client will be constructed.
### Programming example
```go
import {
    "github.com/IBM/go-sdk-core/core"
    "<appropriate-git-repo-url>/exampleservicev1"
}
...
// Create the authenticator.
authenticator := &core.IamAuthenticator{
    ApiKey: "myapikey",
}

// Create the service options struct.
options := &exampleservicev1.ExampleServiceV1Options{
    Authenticator: authenticator,
}

// Construct the service instance.
service := exampleservicev1.NewExampleServiceV1(options)

// 'service' can now be used to invoke operations.
```
### Configuration example
External configuration:
```
export EXAMPLE_SERVICE_AUTH_TYPE=iam
export EXAMPLE_SERVICE_APIKEY=myapikey
```
Application code:
```go
import {
    "github.com/IBM/go-sdk-core/core"
    "<appropriate-git-repo-url>/exampleservicev1"
}
...
// Construct the authenticator from external configuration information for service "example_service".
authenticator := &core.GetAuthenticatorFromEnvironment("example_service")

// Create the service options struct.
options := &exampleservicev1.ExampleServiceV1Options{
    Authenticator: authenticator,
}

// Construct the service instance.
service := exampleservicev1.NewExampleServiceV1(options)

// 'service' can now be used to invoke operations.
```

##  Cloud Pak for Data
The `CloudPakForDataAuthenticator` will accept user-supplied username and password values, and will 
perform the necessary interactions with the Cloud Pak for Data token service to obtain a suitable
bearer token.  The authenticator will also obtain a new bearer token when the current token expires.
The bearer token is then added to each outbound request in the `Authorization` header in the
form:
```
   Authorization: Bearer <bearer-token>
```
### Properties
- Username: (required) the username used to obtain a bearer token.
- Password: (required) the password used to obtain a bearer token.
- URL: (required) The URL representing the Cloud Pak for Data token service endpoint.
- DisableSSLVerification: (optional) A flag that indicates whether verificaton of the server's SSL 
certificate should be disabled or not. The default value is `false`.
- Headers: (optional) A set of key/value pairs that will be sent as HTTP headers in requests
made to the IAM token service.
- Client: (Optional) The `http.Client` object used to invoke token servive requests. If not specified
by the user, a suitable default Client will be constructed.
### Programming example
```go
import {
    "github.com/IBM/go-sdk-core/core"
    "<appropriate-git-repo-url>/exampleservicev1"
}
...
// Create the authenticator.
authenticator := &core.CloudPakForDataAuthenticator{
    Username: "myuser",
    Password: "mypassword",
    URL: "https://mycp4dhost.com/",
}

// Create the service options struct.
options := &exampleservicev1.ExampleServiceV1Options{
    Authenticator: authenticator,
}

// Construct the service instance.
service := exampleservicev1.NewExampleServiceV1(options)

// 'service' can now be used to invoke operations.
```
### Configuration example
External configuration:
```
export EXAMPLE_SERVICE_AUTH_TYPE=cp4d
export EXAMPLE_SERVICE_USERNAME=myuser
export EXAMPLE_SERVICE_PASSWORD=mypassword
export EXAMPLE_SERVICE_URL=https://mycp4dhost.com/
```
Application code:
```go
import {
    "github.com/IBM/go-sdk-core/core"
    "<appropriate-git-repo-url>/exampleservicev1"
}
...
// Construct the authenticator from external configuration information for service "example_service".
authenticator := &core.GetAuthenticatorFromEnvironment("example_service")

// Create the service options struct.
options := &exampleservicev1.ExampleServiceV1Options{
    Authenticator: authenticator,
}

// Construct the service instance.
service := exampleservicev1.NewExampleServiceV1(options)

// 'service' can now be used to invoke operations.
```
## No Auth Authentication
The `NoAuthAuthenticator` is a placeholder authenticator which performs no actual authentication function.   It can be used in situations where authentication needs to be bypassed, perhaps while developing or debugging an application or service.
### Properties
None
### Programming example
```go
import {
    "github.com/IBM/go-sdk-core/core"
    "<appropriate-git-repo-url>/exampleservicev1"
}
...
// Create the authenticator.
authenticator := &core.NoAuthAuthenticator{}

// Create the service options struct.
options := &exampleservicev1.ExampleServiceV1Options{
    Authenticator: authenticator,
}

// Construct the service instance.
service := exampleservicev1.NewExampleServiceV1(options)

// 'service' can now be used to invoke operations.
```
### Configuration example
External configuration:
```
export EXAMPLE_SERVICE_AUTH_TYPE=noauth
```
Application code:
```go
import {
    "github.com/IBM/go-sdk-core/core"
    "<appropriate-git-repo-url>/exampleservicev1"
}
...
// Construct the authenticator from external configuration information for service "example_service".
authenticator := &core.GetAuthenticatorFromEnvironment("example_service")

// Create the service options struct.
options := &exampleservicev1.ExampleServiceV1Options{
    Authenticator: authenticator,
}

// Construct the service instance.
service := exampleservicev1.NewExampleServiceV1(options)

// 'service' can now be used to invoke operations.
```
