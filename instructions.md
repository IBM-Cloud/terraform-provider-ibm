# Instructions

To add this generated code into the IBM Terraform Provider:

### Changes to `provider.go`

- Add the following entry to `import`:
```

	"github.ibm.com/cloudengineering/terraform-provider-template/ibm/service/continuousdeliverypipeline"
```

- Add the following entries to `DataSourcesMap`:
```
    "ibm_tekton_pipeline_definition": continuousdeliverypipeline.DataSourceIBMTektonPipelineDefinition(),
    "ibm_tekton_pipeline_trigger_property": continuousdeliverypipeline.DataSourceIBMTektonPipelineTriggerProperty(),
    "ibm_tekton_pipeline_property": continuousdeliverypipeline.DataSourceIBMTektonPipelineProperty(),
    "ibm_tekton_pipeline_trigger": continuousdeliverypipeline.DataSourceIBMTektonPipelineTrigger(),
	"github.ibm.com/cloudengineering/terraform-provider-template/ibm/service/ibmtoolchainapi"

```

- Add the following entries to `ResourcesMap`:


### Changes to `config.go`

- Add an import for the generated Go SDK:
```

    "github.ibm.com/org-ids/tekton-pipeline-go-sdk/continuousdeliverypipelinev2"

    "github.ibm.com/org-ids/toolchain-go-sdk/ibmtoolchainapiv2"

```

- Add a method to the `ClientSession interface`:
```

    ContinuousDeliveryPipelineV2()   (*continuousdeliverypipelinev2.ContinuousDeliveryPipelineV2, error)

    IbmToolchainApiV2()   (*ibmtoolchainapiv2.IbmToolchainApiV2, error)

```

- Add two fields to the `clientSession struct`:
```

    continuousDeliveryPipelineClient     *continuousdeliverypipelinev2.ContinuousDeliveryPipelineV2
    continuousDeliveryPipelineClientErr  error

    ibmToolchainApiClient     *ibmtoolchainapiv2.IbmToolchainApiV2
    ibmToolchainApiClientErr  error

```

- Implement a new method on the `clientSession struct`:
```

// Continuous Delivery Pipeline
func (session clientSession) ContinuousDeliveryPipelineV2() (*continuousdeliverypipelinev2.ContinuousDeliveryPipelineV2, error) {
    return session.continuousDeliveryPipelineClient, session.continuousDeliveryPipelineClientErr

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

    var continuousDeliveryPipelineClientURL string
    if c.Visibility == "private" || c.Visibility == "public-and-private" {
        continuousDeliveryPipelineClientURL, err = continuousdeliverypipelinev2.GetServiceURLForRegion("private." + c.Region)
        if err != nil && c.Visibility == "public-and-private" {
            continuousDeliveryPipelineClientURL, err = continuousdeliverypipelinev2.GetServiceURLForRegion(c.Region)
        }
    } else {
        continuousDeliveryPipelineClientURL, err = continuousdeliverypipelinev2.GetServiceURLForRegion(c.Region)
    }
    if err != nil {
        continuousDeliveryPipelineClientURL = continuousdeliverypipelinev2.DefaultServiceURL
    }
    if fileMap != nil && c.Visibility != "public-and-private" {
		continuousDeliveryPipelineClientURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_CR_API_ENDPOINT", c.Region, continuousDeliveryPipelineClientURL)
	}
    continuousDeliveryPipelineClientOptions := &continuousdeliverypipelinev2.ContinuousDeliveryPipelineV2Options{

    ibmToolchainApiClientOptions := &ibmtoolchainapiv2.IbmToolchainApiV2Options{

        Authenticator: authenticator,
    }

    // Construct the service client.

    session.continuousDeliveryPipelineClient, err = continuousdeliverypipelinev2.NewContinuousDeliveryPipelineV2(continuousDeliveryPipelineClientOptions)
    if err == nil {
        // Enable retries for API calls
        session.continuousDeliveryPipelineClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
        // Add custom header for analytics
        session.continuousDeliveryPipelineClient.SetDefaultHeaders(gohttp.Header{
            "X-Original-User-Agent": { fmt.Sprintf("terraform-provider-ibm/%s", version.Version) },
        })
    } else {
        session.continuousDeliveryPipelineClientErr = fmt.Errorf("Error occurred while configuring Continuous Delivery Pipeline service: %q", err)

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

Continuous Delivery Pipeline

IBM Toolchain API

``` 
