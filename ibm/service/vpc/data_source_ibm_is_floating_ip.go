// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	floatingIPName    = "name"
	floatingIPAddress = "address"
	floatingIPStatus  = "status"
	floatingIPZone    = "zone"
	floatingIPTarget  = "target"
	floatingIPTags    = "tags"
	floatingIPCRN     = "crn"
)

func DataSourceIBMISFloatingIP() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISFloatingIPRead,

		Schema: map[string]*schema.Schema{

			floatingIPName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the floating IP",
			},

			floatingIPAddress: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Floating IP address",
			},

			floatingIPStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Floating IP status",
			},

			floatingIPZone: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone name",
			},

			floatingIPTarget: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Target info",
			},

			floatingIPCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Floating IP crn",
			},

			floatingIPTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "Floating IP tags",
			},
		},
	}
}

func dataSourceIBMISFloatingIPRead(d *schema.ResourceData, meta interface{}) error {
	floatingIPName := d.Get(isFloatingIPName).(string)
	err := floatingIPGet(d, meta, floatingIPName)
	if err != nil {
		return err
	}
	return nil
}

func floatingIPGet(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	start := ""
	allFloatingIPs := []vpcv1.FloatingIP{}
	for {
		floatingIPOptions := &vpcv1.ListFloatingIpsOptions{}
		if start != "" {
			floatingIPOptions.Start = &start
		}
		floatingIPs, response, err := sess.ListFloatingIps(floatingIPOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error Fetching floating IPs %s\n%s", err, response)
		}
		start = flex.GetNext(floatingIPs.Next)
		allFloatingIPs = append(allFloatingIPs, floatingIPs.FloatingIps...)
		if start == "" {
			break
		}
	}

	for _, ip := range allFloatingIPs {
		if *ip.Name == name {

			d.Set(floatingIPName, *ip.Name)
			d.Set(floatingIPAddress, *ip.Address)
			d.Set(floatingIPStatus, *ip.Status)
			d.Set(floatingIPZone, *ip.Zone.Name)

			d.Set(floatingIPCRN, *ip.CRN)

			target, ok := ip.Target.(*vpcv1.FloatingIPTarget)
			if ok {
				d.Set(floatingIPTarget, target.ID)
			}

			tags, err := flex.GetTagsUsingCRN(meta, *ip.CRN)
			if err != nil {
				fmt.Printf("Error on get of vpc Floating IP (%s) tags: %s", *ip.Address, err)
			}

			d.Set(floatingIPTags, tags)
			d.SetId(*ip.ID)

			return nil
		}
	}

	return fmt.Errorf("[ERROR] No floatingIP found with name  %s", name)

}
