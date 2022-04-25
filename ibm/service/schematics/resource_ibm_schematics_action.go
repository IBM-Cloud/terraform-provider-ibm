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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func ResourceIBMSchematicsAction() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIBMSchematicsActionCreate,
		ReadContext:   ResourceIBMSchematicsActionRead,
		UpdateContext: ResourceIBMSchematicsActionUpdate,
		DeleteContext: ResourceIBMSchematicsActionDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The unique name of your action. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. **Example** you can use the name to stop action.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Action description.",
			},
			"location": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_schematics_action", "location"),
				Description:  "List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Resource-group name for an action. By default, an action is created in `Default` resource group.",
			},
			"bastion_connection_type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_schematics_action", "bastion_connection_type"),
				Description:  "Type of connection to be used when connecting to bastion host. If the `inventory_connection_type=winrm`, then `bastion_connection_type` is not supported.",
			},
			"inventory_connection_type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_schematics_action", "inventory_connection_type"),
				Description:  "Type of connection to be used when connecting to remote host. **Note** Currently, WinRM supports only Windows system with the public IPs and do not support Bastion host.",
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Action tags.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"user_state": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "User defined status of the Schematics object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "User-defined states  * `draft` Object can be modified; can be used by Jobs run by the author, during execution  * `live` Object can be modified; can be used by Jobs during execution  * `locked` Object cannot be modified; can be used by Jobs during execution  * `disable` Object can be modified. cannot be used by Jobs during execution.",
						},
						"set_by": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Name of the User who set the state of the Object.",
						},
						"set_at": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "When the User who set the state of the Object.",
						},
					},
				},
			},
			"source_readme_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL of the `README` file, for the source URL.",
			},
			"source": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Source of templates, playbooks, or controls.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Type of source for the Template.",
						},
						"git": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The connection details to the Git source repository.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"computed_git_repo_url": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The complete URL which is computed by the **git_repo_url**, **git_repo_folder**, and **branch**.",
									},
									"git_repo_url": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The URL to the Git repository that can be used to clone the template.",
									},
									"git_token": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Personal Access Token (PAT) to connect to the Git URLs.",
									},
									"git_repo_folder": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of the folder in the Git repository, that contains the template.",
									},
									"git_release": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of the release tag that are used to fetch the Git repository.",
									},
									"git_branch": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of the branch that are used to fetch the Git repository.",
									},
								},
							},
						},
						"catalog": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The connection details to the IBM Cloud Catalog source.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"catalog_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of the private catalog.",
									},
									"offering_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of an offering in the IBM Cloud Catalog.",
									},
									"offering_version": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The version string of an offering in the IBM Cloud Catalog.",
									},
									"offering_kind": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The type of an offering, in the IBM Cloud Catalog.",
									},
									"offering_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The ID of an offering in the IBM Cloud Catalog.",
									},
									"offering_version_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The ID of an offering version the IBM Cloud Catalog.",
									},
									"offering_repo_url": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The repository URL of an offering, in the IBM Cloud Catalog.",
									},
								},
							},
						},
					},
				},
			},
			"source_type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_schematics_action", "source_type"),
				Description:  "Type of source for the Template.",
			},
			"command_parameter": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Schematics job command parameter (playbook-name).",
			},
			"inventory": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Target inventory record ID, used by the action or ansible playbook.",
			},
			"credentials": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "credentials of the Action.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the credential variable.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The credential value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
						},
						"use_default": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "An user editable metadata for the credential variables.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Type of the variable.",
									},
									"aliases": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of aliases for the variable name.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The description of the meta data.",
									},
									"cloud_data_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Cloud data type of the credential variable. eg. api_key, iam_token, profile_id.",
									},
									"default_value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Default value for the variable only if the override value is not specified.",
									},
									"link_status": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The status of the link.",
									},
									"immutable": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Is the variable readonly ?.",
									},
									"hidden": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If **true**, the variable is not displayed on UI or Command line.",
									},
									"required": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If the variable required?.",
									},
									"position": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The relative position of this variable in a list.",
									},
									"group_by": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The display name of the group this variable belongs to.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The source of this meta-data.",
									},
								},
							},
						},
						"link": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The reference link to the variable value By default the expression points to `$self.value`.",
						},
					},
				},
			},
			"bastion": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Describes a bastion resource.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Bastion Name(Unique).",
						},
						"host": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Reference to the Inventory resource definition.",
						},
					},
				},
			},
			"bastion_credential": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "User editable credential variable data and system generated reference to the value.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the credential variable.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The credential value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
						},
						"use_default": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "An user editable metadata for the credential variables.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Type of the variable.",
									},
									"aliases": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of aliases for the variable name.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The description of the meta data.",
									},
									"cloud_data_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Cloud data type of the credential variable. eg. api_key, iam_token, profile_id.",
									},
									"default_value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Default value for the variable only if the override value is not specified.",
									},
									"link_status": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The status of the link.",
									},
									"immutable": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Is the variable readonly ?.",
									},
									"hidden": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If **true**, the variable is not displayed on UI or Command line.",
									},
									"required": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If the variable required?.",
									},
									"position": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The relative position of this variable in a list.",
									},
									"group_by": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The display name of the group this variable belongs to.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The source of this meta-data.",
									},
								},
							},
						},
						"link": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The reference link to the variable value By default the expression points to `$self.value`.",
						},
					},
				},
			},
			"targets_ini": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Inventory of host and host group for the playbook in `INI` file format. For example, `\"targets_ini\": \"[webserverhost]  172.22.192.6  [dbhost] 172.22.192.5\"`. For more information, about an inventory host group syntax, see [Inventory host groups](https://cloud.ibm.com/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps).",
			},
			"action_inputs": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Input variables for the Action.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the variable. For example, `name = \"inventory username\"`.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
						},
						"use_default": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "An user editable metadata for the variables.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Type of the variable.",
									},
									"aliases": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of aliases for the variable name.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The description of the meta data.",
									},
									"cloud_data_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
									},
									"default_value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Default value for the variable only if the override value is not specified.",
									},
									"link_status": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The status of the link.",
									},
									"secure": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Is the variable secure or sensitive ?.",
									},
									"immutable": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Is the variable readonly ?.",
									},
									"hidden": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If **true**, the variable is not displayed on UI or Command line.",
									},
									"required": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If the variable required?.",
									},
									"options": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"min_value": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The minimum value of the variable. Applicable for the integer type.",
									},
									"max_value": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The maximum value of the variable. Applicable for the integer type.",
									},
									"min_length": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The minimum length of the variable value. Applicable for the string type.",
									},
									"max_length": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The maximum length of the variable value. Applicable for the string type.",
									},
									"matches": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The regex for the variable value.",
									},
									"position": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The relative position of this variable in a list.",
									},
									"group_by": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The display name of the group this variable belongs to.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The source of this meta-data.",
									},
								},
							},
						},
						"link": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The reference link to the variable value By default the expression points to `$self.value`.",
						},
					},
				},
			},
			"action_outputs": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Output variables for the Action.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the variable. For example, `name = \"inventory username\"`.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
						},
						"use_default": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "An user editable metadata for the variables.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Type of the variable.",
									},
									"aliases": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of aliases for the variable name.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The description of the meta data.",
									},
									"cloud_data_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
									},
									"default_value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Default value for the variable only if the override value is not specified.",
									},
									"link_status": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The status of the link.",
									},
									"secure": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Is the variable secure or sensitive ?.",
									},
									"immutable": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Is the variable readonly ?.",
									},
									"hidden": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If **true**, the variable is not displayed on UI or Command line.",
									},
									"required": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If the variable required?.",
									},
									"options": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"min_value": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The minimum value of the variable. Applicable for the integer type.",
									},
									"max_value": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The maximum value of the variable. Applicable for the integer type.",
									},
									"min_length": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The minimum length of the variable value. Applicable for the string type.",
									},
									"max_length": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The maximum length of the variable value. Applicable for the string type.",
									},
									"matches": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The regex for the variable value.",
									},
									"position": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The relative position of this variable in a list.",
									},
									"group_by": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The display name of the group this variable belongs to.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The source of this meta-data.",
									},
								},
							},
						},
						"link": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The reference link to the variable value By default the expression points to `$self.value`.",
						},
					},
				},
			},
			"settings": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Environment variables for the Action.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the variable. For example, `name = \"inventory username\"`.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
						},
						"use_default": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "An user editable metadata for the variables.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Type of the variable.",
									},
									"aliases": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of aliases for the variable name.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The description of the meta data.",
									},
									"cloud_data_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
									},
									"default_value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Default value for the variable only if the override value is not specified.",
									},
									"link_status": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The status of the link.",
									},
									"secure": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Is the variable secure or sensitive ?.",
									},
									"immutable": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Is the variable readonly ?.",
									},
									"hidden": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If **true**, the variable is not displayed on UI or Command line.",
									},
									"required": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If the variable required?.",
									},
									"options": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"min_value": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The minimum value of the variable. Applicable for the integer type.",
									},
									"max_value": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The maximum value of the variable. Applicable for the integer type.",
									},
									"min_length": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The minimum length of the variable value. Applicable for the string type.",
									},
									"max_length": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The maximum length of the variable value. Applicable for the string type.",
									},
									"matches": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The regex for the variable value.",
									},
									"position": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The relative position of this variable in a list.",
									},
									"group_by": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The display name of the group this variable belongs to.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The source of this meta-data.",
									},
								},
							},
						},
						"link": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The reference link to the variable value By default the expression points to `$self.value`.",
						},
					},
				},
			},
			"state": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Computed state of the Action.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status_code": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Status of automation (workspace or action).",
						},
						"status_job_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Job id reference for this status.",
						},
						"status_message": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Automation status message - to be displayed along with the status_code.",
						},
					},
				},
			},
			"sys_lock": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "System lock status.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sys_locked": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Is the automation locked by a Schematic job ?.",
						},
						"sys_locked_by": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the User who performed the job, that lead to the locking of the automation.",
						},
						"sys_locked_at": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "When the User performed the job that lead to locking of the automation ?.",
						},
					},
				},
			},
			"x_github_token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template.",
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
			"playbook_names": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Playbook names retrieved from the repository.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceIBMSchematicsActionValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 1)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "location",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "eu-de, eu-gb, us-east, us-south",
		},
		validate.ValidateSchema{
			Identifier:                 "bastion_connection_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "ssh",
		},
		validate.ValidateSchema{
			Identifier:                 "inventory_connection_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "ssh, winrm",
		},
		validate.ValidateSchema{
			Identifier:                 "source_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "external_scm, git_hub, git_hub_enterprise, git_lab, ibm_cloud_catalog, ibm_git_lab, local",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_schematics_action", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIBMSchematicsActionCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createActionOptions := &schematicsv1.CreateActionOptions{}

	if _, ok := d.GetOk("name"); ok {
		createActionOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("description"); ok {
		createActionOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("location"); ok {
		createActionOptions.SetLocation(d.Get("location").(string))
	}
	if _, ok := d.GetOk("resource_group"); ok {
		createActionOptions.SetResourceGroup(d.Get("resource_group").(string))
	}
	if _, ok := d.GetOk("bastion_connection_type"); ok {
		createActionOptions.SetBastionConnectionType(d.Get("bastion_connection_type").(string))
	}
	if _, ok := d.GetOk("inventory_connection_type"); ok {
		createActionOptions.SetInventoryConnectionType(d.Get("inventory_connection_type").(string))
	}
	if _, ok := d.GetOk("tags"); ok {
		createActionOptions.SetTags(d.Get("tags").([]string))
	}
	if _, ok := d.GetOk("user_state"); ok {
		userStateModel, err := ResourceIBMSchematicsActionMapToUserState(d.Get("user_state.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createActionOptions.SetUserState(userStateModel)
	}
	if _, ok := d.GetOk("source_readme_url"); ok {
		createActionOptions.SetSourceReadmeURL(d.Get("source_readme_url").(string))
	}
	if _, ok := d.GetOk("source"); ok {
		sourceModel, err := ResourceIBMSchematicsActionMapToExternalSource(d.Get("source.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createActionOptions.SetSource(sourceModel)
	}
	if _, ok := d.GetOk("source_type"); ok {
		createActionOptions.SetSourceType(d.Get("source_type").(string))
	}
	if _, ok := d.GetOk("command_parameter"); ok {
		createActionOptions.SetCommandParameter(d.Get("command_parameter").(string))
	}
	if _, ok := d.GetOk("inventory"); ok {
		createActionOptions.SetInventory(d.Get("inventory").(string))
	}
	if _, ok := d.GetOk("credentials"); ok {
		var credentials []schematicsv1.VariableData
		for _, e := range d.Get("credentials").([]interface{}) {
			value := e.(map[string]interface{})
			credentialsItem, err := ResourceIBMSchematicsActionMapToCredentialVariableData(value)
			if err != nil {
				return diag.FromErr(err)
			}
			credentials = append(credentials, *credentialsItem)
		}
		createActionOptions.SetCredentials(credentials)
	}
	if _, ok := d.GetOk("bastion"); ok {
		bastionModel, err := ResourceIBMSchematicsActionMapToBastionResourceDefinition(d.Get("bastion.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createActionOptions.SetBastion(bastionModel)
	}
	if _, ok := d.GetOk("bastion_credential"); ok {
		bastionCredentialModel, err := ResourceIBMSchematicsActionMapToCredentialVariableData(d.Get("bastion_credential.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createActionOptions.SetBastionCredential(bastionCredentialModel)
	}
	if _, ok := d.GetOk("targets_ini"); ok {
		createActionOptions.SetTargetsIni(d.Get("targets_ini").(string))
	}
	if _, ok := d.GetOk("action_inputs"); ok {
		var actionInputs []schematicsv1.VariableData
		for _, e := range d.Get("action_inputs").([]interface{}) {
			value := e.(map[string]interface{})
			actionInputsItem, err := ResourceIBMSchematicsActionMapToVariableData(value)
			if err != nil {
				return diag.FromErr(err)
			}
			actionInputs = append(actionInputs, *actionInputsItem)
		}
		createActionOptions.SetInputs(actionInputs)
	}
	if _, ok := d.GetOk("action_outputs"); ok {
		var actionOutputs []schematicsv1.VariableData
		for _, e := range d.Get("action_outputs").([]interface{}) {
			value := e.(map[string]interface{})
			actionOutputsItem, err := ResourceIBMSchematicsActionMapToVariableData(value)
			if err != nil {
				return diag.FromErr(err)
			}
			actionOutputs = append(actionOutputs, *actionOutputsItem)
		}
		createActionOptions.SetOutputs(actionOutputs)
	}
	if _, ok := d.GetOk("settings"); ok {
		var settings []schematicsv1.VariableData
		for _, e := range d.Get("settings").([]interface{}) {
			value := e.(map[string]interface{})
			settingsItem, err := ResourceIBMSchematicsActionMapToVariableData(value)
			if err != nil {
				return diag.FromErr(err)
			}
			settings = append(settings, *settingsItem)
		}
		createActionOptions.SetSettings(settings)
	}
	if _, ok := d.GetOk("state"); ok {
		stateModel, err := ResourceIBMSchematicsActionMapToActionState(d.Get("state.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createActionOptions.SetState(stateModel)
	}
	if _, ok := d.GetOk("sys_lock"); ok {
		sysLockModel, err := ResourceIBMSchematicsActionMapToSystemLock(d.Get("sys_lock.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createActionOptions.SetSysLock(sysLockModel)
	}
	if _, ok := d.GetOk("x_github_token"); ok {
		createActionOptions.SetXGithubToken(d.Get("x_github_token").(string))
	}

	action, response, err := schematicsClient.CreateActionWithContext(context, createActionOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateActionWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateActionWithContext failed %s\n%s", err, response))
	}

	d.SetId(*action.ID)

	return ResourceIBMSchematicsActionRead(context, d, meta)
}

func ResourceIBMSchematicsActionRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getActionOptions := &schematicsv1.GetActionOptions{}

	getActionOptions.SetActionID(d.Id())

	action, response, err := schematicsClient.GetActionWithContext(context, getActionOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetActionWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetActionWithContext failed %s\n%s", err, response))
	}

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
	if action.Tags != nil {
		if err = d.Set("tags", action.Tags); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting tags: %s", err))
		}
	}
	if action.UserState != nil {
		userStateMap, err := ResourceIBMSchematicsActionUserStateToMap(action.UserState)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("user_state", []map[string]interface{}{userStateMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting user_state: %s", err))
		}
	}
	if err = d.Set("source_readme_url", action.SourceReadmeURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting source_readme_url: %s", err))
	}
	if action.Source != nil {
		sourceMap, err := ResourceIBMSchematicsActionExternalSourceToMap(action.Source)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("source", []map[string]interface{}{sourceMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting source: %s", err))
		}
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
		for _, credentialsItem := range action.Credentials {
			credentialsItemMap, err := ResourceIBMSchematicsActionCredentialVariableDataToMap(&credentialsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			credentials = append(credentials, credentialsItemMap)
		}
	}
	if err = d.Set("credentials", credentials); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting credentials: %s", err))
	}
	if action.Bastion != nil {
		bastionMap, err := ResourceIBMSchematicsActionBastionResourceDefinitionToMap(action.Bastion)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("bastion", []map[string]interface{}{bastionMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting bastion: %s", err))
		}
	}
	if action.BastionCredential != nil {
		bastionCredentialMap, err := ResourceIBMSchematicsActionCredentialVariableDataToMap(action.BastionCredential)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("bastion_credential", []map[string]interface{}{bastionCredentialMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting bastion_credential: %s", err))
		}
	}
	if err = d.Set("targets_ini", action.TargetsIni); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting targets_ini: %s", err))
	}
	actionInputs := []map[string]interface{}{}
	if action.Inputs != nil {
		for _, actionInputsItem := range action.Inputs {
			actionInputsItemMap, err := ResourceIBMSchematicsActionVariableDataToMap(&actionInputsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			actionInputs = append(actionInputs, actionInputsItemMap)
		}
	}
	if err = d.Set("action_inputs", actionInputs); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting action_inputs: %s", err))
	}
	actionOutputs := []map[string]interface{}{}
	if action.Outputs != nil {
		for _, actionOutputsItem := range action.Outputs {
			actionOutputsItemMap, err := ResourceIBMSchematicsActionVariableDataToMap(&actionOutputsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			actionOutputs = append(actionOutputs, actionOutputsItemMap)
		}
	}
	if err = d.Set("action_outputs", actionOutputs); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting action_outputs: %s", err))
	}
	settings := []map[string]interface{}{}
	if action.Settings != nil {
		for _, settingsItem := range action.Settings {
			settingsItemMap, err := ResourceIBMSchematicsActionVariableDataToMap(&settingsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			settings = append(settings, settingsItemMap)
		}
	}
	if err = d.Set("settings", settings); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting settings: %s", err))
	}
	if action.State != nil {
		stateMap, err := ResourceIBMSchematicsActionActionStateToMap(action.State)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("state", []map[string]interface{}{stateMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting state: %s", err))
		}
	}
	if action.SysLock != nil {
		sysLockMap, err := ResourceIBMSchematicsActionSystemLockToMap(action.SysLock)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("sys_lock", []map[string]interface{}{sysLockMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting sys_lock: %s", err))
		}
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
	if action.PlaybookNames != nil {
		if err = d.Set("playbook_names", action.PlaybookNames); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting playbook_names: %s", err))
		}
	}

	return nil
}

func ResourceIBMSchematicsActionUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	updateActionOptions := &schematicsv1.UpdateActionOptions{}

	updateActionOptions.SetActionID(d.Id())

	hasChange := false

	if d.HasChange("name") {
		updateActionOptions.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("description") {
		updateActionOptions.SetDescription(d.Get("description").(string))
		hasChange = true
	}
	if d.HasChange("location") {
		updateActionOptions.SetLocation(d.Get("location").(string))
		hasChange = true
	}
	if d.HasChange("resource_group") {
		updateActionOptions.SetResourceGroup(d.Get("resource_group").(string))
		hasChange = true
	}
	if d.HasChange("bastion_connection_type") {
		updateActionOptions.SetBastionConnectionType(d.Get("bastion_connection_type").(string))
		hasChange = true
	}
	if d.HasChange("inventory_connection_type") {
		updateActionOptions.SetInventoryConnectionType(d.Get("inventory_connection_type").(string))
		hasChange = true
	}
	if d.HasChange("tags") {
		// TODO: handle Tags of type TypeList -- not primitive, not model
		hasChange = true
	}
	if d.HasChange("user_state") {
		userState, err := ResourceIBMSchematicsActionMapToUserState(d.Get("user_state.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateActionOptions.SetUserState(userState)
		hasChange = true
	}
	if d.HasChange("source_readme_url") {
		updateActionOptions.SetSourceReadmeURL(d.Get("source_readme_url").(string))
		hasChange = true
	}
	if d.HasChange("source") {
		source, err := ResourceIBMSchematicsActionMapToExternalSource(d.Get("source.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateActionOptions.SetSource(source)
		hasChange = true
	}
	if d.HasChange("source_type") {
		updateActionOptions.SetSourceType(d.Get("source_type").(string))
		hasChange = true
	}
	if d.HasChange("command_parameter") {
		updateActionOptions.SetCommandParameter(d.Get("command_parameter").(string))
		hasChange = true
	}
	if d.HasChange("inventory") {
		updateActionOptions.SetInventory(d.Get("inventory").(string))
		hasChange = true
	}
	if d.HasChange("credentials") {
		// TODO: handle Credentials of type TypeList -- not primitive, not model
		hasChange = true
	}
	if d.HasChange("bastion") {
		bastion, err := ResourceIBMSchematicsActionMapToBastionResourceDefinition(d.Get("bastion.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateActionOptions.SetBastion(bastion)
		hasChange = true
	}
	if d.HasChange("bastion_credential") {
		bastionCredential, err := ResourceIBMSchematicsActionMapToCredentialVariableData(d.Get("bastion_credential.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateActionOptions.SetBastionCredential(bastionCredential)
		hasChange = true
	}
	if d.HasChange("targets_ini") {
		updateActionOptions.SetTargetsIni(d.Get("targets_ini").(string))
		hasChange = true
	}
	if d.HasChange("action_inputs") {
		// TODO: handle Inputs of type TypeList -- not primitive, not model
		hasChange = true
	}
	if d.HasChange("action_outputs") {
		// TODO: handle Outputs of type TypeList -- not primitive, not model
		hasChange = true
	}
	if d.HasChange("settings") {
		// TODO: handle Settings of type TypeList -- not primitive, not model
		hasChange = true
	}
	if d.HasChange("state") {
		state, err := ResourceIBMSchematicsActionMapToActionState(d.Get("state.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateActionOptions.SetState(state)
		hasChange = true
	}
	if d.HasChange("sys_lock") {
		sysLock, err := ResourceIBMSchematicsActionMapToSystemLock(d.Get("sys_lock.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateActionOptions.SetSysLock(sysLock)
		hasChange = true
	}
	if d.HasChange("x_github_token") {
		updateActionOptions.SetXGithubToken(d.Get("x_github_token").(string))
		hasChange = true
	}

	if hasChange {
		_, response, err := schematicsClient.UpdateActionWithContext(context, updateActionOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateActionWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateActionWithContext failed %s\n%s", err, response))
		}
	}

	return ResourceIBMSchematicsActionRead(context, d, meta)
}

func ResourceIBMSchematicsActionDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteActionOptions := &schematicsv1.DeleteActionOptions{}

	deleteActionOptions.SetActionID(d.Id())

	response, err := schematicsClient.DeleteActionWithContext(context, deleteActionOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteActionWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteActionWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func ResourceIBMSchematicsActionMapToUserState(modelMap map[string]interface{}) (*schematicsv1.UserState, error) {
	model := &schematicsv1.UserState{}
	if modelMap["state"] != nil && modelMap["state"].(string) != "" {
		model.State = core.StringPtr(modelMap["state"].(string))
	}
	if modelMap["set_by"] != nil && modelMap["set_by"].(string) != "" {
		model.SetBy = core.StringPtr(modelMap["set_by"].(string))
	}
	if modelMap["set_at"] != nil {

	}
	return model, nil
}

func ResourceIBMSchematicsActionMapToExternalSource(modelMap map[string]interface{}) (*schematicsv1.ExternalSource, error) {
	model := &schematicsv1.ExternalSource{}
	model.SourceType = core.StringPtr(modelMap["source_type"].(string))
	if modelMap["git"] != nil && len(modelMap["git"].([]interface{})) > 0 {
		GitModel, err := ResourceIBMSchematicsActionMapToExternalSourceGit(modelMap["git"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Git = GitModel
	}
	if modelMap["catalog"] != nil && len(modelMap["catalog"].([]interface{})) > 0 {
		CatalogModel, err := ResourceIBMSchematicsActionMapToExternalSourceCatalog(modelMap["catalog"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Catalog = CatalogModel
	}
	return model, nil
}

func ResourceIBMSchematicsActionMapToExternalSourceGit(modelMap map[string]interface{}) (*schematicsv1.ExternalSourceGit, error) {
	model := &schematicsv1.ExternalSourceGit{}
	if modelMap["computed_git_repo_url"] != nil && modelMap["computed_git_repo_url"].(string) != "" {
		model.ComputedGitRepoURL = core.StringPtr(modelMap["computed_git_repo_url"].(string))
	}
	if modelMap["git_repo_url"] != nil && modelMap["git_repo_url"].(string) != "" {
		model.GitRepoURL = core.StringPtr(modelMap["git_repo_url"].(string))
	}
	if modelMap["git_token"] != nil && modelMap["git_token"].(string) != "" {
		model.GitToken = core.StringPtr(modelMap["git_token"].(string))
	}
	if modelMap["git_repo_folder"] != nil && modelMap["git_repo_folder"].(string) != "" {
		model.GitRepoFolder = core.StringPtr(modelMap["git_repo_folder"].(string))
	}
	if modelMap["git_release"] != nil && modelMap["git_release"].(string) != "" {
		model.GitRelease = core.StringPtr(modelMap["git_release"].(string))
	}
	if modelMap["git_branch"] != nil && modelMap["git_branch"].(string) != "" {
		model.GitBranch = core.StringPtr(modelMap["git_branch"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsActionMapToExternalSourceCatalog(modelMap map[string]interface{}) (*schematicsv1.ExternalSourceCatalog, error) {
	model := &schematicsv1.ExternalSourceCatalog{}
	if modelMap["catalog_name"] != nil && modelMap["catalog_name"].(string) != "" {
		model.CatalogName = core.StringPtr(modelMap["catalog_name"].(string))
	}
	if modelMap["offering_name"] != nil && modelMap["offering_name"].(string) != "" {
		model.OfferingName = core.StringPtr(modelMap["offering_name"].(string))
	}
	if modelMap["offering_version"] != nil && modelMap["offering_version"].(string) != "" {
		model.OfferingVersion = core.StringPtr(modelMap["offering_version"].(string))
	}
	if modelMap["offering_kind"] != nil && modelMap["offering_kind"].(string) != "" {
		model.OfferingKind = core.StringPtr(modelMap["offering_kind"].(string))
	}
	if modelMap["offering_id"] != nil && modelMap["offering_id"].(string) != "" {
		model.OfferingID = core.StringPtr(modelMap["offering_id"].(string))
	}
	if modelMap["offering_version_id"] != nil && modelMap["offering_version_id"].(string) != "" {
		model.OfferingVersionID = core.StringPtr(modelMap["offering_version_id"].(string))
	}
	if modelMap["offering_repo_url"] != nil && modelMap["offering_repo_url"].(string) != "" {
		model.OfferingRepoURL = core.StringPtr(modelMap["offering_repo_url"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsActionMapToCredentialVariableData(modelMap map[string]interface{}) (*schematicsv1.VariableData, error) {
	model := &schematicsv1.VariableData{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["use_default"] != nil {
		model.UseDefault = core.BoolPtr(modelMap["use_default"].(bool))
	}
	if modelMap["metadata"] != nil && len(modelMap["metadata"].([]interface{})) > 0 {
		MetadataModel, err := ResourceIBMSchematicsActionMapToCredentialVariableMetadata(modelMap["metadata"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Metadata = MetadataModel
	}
	if modelMap["link"] != nil && modelMap["link"].(string) != "" {
		model.Link = core.StringPtr(modelMap["link"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsActionMapToCredentialVariableMetadata(modelMap map[string]interface{}) (*schematicsv1.VariableMetadata, error) {
	model := &schematicsv1.VariableMetadata{}
	if modelMap["type"] != nil && modelMap["type"].(string) != "" {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["aliases"] != nil {
		aliases := []string{}
		for _, aliasesItem := range modelMap["aliases"].([]interface{}) {
			aliases = append(aliases, aliasesItem.(string))
		}
		model.Aliases = aliases
	}
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["cloud_data_type"] != nil && modelMap["cloud_data_type"].(string) != "" {
		model.CloudDataType = core.StringPtr(modelMap["cloud_data_type"].(string))
	}
	if modelMap["default_value"] != nil && modelMap["default_value"].(string) != "" {
		model.DefaultValue = core.StringPtr(modelMap["default_value"].(string))
	}
	if modelMap["link_status"] != nil && modelMap["link_status"].(string) != "" {
		model.LinkStatus = core.StringPtr(modelMap["link_status"].(string))
	}
	if modelMap["immutable"] != nil {
		model.Immutable = core.BoolPtr(modelMap["immutable"].(bool))
	}
	if modelMap["hidden"] != nil {
		model.Hidden = core.BoolPtr(modelMap["hidden"].(bool))
	}
	if modelMap["required"] != nil {
		model.Required = core.BoolPtr(modelMap["required"].(bool))
	}
	if modelMap["position"] != nil {
		model.Position = core.Int64Ptr(int64(modelMap["position"].(int)))
	}
	if modelMap["group_by"] != nil && modelMap["group_by"].(string) != "" {
		model.GroupBy = core.StringPtr(modelMap["group_by"].(string))
	}
	if modelMap["source"] != nil && modelMap["source"].(string) != "" {
		model.Source = core.StringPtr(modelMap["source"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsActionMapToBastionResourceDefinition(modelMap map[string]interface{}) (*schematicsv1.BastionResourceDefinition, error) {
	model := &schematicsv1.BastionResourceDefinition{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["host"] != nil && modelMap["host"].(string) != "" {
		model.Host = core.StringPtr(modelMap["host"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsActionMapToVariableData(modelMap map[string]interface{}) (*schematicsv1.VariableData, error) {
	model := &schematicsv1.VariableData{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["use_default"] != nil {
		model.UseDefault = core.BoolPtr(modelMap["use_default"].(bool))
	}
	if modelMap["metadata"] != nil && len(modelMap["metadata"].([]interface{})) > 0 {
		MetadataModel, err := ResourceIBMSchematicsActionMapToVariableMetadata(modelMap["metadata"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Metadata = MetadataModel
	}
	if modelMap["link"] != nil && modelMap["link"].(string) != "" {
		model.Link = core.StringPtr(modelMap["link"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsActionMapToVariableMetadata(modelMap map[string]interface{}) (*schematicsv1.VariableMetadata, error) {
	model := &schematicsv1.VariableMetadata{}
	if modelMap["type"] != nil && modelMap["type"].(string) != "" {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["aliases"] != nil {
		aliases := []string{}
		for _, aliasesItem := range modelMap["aliases"].([]interface{}) {
			aliases = append(aliases, aliasesItem.(string))
		}
		model.Aliases = aliases
	}
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["cloud_data_type"] != nil && modelMap["cloud_data_type"].(string) != "" {
		model.CloudDataType = core.StringPtr(modelMap["cloud_data_type"].(string))
	}
	if modelMap["default_value"] != nil && modelMap["default_value"].(string) != "" {
		model.DefaultValue = core.StringPtr(modelMap["default_value"].(string))
	}
	if modelMap["link_status"] != nil && modelMap["link_status"].(string) != "" {
		model.LinkStatus = core.StringPtr(modelMap["link_status"].(string))
	}
	if modelMap["secure"] != nil {
		model.Secure = core.BoolPtr(modelMap["secure"].(bool))
	}
	if modelMap["immutable"] != nil {
		model.Immutable = core.BoolPtr(modelMap["immutable"].(bool))
	}
	if modelMap["hidden"] != nil {
		model.Hidden = core.BoolPtr(modelMap["hidden"].(bool))
	}
	if modelMap["required"] != nil {
		model.Required = core.BoolPtr(modelMap["required"].(bool))
	}
	if modelMap["options"] != nil {
		options := []string{}
		for _, optionsItem := range modelMap["options"].([]interface{}) {
			options = append(options, optionsItem.(string))
		}
		model.Options = options
	}
	if modelMap["min_value"] != nil {
		model.MinValue = core.Int64Ptr(int64(modelMap["min_value"].(int)))
	}
	if modelMap["max_value"] != nil {
		model.MaxValue = core.Int64Ptr(int64(modelMap["max_value"].(int)))
	}
	if modelMap["min_length"] != nil {
		model.MinLength = core.Int64Ptr(int64(modelMap["min_length"].(int)))
	}
	if modelMap["max_length"] != nil {
		model.MaxLength = core.Int64Ptr(int64(modelMap["max_length"].(int)))
	}
	if modelMap["matches"] != nil && modelMap["matches"].(string) != "" {
		model.Matches = core.StringPtr(modelMap["matches"].(string))
	}
	if modelMap["position"] != nil {
		model.Position = core.Int64Ptr(int64(modelMap["position"].(int)))
	}
	if modelMap["group_by"] != nil && modelMap["group_by"].(string) != "" {
		model.GroupBy = core.StringPtr(modelMap["group_by"].(string))
	}
	if modelMap["source"] != nil && modelMap["source"].(string) != "" {
		model.Source = core.StringPtr(modelMap["source"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsActionMapToActionState(modelMap map[string]interface{}) (*schematicsv1.ActionState, error) {
	model := &schematicsv1.ActionState{}
	if modelMap["status_code"] != nil && modelMap["status_code"].(string) != "" {
		model.StatusCode = core.StringPtr(modelMap["status_code"].(string))
	}
	if modelMap["status_job_id"] != nil && modelMap["status_job_id"].(string) != "" {
		model.StatusJobID = core.StringPtr(modelMap["status_job_id"].(string))
	}
	if modelMap["status_message"] != nil && modelMap["status_message"].(string) != "" {
		model.StatusMessage = core.StringPtr(modelMap["status_message"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsActionMapToSystemLock(modelMap map[string]interface{}) (*schematicsv1.SystemLock, error) {
	model := &schematicsv1.SystemLock{}
	if modelMap["sys_locked"] != nil {
		model.SysLocked = core.BoolPtr(modelMap["sys_locked"].(bool))
	}
	if modelMap["sys_locked_by"] != nil && modelMap["sys_locked_by"].(string) != "" {
		model.SysLockedBy = core.StringPtr(modelMap["sys_locked_by"].(string))
	}
	if modelMap["sys_locked_at"] != nil {

	}
	return model, nil
}

func ResourceIBMSchematicsActionUserStateToMap(model *schematicsv1.UserState) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.State != nil {
		modelMap["state"] = model.State
	}
	if model.SetBy != nil {
		modelMap["set_by"] = model.SetBy
	}
	if model.SetAt != nil {
		modelMap["set_at"] = model.SetAt.String()
	}
	return modelMap, nil
}

func ResourceIBMSchematicsActionExternalSourceToMap(model *schematicsv1.ExternalSource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["source_type"] = model.SourceType
	if model.Git != nil {
		gitMap, err := ResourceIBMSchematicsActionExternalSourceGitToMap(model.Git)
		if err != nil {
			return modelMap, err
		}
		modelMap["git"] = []map[string]interface{}{gitMap}
	}
	if model.Catalog != nil {
		catalogMap, err := ResourceIBMSchematicsActionExternalSourceCatalogToMap(model.Catalog)
		if err != nil {
			return modelMap, err
		}
		modelMap["catalog"] = []map[string]interface{}{catalogMap}
	}
	return modelMap, nil
}

func ResourceIBMSchematicsActionExternalSourceGitToMap(model *schematicsv1.ExternalSourceGit) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ComputedGitRepoURL != nil {
		modelMap["computed_git_repo_url"] = model.ComputedGitRepoURL
	}
	if model.GitRepoURL != nil {
		modelMap["git_repo_url"] = model.GitRepoURL
	}
	if model.GitToken != nil {
		modelMap["git_token"] = model.GitToken
	}
	if model.GitRepoFolder != nil {
		modelMap["git_repo_folder"] = model.GitRepoFolder
	}
	if model.GitRelease != nil {
		modelMap["git_release"] = model.GitRelease
	}
	if model.GitBranch != nil {
		modelMap["git_branch"] = model.GitBranch
	}
	return modelMap, nil
}

func ResourceIBMSchematicsActionExternalSourceCatalogToMap(model *schematicsv1.ExternalSourceCatalog) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CatalogName != nil {
		modelMap["catalog_name"] = model.CatalogName
	}
	if model.OfferingName != nil {
		modelMap["offering_name"] = model.OfferingName
	}
	if model.OfferingVersion != nil {
		modelMap["offering_version"] = model.OfferingVersion
	}
	if model.OfferingKind != nil {
		modelMap["offering_kind"] = model.OfferingKind
	}
	if model.OfferingID != nil {
		modelMap["offering_id"] = model.OfferingID
	}
	if model.OfferingVersionID != nil {
		modelMap["offering_version_id"] = model.OfferingVersionID
	}
	if model.OfferingRepoURL != nil {
		modelMap["offering_repo_url"] = model.OfferingRepoURL
	}
	return modelMap, nil
}

func ResourceIBMSchematicsActionCredentialVariableDataToMap(model *schematicsv1.VariableData) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.UseDefault != nil {
		modelMap["use_default"] = model.UseDefault
	}
	if model.Metadata != nil {
		metadataMap, err := ResourceIBMSchematicsActionCredentialVariableMetadataToMap(model.Metadata)
		if err != nil {
			return modelMap, err
		}
		modelMap["metadata"] = []map[string]interface{}{metadataMap}
	}
	if model.Link != nil {
		modelMap["link"] = model.Link
	}
	return modelMap, nil
}

func ResourceIBMSchematicsActionCredentialVariableMetadataToMap(model *schematicsv1.VariableMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Type != nil {
		modelMap["type"] = model.Type
	}
	if model.Aliases != nil {
		modelMap["aliases"] = model.Aliases
	}
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.CloudDataType != nil {
		modelMap["cloud_data_type"] = model.CloudDataType
	}
	if model.DefaultValue != nil {
		modelMap["default_value"] = model.DefaultValue
	}
	if model.LinkStatus != nil {
		modelMap["link_status"] = model.LinkStatus
	}
	if model.Immutable != nil {
		modelMap["immutable"] = model.Immutable
	}
	if model.Hidden != nil {
		modelMap["hidden"] = model.Hidden
	}
	if model.Required != nil {
		modelMap["required"] = model.Required
	}
	if model.Position != nil {
		modelMap["position"] = flex.IntValue(model.Position)
	}
	if model.GroupBy != nil {
		modelMap["group_by"] = model.GroupBy
	}
	if model.Source != nil {
		modelMap["source"] = model.Source
	}
	return modelMap, nil
}

func ResourceIBMSchematicsActionBastionResourceDefinitionToMap(model *schematicsv1.BastionResourceDefinition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Host != nil {
		modelMap["host"] = model.Host
	}
	return modelMap, nil
}

func ResourceIBMSchematicsActionVariableDataToMap(model *schematicsv1.VariableData) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.UseDefault != nil {
		modelMap["use_default"] = model.UseDefault
	}
	if model.Metadata != nil {
		metadataMap, err := ResourceIBMSchematicsActionVariableMetadataToMap(model.Metadata)
		if err != nil {
			return modelMap, err
		}
		modelMap["metadata"] = []map[string]interface{}{metadataMap}
	}
	if model.Link != nil {
		modelMap["link"] = model.Link
	}
	return modelMap, nil
}

func ResourceIBMSchematicsActionVariableMetadataToMap(model *schematicsv1.VariableMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Type != nil {
		modelMap["type"] = model.Type
	}
	if model.Aliases != nil {
		modelMap["aliases"] = model.Aliases
	}
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.CloudDataType != nil {
		modelMap["cloud_data_type"] = model.CloudDataType
	}
	if model.DefaultValue != nil {
		modelMap["default_value"] = model.DefaultValue
	}
	if model.LinkStatus != nil {
		modelMap["link_status"] = model.LinkStatus
	}
	if model.Secure != nil {
		modelMap["secure"] = model.Secure
	}
	if model.Immutable != nil {
		modelMap["immutable"] = model.Immutable
	}
	if model.Hidden != nil {
		modelMap["hidden"] = model.Hidden
	}
	if model.Required != nil {
		modelMap["required"] = model.Required
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	if model.MinValue != nil {
		modelMap["min_value"] = flex.IntValue(model.MinValue)
	}
	if model.MaxValue != nil {
		modelMap["max_value"] = flex.IntValue(model.MaxValue)
	}
	if model.MinLength != nil {
		modelMap["min_length"] = flex.IntValue(model.MinLength)
	}
	if model.MaxLength != nil {
		modelMap["max_length"] = flex.IntValue(model.MaxLength)
	}
	if model.Matches != nil {
		modelMap["matches"] = model.Matches
	}
	if model.Position != nil {
		modelMap["position"] = flex.IntValue(model.Position)
	}
	if model.GroupBy != nil {
		modelMap["group_by"] = model.GroupBy
	}
	if model.Source != nil {
		modelMap["source"] = model.Source
	}
	return modelMap, nil
}

func ResourceIBMSchematicsActionActionStateToMap(model *schematicsv1.ActionState) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.StatusCode != nil {
		modelMap["status_code"] = model.StatusCode
	}
	if model.StatusJobID != nil {
		modelMap["status_job_id"] = model.StatusJobID
	}
	if model.StatusMessage != nil {
		modelMap["status_message"] = model.StatusMessage
	}
	return modelMap, nil
}

func ResourceIBMSchematicsActionSystemLockToMap(model *schematicsv1.SystemLock) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SysLocked != nil {
		modelMap["sys_locked"] = model.SysLocked
	}
	if model.SysLockedBy != nil {
		modelMap["sys_locked_by"] = model.SysLockedBy
	}
	if model.SysLockedAt != nil {
		modelMap["sys_locked_at"] = model.SysLockedAt.String()
	}
	return modelMap, nil
}
