---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_connector_get_users"
description: |-
  Get information about backup_recovery_connector_get_users
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_connector_get_users

Provides a read-only data source to retrieve information about backup_recovery_connector_get_users. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_connector_get_users" "backup_recovery_connector_get_users" {
	session_name = ibm_backup_recovery_connector_update_user.backup_recovery_connector_update_user_instance.session_name
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `all_under_hierarchy` - (Optional, Boolean) AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned.TenantIds contains ids of the tenants for which objects are to be returned.
* `domain` - (Optional, String) Optionally specify a domain to filter by. If no domain is specified, all users on the Cohesity Cluster are searched. If a domain is specified, only users on the Cohesity Cluster associated with that domain are searched.
* `email_addresses` - (Optional, List) Optionally specify a list of email addresses to filter by.
* `partial_match` - (Optional, Boolean) Optionally specify whether to enable partial match. If set, all users with name containing Usernames will be returned. If set to false, only users with exact the same name as Usernames will be returned. By default this parameter is set to true.
`session_name` - (Required, Forces new resource, String) To be obtained from login API. Login is not yet supported in terraform. User needs to fetch this token manually by making a POST call to connector-url/login.
`curl --location --request POST 'https://150.240.36.117/login' --header 'Content-Type: application/json'  --data-raw '{ "username": "admin","password": "cohesitys7"}' -k -v`
* `tenant_ids` - (Optional, List) TenantIds contains ids of the tenants for which objects are to be returned.
* `usernames` - (Optional, List) Optionally specify a list of usernames to filter by. All users containing username will be returned.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_connector_get_users.
* `users` - (List) Specifies list of users.
Nested schema for **users**:
	* `ad_user_info` - (List) Specifies an AD User's information logged in using an active directory. This information is not stored on the Cluster.
	Nested schema for **ad_user_info**:
		* `group_sids` - (List) Specifies the SIDs of the groups.
		* `groups` - (List) Specifies the groups this user is a part of.
		* `is_floating_user` - (Boolean) Specifies whether this is a floating user or not.
	* `additional_group_names` - (List) Specifies the names of additional groups this User may belong to.
	* `allow_dso_modify` - (Boolean) Specifies if the data security user can be modified by the admin users.
	* `audit_log_settings` - (List) AuditLogSettings specifies struct with audt log configuration. Make these settings in such a way that zero values are cluster default when bb is not present.
	Nested schema for **audit_log_settings**:
		* `read_logging` - (Boolean) ReadLogging specifies whether read logs needs to be captured.
	* `authentication_type` - (String) Specifies the authentication type of the user. 'kAuthLocal' implies authenticated user is a local user. 'kAuthAd' implies authenticated user is an Active Directory user. 'kAuthSalesforce' implies authenticated user is a Salesforce user. 'kAuthGoogle' implies authenticated user is a Google user. 'kAuthSso' implies authenticated user is an SSO user.
	  * Constraints: Allowable values are: `kAuthLocal`, `kAuthAd`, `kAuthSalesforce`, `kAuthGoogle`, `kAuthSso`.
	* `cluster_identifiers` - (List) Specifies the list of clusters this user has access to. If this is not specified, access will be granted to all clusters.
	Nested schema for **cluster_identifiers**:
		* `cluster_id` - (Integer) Specifies the id of the cluster.
		* `cluster_incarnation_id` - (Integer) Specifies the incarnation id of the cluster.
	* `created_time_msecs` - (Integer) Specifies the epoch time in milliseconds when the user account was created on the Cohesity Cluster.
	* `current_password` - (String) Specifies the current password when updating the password.
	* `description` - (String) Specifies a description about the user.
	* `domain` - (String) Specifies the fully qualified domain name (FQDN) of an Active Directory or LOCAL for the default LOCAL domain on the Cohesity Cluster. A user is uniquely identified by combination of the username and the domain.
	* `effective_time_msecs` - (Integer) Specifies the epoch time in milliseconds when the user becomes effective. Until that time, the user cannot log in.
	* `email_address` - (String) Specifies the email address of the user.
	* `expired_time_msecs` - (Integer) Specifies the epoch time in milliseconds when the user becomes expired. After that, the user cannot log in.
	* `force_password_change` - (Boolean) Specifies whether to force user to change password.
	* `google_account` - (List) Google Account Information of a Helios BaaS user.
	Nested schema for **google_account**:
		* `account_id` - (String) Specifies the Account Id assigned by Google.
		* `user_id` - (String) Specifies the User Id assigned by Google.
	* `group_roles` - (List) Specifies the Cohesity roles to associate with the user' group. These roles can only be edited from group.
	* `id` - (String) The unique ID.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `idp_user_info` - (List) Specifies an IdP User's information logged in using an IdP. This information is not stored on the Cluster.
	Nested schema for **idp_user_info**:
		* `group_sids` - (List) Specifies the SIDs of the groups.
		* `groups` - (List) Specifies the Idp groups that the user is part of. As the user may not be registered on the cluster, we may have to capture the idp group membership. This way, if a group is created on the cluster later, users will instantly have access to tenantIds from that group as well.
		* `idp_id` - (Integer) Specifies the unique Id assigned by the Cluster for the IdP.
		* `is_floating_user` - (Boolean) Specifies whether or not this is a floating user.
		* `issuer_id` - (String) Specifies the unique identifier assigned by the vendor for this Cluster.
		* `user_id` - (String) Specifies the unique identifier assigned by the vendor for the user.
		* `vendor` - (String) Specifies the vendor providing the IdP service.
	* `intercom_messenger_token` - (String) Specifies the messenger token for intercom identity verification.
	* `is_account_locked` - (Boolean) Specifies whether the user account is locked.
	* `is_account_mfa_enabled` - (Boolean) Specifies if MFA is enabled for the Helios Account.
	* `is_active` - (Boolean) IsActive specifies whether or not a user is active, or has been disactivated by the customer. The default behavior is 'true'.
	* `is_cluster_mfa_enabled` - (Boolean) Specifies if MFA is enabled on cluster.
	* `last_successful_login_time_msecs` - (Integer) Specifies the epoch time in milliseconds when the user was last logged in successfully.
	* `last_updated_time_msecs` - (Integer) Specifies the epoch time in milliseconds when the user account was last modified on the Cohesity Cluster.
	* `mfa_info` - (List) Specifies information about MFA.
	Nested schema for **mfa_info**:
		* `is_email_otp_setup_done` - (Boolean) Specifies if email OTP setup is done on the user.
		* `is_totp_setup_done` - (Boolean) Specifies if TOTP setup is done on the user.
		* `is_user_exempt_from_mfa` - (Boolean) Specifies if MFA is disabled on the user.
	* `mfa_methods` - (List) Specifies MFA methods that enabled on the cluster.
	* `object_class` - (String) Specifies object class of user, could be either user or group.
	* `org_membership` - (List) OrgMembership contains the list of all available tenantIds for this user to switch to. Only when creating the session user, this field is populated on the fly. We discover the tenantIds from various groups assigned to the users.
	Nested schema for **org_membership**:
		* `bifrost_enabled` - (Boolean) Specifies if this tenant is bifrost enabled or not.
		* `is_managed_on_helios` - (Boolean) Specifies whether this tenant is manged on helios.
		* `name` - (String) Specifies name of the tenant.
		* `restricted` - (Boolean) Whether the user is a restricted user. A restricted user can only view the objects he has permissions to.
		* `roles` - (List) Specifies the Cohesity roles to associate with the user such as such as 'Admin', 'Ops' or 'View'. The Cohesity roles determine privileges on the Cohesity Cluster for this user.
		* `tenant_id` - (String) Specifies the unique id of the tenant.
	* `password` - (String) Specifies the password of this user.
	* `preferences` - (List) Specifies the preferences of this user.
	Nested schema for **preferences**:
		* `locale` - (String) Locale reflects the language settings of the user. Populate using the user preferences stored in Scribe for the user wherever needed.
	* `previous_login_time_msecs` - (Integer) Specifies the epoch time in milliseconds of previous user login.
	* `primary_group_name` - (String) Specifies the name of the primary group of this User.
	* `privilege_ids` - (List) Specifies the Cohesity privileges from the roles. This will be populated based on the union of all privileges in roles. Type for unique privilege Id values. All below enum values specify a value for all uniquely defined privileges in Cohesity.
	  * Constraints: Allowable list items are: `kPrincipalView`, `kPrincipalModify`, `kAppLaunch`, `kAppsManagement`, `kOrganizationView`, `kOrganizationModify`, `kOrganizationImpersonate`, `kCloneView`, `kCloneModify`, `kClusterView`, `kClusterModify`, `kClusterCreate`, `kClusterSupport`, `kClusterUpgrade`, `kClusterRemoteView`, `kClusterRemoteModify`, `kClusterExternalTargetView`, `kClusterExternalTargetModify`, `kClusterAudit`, `kAlertView`, `kAlertModify`, `kVlanView`, `kVlanModify`, `kHybridExtenderView`, `kHybridExtenderDownload`, `kAdLdapView`, `kAdLdapModify`, `kSchedulerView`, `kSchedulerModify`, `kProtectionView`, `kProtectionModify`, `kProtectionJobOperate`, `kProtectionSourceModify`, `kProtectionPolicyView`, `kProtectionPolicyModify`, `kRestoreView`, `kRestoreModify`, `kRestoreDownload`, `kRemoteRestore`, `kStorageView`, `kStorageModify`, `kStorageDomainView`, `kStorageDomainModify`, `kAnalyticsView`, `kAnalyticsModify`, `kReportsView`, `kMcmModify`, `kDataSecurity`, `kSmbBackup`, `kSmbRestore`, `kSmbTakeOwnership`, `kSmbAuditing`, `kMcmUnregister`, `kMcmUpgrade`, `kMcmModifySuperAdmin`, `kMcmViewSuperAdmin`, `kMcmModifyCohesityAdmin`, `kMcmViewCohesityAdmin`, `kObjectSearch`, `kFileDatalockExpiryTimeDecrease`.
	* `profiles` - (List) Specifies the user profiles. NOTE:- Currently used for Helios.
	Nested schema for **profiles**:
		* `cluster_identifiers` - (List) Specifies the list of clusters. This is only valid if tenant type is OnPrem.
		Nested schema for **cluster_identifiers**:
			* `cluster_id` - (Integer) Specifies the id of the cluster.
			* `cluster_incarnation_id` - (Integer) Specifies the incarnation id of the cluster.
		* `is_active` - (Boolean) Specifies whether or not the tenant is active.
		* `is_deleted` - (Boolean) Specifies whether or not the tenant is deleted.
		* `region_ids` - (List) Specifies the list of regions. This is only valid if tenant type is Dmaas.
		* `tenant_id` - (String) Specifies the tenant id.
		* `tenant_name` - (String) Specifies the tenant id.
		* `tenant_type` - (String) Specifies the MCM tenant type. 'Dmaas' implies tenant type is DMaaS. 'Mcm' implies tenant is Mcm Cluster tenant.
		  * Constraints: Allowable values are: `Dmaas`, `Mcm`.
	* `restricted` - (Boolean) Whether the user is a restricted user. A restricted user can only view the objects he has permissions to.
	* `roles` - (List) Specifies the Cohesity roles to associate with the user such as such as 'Admin', 'Ops' or 'View'. The Cohesity roles determine privileges on the Cohesity Cluster for this user.
	* `s3_access_key_id` - (String) Specifies the S3 Account Access Key ID.
	* `s3_account_id` - (String) Specifies the S3 Account Canonical User ID.
	* `s3_secret_key` - (String) Specifies the S3 Account Secret Key.
	* `salesforce_account` - (List) Salesforce Account Information of a Helios user.
	Nested schema for **salesforce_account**:
		* `account_id` - (String) Specifies the Account Id assigned by Salesforce.
		* `helios_access_grant_status` - (String) Specifies the status of helios access.
		* `is_d_gaa_s_user` - (Boolean) Specifies whether user is a DGaaS licensed user.
		* `is_d_maa_s_user` - (Boolean) Specifies whether user is a DMaaS licensed user.
		* `is_d_raa_s_user` - (Boolean) Specifies whether user is a DRaaS licensed user.
		* `is_r_paa_s_user` - (Boolean) Specifies whether user is a RPaaS licensed user.
		* `is_sales_user` - (Boolean) Specifies whether user is a Sales person from Cohesity.
		* `is_support_user` - (Boolean) Specifies whether user is a support person from Cohesity.
		* `user_id` - (String) Specifies the User Id assigned by Salesforce.
	* `sid` - (String) Specifies the unique Security ID (SID) of the user. This field is mandatory in modifying user.
	* `spog_context` - (List) SpogContext specifies all of the information about the user and cluster which is performing action on this cluster.
	Nested schema for **spog_context**:
		* `primary_cluster_id` - (Integer) Specifies the ID of the remote cluster which is accessing this cluster via SPOG.
		* `primary_cluster_user_sid` - (String) Specifies the SID of the user who is accessing this cluster via SPOG.
		* `primary_cluster_username` - (String) Specifies the username of the user who is accessing this cluster via SPOG.
	* `subscription_info` - (List) Extends this to have Helios, DRaaS and DSaaS.
	Nested schema for **subscription_info**:
		* `classification` - (List) ClassificationInfo holds information about the Datahawk Classification subscription such as if it is active or not.
		Nested schema for **classification**:
			* `end_date` - (String) Specifies the end date of the subscription.
			* `is_active` - (Boolean) Specifies the end date of the subscription.
			* `is_free_trial` - (Boolean) Specifies the end date of the subscription.
			* `start_date` - (String) Specifies the start date of the subscription.
		* `data_protect` - (List) DMaaSSubscriptionInfo holds information about the Data Protect subscription such as if it is active or not.
		Nested schema for **data_protect**:
			* `end_date` - (String) Specifies the end date of the subscription.
			* `is_active` - (Boolean) Specifies the end date of the subscription.
			* `is_aws_subscription` - (Boolean) Specifies whether the subscription is AWS Subscription.
			* `is_cohesity_subscription` - (Boolean) Specifies whether the subscription is a Cohesity Paid subscription.
			* `is_free_trial` - (Boolean) Specifies the end date of the subscription.
			* `quantity` - (Integer) Specifies the quantity of the subscription.
			* `start_date` - (String) Specifies the start date of the subscription.
			* `tiering` - (List) Specifies the tiering info.
			Nested schema for **tiering**:
				* `backend_tiering` - (Boolean) Specifies whether back-end tiering is enabled.
				* `frontend_tiering` - (Boolean) Specifies whether Front End Tiering Enabled.
				* `max_retention` - (Integer) Specified the max retention for backup policy creation.
		* `data_protect_azure` - (List) ClassificationInfo holds information about the Datahawk Classification subscription such as if it is active or not.
		Nested schema for **data_protect_azure**:
			* `end_date` - (String) Specifies the end date of the subscription.
			* `is_active` - (Boolean) Specifies the end date of the subscription.
			* `is_free_trial` - (Boolean) Specifies the end date of the subscription.
			* `quantity` - (Integer) Specifies the quantity of the subscription.
			* `start_date` - (String) Specifies the start date of the subscription.
			* `tiering` - (List) Specifies the tiering info.
			Nested schema for **tiering**:
				* `backend_tiering` - (Boolean) Specifies whether back-end tiering is enabled.
				* `frontend_tiering` - (Boolean) Specifies whether Front End Tiering Enabled.
				* `max_retention` - (Integer) Specified the max retention for backup policy creation.
		* `fort_knox_azure_cool` - (List) FortKnoxInfo holds information about the Fortknox Azure or Azure FreeTrial or AwsCold or AwsCold FreeTrial subscription such as if it is active.
		Nested schema for **fort_knox_azure_cool**:
			* `end_date` - (String) Specifies the end date of the subscription.
			* `is_active` - (Boolean) Specifies the end date of the subscription.
			* `is_free_trial` - (Boolean) Specifies the end date of the subscription.
			* `quantity` - (Integer) Specifies the quantity of the subscription.
			* `start_date` - (String) Specifies the start date of the subscription.
		* `fort_knox_azure_hot` - (List) FortKnoxInfo holds information about the Fortknox Azure or Azure FreeTrial or AwsCold or AwsCold FreeTrial subscription such as if it is active.
		Nested schema for **fort_knox_azure_hot**:
			* `end_date` - (String) Specifies the end date of the subscription.
			* `is_active` - (Boolean) Specifies the end date of the subscription.
			* `is_free_trial` - (Boolean) Specifies the end date of the subscription.
			* `quantity` - (Integer) Specifies the quantity of the subscription.
			* `start_date` - (String) Specifies the start date of the subscription.
		* `fort_knox_cold` - (List) FortKnoxInfo holds information about the Fortknox Azure or Azure FreeTrial or AwsCold or AwsCold FreeTrial subscription such as if it is active.
		Nested schema for **fort_knox_cold**:
			* `end_date` - (String) Specifies the end date of the subscription.
			* `is_active` - (Boolean) Specifies the end date of the subscription.
			* `is_free_trial` - (Boolean) Specifies the end date of the subscription.
			* `quantity` - (Integer) Specifies the quantity of the subscription.
			* `start_date` - (String) Specifies the start date of the subscription.
		* `ransomware` - (List) RansomwareInfo holds information about the FortKnox/FortKnoxFreeTrial subscription such as if it is active or not.
		Nested schema for **ransomware**:
			* `end_date` - (String) Specifies the end date of the subscription.
			* `is_active` - (Boolean) Specifies the end date of the subscription.
			* `is_free_trial` - (Boolean) Specifies the end date of the subscription.
			* `quantity` - (Integer) Specifies the quantity of the subscription.
			* `start_date` - (String) Specifies the start date of the subscription.
		* `site_continuity` - (List) SiteContinuityInfo holds information about the Site Continuity subscription such as if it is active or not.
		Nested schema for **site_continuity**:
			* `end_date` - (String) Specifies the end date of the subscription.
			* `is_active` - (Boolean) Specifies the end date of the subscription.
			* `is_free_trial` - (Boolean) Specifies the end date of the subscription.
			* `start_date` - (String) Specifies the start date of the subscription.
		* `threat_protection` - (List) ThreatProtectionInfo holds information about the Datahawk ThreatProtection subscription such as if it is active or not.
		Nested schema for **threat_protection**:
			* `end_date` - (String) Specifies the end date of the subscription.
			* `is_active` - (Boolean) Specifies the end date of the subscription.
			* `is_free_trial` - (Boolean) Specifies the end date of the subscription.
			* `start_date` - (String) Specifies the start date of the subscription.
	* `tenant_accesses` - (List) Specfies the Tenant Access for MCM User.
	Nested schema for **tenant_accesses**:
		* `cluster_identifiers` - (List) Specifies the list of clusters.
		Nested schema for **cluster_identifiers**:
			* `cluster_id` - (Integer) Specifies the id of the cluster.
			* `cluster_incarnation_id` - (Integer) Specifies the incarnation id of the cluster.
		* `created_time_msecs` - (Integer) Specifies the epoch time in milliseconds when the tenant access was created.
		* `effective_time_msecs` - (Integer) Specifies the epoch time in milliseconds when the tenant access becomes effective. Until that time, the user cannot log in.
		* `expired_time_msecs` - (Integer) Specifies the epoch time in milliseconds when the tenant access becomes expired. After that, the user cannot log in.
		* `is_access_active` - (Boolean) IsAccessActive specifies whether or not a tenant access is active, or has been deactivated by the customer. The default behavior is 'true'.
		* `is_active` - (Boolean) Specifies whether or not the tenant is active.
		* `is_deleted` - (Boolean) Specifies whether or not the tenant is deleted.
		* `last_updated_time_msecs` - (Integer) Specifies the epoch time in milliseconds when the tenant access was last modified.
		* `roles` - (List) Specifies the Cohesity roles to associate with the user such as such as 'Admin', 'Ops' or 'View'.
		  * Constraints: The minimum length is `1` item.
		* `tenant_id` - (String) Specifies the tenant id.
		* `tenant_name` - (String) Specifies the tenant name.
		* `tenant_type` - (String) Specifies the MCM tenant type. 'Dmaas' implies tenant type is DMaaS. 'Mcm' implies tenant is Mcm Cluster tenant.
		  * Constraints: Allowable values are: `Dmaas`, `Mcm`.
	* `tenant_id` - (String) Specifies the effective Tenant ID of the user.
	* `username` - (String) Specifies the login name of the user.

