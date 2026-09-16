---
layout: "ibm"
page_title: "IBM : ibm_restapi_request"
description: |-
  Manages an object through an arbitrary IBM Cloud REST API.
subcategory: "REST API"
---

# ibm_restapi_request

Create, update, and delete an object through an arbitrary IBM Cloud REST API. Use this resource for IBM Cloud services that do not have a dedicated Terraform resource yet, instead of adding a third party REST provider to your configuration.

The request is authenticated with the credentials that are already configured on the `ibm` provider, so you do not need to manage an API key or a bearer token separately.

## Example usage

Create a resource group through the Resource Controller API.

```hcl
resource "ibm_restapi_request" "resource_group" {
  url          = "https://resource-controller.cloud.ibm.com/v2/resource_groups"
  id_attribute = "id"

  request_body = jsonencode({
    name       = "my-group"
    account_id = var.account_id
  })

  update_method = "PATCH"
  update_request_body = jsonencode({
    name = "my-group"
  })
}

output "resource_group_id" {
  value = ibm_restapi_request.resource_group.object_id
}
```

Call an API that does not return an identifier and does not support a read operation.

```hcl
resource "ibm_restapi_request" "settings" {
  url            = "https://example.cloud.ibm.com/v1/settings"
  create_method  = "PUT"
  destroy_method = "DELETE"
  read_on_create = false

  request_body = jsonencode({
    enabled = true
  })

  headers = {
    "X-Correlation-Id" = "terraform"
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `accept_status_codes` - (Optional, List of Integers) Status codes outside the 2xx range that are treated as success instead of an error, for example `[409]`.
- `create_method` - (Optional, Forces new resource, String) The HTTP method used to create the object. One of `POST`, `PUT`, `PATCH`. The default value is `POST`.
- `destroy_method` - (Optional, String) The HTTP method used to destroy the object. One of `DELETE`, `POST`. The default value is `DELETE`.
- `destroy_request_body` - (Optional, String) The JSON body that is sent when the object is destroyed. Most APIs do not need one.
- `headers` - (Optional, Map of Strings) Additional HTTP headers that are sent with every request. `Accept` defaults to `application/json`, and `Content-Type` defaults to `application/json` when a body is sent.
- `id_attribute` - (Optional, Forces new resource, String) Dot separated path to the identifier inside the create response, for example `id` or `metadata.guid`. When it is set, the object URL that is used for read, update and destroy is the create URL with the identifier appended. When it is not set, the create URL is used unchanged.
- `ignore_read_errors` - (Optional, Boolean) Whether a failing read is ignored instead of failing the run. The last known response body is kept in state. The default value is `false`.
- `query_params` - (Optional, Map of Strings) Query parameters that are appended to every request URL.
- `read_method` - (Optional, String) The HTTP method used to read the object back. One of `GET`, `POST`. The default value is `GET`.
- `read_on_create` - (Optional, Boolean) Whether the object is read back after it is created. Set it to `false` for APIs that do not expose a read operation, in which case the create response is kept as the object state. The default value is `true`.
- `request_body` - (Optional, String) The JSON body that is sent when the object is created, and when it is updated unless `update_request_body` is set. Use the `jsonencode` function to build it.
- `sensitive_headers` - (Optional, Map of Strings) Additional HTTP headers that are sent with every request and are kept out of the Terraform output. Values here override values with the same key in `headers`.
- `timeout_seconds` - (Optional, Integer) How long a single HTTP request may take before it is cancelled. Between `1` and `3600`. The default value is `60`.
- `update_method` - (Optional, String) The HTTP method used to update the object. One of `PUT`, `PATCH`, `POST`. The default value is `PUT`.
- `update_request_body` - (Optional, String) The JSON body that is sent when the object is updated. When it is not set, `request_body` is sent instead.
- `url` - (Required, Forces new resource, String) The URL that the object is created against. For a collection style API this is the collection URL, for example `https://resource-controller.cloud.ibm.com/v2/resource_instances`.
- `use_iam_auth` - (Optional, Boolean) Whether the IBM Cloud credentials that are configured on the provider are used to authenticate the request. Set it to `false` when the target API is not IAM protected, or when you supply an `Authorization` header through `sensitive_headers`. The default value is `true`.

## Attribute reference

In addition to all argument references, you can access the following attribute references after your resource is created.

- `create_response_body` - (String) The body that is returned by the create request. Useful for APIs that return generated values only once.
- `id` - (String) The unique identifier of the request, which is the object URL.
- `object_id` - (String) The identifier that is read out of the create response by using `id_attribute`. Empty when `id_attribute` is not set.
- `object_url` - (String) The URL that is used to read, update and destroy the object.
- `response_body` - (String) The body of the most recent read response, or of the create response when `read_on_create` is `false`. Use the `jsondecode` function to work with it.
- `response_headers` - (Map of Strings) The headers of the most recent response.
- `response_status_code` - (Integer) The status code of the most recent response.

## Behavior notes

- A read that returns `404` or `410` removes the object from the Terraform state, so the next plan recreates it.
- A destroy that returns `404` or `410` is treated as success, because the object is already gone.
- When the object is created but `id_attribute` cannot be resolved in the response, the run fails with an error that names the keys the response did contain. The object exists remotely at that point and may have to be removed manually.
- Only `request_body`, `update_request_body`, `headers`, `sensitive_headers` and `query_params` trigger an update call. Changing `url`, `create_method` or `id_attribute` replaces the object.

## Import

You can import the `ibm_restapi_request` resource by using the object URL. When `id_attribute` is set, the create URL is a collection URL that cannot be derived from the object URL, so pass both, separated by a comma.

```
$ terraform import ibm_restapi_request.settings https://example.cloud.ibm.com/v1/settings
```

```
$ terraform import ibm_restapi_request.resource_group https://resource-controller.cloud.ibm.com/v2/resource_groups,https://resource-controller.cloud.ibm.com/v2/resource_groups/abc123
```

Arguments that the API does not echo back, such as `request_body` and the method overrides, are not recovered by an import. Add them to your configuration to match the imported object.
