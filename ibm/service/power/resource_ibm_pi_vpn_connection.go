// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_v_p_n_connections"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
)

const (
	// Arguments
	PIVpnIkePolicyID        = "pi_ike_policy_id"
	PIVpnIpSecPolicyID      = "pi_ipsec_policy_id"
	PIVpnNetworks           = "pi_networks"
	PIVpnPeerGatewayAddress = "pi_peer_gateway_address"
	PIVpnSubnets            = "pi_peer_subnets"
	PIVpnMode               = "pi_vpn_connection_mode"
	PIVpnName               = "pi_vpn_connection_name"

	// Attributes
	VPNConnectionID        = "connection_id"
	VPNConnectionStatus    = "connection_status"
	VPNDeadPeer            = "dead_peer_detections"
	VPNAction              = "action"
	VPNInterval            = "interval"
	VPNThreshold           = "threshold"
	VPNGatewayAddress      = "gateway_address"
	VPNLocalGatewayAddress = "local_gateway_addreess"
)

func ResourceIBMPIVPNConnection() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIVPNConnectionCreate,
		ReadContext:   resourceIBMPIVPNConnectionRead,
		UpdateContext: resourceIBMPIVPNConnectionUpdate,
		DeleteContext: resourceIBMPIVPNConnectionDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			// Required Attributes
			PICloudInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PI cloud instance ID",
			},
			PIVpnName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the VPN Connection",
			},
			PIVpnIkePolicyID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier of IKE Policy selected for this VPN Connection",
			},
			PIVpnIpSecPolicyID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier of IPSec Policy selected for this VPN Connection",
			},
			PIVpnMode: {
				Type:             schema.TypeString,
				Required:         true,
				ValidateFunc:     validate.ValidateAllowedStringValues([]string{"policy", "route"}),
				Description:      "Mode used by this VPN Connection, either 'policy' or 'route'",
				DiffSuppressFunc: flex.ApplyOnce,
			},
			PIVpnNetworks: {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Set of network IDs to attach to this VPN connection",
			},
			PIVpnPeerGatewayAddress: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Peer Gateway address",
			},
			PIVpnSubnets: {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Set of CIDR of peer subnets",
			},

			//Computed Attributes
			VPNConnectionID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "VPN connection ID",
			},
			VPNLocalGatewayAddress: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Local Gateway address, only in 'route' mode",
			},
			VPNConnectionStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the VPN connection",
			},
			VPNGatewayAddress: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Public IP address of the VPN Gateway (vSRX) attached to this VPN Connection",
			},
			VPNDeadPeer: {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Dead Peer Detection",
			},
		},
	}
}

func resourceIBMPIVPNConnectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	name := d.Get(PIVpnName).(string)
	ikePolicyId := d.Get(PIVpnIkePolicyID).(string)
	ipsecPolicyId := d.Get(PIVpnIpSecPolicyID).(string)
	mode := d.Get(PIVpnMode).(string)
	networks := d.Get(PIVpnNetworks).(*schema.Set)
	peerSubnets := d.Get(PIVpnSubnets).(*schema.Set)
	peerGatewayAddress := d.Get(PIVpnPeerGatewayAddress).(string)
	pga := models.PeerGatewayAddress(peerGatewayAddress)

	body := &models.VPNConnectionCreate{
		IkePolicy:          &ikePolicyId,
		IPSecPolicy:        &ipsecPolicyId,
		Mode:               &mode,
		Name:               &name,
		PeerGatewayAddress: &pga,
	}
	// networks
	if networks.Len() > 0 {
		body.Networks = flex.ExpandStringList(networks.List())
	} else {
		return diag.Errorf("%s is a required field", PIVpnNetworks)
	}
	// peer subnets
	if peerSubnets.Len() > 0 {
		body.PeerSubnets = flex.ExpandStringList(peerSubnets.List())
	} else {
		return diag.Errorf("%s is a required field", PIVpnSubnets)
	}

	client := st.NewIBMPIVpnConnectionClient(ctx, sess, cloudInstanceID)
	vpnConnection, err := client.Create(body)
	if err != nil {
		log.Printf("[DEBUG] create VPN connection failed %v", err)
		return diag.FromErr(err)
	}

	vpnConnectionId := *vpnConnection.ID
	d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, vpnConnectionId))

	if vpnConnection.JobRef != nil {
		jobID := *vpnConnection.JobRef.ID
		jobClient := st.NewIBMPIJobClient(ctx, sess, cloudInstanceID)

		_, err = waitForIBMPIJobCompleted(ctx, jobClient, jobID, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceIBMPIVPNConnectionRead(ctx, d, meta)
}

func resourceIBMPIVPNConnectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, vpnConnectionID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := st.NewIBMPIVpnConnectionClient(ctx, sess, cloudInstanceID)
	jobClient := st.NewIBMPIJobClient(ctx, sess, cloudInstanceID)

	if d.HasChangesExcept(PIVpnNetworks, PIVpnSubnets) {
		body := &models.VPNConnectionUpdate{}

		if d.HasChanges(PIVpnName) {
			name := d.Get(PIVpnName).(string)
			body.Name = name
		}
		if d.HasChanges(PIVpnIkePolicyID) {
			ikePolicyId := d.Get(PIVpnIkePolicyID).(string)
			body.IkePolicy = ikePolicyId
		}
		if d.HasChanges(PIVpnIpSecPolicyID) {
			ipsecPolicyId := d.Get(PIVpnIpSecPolicyID).(string)
			body.IPSecPolicy = ipsecPolicyId
		}
		if d.HasChanges(PIVpnPeerGatewayAddress) {
			peerGatewayAddress := d.Get(PIVpnPeerGatewayAddress).(string)
			body.PeerGatewayAddress = models.PeerGatewayAddress(peerGatewayAddress)
		}

		_, err = client.Update(vpnConnectionID, body)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	if d.HasChanges(PIVpnNetworks) {
		oldRaw, newRaw := d.GetChange(PIVpnNetworks)
		old := oldRaw.(*schema.Set)
		new := newRaw.(*schema.Set)

		toAdd := new.Difference(old)
		toRemove := old.Difference(new)

		for _, n := range flex.ExpandStringList(toAdd.List()) {
			jobReference, err := client.AddNetwork(vpnConnectionID, n)
			if err != nil {
				return diag.FromErr(err)
			}
			if jobReference != nil {
				_, err = waitForIBMPIJobCompleted(ctx, jobClient, *jobReference.ID, d.Timeout(schema.TimeoutUpdate))
				if err != nil {
					return diag.FromErr(err)
				}
			}
		}
		for _, n := range flex.ExpandStringList(toRemove.List()) {
			jobReference, err := client.DeleteNetwork(vpnConnectionID, n)
			if err != nil {
				return diag.FromErr(err)
			}
			if jobReference != nil {
				_, err = waitForIBMPIJobCompleted(ctx, jobClient, *jobReference.ID, d.Timeout(schema.TimeoutUpdate))
				if err != nil {
					return diag.FromErr(err)
				}
			}
		}

	}
	if d.HasChanges(PIVpnSubnets) {
		oldRaw, newRaw := d.GetChange(PIVpnSubnets)
		old := oldRaw.(*schema.Set)
		new := newRaw.(*schema.Set)

		toAdd := new.Difference(old)
		toRemove := old.Difference(new)

		for _, s := range flex.ExpandStringList(toAdd.List()) {
			_, err := client.AddSubnet(vpnConnectionID, s)
			if err != nil {
				return diag.FromErr(err)
			}
		}
		for _, s := range flex.ExpandStringList(toRemove.List()) {
			_, err := client.DeleteSubnet(vpnConnectionID, s)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}
	return resourceIBMPIVPNConnectionRead(ctx, d, meta)
}

func resourceIBMPIVPNConnectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, vpnConnectionID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := st.NewIBMPIVpnConnectionClient(ctx, sess, cloudInstanceID)
	vpnConnection, err := client.Get(vpnConnectionID)
	if err != nil {
		uErr := errors.Unwrap(err)
		switch uErr.(type) {
		case *p_cloud_v_p_n_connections.PcloudVpnconnectionsGetNotFound:
			log.Printf("[DEBUG] VPN connection does not exist %v", err)
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] get VPN connection failed %v", err)
		return diag.FromErr(err)
	}

	d.Set(VPNConnectionID, vpnConnection.ID)
	d.Set(PIVpnName, vpnConnection.Name)
	if vpnConnection.IkePolicy != nil {
		d.Set(PIVpnIkePolicyID, vpnConnection.IkePolicy.ID)
	}
	if vpnConnection.IPSecPolicy != nil {
		d.Set(PIVpnIpSecPolicyID, vpnConnection.IPSecPolicy.ID)
	}
	d.Set(VPNLocalGatewayAddress, vpnConnection.LocalGatewayAddress)
	d.Set(PIVpnMode, vpnConnection.Mode)
	d.Set(PIVpnPeerGatewayAddress, vpnConnection.PeerGatewayAddress)
	d.Set(VPNConnectionStatus, vpnConnection.Status)
	d.Set(VPNGatewayAddress, vpnConnection.VpnGatewayAddress)

	d.Set(PIVpnNetworks, vpnConnection.NetworkIDs)
	d.Set(PIVpnSubnets, vpnConnection.PeerSubnets)

	if vpnConnection.DeadPeerDetection != nil {
		dpc := vpnConnection.DeadPeerDetection
		dpcMap := map[string]interface{}{
			VPNAction:    *dpc.Action,
			VPNInterval:  strconv.FormatInt(*dpc.Interval, 10),
			VPNThreshold: strconv.FormatInt(*dpc.Threshold, 10),
		}
		d.Set(VPNDeadPeer, dpcMap)
	}

	return nil
}

func resourceIBMPIVPNConnectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, vpnConnectionID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := st.NewIBMPIVpnConnectionClient(ctx, sess, cloudInstanceID)
	jobClient := st.NewIBMPIJobClient(ctx, sess, cloudInstanceID)

	jobRef, err := client.Delete(vpnConnectionID)
	if err != nil {
		uErr := errors.Unwrap(err)
		switch uErr.(type) {
		case *p_cloud_v_p_n_connections.PcloudVpnconnectionsDeleteNotFound:
			log.Printf("[DEBUG] VPN connection does not exist %v", err)
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] delete VPN connection failed %v", err)
		return diag.FromErr(err)
	}
	if jobRef != nil {
		jobID := *jobRef.ID
		_, err = waitForIBMPIJobCompleted(ctx, jobClient, jobID, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId("")
	return nil
}
