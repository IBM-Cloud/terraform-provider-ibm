---
layout: "ibm"
page_title: "IBM : resource_alias"
description: |-
  Manages resource_alias.
subcategory: "Resource management"
---

# ibm\_resource_alias

Provides a resource for resource_alias. This allows resource_alias to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_resource_alias" "resource_alias" {
  name = "my-alias"
  source = "a8dff6d3-d287-4668-a81d-c87c55c2656d"
  target = "crn:v1:bluemix:public:cf:us-south:o/5e939cd5-6377-4383-b9e0-9db22cd11753::cf-space:66c8b915-101a-406c-a784-e6636676e4f5"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the alias. Must be 180 characters or less and cannot include any special characters other than `(space) - . _ :`.
  * Constraints: The value must match regular expression `/^([^[:ascii:]]|[a-zA-Z0-9-._: ])+$/`
* `source` - (Required, string) The short or long ID of resource instance.
* `target` - (Required, string) The CRN of target name(space) in a specific environment, for example, space in Dallas YP, CFEE instance etc.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the resource_alias.
* `guid` - When you create a new alias, a globally unique identifier (GUID) is assigned. This GUID is a unique internal indentifier managed by the resource controller that corresponds to the alias.
* `url` - When you created a new alias, a relative URL path is created identifying the location of the alias.
* `created_at` - The date when the alias was created.
* `updated_at` - The date when the alias was last updated.
* `deleted_at` - The date when the alias was deleted.
* `created_by` - The subject who created the alias.
* `updated_by` - The subject who updated the alias.
* `deleted_by` - The subject who deleted the alias.
* `resource_instance_id` - The ID of the resource instance that is being aliased.
* `target_crn` - The CRN of the target namespace in the specific environment.
* `account_id` - An alpha-numeric value identifying the account ID.
* `resource_id` - The unique ID of the offering. This value is provided by and stored in the global catalog.
* `resource_group_id` - The ID of the resource group.
* `crn` - The CRN of the alias. For more information about this format, see [Cloud Resource Names](https://cloud.ibm.com/docs/overview?topic=overview-crn).
* `region_instance_id` - The ID of the instance in the specific target environment, for example, `service_instance_id` in a given IBM Cloud environment.
* `region_instance_crn` - The CRN of the instance in the specific target environment.
* `state` - The state of the alias.
* `migrated` - A boolean that dictates if the alias was migrated from a previous CF instance.
* `resource_instance_url` - The relative path to the resource instance.
* `resource_bindings_url` - The relative path to the resource bindings for the alias.
* `resource_keys_url` - The relative path to the resource keys for the alias.

## Import

You can import the `ibm_resource_alias` resource by using `id`. The ID associated with the alias.

```
$ terraform import ibm_resource_alias.resource_alias <id>
```
