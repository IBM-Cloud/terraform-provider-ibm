---
layout: "ibm"
page_title: "IBM : ibm_schematics_kms"
description: |-
  Get information about schematics_kms
subcategory: "Schematics"
---

# ibm_schematics_kms

Retrieve the KMS (Key Management Service) settings integrated with IBM Cloud Schematics for a given location. This enables retrieval of BYOK (Bring Your Own Key) and KYOK (Keep Your Own Key) configuration details per geographic region. For more information, about Schematics KMS settings, see [Securing your data in Schematics](https://cloud.ibm.com/docs/schematics?topic=schematics-secure-data).

## Example Usage

```hcl
data "ibm_schematics_kms" "schematics_kms" {
	location = "us-south"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `location` - (Optional, String) Location supported by IBM Cloud Schematics service.  

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the schematics_kms.
* `encryption_scheme` - (String) The encryption scheme values. Allowable values: `byok`, `kyok`.
* `resource_group` - (String) The kms instance resource group to integrate.

* `primary_crk` - (List) The primary kms instance details.
Nested scheme for **primary_crk**:
	* `kms_name` - (String) The primary kms instance name.
	* `kms_private_endpoint` - (String) The primary kms instance private endpoint.
	* `key_crn` - (String) The CRN of the primary root key.

* `secondary_crk` - (List) The secondary kms instance details.
Nested scheme for **secondary_crk**:
	* `kms_name` - (String) The secondary kms instance name.
	* `kms_private_endpoint` - (String) The secondary kms instance private endpoint.
	* `key_crn` - (String) The CRN of the secondary key.