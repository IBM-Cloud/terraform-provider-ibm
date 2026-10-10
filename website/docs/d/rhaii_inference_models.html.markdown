---
layout: "ibm"
page_title: "IBM : ibm_rhaii_inference_models"
description: |-
  Get the list of inference models of a Red Hat AI Inference project.
subcategory: "Red Hat AI Inference"
---

# ibm_rhaii_inference_models

Retrieve the list of inference models that are available in a Red Hat AI Inference (RHAII) project. Use the model `id` as the `model` value in chat completion requests.

This data source calls `GET /v1/projects/{project_id}/inference/models`. For more information, see [List inference models](https://cloud.ibm.com/docs/apis/inference).

## Example Usage

```hcl
resource "ibm_rhaii_project" "project" {
  name = "my-rhaii-project"
}

data "ibm_rhaii_inference_models" "models" {
  project_id = ibm_rhaii_project.project.project_id
  location   = ibm_rhaii_project.project.location
}

output "model_ids" {
  value = data.ibm_rhaii_inference_models.models.models[*].id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `project_id` - (Required, String) The ID of the Red Hat AI Inference project. Use the `project_id` attribute of the `ibm_rhaii_project` resource or data source.
* `location` - (Optional, String) The region of the project, for example `us-east`. When not set, the region of the provider is used, or `us-east` when the service is not available in that region.

## Endpoints

The data source calls the Red Hat AI Inference API on the endpoint that matches the `visibility` argument of the provider:

* `public` (default): `https://<location>.rhai.ibm.com/v1`
* `private`: `https://private.<location>.rhai.ibm.com/v1`. Terraform must run on the IBM Cloud private network, for example in a VPC.
* `public-and-private`: the private endpoint when the region has one, otherwise the public endpoint.

To use another endpoint, set the `IBMCLOUD_RHAII_API_ENDPOINT` environment variable, or the `IBMCLOUD_RHAII_API_ENDPOINT` key in the [endpoints file](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/guides/custom-service-endpoints).

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - (String) The ID of the project.
* `models` - (List) The inference models that are available in the project.
  Nested schema for **models**:
  * `id` - (String) The ID of the model. Use this value as `model` in chat completion requests.
  * `object` - (String) The object type, `model`.
  * `created` - (Integer) The Unix timestamp in seconds when the model was created.
  * `owned_by` - (String) The owner of the model.
  * `name` - (String) The name of the model.
  * `display_name` - (String) The display name of the model.
  * `state` - (String) The state of the model, for example `model_loaded`.
  * `status` - (String) The status of the model, for example `loaded`.
  * `total_params` - (Float) The total number of parameters of the model, for example `70553706496`.
  * `config_json` - (String) The model configuration (architecture details and hyperparameters) as a JSON string. Use `jsondecode()` to read it.
  * `pricing` - (List) The pricing measures of the model.
    Nested schema for **pricing**:
    * `input_measure` - (String) The billing measure for input tokens.
    * `output_measure` - (String) The billing measure for output tokens.
  * `validation` - (List) The validation information of a custom model.
    Nested schema for **validation**:
    * `status` - (String) The validation status.
    * `error` - (String) The validation error message, if any.
    * `completed_at` - (String) The date when the validation completed.
