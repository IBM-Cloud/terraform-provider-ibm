# Instructions

To add this generated code into the IBM Terraform Provider:

### Changes to `provider.go`

- Add the following entry to `import`:
```
	"github.ibm.com/cloudengineering/terraform-provider-template/ibm/service/ibmtoolchainapi"
```

- Add the following entries to `ResourcesMap`:
```
    "ibm_toolchain_tool_git": ibmtoolchainapi.ResourceIbmToolchainToolGit(),
```

### Changes to `config.go`

- Add an import for the generated Go SDK:
```
    "github.ibm.com/org-ids/toolchain-go-sdk/ibmtoolchainapiv2"
```

- Add a method to the `ClientSession interface`:
```
    IbmToolchainApiV2()   (*ibmtoolchainapiv2.IbmToolchainApiV2, error)
```

- Add two fields to the `clientSession struct`:
```
    ibmToolchainApiClient     *ibmtoolchainapiv2.IbmToolchainApiV2
    ibmToolchainApiClientErr  error
```

- Implement a new method on the `clientSession struct`:
```
// IBM Toolchain API
func (session clientSession) IbmToolchainApiV2() (*ibmtoolchainapiv2.IbmToolchainApiV2, error) {
    return session.ibmToolchainApiClient, session.ibmToolchainApiClientErr
}
```

- In the `ClientSession()` method of `Config`, below the existing line that creates an authenticator:
```
    var authenticator *core.BearerTokenAuthenticator
```
  add the code to initialize the service client
```
    // Construct an "options" struct for creating the service client.
    ibmToolchainApiClientOptions := &ibmtoolchainapiv2.IbmToolchainApiV2Options{
        Authenticator: authenticator,
    }

    // Construct the service client.
    session.ibmToolchainApiClient, err = ibmtoolchainapiv2.NewIbmToolchainApiV2(ibmToolchainApiClientOptions)
    if err == nil {
        // Enable retries for API calls
        session.ibmToolchainApiClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
        // Add custom header for analytics
        session.ibmToolchainApiClient.SetDefaultHeaders(gohttp.Header{
            "X-Original-User-Agent": { fmt.Sprintf("terraform-provider-ibm/%s", version.Version) },
        })
    } else {
        session.ibmToolchainApiClientErr = fmt.Errorf("Error occurred while configuring IBM Toolchain API service: %q", err)
    }
```

### Changes to website/allowed-subcategories.txt  

Insert the following line into the website/allowed-subcategories.txt file (in alphabetic order):

```
IBM Toolchain API
``` 
