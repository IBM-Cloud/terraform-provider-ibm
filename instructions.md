# Instructions

To add this generated code into the IBM Terraform Provider:

### Changes to `provider.go`

- Add the following entry to `import`:
```
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/toolchain"
```

- Add the following entries to `ResourcesMap`:
```
    "ibm_toolchain_tool_sonarqube": toolchain.ResourceIBMToolchainToolSonarqube(),
```

- Add the following entries to `globalValidatorDict`:
``` 
    "ibm_toolchain_tool_sonarqube": toolchain.ResourceIBMToolchainToolSonarqubeValidator(),
```

### Changes to `config.go`

- Add an import for the generated Go SDK:
```
    "github.ibm.com/org-ids/toolchain-go-sdk/toolchainv2"
```

- Add a method to the `ClientSession interface`:
```
    ToolchainV2()   (*toolchainv2.ToolchainV2, error)
```

- Add two fields to the `clientSession struct`:
```
    toolchainClient     *toolchainv2.ToolchainV2
    toolchainClientErr  error
```

- Implement a new method on the `clientSession struct`:
```
// Toolchain
func (session clientSession) ToolchainV2() (*toolchainv2.ToolchainV2, error) {
    return session.toolchainClient, session.toolchainClientErr
}
```

- In the `ClientSession()` method of `Config`, below the existing line that creates an authenticator:
```
    var authenticator *core.BearerTokenAuthenticator
```
  add the code to initialize the service client
```
    // Construct an "options" struct for creating the service client.
    var toolchainClientURL string
    if c.Visibility == "private" || c.Visibility == "public-and-private" {
        toolchainClientURL, err = toolchainv2.GetServiceURLForRegion("private." + c.Region)
        if err != nil && c.Visibility == "public-and-private" {
            toolchainClientURL, err = toolchainv2.GetServiceURLForRegion(c.Region)
        }
    } else {
        toolchainClientURL, err = toolchainv2.GetServiceURLForRegion(c.Region)
    }
    if err != nil {
        toolchainClientURL = toolchainv2.DefaultServiceURL
    }
    if fileMap != nil && c.Visibility != "public-and-private" {
		toolchainClientURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_CR_API_ENDPOINT", c.Region, toolchainClientURL)
	}
    toolchainClientOptions := &toolchainv2.ToolchainV2Options{
        Authenticator: authenticator,
        URL: EnvFallBack([]string{"IBMCLOUD_TOOLCHAIN_ENDPOINT"}, toolchainClientURL),
    }

    // Construct the service client.
    session.toolchainClient, err = toolchainv2.NewToolchainV2(toolchainClientOptions)
    if err == nil {
        // Enable retries for API calls
        session.toolchainClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
        // Add custom header for analytics
        session.toolchainClient.SetDefaultHeaders(gohttp.Header{
            "X-Original-User-Agent": { fmt.Sprintf("terraform-provider-ibm/%s", version.Version) },
        })
    } else {
        session.toolchainClientErr = fmt.Errorf("Error occurred while configuring Toolchain service: %q", err)
    }
```

### Changes to website/allowed-subcategories.txt  

Insert the following line into the website/allowed-subcategories.txt file (in alphabetic order):

```
Toolchain
``` 
