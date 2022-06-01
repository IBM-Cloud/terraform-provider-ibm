# Instructions

To add this generated code into the IBM Terraform Provider:

### Changes to `provider.go`

- Add the following entry to `import`:
```
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cdtoolchain"
```

- Add the following entries to `DataSourcesMap`:
```
    "ibm_cd_toolchain_tool_sonarqube": cdtoolchain.DataSourceIBMCdToolchainToolSonarqube(),
```

- Add the following entries to `ResourcesMap`:
```
    "ibm_cd_toolchain_tool_sonarqube": cdtoolchain.ResourceIBMCdToolchainToolSonarqube(),
```

- Add the following entries to `globalValidatorDict`:
``` 
    "ibm_cd_toolchain_tool_sonarqube": cdtoolchain.ResourceIBMCdToolchainToolSonarqubeValidator(),
```

### Changes to `config.go`

- Add an import for the generated Go SDK:
```
    "github.ibm.com/org-ids/toolchain-go-sdk/cdtoolchainv2"
```

- Add a method to the `ClientSession interface`:
```
    CdToolchainV2()   (*cdtoolchainv2.CdToolchainV2, error)
```

- Add two fields to the `clientSession struct`:
```
    cdToolchainClient     *cdtoolchainv2.CdToolchainV2
    cdToolchainClientErr  error
```

- Implement a new method on the `clientSession struct`:
```
// CD Toolchain
func (session clientSession) CdToolchainV2() (*cdtoolchainv2.CdToolchainV2, error) {
    return session.cdToolchainClient, session.cdToolchainClientErr
}
```

- In the `ClientSession()` method of `Config`, below the existing line that creates an authenticator:
```
    var authenticator *core.BearerTokenAuthenticator
```
  add the code to initialize the service client
```
    // Construct an "options" struct for creating the service client.
    var cdToolchainClientURL string
    if c.Visibility == "private" || c.Visibility == "public-and-private" {
        cdToolchainClientURL, err = cdtoolchainv2.GetServiceURLForRegion("private." + c.Region)
        if err != nil && c.Visibility == "public-and-private" {
            cdToolchainClientURL, err = cdtoolchainv2.GetServiceURLForRegion(c.Region)
        }
    } else {
        cdToolchainClientURL, err = cdtoolchainv2.GetServiceURLForRegion(c.Region)
    }
    if err != nil {
        cdToolchainClientURL = cdtoolchainv2.DefaultServiceURL
    }
    if fileMap != nil && c.Visibility != "public-and-private" {
		cdToolchainClientURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_CR_API_ENDPOINT", c.Region, cdToolchainClientURL)
	}
    cdToolchainClientOptions := &cdtoolchainv2.CdToolchainV2Options{
        Authenticator: authenticator,
        URL: EnvFallBack([]string{"IBMCLOUD_TOOLCHAIN_ENDPOINT"}, cdToolchainClientURL),
    }

    // Construct the service client.
    session.cdToolchainClient, err = cdtoolchainv2.NewCdToolchainV2(cdToolchainClientOptions)
    if err == nil {
        // Enable retries for API calls
        session.cdToolchainClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
        // Add custom header for analytics
        session.cdToolchainClient.SetDefaultHeaders(gohttp.Header{
            "X-Original-User-Agent": { fmt.Sprintf("terraform-provider-ibm/%s", version.Version) },
        })
    } else {
        session.cdToolchainClientErr = fmt.Errorf("Error occurred while configuring CD Toolchain service: %q", err)
    }
```

### Changes to website/allowed-subcategories.txt  

Insert the following line into the website/allowed-subcategories.txt file (in alphabetic order):

```
CD Toolchain
``` 
