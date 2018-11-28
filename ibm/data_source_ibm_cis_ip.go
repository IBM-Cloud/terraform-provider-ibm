package ibm

import (
	//"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func dataSourceIBMCISIP() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMCISIPRead,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: "A domain record's internal identifier",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"ipv4_cidrs": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ipv6_cidrs": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceIBMCISIPRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	rnd := acctest.RandString(10)
	log.Printf("resourceCISSettingsRead - Getting CIS IPs \n")

	ipsResults, err := cisClient.Ips().ListIps()
	if err != nil {
		log.Printf("resourceCISIPRead - ListIps Failed\n")
		return err
	} else {

		ipsObj := *ipsResults
		d.Set("ipv4_cidrs", ipsObj.Ipv4)
		d.Set("ipv6_cidrs", ipsObj.Ipv6)

		d.SetId(rnd)
	}
	return nil
}
