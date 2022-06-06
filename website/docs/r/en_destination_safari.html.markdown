---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_safari'
description: |-
  Manages Event Notifications Safari destination.
---

# ibm_en_destination_safari

Create, update, or delete Safari destination by using IBM Cloud™ Event Notifications.

## Example usage for Safari Destination

```terraform
resource "ibm_en_destination_safari" "safari_en_destination" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  name                         = "EN Safari Destination"
  type                         = "push_safari"
  certificate                  = "${path.module}/Certificates/safaricert.p12"
  icon_16x16                   = "${path.module}/Certificates/safariicon.png"
  icon_16x16_2x                = "${path.module}/Certificates/safariicon.png"
  icon_32x32                   = "${path.module}/Certificates/safariicon.png"
  icon_32x32_2x                = "${path.module}/Certificates/safariicon.png"
  icon_128x128                 = "${path.module}/Certificates/safariicon.png"
  icon_128x128_2x              = "${path.module}/Certificates/safariicon.png"
  icon_16x16_content_type      = "png"
  icon_16x16_2x_content_type   = "png"
  icon_32x32_content_type      = "png"
  icon_32x32_2x_content_type   = "png"
  icon_128x128_content_type    = "png"
  icon_128x128_2x_content_type = "png"
  description                  = "Safari destination in EN"
  config {
    params {
      cert_type         = "p12"
      password          = "apnscertpassword"
      url_format_string = "https://test.petstorez.com"
      website_name      = "petstore"
      website_push_id   = "petzz"
      website_url       = "https://test.petstorez.com"
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) The Destintion name.

- `description` - (Optional, String) The Destination description.

- `type` - (Required, String) push_safari.

- `certificate` - (Required, binary) Certificate file. The file type allowed is .p8 and .p12

- `icon_16x16` - (Optional, binary) icon file of dimension 16x16

- `icon_16x16_2x` - (Optional, binary) icon file of dimension 16x16x2x

- `icon_32x32` - (Optional, binary) icon file of dimension 32x32

- `icon_32x32_2x` - (Optional, binary) icon file of dimension 32x32x2x

- `icon_128x128` - (Optional, binary) icon file of dimension 128x128

- `icon_128x128_2x` - (Optional, binary) icon file of dimension 128x128x2x

- `icon_16x16_content_type` - (Optional, binary) The extension of icon image of 16x16 dimension. Required in case of passing icon file.

- `icon_16x16_2x_content_type` - (Optional, binary) The extension of icon image of 16x16x2x dimension. Required in case of passing icon file.

- `icon_32x32_content_type` - (Optional, binary) The extension of icon image of 32x32 dimension. Required in case of passing icon file.

- `icon_32x32_2x_content_type` - (Optional, binary) The extension of icon image of 32x32x2x dimension. Required in case of passing icon file.

- `icon_128x128_content_type` - (Optional, binary) The extension of icon image of 128x128 dimension. Required in case of passing icon file.

- `icon_128x128_2x_content_type` - (Optional, binary) The extension of icon image of 128x128x2x dimension. Required in case of passing icon file.

- `config` - (Required, List) Payload describing a destination configuration.

  Nested scheme for **config**:

  - `params` - (Required, List)

  Nested scheme for **params**:

  - `cert_type` - (String) The Certificate type. Value is p12.

  - `password` - (String) The password string for p12 certificate.

  - `url_format_string` - (String) Website formatted url .

  - `website_name` - (String) The name of website.

  - `website_push_id` - (String) Website push ID  .

  - `website_url` - (String) Website url.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `safari_en_destination`.
- `destination_id` - (String) The unique identifier of the created destination.
- `subscription_count` - (Integer) Number of subscriptions.
  - Constraints: The minimum value is `0`.
- `subscription_names` - (List) List of subscriptions.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_destination_safari` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `destination_id` in the following format:

```
<instance_guid>/<destination_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id`: A string. Unique identifier for Destination.

**Example**

```
$ terraform import ibm_en_destination_safari.safari_en_destination <instance_guid>/<destination_id>
```
