---
layout: "ibm"
page_title: "IBM : ibm_sm_en_registration"
description: |-
  Manages NotificationsRegistrationPrototype.
subcategory: "Secrets Manager"
---

# ibm_sm_en_registration

Provides a resource for NotificationsRegistrationPrototype. This allows NotificationsRegistrationPrototype to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_sm_en_registration" "sm_en_registration" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  event_notifications_instance_crn = "crn:v1:bluemix:public:event-notifications:us-south:a/22018f3c34ff4ff193698d15ca316946:578ad1a4-2fd8-4e66-95d5-79a842ba91f8::"
  event_notifications_source_description = "Optional description of this source in an Event Notifications instance."
  event_notifications_source_name = "My Secrets Manager"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
  * Constraints: Allowable values are: `private`, `public`.
* `event_notifications_instance_crn` - (Required, Forces new resource, String) A CRN that uniquely identifies an IBM Cloud resource.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
* `event_notifications_source_description` - (Optional, Forces new resource, String) An optional description for the source  that is in your Event Notifications instance.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/(.*?)/`.
* `event_notifications_source_name` - (Required, Forces new resource, String) The name that is displayed as a source that is in your Event Notifications instance.
  * Constraints: The maximum length is `256` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the NotificationsRegistrationPrototype.

## Provider Configuration

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

To find which credentials are required for this resource, see the service table [here](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).

### Static credentials

You can provide your static credentials by adding the `ibmcloud_api_key`, `iaas_classic_username`, and `iaas_classic_api_key` arguments in the IBM Cloud provider block.

Usage:
```
provider "ibm" {
    ibmcloud_api_key = ""
    iaas_classic_username = ""
    iaas_classic_api_key = ""
}
```

### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```
provider "ibm" {}
```

Usage:
```
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
terraform plan
```

Note:

1. Create or find your `ibmcloud_api_key` and `iaas_classic_api_key` [here](https://cloud.ibm.com/iam/apikeys).
  - Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloud_api_key`
  - Select `Classic Infrastructure API Keys` option from view dropdown for `iaas_classic_api_key`
2. For iaas_classic_username
  - Go to [Users](https://cloud.ibm.com/iam/users)
  - Click on user.
  - Find user name in the `VPN password` section under `User Details` tab

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_sm_en_registration` resource by using `region` and `instance_id`. Note that after the import, the value of the property `event_notifications_source_name`
is `null` in the Terraform state file. Since this is a required property, you need to provide a value for it in
the corresponding configuration block in your Terraform configuration file. If you don't know the actual name of the event notification source, 
you can put any value. The value is ignored when Terraform updates the resource, because the Secrets Manager API does not modify
the source name after the registration resource is created.
For more information, see [the documentation](https://cloud.ibm.com/docs/secrets-manager).

# Syntax
```bash
$ terraform import ibm_sm_en_registration.sm_en_registration <region>/<instance_id>
```

# Example
```bash
$ terraform import ibm_sm_en_registration.sm_en_registration us-east/6ebc4224-e983-496a-8a54-f40a0bfa9175
```
