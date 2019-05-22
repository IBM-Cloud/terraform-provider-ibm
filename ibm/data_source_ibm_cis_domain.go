package ibm

import (
	"fmt"
	"log"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceIBMCISDomain() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMCISDomainRead,

		Schema: map[string]*schema.Schema{
			"cis_id": {
				Type:        schema.TypeString,
				Description: "CIS object id",
				Required:    true,
			},
			"domain": {
				Type:        schema.TypeString,
				Description: "CISzone - Domain",
				Required:    true,
			},
			"paused": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"original_name_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceIBMCISDomainRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}

	cisId := d.Get("cis_id").(string)
	zoneName := d.Get("domain").(string)
	var zones *[]v1.Zone
	var zoneNames []string

	zones, err = cisClient.Zones().ListZones(cisId)
	if err != nil {
		log.Printf("dataSourcCISdomainRead - ListZones Failed %s\n", err)
		return err
	}
	zonesObj := *zones

	for _, zone := range zonesObj {
		zoneNames = append(zoneNames, zone.Name)
	}
	log.Println(string("Existing zone names"))
	log.Println(zoneNames)

	index := indexOf(zoneName, zoneNames)
	if index == -1 {
		return fmt.Errorf("dataSourcCISdomainRead - Domain with name %s does not exist", zoneName)
	}

	zoneObj := zonesObj[index]
	d.SetId(convertCisToTfTwoVar(zoneObj.Id, cisId))
	d.Set("cis_id", cisId)
	d.Set("name", zoneObj.Name)
	d.Set("domain", zoneObj.Name)
	d.Set("status", zoneObj.Status)
	d.Set("paused", zoneObj.Paused)
	d.Set("name_servers", zoneObj.NameServers)
	d.Set("original_name_servers", zoneObj.OriginalNameServer)

	return nil
}
