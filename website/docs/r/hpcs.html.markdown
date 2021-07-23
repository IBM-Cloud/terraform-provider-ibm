---
subcategory: "Hyper Protect Crypto Service (HPCS)"
layout: "ibm"
page_title: "IBM : Hyper Protect Crypto Service instance"
description: |-
  Manages IBM Cloud Hyper Protect Crypto Service Instance.
---

# ibm\_hpcs

Manages HPCS resource. This allows hpcs sub-resources to be added to an existing hpcs instance.

~> **Note:** As recovery crypto units are available only in us-south and us-east, only these two regions are supported if you want to use this resource for instance initialization.

## Example Usage

```terraform
resource ibm_hpcs hpcs {
  location             = "us-south"
  name                 = "test-hpcs"
  plan                 = "standard"
  units                = 2
  signature_threshold  = 1
  revocation_threshold = 1
  admins {
    name  = "admin1"
    key   = "/cloudTKE/1.sigkey"
    token = "<sensitive1234>"
  }
  admins {
    name  = "admin2"
    key   = "/cloudTKE/2.sigkey"
    token = "<sensitive1234>"
  }
}
```

## Argument Reference

The following arguments are supported:
* `admins` - (Required, List) The list of administrators for the instance crypto units. You can set up to 8 administrators and the number needs to be equal to or greater than the thresholds that you specify. The following values need to be set for each administrator: 
  Nested scheme for `admins`:
  * `key` - (Required, string) The path in your local workstation where you store the administrator signature file.
  * `name` - (Required, string) The name of the administrator. name is limited to 30 characters.
  * `token` - (Required, string, Sensitive) The administrator password to access the corresponding signature file.
  
~> **Note:** When a `signature_server_url` (signing service) is used, the `key` parameter of `admins` identifies the signature key to be used.  How the key parameter is defined depends on the signing service.  The character string for the key parameter is appended to a URI and must contain only unreserved characters as defined by section 2.3 of RFC 3986.  The `token` parameter of `admins` authorizes use of the signature key.  How the token parameter is defined depends on the signing service.
* `failover_units` - (Optional, string) The number of failover crypto units for your service instance. Valid values are 2 or 3 but it must be less than or equal to the number of operational crypto units.
* `location` - (Required, string) The region abbreviation, such as us-south, that represents the geographic area where the operational crypto units of your service instance are located. For more information, see Regions and locations. As recovery crypto units are available only in us-south and us-east, only these two regions are supported if you want to use Terraform for instance initialization.
* `name` - (Required, string) The name of your Hyper Protect Crypto Service instance.
* `plan` - (Required, string) The pricing plan for your service instance. Currently, only the standard plan is supportd.
* `resource_group_id` - (Optional, string) The Id of resource group where you want to organize and manage your service instance.
* `revocation_threshold` - (Required, int) The number of administrator signatures that is required to remove an administrator after you leave imprint mode. The valid value is between 1 and 8.
* `service_endpoints` - (Optional, string) Types of the service endpoints that can be set to a resource instance. Possible values are `public-and-private`, `private-only`.
* `signature_server_url` - (Optional, string) The URL of the signing service. If you use a third-party signing service to provide administrator signature keys, you need to specify the URL.
* `signature_threshold`- (Required, int)  The number of administrator signatures that is required to execute administrative commands. The valid value is between 1 and 8. You need to set it to at least 2 to enable quorum authentication.
* `tags` - (Optional, array of strings) Tags that are associated with your instance are used to organize your resources. 
* `units` -(Required, string) The number of operational crypto units for your service instance. Valid values are 2 and 3.
## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `created_at` - (String) The date when the instance was created.
* `created_by` - (String) The subject who created the instance.
* `crn` - (String) CRN of HPCS Instance.
* `deleted_at` - (String) The date when the instance was deleted.
* `deleted_by` - (String) The subject who deleted the instance.
* `extensions` - (List) The extended metadata as a map associated with the resource instance.
* `guid` - (String) Unique identifier of resource instance.
* `hsm_info` - (List) HSM config of HPCS Instance Crypto Units.
  Nested scheme for `hsm_info`:
  * `admins` - (List) List of Admins for Crypto Units
    Nested scheme for `admins`:
      * `name` - (String) Name of Admin.
      * `ski` - (String) Subject Key Identifier of the administrator signature key.
  * `current_mk_status` - (String) Status of Current Master Key Register.
  * `current_mkvp` - (String) Current Master Key Register Verification Pattern.
  * `hsm_id` - (String) HSM ID.
  * `hsm_location` - (String) HSM Location.
  * `hsm_type` - (String) HSM Type.
  * `new_mk_status` - (String) Status of New Master Key Register.
  * `new_mkvp` - (String) New Master Key Register Verification Pattern.
  * `revocation_threshold` - (Int) Revocation Threshold for Crypto Units.
  * `signature_threshold`- (Int) Signature Threshold for Crypto Units.
* `id` - (String) The unique identifier CRN of this HPCS instance.
* `location` - (String) The location for this HPCS instance.
* `plan` - (String) The pricing plan for your service instance.
* `resource_aliases_url` - (String) The relative path to the resource aliases for the instance.
* `resource_bindings_url` - (String) The relative path to the resource bindings for the instance.
* `resource_keys_url` - (String) The relative path to the resource keys for the instance.
* `restored_at` - (String) The date when the instance under reclamation was restored.
* `restored_by` - (String) The subject who restored the instance back from reclamation.
* `scheduled_reclaim_at` - (String) The date when the instance was scheduled for reclamation.
* `scheduled_reclaim_by` - (String) The subject who initiated the instance reclamation.
* `service` - (String) The service type (`hs-crypto`) of the instance.
* `state` - (String) The current state of the instance. For example, if the instance is deleted, it will return removed.
* `status` - (String) Status of the hpcs instance.
* `update_at` - (String) The date when the instance was last updated.
* `update_by` - (String) The subject who updated the instance.

## Import
The `ibm_hpcs` can be imported by using the `crn`.

```bash
terraform import ibm_hpcs.hpcs <crn>
```

**Example**

```
$ terraform import ibm_hpcs.hpcs crn:v1:bluemix:public:hs-crypto:us-south:a/4448261269a14562b839e0a3019ed980:f115115b-5087-4a4e-9cc8-71acf0542c0d::
```
