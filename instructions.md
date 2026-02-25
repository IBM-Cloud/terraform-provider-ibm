# Instructions

To add the generated code into the IBM Terraform Provider, you will need to make the following changes to the project. Note that these changes have already been generated into this local development environment, with the exception of the change to `website/allowed-subcategories.txt` which is unnecessary for local development.

### Changes to `provider.go`

- Add the following entry to `import`:
```
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/iamidentity"
```

- Add the following entries to `DataSourcesMap`:
```
	"ibm_iam_identity_preference": iamidentity.DataSourceIBMIamIdentityPreference(),
```

- Add the following entries to `ResourcesMap`:
```
	"ibm_iam_identity_preference": iamidentity.ResourceIBMIamIdentityPreference(),
```

### Changes to `config.go`

- Add an import for the generated Go SDK:
```
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
```

- Add a method to the `ClientSession interface`:
```
	IamIdentityV1() (*iamidentityv1.IamIdentityV1, error)
```

- Add two fields to the `clientSession struct`:
```
	iamIdentityClient     *iamidentityv1.IamIdentityV1
	iamIdentityClientErr  error
```

- Implement a new method on the `clientSession struct`:
```
	// IAM Identity Services
	func (session clientSession) IamIdentityV1() (*iamidentityv1.IamIdentityV1, error) {
		return session.iamIdentityClient, session.iamIdentityClientErr
	}
```

- In the `ClientSession()` method of `Config`, below the existing block of code that creates an authenticator:
```
	var authenticator core.Authenticator
	if c.BluemixAPIKey != "" {
		authenticator = &core.IamAuthenticator{
			ApiKey: c.BluemixAPIKey,
			URL:    EnvFallBack([]string{"IBMCLOUD_IAM_API_ENDPOINT"}, "https://iam.cloud.ibm.com") + "/identity/token",
		}
	} else if strings.HasPrefix(sess.BluemixSession.Config.IAMAccessToken, "Bearer") {
		authenticator = &core.BearerTokenAuthenticator{
			BearerToken: sess.BluemixSession.Config.IAMAccessToken[7:],
		}
	} else {
		authenticator = &core.BearerTokenAuthenticator{
			BearerToken: sess.BluemixSession.Config.IAMAccessToken,
		}
	}
```

- In the `ClientSession()` method of `Config`, add the code to initialize the service client:
```

	// Construct an instance of the 'IAM Identity Services' service.
	if session.iamIdentityClientErr == nil {
		// Construct the service options.
		iamIdentityClientOptions := &iamidentityv1.IamIdentityV1Options{
			Authenticator: authenticator,
		}

		// Construct the service client.
		session.iamIdentityClient, err = iamidentityv1.NewIamIdentityV1(iamIdentityClientOptions)
		if err == nil {
			// Enable retries for API calls
			session.iamIdentityClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
			// Add custom header for analytics
			session.iamIdentityClient.SetDefaultHeaders(gohttp.Header{
				"X-Original-User-Agent": { fmt.Sprintf("terraform-provider-ibm/%s", version.Version) },
			})
		} else {
			session.iamIdentityClientErr = fmt.Errorf("Error occurred while constructing 'IAM Identity Services' service client: %q", err)
		}
	}
```

### Changes to website/allowed-subcategories.txt  

Insert the following line into the website/allowed-subcategories.txt file (in alphabetic order):

```
IAM Identity Services
``` 
