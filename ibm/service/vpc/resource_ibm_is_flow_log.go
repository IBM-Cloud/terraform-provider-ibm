// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isFlowLogName                  = "name"
	isFlowLogActive                = "active"
	isFlowLogStorageBucket         = "storage_bucket"
	isFlowLogStorageBucketEndPoint = "endpoint"
	isFlowLogTarget                = "target"
	isFlowLogResourceGroup         = "resource_group"
	isFlowLogTargetType            = "resource_type"
	isFlowLogCreatedAt             = "created_at"
	isFlowLogCrn                   = "crn"
	isFlowLogLifecycleState        = "lifecycle_state"
	isFlowLogHref                  = "href"
	isFlowLogAutoDelete            = "auto_delete"
	isFlowLogVpc                   = "vpc"
	isFlowLogTags                  = "tags"
	isFlowLogAccessTags            = "access_tags"
)

func ResourceIBMISFlowLog() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISFlowLogCreate,
		ReadContext:   resourceIBMISFlowLogRead,
		UpdateContext: resourceIBMISFlowLogUpdate,
		DeleteContext: resourceIBMISFlowLogDelete,
		Exists:        resourceIBMISFlowLogExists,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceTagsCustomizeDiff(diff)
				},
			),
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceValidateAccessTags(diff, v)
				}),
		),

		Schema: map[string]*schema.Schema{
			isFlowLogName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				Description:  "Flow Log Collector name",
				ValidateFunc: validate.InvokeValidator("ibm_is_flow_log", isFlowLogName),
			},

			isFlowLogStorageBucket: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Cloud Object Storage bucket name where the collected flows will be logged",
			},

			isFlowLogTarget: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The target id that the flow log collector is to collect flow logs",
			},

			isFlowLogActive: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Indicates whether this collector is active",
			},

			isFlowLogResourceGroup: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "The resource group of flow log",
			},

			isFlowLogCrn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this flow log collector",
			},

			isFlowLogHref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this flow log collector",
			},

			isFlowLogCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time flow log was created",
			},

			isFlowLogVpc: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The VPC this flow log collector is associated with",
			},

			isFlowLogAutoDelete: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If set to true, this flow log collector will be automatically deleted when the target is deleted",
			},

			isFlowLogLifecycleState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the flow log collector",
			},

			isFlowLogTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_flow_log", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "Tags for the VPC Flow logs",
			},

			isFlowLogAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_flow_log", "accesstag")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
			},

			flex.ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},

			flex.ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			flex.ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			flex.ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func ResourceIBMISFlowLogValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isFlowLogName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "accesstag",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	ibmISFlowLogValidator := validate.ResourceValidator{ResourceName: "ibm_is_flow_log", Schema: validateSchema}
	return &ibmISFlowLogValidator
}

func resourceIBMISFlowLogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_ibm_is_flow_log", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createFlowLogCollectorOptionsModel := &vpcv1.CreateFlowLogCollectorOptions{}
	name := d.Get(isFlowLogName).(string)
	createFlowLogCollectorOptionsModel.Name = &name
	if _, ok := d.GetOk(isFlowLogResourceGroup); ok {
		group := d.Get(isFlowLogResourceGroup).(string)
		resourceGroupIdentityModel := new(vpcv1.ResourceGroupIdentityByID)
		resourceGroupIdentityModel.ID = &group
		createFlowLogCollectorOptionsModel.ResourceGroup = resourceGroupIdentityModel
	}

	if v, ok := d.GetOkExists(isFlowLogActive); ok {
		active := v.(bool)
		createFlowLogCollectorOptionsModel.Active = &active
	}

	target := d.Get(isFlowLogTarget).(string)
	FlowLogCollectorTargetModel := &vpcv1.FlowLogCollectorTargetPrototype{}
	FlowLogCollectorTargetModel.ID = &target
	createFlowLogCollectorOptionsModel.Target = FlowLogCollectorTargetModel

	bucketname := d.Get(isFlowLogStorageBucket).(string)
	cloudObjectStorageBucketIdentityModel := new(vpcv1.LegacyCloudObjectStorageBucketIdentityCloudObjectStorageBucketIdentityByName)
	cloudObjectStorageBucketIdentityModel.Name = &bucketname
	createFlowLogCollectorOptionsModel.StorageBucket = cloudObjectStorageBucketIdentityModel

	flowlogCollector, _, err := vpcClient.CreateFlowLogCollectorWithContext(ctx, createFlowLogCollectorOptionsModel)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateFlowLogCollectorWithContext failed: %s", err.Error()), "ibm_ibm_is_flow_log", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*flowlogCollector.ID)

	log.Printf("[INFO] Flow log collector : %s", *flowlogCollector.ID)

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isFlowLogTags); ok || v != "" {
		oldList, newList := d.GetChange(isFlowLogTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *flowlogCollector.CRN, "", isUserTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource vpc flow log (%s) tags: %s", d.Id(), err)
		}
	}
	if _, ok := d.GetOk(isFlowLogAccessTags); ok {
		oldList, newList := d.GetChange(isFlowLogAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *flowlogCollector.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource VPC Flow Log (%s) access tags: %s", d.Id(), err)
		}
	}
	return resourceIBMISFlowLogRead(ctx, d, meta)
}

func resourceIBMISFlowLogRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_ibm_is_flow_log", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	ID := d.Id()

	getOptions := &vpcv1.GetFlowLogCollectorOptions{
		ID: &ID,
	}
	flowLogCollector, response, err := vpcClient.GetFlowLogCollectorWithContext(context, getOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetFlowLogCollectorWithContext failed: %s", err.Error()), "ibm_ibm_is_flow_log", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(flowLogCollector.Name) {
		if err = d.Set("name", flowLogCollector.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_ibm_is_flow_log", "read", "set-name").GetDiag()
		}
	}

	if !core.IsNil(flowLogCollector.Active) {
		if err = d.Set("active", flowLogCollector.Active); err != nil {
			err = fmt.Errorf("Error setting active: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_ibm_is_flow_log", "read", "set-active").GetDiag()
		}
	}

	if err = d.Set("created_at", flex.DateTimeToString(flowLogCollector.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_ibm_is_flow_log", "read", "set-created_at").GetDiag()
	}

	if err = d.Set("href", flowLogCollector.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_ibm_is_flow_log", "read", "set-href").GetDiag()
	}

	if err = d.Set("crn", flowLogCollector.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_ibm_is_flow_log", "read", "set-crn").GetDiag()
	}

	if err = d.Set("lifecycle_state", flowLogCollector.LifecycleState); err != nil {
		err = fmt.Errorf("Error setting lifecycle_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_ibm_is_flow_log", "read", "set-lifecycle_state").GetDiag()
	}

	if flowLogCollector.VPC != nil {
		if err = d.Set(isFlowLogVpc, *flowLogCollector.VPC.ID); err != nil {
			err = fmt.Errorf("Error setting vpc: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_ibm_is_flow_log", "read", "set-vpc").GetDiag()
		}
	}

	if flowLogCollector.Target != nil {
		targetIntf := flowLogCollector.Target
		target := targetIntf.(*vpcv1.FlowLogCollectorTarget)
		if err = d.Set(isFlowLogTarget, *target.ID); err != nil {
			err = fmt.Errorf("Error setting target: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_ibm_is_flow_log", "read", "set-target").GetDiag()
		}
	}

	if flowLogCollector.StorageBucket != nil {
		bucket := flowLogCollector.StorageBucket
		if err = d.Set(isFlowLogStorageBucket, *bucket.Name); err != nil {
			err = fmt.Errorf("Error setting storage_bucket: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_ibm_is_flow_log", "read", "set-storage_bucket").GetDiag()
		}
	}

	tags, err := flex.GetGlobalTagsUsingCRN(meta, *flowLogCollector.CRN, "", isUserTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource vpc flow log (%s) tags: %s", d.Id(), err)
	}
	if err = d.Set(isFlowLogTags, tags); err != nil {
		err = fmt.Errorf("Error setting tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_flow_log", "read", "set-tags").GetDiag()
	}
	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *flowLogCollector.CRN, "", isAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource VPC Flow Log (%s) access tags: %s", d.Id(), err)
	}
	if err = d.Set(isFlowLogAccessTags, accesstags); err != nil {
		err = fmt.Errorf("Error setting access_tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_flow_log", "read", "set-access_tags").GetDiag()
	}

	controller, err := flex.GetBaseController(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/network/flowLogs"); err != nil {
		err = fmt.Errorf("Error setting controller_url: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_flow_log", "read", "set-controller_url").GetDiag()
	}
	if err = d.Set(flex.ResourceName, *flowLogCollector.Name); err != nil {
		err = fmt.Errorf("Error setting resource_name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_flow_log", "read", "set-resource_name").GetDiag()
	}
	if err = d.Set(flex.ResourceCRN, *flowLogCollector.CRN); err != nil {
		err = fmt.Errorf("Error setting resource_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_flow_log", "read", "set-resource_crn").GetDiag()
	}
	if err = d.Set(flex.ResourceStatus, *flowLogCollector.LifecycleState); err != nil {
		err = fmt.Errorf("Error setting resource_status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_flow_log", "read", "set-resource_status").GetDiag()
	}

	if flowLogCollector.ResourceGroup != nil {
		if err = d.Set(isFlowLogResourceGroup, *flowLogCollector.ResourceGroup.ID); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_flow_log", "read", "set-resource_group").GetDiag()
		}
		if err = d.Set(flex.ResourceGroupName, *flowLogCollector.ResourceGroup.ID); err != nil {
			err = fmt.Errorf("Error setting resource_group_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_flow_log", "read", "set-flex_resource_group_name").GetDiag()
		}
	}

	return nil
}

func resourceIBMISFlowLogUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_ibm_is_flow_log", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	ID := d.Id()

	getOptions := &vpcv1.GetFlowLogCollectorOptions{
		ID: &ID,
	}
	flowlogCollector, _, err := vpcClient.GetFlowLogCollectorWithContext(context, getOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetFlowLogCollectorWithContext failed: %s", err.Error()), "ibm_ibm_is_flow_log", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if d.HasChange(isFlowLogTags) {
		oldList, newList := d.GetChange(isFlowLogTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *flowlogCollector.CRN, "", isUserTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource flow log (%s) tags: %s", *flowlogCollector.ID, err)
		}
	}

	if d.HasChange(isFlowLogAccessTags) {
		oldList, newList := d.GetChange(isFlowLogAccessTags)
		err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *flowlogCollector.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource flow log (%s) access tags: %s", d.Id(), err)
		}
	}

	if d.HasChange(isFlowLogActive) || d.HasChange(isFlowLogName) {
		active := d.Get(isFlowLogActive).(bool)
		name := d.Get(isFlowLogName).(string)
		updoptions := &vpcv1.UpdateFlowLogCollectorOptions{
			ID: &ID,
		}
		flowLogCollectorPatchModel := &vpcv1.FlowLogCollectorPatch{
			Active: &active,
			Name:   &name,
		}
		flowLogCollectorPatch, err := flowLogCollectorPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error calling asPatch for FlowLogCollectorPatch: %s", err.Error()), "ibm_ibm_is_flow_log", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updoptions.FlowLogCollectorPatch = flowLogCollectorPatch
		_, _, err = vpcClient.UpdateFlowLogCollectorWithContext(context, updoptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateFlowLogCollectorWithContext failed: %s", err.Error()), "ibm_ibm_is_flow_log", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMISFlowLogRead(context, d, meta)
}

func resourceIBMISFlowLogDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_ibm_is_flow_log", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	ID := d.Id()
	delOptions := &vpcv1.DeleteFlowLogCollectorOptions{
		ID: &ID,
	}
	response, err := vpcClient.DeleteFlowLogCollectorWithContext(context, delOptions)

	if err != nil && response.StatusCode != 404 {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteFlowLogCollectorWithContext failed: %s", err.Error()), "ibm_ibm_is_flow_log", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}

func resourceIBMISFlowLogExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_flow_log", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, fmt.Errorf("[ERROR] Error initializing VPC client: %v", tfErr.GetDiag())
	}

	ID := d.Id()

	getOptions := &vpcv1.GetFlowLogCollectorOptions{
		ID: &ID,
	}
	_, response, err := vpcClient.GetFlowLogCollector(getOptions)
	if err != nil && response.StatusCode != 404 {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetFlowLogCollectorWithContext failed: %s\n%s", err, response), "ibm_is_flow_log", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, fmt.Errorf("[ERROR] Error checking existence of Flow Log Collector: %v", tfErr.GetDiag())
	}
	if response.StatusCode == 404 {
		d.SetId("")
		return false, nil
	}
	return true, nil
}
