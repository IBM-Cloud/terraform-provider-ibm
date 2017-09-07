---
layout: "ibm"
page_title: "IBM: app_domain_private"
sidebar_current: "docs-ibm-resource-app-domain-private"
description: |-
  Manages IBM Application private domain.
---

# ibm\_app_domain_private

Create, update, or delete private domain on IBM Bluemix.

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
* `org_guid` - (Required, string) The GUID of the organization that owns the domain. The values can be retrieved from data source `ibm_org`, or by running the `bx iam orgs --guid` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).
* `tags` - (Optional, array of strings) Set tags on the application private domain instance.

**NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the private domain.
