---
layout: "ibm"
page_title: "IBM : ibm_rhaii_inference_model"
description: |-
  Get information about an inference model of a Red Hat AI Inference project.
subcategory: "Red Hat AI Inference"
---

# ibm_rhaii_inference_model

Retrieve the details of one inference model of a Red Hat AI Inference (RHAII) project, such as its type, status, configuration, pricing measures and model card.

This data source calls `GET /v1/projects/{project_id}/inference/models/{model}`. For more information, see [Get an inference model](https://cloud.ibm.com/docs/apis/inference).

## Example Usage

```hcl
data "ibm_rhaii_project" "project" {
  name = "my-rhaii-project"
}

data "ibm_rhaii_inference_model" "llama" {
  project_id = data.ibm_rhaii_project.project.project_id
  model      = "llama-3-3-70b-instruct"
}

output "llama_context_window" {
  value = jsondecode(data.ibm_rhaii_inference_model.llama.config_json).max_position_embeddings
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `project_id` - (Required, String) The ID of the Red Hat AI Inference project.
* `model` - (Required, String) The ID of the model, for example `llama-3-3-70b-instruct`.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9-]+$/`.
* `location` - (Optional, String) The region of the project, for example `us-east`. When not set, the endpoint from the provider configuration is used.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - (String) The ID of the data source, in `<project_id>/<model>` format.
* `identifier` - (String) The unique identifier of the model.
* `provider_id` - (String) The ID of the provider that owns the model.
* `provider_resource_id` - (String) The ID of the model in the provider.
* `type` - (String) The resource type, `model`.
* `model_type` - (String) The model type, for example `llm`, `embedding` or `rerank`.
* `name` - (String) The name of the model.
* `display_name` - (String) The display name of the model.
* `state` - (String) The state of the model, for example `model_loaded`.
* `status` - (String) The status of the model, for example `loaded`.
* `total_params` - (Integer) The total number of parameters of the model.
* `config_json` - (String) The model configuration (architecture details and hyperparameters) as a JSON string. Use `jsondecode()` to read it.
* `model_card` - (String) The model card in markdown format.
* `source_json` - (String) The source of a custom model (`huggingface`, `ibmcos`, `oci` or `s3`) as a JSON string. Empty for IBM provided models.
* `pricing` - (List) The pricing measures of the model.
  Nested schema for **pricing**:
  * `input_measure` - (String) The billing measure for input tokens.
  * `output_measure` - (String) The billing measure for output tokens.
* `validation` - (List) The validation information of a custom model.
  Nested schema for **validation**:
  * `status` - (String) The validation status.
  * `error` - (String) The validation error message, if any.
  * `completed_at` - (String) The date when the validation completed.
