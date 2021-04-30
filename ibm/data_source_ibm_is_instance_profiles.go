// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isInstanceProfiles = "profiles"
)

func dataSourceIBMISInstanceProfiles() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISInstanceProfilesRead,

		Schema: map[string]*schema.Schema{

			isInstanceProfiles: {
				Type:        schema.TypeList,
				Description: "List of instance profile maps",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"family": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"architecture": {
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
				},
			},
		},
	}
}

func dataSourceIBMISInstanceProfilesRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	if userDetails.generation == 1 {
		err := classicInstanceProfilesList(d, meta)
		if err != nil {
			return err
		}
	} else {
		err := instanceProfilesList(d, meta)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicInstanceProfilesList(d *schema.ResourceData, meta interface{}) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	start := ""
	allrecs := []vpcclassicv1.InstanceProfile{}
	for {
		listInstanceProfilesOptions := &vpcclassicv1.ListInstanceProfilesOptions{}
		if start != "" {
			listInstanceProfilesOptions.Start = &start
		}
		availableProfiles, response, err := sess.ListInstanceProfiles(listInstanceProfilesOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching Instance Profiles %s\n%s", err, response)
		}
		start = GetNext(availableProfiles.Next)
		allrecs = append(allrecs, availableProfiles.Profiles...)
		if start == "" {
			break
		}
	}
	profilesInfo := make([]map[string]interface{}, 0)
	for _, profile := range allrecs {

		l := map[string]interface{}{
			"name":   *profile.Name,
			"family": *profile.Family,
		}
		profilesInfo = append(profilesInfo, l)
	}
	d.SetId(dataSourceIBMISInstanceProfilesID(d))
	d.Set(isInstanceProfiles, profilesInfo)
	return nil
}

func instanceProfilesList(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	listInstanceProfilesOptions := &vpcv1.ListInstanceProfilesOptions{}
	availableProfiles, response, err := sess.ListInstanceProfiles(listInstanceProfilesOptions)
	if err != nil {
		return fmt.Errorf("Error Fetching Instance Profiles %s\n%s", err, response)
	}
	profilesInfo := make([]map[string]interface{}, 0)
	for _, profile := range availableProfiles.Profiles {

		l := map[string]interface{}{
			"name":   *profile.Name,
			"family": *profile.Family,
		}
		if profile.OsArchitecture != nil && profile.OsArchitecture.Default != nil {
			l["architecture"] = *profile.OsArchitecture.Default
		}
		if profile.Disks != nil {
			l[isInstanceDisks] = dataSourceInstanceProfileFlattenDisks(profile.Disks)
			if err != nil {
				return fmt.Errorf("Error setting disks %s", err)
			}
		}
		profilesInfo = append(profilesInfo, l)
	}
	d.SetId(dataSourceIBMISInstanceProfilesID(d))
	d.Set(isInstanceProfiles, profilesInfo)
	return nil
}

// dataSourceIBMISInstanceProfilesID returns a reasonable ID for a Instance Profile list.
func dataSourceIBMISInstanceProfilesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
