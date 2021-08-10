// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMSatelliteLocationNLBDNS() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMSatelliteLocationNLBDNSRead,

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A unique name of the cluster",
			},
			"nlb_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"secret_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"secret_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cluster": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dns_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lb_hostname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"nlb_ips": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"nlb_sub_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"secret_namespace": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMSatelliteLocationNLBDNSRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	location := d.Get("location").(string)

	satClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return diag.FromErr(err)
	}

	nlbData, err := satClient.NlbDns().GetLocationNLBDNSList(location)
	if err != nil || nlbData == nil || len(nlbData) < 1 {
		return diag.FromErr(fmt.Errorf("[ERROR] Error Listing Satellite NLB DNS (%s): %s", location, err))
	}
	d.SetId(location)
	d.Set("nlb_config", flattenNlbConfigs(nlbData))
	return nil
}
