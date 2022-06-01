# Instructions

To add this generated code into the IBM Terraform Provider:

### Changes to `provider.go`

- Add the following entry to `import`:
```
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cdtektonpipeline"
```

- Add the following entries to `DataSourcesMap`:
```
    "ibm_tekton_pipeline_definition": cdtektonpipeline.DataSourceIBMTektonPipelineDefinition(),
    "ibm_tekton_pipeline_trigger_property": cdtektonpipeline.DataSourceIBMTektonPipelineTriggerProperty(),
    "ibm_tekton_pipeline_property": cdtektonpipeline.DataSourceIBMTektonPipelineProperty(),
    "ibm_tekton_pipeline_trigger": cdtektonpipeline.DataSourceIBMTektonPipelineTrigger(),
    "ibm_tekton_pipeline": cdtektonpipeline.DataSourceIBMTektonPipeline(),
```

- Add the following entries to `ResourcesMap`:
```
    "ibm_tekton_pipeline_definition": cdtektonpipeline.ResourceIBMTektonPipelineDefinition(),
    "ibm_tekton_pipeline_trigger_property": cdtektonpipeline.ResourceIBMTektonPipelineTriggerProperty(),
    "ibm_tekton_pipeline_property": cdtektonpipeline.ResourceIBMTektonPipelineProperty(),
    "ibm_tekton_pipeline_trigger": cdtektonpipeline.ResourceIBMTektonPipelineTrigger(),
    "ibm_tekton_pipeline": cdtektonpipeline.ResourceIBMTektonPipeline(),
```

- Add the following entries to `globalValidatorDict`:
``` 
    "ibm_tekton_pipeline_definition": cdtektonpipeline.ResourceIBMTektonPipelineDefinitionValidator(),
    "ibm_tekton_pipeline_trigger_property": cdtektonpipeline.ResourceIBMTektonPipelineTriggerPropertyValidator(),
    "ibm_tekton_pipeline_property": cdtektonpipeline.ResourceIBMTektonPipelinePropertyValidator(),
    "ibm_tekton_pipeline_trigger": cdtektonpipeline.ResourceIBMTektonPipelineTriggerValidator(),
```

### Changes to `config.go`

- Add an import for the generated Go SDK:
```
    "github.ibm.com/org-ids/tekton-pipeline-go-sdk/cdtektonpipelinev2"
```

- Add a method to the `ClientSession interface`:
```
    CdTektonPipelineV2()   (*cdtektonpipelinev2.CdTektonPipelineV2, error)
```

- Add two fields to the `clientSession struct`:
```
    cdTektonPipelineClient     *cdtektonpipelinev2.CdTektonPipelineV2
    cdTektonPipelineClientErr  error
```

- Implement a new method on the `clientSession struct`:
```
// CD Tekton Pipeline
func (session clientSession) CdTektonPipelineV2() (*cdtektonpipelinev2.CdTektonPipelineV2, error) {
    return session.cdTektonPipelineClient, session.cdTektonPipelineClientErr
}
```

- In the `ClientSession()` method of `Config`, below the existing line that creates an authenticator:
```
    var authenticator *core.BearerTokenAuthenticator
```
  add the code to initialize the service client
```
    // Construct an "options" struct for creating the service client.
    var cdTektonPipelineClientURL string
    if c.Visibility == "private" || c.Visibility == "public-and-private" {
        cdTektonPipelineClientURL, err = cdtektonpipelinev2.GetServiceURLForRegion("private." + c.Region)
        if err != nil && c.Visibility == "public-and-private" {
            cdTektonPipelineClientURL, err = cdtektonpipelinev2.GetServiceURLForRegion(c.Region)
        }
    } else {
        cdTektonPipelineClientURL, err = cdtektonpipelinev2.GetServiceURLForRegion(c.Region)
    }
    if err != nil {
        cdTektonPipelineClientURL = cdtektonpipelinev2.DefaultServiceURL
    }
    if fileMap != nil && c.Visibility != "public-and-private" {
		cdTektonPipelineClientURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_CR_API_ENDPOINT", c.Region, cdTektonPipelineClientURL)
	}
    cdTektonPipelineClientOptions := &cdtektonpipelinev2.CdTektonPipelineV2Options{
        Authenticator: authenticator,
        URL: EnvFallBack([]string{"IBMCLOUD_TEKTON_PIPELINE_ENDPOINT"}, cdTektonPipelineClientURL),
    }

    // Construct the service client.
    session.cdTektonPipelineClient, err = cdtektonpipelinev2.NewCdTektonPipelineV2(cdTektonPipelineClientOptions)
    if err == nil {
        // Enable retries for API calls
        session.cdTektonPipelineClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
        // Add custom header for analytics
        session.cdTektonPipelineClient.SetDefaultHeaders(gohttp.Header{
            "X-Original-User-Agent": { fmt.Sprintf("terraform-provider-ibm/%s", version.Version) },
        })
    } else {
        session.cdTektonPipelineClientErr = fmt.Errorf("Error occurred while configuring CD Tekton Pipeline service: %q", err)
    }
```

### Changes to website/allowed-subcategories.txt  

Insert the following line into the website/allowed-subcategories.txt file (in alphabetic order):

```
CD Tekton Pipeline
``` 
