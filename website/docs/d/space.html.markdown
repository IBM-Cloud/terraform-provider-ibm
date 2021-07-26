---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: ibm_space"
description: |-
  Get information about an IBM Cloud space.
---

# ibm_space

Retrieve information about an existing Cloud Foundry space. For more information, about Cloud Foundry space, see [getting started with IBM Cloud Foundry Enterprise Environment](https://cloud.ibm.com/docs/cloud-foundry?topic=cloud-foundry-getting-started).


## Example usage
The following example retrieves information about the `prod` Cloud Foundry space.

```terraform
data "ibm_space" "spaceData" {
  space = "prod"
  org   = "myorg.com"
}
```

The following example shows how you can use the data source to reference the space ID in the `ibm_service_instance` resource.

```terraform
resource "ibm_service_instance" "service_instance" {
  name       = "test"
  space_guid = data.ibm_space.spaceData.id
  service    = "speech_to_text"
  plan       = "lite"
  tags       = ["cluster-service", "cluster-bind"]
}
```

## Argument reference
Review the argument reference that you can specify for your data source. 

- `name` - (Optional, String)  The name of your space.
- `org` - (Deprecated, String) The name of your Cloud Foundry organization that the space belongs to. You can retrieve the value by running the `ibmcloud iam orgs` command in the IBM Cloud CLI.
- `space` - (Deprecated, String)  The name of your Cloud Foundry space. You can retrieve the value by running the `ibmcloud iam spaces` command in the IBM Cloud CLI.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `auditors` - (String) The Email addresses (associated with IBMID) of the users who have an auditor role in this space.
- `developers` - (String) The Email addresses (associated with IBMID) of the users who have a developer role in this space.
- `id` - (String) The unique identifier of the space.
- `managers` - (String) The Email addresses (associated with IBMID) of the users who have a manager role in this space.
