package ibm

import (
	"log"
	"reflect"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/hashicorp/terraform/helper/schema"
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
				ValidateFunc: validateAllowedStringValue([]string{"none", "cookie"}),
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
		// No Exists due to errors in CIS API returning incorrect return codes on 404
		Importer: &schema.ResourceImporter{},
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
	zoneId, _, err := convertTftoCisTwoVar(d.Get("domain_id").(string))
	if err != nil {
		return err
	}

	var glb *v1.Glb
	var glbObj v1.Glb
	glbNew := v1.GlbBody{}
	glbNew.Name = name

	tfDefaultPools := expandStringList(d.Get("default_pool_ids").(*schema.Set).List())
	defaultPoolIds, _, err := convertTfToCisTwoVarSlice(tfDefaultPools)
	glbNew.DefaultPools = defaultPoolIds
	// glbNew.RegionPools
	// glbNew.PopPools

	fbPoolId := d.Get("fallback_pool_id").(string)
	glbNew.FallbackPool, _, err = convertTftoCisTwoVar(fbPoolId)
	glbNew.Proxied = d.Get("proxied").(bool)
	glbNew.SessionAffinity = d.Get("session_affinity").(string)

	if description, ok := d.GetOk("description"); ok {
		glbNew.Desc = description.(string)
	}
	glb, err = cisClient.Glbs().CreateGlb(cisId, zoneId, glbNew)
	if err != nil {
		log.Printf("CreateGlbs Failed %s\n", err)
		return err
	}
	glbObj = *glb
	d.SetId(convertCisToTfThreeVar(glbObj.Id, zoneId, cisId))

	return resourceCISGlbRead(d, meta)
}

func resourceCISGlbRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	// Extract CIS Ids from TF Id
	glbId, zoneId, cisId, err := convertTfToCisThreeVar(d.Id())
	if err != nil {
		log.Printf("resourceCISGlbRead %s\n", err)
		return err
	}
	var glb *v1.Glb

	glb, err = cisClient.Glbs().GetGlb(cisId, zoneId, glbId)
	if err != nil {
		log.Printf("resourceCIGlbRead - ListGlbs Failed %s\n", err)
		return err
	}
	glbObj := *glb
	d.Set("cis_id", cisId)
	d.Set("domain_id", convertCisToTfTwoVar(zoneId, cisId))
	d.Set("name", glbObj.Name)
	d.Set("default_pool_ids", convertCisToTfTwoVarSlice(glbObj.DefaultPools, cisId))
	d.Set("description", glbObj.Desc)
	d.Set("fallback_pool_id", convertCisToTfTwoVar(glbObj.FallbackPool, cisId))
	d.Set("ttl", glbObj.Ttl)
	d.Set("proxied", glbObj.Proxied)
	d.Set("session_affinity", glbObj.SessionAffinity)

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
	glbId, zoneId, cisId, _ := convertTfToCisThreeVar(d.Id())
	var glb *v1.Glb
	emptyGlb := new(v1.Glb)

	log.Println("Getting Glb to delete")
	glb, err = cisClient.Glbs().GetGlb(cisId, zoneId, glbId)
	if err != nil {
		log.Printf("GetGlb Failed %s\n", err)
		return err
	}

	glbObj := *glb
	if !reflect.DeepEqual(emptyGlb, glbObj) {
		log.Println("Deleting Glb")
		err = cisClient.Glbs().DeleteGlb(cisId, zoneId, glbId)
		if err != nil {
			log.Printf("DeleteGlb Failed %s\n", err)
			return err
		}
	}

	d.SetId("")
	return nil
}
