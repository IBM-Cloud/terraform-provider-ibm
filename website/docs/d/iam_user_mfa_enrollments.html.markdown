---
layout: "ibm"
page_title: "IBM : ibm_iam_user_mfa_enrollments"
description: |-
  Get information about iam_user_mfa_enrollments
subcategory: "IAM Identity Services"
---

# ibm_iam_user_mfa_enrollments

Provides a read-only data source for iam_user_mfa_enrollments. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_user_mfa_enrollments" "iam_user_mfa_enrollments" {
	account_id = "account_id"
	iam_id = "iam_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `account_id` - (Required, Forces new resource, String) ID of the account.
* `iam_id` - (Required, String) iam_id of the user. This user must be the member of the account.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the iam_user_mfa_enrollments.
* `account_based_mfa` - (List) 
Nested scheme for **account_based_mfa**:
	* `complies` - (Boolean) The enrollment complies to the effective requirement.
	* `security_questions` - (List)
	Nested scheme for **security_questions**:
		* `enrolled` - (Boolean) Describes whether the enrollment type is enrolled.
		* `required` - (Boolean) Describes whether the enrollment type is required.
	* `totp` - (List)
	Nested scheme for **totp**:
		* `enrolled` - (Boolean) Describes whether the enrollment type is enrolled.
		* `required` - (Boolean) Describes whether the enrollment type is required.
	* `verisign` - (List)
	Nested scheme for **verisign**:
		* `enrolled` - (Boolean) Describes whether the enrollment type is enrolled.
		* `required` - (Boolean) Describes whether the enrollment type is required.

* `effective_mfa_type` - (String) currently effective mfa type i.e. id_based_mfa or account_based_mfa.

* `id_based_mfa` - (List) 
Nested scheme for **id_based_mfa**:
	* `complies` - (Boolean) The enrollment complies to the effective requirement.
	* `trait_account_default` - (String) Defines the MFA trait for the account. Valid values:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.
	  * Constraints: Allowable values are: `NONE`, `NONE_NO_ROPC`, `TOTP`, `TOTP4ALL`, `LEVEL1`, `LEVEL2`, `LEVEL3`.
	* `trait_effective` - (String) Defines the MFA trait for the account. Valid values:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.
	  * Constraints: Allowable values are: `NONE`, `NONE_NO_ROPC`, `TOTP`, `TOTP4ALL`, `LEVEL1`, `LEVEL2`, `LEVEL3`.
	* `trait_user_specific` - (String) Defines the MFA trait for the account. Valid values:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.
	  * Constraints: Allowable values are: `NONE`, `NONE_NO_ROPC`, `TOTP`, `TOTP4ALL`, `LEVEL1`, `LEVEL2`, `LEVEL3`.

