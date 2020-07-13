package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk-scoped/vpcv1"
)

func resourceIBMISInstanceGroup() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISInstanceGroupCreate,
		Read:     resourceIBMISInstanceGroupRead,
		Update:   resourceIBMISInstanceGroupUpdate,
		Delete:   resourceIBMISInstanceGroupDelete,
		Exists:   resourceIBMISInstanceGroupExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The user-defined name for this instance group",
			},

			"instance_template": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance template ID",
			},

			"membership_count": {
				Type:        schema.TypeInt,
				Required:    true,
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
				Required:    true,
				Description: "list of subnet IDs",
			},

			"application_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Used by the instance group when scaling up instances to supply the port for the load balancer pool member.",
			},

			"load_balancer_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "load balancer ID",
			},

			"load_balancer_pool_id": {
				Type:        schema.TypeString,
				Optional:    true,
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

func resourceIBMISInstanceGroupCreate(d *schema.ResourceData, meta interface{}) error {

	name := d.Get("name").(string)
	instanceTemplate := d.Get("instance_template").(string)
	membershipCount := d.Get("membership_count").(int)
	subnets := d.Get("subnets")
	//lbID := d.Get("load_balancer_id").(string)
	//lbPoolID := d.Get("load_balancer_pool_id").(string)

	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}
	// return nil

	var subnetIDs []vpcv1.SubnetIdentityIntf
	for _, s := range subnets.([]interface{}) {
		subnet := s.(string)
		subnetIDs = append(subnetIDs, &vpcv1.SubnetIdentity{ID: &subnet})
	}
	mc := int64(membershipCount)
	instanceGroupOptions := vpcv1.CreateInstanceGroupOptions{
		InstanceTemplate: &vpcv1.InstanceTemplateIdentity{
			ID: &instanceTemplate,
		},
		Subnets:         subnetIDs,
		Name:            &name,
		MembershipCount: &mc,
		//LoadBalancer:     &vpcv1.LoadBalancerIdentity{ID: &lbID},
		//LoadBalancerPool: &vpcv1.LoadBalancerPoolIdentity{ID: &lbPoolID},
		// ResourceGroup: &vpcv1.ResourceGroupIdentity{
		// 	ID: &,
		// }
	}
	instanceGroup, response, err := sess.CreateInstanceGroup(&instanceGroupOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Creating InstanceGroup: %s\n%s", err, response)
	}
	d.SetId(*instanceGroup.ID)

	return resourceIBMISInstanceGroupUpdate(d, meta)

}

func resourceIBMISInstanceGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}

	var changed bool
	instanceGroupUpdateOptions := vpcv1.UpdateInstanceGroupOptions{}

	if d.HasChange("name") && !d.IsNewResource() {
		name := d.Get("name").(string)
		instanceGroupUpdateOptions.Name = &name
		changed = true
	}

	if d.HasChange("instance_template") && !d.IsNewResource() {
		instanceTemplate := d.Get("instance_template").(string)
		instanceGroupUpdateOptions.InstanceTemplate = &vpcv1.InstanceTemplateIdentity{
			ID: &instanceTemplate,
		}
		changed = true
	}

	if d.HasChange("membership_count") && !d.IsNewResource() {
		membershipCount := d.Get("membership_count").(int)
		mc := int64(membershipCount)
		instanceGroupUpdateOptions.MembershipCount = &mc
		changed = true
	}

	if d.HasChange("subnets") && !d.IsNewResource() {
		subnets := d.Get("subnets")
		//oldList, newList := d.GetChange("subnets")
		var subnetIDs []vpcv1.SubnetIdentityIntf
		for _, s := range subnets.([]interface{}) {
			subnet := s.(string)
			subnetIDs = append(subnetIDs, &vpcv1.SubnetIdentity{ID: &subnet})
		}
		instanceGroupUpdateOptions.Subnets = subnetIDs
		changed = true
	}

	if d.HasChange("application_port") && !d.IsNewResource() {
		lbID := d.Get("application_port").(int64)
		instanceGroupUpdateOptions.ApplicationPort = &lbID
		changed = true
	}

	if d.HasChange("load_balancer_id") && !d.IsNewResource() {
		lbID := d.Get("load_balancer_id").(string)
		instanceGroupUpdateOptions.LoadBalancer = &vpcv1.LoadBalancerIdentity{ID: &lbID}
		changed = true
	}

	if d.HasChange("load_balancer_pool_id") && !d.IsNewResource() {
		lbPoolID := d.Get("load_balancer_pool_id").(string)
		instanceGroupUpdateOptions.LoadBalancerPool = &vpcv1.LoadBalancerPoolIdentity{ID: &lbPoolID}
		changed = true
	}

	if changed {
		instanceGroupID := d.Id()
		instanceGroupUpdateOptions.ID = &instanceGroupID
		_, response, err := sess.UpdateInstanceGroup(&instanceGroupUpdateOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			return fmt.Errorf("Error Updating InstanceGroup: %s\n%s", err, response)
		}
	}
	return resourceIBMISInstanceGroupRead(d, meta)
}

func resourceIBMISInstanceGroupRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}

	instanceGroupID := d.Id()
	getInstanceGroupOptions := vpcv1.GetInstanceGroupOptions{ID: &instanceGroupID}
	instanceGroup, response, err := sess.GetInstanceGroup(&getInstanceGroupOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting InstanceGroup: %s\n%s", err, response)
	}
	d.Set("name", *instanceGroup.Name)
	d.Set("instance_template", *instanceGroup.InstanceTemplate)
	d.Set("membership_count", *instanceGroup.MembershipCount)
	d.Set("resource_group", *instanceGroup.ResourceGroup)
	// if *instanceGroup.ApplicationPort != nil {
	// 	d.Set("application_port", *instanceGroup.ApplicationPort)
	// }

	subnets := make([]string, 0)

	for i := 0; i < len(instanceGroup.Subnets); i++ {
		subnets = append(subnets, string(*(instanceGroup.Subnets[i].ID)))
	}
	// if *instanceGroup.LoadBalancerPool != nil {
	// 	d.Set("load_balancer_pool_id", *instanceGroup.LoadBalancerPool)
	// }
	d.Set("subnets", subnets)
	managers := make([]string, 0)

	for i := 0; i < len(instanceGroup.Managers); i++ {
		managers = append(managers, string(*(instanceGroup.Managers[i].ID)))
	}
	d.Set("managers", managers)

	d.Set("vpc", *instanceGroup.VPC.ID)

	return nil
}

func resourceIBMISInstanceGroupDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}
	instanceGroupID := d.Id()
	deleteInstanceGroupOptions := vpcv1.DeleteInstanceGroupOptions{ID: &instanceGroupID}
	response, err := sess.DeleteInstanceGroup(&deleteInstanceGroupOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Deleting the InstanceGroup: %s\n%s", err, response)
	}
	return nil
}

func resourceIBMISInstanceGroupExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := myvpcClient(meta)
	if err != nil {
		return false, err
	}
	instanceGroupID := d.Id()
	getInstanceGroupOptions := vpcv1.GetInstanceGroupOptions{ID: &instanceGroupID}
	_, response, err := sess.GetInstanceGroup(&getInstanceGroupOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error Getting InstanceGroup: %s\n%s", err, response)
	}
	return true, nil
}

func myvpcClient(meta interface{}) (*vpcv1.VpcV1, error) {
	sess, err := meta.(ClientSession).VpcV1APIScoped()
	return sess, err
}
