# IBM IAM Custom Role example

This example shows how to create an IAM custom role in account. Specify a service and Action to be assigned to the Custom Role .


These types of resources and datsources are supported:

* [ Import Roles ](https://cloud.ibm.com/docs/terraform?topic=terraform-iam-data-sources)
* [ Import Actions for Roles ](https://cloud.ibm.com/docs/terraform?topic=terraform-iam-data-sources)
* [ Create Custom Roles ](https://cloud.ibm.com/docs/terraform?topic=terraform-iam-resources#iam-custom-role)


## Terraform versions

Terraform 0.12. Pin module version to `~> v1.7.0`. Branch - `master`.

Terraform 0.11. Pin module version to `~> v0.28.0`. Branch - `terraform_v0.11.x`.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## Custom Role Resources

`Create custom role resource by Importing Action IDs`:
```hcl

data "ibm_iam_role_actions" "test" {
  service = "%s"
}

resource "ibm_iam_custom_role" "customrole" {
    name         = var.name
    display_name = var.displayname
    description  = "Custom Role for test scenario2"
    service = "kms"
    actions      = [data.ibm_iam_role_actions.test.manager.18]
}

```
`Create Custom Roles`:
```hcl

resource "ibm_iam_custom_role" "customrole" {
  name         = var.name
  display_name = var.displayname
  description  = "This is a custom role"
  service = "kms"
  actions      = [var.action]
}

```
##  IAM Roles Data Source
`List all roles and import it to create an access group policy:`

```hcl

data "ibm_iam_roles" "test" {
	service = "%s"
  }

resource "ibm_iam_access_group_policy" "policy" {
	access_group_id = var.accgrpid
	roles           = [data.ibm_iam_roles.test.roles.10.name,"Viewer"]
	tags            = ["tag1"]
	resources {
	  service = "kms"
	}
}
```

## Assumptions

1. The data source ibm_iam_roles will list all kind of roles along with custom roles. Ommit the service to list only platform roles
2. Custom roles can be used in access group policy, user policy, service policy and user invite
3. [ API Documentation for Custom Roles ](https://cloud.ibm.com/apidocs/iam-policy-management#get-roles-by-filters)


## Examples

* [ Create custom role and import in policy example ](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-iam-custom-role)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | n/a |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| displayname | The display name of the custom role| `string` | yes |
| name | The name of the custom role | `string` | yes |
| description | The description of the custom role | `string` | yes |
| actions | The action ID to be associated with the custom role | `string` | yes |
| servicename | The service name for which custom role is being created | `string` | yes |
| agname | The name of the access group to be created | `string` | yes |


## Outputs

| Name | Description |
|------|-------------|
| roles | The list of all iam roles associated with the service and their properties |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
