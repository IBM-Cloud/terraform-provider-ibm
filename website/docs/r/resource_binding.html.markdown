---
layout: "ibm"
page_title: "IBM : resource_binding"
description: |-
  Manages resource_binding.
subcategory: "Resource management"
---

# ibm\_resource_binding

Provides a resource for resource_binding. This allows resource_binding to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_resource_binding" "resource_binding" {
  source = "25eba2a9-beef-450b-82cf-f5ad5e36c6dd"
  target = "crn:v1:bluemix:public:cf:us-south:s/0ba4dba0-a120-4a1e-a124-5a249a904b76::cf-application:a1caa40b-2c24-4da8-8267-ac2c1a42ad0c"
  name = "my-binding"
  role = "Writer"
}
```

## Argument Reference

The following arguments are supported:

* `source` - (Required, string) The short or long ID of resource alias.
* `target` - (Required, string) The CRN of application to bind to in a specific environment, for example, Dallas YP, CFEE instance.
* `name` - (Optional, string) The name of the binding. Must be 180 characters or less and cannot include any special characters other than `(space) - . _ :`.
  * Constraints: The value must match regular expression `/^([^[:ascii:]]|[a-zA-Z0-9-._: ])+$/`
* `parameters` - (Optional, List) Configuration options represented as key-value pairs. Service defined options are passed through to the target resource brokers, whereas platform defined options are not.
  * `serviceid_crn` - (Optional, string) An optional platform defined option to reuse an existing IAM serviceId for the role assignment.
* `role` - (Optional, string) The role name or it's CRN.
  * Constraints: The default value is `Writer`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the resource_binding.
* `guid` - When you create a new binding, a globally unique identifier (GUID) is assigned. This GUID is a unique internal identifier managed by the resource controller that corresponds to the binding.
* `url` - When you provision a new binding, a relative URL path is created identifying the location of the binding.
* `created_at` - The date when the binding was created.
* `updated_at` - The date when the binding was last updated.
* `deleted_at` - The date when the binding was deleted.
* `created_by` - The subject who created the binding.
* `updated_by` - The subject who updated the binding.
* `deleted_by` - The subject who deleted the binding.
* `source_crn` - The CRN of resource alias associated to the binding.
* `target_crn` - The CRN of target resource, for example, application, in a specific environment.
* `crn` - The full Cloud Resource Name (CRN) associated with the binding. For more information about this format, see [Cloud Resource Names](https://cloud.ibm.com/docs/overview?topic=overview-crn).
* `region_binding_id` - The ID of the binding in the specific target environment, for example, `service_binding_id` in a given IBM Cloud environment.
* `region_binding_crn` - The CRN of the binding in the specific target environment.
* `account_id` - An alpha-numeric value identifying the account ID.
* `resource_group_id` - The ID of the resource group.
* `state` - The state of the binding.
* `credentials` - The credentials for the binding. Additional key-value pairs are passed through from the resource brokers.  For additional details, see the service’s documentation.
* `iam_compatible` - Specifies whether the binding’s credentials support IAM.
* `resource_id` - The unique ID of the offering. This value is provided by and stored in the global catalog.
* `migrated` - A boolean that dictates if the alias was migrated from a previous CF instance.
* `resource_alias_url` - The relative path to the resource alias that this binding is associated with.

## Import

You can import the `ibm_resource_binding` resource by using `id`. The ID associated with the binding.

```
$ terraform import ibm_resource_binding.resource_binding <id>
```
