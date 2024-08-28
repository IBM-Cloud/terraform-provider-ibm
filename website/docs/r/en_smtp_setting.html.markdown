---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_smtp_setting'
description: |-
  Manages Event Notifications SMTP Configuration Setting.
---

# ibm_en_smtp_setting

update integration using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_smtp_setting" "en_smtp_setting" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  smtp_config_id = ibm_resource_instance.en_smtp_configuration_instance.en_smtp_configuration_id
  settings {
    subnets = ["44.255.224.210/20","100.113.203.15/26"]
  }
}
```

Note: The support for legacy allowlisting has been deprecated. The support has been enabled via Context-based-restrictions. For detailed information, please refer here: https://cloud.ibm.com/docs/event-notifications?topic=event-notifications-en-smtp-configurations#en-smtp-configurations-cbr. You can get the existing IP's details using data source `ibm_en_smtp_allowed_ips`

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `smtp_config_id` - (Required, String) - SMTP ID.
  * Constraints: The maximum length is `100` characters. The minimum length is `32` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}/`. .

- `settings` - (Required, List)

  Nested scheme for **settings**:

  - `subnets` - (Required, list) The SMTP allowed Ips.


## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `en_smtp_setting`.

## Import

You can import the `ibm_smtp_setting` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `smtp_config_id` in the following format:

```
<instance_guid>/<smtp_config_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `smtp_config_id`: A string. Unique identifier for SMTP Config

**Example**

```
$ terraform import ibm_en_integration.en_integration <instance_guid>/<smtp_config_id>
```
