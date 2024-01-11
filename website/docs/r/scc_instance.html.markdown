---
layout: "ibm"
page_title: "IBM : ibm_scc_instance"
description: |-
  Manages scc_instance.
subcategory: "Security and Compliance Center"
---

# ibm_scc_instance

Create, update, and delete scc_instance with this resource.

~> NOTE: This document details how to use the resource `ibm_resource_instance` targeting the service `Security and Compliance Center`. For more information about the Terraform resource `ibm_resource_instance`, click [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/resource_instance)

## Example Usage

```hcl
data "ibm_resource_group" "group" {
  name = "test"
}

resource "ibm_resource_instance" "scc_instance" {
  name              = "test"
  service           = "compliance"
  plan              = "security-compliance-center-standard-plan" # also support security-compliance-center-trial-plan
  location          = "us-south"
  resource_group_id = data.ibm_resource_group.group.id
  tags              = ["tag1", "tag2"]
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `location` - (Required, Forces new resource, String) Target location or environment to create the resource instance.
- `plan` - (Required, String) The name of the plan type supported by service. You can retrieve the value by running the `ibmcloud catalog service <servicename>` command.
- `name` - (Required, String) A descriptive name used to identify the resource instance.
- `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group where you want to create the service. You can retrieve the value from data source `ibm_resource_group`. If not provided creates the service in `default` resource group.
- `tags` (Optional, Array of Strings) Tags associated with the instance.
- `service` - (Required, Forces new resource, String) The name of the service offering.

