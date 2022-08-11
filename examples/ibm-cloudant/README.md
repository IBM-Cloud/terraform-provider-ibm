# Examples for `ibm_cloudant`

Examples in the subfolders illustrate how to use the `ibm_cloudant`.

These types of resources are supported:

* ibm_cloudant

Each example creates an IBM Cloudant resource instance called `cloudant`.

## Examples

Examples can be found in the subfolders along with the instructions how to run them.

- [Lite plan](lite-plan)
- [Lite plan with legacy credentials](lite-plan-legacy)
- [Lite plan with IAM credentials](lite-plan-iam)
- [Standard plan with custom capacity](standard-plan)
- [Standard plan with a new database](standard-plan-with-database)
- [Standard plan with data event tracking](standard-plan-with-data-events)
- [Standard plan on dedicated hardware](standard-plan-on-dedicated-hw)

## Assumptions

## Notes

1. With `Lite` plan `capacity` can be set no more than 1 throughput blocks.
1. `parameters` can overwrite the previously set arguments named the same way.
1. With [`Standard` plan on dedicated hardware](standard-plan-on-dedicated-hw) the hardware must be ordered separately and provisioning should be completed before using Terraform on it

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |

## Inputs

| Name | Description | Type | Required | Default |
|------|-------------|------|----------|---------|
| ibmcloud_api_key | IBM Cloud API key. | `string` | true | -
| name | A name for the resource instance. | `string` | true | -
| location | Target location or environment to create the resource instance. (Forces new resource.) | `string` | true | -
| plan | The plan type of the service. | string | true | -
| capacity | A number of blocks of throughput units. For more details please read about [`blocks`](https://cloud.ibm.com/apidocs/cloudant#putcapacitythroughputconfiguration) parameter. | `number` | false | `1`
| id | The unique identifier of the new Cloudant resource. | `string` | false | -
| cors_config | Configuration for CORS. The minimum length is `1` item. Can conflict with `enable_cors` if it is set to false. In this case the CORS customization is not allowed. | block `list` (see nested arguments below) | false | -
| cors_config.origins | An array of strings that contain allowed origin domains. You have to specify the full URL including the protocol. It is recommended that only the HTTPS protocol is used. Subdomains count as separate domains, so you have to specify all subdomains used. | `list(string)` | false | -
| cors_config.allow_credentials | Boolean value to allow authentication credentials. If set to true, browser requests must be done by using withCredentials = true. | `bool` | false | `false`
| enable_cors | Boolean value to turn CORS on and off. | `bool` | false | `true`
| environment_crn | CRN of the IBM Cloudant Dedicated Hardware plan instance. | `string` | false | -
| legacy_credentials | Use both legacy credentials and IAM for authentication. | `bool` | false | `false`
| include_data_events | Include `data` event types in events sent to IBM Cloud Activity Tracker with LogDNA for the IBM Cloudant instance. By default only emitted events are of `management` type. | `bool` | false | `false`
| parameters | Arbitrary parameters to pass. Must be a JSON object. | `map(string)` | false | -
| resource_group_id | The resource group id. (Forces new resource.) | `string` | false | -
| service_endpoints | Types of the service endpoints. Possible values are 'public', 'private', 'public-and-private'. | `string` | false | -
| tags | Tags associated with the instance. | `set(string)` | false | -
| db_name | A name of database to create. | `string` | true | -
| timeouts.create<br>timeouts.update<br>timeouts.delete | The operation of the IBM Cloudant instance is considered failed if no response received for the given timeout. | `string` | false | -

## Outputs

| Name | Description |
|------|-------------|
| ibm_cloudant | `ibm_cloudant` terraform resource instance |
