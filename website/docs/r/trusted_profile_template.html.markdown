---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile_template"
description: |-
  Manages IAM trusted profile templates
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_trusted_profile_template

Create, update, commit, and delete trusted profile templates with this resource.

## Example Usage

```hcl
resource "ibm_iam_trusted_profile_template" "trusted_profile_template_instance" {
	name = "${var.trusted_profile_template_name}"
	description = "${var.trusted_profile_template_description}"
	profile {
		name = "Profile from Template"
		description = "description of profile from template"
		rules {
			name = "name"
			type = "Profile-SAML"
			realm_name = "test-realm-979"
			expiration = 1
			conditions {
				claim = "claim"
				operator = "EQUALS"
				value = "\"value\""
			}
		}
		identities {
			iam_id = "IBMid-123456789"
			identifier = "IBMid-123456789"
			type = "user"
			accounts = ["3213asv21s3d2vsd6bv54sfb321dfb"]
		}
	}
}
```

Trusted profile template with a Service Id identity and two access policy templates

```hcl
resource "ibm_iam_trusted_profile_template" "trusted_profile_template_instance" {
	name = "${var.trusted_profile_template_name}"
	description = "${var.trusted_profile_template_description}"
	profile {
		name = "Profile from Template"
		description = "description of profile from template"
		identities {
			iam_id = "iam-ServiceId-abcd83bd-e218-48af-8073-c0c1b3980001"
			identifier = "ServiceId-abcd83bd-e218-48af-8073-c0c1b3980001"
			type = "serviceid"
		}
	}
	policy_template_references {
		id = split("/", ibm_iam_policy_template.Kube_Administrator_ap_template.id)[0]
		version = ibm_iam_policy_template.Kube_Administrator_ap_template.version
  	}
  	policy_template_references {
		id = split("/", ibm_iam_policy_template.Services_Reader_ap_template.id)[0]
		version = ibm_iam_policy_template.Services_Reader_ap_template.version
  	}
  	committed = true
}
```

Trusted Profile template with service identity and one access policy

```hcl
resource "ibm_iam_trusted_profile_template" "trusted_profile_template_instance" {
	name = "${var.trusted_profile_template_name}"
	description = "${var.trusted_profile_template_description}"
	profile {
		name = "Profile from Template"
		description = "description of profile from template"
		identities {
			iam_id = format("crn-%s", var.instance_crn)
      		identifier = var.instance_crn
			type = "crn"
		}
	}
	policy_template_references {
		id = split("/", ibm_iam_policy_template.Kube_Administrator_ap_template.id)[0]
		version = ibm_iam_policy_template.Kube_Administrator_ap_template.version
  	}
  	committed = true
}
```

New Trusted Profile template version

```hcl
resource "ibm_iam_trusted_profile_template" "trusted_profile_template_instance" {
	name = "${var.trusted_profile_template_name}"
	description = "${var.trusted_profile_template_description}"
	profile {
		name = "Profile from Template"
		description = "description of profile from template"
		identities {
			iam_id = "iam-ServiceId-abcd83bd-e218-48af-8073-c0c1b3980001"
			identifier = "ServiceId-abcd83bd-e218-48af-8073-c0c1b3980001"
			type = "serviceid"
		}
	}
	policy_template_references {
		id = split("/", ibm_iam_policy_template.Kube_Administrator_ap_template.id)[0]
		version = ibm_iam_policy_template.Kube_Administrator_ap_template.version
  	}
  	policy_template_references {
		id = split("/", ibm_iam_policy_template.Services_Reader_ap_template.id)[0]
		version = ibm_iam_policy_template.Services_Reader_ap_template.version
  	}
  	committed = true
}

resource "ibm_iam_trusted_profile_template" "trusted_profile_template_v2" {
	template_id = ibm_iam_trusted_profile_template.trusted_profile_template_instance.id
	name = "${var.trusted_profile_template_name}"
	description = "${var.trusted_profile_template_description} v2"
	profile {
		name = "Profile from Template"
		description = "description of profile from template"
		identities {
			iam_id = "iam-ServiceId-abcd83bd-e218-48af-8073-c0c1b3980001"
			identifier = "ServiceId-abcd83bd-e218-48af-8073-c0c1b3980001"
			type = "serviceid"
		}
	}
	policy_template_references {
		id = split("/", ibm_iam_policy_template.Kube_Administrator_ap_template.id)[0]
		version = ibm_iam_policy_template.Kube_Administrator_ap_template.version
  	}
  	committed = true
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `name` - (Required, String) The name of the trusted profile template.
* `description` - (Optional, String) The description of the trusted profile template. Describe the template for enterprise account users.
* `profile` - (Optional, List) Input body parameters for the TemplateProfileComponent.
Nested schema for **profile**:
	* `name` - (Required, String) Name of the Profile.
	* `description` - (Optional, String) Description of the Profile.
	* `rules` - (Optional, List) Rules for the Profile.
	Nested schema for **rules**:
		* `conditions` - (Required, List) Conditions of this claim rule.
		Nested schema for **conditions**:
			* `claim` - (Required, String) The claim to evaluate against.
			* `operator` - (Required, String) The operation to perform on the claim. valid values are EQUALS, NOT_EQUALS, EQUALS_IGNORE_CASE, NOT_EQUALS_IGNORE_CASE, CONTAINS, IN.
			* `value` - (Required, String) The stringified JSON value that the claim is compared to using the operator.
		* `expiration` - (Optional, Integer) Session expiration in seconds, only required if type is 'Profile-SAML'.
		* `name` - (Optional, String) Name of the claim rule to be created or updated.
		* `realm_name` - (Optional, String) The realm name of the Idp this claim rule applies to. This field is required only if the type is specified as 'Profile-SAML'.
		* `type` - (Required, String) Type of the claim rule.
		  * Constraints: Allowable values are: `Profile-SAML`.
	* `identities` - (Optional, List) Identities for the Profile.
	  Nested schema for **identities**:
		* `accounts` - (Optional, List) Only valid for the type user. Accounts from which a user can assume the trusted profile.
		* `description` - (Optional, String) Description of the identity that can assume the trusted profile. This is optional field for all the types of identities. When this field is not set for the identity type 'serviceid' then the description of the service id is used. Description is recommended for the identity type 'crn' E.g. 'Instance 1234 of IBM Cloud Service project'.
		* `iam_id` - (Required, String) IAM ID of the identity.
		* `identifier` - (Required, String) Identifier of the identity that can assume the trusted profiles. This can be a user identifier (IAM id), serviceid or crn. Internally it uses account id of the service id for the identifier 'serviceid' and for the identifier 'crn' it uses account id contained in the CRN.
		* `type` - (Required, String) Type of the identity.
			* Constraints: Allowable values are: `user`, `serviceid`, `crn`.
* `policy_template_references` - (Optional, List) Existing policy templates that you can reference to assign access in the trusted profile component.
  Nested schema for **policy_template_references**:
	* `id` - (Required, String) ID of Access Policy Template.
	* `version` - (Required, String) Version of Access Policy Template.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the trusted_profile_template.
* `committed` - (Boolean) Committed flag determines if the template is ready for assignment.
* `profile` - (List) Input body parameters for the TemplateProfileComponent.
  Nested schema for **profile**:
	* `name` - (String) Name of the Profile.
	* `description` - (String) Description of the Profile.
	* `rules` - (List) Rules for the Profile.
	  Nested schema for **rules**:
		* `conditions` - (List) Conditions of this claim rule.
		  Nested schema for **conditions**:
			* `claim` - (String) The claim to evaluate against.
			* `operator` - (String) The operation to perform on the claim. valid values are EQUALS, NOT_EQUALS, EQUALS_IGNORE_CASE, NOT_EQUALS_IGNORE_CASE, CONTAINS, IN.
			* `value` - (String) The stringified JSON value that the claim is compared to using the operator.
		* `expiration` - (Integer) Session expiration in seconds, only required if type is 'Profile-SAML'.
		* `name` - (String) Name of the claim rule to be created or updated.
		* `realm_name` - (String) The realm name of the Idp this claim rule applies to. This field is required only if the type is specified as 'Profile-SAML'.
		* `type` - (String) Type of the claim rule.
	* `identities` - (Optional, List) Identities for the Profile.
	  Nested schema for **identities**:
		* `accounts` - (List) Only valid for the type user. Accounts from which a user can assume the trusted profile.
		* `description` - (String) Description of the identity that can assume the trusted profile. This is optional field for all the types of identities. When this field is not set for the identity type 'serviceid' then the description of the service id is used. Description is recommended for the identity type 'crn' E.g. 'Instance 1234 of IBM Cloud Service project'.
		* `iam_id` - (String) IAM ID of the identity.
		* `identifier` - (String) Identifier of the identity that can assume the trusted profiles. This can be a user identifier (IAM id), serviceid or crn. Internally it uses account id of the service id for the identifier 'serviceid' and for the identifier 'crn' it uses account id contained in the CRN.
		* `type` - (String) Type of the identity.
* `policy_template_references` - (List) Existing policy templates that you can reference to assign access in the trusted profile component.
  Nested schema for **policy_template_references**:
	* `id` - (String) ID of Access Policy Template.
	* `version` - (String) Version of Access Policy Template.
* `created_at` - (String) Timestamp of when the template was created.
* `created_by_id` - (String) IAMid of the creator.
* `crn` - (String) Cloud resource name.
* `entity_tag` - (String) Entity tag for this templateId-version combination.
* `history` - (List) History of the trusted profile template.
Nested schema for **history**:
	* `action` - (String) Action of the history entry.
	* `iam_id` - (String) IAM ID of the identity which triggered the action.
	* `iam_id_account` - (String) Account of the identity which triggered the action.
	* `message` - (String) Message which summarizes the executed action.
	* `params` - (List) Params of the history entry.
	* `timestamp` - (String) Timestamp when the action was triggered.
* `last_modified_at` - (String) Timestamp of when the template was last modified.
* `last_modified_by_id` - (String) IAMid of the identity that made the latest modification.

## Import

You can import the `ibm_iam_trusted_profile_template` resource by using `version`. Version of the the template.

### Syntax

```bash
$ terraform import ibm_iam_trusted_profile_template.trusted_profile_template $version
```
