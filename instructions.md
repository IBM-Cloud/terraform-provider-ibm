# Instructions

To add the generated code into the IBM Terraform Provider, you will need to make the following changes to the project. Note that these changes have already been generated into this local development environment, with the exception of the change to `website/allowed-subcategories.txt` which is unnecessary for local development.

### Changes to `provider.go`

- Add the following entry to `import`:
```
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/secretsmanager"
```

- Add the following entries to `ResourcesMap`:
```
    "ibm_secret_group": secretsmanager.ResourceIbmSecretGroup(),
```

- Add the following entries to `globalValidatorDict`:
``` 
    "ibm_secret_group": secretsmanager.ResourceIbmSecretGroupValidator(),
```

### Changes to `config.go`

- Add an import for the generated Go SDK:
```
    "github.com/IBM-Cloud/secrets-manager-mt-go-sdk/secretsmanagerv1"
```

- Add a method to the `ClientSession interface`:
```
    SecretsManagerV1()   (*secretsmanagerv1.SecretsManagerV1, error)
```

- Add two fields to the `clientSession struct`:
```
    secretsManagerClient     *secretsmanagerv1.SecretsManagerV1
    secretsManagerClientErr  error
```

- Implement a new method on the `clientSession struct`:
```
// IBM Cloud Secrets Manager Basic API
func (session clientSession) SecretsManagerV1() (*secretsmanagerv1.SecretsManagerV1, error) {
    if session.secretsManagerClientErr != nil {
        return session.secretsManagerClient, session.secretsManagerClientErr
    }
    return session.secretsManagerClient.Clone(), nil
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
  add the code to initialize the service client
```
    // Construct an "options" struct for creating the service client.
    secretsManagerClientOptions := &secretsmanagerv1.SecretsManagerV1Options{
        Authenticator: authenticator,
        URL: EnvFallBack([]string{"MOCK_ENDPOINT"}, secretsmanagerv1.DefaultServiceURL),
        XInstanceCrn: core.StringPtr("crn:v1:bluemix:public:secrets-manager:us-south:a/321f5eb20987423e97aa9876f18b7c11:b49ad24d-71d4-4ebc-b9b9-a0937d1c84d0::"),
    }

    // Construct the service client.
    session.secretsManagerClient, err = secretsmanagerv1.NewSecretsManagerV1(secretsManagerClientOptions)
    if err == nil {
        // Enable retries for API calls
        session.secretsManagerClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
        // Add custom header for analytics
        session.secretsManagerClient.SetDefaultHeaders(gohttp.Header{
            "X-Original-User-Agent": { fmt.Sprintf("terraform-provider-ibm/%s", version.Version) },
        })
    } else {
        session.secretsManagerClientErr = fmt.Errorf("Error occurred while configuring IBM Cloud Secrets Manager Basic API service: %q", err)
    }
```

### Changes to website/allowed-subcategories.txt  

Insert the following line into the website/allowed-subcategories.txt file (in alphabetic order):

```
IBM Cloud Secrets Manager Basic API
``` 
