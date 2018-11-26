---
layout: "ibm"
page_title: "IBM: app_domain_private"
sidebar_current: "docs-ibm-resource-app-domain-private"
description: |-
  Manages IBM Application private domain.
---

# ibm\_app_domain_private

Provides a private domain resource. This allows private domains to be created, updated, and deleted.

## Example Usage

```hcl
data "ibm_org" "orgdata" {
  org = "someexample.com"
}

resource "ibm_app_domain_private" "domain" {
  name     = "foo.com"
  org_guid = "${data.ibm_org.orgdata.id}"
  tags     = ["tag1", "tag2"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the domain.
* `org_guid` - (Required, string) The GUID of the organization that owns the domain. You can retrieve the value from data source `ibm_org` or by running the `ibmcloud iam orgs --guid` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `tags` - (Optional, array of strings) Tags associated with the application private domain instance.  
  **NOTE**: `Tags` are managed locally and are currently not stored on the IBM Cloud service endpoint.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the private domain.
