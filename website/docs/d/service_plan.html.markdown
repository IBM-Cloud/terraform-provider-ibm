---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: ibm_service_plan"
description: |-
  Get information about a service plan from IBM Cloud.
---

# `ibm_service_plan`

Retrieve information about a service plan for a Cloud Foundry service. For more information, about service plan, see [Dependencies on other services](https://cloud.ibm.com/docs/cloud-foundry-public?topic=cloud-foundry-public-dependencies).


## Example usage
The following example retrieves information about the `Lite` service plan for the `CloudantNOSQLDB` service. 

```
data "ibm_service_plan" "service_plan" {
  service = "cloudantNoSQLDB"
  plan    = "Lite"
}
```

## Argument reference
Review the input parameters that you can specify for your data source. 


- `plan` - (Required, String)  The name of the plan type supported by the service. You can retrieve the plan type by running the `ibmcloud service offerings` command in the IBM Cloud command.
- `service` - (String)  Required-The name of the service offering. You can retrieve the name of the service by running the `ibmcloud service offerings` command in the IBM Cloud command.


## Attribute reference
Review the output parameters that you can access after you retrieved your data source. 

- `id` - (String) The unique identifier of the service plan.



