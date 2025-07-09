// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isLBPoolID                       = "pool"
	isLBPoolMemberPort               = "port"
	isLBPoolMemberTargetAddress      = "target_address"
	isLBPoolMemberTargetID           = "target_id"
	isLBPoolMemberWeight             = "weight"
	isLBPoolMemberProvisioningStatus = "provisioning_status"
	isLBPoolMemberHealth             = "health"
	isLBPoolMemberHref               = "href"
	isLBPoolMemberDeletePending      = "delete_pending"
	isLBPoolMemberDeleted            = "done"
	isLBPoolMemberActive             = "active"
	isLBPoolUpdating                 = "updating"
)

func ResourceIBMISLBPoolMember() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISLBPoolMemberCreate,
		ReadContext:   resourceIBMISLBPoolMemberRead,
		UpdateContext: resourceIBMISLBPoolMemberUpdate,
		DeleteContext: resourceIBMISLBPoolMemberDelete,
		Exists:        resourceIBMISLBPoolMemberExists,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			isLBPoolID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, o, n string, d *schema.ResourceData) bool {
					if o == "" {
						return false
					}
					// if state file entry and tf file entry matches
					if strings.Compare(n, o) == 0 {
						return true
					}

					if strings.Contains(n, "/") {
						new := strings.Split(n, "/")
						if strings.Compare(new[1], o) == 0 {
							return true
						}
					}

					return false
				},
				Description: "Loadblancer Poold ID",
			},

			isLBID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Load balancer ID",
			},

			isLBPoolMemberPort: {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Load Balancer Pool port",
			},

			isLBPoolMemberTargetAddress: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{isLBPoolMemberTargetAddress, isLBPoolMemberTargetID},
				Description:  "Load balancer pool member target address",
			},

			isLBPoolMemberTargetID: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{isLBPoolMemberTargetAddress, isLBPoolMemberTargetID},
				Description:  "Load balancer pool member target id",
			},

			isLBPoolMemberWeight: {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_lb_pool_member", isLBPoolMemberWeight),
				Description:  "Load balcner pool member weight",
			},

			isLBPoolMemberProvisioningStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load balancer Pool member provisioning status",
			},

			isLBPoolMemberHealth: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "LB Pool member health",
			},

			isLBPoolMemberHref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "LB pool member Href value",
			},

			flex.RelatedCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the LB resource",
			},
		},
	}
}

func ResourceIBMISLBPoolMemberValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isLBPoolMemberWeight,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			Optional:                   true,
			MinValue:                   "0",
			MaxValue:                   "100"})

	ibmISLBResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_lb_pool_member", Schema: validateSchema}
	return &ibmISLBResourceValidator
}

func resourceIBMISLBPoolMemberCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	log.Printf("[DEBUG] LB Pool create")
	lbPoolID, err := getPoolId(d.Get(isLBPoolID).(string))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "create", "sep-id-parts").GetDiag()
	}

	lbID := d.Get(isLBID).(string)
	port := d.Get(isLBPoolMemberPort).(int)
	port64 := int64(port)

	var weight int64

	isLBKey := "load_balancer_key_" + lbID
	conns.IbmMutexKV.Lock(isLBKey)
	defer conns.IbmMutexKV.Unlock(isLBKey)

	diag := lbpMemberCreate(context, d, meta, lbID, lbPoolID, port64, weight)
	if diag != nil {
		return diag
	}

	return resourceIBMISLBPoolMemberRead(context, d, meta)
}

func lbpMemberCreate(context context.Context, d *schema.ResourceData, meta interface{}, lbID, lbPoolID string, port, weight int64) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForLBPoolActive(sess, lbID, lbPoolID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBPoolActive failed: %s", err.Error()), "ibm_is_lb_pool_member", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBAvailable failed: %s", err.Error()), "ibm_is_lb_pool_member", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	options := &vpcv1.CreateLoadBalancerPoolMemberOptions{
		LoadBalancerID: &lbID,
		PoolID:         &lbPoolID,
		Port:           &port,
	}

	if _, ok := d.GetOk(isLBPoolMemberTargetAddress); ok {
		targetAddress := d.Get(isLBPoolMemberTargetAddress).(string)
		target := &vpcv1.LoadBalancerPoolMemberTargetPrototype{
			Address: &targetAddress,
		}
		options.Target = target
	} else {
		targetID := d.Get(isLBPoolMemberTargetID).(string)
		target := &vpcv1.LoadBalancerPoolMemberTargetPrototype{
			ID: &targetID,
		}
		options.Target = target
	}
	if w, ok := d.GetOkExists(isLBPoolMemberWeight); ok {
		weight = int64(w.(int))
		options.Weight = &weight
	}

	lbPoolMember, _, err := sess.CreateLoadBalancerPoolMemberWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateLoadBalancerPoolMemberWithContext failed: %s", err.Error()), "ibm_is_lb_pool_member", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", lbID, lbPoolID, *lbPoolMember.ID))
	log.Printf("[INFO] lbpool member : %s", *lbPoolMember.ID)

	_, err = isWaitForLBPoolMemberAvailable(sess, lbID, lbPoolID, *lbPoolMember.ID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBPoolMemberAvailable failed: %s", err.Error()), "ibm_is_lb_pool_member", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForLBPoolActive(sess, lbID, lbPoolID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBPoolActive failed: %s", err.Error()), "ibm_is_lb_pool_member", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBAvailable failed: %s", err.Error()), "ibm_is_lb_pool_member", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return nil
}

func isWaitForLBPoolMemberAvailable(lbc *vpcv1.VpcV1, lbID, lbPoolID, lbPoolMemID string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for load balancer pool member(%s) to be available.", lbPoolMemID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"create_pending", "update_pending", "maintenance_pending"},
		Target:     []string{isLBPoolMemberActive, ""},
		Refresh:    isLBPoolMemberRefreshFunc(lbc, lbID, lbPoolID, lbPoolMemID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLBPoolMemberRefreshFunc(lbc *vpcv1.VpcV1, lbID, lbPoolID, lbPoolMemID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		getlbpmoptions := &vpcv1.GetLoadBalancerPoolMemberOptions{
			LoadBalancerID: &lbID,
			PoolID:         &lbPoolID,
			ID:             &lbPoolMemID,
		}
		lbPoolMem, response, err := lbc.GetLoadBalancerPoolMember(getlbpmoptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error Getting Load Balancer Pool Member: %s\n%s", err, response)
		}

		if *lbPoolMem.ProvisioningStatus == isLBPoolMemberActive {
			return lbPoolMem, *lbPoolMem.ProvisioningStatus, nil
		}

		return lbPoolMem, *lbPoolMem.ProvisioningStatus, nil
	}
}

func resourceIBMISLBPoolMemberRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "read", "sep-id-parts").GetDiag()
	}

	if len(parts) < 3 {
		err = fmt.Errorf(
			"The id should contain loadbalancer Id, loadbalancer pool Id and loadbalancer poolmemebr Id")
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "read", "sep-id-parts").GetDiag()

	}

	lbID := parts[0]
	lbPoolID := parts[1]
	lbPoolMemID := parts[2]

	diag := lbpmemberGet(context, d, meta, lbID, lbPoolID, lbPoolMemID)
	if diag != nil {
		return diag
	}

	return nil
}

func lbpmemberGet(context context.Context, d *schema.ResourceData, meta interface{}, lbID, lbPoolID, lbPoolMemID string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getlbpmoptions := &vpcv1.GetLoadBalancerPoolMemberOptions{
		LoadBalancerID: &lbID,
		PoolID:         &lbPoolID,
		ID:             &lbPoolMemID,
	}
	loadBalancerPoolMember, response, err := sess.GetLoadBalancerPoolMemberWithContext(context, getlbpmoptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLoadBalancerPoolMemberWithContext failed: %s", err.Error()), "ibm_is_lb_pool_member", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(isLBPoolID, lbPoolID); err != nil {
		err = fmt.Errorf("Error setting pool: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "read", "set-pool").GetDiag()
	}
	if err = d.Set(isLBID, lbID); err != nil {
		err = fmt.Errorf("Error setting lb: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "read", "set-lb").GetDiag()
	}
	if err = d.Set("port", flex.IntValue(loadBalancerPoolMember.Port)); err != nil {
		err = fmt.Errorf("Error setting port: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "read", "set-port").GetDiag()
	}

	target := loadBalancerPoolMember.Target.(*vpcv1.LoadBalancerPoolMemberTarget)
	if target.Address != nil {
		if err = d.Set(isLBPoolMemberTargetAddress, *target.Address); err != nil {
			err = fmt.Errorf("Error setting target_address: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "read", "set-target_address").GetDiag()
		}
	}
	if target.ID != nil {
		if err = d.Set(isLBPoolMemberTargetID, *target.ID); err != nil {
			err = fmt.Errorf("Error setting target_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "read", "set-target_id").GetDiag()
		}
	}
	if !core.IsNil(loadBalancerPoolMember.Weight) {
		if err = d.Set("weight", flex.IntValue(loadBalancerPoolMember.Weight)); err != nil {
			err = fmt.Errorf("Error setting weight: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "read", "set-weight").GetDiag()
		}
	}
	if err = d.Set("provisioning_status", loadBalancerPoolMember.ProvisioningStatus); err != nil {
		err = fmt.Errorf("Error setting provisioning_status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "read", "set-provisioning_status").GetDiag()
	}
	if err = d.Set("health", loadBalancerPoolMember.Health); err != nil {
		err = fmt.Errorf("Error setting health: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "read", "set-health").GetDiag()
	}
	if err = d.Set("href", loadBalancerPoolMember.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "read", "set-href").GetDiag()
	}
	getLoadBalancerOptions := &vpcv1.GetLoadBalancerOptions{
		ID: &lbID,
	}
	lb, response, err := sess.GetLoadBalancerWithContext(context, getLoadBalancerOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLoadBalancerWithContext failed: %s", err.Error()), "ibm_is_lb_pool_member", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(flex.RelatedCRN, *lb.CRN); err != nil {
		err = fmt.Errorf("Error setting related_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "read", "set-related_crn").GetDiag()
	}
	return nil
}

func resourceIBMISLBPoolMemberUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "update", "sep-id-parts").GetDiag()
	}

	lbID := parts[0]
	lbPoolID := parts[1]
	lbPoolMemID := parts[2]

	diag := lbpmemberUpdate(context, d, meta, lbID, lbPoolID, lbPoolMemID)
	if diag != nil {
		return diag
	}

	return resourceIBMISLBPoolMemberRead(context, d, meta)
}

func lbpmemberUpdate(context context.Context, d *schema.ResourceData, meta interface{}, lbID, lbPoolID, lbPoolMemID string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if d.HasChange(isLBPoolMemberTargetID) || d.HasChange(isLBPoolMemberTargetAddress) || d.HasChange(isLBPoolMemberPort) || d.HasChange(isLBPoolMemberWeight) {

		port := int64(d.Get(isLBPoolMemberPort).(int))
		weight := int64(d.Get(isLBPoolMemberWeight).(int))

		isLBKey := "load_balancer_key_" + lbID
		conns.IbmMutexKV.Lock(isLBKey)
		defer conns.IbmMutexKV.Unlock(isLBKey)

		_, err = isWaitForLBPoolActive(sess, lbID, lbPoolID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBPoolActive failed: %s", err.Error()), "ibm_is_lb_pool_member", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, err = isWaitForLBPoolMemberAvailable(sess, lbID, lbPoolID, lbPoolMemID, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBPoolMemberAvailable failed: %s", err.Error()), "ibm_is_lb_pool_member", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, err = isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBAvailable failed: %s", err.Error()), "ibm_is_lb_pool_member", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		updatelbpmoptions := &vpcv1.UpdateLoadBalancerPoolMemberOptions{
			LoadBalancerID: &lbID,
			PoolID:         &lbPoolID,
			ID:             &lbPoolMemID,
		}

		loadBalancerPoolMemberPatchModel := &vpcv1.LoadBalancerPoolMemberPatch{
			Port:   &port,
			Weight: &weight,
		}

		if d.HasChange(isLBPoolMemberTargetAddress) {
			targetAddress := d.Get(isLBPoolMemberTargetAddress).(string)
			target := &vpcv1.LoadBalancerPoolMemberTargetPrototypeIP{
				Address: &targetAddress,
			}
			loadBalancerPoolMemberPatchModel.Target = target
		} else if d.HasChange(isLBPoolMemberTargetID) {
			targetID := d.Get(isLBPoolMemberTargetID).(string)
			target := &vpcv1.LoadBalancerPoolMemberTargetPrototype{
				ID: &targetID,
			}
			loadBalancerPoolMemberPatchModel.Target = target
		}

		loadBalancerPoolMemberPatch, err := loadBalancerPoolMemberPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("loadBalancerPoolMemberPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_lb_pool_member", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updatelbpmoptions.LoadBalancerPoolMemberPatch = loadBalancerPoolMemberPatch

		_, _, err = sess.UpdateLoadBalancerPoolMemberWithContext(context, updatelbpmoptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateLoadBalancerPoolMemberWithContext failed: %s", err.Error()), "ibm_is_lb_pool_member", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForLBPoolMemberAvailable(sess, lbID, lbPoolID, lbPoolMemID, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBPoolMemberAvailable failed: %s", err.Error()), "ibm_is_lb_pool_member", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, err = isWaitForLBPoolActive(sess, lbID, lbPoolID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBPoolActive failed: %s", err.Error()), "ibm_is_lb_pool_member", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, err = isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBAvailable failed: %s", err.Error()), "ibm_is_lb_pool_member", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return nil
}

func resourceIBMISLBPoolMemberDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "delete", "sep-id-parts").GetDiag()
	}

	lbID := parts[0]
	lbPoolID := parts[1]
	lbPoolMemID := parts[2]

	isLBKey := "load_balancer_key_" + lbID
	conns.IbmMutexKV.Lock(isLBKey)
	defer conns.IbmMutexKV.Unlock(isLBKey)

	diag := lbpmemberDelete(context, d, meta, lbID, lbPoolID, lbPoolMemID)
	if diag != nil {
		return diag
	}

	return nil
}

func lbpmemberDelete(context context.Context, d *schema.ResourceData, meta interface{}, lbID, lbPoolID, lbPoolMemID string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getlbpmoptions := &vpcv1.GetLoadBalancerPoolMemberOptions{
		LoadBalancerID: &lbID,
		PoolID:         &lbPoolID,
		ID:             &lbPoolMemID,
	}
	_, response, err := sess.GetLoadBalancerPoolMemberWithContext(context, getlbpmoptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLoadBalancerPoolMemberWithContext failed: %s", err.Error()), "ibm_is_lb_pool_member", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForLBPoolMemberAvailable(sess, lbID, lbPoolID, lbPoolMemID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBPoolMemberAvailable failed: %s", err.Error()), "ibm_is_lb_pool_member", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForLBPoolActive(sess, lbID, lbPoolID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBPoolActive failed: %s", err.Error()), "ibm_is_lb_pool_member", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBAvailable failed: %s", err.Error()), "ibm_is_lb_pool_member", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dellbpmoptions := &vpcv1.DeleteLoadBalancerPoolMemberOptions{
		LoadBalancerID: &lbID,
		PoolID:         &lbPoolID,
		ID:             &lbPoolMemID,
	}
	response, err = sess.DeleteLoadBalancerPoolMemberWithContext(context, dellbpmoptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteLoadBalancerPoolMemberWithContext failed: %s", err.Error()), "ibm_is_lb_pool_member", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForLBPoolMemberDeleted(sess, lbID, lbPoolID, lbPoolMemID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBPoolMemberDeleted failed: %s", err.Error()), "ibm_is_lb_pool_member", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForLBPoolActive(sess, lbID, lbPoolID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBPoolActive failed: %s", err.Error()), "ibm_is_lb_pool_member", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBAvailable failed: %s", err.Error()), "ibm_is_lb_pool_member", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")
	return nil
}

func isWaitForLBPoolMemberDeleted(lbc *vpcv1.VpcV1, lbID, lbPoolID, lbPoolMemID string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for  (%s) to be deleted.", lbPoolMemID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isLBPoolMemberDeletePending},
		Target:     []string{isLBPoolMemberDeleted, ""},
		Refresh:    isDeleteLBPoolMemberRefreshFunc(lbc, lbID, lbPoolID, lbPoolMemID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isDeleteLBPoolMemberRefreshFunc(lbc *vpcv1.VpcV1, lbID, lbPoolID, lbPoolMemID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		getlbpmoptions := &vpcv1.GetLoadBalancerPoolMemberOptions{
			LoadBalancerID: &lbID,
			PoolID:         &lbPoolID,
			ID:             &lbPoolMemID,
		}
		lbPoolMem, response, err := lbc.GetLoadBalancerPoolMember(getlbpmoptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return lbPoolMem, isLBPoolMemberDeleted, nil
			}
			return nil, "", fmt.Errorf("[ERROR] Error Deleting Load balancer pool member: %s\n%s", err, response)
		}
		return lbPoolMem, isLBPoolMemberDeletePending, nil
	}
}

func resourceIBMISLBPoolMemberExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return false, flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "delete", "sep-id-parts")

	}
	if len(parts) != 3 {
		err = fmt.Errorf(
			"The id should contain loadbalancer Id, loadbalancer pool Id and loadbalancer poolmemebr Id")
		return false, flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "delete", "sep-id-parts")

	}

	lbID := parts[0]
	lbPoolID := parts[1]
	lbPoolMemID := parts[2]

	exists, err := lbpmemberExists(d, meta, lbID, lbPoolID, lbPoolMemID)
	return exists, err

}

func lbpmemberExists(d *schema.ResourceData, meta interface{}, lbID, lbPoolID, lbPoolMemID string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_pool_member", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}

	getlbpmoptions := &vpcv1.GetLoadBalancerPoolMemberOptions{
		LoadBalancerID: &lbID,
		PoolID:         &lbPoolID,
		ID:             &lbPoolMemID,
	}
	_, response, err := sess.GetLoadBalancerPoolMember(getlbpmoptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLoadBalancerPoolMember failed: %s", err.Error()), "ibm_is_lb_pool_member", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}

func getPoolId(id string) (string, error) {
	if strings.Contains(id, "/") {
		parts, err := flex.IdParts(id)
		if err != nil {
			return "", err
		}

		return parts[1], nil
	} else {
		return id, nil
	}
}
