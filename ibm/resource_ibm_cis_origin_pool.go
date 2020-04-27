package ibm

import (
	"log"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
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
				Description: "List of regions",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the CIS Origin Pool",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Boolean value set to true if cis origin pool needs to be enabled",
			},
			"minimum_origins": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "Minimum number of Origins",
			},
			"monitor": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Monitor value",
			},
			"notification_email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Email address configured to recieve the notifications",
			},
			"origins": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Origins info",
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
						"weight": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"healthy": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"health": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Health info",
			},
			"created_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation date info",
			},
			"modified_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Modified date info",
			},
		},

		Create:   resourceCISpoolCreate,
		Read:     resourceCISpoolRead,
		Update:   resourceCISpoolUpdate,
		Delete:   resourceCISpoolDelete,
		Importer: &schema.ResourceImporter{},
	}
}

func resourceCISpoolCreate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}

	cisId := d.Get("cis_id").(string)
	name := d.Get("name").(string)
	origins := d.Get("origins").(*schema.Set)
	checkRegions := d.Get("check_regions").(*schema.Set).List()

	poolNew := v1.PoolBody{}
	poolNew.Name = name
	poolNew.CheckRegions = expandStringList(checkRegions)
	poolNew.Origins = expandOrigins(origins)

	if notEmail, ok := d.GetOk("notification_email"); ok {
		poolNew.NotEmail = notEmail.(string)
	}
	if monitor, ok := d.GetOk("monitor"); ok {
		monitorId, _, _ := convertTftoCisTwoVar(monitor.(string))
		poolNew.Monitor = monitorId
	}
	if enabled, ok := d.GetOk("enabled"); ok {
		poolNew.Enabled = enabled.(bool)
	}
	if minOrigins, ok := d.GetOk("minimum_origins"); ok {
		poolNew.MinOrigins = minOrigins.(int)
	}
	if description, ok := d.GetOk("description"); ok {
		poolNew.Description = description.(string)
	}

	var pool *v1.Pool
	var poolObj v1.Pool

	pool, err = cisClient.Pools().CreatePool(cisId, poolNew)
	if err != nil {
		log.Printf("[WARN] CreatePools Failed %s\n", err)
		return err
	}

	poolObj = *pool

	//Set unique TF Id from concatenated CIS Ids
	d.SetId(convertCisToTfTwoVar(poolObj.Id, cisId))
	d.Set("name", poolObj.Name)

	return resourceCISpoolRead(d, meta)
}

func resourceCISpoolRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	poolId, cisId, err := convertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}
	var pool *v1.Pool
	pool, err = cisClient.Pools().GetPool(cisId, poolId)
	if err != nil {
		if checkCisPoolDeleted(d, meta, err, pool) {
			d.SetId("")
			return nil
		}
		log.Printf("[WARN] Error getting zone during PoolRead %v\n", err)
		return err
	}

	poolObj := *pool
	d.Set("cis_id", cisId)
	d.Set("name", poolObj.Name)
	d.Set("check_regions", poolObj.CheckRegions)
	d.Set("origins", flattenOrigins(poolObj.Origins))
	d.Set("description", poolObj.Description)
	d.Set("enabled", poolObj.Enabled)
	d.Set("minimum_origins", poolObj.MinOrigins)
	d.Set("monitor", convertCisToTfTwoVar(poolObj.Monitor, cisId))
	d.Set("notification_email", poolObj.NotEmail)
	d.Set("health", poolObj.Health)
	d.Set("created_on", poolObj.CreatedOn)
	d.Set("modified_on", poolObj.ModifiedOn)

	return nil
}

func resourceCISpoolUpdate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	poolID, cisID, err := convertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}
	poolUpdate := v1.PoolBody{}
	if d.HasChange("name") || d.HasChange("origins") || d.HasChange("check_regions") || d.HasChange("notification_email") || d.HasChange("monitor") || d.HasChange("enabled") || d.HasChange("minimum_origins") || d.HasChange("description") {
		if name, ok := d.GetOk("name"); ok {
			poolUpdate.Name = name.(string)
		}
		if origin, ok := d.GetOk("origins"); ok {
			origins := origin.(*schema.Set)
			poolUpdate.Origins = expandOrigins(origins)
		}
		if checkregions, ok := d.GetOk("check_regions"); ok {
			checkRegions := checkregions.(*schema.Set).List()
			poolUpdate.CheckRegions = expandStringList(checkRegions)
		}
		if notEmail, ok := d.GetOk("notification_email"); ok {
			poolUpdate.NotEmail = notEmail.(string)
		}
		if monitor, ok := d.GetOk("monitor"); ok {
			monitorID, _, _ := convertTftoCisTwoVar(monitor.(string))
			poolUpdate.Monitor = monitorID
		}
		if enabled, ok := d.GetOk("enabled"); ok {
			poolUpdate.Enabled = enabled.(bool)
		}
		if minOrigins, ok := d.GetOk("minimum_origins"); ok {
			poolUpdate.MinOrigins = minOrigins.(int)
		}
		if description, ok := d.GetOk("description"); ok {
			poolUpdate.Description = description.(string)
		}
		_, err = cisClient.Pools().UpdatePool(cisID, poolID, poolUpdate)
		if err != nil {
			log.Printf("[WARN] Error getting zone during PoolUpdate %v\n", err)
			return err
		}
	}

	return resourceCISpoolRead(d, meta)
}

func resourceCISpoolDelete(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	poolId, cisId, err := convertTftoCisTwoVar(d.Id())
	var pool *v1.Pool
	emptyPool := new(v1.Pool)
	pool, err = cisClient.Pools().GetPool(cisId, poolId)
	if err != nil {
		if checkCisPoolDeleted(d, meta, err, pool) {
			d.SetId("")
			return nil
		}
		log.Printf("[WARN] Error getting zone during PoolRead %v\n", err)
		return err
	}

	poolObj := *pool
	if !reflect.DeepEqual(emptyPool, poolObj) {
		err = cisClient.Pools().DeletePool(cisId, poolId)
		if err != nil {
			log.Printf("[WARN] DeletePool Failed %s\n", err)
			return err
		}
	}

	d.SetId("")
	return nil
}

func checkCisPoolDeleted(d *schema.ResourceData, meta interface{}, errCheck error, pool *v1.Pool) bool {
	// Check if error is due to removal of Cis resource and hence all subresources
	if strings.Contains(errCheck.Error(), "Object not found") ||
		strings.Contains(errCheck.Error(), "status code: 404") ||
		strings.Contains(errCheck.Error(), "Invalid zone identifier") { //code 400
		log.Printf("[WARN] Removing resource from state because it's not found via the CIS API")
		return true
	}
	_, cisId, _ := convertTftoCisTwoVar(d.Id())
	exists, errNew := rcInstanceExists(cisId, "ibm_cis", meta)
	if errNew != nil {
		log.Printf("[WARN] resourceCISpoolRead - Failure validating service exists %s\n", errNew)
		return false
	}
	if !exists {
		log.Printf("[WARN] Removing pool from state because parent cis instance is in removed state")
		return true
	}
	return false
}
