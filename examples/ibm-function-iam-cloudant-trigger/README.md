# IBM Cloud Function example

This example shows how IBM Cloud Functions IAM namespace created and action is triggered when documents in Cloudant NoSQL databases are changed or added.

In this example a Cloudant NoSQL service instance is created. We deploy a python app which creates a database 'databasedemo' in Cloudant  NOSQL. We bind a cloudant package using IBM Cloud Function package and create an action, trigger and rule.

When you change documents or add documents in your Cloudant database you can see logs in IBM Cloud Functions dashboard.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

These types of resources and datasources are supported:

* [ Namespace ](https://cloud.ibm.com/docs/terraform?topic=terraform-function-data-sources#fn-namespace_ds)
* [ Package ](https://cloud.ibm.com/docs/terraform?topic=terraform-function-data-sources#fn-package)
* [ Action ](https://cloud.ibm.com/docs/terraform?topic=terraform-function-data-sources#fn-action)
* [ Trigger ](https://cloud.ibm.com/docs/terraform?topic=terraform-function-data-sources#fn-trigger)
* [ Rule ](https://cloud.ibm.com/docs/terraform?topic=terraform-function-data-sources#fn-rule)

## Terraform versions

Terraform 0.12. Pin module version to `~> v1.10.0`. Branch - `master`.

## Cloud Function Resources

Namespace Resource:
```hcl
data "ibm_resource_group" "resource-group" {
  name = var.resource_group
}

resource "ibm_function_namespace" "namespace" {
  name                = var.namespace
  resource_group_id   = data.ibm_resource_group.resource-group.id
}
```
Package Resource:
```hcl
resource "ibm_function_package" "package" {
  name      = var.packageName
  namespace = var.namespace

  user_defined_parameters = <<EOF
        [
    		{
        		"key":"name",
        		"value":"terraform"
    		},
    		{
        		"key":"place",
        		"value":"India"
    		}
	]
  EOF

}
```
Action Resource:
```hcl
resource "ibm_function_action" "action" {
  name      = var.actionname
  namespace = var.namespace  

  exec {
    kind = "nodejs:10"
    code = file("hello.js")
  }
}
```
Trigger Resource:
```hcl
resource "ibm_function_trigger" "trigger" {
  name      = var.triggername
  namespace = var.namespace  
	
  user_defined_parameters = <<EOF
         [
        	  {
                "key":"place",
                "value":"India"
            }
         ]

  EOF
  user_defined_annotations = <<EOF
         [
              {
                  "key":"Description",
                  "value":"Sample code to display hello"
              }
         ]

  EOF
}
```
Rule Resource:
```hcl
resource "ibm_function_rule" "rule" {
  name         = var.rulename
  namespace    = var.namespace  	
  trigger_name = ibm_function_trigger.trigger.name
  action_name  = ibm_function_action.action.name
}
```

## Cloud Function Data sources

Namespace Data Source:
```hcl
data "ibm_function_namespace" "namespace" {
    name      = var.namespace
}
```
Package Data Source:
```hcl
data "ibm_function_package" "package" {
    name      = var.packagename
    namespace = var.namespace
}
```
Action Data Source:
```hcl
data "ibm_function_action" "action" {
    name      = var.actionname
    namespace = var.namespace
}
```
Trigger Data Source:
```hcl
data "ibm_function_trigger" "trigger" {
    name      = var.triggername
    namespace = var.namespace
}
```
Rule Data Source:
```hcl
data "ibm_function_rule" "rule" {
    name      = var.rulename
    namespace = var.namespace
}
```

## Notes

1. 'ibm_function_namespace' resource supports only IAM based namespace creation.

## Examples

* [Cloud Functions resources](https://github.com/Mavrickk3/terraform-provider-ibm/tree/master/examples/ibm-function-iam-cloudant-trigger)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

Single OpenAPI document or directory of documents.

## Providers

| Name | Version |
|------|---------|
| ibm | n/a |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| namespace | The cloud function namespace.| `string` | yes |
| resource_group| The resource group name.| `string` | yes |
| packageName | The name of the Package Instance. | `string` | yes |
| actionName | The name of the Action Instance. | `string` | yes |
| boundPackageName | The name of the boundpackage instance. | `string` | yes |
| triggerName | The name of the trigger instance. | `string` | yes |
| ruleName |The name of the rule instance.| `string` | yes |
| dbname | The name of the database. | `string` | yes |
| space | The Cloud Foundry Enterprise Environment are scoped within specific spaces. | `string` | yes. |
| org | A space exists within a specific organization. | `string` | yes |
| service | The name of the cf service. | `string` | yes |
| plan | The type of the account plan. Supported values are [`Lite`] and [`Standard`]. | `string` | yes |
| service_instance_name | The name of the cf service instance. | `string` | yes |
| service_key_name | The name of the service key. | `string` | yes |
| app_version | The cf application version. | `string` | yes |
| git_repo | The git repo location of application. | `string` | yes |
| dir_to_clone | Directory to clone the git repo. | `string` | yes |
| app_zip | The zip file contains the application to install. | `string` | yes |
| app_name | The name of the application. | `string` | yes |
| app_command | The command to install application. | `string` | yes |
| buildpack | The build pack name. | `string` | no |

## Outputs

| Name | Description |
|------|-------------|
| namespaceID | Namespace ID |
| packageID | Package ID |
| actionID | Action ID |
| ruleID | Rule ID |
| triggerID | Trigger ID |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
