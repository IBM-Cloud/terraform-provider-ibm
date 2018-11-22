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

func resourceIBMCISPool() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cis_id": {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "name",
				Required:    true,
				//ValidateFunc: validation.StringLenBetween(1, 32),
			},
			"check_regions": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			// "check_regions": {
			//     Type:     schema.TypeList,
			//     Required: true,
			//     Elem:     &schema.Schema{
			//         Type: schema.TypeString},
			//     },
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"minimum_origins": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"monitor": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"notification_email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"origins": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"address": {
							Type:     schema.TypeString,
							Required: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
			"healthy": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},

		Create: resourceCISpoolCreate,
		Read:   resourceCISpoolRead,
		Update: resourceCISpoolUpdate,
		Delete: resourceCISpoolDelete,
	}
}

func resourceCISpoolCreate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	log.Printf("   client %v\n", cisClient)
	if err != nil {
		return err
	}

	cisId := d.Get("cis_id").(string)
	name := d.Get("name").(string)
	origins := d.Get("origins").(*schema.Set)
	checkRegions := d.Get("check_regions").(*schema.Set).List()
	log.Printf("CheckRegions are >>>>>>>>>>>>>>> %v\n", checkRegions)

	var pools *[]v1.Pool
	pools, err = cisClient.Pools().ListPools(cisId)
	if err != nil {
		log.Printf("ListPools Failed %s\n", err)
		return err
	}

	var poolNames []string
	poolsObj := *pools
	for _, pool := range poolsObj {
		poolNames = append(poolNames, pool.Name)
	}

	log.Println(string("Existing pool names"))
	log.Println(poolNames)

	var pool *v1.Pool
	var poolObj v1.Pool

	// Handle pool existing case gracefully, by just populating Terraform values
	// If zome does not exist, create
	index := indexOf(name, poolNames)
	//indexOf returns -1 if the pool is not found in the list of pools, so we create it
	// Otherwise it returns the index in the paths array.
	if index == -1 {

		poolNew := v1.PoolBody{}
		poolNew.Name = name
		poolNew.CheckRegions = expandToStringList(checkRegions)
		poolNew.Origins = expandOrigins(origins)

		if notEmail, ok := d.GetOk("notEmail"); ok {
			poolNew.NotEmail = notEmail.(string)
		}
		if monitor, ok := d.GetOk("monitor"); ok {
			log.Printf(">>>>>> monitor added with value >>%s<<", monitor.(string))
			poolNew.Monitor = monitor.(string)
		}
		if enabled, ok := d.GetOk("enabled"); ok {
			poolNew.Enabled = enabled.(bool)
		}
		if minOrigins, ok := d.GetOk("minOrigins"); ok {
			poolNew.MinOrigins = minOrigins.(int)
		}
		if description, ok := d.GetOk("description"); ok {
			poolNew.Description = description.(string)
		}

		pool, err = cisClient.Pools().CreatePool(cisId, poolNew)
		if err != nil {
			log.Printf("CreatePools Failed %s\n", err)
			return err
		}
		poolObj = *pool
	} else {
		// If pool already exists retrieve existing pool from array of pools.
		poolObj = poolsObj[index]
	}

	d.SetId(poolObj.Id)
	d.Set("name", poolObj.Name)

	return resourceCISpoolRead(d, meta)
}

func resourceCISpoolRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}

	var poolId string

	poolId = d.Id()
	cisId := d.Get("cis_id").(string)
	//emptyPool := v1.Pool{}

	log.Printf("resourceCISpoolRead - Getting Pool %v\n", poolId)
	var pool *v1.Pool

	pool, err = cisClient.Pools().GetPool(cisId, poolId)
	if err != nil {
		log.Printf("resourceCIpoolRead - ListPools Failed %s\n", err)
		return err
	} else {
		log.Printf("resourceCISpoolRead - Retrieved Pool %v\n", pool)

		// if cmp.Equal(pool, emptyPool) {
		// 	log.Printf("resourceCIShealthCheckRead - No pool returned. Delete")

		// 	// Contrary to the doc. SetId("") does not delete the object on a Read
		// 	//   d.SetId("")
		// } else {

		poolObj := *pool
		d.Set("name", poolObj.Name)
		d.Set("check_regions", poolObj.CheckRegions)
		d.Set("origins", poolObj.Origins)
		d.Set("description", poolObj.Description)
		d.Set("enabled", poolObj.Enabled)
		d.Set("minimum_origins", poolObj.MinOrigins)
		d.Set("monitor", poolObj.Monitor)
		d.Set("notification_email", poolObj.NotEmail)

	}
	return nil
}

func expandToStringList(list interface{}) []string {
	iList := list.([]interface{})
	checkRegions := make([]string, 0, len(iList))
	for _, region := range iList {
		checkRegions = append(checkRegions, region.(string))
	}
	return checkRegions
}

func expandOrigins(originsList *schema.Set) (origins []v1.Origin) {
	for _, iface := range originsList.List() {
		orig := iface.(map[string]interface{})
		origin := v1.Origin{
			Name:    orig["name"].(string),
			Address: orig["address"].(string),
			Enabled: orig["enabled"].(bool),
		}
		origins = append(origins, origin)
	}
	return
}

func resourceCISpoolUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceCISpoolRead(d, meta)
}

func resourceCISpoolDelete(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	poolId := d.Id()
	cisId := d.Get("cis_id").(string)
	var pool *v1.Pool
	emptyPool := new(v1.Pool)

	log.Println("Getting Pool to delete")
	pool, err = cisClient.Pools().GetPool(cisId, poolId)
	if err != nil {
		log.Printf("GetPool Failed %s\n", err)
		return err
	}

	poolObj := *pool
	if !cmp.Equal(poolObj, emptyPool) {
		log.Println("Deleting Pool")
		err = cisClient.Pools().DeletePool(cisId, poolId)
		if err != nil {
			log.Printf("DeletePool Failed %s\n", err)
			return err
		}
	}

	d.SetId("")
	return nil
}
