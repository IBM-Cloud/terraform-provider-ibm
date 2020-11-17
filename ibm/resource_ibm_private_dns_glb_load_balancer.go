package ibm

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM/networking-go-sdk/dnssvcsv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	pdnsGlbLoadBalancerName             = "name"
	pdnsGlbLoadBalancerID               = "loadbalancer_id"
	pdnsGlbLoadBalancerDescription      = "description"
	pdnsGlbLoadBalancerEnabled          = "enabled"
	pdnsGlbLoadBalancerTTL              = "ttl"
	pdnsGlbLoadBalancerHealth           = "health"
	pdnsGlbLoadBalancerFallbackPool     = "fallback_pool"
	pdnsGlbLoadBalancerDefaultPool      = "default_pools"
	pdnsGlbLoadBalancerAZPools          = "az_pools"
	pdnsGlbLoadBalancerAvailabilityZone = "availability_zone"
	pdnsGlbLoadBalancerPools            = "pools"
	pdnsGlbLoadBalancerCreatedOn        = "created_on"
	pdnsGlbLoadBalancerModifiedOn       = "modified_on"
)

func resourceIBMPrivateDNSGLBLoadbalancer() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMPrivateDNSGLBLoadbalancerCreate,
		Read:     resourceIBMPrivateDNSGLBLoadbalancerRead,
		Update:   resourceIBMPrivateDNSGLBLoadbalancerUpdate,
		Delete:   resourceIBMPrivateDNSGLBLoadbalancerDelete,
		Exists:   resourceIBMPrivateDNSGLBLoadbalancerExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			pdnsGlbLoadBalancerID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load balancer Id",
			},

			pdnsInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The GUID of the private DNS.",
			},

			pdnsZoneID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Zone Id",
			},

			pdnsGlbLoadBalancerName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the load balancer",
			},

			pdnsGlbLoadBalancerDescription: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Descriptive text of the load balancer",
			},

			pdnsGlbLoadBalancerEnabled: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the load balancer is enabled",
			},

			pdnsGlbLoadBalancerTTL: {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Time to live in second",
			},

			pdnsGlbLoadBalancerHealth: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load balancer Id",
			},

			pdnsGlbLoadBalancerFallbackPool: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The pool ID to use when all other pools are detected as unhealthy",
			},

			pdnsGlbLoadBalancerDefaultPool: {
				Type:        schema.TypeList,
				Required:    true,
				Description: "A list of pool IDs ordered by their failover priority",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			pdnsGlbLoadBalancerAZPools: {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Map availability zones to pool ID's.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						pdnsGlbLoadBalancerAvailabilityZone: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Availability zone.",
						},

						pdnsGlbLoadBalancerPools: {
							Type:        schema.TypeList,
							Required:    true,
							Description: "List of load balancer pools",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			pdnsGlbLoadBalancerCreatedOn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GLB Load Balancer creation date",
			},

			pdnsGlbLoadBalancerModifiedOn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GLB Load Balancer Modification date",
			},
		},
	}
}

func resourceIBMPrivateDNSGLBLoadbalancerCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDNSClientSession()
	if err != nil {
		return err
	}
	instanceID := d.Get(pdnsInstanceID).(string)
	zoneID := d.Get(pdnsZoneID).(string)
	createlbOptions := sess.NewCreateLoadBalancerOptions(instanceID, zoneID)

	lbname := d.Get(pdnsGlbLoadBalancerName).(string)
	createlbOptions.SetName(lbname)

	if description, ok := d.GetOk(pdnsGlbLoadBalancerDescription); ok {
		createlbOptions.SetDescription(description.(string))
	}
	if enable, ok := d.GetOkExists(pdnsGlbLoadBalancerEnabled); ok {
		createlbOptions.SetEnabled(enable.(bool))
	}
	if ttl, ok := d.GetOk(pdnsGlbLoadBalancerTTL); ok {
		createlbOptions.SetTTL(int64(ttl.(int)))
	}
	if flbpool, ok := d.GetOk(pdnsGlbLoadBalancerFallbackPool); ok {
		createlbOptions.SetFallbackPool(flbpool.(string))
	}

	createlbOptions.SetDefaultPools(expandStringList(d.Get(pdnsGlbLoadBalancerDefaultPool).([]interface{})))

	if AZpools, ok := d.GetOk(pdnsGlbLoadBalancerAZPools); ok {
		expandedAzpools, err := expandGLBAZPools(AZpools)
		if err != nil {
			return err
		}
		createlbOptions.SetAzPools(expandedAzpools)
	}

	result, resp, err := sess.CreateLoadBalancer(createlbOptions)
	if err != nil {
		log.Printf("create global load balancer failed %s", resp)
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", instanceID, zoneID, *result.ID))
	return resourceIBMPrivateDNSGLBLoadbalancerRead(d, meta)
}

func expandGLBAZPools(azpool interface{}) ([]dnssvcsv1.LoadBalancerAzPoolsItem, error) {
	azpools := azpool.(*schema.Set).List()
	expandAZpools := make([]dnssvcsv1.LoadBalancerAzPoolsItem, 0)
	for _, v := range azpools {
		locationConfig := v.(map[string]interface{})
		avzone := locationConfig[pdnsGlbLoadBalancerAvailabilityZone].(string)
		pools := expandStringList(locationConfig[pdnsGlbLoadBalancerPools].([]interface{}))
		aZItem := dnssvcsv1.LoadBalancerAzPoolsItem{
			AvailabilityZone: &avzone,
			Pools:            pools,
		}
		expandAZpools = append(expandAZpools, aZItem)
	}
	return expandAZpools, nil
}

func resourceIBMPrivateDNSGLBLoadbalancerRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDNSClientSession()
	if err != nil {
		return err
	}
	idset := strings.Split(d.Id(), "/")

	getlbOptions := sess.NewGetLoadBalancerOptions(idset[0], idset[1], idset[2])
	presponse, resp, err := sess.GetLoadBalancer(getlbOptions)
	if err != nil {
		return fmt.Errorf("Error fetching pdns GLB :%s\n%s", err, resp)
	}

	response := *presponse
	d.Set(pdnsGlbLoadBalancerName, response.Name)
	d.Set(pdnsGlbLoadBalancerID, response.ID)
	d.Set(pdnsGlbLoadBalancerDescription, response.Description)
	d.Set(pdnsGlbLoadBalancerEnabled, response.Enabled)
	d.Set(pdnsGlbLoadBalancerTTL, response.TTL)
	d.Set(pdnsGlbLoadBalancerHealth, response.Health)
	d.Set(pdnsGlbLoadBalancerFallbackPool, response.FallbackPool)
	d.Set(pdnsGlbLoadBalancerDefaultPool, response.DefaultPools)
	d.Set(pdnsGlbLoadBalancerCreatedOn, response.CreatedOn)
	d.Set(pdnsGlbLoadBalancerModifiedOn, response.ModifiedOn)
	d.Set(pdnsGlbLoadBalancerAZPools, flattenDataSourceLoadBalancerAZpool(response.AzPools))

	log.Printf("global load balancer pool created successfully : %s", idset[1])

	return nil
}

func flattenDataSourceLoadBalancerAZpool(azpool []dnssvcsv1.LoadBalancerAzPoolsItem) interface{} {
	flattened := make([]interface{}, 0)
	for _, v := range azpool {
		cfg := map[string]interface{}{
			pdnsGlbLoadBalancerAvailabilityZone: v.AvailabilityZone,
			pdnsGlbLoadBalancerPools:            flattenStringList(v.Pools),
		}
		flattened = append(flattened, cfg)
	}
	return flattened
}

func resourceIBMPrivateDNSGLBLoadbalancerUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDNSClientSession()
	if err != nil {
		return err
	}

	idset := strings.Split(d.Id(), "/")

	updatelbOptions := sess.NewUpdateLoadBalancerOptions(idset[0], idset[1], idset[2])

	if d.HasChange(pdnsGlbLoadBalancerName) ||
		d.HasChange(pdnsGlbLoadBalancerDescription) ||
		d.HasChange(pdnsGlbLoadBalancerEnabled) ||
		d.HasChange(pdnsGlbLoadBalancerTTL) ||
		d.HasChange(pdnsGlbLoadBalancerFallbackPool) ||
		d.HasChange(pdnsGlbLoadBalancerDefaultPool) ||
		d.HasChange(pdnsGlbLoadBalancerAZPools) {

		if name, ok := d.GetOk(pdnsGlbLoadBalancerName); ok {
			updatelbOptions.SetName(name.(string))
		}
		if description, ok := d.GetOk(pdnsGlbLoadBalancerDescription); ok {
			updatelbOptions.SetDescription(description.(string))
		}
		if enable, ok := d.GetOkExists(pdnsGlbLoadBalancerEnabled); ok {
			updatelbOptions.SetEnabled(enable.(bool))
		}
		if ttl, ok := d.GetOk(pdnsGlbLoadBalancerTTL); ok {
			updatelbOptions.SetTTL(int64(ttl.(int)))
		}
		if flbpool, ok := d.GetOk(pdnsGlbLoadBalancerFallbackPool); ok {
			updatelbOptions.SetFallbackPool(flbpool.(string))
		}

		if _, ok := d.GetOk(pdnsGlbLoadBalancerDefaultPool); ok {
			updatelbOptions.SetDefaultPools(expandStringList(d.Get(pdnsGlbLoadBalancerDefaultPool).([]interface{})))
		}

		if AZpools, ok := d.GetOk(pdnsGlbLoadBalancerAZPools); ok {
			expandedAzpools, err := expandGLBAZPools(AZpools)
			if err != nil {
				return err
			}
			updatelbOptions.SetAzPools(expandedAzpools)
		}

		result, detail, err := sess.UpdateLoadBalancer(updatelbOptions)
		if err != nil {
			return fmt.Errorf("Error updating pdns GLB :%s\n%s", err, detail)
		}
		log.Printf("Load Balancer update succesful : %s", *result.ID)
	}

	return resourceIBMPrivateDNSGLBLoadbalancerRead(d, meta)
}

func resourceIBMPrivateDNSGLBLoadbalancerDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDNSClientSession()
	if err != nil {
		return err
	}

	idset := strings.Split(d.Id(), "/")

	DeletelbOptions := sess.NewDeleteLoadBalancerOptions(idset[0], idset[1], idset[2])
	response, err := sess.DeleteLoadBalancer(DeletelbOptions)

	if err != nil {
		return fmt.Errorf("Error deleting pdns GLB :%s\n%s", err, response)
	}

	d.SetId("")
	return nil
}

func resourceIBMPrivateDNSGLBLoadbalancerExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).PrivateDNSClientSession()
	if err != nil {
		return false, err
	}

	idset := strings.Split(d.Id(), "/")

	getlbOptions := sess.NewGetLoadBalancerOptions(idset[0], idset[1], idset[1])
	response, detail, err := sess.GetLoadBalancer(getlbOptions)
	if err != nil {
		if response != nil && detail != nil && detail.StatusCode == 404 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
