# Instructions

To add this generated code into the IBM Terraform Provider:

### Changes to `provider.go`

- Add the following entries to `ResourcesMap`:
```
    "ibm_is_dedicated_host_group": resourceIbmIsDedicatedHostGroup(),
```

- Add the following entries to `globalValidatorDict`:
```
    "ibm_is_dedicated_host_group": resourceIbmIsDedicatedHostGroupValidator(),
```

### Changes to `config.go`

- Add an import for the generated Go SDK:
```
    "github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
```

- Add a method to the `ClientSession interface`:
```
    VpcV1()   (*vpcv1.VpcV1, error)
```

- Add two fields to the `clientSession struct`:
```
    vpcClient     *vpcv1.VpcV1
    vpcClientErr  error
```

- Implement a new method on the `clientSession struct`:
```
    func (session clientSession) VpcV1() (*vpcv1.VpcV1, error) {
        return session.vpcClient, session.vpcClientErr
    }
```

- In the `ClientSession()` method of `Config`, below the existing line that creates an authenticator:
```
    var authenticator *core.BearerTokenAuthenticator
```
  add the code to initialize the service client
```
    // Construct an "options" struct for creating the service client.
    vpcClientOptions := &vpcv1.VpcV1Options{
        Authenticator: authenticator,
        Version: core.TypeStringPtr("testString"),
        generation: core.TypeIntPtr(int64(2)),
    }

    // Construct the service client.
    session.vpcClient, err = vpcv1.NewVpcV1(vpcClientOptions)
    if err == nil {
        // Enable retries for API calls
        session.vpcClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
        // Add custom header for analytics
        session.vpcClient.SetDefaultHeaders(gohttp.Header{
            "X-Original-User-Agent": { fmt.Sprintf("terraform-provider-ibm/%s", version.Version) },
        })
    } else {
        session.vpcClientErr = fmt.Errorf("Error occurred while configuring Virtual Private Cloud API service: %q", err)
    }
```
