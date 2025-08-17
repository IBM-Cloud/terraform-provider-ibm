---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_connector_update_user"
description: |-
  Manages backup_recovery_connector_update_user.
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_connector_update_user

Create backup_recovery_connector_update_users with this resource.

**Note**
ibm_backup_recovery_connector_update_user resource does not support update or delete operations due to the absence of corresponding API endpoints. As a result, Terraform cannot manage these operations for those resources. Users should be aware that removing these resources from the configuration (main.tf) will only remove them from the Terraform state and will not affect the actual resources in the backend. Similarly updating these resources will throw an error in the plan phase stating that the resource cannot be updated.


## Example Usage

```hcl
resource "ibm_backup_recovery_connector_update_user" "backup_recovery_connector_update_user_instance" {
  ad_user_info {
		group_sids = [ "groupSids" ]
		groups = [ "groups" ]
		is_floating_user = true
  }
  audit_log_settings {
		read_logging = true
  }
  cluster_identifiers {
		cluster_id = 1
		cluster_incarnation_id = 1
  }
  google_account {
		account_id = "account_id"
		user_id = "user_id"
  }
  idp_user_info {
		group_sids = [ "groupSids" ]
		groups = [ "groups" ]
		idp_id = 1
		is_floating_user = true
		issuer_id = "issuer_id"
		user_id = "user_id"
		vendor = "vendor"
  }
  mfa_info {
		is_email_otp_setup_done = true
		is_totp_setup_done = true
		is_user_exempt_from_mfa = true
  }
  org_membership {
		bifrost_enabled = true
		is_managed_on_helios = true
		name = "name"
		restricted = true
		roles = [ "roles" ]
		tenant_id = "tenant_id"
  }
  preferences {
		locale = "locale"
  }
  profiles {
		cluster_identifiers {
			cluster_id = 1
			cluster_incarnation_id = 1
		}
		is_active = true
		is_deleted = true
		region_ids = [ "regionIds" ]
		tenant_id = "tenant_id"
		tenant_name = "tenant_name"
		tenant_type = "Dmaas"
  }
  salesforce_account {
		account_id = "account_id"
		helios_access_grant_status = "helios_access_grant_status"
		is_d_gaa_s_user = true
		is_d_maa_s_user = true
		is_d_raa_s_user = true
		is_r_paa_s_user = true
		is_sales_user = true
		is_support_user = true
		user_id = "user_id"
  }
  session_name = "MTczNjc0NzY1OHxEWDhFQVFMX2dBQUJFQUVRQUFELUFZWF9nQUFKQm5OMGNtbHVad3dLQUFoMWMyVnlibUZ0WlFaemRISnBibWNNQndBRllXUnRhVzRHYzNSeWFXNW5EQWNBQlhKdmJHVnpCbk4wY21sdVp3d1FBQTVEVDBoRlUwbFVXVjlCUkUxSlRnWnpkSEpwYm1jTUN3QUpjMmxrY3kxb1lYTm9Cbk4wY21sdVp3d3RBQ3RTYVV4ZmFqQmZOVGxxZFZJeWVIVlZhREJ2UVZGNlUxcEhTVWc1TlZVdFlVWTBjV1JNUjNaTk9VUTBCbk4wY21sdVp3d01BQXBwYmkxamJIVnpkR1Z5QkdKdmIyd0NBZ0FCQm5OMGNtbHVad3dMQUFsaGRYUm9MWFI1Y0dVR2MzUnlhVzVuREFNQUFURUdjM1J5YVc1bkRCRUFEMlY0Y0dseVlYUnBiMjR0ZEdsdFpRWnpkSEpwYm1jTURBQUtNVGN6Tmpnek5EQTFPQVp6ZEhKcGJtY01DZ0FJZFhObGNpMXphV1FHYzNSeWFXNW5EQ0FBSGxNdE1TMHhNREF0TWpFdE16YzRNVFkyTXpVdE1qUXhPRFk1TXpVdE1RWnpkSEpwYm1jTUNBQUdaRzl0WVdsdUJuTjBjbWx1Wnd3SEFBVk1UME5CVEFaemRISnBibWNNQ0FBR2JHOWpZV3hsQm5OMGNtbHVad3dIQUFWbGJpMTFjdz09fGXFZlPU_3Nl46_gPKAw619qs6Pl7PX453Y_lf5BvBBo"
  spog_context {
		primary_cluster_id = 1
		primary_cluster_user_sid = "primary_cluster_user_sid"
		primary_cluster_username = "primary_cluster_username"
  }
  subscription_info {
		classification {
			end_date = "end_date"
			is_active = true
			is_free_trial = true
			start_date = "start_date"
		}
		data_protect {
			end_date = "end_date"
			is_active = true
			is_free_trial = true
			is_aws_subscription = true
			is_cohesity_subscription = true
			quantity = 1
			start_date = "start_date"
			tiering {
				backend_tiering = true
				frontend_tiering = true
				max_retention = 1
			}
		}
		data_protect_azure {
			end_date = "end_date"
			is_active = true
			is_free_trial = true
			quantity = 1
			start_date = "start_date"
			tiering {
				backend_tiering = true
				frontend_tiering = true
				max_retention = 1
			}
		}
		fort_knox_azure_cool {
			end_date = "end_date"
			is_active = true
			is_free_trial = true
			quantity = 1
			start_date = "start_date"
		}
		fort_knox_azure_hot {
			end_date = "end_date"
			is_active = true
			is_free_trial = true
			quantity = 1
			start_date = "start_date"
		}
		fort_knox_cold {
			end_date = "end_date"
			is_active = true
			is_free_trial = true
			quantity = 1
			start_date = "start_date"
		}
		ransomware {
			end_date = "end_date"
			is_active = true
			is_free_trial = true
			quantity = 1
			start_date = "start_date"
		}
		site_continuity {
			end_date = "end_date"
			is_active = true
			is_free_trial = true
			start_date = "start_date"
		}
		threat_protection {
			end_date = "end_date"
			is_active = true
			is_free_trial = true
			start_date = "start_date"
		}
  }
  tenant_accesses {
		cluster_identifiers {
			cluster_id = 1
			cluster_incarnation_id = 1
		}
		created_time_msecs = 1
		effective_time_msecs = 1
		expired_time_msecs = 1
		is_access_active = true
		is_active = true
		is_deleted = true
		last_updated_time_msecs = 1
		roles = [ "roles" ]
		tenant_id = "tenant_id"
		tenant_name = "tenant_name"
		tenant_type = "Dmaas"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `ad_user_info` - (Optional, Forces new resource, List) Specifies an AD User's information logged in using an active directory. This information is not stored on the Cluster.
Nested schema for **ad_user_info**:
	* `group_sids` - (Optional, List) Specifies the SIDs of the groups.
	* `groups` - (Optional, List) Specifies the groups this user is a part of.
	* `is_floating_user` - (Optional, Boolean) Specifies whether this is a floating user or not.
* `additional_group_names` - (Optional, Forces new resource, List) Specifies the names of additional groups this User may belong to.
* `allow_dso_modify` - (Optional, Forces new resource, Boolean) Specifies if the data security user can be modified by the admin users.
* `audit_log_settings` - (Optional, Forces new resource, List) AuditLogSettings specifies struct with audt log configuration. Make these settings in such a way that zero values are cluster default when bb is not present.
Nested schema for **audit_log_settings**:
	* `read_logging` - (Optional, Boolean) ReadLogging specifies whether read logs needs to be captured.
* `authentication_type` - (Optional, Forces new resource, String) Specifies the authentication type of the user. 'kAuthLocal' implies authenticated user is a local user. 'kAuthAd' implies authenticated user is an Active Directory user. 'kAuthSalesforce' implies authenticated user is a Salesforce user. 'kAuthGoogle' implies authenticated user is a Google user. 'kAuthSso' implies authenticated user is an SSO user.
  * Constraints: Allowable values are: `kAuthLocal`, `kAuthAd`, `kAuthSalesforce`, `kAuthGoogle`, `kAuthSso`.
* `cluster_identifiers` - (Optional, Forces new resource, List) Specifies the list of clusters this user has access to. If this is not specified, access will be granted to all clusters.
Nested schema for **cluster_identifiers**:
	* `cluster_id` - (Optional, Integer) Specifies the id of the cluster.
	* `cluster_incarnation_id` - (Optional, Integer) Specifies the incarnation id of the cluster.
* `created_time_msecs` - (Optional, Forces new resource, Integer) Specifies the epoch time in milliseconds when the user account was created on the Cohesity Cluster.
* `current_password` - (Optional, Forces new resource, String) Specifies the current password when updating the password.
* `description` - (Optional, Forces new resource, String) Specifies a description about the user.
* `domain` - (Optional, Forces new resource, String) Specifies the fully qualified domain name (FQDN) of an Active Directory or LOCAL for the default LOCAL domain on the Cohesity Cluster. A user is uniquely identified by combination of the username and the domain.
* `effective_time_msecs` - (Optional, Forces new resource, Integer) Specifies the epoch time in milliseconds when the user becomes effective. Until that time, the user cannot log in.
* `email_address` - (Optional, Forces new resource, String) Specifies the email address of the user.
* `expired_time_msecs` - (Optional, Forces new resource, Integer) Specifies the epoch time in milliseconds when the user becomes expired. After that, the user cannot log in.
* `force_password_change` - (Optional, Forces new resource, Boolean) Specifies whether to force user to change password.
* `google_account` - (Optional, Forces new resource, List) Google Account Information of a Helios BaaS user.
Nested schema for **google_account**:
	* `account_id` - (Optional, String) Specifies the Account Id assigned by Google.
	* `user_id` - (Optional, String) Specifies the User Id assigned by Google.
* `idp_user_info` - (Optional, Forces new resource, List) Specifies an IdP User's information logged in using an IdP. This information is not stored on the Cluster.
Nested schema for **idp_user_info**:
	* `group_sids` - (Optional, List) Specifies the SIDs of the groups.
	* `groups` - (Optional, List) Specifies the Idp groups that the user is part of. As the user may not be registered on the cluster, we may have to capture the idp group membership. This way, if a group is created on the cluster later, users will instantly have access to tenantIds from that group as well.
	* `idp_id` - (Optional, Integer) Specifies the unique Id assigned by the Cluster for the IdP.
	* `is_floating_user` - (Optional, Boolean) Specifies whether or not this is a floating user.
	* `issuer_id` - (Optional, String) Specifies the unique identifier assigned by the vendor for this Cluster.
	* `user_id` - (Optional, String) Specifies the unique identifier assigned by the vendor for the user.
	* `vendor` - (Optional, String) Specifies the vendor providing the IdP service.
* `intercom_messenger_token` - (Optional, Forces new resource, String) Specifies the messenger token for intercom identity verification.
* `is_account_locked` - (Optional, Forces new resource, Boolean) Specifies whether the user account is locked.
* `is_active` - (Optional, Forces new resource, Boolean) IsActive specifies whether or not a user is active, or has been disactivated by the customer. The default behavior is 'true'.
* `last_successful_login_time_msecs` - (Optional, Forces new resource, Integer) Specifies the epoch time in milliseconds when the user was last logged in successfully.
* `last_updated_time_msecs` - (Optional, Forces new resource, Integer) Specifies the epoch time in milliseconds when the user account was last modified on the Cohesity Cluster.
* `mfa_info` - (Optional, Forces new resource, List) Specifies information about MFA.
Nested schema for **mfa_info**:
	* `is_email_otp_setup_done` - (Computed, Boolean) Specifies if email OTP setup is done on the user.
	* `is_totp_setup_done` - (Computed, Boolean) Specifies if TOTP setup is done on the user.
	* `is_user_exempt_from_mfa` - (Optional, Boolean) Specifies if MFA is disabled on the user.
* `mfa_methods` - (Optional, Forces new resource, List) Specifies MFA methods that enabled on the cluster.
* `object_class` - (Optional, Forces new resource, String) Specifies object class of user, could be either user or group.
* `org_membership` - (Optional, Forces new resource, List) OrgMembership contains the list of all available tenantIds for this user to switch to. Only when creating the session user, this field is populated on the fly. We discover the tenantIds from various groups assigned to the users.
Nested schema for **org_membership**:
	* `bifrost_enabled` - (Optional, Boolean) Specifies if this tenant is bifrost enabled or not.
	* `is_managed_on_helios` - (Optional, Boolean) Specifies whether this tenant is manged on helios.
	* `name` - (Optional, String) Specifies name of the tenant.
	* `restricted` - (Optional, Boolean) Whether the user is a restricted user. A restricted user can only view the objects he has permissions to.
	* `roles` - (Optional, List) Specifies the Cohesity roles to associate with the user such as such as 'Admin', 'Ops' or 'View'. The Cohesity roles determine privileges on the Cohesity Cluster for this user.
	* `tenant_id` - (Optional, String) Specifies the unique id of the tenant.
* `password` - (Optional, Forces new resource, String) Specifies the password of this user.
* `preferences` - (Optional, Forces new resource, List) Specifies the preferences of this user.
Nested schema for **preferences**:
	* `locale` - (Optional, String) Locale reflects the language settings of the user. Populate using the user preferences stored in Scribe for the user wherever needed.
* `previous_login_time_msecs` - (Optional, Forces new resource, Integer) Specifies the epoch time in milliseconds of previous user login.
* `primary_group_name` - (Optional, Forces new resource, String) Specifies the name of the primary group of this User.
* `privilege_ids` - (Optional, Forces new resource, List) Specifies the Cohesity privileges from the roles. This will be populated based on the union of all privileges in roles. Type for unique privilege Id values. All below enum values specify a value for all uniquely defined privileges in Cohesity.
  * Constraints: Allowable list items are: `kPrincipalView`, `kPrincipalModify`, `kAppLaunch`, `kAppsManagement`, `kOrganizationView`, `kOrganizationModify`, `kOrganizationImpersonate`, `kCloneView`, `kCloneModify`, `kClusterView`, `kClusterModify`, `kClusterCreate`, `kClusterSupport`, `kClusterUpgrade`, `kClusterRemoteView`, `kClusterRemoteModify`, `kClusterExternalTargetView`, `kClusterExternalTargetModify`, `kClusterAudit`, `kAlertView`, `kAlertModify`, `kVlanView`, `kVlanModify`, `kHybridExtenderView`, `kHybridExtenderDownload`, `kAdLdapView`, `kAdLdapModify`, `kSchedulerView`, `kSchedulerModify`, `kProtectionView`, `kProtectionModify`, `kProtectionJobOperate`, `kProtectionSourceModify`, `kProtectionPolicyView`, `kProtectionPolicyModify`, `kRestoreView`, `kRestoreModify`, `kRestoreDownload`, `kRemoteRestore`, `kStorageView`, `kStorageModify`, `kStorageDomainView`, `kStorageDomainModify`, `kAnalyticsView`, `kAnalyticsModify`, `kReportsView`, `kMcmModify`, `kDataSecurity`, `kSmbBackup`, `kSmbRestore`, `kSmbTakeOwnership`, `kSmbAuditing`, `kMcmUnregister`, `kMcmUpgrade`, `kMcmModifySuperAdmin`, `kMcmViewSuperAdmin`, `kMcmModifyCohesityAdmin`, `kMcmViewCohesityAdmin`, `kObjectSearch`, `kFileDatalockExpiryTimeDecrease`.
* `profiles` - (Optional, Forces new resource, List) Specifies the user profiles. NOTE:- Currently used for Helios.
Nested schema for **profiles**:
	* `cluster_identifiers` - (Optional, List) Specifies the list of clusters. This is only valid if tenant type is OnPrem.
	Nested schema for **cluster_identifiers**:
		* `cluster_id` - (Optional, Integer) Specifies the id of the cluster.
		* `cluster_incarnation_id` - (Optional, Integer) Specifies the incarnation id of the cluster.
	* `is_active` - (Optional, Boolean) Specifies whether or not the tenant is active.
	* `is_deleted` - (Optional, Boolean) Specifies whether or not the tenant is deleted.
	* `region_ids` - (Optional, List) Specifies the list of regions. This is only valid if tenant type is Dmaas.
	* `tenant_id` - (Optional, String) Specifies the tenant id.
	* `tenant_name` - (Optional, String) Specifies the tenant id.
	* `tenant_type` - (Optional, String) Specifies the MCM tenant type. 'Dmaas' implies tenant type is DMaaS. 'Mcm' implies tenant is Mcm Cluster tenant.
	  * Constraints: Allowable values are: `Dmaas`, `Mcm`.
* `restricted` - (Optional, Forces new resource, Boolean) Whether the user is a restricted user. A restricted user can only view the objects he has permissions to.
* `roles` - (Optional, Forces new resource, List) Specifies the Cohesity roles to associate with the user such as such as 'Admin', 'Ops' or 'View'. The Cohesity roles determine privileges on the Cohesity Cluster for this user.
* `s3_access_key_id` - (Optional, Forces new resource, String) Specifies the S3 Account Access Key ID.
* `s3_account_id` - (Optional, Forces new resource, String) Specifies the S3 Account Canonical User ID.
* `s3_secret_key` - (Optional, Forces new resource, String) Specifies the S3 Account Secret Key.
* `salesforce_account` - (Optional, Forces new resource, List) Salesforce Account Information of a Helios user.
Nested schema for **salesforce_account**:
	* `account_id` - (Optional, String) Specifies the Account Id assigned by Salesforce.
	* `helios_access_grant_status` - (Optional, String) Specifies the status of helios access.
	* `is_d_gaa_s_user` - (Optional, Boolean) Specifies whether user is a DGaaS licensed user.
	* `is_d_maa_s_user` - (Optional, Boolean) Specifies whether user is a DMaaS licensed user.
	* `is_d_raa_s_user` - (Optional, Boolean) Specifies whether user is a DRaaS licensed user.
	* `is_r_paa_s_user` - (Optional, Boolean) Specifies whether user is a RPaaS licensed user.
	* `is_sales_user` - (Optional, Boolean) Specifies whether user is a Sales person from Cohesity.
	* `is_support_user` - (Optional, Boolean) Specifies whether user is a support person from Cohesity.
	* `user_id` - (Optional, String) Specifies the User Id assigned by Salesforce.
* `session_name` - (Required, Forces new resource, String) To be obtained from login API. Login is not yet supported in terraform. User needs to fetch this token manually by making a POST call to connector-url/login.
`curl --location --request POST 'https://150.240.36.117/login' --header 'Content-Type: application/json'  --data-raw '{ "username": "admin","password": "cohesitys7"}' -k -v`
* `sid` - (Optional, Forces new resource, String) Specifies the unique Security ID (SID) of the user. This field is mandatory in modifying user.
* `spog_context` - (Optional, Forces new resource, List) SpogContext specifies all of the information about the user and cluster which is performing action on this cluster.
Nested schema for **spog_context**:
	* `primary_cluster_id` - (Optional, Integer) Specifies the ID of the remote cluster which is accessing this cluster via SPOG.
	* `primary_cluster_user_sid` - (Optional, String) Specifies the SID of the user who is accessing this cluster via SPOG.
	* `primary_cluster_username` - (Optional, String) Specifies the username of the user who is accessing this cluster via SPOG.
* `subscription_info` - (Optional, Forces new resource, List) Extends this to have Helios, DRaaS and DSaaS.
Nested schema for **subscription_info**:
	* `classification` - (Optional, List) ClassificationInfo holds information about the Datahawk Classification subscription such as if it is active or not.
	Nested schema for **classification**:
		* `end_date` - (Optional, String) Specifies the end date of the subscription.
		* `is_active` - (Optional, Boolean) Specifies the end date of the subscription.
		* `is_free_trial` - (Optional, Boolean) Specifies the end date of the subscription.
		* `start_date` - (Optional, String) Specifies the start date of the subscription.
	* `data_protect` - (Optional, List) DMaaSSubscriptionInfo holds information about the Data Protect subscription such as if it is active or not.
	Nested schema for **data_protect**:
		* `end_date` - (Optional, String) Specifies the end date of the subscription.
		* `is_active` - (Optional, Boolean) Specifies the end date of the subscription.
		* `is_aws_subscription` - (Optional, Boolean) Specifies whether the subscription is AWS Subscription.
		* `is_cohesity_subscription` - (Optional, Boolean) Specifies whether the subscription is a Cohesity Paid subscription.
		* `is_free_trial` - (Optional, Boolean) Specifies the end date of the subscription.
		* `quantity` - (Optional, Integer) Specifies the quantity of the subscription.
		* `start_date` - (Optional, String) Specifies the start date of the subscription.
		* `tiering` - (Optional, List) Specifies the tiering info.
		Nested schema for **tiering**:
			* `backend_tiering` - (Optional, Boolean) Specifies whether back-end tiering is enabled.
			* `frontend_tiering` - (Optional, Boolean) Specifies whether Front End Tiering Enabled.
			* `max_retention` - (Optional, Integer) Specified the max retention for backup policy creation.
	* `data_protect_azure` - (Optional, List) ClassificationInfo holds information about the Datahawk Classification subscription such as if it is active or not.
	Nested schema for **data_protect_azure**:
		* `end_date` - (Optional, String) Specifies the end date of the subscription.
		* `is_active` - (Optional, Boolean) Specifies the end date of the subscription.
		* `is_free_trial` - (Optional, Boolean) Specifies the end date of the subscription.
		* `quantity` - (Optional, Integer) Specifies the quantity of the subscription.
		* `start_date` - (Optional, String) Specifies the start date of the subscription.
		* `tiering` - (Optional, List) Specifies the tiering info.
		Nested schema for **tiering**:
			* `backend_tiering` - (Optional, Boolean) Specifies whether back-end tiering is enabled.
			* `frontend_tiering` - (Optional, Boolean) Specifies whether Front End Tiering Enabled.
			* `max_retention` - (Optional, Integer) Specified the max retention for backup policy creation.
	* `fort_knox_azure_cool` - (Optional, List) FortKnoxInfo holds information about the Fortknox Azure or Azure FreeTrial or AwsCold or AwsCold FreeTrial subscription such as if it is active.
	Nested schema for **fort_knox_azure_cool**:
		* `end_date` - (Optional, String) Specifies the end date of the subscription.
		* `is_active` - (Optional, Boolean) Specifies the end date of the subscription.
		* `is_free_trial` - (Optional, Boolean) Specifies the end date of the subscription.
		* `quantity` - (Optional, Integer) Specifies the quantity of the subscription.
		* `start_date` - (Optional, String) Specifies the start date of the subscription.
	* `fort_knox_azure_hot` - (Optional, List) FortKnoxInfo holds information about the Fortknox Azure or Azure FreeTrial or AwsCold or AwsCold FreeTrial subscription such as if it is active.
	Nested schema for **fort_knox_azure_hot**:
		* `end_date` - (Optional, String) Specifies the end date of the subscription.
		* `is_active` - (Optional, Boolean) Specifies the end date of the subscription.
		* `is_free_trial` - (Optional, Boolean) Specifies the end date of the subscription.
		* `quantity` - (Optional, Integer) Specifies the quantity of the subscription.
		* `start_date` - (Optional, String) Specifies the start date of the subscription.
	* `fort_knox_cold` - (Optional, List) FortKnoxInfo holds information about the Fortknox Azure or Azure FreeTrial or AwsCold or AwsCold FreeTrial subscription such as if it is active.
	Nested schema for **fort_knox_cold**:
		* `end_date` - (Optional, String) Specifies the end date of the subscription.
		* `is_active` - (Optional, Boolean) Specifies the end date of the subscription.
		* `is_free_trial` - (Optional, Boolean) Specifies the end date of the subscription.
		* `quantity` - (Optional, Integer) Specifies the quantity of the subscription.
		* `start_date` - (Optional, String) Specifies the start date of the subscription.
	* `ransomware` - (Optional, List) RansomwareInfo holds information about the FortKnox/FortKnoxFreeTrial subscription such as if it is active or not.
	Nested schema for **ransomware**:
		* `end_date` - (Optional, String) Specifies the end date of the subscription.
		* `is_active` - (Optional, Boolean) Specifies the end date of the subscription.
		* `is_free_trial` - (Optional, Boolean) Specifies the end date of the subscription.
		* `quantity` - (Optional, Integer) Specifies the quantity of the subscription.
		* `start_date` - (Optional, String) Specifies the start date of the subscription.
	* `site_continuity` - (Optional, List) SiteContinuityInfo holds information about the Site Continuity subscription such as if it is active or not.
	Nested schema for **site_continuity**:
		* `end_date` - (Optional, String) Specifies the end date of the subscription.
		* `is_active` - (Optional, Boolean) Specifies the end date of the subscription.
		* `is_free_trial` - (Optional, Boolean) Specifies the end date of the subscription.
		* `start_date` - (Optional, String) Specifies the start date of the subscription.
	* `threat_protection` - (Optional, List) ThreatProtectionInfo holds information about the Datahawk ThreatProtection subscription such as if it is active or not.
	Nested schema for **threat_protection**:
		* `end_date` - (Optional, String) Specifies the end date of the subscription.
		* `is_active` - (Optional, Boolean) Specifies the end date of the subscription.
		* `is_free_trial` - (Optional, Boolean) Specifies the end date of the subscription.
		* `start_date` - (Optional, String) Specifies the start date of the subscription.
* `tenant_accesses` - (Optional, Forces new resource, List) Specfies the Tenant Access for MCM User.
Nested schema for **tenant_accesses**:
	* `cluster_identifiers` - (Optional, List) Specifies the list of clusters.
	Nested schema for **cluster_identifiers**:
		* `cluster_id` - (Optional, Integer) Specifies the id of the cluster.
		* `cluster_incarnation_id` - (Optional, Integer) Specifies the incarnation id of the cluster.
	* `created_time_msecs` - (Optional, Integer) Specifies the epoch time in milliseconds when the tenant access was created.
	* `effective_time_msecs` - (Optional, Integer) Specifies the epoch time in milliseconds when the tenant access becomes effective. Until that time, the user cannot log in.
	* `expired_time_msecs` - (Optional, Integer) Specifies the epoch time in milliseconds when the tenant access becomes expired. After that, the user cannot log in.
	* `is_access_active` - (Optional, Boolean) IsAccessActive specifies whether or not a tenant access is active, or has been deactivated by the customer. The default behavior is 'true'.
	* `is_active` - (Optional, Boolean) Specifies whether or not the tenant is active.
	* `is_deleted` - (Optional, Boolean) Specifies whether or not the tenant is deleted.
	* `last_updated_time_msecs` - (Optional, Integer) Specifies the epoch time in milliseconds when the tenant access was last modified.
	* `roles` - (Optional, List) Specifies the Cohesity roles to associate with the user such as such as 'Admin', 'Ops' or 'View'.
	  * Constraints: The minimum length is `1` item.
	* `tenant_id` - (Optional, String) Specifies the tenant id.
	* `tenant_name` - (Optional, String) Specifies the tenant name.
	* `tenant_type` - (Optional, String) Specifies the MCM tenant type. 'Dmaas' implies tenant type is DMaaS. 'Mcm' implies tenant is Mcm Cluster tenant.
	  * Constraints: Allowable values are: `Dmaas`, `Mcm`.
* `tenant_id` - (Optional, Forces new resource, String) Specifies the effective Tenant ID of the user.
* `username` - (Optional, Forces new resource, String) Specifies the login name of the user.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the backup_recovery_connector_update_user.
* `group_roles` - (List) Specifies the Cohesity roles to associate with the user' group. These roles can only be edited from group.
* `is_account_mfa_enabled` - (Boolean) Specifies if MFA is enabled for the Helios Account.
* `is_cluster_mfa_enabled` - (Boolean) Specifies if MFA is enabled on cluster.

