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
	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	isVolumeProfile       = "name"
	isVolumeProfileFamily = "family"
)

func dataSourceIBMISVolumeProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISVolumeProfileRead,

		Schema: map[string]*schema.Schema{

			isVolumeProfile: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Volume profile name",
			},

			isVolumeProfileFamily: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Volume profile family",
			},
		},
	}
}

func dataSourceIBMISVolumeProfileRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	name := d.Get(isVolumeProfile).(string)
	if userDetails.generation == 1 {
		err := classicVolumeProfileGet(d, meta, name)
		if err != nil {
			return err
		}
	} else {
		err := volumeProfileGet(d, meta, name)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicVolumeProfileGet(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getVolumeProfileOptions := &vpcclassicv1.GetVolumeProfileOptions{
		Name: &name,
	}
	profile, _, err := sess.GetVolumeProfile(getVolumeProfileOptions)
	if err != nil {
		return err
	}
	// For lack of anything better, compose our id from profile name.
	d.SetId(*profile.Name)
	d.Set(isVolumeProfile, *profile.Name)
	d.Set(isVolumeProfileFamily, *profile.Family)
	return nil
}

func volumeProfileGet(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getVolumeProfileOptions := &vpcv1.GetVolumeProfileOptions{
		Name: &name,
	}
	profile, _, err := sess.GetVolumeProfile(getVolumeProfileOptions)
	if err != nil {
		return err
	}
	// For lack of anything better, compose our id from profile name.
	d.SetId(*profile.Name)
	d.Set(isVolumeProfile, *profile.Name)
	d.Set(isVolumeProfileFamily, *profile.Family)
	return nil
}
