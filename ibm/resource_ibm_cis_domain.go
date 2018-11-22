package ibm

import (
	//"fmt"
	"log"
	//"strings"
	//"time"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	//"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/google/go-cmp/cmp"
	//"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
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
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
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

	//var zoneList []string
	log.Printf(">>>> cisId %s  domain %s\n", cisId, zoneName)

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

	log.Println(string("Existing zone names"))
	log.Println(zoneNames)

	//var nameServers []string
	zoneNew := v1.ZoneBody{Name: zoneName}

	var zone *v1.Zone
	var zoneObj v1.Zone

	// Handle zone existing case gracefully, by just populating Terraform values
	// If zome does not exist, create
	index := indexOf(zoneName, zoneNames)
	//indexOf returns -1 if the zone is not found in the list of zones, so we create it
	// Otherwise it returns the index in the names array.
	if index == -1 {
		zone, err = cisClient.Zones().CreateZone(cisId, zoneNew)
		if err != nil {
			log.Printf("CreateZones Failed %s\n", err)
			return err
		}
		zoneObj = *zone
	} else {
		// If zone already exists retrieve existing zone from array of zones.
		zoneObj = zonesObj[index]
	}

	d.SetId(zoneObj.Id)
	d.Set("name", zoneObj.Name)

	// id1 := d.Get("id").(string)
	// log.Printf("Getting first Zone ID %v\n", id1)

	id2 := d.Id()
	log.Printf("Getting first Zone ID %v\n", id2)

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
	//emptyZone := v1.Zone{}

	log.Printf("resourceCISdomainRead - Getting Zone %v\n", zoneId)
	var zone *v1.Zone

	zone, err = cisClient.Zones().GetZone(cisId, zoneId)
	if err != nil {
		log.Printf("resourceCISdomainRead - ListZones Failed %s\n", err)
		return err
	} else {
		log.Printf("resourceCISdomainRead - Retrieved Zone %v\n", zone)

		// if cmp.Equal(zone, emptyZone) {
		// 	log.Printf("resourceCISdomainRead - No zone returned. Delete")

		// 	// Contrary to the doc. SetId("") does not delete the object on a Read
		// 	//   d.SetId("")
		// } else {

		zoneObj := *zone

		d.Set("name", zoneObj.Name)
		d.Set("status", zoneObj.Status)
		d.Set("paused", zoneObj.Paused)
		d.Set("name_servers", zoneObj.NameServers)
		d.Set("original_name_servers", zoneObj.OriginalNameServer)

		// }
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
	if !cmp.Equal(zoneObj, emptyZone) {
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
