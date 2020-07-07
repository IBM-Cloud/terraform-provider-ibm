package ibm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

const (
	isInstanceProfileName   = "name"
	isInstanceProfileFamily = "family"
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
	return nil
}
