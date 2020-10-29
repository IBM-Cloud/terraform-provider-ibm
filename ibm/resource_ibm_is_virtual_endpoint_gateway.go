package ibm

import (
	"log"
	"time"

	"github.com/IBM/go-sdk-core/core"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	vpcscoped "github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

const (
	isVirtualEndpointGatewayName               = "name"
	isVirtualEndpointGatewayResourceType       = "resource_type"
	isVirtualEndpointGatewayResourceGroupID    = "resource_group"
	isVirtualEndpointGatewayCreatedAt          = "created_at"
	isVirtualEndpointGatewayIPs                = "ips"
	isVirtualEndpointGatewayIPsID              = "id"
	isVirtualEndpointGatewayIPsAddress         = "address"
	isVirtualEndpointGatewayIPsName            = "name"
	isVirtualEndpointGatewayIPsSubnet          = "subnet"
	isVirtualEndpointGatewayIPsSubnetID        = "subnet_id"
	isVirtualEndpointGatewayIPsResourceType    = "resource_type"
	isVirtualEndpointGatewayHealthState        = "health_state"
	isVirtualEndpointGatewayLifecycleState     = "lifecycle_state"
	isVirtualEndpointGatewayTarget             = "target"
	isVirtualEndpointGatewayTargetName         = "name"
	isVirtualEndpointGatewayTargetResourceType = "resource_type"
	isVirtualEndpointGatewayVpcID              = "vpc"
)

func resourceIBMISEndpointGateway() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMisVirtualEndpointGatewayCreate,
		Read:     resourceIBMisVirtualEndpointGatewayRead,
		Update:   resourceIBMisVirtualEndpointGatewayUpdate,
		Delete:   resourceIBMisVirtualEndpointGatewayDelete,
		Exists:   resourceIBMisVirtualEndpointGatewayExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			isVirtualEndpointGatewayName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateISName,
				Description:  "Endpoint gateway name",
			},
			isVirtualEndpointGatewayResourceType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Endpoint gateway resource type",
			},
			isVirtualEndpointGatewayResourceGroupID: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The resource group id",
			},
			isVirtualEndpointGatewayCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Endpoint gateway created date and time",
			},
			isVirtualEndpointGatewayHealthState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Endpoint gateway health state",
			},
			isVirtualEndpointGatewayLifecycleState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Endpoint gateway lifecycle state",
			},
			isVirtualEndpointGatewayIPs: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Endpoint gateway resource group",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isVirtualEndpointGatewayIPsID: {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The IPs id",
						},
						isVirtualEndpointGatewayIPsName: {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The IPs name",
						},
						isVirtualEndpointGatewayIPsSubnetID: {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The Subnet id",
						},
						isVirtualEndpointGatewayIPsResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The VPC Resource Type",
						},
					},
				},
			},
			isVirtualEndpointGatewayTarget: {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: "Endpoint gateway target",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isVirtualEndpointGatewayTargetName: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The target name",
						},
						isVirtualEndpointGatewayTargetResourceType: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The target resource type",
						},
					},
				},
			},
			isVirtualEndpointGatewayVpcID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The VPC id",
			},
		},
	}
}

func resourceIBMisVirtualEndpointGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}

	name := d.Get(isVirtualEndpointGatewayName).(string)

	// target opiton
	target := (d.Get(isVirtualEndpointGatewayTarget).([]interface{}))[0].(map[string]interface{})
	taretName := target[isVirtualEndpointGatewayTargetName].(string)
	targetResourceType := target[isVirtualEndpointGatewayResourceType].(string)
	targetOpt := &vpcscoped.EndpointGatewayTargetPrototype{
		Name:         core.StringPtr(taretName),
		ResourceType: core.StringPtr(targetResourceType),
	}

	// vpc option
	vpcID := d.Get(isVirtualEndpointGatewayVpcID).(string)
	vpcOpt := &vpcscoped.VPCIdentity{
		ID: core.StringPtr(vpcID),
	}

	// update option
	opt := sess.NewCreateEndpointGatewayOptions(targetOpt, vpcOpt)
	opt.SetName(name)
	opt.SetTarget(targetOpt)
	opt.SetVPC(vpcOpt)

	// IPs option
	if ips, ok := d.GetOk(isVirtualEndpointGatewayIPs); ok {
		opt.SetIps(expandIPs(ips.([]interface{})))
	}

	// Resource group option
	if resourceGroup, ok := d.GetOk(isVirtualEndpointGatewayResourceGroupID); ok {
		resourceGroupID := resourceGroup.(string)

		resourceGroupOpt := &vpcscoped.ResourceGroupIdentity{
			ID: core.StringPtr(resourceGroupID),
		}
		opt.SetResourceGroup(resourceGroupOpt)

	}
	result, response, err := sess.CreateEndpointGateway(opt)
	if err != nil {
		log.Printf("Create Endpoint Gateway failed: %v", response)
		return err
	}

	d.SetId(*result.ID)
	return resourceIBMisVirtualEndpointGatewayRead(d, meta)
}

func resourceIBMisVirtualEndpointGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}

	if d.HasChange(isVirtualEndpointGatewayName) {
		name := d.Get(isVirtualEndpointGatewayName).(string)

		// create option
		endpointGatewayPatchModel := new(vpcscoped.EndpointGatewayPatch)
		endpointGatewayPatchModel.Name = core.StringPtr(name)
		endpointGatewayPatchModelAsPatch, _ := endpointGatewayPatchModel.AsPatch()
		opt := sess.NewUpdateEndpointGatewayOptions(d.Id(), endpointGatewayPatchModelAsPatch)
		_, response, err := sess.UpdateEndpointGateway(opt)
		if err != nil {
			log.Printf("Update Endpoint Gateway failed: %v", response)
			return err
		}

	}
	return resourceIBMisVirtualEndpointGatewayRead(d, meta)
}

func resourceIBMisVirtualEndpointGatewayRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}
	// read option
	opt := sess.NewGetEndpointGatewayOptions(d.Id())
	result, response, err := sess.GetEndpointGateway(opt)
	if err != nil {
		log.Printf("Get Endpoint Gateway failed: %v", response)
		return err
	}
	d.Set(isVirtualEndpointGatewayName, result.Name)
	d.Set(isVirtualEndpointGatewayCreatedAt, result.CreatedAt)
	d.Set(isVirtualEndpointGatewayHealthState, result.HealthState)
	d.Set(isVirtualEndpointGatewayCreatedAt, result.CreatedAt.String())
	d.Set(isVirtualEndpointGatewayLifecycleState, result.LifecycleState)
	d.Set(isVirtualEndpointGatewayResourceType, result.ResourceType)
	d.Set(isVirtualEndpointGatewayIPs, flattenIPs(result.Ips))
	d.Set(isVirtualEndpointGatewayResourceGroupID, result.ResourceGroup.ID)
	d.Set(isVirtualEndpointGatewayTarget, flattenEndpointGatewayTarget(result.Target.(*vpcscoped.EndpointGatewayTarget)))
	d.Set(isVirtualEndpointGatewayVpcID, result.VPC.ID)
	return nil
}

func resourceIBMisVirtualEndpointGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}
	// read option
	opt := sess.NewDeleteEndpointGatewayOptions(d.Id())
	response, err := sess.DeleteEndpointGateway(opt)
	if err != nil {
		log.Printf("Delete Endpoint Gateway failed: %v", response)
	}
	return nil
}

func resourceIBMisVirtualEndpointGatewayExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := myvpcClient(meta)
	if err != nil {
		return false, err
	}
	// read option
	opt := sess.NewGetEndpointGatewayOptions(d.Id())
	_, response, err := sess.GetEndpointGateway(opt)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			log.Printf("Endpoint Gateway does not exist.")
			return false, nil
		}
		log.Printf("Error : %s", response)
		return false, err
	}
	return true, nil
}

func expandIPs(ipsSet []interface{}) (ipsOptions []vpcscoped.EndpointGatewayReservedIPIntf) {
	ipsList := ipsSet
	for _, item := range ipsList {
		ips := item.(map[string]interface{})
		// IPs option
		ipsID := ips[isVirtualEndpointGatewayIPsID].(string)
		ipsName := ips[isVirtualEndpointGatewayIPsName].(string)

		// IPs subnet option
		ipsSubnetID := ips[isVirtualEndpointGatewayIPsSubnetID].(string)

		ipsSubnetOpt := &vpcscoped.SubnetIdentity{
			ID: &ipsSubnetID,
		}

		ipsOpt := &vpcscoped.EndpointGatewayReservedIP{
			ID:     core.StringPtr(ipsID),
			Name:   core.StringPtr(ipsName),
			Subnet: ipsSubnetOpt,
		}
		ipsOptions = append(ipsOptions, ipsOpt)
	}
	return ipsOptions
}

func flattenIPs(ipsList []vpcscoped.ReservedIPReference) interface{} {
	ipsListOutput := make([]interface{}, 0)
	for _, item := range ipsList {
		ips := make(map[string]interface{}, 0)
		ips[isVirtualEndpointGatewayIPsID] = *item.ID
		ips[isVirtualEndpointGatewayIPsName] = *item.Name
		ips[isVirtualEndpointGatewayIPsResourceType] = *item.ResourceType

		ipsListOutput = append(ipsListOutput, ips)
	}
	return ipsListOutput
}

func flattenEndpointGatewayTarget(target *vpcscoped.EndpointGatewayTarget) interface{} {
	targetSlice := []interface{}{}
	targetOutput := map[string]string{}
	if target == nil {
		return targetOutput
	}
	if target.Name != nil {
		targetOutput[isVirtualEndpointGatewayTargetName] = *target.Name
	}
	if target.ResourceType != nil {
		targetOutput[isVirtualEndpointGatewayTargetResourceType] = *target.ResourceType
	}
	targetSlice = append(targetSlice, targetOutput)
	return targetSlice
}
