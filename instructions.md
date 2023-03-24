# Instructions

To add the generated code into the IBM Terraform Provider, you will need to make the following changes to the project. Note that these changes have already been generated into this local development environment, with the exception of the change to `website/allowed-subcategories.txt` which is unnecessary for local development.

### Changes to `provider.go`

- Add the following entry to `import`:
```
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/project"
```

### Changes to `config.go`

- Add an import for the generated Go SDK:
```
    "github.com/damianovesperini/platform-services-go-sdk/projectv1"
```

- Add a method to the `ClientSession interface`:
```
    ProjectV1()   (*projectv1.ProjectV1, error)
```

- Add two fields to the `clientSession struct`:
```
    projectClient     *projectv1.ProjectV1
    projectClientErr  error
```

- Implement a new method on the `clientSession struct`:
```
// Projects API Specification
func (session clientSession) ProjectV1() (*projectv1.ProjectV1, error) {
    return session.projectClient, session.projectClientErr
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
    projectClientOptions := &projectv1.ProjectV1Options{
        Authenticator: authenticator,
    }

    // Construct the service client.
    session.projectClient, err = projectv1.NewProjectV1(projectClientOptions)
    if err == nil {
        // Enable retries for API calls
        session.projectClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
        // Add custom header for analytics
        session.projectClient.SetDefaultHeaders(gohttp.Header{
            "X-Original-User-Agent": { fmt.Sprintf("terraform-provider-ibm/%s", version.Version) },
        })
    } else {
        session.projectClientErr = fmt.Errorf("Error occurred while configuring Projects API Specification service: %q", err)
    }
```

### Changes to website/allowed-subcategories.txt  

Insert the following line into the website/allowed-subcategories.txt file (in alphabetic order):

```
Projects API Specification
``` 
