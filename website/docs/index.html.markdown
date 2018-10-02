---
layout: "ibm"
page_title: "Provider: IBM"
sidebar_current: "docs-ibm-index"
description: |-
  The IBM Cloud provider is used to interact with IBM Cloud resources.
---

# IBM Cloud Provider

The IBM Cloud provider is used to manage IBM Cloud resources. The provider must be configured with the proper credentials before it can be used.

Use the navigation menu on the left to read about the available data sources and resources.

## Example Usage


```hcl
# Configure the IBM Cloud Provider
provider "ibm" {
  bluemix_api_key    = "${var.ibm_bmx_api_key}"
  softlayer_username = "${var.ibm_sl_username}"
  softlayer_api_key  = "${var.ibm_sl_api_key}"
}

# Create an IBM Cloud infrastructure SSH key. You can find the SSH key surfaces in the infrastructure console under Devices > Manage > SSH Keys
resource "ibm_compute_ssh_key" "test_key_1" {
  label      = "test_key_1"
  public_key = "${var.ssh_public_key}"
}

# Create a virtual server with the SSH key
resource "ibm_compute_vm_instance" "my_server_2" {
  hostname          = "host-b.example.com"
  domain            = "example.com"
  ssh_key_ids       = [123456, "${ibm_compute_ssh_key.test_key_1.id}"]
  os_reference_code = "CENTOS_6_64"
  datacenter        = "ams01"
  network_speed     = 10
  cores             = 1
  memory            = 1024
}

# Reference details of the IBM Cloud space
data "ibm_space" "space" {
  space = "${var.space}"
  org   = "${var.org}"
}

# Create an instance of an IBM service
resource "ibm_service_instance" "service" {
  name       = "${var.instance_name}"
  space_guid = "${data.ibm_space.space.id}"
  service    = "speech_to_text"
  plan       = "lite"
  tags       = ["cluster-service", "cluster-bind"]
}

# Create a Cloud Functions action
resource "ibm_function_action" "nodehello" {
  name = "action-name"
  exec = {
    kind = "nodejs:6"
    code = "${file("hellonode.js")}"
  }
}

```

## Authentication

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

### Static credentials ###

You can provide your static credentials by adding the `bluemix_api_key`, `softlayer_username`, and `softlayer_api_key` arguments in the IBM Cloud provider block.

Usage:

```hcl
provider "ibm" {
    bluemix_api_key = ""
    softlayer_username = ""
    softlayer_api_key = ""
}
```


### Environment variables

You can provide your credentials by exporting the `BM_API_KEY`, `SL_USERNAME`, and `SL_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```hcl
provider "ibm" {}
```

Usage:

```shell
export BM_API_KEY="bmx_api_key"
export SL_USERNAME="sl_username"
export SL_API_KEY="sl_api_key"
terraform plan
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `bluemix_api_key` - (optional) The IBM Cloud platform API key. You must either add it as a credential in the provider block or source it from the `BM_API_KEY` (higher precedence) or `BLUEMIX_API_KEY` environment variable. The key is required to provision Cloud Foundry or IBM Cloud Container Service resources, such as any resource that begins with `ibm` or `ibm_container`.

* `bluemix_timeout` - (optional) The timeout, expressed in seconds, for interacting with IBM Cloud APIs. You can also source the timeout from the `BM_TIMEOUT` (higher precedence) or `BLUEMIX_TIMEOUT` environment variable. The default value is `60`.

* `softlayer_username` - (optional) The IBM Cloud infrastructure (SoftLayer) user name. You must either add it as a credential in the provider block or source it from the `SL_USERNAME` (higher precedence) or `SOFTLAYER_USERNAME` environment variable.

* `softlayer_api_key` - (optional) The IBM Cloud infrastructure API key. You must either add it as a credential in the provider block or source it from the `SL_API_KEY` (higher precedence) or `SOFTLAYER_API_KEY` environment variable. The key is required to provision infrastructure resources, such as any resource that begins with `ibm_compute`.

* `softlayer_endpoint_url` - (optional) The IBM Cloud infrastructure endpoint url. You can also source it from the `SL_ENDPOINT_URL` (higher precedence) or `SOFTLAYER_ENDPOINT_URL` environment variable. The default value is `https://api.softlayer.com/rest/v3`.

* `softlayer_timeout` - (optional) The timeout, expressed in seconds, for the IBM Cloud infrastructure API key. You can also source the timeout from the `SL_TIMEOUT` (higher precedence) or `SOFTLAYER_TIMEOUT` environment variable. The default value is `60`.

* `region` - (optional) The IBM Cloud region. You can also source it from the `BM_REGION` (higher precedence) or `BLUEMIX_REGION` environment variable. The default value is `us-south`.

* `max_retries` - (Optional) This is the maximum number of times an IBM Cloud infrastructure API call is retried, in the case where requests are getting network related timeout and rate limit exceeded error code. You can also source it from the `MAX_RETRIES` environment variable. The default value is `5`.

* `function_namespace` - (Optional) Your Cloud Functions namespace is composed from your IBM Cloud org and space like \<org\>_\<space\>. This attribute is required only when creating a Cloud Functions resource. It must be provided when you are creating such resources in IBM Cloud. You can also source it from the FUNCTION_NAMESPACE environment variable.
