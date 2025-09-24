// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.107.1-41b0fbd0-20250825-080732
 */

package project

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
	"github.com/IBM/project-go-sdk/projectv1"
)

func ResourceIbmProjectConfig() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmProjectConfigCreate,
		ReadContext:   resourceIbmProjectConfigRead,
		UpdateContext: resourceIbmProjectConfigUpdate,
		DeleteContext: resourceIbmProjectConfigDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_project_config", "project_id"),
				Description:  "The unique project ID.",
			},
			"schematics": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "A Schematics workspace that is associated to a project configuration, with scripts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"workspace_crn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "An IBM Cloud resource name that uniquely identifies a resource.",
						},
						"validate_pre_script": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The type of the script.",
									},
									"path": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The path to this script is within the current version source.",
									},
									"short_description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The short description for this script.",
									},
								},
							},
						},
						"validate_post_script": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The type of the script.",
									},
									"path": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The path to this script is within the current version source.",
									},
									"short_description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The short description for this script.",
									},
								},
							},
						},
						"deploy_pre_script": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The type of the script.",
									},
									"path": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The path to this script is within the current version source.",
									},
									"short_description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The short description for this script.",
									},
								},
							},
						},
						"deploy_post_script": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The type of the script.",
									},
									"path": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The path to this script is within the current version source.",
									},
									"short_description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The short description for this script.",
									},
								},
							},
						},
						"undeploy_pre_script": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The type of the script.",
									},
									"path": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The path to this script is within the current version source.",
									},
									"short_description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The short description for this script.",
									},
								},
							},
						},
						"undeploy_post_script": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The type of the script.",
									},
									"path": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The path to this script is within the current version source.",
									},
									"short_description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The short description for this script.",
									},
								},
							},
						},
					},
				},
			},
			"definition": &schema.Schema{
				Type:     schema.TypeList,
				MinItems: 1,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compliance_profile": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The profile that is required for compliance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The unique ID for the compliance profile.",
									},
									"instance_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A unique ID for the instance of a compliance profile.",
									},
									"instance_location": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The location of the compliance instance.",
									},
									"attachment_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A unique ID for the attachment to a compliance profile.",
									},
									"profile_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of the compliance profile.",
									},
									"wp_policy_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The unique ID for the Workload Protection policy.",
									},
									"wp_instance_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A unique ID for the instance of a Workload Protection.",
									},
									"wp_instance_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of the Workload Protection instance.",
									},
									"wp_instance_location": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The location of the compliance instance.",
									},
									"wp_zone_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A unique ID for the zone to a Workload Protection policy.",
									},
									"wp_zone_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A unique ID for the zone to a Workload Protection policy.",
									},
									"wp_policy_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of the Workload Protection policy.",
									},
								},
							},
						},
						"locator_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "A unique concatenation of the catalog ID and the version ID that identify the deployable architecture in the catalog. If importing from an existing Schematics workspace that is not backed by cart, a `locator_id` is required. If using a Schematics workspace that is backed by cart, a `locator_id` is not necessary because the Schematics workspace has one.> There are 3 scenarios:> 1. If only a `locator_id` is specified, a new Schematics workspace is instantiated with that `locator_id`.> 2. If only a schematics `workspace_crn` is specified, a `400` is returned if a `locator_id` is not found in the existing schematics workspace.> 3. If both a Schematics `workspace_crn` and a `locator_id` are specified, a `400` message is returned if the specified `locator_id` does not agree with the `locator_id` in the existing Schematics workspace.> For more information of creating a Schematics workspace, see [Creating workspaces and importing the Terraform template](/docs/schematics?topic=schematics-sch-create-wks).",
						},
						"members": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The member deployabe architectures that are included in the stack.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name matching the alias in the stack definition.",
									},
									"config_id": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The unique ID.",
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "A project configuration description.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The configuration name. It's unique within the account across projects and regions.",
						},
						"authorizations": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The authorization details. It can authorize by using a trusted profile or an API key in Secrets Manager.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trusted_profile_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The trusted profile ID.",
									},
									"method": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The authorization method. It can authorize by using a trusted profile or an API key in Secrets Manager.",
									},
									"api_key": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Sensitive:   true,
										Description: "The IBM Cloud API Key. It can be either raw or pulled from the catalog via a `CRN` or `JSON` blob.",
									},
								},
							},
						},
						"inputs": &schema.Schema{
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "The input variables that are used for configuration definition and environment.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"settings": &schema.Schema{
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "The Schematics environment variables to use to deploy the configuration. Settings are only available if they are specified when the configuration is initially created.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"environment_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ID of the project environment.",
						},
						"resource_crns": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The CRNs of the resources that are associated with this configuration.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"version": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The version of the configuration.",
			},
			"needs_attention_state": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The needs attention state of a configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"event_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the event.",
						},
						"event": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the event.",
						},
						"severity": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The severity of the event. This is a system generated field. For user triggered events the field is not present.",
						},
						"action_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "An actionable Url that users can access in response to the event. This is a system generated field. For user triggered events the field is not present.",
						},
						"target": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The configuration id and version for which the event occurred. This field is only available for user generated events. For system triggered events the field is not present.",
						},
						"triggered_by": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The IAM id of the user that triggered the event. This field is only available for user generated events. For system triggered events the field is not present.",
						},
						"timestamp": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time at which the event was triggered.",
						},
					},
				},
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339.",
			},
			"modified_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339.",
			},
			"outputs": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The outputs of a Schematics template property.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The variable name.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "A short explanation of the output value.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "This property can be any value - a string, number, boolean, array, or object.",
						},
					},
				},
			},
			"references": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The resolved references that are used by the configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the configuration.",
			},
			"state_code": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Computed state code clarifying the prerequisites for validation for the configuration.",
			},
			"config_error": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: "The error from config actions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The error message from config actions.",
						},
						"details": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "The error details from config actions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
					},
				},
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A Url.",
			},
			"is_draft": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The flag that indicates whether the version of the configuration is draft, or active.",
			},
			"last_saved_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339.",
			},
			"project": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The project that is referenced by this resource.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique ID.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A Url.",
						},
						"definition": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The definition of the project reference.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the project.",
									},
								},
							},
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An IBM Cloud resource name that uniquely identifies a resource.",
						},
					},
				},
			},
			"update_available": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The flag that indicates whether a configuration update is available.",
			},
			"template_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The stack definition identifier.",
			},
			"member_of": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: "The stack config parent of which this configuration is a member of.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique ID.",
						},
						"definition": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The definition summary of the stack configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The configuration name. It's unique within the account across projects and regions.",
									},
									"members": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The member deployable architectures that are included in the stack.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name matching the alias in the stack definition.",
												},
												"config_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique ID.",
												},
											},
										},
									},
								},
							},
						},
						"version": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The version of the stack configuration.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A Url.",
						},
					},
				},
			},
			"deployment_model": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The configuration type.",
			},
			"approved_version": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: "A summary of a project configuration version.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"definition": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A summary of the definition in a project configuration version.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"environment_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The ID of the project environment.",
									},
									"locator_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
										Description: "A unique concatenation of the catalog ID and the version ID that identify the deployable architecture in the catalog. If importing from an existing Schematics workspace that is not backed by cart, a `locator_id` is required. If using a Schematics workspace that is backed by cart, a `locator_id` is not necessary because the Schematics workspace has one.> There are 3 scenarios:> 1. If only a `locator_id` is specified, a new Schematics workspace is instantiated with that `locator_id`.> 2. If only a schematics `workspace_crn` is specified, a `400` is returned if a `locator_id` is not found in the existing schematics workspace.> 3. If both a Schematics `workspace_crn` and a `locator_id` are specified, a `400` message is returned if the specified `locator_id` does not agree with the `locator_id` in the existing Schematics workspace.> For more information of creating a Schematics workspace, see [Creating workspaces and importing the Terraform template](/docs/schematics?topic=schematics-sch-create-wks).",
									},
								},
							},
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The state of the configuration.",
						},
						"version": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The version number of the configuration.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A Url.",
						},
					},
				},
			},
			"deployed_version": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: "A summary of a project configuration version.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"definition": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A summary of the definition in a project configuration version.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"environment_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The ID of the project environment.",
									},
									"locator_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
										Description: "A unique concatenation of the catalog ID and the version ID that identify the deployable architecture in the catalog. If importing from an existing Schematics workspace that is not backed by cart, a `locator_id` is required. If using a Schematics workspace that is backed by cart, a `locator_id` is not necessary because the Schematics workspace has one.> There are 3 scenarios:> 1. If only a `locator_id` is specified, a new Schematics workspace is instantiated with that `locator_id`.> 2. If only a schematics `workspace_crn` is specified, a `400` is returned if a `locator_id` is not found in the existing schematics workspace.> 3. If both a Schematics `workspace_crn` and a `locator_id` are specified, a `400` message is returned if the specified `locator_id` does not agree with the `locator_id` in the existing Schematics workspace.> For more information of creating a Schematics workspace, see [Creating workspaces and importing the Terraform template](/docs/schematics?topic=schematics-sch-create-wks).",
									},
								},
							},
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The state of the configuration.",
						},
						"version": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The version number of the configuration.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A Url.",
						},
					},
				},
			},
			"project_config_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.",
			},
		},
	}
}

func ResourceIbmProjectConfigValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "project_id",
			ValidateFunctionIdentifier: validate.ValidateRegexp,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[\.\-0-9a-zA-Z]+$`,
			MaxValueLength:             128,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_project_config", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmProjectConfigCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createConfigOptions := &projectv1.CreateConfigOptions{}

	createConfigOptions.SetProjectID(d.Get("project_id").(string))
	definitionModel, err := ResourceIbmProjectConfigMapToProjectConfigDefinitionPrototype(d.Get("definition.0").(map[string]interface{}))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "create", "parse-definition").GetDiag()
	}
	createConfigOptions.SetDefinition(definitionModel)
	if _, ok := d.GetOk("schematics"); ok {
		schematicsModel, err := ResourceIbmProjectConfigMapToSchematicsWorkspace(d.Get("schematics.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "create", "parse-schematics").GetDiag()
		}
		createConfigOptions.SetSchematics(schematicsModel)
	}

	projectConfig, _, err := projectClient.CreateConfigWithContext(context, createConfigOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateConfigWithContext failed: %s", err.Error()), "ibm_project_config", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createConfigOptions.ProjectID, *projectConfig.ID))

	return resourceIbmProjectConfigRead(context, d, meta)
}

func resourceIbmProjectConfigRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getConfigOptions := &projectv1.GetConfigOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "sep-id-parts").GetDiag()
	}

	getConfigOptions.SetProjectID(parts[0])
	getConfigOptions.SetID(parts[1])

	projectConfig, response, err := projectClient.GetConfigWithContext(context, getConfigOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetConfigWithContext failed: %s", err.Error()), "ibm_project_config", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	definitionMap, err := ResourceIbmProjectConfigProjectConfigDefinitionResponseToMap(projectConfig.Definition)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "definition-to-map").GetDiag()
	}
	if err = d.Set("definition", []map[string]interface{}{definitionMap}); err != nil {
		err = fmt.Errorf("Error setting definition: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-definition").GetDiag()
	}
	if err = d.Set("version", flex.IntValue(projectConfig.Version)); err != nil {
		err = fmt.Errorf("Error setting version: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-version").GetDiag()
	}
	needsAttentionState := []map[string]interface{}{}
	for _, needsAttentionStateItem := range projectConfig.NeedsAttentionState {
		needsAttentionStateItemMap, err := ResourceIbmProjectConfigProjectConfigNeedsAttentionStateToMap(&needsAttentionStateItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "needs_attention_state-to-map").GetDiag()
		}
		needsAttentionState = append(needsAttentionState, needsAttentionStateItemMap)
	}
	if err = d.Set("needs_attention_state", needsAttentionState); err != nil {
		err = fmt.Errorf("Error setting needs_attention_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-needs_attention_state").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(projectConfig.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("modified_at", flex.DateTimeToString(projectConfig.ModifiedAt)); err != nil {
		err = fmt.Errorf("Error setting modified_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-modified_at").GetDiag()
	}
	outputs := []map[string]interface{}{}
	for _, outputsItem := range projectConfig.Outputs {
		outputsItemMap, err := ResourceIbmProjectConfigOutputValueToMap(&outputsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "outputs-to-map").GetDiag()
		}
		outputs = append(outputs, outputsItemMap)
	}
	if err = d.Set("outputs", outputs); err != nil {
		err = fmt.Errorf("Error setting outputs: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-outputs").GetDiag()
	}
	referencesMap, err := ResourceIbmProjectConfigReferenceValueToMap(projectConfig.References)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "references-to-map").GetDiag()
	}
	if err = d.Set("references", []map[string]interface{}{referencesMap}); err != nil {
		err = fmt.Errorf("Error setting references: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-references").GetDiag()
	}
	if err = d.Set("state", projectConfig.State); err != nil {
		err = fmt.Errorf("Error setting state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-state").GetDiag()
	}
	if !core.IsNil(projectConfig.StateCode) {
		if err = d.Set("state_code", projectConfig.StateCode); err != nil {
			err = fmt.Errorf("Error setting state_code: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-state_code").GetDiag()
		}
	}
	if !core.IsNil(projectConfig.ConfigError) {
		configErrorMap, err := ResourceIbmProjectConfigProjectConfigErrorToMap(projectConfig.ConfigError)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "config_error-to-map").GetDiag()
		}
		if err = d.Set("config_error", []map[string]interface{}{configErrorMap}); err != nil {
			err = fmt.Errorf("Error setting config_error: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-config_error").GetDiag()
		}
	}
	if err = d.Set("href", projectConfig.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-href").GetDiag()
	}
	if err = d.Set("is_draft", projectConfig.IsDraft); err != nil {
		err = fmt.Errorf("Error setting is_draft: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-is_draft").GetDiag()
	}
	if !core.IsNil(projectConfig.LastSavedAt) {
		if err = d.Set("last_saved_at", flex.DateTimeToString(projectConfig.LastSavedAt)); err != nil {
			err = fmt.Errorf("Error setting last_saved_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-last_saved_at").GetDiag()
		}
	}
	projectMap, err := ResourceIbmProjectConfigProjectReferenceToMap(projectConfig.Project)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "project-to-map").GetDiag()
	}
	if err = d.Set("project", []map[string]interface{}{projectMap}); err != nil {
		err = fmt.Errorf("Error setting project: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-project").GetDiag()
	}
	if !core.IsNil(projectConfig.UpdateAvailable) {
		if err = d.Set("update_available", projectConfig.UpdateAvailable); err != nil {
			err = fmt.Errorf("Error setting update_available: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-update_available").GetDiag()
		}
	}
	if !core.IsNil(projectConfig.TemplateID) {
		if err = d.Set("template_id", projectConfig.TemplateID); err != nil {
			err = fmt.Errorf("Error setting template_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-template_id").GetDiag()
		}
	}
	if !core.IsNil(projectConfig.MemberOf) {
		memberOfMap, err := ResourceIbmProjectConfigMemberOfDefinitionToMap(projectConfig.MemberOf)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "member_of-to-map").GetDiag()
		}
		if err = d.Set("member_of", []map[string]interface{}{memberOfMap}); err != nil {
			err = fmt.Errorf("Error setting member_of: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-member_of").GetDiag()
		}
	}
	if err = d.Set("deployment_model", projectConfig.DeploymentModel); err != nil {
		err = fmt.Errorf("Error setting deployment_model: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-deployment_model").GetDiag()
	}
	if !core.IsNil(projectConfig.ApprovedVersion) {
		approvedVersionMap, err := ResourceIbmProjectConfigProjectConfigVersionSummaryToMap(projectConfig.ApprovedVersion)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "approved_version-to-map").GetDiag()
		}
		if err = d.Set("approved_version", []map[string]interface{}{approvedVersionMap}); err != nil {
			err = fmt.Errorf("Error setting approved_version: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-approved_version").GetDiag()
		}
	}
	if !core.IsNil(projectConfig.DeployedVersion) {
		deployedVersionMap, err := ResourceIbmProjectConfigProjectConfigVersionSummaryToMap(projectConfig.DeployedVersion)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "deployed_version-to-map").GetDiag()
		}
		if err = d.Set("deployed_version", []map[string]interface{}{deployedVersionMap}); err != nil {
			err = fmt.Errorf("Error setting deployed_version: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-deployed_version").GetDiag()
		}
	}
	if err = d.Set("project_config_id", projectConfig.ID); err != nil {
		err = fmt.Errorf("Error setting project_config_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "read", "set-project_config_id").GetDiag()
	}

	return nil
}

func resourceIbmProjectConfigUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateConfigOptions := &projectv1.UpdateConfigOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "update", "sep-id-parts").GetDiag()
	}

	updateConfigOptions.SetProjectID(parts[0])
	updateConfigOptions.SetID(parts[1])

	hasChange := false

	if d.HasChange("project_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "project_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_project_config", "update", "project_id-forces-new").GetDiag()
	}
	if d.HasChange("definition") {
		definition, err := ResourceIbmProjectConfigMapToProjectConfigDefinitionPatch(d.Get("definition.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "update", "parse-definition").GetDiag()
		}
		updateConfigOptions.SetDefinition(definition)
		hasChange = true
	}

	if hasChange {
		_, _, err = projectClient.UpdateConfigWithContext(context, updateConfigOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateConfigWithContext failed: %s", err.Error()), "ibm_project_config", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIbmProjectConfigRead(context, d, meta)
}

func resourceIbmProjectConfigDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteConfigOptions := &projectv1.DeleteConfigOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_project_config", "delete", "sep-id-parts").GetDiag()
	}

	deleteConfigOptions.SetProjectID(parts[0])
	deleteConfigOptions.SetID(parts[1])

	_, _, err = projectClient.DeleteConfigWithContext(context, deleteConfigOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteConfigWithContext failed: %s", err.Error()), "ibm_project_config", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIbmProjectConfigMapToProjectConfigDefinitionPrototype(modelMap map[string]interface{}) (projectv1.ProjectConfigDefinitionPrototypeIntf, error) {
	model := &projectv1.ProjectConfigDefinitionPrototype{}
	if modelMap["compliance_profile"] != nil && len(modelMap["compliance_profile"].([]interface{})) > 0 {
		ComplianceProfileModel, err := ResourceIbmProjectConfigMapToProjectComplianceProfile(modelMap["compliance_profile"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ComplianceProfile = ComplianceProfileModel
	}
	if modelMap["locator_id"] != nil && modelMap["locator_id"].(string) != "" {
		model.LocatorID = core.StringPtr(modelMap["locator_id"].(string))
	}
	if modelMap["members"] != nil {
		members := []projectv1.StackMember{}
		for _, membersItem := range modelMap["members"].([]interface{}) {
			membersItemModel, err := ResourceIbmProjectConfigMapToStackMember(membersItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			members = append(members, *membersItemModel)
		}
		model.Members = members
	}
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["authorizations"] != nil && len(modelMap["authorizations"].([]interface{})) > 0 {
		AuthorizationsModel, err := ResourceIbmProjectConfigMapToProjectConfigAuth(modelMap["authorizations"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Authorizations = AuthorizationsModel
	}
	if modelMap["inputs"] != nil {
		model.Inputs = modelMap["inputs"].(map[string]interface{})
	}
	if modelMap["settings"] != nil {
		model.Settings = modelMap["settings"].(map[string]interface{})
	}
	if modelMap["environment_id"] != nil && modelMap["environment_id"].(string) != "" {
		model.EnvironmentID = core.StringPtr(modelMap["environment_id"].(string))
	}
	if modelMap["resource_crns"] != nil {
		resourceCrns := []string{}
		for _, resourceCrnsItem := range modelMap["resource_crns"].([]interface{}) {
			resourceCrns = append(resourceCrns, resourceCrnsItem.(string))
		}
		model.ResourceCrns = resourceCrns
	}
	return model, nil
}

func ResourceIbmProjectConfigMapToProjectComplianceProfile(modelMap map[string]interface{}) (projectv1.ProjectComplianceProfileIntf, error) {
	model := &projectv1.ProjectComplianceProfile{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["instance_id"] != nil && modelMap["instance_id"].(string) != "" {
		model.InstanceID = core.StringPtr(modelMap["instance_id"].(string))
	}
	if modelMap["instance_location"] != nil && modelMap["instance_location"].(string) != "" {
		model.InstanceLocation = core.StringPtr(modelMap["instance_location"].(string))
	}
	if modelMap["attachment_id"] != nil && modelMap["attachment_id"].(string) != "" {
		model.AttachmentID = core.StringPtr(modelMap["attachment_id"].(string))
	}
	if modelMap["profile_name"] != nil && modelMap["profile_name"].(string) != "" {
		model.ProfileName = core.StringPtr(modelMap["profile_name"].(string))
	}
	if modelMap["wp_policy_id"] != nil && modelMap["wp_policy_id"].(string) != "" {
		model.WpPolicyID = core.StringPtr(modelMap["wp_policy_id"].(string))
	}
	if modelMap["wp_instance_id"] != nil && modelMap["wp_instance_id"].(string) != "" {
		model.WpInstanceID = core.StringPtr(modelMap["wp_instance_id"].(string))
	}
	if modelMap["wp_instance_name"] != nil && modelMap["wp_instance_name"].(string) != "" {
		model.WpInstanceName = core.StringPtr(modelMap["wp_instance_name"].(string))
	}
	if modelMap["wp_instance_location"] != nil && modelMap["wp_instance_location"].(string) != "" {
		model.WpInstanceLocation = core.StringPtr(modelMap["wp_instance_location"].(string))
	}
	if modelMap["wp_zone_id"] != nil && modelMap["wp_zone_id"].(string) != "" {
		model.WpZoneID = core.StringPtr(modelMap["wp_zone_id"].(string))
	}
	if modelMap["wp_zone_name"] != nil && modelMap["wp_zone_name"].(string) != "" {
		model.WpZoneName = core.StringPtr(modelMap["wp_zone_name"].(string))
	}
	if modelMap["wp_policy_name"] != nil && modelMap["wp_policy_name"].(string) != "" {
		model.WpPolicyName = core.StringPtr(modelMap["wp_policy_name"].(string))
	}
	return model, nil
}

func ResourceIbmProjectConfigMapToProjectComplianceProfileNullableObject(modelMap map[string]interface{}) (*projectv1.ProjectComplianceProfileNullableObject, error) {
	model := &projectv1.ProjectComplianceProfileNullableObject{}
	return model, nil
}

func ResourceIbmProjectConfigMapToProjectComplianceProfileV1(modelMap map[string]interface{}) (*projectv1.ProjectComplianceProfileV1, error) {
	model := &projectv1.ProjectComplianceProfileV1{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["instance_id"] != nil && modelMap["instance_id"].(string) != "" {
		model.InstanceID = core.StringPtr(modelMap["instance_id"].(string))
	}
	if modelMap["instance_location"] != nil && modelMap["instance_location"].(string) != "" {
		model.InstanceLocation = core.StringPtr(modelMap["instance_location"].(string))
	}
	if modelMap["attachment_id"] != nil && modelMap["attachment_id"].(string) != "" {
		model.AttachmentID = core.StringPtr(modelMap["attachment_id"].(string))
	}
	if modelMap["profile_name"] != nil && modelMap["profile_name"].(string) != "" {
		model.ProfileName = core.StringPtr(modelMap["profile_name"].(string))
	}
	if modelMap["wp_policy_id"] != nil && modelMap["wp_policy_id"].(string) != "" {
		model.WpPolicyID = core.StringPtr(modelMap["wp_policy_id"].(string))
	}
	if modelMap["wp_instance_id"] != nil && modelMap["wp_instance_id"].(string) != "" {
		model.WpInstanceID = core.StringPtr(modelMap["wp_instance_id"].(string))
	}
	if modelMap["wp_instance_name"] != nil && modelMap["wp_instance_name"].(string) != "" {
		model.WpInstanceName = core.StringPtr(modelMap["wp_instance_name"].(string))
	}
	if modelMap["wp_instance_location"] != nil && modelMap["wp_instance_location"].(string) != "" {
		model.WpInstanceLocation = core.StringPtr(modelMap["wp_instance_location"].(string))
	}
	if modelMap["wp_zone_id"] != nil && modelMap["wp_zone_id"].(string) != "" {
		model.WpZoneID = core.StringPtr(modelMap["wp_zone_id"].(string))
	}
	if modelMap["wp_zone_name"] != nil && modelMap["wp_zone_name"].(string) != "" {
		model.WpZoneName = core.StringPtr(modelMap["wp_zone_name"].(string))
	}
	if modelMap["wp_policy_name"] != nil && modelMap["wp_policy_name"].(string) != "" {
		model.WpPolicyName = core.StringPtr(modelMap["wp_policy_name"].(string))
	}
	return model, nil
}

func ResourceIbmProjectConfigMapToStackMember(modelMap map[string]interface{}) (*projectv1.StackMember, error) {
	model := &projectv1.StackMember{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.ConfigID = core.StringPtr(modelMap["config_id"].(string))
	return model, nil
}

func ResourceIbmProjectConfigMapToProjectConfigAuth(modelMap map[string]interface{}) (*projectv1.ProjectConfigAuth, error) {
	model := &projectv1.ProjectConfigAuth{}
	if modelMap["trusted_profile_id"] != nil && modelMap["trusted_profile_id"].(string) != "" {
		model.TrustedProfileID = core.StringPtr(modelMap["trusted_profile_id"].(string))
	}
	if modelMap["method"] != nil && modelMap["method"].(string) != "" {
		model.Method = core.StringPtr(modelMap["method"].(string))
	}
	if modelMap["api_key"] != nil && modelMap["api_key"].(string) != "" {
		model.ApiKey = core.StringPtr(modelMap["api_key"].(string))
	}
	return model, nil
}

func ResourceIbmProjectConfigMapToProjectConfigDefinitionPrototypeDAConfigDefinitionPropertiesPrototype(modelMap map[string]interface{}) (*projectv1.ProjectConfigDefinitionPrototypeDAConfigDefinitionPropertiesPrototype, error) {
	model := &projectv1.ProjectConfigDefinitionPrototypeDAConfigDefinitionPropertiesPrototype{}
	if modelMap["compliance_profile"] != nil && len(modelMap["compliance_profile"].([]interface{})) > 0 {
		ComplianceProfileModel, err := ResourceIbmProjectConfigMapToProjectComplianceProfile(modelMap["compliance_profile"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ComplianceProfile = ComplianceProfileModel
	}
	if modelMap["locator_id"] != nil && modelMap["locator_id"].(string) != "" {
		model.LocatorID = core.StringPtr(modelMap["locator_id"].(string))
	}
	if modelMap["members"] != nil {
		members := []projectv1.StackMember{}
		for _, membersItem := range modelMap["members"].([]interface{}) {
			membersItemModel, err := ResourceIbmProjectConfigMapToStackMember(membersItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			members = append(members, *membersItemModel)
		}
		model.Members = members
	}
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["authorizations"] != nil && len(modelMap["authorizations"].([]interface{})) > 0 {
		AuthorizationsModel, err := ResourceIbmProjectConfigMapToProjectConfigAuth(modelMap["authorizations"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Authorizations = AuthorizationsModel
	}
	if modelMap["inputs"] != nil {
		model.Inputs = modelMap["inputs"].(map[string]interface{})
	}
	if modelMap["settings"] != nil {
		model.Settings = modelMap["settings"].(map[string]interface{})
	}
	if modelMap["environment_id"] != nil && modelMap["environment_id"].(string) != "" {
		model.EnvironmentID = core.StringPtr(modelMap["environment_id"].(string))
	}
	return model, nil
}

func ResourceIbmProjectConfigMapToProjectConfigDefinitionPrototypeResourceConfigDefinitionPropertiesPrototype(modelMap map[string]interface{}) (*projectv1.ProjectConfigDefinitionPrototypeResourceConfigDefinitionPropertiesPrototype, error) {
	model := &projectv1.ProjectConfigDefinitionPrototypeResourceConfigDefinitionPropertiesPrototype{}
	if modelMap["resource_crns"] != nil {
		resourceCrns := []string{}
		for _, resourceCrnsItem := range modelMap["resource_crns"].([]interface{}) {
			resourceCrns = append(resourceCrns, resourceCrnsItem.(string))
		}
		model.ResourceCrns = resourceCrns
	}
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["authorizations"] != nil && len(modelMap["authorizations"].([]interface{})) > 0 {
		AuthorizationsModel, err := ResourceIbmProjectConfigMapToProjectConfigAuth(modelMap["authorizations"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Authorizations = AuthorizationsModel
	}
	if modelMap["inputs"] != nil {
		model.Inputs = modelMap["inputs"].(map[string]interface{})
	}
	if modelMap["settings"] != nil {
		model.Settings = modelMap["settings"].(map[string]interface{})
	}
	if modelMap["environment_id"] != nil && modelMap["environment_id"].(string) != "" {
		model.EnvironmentID = core.StringPtr(modelMap["environment_id"].(string))
	}
	return model, nil
}

func ResourceIbmProjectConfigMapToSchematicsWorkspace(modelMap map[string]interface{}) (*projectv1.SchematicsWorkspace, error) {
	model := &projectv1.SchematicsWorkspace{}
	if modelMap["workspace_crn"] != nil && modelMap["workspace_crn"].(string) != "" {
		model.WorkspaceCrn = core.StringPtr(modelMap["workspace_crn"].(string))
	}
	return model, nil
}

func ResourceIbmProjectConfigMapToProjectConfigDefinitionPatch(modelMap map[string]interface{}) (projectv1.ProjectConfigDefinitionPatchIntf, error) {
	model := &projectv1.ProjectConfigDefinitionPatch{}
	if modelMap["compliance_profile"] != nil && len(modelMap["compliance_profile"].([]interface{})) > 0 {
		ComplianceProfileModel, err := ResourceIbmProjectConfigMapToProjectComplianceProfile(modelMap["compliance_profile"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ComplianceProfile = ComplianceProfileModel
	}
	if modelMap["locator_id"] != nil && modelMap["locator_id"].(string) != "" {
		model.LocatorID = core.StringPtr(modelMap["locator_id"].(string))
	}
	if modelMap["members"] != nil {
		members := []projectv1.StackMember{}
		for _, membersItem := range modelMap["members"].([]interface{}) {
			membersItemModel, err := ResourceIbmProjectConfigMapToStackMember(membersItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			members = append(members, *membersItemModel)
		}
		model.Members = members
	}
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["authorizations"] != nil && len(modelMap["authorizations"].([]interface{})) > 0 {
		AuthorizationsModel, err := ResourceIbmProjectConfigMapToProjectConfigAuth(modelMap["authorizations"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Authorizations = AuthorizationsModel
	}
	if modelMap["inputs"] != nil {
		model.Inputs = modelMap["inputs"].(map[string]interface{})
	}
	if modelMap["settings"] != nil {
		model.Settings = modelMap["settings"].(map[string]interface{})
	}
	if modelMap["environment_id"] != nil && modelMap["environment_id"].(string) != "" {
		model.EnvironmentID = core.StringPtr(modelMap["environment_id"].(string))
	}
	if modelMap["resource_crns"] != nil {
		resourceCrns := []string{}
		for _, resourceCrnsItem := range modelMap["resource_crns"].([]interface{}) {
			resourceCrns = append(resourceCrns, resourceCrnsItem.(string))
		}
		model.ResourceCrns = resourceCrns
	}
	return model, nil
}

func ResourceIbmProjectConfigMapToProjectConfigDefinitionPatchDAConfigDefinitionPropertiesPatch(modelMap map[string]interface{}) (*projectv1.ProjectConfigDefinitionPatchDAConfigDefinitionPropertiesPatch, error) {
	model := &projectv1.ProjectConfigDefinitionPatchDAConfigDefinitionPropertiesPatch{}
	if modelMap["compliance_profile"] != nil && len(modelMap["compliance_profile"].([]interface{})) > 0 {
		ComplianceProfileModel, err := ResourceIbmProjectConfigMapToProjectComplianceProfile(modelMap["compliance_profile"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ComplianceProfile = ComplianceProfileModel
	}
	if modelMap["locator_id"] != nil && modelMap["locator_id"].(string) != "" {
		model.LocatorID = core.StringPtr(modelMap["locator_id"].(string))
	}
	if modelMap["members"] != nil {
		members := []projectv1.StackMember{}
		for _, membersItem := range modelMap["members"].([]interface{}) {
			membersItemModel, err := ResourceIbmProjectConfigMapToStackMember(membersItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			members = append(members, *membersItemModel)
		}
		model.Members = members
	}
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["authorizations"] != nil && len(modelMap["authorizations"].([]interface{})) > 0 {
		AuthorizationsModel, err := ResourceIbmProjectConfigMapToProjectConfigAuth(modelMap["authorizations"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Authorizations = AuthorizationsModel
	}
	if modelMap["inputs"] != nil {
		model.Inputs = modelMap["inputs"].(map[string]interface{})
	}
	if modelMap["settings"] != nil {
		model.Settings = modelMap["settings"].(map[string]interface{})
	}
	if modelMap["environment_id"] != nil && modelMap["environment_id"].(string) != "" {
		model.EnvironmentID = core.StringPtr(modelMap["environment_id"].(string))
	}
	return model, nil
}

func ResourceIbmProjectConfigMapToProjectConfigDefinitionPatchResourceConfigDefinitionPropertiesPatch(modelMap map[string]interface{}) (*projectv1.ProjectConfigDefinitionPatchResourceConfigDefinitionPropertiesPatch, error) {
	model := &projectv1.ProjectConfigDefinitionPatchResourceConfigDefinitionPropertiesPatch{}
	if modelMap["resource_crns"] != nil {
		resourceCrns := []string{}
		for _, resourceCrnsItem := range modelMap["resource_crns"].([]interface{}) {
			resourceCrns = append(resourceCrns, resourceCrnsItem.(string))
		}
		model.ResourceCrns = resourceCrns
	}
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["authorizations"] != nil && len(modelMap["authorizations"].([]interface{})) > 0 {
		AuthorizationsModel, err := ResourceIbmProjectConfigMapToProjectConfigAuth(modelMap["authorizations"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Authorizations = AuthorizationsModel
	}
	if modelMap["inputs"] != nil {
		model.Inputs = modelMap["inputs"].(map[string]interface{})
	}
	if modelMap["settings"] != nil {
		model.Settings = modelMap["settings"].(map[string]interface{})
	}
	if modelMap["environment_id"] != nil && modelMap["environment_id"].(string) != "" {
		model.EnvironmentID = core.StringPtr(modelMap["environment_id"].(string))
	}
	return model, nil
}

func ResourceIbmProjectConfigSchematicsMetadataToMap(model *projectv1.SchematicsMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.WorkspaceCrn != nil {
		modelMap["workspace_crn"] = *model.WorkspaceCrn
	}
	if model.ValidatePreScript != nil {
		validatePreScriptMap, err := ResourceIbmProjectConfigScriptToMap(model.ValidatePreScript)
		if err != nil {
			return modelMap, err
		}
		modelMap["validate_pre_script"] = []map[string]interface{}{validatePreScriptMap}
	}
	if model.ValidatePostScript != nil {
		validatePostScriptMap, err := ResourceIbmProjectConfigScriptToMap(model.ValidatePostScript)
		if err != nil {
			return modelMap, err
		}
		modelMap["validate_post_script"] = []map[string]interface{}{validatePostScriptMap}
	}
	if model.DeployPreScript != nil {
		deployPreScriptMap, err := ResourceIbmProjectConfigScriptToMap(model.DeployPreScript)
		if err != nil {
			return modelMap, err
		}
		modelMap["deploy_pre_script"] = []map[string]interface{}{deployPreScriptMap}
	}
	if model.DeployPostScript != nil {
		deployPostScriptMap, err := ResourceIbmProjectConfigScriptToMap(model.DeployPostScript)
		if err != nil {
			return modelMap, err
		}
		modelMap["deploy_post_script"] = []map[string]interface{}{deployPostScriptMap}
	}
	if model.UndeployPreScript != nil {
		undeployPreScriptMap, err := ResourceIbmProjectConfigScriptToMap(model.UndeployPreScript)
		if err != nil {
			return modelMap, err
		}
		modelMap["undeploy_pre_script"] = []map[string]interface{}{undeployPreScriptMap}
	}
	if model.UndeployPostScript != nil {
		undeployPostScriptMap, err := ResourceIbmProjectConfigScriptToMap(model.UndeployPostScript)
		if err != nil {
			return modelMap, err
		}
		modelMap["undeploy_post_script"] = []map[string]interface{}{undeployPostScriptMap}
	}
	return modelMap, nil
}

func ResourceIbmProjectConfigScriptToMap(model *projectv1.Script) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Path != nil {
		modelMap["path"] = *model.Path
	}
	if model.ShortDescription != nil {
		modelMap["short_description"] = *model.ShortDescription
	}
	return modelMap, nil
}

func ResourceIbmProjectConfigProjectConfigDefinitionResponseToMap(model projectv1.ProjectConfigDefinitionResponseIntf) (map[string]interface{}, error) {
	if _, ok := model.(*projectv1.ProjectConfigDefinitionResponseDAConfigDefinitionPropertiesResponse); ok {
		return ResourceIbmProjectConfigProjectConfigDefinitionResponseDAConfigDefinitionPropertiesResponseToMap(model.(*projectv1.ProjectConfigDefinitionResponseDAConfigDefinitionPropertiesResponse))
	} else if _, ok := model.(*projectv1.ProjectConfigDefinitionResponseResourceConfigDefinitionPropertiesResponse); ok {
		return ResourceIbmProjectConfigProjectConfigDefinitionResponseResourceConfigDefinitionPropertiesResponseToMap(model.(*projectv1.ProjectConfigDefinitionResponseResourceConfigDefinitionPropertiesResponse))
	} else if _, ok := model.(*projectv1.ProjectConfigDefinitionResponse); ok {
		modelMap := make(map[string]interface{})
		model := model.(*projectv1.ProjectConfigDefinitionResponse)
		if model.ComplianceProfile != nil {
			complianceProfileMap, err := ResourceIbmProjectConfigProjectComplianceProfileToMap(model.ComplianceProfile)
			if err != nil {
				return modelMap, err
			}
			if len(complianceProfileMap) > 0 {
				modelMap["compliance_profile"] = []map[string]interface{}{complianceProfileMap}
			}
		}
		if model.LocatorID != nil {
			modelMap["locator_id"] = *model.LocatorID
		}
		if model.Members != nil {
			members := []map[string]interface{}{}
			for _, membersItem := range model.Members {
				membersItemMap, err := ResourceIbmProjectConfigStackMemberToMap(&membersItem) // #nosec G601
				if err != nil {
					return modelMap, err
				}
				members = append(members, membersItemMap)
			}
			modelMap["members"] = members
		}
		if model.Description != nil {
			modelMap["description"] = *model.Description
		}
		if model.Name != nil {
			modelMap["name"] = *model.Name
		}
		if model.Authorizations != nil {
			authorizationsMap, err := ResourceIbmProjectConfigProjectConfigAuthToMap(model.Authorizations)
			if err != nil {
				return modelMap, err
			}
			modelMap["authorizations"] = []map[string]interface{}{authorizationsMap}
		}
		if model.Inputs != nil {
			inputs := make(map[string]interface{})
			for k, v := range model.Inputs {
				inputs[k] = flex.Stringify(v)
			}
			modelMap["inputs"] = inputs
		}
		if model.Settings != nil {
			settings := make(map[string]interface{})
			for k, v := range model.Settings {
				settings[k] = flex.Stringify(v)
			}
			modelMap["settings"] = settings
		}
		if model.EnvironmentID != nil {
			modelMap["environment_id"] = *model.EnvironmentID
		}
		if model.ResourceCrns != nil {
			modelMap["resource_crns"] = model.ResourceCrns
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized projectv1.ProjectConfigDefinitionResponseIntf subtype encountered")
	}
}

func ResourceIbmProjectConfigProjectComplianceProfileToMap(model projectv1.ProjectComplianceProfileIntf) (map[string]interface{}, error) {
	if _, ok := model.(*projectv1.ProjectComplianceProfileNullableObject); ok {
		return ResourceIbmProjectConfigProjectComplianceProfileNullableObjectToMap(model.(*projectv1.ProjectComplianceProfileNullableObject))
	} else if _, ok := model.(*projectv1.ProjectComplianceProfileV1); ok {
		return ResourceIbmProjectConfigProjectComplianceProfileV1ToMap(model.(*projectv1.ProjectComplianceProfileV1))
	} else if _, ok := model.(*projectv1.ProjectComplianceProfile); ok {
		modelMap := make(map[string]interface{})
		model := model.(*projectv1.ProjectComplianceProfile)
		if model.ID != nil {
			modelMap["id"] = *model.ID
		}
		if model.InstanceID != nil {
			modelMap["instance_id"] = *model.InstanceID
		}
		if model.InstanceLocation != nil {
			modelMap["instance_location"] = *model.InstanceLocation
		}
		if model.AttachmentID != nil {
			modelMap["attachment_id"] = *model.AttachmentID
		}
		if model.ProfileName != nil {
			modelMap["profile_name"] = *model.ProfileName
		}
		if model.WpPolicyID != nil {
			modelMap["wp_policy_id"] = *model.WpPolicyID
		}
		if model.WpInstanceID != nil {
			modelMap["wp_instance_id"] = *model.WpInstanceID
		}
		if model.WpInstanceName != nil {
			modelMap["wp_instance_name"] = *model.WpInstanceName
		}
		if model.WpInstanceLocation != nil {
			modelMap["wp_instance_location"] = *model.WpInstanceLocation
		}
		if model.WpZoneID != nil {
			modelMap["wp_zone_id"] = *model.WpZoneID
		}
		if model.WpZoneName != nil {
			modelMap["wp_zone_name"] = *model.WpZoneName
		}
		if model.WpPolicyName != nil {
			modelMap["wp_policy_name"] = *model.WpPolicyName
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized projectv1.ProjectComplianceProfileIntf subtype encountered")
	}
}

func ResourceIbmProjectConfigProjectComplianceProfileNullableObjectToMap(model *projectv1.ProjectComplianceProfileNullableObject) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	return modelMap, nil
}

func ResourceIbmProjectConfigProjectComplianceProfileV1ToMap(model *projectv1.ProjectComplianceProfileV1) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.InstanceID != nil {
		modelMap["instance_id"] = *model.InstanceID
	}
	if model.InstanceLocation != nil {
		modelMap["instance_location"] = *model.InstanceLocation
	}
	if model.AttachmentID != nil {
		modelMap["attachment_id"] = *model.AttachmentID
	}
	if model.ProfileName != nil {
		modelMap["profile_name"] = *model.ProfileName
	}
	if model.WpPolicyID != nil {
		modelMap["wp_policy_id"] = *model.WpPolicyID
	}
	if model.WpInstanceID != nil {
		modelMap["wp_instance_id"] = *model.WpInstanceID
	}
	if model.WpInstanceName != nil {
		modelMap["wp_instance_name"] = *model.WpInstanceName
	}
	if model.WpInstanceLocation != nil {
		modelMap["wp_instance_location"] = *model.WpInstanceLocation
	}
	if model.WpZoneID != nil {
		modelMap["wp_zone_id"] = *model.WpZoneID
	}
	if model.WpZoneName != nil {
		modelMap["wp_zone_name"] = *model.WpZoneName
	}
	if model.WpPolicyName != nil {
		modelMap["wp_policy_name"] = *model.WpPolicyName
	}
	return modelMap, nil
}

func ResourceIbmProjectConfigStackMemberToMap(model *projectv1.StackMember) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = *model.Name
	modelMap["config_id"] = *model.ConfigID
	return modelMap, nil
}

func ResourceIbmProjectConfigProjectConfigAuthToMap(model *projectv1.ProjectConfigAuth) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TrustedProfileID != nil {
		modelMap["trusted_profile_id"] = *model.TrustedProfileID
	}
	if model.Method != nil {
		modelMap["method"] = *model.Method
	}
	if model.ApiKey != nil {
		modelMap["api_key"] = *model.ApiKey
	}
	return modelMap, nil
}

func ResourceIbmProjectConfigProjectConfigDefinitionResponseDAConfigDefinitionPropertiesResponseToMap(model *projectv1.ProjectConfigDefinitionResponseDAConfigDefinitionPropertiesResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ComplianceProfile != nil {
		complianceProfileMap, err := ResourceIbmProjectConfigProjectComplianceProfileToMap(model.ComplianceProfile)
		if err != nil {
			return modelMap, err
		}
		modelMap["compliance_profile"] = []map[string]interface{}{complianceProfileMap}
	}
	if model.LocatorID != nil {
		modelMap["locator_id"] = *model.LocatorID
	}
	if model.Members != nil {
		members := []map[string]interface{}{}
		for _, membersItem := range model.Members {
			membersItemMap, err := ResourceIbmProjectConfigStackMemberToMap(&membersItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			members = append(members, membersItemMap)
		}
		modelMap["members"] = members
	}
	modelMap["description"] = *model.Description
	modelMap["name"] = *model.Name
	if model.Authorizations != nil {
		authorizationsMap, err := ResourceIbmProjectConfigProjectConfigAuthToMap(model.Authorizations)
		if err != nil {
			return modelMap, err
		}
		modelMap["authorizations"] = []map[string]interface{}{authorizationsMap}
	}
	if model.Inputs != nil {
		inputs := make(map[string]interface{})
		for k, v := range model.Inputs {
			inputs[k] = flex.Stringify(v)
		}
		modelMap["inputs"] = inputs
	}
	if model.Settings != nil {
		settings := make(map[string]interface{})
		for k, v := range model.Settings {
			settings[k] = flex.Stringify(v)
		}
		modelMap["settings"] = settings
	}
	if model.EnvironmentID != nil {
		modelMap["environment_id"] = *model.EnvironmentID
	}
	return modelMap, nil
}

func ResourceIbmProjectConfigProjectConfigDefinitionResponseResourceConfigDefinitionPropertiesResponseToMap(model *projectv1.ProjectConfigDefinitionResponseResourceConfigDefinitionPropertiesResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ResourceCrns != nil {
		modelMap["resource_crns"] = model.ResourceCrns
	}
	modelMap["description"] = *model.Description
	modelMap["name"] = *model.Name
	if model.Authorizations != nil {
		authorizationsMap, err := ResourceIbmProjectConfigProjectConfigAuthToMap(model.Authorizations)
		if err != nil {
			return modelMap, err
		}
		modelMap["authorizations"] = []map[string]interface{}{authorizationsMap}
	}
	if model.Inputs != nil {
		inputs := make(map[string]interface{})
		for k, v := range model.Inputs {
			inputs[k] = flex.Stringify(v)
		}
		modelMap["inputs"] = inputs
	}
	if model.Settings != nil {
		settings := make(map[string]interface{})
		for k, v := range model.Settings {
			settings[k] = flex.Stringify(v)
		}
		modelMap["settings"] = settings
	}
	if model.EnvironmentID != nil {
		modelMap["environment_id"] = *model.EnvironmentID
	}
	return modelMap, nil
}

func ResourceIbmProjectConfigProjectConfigNeedsAttentionStateToMap(model *projectv1.ProjectConfigNeedsAttentionState) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["event_id"] = *model.EventID
	modelMap["event"] = *model.Event
	modelMap["severity"] = *model.Severity
	if model.ActionURL != nil {
		modelMap["action_url"] = *model.ActionURL
	}
	if model.Target != nil {
		modelMap["target"] = *model.Target
	}
	if model.TriggeredBy != nil {
		modelMap["triggered_by"] = *model.TriggeredBy
	}
	modelMap["timestamp"] = model.Timestamp.String()
	return modelMap, nil
}

func ResourceIbmProjectConfigOutputValueToMap(model *projectv1.OutputValue) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Value != nil {
		modelMap["value"] = flex.Stringify(model.Value)
	}
	return modelMap, nil
}

func ResourceIbmProjectConfigReferenceValueToMap(model *projectv1.ReferenceValue) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	return modelMap, nil
}

func ResourceIbmProjectConfigProjectConfigErrorToMap(model *projectv1.ProjectConfigError) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	if model.Details != nil {
		detailsMap, err := ResourceIbmProjectConfigProjectConfigErrorDetailsToMap(model.Details)
		if err != nil {
			return modelMap, err
		}
		modelMap["details"] = []map[string]interface{}{detailsMap}
	}
	return modelMap, nil
}

func ResourceIbmProjectConfigProjectConfigErrorDetailsToMap(model *projectv1.ProjectConfigErrorDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	return modelMap, nil
}

func ResourceIbmProjectConfigProjectReferenceToMap(model *projectv1.ProjectReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["href"] = *model.Href
	definitionMap, err := ResourceIbmProjectConfigProjectDefinitionReferenceToMap(model.Definition)
	if err != nil {
		return modelMap, err
	}
	modelMap["definition"] = []map[string]interface{}{definitionMap}
	modelMap["crn"] = *model.Crn
	return modelMap, nil
}

func ResourceIbmProjectConfigProjectDefinitionReferenceToMap(model *projectv1.ProjectDefinitionReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = *model.Name
	return modelMap, nil
}

func ResourceIbmProjectConfigMemberOfDefinitionToMap(model *projectv1.MemberOfDefinition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	definitionMap, err := ResourceIbmProjectConfigStackConfigDefinitionSummaryToMap(model.Definition)
	if err != nil {
		return modelMap, err
	}
	modelMap["definition"] = []map[string]interface{}{definitionMap}
	modelMap["version"] = flex.IntValue(model.Version)
	modelMap["href"] = *model.Href
	return modelMap, nil
}

func ResourceIbmProjectConfigStackConfigDefinitionSummaryToMap(model *projectv1.StackConfigDefinitionSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = *model.Name
	members := []map[string]interface{}{}
	for _, membersItem := range model.Members {
		membersItemMap, err := ResourceIbmProjectConfigStackMemberToMap(&membersItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		members = append(members, membersItemMap)
	}
	modelMap["members"] = members
	return modelMap, nil
}

func ResourceIbmProjectConfigProjectConfigVersionSummaryToMap(model *projectv1.ProjectConfigVersionSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	definitionMap, err := ResourceIbmProjectConfigProjectConfigVersionDefinitionSummaryToMap(model.Definition)
	if err != nil {
		return modelMap, err
	}
	modelMap["definition"] = []map[string]interface{}{definitionMap}
	modelMap["state"] = *model.State
	modelMap["version"] = flex.IntValue(model.Version)
	modelMap["href"] = *model.Href
	return modelMap, nil
}

func ResourceIbmProjectConfigProjectConfigVersionDefinitionSummaryToMap(model *projectv1.ProjectConfigVersionDefinitionSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EnvironmentID != nil {
		modelMap["environment_id"] = *model.EnvironmentID
	}
	if model.LocatorID != nil {
		modelMap["locator_id"] = *model.LocatorID
	}
	return modelMap, nil
}
