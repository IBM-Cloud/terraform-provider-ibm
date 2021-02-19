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
	isRegionEndpoint = "endpoint"
	isRegionName     = "name"
	isRegionStatus   = "status"
)

func dataSourceIBMISRegion() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISRegionRead,

		Schema: map[string]*schema.Schema{

			isRegionEndpoint: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isRegionName: {
				Type:     schema.TypeString,
				Required: true,
			},

			isRegionStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMISRegionRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	name := d.Get("name").(string)
	if userDetails.generation == 1 {
		err := classicRegionGet(d, meta, name)
		if err != nil {
			return err
		}
	} else {
		err := regionGet(d, meta, name)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicRegionGet(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getRegionOptions := &vpcclassicv1.GetRegionOptions{
		Name: &name,
	}
	region, _, err := sess.GetRegion(getRegionOptions)
	if err != nil {
		return err
	}
	// For lack of anything better, compose our id from region name.
	d.SetId(*region.Name)
	d.Set(isRegionEndpoint, *region.Endpoint)
	d.Set(isRegionName, *region.Name)
	d.Set(isRegionStatus, *region.Status)
	return nil
}

func regionGet(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getRegionOptions := &vpcv1.GetRegionOptions{
		Name: &name,
	}
	region, _, err := sess.GetRegion(getRegionOptions)
	if err != nil {
		return err
	}
	// For lack of anything better, compose our id from region name.
	d.SetId(*region.Name)
	d.Set(isRegionEndpoint, *region.Endpoint)
	d.Set(isRegionName, *region.Name)
	d.Set(isRegionStatus, *region.Status)
	return nil
}
