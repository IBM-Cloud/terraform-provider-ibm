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
    "ibm_tekton_pipeline_workers": continuousdeliverypipeline.DataSourceIBMTektonPipelineWorkers(),
    "ibm_tekton_pipeline_trigger": continuousdeliverypipeline.DataSourceIBMTektonPipelineTrigger(),
    "ibm_tekton_pipeline": continuousdeliverypipeline.DataSourceIBMTektonPipeline(),
```

- Add the following entries to `ResourcesMap`:
```
    "ibm_tekton_pipeline_definition": continuousdeliverypipeline.ResourceIBMTektonPipelineDefinition(),
    "ibm_tekton_pipeline_trigger_property": continuousdeliverypipeline.ResourceIBMTektonPipelineTriggerProperty(),
    "ibm_tekton_pipeline_property": continuousdeliverypipeline.ResourceIBMTektonPipelineProperty(),
    "ibm_tekton_pipeline_trigger": continuousdeliverypipeline.ResourceIBMTektonPipelineTrigger(),
    "ibm_tekton_pipeline": continuousdeliverypipeline.ResourceIBMTektonPipeline(),
```

- Add the following entries to `globalValidatorDict`:
``` 
    "ibm_tekton_pipeline_definition": continuousdeliverypipeline.ResourceIBMTektonPipelineDefinitionValidator(),
    "ibm_tekton_pipeline_trigger_property": continuousdeliverypipeline.ResourceIBMTektonPipelineTriggerPropertyValidator(),
    "ibm_tekton_pipeline_property": continuousdeliverypipeline.ResourceIBMTektonPipelinePropertyValidator(),
    "ibm_tekton_pipeline_trigger": continuousdeliverypipeline.ResourceIBMTektonPipelineTriggerValidator(),
    "ibm_tekton_pipeline": continuousdeliverypipeline.ResourceIBMTektonPipelineValidator(),
```

### Changes to `config.go`

- Add an import for the generated Go SDK:
```
    "github.ibm.com/org-ids/tekton-pipeline-go-sdk/continuousdeliverypipelinev2"
```

- Add a method to the `ClientSession interface`:
```
    ContinuousDeliveryPipelineV2()   (*continuousdeliverypipelinev2.ContinuousDeliveryPipelineV2, error)
```

- Add two fields to the `clientSession struct`:
```
    continuousDeliveryPipelineClient     *continuousdeliverypipelinev2.ContinuousDeliveryPipelineV2
    continuousDeliveryPipelineClientErr  error
```

- Implement a new method on the `clientSession struct`:
```
// Continuous Delivery Pipeline
func (session clientSession) ContinuousDeliveryPipelineV2() (*continuousdeliverypipelinev2.ContinuousDeliveryPipelineV2, error) {
    return session.continuousDeliveryPipelineClient, session.continuousDeliveryPipelineClientErr
}
```

- In the `ClientSession()` method of `Config`, below the existing line that creates an authenticator:
```
    var authenticator *core.BearerTokenAuthenticator
```
  add the code to initialize the service client
```
    // Construct an "options" struct for creating the service client.
    continuousDeliveryPipelineClientOptions := &continuousdeliverypipelinev2.ContinuousDeliveryPipelineV2Options{
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
    }
```

### Changes to website/allowed-subcategories.txt  

Insert the following line into the website/allowed-subcategories.txt file (in alphabetic order):

```
Continuous Delivery Pipeline
``` 
