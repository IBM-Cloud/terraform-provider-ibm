package ibm

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"math/rand"
	"strconv"
)

func dataSourceIBMCISIP() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMCISIPRead,

		Schema: map[string]*schema.Schema{
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
	rnd := rand.Intn(8999999) + 1000000

	ipsResults, err := cisClient.Ips().ListIps()
	if err != nil {
		log.Printf("resourceCISIPRead - ListIps Failed\n")
		return err
	} else {

		ipsObj := *ipsResults
		d.Set("ipv4_cidrs", ipsObj.Ipv4)
		d.Set("ipv6_cidrs", ipsObj.Ipv6)

		d.SetId(strconv.Itoa(rnd))
	}
	return nil
}
