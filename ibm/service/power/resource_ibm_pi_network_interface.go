// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"

	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
)

func ResourceIBMPINetworkInterface() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPINetworkInterfaceCreate,
		ReadContext:   resourceIBMPINetworkInterfaceRead,
		UpdateContext: resourceIBMPINetworkInterfaceUpdate,
		DeleteContext: resourceIBMPINetworkInterfaceDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
		},
		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v any) error {
				return flex.ResourcePowerUserTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				ForceNew:     true,
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_InstanceID: {
				Description: "If supplied populated it attaches to the instance ID, if empty detaches from the instance ID.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			Arg_IPAddress: {
				Description: "The requested IP address of this network interface.",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeString,
			},
			Arg_Name: {
				Description: "Name of the network interface.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			Arg_NetworkID: {
				Description:  "Network ID.",
				ForceNew:     true,
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_UserTags: {
				Computed:    true,
				Description: "The user tags attached to this resource.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Set:         schema.HashString,
				Type:        schema.TypeSet,
			},
			// Attributes
			Attr_CRN: {
				Computed:    true,
				Description: "The network interface's crn.",
				Type:        schema.TypeString,
			},
			Attr_ExternalIP: {
				Computed:    true,
				Description: "The external ip address for pub-vlan networks.",
				Type:        schema.TypeString,
			},
			Attr_Instance: {
				Computed:    true,
				Optional:    true,
				Description: "The attached instance to this network interface.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_Href: {
							Computed:    true,
							Description: "Link to instance resource.",
							Type:        schema.TypeString,
						},
						Attr_InstanceID: {
							Computed:    true,
							Description: "The attached instance ID.",
							Type:        schema.TypeString,
						},
					},
				},
				Type: schema.TypeList,
			},
			Attr_IPAddress: {
				Computed:    true,
				Description: "The ip address of this network interface.",
				Type:        schema.TypeString,
			},
			Attr_MacAddress: {
				Computed:    true,
				Description: "The mac address of the network interface.",
				Type:        schema.TypeString,
			},
			Attr_Name: {
				Computed:    true,
				Description: "Name of the network interface (not unique or indexable).",
				Type:        schema.TypeString,
			},
			Attr_NetworkInterfaceID: {
				Computed:    true,
				Description: "The unique identifier of the network interface.",
				Type:        schema.TypeString,
			},
			Attr_NetworkSecurityGroupID: {
				Computed:    true,
				Deprecated:  "Deprecated, use network_security_group_ids instead.",
				Description: "ID of the network security group the network interface will be added to.",
				Type:        schema.TypeString,
			},
			Attr_NetworkSecurityGroupIDs: {
				Computed:    true,
				Description: "List of network security groups that the network interface is a member of.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Type:        schema.TypeSet,
			},
			Attr_Status: {
				Computed:    true,
				Description: "The status of the network interface.",
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceIBMPINetworkInterfaceCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "ibm_pi_network_interface", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	networkID := d.Get(Arg_NetworkID).(string)
	networkC := instance.NewIBMPINetworkClient(ctx, sess, cloudInstanceID)
	var body = &models.NetworkInterfaceCreate{}
	if v, ok := d.GetOk(Arg_IPAddress); ok {
		body.IPAddress = v.(string)
	}
	if v, ok := d.GetOk(Arg_Name); ok {
		body.Name = v.(string)
	}
	if v, ok := d.GetOk(Arg_UserTags); ok {
		userTags := flex.FlattenSet(v.(*schema.Set))
		body.UserTags = userTags
	}
	networkInterface, err := networkC.CreateNetworkInterface(networkID, body)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateNetworkInterface failed: %s", err.Error()), "ibm_pi_network_interface", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	networkInterfaceID := *networkInterface.ID
	_, err = isWaitForIBMPINetworkInterfaceAvailable(ctx, networkC, networkID, networkInterfaceID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForIBMPINetworkInterfaceAvailable failed: %s", err.Error()), "ibm_pi_network_interface", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	crn := networkInterface.Crn
	if _, ok := d.GetOk(Arg_UserTags); ok {
		if crn != nil {
			oldList, newList := d.GetChange(Arg_UserTags)
			err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *crn, "", UserTagType)
			if err != nil {
				log.Printf("Error on update of network interface (%s) pi_user_tags: %s", networkInterfaceID, err)
			}
		}
	}
	if v, ok := d.GetOk(Arg_InstanceID); ok {
		instanceID := v.(string)
		body := &models.NetworkInterfaceUpdate{
			InstanceID: &instanceID,
		}
		_, err = networkC.UpdateNetworkInterface(networkID, networkInterfaceID, body)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateNetworkInterface failed: %s", err.Error()), "ibm_pi_network_interface", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForIBMPINetworkInterfaceUpdate(ctx, networkC, networkID, networkInterfaceID, "", instanceID, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForIBMPINetworkInterfaceUpdate failed: %s", err.Error()), "ibm_pi_network_interface", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	d.SetId(fmt.Sprintf("%s/%s/%s", cloudInstanceID, networkID, networkInterfaceID))

	return resourceIBMPINetworkInterfaceRead(ctx, d, meta)
}

func resourceIBMPINetworkInterfaceRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "ibm_pi_network_interface", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IdParts failed: %s", err.Error()), "ibm_pi_network_interface", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	networkC := instance.NewIBMPINetworkClient(ctx, sess, parts[0])
	networkInterface, err := networkC.GetNetworkInterface(parts[1], parts[2])
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), NotFound) {
			log.Printf("[WARNING] network interface resource was not found or removed outside of terraform\n")
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetNetworkInterface failed: %s", err.Error()), "ibm_pi_network_interface", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.Set(Attr_ExternalIP, networkInterface.ExternalIP)
	d.Set(Attr_IPAddress, networkInterface.IPAddress)
	d.Set(Attr_MacAddress, networkInterface.MacAddress)
	d.Set(Attr_Name, networkInterface.Name)
	d.Set(Attr_NetworkInterfaceID, networkInterface.ID)
	d.Set(Attr_NetworkSecurityGroupID, networkInterface.NetworkSecurityGroupID)
	d.Set(Attr_NetworkSecurityGroupIDs, networkInterface.NetworkSecurityGroupIDs)
	if networkInterface.Instance != nil {
		pvmInstance := []map[string]any{}
		instanceMap := pvmInstanceToMap(networkInterface.Instance)
		pvmInstance = append(pvmInstance, instanceMap)
		d.Set(Attr_Instance, pvmInstance)
	} else {
		d.Set(Attr_Instance, nil)
	}
	d.Set(Attr_Status, networkInterface.Status)
	if networkInterface.Crn != nil {
		d.Set(Attr_CRN, networkInterface.Crn)
		tags, err := flex.GetGlobalTagsUsingCRN(meta, string(*networkInterface.Crn), "", UserTagType)
		if err != nil {
			log.Printf("Error on get of network interface (%s) pi_user_tags: %s", *networkInterface.ID, err)
		}
		d.Set(Arg_UserTags, tags)
	}

	return nil
}

func resourceIBMPINetworkInterfaceUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "ibm_pi_network_interface", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IdParts failed: %s", err.Error()), "ibm_pi_network_interface", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	networkC := instance.NewIBMPINetworkClient(ctx, sess, parts[0])
	body := &models.NetworkInterfaceUpdate{}

	hasChange := false
	if d.HasChange(Arg_UserTags) {
		if crn, ok := d.GetOk(Attr_CRN); ok {
			oldList, newList := d.GetChange(Arg_UserTags)
			err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, crn.(string), "", UserTagType)
			if err != nil {
				log.Printf("Error on update of network interface (%s) pi_user_tags: %s", parts[2], err)
			}
		}
	}
	if d.HasChange(Arg_Name) {
		name := d.Get(Arg_Name).(string)
		body.Name = &name
		hasChange = true
	}
	if d.HasChange(Arg_InstanceID) {
		instanceID := d.Get(Arg_InstanceID).(string)
		body.InstanceID = &instanceID
		hasChange = true
	}

	if hasChange {
		_, err = networkC.UpdateNetworkInterface(parts[1], parts[2], body)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateNetworkInterface failed: %s", err.Error()), "ibm_pi_network_interface", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		expectedName := ""
		if d.HasChange(Arg_Name) {
			expectedName = d.Get(Arg_Name).(string)
		}
		expectedInstanceID := ""
		if d.HasChange(Arg_InstanceID) {
			expectedInstanceID = d.Get(Arg_InstanceID).(string)
		}
		_, err = isWaitForIBMPINetworkInterfaceUpdate(ctx, networkC, parts[1], parts[2], expectedName, expectedInstanceID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForIBMPINetworkInterfaceUpdate failed: %s", err.Error()), "ibm_pi_network_interface", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMPINetworkInterfaceRead(ctx, d, meta)
}

func resourceIBMPINetworkInterfaceDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "ibm_pi_network_interface", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IdParts failed: %s", err.Error()), "ibm_pi_network_interface", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	networkC := instance.NewIBMPINetworkClient(ctx, sess, parts[0])
	err = networkC.DeleteNetworkInterface(parts[1], parts[2])
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteNetworkInterface failed: %s", err.Error()), "ibm_pi_network_interface", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForIBMPINetworkInterfaceDeleted(ctx, networkC, parts[1], parts[2], d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForIBMPINetworkInterfaceDeleted failed: %s", err.Error()), "ibm_pi_network_interface", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")

	return nil
}

func isWaitForIBMPINetworkInterfaceAvailable(ctx context.Context, client *instance.IBMPINetworkClient, networkID string, networkInterfaceID string, timeout time.Duration) (any, error) {
	stateConf := &retry.StateChangeConf{
		Pending:    []string{State_Build},
		Target:     []string{State_Down},
		Refresh:    isIBMPINetworkInterfaceRefreshFunc(client, networkID, networkInterfaceID),
		Timeout:    timeout,
		Delay:      5 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isIBMPINetworkInterfaceRefreshFunc(client *instance.IBMPINetworkClient, networkID, networkInterfaceID string) retry.StateRefreshFunc {
	return func() (any, string, error) {
		networkInterface, err := client.GetNetworkInterface(networkID, networkInterfaceID)
		if err != nil {
			return nil, "", err
		}
		if strings.ToLower(*networkInterface.Status) == State_Down {
			return networkInterface, State_Down, nil
		}
		return networkInterface, State_Build, nil
	}
}

func isWaitForIBMPINetworkInterfaceUpdate(ctx context.Context, client *instance.IBMPINetworkClient, networkID, networkInterfaceID, expectedName, expectedInstanceID string, timeout time.Duration) (any, error) {
	stateConf := &retry.StateChangeConf{
		Pending:    []string{State_Build},
		Target:     []string{State_Active},
		Refresh:    isIBMPINetworkInterfaceUpdateRefreshFunc(client, networkID, networkInterfaceID, expectedName, expectedInstanceID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForStateContext(ctx)
}

func isIBMPINetworkInterfaceUpdateRefreshFunc(client *instance.IBMPINetworkClient, networkID, networkInterfaceID, expectedName, expectedInstanceID string) retry.StateRefreshFunc {
	return func() (any, string, error) {
		networkInterface, err := client.GetNetworkInterface(networkID, networkInterfaceID)
		if err != nil {
			return nil, "", err
		}
		nameOK := expectedName == "" || (networkInterface.Name != nil && *networkInterface.Name == expectedName)
		instanceOK := expectedInstanceID == "" || (strings.ToLower(*networkInterface.Status) == State_Active && networkInterface.Instance != nil && networkInterface.Instance.InstanceID == expectedInstanceID)
		if nameOK && instanceOK {
			return networkInterface, State_Active, nil
		}
		return networkInterface, State_Build, nil
	}
}

func isWaitForIBMPINetworkInterfaceDeleted(ctx context.Context, client *instance.IBMPINetworkClient, networkID, networkInterfaceID string, timeout time.Duration) (any, error) {
	stateConf := &retry.StateChangeConf{
		Pending:    []string{State_Found},
		Target:     []string{State_NotFound},
		Refresh:    isIBMPINetworkInterfaceDeleteRefreshFunc(client, networkID, networkInterfaceID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForStateContext(ctx)
}

func isIBMPINetworkInterfaceDeleteRefreshFunc(client *instance.IBMPINetworkClient, networkID, networkInterfaceID string) retry.StateRefreshFunc {
	return func() (any, string, error) {
		networkInterface, err := client.GetNetworkInterface(networkID, networkInterfaceID)
		if err != nil {
			return networkInterface, State_NotFound, nil
		}
		return networkInterface, State_Found, nil
	}
}
