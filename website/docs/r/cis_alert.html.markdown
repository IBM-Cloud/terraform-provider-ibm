---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_alert"
description: |-
  Provides a IBM Cloud CIS Alert Policy.
---

# ibm_cis_alert

Provides a IBM CIS Alert Policy. This resource is associated with an IBM Cloud Internet Services (CIS) instance and a CIS Domain resource. It allows to create, update, delete Alert Policy of a CIS instance. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis?topic=cis-about-ibm-cloud-internet-services-cis).

## Example usage

```terraform

resource "ibm_cis_webhook" "test" {
    cis_id  = data.ibm_cis.cis.id
    name = "My Slack Alert Webhook",
    url ="https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd",
    secret = "ff1d9b80-b51d-4a06-bf67-6752fae1eb74"
}
resource "ibm_cis_alert" "test" {
  depends_on  = [ibm_cis_webhook.test]
  cis_id  = data.ibm_cis.cis.id
  name        = "test-alert-police"
  description = "alert policy description"
  enabled     = true
  alert_type = "g6_pool_toggle_alert"
  mechanisms {
    email    = ["mynotifications@email.com"]
    webhooks = [ibm_cis_webhook.test.webhook_id]
  }
  filters = <<FILTERS
  {
			"enabled": [
				"true"
			],
			"pool_id": [
				"9984f902f29adfc9bb8a5e42b7b5c592"
			]
	}
  FILTERS

	conditions = <<CONDITIONS
  {
  "and": [
    {
      "or": [
        {
          "==": [
            {
              "var": "pool_id"
            },
            "9984f902f29adfc9bb8a5e42b7b5c592"
          ]
        }
      ]
    },
    {
      "or": [
        {
          "==": [
            {
              "var": "enabled"
            },
            "true"
          ]
        }
      ]
    }
  ]
}
CONDITIONS
} 
```

## Argument reference
Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `name` - (Required, String) Name of the Alert Policy.
- `description` - (Optional, String) Description of the Alert Policy.
- `enabled` - (Required, Boolean) Alert Policy status enabled/disbaled.
- `alert_type` - (Required, String) Condition for the alert.
- `filters` - (Required, String) Must provided in JSON format. filter is the list of all enablement statuses and pool IDs for the pool toggle alert. Empty filters depending for the alert type. HTTP DDOS Attack Alerter does not require any filters. The Load Balancing Pool Enablement Alerter requires a list of IDs for the pools and their corresponding alert trigger (set whether alerts are recieved on disablement, enablement, or both). The basic WAF Alerter requires a list of zones to be monitored. The Advanced Security Alerter requires a list of zones to be monitored as well as a list of services to monitor.(https://cloud.ibm.com/docs/cis?topic=cis-configuring-notifications&interface=api)
- `conditions` - (Required, String) The conditions in JSON format. Required field when updating the Alert policy. Conditions depending on the alert type. HTTP DDOS Attack Alerter does not have any conditions. The Load Balancing Pool Enablement Alerter takes conditions that describe for all pools whether the pool is being enabled, disabled, or both. This field is not required when creating a new alert.(https://cloud.ibm.com/docs/cis?topic=cis-configuring-notifications&interface=api)
- `mechanisms` - (Required, List) Delivery mechanisms for the alert, can include an email, a webhook, or both.
	
  Nested scheme for `mechanisms`:
	- `email` - (Set, String) Provide at least one of email id.
	- `webhooks` - (Set, String) Provide at least one of Webhook id.

## Attributes Reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of alert resource. It is a combination of <`alert-id`>:<`crn`> attributes concatenated with ":".
- `alert_id` - (String) Unique identifier for the each Alert.


## Import

The `ibm_cis_alert` resource can be imported using the `id`. The ID is formed from the `Alert ID`and the `CRN` (Cloud Resource Name) concatentated usinga `:` character.

The CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Alert ID** is a 32 digit character string of the form: `52bfa670237f49ecb68473033c569649`. 

**Syntax**

```
$ terraform import ibm_cis_alert.myorg <alert_id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_alert.myorg
crn:v1:bluemix:public:internet-svcs-ci:global:a/01652b251c3ae2787110a995d8db0135:9054ad06-3485-421a-9300-fe3fb4b79e1d::
```