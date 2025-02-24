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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmBackupRecoveryConnectorUpdateUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmBackupRecoveryConnectorUpdateUserCreate,
		ReadContext:   resourceIbmBackupRecoveryConnectorUpdateUserRead,
		DeleteContext: resourceIbmBackupRecoveryConnectorUpdateUserDelete,
		UpdateContext: resourceIbmBackupRecoveryConnectorUpdateUserUpdate,
		Importer:      &schema.ResourceImporter{},
		CustomizeDiff: checkDiffResourceIbmBackupRecoveryConnectorUpdateUser,
		Schema: map[string]*schema.Schema{
			"session_name_cookie": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the session name cookie of the Cohesity user.",
			},
			"ad_user_info": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies an AD User's information logged in using an active directory. This information is not stored on the Cluster.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_sids": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the SIDs of the groups.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"groups": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the groups this user is a part of.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"is_floating_user": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether this is a floating user or not.",
						},
					},
				},
			},
			"additional_group_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				// ForceNew:    true
				Description: "Specifies the names of additional groups this User may belong to.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"allow_dso_modify": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies if the data security user can be modified by the admin users.",
			},
			"audit_log_settings": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "AuditLogSettings specifies struct with audt log configuration. Make these settings in such a way that zero values are cluster default when bb is not present.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"read_logging": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "ReadLogging specifies whether read logs needs to be captured.",
						},
					},
				},
			},
			"authentication_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the authentication type of the user. 'kAuthLocal' implies authenticated user is a local user. 'kAuthAd' implies authenticated user is an Active Directory user. 'kAuthSalesforce' implies authenticated user is a Salesforce user. 'kAuthGoogle' implies authenticated user is a Google user. 'kAuthSso' implies authenticated user is an SSO user.",
			},
			"cluster_identifiers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the list of clusters this user has access to. If this is not specified, access will be granted to all clusters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the id of the cluster.",
						},
						"cluster_incarnation_id": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the incarnation id of the cluster.",
						},
					},
				},
			},
			"created_time_msecs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the epoch time in milliseconds when the user account was created on the Cohesity Cluster.",
			},
			"current_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				// ForceNew:    true
				Description: "Specifies the current password when updating the password.",
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies a description about the user.",
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the fully qualified domain name (FQDN) of an Active Directory or LOCAL for the default LOCAL domain on the Cohesity Cluster. A user is uniquely identified by combination of the username and the domain.",
			},
			"effective_time_msecs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the epoch time in milliseconds when the user becomes effective. Until that time, the user cannot log in.",
			},
			"email_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the email address of the user.",
			},
			"expired_time_msecs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the epoch time in milliseconds when the user becomes expired. After that, the user cannot log in.",
			},
			"force_password_change": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies whether to force user to change password.",
			},
			"google_account": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Google Account Information of a Helios BaaS user.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the Account Id assigned by Google.",
						},
						"user_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the User Id assigned by Google.",
						},
					},
				},
			},
			"idp_user_info": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies an IdP User's information logged in using an IdP. This information is not stored on the Cluster.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_sids": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the SIDs of the groups.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"groups": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the Idp groups that the user is part of. As the user may not be registered on the cluster, we may have to capture the idp group membership. This way, if a group is created on the cluster later, users will instantly have access to tenantIds from that group as well.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"idp_id": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the unique Id assigned by the Cluster for the IdP.",
						},
						"is_floating_user": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether or not this is a floating user.",
						},
						"issuer_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the unique identifier assigned by the vendor for this Cluster.",
						},
						"user_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the unique identifier assigned by the vendor for the user.",
						},
						"vendor": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the vendor providing the IdP service.",
						},
					},
				},
			},
			"intercom_messenger_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the messenger token for intercom identity verification.",
			},
			"is_account_locked": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies whether the user account is locked.",
			},
			"is_active": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "IsActive specifies whether or not a user is active, or has been disactivated by the customer. The default behavior is 'true'.",
			},
			"last_successful_login_time_msecs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the epoch time in milliseconds when the user was last logged in successfully.",
			},
			"last_updated_time_msecs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the epoch time in milliseconds when the user account was last modified on the Cohesity Cluster.",
			},
			"mfa_info": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				// ForceNew:    true
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
							Optional:    true,
							Description: "Specifies if MFA is disabled on the user.",
						},
					},
				},
			},
			"mfa_methods": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				// ForceNew:    true
				Description: "Specifies MFA methods that enabled on the cluster.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"object_class": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies object class of user, could be either user or group.",
			},
			"org_membership": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "OrgMembership contains the list of all available tenantIds for this user to switch to. Only when creating the session user, this field is populated on the fly. We discover the tenantIds from various groups assigned to the users.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bifrost_enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies if this tenant is bifrost enabled or not.",
						},
						"is_managed_on_helios": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether this tenant is manged on helios.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies name of the tenant.",
						},
						"restricted": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether the user is a restricted user. A restricted user can only view the objects he has permissions to.",
						},
						"roles": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the Cohesity roles to associate with the user such as such as 'Admin', 'Ops' or 'View'. The Cohesity roles determine privileges on the Cohesity Cluster for this user.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"tenant_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the unique id of the tenant.",
						},
					},
				},
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				// ForceNew:    true
				Description: "Specifies the password of this user.",
			},
			"preferences": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the preferences of this user.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"locale": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Locale reflects the language settings of the user. Populate using the user preferences stored in Scribe for the user wherever needed.",
						},
					},
				},
			},
			"previous_login_time_msecs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the epoch time in milliseconds of previous user login.",
			},
			"primary_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the name of the primary group of this User.",
			},
			"privilege_ids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				// ForceNew:    true
				Description: "Specifies the Cohesity privileges from the roles. This will be populated based on the union of all privileges in roles. Type for unique privilege Id values. All below enum values specify a value for all uniquely defined privileges in Cohesity.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"profiles": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the user profiles. NOTE:- Currently used for Helios.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_identifiers": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the list of clusters. This is only valid if tenant type is OnPrem.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster_id": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies the id of the cluster.",
									},
									"cluster_incarnation_id": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies the incarnation id of the cluster.",
									},
								},
							},
						},
						"is_active": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether or not the tenant is active.",
						},
						"is_deleted": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether or not the tenant is deleted.",
						},
						"region_ids": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the list of regions. This is only valid if tenant type is Dmaas.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"tenant_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the tenant id.",
						},
						"tenant_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the tenant id.",
						},
						"tenant_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the MCM tenant type. 'Dmaas' implies tenant type is DMaaS. 'Mcm' implies tenant is Mcm Cluster tenant.",
						},
					},
				},
			},
			"restricted": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Whether the user is a restricted user. A restricted user can only view the objects he has permissions to.",
			},
			"roles": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the Cohesity roles to associate with the user such as such as 'Admin', 'Ops' or 'View'. The Cohesity roles determine privileges on the Cohesity Cluster for this user.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"s3_access_key_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the S3 Account Access Key ID.",
			},
			"s3_account_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the S3 Account Canonical User ID.",
			},
			"s3_secret_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the S3 Account Secret Key.",
			},
			"salesforce_account": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Salesforce Account Information of a Helios user.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the Account Id assigned by Salesforce.",
						},
						"helios_access_grant_status": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the status of helios access.",
						},
						"is_d_gaa_s_user": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether user is a DGaaS licensed user.",
						},
						"is_d_maa_s_user": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether user is a DMaaS licensed user.",
						},
						"is_d_raa_s_user": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether user is a DRaaS licensed user.",
						},
						"is_r_paa_s_user": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether user is a RPaaS licensed user.",
						},
						"is_sales_user": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether user is a Sales person from Cohesity.",
						},
						"is_support_user": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether user is a support person from Cohesity.",
						},
						"user_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the User Id assigned by Salesforce.",
						},
					},
				},
			},
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the unique Security ID (SID) of the user. This field is mandatory in modifying user.",
			},
			"spog_context": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "SpogContext specifies all of the information about the user and cluster which is performing action on this cluster.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"primary_cluster_id": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the ID of the remote cluster which is accessing this cluster via SPOG.",
						},
						"primary_cluster_user_sid": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the SID of the user who is accessing this cluster via SPOG.",
						},
						"primary_cluster_username": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the username of the user who is accessing this cluster via SPOG.",
						},
					},
				},
			},
			"subscription_info": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Extends this to have Helios, DRaaS and DSaaS.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"classification": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "ClassificationInfo holds information about the Datahawk Classification subscription such as if it is active or not.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_active": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_free_trial": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"start_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the start date of the subscription.",
									},
								},
							},
						},
						"data_protect": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "DMaaSSubscriptionInfo holds information about the Data Protect subscription such as if it is active or not.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_active": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_free_trial": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_aws_subscription": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies whether the subscription is AWS Subscription.",
									},
									"is_cohesity_subscription": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies whether the subscription is a Cohesity Paid subscription.",
									},
									"quantity": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies the quantity of the subscription.",
									},
									"start_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the start date of the subscription.",
									},
									"tiering": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the tiering info.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"backend_tiering": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies whether back-end tiering is enabled.",
												},
												"frontend_tiering": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies whether Front End Tiering Enabled.",
												},
												"max_retention": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
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
							MaxItems:    1,
							Optional:    true,
							Description: "ClassificationInfo holds information about the Datahawk Classification subscription such as if it is active or not.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_active": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_free_trial": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"quantity": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies the quantity of the subscription.",
									},
									"start_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the start date of the subscription.",
									},
									"tiering": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the tiering info.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"backend_tiering": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies whether back-end tiering is enabled.",
												},
												"frontend_tiering": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies whether Front End Tiering Enabled.",
												},
												"max_retention": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
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
							MaxItems:    1,
							Optional:    true,
							Description: "FortKnoxInfo holds information about the Fortknox Azure or Azure FreeTrial or AwsCold or AwsCold FreeTrial subscription such as if it is active.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_active": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_free_trial": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"quantity": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies the quantity of the subscription.",
									},
									"start_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the start date of the subscription.",
									},
								},
							},
						},
						"fort_knox_azure_hot": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "FortKnoxInfo holds information about the Fortknox Azure or Azure FreeTrial or AwsCold or AwsCold FreeTrial subscription such as if it is active.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_active": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_free_trial": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"quantity": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies the quantity of the subscription.",
									},
									"start_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the start date of the subscription.",
									},
								},
							},
						},
						"fort_knox_cold": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "FortKnoxInfo holds information about the Fortknox Azure or Azure FreeTrial or AwsCold or AwsCold FreeTrial subscription such as if it is active.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_active": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_free_trial": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"quantity": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies the quantity of the subscription.",
									},
									"start_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the start date of the subscription.",
									},
								},
							},
						},
						"ransomware": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "RansomwareInfo holds information about the FortKnox/FortKnoxFreeTrial subscription such as if it is active or not.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_active": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_free_trial": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"quantity": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies the quantity of the subscription.",
									},
									"start_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the start date of the subscription.",
									},
								},
							},
						},
						"site_continuity": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "SiteContinuityInfo holds information about the Site Continuity subscription such as if it is active or not.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_active": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_free_trial": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"start_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the start date of the subscription.",
									},
								},
							},
						},
						"threat_protection": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "ThreatProtectionInfo holds information about the Datahawk ThreatProtection subscription such as if it is active or not.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_active": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"is_free_trial": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies the end date of the subscription.",
									},
									"start_date": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the start date of the subscription.",
									},
								},
							},
						},
					},
				},
			},
			"tenant_accesses": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specfies the Tenant Access for MCM User.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_identifiers": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the list of clusters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster_id": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies the id of the cluster.",
									},
									"cluster_incarnation_id": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies the incarnation id of the cluster.",
									},
								},
							},
						},
						"created_time_msecs": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the epoch time in milliseconds when the tenant access was created.",
						},
						"effective_time_msecs": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the epoch time in milliseconds when the tenant access becomes effective. Until that time, the user cannot log in.",
						},
						"expired_time_msecs": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the epoch time in milliseconds when the tenant access becomes expired. After that, the user cannot log in.",
						},
						"is_access_active": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "IsAccessActive specifies whether or not a tenant access is active, or has been deactivated by the customer. The default behavior is 'true'.",
						},
						"is_active": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether or not the tenant is active.",
						},
						"is_deleted": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether or not the tenant is deleted.",
						},
						"last_updated_time_msecs": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the epoch time in milliseconds when the tenant access was last modified.",
						},
						"roles": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the Cohesity roles to associate with the user such as such as 'Admin', 'Ops' or 'View'.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"tenant_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the tenant id.",
						},
						"tenant_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the tenant name.",
						},
						"tenant_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the MCM tenant type. 'Dmaas' implies tenant type is DMaaS. 'Mcm' implies tenant is Mcm Cluster tenant.",
						},
					},
				},
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// ForceNew:    true
				Description: "Specifies the effective Tenant ID of the user.",
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				// ForceNew:    true
				Description: "Specifies the login name of the user.",
			},
			"group_roles": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the Cohesity roles to associate with the user' group. These roles can only be edited from group.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"is_account_mfa_enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Specifies if MFA is enabled for the Helios Account.",
			},
			"is_cluster_mfa_enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Specifies if MFA is enabled on cluster.",
			},
		},
	}
}

func checkDiffResourceIbmBackupRecoveryConnectorUpdateUser(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	// oldId, _ := d.GetChange("x_ibm_tenant_id")
	// if oldId == "" {
	// 	return nil
	// }

	// return if it's a new resource
	if d.Id() == "" {
		return nil
	}

	for fieldName := range ResourceIbmBackupRecoveryConnectorUpdateUser().Schema {
		if d.HasChange(fieldName) {
			return fmt.Errorf("[ERROR] Resource ibm_backup_recovery_connector_update_user cannot be updated :%s", fieldName)
		}
	}
	return nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "authentication_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "kAuthAd, kAuthGoogle, kAuthLocal, kAuthSalesforce, kAuthSso",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_backup_recovery_connector_update_user", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmBackupRecoveryConnectorUpdateUserCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1Connector()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateUserOptions := &backuprecoveryv1.UpdateUserOptions{}

	updateUserOptions.SetSessionName(d.Get("session_name_cookie").(string))

	if _, ok := d.GetOk("ad_user_info"); ok {
		adUserInfoModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToAdUserInfo(d.Get("ad_user_info.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "create", "parse-ad_user_info").GetDiag()
		}
		updateUserOptions.SetAdUserInfo(adUserInfoModel)
	}
	if _, ok := d.GetOk("additional_group_names"); ok {
		var additionalGroupNames []string
		for _, v := range d.Get("additional_group_names").([]interface{}) {
			additionalGroupNamesItem := v.(string)
			additionalGroupNames = append(additionalGroupNames, additionalGroupNamesItem)
		}
		updateUserOptions.SetAdditionalGroupNames(additionalGroupNames)
	}
	if _, ok := d.GetOk("allow_dso_modify"); ok {
		updateUserOptions.SetAllowDsoModify(d.Get("allow_dso_modify").(bool))
	}
	if _, ok := d.GetOk("audit_log_settings"); ok {
		auditLogSettingsModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToAuditLogSettings(d.Get("audit_log_settings.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "create", "parse-audit_log_settings").GetDiag()
		}
		updateUserOptions.SetAuditLogSettings(auditLogSettingsModel)
	}
	if _, ok := d.GetOk("authentication_type"); ok {
		updateUserOptions.SetAuthenticationType(d.Get("authentication_type").(string))
	}
	if _, ok := d.GetOk("cluster_identifiers"); ok {
		var clusterIdentifiers []backuprecoveryv1.UserClusterIdentifier
		for _, v := range d.Get("cluster_identifiers").([]interface{}) {
			value := v.(map[string]interface{})
			clusterIdentifiersItem, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToUserClusterIdentifier(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "create", "parse-cluster_identifiers").GetDiag()
			}
			clusterIdentifiers = append(clusterIdentifiers, *clusterIdentifiersItem)
		}
		updateUserOptions.SetClusterIdentifiers(clusterIdentifiers)
	}
	if _, ok := d.GetOk("created_time_msecs"); ok {
		updateUserOptions.SetCreatedTimeMsecs(int64(d.Get("created_time_msecs").(int)))
	}
	if _, ok := d.GetOk("current_password"); ok {
		updateUserOptions.SetCurrentPassword(d.Get("current_password").(string))
	}
	if _, ok := d.GetOk("description"); ok {
		updateUserOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("domain"); ok {
		updateUserOptions.SetDomain(d.Get("domain").(string))
	}
	if _, ok := d.GetOk("effective_time_msecs"); ok {
		updateUserOptions.SetEffectiveTimeMsecs(int64(d.Get("effective_time_msecs").(int)))
	}
	if _, ok := d.GetOk("email_address"); ok {
		updateUserOptions.SetEmailAddress(d.Get("email_address").(string))
	}
	if _, ok := d.GetOk("expired_time_msecs"); ok {
		updateUserOptions.SetExpiredTimeMsecs(int64(d.Get("expired_time_msecs").(int)))
	}
	if _, ok := d.GetOk("force_password_change"); ok {
		updateUserOptions.SetForcePasswordChange(d.Get("force_password_change").(bool))
	}
	if _, ok := d.GetOk("google_account"); ok {
		googleAccountModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToGoogleAccountInfo(d.Get("google_account.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "create", "parse-google_account").GetDiag()
		}
		updateUserOptions.SetGoogleAccount(googleAccountModel)
	}
	if _, ok := d.GetOk("idp_user_info"); ok {
		idpUserInfoModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToIdpUserInfo(d.Get("idp_user_info.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "create", "parse-idp_user_info").GetDiag()
		}
		updateUserOptions.SetIdpUserInfo(idpUserInfoModel)
	}
	if _, ok := d.GetOk("intercom_messenger_token"); ok {
		updateUserOptions.SetIntercomMessengerToken(d.Get("intercom_messenger_token").(string))
	}
	if _, ok := d.GetOk("is_account_locked"); ok {
		updateUserOptions.SetIsAccountLocked(d.Get("is_account_locked").(bool))
	}
	if _, ok := d.GetOk("is_active"); ok {
		updateUserOptions.SetIsActive(d.Get("is_active").(bool))
	}
	if _, ok := d.GetOk("last_successful_login_time_msecs"); ok {
		updateUserOptions.SetLastSuccessfulLoginTimeMsecs(int64(d.Get("last_successful_login_time_msecs").(int)))
	}
	if _, ok := d.GetOk("last_updated_time_msecs"); ok {
		updateUserOptions.SetLastUpdatedTimeMsecs(int64(d.Get("last_updated_time_msecs").(int)))
	}
	if _, ok := d.GetOk("mfa_info"); ok {
		mfaInfoModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToMfaInfo(d.Get("mfa_info.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "create", "parse-mfa_info").GetDiag()
		}
		updateUserOptions.SetMfaInfo(mfaInfoModel)
	}
	if _, ok := d.GetOk("mfa_methods"); ok {
		var mfaMethods []string
		for _, v := range d.Get("mfa_methods").([]interface{}) {
			mfaMethodsItem := v.(string)
			mfaMethods = append(mfaMethods, mfaMethodsItem)
		}
		updateUserOptions.SetMfaMethods(mfaMethods)
	}
	if _, ok := d.GetOk("object_class"); ok {
		updateUserOptions.SetObjectClass(d.Get("object_class").(string))
	}
	if _, ok := d.GetOk("org_membership"); ok {
		var orgMembership []backuprecoveryv1.TenantConfig
		for _, v := range d.Get("org_membership").([]interface{}) {
			value := v.(map[string]interface{})
			orgMembershipItem, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToTenantConfig(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "create", "parse-org_membership").GetDiag()
			}
			orgMembership = append(orgMembership, *orgMembershipItem)
		}
		updateUserOptions.SetOrgMembership(orgMembership)
	}
	if _, ok := d.GetOk("password"); ok {
		updateUserOptions.SetPassword(d.Get("password").(string))
	}
	if _, ok := d.GetOk("preferences"); ok {
		preferencesModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToUsersPreferences(d.Get("preferences.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "create", "parse-preferences").GetDiag()
		}
		updateUserOptions.SetPreferences(preferencesModel)
	}
	if _, ok := d.GetOk("previous_login_time_msecs"); ok {
		updateUserOptions.SetPreviousLoginTimeMsecs(int64(d.Get("previous_login_time_msecs").(int)))
	}
	if _, ok := d.GetOk("primary_group_name"); ok {
		updateUserOptions.SetPrimaryGroupName(d.Get("primary_group_name").(string))
	}
	if _, ok := d.GetOk("privilege_ids"); ok {
		var privilegeIds []string
		for _, v := range d.Get("privilege_ids").([]interface{}) {
			privilegeIdsItem := v.(string)
			privilegeIds = append(privilegeIds, privilegeIdsItem)
		}
		updateUserOptions.SetPrivilegeIds(privilegeIds)
	}
	if _, ok := d.GetOk("profiles"); ok {
		var profiles []backuprecoveryv1.UserProfile
		for _, v := range d.Get("profiles").([]interface{}) {
			value := v.(map[string]interface{})
			profilesItem, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToUserProfile(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "create", "parse-profiles").GetDiag()
			}
			profiles = append(profiles, *profilesItem)
		}
		updateUserOptions.SetProfiles(profiles)
	}
	if _, ok := d.GetOk("restricted"); ok {
		updateUserOptions.SetRestricted(d.Get("restricted").(bool))
	}
	if _, ok := d.GetOk("roles"); ok {
		var roles []string
		for _, v := range d.Get("roles").([]interface{}) {
			rolesItem := v.(string)
			roles = append(roles, rolesItem)
		}
		updateUserOptions.SetRoles(roles)
	}
	if _, ok := d.GetOk("s3_access_key_id"); ok {
		updateUserOptions.SetS3AccessKeyID(d.Get("s3_access_key_id").(string))
	}
	if _, ok := d.GetOk("s3_account_id"); ok {
		updateUserOptions.SetS3AccountID(d.Get("s3_account_id").(string))
	}
	if _, ok := d.GetOk("s3_secret_key"); ok {
		updateUserOptions.SetS3SecretKey(d.Get("s3_secret_key").(string))
	}
	if _, ok := d.GetOk("salesforce_account"); ok {
		salesforceAccountModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToSalesforceAccountInfo(d.Get("salesforce_account.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "create", "parse-salesforce_account").GetDiag()
		}
		updateUserOptions.SetSalesforceAccount(salesforceAccountModel)
	}
	if _, ok := d.GetOk("sid"); ok {
		updateUserOptions.SetSid(d.Get("sid").(string))
	}
	if _, ok := d.GetOk("spog_context"); ok {
		spogContextModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToSpogContext(d.Get("spog_context.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "create", "parse-spog_context").GetDiag()
		}
		updateUserOptions.SetSpogContext(spogContextModel)
	}
	if _, ok := d.GetOk("subscription_info"); ok {
		subscriptionInfoModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToSubscriptionInfo(d.Get("subscription_info.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "create", "parse-subscription_info").GetDiag()
		}
		updateUserOptions.SetSubscriptionInfo(subscriptionInfoModel)
	}
	if _, ok := d.GetOk("tenant_accesses"); ok {
		var tenantAccesses []backuprecoveryv1.TenantAccesses
		for _, v := range d.Get("tenant_accesses").([]interface{}) {
			value := v.(map[string]interface{})
			tenantAccessesItem, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToTenantAccesses(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "create", "parse-tenant_accesses").GetDiag()
			}
			tenantAccesses = append(tenantAccesses, *tenantAccessesItem)
		}
		updateUserOptions.SetTenantAccesses(tenantAccesses)
	}
	if _, ok := d.GetOk("tenant_id"); ok {
		updateUserOptions.SetTenantID(d.Get("tenant_id").(string))
	}
	if _, ok := d.GetOk("username"); ok {
		updateUserOptions.SetUsername(d.Get("username").(string))
	}

	userDetails, _, err := backupRecoveryClient.UpdateUserWithContext(context, updateUserOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateUserWithContext failed: %s", err.Error()), "ibm_backup_recovery_connector_update_user", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	id := fmt.Sprintf("%s:%s", *userDetails.Sid, *userDetails.Username)
	d.SetId(id)

	return resourceIbmBackupRecoveryConnectorUpdateUserRead(context, d, meta)
}

func resourceIbmBackupRecoveryConnectorUpdateUserRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1Connector()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	username := strings.Split(d.Id(), ":")[1]
	getUsersOptions := &backuprecoveryv1.GetUsersOptions{Usernames: []string{username}}

	getUsersOptions.SetSessionName(d.Get("session_name_cookie").(string))

	getUsersResponse, response, err := backupRecoveryClient.GetUsersWithContext(context, getUsersOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetUsersWithContext failed: %s", err.Error()), "ibm_backup_recovery_connector_update_user", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(getUsersResponse[0].AdUserInfo) {
		adUserInfoMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserAdUserInfoToMap(getUsersResponse[0].AdUserInfo)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "ad_user_info-to-map").GetDiag()
		}
		if err = d.Set("ad_user_info", []map[string]interface{}{adUserInfoMap}); err != nil {
			err = fmt.Errorf("Error setting ad_user_info: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-ad_user_info").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].AdditionalGroupNames) {
		if err = d.Set("additional_group_names", getUsersResponse[0].AdditionalGroupNames); err != nil {
			err = fmt.Errorf("Error setting additional_group_names: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-additional_group_names").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].AllowDsoModify) {
		if err = d.Set("allow_dso_modify", getUsersResponse[0].AllowDsoModify); err != nil {
			err = fmt.Errorf("Error setting allow_dso_modify: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-allow_dso_modify").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].AuditLogSettings) {
		auditLogSettingsMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserAuditLogSettingsToMap(getUsersResponse[0].AuditLogSettings)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "audit_log_settings-to-map").GetDiag()
		}
		if err = d.Set("audit_log_settings", []map[string]interface{}{auditLogSettingsMap}); err != nil {
			err = fmt.Errorf("Error setting audit_log_settings: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-audit_log_settings").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].AuthenticationType) {
		if err = d.Set("authentication_type", getUsersResponse[0].AuthenticationType); err != nil {
			err = fmt.Errorf("Error setting authentication_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-authentication_type").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].ClusterIdentifiers) {
		clusterIdentifiers := []map[string]interface{}{}
		for _, clusterIdentifiersItem := range getUsersResponse[0].ClusterIdentifiers {
			clusterIdentifiersItemMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserUserClusterIdentifierToMap(&clusterIdentifiersItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "cluster_identifiers-to-map").GetDiag()
			}
			clusterIdentifiers = append(clusterIdentifiers, clusterIdentifiersItemMap)
		}
		if err = d.Set("cluster_identifiers", clusterIdentifiers); err != nil {
			err = fmt.Errorf("Error setting cluster_identifiers: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-cluster_identifiers").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].CreatedTimeMsecs) {
		if err = d.Set("created_time_msecs", flex.IntValue(getUsersResponse[0].CreatedTimeMsecs)); err != nil {
			err = fmt.Errorf("Error setting created_time_msecs: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-created_time_msecs").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].CurrentPassword) {
		if err = d.Set("current_password", getUsersResponse[0].CurrentPassword); err != nil {
			err = fmt.Errorf("Error setting current_password: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-current_password").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].Description) {
		if err = d.Set("description", getUsersResponse[0].Description); err != nil {
			err = fmt.Errorf("Error setting description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-description").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].Domain) {
		if err = d.Set("domain", getUsersResponse[0].Domain); err != nil {
			err = fmt.Errorf("Error setting domain: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-domain").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].EffectiveTimeMsecs) {
		if err = d.Set("effective_time_msecs", flex.IntValue(getUsersResponse[0].EffectiveTimeMsecs)); err != nil {
			err = fmt.Errorf("Error setting effective_time_msecs: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-effective_time_msecs").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].EmailAddress) {
		if err = d.Set("email_address", getUsersResponse[0].EmailAddress); err != nil {
			err = fmt.Errorf("Error setting email_address: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-email_address").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].ExpiredTimeMsecs) {
		if err = d.Set("expired_time_msecs", flex.IntValue(getUsersResponse[0].ExpiredTimeMsecs)); err != nil {
			err = fmt.Errorf("Error setting expired_time_msecs: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-expired_time_msecs").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].ForcePasswordChange) {
		if err = d.Set("force_password_change", getUsersResponse[0].ForcePasswordChange); err != nil {
			err = fmt.Errorf("Error setting force_password_change: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-force_password_change").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].GoogleAccount) {
		googleAccountMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserGoogleAccountInfoToMap(getUsersResponse[0].GoogleAccount)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "google_account-to-map").GetDiag()
		}
		if err = d.Set("google_account", []map[string]interface{}{googleAccountMap}); err != nil {
			err = fmt.Errorf("Error setting google_account: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-google_account").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].IdpUserInfo) {
		idpUserInfoMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserIdpUserInfoToMap(getUsersResponse[0].IdpUserInfo)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "idp_user_info-to-map").GetDiag()
		}
		if err = d.Set("idp_user_info", []map[string]interface{}{idpUserInfoMap}); err != nil {
			err = fmt.Errorf("Error setting idp_user_info: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-idp_user_info").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].IntercomMessengerToken) {
		if err = d.Set("intercom_messenger_token", getUsersResponse[0].IntercomMessengerToken); err != nil {
			err = fmt.Errorf("Error setting intercom_messenger_token: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-intercom_messenger_token").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].IsAccountLocked) {
		if err = d.Set("is_account_locked", getUsersResponse[0].IsAccountLocked); err != nil {
			err = fmt.Errorf("Error setting is_account_locked: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-is_account_locked").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].IsActive) {
		if err = d.Set("is_active", getUsersResponse[0].IsActive); err != nil {
			err = fmt.Errorf("Error setting is_active: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-is_active").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].LastSuccessfulLoginTimeMsecs) {
		if err = d.Set("last_successful_login_time_msecs", flex.IntValue(getUsersResponse[0].LastSuccessfulLoginTimeMsecs)); err != nil {
			err = fmt.Errorf("Error setting last_successful_login_time_msecs: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-last_successful_login_time_msecs").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].LastUpdatedTimeMsecs) {
		if err = d.Set("last_updated_time_msecs", flex.IntValue(getUsersResponse[0].LastUpdatedTimeMsecs)); err != nil {
			err = fmt.Errorf("Error setting last_updated_time_msecs: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-last_updated_time_msecs").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].MfaInfo) {
		mfaInfoMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserMfaInfoToMap(getUsersResponse[0].MfaInfo)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "mfa_info-to-map").GetDiag()
		}
		if err = d.Set("mfa_info", []map[string]interface{}{mfaInfoMap}); err != nil {
			err = fmt.Errorf("Error setting mfa_info: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-mfa_info").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].MfaMethods) {
		if err = d.Set("mfa_methods", getUsersResponse[0].MfaMethods); err != nil {
			err = fmt.Errorf("Error setting mfa_methods: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-mfa_methods").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].ObjectClass) {
		if err = d.Set("object_class", getUsersResponse[0].ObjectClass); err != nil {
			err = fmt.Errorf("Error setting object_class: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-object_class").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].OrgMembership) {
		orgMembership := []map[string]interface{}{}
		for _, orgMembershipItem := range getUsersResponse[0].OrgMembership {
			orgMembershipItemMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserTenantConfigToMap(&orgMembershipItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "org_membership-to-map").GetDiag()
			}
			orgMembership = append(orgMembership, orgMembershipItemMap)
		}
		if err = d.Set("org_membership", orgMembership); err != nil {
			err = fmt.Errorf("Error setting org_membership: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-org_membership").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].Password) {
		if err = d.Set("password", getUsersResponse[0].Password); err != nil {
			err = fmt.Errorf("Error setting password: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-password").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].Preferences) {
		preferencesMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserUsersPreferencesToMap(getUsersResponse[0].Preferences)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "preferences-to-map").GetDiag()
		}
		if err = d.Set("preferences", []map[string]interface{}{preferencesMap}); err != nil {
			err = fmt.Errorf("Error setting preferences: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-preferences").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].PreviousLoginTimeMsecs) {
		if err = d.Set("previous_login_time_msecs", flex.IntValue(getUsersResponse[0].PreviousLoginTimeMsecs)); err != nil {
			err = fmt.Errorf("Error setting previous_login_time_msecs: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-previous_login_time_msecs").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].PrimaryGroupName) {
		if err = d.Set("primary_group_name", getUsersResponse[0].PrimaryGroupName); err != nil {
			err = fmt.Errorf("Error setting primary_group_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-primary_group_name").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].PrivilegeIds) {
		if err = d.Set("privilege_ids", getUsersResponse[0].PrivilegeIds); err != nil {
			err = fmt.Errorf("Error setting privilege_ids: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-privilege_ids").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].Profiles) {
		profiles := []map[string]interface{}{}
		for _, profilesItem := range getUsersResponse[0].Profiles {
			profilesItemMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserUserProfileToMap(&profilesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "profiles-to-map").GetDiag()
			}
			profiles = append(profiles, profilesItemMap)
		}
		if err = d.Set("profiles", profiles); err != nil {
			err = fmt.Errorf("Error setting profiles: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-profiles").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].Restricted) {
		if err = d.Set("restricted", getUsersResponse[0].Restricted); err != nil {
			err = fmt.Errorf("Error setting restricted: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-restricted").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].Roles) {
		if err = d.Set("roles", getUsersResponse[0].Roles); err != nil {
			err = fmt.Errorf("Error setting roles: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-roles").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].S3AccessKeyID) {
		if err = d.Set("s3_access_key_id", getUsersResponse[0].S3AccessKeyID); err != nil {
			err = fmt.Errorf("Error setting s3_access_key_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-s3_access_key_id").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].S3AccountID) {
		if err = d.Set("s3_account_id", getUsersResponse[0].S3AccountID); err != nil {
			err = fmt.Errorf("Error setting s3_account_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-s3_account_id").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].S3SecretKey) {
		if err = d.Set("s3_secret_key", getUsersResponse[0].S3SecretKey); err != nil {
			err = fmt.Errorf("Error setting s3_secret_key: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-s3_secret_key").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].SalesforceAccount) {
		salesforceAccountMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserSalesforceAccountInfoToMap(getUsersResponse[0].SalesforceAccount)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "salesforce_account-to-map").GetDiag()
		}
		if err = d.Set("salesforce_account", []map[string]interface{}{salesforceAccountMap}); err != nil {
			err = fmt.Errorf("Error setting salesforce_account: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-salesforce_account").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].Sid) {
		if err = d.Set("sid", getUsersResponse[0].Sid); err != nil {
			err = fmt.Errorf("Error setting sid: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-sid").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].SpogContext) {
		spogContextMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserSpogContextToMap(getUsersResponse[0].SpogContext)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "spog_context-to-map").GetDiag()
		}
		if err = d.Set("spog_context", []map[string]interface{}{spogContextMap}); err != nil {
			err = fmt.Errorf("Error setting spog_context: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-spog_context").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].SubscriptionInfo) {
		subscriptionInfoMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserSubscriptionInfoToMap(getUsersResponse[0].SubscriptionInfo)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "subscription_info-to-map").GetDiag()
		}
		if err = d.Set("subscription_info", []map[string]interface{}{subscriptionInfoMap}); err != nil {
			err = fmt.Errorf("Error setting subscription_info: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-subscription_info").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].TenantAccesses) {
		tenantAccesses := []map[string]interface{}{}
		for _, tenantAccessesItem := range getUsersResponse[0].TenantAccesses {
			tenantAccessesItemMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserTenantAccessesToMap(&tenantAccessesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "tenant_accesses-to-map").GetDiag()
			}
			tenantAccesses = append(tenantAccesses, tenantAccessesItemMap)
		}
		if err = d.Set("tenant_accesses", tenantAccesses); err != nil {
			err = fmt.Errorf("Error setting tenant_accesses: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-tenant_accesses").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].TenantID) {
		if err = d.Set("tenant_id", getUsersResponse[0].TenantID); err != nil {
			err = fmt.Errorf("Error setting tenant_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-tenant_id").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].Username) {
		if err = d.Set("username", getUsersResponse[0].Username); err != nil {
			err = fmt.Errorf("Error setting username: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-username").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].GroupRoles) {
		if err = d.Set("group_roles", getUsersResponse[0].GroupRoles); err != nil {
			err = fmt.Errorf("Error setting group_roles: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-group_roles").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].IsAccountMfaEnabled) {
		if err = d.Set("is_account_mfa_enabled", getUsersResponse[0].IsAccountMfaEnabled); err != nil {
			err = fmt.Errorf("Error setting is_account_mfa_enabled: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-is_account_mfa_enabled").GetDiag()
		}
	}
	if !core.IsNil(getUsersResponse[0].IsClusterMfaEnabled) {
		if err = d.Set("is_cluster_mfa_enabled", getUsersResponse[0].IsClusterMfaEnabled); err != nil {
			err = fmt.Errorf("Error setting is_cluster_mfa_enabled: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_update_user", "read", "set-is_cluster_mfa_enabled").GetDiag()
		}
	}

	return nil
}

func resourceIbmBackupRecoveryConnectorUpdateUserDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.

	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Delete Not Supported",
		Detail:   "The resource definition will be only be removed from the terraform statefile. This resource cannot be deleted from the backend. ",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToAdUserInfo(modelMap map[string]interface{}) (*backuprecoveryv1.AdUserInfo, error) {
	model := &backuprecoveryv1.AdUserInfo{}
	if modelMap["group_sids"] != nil {
		groupSids := []string{}
		for _, groupSidsItem := range modelMap["group_sids"].([]interface{}) {
			groupSids = append(groupSids, groupSidsItem.(string))
		}
		model.GroupSids = groupSids
	}
	if modelMap["groups"] != nil {
		groups := []string{}
		for _, groupsItem := range modelMap["groups"].([]interface{}) {
			groups = append(groups, groupsItem.(string))
		}
		model.Groups = groups
	}
	if modelMap["is_floating_user"] != nil {
		model.IsFloatingUser = core.BoolPtr(modelMap["is_floating_user"].(bool))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToAuditLogSettings(modelMap map[string]interface{}) (*backuprecoveryv1.AuditLogSettings, error) {
	model := &backuprecoveryv1.AuditLogSettings{}
	if modelMap["read_logging"] != nil {
		model.ReadLogging = core.BoolPtr(modelMap["read_logging"].(bool))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToUserClusterIdentifier(modelMap map[string]interface{}) (*backuprecoveryv1.UserClusterIdentifier, error) {
	model := &backuprecoveryv1.UserClusterIdentifier{}
	if modelMap["cluster_id"] != nil {
		model.ClusterID = core.Int64Ptr(int64(modelMap["cluster_id"].(int)))
	}
	if modelMap["cluster_incarnation_id"] != nil {
		model.ClusterIncarnationID = core.Int64Ptr(int64(modelMap["cluster_incarnation_id"].(int)))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToGoogleAccountInfo(modelMap map[string]interface{}) (*backuprecoveryv1.GoogleAccountInfo, error) {
	model := &backuprecoveryv1.GoogleAccountInfo{}
	if modelMap["account_id"] != nil && modelMap["account_id"].(string) != "" {
		model.AccountID = core.StringPtr(modelMap["account_id"].(string))
	}
	if modelMap["user_id"] != nil && modelMap["user_id"].(string) != "" {
		model.UserID = core.StringPtr(modelMap["user_id"].(string))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToIdpUserInfo(modelMap map[string]interface{}) (*backuprecoveryv1.IdpUserInfo, error) {
	model := &backuprecoveryv1.IdpUserInfo{}
	if modelMap["group_sids"] != nil {
		groupSids := []string{}
		for _, groupSidsItem := range modelMap["group_sids"].([]interface{}) {
			groupSids = append(groupSids, groupSidsItem.(string))
		}
		model.GroupSids = groupSids
	}
	if modelMap["groups"] != nil {
		groups := []string{}
		for _, groupsItem := range modelMap["groups"].([]interface{}) {
			groups = append(groups, groupsItem.(string))
		}
		model.Groups = groups
	}
	if modelMap["idp_id"] != nil {
		model.IdpID = core.Int64Ptr(int64(modelMap["idp_id"].(int)))
	}
	if modelMap["is_floating_user"] != nil {
		model.IsFloatingUser = core.BoolPtr(modelMap["is_floating_user"].(bool))
	}
	if modelMap["issuer_id"] != nil && modelMap["issuer_id"].(string) != "" {
		model.IssuerID = core.StringPtr(modelMap["issuer_id"].(string))
	}
	if modelMap["user_id"] != nil && modelMap["user_id"].(string) != "" {
		model.UserID = core.StringPtr(modelMap["user_id"].(string))
	}
	if modelMap["vendor"] != nil && modelMap["vendor"].(string) != "" {
		model.Vendor = core.StringPtr(modelMap["vendor"].(string))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToMfaInfo(modelMap map[string]interface{}) (*backuprecoveryv1.MfaInfo, error) {
	model := &backuprecoveryv1.MfaInfo{}
	if modelMap["is_email_otp_setup_done"] != nil {
		model.IsEmailOtpSetupDone = core.BoolPtr(modelMap["is_email_otp_setup_done"].(bool))
	}
	if modelMap["is_totp_setup_done"] != nil {
		model.IsTotpSetupDone = core.BoolPtr(modelMap["is_totp_setup_done"].(bool))
	}
	if modelMap["is_user_exempt_from_mfa"] != nil {
		model.IsUserExemptFromMfa = core.BoolPtr(modelMap["is_user_exempt_from_mfa"].(bool))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToTenantConfig(modelMap map[string]interface{}) (*backuprecoveryv1.TenantConfig, error) {
	model := &backuprecoveryv1.TenantConfig{}
	if modelMap["bifrost_enabled"] != nil {
		model.BifrostEnabled = core.BoolPtr(modelMap["bifrost_enabled"].(bool))
	}
	if modelMap["is_managed_on_helios"] != nil {
		model.IsManagedOnHelios = core.BoolPtr(modelMap["is_managed_on_helios"].(bool))
	}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["restricted"] != nil {
		model.Restricted = core.BoolPtr(modelMap["restricted"].(bool))
	}
	if modelMap["roles"] != nil {
		roles := []string{}
		for _, rolesItem := range modelMap["roles"].([]interface{}) {
			roles = append(roles, rolesItem.(string))
		}
		model.Roles = roles
	}
	if modelMap["tenant_id"] != nil && modelMap["tenant_id"].(string) != "" {
		model.TenantID = core.StringPtr(modelMap["tenant_id"].(string))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToUsersPreferences(modelMap map[string]interface{}) (*backuprecoveryv1.UsersPreferences, error) {
	model := &backuprecoveryv1.UsersPreferences{}
	if modelMap["locale"] != nil && modelMap["locale"].(string) != "" {
		model.Locale = core.StringPtr(modelMap["locale"].(string))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToUserProfile(modelMap map[string]interface{}) (*backuprecoveryv1.UserProfile, error) {
	model := &backuprecoveryv1.UserProfile{}
	if modelMap["cluster_identifiers"] != nil {
		clusterIdentifiers := []backuprecoveryv1.UserClusterIdentifier{}
		for _, clusterIdentifiersItem := range modelMap["cluster_identifiers"].([]interface{}) {
			clusterIdentifiersItemModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToUserClusterIdentifier(clusterIdentifiersItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			clusterIdentifiers = append(clusterIdentifiers, *clusterIdentifiersItemModel)
		}
		model.ClusterIdentifiers = clusterIdentifiers
	}
	if modelMap["is_active"] != nil {
		model.IsActive = core.BoolPtr(modelMap["is_active"].(bool))
	}
	if modelMap["is_deleted"] != nil {
		model.IsDeleted = core.BoolPtr(modelMap["is_deleted"].(bool))
	}
	if modelMap["region_ids"] != nil {
		regionIds := []string{}
		for _, regionIdsItem := range modelMap["region_ids"].([]interface{}) {
			regionIds = append(regionIds, regionIdsItem.(string))
		}
		model.RegionIds = regionIds
	}
	if modelMap["tenant_id"] != nil && modelMap["tenant_id"].(string) != "" {
		model.TenantID = core.StringPtr(modelMap["tenant_id"].(string))
	}
	if modelMap["tenant_name"] != nil && modelMap["tenant_name"].(string) != "" {
		model.TenantName = core.StringPtr(modelMap["tenant_name"].(string))
	}
	if modelMap["tenant_type"] != nil && modelMap["tenant_type"].(string) != "" {
		model.TenantType = core.StringPtr(modelMap["tenant_type"].(string))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToSalesforceAccountInfo(modelMap map[string]interface{}) (*backuprecoveryv1.SalesforceAccountInfo, error) {
	model := &backuprecoveryv1.SalesforceAccountInfo{}
	if modelMap["account_id"] != nil && modelMap["account_id"].(string) != "" {
		model.AccountID = core.StringPtr(modelMap["account_id"].(string))
	}
	if modelMap["helios_access_grant_status"] != nil && modelMap["helios_access_grant_status"].(string) != "" {
		model.HeliosAccessGrantStatus = core.StringPtr(modelMap["helios_access_grant_status"].(string))
	}
	if modelMap["is_d_gaa_s_user"] != nil {
		model.IsDGaaSUser = core.BoolPtr(modelMap["is_d_gaa_s_user"].(bool))
	}
	if modelMap["is_d_maa_s_user"] != nil {
		model.IsDMaaSUser = core.BoolPtr(modelMap["is_d_maa_s_user"].(bool))
	}
	if modelMap["is_d_raa_s_user"] != nil {
		model.IsDRaaSUser = core.BoolPtr(modelMap["is_d_raa_s_user"].(bool))
	}
	if modelMap["is_r_paa_s_user"] != nil {
		model.IsRPaaSUser = core.BoolPtr(modelMap["is_r_paa_s_user"].(bool))
	}
	if modelMap["is_sales_user"] != nil {
		model.IsSalesUser = core.BoolPtr(modelMap["is_sales_user"].(bool))
	}
	if modelMap["is_support_user"] != nil {
		model.IsSupportUser = core.BoolPtr(modelMap["is_support_user"].(bool))
	}
	if modelMap["user_id"] != nil && modelMap["user_id"].(string) != "" {
		model.UserID = core.StringPtr(modelMap["user_id"].(string))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToSpogContext(modelMap map[string]interface{}) (*backuprecoveryv1.SpogContext, error) {
	model := &backuprecoveryv1.SpogContext{}
	if modelMap["primary_cluster_id"] != nil {
		model.PrimaryClusterID = core.Int64Ptr(int64(modelMap["primary_cluster_id"].(int)))
	}
	if modelMap["primary_cluster_user_sid"] != nil && modelMap["primary_cluster_user_sid"].(string) != "" {
		model.PrimaryClusterUserSid = core.StringPtr(modelMap["primary_cluster_user_sid"].(string))
	}
	if modelMap["primary_cluster_username"] != nil && modelMap["primary_cluster_username"].(string) != "" {
		model.PrimaryClusterUsername = core.StringPtr(modelMap["primary_cluster_username"].(string))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToSubscriptionInfo(modelMap map[string]interface{}) (*backuprecoveryv1.SubscriptionInfo, error) {
	model := &backuprecoveryv1.SubscriptionInfo{}
	if modelMap["classification"] != nil && len(modelMap["classification"].([]interface{})) > 0 {
		ClassificationModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToClassificationInfo(modelMap["classification"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Classification = ClassificationModel
	}
	if modelMap["data_protect"] != nil && len(modelMap["data_protect"].([]interface{})) > 0 {
		DataProtectModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToDataProtectInfo(modelMap["data_protect"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.DataProtect = DataProtectModel
	}
	if modelMap["data_protect_azure"] != nil && len(modelMap["data_protect_azure"].([]interface{})) > 0 {
		DataProtectAzureModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToDataProtectAzureInfo(modelMap["data_protect_azure"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.DataProtectAzure = DataProtectAzureModel
	}
	if modelMap["fort_knox_azure_cool"] != nil && len(modelMap["fort_knox_azure_cool"].([]interface{})) > 0 {
		FortKnoxAzureCoolModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToFortKnoxInfo(modelMap["fort_knox_azure_cool"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.FortKnoxAzureCool = FortKnoxAzureCoolModel
	}
	if modelMap["fort_knox_azure_hot"] != nil && len(modelMap["fort_knox_azure_hot"].([]interface{})) > 0 {
		FortKnoxAzureHotModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToFortKnoxInfo(modelMap["fort_knox_azure_hot"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.FortKnoxAzureHot = FortKnoxAzureHotModel
	}
	if modelMap["fort_knox_cold"] != nil && len(modelMap["fort_knox_cold"].([]interface{})) > 0 {
		FortKnoxColdModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToFortKnoxInfo(modelMap["fort_knox_cold"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.FortKnoxCold = FortKnoxColdModel
	}
	if modelMap["ransomware"] != nil && len(modelMap["ransomware"].([]interface{})) > 0 {
		RansomwareModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToFortKnoxInfo(modelMap["ransomware"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Ransomware = RansomwareModel
	}
	if modelMap["site_continuity"] != nil && len(modelMap["site_continuity"].([]interface{})) > 0 {
		SiteContinuityModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToClassificationInfo(modelMap["site_continuity"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.SiteContinuity = SiteContinuityModel
	}
	if modelMap["threat_protection"] != nil && len(modelMap["threat_protection"].([]interface{})) > 0 {
		ThreatProtectionModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToClassificationInfo(modelMap["threat_protection"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ThreatProtection = ThreatProtectionModel
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToClassificationInfo(modelMap map[string]interface{}) (*backuprecoveryv1.ClassificationInfo, error) {
	model := &backuprecoveryv1.ClassificationInfo{}
	if modelMap["end_date"] != nil && modelMap["end_date"].(string) != "" {
		model.EndDate = core.StringPtr(modelMap["end_date"].(string))
	}
	if modelMap["is_active"] != nil {
		model.IsActive = core.BoolPtr(modelMap["is_active"].(bool))
	}
	if modelMap["is_free_trial"] != nil {
		model.IsFreeTrial = core.BoolPtr(modelMap["is_free_trial"].(bool))
	}
	if modelMap["start_date"] != nil && modelMap["start_date"].(string) != "" {
		model.StartDate = core.StringPtr(modelMap["start_date"].(string))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToDataProtectInfo(modelMap map[string]interface{}) (*backuprecoveryv1.DataProtectInfo, error) {
	model := &backuprecoveryv1.DataProtectInfo{}
	if modelMap["end_date"] != nil && modelMap["end_date"].(string) != "" {
		model.EndDate = core.StringPtr(modelMap["end_date"].(string))
	}
	if modelMap["is_active"] != nil {
		model.IsActive = core.BoolPtr(modelMap["is_active"].(bool))
	}
	if modelMap["is_free_trial"] != nil {
		model.IsFreeTrial = core.BoolPtr(modelMap["is_free_trial"].(bool))
	}
	if modelMap["is_aws_subscription"] != nil {
		model.IsAwsSubscription = core.BoolPtr(modelMap["is_aws_subscription"].(bool))
	}
	if modelMap["is_cohesity_subscription"] != nil {
		model.IsCohesitySubscription = core.BoolPtr(modelMap["is_cohesity_subscription"].(bool))
	}
	if modelMap["quantity"] != nil {
		model.Quantity = core.Int64Ptr(int64(modelMap["quantity"].(int)))
	}
	if modelMap["start_date"] != nil && modelMap["start_date"].(string) != "" {
		model.StartDate = core.StringPtr(modelMap["start_date"].(string))
	}
	if modelMap["tiering"] != nil && len(modelMap["tiering"].([]interface{})) > 0 {
		TieringModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToTieringInfo(modelMap["tiering"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Tiering = TieringModel
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToTieringInfo(modelMap map[string]interface{}) (*backuprecoveryv1.TieringInfo, error) {
	model := &backuprecoveryv1.TieringInfo{}
	if modelMap["backend_tiering"] != nil {
		model.BackendTiering = core.BoolPtr(modelMap["backend_tiering"].(bool))
	}
	if modelMap["frontend_tiering"] != nil {
		model.FrontendTiering = core.BoolPtr(modelMap["frontend_tiering"].(bool))
	}
	if modelMap["max_retention"] != nil {
		model.MaxRetention = core.Int64Ptr(int64(modelMap["max_retention"].(int)))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToDataProtectAzureInfo(modelMap map[string]interface{}) (*backuprecoveryv1.DataProtectAzureInfo, error) {
	model := &backuprecoveryv1.DataProtectAzureInfo{}
	if modelMap["end_date"] != nil && modelMap["end_date"].(string) != "" {
		model.EndDate = core.StringPtr(modelMap["end_date"].(string))
	}
	if modelMap["is_active"] != nil {
		model.IsActive = core.BoolPtr(modelMap["is_active"].(bool))
	}
	if modelMap["is_free_trial"] != nil {
		model.IsFreeTrial = core.BoolPtr(modelMap["is_free_trial"].(bool))
	}
	if modelMap["quantity"] != nil {
		model.Quantity = core.Int64Ptr(int64(modelMap["quantity"].(int)))
	}
	if modelMap["start_date"] != nil && modelMap["start_date"].(string) != "" {
		model.StartDate = core.StringPtr(modelMap["start_date"].(string))
	}
	if modelMap["tiering"] != nil && len(modelMap["tiering"].([]interface{})) > 0 {
		TieringModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToTieringInfo(modelMap["tiering"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Tiering = TieringModel
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToFortKnoxInfo(modelMap map[string]interface{}) (*backuprecoveryv1.FortKnoxInfo, error) {
	model := &backuprecoveryv1.FortKnoxInfo{}
	if modelMap["end_date"] != nil && modelMap["end_date"].(string) != "" {
		model.EndDate = core.StringPtr(modelMap["end_date"].(string))
	}
	if modelMap["is_active"] != nil {
		model.IsActive = core.BoolPtr(modelMap["is_active"].(bool))
	}
	if modelMap["is_free_trial"] != nil {
		model.IsFreeTrial = core.BoolPtr(modelMap["is_free_trial"].(bool))
	}
	if modelMap["quantity"] != nil {
		model.Quantity = core.Int64Ptr(int64(modelMap["quantity"].(int)))
	}
	if modelMap["start_date"] != nil && modelMap["start_date"].(string) != "" {
		model.StartDate = core.StringPtr(modelMap["start_date"].(string))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserMapToTenantAccesses(modelMap map[string]interface{}) (*backuprecoveryv1.TenantAccesses, error) {
	model := &backuprecoveryv1.TenantAccesses{}
	if modelMap["cluster_identifiers"] != nil {
		clusterIdentifiers := []backuprecoveryv1.UserClusterIdentifier{}
		for _, clusterIdentifiersItem := range modelMap["cluster_identifiers"].([]interface{}) {
			clusterIdentifiersItemModel, err := ResourceIbmBackupRecoveryConnectorUpdateUserMapToUserClusterIdentifier(clusterIdentifiersItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			clusterIdentifiers = append(clusterIdentifiers, *clusterIdentifiersItemModel)
		}
		model.ClusterIdentifiers = clusterIdentifiers
	}
	if modelMap["created_time_msecs"] != nil {
		model.CreatedTimeMsecs = core.Int64Ptr(int64(modelMap["created_time_msecs"].(int)))
	}
	if modelMap["effective_time_msecs"] != nil {
		model.EffectiveTimeMsecs = core.Int64Ptr(int64(modelMap["effective_time_msecs"].(int)))
	}
	if modelMap["expired_time_msecs"] != nil {
		model.ExpiredTimeMsecs = core.Int64Ptr(int64(modelMap["expired_time_msecs"].(int)))
	}
	if modelMap["is_access_active"] != nil {
		model.IsAccessActive = core.BoolPtr(modelMap["is_access_active"].(bool))
	}
	if modelMap["is_active"] != nil {
		model.IsActive = core.BoolPtr(modelMap["is_active"].(bool))
	}
	if modelMap["is_deleted"] != nil {
		model.IsDeleted = core.BoolPtr(modelMap["is_deleted"].(bool))
	}
	if modelMap["last_updated_time_msecs"] != nil {
		model.LastUpdatedTimeMsecs = core.Int64Ptr(int64(modelMap["last_updated_time_msecs"].(int)))
	}
	if modelMap["roles"] != nil {
		roles := []string{}
		for _, rolesItem := range modelMap["roles"].([]interface{}) {
			roles = append(roles, rolesItem.(string))
		}
		model.Roles = roles
	}
	if modelMap["tenant_id"] != nil && modelMap["tenant_id"].(string) != "" {
		model.TenantID = core.StringPtr(modelMap["tenant_id"].(string))
	}
	if modelMap["tenant_name"] != nil && modelMap["tenant_name"].(string) != "" {
		model.TenantName = core.StringPtr(modelMap["tenant_name"].(string))
	}
	if modelMap["tenant_type"] != nil && modelMap["tenant_type"].(string) != "" {
		model.TenantType = core.StringPtr(modelMap["tenant_type"].(string))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserAdUserInfoToMap(model *backuprecoveryv1.AdUserInfo) (map[string]interface{}, error) {
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

func ResourceIbmBackupRecoveryConnectorUpdateUserAuditLogSettingsToMap(model *backuprecoveryv1.AuditLogSettings) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ReadLogging != nil {
		modelMap["read_logging"] = *model.ReadLogging
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserUserClusterIdentifierToMap(model *backuprecoveryv1.UserClusterIdentifier) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClusterID != nil {
		modelMap["cluster_id"] = flex.IntValue(model.ClusterID)
	}
	if model.ClusterIncarnationID != nil {
		modelMap["cluster_incarnation_id"] = flex.IntValue(model.ClusterIncarnationID)
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserGoogleAccountInfoToMap(model *backuprecoveryv1.GoogleAccountInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AccountID != nil {
		modelMap["account_id"] = *model.AccountID
	}
	if model.UserID != nil {
		modelMap["user_id"] = *model.UserID
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserIdpUserInfoToMap(model *backuprecoveryv1.IdpUserInfo) (map[string]interface{}, error) {
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

func ResourceIbmBackupRecoveryConnectorUpdateUserMfaInfoToMap(model *backuprecoveryv1.MfaInfo) (map[string]interface{}, error) {
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

func ResourceIbmBackupRecoveryConnectorUpdateUserTenantConfigToMap(model *backuprecoveryv1.TenantConfig) (map[string]interface{}, error) {
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

func ResourceIbmBackupRecoveryConnectorUpdateUserUsersPreferencesToMap(model *backuprecoveryv1.UsersPreferences) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Locale != nil {
		modelMap["locale"] = *model.Locale
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserUserProfileToMap(model *backuprecoveryv1.UserProfile) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClusterIdentifiers != nil {
		clusterIdentifiers := []map[string]interface{}{}
		for _, clusterIdentifiersItem := range model.ClusterIdentifiers {
			clusterIdentifiersItemMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserUserClusterIdentifierToMap(&clusterIdentifiersItem) // #nosec G601
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

func ResourceIbmBackupRecoveryConnectorUpdateUserSalesforceAccountInfoToMap(model *backuprecoveryv1.SalesforceAccountInfo) (map[string]interface{}, error) {
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

func ResourceIbmBackupRecoveryConnectorUpdateUserSpogContextToMap(model *backuprecoveryv1.SpogContext) (map[string]interface{}, error) {
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

func ResourceIbmBackupRecoveryConnectorUpdateUserSubscriptionInfoToMap(model *backuprecoveryv1.SubscriptionInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Classification != nil {
		classificationMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserClassificationInfoToMap(model.Classification)
		if err != nil {
			return modelMap, err
		}
		modelMap["classification"] = []map[string]interface{}{classificationMap}
	}
	if model.DataProtect != nil {
		dataProtectMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserDataProtectInfoToMap(model.DataProtect)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_protect"] = []map[string]interface{}{dataProtectMap}
	}
	if model.DataProtectAzure != nil {
		dataProtectAzureMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserDataProtectAzureInfoToMap(model.DataProtectAzure)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_protect_azure"] = []map[string]interface{}{dataProtectAzureMap}
	}
	if model.FortKnoxAzureCool != nil {
		fortKnoxAzureCoolMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserFortKnoxInfoToMap(model.FortKnoxAzureCool)
		if err != nil {
			return modelMap, err
		}
		modelMap["fort_knox_azure_cool"] = []map[string]interface{}{fortKnoxAzureCoolMap}
	}
	if model.FortKnoxAzureHot != nil {
		fortKnoxAzureHotMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserFortKnoxInfoToMap(model.FortKnoxAzureHot)
		if err != nil {
			return modelMap, err
		}
		modelMap["fort_knox_azure_hot"] = []map[string]interface{}{fortKnoxAzureHotMap}
	}
	if model.FortKnoxCold != nil {
		fortKnoxColdMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserFortKnoxInfoToMap(model.FortKnoxCold)
		if err != nil {
			return modelMap, err
		}
		modelMap["fort_knox_cold"] = []map[string]interface{}{fortKnoxColdMap}
	}
	if model.Ransomware != nil {
		ransomwareMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserFortKnoxInfoToMap(model.Ransomware)
		if err != nil {
			return modelMap, err
		}
		modelMap["ransomware"] = []map[string]interface{}{ransomwareMap}
	}
	if model.SiteContinuity != nil {
		siteContinuityMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserClassificationInfoToMap(model.SiteContinuity)
		if err != nil {
			return modelMap, err
		}
		modelMap["site_continuity"] = []map[string]interface{}{siteContinuityMap}
	}
	if model.ThreatProtection != nil {
		threatProtectionMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserClassificationInfoToMap(model.ThreatProtection)
		if err != nil {
			return modelMap, err
		}
		modelMap["threat_protection"] = []map[string]interface{}{threatProtectionMap}
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserClassificationInfoToMap(model *backuprecoveryv1.ClassificationInfo) (map[string]interface{}, error) {
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

func ResourceIbmBackupRecoveryConnectorUpdateUserDataProtectInfoToMap(model *backuprecoveryv1.DataProtectInfo) (map[string]interface{}, error) {
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
		tieringMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserTieringInfoToMap(model.Tiering)
		if err != nil {
			return modelMap, err
		}
		modelMap["tiering"] = []map[string]interface{}{tieringMap}
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserTieringInfoToMap(model *backuprecoveryv1.TieringInfo) (map[string]interface{}, error) {
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

func ResourceIbmBackupRecoveryConnectorUpdateUserDataProtectAzureInfoToMap(model *backuprecoveryv1.DataProtectAzureInfo) (map[string]interface{}, error) {
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
		tieringMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserTieringInfoToMap(model.Tiering)
		if err != nil {
			return modelMap, err
		}
		modelMap["tiering"] = []map[string]interface{}{tieringMap}
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoveryConnectorUpdateUserFortKnoxInfoToMap(model *backuprecoveryv1.FortKnoxInfo) (map[string]interface{}, error) {
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

func ResourceIbmBackupRecoveryConnectorUpdateUserTenantAccessesToMap(model *backuprecoveryv1.TenantAccesses) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClusterIdentifiers != nil {
		clusterIdentifiers := []map[string]interface{}{}
		for _, clusterIdentifiersItem := range model.ClusterIdentifiers {
			clusterIdentifiersItemMap, err := ResourceIbmBackupRecoveryConnectorUpdateUserUserClusterIdentifierToMap(&clusterIdentifiersItem) // #nosec G601
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

func resourceIbmBackupRecoveryConnectorUpdateUserUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Resource update will only affect terraform state and not the actual backend resource",
		Detail:   "Update operation for this resource is not supported and will only affect the terraform statefile. No changes will be made to the backend resource.",
	}
	diags = append(diags, warning)
	// d.SetId("")
	return diags
}
