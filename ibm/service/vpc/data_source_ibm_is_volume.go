// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMISVolume() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISVolumeRead,

		Schema: map[string]*schema.Schema{

			isVolumeName: {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{isVolumeName, "identifier"},
				ValidateFunc: validate.InvokeDataSourceValidator("ibm_is_volume", isVolumeName),
				Description:  "Volume name",
			},
			"identifier": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{isVolumeName, "identifier"},
				ValidateFunc: validate.InvokeDataSourceValidator("ibm_is_volume", "identifier"),
				Description:  "Volume name",
			},

			isVolumeZone: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Zone name",
			},
			isVolumesActive: &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether a running virtual server instance has an attachment to this volume.",
			},
			// defined_performance changes
			"adjustable_capacity_states": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The attachment states that support adjustable capacity for this volume.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"adjustable_iops_states": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The attachment states that support adjustable IOPS for this volume.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			isVolumeAttachmentState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The attachment state of the volume.",
			},
			isVolumeBandwidth: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The maximum bandwidth (in megabits per second) for the volume",
			},
			isVolumesBusy: &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this volume is performing an operation that must be serialized. If an operation specifies that it requires serialization, the operation will fail unless this property is `false`.",
			},
			isVolumesCreatedAt: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the volume was created.",
			},
			isVolumeResourceGroup: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource group name",
			},
			isVolumesOperatingSystem: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The operating system associated with this volume. If absent, this volume was notcreated from an image, or the image did not include an operating system.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isVolumeArchitecture: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The operating system architecture.",
						},
						isVolumeDHOnly: &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Images with this operating system can only be used on dedicated hosts or dedicated host groups.",
						},
						isVolumeDisplayName: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A unique, display-friendly name for the operating system.",
						},
						isVolumeOSFamily: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The software family for this operating system.",
						},

						isVolumesOperatingSystemHref: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this operating system.",
						},
						isVolumesOperatingSystemName: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The globally unique name for this operating system.",
						},
						isVolumeOSVendor: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vendor of the operating system.",
						},
						isVolumeOSVersion: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The major release version of this operating system.",
						},
					},
				},
			},
			isVolumeProfileName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Volume profile name",
			},

			isVolumeEncryptionKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Volume encryption key info",
			},

			isVolumeEncryptionType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Volume encryption type info",
			},

			isVolumeCapacity: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Vloume capacity value",
			},

			isVolumeIops: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "IOPS value for the Volume",
			},

			"storage_generation": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "storage_generation indicates which generation the profile family belongs to. For the custom and tiered profiles, this value is 1.",
			},

			isVolumeCrn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN value for the volume instance",
			},

			isVolumeStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Volume status",
			},
			isVolumeHealthReasons: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isVolumeHealthReasonsCode: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the reason for this health state.",
						},

						isVolumeHealthReasonsMessage: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the reason for this health state.",
						},

						isVolumeHealthReasonsMoreInfo: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about the reason for this health state.",
						},
					},
				},
			},
			isVolumeCatalogOffering: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The catalog offering this volume was created from. If a virtual server instance is provisioned with a boot_volume_attachment specifying this volume, the virtual server instance will use this volume's catalog offering, including its pricing plan.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isVolumeCatalogOfferingPlanCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this catalog offering version's billing plan",
						},
						"deleted": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and provides some supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						isVolumeCatalogOfferingVersionCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this version of a catalog offering",
						},
					},
				},
			},

			isVolumeHealthState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The health of this resource.",
			},
			isVolumeStatusReasons: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isVolumeStatusReasonsCode: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the status reason",
						},

						isVolumeStatusReasonsMessage: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the status reason",
						},

						isVolumeStatusReasonsMoreInfo: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about this status reason",
						},
					},
				},
			},

			isVolumeTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "Tags for the volume instance",
			},

			isVolumeAccessTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "Access management tags for the volume instance",
			},

			isVolumeSourceSnapshot: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Identifier of the snapshot from which this volume was cloned",
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
			"allowed_use": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The usage constraints to be matched against the requested instance or bare metal server properties to determine compatibility.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bare_metal_server": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expression that must be satisfied by the properties of a bare metal server provisioned using the image data in this volume.",
						},
						"instance": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expression that must be satisfied by the properties of a virtual server instance provisioned using this volume.",
						},
						"api_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The API version with which to evaluate the expressions.",
						},
					},
				},
			},
		},
	}
}

func DataSourceIBMISVolumeValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "identifier",
			ValidateFunctionIdentifier: validate.ValidateNoZeroValues,
			Type:                       validate.TypeString})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isVolumeName,
			ValidateFunctionIdentifier: validate.ValidateNoZeroValues,
			Type:                       validate.TypeString})

	ibmISVoulmeDataSourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_volume", Schema: validateSchema}
	return &ibmISVoulmeDataSourceValidator
}

func dataSourceIBMISVolumeRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	err := volumeGet(context, d, meta)
	if err != nil {
		return err
	}
	return nil
}

func volumeGet(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_volume", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	var volume vpcv1.Volume
	if volName, ok := d.GetOk(isVolumeName); ok {
		name := volName.(string)
		zone := ""
		if zname, ok := d.GetOk(isVolumeZone); ok {
			zone = zname.(string)
		}
		listVolumesOptions := &vpcv1.ListVolumesOptions{
			Name: &name,
		}

		if zone != "" {
			listVolumesOptions.ZoneName = &zone
		}
		listVolumesOptions.Name = &name
		vols, _, err := sess.ListVolumesWithContext(context, listVolumesOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListVolumesWithContext failed: %s", err.Error()), "(Data) ibm_is_volume", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		allrecs := vols.Volumes

		if len(allrecs) == 0 {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("No Volume found with name: %s", name), "(Data) ibm_is_volume", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		volume = allrecs[0]
	} else {
		identifier := d.Get("identifier").(string)
		getVolumeOptions := &vpcv1.GetVolumeOptions{
			ID: &identifier,
		}

		volPtr, _, err := sess.GetVolumeWithContext(context, getVolumeOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "(Data) ibm_is_volume", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		volume = *volPtr
	}
	d.SetId(*volume.ID)
	if err = d.Set("active", volume.Active); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting active: %s", err), "(Data) ibm_is_volume", "read", "set-active").GetDiag()
	}
	if err = d.Set("attachment_state", volume.AttachmentState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting attachment_state: %s", err), "(Data) ibm_is_volume", "read", "set-attachment_state").GetDiag()
	}
	if err = d.Set("bandwidth", flex.IntValue(volume.Bandwidth)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting bandwidth: %s", err), "(Data) ibm_is_volume", "read", "set-bandwidth").GetDiag()
	}
	if err = d.Set("busy", volume.Busy); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting busy: %s", err), "(Data) ibm_is_volume", "read", "set-busy").GetDiag()
	}
	if err = d.Set("capacity", flex.IntValue(volume.Capacity)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting capacity: %s", err), "(Data) ibm_is_volume", "read", "set-capacity").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(volume.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_volume", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("name", volume.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_volume", "read", "set-name").GetDiag()
	}
	d.Set("identifier", *volume.ID)
	if volume.OperatingSystem != nil {
		operatingSystemList := []map[string]interface{}{}
		operatingSystemMap := dataSourceVolumeCollectionVolumesOperatingSystemToMap(*volume.OperatingSystem)
		operatingSystemList = append(operatingSystemList, operatingSystemMap)
		if err = d.Set(isVolumesOperatingSystem, operatingSystemList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting operating_system: %s", err), "(Data) ibm_is_volume", "read", "set-operating_system").GetDiag()
		}
	}
	if err = d.Set(isVolumeProfileName, *volume.Profile.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting profile: %s", err), "(Data) ibm_is_volume", "read", "set-profile").GetDiag()
	}
	if err = d.Set(isVolumeZone, *volume.Zone.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting zone: %s", err), "(Data) ibm_is_volume", "read", "set-zone").GetDiag()
	}
	if volume.EncryptionKey != nil {
		if err = d.Set(isVolumeEncryptionKey, volume.EncryptionKey.CRN); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting encryption_key: %s", err), "(Data) ibm_is_volume", "read", "set-encryption_key").GetDiag()
		}
	}
	if err = d.Set("encryption_type", volume.Encryption); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting encryption_type: %s", err), "(Data) ibm_is_volume", "read", "set-encryption_type").GetDiag()
	}
	if volume.SourceSnapshot != nil {
		if err = d.Set(isVolumeSourceSnapshot, *volume.SourceSnapshot.ID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting source_snapshot: %s", err), "(Data) ibm_is_volume", "read", "set-source_snapshot").GetDiag()
		}
	}
	if err = d.Set("iops", flex.IntValue(volume.Iops)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting iops: %s", err), "(Data) ibm_is_volume", "read", "set-iops").GetDiag()
	}

	if err = d.Set("storage_generation", flex.IntValue(volume.StorageGeneration)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting storage_generation: %s", err), "(Data) ibm_is_volume", "read", "set-storage_generation").GetDiag()
	}

	if err = d.Set("capacity", flex.IntValue(volume.Capacity)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting capacity: %s", err), "(Data) ibm_is_volume", "read", "set-capacity").GetDiag()
	}
	if err = d.Set("crn", volume.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_volume", "read", "set-crn").GetDiag()
	}
	if err = d.Set("status", volume.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_volume", "read", "set-status").GetDiag()
	}
	if volume.StatusReasons != nil {
		statusReasonsList := make([]map[string]interface{}, 0)
		for _, sr := range volume.StatusReasons {
			currentSR := map[string]interface{}{}
			if sr.Code != nil && sr.Message != nil {
				currentSR[isVolumeStatusReasonsCode] = *sr.Code
				currentSR[isVolumeStatusReasonsMessage] = *sr.Message
				if sr.MoreInfo != nil {
					currentSR[isVolumeStatusReasonsMoreInfo] = *sr.Message
				}
				statusReasonsList = append(statusReasonsList, currentSR)
			}
			if err = d.Set(isVolumeStatusReasons, statusReasonsList); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status_reasons: %s", err), "(Data) ibm_is_volume", "read", "set-status_reasons").GetDiag()
			}
		}
	}
	if err = d.Set(isVolumeTags, volume.UserTags); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tags: %s", err), "(Data) ibm_is_volume", "read", "set-tags").GetDiag()
	}
	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *volume.CRN, "", isVolumeAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource vpc volume (%s) access tags: %s", d.Id(), err)
	}
	if err = d.Set(isVolumeAccessTags, accesstags); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_tags: %s", err), "(Data) ibm_is_volume", "read", "set-access_tags").GetDiag()
	}
	controller, err := flex.GetBaseController(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBaseController failed: %s", err.Error()), "(Data) ibm_is_volume", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/storage/storageVolumes"); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_controller_url: %s", err), "(Data) ibm_is_volume", "read", "set-resource_controller_url").GetDiag()
	}

	if err = d.Set(flex.ResourceName, *volume.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_name: %s", err), "(Data) ibm_is_volume", "read", "set-resource_name").GetDiag()
	}

	if err = d.Set(flex.ResourceCRN, *volume.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_crn: %s", err), "(Data) ibm_is_volume", "read", "set-resource_crn").GetDiag()
	}
	if err = d.Set("resource_status", volume.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_status: %s", err), "(Data) ibm_is_volume", "read", "set-resource_status").GetDiag()
	}
	if volume.ResourceGroup != nil {
		if err = d.Set(flex.ResourceGroupName, volume.ResourceGroup.Name); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group_name: %s", err), "(Data) ibm_is_volume", "read", "set-resource_group_name").GetDiag()
		}
		if err = d.Set(isVolumeResourceGroup, *volume.ResourceGroup.ID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_volume", "read", "set-resource_group").GetDiag()
		}
	}

	if volume.HealthReasons != nil {
		healthReasonsList := make([]map[string]interface{}, 0)
		for _, sr := range volume.HealthReasons {
			currentSR := map[string]interface{}{}
			if sr.Code != nil && sr.Message != nil {
				currentSR[isVolumeHealthReasonsCode] = *sr.Code
				currentSR[isVolumeHealthReasonsMessage] = *sr.Message
				if sr.MoreInfo != nil {
					currentSR[isVolumeHealthReasonsMoreInfo] = *sr.Message
				}
				healthReasonsList = append(healthReasonsList, currentSR)
			}
		}
		if err = d.Set(isVolumeHealthReasons, healthReasonsList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting health_reasons: %s", err), "(Data) ibm_is_volume", "read", "set-health_reasons").GetDiag()
		}
	}
	// catalog
	if volume.CatalogOffering != nil {
		versionCrn := ""
		if volume.CatalogOffering.Version != nil && volume.CatalogOffering.Version.CRN != nil {
			versionCrn = *volume.CatalogOffering.Version.CRN
		}
		catalogList := make([]map[string]interface{}, 0)
		catalogMap := map[string]interface{}{}
		if versionCrn != "" {
			catalogMap[isVolumeCatalogOfferingVersionCrn] = versionCrn
		}
		if volume.CatalogOffering.Plan != nil {
			planCrn := ""
			if volume.CatalogOffering.Plan.CRN != nil {
				planCrn = *volume.CatalogOffering.Plan.CRN
			}
			if planCrn != "" {
				catalogMap[isVolumeCatalogOfferingPlanCrn] = *volume.CatalogOffering.Plan.CRN
			}
			if volume.CatalogOffering.Plan.Deleted != nil {
				deletedMap := resourceIbmIsVolumeCatalogOfferingVersionPlanReferenceDeletedToMap(*volume.CatalogOffering.Plan.Deleted)
				catalogMap["deleted"] = []map[string]interface{}{deletedMap}
			}
		}
		catalogList = append(catalogList, catalogMap)

		if err = d.Set(isVolumeCatalogOffering, catalogList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting catalog_offering: %s", err), "(Data) ibm_is_volume", "read", "set-catalog_offering").GetDiag()
		}
	}
	if err = d.Set("health_state", volume.HealthState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting health_state: %s", err), "(Data) ibm_is_volume", "read", "set-health_state").GetDiag()
	}

	if err = d.Set("adjustable_capacity_states", volume.AdjustableCapacityStates); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting adjustable_capacity_states: %s", err), "(Data) ibm_is_volume", "read", "set-adjustable_capacity_states").GetDiag()
	}
	if err = d.Set("adjustable_iops_states", volume.AdjustableIopsStates); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting adjustable_iops_states: %s", err), "(Data) ibm_is_volume", "read", "set-adjustable_iops_states").GetDiag()
	}

	allowedUses := []map[string]interface{}{}
	if volume.AllowedUse != nil {
		modelMap, err := ResourceceIBMIsVolumeAllowedUseToMap(volume.AllowedUse)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allowed_use: %s", err), "(Data) ibm_is_volume", "read", "set-allowed_use").GetDiag()
		}
		allowedUses = append(allowedUses, modelMap)
	}
	if err = d.Set("allowed_use", allowedUses); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allowed_use: %s", err), "(Data) ibm_is_volume", "read", "set-allowed_use").GetDiag()
	}
	return nil
}

func resourceIbmIsVolumeCatalogOfferingVersionPlanReferenceDeletedToMap(catalogOfferingVersionPlanReferenceDeleted vpcv1.Deleted) map[string]interface{} {
	catalogOfferingVersionPlanReferenceDeletedMap := map[string]interface{}{}

	catalogOfferingVersionPlanReferenceDeletedMap["more_info"] = catalogOfferingVersionPlanReferenceDeleted.MoreInfo

	return catalogOfferingVersionPlanReferenceDeletedMap
}
