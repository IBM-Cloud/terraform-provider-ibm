package ibm

import (
	"fmt"
	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"reflect"
)

func resourceIBMCISDomain() *schema.Resource {
	return &schema.Resource{
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
		Create: resourceCISdomainCreate,
		Read:   resourceCISdomainRead,
		Update: resourceCISdomainUpdate,
		Delete: resourceCISdomainDelete,
		// No Exists due to errors in CIS API returning incorrect return codes on 404
	}
}

func resourceCISdomainCreate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	log.Printf("   client %v\n", cisClient)
	if err != nil {
		return err
	}

	cisId := d.Get("cis_id").(string)
	zoneName := d.Get("domain").(string)

	var zones *[]v1.Zone
	zones, err = cisClient.Zones().ListZones(cisId)
	if err != nil {
		log.Printf("ListZones Failed %s\n", err)
		return err
	}

	var zoneNames []string
	zonesObj := *zones
	for _, zone := range zonesObj {
		zoneNames = append(zoneNames, zone.Name)
	}
	zoneNew := v1.ZoneBody{Name: zoneName}

	var zone *v1.Zone
	var zoneObj v1.Zone

	index := indexOf(zoneName, zoneNames)
	if index == -1 {
		zone, err = cisClient.Zones().CreateZone(cisId, zoneNew)
		if err != nil {
			log.Printf("CreateZones Failed %s\n", err)
			return err
		}
		zoneObj = *zone
	} else {
		return fmt.Errorf("resource with name %s already exists", zoneName)
	}

	d.SetId(zoneObj.Id)
	d.Set("name", zoneObj.Name)
	return resourceCISdomainRead(d, meta)
}

func resourceCISdomainRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}

	var zoneId string

	zoneId = d.Id()
	cisId := d.Get("cis_id").(string)
	zoneId = d.Id()

	log.Printf("resourceCISdomainRead - Getting Zone %v\n", zoneId)
	var zone *v1.Zone

	zone, err = cisClient.Zones().GetZone(cisId, zoneId)
	if err != nil {
		log.Printf("resourceCISdomainRead - ListZones Failed %s\n", err)
		return err
	} else {
		log.Printf("resourceCISdomainRead - Retrieved Zone %v\n", zone)

		zoneObj := *zone

		d.Set("name", zoneObj.Name)
		d.Set("status", zoneObj.Status)
		d.Set("paused", zoneObj.Paused)
		d.Set("name_servers", zoneObj.NameServers)
		d.Set("original_name_servers", zoneObj.OriginalNameServer)

	}
	return nil
}

func resourceCISdomainUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceCISdomainRead(d, meta)
}

func resourceCISdomainDelete(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	zoneId := d.Id()
	cisId := d.Get("cis_id").(string)
	var zone *v1.Zone
	emptyZone := new(v1.Zone)

	log.Println("Getting Zone to delete")
	zone, err = cisClient.Zones().GetZone(cisId, zoneId)
	if err != nil {
		log.Printf("GetZone Failed %s\n", err)
		return err
	}

	zoneObj := *zone
	if !reflect.DeepEqual(emptyZone, zoneObj) {
		log.Println("Deleting Zone")
		err = cisClient.Zones().DeleteZone(cisId, zoneId)
		if err != nil {
			log.Printf("DeleteZone Failed %s\n", err)
			return err
		}
	}

	d.SetId("")
	return nil
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
