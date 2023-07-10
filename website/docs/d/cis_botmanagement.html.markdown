---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_bot_managements"
description: |-
  Get information on an IBM Cloud Internet Services Bot Management APIs.
---

# ibm_cis_bot_managements

Retrieve information about an IBM Cloud Internet Services Bot Management data sources for a zone. For more information, see [IBM Cloud Internet Services Bot Management](https://cloud.ibm.com/docs/cis?topic=cis-about-bot-mgmt).
a
## Example usage

```terraform
data "ibm_cis_bot_managements" "tests" {
    cis_id                          = data.ibm_cis.cis.id
    domain = data.ibm_cis_domain.cis_domain.domain

}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain` - (Required, String) The Domain of the CIS service instance.


## Attributes reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The ID of the CIS service instance.
- `domain` - (String) The Domain of the CIS service instance.
- `fight_mode` - (Boolean) Fight mode enable/disable
- `enable_js` - (Boolean) Use lightweight, invisible JavaScript detections to improve Bot Management. Learn more about [JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/)
- `session_score` - (Boolean) Session score enable/disable
- `auth_id_logging` - (Boolean) Auth ID Logging enable/disable
- `use_latest_model` - (Boolean) Use Latest Model enable/disable
