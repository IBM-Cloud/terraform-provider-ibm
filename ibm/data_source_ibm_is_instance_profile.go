// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isInstanceProfileName         = "name"
	isInstanceProfileFamily       = "family"
	isInstanceProfileArchitecture = "architecture"
)

func dataSourceIBMISInstanceProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISInstanceProfileRead,

		Schema: map[string]*schema.Schema{

			isInstanceProfileName: {
				Type:     schema.TypeString,
				Required: true,
			},

			isInstanceProfileFamily: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isInstanceProfileArchitecture: {
				Type:     schema.TypeString,
				Computed: true,
			},
			isInstanceDisks: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of the instance profile's disks.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"quantity": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The value for this profile field.",
									},
									"default": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The default value for this profile field.",
									},
									"max": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum value for this profile field.",
									},
									"min": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum value for this profile field.",
									},
									"step": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The increment step value for this profile field.",
									},
									"values": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The permitted values for this profile field.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
								},
							},
						},
						"size": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The value for this profile field.",
									},
									"default": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The default value for this profile field.",
									},
									"max": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum value for this profile field.",
									},
									"min": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum value for this profile field.",
									},
									"step": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The increment step value for this profile field.",
									},
									"values": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The permitted values for this profile field.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
								},
							},
						},
						"supported_interface_types": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The disk interface used for attaching the disk.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type for this profile field.",
									},
									"values": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The supported disk interfaces used for attaching the disk.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
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

func dataSourceIBMISInstanceProfileRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	name := d.Get(isInstanceProfileName).(string)
	if userDetails.generation == 1 {
		err := classicInstanceProfileGet(d, meta, name)
		if err != nil {
			return err
		}
	} else {
		err := instanceProfileGet(d, meta, name)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicInstanceProfileGet(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getInstanceProfileOptions := &vpcclassicv1.GetInstanceProfileOptions{
		Name: &name,
	}
	profile, _, err := sess.GetInstanceProfile(getInstanceProfileOptions)
	if err != nil {
		return err
	}
	// For lack of anything better, compose our id from profile name.
	d.SetId(*profile.Name)
	d.Set(isInstanceProfileName, *profile.Name)
	d.Set(isInstanceProfileFamily, *profile.Family)
	return nil
}

func instanceProfileGet(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getInstanceProfileOptions := &vpcv1.GetInstanceProfileOptions{
		Name: &name,
	}
	profile, _, err := sess.GetInstanceProfile(getInstanceProfileOptions)
	if err != nil {
		return err
	}
	// For lack of anything better, compose our id from profile name.
	d.SetId(*profile.Name)
	d.Set(isInstanceProfileName, *profile.Name)
	d.Set(isInstanceProfileFamily, *profile.Family)
	if profile.OsArchitecture != nil && profile.OsArchitecture.Default != nil {
		d.Set(isInstanceProfileArchitecture, *profile.OsArchitecture.Default)
	}

	if profile.Disks != nil {
		err = d.Set(isInstanceDisks, dataSourceInstanceProfileFlattenDisks(profile.Disks))
	}

	return nil
}

func dataSourceInstanceProfileFlattenDisks(result []vpcv1.InstanceProfileDisk) (disks []map[string]interface{}) {
	for _, disksItem := range result {
		disks = append(disks, dataSourceInstanceProfileDisksToMap(disksItem))
	}

	return disks
}

func dataSourceInstanceProfileDisksToMap(disksItem vpcv1.InstanceProfileDisk) (disksMap map[string]interface{}) {
	disksMap = map[string]interface{}{}

	if disksItem.Quantity != nil {
		quantityList := []map[string]interface{}{}
		quantityMap := dataSourceInstanceProfileDisksQuantityToMap(*disksItem.Quantity.(*vpcv1.InstanceProfileDiskQuantity))
		quantityList = append(quantityList, quantityMap)
		disksMap["quantity"] = quantityList
	}
	if disksItem.Size != nil {
		sizeList := []map[string]interface{}{}
		sizeMap := dataSourceInstanceProfileDisksSizeToMap(*disksItem.Size.(*vpcv1.InstanceProfileDiskSize))
		sizeList = append(sizeList, sizeMap)
		disksMap["size"] = sizeList
	}
	if disksItem.SupportedInterfaceTypes != nil {
		supportedInterfaceTypesList := []map[string]interface{}{}
		supportedInterfaceTypesMap := dataSourceInstanceProfileDisksSupportedInterfaceTypesToMap(*disksItem.SupportedInterfaceTypes)
		supportedInterfaceTypesList = append(supportedInterfaceTypesList, supportedInterfaceTypesMap)
		disksMap["supported_interface_types"] = supportedInterfaceTypesList
	}

	return disksMap
}

func dataSourceInstanceProfileDisksQuantityToMap(quantityItem vpcv1.InstanceProfileDiskQuantity) (quantityMap map[string]interface{}) {
	quantityMap = map[string]interface{}{}

	if quantityItem.Type != nil {
		quantityMap["type"] = quantityItem.Type
	}
	if quantityItem.Value != nil {
		quantityMap["value"] = quantityItem.Value
	}
	if quantityItem.Default != nil {
		quantityMap["default"] = quantityItem.Default
	}
	if quantityItem.Max != nil {
		quantityMap["max"] = quantityItem.Max
	}
	if quantityItem.Min != nil {
		quantityMap["min"] = quantityItem.Min
	}
	if quantityItem.Step != nil {
		quantityMap["step"] = quantityItem.Step
	}
	if quantityItem.Values != nil {
		quantityMap["values"] = quantityItem.Values
	}

	return quantityMap
}

func dataSourceInstanceProfileDisksSizeToMap(sizeItem vpcv1.InstanceProfileDiskSize) (sizeMap map[string]interface{}) {
	sizeMap = map[string]interface{}{}

	if sizeItem.Type != nil {
		sizeMap["type"] = sizeItem.Type
	}
	if sizeItem.Value != nil {
		sizeMap["value"] = sizeItem.Value
	}
	if sizeItem.Default != nil {
		sizeMap["default"] = sizeItem.Default
	}
	if sizeItem.Max != nil {
		sizeMap["max"] = sizeItem.Max
	}
	if sizeItem.Min != nil {
		sizeMap["min"] = sizeItem.Min
	}
	if sizeItem.Step != nil {
		sizeMap["step"] = sizeItem.Step
	}
	if sizeItem.Values != nil {
		sizeMap["values"] = sizeItem.Values
	}

	return sizeMap
}

func dataSourceInstanceProfileDisksSupportedInterfaceTypesToMap(supportedInterfaceTypesItem vpcv1.InstanceProfileDiskSupportedInterfaces) (supportedInterfaceTypesMap map[string]interface{}) {
	supportedInterfaceTypesMap = map[string]interface{}{}

	if supportedInterfaceTypesItem.Default != nil {
		supportedInterfaceTypesMap["default"] = supportedInterfaceTypesItem.Default
	}
	if supportedInterfaceTypesItem.Type != nil {
		supportedInterfaceTypesMap["type"] = supportedInterfaceTypesItem.Type
	}
	if supportedInterfaceTypesItem.Values != nil {
		supportedInterfaceTypesMap["values"] = supportedInterfaceTypesItem.Values
	}

	return supportedInterfaceTypesMap
}
