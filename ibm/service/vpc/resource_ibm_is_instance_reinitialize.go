// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMISInstanceReinitialize() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISInstanceReinitializeCreate,
		ReadContext:   resourceIBMISInstanceReinitializeRead,
		DeleteContext: resourceIBMISInstanceReinitializeDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(15 * time.Minute),
			Update: schema.DefaultTimeout(15 * time.Minute),
			Delete: schema.DefaultTimeout(15 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Instance identifier",
			},
			"image": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Description:   "The image to be used when reinitializing the instance",
				ConflictsWith: []string{"boot_volume_attachment"},
			},
			"boot_volume_attachment": {
				Type:          schema.TypeList,
				Optional:      true,
				ForceNew:      true,
				MaxItems:      1,
				Description:   "The boot volume attachment for the virtual server instance. If not specified, a new boot volume attachment will be created.",
				ConflictsWith: []string{"image"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"volume": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							Description: "The boot volume attachment configuration for reinitialization by volume",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:          schema.TypeString,
										Optional:      true,
										ForceNew:      true,
										Description:   "The ID of the volume to attach as boot volume",
										ConflictsWith: []string{"boot_volume_attachment.0.volume.0.source_snapshot"},
									},
									"source_snapshot": {
										Type:          schema.TypeList,
										Optional:      true,
										ForceNew:      true,
										MaxItems:      1,
										Description:   "The snapshot to use as a source for the volume's data. The specified snapshot may be in a different account, subject to IAM policies.",
										ConflictsWith: []string{"boot_volume_attachment.0.volume.0.id"},
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeString,
													Required:    true,
													ForceNew:    true,
													Description: "The ID of the snapshot",
												},
											},
										},
									},
									"allowed_use": {
										Type:        schema.TypeList,
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										Description: "The allowed use configuration for this volume",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"api_version": {
													Type:        schema.TypeString,
													Required:    true,
													ForceNew:    true,
													Description: "The API version with which to evaluate the expressions.",
												},
												"bare_metal_server": {
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Description: "The expression that must be satisfied by the properties of a bare metal server provisioned using the image data in this volume.",
												},
												"instance": {
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Description: "The expression that must be satisfied by the properties of a virtual server instance provisioned using this volume.",
												},
											},
										},
									},
									"bandwidth": {
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Description: "The maximum bandwidth (in megabits per second) for the volume.",
									},
									"capacity": {
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Description: "The capacity to use for the volume (in gigabytes).",
									},
									"encryption_key": {
										Type:        schema.TypeList,
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										Description: "The root key to use to wrap the data encryption key for the volume.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"crn": {
													Type:        schema.TypeString,
													Required:    true,
													ForceNew:    true,
													Description: "The CRN of the Key Protect Root Key for this resource.",
												},
											},
										},
									},
									"iops": {
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Description: "The maximum I/O operations per second (IOPS) to use for this volume.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "The name for this volume.",
									},
									"profile": {
										Type:        schema.TypeList,
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										Description: "The profile for this volume.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Required:    true,
													ForceNew:    true,
													Description: "The globally unique name for this volume profile",
												},
											},
										},
									},
									"resource_group": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "The resource group to use for this volume.",
									},
									"user_tags": {
										Type:        schema.TypeSet,
										Optional:    true,
										ForceNew:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Set:         schema.HashString,
										Description: "The user tags associated with this volume.",
									},
								},
							},
						},
						"delete_volume_on_instance_delete": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Description: "Indicates whether the volume will be deleted when the instance is deleted",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The name of the boot volume attachment",
						},
					},
				},
			},
			"default_trusted_profile": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: "The default trusted profile configuration to use for this virtual server instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_link": {
							Type:        schema.TypeBool,
							Required:    true,
							ForceNew:    true,
							Description: "If set to true, the system will create a link to the specified target trusted profile.",
						},
						"target": {
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							MaxItems:    1,
							Description: "The default trusted profile configuration to use for this virtual server instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "The unique identifier for this trusted profile",
									},
									"crn": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "The CRN for this trusted profile",
									},
								},
							},
						},
					},
				},
			},
			"keys": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "SSH key Ids for the instance",
			},
			"user_data": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Instance user data to replace initialization",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the instance after reinitialization",
			},
		},
	}
}

func resourceIBMISInstanceReinitializeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var instanceId string
	if instId, ok := d.GetOk("instance_id"); ok {
		instanceId = instId.(string)
	}

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_reinitialize", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Stop instance if running
	stopInstanceIfRunning := false
	stopInstanceIfRunning, err = resourceStopInstanceIfRunning(ctx, d, sess, instanceId)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("resourceStopInstanceIfRunning failed: %s", err.Error()), "ibm_is_instance_reinitialize", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Build the reinitialize options
	reinitializeOptions := &vpcv1.ReinitializeInstanceOptions{
		ID: &instanceId,
	}

	// Determine which reinitialization prototype to use
	var instanceReinitializePrototype vpcv1.InstanceReinitializePrototypeIntf

	// Build common fields
	var userData *string
	if userDataVal, ok := d.GetOk("user_data"); ok {
		userDataStr := userDataVal.(string)
		userData = &userDataStr
	}

	var keys []vpcv1.KeyIdentityIntf
	if keysSet, ok := d.GetOk("keys"); ok {
		keySet := keysSet.(*schema.Set)
		if keySet.Len() > 0 {
			keys = make([]vpcv1.KeyIdentityIntf, keySet.Len())
			for i, key := range keySet.List() {
				keystr := key.(string)
				keys[i] = &vpcv1.KeyIdentity{
					ID: &keystr,
				}
			}
		}
	}

	var defaultTrustedProfile *vpcv1.InstanceDefaultTrustedProfilePrototype
	if defaultTrustedProfileList, ok := d.GetOk("default_trusted_profile"); ok && len(defaultTrustedProfileList.([]interface{})) > 0 {
		trustedProfileConfig := defaultTrustedProfileList.([]interface{})[0].(map[string]interface{})
		trustedProfilePrototype := &vpcv1.InstanceDefaultTrustedProfilePrototype{}

		if autoLink, ok := trustedProfileConfig["auto_link"]; ok && autoLink.(string) != "" {
			autoLinkStr := autoLink.(bool)
			trustedProfilePrototype.AutoLink = &autoLinkStr
		}

		if targetList, ok := trustedProfileConfig["target"]; ok && len(targetList.([]interface{})) > 0 {
			targetConfig := targetList.([]interface{})[0].(map[string]interface{})
			targetIdentity := &vpcv1.TrustedProfileIdentity{}

			if id, ok := targetConfig["id"]; ok && id.(string) != "" {
				idStr := id.(string)
				targetIdentity.ID = &idStr
			}
			if crn, ok := targetConfig["crn"]; ok && crn.(string) != "" {
				crnStr := crn.(string)
				targetIdentity.CRN = &crnStr
			}
			trustedProfilePrototype.Target = targetIdentity
		}
		defaultTrustedProfile = trustedProfilePrototype
	}

	// Check for reinitialize by image
	if imageId, ok := d.GetOk("image"); ok {
		imageStr := imageId.(string)

		instanceReinitializePrototype = &vpcv1.InstanceReinitializePrototypeInstanceReinitializeByImage{
			Image: &vpcv1.ImageIdentityByID{
				ID: &imageStr,
			},
			UserData:              userData,
			Keys:                  keys,
			DefaultTrustedProfile: defaultTrustedProfile,
		}
		// Reinitialize by boot volume attachment
		if bootVolumeAttachmentList, ok := d.GetOk("boot_volume_attachment"); ok && len(bootVolumeAttachmentList.([]interface{})) > 0 {

			bootVolumeAttachmentConfig := bootVolumeAttachmentList.([]interface{})[0].(map[string]interface{})
			bootvolumeAttachment := &vpcv1.VolumeAttachmentPrototypeInstanceByImageContext{}

			// Set volume attachment name
			if name, ok := bootVolumeAttachmentConfig["name"]; ok && name.(string) != "" {
				nameStr := name.(string)
				bootvolumeAttachment.Name = &nameStr
			}

			// Set delete volume on instance delete
			if deleteOnDelete, ok := bootVolumeAttachmentConfig["delete_volume_on_instance_delete"]; ok {
				deleteBool := deleteOnDelete.(bool)
				bootvolumeAttachment.DeleteVolumeOnInstanceDelete = &deleteBool
			}

			// Build volume prototype
			if volumeList, ok := bootVolumeAttachmentConfig["volume"]; ok && len(volumeList.([]interface{})) > 0 {
				volumeConfig := volumeList.([]interface{})[0].(map[string]interface{})
				volumePrototype := &vpcv1.VolumePrototypeInstanceByImageContext{}

				// Set volume name
				if volumeName, ok := volumeConfig["name"]; ok && volumeName.(string) != "" {
					volumeNameStr := volumeName.(string)
					volumePrototype.Name = &volumeNameStr
				}

				// Set capacity
				if capacity, ok := volumeConfig["capacity"]; ok && capacity.(int) > 0 {
					capacityInt64 := int64(capacity.(int))
					volumePrototype.Capacity = &capacityInt64
				}

				// Set IOPS
				if iops, ok := volumeConfig["iops"]; ok && iops.(int) > 0 {
					iopsInt64 := int64(iops.(int))
					volumePrototype.Iops = &iopsInt64
				}

				// Set bandwidth
				if bandwidth, ok := volumeConfig["bandwidth"]; ok && bandwidth.(int) > 0 {
					bandwidthInt64 := int64(bandwidth.(int))
					volumePrototype.Bandwidth = &bandwidthInt64
				}

				// Set profile
				if profileList, ok := volumeConfig["profile"]; ok && len(profileList.([]interface{})) > 0 {
					profileConfig := profileList.([]interface{})[0].(map[string]interface{})
					if profileName, ok := profileConfig["name"]; ok && profileName.(string) != "" {
						profileNameStr := profileName.(string)
						volumePrototype.Profile = &vpcv1.VolumeProfileIdentity{
							Name: &profileNameStr,
						}
					}
				}

				// Set encryption key
				if encryptionKeyList, ok := volumeConfig["encryption_key"]; ok && len(encryptionKeyList.([]interface{})) > 0 {
					encryptionKeyConfig := encryptionKeyList.([]interface{})[0].(map[string]interface{})
					if crn, ok := encryptionKeyConfig["crn"]; ok && crn.(string) != "" {
						crnStr := crn.(string)
						volumePrototype.EncryptionKey = &vpcv1.EncryptionKeyIdentity{
							CRN: &crnStr,
						}
					}
				}

				// Set resource group
				if resourceGroup, ok := volumeConfig["resource_group"]; ok && resourceGroup.(string) != "" {
					resourceGroupStr := resourceGroup.(string)
					volumePrototype.ResourceGroup = &vpcv1.ResourceGroupIdentity{
						ID: &resourceGroupStr,
					}
				}

				// Set user tags
				if userTagsSet, ok := volumeConfig["user_tags"]; ok {
					userTags := userTagsSet.(*schema.Set)
					if userTags.Len() > 0 {
						tags := make([]string, userTags.Len())
						for i, tag := range userTags.List() {
							tags[i] = tag.(string)
						}
						volumePrototype.UserTags = tags
					}
				}

				// Set allowed use
				if allowedUseList, ok := volumeConfig["allowed_use"]; ok && len(allowedUseList.([]interface{})) > 0 {
					allowedUseConfig := allowedUseList.([]interface{})[0].(map[string]interface{})
					allowedUse := &vpcv1.VolumeAllowedUsePrototype{}

					if apiVersion, ok := allowedUseConfig["api_version"]; ok && apiVersion.(string) != "" {
						apiVersionStr := apiVersion.(string)
						allowedUse.ApiVersion = &apiVersionStr
					}
					if bareMetalServer, ok := allowedUseConfig["bare_metal_server"]; ok && bareMetalServer.(string) != "" {
						bareMetalServerStr := bareMetalServer.(string)
						allowedUse.BareMetalServer = &bareMetalServerStr
					}
					if instance, ok := allowedUseConfig["instance"]; ok && instance.(string) != "" {
						instanceStr := instance.(string)
						allowedUse.Instance = &instanceStr
					}
					volumePrototype.AllowedUse = allowedUse
				}

				bootvolumeAttachment.Volume = volumePrototype
			}

			instanceReinitializePrototype = &vpcv1.InstanceReinitializePrototypeInstanceReinitializeByImage{
				BootVolumeAttachment:  bootvolumeAttachment,
				UserData:              userData,
				Keys:                  keys,
				DefaultTrustedProfile: defaultTrustedProfile,
			}
			reinitializeOptions.InstanceReinitializePrototype = instanceReinitializePrototype
			log.Printf("[INFO] Reinitializing instance %s by boot volume attachment", instanceId)
			log.Printf("[INFO] Reinitializing instance %s by image: %s", instanceId, imageStr)
		}

	} else if bootVolumeAttachmentList, ok := d.GetOk("boot_volume_attachment"); ok && len(bootVolumeAttachmentList.([]interface{})) > 0 {
		// Reinitialize by boot volume attachment
		bootVolumeAttachmentConfig := bootVolumeAttachmentList.([]interface{})[0].(map[string]interface{})

		// Build volume prototype
		if volumeList, ok := bootVolumeAttachmentConfig["volume"]; ok && len(volumeList.([]interface{})) > 0 {
			volumeConfig := volumeList.([]interface{})[0].(map[string]interface{})

			// Check for volume ID or source snapshot
			if volumeId, ok := volumeConfig["id"]; ok && volumeId.(string) != "" {
				bootvolumeAttachment := &vpcv1.VolumeAttachmentPrototypeInstanceByVolumeContext{}
				// Set volume attachment name
				if name, ok := bootVolumeAttachmentConfig["name"]; ok && name.(string) != "" {
					nameStr := name.(string)
					bootvolumeAttachment.Name = &nameStr
				}

				// Set delete volume on instance delete
				if deleteOnDelete, ok := bootVolumeAttachmentConfig["delete_volume_on_instance_delete"]; ok {
					deleteBool := deleteOnDelete.(bool)
					bootvolumeAttachment.DeleteVolumeOnInstanceDelete = &deleteBool
				}
				volumeIdStr := volumeId.(string)
				bootvolumeAttachment.Volume = &vpcv1.VolumeIdentityByID{
					ID: &volumeIdStr,
				}
				instanceReinitializePrototype = &vpcv1.InstanceReinitializePrototypeInstanceReinitializeByVolume{
					BootVolumeAttachment:  bootvolumeAttachment,
					UserData:              userData,
					Keys:                  keys,
					DefaultTrustedProfile: defaultTrustedProfile,
				}
				reinitializeOptions.InstanceReinitializePrototype = instanceReinitializePrototype
				log.Printf("[INFO] Reinitializing instance %s by volume ID: %s", instanceId, volumeIdStr)

			} else if sourceSnapshotList, ok := volumeConfig["source_snapshot"]; ok && len(sourceSnapshotList.([]interface{})) > 0 {
				sourceSnapshotConfig := sourceSnapshotList.([]interface{})[0].(map[string]interface{})
				if snapshotId, ok := sourceSnapshotConfig["id"]; ok && snapshotId.(string) != "" {
					bootvolumeAttachment := &vpcv1.VolumeAttachmentPrototypeInstanceBySourceSnapshotContext{}
					// Set volume attachment name
					if name, ok := bootVolumeAttachmentConfig["name"]; ok && name.(string) != "" {
						nameStr := name.(string)
						bootvolumeAttachment.Name = &nameStr
					}

					// Set delete volume on instance delete
					if deleteOnDelete, ok := bootVolumeAttachmentConfig["delete_volume_on_instance_delete"]; ok {
						deleteBool := deleteOnDelete.(bool)
						bootvolumeAttachment.DeleteVolumeOnInstanceDelete = &deleteBool
					}
					snapshotIdStr := snapshotId.(string)
					volumePrototype := &vpcv1.VolumePrototypeInstanceBySourceSnapshotContext{}
					volumePrototype.SourceSnapshot = &vpcv1.SnapshotIdentityByID{
						ID: &snapshotIdStr,
					}
					// Set volume name
					if volumeName, ok := volumeConfig["name"]; ok && volumeName.(string) != "" {
						volumeNameStr := volumeName.(string)
						volumePrototype.Name = &volumeNameStr
					}

					// Set capacity
					if capacity, ok := volumeConfig["capacity"]; ok && capacity.(int) > 0 {
						capacityInt64 := int64(capacity.(int))
						volumePrototype.Capacity = &capacityInt64
					}

					// Set IOPS
					if iops, ok := volumeConfig["iops"]; ok && iops.(int) > 0 {
						iopsInt64 := int64(iops.(int))
						volumePrototype.Iops = &iopsInt64
					}

					// Set bandwidth
					if bandwidth, ok := volumeConfig["bandwidth"]; ok && bandwidth.(int) > 0 {
						bandwidthInt64 := int64(bandwidth.(int))
						volumePrototype.Bandwidth = &bandwidthInt64
					}

					// Set profile
					if profileList, ok := volumeConfig["profile"]; ok && len(profileList.([]interface{})) > 0 {
						profileConfig := profileList.([]interface{})[0].(map[string]interface{})
						if profileName, ok := profileConfig["name"]; ok && profileName.(string) != "" {
							profileNameStr := profileName.(string)
							volumePrototype.Profile = &vpcv1.VolumeProfileIdentity{
								Name: &profileNameStr,
							}
						}
					}

					// Set encryption key
					if encryptionKeyList, ok := volumeConfig["encryption_key"]; ok && len(encryptionKeyList.([]interface{})) > 0 {
						encryptionKeyConfig := encryptionKeyList.([]interface{})[0].(map[string]interface{})
						if crn, ok := encryptionKeyConfig["crn"]; ok && crn.(string) != "" {
							crnStr := crn.(string)
							volumePrototype.EncryptionKey = &vpcv1.EncryptionKeyIdentity{
								CRN: &crnStr,
							}
						}
					}

					// Set resource group
					if resourceGroup, ok := volumeConfig["resource_group"]; ok && resourceGroup.(string) != "" {
						resourceGroupStr := resourceGroup.(string)
						volumePrototype.ResourceGroup = &vpcv1.ResourceGroupIdentity{
							ID: &resourceGroupStr,
						}
					}

					// Set user tags
					if userTagsSet, ok := volumeConfig["user_tags"]; ok {
						userTags := userTagsSet.(*schema.Set)
						if userTags.Len() > 0 {
							tags := make([]string, userTags.Len())
							for i, tag := range userTags.List() {
								tags[i] = tag.(string)
							}
							volumePrototype.UserTags = tags
						}
					}

					// Set allowed use
					if allowedUseList, ok := volumeConfig["allowed_use"]; ok && len(allowedUseList.([]interface{})) > 0 {
						allowedUseConfig := allowedUseList.([]interface{})[0].(map[string]interface{})
						allowedUse := &vpcv1.VolumeAllowedUsePrototype{}

						if apiVersion, ok := allowedUseConfig["api_version"]; ok && apiVersion.(string) != "" {
							apiVersionStr := apiVersion.(string)
							allowedUse.ApiVersion = &apiVersionStr
						}
						if bareMetalServer, ok := allowedUseConfig["bare_metal_server"]; ok && bareMetalServer.(string) != "" {
							bareMetalServerStr := bareMetalServer.(string)
							allowedUse.BareMetalServer = &bareMetalServerStr
						}
						if instance, ok := allowedUseConfig["instance"]; ok && instance.(string) != "" {
							instanceStr := instance.(string)
							allowedUse.Instance = &instanceStr
						}
						volumePrototype.AllowedUse = allowedUse
					}
					bootvolumeAttachment.Volume = volumePrototype
					instanceReinitializePrototype = &vpcv1.InstanceReinitializePrototypeInstanceReinitializeBySnapshot{
						BootVolumeAttachment:  bootvolumeAttachment,
						UserData:              userData,
						Keys:                  keys,
						DefaultTrustedProfile: defaultTrustedProfile,
					}
					reinitializeOptions.InstanceReinitializePrototype = instanceReinitializePrototype
					log.Printf("[INFO] Reinitializing instance %s by source snapshot: %s", instanceId, snapshotIdStr)
				}
			}
		}
	} else {
		tfErr := flex.TerraformErrorf(fmt.Errorf("one of image or boot_volume_attachment must be provided"),
			"Missing boot source configuration", "ibm_is_instance_reinitialize", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Call the reinitialize API
	response, err := sess.ReinitializeInstanceWithContext(ctx, reinitializeOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ReinitializeInstanceWithContext failed: %s", err.Error()), "ibm_is_instance_reinitialize", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Check for success response code 204
	if response.StatusCode != 204 {
		tfErr := flex.TerraformErrorf(fmt.Errorf("reinitialization failed with status code: %d", response.StatusCode),
			"Unexpected response status", "ibm_is_instance_reinitialize", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(instanceId)

	// Wait for instance to be ready
	_, err = isWaitForInstanceReinitializationComplete(sess, instanceId, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceReinitializationComplete failed: %s", err.Error()),
			"ibm_is_instance_reinitialize", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Restart instance if it was stopped before reinitialization
	if stopInstanceIfRunning {
		err = resourceStartInstanceIfStopped(ctx, d, sess, instanceId)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("resourceStartInstanceIfStopped failed: %s", err.Error()),
				"ibm_is_instance_reinitialize", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMISInstanceReinitializeRead(ctx, d, meta)
}

func resourceIBMISInstanceReinitializeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	instanceId := d.Id()
	if instanceId == "" {
		return nil
	}

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_reinitialize", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Get instance status
	options := &vpcv1.GetInstanceOptions{
		ID: &instanceId,
	}
	instance, response, err := sess.GetInstanceWithContext(ctx, options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceWithContext failed: %s", err.Error()),
			"ibm_is_instance_reinitialize", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.Set("instance_id", instanceId)
	if instance.Status != nil {
		d.Set("status", *instance.Status)
	}

	return nil
}

func resourceIBMISInstanceReinitializeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Reinitialization is a one-time action, just remove from state
	d.SetId("")
	return nil
}

// resourceStopInstanceIfRunning stops the instance if it's in running state
func resourceStopInstanceIfRunning(ctx context.Context, d *schema.ResourceData, sess *vpcv1.VpcV1, instanceId string) (bool, error) {
	// Get current instance status
	getInstanceOptions := &vpcv1.GetInstanceOptions{
		ID: &instanceId,
	}
	instance, response, err := sess.GetInstanceWithContext(ctx, getInstanceOptions)
	if err != nil {
		return false, fmt.Errorf("error getting instance %s: %s\n%s", instanceId, err, response)
	}

	// Check if instance is running
	if instance.Status != nil && *instance.Status == "running" {
		log.Printf("[INFO] Stopping instance %s before reinitialization", instanceId)

		// Stop the instance
		stopInstanceOptions := &vpcv1.CreateInstanceActionOptions{
			InstanceID: &instanceId,
		}

		// Use hard stop for reinitialization
		forceStop := true
		stopInstanceOptions.Force = &forceStop
		typeStop := "hard"
		stopInstanceOptions.Type = &typeStop

		_, response, err := sess.CreateInstanceActionWithContext(ctx, stopInstanceOptions)
		if err != nil {
			return false, fmt.Errorf("error stopping instance %s: %s\n%s", instanceId, err, response)
		}

		// Wait for instance to be stopped
		_, err = waitForInstanceStopped(sess, instanceId, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return false, fmt.Errorf("error waiting for instance %s to stop: %s", instanceId, err)
		}

		log.Printf("[INFO] Instance %s stopped successfully", instanceId)
		return true, nil
	}

	return false, nil
}

// resourceStartInstanceIfStopped starts the instance if it was previously stopped
func resourceStartInstanceIfStopped(ctx context.Context, d *schema.ResourceData, sess *vpcv1.VpcV1, instanceId string) error {
	// Get current instance status
	getInstanceOptions := &vpcv1.GetInstanceOptions{
		ID: &instanceId,
	}
	instance, response, err := sess.GetInstanceWithContext(ctx, getInstanceOptions)
	if err != nil {
		return fmt.Errorf("error getting instance %s: %s\n%s", instanceId, err, response)
	}

	// Check if instance is stopped
	if instance.Status != nil && *instance.Status == "stopped" {
		log.Printf("[INFO] Starting instance %s after reinitialization", instanceId)

		// Start the instance
		startInstanceOptions := &vpcv1.CreateInstanceActionOptions{
			InstanceID: &instanceId,
		}

		_, response, err := sess.CreateInstanceActionWithContext(ctx, startInstanceOptions)
		if err != nil {
			return fmt.Errorf("error starting instance %s: %s\n%s", instanceId, err, response)
		}

		// Wait for instance to be running
		_, err = waitForInstanceRunning(sess, instanceId, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return fmt.Errorf("error waiting for instance %s to start: %s", instanceId, err)
		}

		log.Printf("[INFO] Instance %s started successfully", instanceId)
	}

	return nil
}

// waitForInstanceStopped waits for the instance to reach stopped state
func waitForInstanceStopped(client *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for instance (%s) to be stopped", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"running", "starting", "stopping", "pending"},
		Target:     []string{"stopped"},
		Refresh:    instanceStateRefreshFunc(client, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

// waitForInstanceRunning waits for the instance to reach running state
func waitForInstanceRunning(client *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for instance (%s) to be running", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"stopped", "starting", "pending"},
		Target:     []string{"running"},
		Refresh:    instanceStateRefreshFunc(client, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

// instanceStateRefreshFunc returns a refresh function for instance state
func instanceStateRefreshFunc(client *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getInstanceOptions := &vpcv1.GetInstanceOptions{
			ID: &id,
		}
		instance, response, err := client.GetInstance(getInstanceOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return nil, "deleted", nil
			}
			return nil, "", fmt.Errorf("[ERROR] Error getting instance: %s\n%s", err, response)
		}

		// Check for failed status
		if instance.Status != nil && *instance.Status == "failed" {
			statusReasons := instance.StatusReasons
			out, err := json.MarshalIndent(statusReasons, "", "    ")
			if err == nil {
				return instance, *instance.Status, fmt.Errorf("[ERROR] Instance (%s) went into failed state: %s", id, string(out))
			}
			return instance, *instance.Status, fmt.Errorf("[ERROR] Instance (%s) went into failed state", id)
		}

		if instance.Status != nil {
			return instance, *instance.Status, nil
		}
		return instance, "pending", nil
	}
}

// isWaitForInstanceReinitializationComplete waits for instance reinitialization to complete
func isWaitForInstanceReinitializationComplete(client *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for instance (%s) to complete reinitialization", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"pending", "starting", "stopping", "updating", "reinitializing"},
		Target:     []string{"running", "stopped"},
		Refresh:    instanceStateRefreshFunc(client, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}
