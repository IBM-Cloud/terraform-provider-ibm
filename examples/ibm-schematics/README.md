# Example for IBM Schematics Datasource

This example illustrates how to retrieve information about a Schematics workspace, state information for a Schematics workspace and information about the Terraform state file for a Schematics workspace.

These types of datasources are supported:

* [ ibm_schematics_workspace ](https://cloud.ibm.com/docs/terraform?topic=terraform-schematics-data-sources#schematics-workspace)
* [ ibm_schematics_output](https://cloud.ibm.com/docs/terraform?topic=terraform-schematics-data-sources#schematics-output)
* [ ibm_schematics_state ](https://cloud.ibm.com/docs/terraform?topic=terraform-schematics-data-sources#schematics-state)

## Terraform versions

Terraform 0.12. Pin module version to `~> v1.4.0`. Branch - `master`.

Terraform 0.11. Pin module version to `~> v0.25.0`. Branch - `terraform_v0.11.x`.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## ibm_schematics_workspace datasource

`Retrieve information about a Schematics workspace`:
```hcl

data "ibm_schematics_workspace" "test" {
  workspace_id = var.workspaceID
}

```
## ibm_schematics_output datasource

`Retrieve state information for a Schematics workspace.`:
```hcl

data "ibm_schematics_workspace" "vpc" {
  workspace_id = var.workspaceID
}

data "ibm_schematics_output" "test" {
  workspace_id = var.workspaceID
  template_id= data.ibm_schematics_workspace.vpc.template_id.0
}

```

## ibm_schematics_state datasource

`Retrieve information about the Terraform state file for a Schematics workspace.`:
```hcl

data "ibm_schematics_state" "test" {
  workspace_id = var.workspaceID
  template_id= var.templateID
}

```

## Examples

* [ Schematics data sources ](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-schematics)


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
| workspaceID | ID of the Schematics workspace.| `string` | yes |
| templateID | ID of the template that the workspace uses.| `string` | yes |

## Outputs

| Name    | Description                            |
|---------|----------------------------------------|
| status  | The status of the workspace.           |
| tags    | tags that were added to the workspace. |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
