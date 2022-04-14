---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_credential"
description: |-
  Manages credentials.
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_credential

Provides a resource for credentials. This allows credentials to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_scc_posture_credential" "credentials" {
  description = "This credential is used for testing."
  display_fields = {"password":"testpassword","username":"test"}
  enabled = true
  group = {"id":"1","passphrase":"passphrase"}
  name = "test_create"
  purpose = "discovery_fact_collection_remediation"
  type = "username_password"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `description` - (Required, String) Credentials description.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\._,\\s]*$/`.
* `display_fields` - (Required, List) Details the fields on the credential. This will change as per credential type selected.
Nested scheme for **display_fields**:
	* `auth_url` - (Optional, String) auth url of the Open Stack cloud.This is mandatory for Open Stack Credential type ie when type=openstack_cloud.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `aws_arn` - (Optional, String) AWS arn value.This is used for AWS Cloud ie when type=aws_cloud.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `aws_client_id` - (Optional, String) AWS client Id.This is mandatory for AWS Cloud ie when type=aws_cloud.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `aws_client_secret` - (Optional, String) AWS client secret.This is mandatory for AWS Cloud ie when type=aws_cloud.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `aws_region` - (Optional, String) AWS region.This is used for AWS Cloud ie when type=aws_cloud.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `azure_client_id` - (Optional, String) Azure client Id. This is mandatory for Azure Credential type ie when type=azure_cloud.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `azure_client_secret` - (Optional, String) Azure client secret.This is mandatory for Azure Credential type ie when type=azure_cloud.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `azure_resource_group` - (Optional, String) Azure resource group.This field is used for Azure Credential type ie when type=azure_cloud.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `azure_subscription_id` - (Optional, String) Azure subscription Id.This is mandatory for Azure Credential type ie when type=azure_cloud.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `database_name` - (Optional, String) Database name.This is mandatory for Database Credential type ie when type=database.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `ibm_api_key` - (Optional, String) The IBM Cloud API Key. This is mandatory for IBM Credential Type ie when type=ibm_cloud.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `ms_365_client_id` - (Optional, String) The MS365 client Id.This is mandatory for Windows MS365 Credential type ie when type=ms_365.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `ms_365_client_secret` - (Optional, String) The MS365 client secret.This is mandatory for Windows MS365 Credential type ie when type=ms_365.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `ms_365_tenant_id` - (Optional, String) The MS365 tenantId.This is mandatory for Windows MS365 Credential type ie when type=ms_365.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `password` - (Optional, String) password of the user.This is mandatory for DataBase(ie type=database), Kerbros(ie type=kerberos_windows),OpenStack(ie type=openstack_cloud) and Username-Password(type=username_password) Credentials.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `pem_data` - (Optional, String) The base64 encoded data to associate with the PEM file.
	  * Constraints: The maximum length is `4000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `pem_file_name` - (Optional, String) The name of the PEM file.
	  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.  
	* `project_domain_name` - (Optional, String) project domain name of the Open Stack cloud.This is mandatory for Open Stack Credential type ie when type=openstack_cloud.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `project_name` - (Optional, String) Project name of the Open Stack cloud.This is mandatory for Open Stack Credential type ie when type=openstack_cloud.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `user_domain_name` - (Optional, String) user domain name of the Open Stack cloud.This is mandatory for Open Stack Credential type ie when type=openstack_cloud.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `username` - (Optional, String) username of the user.This is mandatory for DataBase(ie type=database), Kerbros(ie type=kerberos_windows),OpenStack(ie type=openstack_cloud) and Username-Password(type=username_password) Credentials.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `winrm_authtype` - (Optional, String) Kerberos windows auth type.This is mandatory for Windows Kerberos Credential type ie when type=kerberos_windows.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
	* `winrm_port` - (Optional, String) Kerberos windows port.This is mandatory for Windows Kerberos Credential type ie when type=kerberos_windows.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.
	* `winrm_usessl` - (Optional, String) Kerberos windows ssl.This is mandatory for Windows Kerberos Credential type ie when type=kerberos_windows.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
* `enabled` - (Required, Boolean) Credentials status enabled/disbaled.
* `group` - (Required, List) Credential group details.
Nested scheme for **group**:
	* `id` - (Required, String) credential group id.
	  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.
	* `passphrase` - (Required, String) passphase of the credential.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.\\*,_\\s]*$/`.
* `name` - (Required, String) Credentials name.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\._,\\s]*$/`.
* `purpose` - (Required, String) Purpose for which the credential is created.
  * Constraints: Allowable values are: `discovery_collection`, `discovery_fact_collection`, `remediation`, `discovery_collection_remediation`, `discovery_fact_collection_remediation`. The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
* `type` - (Required, String) Credentials type.
  * Constraints: Allowable values are: `username_password`, `aws_cloud`, `azure_cloud`, `database`, `kerberos_windows`, `ms_365`, `openstack_cloud`, `ibm_cloud`, `user_name_pem`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the credentials.

## Import

You can import the `ibm_scc_posture_credential` resource by using `id`. An identifier of the credential.

# Syntax
```
$ terraform import ibm_scc_posture_credential.credentials <id>
```

# Example
```
$ terraform import ibm_scc_posture_credential.credentials 1
```
