---
layout: "ibm"
page_title: "IBM : ibm_code_engine_secret"
description: |-
  Manages code_engine_secret.
subcategory: "Code Engine"
---

# ibm_code_engine_secret

Provides a resource for code_engine_secret. This allows code_engine_secret to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_code_engine_secret" "code_engine_secret_instance" {
  project_id = "15314cc3-85b4-4338-903f-c28cdee6d005"
  name = "my-secret"
  format = "generic"

  data {
		key1 = "value1"
		key2 = "value2"
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `data` - (Optional, Map) The key-value pair for the secret. Values must be specified in `KEY=VALUE` format. Each `KEY` field must consist of alphanumeric characters, `-`, `_` or `.` and must not be exceed a max length of 253 characters. Each `VALUE` field can consists of any character and must not be exceed a max length of 1048576 characters. Depending on the `format`, certain `KEY=VALUE` are mandatory. For format `ssh_auth` you must provide the fields `ssh_key` and `known_hosts`. For format registry
  * Constraints: Depending on the `format`, certain key-value pairs are mandatory:
    * For format `generic` there are no constraints.
    * For format `ssh_auth` you must provide the fields `ssh_key` and `known_hosts`.
    * For format `basic_auth` you must provide the fields `username` and `password`.
    * For format `tls` you must provide the fields `tls_cert` and `tls_key`.
    * For format `registry` you must provide the fields `username`,  `password`, `server` and `email`.
* `format` - (Required, Forces new resource, String) Specify the format of the secret.
  * Constraints: Allowable values are: `generic`, `ssh_auth`, `basic_auth`, `tls`, `registry`. The value must match regular expression `/^(generic|ssh_auth|basic_auth|tls|registry)$/`.
* `name` - (Required, String) The name of the secret.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `secret_id` - (String) The identifier of the resource.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `created_at` - (String) The timestamp when the resource was created.
* `entity_tag` - (String) The version of the secret instance, which is used to achieve optimistic locking.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[\\*\\-a-z0-9]+$/`.
* `href` - (String) When you provision a new secret,  a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `id` - (String) The identifier of the resource.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `resource_type` - (String) The type of the secret.
* `etag` - ETag identifier for code_engine_secret.

## Import

You can import the `ibm_code_engine_secret` resource by using `name`.
The `name` property can be formed from `project_id`, and `name` in the following format:

```
<project_id>/<name>
```
* `project_id`: A string in the format `15314cc3-85b4-4338-903f-c28cdee6d005`. The ID of the project.
* `name`: A string in the format `my-secret`. The name of your secret.

# Syntax
```
$ terraform import ibm_code_engine_secret.code_engine_secret <project_id>/<name>
```

# Example
```
$ terraform import ibm_code_engine_project.code_engine_project "15314cc3-85b4-4338-903f-c28cdee6d005/my-secret"
```
