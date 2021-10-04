# Instructions

To add this generated code into the IBM Terraform Provider:

### Changes to `provider.go`

- Add the following entries to `DataSourcesMap`:
```
    "ibm_cbr_zone": dataSourceIBMCbrZone(),
    "ibm_cbr_rule": dataSourceIBMCbrRule(),
```

- Add the following entries to `ResourcesMap`:
```
    "ibm_cbr_zone": resourceIBMCbrZone(),
    "ibm_cbr_rule": resourceIBMCbrRule(),
```

- Add the following entries to `globalValidatorDict`:
```
    "ibm_cbr_zone": resourceIBMCbrZoneValidator(),
    "ibm_cbr_rule": resourceIBMCbrRuleValidator(),
```

### Changes to `config.go`

- Add an import for the generated Go SDK:
```
    "github.com/IBM/platform-services-go-sdk/contextbasedrestrictionsv1"
```

- Add a method to the `ClientSession interface`:
```
    ContextBasedRestrictionsV1()   (*contextbasedrestrictionsv1.ContextBasedRestrictionsV1, error)
```

- Add two fields to the `clientSession struct`:
```
    contextBasedRestrictionsClient     *contextbasedrestrictionsv1.ContextBasedRestrictionsV1
    contextBasedRestrictionsClientErr  error
```

- Implement a new method on the `clientSession struct`:
```
// Context Based Restrictions
func (session clientSession) ContextBasedRestrictionsV1() (*contextbasedrestrictionsv1.ContextBasedRestrictionsV1, error) {
    return session.contextBasedRestrictionsClient, session.contextBasedRestrictionsClientErr
}
```

- In the `ClientSession()` method of `Config`, below the existing line that creates an authenticator:
```
    var authenticator *core.BearerTokenAuthenticator
```
  add the code to initialize the service client
```
    // Construct an "options" struct for creating the service client.
    // zn -- GO SDK changed ContextBasedRestrictionsV1Options changed to Options 
    contextBasedRestrictionsClientOptions := &contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
        Authenticator: authenticator,
    }

    // Construct the service client.
    session.contextBasedRestrictionsClient, err = contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(contextBasedRestrictionsClientOptions)
    if err == nil {
        // Enable retries for API calls
        session.contextBasedRestrictionsClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
        // Add custom header for analytics
        session.contextBasedRestrictionsClient.SetDefaultHeaders(gohttp.Header{
            "X-Original-User-Agent": { fmt.Sprintf("terraform-provider-ibm/%s", version.Version) },
        })
    } else {
        session.contextBasedRestrictionsClientErr = fmt.Errorf("Error occurred while configuring Context Based Restrictions service: %q", err)
    }
```

### Changes to website/allowed-subcategories.txt  

Insert the following line into the website/allowed-subcategories.txt file (in alphabetic order):

```
Context Based Restrictions
``` 
