---
layout: "ibm"
page_title: "IBM : ibm_scc_addon_activity_insights_cos_details"
description: |-
  Manages scc_addon_activity_insights_cos_details.
subcategory: "Security and Compliance Center"
---

# ibm_scc_addon_activity_insights_cos_details

Provides a resource for scc_addon_activity_insights_cos_details. This allows scc_addon_activity_insights_cos_details to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_scc_addon_activity_insights_cos_details" "scc_addon_activity_insights_cos_details" {
  cos_details {
    cos_instance = "cos_instance_1"
    bucket_name = "us-south-collection"
    description = "Holds SCC Insights"
    cos_bucket_url = "https://webhook.site/96cf2ebe-7caa-409b-89be-edf89a25f5db"
  }
  region_id = "us"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `account_id` - (Optional, Forces new resource, String) Account ID is optional, if not provided value will be inferred from the token retrieved from the IBM Cloud API key.
* `cos_details` - (Required, List) 
Nested scheme for **cos_details**:
	* `bucket_name` - (Required, String)
	* `cos_bucket_url` - (Required, String) cos bucket url.
	* `cos_instance` - (Required, String)
	* `description` - (Required, String)
* `region_id` - (Required, String) Region for example - us-south, eu-gb.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the scc_addon_activity_insights_cos_details.

## Import

You can import the `ibm_scc_addon_activity_insights_cos_details` resource by using `region_id`.
The `region_id` property can be formed from 

```
<account_id>/<bucket_id>
```
* `account_id` - A string. AccountID from the resource has to be imported.
* `bucket_id` - A string. ID of the bucket whose COS details has to be imported.

# Syntax
```
$ terraform import ibm_scc_addon_activity_insights_cos_details.scc_addon_activity_insights_cos_details 
```
