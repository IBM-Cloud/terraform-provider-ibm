package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk-scoped/vpcv1"
)

func dataSourceIBMISInstanceGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISInstanceGroupRead,

		Schema: map[string]*schema.Schema{

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user-defined name for this instance group",
			},

			"instance_template": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "instance template ID",
			},

			"membership_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of instances in the instance group",
			},

			"resource_group": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Resource group ID",
			},

			"subnets": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
				Description: "list of subnet IDs",
			},

			"application_port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Used by the instance group when scaling up instances to supply the port for the load balancer pool member.",
			},

			"load_balancer_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "load balancer ID",
			},

			"load_balancer_pool_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "load balancer pool ID",
			},

			"managers": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
				Description: "list of Managers associated with instancegroup",
			},

			"vpc": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "vpc instance",
			},
		},
	}
}

func dataSourceIBMISInstanceGroupRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}

	listInstanceGroupOptions := vpcv1.ListInstanceGroupsOptions{}
	instanceGroupsCollection, response, err := sess.ListInstanceGroups(&listInstanceGroupOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting InstanceGroup: %s\n%s", err, response)
	}
	name := d.Get("name")
	for _, instanceGroup := range instanceGroupsCollection.InstanceGroups {
		if *instanceGroup.Name == name {
			d.Set("name", *instanceGroup.Name)
			d.Set("instance_template", *instanceGroup.InstanceTemplate.ID)
			d.Set("membership_count", *instanceGroup.MembershipCount)
			d.Set("resource_group", *instanceGroup.ResourceGroup.ID)
			d.SetId(*instanceGroup.ID)
			if instanceGroup.ApplicationPort != nil {
				d.Set("application_port", *instanceGroup.ApplicationPort)
			}
			subnets := make([]string, 0)
			for i := 0; i < len(instanceGroup.Subnets); i++ {
				subnets = append(subnets, string(*(instanceGroup.Subnets[i].ID)))
			}
			if instanceGroup.LoadBalancerPool != nil {
				d.Set("load_balancer_pool_id", *instanceGroup.LoadBalancerPool)
			}
			d.Set("subnets", subnets)
			managers := make([]string, 0)
			for i := 0; i < len(instanceGroup.Managers); i++ {
				managers = append(managers, string(*(instanceGroup.Managers[i].ID)))
			}
			d.Set("managers", managers)
			d.Set("vpc", *instanceGroup.VPC.ID)
		}
	}
	return nil
}
