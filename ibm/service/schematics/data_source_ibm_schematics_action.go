// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func DataSourceIBMSchematicsAction() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceIBMSchematicsActionRead,

		Schema: map[string]*schema.Schema{
			"action_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Action Id.  Use GET /actions API to look up the Action Ids in your IBM Cloud account.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique name of your action. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. **Example** you can use the name to stop action.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Action description.",
			},
			"location": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource-group name for an action. By default, an action is created in `Default` resource group.",
			},
			"bastion_connection_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of connection to be used when connecting to bastion host. If the `inventory_connection_type=winrm`, then `bastion_connection_type` is not supported.",
			},
			"inventory_connection_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of connection to be used when connecting to remote host. **Note** Currently, WinRM supports only Windows system with the public IPs and do not support Bastion host.",
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Action tags.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_state": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "User defined status of the Schematics object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User-defined states  * `draft` Object can be modified; can be used by Jobs run by the author, during execution  * `live` Object can be modified; can be used by Jobs during execution  * `locked` Object cannot be modified; can be used by Jobs during execution  * `disable` Object can be modified. cannot be used by Jobs during execution.",
						},
						"set_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the User who set the state of the Object.",
						},
						"set_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "When the User who set the state of the Object.",
						},
					},
				},
			},
			"source_readme_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of the `README` file, for the source URL.",
			},
			"source": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Source of templates, playbooks, or controls.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of source for the Template.",
						},
						"git": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The connection details to the Git source repository.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"computed_git_repo_url": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The complete URL which is computed by the **git_repo_url**, **git_repo_folder**, and **branch**.",
									},
									"git_repo_url": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL to the Git repository that can be used to clone the template.",
									},
									"git_token": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Personal Access Token (PAT) to connect to the Git URLs.",
									},
									"git_repo_folder": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the folder in the Git repository, that contains the template.",
									},
									"git_release": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the release tag that are used to fetch the Git repository.",
									},
									"git_branch": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the branch that are used to fetch the Git repository.",
									},
								},
							},
						},
						"catalog": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The connection details to the IBM Cloud Catalog source.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"catalog_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the private catalog.",
									},
									"offering_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of an offering in the IBM Cloud Catalog.",
									},
									"offering_version": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The version string of an offering in the IBM Cloud Catalog.",
									},
									"offering_kind": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of an offering, in the IBM Cloud Catalog.",
									},
									"offering_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of an offering in the IBM Cloud Catalog.",
									},
									"offering_version_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of an offering version the IBM Cloud Catalog.",
									},
									"offering_repo_url": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The repository URL of an offering, in the IBM Cloud Catalog.",
									},
								},
							},
						},
					},
				},
			},
			"source_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of source for the Template.",
			},
			"command_parameter": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Schematics job command parameter (playbook-name).",
			},
			"inventory": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Target inventory record ID, used by the action or ansible playbook.",
			},
			"credentials": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "credentials of the Action.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the credential variable.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The credential value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
						},
						"use_default": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "An user editable metadata for the credential variables.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Type of the variable.",
									},
									"aliases": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of aliases for the variable name.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the meta data.",
									},
									"cloud_data_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cloud data type of the credential variable. eg. api_key, iam_token, profile_id.",
									},
									"default_value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Default value for the variable only if the override value is not specified.",
									},
									"link_status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the link.",
									},
									"immutable": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is the variable readonly ?.",
									},
									"hidden": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If **true**, the variable is not displayed on UI or Command line.",
									},
									"required": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If the variable required?.",
									},
									"position": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The relative position of this variable in a list.",
									},
									"group_by": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The display name of the group this variable belongs to.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The source of this meta-data.",
									},
								},
							},
						},
						"link": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reference link to the variable value By default the expression points to `$self.value`.",
						},
					},
				},
			},
			"bastion": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Describes a bastion resource.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Bastion Name(Unique).",
						},
						"host": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Reference to the Inventory resource definition.",
						},
					},
				},
			},
			"bastion_credential": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "User editable credential variable data and system generated reference to the value.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the credential variable.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The credential value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
						},
						"use_default": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "An user editable metadata for the credential variables.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Type of the variable.",
									},
									"aliases": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of aliases for the variable name.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the meta data.",
									},
									"cloud_data_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cloud data type of the credential variable. eg. api_key, iam_token, profile_id.",
									},
									"default_value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Default value for the variable only if the override value is not specified.",
									},
									"link_status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the link.",
									},
									"immutable": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is the variable readonly ?.",
									},
									"hidden": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If **true**, the variable is not displayed on UI or Command line.",
									},
									"required": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If the variable required?.",
									},
									"position": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The relative position of this variable in a list.",
									},
									"group_by": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The display name of the group this variable belongs to.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The source of this meta-data.",
									},
								},
							},
						},
						"link": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reference link to the variable value By default the expression points to `$self.value`.",
						},
					},
				},
			},
			"targets_ini": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Inventory of host and host group for the playbook in `INI` file format. For example, `\"targets_ini\": \"[webserverhost]  172.22.192.6  [dbhost] 172.22.192.5\"`. For more information, about an inventory host group syntax, see [Inventory host groups](https://cloud.ibm.com/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps).",
			},
			"action_inputs": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Input variables for the Action.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the variable. For example, `name = \"inventory username\"`.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
						},
						"use_default": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "An user editable metadata for the variables.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Type of the variable.",
									},
									"aliases": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of aliases for the variable name.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the meta data.",
									},
									"cloud_data_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
									},
									"default_value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Default value for the variable only if the override value is not specified.",
									},
									"link_status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the link.",
									},
									"secure": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is the variable secure or sensitive ?.",
									},
									"immutable": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is the variable readonly ?.",
									},
									"hidden": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If **true**, the variable is not displayed on UI or Command line.",
									},
									"required": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If the variable required?.",
									},
									"options": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"min_value": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum value of the variable. Applicable for the integer type.",
									},
									"max_value": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum value of the variable. Applicable for the integer type.",
									},
									"min_length": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum length of the variable value. Applicable for the string type.",
									},
									"max_length": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum length of the variable value. Applicable for the string type.",
									},
									"matches": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The regex for the variable value.",
									},
									"position": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The relative position of this variable in a list.",
									},
									"group_by": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The display name of the group this variable belongs to.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The source of this meta-data.",
									},
								},
							},
						},
						"link": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reference link to the variable value By default the expression points to `$self.value`.",
						},
					},
				},
			},
			"action_outputs": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output variables for the Action.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the variable. For example, `name = \"inventory username\"`.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
						},
						"use_default": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "An user editable metadata for the variables.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Type of the variable.",
									},
									"aliases": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of aliases for the variable name.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the meta data.",
									},
									"cloud_data_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
									},
									"default_value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Default value for the variable only if the override value is not specified.",
									},
									"link_status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the link.",
									},
									"secure": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is the variable secure or sensitive ?.",
									},
									"immutable": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is the variable readonly ?.",
									},
									"hidden": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If **true**, the variable is not displayed on UI or Command line.",
									},
									"required": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If the variable required?.",
									},
									"options": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"min_value": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum value of the variable. Applicable for the integer type.",
									},
									"max_value": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum value of the variable. Applicable for the integer type.",
									},
									"min_length": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum length of the variable value. Applicable for the string type.",
									},
									"max_length": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum length of the variable value. Applicable for the string type.",
									},
									"matches": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The regex for the variable value.",
									},
									"position": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The relative position of this variable in a list.",
									},
									"group_by": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The display name of the group this variable belongs to.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The source of this meta-data.",
									},
								},
							},
						},
						"link": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reference link to the variable value By default the expression points to `$self.value`.",
						},
					},
				},
			},
			"settings": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Environment variables for the Action.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the variable. For example, `name = \"inventory username\"`.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
						},
						"use_default": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "An user editable metadata for the variables.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Type of the variable.",
									},
									"aliases": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of aliases for the variable name.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the meta data.",
									},
									"cloud_data_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
									},
									"default_value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Default value for the variable only if the override value is not specified.",
									},
									"link_status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the link.",
									},
									"secure": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is the variable secure or sensitive ?.",
									},
									"immutable": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is the variable readonly ?.",
									},
									"hidden": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If **true**, the variable is not displayed on UI or Command line.",
									},
									"required": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If the variable required?.",
									},
									"options": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"min_value": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum value of the variable. Applicable for the integer type.",
									},
									"max_value": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum value of the variable. Applicable for the integer type.",
									},
									"min_length": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum length of the variable value. Applicable for the string type.",
									},
									"max_length": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum length of the variable value. Applicable for the string type.",
									},
									"matches": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The regex for the variable value.",
									},
									"position": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The relative position of this variable in a list.",
									},
									"group_by": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The display name of the group this variable belongs to.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The source of this meta-data.",
									},
								},
							},
						},
						"link": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reference link to the variable value By default the expression points to `$self.value`.",
						},
					},
				},
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Action ID.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Action Cloud Resource Name.",
			},
			"account": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Action account ID.",
			},
			"source_created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Action Playbook Source creation time.",
			},
			"source_created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "E-mail address of user who created the Action Playbook Source.",
			},
			"source_updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The action playbook updation time.",
			},
			"source_updated_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "E-mail address of user who updated the action playbook source.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Action creation time.",
			},
			"created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "E-mail address of the user who created an action.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Action updation time.",
			},
			"updated_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "E-mail address of the user who updated an action.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Computed state of the Action.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status_code": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of automation (workspace or action).",
						},
						"status_job_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Job id reference for this status.",
						},
						"status_message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Automation status message - to be displayed along with the status_code.",
						},
					},
				},
			},
			"playbook_names": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Playbook names retrieved from the repository.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sys_lock": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "System lock status.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sys_locked": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is the automation locked by a Schematic job ?.",
						},
						"sys_locked_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the User who performed the job, that lead to the locking of the automation.",
						},
						"sys_locked_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "When the User performed the job that lead to locking of the automation ?.",
						},
					},
				},
			},
		},
	}
}

func DataSourceIBMSchematicsActionRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getActionOptions := &schematicsv1.GetActionOptions{}

	getActionOptions.SetActionID(d.Get("action_id").(string))

	action, response, err := schematicsClient.GetActionWithContext(context, getActionOptions)
	if err != nil {
		log.Printf("[DEBUG] GetActionWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetActionWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s", *getActionOptions.ActionID))

	if err = d.Set("name", action.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}

	if err = d.Set("description", action.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}

	if err = d.Set("location", action.Location); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting location: %s", err))
	}

	if err = d.Set("resource_group", action.ResourceGroup); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group: %s", err))
	}

	if err = d.Set("bastion_connection_type", action.BastionConnectionType); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting bastion_connection_type: %s", err))
	}

	if err = d.Set("inventory_connection_type", action.InventoryConnectionType); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting inventory_connection_type: %s", err))
	}

	userState := []map[string]interface{}{}
	if action.UserState != nil {
		modelMap, err := DataSourceIBMSchematicsActionUserStateToMap(action.UserState)
		if err != nil {
			return diag.FromErr(err)
		}
		userState = append(userState, modelMap)
	}
	if err = d.Set("user_state", userState); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting user_state %s", err))
	}

	if err = d.Set("source_readme_url", action.SourceReadmeURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting source_readme_url: %s", err))
	}

	source := []map[string]interface{}{}
	if action.Source != nil {
		modelMap, err := DataSourceIBMSchematicsActionExternalSourceToMap(action.Source)
		if err != nil {
			return diag.FromErr(err)
		}
		source = append(source, modelMap)
	}
	if err = d.Set("source", source); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting source %s", err))
	}

	if err = d.Set("source_type", action.SourceType); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting source_type: %s", err))
	}

	if err = d.Set("command_parameter", action.CommandParameter); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting command_parameter: %s", err))
	}

	if err = d.Set("inventory", action.Inventory); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting inventory: %s", err))
	}

	credentials := []map[string]interface{}{}
	if action.Credentials != nil {
		for _, modelItem := range action.Credentials {
			modelMap, err := DataSourceIBMSchematicsActionCredentialVariableDataToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			credentials = append(credentials, modelMap)
		}
	}
	if err = d.Set("credentials", credentials); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting credentials %s", err))
	}

	bastion := []map[string]interface{}{}
	if action.Bastion != nil {
		modelMap, err := DataSourceIBMSchematicsActionBastionResourceDefinitionToMap(action.Bastion)
		if err != nil {
			return diag.FromErr(err)
		}
		bastion = append(bastion, modelMap)
	}
	if err = d.Set("bastion", bastion); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting bastion %s", err))
	}

	bastionCredential := []map[string]interface{}{}
	if action.BastionCredential != nil {
		modelMap, err := DataSourceIBMSchematicsActionCredentialVariableDataToMap(action.BastionCredential)
		if err != nil {
			return diag.FromErr(err)
		}
		bastionCredential = append(bastionCredential, modelMap)
	}
	if err = d.Set("bastion_credential", bastionCredential); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting bastion_credential %s", err))
	}

	if err = d.Set("targets_ini", action.TargetsIni); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting targets_ini: %s", err))
	}

	actionInputs := []map[string]interface{}{}
	if action.Inputs != nil {
		for _, modelItem := range action.Inputs {
			modelMap, err := DataSourceIBMSchematicsActionVariableDataToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			actionInputs = append(actionInputs, modelMap)
		}
	}
	if err = d.Set("action_inputs", actionInputs); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting action_inputs %s", err))
	}

	actionOutputs := []map[string]interface{}{}
	if action.Outputs != nil {
		for _, modelItem := range action.Outputs {
			modelMap, err := DataSourceIBMSchematicsActionVariableDataToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			actionOutputs = append(actionOutputs, modelMap)
		}
	}
	if err = d.Set("action_outputs", actionOutputs); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting action_outputs %s", err))
	}

	settings := []map[string]interface{}{}
	if action.Settings != nil {
		for _, modelItem := range action.Settings {
			modelMap, err := DataSourceIBMSchematicsActionVariableDataToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			settings = append(settings, modelMap)
		}
	}
	if err = d.Set("settings", settings); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting settings %s", err))
	}

	if err = d.Set("id", action.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting id: %s", err))
	}

	if err = d.Set("crn", action.Crn); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}

	if err = d.Set("account", action.Account); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting account: %s", err))
	}

	if err = d.Set("source_created_at", flex.DateTimeToString(action.SourceCreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting source_created_at: %s", err))
	}

	if err = d.Set("source_created_by", action.SourceCreatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting source_created_by: %s", err))
	}

	if err = d.Set("source_updated_at", flex.DateTimeToString(action.SourceUpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting source_updated_at: %s", err))
	}

	if err = d.Set("source_updated_by", action.SourceUpdatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting source_updated_by: %s", err))
	}

	if err = d.Set("created_at", flex.DateTimeToString(action.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}

	if err = d.Set("created_by", action.CreatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_by: %s", err))
	}

	if err = d.Set("updated_at", flex.DateTimeToString(action.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}

	if err = d.Set("updated_by", action.UpdatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_by: %s", err))
	}

	state := []map[string]interface{}{}
	if action.State != nil {
		modelMap, err := DataSourceIBMSchematicsActionActionStateToMap(action.State)
		if err != nil {
			return diag.FromErr(err)
		}
		state = append(state, modelMap)
	}
	if err = d.Set("state", state); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting state %s", err))
	}

	sysLock := []map[string]interface{}{}
	if action.SysLock != nil {
		modelMap, err := DataSourceIBMSchematicsActionSystemLockToMap(action.SysLock)
		if err != nil {
			return diag.FromErr(err)
		}
		sysLock = append(sysLock, modelMap)
	}
	if err = d.Set("sys_lock", sysLock); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting sys_lock %s", err))
	}

	return nil
}

func DataSourceIBMSchematicsActionUserStateToMap(model *schematicsv1.UserState) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.State != nil {
		modelMap["state"] = *model.State
	}
	if model.SetBy != nil {
		modelMap["set_by"] = *model.SetBy
	}
	if model.SetAt != nil {
		modelMap["set_at"] = model.SetAt.String()
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsActionExternalSourceToMap(model *schematicsv1.ExternalSource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SourceType != nil {
		modelMap["source_type"] = *model.SourceType
	}
	if model.Git != nil {
		gitMap, err := DataSourceIBMSchematicsActionExternalSourceGitToMap(model.Git)
		if err != nil {
			return modelMap, err
		}
		modelMap["git"] = []map[string]interface{}{gitMap}
	}
	if model.Catalog != nil {
		catalogMap, err := DataSourceIBMSchematicsActionExternalSourceCatalogToMap(model.Catalog)
		if err != nil {
			return modelMap, err
		}
		modelMap["catalog"] = []map[string]interface{}{catalogMap}
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsActionExternalSourceGitToMap(model *schematicsv1.ExternalSourceGit) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ComputedGitRepoURL != nil {
		modelMap["computed_git_repo_url"] = *model.ComputedGitRepoURL
	}
	if model.GitRepoURL != nil {
		modelMap["git_repo_url"] = *model.GitRepoURL
	}
	if model.GitToken != nil {
		modelMap["git_token"] = *model.GitToken
	}
	if model.GitRepoFolder != nil {
		modelMap["git_repo_folder"] = *model.GitRepoFolder
	}
	if model.GitRelease != nil {
		modelMap["git_release"] = *model.GitRelease
	}
	if model.GitBranch != nil {
		modelMap["git_branch"] = *model.GitBranch
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsActionExternalSourceCatalogToMap(model *schematicsv1.ExternalSourceCatalog) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CatalogName != nil {
		modelMap["catalog_name"] = *model.CatalogName
	}
	if model.OfferingName != nil {
		modelMap["offering_name"] = *model.OfferingName
	}
	if model.OfferingVersion != nil {
		modelMap["offering_version"] = *model.OfferingVersion
	}
	if model.OfferingKind != nil {
		modelMap["offering_kind"] = *model.OfferingKind
	}
	if model.OfferingID != nil {
		modelMap["offering_id"] = *model.OfferingID
	}
	if model.OfferingVersionID != nil {
		modelMap["offering_version_id"] = *model.OfferingVersionID
	}
	if model.OfferingRepoURL != nil {
		modelMap["offering_repo_url"] = *model.OfferingRepoURL
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsActionCredentialVariableDataToMap(model *schematicsv1.VariableData) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	if model.UseDefault != nil {
		modelMap["use_default"] = *model.UseDefault
	}
	if model.Metadata != nil {
		metadataMap, err := DataSourceIBMSchematicsActionCredentialVariableMetadataToMap(model.Metadata)
		if err != nil {
			return modelMap, err
		}
		modelMap["metadata"] = []map[string]interface{}{metadataMap}
	}
	if model.Link != nil {
		modelMap["link"] = *model.Link
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsActionCredentialVariableMetadataToMap(model *schematicsv1.VariableMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Aliases != nil {
		modelMap["aliases"] = model.Aliases
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.CloudDataType != nil {
		modelMap["cloud_data_type"] = *model.CloudDataType
	}
	if model.DefaultValue != nil {
		modelMap["default_value"] = *model.DefaultValue
	}
	if model.LinkStatus != nil {
		modelMap["link_status"] = *model.LinkStatus
	}
	if model.Immutable != nil {
		modelMap["immutable"] = *model.Immutable
	}
	if model.Hidden != nil {
		modelMap["hidden"] = *model.Hidden
	}
	if model.Required != nil {
		modelMap["required"] = *model.Required
	}
	if model.Position != nil {
		modelMap["position"] = *model.Position
	}
	if model.GroupBy != nil {
		modelMap["group_by"] = *model.GroupBy
	}
	if model.Source != nil {
		modelMap["source"] = *model.Source
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsActionBastionResourceDefinitionToMap(model *schematicsv1.BastionResourceDefinition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Host != nil {
		modelMap["host"] = *model.Host
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsActionVariableDataToMap(model *schematicsv1.VariableData) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	if model.UseDefault != nil {
		modelMap["use_default"] = *model.UseDefault
	}
	if model.Metadata != nil {
		metadataMap, err := DataSourceIBMSchematicsActionVariableMetadataToMap(model.Metadata)
		if err != nil {
			return modelMap, err
		}
		modelMap["metadata"] = []map[string]interface{}{metadataMap}
	}
	if model.Link != nil {
		modelMap["link"] = *model.Link
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsActionVariableMetadataToMap(model *schematicsv1.VariableMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Aliases != nil {
		modelMap["aliases"] = model.Aliases
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.CloudDataType != nil {
		modelMap["cloud_data_type"] = *model.CloudDataType
	}
	if model.DefaultValue != nil {
		modelMap["default_value"] = *model.DefaultValue
	}
	if model.LinkStatus != nil {
		modelMap["link_status"] = *model.LinkStatus
	}
	if model.Secure != nil {
		modelMap["secure"] = *model.Secure
	}
	if model.Immutable != nil {
		modelMap["immutable"] = *model.Immutable
	}
	if model.Hidden != nil {
		modelMap["hidden"] = *model.Hidden
	}
	if model.Required != nil {
		modelMap["required"] = *model.Required
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	if model.MinValue != nil {
		modelMap["min_value"] = *model.MinValue
	}
	if model.MaxValue != nil {
		modelMap["max_value"] = *model.MaxValue
	}
	if model.MinLength != nil {
		modelMap["min_length"] = *model.MinLength
	}
	if model.MaxLength != nil {
		modelMap["max_length"] = *model.MaxLength
	}
	if model.Matches != nil {
		modelMap["matches"] = *model.Matches
	}
	if model.Position != nil {
		modelMap["position"] = *model.Position
	}
	if model.GroupBy != nil {
		modelMap["group_by"] = *model.GroupBy
	}
	if model.Source != nil {
		modelMap["source"] = *model.Source
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsActionActionStateToMap(model *schematicsv1.ActionState) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.StatusCode != nil {
		modelMap["status_code"] = *model.StatusCode
	}
	if model.StatusJobID != nil {
		modelMap["status_job_id"] = *model.StatusJobID
	}
	if model.StatusMessage != nil {
		modelMap["status_message"] = *model.StatusMessage
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsActionSystemLockToMap(model *schematicsv1.SystemLock) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SysLocked != nil {
		modelMap["sys_locked"] = *model.SysLocked
	}
	if model.SysLockedBy != nil {
		modelMap["sys_locked_by"] = *model.SysLockedBy
	}
	if model.SysLockedAt != nil {
		modelMap["sys_locked_at"] = model.SysLockedAt.String()
	}
	return modelMap, nil
}
