// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMContainerNLBDNS() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMContainerNLBDNSRead,

		Schema: map[string]*schema.Schema{
			"name": {
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

func dataSourceIBMContainerNLBDNSRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	name := d.Get("name").(string)

	kubeClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return diag.FromErr(err)
	}

	nlbData, err := kubeClient.NlbDns().GetNLBDNSList(name)
	if err != nil || nlbData == nil || len(nlbData) < 1 {
		return diag.FromErr(fmt.Errorf("[ERROR] Error Listing NLB DNS (%s): %s", name, err))
	}
	d.SetId(name)
	d.Set("nlb_config", flattenNlbConfigs(nlbData))
	return nil
}
func flattenNlbConfigs(nlbData []containerv2.NlbVPCListConfig) []map[string]interface{} {
	nlbConfigList := make([]map[string]interface{}, 0)
	for _, n := range nlbData {
		nlbConfig := make(map[string]interface{})
		nlbConfig["secret_name"] = n.SecretName
		nlbConfig["secret_status"] = n.SecretStatus
		c := n.Nlb
		nlbConfig["cluster"] = c.Cluster
		nlbConfig["dns_type"] = c.DnsType
		nlbConfig["lb_hostname"] = c.LbHostname
		nlbConfig["nlb_ips"] = c.NlbIPArray
		nlbConfig["nlb_sub_domain"] = c.NlbSubdomain
		nlbConfig["secret_namespace"] = c.SecretNamespace
		nlbConfig["type"] = c.Type
		nlbConfigList = append(nlbConfigList, nlbConfig)
	}

	return nlbConfigList
}
