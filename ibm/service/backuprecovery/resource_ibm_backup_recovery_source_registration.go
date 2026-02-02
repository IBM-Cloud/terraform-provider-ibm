// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmBackupRecoverySourceRegistration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmBackupRecoverySourceRegistrationCreate,
		ReadContext:   resourceIbmBackupRecoverySourceRegistrationRead,
		UpdateContext: resourceIbmBackupRecoverySourceRegistrationUpdate,
		DeleteContext: resourceIbmBackupRecoverySourceRegistrationDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"x_ibm_tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.",
			},
			"environment": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// ValidateFunc: validate.InvokeValidator("ibm_backup_recovery_source_registration", "environment"),
				Description: "Specifies the environment type of the Protection Source.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The user specified name for this source.",
			},
			"connection_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. This field will be depricated in future. Use connections field.",
			},
			"is_internal_encrypted": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies if credentials are encrypted by internal key.",
			},
			"encryption_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the key that user has encrypted the credential with.",
			},
			"connections": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Specfies the list of connections for the source.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the id of the connection.",
						},
						"entity_id": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the entity id of the source. The source can a non-root entity.",
						},
						"connector_group_id": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the connector group id of connector groups.",
						},
						"data_source_connection_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the id of the connection in string format.",
						},
					},
				},
			},
			"connector_group_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the connector group id of connector groups.",
			},
			"data_source_connection_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. Also, this is the 'string' of connectionId. This property was added to accommodate for ID values that exceed 2^53 - 1, which is the max value for which JS maintains precision.",
			},
			"advanced_configs": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the advanced configuration for a protection source.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "key.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "value.",
						},
					},
				},
			},
			"physical_params": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Specifies parameters to register physical server.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"endpoint": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specifies the endpoint IPaddress, URL or hostname of the physical host.",
						},
						"force_register": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The agent running on a physical host will fail the registration if it is already registered as part of another cluster. By setting this option to true, agent can be forced to register with the current cluster.",
						},
						"host_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the type of host.",
						},
						"physical_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the type of physical server.",
						},
						"applications": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the list of applications to be registered with Physical Source.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"kubernetes_params": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Specifies the parameters to register a Kubernetes source.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_protect_config": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the parameters to auto protect the source after registration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error_message": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the error message in case source registration is successful but protection job creation fails.",
									},
									"is_default_auto_protected": &schema.Schema{
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Specifies if entire source should be auto protected after registration. Default: False.",
									},
									"policy_id": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specifies the protection policy to auto protect the source with.",
									},
									"protection_group_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the protection group Id after it is successfully created.",
									},
									"storage_domain_id": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the storage domain id for the protection job.",
									},
								},
							},
						},
						"client_private_key": &schema.Schema{
							Type:             schema.TypeString,
							Required:         true,
							Sensitive:        true,
							DiffSuppressFunc: suppressParameterDuringRefresh,
							Description:      "Specifies the bearer token or private key of Kubernetes source.",
						},
						"cohesity_dataprotect_plugin_image_location": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the custom Cohesity Dataprotect plugin image location of the Kubernetes source.",
						},
						"data_mover_image_location": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specifies the datamover image location of Kubernetes source.",
						},
						"datamover_service_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the data mover service type of Kubernetes source.",
						},
						"default_vlan_params": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies VLAN params associated with the backup/restore operation.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"disable_vlan": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the restore.",
									},
									"interface_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Interface group to use for backup/restore. If this is not specified, primary interface group for the cluster will be used.",
									},
									"vlan_id": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "If this is set, then the Cohesity host name or the IP address associated with this VLAN is used for mounting Cohesity's view on the remote host.",
									},
								},
							},
						},
						"endpoint": &schema.Schema{
							Type:             schema.TypeString,
							Required:         true,
							Sensitive:        true,
							DiffSuppressFunc: suppressParameterDuringRefresh,
							Description:      "Specifies the endpoint of Kubernetes source.",
						},
						"init_container_image_location": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the initial container image location of Kubernetes source.",
						},
						"kubernetes_distribution": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specifies the distribution type of Kubernetes source.",
						},
						"kubernetes_type": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: suppressParameterDuringRefresh,
							Description:      "Specifies the type of kubernetes source.",
						},
						"priority_class_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the priority class name for cohesity resources.",
						},
						"resource_annotations": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies resource annotations to be applied on cohesity resources.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specifies the label key.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the label value.",
									},
								},
							},
						},
						"resource_labels": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies resource label to be applied on cohesity resources.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specifies the label key.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the label value.",
									},
								},
							},
						},
						"san_fields": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the SAN field for agent certificate.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"service_annotations": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the service annotation object of Kubernetes source.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the service annotation key value.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the service annotation value.",
									},
								},
							},
						},
						"velero_aws_plugin_image_location": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the velero AWS plugin image location of the Kubernetes source.",
						},
						"velero_image_location": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the velero image location of the Kubernetes source.",
						},
						"velero_openshift_plugin_image_location": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the velero open shift plugin image for the Kubernetes source.",
						},
						"vlan_info_vec": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies VLAN information provided during registration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_annotations": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies annotations to be put on services for IP allocation. Applicable only when service is of type LoadBalancer.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the service annotation key value.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the service annotation value.",
												},
											},
										},
									},
									"vlan_params": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies VLAN params associated with the backup/restore operation.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"disable_vlan": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the restore.",
												},
												"interface_name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Interface group to use for backup/restore. If this is not specified, primary interface group for the cluster will be used.",
												},
												"vlan_id": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "If this is set, then the Cohesity host name or the IP address associated with this VLAN is used for mounting Cohesity's view on the remote host.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"auto_proetction_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user specified name for this source.",
			},
			"source_id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of top level source object discovered after the registration.",
			},
			"source_info": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies information about an object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Specifies object id.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the name of the object.",
						},
						"source_id": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Specifies registered source id to which object belongs.",
						},
						"source_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Specifies registered source name to which object belongs.",
						},
						"environment": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the environment of the object.",
						},
						"object_hash": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the hash identifier of the object.",
						},
						"object_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the type of the object.",
						},
						"logical_size_bytes": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the logical size of object in bytes.",
						},
						"uuid": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the uuid which is a unique identifier of the object.",
						},
						"global_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the global id which is a unique identifier of the object.",
						},
						"protection_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the protection type of the object if any.",
						},
						"sharepoint_site_summary": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the common parameters for Sharepoint site objects.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"site_web_url": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the web url for the Sharepoint site.",
									},
								},
							},
						},
						"os_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the operating system type of the object.",
						},
						"child_objects": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "Specifies child object details.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Specifies object id.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the name of the object.",
									},
									"source_id": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Specifies registered source id to which object belongs.",
									},
									"source_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Specifies registered source name to which object belongs.",
									},
									"environment": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the environment of the object.",
									},
									"object_hash": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the hash identifier of the object.",
									},
									"object_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the type of the object.",
									},
									"logical_size_bytes": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the logical size of object in bytes.",
									},
									"uuid": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the uuid which is a unique identifier of the object.",
									},
									"global_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the global id which is a unique identifier of the object.",
									},
									"protection_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the protection type of the object if any.",
									},
									"sharepoint_site_summary": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the common parameters for Sharepoint site objects.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"site_web_url": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Specifies the web url for the Sharepoint site.",
												},
											},
										},
									},
									"os_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the operating system type of the object.",
									},
									"child_objects": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "Specifies child object details.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{},
										},
									},
									"v_center_summary": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"is_cloud_env": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Computed:    true,
													Description: "Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.",
												},
											},
										},
									},
									"windows_cluster_summary": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cluster_source_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Specifies the type of cluster resource this source represents.",
												},
											},
										},
									},
								},
							},
						},
						"v_center_summary": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_cloud_env": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Computed:    true,
										Description: "Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.",
									},
								},
							},
						},
						"windows_cluster_summary": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster_source_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the type of cluster resource this source represents.",
									},
								},
							},
						},
						"protection_stats": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the count and size of protected and unprotected objects for the size.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"environment": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the environment of the object.",
									},
									"protected_count": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the count of the protected leaf objects.",
									},
									"unprotected_count": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the count of the unprotected leaf objects.",
									},
									"deleted_protected_count": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the count of protected leaf objects which were deleted from the source after being protected.",
									},
									"protected_size_bytes": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the protected logical size in bytes.",
									},
									"unprotected_size_bytes": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the unprotected logical size in bytes.",
									},
								},
							},
						},
						"permissions": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the list of users, groups and users that have permissions for a given object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"object_id": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the id of the object.",
									},
									"users": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the list of users which has the permissions to the object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Specifies the name of the user.",
												},
												"sid": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Specifies the sid of the user.",
												},
												"domain": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Specifies the domain of the user.",
												},
											},
										},
									},
									"groups": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the list of user groups which has permissions to the object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Specifies the name of the user group.",
												},
												"sid": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Specifies the sid of the user group.",
												},
												"domain": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Specifies the domain of the user group.",
												},
											},
										},
									},
									"tenant": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "Specifies a tenant object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"created_at_time_msecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Epoch time when tenant was created.",
												},
												"deleted_at_time_msecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Epoch time when tenant was last updated.",
												},
												"description": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Description about the tenant.",
												},
												"external_vendor_metadata": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Description: "Specifies the additional metadata for the tenant that is specifically set by the external vendors who are responsible for managing tenants. This field will only applicable if tenant creation is happening for a specially provisioned clusters for external vendors.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ibm_tenant_metadata_params": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Computed:    true,
																Description: "Specifies the additional metadata for the tenant that is specifically set by the external vendor of type 'IBM'.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"account_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																			Description: "Specifies the unique identifier of the IBM's account ID.",
																		},
																		"crn": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																			Description: "Specifies the unique CRN associated with the tenant.",
																		},
																		"custom_properties": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Computed:    true,
																			Description: "Specifies the list of custom properties associated with the tenant. External vendors can choose to set any properties inside following list. Note that the fields set inside the following will not be available for direct filtering. API callers should make sure that no sensitive information such as passwords is sent in these fields.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"key": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Computed:    true,
																						Description: "Specifies the unique key for custom property.",
																					},
																					"value": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Computed:    true,
																						Description: "Specifies the value for the above custom key.",
																					},
																				},
																			},
																		},
																		"liveness_mode": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																			Description: "Specifies the current liveness mode of the tenant. This mode may change based on AZ failures when vendor chooses to failover or failback the tenants to other AZs.",
																		},
																		"metrics_config": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Computed:    true,
																			Description: "Specifies the metadata for metrics configuration. The metadata defined here will be used by cluster to send the usgae metrics to IBM cloud metering service for calculating the tenant billing.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"cos_resource_config": &schema.Schema{
																						Type:        schema.TypeList,
																						Optional:    true,
																						Computed:    true,
																						Description: "Specifies the details of COS resource configuration required for posting metrics and trackinb billing information for IBM tenants.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"resource_url": &schema.Schema{
																									Type:        schema.TypeString,
																									Optional:    true,
																									Computed:    true,
																									Description: "Specifies the resource COS resource configuration endpoint that will be used for fetching bucket usage for a given tenant.",
																								},
																							},
																						},
																					},
																					"iam_metrics_config": &schema.Schema{
																						Type:        schema.TypeList,
																						Optional:    true,
																						Computed:    true,
																						Description: "Specifies the IAM configuration that will be used for accessing the billing service in IBM cloud.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"iam_url": &schema.Schema{
																									Type:        schema.TypeString,
																									Optional:    true,
																									Computed:    true,
																									Description: "Specifies the IAM URL needed to fetch the operator token from IBM. The operator token is needed to make service API calls to IBM billing service.",
																								},
																								"billing_api_key_secret_id": &schema.Schema{
																									Type:        schema.TypeString,
																									Optional:    true,
																									Computed:    true,
																									Description: "Specifies Id of the secret that contains the API key.",
																								},
																							},
																						},
																					},
																					"metering_config": &schema.Schema{
																						Type:        schema.TypeList,
																						Optional:    true,
																						Computed:    true,
																						Description: "Specifies the metering configuration that will be used for IBM cluster to send the billing details to IBM billing service.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"part_ids": &schema.Schema{
																									Type:        schema.TypeList,
																									Optional:    true,
																									Computed:    true,
																									Description: "Specifies the list of part identifiers used for metrics identification.",
																									Elem:        &schema.Schema{Type: schema.TypeString},
																								},
																								"submission_interval_in_secs": &schema.Schema{
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Computed:    true,
																									Description: "Specifies the frequency in seconds at which the metrics will be pushed to IBM billing service from cluster.",
																								},
																								"url": &schema.Schema{
																									Type:        schema.TypeString,
																									Optional:    true,
																									Computed:    true,
																									Description: "Specifies the base metering URL that will be used by cluster to send the billing information.",
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																		"ownership_mode": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																			Description: "Specifies the current ownership mode for the tenant. The ownership of the tenant represents the active role for functioning of the tenant.",
																		},
																		"plan_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																			Description: "Specifies the Plan Id associated with the tenant. This field is introduced for tracking purposes inside IBM enviournment.",
																		},
																		"resource_group_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																			Description: "Specifies the Resource Group ID associated with the tenant.",
																		},
																		"resource_instance_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																			Description: "Specifies the Resource Instance ID associated with the tenant. This field is introduced for tracking purposes inside IBM enviournment.",
																		},
																	},
																},
															},
															"type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the type of the external vendor. The type specific parameters must be specified the provided type.",
															},
														},
													},
												},
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "The tenant id.",
												},
												"is_managed_on_helios": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Computed:    true,
													Description: "Flag to indicate if tenant is managed on helios.",
												},
												"last_updated_at_time_msecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Epoch time when tenant was last updated.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Name of the Tenant.",
												},
												"network": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Description: "Networking information about a Tenant on a Cluster.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"connector_enabled": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Whether connector (hybrid extender) is enabled.",
															},
															"cluster_hostname": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "The hostname for Cohesity cluster as seen by tenants and as is routable from the tenant's network. Tenant's VLAN's hostname, if available can be used instead but it is mandatory to provide this value if there's no VLAN hostname to use. Also, when set, this field would take precedence over VLAN hostname.",
															},
															"cluster_ips": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Computed:    true,
																Description: "Set of IPs as seen from the tenant's network for the Cohesity cluster. Only one from 'clusterHostname' and 'clusterIps' is needed.",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
														},
													},
												},
												"status": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Current Status of the Tenant.",
												},
											},
										},
									},
								},
							},
						},
						"mssql_params": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the parameters for Msssql object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aag_info": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "Object details for Mssql.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Specifies the AAG name.",
												},
												"object_id": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Computed:    true,
													Description: "Specifies the AAG object Id.",
												},
											},
										},
									},
									"host_info": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "Specifies the host information for a objects. This is mainly populated in case of App objects where app object is hosted by another object such as VM or physical server.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Specifies the id of the host object.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Specifies the name of the host object.",
												},
												"environment": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Specifies the environment of the object.",
												},
											},
										},
									},
									"is_encrypted": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Computed:    true,
										Description: "Specifies whether the database is TDE enabled.",
									},
								},
							},
						},
						"physical_params": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the parameters for Physical object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_system_backup": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Computed:    true,
										Description: "Specifies if system backup was enabled for the source in a particular run.",
									},
								},
							},
						},
					},
				},
			},
			"authentication_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the status of the authentication during the registration of a Protection Source. 'Pending' indicates the authentication is in progress. 'Scheduled' indicates the authentication is scheduled. 'Finished' indicates the authentication is completed. 'RefreshInProgress' indicates the refresh is in progress.",
			},
			"registration_time_msecs": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Specifies the time when the source was registered in milliseconds.",
			},
			"last_refreshed_time_msecs": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Specifies the time when the source was last refreshed in milliseconds.",
			},
			"external_metadata": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the External metadata of an entity.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"maintenance_mode_config": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the entity metadata for maintenance mode.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"activation_time_intervals": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the absolute intervals where the maintenance schedule is valid, i.e. maintenance_shedule is considered only for these time ranges. (For example, if there is one time range with [now_usecs, now_usecs + 10 days], the action will be done during the maintenance_schedule for the next 10 days.)The start time must be specified. The end time can be -1 which would denote an indefinite maintenance mode.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"end_time_usecs": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,

													Description: "Specifies the end time of this time range.",
												},
												"start_time_usecs": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,

													Description: "Specifies the start time of this time range.",
												},
											},
										},
									},
									"maintenance_schedule": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies a schedule for actions to be taken.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"periodic_time_windows": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the time range within the days of the week.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"day_of_the_week": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the week day.",
															},
															"end_time": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the time in hours and minutes.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"hour": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,

																			Description: "Specifies the hour of this time.",
																		},
																		"minute": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,

																			Description: "Specifies the minute of this time.",
																		},
																	},
																},
															},
															"start_time": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,

																Description: "Specifies the time in hours and minutes.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"hour": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,

																			Description: "Specifies the hour of this time.",
																		},
																		"minute": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,

																			Description: "Specifies the minute of this time.",
																		},
																	},
																},
															},
														},
													},
												},
												"schedule_type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,

													Description: "Specifies the type of schedule for this ScheduleProto.",
												},
												"time_ranges": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,

													Description: "Specifies the time ranges in usecs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"end_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the end time of this time range.",
															},
															"start_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the start time of this time range.",
															},
														},
													},
												},
												"timezone": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,

													Description: "Specifies the timezone of the user of this ScheduleProto. The timezones have unique names of the form 'Area/Location'.",
												},
											},
										},
									},
									"user_message": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,

										Description: "User provided message associated with this maintenance mode.",
									},
									"workflow_intervention_spec_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,

										Description: "Specifies the type of intervention for different workflows when the source goes into maintenance mode.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"intervention": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,

													Description: "Specifies the intervention type for ongoing tasks.",
												},
												"workflow_type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,

													Description: "Specifies the workflow type for which an intervention would be needed when maintenance mode begins.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func suppressParameterDuringRefresh(k, o, n string, d *schema.ResourceData) bool {
	if len(d.Id()) == 0 {
		return false
	}
	if len(d.Id()) != 0 {
		if d.HasChange("kubernetes_params.0.data_mover_image_location") ||
			d.HasChange("kubernetes_params.0.datamover_service_type") ||
			d.HasChange("kubernetes_params.0.vlan_info_vec") ||
			d.HasChange("kubernetes_params.0.resource_annotations") ||
			d.HasChange("kubernetes_params.0.resource_labels") ||
			d.HasChange("kubernetes_params.0.velero_openshift_plugin_image_location") ||
			d.HasChange("kubernetes_params.0.cohesity_dataprotect_plugin_image_location") ||
			d.HasChange("kubernetes_params.0.velero_image_location") ||
			d.HasChange("kubernetes_params.0.velero_aws_plugin_image_location") ||
			d.HasChange("kubernetes_params.0.san_fields") ||
			d.HasChange("kubernetes_params.0.service_annotations") ||
			d.HasChange("kubernetes_params.0.priority_class_name") ||
			d.HasChange("kubernetes_params.0.kubernetes_distribution") ||
			d.HasChange("kubernetes_params.0.init_container_image_location") ||
			d.HasChange("kubernetes_params.0.auto_protect_config") ||
			d.HasChange("kubernetes_params.0.default_vlan_params") {
			return false
		}
	}
	return true
}

func resourceIbmBackupRecoverySourceRegistrationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	endpointType := d.Get("endpoint_type").(string)
	instanceId, region := getInstanceIdAndRegion(d)
	if instanceId != "" && region != "" {
		bmxsession, err := meta.(conns.ClientSession).BluemixSession()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to get clientSession"), "ibm_backup_recovery", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		backupRecoveryClient = getClientWithInstanceEndpoint(backupRecoveryClient, bmxsession, instanceId, region, endpointType)
	}

	registerProtectionSourceOptions := &backuprecoveryv1.RegisterProtectionSourceOptions{}

	tenantId := d.Get("x_ibm_tenant_id").(string)
	registerProtectionSourceOptions.SetXIBMTenantID(tenantId)
	registerProtectionSourceOptions.SetEnvironment(d.Get("environment").(string))
	if _, ok := d.GetOk("name"); ok {
		registerProtectionSourceOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("is_internal_encrypted"); ok {
		registerProtectionSourceOptions.SetIsInternalEncrypted(d.Get("is_internal_encrypted").(bool))
	}
	if _, ok := d.GetOk("encryption_key"); ok {
		registerProtectionSourceOptions.SetEncryptionKey(d.Get("encryption_key").(string))
	}
	if _, ok := d.GetOk("connection_id"); ok {
		connId, err := strconv.ParseInt(d.Get("connection_id").(string), 10, 64)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "create", "parse-connection-id").GetDiag()
		}
		registerProtectionSourceOptions.SetConnectionID(connId)
	}
	if _, ok := d.GetOk("connections"); ok {
		var connections []backuprecoveryv1.ConnectionConfig
		for _, v := range d.Get("connections").([]interface{}) {
			value := v.(map[string]interface{})
			connectionsItem, err := ResourceIbmBackupRecoverySourceRegistrationMapToConnectionConfig(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "create", "parse-connections").GetDiag()
			}
			connections = append(connections, *connectionsItem)
		}
		registerProtectionSourceOptions.SetConnections(connections)
	}
	if _, ok := d.GetOk("connector_group_id"); ok {
		registerProtectionSourceOptions.SetConnectorGroupID(int64(d.Get("connector_group_id").(int)))
	}
	if _, ok := d.GetOk("advanced_configs"); ok {
		var advancedConfigs []backuprecoveryv1.KeyValuePair
		for _, v := range d.Get("advanced_configs").([]interface{}) {
			value := v.(map[string]interface{})
			advancedConfigsItem, err := ResourceIbmBackupRecoverySourceRegistrationMapToKeyValuePair(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "create", "parse-advanced_configs").GetDiag()
			}
			advancedConfigs = append(advancedConfigs, *advancedConfigsItem)
		}
		registerProtectionSourceOptions.SetAdvancedConfigs(advancedConfigs)
	}
	if _, ok := d.GetOk("data_source_connection_id"); ok {
		registerProtectionSourceOptions.SetDataSourceConnectionID(d.Get("data_source_connection_id").(string))
	}
	if _, ok := d.GetOk("kubernetes_params"); ok {
		kubernetesParamsModel, err := ResourceIbmBackupRecoverySourceRegistrationMapToKubernetesSourceRegistrationParams(d.Get("kubernetes_params.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "create", "parse-kubernetes_params").GetDiag()
		}
		registerProtectionSourceOptions.SetKubernetesParams(kubernetesParamsModel)
	}
	if _, ok := d.GetOk("physical_params"); ok {
		physicalParamsModel, err := ResourceIbmBackupRecoverySourceRegistrationMapToPhysicalSourceRegistrationParams(d.Get("physical_params.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "create", "parse-physical_params").GetDiag()
		}
		registerProtectionSourceOptions.SetPhysicalParams(physicalParamsModel)
	}

	sourceRegistrationReponseParams, _, err := backupRecoveryClient.RegisterProtectionSourceWithContext(context, registerProtectionSourceOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("RegisterProtectionSourceWithContext failed: %s", err.Error()), "ibm_backup_recovery_source_registration", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	registrationId := fmt.Sprintf("%s::%s", tenantId, strconv.Itoa(int(*sourceRegistrationReponseParams.ID)))
	d.SetId(registrationId)

	if sourceRegistrationReponseParams.KubernetesParams != nil && sourceRegistrationReponseParams.KubernetesParams.AutoProtectConfig != nil && sourceRegistrationReponseParams.KubernetesParams.AutoProtectConfig.ProtectionGroupID != nil {

		err := d.Set("auto_proetction_group_id", *sourceRegistrationReponseParams.KubernetesParams.AutoProtectConfig.ProtectionGroupID)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("RegisterProtectionSourceWithContext failed: %s", err.Error()), "ibm_backup_recovery_source_registration", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return resourceIbmBackupRecoverySourceRegistrationRead(context, d, meta)
}

func resourceIbmBackupRecoverySourceRegistrationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tenantId := d.Get("x_ibm_tenant_id").(string)
	registrationId := d.Id()
	if strings.Contains(d.Id(), "::") {
		tenantId = ParseId(d.Id(), "tenantId")
		registrationId = ParseId(d.Id(), "id")
	}
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	endpointType := d.Get("endpoint_type").(string)
	instanceId, region := getInstanceIdAndRegion(d)
	if instanceId != "" && region != "" {
		bmxsession, err := meta.(conns.ClientSession).BluemixSession()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to get clientSession"), "ibm_backup_recovery", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		backupRecoveryClient = getClientWithInstanceEndpoint(backupRecoveryClient, bmxsession, instanceId, region, endpointType)
	}
	getProtectionSourceRegistrationOptions := &backuprecoveryv1.GetProtectionSourceRegistrationOptions{}

	id, err := strconv.Atoi(registrationId)
	if err != nil {
		return diag.FromErr(err)
	}

	getProtectionSourceRegistrationOptions.SetID(int64(id))
	getProtectionSourceRegistrationOptions.SetXIBMTenantID(tenantId)

	sourceRegistrationReponseParams, response, err := backupRecoveryClient.GetProtectionSourceRegistrationWithContext(context, getProtectionSourceRegistrationOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetProtectionSourceRegistrationWithContext failed: %s", err.Error()), "ibm_backup_recovery_source_registration", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if instanceId != "" {
		if err := d.Set("instance_id", instanceId); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting instance_id: %s", err), "(Resource) ibm_backup_recovery_source_registration", "read", "set-instance-id").GetDiag()
		}
	}
	if region != "" {
		if err := d.Set("region", region); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting region: %s", err), "(Resource) ibm_backup_recovery_source_registration", "read", "set--region").GetDiag()
		}
	}

	if err = d.Set("endpoint_type", d.Get("endpoint_type").(string)); err != nil {
		err = fmt.Errorf("Error setting endpoint_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-endpoint-type").GetDiag()
	}

	if err = d.Set("environment", sourceRegistrationReponseParams.Environment); err != nil {
		err = fmt.Errorf("Error setting environment: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-environment").GetDiag()
	}

	if err = d.Set("x_ibm_tenant_id", tenantId); err != nil {
		err = fmt.Errorf("Error setting x_ibm_tenant_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-x_ibm_tenant_id").GetDiag()
	}

	if !core.IsNil(sourceRegistrationReponseParams.Name) {
		if err = d.Set("name", sourceRegistrationReponseParams.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.ConnectionID) {
		if err = d.Set("connection_id", strconv.Itoa(int(*sourceRegistrationReponseParams.ConnectionID))); err != nil {
			err = fmt.Errorf("Error setting connection_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-connection_id").GetDiag()
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.Connections) {
		connections := []map[string]interface{}{}
		for _, connectionsItem := range sourceRegistrationReponseParams.Connections {
			connectionsItemMap, err := ResourceIbmBackupRecoverySourceRegistrationConnectionConfigToMap(&connectionsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "connections-to-map").GetDiag()
			}
			connections = append(connections, connectionsItemMap)
		}
		if err = d.Set("connections", connections); err != nil {
			err = fmt.Errorf("Error setting connections: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-connections").GetDiag()
		}
	} else {
		if err = d.Set("connections", []interface{}{}); err != nil {
			err = fmt.Errorf("Error setting external_metadata: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-external_metadata").GetDiag()
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.ConnectorGroupID) {
		if err = d.Set("connector_group_id", flex.IntValue(sourceRegistrationReponseParams.ConnectorGroupID)); err != nil {
			err = fmt.Errorf("Error setting connector_group_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-connector_group_id").GetDiag()
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.DataSourceConnectionID) {
		if err = d.Set("data_source_connection_id", sourceRegistrationReponseParams.DataSourceConnectionID); err != nil {
			err = fmt.Errorf("Error setting data_source_connection_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-data_source_connection_id").GetDiag()
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.AdvancedConfigs) {
		advancedConfigs := []map[string]interface{}{}
		for _, advancedConfigsItem := range sourceRegistrationReponseParams.AdvancedConfigs {
			advancedConfigsItemMap, err := ResourceIbmBackupRecoverySourceRegistrationKeyValuePairToMap(&advancedConfigsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "advanced_configs-to-map").GetDiag()
			}
			advancedConfigs = append(advancedConfigs, advancedConfigsItemMap)
		}
		if err = d.Set("advanced_configs", advancedConfigs); err != nil {
			err = fmt.Errorf("Error setting advanced_configs: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-advanced_configs").GetDiag()
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.PhysicalParams) {
		physicalParamsMap, err := ResourceIbmBackupRecoverySourceRegistrationPhysicalSourceRegistrationParamsToMap(sourceRegistrationReponseParams.PhysicalParams)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "physical_params-to-map").GetDiag()
		}
		if err = d.Set("physical_params", []map[string]interface{}{physicalParamsMap}); err != nil {
			err = fmt.Errorf("Error setting physical_params: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-physical_params").GetDiag()
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.KubernetesParams) {
		group_id, ok := d.GetOk("auto_proetction_group_id")
		if ok && group_id.(string) != "" {
			autoProetctionGroupId := group_id.(string)
			sourceRegistrationReponseParams.KubernetesParams.AutoProtectConfig.ProtectionGroupID = &autoProetctionGroupId
		}
		kubernetesParamsMap, err := ResourceIbmBackupRecoverySourceRegistrationKubernetesSourceRegistrationParamsToMap(sourceRegistrationReponseParams.KubernetesParams)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "kubernetes_params-to-map").GetDiag()
		}
		if err = d.Set("kubernetes_params", []map[string]interface{}{kubernetesParamsMap}); err != nil {
			err = fmt.Errorf("Error setting kubernetes_params: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-kubernetes_params").GetDiag()
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.SourceID) {
		if err = d.Set("source_id", flex.IntValue(sourceRegistrationReponseParams.SourceID)); err != nil {
			err = fmt.Errorf("Error setting source_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-source_id").GetDiag()
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.SourceInfo) {
		sourceInfoMap, err := ResourceIbmBackupRecoverySourceRegistrationObjectToMap(sourceRegistrationReponseParams.SourceInfo)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "source_info-to-map").GetDiag()
		}
		if err = d.Set("source_info", []map[string]interface{}{sourceInfoMap}); err != nil {
			err = fmt.Errorf("Error setting source_info: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-source_info").GetDiag()
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.AuthenticationStatus) {
		if err = d.Set("authentication_status", sourceRegistrationReponseParams.AuthenticationStatus); err != nil {
			err = fmt.Errorf("Error setting authentication_status: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-authentication_status").GetDiag()
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.RegistrationTimeMsecs) {
		if err = d.Set("registration_time_msecs", flex.IntValue(sourceRegistrationReponseParams.RegistrationTimeMsecs)); err != nil {
			err = fmt.Errorf("Error setting registration_time_msecs: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-registration_time_msecs").GetDiag()
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.LastRefreshedTimeMsecs) {
		if err = d.Set("last_refreshed_time_msecs", flex.IntValue(sourceRegistrationReponseParams.LastRefreshedTimeMsecs)); err != nil {
			err = fmt.Errorf("Error setting last_refreshed_time_msecs: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-last_refreshed_time_msecs").GetDiag()
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.ExternalMetadata) {
		externalMetadataMap, err := ResourceIbmBackupRecoverySourceRegistrationEntityExternalMetadataToMap(sourceRegistrationReponseParams.ExternalMetadata)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "external_metadata-to-map").GetDiag()
		}
		if err = d.Set("external_metadata", []map[string]interface{}{externalMetadataMap}); err != nil {
			err = fmt.Errorf("Error setting external_metadata: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-external_metadata").GetDiag()
		}
	} else {
		if err = d.Set("external_metadata", []interface{}{}); err != nil {
			err = fmt.Errorf("Error setting external_metadata: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "read", "set-external_metadata").GetDiag()
		}
	}

	return nil
}

func resourceIbmBackupRecoverySourceRegistrationUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	endpointType := d.Get("endpoint_type").(string)
	instanceId, region := getInstanceIdAndRegion(d)
	if instanceId != "" && region != "" {
		bmxsession, err := meta.(conns.ClientSession).BluemixSession()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to get clientSession"), "ibm_backup_recovery", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		backupRecoveryClient = getClientWithInstanceEndpoint(backupRecoveryClient, bmxsession, instanceId, region, endpointType)
	}

	tenantId := d.Get("x_ibm_tenant_id").(string)
	registrationId := d.Id()
	if strings.Contains(d.Id(), "::") {
		tenantId = ParseId(d.Id(), "tenantId")
		registrationId = ParseId(d.Id(), "id")
	}

	id, err := strconv.Atoi(registrationId)
	if err != nil {
		return diag.FromErr(err)
	}

	patchData := false
	putData := false

	if d.HasChange("environment") {
		patchData = true
	}

	if d.HasChange("name") ||
		d.HasChange("is_internal_encrypted") ||
		d.HasChange("encryption_key") ||
		d.HasChange("connection_id") ||
		d.HasChange("connections") ||
		d.HasChange("connector_group_id") ||
		d.HasChange("advanced_configs") ||
		d.HasChange("data_source_connection_id") ||
		d.HasChange("kubernetes_params") ||
		d.HasChange("physical_params") {
		putData = true
	}

	if patchData && !putData {
		patchProtectionSourceRegistrationOptions := &backuprecoveryv1.PatchProtectionSourceRegistrationOptions{}
		patchProtectionSourceRegistrationOptions.SetXIBMTenantID(tenantId)

		patchProtectionSourceRegistrationOptions.SetEnvironment(d.Get("environment").(string))
		patchProtectionSourceRegistrationOptions.SetID(int64(id))
		_, _, err = backupRecoveryClient.PatchProtectionSourceRegistrationWithContext(context, patchProtectionSourceRegistrationOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PatchProtectionSourceRegistrationWithContext failed: %s", err.Error()), "ibm_backup_recovery_source_registration", "patch")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	} else {

		updateProtectionSourceRegistrationOptions := &backuprecoveryv1.UpdateProtectionSourceRegistrationOptions{}
		updateProtectionSourceRegistrationOptions.SetXIBMTenantID(tenantId)

		updateProtectionSourceRegistrationOptions.SetEnvironment(d.Get("environment").(string))
		if _, ok := d.GetOk("name"); ok {
			updateProtectionSourceRegistrationOptions.SetName(d.Get("name").(string))
		}
		if _, ok := d.GetOk("is_internal_encrypted"); ok {
			updateProtectionSourceRegistrationOptions.SetIsInternalEncrypted(d.Get("is_internal_encrypted").(bool))
		}
		if _, ok := d.GetOk("encryption_key"); ok {
			updateProtectionSourceRegistrationOptions.SetEncryptionKey(d.Get("encryption_key").(string))
		}

		if _, ok := d.GetOk("connection_id"); ok {
			connId, err := strconv.ParseInt(d.Get("connection_id").(string), 10, 64)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "create", "parse-connection-id").GetDiag()
			}
			updateProtectionSourceRegistrationOptions.SetConnectionID(connId)
		}

		if !d.HasChange("data_source_connection_id") {
			if _, ok := d.GetOk("connection_id"); ok {
				connId, err := strconv.ParseInt(d.Get("connection_id").(string), 10, 64)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("setting connection_id failed: %s", err.Error()), "ibm_backup_recovery_source_registration", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				updateProtectionSourceRegistrationOptions.SetConnectionID(connId)
			}
		} else {
			if _, ok := d.GetOk("data_source_connection_id"); ok {
				updateProtectionSourceRegistrationOptions.SetDataSourceConnectionID(d.Get("data_source_connection_id").(string))
			}
		}

		if _, ok := d.GetOk("connections"); ok {
			var connections []backuprecoveryv1.ConnectionConfig
			for _, v := range d.Get("connections").([]interface{}) {
				value := v.(map[string]interface{})
				connectionsItem, err := ResourceIbmBackupRecoverySourceRegistrationMapToConnectionConfig(value)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "update", "parse-connections").GetDiag()
				}
				connections = append(connections, *connectionsItem)
			}
			updateProtectionSourceRegistrationOptions.SetConnections(connections)
		}
		if _, ok := d.GetOk("connector_group_id"); ok {
			updateProtectionSourceRegistrationOptions.SetConnectorGroupID(int64(d.Get("connector_group_id").(int)))
		}
		if _, ok := d.GetOk("advanced_configs"); ok {
			var advancedConfigs []backuprecoveryv1.KeyValuePair
			for _, v := range d.Get("advanced_configs").([]interface{}) {
				value := v.(map[string]interface{})
				advancedConfigsItem, err := ResourceIbmBackupRecoverySourceRegistrationMapToKeyValuePair(value)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "update", "parse-advanced_configs").GetDiag()
				}
				advancedConfigs = append(advancedConfigs, *advancedConfigsItem)
			}
			updateProtectionSourceRegistrationOptions.SetAdvancedConfigs(advancedConfigs)
		}

		if _, ok := d.GetOk("physical_params"); ok {
			physicalParams, err := ResourceIbmBackupRecoverySourceRegistrationMapToPhysicalSourceRegistrationParams(d.Get("physical_params.0").(map[string]interface{}))
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "update", "parse-physical_params").GetDiag()
			}
			updateProtectionSourceRegistrationOptions.SetPhysicalParams(physicalParams)
		}
		if _, ok := d.GetOk("kubernetes_params"); ok {
			kubernetesParams, err := ResourceIbmBackupRecoverySourceRegistrationMapToKubernetesSourceRegistrationParams(d.Get("kubernetes_params.0").(map[string]interface{}))
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "update", "parse-kubernetes_params").GetDiag()
			}
			updateProtectionSourceRegistrationOptions.SetKubernetesParams(kubernetesParams)
		}
		updateProtectionSourceRegistrationOptions.SetID(int64(id))
		_, _, err = backupRecoveryClient.UpdateProtectionSourceRegistrationWithContext(context, updateProtectionSourceRegistrationOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateProtectionSourceRegistrationWithContext failed: %s", err.Error()), "ibm_backup_recovery_source_registration", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	d.SetId(d.Id())
	return resourceIbmBackupRecoverySourceRegistrationRead(context, d, meta)
}

func ParseId(id string, info string) string {
	if info == "tenantId" {
		return strings.Split(id, "::")[0]
	}
	if info == "id" {
		return strings.Split(id, "::")[1]
	}
	return ""
}

func resourceIbmBackupRecoverySourceRegistrationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_source_registration", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	endpointType := d.Get("endpoint_type").(string)
	instanceId, region := getInstanceIdAndRegion(d)
	if instanceId != "" && region != "" {
		bmxsession, err := meta.(conns.ClientSession).BluemixSession()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to get clientSession"), "ibm_backup_recovery", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		backupRecoveryClient = getClientWithInstanceEndpoint(backupRecoveryClient, bmxsession, instanceId, region, endpointType)
	}

	deleteProtectionSourceRegistrationOptions := &backuprecoveryv1.DeleteProtectionSourceRegistrationOptions{}

	tenantId := d.Get("x_ibm_tenant_id").(string)
	registrationId := d.Id()
	if strings.Contains(d.Id(), "::") {
		tenantId = ParseId(d.Id(), "tenantId")
		registrationId = ParseId(d.Id(), "id")
	}

	id, err := strconv.Atoi(registrationId)
	if err != nil {
		return diag.FromErr(err)
	}

	deleteProtectionSourceRegistrationOptions.SetID(int64(id))
	deleteProtectionSourceRegistrationOptions.SetXIBMTenantID(tenantId)

	_, err = backupRecoveryClient.DeleteProtectionSourceRegistrationWithContext(context, deleteProtectionSourceRegistrationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteProtectionSourceRegistrationWithContext failed: %s", err.Error()), "ibm_backup_recovery_source_registration", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIbmBackupRecoverySourceRegistrationMapToConnectionConfig(modelMap map[string]interface{}) (*backuprecoveryv1.ConnectionConfig, error) {
	model := &backuprecoveryv1.ConnectionConfig{}
	if modelMap["connection_id"] != nil {
		ConnectionID, err := strconv.ParseInt(modelMap["connection_id"].(string), 10, 64)
		if err != nil {
			return model, err
		}
		model.ConnectionID = &ConnectionID
	}
	if modelMap["entity_id"] != nil {
		model.EntityID = core.Int64Ptr(int64(modelMap["entity_id"].(int)))
	}
	if modelMap["connector_group_id"] != nil {
		model.ConnectorGroupID = core.Int64Ptr(int64(modelMap["connector_group_id"].(int)))
	}
	if modelMap["data_source_connection_id"] != nil && modelMap["data_source_connection_id"].(string) != "" {
		model.DataSourceConnectionID = core.StringPtr(modelMap["data_source_connection_id"].(string))
	}
	return model, nil
}

func ResourceIbmBackupRecoverySourceRegistrationMapToKeyValuePair(modelMap map[string]interface{}) (*backuprecoveryv1.KeyValuePair, error) {
	model := &backuprecoveryv1.KeyValuePair{}
	model.Key = core.StringPtr(modelMap["key"].(string))
	model.Value = core.StringPtr(modelMap["value"].(string))
	return model, nil
}

func ResourceIbmBackupRecoverySourceRegistrationMapToKubernetesSourceRegistrationParams(modelMap map[string]interface{}) (*backuprecoveryv1.KubernetesSourceRegistrationParams, error) {
	model := &backuprecoveryv1.KubernetesSourceRegistrationParams{}
	if modelMap["auto_protect_config"] != nil && len(modelMap["auto_protect_config"].([]interface{})) > 0 && modelMap["auto_protect_config"].([]interface{})[0] != nil {
		AutoProtectConfigModel, err := ResourceIbmBackupRecoverySourceRegistrationMapToKubernetesAutoProtectConfig(modelMap["auto_protect_config"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.AutoProtectConfig = AutoProtectConfigModel
	}
	model.ClientPrivateKey = core.StringPtr(modelMap["client_private_key"].(string))
	if modelMap["cohesity_dataprotect_plugin_image_location"] != nil && modelMap["cohesity_dataprotect_plugin_image_location"].(string) != "" {
		model.CohesityDataprotectPluginImageLocation = core.StringPtr(modelMap["cohesity_dataprotect_plugin_image_location"].(string))
	}
	model.DataMoverImageLocation = core.StringPtr(modelMap["data_mover_image_location"].(string))
	if modelMap["datamover_service_type"] != nil && modelMap["datamover_service_type"].(string) != "" {
		model.DatamoverServiceType = core.StringPtr(modelMap["datamover_service_type"].(string))
	}
	if modelMap["default_vlan_params"] != nil && len(modelMap["default_vlan_params"].([]interface{})) > 0 {
		DefaultVlanParamsModel, err := ResourceIbmBackupRecoverySourceRegistrationMapToVlanParams(modelMap["default_vlan_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.DefaultVlanParams = DefaultVlanParamsModel
	}
	model.Endpoint = core.StringPtr(modelMap["endpoint"].(string))
	if modelMap["init_container_image_location"] != nil && modelMap["init_container_image_location"].(string) != "" {
		model.InitContainerImageLocation = core.StringPtr(modelMap["init_container_image_location"].(string))
	}
	model.KubernetesDistribution = core.StringPtr(modelMap["kubernetes_distribution"].(string))
	if modelMap["kubernetes_type"] != nil && modelMap["kubernetes_type"].(string) != "" {
		model.KubernetesType = core.StringPtr(modelMap["kubernetes_type"].(string))
	}
	if modelMap["priority_class_name"] != nil && modelMap["priority_class_name"].(string) != "" {
		model.PriorityClassName = core.StringPtr(modelMap["priority_class_name"].(string))
	}
	if modelMap["resource_annotations"] != nil {
		resourceAnnotations := []backuprecoveryv1.KubernetesLabelObject{}
		for _, resourceAnnotationsItem := range modelMap["resource_annotations"].([]interface{}) {
			resourceAnnotationsItemModel, err := ResourceIbmBackupRecoverySourceRegistrationMapToKubernetesLabelObject(resourceAnnotationsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			resourceAnnotations = append(resourceAnnotations, *resourceAnnotationsItemModel)
		}
		model.ResourceAnnotations = resourceAnnotations
	}
	if modelMap["resource_labels"] != nil {
		resourceLabels := []backuprecoveryv1.KubernetesLabelObject{}
		for _, resourceLabelsItem := range modelMap["resource_labels"].([]interface{}) {
			resourceLabelsItemModel, err := ResourceIbmBackupRecoverySourceRegistrationMapToKubernetesLabelObject(resourceLabelsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			resourceLabels = append(resourceLabels, *resourceLabelsItemModel)
		}
		model.ResourceLabels = resourceLabels
	}
	if modelMap["san_fields"] != nil {
		sanFields := []string{}
		for _, sanFieldsItem := range modelMap["san_fields"].([]interface{}) {
			sanFields = append(sanFields, sanFieldsItem.(string))
		}
		model.SanFields = sanFields
	}
	if modelMap["service_annotations"] != nil {
		serviceAnnotations := []backuprecoveryv1.KubernetesServiceAnnotationObject{}
		for _, serviceAnnotationsItem := range modelMap["service_annotations"].([]interface{}) {
			serviceAnnotationsItemModel, err := ResourceIbmBackupRecoverySourceRegistrationMapToKubernetesServiceAnnotationObject(serviceAnnotationsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			serviceAnnotations = append(serviceAnnotations, *serviceAnnotationsItemModel)
		}
		model.ServiceAnnotations = serviceAnnotations
	}
	if modelMap["velero_aws_plugin_image_location"] != nil && modelMap["velero_aws_plugin_image_location"].(string) != "" {
		model.VeleroAwsPluginImageLocation = core.StringPtr(modelMap["velero_aws_plugin_image_location"].(string))
	}
	if modelMap["velero_image_location"] != nil && modelMap["velero_image_location"].(string) != "" {
		model.VeleroImageLocation = core.StringPtr(modelMap["velero_image_location"].(string))
	}
	if modelMap["velero_openshift_plugin_image_location"] != nil && modelMap["velero_openshift_plugin_image_location"].(string) != "" {
		model.VeleroOpenshiftPluginImageLocation = core.StringPtr(modelMap["velero_openshift_plugin_image_location"].(string))
	}
	if modelMap["vlan_info_vec"] != nil {
		vlanInfoVec := []backuprecoveryv1.KubernetesVlanInfo{}
		for _, vlanInfoVecItem := range modelMap["vlan_info_vec"].([]interface{}) {
			vlanInfoVecItemModel, err := ResourceIbmBackupRecoverySourceRegistrationMapToKubernetesVlanInfo(vlanInfoVecItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			vlanInfoVec = append(vlanInfoVec, *vlanInfoVecItemModel)
		}
		model.VlanInfoVec = vlanInfoVec
	}
	return model, nil
}

func ResourceIbmBackupRecoverySourceRegistrationMapToKubernetesAutoProtectConfig(modelMap map[string]interface{}) (*backuprecoveryv1.KubernetesAutoProtectConfig, error) {
	model := &backuprecoveryv1.KubernetesAutoProtectConfig{}
	if modelMap["error_message"] != nil && modelMap["error_message"].(string) != "" {
		model.ErrorMessage = core.StringPtr(modelMap["error_message"].(string))
	}
	model.IsDefaultAutoProtected = core.BoolPtr(modelMap["is_default_auto_protected"].(bool))
	model.PolicyID = core.StringPtr(modelMap["policy_id"].(string))
	if modelMap["protection_group_id"] != nil && modelMap["protection_group_id"].(string) != "" {
		model.ProtectionGroupID = core.StringPtr(modelMap["protection_group_id"].(string))
	}
	if modelMap["storage_domain_id"] != nil {
		model.StorageDomainID = core.Int64Ptr(int64(modelMap["storage_domain_id"].(int)))
	}
	return model, nil
}

func ResourceIbmBackupRecoverySourceRegistrationMapToVlanParams(modelMap map[string]interface{}) (*backuprecoveryv1.VlanParams, error) {
	model := &backuprecoveryv1.VlanParams{}
	if modelMap["disable_vlan"] != nil {
		model.DisableVlan = core.BoolPtr(modelMap["disable_vlan"].(bool))
	}
	if modelMap["interface_name"] != nil && modelMap["interface_name"].(string) != "" {
		model.InterfaceName = core.StringPtr(modelMap["interface_name"].(string))
	}
	if modelMap["vlan_id"] != nil {
		model.VlanID = core.Int64Ptr(int64(modelMap["vlan_id"].(int)))
	}
	return model, nil
}

func ResourceIbmBackupRecoverySourceRegistrationMapToKubernetesLabelObject(modelMap map[string]interface{}) (*backuprecoveryv1.KubernetesLabelObject, error) {
	model := &backuprecoveryv1.KubernetesLabelObject{}
	model.Key = core.StringPtr(modelMap["key"].(string))
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	return model, nil
}

func ResourceIbmBackupRecoverySourceRegistrationMapToKubernetesServiceAnnotationObject(modelMap map[string]interface{}) (*backuprecoveryv1.KubernetesServiceAnnotationObject, error) {
	model := &backuprecoveryv1.KubernetesServiceAnnotationObject{}
	if modelMap["key"] != nil && modelMap["key"].(string) != "" {
		model.Key = core.StringPtr(modelMap["key"].(string))
	}
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	return model, nil
}

func ResourceIbmBackupRecoverySourceRegistrationMapToKubernetesVlanInfo(modelMap map[string]interface{}) (*backuprecoveryv1.KubernetesVlanInfo, error) {
	model := &backuprecoveryv1.KubernetesVlanInfo{}
	if modelMap["service_annotations"] != nil {
		serviceAnnotations := []backuprecoveryv1.KubernetesServiceAnnotationObject{}
		for _, serviceAnnotationsItem := range modelMap["service_annotations"].([]interface{}) {
			serviceAnnotationsItemModel, err := ResourceIbmBackupRecoverySourceRegistrationMapToKubernetesServiceAnnotationObject(serviceAnnotationsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			serviceAnnotations = append(serviceAnnotations, *serviceAnnotationsItemModel)
		}
		model.ServiceAnnotations = serviceAnnotations
	}
	if modelMap["vlan_params"] != nil && len(modelMap["vlan_params"].([]interface{})) > 0 {
		VlanParamsModel, err := ResourceIbmBackupRecoverySourceRegistrationMapToVlanParams(modelMap["vlan_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.VlanParams = VlanParamsModel
	}
	return model, nil
}

func ResourceIbmBackupRecoverySourceRegistrationMapToPhysicalSourceRegistrationParams(modelMap map[string]interface{}) (*backuprecoveryv1.PhysicalSourceRegistrationParams, error) {
	model := &backuprecoveryv1.PhysicalSourceRegistrationParams{}
	model.Endpoint = core.StringPtr(modelMap["endpoint"].(string))
	if modelMap["force_register"] != nil {
		model.ForceRegister = core.BoolPtr(modelMap["force_register"].(bool))
	}
	if modelMap["host_type"] != nil && modelMap["host_type"].(string) != "" {
		model.HostType = core.StringPtr(modelMap["host_type"].(string))
	}
	if modelMap["physical_type"] != nil && modelMap["physical_type"].(string) != "" {
		model.PhysicalType = core.StringPtr(modelMap["physical_type"].(string))
	}
	if modelMap["applications"] != nil {
		applications := []string{}
		for _, applicationsItem := range modelMap["applications"].([]interface{}) {
			applications = append(applications, applicationsItem.(string))
		}
		model.Applications = applications
	}
	return model, nil
}

func ResourceIbmBackupRecoverySourceRegistrationConnectionConfigToMap(model *backuprecoveryv1.ConnectionConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConnectionID != nil {
		connId := strconv.Itoa(int(*model.ConnectionID))
		modelMap["connection_id"] = flex.StringValue(&connId)
	}
	if model.EntityID != nil {
		modelMap["entity_id"] = flex.IntValue(model.EntityID)
	}
	if model.ConnectorGroupID != nil {
		modelMap["connector_group_id"] = flex.IntValue(model.ConnectorGroupID)
	}
	if model.DataSourceConnectionID != nil {
		modelMap["data_source_connection_id"] = *model.DataSourceConnectionID
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationKeyValuePairToMap(model *backuprecoveryv1.KeyValuePair) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = *model.Key
	modelMap["value"] = *model.Value
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationPhysicalSourceRegistrationParamsToMap(model *backuprecoveryv1.PhysicalSourceRegistrationParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["endpoint"] = *model.Endpoint
	if model.ForceRegister != nil {
		modelMap["force_register"] = *model.ForceRegister
	}
	if model.HostType != nil {
		modelMap["host_type"] = *model.HostType
	}
	if model.PhysicalType != nil {
		modelMap["physical_type"] = *model.PhysicalType
	}
	if model.Applications != nil {
		modelMap["applications"] = model.Applications
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationKubernetesSourceRegistrationParamsToMap(model *backuprecoveryv1.KubernetesSourceRegistrationParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AutoProtectConfig != nil {
		autoProtectConfigMap, err := ResourceIbmBackupRecoverySourceRegistrationKubernetesAutoProtectConfigToMap(model.AutoProtectConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["auto_protect_config"] = []map[string]interface{}{autoProtectConfigMap}
	}
	if model.ClientPrivateKey != nil {
		modelMap["client_private_key"] = *model.ClientPrivateKey
	}
	if model.CohesityDataprotectPluginImageLocation != nil {
		modelMap["cohesity_dataprotect_plugin_image_location"] = *model.CohesityDataprotectPluginImageLocation
	}
	if model.DataMoverImageLocation != nil {
		modelMap["data_mover_image_location"] = *model.DataMoverImageLocation
	}
	if model.DatamoverServiceType != nil {
		modelMap["datamover_service_type"] = *model.DatamoverServiceType
	}
	if model.DefaultVlanParams != nil {
		defaultVlanParamsMap, err := ResourceIbmBackupRecoverySourceRegistrationVlanParamsToMap(model.DefaultVlanParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["default_vlan_params"] = []map[string]interface{}{defaultVlanParamsMap}
	}
	if model.Endpoint != nil {
		modelMap["endpoint"] = *model.Endpoint
	}
	if model.InitContainerImageLocation != nil {
		modelMap["init_container_image_location"] = *model.InitContainerImageLocation
	}
	if model.KubernetesDistribution != nil {
		modelMap["kubernetes_distribution"] = *model.KubernetesDistribution
	}
	if model.KubernetesType != nil {
		modelMap["kubernetes_type"] = *model.KubernetesType
	}
	if model.PriorityClassName != nil {
		modelMap["priority_class_name"] = *model.PriorityClassName
	}
	if model.ResourceAnnotations != nil {
		resourceAnnotations := []map[string]interface{}{}
		for _, resourceAnnotationsItem := range model.ResourceAnnotations {
			resourceAnnotationsItemMap, err := ResourceIbmBackupRecoverySourceRegistrationKubernetesLabelObjectToMap(&resourceAnnotationsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			resourceAnnotations = append(resourceAnnotations, resourceAnnotationsItemMap)
		}
		modelMap["resource_annotations"] = resourceAnnotations
	}
	if model.ResourceLabels != nil {
		resourceLabels := []map[string]interface{}{}
		for _, resourceLabelsItem := range model.ResourceLabels {
			resourceLabelsItemMap, err := ResourceIbmBackupRecoverySourceRegistrationKubernetesLabelObjectToMap(&resourceLabelsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			resourceLabels = append(resourceLabels, resourceLabelsItemMap)
		}
		modelMap["resource_labels"] = resourceLabels
	}
	if model.SanFields != nil {
		modelMap["san_fields"] = model.SanFields
	}
	if model.ServiceAnnotations != nil {
		serviceAnnotations := []map[string]interface{}{}
		for _, serviceAnnotationsItem := range model.ServiceAnnotations {
			serviceAnnotationsItemMap, err := ResourceIbmBackupRecoverySourceRegistrationKubernetesServiceAnnotationObjectToMap(&serviceAnnotationsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			serviceAnnotations = append(serviceAnnotations, serviceAnnotationsItemMap)
		}
		modelMap["service_annotations"] = serviceAnnotations
	}
	if model.VeleroAwsPluginImageLocation != nil {
		modelMap["velero_aws_plugin_image_location"] = *model.VeleroAwsPluginImageLocation
	}
	if model.VeleroImageLocation != nil {
		modelMap["velero_image_location"] = *model.VeleroImageLocation
	}
	if model.VeleroOpenshiftPluginImageLocation != nil {
		modelMap["velero_openshift_plugin_image_location"] = *model.VeleroOpenshiftPluginImageLocation
	}
	if model.VlanInfoVec != nil {
		vlanInfoVec := []map[string]interface{}{}
		for _, vlanInfoVecItem := range model.VlanInfoVec {
			vlanInfoVecItemMap, err := ResourceIbmBackupRecoverySourceRegistrationKubernetesVlanInfoToMap(&vlanInfoVecItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			vlanInfoVec = append(vlanInfoVec, vlanInfoVecItemMap)
		}
		modelMap["vlan_info_vec"] = vlanInfoVec
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationKubernetesAutoProtectConfigToMap(model *backuprecoveryv1.KubernetesAutoProtectConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ErrorMessage != nil {
		modelMap["error_message"] = *model.ErrorMessage
	}
	if model.IsDefaultAutoProtected != nil {
		modelMap["is_default_auto_protected"] = *model.IsDefaultAutoProtected
	}
	if model.PolicyID != nil {
		modelMap["policy_id"] = *model.PolicyID
	}
	if model.ProtectionGroupID != nil {
		modelMap["protection_group_id"] = *model.ProtectionGroupID
	}
	if model.StorageDomainID != nil {
		modelMap["storage_domain_id"] = flex.IntValue(model.StorageDomainID)
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationVlanParamsToMap(model *backuprecoveryv1.VlanParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DisableVlan != nil {
		modelMap["disable_vlan"] = *model.DisableVlan
	}
	if model.InterfaceName != nil {
		modelMap["interface_name"] = *model.InterfaceName
	}
	if model.VlanID != nil {
		modelMap["vlan_id"] = flex.IntValue(model.VlanID)
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationKubernetesLabelObjectToMap(model *backuprecoveryv1.KubernetesLabelObject) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = *model.Key
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationKubernetesServiceAnnotationObjectToMap(model *backuprecoveryv1.KubernetesServiceAnnotationObject) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Key != nil {
		modelMap["key"] = *model.Key
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationKubernetesVlanInfoToMap(model *backuprecoveryv1.KubernetesVlanInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ServiceAnnotations != nil {
		serviceAnnotations := []map[string]interface{}{}
		for _, serviceAnnotationsItem := range model.ServiceAnnotations {
			serviceAnnotationsItemMap, err := ResourceIbmBackupRecoverySourceRegistrationKubernetesServiceAnnotationObjectToMap(&serviceAnnotationsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			serviceAnnotations = append(serviceAnnotations, serviceAnnotationsItemMap)
		}
		modelMap["service_annotations"] = serviceAnnotations
	}
	if model.VlanParams != nil {
		vlanParamsMap, err := ResourceIbmBackupRecoverySourceRegistrationVlanParamsToMap(model.VlanParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["vlan_params"] = []map[string]interface{}{vlanParamsMap}
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationObjectToMap(model *backuprecoveryv1.Object) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.SourceID != nil {
		modelMap["source_id"] = flex.IntValue(model.SourceID)
	}
	if model.SourceName != nil {
		modelMap["source_name"] = *model.SourceName
	}
	if model.Environment != nil {
		modelMap["environment"] = *model.Environment
	}
	if model.ObjectHash != nil {
		modelMap["object_hash"] = *model.ObjectHash
	}
	if model.ObjectType != nil {
		modelMap["object_type"] = *model.ObjectType
	}
	if model.LogicalSizeBytes != nil {
		modelMap["logical_size_bytes"] = flex.IntValue(model.LogicalSizeBytes)
	}
	if model.UUID != nil {
		modelMap["uuid"] = *model.UUID
	}
	if model.GlobalID != nil {
		modelMap["global_id"] = *model.GlobalID
	}
	if model.ProtectionType != nil {
		modelMap["protection_type"] = *model.ProtectionType
	}
	if model.SharepointSiteSummary != nil {
		sharepointSiteSummaryMap, err := ResourceIbmBackupRecoverySourceRegistrationSharepointObjectParamsToMap(model.SharepointSiteSummary)
		if err != nil {
			return modelMap, err
		}
		modelMap["sharepoint_site_summary"] = []map[string]interface{}{sharepointSiteSummaryMap}
	}
	if model.OsType != nil {
		modelMap["os_type"] = *model.OsType
	}
	if model.ChildObjects != nil {
		childObjects := []map[string]interface{}{}
		for _, childObjectsItem := range model.ChildObjects {
			childObjectsItemMap, err := ResourceIbmBackupRecoverySourceRegistrationObjectSummaryToMap(&childObjectsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			childObjects = append(childObjects, childObjectsItemMap)
		}
		modelMap["child_objects"] = childObjects
	}
	if model.VCenterSummary != nil {
		vCenterSummaryMap, err := ResourceIbmBackupRecoverySourceRegistrationObjectTypeVCenterParamsToMap(model.VCenterSummary)
		if err != nil {
			return modelMap, err
		}
		modelMap["v_center_summary"] = []map[string]interface{}{vCenterSummaryMap}
	}
	if model.WindowsClusterSummary != nil {
		windowsClusterSummaryMap, err := ResourceIbmBackupRecoverySourceRegistrationObjectTypeWindowsClusterParamsToMap(model.WindowsClusterSummary)
		if err != nil {
			return modelMap, err
		}
		modelMap["windows_cluster_summary"] = []map[string]interface{}{windowsClusterSummaryMap}
	}
	if model.ProtectionStats != nil {
		protectionStats := []map[string]interface{}{}
		for _, protectionStatsItem := range model.ProtectionStats {
			protectionStatsItemMap, err := ResourceIbmBackupRecoverySourceRegistrationObjectProtectionStatsSummaryToMap(&protectionStatsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			protectionStats = append(protectionStats, protectionStatsItemMap)
		}
		modelMap["protection_stats"] = protectionStats
	}
	if model.Permissions != nil {
		permissionsMap, err := ResourceIbmBackupRecoverySourceRegistrationPermissionInfoToMap(model.Permissions)
		if err != nil {
			return modelMap, err
		}
		modelMap["permissions"] = []map[string]interface{}{permissionsMap}
	}
	if model.MssqlParams != nil {
		mssqlParamsMap, err := ResourceIbmBackupRecoverySourceRegistrationObjectMssqlParamsToMap(model.MssqlParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["mssql_params"] = []map[string]interface{}{mssqlParamsMap}
	}
	if model.PhysicalParams != nil {
		physicalParamsMap, err := ResourceIbmBackupRecoverySourceRegistrationObjectPhysicalParamsToMap(model.PhysicalParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_params"] = []map[string]interface{}{physicalParamsMap}
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationSharepointObjectParamsToMap(model *backuprecoveryv1.SharepointObjectParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SiteWebURL != nil {
		modelMap["site_web_url"] = *model.SiteWebURL
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationObjectSummaryToMap(model *backuprecoveryv1.ObjectSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.SourceID != nil {
		modelMap["source_id"] = flex.IntValue(model.SourceID)
	}
	if model.SourceName != nil {
		modelMap["source_name"] = *model.SourceName
	}
	if model.Environment != nil {
		modelMap["environment"] = *model.Environment
	}
	if model.ObjectHash != nil {
		modelMap["object_hash"] = *model.ObjectHash
	}
	if model.ObjectType != nil {
		modelMap["object_type"] = *model.ObjectType
	}
	if model.LogicalSizeBytes != nil {
		modelMap["logical_size_bytes"] = flex.IntValue(model.LogicalSizeBytes)
	}
	if model.UUID != nil {
		modelMap["uuid"] = *model.UUID
	}
	if model.GlobalID != nil {
		modelMap["global_id"] = *model.GlobalID
	}
	if model.ProtectionType != nil {
		modelMap["protection_type"] = *model.ProtectionType
	}
	if model.SharepointSiteSummary != nil {
		sharepointSiteSummaryMap, err := ResourceIbmBackupRecoverySourceRegistrationSharepointObjectParamsToMap(model.SharepointSiteSummary)
		if err != nil {
			return modelMap, err
		}
		modelMap["sharepoint_site_summary"] = []map[string]interface{}{sharepointSiteSummaryMap}
	}
	if model.OsType != nil {
		modelMap["os_type"] = *model.OsType
	}
	if model.ChildObjects != nil {
		childObjects := []map[string]interface{}{}
		for _, childObjectsItem := range model.ChildObjects {
			childObjectsItemMap, err := ResourceIbmBackupRecoverySourceRegistrationObjectSummaryToMap(&childObjectsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			childObjects = append(childObjects, childObjectsItemMap)
		}
		modelMap["child_objects"] = childObjects
	}
	if model.VCenterSummary != nil {
		vCenterSummaryMap, err := ResourceIbmBackupRecoverySourceRegistrationObjectTypeVCenterParamsToMap(model.VCenterSummary)
		if err != nil {
			return modelMap, err
		}
		modelMap["v_center_summary"] = []map[string]interface{}{vCenterSummaryMap}
	}
	if model.WindowsClusterSummary != nil {
		windowsClusterSummaryMap, err := ResourceIbmBackupRecoverySourceRegistrationObjectTypeWindowsClusterParamsToMap(model.WindowsClusterSummary)
		if err != nil {
			return modelMap, err
		}
		modelMap["windows_cluster_summary"] = []map[string]interface{}{windowsClusterSummaryMap}
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationObjectTypeVCenterParamsToMap(model *backuprecoveryv1.ObjectTypeVCenterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.IsCloudEnv != nil {
		modelMap["is_cloud_env"] = *model.IsCloudEnv
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationObjectTypeWindowsClusterParamsToMap(model *backuprecoveryv1.ObjectTypeWindowsClusterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClusterSourceType != nil {
		modelMap["cluster_source_type"] = *model.ClusterSourceType
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationObjectProtectionStatsSummaryToMap(model *backuprecoveryv1.ObjectProtectionStatsSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Environment != nil {
		modelMap["environment"] = *model.Environment
	}
	if model.ProtectedCount != nil {
		modelMap["protected_count"] = flex.IntValue(model.ProtectedCount)
	}
	if model.UnprotectedCount != nil {
		modelMap["unprotected_count"] = flex.IntValue(model.UnprotectedCount)
	}
	if model.DeletedProtectedCount != nil {
		modelMap["deleted_protected_count"] = flex.IntValue(model.DeletedProtectedCount)
	}
	if model.ProtectedSizeBytes != nil {
		modelMap["protected_size_bytes"] = flex.IntValue(model.ProtectedSizeBytes)
	}
	if model.UnprotectedSizeBytes != nil {
		modelMap["unprotected_size_bytes"] = flex.IntValue(model.UnprotectedSizeBytes)
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationPermissionInfoToMap(model *backuprecoveryv1.PermissionInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ObjectID != nil {
		modelMap["object_id"] = flex.IntValue(model.ObjectID)
	}
	if model.Users != nil {
		users := []map[string]interface{}{}
		for _, usersItem := range model.Users {
			usersItemMap, err := ResourceIbmBackupRecoverySourceRegistrationUserToMap(&usersItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			users = append(users, usersItemMap)
		}
		modelMap["users"] = users
	}
	if model.Groups != nil {
		groups := []map[string]interface{}{}
		for _, groupsItem := range model.Groups {
			groupsItemMap, err := ResourceIbmBackupRecoverySourceRegistrationGroupToMap(&groupsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			groups = append(groups, groupsItemMap)
		}
		modelMap["groups"] = groups
	}
	if model.Tenant != nil {
		tenantMap, err := ResourceIbmBackupRecoverySourceRegistrationTenantToMap(model.Tenant)
		if err != nil {
			return modelMap, err
		}
		modelMap["tenant"] = []map[string]interface{}{tenantMap}
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationUserToMap(model *backuprecoveryv1.User) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Sid != nil {
		modelMap["sid"] = *model.Sid
	}
	if model.Domain != nil {
		modelMap["domain"] = *model.Domain
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationGroupToMap(model *backuprecoveryv1.Group) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Sid != nil {
		modelMap["sid"] = *model.Sid
	}
	if model.Domain != nil {
		modelMap["domain"] = *model.Domain
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationTenantToMap(model *backuprecoveryv1.Tenant) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CreatedAtTimeMsecs != nil && *(model.CreatedAtTimeMsecs) != 0 {
		modelMap["created_at_time_msecs"] = flex.IntValue(model.CreatedAtTimeMsecs)
	}
	if model.DeletedAtTimeMsecs != nil && *(model.DeletedAtTimeMsecs) != 0 {
		modelMap["deleted_at_time_msecs"] = flex.IntValue(model.DeletedAtTimeMsecs)
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.ExternalVendorMetadata != nil {
		externalVendorMetadataMap, err := ResourceIbmBackupRecoverySourceRegistrationExternalVendorTenantMetadataToMap(model.ExternalVendorMetadata)
		if err != nil {
			return modelMap, err
		}
		modelMap["external_vendor_metadata"] = []map[string]interface{}{externalVendorMetadataMap}
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.IsManagedOnHelios != nil {
		modelMap["is_managed_on_helios"] = *model.IsManagedOnHelios
	}
	if model.LastUpdatedAtTimeMsecs != nil && *(model.LastUpdatedAtTimeMsecs) != 0 {
		modelMap["last_updated_at_time_msecs"] = flex.IntValue(model.LastUpdatedAtTimeMsecs)
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Network != nil {
		networkMap, err := ResourceIbmBackupRecoverySourceRegistrationTenantNetworkToMap(model.Network)
		if err != nil {
			return modelMap, err
		}
		modelMap["network"] = []map[string]interface{}{networkMap}
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationExternalVendorTenantMetadataToMap(model *backuprecoveryv1.ExternalVendorTenantMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.IbmTenantMetadataParams != nil {
		ibmTenantMetadataParamsMap, err := ResourceIbmBackupRecoverySourceRegistrationIbmTenantMetadataParamsToMap(model.IbmTenantMetadataParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["ibm_tenant_metadata_params"] = []map[string]interface{}{ibmTenantMetadataParamsMap}
	}
	modelMap["type"] = *model.Type
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationIbmTenantMetadataParamsToMap(model *backuprecoveryv1.IbmTenantMetadataParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AccountID != nil {
		modelMap["account_id"] = *model.AccountID
	}
	if model.Crn != nil {
		modelMap["crn"] = *model.Crn
	}
	if model.CustomProperties != nil {
		customProperties := []map[string]interface{}{}
		for _, customPropertiesItem := range model.CustomProperties {
			customPropertiesItemMap, err := ResourceIbmBackupRecoverySourceRegistrationExternalVendorCustomPropertiesToMap(&customPropertiesItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			customProperties = append(customProperties, customPropertiesItemMap)
		}
		modelMap["custom_properties"] = customProperties
	}
	if model.LivenessMode != nil {
		modelMap["liveness_mode"] = *model.LivenessMode
	}
	if model.MetricsConfig != nil {
		metricsConfigMap, err := ResourceIbmBackupRecoverySourceRegistrationIbmTenantMetricsConfigToMap(model.MetricsConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["metrics_config"] = []map[string]interface{}{metricsConfigMap}
	}
	if model.OwnershipMode != nil {
		modelMap["ownership_mode"] = *model.OwnershipMode
	}
	if model.PlanID != nil {
		modelMap["plan_id"] = *model.PlanID
	}
	if model.ResourceGroupID != nil {
		modelMap["resource_group_id"] = *model.ResourceGroupID
	}
	if model.ResourceInstanceID != nil {
		modelMap["resource_instance_id"] = *model.ResourceInstanceID
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationExternalVendorCustomPropertiesToMap(model *backuprecoveryv1.ExternalVendorCustomProperties) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Key != nil {
		modelMap["key"] = *model.Key
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationIbmTenantMetricsConfigToMap(model *backuprecoveryv1.IbmTenantMetricsConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CosResourceConfig != nil {
		cosResourceConfigMap, err := ResourceIbmBackupRecoverySourceRegistrationIbmTenantCOSResourceConfigToMap(model.CosResourceConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["cos_resource_config"] = []map[string]interface{}{cosResourceConfigMap}
	}
	if model.IamMetricsConfig != nil {
		iamMetricsConfigMap, err := ResourceIbmBackupRecoverySourceRegistrationIbmTenantIAMMetricsConfigToMap(model.IamMetricsConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["iam_metrics_config"] = []map[string]interface{}{iamMetricsConfigMap}
	}
	if model.MeteringConfig != nil {
		meteringConfigMap, err := ResourceIbmBackupRecoverySourceRegistrationIbmTenantMeteringConfigToMap(model.MeteringConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["metering_config"] = []map[string]interface{}{meteringConfigMap}
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationIbmTenantCOSResourceConfigToMap(model *backuprecoveryv1.IbmTenantCOSResourceConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ResourceURL != nil {
		modelMap["resource_url"] = *model.ResourceURL
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationIbmTenantIAMMetricsConfigToMap(model *backuprecoveryv1.IbmTenantIAMMetricsConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.IAMURL != nil {
		modelMap["iam_url"] = *model.IAMURL
	}
	if model.BillingApiKeySecretID != nil {
		modelMap["billing_api_key_secret_id"] = *model.BillingApiKeySecretID
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationIbmTenantMeteringConfigToMap(model *backuprecoveryv1.IbmTenantMeteringConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PartIds != nil {
		modelMap["part_ids"] = model.PartIds
	}
	if model.SubmissionIntervalInSecs != nil {
		modelMap["submission_interval_in_secs"] = flex.IntValue(model.SubmissionIntervalInSecs)
	}
	if model.URL != nil {
		modelMap["url"] = *model.URL
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationTenantNetworkToMap(model *backuprecoveryv1.TenantNetwork) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["connector_enabled"] = *model.ConnectorEnabled
	if model.ClusterHostname != nil {
		modelMap["cluster_hostname"] = *model.ClusterHostname
	}
	if model.ClusterIps != nil {
		modelMap["cluster_ips"] = model.ClusterIps
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationObjectMssqlParamsToMap(model *backuprecoveryv1.ObjectMssqlParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AagInfo != nil {
		aagInfoMap, err := ResourceIbmBackupRecoverySourceRegistrationAAGInfoToMap(model.AagInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["aag_info"] = []map[string]interface{}{aagInfoMap}
	}
	if model.HostInfo != nil {
		hostInfoMap, err := ResourceIbmBackupRecoverySourceRegistrationHostInformationToMap(model.HostInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["host_info"] = []map[string]interface{}{hostInfoMap}
	}
	if model.IsEncrypted != nil {
		modelMap["is_encrypted"] = *model.IsEncrypted
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationAAGInfoToMap(model *backuprecoveryv1.AAGInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.ObjectID != nil {
		modelMap["object_id"] = flex.IntValue(model.ObjectID)
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationHostInformationToMap(model *backuprecoveryv1.HostInformation) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Environment != nil {
		modelMap["environment"] = *model.Environment
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationObjectPhysicalParamsToMap(model *backuprecoveryv1.ObjectPhysicalParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EnableSystemBackup != nil {
		modelMap["enable_system_backup"] = *model.EnableSystemBackup
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationEntityExternalMetadataToMap(model *backuprecoveryv1.EntityExternalMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaintenanceModeConfig != nil {
		maintenanceModeConfigMap, err := ResourceIbmBackupRecoverySourceRegistrationMaintenanceModeConfigToMap(model.MaintenanceModeConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["maintenance_mode_config"] = []map[string]interface{}{maintenanceModeConfigMap}
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationMaintenanceModeConfigToMap(model *backuprecoveryv1.MaintenanceModeConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ActivationTimeIntervals != nil {
		activationTimeIntervals := []map[string]interface{}{}
		for _, activationTimeIntervalsItem := range model.ActivationTimeIntervals {
			activationTimeIntervalsItemMap, err := ResourceIbmBackupRecoverySourceRegistrationTimeRangeUsecsToMap(&activationTimeIntervalsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			activationTimeIntervals = append(activationTimeIntervals, activationTimeIntervalsItemMap)
		}
		modelMap["activation_time_intervals"] = activationTimeIntervals
	}
	if model.MaintenanceSchedule != nil {
		maintenanceScheduleMap, err := ResourceIbmBackupRecoverySourceRegistrationScheduleToMap(model.MaintenanceSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["maintenance_schedule"] = []map[string]interface{}{maintenanceScheduleMap}
	}
	if model.UserMessage != nil {
		modelMap["user_message"] = *model.UserMessage
	}
	if model.WorkflowInterventionSpecList != nil {
		workflowInterventionSpecList := []map[string]interface{}{}
		for _, workflowInterventionSpecListItem := range model.WorkflowInterventionSpecList {
			workflowInterventionSpecListItemMap, err := ResourceIbmBackupRecoverySourceRegistrationWorkflowInterventionSpecToMap(&workflowInterventionSpecListItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			workflowInterventionSpecList = append(workflowInterventionSpecList, workflowInterventionSpecListItemMap)
		}
		modelMap["workflow_intervention_spec_list"] = workflowInterventionSpecList
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationTimeRangeUsecsToMap(model *backuprecoveryv1.TimeRangeUsecs) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationScheduleToMap(model *backuprecoveryv1.Schedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PeriodicTimeWindows != nil {
		periodicTimeWindows := []map[string]interface{}{}
		for _, periodicTimeWindowsItem := range model.PeriodicTimeWindows {
			periodicTimeWindowsItemMap, err := ResourceIbmBackupRecoverySourceRegistrationTimeWindowToMap(&periodicTimeWindowsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			periodicTimeWindows = append(periodicTimeWindows, periodicTimeWindowsItemMap)
		}
		modelMap["periodic_time_windows"] = periodicTimeWindows
	}
	if model.ScheduleType != nil {
		modelMap["schedule_type"] = *model.ScheduleType
	}
	if model.TimeRanges != nil {
		timeRanges := []map[string]interface{}{}
		for _, timeRangesItem := range model.TimeRanges {
			timeRangesItemMap, err := ResourceIbmBackupRecoverySourceRegistrationTimeRangeUsecsToMap(&timeRangesItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			timeRanges = append(timeRanges, timeRangesItemMap)
		}
		modelMap["time_ranges"] = timeRanges
	}
	if model.Timezone != nil {
		modelMap["timezone"] = *model.Timezone
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationTimeWindowToMap(model *backuprecoveryv1.TimeWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DayOfTheWeek != nil {
		modelMap["day_of_the_week"] = *model.DayOfTheWeek
	}
	if model.EndTime != nil {
		endTimeMap, err := ResourceIbmBackupRecoverySourceRegistrationTimeToMap(model.EndTime)
		if err != nil {
			return modelMap, err
		}
		modelMap["end_time"] = []map[string]interface{}{endTimeMap}
	}
	if model.StartTime != nil {
		startTimeMap, err := ResourceIbmBackupRecoverySourceRegistrationTimeToMap(model.StartTime)
		if err != nil {
			return modelMap, err
		}
		modelMap["start_time"] = []map[string]interface{}{startTimeMap}
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationTimeToMap(model *backuprecoveryv1.Time) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Hour != nil {
		modelMap["hour"] = flex.IntValue(model.Hour)
	}
	if model.Minute != nil {
		modelMap["minute"] = flex.IntValue(model.Minute)
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoverySourceRegistrationWorkflowInterventionSpecToMap(model *backuprecoveryv1.WorkflowInterventionSpec) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["intervention"] = *model.Intervention
	modelMap["workflow_type"] = *model.WorkflowType
	return modelMap, nil
}
