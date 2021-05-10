---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: ibm_space"
description: |-
  Get information about an IBM Cloud space.
---

# `ibm_space`

Retrieve information about an existing Cloud Foundry space. For more information, about Cloud Foundry space, see [Getting started with IBM Cloud Foundry Enterprise Environment](https://cloud.ibm.com/docs/cloud-foundry?topic=cloud-foundry-getting-started).


## Example usage
The following example retrieves information about the `prod` Cloud Foundry space.

```
data "ibm_space" "spaceData" {
  space = "prod"
  org   = "myorg.com"
}
```

The following example shows how you can use the data source to reference the space ID in the `ibm_service_instance` resource.

```
resource "ibm_service_instance" "service_instance" {
  name       = "test"
  space_guid = data.ibm_space.spaceData.id
  service    = "speech_to_text"
  plan       = "lite"
  tags       = ["cluster-service", "cluster-bind"]
}
```

## Argument reference
Review the input parameters that you can specify for your data source. 

- `name` - (Optional, String)  The name of your space.
- `org` - (Deprecated, String) The name of your Cloud Foundry organization that the space belongs to. You can retrieve the value by running the `ibmcloud iam orgs` command in the IBM Cloud CLI.
- `space` - (Deprecated, String)  The name of your Cloud Foundry space. You can retrieve the value by running the `ibmcloud iam spaces` command in the IBM Cloud CLI.


## Attribute reference
Review the output parameters that you can access after you retrieved your data source. 

- `auditors` - (String) The email addresses (associated with IBMID) of the users who have an auditor role in this space.
- `developers` - (String) The email addresses (associated with IBMID) of the users who have a developer role in this space.
- `id` - (String) The unique identifier of the space.
- `managers` - (String) The email addresses (associated with IBMID) of the users who have a manager role in this space.
