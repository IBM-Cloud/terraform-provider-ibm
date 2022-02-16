---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_webhook"
description: |-
  Provides a IBM Cloud CIS Webhook.
---

# ibm_cis_webhook

Provides a IBM CIS Webhook. This resource is associated with an IBM Cloud Internet Services (CIS) instance and a CIS Domain resource. It allows to create, update, delete webhook of a CIS instance. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis?topic=cis-about-ibm-cloud-internet-services-cis).

## Example usage

```terraform

resource "ibm_cis_webhook" "test" {
    cis_id  = data.ibm_cis.cis.id
    name = "My Slack Alert Webhook",
    url ="https://hooks.slack.com/services/Ds3fdBFbV/456464Gdd",
    secret = "ff1d9b80-b51d-4a06-bf67-6752fae1eb74"
}

```

## Argument reference
Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `name` - (Required, String) Name of the Webhook.
- `secret` - (Optional, String) This is Sensitive and optional secret or API key needed to use the webhook.
- `url` - (Required, string) Webhook Url.

## Attributes Reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of webhook resource. It is a combination of <`webhook-id`>:<`crn`> attributes concatenated with ":".
- `webhook_id` - (String) Unique identifier for the Webhook.

## Import

The `ibm_cis_webhook` resource can be imported using the `id`. The ID is formed from the `Webhook ID`and the `CRN` (Cloud Resource Name) concatentated usinga `:` character.

The CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Webhook ID** is a 32 digit character string of the form: `d72c91492cc24d8286fb713d406abe91`. 

**Syntax**

```
$ terraform import ibm_cis_webhook.myorg <webhook_id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_webhook.myorg
crn:v1:bluemix:public:internet-svcs-ci:global:a/01652b251c3ae2787110a995d8db0135:9054ad06-3485-421a-9300-fe3fb4b79e1d::
```