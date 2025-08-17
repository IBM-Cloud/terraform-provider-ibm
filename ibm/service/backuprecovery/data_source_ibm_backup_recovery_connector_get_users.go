// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmBackupRecoveryConnectorGetUsers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryConnectorGetUsersRead,

		Schema: map[string]*schema.Schema{
			"session_name_cookie": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Sensitive:   true,
				Description: "Specifies the session name cookie of the Cohesity user.",
			},
			"tenant_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "TenantIds contains ids of the tenants for which objects are to be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"all_under_hierarchy": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned.TenantIds contains ids of the tenants for which objects are to be returned.",
			},
			"usernames": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Optionally specify a list of usernames to filter by. All users containing username will be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"email_addresses": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Optionally specify a list of email addresses to filter by.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"domain": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optionally specify a domain to filter by. If no domain is specified, all users on the Cohesity Cluster are searched. If a domain is specified, only users on the Cohesity Cluster associated with that domain are searched.",
			},
			"partial_match": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optionally specify whether to enable partial match. If set, all users with name containing Usernames will be returned. If set to false, only users with exact the same name as Usernames will be returned. By default this parameter is set to true.",
			},
			"users": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies list of users.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique ID.",
						},
						"ad_user_info": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies an AD User's information logged in using an active directory. This information is not stored on the Cluster.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"group_sids": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the SIDs of the groups.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"groups": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the groups this user is a part of.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"is_floating_user": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether this is a floating user or not.",
									},
								},
							},
						},
						"additional_group_names": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the names of additional groups this User may belong to.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"allow_dso_modify": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies if the data security user can be modified by the admin users.",
						},
						"audit_log_settings": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "AuditLogSettings specifies struct with audt log configuration. Make these settings in such a way that zero values are cluster default when bb is not present.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"read_logging": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "ReadLogging specifies whether read logs needs to be captured.",
									},
								},
							},
						},
						"authentication_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the authentication type of the user. 'kAuthLocal' implies authenticated user is a local user. 'kAuthAd' implies authenticated user is an Active Directory user. 'kAuthSalesforce' implies authenticated user is a Salesforce user. 'kAuthGoogle' implies authenticated user is a Google user. 'kAuthSso' implies authenticated user is an SSO user.",
						},
						"cluster_identifiers": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the list of clusters this user has access to. If this is not specified, access will be granted to all clusters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the id of the cluster.",
									},
									"cluster_incarnation_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the incarnation id of the cluster.",
									},
								},
							},
						},
						"created_time_msecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the epoch time in milliseconds when the user account was created on the Cohesity Cluster.",
						},
						"current_password": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the current password when updating the password.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies a description about the user.",
						},
						"domain": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the fully qualified domain name (FQDN) of an Active Directory or LOCAL for the default LOCAL domain on the Cohesity Cluster. A user is uniquely identified by combination of the username and the domain.",
						},
						"effective_time_msecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the epoch time in milliseconds when the user becomes effective. Until that time, the user cannot log in.",
						},
						"email_address": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the email address of the user.",
						},
						"expired_time_msecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the epoch time in milliseconds when the user becomes expired. After that, the user cannot log in.",
						},
						"force_password_change": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies whether to force user to change password.",
						},
						"google_account": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Google Account Information of a Helios BaaS user.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the Account Id assigned by Google.",
									},
									"user_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the User Id assigned by Google.",
									},
								},
							},
						},
						"group_roles": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the Cohesity roles to associate with the user' group. These roles can only be edited from group.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"idp_user_info": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies an IdP User's information logged in using an IdP. This information is not stored on the Cluster.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"group_sids": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the SIDs of the groups.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"groups": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the Idp groups that the user is part of. As the user may not be registered on the cluster, we may have to capture the idp group membership. This way, if a group is created on the cluster later, users will instantly have access to tenantIds from that group as well.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"idp_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the unique Id assigned by the Cluster for the IdP.",
									},
									"is_floating_user": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether or not this is a floating user.",
									},
									"issuer_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the unique identifier assigned by the vendor for this Cluster.",
									},
									"user_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the unique identifier assigned by the vendor for the user.",
									},
									"vendor": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the vendor providing the IdP service.",
									},
								},
							},
						},
						"intercom_messenger_token": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the messenger token for intercom identity verification.",
						},
						"is_account_locked": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies whether the user account is locked.",
						},
						"is_account_mfa_enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies if MFA is enabled for the Helios Account.",
						},
						"is_active": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "IsActive specifies whether or not a user is active, or has been disactivated by the customer. The default behavior is 'true'.",
						},
						"is_cluster_mfa_enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies if MFA is enabled on cluster.",
						},
						"last_successful_login_time_msecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the epoch time in milliseconds when the user was last logged in successfully.",
						},
						"last_updated_time_msecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the epoch time in milliseconds when the user account was last modified on the Cohesity Cluster.",
						},
						"mfa_info": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies information about MFA.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_email_otp_setup_done": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies if email OTP setup is done on the user.",
									},
									"is_totp_setup_done": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies if TOTP setup is done on the user.",
									},
									"is_user_exempt_from_mfa": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies if MFA is disabled on the user.",
									},
								},
							},
						},
						"mfa_methods": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies MFA methods that enabled on the cluster.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"object_class": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies object class of user, could be either user or group.",
						},
						"org_membership": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "OrgMembership contains the list of all available tenantIds for this user to switch to. Only when creating the session user, this field is populated on the fly. We discover the tenantIds from various groups assigned to the users.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bifrost_enabled": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies if this tenant is bifrost enabled or not.",
									},
									"is_managed_on_helios": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether this tenant is manged on helios.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies name of the tenant.",
									},
									"restricted": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the user is a restricted user. A restricted user can only view the objects he has permissions to.",
									},
									"roles": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the Cohesity roles to associate with the user such as such as 'Admin', 'Ops' or 'View'. The Cohesity roles determine privileges on the Cohesity Cluster for this user.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"tenant_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the unique id of the tenant.",
									},
								},
							},
						},
						"password": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the password of this user.",
						},
						"preferences": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the preferences of this user.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"locale": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Locale reflects the language settings of the user. Populate using the user preferences stored in Scribe for the user wherever needed.",
									},
								},
							},
						},
						"previous_login_time_msecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the epoch time in milliseconds of previous user login.",
						},
						"primary_group_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of the primary group of this User.",
						},
						"privilege_ids": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the Cohesity privileges from the roles. This will be populated based on the union of all privileges in roles. Type for unique privilege Id values. All below enum values specify a value for all uniquely defined privileges in Cohesity.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"profiles": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the user profiles. NOTE:- Currently used for Helios.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster_identifiers": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the list of clusters. This is only valid if tenant type is OnPrem.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cluster_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the id of the cluster.",
												},
												"cluster_incarnation_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the incarnation id of the cluster.",
												},
											},
										},
									},
									"is_active": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether or not the tenant is active.",
									},
									"is_deleted": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether or not the tenant is deleted.",
									},
									"region_ids": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the list of regions. This is only valid if tenant type is Dmaas.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"tenant_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the tenant id.",
									},
									"tenant_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the tenant id.",
									},
									"tenant_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the MCM tenant type. 'Dmaas' implies tenant type is DMaaS. 'Mcm' implies tenant is Mcm Cluster tenant.",
									},
								},
							},
						},
						"restricted": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the user is a restricted user. A restricted user can only view the objects he has permissions to.",
						},
						"roles": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the Cohesity roles to associate with the user such as such as 'Admin', 'Ops' or 'View'. The Cohesity roles determine privileges on the Cohesity Cluster for this user.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"s3_access_key_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the S3 Account Access Key ID.",
						},
						"s3_account_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the S3 Account Canonical User ID.",
						},
						"s3_secret_key": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the S3 Account Secret Key.",
						},
						"salesforce_account": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Salesforce Account Information of a Helios user.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the Account Id assigned by Salesforce.",
									},
									"helios_access_grant_status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the status of helios access.",
									},
									"is_d_gaa_s_user": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether user is a DGaaS licensed user.",
									},
									"is_d_maa_s_user": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether user is a DMaaS licensed user.",
									},
									"is_d_raa_s_user": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether user is a DRaaS licensed user.",
									},
									"is_r_paa_s_user": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether user is a RPaaS licensed user.",
									},
									"is_sales_user": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether user is a Sales person from Cohesity.",
									},
									"is_support_user": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether user is a support person from Cohesity.",
									},
									"user_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the User Id assigned by Salesforce.",
									},
								},
							},
						},
						"sid": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the unique Security ID (SID) of the user. This field is mandatory in modifying user.",
						},
						"spog_context": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "SpogContext specifies all of the information about the user and cluster which is performing action on this cluster.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"primary_cluster_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the ID of the remote cluster which is accessing this cluster via SPOG.",
									},
									"primary_cluster_user_sid": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the SID of the user who is accessing this cluster via SPOG.",
									},
									"primary_cluster_username": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the username of the user who is accessing this cluster via SPOG.",
									},
								},
							},
						},
						"subscription_info": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Extends this to have Helios, DRaaS and DSaaS.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"classification": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "ClassificationInfo holds information about the Datahawk Classification subscription such as if it is active or not.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"end_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_active": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_free_trial": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"start_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the start date of the subscription.",
												},
											},
										},
									},
									"data_protect": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "DMaaSSubscriptionInfo holds information about the Data Protect subscription such as if it is active or not.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"end_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_active": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_free_trial": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_aws_subscription": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether the subscription is AWS Subscription.",
												},
												"is_cohesity_subscription": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether the subscription is a Cohesity Paid subscription.",
												},
												"quantity": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the quantity of the subscription.",
												},
												"start_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the start date of the subscription.",
												},
												"tiering": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the tiering info.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"backend_tiering": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether back-end tiering is enabled.",
															},
															"frontend_tiering": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether Front End Tiering Enabled.",
															},
															"max_retention": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specified the max retention for backup policy creation.",
															},
														},
													},
												},
											},
										},
									},
									"data_protect_azure": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "ClassificationInfo holds information about the Datahawk Classification subscription such as if it is active or not.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"end_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_active": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_free_trial": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"quantity": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the quantity of the subscription.",
												},
												"start_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the start date of the subscription.",
												},
												"tiering": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the tiering info.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"backend_tiering": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether back-end tiering is enabled.",
															},
															"frontend_tiering": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether Front End Tiering Enabled.",
															},
															"max_retention": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specified the max retention for backup policy creation.",
															},
														},
													},
												},
											},
										},
									},
									"fort_knox_azure_cool": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "FortKnoxInfo holds information about the Fortknox Azure or Azure FreeTrial or AwsCold or AwsCold FreeTrial subscription such as if it is active.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"end_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_active": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_free_trial": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"quantity": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the quantity of the subscription.",
												},
												"start_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the start date of the subscription.",
												},
											},
										},
									},
									"fort_knox_azure_hot": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "FortKnoxInfo holds information about the Fortknox Azure or Azure FreeTrial or AwsCold or AwsCold FreeTrial subscription such as if it is active.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"end_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_active": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_free_trial": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"quantity": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the quantity of the subscription.",
												},
												"start_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the start date of the subscription.",
												},
											},
										},
									},
									"fort_knox_cold": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "FortKnoxInfo holds information about the Fortknox Azure or Azure FreeTrial or AwsCold or AwsCold FreeTrial subscription such as if it is active.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"end_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_active": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_free_trial": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"quantity": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the quantity of the subscription.",
												},
												"start_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the start date of the subscription.",
												},
											},
										},
									},
									"ransomware": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "RansomwareInfo holds information about the FortKnox/FortKnoxFreeTrial subscription such as if it is active or not.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"end_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_active": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_free_trial": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"quantity": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the quantity of the subscription.",
												},
												"start_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the start date of the subscription.",
												},
											},
										},
									},
									"site_continuity": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "SiteContinuityInfo holds information about the Site Continuity subscription such as if it is active or not.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"end_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_active": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_free_trial": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"start_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the start date of the subscription.",
												},
											},
										},
									},
									"threat_protection": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "ThreatProtectionInfo holds information about the Datahawk ThreatProtection subscription such as if it is active or not.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"end_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_active": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"is_free_trial": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies the end date of the subscription.",
												},
												"start_date": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the start date of the subscription.",
												},
											},
										},
									},
								},
							},
						},
						"tenant_accesses": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specfies the Tenant Access for MCM User.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster_identifiers": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the list of clusters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cluster_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the id of the cluster.",
												},
												"cluster_incarnation_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the incarnation id of the cluster.",
												},
											},
										},
									},
									"created_time_msecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the epoch time in milliseconds when the tenant access was created.",
									},
									"effective_time_msecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the epoch time in milliseconds when the tenant access becomes effective. Until that time, the user cannot log in.",
									},
									"expired_time_msecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the epoch time in milliseconds when the tenant access becomes expired. After that, the user cannot log in.",
									},
									"is_access_active": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "IsAccessActive specifies whether or not a tenant access is active, or has been deactivated by the customer. The default behavior is 'true'.",
									},
									"is_active": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether or not the tenant is active.",
									},
									"is_deleted": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether or not the tenant is deleted.",
									},
									"last_updated_time_msecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the epoch time in milliseconds when the tenant access was last modified.",
									},
									"roles": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the Cohesity roles to associate with the user such as such as 'Admin', 'Ops' or 'View'.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"tenant_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the tenant id.",
									},
									"tenant_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the tenant name.",
									},
									"tenant_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the MCM tenant type. 'Dmaas' implies tenant type is DMaaS. 'Mcm' implies tenant is Mcm Cluster tenant.",
									},
								},
							},
						},
						"tenant_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the effective Tenant ID of the user.",
						},
						"username": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the login name of the user.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryConnectorGetUsersRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1Connector()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_connector_get_users", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getUsersOptions := &backuprecoveryv1.GetUsersOptions{}

	getUsersOptions.SetSessionName(d.Get("session_name_cookie").(string))

	if _, ok := d.GetOk("tenant_ids"); ok {
		var tenantIds []string
		for _, v := range d.Get("tenant_ids").([]interface{}) {
			tenantIdsItem := v.(string)
			tenantIds = append(tenantIds, tenantIdsItem)
		}
		getUsersOptions.SetTenantIds(tenantIds)
	}
	if _, ok := d.GetOk("all_under_hierarchy"); ok {
		getUsersOptions.SetAllUnderHierarchy(d.Get("all_under_hierarchy").(bool))
	}
	if _, ok := d.GetOk("usernames"); ok {
		var usernames []string
		for _, v := range d.Get("usernames").([]interface{}) {
			usernamesItem := v.(string)
			usernames = append(usernames, usernamesItem)
		}
		getUsersOptions.SetUsernames(usernames)
	}
	if _, ok := d.GetOk("email_addresses"); ok {
		var emailAddresses []string
		for _, v := range d.Get("email_addresses").([]interface{}) {
			emailAddressesItem := v.(string)
			emailAddresses = append(emailAddresses, emailAddressesItem)
		}
		getUsersOptions.SetEmailAddresses(emailAddresses)
	}
	if _, ok := d.GetOk("domain"); ok {
		getUsersOptions.SetDomain(d.Get("domain").(string))
	}
	if _, ok := d.GetOk("partial_match"); ok {
		getUsersOptions.SetPartialMatch(d.Get("partial_match").(bool))
	}

	getUsersResponse, _, err := backupRecoveryClient.GetUsersWithContext(context, getUsersOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetUsersWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_connector_get_users", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryConnectorGetUsersID(d))

	if !core.IsNil(getUsersResponse) {
		users := []map[string]interface{}{}
		for _, usersItem := range getUsersResponse {
			usersItemMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersUserDetailsToMap(&usersItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_connector_get_users", "read", "users-to-map").GetDiag()
			}
			users = append(users, usersItemMap)
		}
		if err = d.Set("users", users); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting users: %s", err), "(Data) ibm_backup_recovery_connector_get_users", "read", "set-users").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryConnectorGetUsersID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryConnectorGetUsersID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryConnectorGetUsersUserDetailsToMap(model *backuprecoveryv1.UserDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})

	if model.AdUserInfo != nil {
		adUserInfoMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersAdUserInfoToMap(model.AdUserInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["ad_user_info"] = []map[string]interface{}{adUserInfoMap}
	}
	if model.AdditionalGroupNames != nil {
		modelMap["additional_group_names"] = model.AdditionalGroupNames
	}
	if model.AllowDsoModify != nil {
		modelMap["allow_dso_modify"] = *model.AllowDsoModify
	}
	if model.AuditLogSettings != nil {
		auditLogSettingsMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersAuditLogSettingsToMap(model.AuditLogSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["audit_log_settings"] = []map[string]interface{}{auditLogSettingsMap}
	}
	if model.AuthenticationType != nil {
		modelMap["authentication_type"] = *model.AuthenticationType
	}
	if model.ClusterIdentifiers != nil {
		clusterIdentifiers := []map[string]interface{}{}
		for _, clusterIdentifiersItem := range model.ClusterIdentifiers {
			clusterIdentifiersItemMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersUserClusterIdentifierToMap(&clusterIdentifiersItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			clusterIdentifiers = append(clusterIdentifiers, clusterIdentifiersItemMap)
		}
		modelMap["cluster_identifiers"] = clusterIdentifiers
	}
	if model.CreatedTimeMsecs != nil {
		modelMap["created_time_msecs"] = flex.IntValue(model.CreatedTimeMsecs)
	}
	if model.CurrentPassword != nil {
		modelMap["current_password"] = *model.CurrentPassword
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Domain != nil {
		modelMap["domain"] = *model.Domain
	}
	if model.EffectiveTimeMsecs != nil {
		modelMap["effective_time_msecs"] = flex.IntValue(model.EffectiveTimeMsecs)
	}
	if model.EmailAddress != nil {
		modelMap["email_address"] = *model.EmailAddress
	}
	if model.ExpiredTimeMsecs != nil {
		modelMap["expired_time_msecs"] = flex.IntValue(model.ExpiredTimeMsecs)
	}
	if model.ForcePasswordChange != nil {
		modelMap["force_password_change"] = *model.ForcePasswordChange
	}
	if model.GoogleAccount != nil {
		googleAccountMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersGoogleAccountInfoToMap(model.GoogleAccount)
		if err != nil {
			return modelMap, err
		}
		modelMap["google_account"] = []map[string]interface{}{googleAccountMap}
	}
	if model.GroupRoles != nil {
		modelMap["group_roles"] = model.GroupRoles
	}
	if model.IdpUserInfo != nil {
		idpUserInfoMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersIdpUserInfoToMap(model.IdpUserInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["idp_user_info"] = []map[string]interface{}{idpUserInfoMap}
	}
	if model.IntercomMessengerToken != nil {
		modelMap["intercom_messenger_token"] = *model.IntercomMessengerToken
	}
	if model.IsAccountLocked != nil {
		modelMap["is_account_locked"] = *model.IsAccountLocked
	}
	if model.IsAccountMfaEnabled != nil {
		modelMap["is_account_mfa_enabled"] = *model.IsAccountMfaEnabled
	}
	if model.IsActive != nil {
		modelMap["is_active"] = *model.IsActive
	}
	if model.IsClusterMfaEnabled != nil {
		modelMap["is_cluster_mfa_enabled"] = *model.IsClusterMfaEnabled
	}
	if model.LastSuccessfulLoginTimeMsecs != nil {
		modelMap["last_successful_login_time_msecs"] = flex.IntValue(model.LastSuccessfulLoginTimeMsecs)
	}
	if model.LastUpdatedTimeMsecs != nil {
		modelMap["last_updated_time_msecs"] = flex.IntValue(model.LastUpdatedTimeMsecs)
	}
	if model.MfaInfo != nil {
		mfaInfoMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersMfaInfoToMap(model.MfaInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["mfa_info"] = []map[string]interface{}{mfaInfoMap}
	}
	if model.MfaMethods != nil {
		modelMap["mfa_methods"] = model.MfaMethods
	}
	if model.ObjectClass != nil {
		modelMap["object_class"] = *model.ObjectClass
	}
	if model.OrgMembership != nil {
		orgMembership := []map[string]interface{}{}
		for _, orgMembershipItem := range model.OrgMembership {
			orgMembershipItemMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersTenantConfigToMap(&orgMembershipItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			orgMembership = append(orgMembership, orgMembershipItemMap)
		}
		modelMap["org_membership"] = orgMembership
	}
	if model.Password != nil {
		modelMap["password"] = *model.Password
	}
	if model.Preferences != nil {
		preferencesMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersUsersPreferencesToMap(model.Preferences)
		if err != nil {
			return modelMap, err
		}
		modelMap["preferences"] = []map[string]interface{}{preferencesMap}
	}
	if model.PreviousLoginTimeMsecs != nil {
		modelMap["previous_login_time_msecs"] = flex.IntValue(model.PreviousLoginTimeMsecs)
	}
	if model.PrimaryGroupName != nil {
		modelMap["primary_group_name"] = *model.PrimaryGroupName
	}
	if model.PrivilegeIds != nil {
		modelMap["privilege_ids"] = model.PrivilegeIds
	}
	if model.Profiles != nil {
		profiles := []map[string]interface{}{}
		for _, profilesItem := range model.Profiles {
			profilesItemMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersUserProfileToMap(&profilesItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			profiles = append(profiles, profilesItemMap)
		}
		modelMap["profiles"] = profiles
	}
	if model.Restricted != nil {
		modelMap["restricted"] = *model.Restricted
	}
	if model.Roles != nil {
		modelMap["roles"] = model.Roles
	}
	if model.S3AccessKeyID != nil {
		modelMap["s3_access_key_id"] = *model.S3AccessKeyID
	}
	if model.S3AccountID != nil {
		modelMap["s3_account_id"] = *model.S3AccountID
	}
	if model.S3SecretKey != nil {
		modelMap["s3_secret_key"] = *model.S3SecretKey
	}
	if model.SalesforceAccount != nil {
		salesforceAccountMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersSalesforceAccountInfoToMap(model.SalesforceAccount)
		if err != nil {
			return modelMap, err
		}
		modelMap["salesforce_account"] = []map[string]interface{}{salesforceAccountMap}
	}
	if model.Sid != nil {
		modelMap["sid"] = *model.Sid
	}
	if model.SpogContext != nil {
		spogContextMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersSpogContextToMap(model.SpogContext)
		if err != nil {
			return modelMap, err
		}
		modelMap["spog_context"] = []map[string]interface{}{spogContextMap}
	}
	if model.SubscriptionInfo != nil {
		subscriptionInfoMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersSubscriptionInfoToMap(model.SubscriptionInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["subscription_info"] = []map[string]interface{}{subscriptionInfoMap}
	}
	if model.TenantAccesses != nil {
		tenantAccesses := []map[string]interface{}{}
		for _, tenantAccessesItem := range model.TenantAccesses {
			tenantAccessesItemMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersTenantAccessesToMap(&tenantAccessesItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			tenantAccesses = append(tenantAccesses, tenantAccessesItemMap)
		}
		modelMap["tenant_accesses"] = tenantAccesses
	}
	if model.TenantID != nil {
		modelMap["tenant_id"] = *model.TenantID
	}
	modelMap["username"] = *model.Username
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersAdUserInfoToMap(model *backuprecoveryv1.AdUserInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.GroupSids != nil {
		modelMap["group_sids"] = model.GroupSids
	}
	if model.Groups != nil {
		modelMap["groups"] = model.Groups
	}
	if model.IsFloatingUser != nil {
		modelMap["is_floating_user"] = *model.IsFloatingUser
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersAuditLogSettingsToMap(model *backuprecoveryv1.AuditLogSettings) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ReadLogging != nil {
		modelMap["read_logging"] = *model.ReadLogging
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersUserClusterIdentifierToMap(model *backuprecoveryv1.UserClusterIdentifier) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClusterID != nil {
		modelMap["cluster_id"] = flex.IntValue(model.ClusterID)
	}
	if model.ClusterIncarnationID != nil {
		modelMap["cluster_incarnation_id"] = flex.IntValue(model.ClusterIncarnationID)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersGoogleAccountInfoToMap(model *backuprecoveryv1.GoogleAccountInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AccountID != nil {
		modelMap["account_id"] = *model.AccountID
	}
	if model.UserID != nil {
		modelMap["user_id"] = *model.UserID
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersIdpUserInfoToMap(model *backuprecoveryv1.IdpUserInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.GroupSids != nil {
		modelMap["group_sids"] = model.GroupSids
	}
	if model.Groups != nil {
		modelMap["groups"] = model.Groups
	}
	if model.IdpID != nil {
		modelMap["idp_id"] = flex.IntValue(model.IdpID)
	}
	if model.IsFloatingUser != nil {
		modelMap["is_floating_user"] = *model.IsFloatingUser
	}
	if model.IssuerID != nil {
		modelMap["issuer_id"] = *model.IssuerID
	}
	if model.UserID != nil {
		modelMap["user_id"] = *model.UserID
	}
	if model.Vendor != nil {
		modelMap["vendor"] = *model.Vendor
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersMfaInfoToMap(model *backuprecoveryv1.MfaInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.IsEmailOtpSetupDone != nil {
		modelMap["is_email_otp_setup_done"] = *model.IsEmailOtpSetupDone
	}
	if model.IsTotpSetupDone != nil {
		modelMap["is_totp_setup_done"] = *model.IsTotpSetupDone
	}
	if model.IsUserExemptFromMfa != nil {
		modelMap["is_user_exempt_from_mfa"] = *model.IsUserExemptFromMfa
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersTenantConfigToMap(model *backuprecoveryv1.TenantConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BifrostEnabled != nil {
		modelMap["bifrost_enabled"] = *model.BifrostEnabled
	}
	if model.IsManagedOnHelios != nil {
		modelMap["is_managed_on_helios"] = *model.IsManagedOnHelios
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Restricted != nil {
		modelMap["restricted"] = *model.Restricted
	}
	if model.Roles != nil {
		modelMap["roles"] = model.Roles
	}
	if model.TenantID != nil {
		modelMap["tenant_id"] = *model.TenantID
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersUsersPreferencesToMap(model *backuprecoveryv1.UsersPreferences) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Locale != nil {
		modelMap["locale"] = *model.Locale
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersUserProfileToMap(model *backuprecoveryv1.UserProfile) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClusterIdentifiers != nil {
		clusterIdentifiers := []map[string]interface{}{}
		for _, clusterIdentifiersItem := range model.ClusterIdentifiers {
			clusterIdentifiersItemMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersUserClusterIdentifierToMap(&clusterIdentifiersItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			clusterIdentifiers = append(clusterIdentifiers, clusterIdentifiersItemMap)
		}
		modelMap["cluster_identifiers"] = clusterIdentifiers
	}
	if model.IsActive != nil {
		modelMap["is_active"] = *model.IsActive
	}
	if model.IsDeleted != nil {
		modelMap["is_deleted"] = *model.IsDeleted
	}
	if model.RegionIds != nil {
		modelMap["region_ids"] = model.RegionIds
	}
	if model.TenantID != nil {
		modelMap["tenant_id"] = *model.TenantID
	}
	if model.TenantName != nil {
		modelMap["tenant_name"] = *model.TenantName
	}
	if model.TenantType != nil {
		modelMap["tenant_type"] = *model.TenantType
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersSalesforceAccountInfoToMap(model *backuprecoveryv1.SalesforceAccountInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AccountID != nil {
		modelMap["account_id"] = *model.AccountID
	}
	if model.HeliosAccessGrantStatus != nil {
		modelMap["helios_access_grant_status"] = *model.HeliosAccessGrantStatus
	}
	if model.IsDGaaSUser != nil {
		modelMap["is_d_gaa_s_user"] = *model.IsDGaaSUser
	}
	if model.IsDMaaSUser != nil {
		modelMap["is_d_maa_s_user"] = *model.IsDMaaSUser
	}
	if model.IsDRaaSUser != nil {
		modelMap["is_d_raa_s_user"] = *model.IsDRaaSUser
	}
	if model.IsRPaaSUser != nil {
		modelMap["is_r_paa_s_user"] = *model.IsRPaaSUser
	}
	if model.IsSalesUser != nil {
		modelMap["is_sales_user"] = *model.IsSalesUser
	}
	if model.IsSupportUser != nil {
		modelMap["is_support_user"] = *model.IsSupportUser
	}
	if model.UserID != nil {
		modelMap["user_id"] = *model.UserID
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersSpogContextToMap(model *backuprecoveryv1.SpogContext) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PrimaryClusterID != nil {
		modelMap["primary_cluster_id"] = flex.IntValue(model.PrimaryClusterID)
	}
	if model.PrimaryClusterUserSid != nil {
		modelMap["primary_cluster_user_sid"] = *model.PrimaryClusterUserSid
	}
	if model.PrimaryClusterUsername != nil {
		modelMap["primary_cluster_username"] = *model.PrimaryClusterUsername
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersSubscriptionInfoToMap(model *backuprecoveryv1.SubscriptionInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Classification != nil {
		classificationMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersClassificationInfoToMap(model.Classification)
		if err != nil {
			return modelMap, err
		}
		modelMap["classification"] = []map[string]interface{}{classificationMap}
	}
	if model.DataProtect != nil {
		dataProtectMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersDataProtectInfoToMap(model.DataProtect)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_protect"] = []map[string]interface{}{dataProtectMap}
	}
	if model.DataProtectAzure != nil {
		dataProtectAzureMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersDataProtectAzureInfoToMap(model.DataProtectAzure)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_protect_azure"] = []map[string]interface{}{dataProtectAzureMap}
	}
	if model.FortKnoxAzureCool != nil {
		fortKnoxAzureCoolMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersFortKnoxInfoToMap(model.FortKnoxAzureCool)
		if err != nil {
			return modelMap, err
		}
		modelMap["fort_knox_azure_cool"] = []map[string]interface{}{fortKnoxAzureCoolMap}
	}
	if model.FortKnoxAzureHot != nil {
		fortKnoxAzureHotMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersFortKnoxInfoToMap(model.FortKnoxAzureHot)
		if err != nil {
			return modelMap, err
		}
		modelMap["fort_knox_azure_hot"] = []map[string]interface{}{fortKnoxAzureHotMap}
	}
	if model.FortKnoxCold != nil {
		fortKnoxColdMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersFortKnoxInfoToMap(model.FortKnoxCold)
		if err != nil {
			return modelMap, err
		}
		modelMap["fort_knox_cold"] = []map[string]interface{}{fortKnoxColdMap}
	}
	if model.Ransomware != nil {
		ransomwareMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersFortKnoxInfoToMap(model.Ransomware)
		if err != nil {
			return modelMap, err
		}
		modelMap["ransomware"] = []map[string]interface{}{ransomwareMap}
	}
	if model.SiteContinuity != nil {
		siteContinuityMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersClassificationInfoToMap(model.SiteContinuity)
		if err != nil {
			return modelMap, err
		}
		modelMap["site_continuity"] = []map[string]interface{}{siteContinuityMap}
	}
	if model.ThreatProtection != nil {
		threatProtectionMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersClassificationInfoToMap(model.ThreatProtection)
		if err != nil {
			return modelMap, err
		}
		modelMap["threat_protection"] = []map[string]interface{}{threatProtectionMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersClassificationInfoToMap(model *backuprecoveryv1.ClassificationInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EndDate != nil {
		modelMap["end_date"] = *model.EndDate
	}
	if model.IsActive != nil {
		modelMap["is_active"] = *model.IsActive
	}
	if model.IsFreeTrial != nil {
		modelMap["is_free_trial"] = *model.IsFreeTrial
	}
	if model.StartDate != nil {
		modelMap["start_date"] = *model.StartDate
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersDataProtectInfoToMap(model *backuprecoveryv1.DataProtectInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EndDate != nil {
		modelMap["end_date"] = *model.EndDate
	}
	if model.IsActive != nil {
		modelMap["is_active"] = *model.IsActive
	}
	if model.IsFreeTrial != nil {
		modelMap["is_free_trial"] = *model.IsFreeTrial
	}
	if model.IsAwsSubscription != nil {
		modelMap["is_aws_subscription"] = *model.IsAwsSubscription
	}
	if model.IsCohesitySubscription != nil {
		modelMap["is_cohesity_subscription"] = *model.IsCohesitySubscription
	}
	if model.Quantity != nil {
		modelMap["quantity"] = flex.IntValue(model.Quantity)
	}
	if model.StartDate != nil {
		modelMap["start_date"] = *model.StartDate
	}
	if model.Tiering != nil {
		tieringMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersTieringInfoToMap(model.Tiering)
		if err != nil {
			return modelMap, err
		}
		modelMap["tiering"] = []map[string]interface{}{tieringMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersTieringInfoToMap(model *backuprecoveryv1.TieringInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BackendTiering != nil {
		modelMap["backend_tiering"] = *model.BackendTiering
	}
	if model.FrontendTiering != nil {
		modelMap["frontend_tiering"] = *model.FrontendTiering
	}
	if model.MaxRetention != nil {
		modelMap["max_retention"] = flex.IntValue(model.MaxRetention)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersDataProtectAzureInfoToMap(model *backuprecoveryv1.DataProtectAzureInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EndDate != nil {
		modelMap["end_date"] = *model.EndDate
	}
	if model.IsActive != nil {
		modelMap["is_active"] = *model.IsActive
	}
	if model.IsFreeTrial != nil {
		modelMap["is_free_trial"] = *model.IsFreeTrial
	}
	if model.Quantity != nil {
		modelMap["quantity"] = flex.IntValue(model.Quantity)
	}
	if model.StartDate != nil {
		modelMap["start_date"] = *model.StartDate
	}
	if model.Tiering != nil {
		tieringMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersTieringInfoToMap(model.Tiering)
		if err != nil {
			return modelMap, err
		}
		modelMap["tiering"] = []map[string]interface{}{tieringMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersFortKnoxInfoToMap(model *backuprecoveryv1.FortKnoxInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EndDate != nil {
		modelMap["end_date"] = *model.EndDate
	}
	if model.IsActive != nil {
		modelMap["is_active"] = *model.IsActive
	}
	if model.IsFreeTrial != nil {
		modelMap["is_free_trial"] = *model.IsFreeTrial
	}
	if model.Quantity != nil {
		modelMap["quantity"] = flex.IntValue(model.Quantity)
	}
	if model.StartDate != nil {
		modelMap["start_date"] = *model.StartDate
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorGetUsersTenantAccessesToMap(model *backuprecoveryv1.TenantAccesses) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClusterIdentifiers != nil {
		clusterIdentifiers := []map[string]interface{}{}
		for _, clusterIdentifiersItem := range model.ClusterIdentifiers {
			clusterIdentifiersItemMap, err := DataSourceIbmBackupRecoveryConnectorGetUsersUserClusterIdentifierToMap(&clusterIdentifiersItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			clusterIdentifiers = append(clusterIdentifiers, clusterIdentifiersItemMap)
		}
		modelMap["cluster_identifiers"] = clusterIdentifiers
	}
	if model.CreatedTimeMsecs != nil {
		modelMap["created_time_msecs"] = flex.IntValue(model.CreatedTimeMsecs)
	}
	if model.EffectiveTimeMsecs != nil {
		modelMap["effective_time_msecs"] = flex.IntValue(model.EffectiveTimeMsecs)
	}
	if model.ExpiredTimeMsecs != nil {
		modelMap["expired_time_msecs"] = flex.IntValue(model.ExpiredTimeMsecs)
	}
	if model.IsAccessActive != nil {
		modelMap["is_access_active"] = *model.IsAccessActive
	}
	if model.IsActive != nil {
		modelMap["is_active"] = *model.IsActive
	}
	if model.IsDeleted != nil {
		modelMap["is_deleted"] = *model.IsDeleted
	}
	if model.LastUpdatedTimeMsecs != nil {
		modelMap["last_updated_time_msecs"] = flex.IntValue(model.LastUpdatedTimeMsecs)
	}
	if model.Roles != nil {
		modelMap["roles"] = model.Roles
	}
	if model.TenantID != nil {
		modelMap["tenant_id"] = *model.TenantID
	}
	if model.TenantName != nil {
		modelMap["tenant_name"] = *model.TenantName
	}
	if model.TenantType != nil {
		modelMap["tenant_type"] = *model.TenantType
	}
	return modelMap, nil
}
