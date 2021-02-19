/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

package ibm

import (
	"fmt"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
