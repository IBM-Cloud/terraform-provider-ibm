// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.3-943fbc81-20260603-173645
*/

package brsmigration

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
	"github.com/IBM/ibm-brs-migration-sdk-go/brsmigrationv1"
)

func ResourceIbmBrsMigrationDiscover() *schema.Resource {
	return &schema.Resource{
		CreateContext:   resourceIbmBrsMigrationDiscoverCreate,
		ReadContext:     resourceIbmBrsMigrationDiscoverRead,
		DeleteContext:   resourceIbmBrsMigrationDiscoverDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"brs_instance": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validate.InvokeValidator("ibm_brs_migration_discover", "brs_instance"),
				Description: "The BRS service instance identifier — either a bare UUID (instance ID) or a full BRS service instance CRN (percent-encoded when supplied in a URL, since CRNs contain colons and forward-slashes).",
			},
			"migration_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validate.InvokeValidator("ibm_brs_migration_discover", "migration_id"),
				Description: "The migration project ID (uuid4 format).",
			},
			"env": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validate.InvokeValidator("ibm_brs_migration_discover", "env"),
				Description: "Infrastructure environment being discovered.",
			},
			"location": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validate.InvokeValidator("ibm_brs_migration_discover", "location"),
				Description: "Location scope used for this discovery job (datacenter or zone).",
			},
			"vpc_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				ValidateFunc: validate.InvokeValidator("ibm_brs_migration_discover", "vpc_id"),
				Description: "VPC ID used to scope discovery. Present for env=vpc only.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current lifecycle state of the discovery job.",
			},
			"start_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Start of the time window used for this discovery run.",
			},
			"end_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "End of the time window used for this discovery run.",
			},
			"message": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Human-readable status or error message.",
			},
			"summary": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Counts of discovered resources by compute and storage type.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Total number of compute resources discovered.",
						},
						"compute": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "Compute resource counts by type.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"virtual_server": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Number of Virtual Server Instances discovered.",
									},
									"bare_metal": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Number of bare metal servers discovered.",
									},
								},
							},
						},
						"storage": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "Storage volume counts by type.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"block": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Number of block volumes discovered.",
									},
									"file": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Number of file shares discovered.",
									},
									"san": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Number of SAN volumes discovered (Classic only).",
									},
									"local": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Number of local disks discovered.",
									},
								},
							},
						},
					},
				},
			},
			"brs_migration_discover_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique discovery job ID (job-{uuid4} format).",
			},
		},
	}
}

func ResourceIbmBrsMigrationDiscoverValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "brs_instance",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|crn:v1:bluemix:public:backup-recovery:[a-z0-9-]+:a\/[a-zA-Z0-9]+:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}::)$`,
			MinValueLength:             1,
			MaxValueLength:             512,
		},
		validate.ValidateSchema{
			Identifier:                 "migration_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		},
		validate.ValidateSchema{
			Identifier:                 "env",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "classic, vpc",
		},
		validate.ValidateSchema{
			Identifier:                 "location",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-z][a-z0-9-]+[0-9]$`,
			MinValueLength:             3,
			MaxValueLength:             32,
		},
		validate.ValidateSchema{
			Identifier:                 "vpc_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^r[0-9]+-[0-9a-zA-Z-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_brs_migration_discover", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmBrsMigrationDiscoverCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_discover", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createDiscoverOptions := &brsmigrationv1.CreateDiscoverOptions{}

	createDiscoverOptions.SetBrsInstance(d.Get("brs_instance").(string))
	createDiscoverOptions.SetMigrationID(d.Get("migration_id").(string))
	createDiscoverOptions.SetEnv(d.Get("env").(string))
	createDiscoverOptions.SetLocation(d.Get("location").(string))
	if _, ok := d.GetOk("vpc_id"); ok {
		createDiscoverOptions.SetVpcID(d.Get("vpc_id").(string))
	}
	if _, ok := d.GetOk("credentials_crn"); ok {
		createDiscoverOptions.SetCredentialsCrn(d.Get("credentials_crn").(string))
	}
	if _, ok := d.GetOk("compute"); ok {
		computeModel, err := ResourceIbmBrsMigrationDiscoverMapToDiscoverJobPrototypeCompute(d.Get("compute.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_discover", "create", "parse-compute").GetDiag()
		}
		createDiscoverOptions.SetCompute(computeModel)
	}
	if _, ok := d.GetOk("storage"); ok {
		storageModel, err := ResourceIbmBrsMigrationDiscoverMapToDiscoverJobPrototypeStorage(d.Get("storage.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_discover", "create", "parse-storage").GetDiag()
		}
		createDiscoverOptions.SetStorage(storageModel)
	}

	discoverJobAccepted, _, err := brsMigrationClient.CreateDiscoverWithContext(context, createDiscoverOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDiscoverWithContext failed: %s", err.Error()), "ibm_brs_migration_discover", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", *createDiscoverOptions.BrsInstance, *createDiscoverOptions.MigrationID, *discoverJobAccepted.ID))

	return resourceIbmBrsMigrationDiscoverRead(context, d, meta)
}

func resourceIbmBrsMigrationDiscoverRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_discover", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDiscoverOptions := &brsmigrationv1.GetDiscoverOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_discover", "read", "sep-id-parts").GetDiag()
	}

	getDiscoverOptions.SetBrsInstance(parts[0])
	getDiscoverOptions.SetMigrationID(parts[1])

	discoverJob, response, err := brsMigrationClient.GetDiscoverWithContext(context, getDiscoverOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDiscoverWithContext failed: %s", err.Error()), "ibm_brs_migration_discover", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("env", discoverJob.Env); err != nil {
		err = fmt.Errorf("Error setting env: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_discover", "read", "set-env").GetDiag()
	}
	if err = d.Set("location", discoverJob.Location); err != nil {
		err = fmt.Errorf("Error setting location: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_discover", "read", "set-location").GetDiag()
	}
	if !core.IsNil(discoverJob.VpcID) {
		if err = d.Set("vpc_id", discoverJob.VpcID); err != nil {
			err = fmt.Errorf("Error setting vpc_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_discover", "read", "set-vpc_id").GetDiag()
		}
	}
	if err = d.Set("state", discoverJob.State); err != nil {
		err = fmt.Errorf("Error setting state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_discover", "read", "set-state").GetDiag()
	}
	if !core.IsNil(discoverJob.StartTime) {
		if err = d.Set("start_time", flex.DateTimeToString(discoverJob.StartTime)); err != nil {
			err = fmt.Errorf("Error setting start_time: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_discover", "read", "set-start_time").GetDiag()
		}
	}
	if !core.IsNil(discoverJob.EndTime) {
		if err = d.Set("end_time", flex.DateTimeToString(discoverJob.EndTime)); err != nil {
			err = fmt.Errorf("Error setting end_time: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_discover", "read", "set-end_time").GetDiag()
		}
	}
	if !core.IsNil(discoverJob.Message) {
		if err = d.Set("message", discoverJob.Message); err != nil {
			err = fmt.Errorf("Error setting message: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_discover", "read", "set-message").GetDiag()
		}
	}
	if !core.IsNil(discoverJob.Summary) {
		summaryMap, err := ResourceIbmBrsMigrationDiscoverDiscoverJobSummaryToMap(discoverJob.Summary)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_discover", "read", "summary-to-map").GetDiag()
		}
		if err = d.Set("summary", []map[string]interface{}{summaryMap}); err != nil {
			err = fmt.Errorf("Error setting summary: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_discover", "read", "set-summary").GetDiag()
		}
	}
	if err = d.Set("brs_migration_discover_id", discoverJob.ID); err != nil {
		err = fmt.Errorf("Error setting brs_migration_discover_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_discover", "read", "set-brs_migration_discover_id").GetDiag()
	}

	return nil
}

func resourceIbmBrsMigrationDiscoverDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_discover", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	cancelDiscoverOptions := &brsmigrationv1.CancelDiscoverOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_discover", "delete", "sep-id-parts").GetDiag()
	}

	cancelDiscoverOptions.SetBrsInstance(parts[0])
	cancelDiscoverOptions.SetMigrationID(parts[1])

	_, err = brsMigrationClient.CancelDiscoverWithContext(context, cancelDiscoverOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CancelDiscoverWithContext failed: %s", err.Error()), "ibm_brs_migration_discover", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIbmBrsMigrationDiscoverMapToDiscoverJobPrototypeCompute(modelMap map[string]interface{}) (*brsmigrationv1.DiscoverJobPrototypeCompute, error) {
	model := &brsmigrationv1.DiscoverJobPrototypeCompute{}
	if modelMap["types"] != nil {
		types := []string{}
		for _, typesItem := range modelMap["types"].([]interface{}) {
			types = append(types, typesItem.(string))
		}
		model.Types = types
	}
	return model, nil
}

func ResourceIbmBrsMigrationDiscoverMapToDiscoverJobPrototypeStorage(modelMap map[string]interface{}) (*brsmigrationv1.DiscoverJobPrototypeStorage, error) {
	model := &brsmigrationv1.DiscoverJobPrototypeStorage{}
	if modelMap["types"] != nil {
		types := []string{}
		for _, typesItem := range modelMap["types"].([]interface{}) {
			types = append(types, typesItem.(string))
		}
		model.Types = types
	}
	return model, nil
}

func ResourceIbmBrsMigrationDiscoverDiscoverJobSummaryToMap(model *brsmigrationv1.DiscoverJobSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Total != nil {
		modelMap["total"] = flex.IntValue(model.Total)
	}
	if model.Compute != nil {
		computeMap, err := ResourceIbmBrsMigrationDiscoverDiscoverJobSummaryComputeToMap(model.Compute)
		if err != nil {
			return modelMap, err
		}
		modelMap["compute"] = []map[string]interface{}{computeMap}
	}
	if model.Storage != nil {
		storageMap, err := ResourceIbmBrsMigrationDiscoverDiscoverJobSummaryStorageToMap(model.Storage)
		if err != nil {
			return modelMap, err
		}
		modelMap["storage"] = []map[string]interface{}{storageMap}
	}
	return modelMap, nil
}

func ResourceIbmBrsMigrationDiscoverDiscoverJobSummaryComputeToMap(model *brsmigrationv1.DiscoverJobSummaryCompute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.VirtualServer != nil {
		modelMap["virtual_server"] = flex.IntValue(model.VirtualServer)
	}
	if model.BareMetal != nil {
		modelMap["bare_metal"] = flex.IntValue(model.BareMetal)
	}
	return modelMap, nil
}

func ResourceIbmBrsMigrationDiscoverDiscoverJobSummaryStorageToMap(model *brsmigrationv1.DiscoverJobSummaryStorage) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Block != nil {
		modelMap["block"] = flex.IntValue(model.Block)
	}
	if model.File != nil {
		modelMap["file"] = flex.IntValue(model.File)
	}
	if model.San != nil {
		modelMap["san"] = flex.IntValue(model.San)
	}
	if model.Local != nil {
		modelMap["local"] = flex.IntValue(model.Local)
	}
	return modelMap, nil
}
