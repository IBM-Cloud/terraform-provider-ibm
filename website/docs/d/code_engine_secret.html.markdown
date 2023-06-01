---
layout: "ibm"
page_title: "IBM : ibm_code_engine_secret"
description: |-
  Get information about code_engine_secret
subcategory: "Code Engine"
---

# ibm_code_engine_secret

Provides a read-only data source for code_engine_secret. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_code_engine_secret" "code_engine_secret" {
	project_id = data.ibm_code_engine_project.code_engine_project.project_id
	name       = "my-secret"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `name` - (Required, Forces new resource, String) The name of your secret.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the code_engine_secret.
* `created_at` - (String) The timestamp when the resource was created.

* `data` - (Map) Data container that allows to specify config parameters and their values as a key-value map. Each key field must consist of alphanumeric characters, `-`, `_` or `.` and must not be exceed a max length of 253 characters. Each value field can consists of any character and must not be exceed a max length of 1048576 characters.

* `entity_tag` - (String) The version of the secret instance, which is used to achieve optimistic locking.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[\\*\\-a-z0-9]+$/`.

* `format` - (Forces new resource, String) Specify the format of the secret.
  * Constraints: Allowable values are: `generic`, `ssh_auth`, `basic_auth`, `tls`, `service_access`, `registry`, `other`. The value must match regular expression `/^(generic|ssh_auth|basic_auth|tls|service_access|registry|other)$/`.

* `href` - (String) When you provision a new secret,  a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `resource_type` - (String) The type of the secret.

* `secret_id` - (String) The identifier of the resource.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

