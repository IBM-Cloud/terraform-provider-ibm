---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_integration'
description: |-
  Manages Event Notifications Integrations.
---

# ibm_en_integration

update integration using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_integration" "en_kms_integration" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  integration_id = "xyz-rdserr-froyth-lowhbw"
  type = "kms"
  metadata {
    endpoint = "https://us-south.kms.cloud.ibm.com"
    crn = "crn:v1:bluemix:public:kms:us-south:a/tyyeeuuii2637390003hehhhhi:fgsyysbnjiios::"
    root_key_id = "gyyebvhy-34673783-nshuwubw"
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `integration_id` - (Required, String) Unique identifier for Integration created with .

- `type` - (Required, String) The integration type kms/hs-crypto.

- `metadata` - (Required, List)

  Nested scheme for **params**:

  - `endpoint` - (Required, String) key protect/hyper protect service endpoint.
  - `crn` - (Required, String) crn of key protect/ hyper protect instance.
  - `root_key_id` - (Required, String) Root key id.


## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `en_integration`.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_integration` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `integration_id` in the following format:

```
<instance_guid>/<integration_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `integration_id`: A string. Unique identifier for Destination.

**Example**

```
$ terraform import ibm_en_integration.en_integration <instance_guid>/<integration_id>
```
