package ibm

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/compute"
)

const (
	isInstanceProfileName       = "name"
	isInstanceProfileFamily     = "family"
	isInstanceProfileGeneration = "generation"
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

			isInstanceProfileGeneration: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMISInstanceProfileRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	instanceC := compute.NewInstanceClient(sess)
	profile, err := instanceC.GetProfile(d.Get(isInstanceProfileName).(string))
	if err != nil {
		return err
	}
	// For lack of anything better, compose our id from region name + zone name.
	id := profile.Name
	d.SetId(id)
	d.Set(isInstanceProfileName, profile.Name)
	d.Set(isInstanceProfileFamily, profile.Family)
	d.Set(isInstanceProfileGeneration, profile.Generation)
	return nil
}
