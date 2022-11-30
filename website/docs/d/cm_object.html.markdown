---
layout: "ibm"
page_title: "IBM : ibm_cm_object"
description: |-
  Get information about ibm_cm_object
subcategory: "Catalog Management"
---

# ibm_cm_object

Provides a read-only data source for ibm_cm_object. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cm_object" "cm_object" {
	catalog_id = ibm_cm_object.cm_object.catalog_id
	object_id = ibm_cm_object.cm_object.id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `catalog_id` - (Required, Forces new resource, String) Catalog identifier.
* `object_id` - (Required, Forces new resource, String) Object identification.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the ibm_cm_object.
* `name` - (String) The programmatic name of this object.
* `rev` - (String) Cloudant revision.
* `crn` - (String) The crn for this specific object.
* `url` - (String) The url for this specific object.
* `parent_id` - (String) The parent for this specific object.
* `label` - (String) Display name in the requested language.
* `tags` - (List) List of tags associated with this catalog.
* `created` - (String) The date and time this catalog was created.
* `updated` - (String) The data and time this catalog was last updated.
* `short_description` - (String) Short description in the requested language.
* `kind` - (String) Kind of object.
* `publish` - Publish information.
* Nested scheme for **publish**:
	* `permit_ibm_public_publish` - (Boolean) Is it permitted to request publishing to IBM or Public.
	* `ibm_approved` - (Boolean) Indicates if this offering has been approved for use by all IBMers.
	* `public_approved` - (Boolean) Indicates if this offering has been approved for use by all IBM Cloud users.
	* `portal_approval_record` - (String) The portal's approval record ID.
	* `portal_url` - (String) The portal UI URL.
* `state` - Object state.
* Nested scheme for **state**:
	* `current` - (String) one of: new, validated, account-published, ibm-published, public-published.
	* `current_entered` - (String) Date and time of current request.
	* `pending` - (String) one of: new, validated, account-published, ibm-published, public-published.
	* `pending_requested` - (String) Date and time of pending request.
	* `previous` - (String) one of: new, validated, account-published, ibm-published, public-published.
* `catalog_name` - (String) The name of the catalog.
* `data` - (String) Stringified map of object data.
