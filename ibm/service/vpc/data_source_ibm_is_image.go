// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isImageCatalogOffering         = "catalog_offering"
	isImageCatalogOfferingManaged  = "managed"
	isImageCatalogOfferingVersion  = "version"
	isImageCatalogOfferingCrn      = "crn"
	isImageCatalogOfferingDeleted  = "deleted"
	isImageCatalogOfferingMoreInfo = "more_info"
)

func DataSourceIBMISImage() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISImageRead,

		Schema: map[string]*schema.Schema{

			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"identifier", "name"},
				Description:  "Image name",
			},

			"identifier": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"identifier", "name"},
				Description:  "Image id",
			},

			"visibility": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"public", "private"}),
				Description:  "Whether the image is publicly visible or private to the account",
			},
			"resource_group": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The resource group for this IPsec policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this resource group.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this resource group.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this resource group.",
						},
					},
				},
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of this image",
			},
			"status_reasons": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current status (if any).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the status reason.",
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the status reason.",
						},
						"more_info": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about this status reason.",
						},
					},
				},
			},
			"operating_system": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isOperatingSystemAllowUserImageCreation: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Users may create new images with this operating system",
						},
						"architecture": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The operating system architecture",
						},
						"dedicated_host_only": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Images with this operating system can only be used on dedicated hosts or dedicated host groups",
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A unique, display-friendly name for the operating system",
						},
						"family": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The software family for this operating system",
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this operating system",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The globally unique name for this operating system",
						},
						"vendor": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vendor of the operating system",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The major release version of this operating system",
						},
						isOperatingSystemUserDataFormat: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user data format for this image",
						},
					},
				},
			},
			"os": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Image Operating system",
			},
			isImageUserDataFormat: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user data format for this image",
			},
			"architecture": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The operating system architecture",
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this image",
			},
			isImageCheckSum: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SHA256 Checksum for this image",
			},
			isImageEncryptionKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of the Key Protect Root Key or Hyper Protect Crypto Service Root Key for this resource",
			},
			isImageEncryption: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of encryption used on the image",
			},
			"remote": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "If present, this property indicates that the resource associated with this reference is remote and therefore may not be directly retrievable.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates that the referenced resource is remote to this account, and identifies the owning account.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this resource group.",
									},
									"resource_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type.",
									},
								},
							},
						},
					},
				},
			},
			"source_volume": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Source volume id of the image",
			},
			isImageCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the image was created",
			},
			isImageDeprecationAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The deprecation date and time (UTC) for this image. If absent, no deprecation date and time has been set.",
			},
			isImageObsolescenceAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The obsolescence date and time (UTC) for this image. If absent, no obsolescence date and time has been set.",
			},
			isImageCatalogOffering: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isImageCatalogOfferingManaged: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether this image is managed as part of a catalog offering. A managed image can be provisioned using its catalog offering CRN or catalog offering version CRN.",
						},
						isImageCatalogOfferingVersion: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The catalog offering version associated with this image.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// isImageCatalogOfferingDeleted: {
									// 	Type:        schema.TypeList,
									// 	Computed:    true,
									// 	Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
									// 	Elem: &schema.Resource{
									// 		Schema: map[string]*schema.Schema{
									// 			isImageCatalogOfferingMoreInfo: {
									// 				Type:        schema.TypeString,
									// 				Computed:    true,
									// 				Description: "Link to documentation about deleted resources.",
									// 			},
									// 		},
									// 	},
									// },
									isImageCatalogOfferingCrn: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this version of the IBM Cloud catalog offering.",
									},
								},
							},
						},
					},
				},
			},
			"allowed_use": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The usage constraints to match against the requested instance or bare metal server properties to determine compatibility.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"api_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The API version with which to evaluate the expressions.",
						},
						"bare_metal_server": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expression that must be satisfied by the properties of a bare metal server provisioned using this image.",
						},
						"instance": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expression that must be satisfied by the properties of a virtual server instance provisioned using this image.",
						},
					},
				},
			},
			isImageAccessTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access tags",
			},
		},
	}
}

func dataSourceIBMISImageRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	name := d.Get("name").(string)
	identifier := d.Get("identifier").(string)
	var visibility string
	if v, ok := d.GetOk("visibility"); ok {
		visibility = v.(string)
	}
	if name != "" {
		err := imageGetByName(context, d, meta, name, visibility)
		if err != nil {
			return err
		}
	} else if identifier != "" {
		err := imageGetById(context, d, meta, identifier)
		if err != nil {
			return err
		}
	}

	return nil
}

func imageGetByName(context context.Context, d *schema.ResourceData, meta interface{}, name, visibility string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_image", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	listImagesOptions := &vpcv1.ListImagesOptions{
		Name: &name,
	}

	if visibility != "" {
		listImagesOptions.Visibility = &visibility
	}
	availableImages, _, err := sess.ListImagesWithContext(context, listImagesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListImagesWithContext failed: %s", err.Error()), "(Data) ibm_is_image", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	allrecs := availableImages.Images

	if len(allrecs) == 0 {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("No image found with name: %s", name), "(Data) ibm_is_image", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	image := allrecs[0]
	d.SetId(*image.ID)
	if err = d.Set("user_data_format", image.UserDataFormat); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting user_data_format: %s", err), "(Data) ibm_is_image", "read", "set-user_data_format").GetDiag()
	}
	if err = d.Set("status", image.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_image", "read", "set-status").GetDiag()
	}
	if *image.Status == "deprecated" {
		fmt.Printf("[WARN] Given image %s is deprecated and soon will be obsolete.", name)
	}
	if len(image.StatusReasons) > 0 {
		if err = d.Set("status_reasons", dataSourceIBMIsImageFlattenStatusReasons(image.StatusReasons)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status_reasons: %s", err), "(Data) ibm_is_image", "read", "set-status_reasons").GetDiag()
		}
	}
	if err = d.Set("name", image.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_image", "read", "set-name").GetDiag()
	}
	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *image.CRN, "", isImageAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource image (%s) access tags: %s", d.Id(), err)
	}
	if err = d.Set("access_tags", accesstags); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_tags: %s", err), "(Data) ibm_is_image", "read", "set-access_tags").GetDiag()
	}
	if err = d.Set("visibility", image.Visibility); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting visibility: %s", err), "(Data) ibm_is_image", "read", "set-visibility").GetDiag()
	}

	if image.OperatingSystem != nil {
		operatingSystemList := []map[string]interface{}{}
		operatingSystemMap := dataSourceIBMISImageOperatingSystemToMap(*image.OperatingSystem)
		operatingSystemList = append(operatingSystemList, operatingSystemMap)
		if err = d.Set("operating_system", operatingSystemList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting operating_system: %s", err), "(Data) ibm_is_image", "read", "set-operating_system").GetDiag()
		}
	}
	if image.ResourceGroup != nil {
		resourceGroupList := []map[string]interface{}{}
		resourceGroupMap := dataSourceImageResourceGroupToMap(*image.ResourceGroup)
		resourceGroupList = append(resourceGroupList, resourceGroupMap)
		if err = d.Set("resource_group", resourceGroupList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_image", "read", "set-resource_group").GetDiag()
		}
	}
	if image.Remote != nil {
		imageRemoteList := []map[string]interface{}{}
		imageRemoteMap, err := dataSourceImageRemote(image)
		if err != nil {
			if err != nil {
				tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_image", "read", "initialize-client")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		}
		imageRemoteList = append(imageRemoteList, imageRemoteMap)
		if err = d.Set(isImageRemote, imageRemoteList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting remote: %s", err), "(Data) ibm_is_image", "read", "set-remote").GetDiag()
		}
	}

	if err = d.Set("os", image.OperatingSystem.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting os: %s", err), "(Data) ibm_is_image", "read", "set-os").GetDiag()
	}
	if err = d.Set("architecture", image.OperatingSystem.Architecture); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting architecture: %s", err), "(Data) ibm_is_image", "read", "set-architecture").GetDiag()
	}
	if err = d.Set("crn", image.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_image", "read", "set-crn").GetDiag()
	}
	if err = d.Set("encryption", image.Encryption); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting encryption: %s", err), "(Data) ibm_is_image", "read", "set-encryption").GetDiag()
	}
	if image.EncryptionKey != nil {
		if err = d.Set("encryption_key", *image.EncryptionKey.CRN); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting encryption_key: %s", err), "(Data) ibm_is_image", "read", "set-encryption_key").GetDiag()
		}
	}
	if image.File != nil && image.File.Checksums != nil {
		if err = d.Set("checksum", *image.File.Checksums.Sha256); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting checksum: %s", err), "(Data) ibm_is_image", "read", "set-checksum").GetDiag()
		}
	}
	if image.SourceVolume != nil {
		if err = d.Set("source_volume", *image.SourceVolume.ID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting source_volume: %s", err), "(Data) ibm_is_image", "read", "set-source_volume").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(image.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_image", "read", "set-created_at").GetDiag()
	}
	if !core.IsNil(image.DeprecationAt) {
		if err = d.Set("deprecation_at", flex.DateTimeToString(image.DeprecationAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting deprecation_at: %s", err), "(Data) ibm_is_image", "read", "set-deprecation_at").GetDiag()
		}
	}
	if !core.IsNil(image.ObsolescenceAt) {
		if err = d.Set("obsolescence_at", flex.DateTimeToString(image.ObsolescenceAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting obsolescence_at: %s", err), "(Data) ibm_is_image", "read", "set-obsolescence_at").GetDiag()
		}
	}
	if image.CatalogOffering != nil {
		catalogOfferingList := []map[string]interface{}{}
		catalogOfferingMap := dataSourceImageCollectionCatalogOfferingToMap(*image.CatalogOffering)
		catalogOfferingList = append(catalogOfferingList, catalogOfferingMap)
		if err = d.Set("catalog_offering", catalogOfferingList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting catalog_offering: %s", err), "(Data) ibm_is_image", "read", "set-catalog_offering").GetDiag()
		}
	}
	allowedUse := []map[string]interface{}{}
	if image.AllowedUse != nil {
		modelMap, err := DataSourceIBMIsImageAllowedUseToMap(image.AllowedUse)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_image", "read")
			log.Println(tfErr.GetDiag())
		}
		allowedUse = append(allowedUse, modelMap)
	}
	if err = d.Set("allowed_use", allowedUse); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting allowed_use: %s", err), "(Data) ibm_is_image", "read")
		log.Println(tfErr.GetDiag())
	}
	return nil

}
func imageGetById(context context.Context, d *schema.ResourceData, meta interface{}, identifier string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_image", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getImageOptions := &vpcv1.GetImageOptions{
		ID: &identifier,
	}

	image, _, err := sess.GetImageWithContext(context, getImageOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetImageWithContext failed: %s", err.Error()), "(Data) ibm_is_image", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*image.ID)
	if err = d.Set("status", image.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_image", "read", "set-status").GetDiag()
	}
	if *image.Status == "deprecated" {
		fmt.Printf("[WARN] Given image %s is deprecated and soon will be obsolete.", name)
	}

	if image.Remote != nil {
		imageRemoteList := []map[string]interface{}{}
		imageRemoteMap, err := dataSourceImageRemote(*image)
		if err != nil {
			tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_legacy_vendor_images", "read", "initialize-client")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		imageRemoteList = append(imageRemoteList, imageRemoteMap)
		if err = d.Set(isImageRemote, imageRemoteList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting remote: %s", err), "(Data) ibm_is_image", "read", "set-remote").GetDiag()
		}
	}

	if len(image.StatusReasons) > 0 {
		if err = d.Set("status_reasons", dataSourceIBMIsImageFlattenStatusReasons(image.StatusReasons)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status_reasons: %s", err), "(Data) ibm_is_image", "read", "set-status_reasons").GetDiag()
		}
	}
	if err = d.Set("name", image.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_image", "read", "set-name").GetDiag()
	}
	if err = d.Set("user_data_format", image.UserDataFormat); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting user_data_format: %s", err), "(Data) ibm_is_image", "read", "set-user_data_format").GetDiag()
	}
	if err = d.Set("visibility", image.Visibility); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting visibility: %s", err), "(Data) ibm_is_image", "read", "set-visibility").GetDiag()
	}
	if image.OperatingSystem != nil {
		operatingSystemList := []map[string]interface{}{}
		operatingSystemMap := dataSourceIBMISImageOperatingSystemToMap(*image.OperatingSystem)
		operatingSystemList = append(operatingSystemList, operatingSystemMap)
		if err = d.Set("operating_system", operatingSystemList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting operating_system: %s", err), "(Data) ibm_is_image", "read", "set-operating_system").GetDiag()
		}
	}
	if image.ResourceGroup != nil {
		resourceGroupList := []map[string]interface{}{}
		resourceGroupMap := dataSourceImageResourceGroupToMap(*image.ResourceGroup)
		resourceGroupList = append(resourceGroupList, resourceGroupMap)
		if err = d.Set("resource_group", resourceGroupList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_image", "read", "set-resource_group").GetDiag()
		}
	}
	if err = d.Set("os", image.OperatingSystem.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting os: %s", err), "(Data) ibm_is_image", "read", "set-os").GetDiag()
	}
	if err = d.Set("architecture", image.OperatingSystem.Architecture); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting architecture: %s", err), "(Data) ibm_is_image", "read", "set-architecture").GetDiag()
	}
	if err = d.Set("crn", image.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_image", "read", "set-crn").GetDiag()
	}

	if err = d.Set("encryption", image.Encryption); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting encryption: %s", err), "(Data) ibm_is_image", "read", "set-encryption").GetDiag()
	}
	if image.EncryptionKey != nil {
		if err = d.Set("encryption_key", *image.EncryptionKey.CRN); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting encryption_key: %s", err), "(Data) ibm_is_image", "read", "set-encryption_key").GetDiag()
		}
	}
	if image.File != nil && image.File.Checksums != nil {
		if err = d.Set("checksum", *image.File.Checksums.Sha256); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting checksum: %s", err), "(Data) ibm_is_image", "read", "set-checksum").GetDiag()
		}
	}
	if image.SourceVolume != nil {
		if err = d.Set("source_volume", *image.SourceVolume.ID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting source_volume: %s", err), "(Data) ibm_is_image", "read", "set-source_volume").GetDiag()
		}
	}
	if image.CatalogOffering != nil {
		catalogOfferingList := []map[string]interface{}{}
		catalogOfferingMap := dataSourceImageCollectionCatalogOfferingToMap(*image.CatalogOffering)
		catalogOfferingList = append(catalogOfferingList, catalogOfferingMap)
		if err = d.Set("catalog_offering", catalogOfferingList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting catalog_offering: %s", err), "(Data) ibm_is_image", "read", "set-catalog_offering").GetDiag()
		}
	}
	allowedUse := []map[string]interface{}{}
	if image.AllowedUse != nil {
		modelMap, err := DataSourceIBMIsImageAllowedUseToMap(image.AllowedUse)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_image", "read")
			log.Println(tfErr.GetDiag())
		}
		allowedUse = append(allowedUse, modelMap)
	}
	if err = d.Set("allowed_use", allowedUse); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting allowed_use: %s", err), "(Data) ibm_is_image", "read")
		log.Println(tfErr.GetDiag())
	}
	return nil
}

func dataSourceIBMISImageOperatingSystemToMap(operatingSystemItem vpcv1.OperatingSystem) (operatingSystemMap map[string]interface{}) {
	operatingSystemMap = map[string]interface{}{}
	if operatingSystemItem.AllowUserImageCreation != nil {
		operatingSystemMap[isOperatingSystemAllowUserImageCreation] = operatingSystemItem.AllowUserImageCreation
	}
	if operatingSystemItem.Architecture != nil {
		operatingSystemMap["architecture"] = operatingSystemItem.Architecture
	}
	if operatingSystemItem.DedicatedHostOnly != nil {
		operatingSystemMap["dedicated_host_only"] = operatingSystemItem.DedicatedHostOnly
	}
	if operatingSystemItem.DisplayName != nil {
		operatingSystemMap["display_name"] = operatingSystemItem.DisplayName
	}
	if operatingSystemItem.Family != nil {
		operatingSystemMap["family"] = operatingSystemItem.Family
	}
	if operatingSystemItem.Href != nil {
		operatingSystemMap["href"] = operatingSystemItem.Href
	}
	if operatingSystemItem.Name != nil {
		operatingSystemMap["name"] = operatingSystemItem.Name
	}
	if operatingSystemItem.Vendor != nil {
		operatingSystemMap["vendor"] = operatingSystemItem.Vendor
	}
	if operatingSystemItem.Version != nil {
		operatingSystemMap["version"] = operatingSystemItem.Version
	}
	if operatingSystemItem.UserDataFormat != nil {
		operatingSystemMap[isOperatingSystemUserDataFormat] = operatingSystemItem.UserDataFormat
	}
	return operatingSystemMap
}

func dataSourceImageCollectionCatalogOfferingToMap(imageCatalogOfferingItem vpcv1.ImageCatalogOffering) (imageCatalogOfferingMap map[string]interface{}) {
	imageCatalogOfferingMap = map[string]interface{}{}
	if imageCatalogOfferingItem.Managed != nil {
		imageCatalogOfferingMap[isImageCatalogOfferingManaged] = imageCatalogOfferingItem.Managed
	}
	if imageCatalogOfferingItem.Version != nil {
		imageCatalogOfferingVersionList := []map[string]interface{}{}
		imageCatalogOfferingVersionMap := map[string]interface{}{}
		imageCatalogOfferingVersionMap[isImageCatalogOfferingCrn] = imageCatalogOfferingItem.Version.CRN

		// if imageCatalogOfferingItem.Version.Deleted != nil {
		// 	imageCatalogOfferingVersionDeletedList := []map[string]interface{}{}
		// 	imageCatalogOfferingVersionDeletedMap := map[string]interface{}{}
		// 	imageCatalogOfferingVersionDeletedMap[isImageCatalogOfferingMoreInfo] = imageCatalogOfferingItem.Version.Deleted.MoreInfo
		// 	imageCatalogOfferingVersionDeletedList = append(imageCatalogOfferingVersionDeletedList, imageCatalogOfferingVersionDeletedMap)
		// 	imageCatalogOfferingVersionMap[isImageCatalogOfferingDeleted] = imageCatalogOfferingVersionDeletedList
		// }
		imageCatalogOfferingVersionList = append(imageCatalogOfferingVersionList, imageCatalogOfferingVersionMap)
		imageCatalogOfferingMap[isImageCatalogOfferingVersion] = imageCatalogOfferingVersionList
	}

	return imageCatalogOfferingMap
}

func dataSourceImageRemote(imageRemote vpcv1.Image) (map[string]interface{}, error) {
	if imageRemote.Remote == nil || imageRemote.Remote.Account == nil {
		return nil, nil
	}

	accountMap := map[string]interface{}{}

	if imageRemote.Remote.Account.ID != nil {
		accountMap["id"] = *imageRemote.Remote.Account.ID
	}
	if imageRemote.Remote.Account.ResourceType != nil {
		accountMap["resource_type"] = *imageRemote.Remote.Account.ResourceType
	}

	remoteMap := map[string]interface{}{
		"account": []interface{}{accountMap},
	}

	return remoteMap, nil
}

func dataSourceIBMIsImageFlattenStatusReasons(result []vpcv1.ImageStatusReason) (statusReasons []map[string]interface{}) {
	for _, statusReasonsItem := range result {
		statusReasons = append(statusReasons, dataSourceIBMIsImageStatusReasonToMap(&statusReasonsItem))
	}

	return statusReasons
}

func dataSourceIBMIsImageStatusReasonToMap(model *vpcv1.ImageStatusReason) map[string]interface{} {
	modelMap := make(map[string]interface{})
	if model.Code != nil {
		modelMap["code"] = *model.Code
	}
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap
}
func dataSourceImageResourceGroupToMap(resourceGroupItem vpcv1.ResourceGroupReference) (resourceGroupMap map[string]interface{}) {
	resourceGroupMap = map[string]interface{}{}

	if resourceGroupItem.Href != nil {
		resourceGroupMap["href"] = resourceGroupItem.Href
	}
	if resourceGroupItem.ID != nil {
		resourceGroupMap["id"] = resourceGroupItem.ID
	}
	if resourceGroupItem.Name != nil {
		resourceGroupMap["name"] = resourceGroupItem.Name
	}

	return resourceGroupMap
}

func DataSourceIBMIsImageAllowedUseToMap(model *vpcv1.ImageAllowedUse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BareMetalServer != nil {
		modelMap["bare_metal_server"] = *model.BareMetalServer
	}
	if model.Instance != nil {
		modelMap["instance"] = *model.Instance
	}
	if model.ApiVersion != nil {
		modelMap["api_version"] = *model.ApiVersion
	}
	return modelMap, nil
}
