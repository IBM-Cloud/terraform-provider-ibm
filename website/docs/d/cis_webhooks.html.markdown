---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_webhooks"
description: |-
  Get information on an IBM Cloud Internet Services webhooks.
---

# ibm_cis_webhooks

Retrieve information about an IBM Cloud Internet Services webhooks data sources. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis?topic=cis-about-ibm-cloud-internet-services-cis).

## Example usage

```terraform
data "ibm_cis_webhooks" "test" {
  cis_id    = ibm_cis.instance.id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS service instance.


## Attributes reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The Webhook ID. It is a combination of <`webhook_id`>,<`cis_id`> attributes concatenated with ":"
- `cis_webhooks` - (List)
   - `name` - (String) The name of webhook.
   - `url` - (Boolean). Whether this webhook is currently disabled.
   - `type` - (String) The information about this webhook to help identify the purpose of it.
   - `webhook_id` - (String) The Webhook ID.

