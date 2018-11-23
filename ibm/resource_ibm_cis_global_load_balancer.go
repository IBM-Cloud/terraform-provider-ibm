package ibm

import (
	//"fmt"
	"log"
	//"strings"
	//"time"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/google/go-cmp/cmp"
	//"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	//"github.com/hashicorp/terraform/helper/validation"
)

func resourceIBMCISGlb() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cis_id": {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
			},
			"domain_id": {
				Type:        schema.TypeString,
				Description: "Associated CIS domain",
				Required:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "name",
				Required:    true,
			},
			"fallback_pool_id": {
				Type:        schema.TypeString,
				Description: "name",
				Required:    true,
				//ValidateFunc: validation.StringLenBetween(1, 32),
			},
			"default_pool_ids": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				//ValidateFunc: validation.StringLenBetween(1, 32),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				//ValidateFunc: validation.StringLenBetween(0, 1024),
			},
			"ttl": {
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"proxied"}, // this is set to zero regardless of config when proxied=true

			},
			"proxied": {
				Type:          schema.TypeBool,
				Optional:      true,
				Default:       false,
				ConflictsWith: []string{"ttl"},
			},
			"session_affinity": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "none",
				// Set to cookie when proxy=true
				//ValidateFunc: validation.StringInSlice([]string{"none", "cookie"}, false),
			},

			// "region_pools": &schema.Schema{
			// 	Type:     schema.TypeMap,
			// 	Optional: true,
			// 	Computed: true,
			// 	Elem:     &schema.Schema{Type: schema.TypeString},
			// },
			// "pop_pools": &schema.Schema{
			// 	Type:     schema.TypeMap,
			// 	Optional: true,
			// 	Computed: true,
			// 	Elem:     &schema.Schema{Type: schema.TypeString},
			// },
			"created_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},

		Create: resourceCISGlbCreate,
		Read:   resourceCISGlbRead,
		Update: resourceCISGlbUpdate,
		Delete: resourceCISGlbDelete,
	}
}

func resourceCISGlbCreate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	log.Printf("   client %v\n", cisClient)
	if err != nil {
		return err
	}

	cisId := d.Get("cis_id").(string)
	name := d.Get("name").(string)
	zoneId := d.Get("domain_id").(string)
	defaultPools := d.Get("default_pool_ids").(*schema.Set).List()

	var Glbs *[]v1.Glb
	Glbs, err = cisClient.Glbs().ListGlbs(cisId, zoneId)
	if err != nil {
		log.Printf("ListGlbs Failed %s\n", err)
		return err
	}

	var GlbNames []string
	GlbsObj := *Glbs
	for _, Glb := range GlbsObj {
		GlbNames = append(GlbNames, Glb.Name)
	}

	log.Println(string("Existing Glb names"))
	log.Println(GlbNames)

	var Glb *v1.Glb
	var GlbObj v1.Glb

	// Handle Glb existing case gracefully, by just populating Terraform values
	// If zome does not exist, create
	index := indexOf(name, GlbNames)
	//indexOf returns -1 if the Glb is not found in the list of Glbs, so we create it
	// Otherwise it returns the index in the paths array.
	if index == -1 {

		GlbNew := v1.GlbBody{}
		GlbNew.Name = name

		GlbNew.DefaultPools = expandToStringList(defaultPools)
		GlbNew.FallbackPool = d.Get("fallback_pool_id").(string)
		GlbNew.Proxied = d.Get("proxied").(bool)
		GlbNew.SessionAffinity = d.Get("session_affinity").(string)

		if description, ok := d.GetOk("description"); ok {
			GlbNew.Desc = description.(string)
		}
		Glb, err = cisClient.Glbs().CreateGlb(cisId, zoneId, GlbNew)
		if err != nil {
			log.Printf("CreateGlbs Failed %s\n", err)
			return err
		}
		GlbObj = *Glb
	} else {
		// If Glb already exists retrieve existing Glb from array of Glbs.
		GlbObj = GlbsObj[index]
	}

	d.SetId(GlbObj.Id)
	d.Set("domain_id", zoneId)

	return resourceCISGlbRead(d, meta)
}

func resourceCISGlbRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}

	var GlbId string

	GlbId = d.Id()
	cisId := d.Get("cis_id").(string)
	zoneId := d.Get("domain_id").(string)
	//emptyGlb := v1.Glb{}

	log.Printf("resourceCISGlbRead - Getting Glb %v\n", GlbId)
	var Glb *v1.Glb

	Glb, err = cisClient.Glbs().GetGlb(cisId, zoneId, GlbId)
	if err != nil {
		log.Printf("resourceCIGlbRead - ListGlbs Failed %s\n", err)
		return err
	} else {
		log.Printf("resourceCISGlbRead - Retrieved Glb %v\n", Glb)

		// if cmp.Equal(Glb, emptyGlb) {
		// 	log.Printf("resourceCIShealthCheckRead - No Glb returned. Delete")

		// 	// Contrary to the doc. SetId("") does not delete the object on a Read
		// 	//   d.SetId("")
		// } else {

		GlbObj := *Glb
		d.Set("name", GlbObj.Name)
		d.Set("default_pool_ids", GlbObj.DefaultPools)
		d.Set("description", GlbObj.Desc)
		d.Set("fallback_pool_id", GlbObj.FallbackPool)
		d.Set("ttl", GlbObj.Ttl)
		d.Set("proxied", GlbObj.Proxied)
		d.Set("session_affinity", GlbObj.SessionAffinity)

	}
	return nil
}

func resourceCISGlbUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceCISGlbRead(d, meta)
}

func resourceCISGlbDelete(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	GlbId := d.Id()
	cisId := d.Get("cis_id").(string)
	zoneId := d.Get("domain_id").(string)
	var Glb *v1.Glb
	emptyGlb := new(v1.Glb)

	log.Println("Getting Glb to delete")
	Glb, err = cisClient.Glbs().GetGlb(cisId, zoneId, GlbId)
	if err != nil {
		log.Printf("GetGlb Failed %s\n", err)
		return err
	}

	GlbObj := *Glb
	if !cmp.Equal(GlbObj, emptyGlb) {
		log.Println("Deleting Glb")
		err = cisClient.Glbs().DeleteGlb(cisId, zoneId, GlbId)
		if err != nil {
			log.Printf("DeleteGlb Failed %s\n", err)
			return err
		}
	}

	d.SetId("")
	return nil
}
