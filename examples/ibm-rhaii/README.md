# Red Hat AI Inference project example

This example creates a Red Hat AI Inference (RHAII) project, reads it back with the data source, and lists the inference models that are available in it.

It is the Terraform equivalent of:

```sh
ibmcloud resource service-instance-create my-rhaii-project instructlab instructlab-pricing-plan us-east -g Default
```

## Usage

```sh
export TF_VAR_ibmcloud_api_key=<your_api_key>
terraform init
terraform plan
terraform apply
```

Run `terraform destroy` when you no longer need the project.

## Resources and data sources

* `ibm_rhaii_project` (resource)
* `ibm_rhaii_project` (data source)
* `ibm_rhaii_inference_models` (data source)
* `ibm_rhaii_inference_model` (data source)

## Inputs

| Name | Description | Type | Default |
|------|-------------|------|---------|
| ibmcloud_api_key | IBM Cloud API key | `string` | n/a |
| resource_group | Name of the resource group for the project | `string` | `Default` |
| project_name | Name of the project | `string` | `my-rhaii-project` |
| plan_name | Name of the pricing plan of the project | `string` | `instructlab-pricing-plan` |
| location | Region of the project | `string` | `us-east` |
| tags | User tags for the project | `list(string)` | `[]` |
| access_tags | Access management tags for the project. The tags must already exist. | `list(string)` | `[]` |
| model | ID of the inference model to read | `string` | `llama-3-3-70b-instruct` |

## Outputs

| Name | Description |
|------|-------------|
| project_id | ID of the project. Use it as `project_id` in the Red Hat AI Inference API. |
| endpoint | Public base URL of the Red Hat AI Inference API for this project |
| private_endpoint | Private base URL of the Red Hat AI Inference API for this project |
| plan_id | Global catalog ID of the pricing plan of the project |
| crn | CRN of the project |
| model_ids | IDs of the inference models available in the project |
| model_status | Status of the selected inference model |
