---

subcategory: "Schematics"
layout: "ibm"
page_title: "IBM: ibm_schematics_state"
sidebar_current: "docs-ibm-datasource-schematics-state"
description: |-
  Get information about Schematics state
---

# ibm_schematics_state
Retrieve information about the  Terraform state file for a Schematics workspace. For more information, about Schematics workspace state, see [workspace state](https://cloud.ibm.com/docs/schematics?topic=schematics-workspace-setup#wks-state).

## Example usage

```terraform
data "ibm_schematics_state" "schematics_state" {
	workspace_id = "workspace_id"
	template_id = "template_id"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

* `location` - (Optional,String) Location supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
  * Constraints: Allowable values are: us-south, us-east, eu-gb, eu-de
- `template_id` - (Required, String) The ID of the Terraform template for which you want to retrieve the Terraform statefile. When you create a workspace, the Terraform template that your workspace points to is assigned a unique ID. To find this ID, use the `GET /v1/workspaces` API and review the `template_data.id` value.
- `workspace_id` - (Required, String) The workspace ID for which you want to retrieve the Terraform statefile. To find the workspace ID, use the `GET /v1/workspaces` API.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `id` - (String) The unique ID of the Schematics state.
- `version`-  (String) The Schematics version.
- `terraform_version`-  (String) The Terraform version.
- `serial` - (String) The state store serial number details.
- `lineage`- (Integer) The state store lineage number details.
- `modules`-  (String) The state store module details.
