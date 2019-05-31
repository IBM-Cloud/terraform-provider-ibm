package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/vpn"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

const (
	isVPNGatewayConnectionAdminStateup              = "admin_state_up"
	isVPNGatewayConnectionName                      = "name"
	isVPNGatewayConnectionVPNGateway                = "vpn_gateway"
	isVPNGatewayConnectionPeerAddress               = "peer_address"
	isVPNGatewayConnectionPreSharedKey              = "preshared_key"
	isVPNGatewayConnectionLocalCIDRS                = "local_cidrs"
	isVPNGatewayConnectionPeerCIDRS                 = "peer_cidrs"
	isVPNGatewayConnectionIKEPolicy                 = "ike_policy"
	isVPNGatewayConnectionIPSECPolicy               = "ipsec_policy"
	isVPNGatewayConnectionDeadPeerDetectionAction   = "action"
	isVPNGatewayConnectionDeadPeerDetectionInterval = "interval"
	isVPNGatewayConnectionDeadPeerDetectionTimeout  = "timeout"
	isVPNGatewayConnectionStatus                    = "status"
	isVPNGatewayConnectionDeleting                  = "deleting"
	isVPNGatewayConnectionDeleted                   = "done"
	isVPNGatewayConnectionProvisioning              = "provisioning"
	isVPNGatewayConnectionProvisioningDone          = "done"
)

func resourceIBMISVPNGatewayConnection() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISVPNGatewayConnectionCreate,
		Read:     resourceIBMISVPNGatewayConnectionRead,
		Update:   resourceIBMISVPNGatewayConnectionUpdate,
		Delete:   resourceIBMISVPNGatewayConnectionDelete,
		Exists:   resourceIBMISVPNGatewayConnectionExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			isVPNGatewayConnectionName: {
				Type:     schema.TypeString,
				Required: true,
			},

			isVPNGatewayConnectionVPNGateway: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			isVPNGatewayConnectionPeerAddress: {
				Type:     schema.TypeString,
				Required: true,
			},

			isVPNGatewayConnectionPreSharedKey: {
				Type:     schema.TypeString,
				Required: true,
			},

			isVPNGatewayConnectionAdminStateup: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			isVPNGatewayConnectionLocalCIDRS: {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			isVPNGatewayConnectionPeerCIDRS: {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			isVPNGatewayConnectionDeadPeerDetectionAction: {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: validateAllowedStringValue([]string{"restart", "clear", "hold", "none"}),
			},
			isVPNGatewayConnectionDeadPeerDetectionInterval: {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      30,
				ValidateFunc: validateDeadPeerDetectionInterval,
			},
			isVPNGatewayConnectionDeadPeerDetectionTimeout: {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      120,
				ValidateFunc: validateDeadPeerDetectionInterval,
			},

			isVPNGatewayConnectionIPSECPolicy: {
				Type:     schema.TypeString,
				Optional: true,
			},

			isVPNGatewayConnectionIKEPolicy: {
				Type:     schema.TypeString,
				Optional: true,
			},

			isVPNGatewayConnectionStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMISVPNGatewayConnectionCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] VPNGatewayConnection create")
	name := d.Get(isVPNGatewayConnectionName).(string)
	gatewayID := d.Get(isVPNGatewayConnectionVPNGateway).(string)
	peerAddress := d.Get(isVPNGatewayConnectionPeerAddress).(string)
	prephasedKey := d.Get(isVPNGatewayConnectionPreSharedKey).(string)
	stateUp := d.Get(isVPNGatewayConnectionAdminStateup).(bool)
	action := d.Get(isVPNGatewayConnectionDeadPeerDetectionAction).(string)
	interval := int64(d.Get(isVPNGatewayConnectionDeadPeerDetectionInterval).(int))
	timeout := int64(d.Get(isVPNGatewayConnectionDeadPeerDetectionTimeout).(int))

	localCidrs := expandStringList((d.Get(isVPNGatewayConnectionLocalCIDRS).(*schema.Set)).List())
	peerCidrs := expandStringList((d.Get(isVPNGatewayConnectionPeerCIDRS).(*schema.Set)).List())

	deadPeerDetection := &models.VPNGatewayConnectionDPD{
		Action:   &action,
		Interval: &interval,
		Timeout:  &timeout,
	}

	var ikePolicyIdentity *models.IKEPolicyIdentity
	var ipsecPolicyIdentity *models.IpsecPolicyIdentity

	if ikePolicy, ok := d.GetOk(isVPNGatewayConnectionIKEPolicy); ok {
		ikePolicyIdentity = &models.IKEPolicyIdentity{
			ID: strfmt.UUID(ikePolicy.(string)),
		}
	} else {
		ikePolicyIdentity = nil
	}

	if ipsecPolicy, ok := d.GetOk(isVPNGatewayConnectionIPSECPolicy); ok {
		ipsecPolicyIdentity = &models.IpsecPolicyIdentity{
			ID: strfmt.UUID(ipsecPolicy.(string)),
		}
	} else {
		ipsecPolicyIdentity = nil
	}

	VPNGatewayConnectionC := vpn.NewVpnClient(sess)

	VPNGatewayConnection, err := VPNGatewayConnectionC.CreateConnections(gatewayID, name, peerAddress, prephasedKey, peerCidrs, localCidrs, stateUp, deadPeerDetection, ikePolicyIdentity, ipsecPolicyIdentity)
	if err != nil {
		log.Printf("[DEBUG] VPNGatewayConnection err %s", isErrorToString(err))
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", gatewayID, VPNGatewayConnection.ID.String()))
	log.Printf("[INFO] VPNGatewayConnection : %s/%s", gatewayID, VPNGatewayConnection.ID.String())
	return resourceIBMISVPNGatewayConnectionRead(d, meta)
}

func resourceIBMISVPNGatewayConnectionRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	gID := parts[0]
	gConnID := parts[1]

	VPNGatewayConnectionC := vpn.NewVpnClient(sess)

	VPNGatewayConnection, err := VPNGatewayConnectionC.GetConnection(gID, gConnID)
	if err != nil {
		return err
	}

	d.Set(isVPNGatewayConnectionName, VPNGatewayConnection.Name)
	d.Set(isVPNGatewayConnectionVPNGateway, gID)
	d.Set(isVPNGatewayConnectionAdminStateup, VPNGatewayConnection.AdminStateUp)
	d.Set(isVPNGatewayConnectionPeerAddress, VPNGatewayConnection.PeerAddress)
	d.Set(isVPNGatewayConnectionPreSharedKey, VPNGatewayConnection.Psk)
	d.Set(isVPNGatewayConnectionLocalCIDRS, flattenStringList(VPNGatewayConnection.LocalCidrs))
	d.Set(isVPNGatewayConnectionPeerCIDRS, flattenStringList(VPNGatewayConnection.PeerCidrs))
	if VPNGatewayConnection.IkePolicy != nil {
		d.Set(isVPNGatewayConnectionIKEPolicy, VPNGatewayConnection.IkePolicy.ID.String())
	}
	if VPNGatewayConnection.IpsecPolicy != nil {
		d.Set(isVPNGatewayConnectionIPSECPolicy, VPNGatewayConnection.IpsecPolicy.ID.String())
	}
	d.Set(isVPNGatewayConnectionDeadPeerDetectionAction, VPNGatewayConnection.DeadPeerDetection.Action)
	d.Set(isVPNGatewayConnectionDeadPeerDetectionInterval, VPNGatewayConnection.DeadPeerDetection.Interval)
	d.Set(isVPNGatewayConnectionDeadPeerDetectionTimeout, VPNGatewayConnection.DeadPeerDetection.Timeout)
	return nil
}

func resourceIBMISVPNGatewayConnectionUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	VPNGatewayConnectionC := vpn.NewVpnClient(sess)

	var name, peerAddress, psk string
	var deadPeerDetection *models.VPNGatewayConnectionDPD
	var ikePolicy *models.IKEPolicyIdentity
	var ipsecPolicy *models.IpsecPolicyIdentity
	hasChange := false
	adminStateUp := d.Get(isVPNGatewayConnectionAdminStateup).(bool)

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	gID := parts[0]
	gConnID := parts[1]

	if d.HasChange(isVPNGatewayConnectionName) {
		name = d.Get(isVPNGatewayConnectionName).(string)
		hasChange = true
	}

	if d.HasChange(isVPNGatewayConnectionPeerAddress) {
		peerAddress = d.Get(isVPNGatewayConnectionPeerAddress).(string)
		hasChange = true
	}

	if d.HasChange(isVPNGatewayConnectionPreSharedKey) {
		psk = d.Get(isVPNGatewayConnectionPreSharedKey).(string)
		hasChange = true
	}

	if d.HasChange(isVPNGatewayConnectionDeadPeerDetectionAction) || d.HasChange(isVPNGatewayConnectionDeadPeerDetectionInterval) || d.HasChange(isVPNGatewayConnectionDeadPeerDetectionTimeout) {
		action := d.Get(isVPNGatewayConnectionDeadPeerDetectionAction).(string)
		interval := int64(d.Get(isVPNGatewayConnectionDeadPeerDetectionInterval).(int))
		timeout := int64(d.Get(isVPNGatewayConnectionDeadPeerDetectionTimeout).(int))
		deadPeerDetection = &models.VPNGatewayConnectionDPD{
			Action:   &action,
			Interval: &interval,
			Timeout:  &timeout,
		}
		hasChange = true
	} else {
		deadPeerDetection = nil
	}

	if d.HasChange(isVPNGatewayConnectionIKEPolicy) {
		ikePolicy = &models.IKEPolicyIdentity{
			ID: strfmt.UUID(d.Get(isVPNGatewayConnectionIKEPolicy).(string)),
		}
		hasChange = true
	} else {
		ikePolicy = nil
	}

	if d.HasChange(isVPNGatewayConnectionIPSECPolicy) {
		ipsecPolicy = &models.IpsecPolicyIdentity{
			ID: strfmt.UUID(d.Get(isVPNGatewayConnectionIPSECPolicy).(string)),
		}
		hasChange = true
	} else {
		ipsecPolicy = nil
	}

	if d.HasChange(isVPNGatewayConnectionAdminStateup) {
		adminStateUp = d.Get(isVPNGatewayConnectionAdminStateup).(bool)
		hasChange = true

	}

	if hasChange {
		_, err = VPNGatewayConnectionC.UpdateConnection(gConnID, gID, name, peerAddress, psk, adminStateUp, deadPeerDetection, ikePolicy, ipsecPolicy)
		if err != nil {
			return err
		}

	}

	return resourceIBMISVPNGatewayConnectionRead(d, meta)
}

func resourceIBMISVPNGatewayConnectionDelete(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	gID := parts[0]
	gConnID := parts[1]

	VPNGatewayConnectionC := vpn.NewVpnClient(sess)
	err = VPNGatewayConnectionC.DeleteConnection(gID, gConnID)
	if err != nil {
		return err
	}

	_, err = isWaitForVPNGatewayConnectionDeleted(VPNGatewayConnectionC, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func isWaitForVPNGatewayConnectionDeleted(VPNGatewayConnection *vpn.VpnClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for VPNGatewayConnection (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVPNGatewayConnectionDeleting},
		Target:     []string{},
		Refresh:    isVPNGatewayConnectionDeleteRefreshFunc(VPNGatewayConnection, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVPNGatewayConnectionDeleteRefreshFunc(VPNGatewayConnection *vpn.VpnClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		parts, err := idParts(id)
		if err != nil {
			return nil, isVPNGatewayConnectionDeleting, err
		}

		gID := parts[0]
		gConnID := parts[1]
		VPNGatewayConnection, err := VPNGatewayConnection.GetConnection(gID, gConnID)
		if err == nil {
			return VPNGatewayConnection, isVPNGatewayConnectionDeleting, nil
		}

		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			log.Printf("[DEBUG] %s", iserror.Error())
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "vpn_connection_not_found" {
				return nil, isVPNGatewayConnectionDeleted, nil
			}
		}
		return nil, isVPNGatewayConnectionDeleting, err
	}
}

func resourceIBMISVPNGatewayConnectionExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}

	gID := parts[0]
	gConnID := parts[1]
	VPNGatewayConnectionC := vpn.NewVpnClient(sess)

	_, err = VPNGatewayConnectionC.GetConnection(gID, gConnID)
	if err != nil {
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}
