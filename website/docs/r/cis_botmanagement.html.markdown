---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_bot_management"
description: |-
  Get information on an IBM Cloud Internet Services Bot Management APIs.
---

# ibm_cis_bot_management

 Provides IBM Cloud Internet Services Bot Management resource. The resource allows to update Bot Management settings of a domain of an IBM Cloud Internet Services CIS instance. For more information, see [IBM Cloud Internet Services Bot Management](https://cloud.ibm.com/docs/cis?topic=cis-about-bot-mgmt).

## Example usage
```terraform
# Change Bot Management setting of CIS instance

resource "ibm_cis_bot_management" "test" {
    cis_id                          = data.ibm_cis.cis.id
    domain = data.ibm_cis_domain.cis_domain.domain
    fight_mode				= false
    session_score			= false
    enable_js				= false
    auth_id_logging			= false
    use_latest_model 		= false

}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain` - (Required, String) The Domain of the CIS service instance.
- `fight_mode` - (Required, Boolean) Fight mode enable/disable
- `enable_js` - (Required, Boolean) Use lightweight, invisible JavaScript detections to improve Bot Management. Learn more about [JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/)
- `session_score` - (Required, Boolean) Session score enable/disable
- `auth_id_logging` - (Required, Boolean) Auth ID Logging enable/disable
- `use_latest_model` - (Required, Boolean) Use Latest Model enable/disable



