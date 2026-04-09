// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.102.0-615ec964-20250307-203034
 */

package codeengine

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
)

func ResourceIbmCodeEngineAllowedOutboundDestination() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmCodeEngineAllowedOutboundDestinationCreate,
		ReadContext:   resourceIbmCodeEngineAllowedOutboundDestinationRead,
		UpdateContext: resourceIbmCodeEngineAllowedOutboundDestinationUpdate,
		DeleteContext: resourceIbmCodeEngineAllowedOutboundDestinationDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_allowed_outbound_destination", "project_id"),
				Description:  "The ID of the project.",
			},
			"type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_allowed_outbound_destination", "type"),
				Description:  "Specify the type of the allowed outbound destination. Allowed types are: `cidr_block` and `private_path_service_gateway`.",
			},
			"cidr_block": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_allowed_outbound_destination", "cidr_block"),
				Description:  "The IPv4 address range.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_allowed_outbound_destination", "name"),
				Description:  "The name of the allowed outbound destination.",
			},
			"private_path_service_gateway_crn": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_allowed_outbound_destination", "private_path_service_gateway_crn"),
				Description:  "The CRN of the Private Path service.",
			},
			"isolation_policy": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "shared",
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_allowed_outbound_destination", "isolation_policy"),
				Description:  "Optional property to specify the isolation policy of the private path service gateway. If set to `shared`, other projects within the same account or enterprise account family can connect to Private Path service, too. If set to `dedicated` the gateway can only be used by a single Code Engine project. If not specified the isolation policy will be set to `shared`.",
			},
			"entity_tag": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version of the allowed outbound destination, which is used to achieve optimistic locking.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current status of the outbound destination.",
			},
			"status_details": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"endpoint_gateway": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Optional information about the endpoint gateway located in the Code Engine VPC that connects to the private path service gateway.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The account that created the endpoint gateway.",
									},
									"created_at": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The timestamp when the endpoint gateway was created.",
									},
									"ips": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The reserved IPs bound to this endpoint gateway.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name for this endpoint gateway. The name is unique across all endpoint gateways in the VPC.",
									},
								},
							},
						},
						"private_path_service_gateway": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Optional information about the private path service gateway that this allowed outbound destination points to.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The private path service gateway identifier.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of private path service gateway.",
									},
									"service_endpoints": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The fully qualified domain names for this private path service gateway. The domains are used for endpoint gateways to connect to the service and are configured in the VPC for each endpoint gateway.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"reason": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Optional information to provide more context in case of a 'failed' or 'deploying' status.",
						},
					},
				},
			},
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIbmCodeEngineAllowedOutboundDestinationValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "project_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		},
		validate.ValidateSchema{
			Identifier:                 "type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "cidr_block, private_path_service_gateway",
			Regexp:                     `^(cidr_block|private_path_service_gateway)$`,
		},
		validate.ValidateSchema{
			Identifier:                 "cidr_block",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$`,
			MinValueLength:             9,
			MaxValueLength:             18,
		},
		validate.ValidateSchema{
			Identifier:                 "private_path_service_gateway_crn",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^crn\:v1\:[a-zA-Z0-9]*\:(public|dedicated|local)\:is\:([a-z][\-a-z0-9_]*[a-z0-9])?\:((a|o|s)\/[\-a-z0-9]+)?\:\:private-path-service-gateway\:[\-a-zA-Z0-9\/.]*$`,
			MinValueLength:             20,
			MaxValueLength:             253,
		},
		validate.ValidateSchema{
			Identifier:                 "isolation_policy",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "dedicated, shared",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_code_engine_allowed_outbound_destination", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmCodeEngineAllowedOutboundDestinationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bodyModelMap := map[string]interface{}{}
	createAllowedOutboundDestinationOptions := &codeenginev2.CreateAllowedOutboundDestinationOptions{}

	bodyModelMap["name"] = d.Get("name")
	bodyModelMap["type"] = d.Get("type")
	if _, ok := d.GetOk("cidr_block"); ok {
		bodyModelMap["cidr_block"] = d.Get("cidr_block")
	}
	if _, ok := d.GetOk("private_path_service_gateway_crn"); ok {
		bodyModelMap["private_path_service_gateway_crn"] = d.Get("private_path_service_gateway_crn")
	}
	if _, ok := d.GetOk("isolation_policy"); ok {
		bodyModelMap["isolation_policy"] = d.Get("isolation_policy")
	}
	createAllowedOutboundDestinationOptions.SetProjectID(d.Get("project_id").(string))
	convertedModel, err := ResourceIbmCodeEngineAllowedOutboundDestinationMapToAllowedOutboundDestinationPrototype(bodyModelMap)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "create", "parse-request-body").GetDiag()
	}
	createAllowedOutboundDestinationOptions.AllowedOutboundDestination = convertedModel

	allowedOutboundDestinationIntf, _, err := codeEngineClient.CreateAllowedOutboundDestinationWithContext(context, createAllowedOutboundDestinationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateAllowedOutboundDestinationWithContext failed: %s", err.Error()), "ibm_code_engine_allowed_outbound_destination", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allowedOutboundDestination := allowedOutboundDestinationIntf.(*codeenginev2.AllowedOutboundDestination)
	d.SetId(fmt.Sprintf("%s/%s", *createAllowedOutboundDestinationOptions.ProjectID, *allowedOutboundDestination.Name))

	_, err = waitForIbmCodeEngineAllowedOutboundDestinationCreate(d, meta)
	if err != nil {
		errMsg := fmt.Sprintf("Error waiting for resource IbmCodeEngineAllowedOutboundDestination (%s) to be created: %s", d.Id(), err)
		return flex.DiscriminatedTerraformErrorf(err, errMsg, "ibm_code_engine_allowed_outbound_destination", "create", "wait-for-state").GetDiag()
	}

	return resourceIbmCodeEngineAllowedOutboundDestinationRead(context, d, meta)
}

func waitForIbmCodeEngineAllowedOutboundDestinationCreate(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return false, err
	}
	getAllowedOutboundDestinationOptions := &codeenginev2.GetAllowedOutboundDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return false, err
	}

	getAllowedOutboundDestinationOptions.SetProjectID(parts[0])
	getAllowedOutboundDestinationOptions.SetName(parts[1])

	stateConf := &resource.StateChangeConf{
		Pending: []string{"deploying"},
		Target:  []string{"ready"},
		Refresh: func() (interface{}, string, error) {
			stateObj, response, err := codeEngineClient.GetAllowedOutboundDestination(getAllowedOutboundDestinationOptions)
			if err != nil {
				if sdkErr, ok := err.(*core.SDKProblem); ok && response.GetStatusCode() == 404 {
					sdkErr.Summary = fmt.Sprintf("The instance %s does not exist anymore: %s", "getAllowedOutboundDestinationOptions", err)
					return nil, "", sdkErr
				}
				return nil, "", err
			}
			failStates := map[string]bool{"failed": true}
			stateObjF := stateObj.(*codeenginev2.AllowedOutboundDestination)
			if failStates[*stateObjF.Status] {
				return stateObj, *stateObjF.Status, fmt.Errorf("The instance %s failed: %s", "getAllowedOutboundDestinationOptions", err)
			}
			return stateObj, *stateObjF.Status, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 60 * time.Second,
	}

	return stateConf.WaitForStateContext(context.Background())
}

func resourceIbmCodeEngineAllowedOutboundDestinationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getAllowedOutboundDestinationOptions := &codeenginev2.GetAllowedOutboundDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "read", "sep-id-parts").GetDiag()
	}

	getAllowedOutboundDestinationOptions.SetProjectID(parts[0])
	getAllowedOutboundDestinationOptions.SetName(parts[1])

	allowedOutboundDestinationIntf, response, err := codeEngineClient.GetAllowedOutboundDestinationWithContext(context, getAllowedOutboundDestinationOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAllowedOutboundDestinationWithContext failed: %s", err.Error()), "ibm_code_engine_allowed_outbound_destination", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allowedOutboundDestination := allowedOutboundDestinationIntf.(*codeenginev2.AllowedOutboundDestination)
	if err = d.Set("type", allowedOutboundDestination.Type); err != nil {
		err = fmt.Errorf("Error setting type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "read", "set-type").GetDiag()
	}
	if !core.IsNil(allowedOutboundDestination.CidrBlock) {
		if err = d.Set("cidr_block", allowedOutboundDestination.CidrBlock); err != nil {
			err = fmt.Errorf("Error setting cidr_block: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "read", "set-cidr_block").GetDiag()
		}
	}
	if err = d.Set("name", allowedOutboundDestination.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "read", "set-name").GetDiag()
	}
	if !core.IsNil(allowedOutboundDestination.PrivatePathServiceGatewayCrn) {
		if err = d.Set("private_path_service_gateway_crn", allowedOutboundDestination.PrivatePathServiceGatewayCrn); err != nil {
			err = fmt.Errorf("Error setting private_path_service_gateway_crn: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "read", "set-private_path_service_gateway_crn").GetDiag()
		}
	}
	if !core.IsNil(allowedOutboundDestination.IsolationPolicy) {
		if err = d.Set("isolation_policy", allowedOutboundDestination.IsolationPolicy); err != nil {
			err = fmt.Errorf("Error setting isolation_policy: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "read", "set-isolation_policy").GetDiag()
		}
	}
	if !core.IsNil(allowedOutboundDestination.EntityTag) {
		if err = d.Set("entity_tag", allowedOutboundDestination.EntityTag); err != nil {
			err = fmt.Errorf("Error setting entity_tag: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "read", "set-entity_tag").GetDiag()
		}
	}
	if !core.IsNil(allowedOutboundDestination.Status) {
		if err = d.Set("status", allowedOutboundDestination.Status); err != nil {
			err = fmt.Errorf("Error setting status: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "read", "set-status").GetDiag()
		}
	}

	if !core.IsNil(allowedOutboundDestination.StatusDetails) {
		statusDetailsMap, err := ResourceIbmCodeEngineAllowedOutboundDestinationAllowedOutboundStatusDetailsToMap(allowedOutboundDestination.StatusDetails)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "read", "status_details-to-map").GetDiag()
		}
		// Only set status_details if the map contains actual data
		if len(statusDetailsMap) > 0 {
			if err = d.Set("status_details", []map[string]interface{}{statusDetailsMap}); err != nil {
				err = fmt.Errorf("Error setting status_details: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "read", "set-status_details").GetDiag()
			}
		} else {
			// Explicitly set to empty list when StatusDetails exists but has no data
			if err = d.Set("status_details", []map[string]interface{}{}); err != nil {
				err = fmt.Errorf("Error setting status_details: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "read", "set-status_details").GetDiag()
			}
		}
	} else {
		// Explicitly set to empty list when StatusDetails is nil
		if err = d.Set("status_details", []map[string]interface{}{}); err != nil {
			err = fmt.Errorf("Error setting status_details: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "read", "set-status_details").GetDiag()
		}
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_code_engine_allowed_outbound_destination", "read", "set-etag").GetDiag()
	}

	return nil
}

func resourceIbmCodeEngineAllowedOutboundDestinationUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateAllowedOutboundDestinationOptions := &codeenginev2.UpdateAllowedOutboundDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "update", "sep-id-parts").GetDiag()
	}

	updateAllowedOutboundDestinationOptions.SetProjectID(parts[0])
	updateAllowedOutboundDestinationOptions.SetName(parts[1])

	hasChange := false

	patchVals := &codeenginev2.AllowedOutboundDestinationPatch{}
	if d.HasChange("project_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "project_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_code_engine_allowed_outbound_destination", "update", "project_id-forces-new").GetDiag()
	}
	if d.HasChange("cidr_block") {
		newCidrBlock := d.Get("cidr_block").(string)
		patchVals.CidrBlock = &newCidrBlock
		hasChange = true
	}
	if d.HasChange("isolation_policy") {
		newIsolationPolicy := d.Get("isolation_policy").(string)
		patchVals.IsolationPolicy = &newIsolationPolicy
		hasChange = true
	}
	updateAllowedOutboundDestinationOptions.SetIfMatch(d.Get("etag").(string))

	if hasChange {
		// Fields with `nil` values are omitted from the generic map,
		// so we need to re-add them to support removing arguments
		// in merge-patch operations sent to the service.
		updateAllowedOutboundDestinationOptions.AllowedOutboundDestination = ResourceIbmCodeEngineAllowedOutboundDestinationAllowedOutboundDestinationPatchAsPatch(patchVals, d)

		_, _, err = codeEngineClient.UpdateAllowedOutboundDestinationWithContext(context, updateAllowedOutboundDestinationOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateAllowedOutboundDestinationWithContext failed: %s", err.Error()), "ibm_code_engine_allowed_outbound_destination", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIbmCodeEngineAllowedOutboundDestinationRead(context, d, meta)
}

func resourceIbmCodeEngineAllowedOutboundDestinationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteAllowedOutboundDestinationOptions := &codeenginev2.DeleteAllowedOutboundDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_allowed_outbound_destination", "delete", "sep-id-parts").GetDiag()
	}

	deleteAllowedOutboundDestinationOptions.SetProjectID(parts[0])
	deleteAllowedOutboundDestinationOptions.SetName(parts[1])

	_, err = codeEngineClient.DeleteAllowedOutboundDestinationWithContext(context, deleteAllowedOutboundDestinationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteAllowedOutboundDestinationWithContext failed: %s", err.Error()), "ibm_code_engine_allowed_outbound_destination", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIbmCodeEngineAllowedOutboundDestinationMapToAllowedOutboundDestinationPrototype(modelMap map[string]interface{}) (codeenginev2.AllowedOutboundDestinationPrototypeIntf, error) {
	model := &codeenginev2.AllowedOutboundDestinationPrototype{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["cidr_block"] != nil && modelMap["cidr_block"].(string) != "" {
		model.CidrBlock = core.StringPtr(modelMap["cidr_block"].(string))
	}
	if modelMap["private_path_service_gateway_crn"] != nil && modelMap["private_path_service_gateway_crn"].(string) != "" {
		model.PrivatePathServiceGatewayCrn = core.StringPtr(modelMap["private_path_service_gateway_crn"].(string))
	}
	if modelMap["isolation_policy"] != nil && modelMap["isolation_policy"].(string) != "" {
		model.IsolationPolicy = core.StringPtr(modelMap["isolation_policy"].(string))
	}
	return model, nil
}

func ResourceIbmCodeEngineAllowedOutboundDestinationMapToAllowedOutboundDestinationPrototypeCidrBlockDataPrototype(modelMap map[string]interface{}) (*codeenginev2.AllowedOutboundDestinationPrototypeCidrBlockDataPrototype, error) {
	model := &codeenginev2.AllowedOutboundDestinationPrototypeCidrBlockDataPrototype{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.CidrBlock = core.StringPtr(modelMap["cidr_block"].(string))
	return model, nil
}

func ResourceIbmCodeEngineAllowedOutboundDestinationMapToAllowedOutboundDestinationPrototypePrivatePathServiceGatewayDataPrototype(modelMap map[string]interface{}) (*codeenginev2.AllowedOutboundDestinationPrototypePrivatePathServiceGatewayDataPrototype, error) {
	model := &codeenginev2.AllowedOutboundDestinationPrototypePrivatePathServiceGatewayDataPrototype{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.PrivatePathServiceGatewayCrn = core.StringPtr(modelMap["private_path_service_gateway_crn"].(string))
	if modelMap["isolation_policy"] != nil && modelMap["isolation_policy"].(string) != "" {
		model.IsolationPolicy = core.StringPtr(modelMap["isolation_policy"].(string))
	}
	return model, nil
}

func ResourceIbmCodeEngineAllowedOutboundDestinationAllowedOutboundStatusDetailsToMap(model codeenginev2.AllowedOutboundStatusDetailsIntf) (map[string]interface{}, error) {
	if _, ok := model.(*codeenginev2.AllowedOutboundStatusDetailsPrivatePathServiceGatewayStatusDetails); ok {
		return ResourceIbmCodeEngineAllowedOutboundDestinationAllowedOutboundStatusDetailsPrivatePathServiceGatewayStatusDetailsToMap(model.(*codeenginev2.AllowedOutboundStatusDetailsPrivatePathServiceGatewayStatusDetails))
	} else if _, ok := model.(*codeenginev2.AllowedOutboundStatusDetails); ok {
		modelMap := make(map[string]interface{})
		model := model.(*codeenginev2.AllowedOutboundStatusDetails)
		if model.EndpointGateway != nil {
			endpointGatewayMap, err := ResourceIbmCodeEngineAllowedOutboundDestinationEndpointGatewayDetailsToMap(model.EndpointGateway)
			if err != nil {
				return modelMap, err
			}
			modelMap["endpoint_gateway"] = []map[string]interface{}{endpointGatewayMap}
		}
		if model.PrivatePathServiceGateway != nil {
			privatePathServiceGatewayMap, err := ResourceIbmCodeEngineAllowedOutboundDestinationPrivatePathServiceGatewayDetailsToMap(model.PrivatePathServiceGateway)
			if err != nil {
				return modelMap, err
			}
			modelMap["private_path_service_gateway"] = []map[string]interface{}{privatePathServiceGatewayMap}
		}
		if model.Reason != nil {
			modelMap["reason"] = *model.Reason
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized codeenginev2.AllowedOutboundStatusDetailsIntf subtype encountered")
	}
}

func ResourceIbmCodeEngineAllowedOutboundDestinationEndpointGatewayDetailsToMap(model *codeenginev2.EndpointGatewayDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AccountID != nil {
		modelMap["account_id"] = *model.AccountID
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = *model.CreatedAt
	}
	if model.Ips != nil {
		modelMap["ips"] = model.Ips
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap, nil
}

func ResourceIbmCodeEngineAllowedOutboundDestinationPrivatePathServiceGatewayDetailsToMap(model *codeenginev2.PrivatePathServiceGatewayDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.ServiceEndpoints != nil {
		modelMap["service_endpoints"] = model.ServiceEndpoints
	}
	return modelMap, nil
}

func ResourceIbmCodeEngineAllowedOutboundDestinationAllowedOutboundStatusDetailsPrivatePathServiceGatewayStatusDetailsToMap(model *codeenginev2.AllowedOutboundStatusDetailsPrivatePathServiceGatewayStatusDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EndpointGateway != nil {
		endpointGatewayMap, err := ResourceIbmCodeEngineAllowedOutboundDestinationEndpointGatewayDetailsToMap(model.EndpointGateway)
		if err != nil {
			return modelMap, err
		}
		modelMap["endpoint_gateway"] = []map[string]interface{}{endpointGatewayMap}
	}
	if model.PrivatePathServiceGateway != nil {
		privatePathServiceGatewayMap, err := ResourceIbmCodeEngineAllowedOutboundDestinationPrivatePathServiceGatewayDetailsToMap(model.PrivatePathServiceGateway)
		if err != nil {
			return modelMap, err
		}
		modelMap["private_path_service_gateway"] = []map[string]interface{}{privatePathServiceGatewayMap}
	}
	if model.Reason != nil {
		modelMap["reason"] = *model.Reason
	}
	return modelMap, nil
}

func ResourceIbmCodeEngineAllowedOutboundDestinationAllowedOutboundDestinationPatchAsPatch(patchVals *codeenginev2.AllowedOutboundDestinationPatch, d *schema.ResourceData) map[string]interface{} {
	patch, _ := patchVals.AsPatch()
	var path string

	path = "cidr_block"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["cidr_block"] = nil
	} else if !exists {
		delete(patch, "cidr_block")
	}
	path = "isolation_policy"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["isolation_policy"] = nil
	} else if !exists {
		delete(patch, "isolation_policy")
	}

	return patch
}
