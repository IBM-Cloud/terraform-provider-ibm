// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isVPNGatewayConnectionAdminStateup              = "admin_state_up"
	isVPNGatewayConnectionAdminAuthenticationmode   = "authentication_mode"
	isVPNGatewayConnectionName                      = "name"
	isVPNGatewayConnectionVPNGateway                = "vpn_gateway"
	isVPNGatewayConnection                          = "gateway_connection"
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
	isVPNGatewayConnectionMode                      = "mode"
	isVPNGatewayConnectionTunnels                   = "tunnels"
	isVPNGatewayConnectionResourcetype              = "resource_type"
	isVPNGatewayConnectionCreatedat                 = "created_at"
	isVPNGatewayConnectionStatusreasons             = "status_reasons"
)

func ResourceIBMISVPNGatewayConnection() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISVPNGatewayConnectionCreate,
		ReadContext:   resourceIBMISVPNGatewayConnectionRead,
		UpdateContext: resourceIBMISVPNGatewayConnectionUpdate,
		DeleteContext: resourceIBMISVPNGatewayConnectionDelete,
		Exists:        resourceIBMISVPNGatewayConnectionExists,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			isVPNGatewayConnectionName: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_vpn_gateway_connection", isVPNGatewayConnectionName),
				Description:  "VPN Gateway connection name",
			},

			isVPNGatewayConnectionVPNGateway: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "VPN Gateway info",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Href of the VPN Gateway connection",
			},
			// Deprecated
			isVPNGatewayConnectionPeerAddress: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "VPN gateway connection peer address",
				Deprecated:  "peer_address is deprecated, use peer instead",
			},

			// distribute traffic
			"distribute_traffic": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Indicates whether the traffic is distributed between the `up` tunnels of the VPN gateway connection when the VPC route's next hop is a VPN connection. If `false`, the traffic is only routed through the `up` tunnel with the lower `public_ip` address.",
			},

			// new breaking changes
			"establish_mode": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "bidirectional",
				ValidateFunc: validate.InvokeValidator("ibm_is_vpn_gateway_connection", "establish_mode"),
				Description:  "The establish mode of the VPN gateway connection:- `bidirectional`: Either side of the VPN gateway can initiate IKE protocol   negotiations or rekeying processes.- `peer_only`: Only the peer can initiate IKE protocol negotiations for this VPN gateway   connection. Additionally, the peer is responsible for initiating the rekeying process   after the connection is established. If rekeying does not occur, the VPN gateway   connection will be brought down after its lifetime expires.",
			},
			"local": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ike_identities": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "The local IKE identities.A VPN gateway in static route mode consists of two members in active-active mode. The first identity applies to the first member, and the second identity applies to the second member.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The IKE identity type.The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the backup policy on which the unexpected property value was encountered.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The IKE identity FQDN value.",
									},
								},
							},
						},
						"cidrs": {
							Type:          schema.TypeSet,
							Optional:      true,
							Computed:      true,
							ConflictsWith: []string{"local_cidrs"},
							Elem:          &schema.Schema{Type: schema.TypeString},
							Set:           schema.HashString,
							Description:   "VPN gateway connection local CIDRs",
						},
					},
				},
			},
			"peer": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ike_identity": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Description: "The peer IKE identity.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The IKE identity type.The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the backup policy on which the unexpected property value was encountered.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The IKE identity FQDN value.",
									},
								},
							},
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicates whether `peer.address` or `peer.fqdn` is used.",
						},
						"address": &schema.Schema{
							Type:          schema.TypeString,
							Optional:      true,
							Computed:      true,
							ConflictsWith: []string{"peer_address"},
							Description:   "The IP address of the peer VPN gateway for this connection.",
						},
						"fqdn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The FQDN of the peer VPN gateway for this connection.",
						},
						"cidrs": {
							Type:          schema.TypeSet,
							Optional:      true,
							Computed:      true,
							ConflictsWith: []string{"peer_cidrs"},
							Elem:          &schema.Schema{Type: schema.TypeString},
							Set:           schema.HashString,
							Description:   "VPN gateway connection peer CIDRs",
						},
						"asn": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "The peer autonomous system number (ASN) for this VPN gateway connection.",
						},
					},
				},
			},
			isVPNGatewayConnectionPreSharedKey: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "vpn gateway",
			},

			isVPNGatewayConnectionAdminStateup: {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "VPN gateway connection admin state",
			},
			// deprecated
			isVPNGatewayConnectionLocalCIDRS: {
				Type:          schema.TypeSet,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"local"},
				Elem:          &schema.Schema{Type: schema.TypeString},
				Set:           schema.HashString,
				Description:   "VPN gateway connection local CIDRs",
				Deprecated:    "local_cidrs is deprecated, use local instead",
			},
			// deprecated
			isVPNGatewayConnectionPeerCIDRS: {
				Type:          schema.TypeSet,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"peer"},
				Elem:          &schema.Schema{Type: schema.TypeString},
				Set:           schema.HashString,
				Description:   "VPN gateway connection peer CIDRs",
				Deprecated:    "peer_cidrs is deprecated, use peer instead",
			},

			isVPNGatewayConnectionDeadPeerDetectionAction: {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "restart",
				ValidateFunc: validate.InvokeValidator("ibm_is_vpn_gateway_connection", isVPNGatewayConnectionDeadPeerDetectionAction),
				Description:  "Action detection for dead peer detection action",
			},
			isVPNGatewayConnectionDeadPeerDetectionInterval: {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      2,
				ValidateFunc: validate.InvokeValidator("ibm_is_vpn_gateway_connection", isVPNGatewayConnectionDeadPeerDetectionInterval),
				Description:  "Interval for dead peer detection interval",
			},
			isVPNGatewayConnectionDeadPeerDetectionTimeout: {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      10,
				ValidateFunc: validate.InvokeValidator("ibm_is_vpn_gateway_connection", isVPNGatewayConnectionDeadPeerDetectionTimeout),
				Description:  "Timeout for dead peer detection",
			},

			isVPNGatewayConnectionIPSECPolicy: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP security policy for vpn gateway connection",
			},

			isVPNGatewayConnectionIKEPolicy: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "VPN gateway connection IKE Policy",
			},

			isVPNGatewayConnection: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this VPN gateway connection",
			},

			isVPNGatewayConnectionStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "VPN gateway connection status",
			},
			isVPNGatewayConnectionStatusreasons: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current status (if any).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the status reason.",
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the status reason.",
						},
						"more_info": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about this status reason.",
						},
					},
				},
			},
			flex.RelatedCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the VPN Gateway resource",
			},

			isVPNGatewayConnectionAdminAuthenticationmode: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The authentication mode",
			},

			isVPNGatewayConnectionResourcetype: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type",
			},
			"routing_protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Routing protocols for this VPN gateway connection.",
			},
			"tunnel": {
				Type:     schema.TypeList,
				MinItems: 2,
				MaxItems: 2,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"neighbor_ip": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The IP address of the neighbor on the virtual tunnel interface.",
						},
						"tunnel_interface_ip": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The IP address of the virtual tunnel interface.",
						},
					},
				},
			},

			isVPNGatewayConnectionCreatedat: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that this VPN gateway connection was created",
			},

			isVPNGatewayConnectionMode: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The mode of the VPN gateway",
			},

			isVPNGatewayConnectionTunnels: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The VPN tunnel configuration for this VPN gateway connection (in static route mode)",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address of the VPN gateway member in which the tunnel resides",
						},

						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the VPN Tunnel",
						},
						"neighbor_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address of the neighbor on the virtual tunnel interface.",
						},
						"tunnel_interface_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address of the virtual tunnel interface.",
						},
						"protocol_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "BGP routing protocol state.",
						},
					},
				},
			},
		},
	}
}

func ResourceIBMISVPNGatewayConnectionValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	action := "restart, clear, hold, none"

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isVPNGatewayConnectionName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
		validate.ValidateSchema{
			Identifier:                 "establish_mode",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "bidirectional, peer_only",
			Regexp:                     `^[a-z][a-z0-9]*(_[a-z0-9]+)*$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
	)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isVPNGatewayConnectionDeadPeerDetectionAction,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              action})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isVPNGatewayConnectionDeadPeerDetectionInterval,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "1",
			MaxValue:                   "86399"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isVPNGatewayConnectionDeadPeerDetectionTimeout,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "2",
			MaxValue:                   "86399"})

	ibmISVPNGatewayConnectionResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_vpn_gateway_connection", Schema: validateSchema}
	return &ibmISVPNGatewayConnectionResourceValidator
}

func resourceIBMISVPNGatewayConnectionCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	log.Printf("[DEBUG] VPNGatewayConnection create")
	name := d.Get(isVPNGatewayConnectionName).(string)
	gatewayID := d.Get(isVPNGatewayConnectionVPNGateway).(string)
	peerAddress := d.Get(isVPNGatewayConnectionPeerAddress).(string)
	prephasedKey := d.Get(isVPNGatewayConnectionPreSharedKey).(string)

	var interval, timeout int64
	if intvl, ok := d.GetOk(isVPNGatewayConnectionDeadPeerDetectionInterval); ok {
		interval = int64(intvl.(int))
	} else {
		interval = 30
	}

	if tout, ok := d.GetOk(isVPNGatewayConnectionDeadPeerDetectionTimeout); ok {
		timeout = int64(tout.(int))
	} else {
		timeout = 120
	}
	var action string
	if act, ok := d.GetOk(isVPNGatewayConnectionDeadPeerDetectionAction); ok {
		action = act.(string)
	} else {
		action = "none"
	}

	err := vpngwconCreate(context, d, meta, name, gatewayID, peerAddress, prephasedKey, action, interval, timeout)
	if err != nil {
		return err
	}
	return resourceIBMISVPNGatewayConnectionRead(context, d, meta)
}

func vpngwconCreate(context context.Context, d *schema.ResourceData, meta interface{}, name, gatewayID, peerAddress, prephasedKey, action string, interval, timeout int64) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	vpngateway, response, err := sess.GetVPNGatewayWithContext(context, &vpcv1.GetVPNGatewayOptions{
		ID: &gatewayID,
	})
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPNGatewayWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	var protocol string
	if routingProtocolOk, ok := d.GetOk("routing_protocol"); ok {
		protocol = routingProtocolOk.(string)
	}
	if *vpngateway.(*vpcv1.VPNGateway).Mode == "policy" {

		vpnGatewayConnectionPrototypeModel := &vpcv1.VPNGatewayConnectionPrototypeVPNGatewayConnectionPolicyModePrototype{
			Psk: &prephasedKey,
			DeadPeerDetection: &vpcv1.VPNGatewayConnectionDpdPrototype{
				Action:   &action,
				Interval: &interval,
				Timeout:  &timeout,
			},
			Name: &name,
		}

		if _, ok := d.GetOkExists(isVPNGatewayConnectionAdminStateup); ok {
			stateUp := d.Get(isVPNGatewayConnectionAdminStateup).(bool)
			vpnGatewayConnectionPrototypeModel.AdminStateUp = core.BoolPtr(stateUp)
		}

		var ikePolicyIdentity, ipsecPolicyIdentity string
		// new breaking changes
		if establishModeOk, ok := d.GetOk("establish_mode"); ok {
			vpnGatewayConnectionPrototypeModel.EstablishMode = core.StringPtr(establishModeOk.(string))
		}

		if localOk, ok := d.GetOk("local"); ok && len(localOk.([]interface{})) > 0 {
			log.Println("[INFO] inside local block")
			LocalModel, err := resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionPolicyModeLocalPrototype(localOk.([]interface{})[0].(map[string]interface{}))
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "create", "parse-local").GetDiag()
			}
			vpnGatewayConnectionPrototypeModel.Local = LocalModel
		} else if _, ok := d.GetOk(isVPNGatewayConnectionLocalCIDRS); ok {
			log.Println("[INFO] inside local cidrs block")
			localCidrs := flex.ExpandStringList((d.Get(isVPNGatewayConnectionLocalCIDRS).(*schema.Set)).List())
			model := &vpcv1.VPNGatewayConnectionPolicyModeLocalPrototype{}
			model.CIDRs = localCidrs
			vpnGatewayConnectionPrototypeModel.Local = model
		}
		if peerOk, ok := d.GetOk("peer"); ok && len(peerOk.([]interface{})) > 0 {
			PeerModel, err := resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionPolicyModePeerPrototype(peerOk.([]interface{})[0].(map[string]interface{}))
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "create", "parse-peer").GetDiag()
			}
			vpnGatewayConnectionPrototypeModel.Peer = PeerModel
		} else if _, ok := d.GetOk(isVPNGatewayConnectionPeerCIDRS); ok || peerAddress != "" {
			model := &vpcv1.VPNGatewayConnectionPolicyModePeerPrototype{}
			if ok {
				peerCidrs := flex.ExpandStringList((d.Get(isVPNGatewayConnectionPeerCIDRS).(*schema.Set)).List())
				model.CIDRs = peerCidrs
			}
			if peerAddress != "" {
				model.Address = &peerAddress
			}
			vpnGatewayConnectionPrototypeModel.Peer = model
		}

		if ikePolicy, ok := d.GetOk(isVPNGatewayConnectionIKEPolicy); ok {
			ikePolicyIdentity = ikePolicy.(string)
			vpnGatewayConnectionPrototypeModel.IkePolicy = &vpcv1.VPNGatewayConnectionIkePolicyPrototype{
				ID: &ikePolicyIdentity,
			}
		} else {
			vpnGatewayConnectionPrototypeModel.IkePolicy = nil
		}
		if ipsecPolicy, ok := d.GetOk(isVPNGatewayConnectionIPSECPolicy); ok {
			ipsecPolicyIdentity = ipsecPolicy.(string)
			vpnGatewayConnectionPrototypeModel.IpsecPolicy = &vpcv1.VPNGatewayConnectionIPsecPolicyPrototype{
				ID: &ipsecPolicyIdentity,
			}
		} else {
			vpnGatewayConnectionPrototypeModel.IpsecPolicy = nil
		}

		options := &vpcv1.CreateVPNGatewayConnectionOptions{
			VPNGatewayID:                  &gatewayID,
			VPNGatewayConnectionPrototype: vpnGatewayConnectionPrototypeModel,
		}

		vpnGatewayConnectionIntf, _, err := sess.CreateVPNGatewayConnectionWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateVPNGatewayConnectionWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnection); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnection)
			d.SetId(fmt.Sprintf("%s/%s", gatewayID, *vpnGatewayConnection.ID))
			log.Printf("[INFO] VPNGatewayConnection : %s/%s", gatewayID, *vpnGatewayConnection.ID)
		} else if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteMode); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteMode)
			d.SetId(fmt.Sprintf("%s/%s", gatewayID, *vpnGatewayConnection.ID))
			log.Printf("[INFO] VPNGatewayConnection : %s/%s", gatewayID, *vpnGatewayConnection.ID)
		} else if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionPolicyMode); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionPolicyMode)
			d.SetId(fmt.Sprintf("%s/%s", gatewayID, *vpnGatewayConnection.ID))
			log.Printf("[INFO] VPNGatewayConnection : %s/%s", gatewayID, *vpnGatewayConnection.ID)
		} else if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionStaticRouteMode); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionStaticRouteMode)
			d.SetId(fmt.Sprintf("%s/%s", gatewayID, *vpnGatewayConnection.ID))
			log.Printf("[INFO] VPNGatewayConnection : %s/%s", gatewayID, *vpnGatewayConnection.ID)
		} else if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionDynamicRouteMode); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionDynamicRouteMode)
			d.SetId(fmt.Sprintf("%s/%s", gatewayID, *vpnGatewayConnection.ID))
			log.Printf("[INFO] VPNGatewayConnection : %s/%s", gatewayID, *vpnGatewayConnection.ID)
		} else {
			return flex.DiscriminatedTerraformErrorf(nil, fmt.Sprintf("Unrecognized vpcv1.VPNGatewayConnectionIntf subtype encountered"), "ibm_is_vpn_gateway_connection", "create", "unrecognized-subtype-of-VPNGatewayConnection").GetDiag()
		}
	} else if *vpngateway.(*vpcv1.VPNGateway).Mode == "route" && protocol == "bgp" {

		vpnGatewayConnectionPrototypeModel := &vpcv1.VPNGatewayConnectionPrototypeVPNGatewayConnectionDynamicRouteModePrototype{
			Psk: &prephasedKey,
			DeadPeerDetection: &vpcv1.VPNGatewayConnectionDpdPrototype{
				Action:   &action,
				Interval: &interval,
				Timeout:  &timeout,
			},
			Name: &name,
		}
		if _, ok := d.GetOkExists(isVPNGatewayConnectionAdminStateup); ok {
			stateUp := d.Get(isVPNGatewayConnectionAdminStateup).(bool)
			vpnGatewayConnectionPrototypeModel.AdminStateUp = core.BoolPtr(stateUp)
		}
		var ikePolicyIdentity, ipsecPolicyIdentity string
		if establishModeOk, ok := d.GetOk("establish_mode"); ok {
			vpnGatewayConnectionPrototypeModel.EstablishMode = core.StringPtr(establishModeOk.(string))
		}

		if localOk, ok := d.GetOk("local"); ok && len(localOk.([]interface{})) > 0 {
			log.Println("[INFO] inside local block")
			LocalModel, err := resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionDynamicRouteModeLocalPrototype(localOk.([]interface{})[0].(map[string]interface{}))
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "create", "parse-local").GetDiag()
			}
			vpnGatewayConnectionPrototypeModel.Local = LocalModel
		}
		if peerOk, ok := d.GetOk("peer"); ok && len(peerOk.([]interface{})) > 0 {
			PeerModel, err := resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionDynamicRouteModePeerPrototype(peerOk.([]interface{})[0].(map[string]interface{}))
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "create", "parse-peer").GetDiag()
			}
			vpnGatewayConnectionPrototypeModel.Peer = PeerModel
		} else if peerAddress != "" {
			model := &vpcv1.VPNGatewayConnectionDynamicRouteModePeerPrototype{}
			if peerAddress != "" {
				model.Address = &peerAddress
			}
			vpnGatewayConnectionPrototypeModel.Peer = model
		}

		if ikePolicy, ok := d.GetOk(isVPNGatewayConnectionIKEPolicy); ok {
			ikePolicyIdentity = ikePolicy.(string)
			vpnGatewayConnectionPrototypeModel.IkePolicy = &vpcv1.VPNGatewayConnectionIkePolicyPrototype{
				ID: &ikePolicyIdentity,
			}
		} else {
			vpnGatewayConnectionPrototypeModel.IkePolicy = nil
		}

		if ipsecPolicy, ok := d.GetOk(isVPNGatewayConnectionIPSECPolicy); ok {
			ipsecPolicyIdentity = ipsecPolicy.(string)
			vpnGatewayConnectionPrototypeModel.IpsecPolicy = &vpcv1.VPNGatewayConnectionIPsecPolicyPrototype{
				ID: &ipsecPolicyIdentity,
			}
		} else {
			vpnGatewayConnectionPrototypeModel.IpsecPolicy = nil
		}
		if distributeTrafficOk, ok := d.GetOkExists("distribute_traffic"); ok {
			vpnGatewayConnectionPrototypeModel.DistributeTraffic = core.BoolPtr(distributeTrafficOk.(bool))
		}
		if routingProtocolOk, ok := d.GetOk("routing_protocol"); ok {
			vpnGatewayConnectionPrototypeModel.RoutingProtocol = core.StringPtr(routingProtocolOk.(string))
		}

		if tunnelOk, ok := d.GetOk("tunnel"); ok && len(tunnelOk.([]interface{})) > 0 {
			log.Println("[INFO] inside tunnel block")
			tunnelModels := []vpcv1.VPNGatewayConnectionTunnelPrototype{}
			for _, tunnelItem := range tunnelOk.([]interface{}) {
				tunnelModel := resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionTunnelPrototype(tunnelItem.(map[string]interface{}))
				tunnelModels = append(tunnelModels, tunnelModel)
			}
			vpnGatewayConnectionPrototypeModel.Tunnels = tunnelModels
		}

		options := &vpcv1.CreateVPNGatewayConnectionOptions{
			VPNGatewayID:                  &gatewayID,
			VPNGatewayConnectionPrototype: vpnGatewayConnectionPrototypeModel,
		}

		vpnGatewayConnectionIntf, _, err := sess.CreateVPNGatewayConnectionWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateVPNGatewayConnectionWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnection); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnection)
			d.SetId(fmt.Sprintf("%s/%s", gatewayID, *vpnGatewayConnection.ID))
			log.Printf("[INFO] VPNGatewayConnection : %s/%s", gatewayID, *vpnGatewayConnection.ID)
		} else if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteMode); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteMode)
			d.SetId(fmt.Sprintf("%s/%s", gatewayID, *vpnGatewayConnection.ID))
			log.Printf("[INFO] VPNGatewayConnection : %s/%s", gatewayID, *vpnGatewayConnection.ID)
		} else if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionPolicyMode); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionPolicyMode)
			d.SetId(fmt.Sprintf("%s/%s", gatewayID, *vpnGatewayConnection.ID))
			log.Printf("[INFO] VPNGatewayConnection : %s/%s", gatewayID, *vpnGatewayConnection.ID)
		} else if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionStaticRouteMode); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionStaticRouteMode)
			d.SetId(fmt.Sprintf("%s/%s", gatewayID, *vpnGatewayConnection.ID))
			log.Printf("[INFO] VPNGatewayConnection : %s/%s", gatewayID, *vpnGatewayConnection.ID)
		} else if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionDynamicRouteMode); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionDynamicRouteMode)
			d.SetId(fmt.Sprintf("%s/%s", gatewayID, *vpnGatewayConnection.ID))
			log.Printf("[INFO] VPNGatewayConnection : %s/%s", gatewayID, *vpnGatewayConnection.ID)
		} else {
			return flex.DiscriminatedTerraformErrorf(nil, fmt.Sprintf("Unrecognized vpcv1.VPNGatewayConnectionIntf subtype encountered"), "ibm_is_vpn_gateway_connection", "create", "unrecognized-subtype-of-VPNGatewayConnection").GetDiag()
		}
	} else {

		vpnGatewayConnectionPrototypeModel := &vpcv1.VPNGatewayConnectionPrototypeVPNGatewayConnectionStaticRouteModePrototype{
			Psk: &prephasedKey,
			DeadPeerDetection: &vpcv1.VPNGatewayConnectionDpdPrototype{
				Action:   &action,
				Interval: &interval,
				Timeout:  &timeout,
			},
			Name: &name,
		}
		if _, ok := d.GetOkExists(isVPNGatewayConnectionAdminStateup); ok {
			stateUp := d.Get(isVPNGatewayConnectionAdminStateup).(bool)
			vpnGatewayConnectionPrototypeModel.AdminStateUp = core.BoolPtr(stateUp)
		}
		var ikePolicyIdentity, ipsecPolicyIdentity string
		// new breaking changes
		if establishModeOk, ok := d.GetOk("establish_mode"); ok {
			vpnGatewayConnectionPrototypeModel.EstablishMode = core.StringPtr(establishModeOk.(string))
		}

		if localOk, ok := d.GetOk("local"); ok && len(localOk.([]interface{})) > 0 {
			log.Println("[INFO] inside local block")
			LocalModel, err := resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionStaticRouteModeLocalPrototype(localOk.([]interface{})[0].(map[string]interface{}))
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "create", "parse-local").GetDiag()
			}
			vpnGatewayConnectionPrototypeModel.Local = LocalModel
		}
		if peerOk, ok := d.GetOk("peer"); ok && len(peerOk.([]interface{})) > 0 {
			PeerModel, err := resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionStaticRouteModePeerPrototype(peerOk.([]interface{})[0].(map[string]interface{}))
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "create", "parse-peer").GetDiag()
			}
			vpnGatewayConnectionPrototypeModel.Peer = PeerModel
		} else if peerAddress != "" {
			model := &vpcv1.VPNGatewayConnectionStaticRouteModePeerPrototype{}
			if peerAddress != "" {
				model.Address = &peerAddress
			}
			vpnGatewayConnectionPrototypeModel.Peer = model
		}

		if ikePolicy, ok := d.GetOk(isVPNGatewayConnectionIKEPolicy); ok {
			ikePolicyIdentity = ikePolicy.(string)
			vpnGatewayConnectionPrototypeModel.IkePolicy = &vpcv1.VPNGatewayConnectionIkePolicyPrototype{
				ID: &ikePolicyIdentity,
			}
		} else {
			vpnGatewayConnectionPrototypeModel.IkePolicy = nil
		}

		if ipsecPolicy, ok := d.GetOk(isVPNGatewayConnectionIPSECPolicy); ok {
			ipsecPolicyIdentity = ipsecPolicy.(string)
			vpnGatewayConnectionPrototypeModel.IpsecPolicy = &vpcv1.VPNGatewayConnectionIPsecPolicyPrototype{
				ID: &ipsecPolicyIdentity,
			}
		} else {
			vpnGatewayConnectionPrototypeModel.IpsecPolicy = nil
		}
		if distributeTrafficOk, ok := d.GetOkExists("distribute_traffic"); ok {
			vpnGatewayConnectionPrototypeModel.DistributeTraffic = core.BoolPtr(distributeTrafficOk.(bool))
		}
		options := &vpcv1.CreateVPNGatewayConnectionOptions{
			VPNGatewayID:                  &gatewayID,
			VPNGatewayConnectionPrototype: vpnGatewayConnectionPrototypeModel,
		}

		vpnGatewayConnectionIntf, _, err := sess.CreateVPNGatewayConnectionWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateVPNGatewayConnectionWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnection); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnection)
			d.SetId(fmt.Sprintf("%s/%s", gatewayID, *vpnGatewayConnection.ID))
			log.Printf("[INFO] VPNGatewayConnection : %s/%s", gatewayID, *vpnGatewayConnection.ID)
		} else if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteMode); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteMode)
			d.SetId(fmt.Sprintf("%s/%s", gatewayID, *vpnGatewayConnection.ID))
			log.Printf("[INFO] VPNGatewayConnection : %s/%s", gatewayID, *vpnGatewayConnection.ID)
		} else if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionPolicyMode); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionPolicyMode)
			d.SetId(fmt.Sprintf("%s/%s", gatewayID, *vpnGatewayConnection.ID))
			log.Printf("[INFO] VPNGatewayConnection : %s/%s", gatewayID, *vpnGatewayConnection.ID)
		} else if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionStaticRouteMode); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionStaticRouteMode)
			d.SetId(fmt.Sprintf("%s/%s", gatewayID, *vpnGatewayConnection.ID))
			log.Printf("[INFO] VPNGatewayConnection : %s/%s", gatewayID, *vpnGatewayConnection.ID)
		} else if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionDynamicRouteMode); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionDynamicRouteMode)
			d.SetId(fmt.Sprintf("%s/%s", gatewayID, *vpnGatewayConnection.ID))
			log.Printf("[INFO] VPNGatewayConnection : %s/%s", gatewayID, *vpnGatewayConnection.ID)
		} else {
			return flex.DiscriminatedTerraformErrorf(nil, fmt.Sprintf("Unrecognized vpcv1.VPNGatewayConnectionIntf subtype encountered"), "ibm_is_vpn_gateway_connection", "create", "unrecognized-subtype-of-VPNGatewayConnection").GetDiag()
		}
	}

	return nil
}

func resourceIBMISVPNGatewayConnectionRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "sep-id-parts").GetDiag()
	}

	gID := parts[0]
	gConnID := parts[1]

	diagErr := vpngwconGet(context, d, meta, gID, gConnID)
	if diagErr != nil {
		return diagErr
	}
	return nil
}

func vpngwconGet(context context.Context, d *schema.ResourceData, meta interface{}, gID, gConnID string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.GetVPNGatewayConnectionOptions{
		VPNGatewayID: &gID,
		ID:           &gConnID,
	}
	vpnGatewayConnectionIntf, response, err := sess.GetVPNGatewayConnectionWithContext(context, options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPNGatewayConnectionWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(isVPNGatewayConnection, gConnID); err != nil {
		err = fmt.Errorf("Error setting gateway_connection: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-gateway_connection").GetDiag()
	}
	setvpnGatewayConnectionIntfResource(context, d, gID, vpnGatewayConnectionIntf)
	getVPNGatewayOptions := &vpcv1.GetVPNGatewayOptions{
		ID: &gID,
	}
	vpngatewayIntf, response, err := sess.GetVPNGatewayWithContext(context, getVPNGatewayOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPNGatewayWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	vpngateway := vpngatewayIntf.(*vpcv1.VPNGateway)
	if err = d.Set(flex.RelatedCRN, *vpngateway.CRN); err != nil {
		err = fmt.Errorf("Error setting related_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-related_crn").GetDiag()
	}

	return nil
}

func resourceIBMISVPNGatewayConnectionUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	hasChanged := false

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "update", "sep-id-parts").GetDiag()
	}

	gID := parts[0]
	gConnID := parts[1]
	diagErr := vpngwconUpdate(context, d, meta, gID, gConnID, hasChanged)
	if diagErr != nil {
		return diagErr
	}
	return resourceIBMISVPNGatewayConnectionRead(context, d, meta)
}

func vpngwconUpdate(context context.Context, d *schema.ResourceData, meta interface{}, gID, gConnID string, hasChanged bool) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateVpnGatewayConnectionOptions := &vpcv1.UpdateVPNGatewayConnectionOptions{
		VPNGatewayID: &gID,
		ID:           &gConnID,
	}
	vpnGatewayConnectionPatchModel := &vpcv1.VPNGatewayConnectionPatch{}

	if d.HasChange("distribute_traffic") {
		vpnGatewayConnectionPatchModel.DistributeTraffic = core.BoolPtr(d.Get("distribute_traffic").(bool))
		hasChanged = true
	}
	if d.HasChange("routing_protocol") {
		vpnGatewayConnectionPatchModel.RoutingProtocol = core.StringPtr(d.Get("routing_protocol").(string))
		hasChanged = true
	}
	if d.HasChange("tunnel_neighbor_address") {
		vpnGatewayConnectionPatchModel.Tunnels[0].NeighborIP.Address = core.StringPtr(d.Get("tunnel_neighbor_address").(string))
		hasChanged = true
	}
	if d.HasChange("tunnel_interface_address") {
		vpnGatewayConnectionPatchModel.Tunnels[0].TunnelInterfaceIP.Address = core.StringPtr(d.Get("tunnel_interface_address").(string))
		hasChanged = true
	}

	if d.HasChange("tunnel") {
		log.Println("[INFO] inside tunnel block")
		options := &vpcv1.GetVPNGatewayConnectionOptions{
			VPNGatewayID: &gID,
			ID:           &gConnID,
		}
		_, response, err := sess.GetVPNGatewayConnectionWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPNGatewayConnectionWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "get")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		eTag := response.Headers.Get("ETag")
		updateVpnGatewayConnectionOptions.IfMatch = &eTag

		tunnelModels := []vpcv1.VPNGatewayConnectionTunnel{}
		tunnelOk, _ := d.GetOk("tunnel")
		for _, tunnelItem := range tunnelOk.([]interface{}) {
			tunnelModel := resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionTunnelPrototypePatch(tunnelItem.(map[string]interface{}))
			tunnelModels = append(tunnelModels, tunnelModel)
		}
		vpnGatewayConnectionPatchModel.Tunnels = tunnelModels
	}

	if d.HasChange(isVPNGatewayConnectionName) {
		name := d.Get(isVPNGatewayConnectionName).(string)
		vpnGatewayConnectionPatchModel.Name = &name
		hasChanged = true
	}
	if d.HasChange("establish_mode") {
		newEstablishMode := d.Get("establish_mode").(string)
		vpnGatewayConnectionPatchModel.EstablishMode = &newEstablishMode
		hasChanged = true
	}

	if d.HasChange("local.0.cidrs") {
		o, n := d.GetChange("local.0.cidrs")
		oldSet := o.(*schema.Set)
		newSet := n.(*schema.Set)

		// Find items to remove (present in old but not in new)
		toRemove := oldSet.Difference(newSet)
		if toRemove.Len() > 0 {
			for _, cidr := range toRemove.List() {
				cidrStr := cidr.(string)
				removeVPNGatewayConnectionsLocalCIDROptions := &vpcv1.RemoveVPNGatewayConnectionsLocalCIDROptions{
					VPNGatewayID: &gID,
					ID:           &gConnID,
					CIDR:         &cidrStr,
				}

				res, err := sess.RemoveVPNGatewayConnectionsLocalCIDRWithContext(context, removeVPNGatewayConnectionsLocalCIDROptions)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("RemoveVPNGatewayConnectionsLocalCIDRWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}

				if res.StatusCode != 201 && res.StatusCode != 204 {
					err = fmt.Errorf("unexpected status code %d while removing Local CIDR %s", res.StatusCode, cidrStr)
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("RemoveVPNGatewayConnectionsLocalCIDRWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		}

		// Find items to add (present in new but not in old)
		toAdd := newSet.Difference(oldSet)
		if toAdd.Len() > 0 {
			for _, cidr := range toAdd.List() {
				cidrStr := cidr.(string)
				addVPNGatewayConnectionsLocalCIDROptions := &vpcv1.AddVPNGatewayConnectionsLocalCIDROptions{
					VPNGatewayID: &gID,
					ID:           &gConnID,
					CIDR:         &cidrStr,
				}

				res, err := sess.AddVPNGatewayConnectionsLocalCIDRWithContext(context, addVPNGatewayConnectionsLocalCIDROptions)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("AddVPNGatewayConnectionsLocalCIDRWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}

				if res.StatusCode != 201 && res.StatusCode != 204 {
					err = fmt.Errorf("unexpected status code %d while adding Local CIDR %s", res.StatusCode, cidrStr)
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("AddVPNGatewayConnectionsLocalCIDRWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		}
	}

	if d.HasChange("peer") {
		peer, err := resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionPeerPatch(d, d.Get("peer.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "update", "parse-peer").GetDiag()
		}
		if d.HasChange("peer.0.cidrs") {
			o, n := d.GetChange("peer.0.cidrs")
			oldSet := o.(*schema.Set)
			newSet := n.(*schema.Set)

			// Find items to remove (present in old but not in new)
			toRemove := oldSet.Difference(newSet)
			if toRemove.Len() > 0 {
				for _, cidr := range toRemove.List() {
					cidrStr := cidr.(string)
					removeVPNGatewayConnectionsPeerCIDROptions := &vpcv1.RemoveVPNGatewayConnectionsPeerCIDROptions{
						VPNGatewayID: &gID,
						ID:           &gConnID,
						CIDR:         &cidrStr,
					}

					res, err := sess.RemoveVPNGatewayConnectionsPeerCIDRWithContext(context, removeVPNGatewayConnectionsPeerCIDROptions)
					if err != nil {
						tfErr := flex.TerraformErrorf(err, fmt.Sprintf("RemoveVPNGatewayConnectionsPeerCIDRWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "update")
						log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
						return tfErr.GetDiag()
					}

					if res.StatusCode != 201 && res.StatusCode != 204 {
						err = fmt.Errorf("unexpected status code %d while removing CIDR %s", res.StatusCode, cidrStr)
						tfErr := flex.TerraformErrorf(err, fmt.Sprintf("RemoveVPNGatewayConnectionsPeerCIDRWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "update")
						log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
						return tfErr.GetDiag()
					}
				}
			}

			// Find items to add (present in new but not in old)
			toAdd := newSet.Difference(oldSet)
			if toAdd.Len() > 0 {
				for _, cidr := range toAdd.List() {
					cidrStr := cidr.(string)
					addVPNGatewayConnectionsPeerCIDROptions := &vpcv1.AddVPNGatewayConnectionsPeerCIDROptions{
						VPNGatewayID: &gID,
						ID:           &gConnID,
						CIDR:         &cidrStr,
					}

					res, err := sess.AddVPNGatewayConnectionsPeerCIDRWithContext(context, addVPNGatewayConnectionsPeerCIDROptions)
					if err != nil {
						tfErr := flex.TerraformErrorf(err, fmt.Sprintf("AddVPNGatewayConnectionsPeerCIDRWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "update")
						log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
						return tfErr.GetDiag()
					}

					if res.StatusCode != 201 && res.StatusCode != 204 {
						err = fmt.Errorf("unexpected status code %d while adding CIDR %s", res.StatusCode, cidrStr)
						tfErr := flex.TerraformErrorf(err, fmt.Sprintf("AddVPNGatewayConnectionsPeerCIDRWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "update")
						log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
						return tfErr.GetDiag()
					}
				}
			}

		}
		vpnGatewayConnectionPatchModel.Peer = peer
		hasChanged = true
	}
	// Deprecated
	if d.HasChange(isVPNGatewayConnectionPeerAddress) {
		peerAddress := d.Get(isVPNGatewayConnectionPeerAddress).(string)
		model := &vpcv1.VPNGatewayConnectionPeerPatch{}
		model.Address = &peerAddress
		hasChanged = true
	}

	if d.HasChange(isVPNGatewayConnectionPreSharedKey) {
		psk := d.Get(isVPNGatewayConnectionPreSharedKey).(string)
		vpnGatewayConnectionPatchModel.Psk = &psk
		hasChanged = true
	}

	if d.HasChange(isVPNGatewayConnectionDeadPeerDetectionAction) || d.HasChange(isVPNGatewayConnectionDeadPeerDetectionInterval) || d.HasChange(isVPNGatewayConnectionDeadPeerDetectionTimeout) {
		action := d.Get(isVPNGatewayConnectionDeadPeerDetectionAction).(string)
		interval := int64(d.Get(isVPNGatewayConnectionDeadPeerDetectionInterval).(int))
		timeout := int64(d.Get(isVPNGatewayConnectionDeadPeerDetectionTimeout).(int))

		// Construct an instance of the VPNGatewayConnectionDpdPatch model
		vpnGatewayConnectionDpdPatchModel := new(vpcv1.VPNGatewayConnectionDpdPatch)
		vpnGatewayConnectionDpdPatchModel.Action = &action
		vpnGatewayConnectionDpdPatchModel.Interval = &interval
		vpnGatewayConnectionDpdPatchModel.Timeout = &timeout
		vpnGatewayConnectionPatchModel.DeadPeerDetection = vpnGatewayConnectionDpdPatchModel
		hasChanged = true
	}

	if d.HasChange(isVPNGatewayConnectionIKEPolicy) {
		ikePolicyIdentity := d.Get(isVPNGatewayConnectionIKEPolicy).(string)
		if ikePolicyIdentity == "" {
			var nullPatch *vpcv1.VPNGatewayConnectionIkePolicyPatch
			vpnGatewayConnectionPatchModel.IkePolicy = nullPatch
		} else {
			vpnGatewayConnectionPatchModel.IkePolicy = &vpcv1.VPNGatewayConnectionIkePolicyPatch{
				ID: &ikePolicyIdentity,
			}
		}
		hasChanged = true
	} else {
		vpnGatewayConnectionPatchModel.IkePolicy = nil
	}

	if d.HasChange(isVPNGatewayConnectionIPSECPolicy) {
		ipsecPolicyIdentity := d.Get(isVPNGatewayConnectionIPSECPolicy).(string)
		if ipsecPolicyIdentity == "" {
			var nullPatch *vpcv1.VPNGatewayConnectionIPsecPolicyPatch
			vpnGatewayConnectionPatchModel.IpsecPolicy = nullPatch
		} else {
			vpnGatewayConnectionPatchModel.IpsecPolicy = &vpcv1.VPNGatewayConnectionIPsecPolicyPatch{
				ID: &ipsecPolicyIdentity,
			}
		}
		hasChanged = true
	} else {
		vpnGatewayConnectionPatchModel.IpsecPolicy = nil
	}

	if d.HasChange(isVPNGatewayConnectionAdminStateup) {
		adminStateUp := d.Get(isVPNGatewayConnectionAdminStateup).(bool)
		vpnGatewayConnectionPatchModel.AdminStateUp = &adminStateUp
		hasChanged = true
	}

	if hasChanged {
		vpnGatewayConnectionPatch, err := vpnGatewayConnectionPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("vpnGatewayConnectionPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateVpnGatewayConnectionOptions.VPNGatewayConnectionPatch = vpnGatewayConnectionPatch
		_, _, err = sess.UpdateVPNGatewayConnectionWithContext(context, updateVpnGatewayConnectionOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVPNGatewayConnectionWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return nil
}

func resourceIBMISVPNGatewayConnectionDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "delete", "sep-id-parts").GetDiag()
	}

	gID := parts[0]
	gConnID := parts[1]

	diagErr := vpngwconDelete(context, d, meta, gID, gConnID)
	if diagErr != nil {
		return diagErr
	}
	return nil
}

func vpngwconDelete(context context.Context, d *schema.ResourceData, meta interface{}, gID, gConnID string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getVpnGatewayConnectionOptions := &vpcv1.GetVPNGatewayConnectionOptions{
		VPNGatewayID: &gID,
		ID:           &gConnID,
	}
	_, response, err := sess.GetVPNGatewayConnectionWithContext(context, getVpnGatewayConnectionOptions)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPNGatewayConnectionWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteVpnGatewayConnectionOptions := &vpcv1.DeleteVPNGatewayConnectionOptions{
		VPNGatewayID: &gID,
		ID:           &gConnID,
	}
	response, err = sess.DeleteVPNGatewayConnectionWithContext(context, deleteVpnGatewayConnectionOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteVPNGatewayConnectionWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForVPNGatewayConnectionDeleted(sess, gID, gConnID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVPNGatewayConnectionDeleted failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")
	return nil
}

func isWaitForVPNGatewayConnectionDeleted(vpnGatewayConnection *vpcv1.VpcV1, gID, gConnID string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for VPNGatewayConnection (%s) to be deleted.", gConnID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVPNGatewayConnectionDeleting},
		Target:     []string{"", isVPNGatewayConnectionDeleted},
		Refresh:    isVPNGatewayConnectionDeleteRefreshFunc(vpnGatewayConnection, gID, gConnID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVPNGatewayConnectionDeleteRefreshFunc(vpnGatewayConnection *vpcv1.VpcV1, gID, gConnID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getVpnGatewayConnectionOptions := &vpcv1.GetVPNGatewayConnectionOptions{
			VPNGatewayID: &gID,
			ID:           &gConnID,
		}
		vpngwcon, response, err := vpnGatewayConnection.GetVPNGatewayConnection(getVpnGatewayConnectionOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return "", isVPNGatewayConnectionDeleted, nil
			}
			return "", "", fmt.Errorf("[ERROR] The Vpn Gateway Connection %s failed to delete: %s\n%s", gConnID, err, response)
		}
		return vpngwcon, isVPNGatewayConnectionDeleting, nil
	}
}

func resourceIBMISVPNGatewayConnectionExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return false, err
	}
	if len(parts) != 2 {
		return false, fmt.Errorf("[ERROR] Incorrect ID %s: ID should be a combination of gID/gConnID", d.Id())
	}

	gID := parts[0]
	gConnID := parts[1]
	exists, err := vpngwconExists(d, meta, gID, gConnID)
	return exists, err
}

func vpngwconExists(d *schema.ResourceData, meta interface{}, gID, gConnID string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}

	getVpnGatewayConnectionOptions := &vpcv1.GetVPNGatewayConnectionOptions{
		VPNGatewayID: &gID,
		ID:           &gConnID,
	}
	_, response, err := sess.GetVPNGatewayConnection(getVpnGatewayConnectionOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("[ERROR] Error getting Vpn Gateway Connection: %s\n%s", err, response)
	}
	return true, nil
}

func resourceVPNGatewayConnectionFlattenLifecycleReasons(statusReasons []vpcv1.VPNGatewayConnectionStatusReason) (statusReasonsList []map[string]interface{}) {
	statusReasonsList = make([]map[string]interface{}, 0)
	for _, lr := range statusReasons {
		currentLR := map[string]interface{}{}
		if lr.Code != nil && lr.Message != nil {
			currentLR[isInstanceLifecycleReasonsCode] = *lr.Code
			currentLR[isInstanceLifecycleReasonsMessage] = *lr.Message
			if lr.MoreInfo != nil {
				currentLR[isInstanceLifecycleReasonsMoreInfo] = *lr.MoreInfo
			}
			statusReasonsList = append(statusReasonsList, currentLR)
		}
	}
	return statusReasonsList
}

// helper functions

func resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionPolicyModeLocalPrototype(modelMap map[string]interface{}) (*vpcv1.VPNGatewayConnectionPolicyModeLocalPrototype, error) {
	model := &vpcv1.VPNGatewayConnectionPolicyModeLocalPrototype{}
	if modelMap["ike_identities"] != nil {
		ikeIdentities := []vpcv1.VPNGatewayConnectionIkeIdentityPrototypeIntf{}
		for _, ikeIdentitiesItem := range modelMap["ike_identities"].([]interface{}) {
			ikeIdentitiesItemModel, err := resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionIkeIdentityPrototype(ikeIdentitiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			ikeIdentities = append(ikeIdentities, ikeIdentitiesItemModel)
		}
		model.IkeIdentities = ikeIdentities
	}
	if modelMap["cidrs"] != nil && modelMap["cidrs"].(*schema.Set).Len() > 0 {
		localCidrs := flex.ExpandStringList((modelMap["cidrs"].(*schema.Set)).List())
		model.CIDRs = localCidrs
	}
	return model, nil
}
func resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionStaticRouteModeLocalPrototype(modelMap map[string]interface{}) (*vpcv1.VPNGatewayConnectionStaticRouteModeLocalPrototype, error) {
	model := &vpcv1.VPNGatewayConnectionStaticRouteModeLocalPrototype{}
	if modelMap["ike_identities"] != nil {
		ikeIdentities := []vpcv1.VPNGatewayConnectionIkeIdentityPrototypeIntf{}
		for _, ikeIdentitiesItem := range modelMap["ike_identities"].([]interface{}) {
			ikeIdentitiesItemModel, err := resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionIkeIdentityPrototype(ikeIdentitiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			ikeIdentities = append(ikeIdentities, ikeIdentitiesItemModel)
		}
		model.IkeIdentities = ikeIdentities
	}
	return model, nil
}

func resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionDynamicRouteModeLocalPrototype(modelMap map[string]interface{}) (*vpcv1.VPNGatewayConnectionDynamicRouteModeLocalPrototype, error) {
	model := &vpcv1.VPNGatewayConnectionDynamicRouteModeLocalPrototype{}
	if modelMap["ike_identities"] != nil {
		ikeIdentities := []vpcv1.VPNGatewayConnectionIkeIdentityPrototypeIntf{}
		for _, ikeIdentitiesItem := range modelMap["ike_identities"].([]interface{}) {
			ikeIdentitiesItemModel, err := resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionIkeIdentityPrototype(ikeIdentitiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			ikeIdentities = append(ikeIdentities, ikeIdentitiesItemModel)
		}
		model.IkeIdentities = ikeIdentities
	}
	return model, nil
}

func resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionTunnelPrototype(modelMap map[string]interface{}) vpcv1.VPNGatewayConnectionTunnelPrototype {
	model := vpcv1.VPNGatewayConnectionTunnelPrototype{}
	if model.NeighborIP == nil {
		model.NeighborIP = &vpcv1.IP{}
	}
	if model.TunnelInterfaceIP == nil {
		model.TunnelInterfaceIP = &vpcv1.IP{}
	}
	if neighborIP, ok := modelMap["neighbor_ip"].(string); ok {
		model.NeighborIP.Address = core.StringPtr(neighborIP)
	}

	if tunnelInterfaceIP, ok := modelMap["tunnel_interface_ip"].(string); ok {
		model.TunnelInterfaceIP.Address = core.StringPtr(tunnelInterfaceIP)
	}
	return model
}

func resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionTunnelPrototypePatch(modelMap map[string]interface{}) vpcv1.VPNGatewayConnectionTunnel {
	model := vpcv1.VPNGatewayConnectionTunnel{}
	if model.NeighborIP == nil {
		model.NeighborIP = &vpcv1.IP{}
	}
	if model.TunnelInterfaceIP == nil {
		model.TunnelInterfaceIP = &vpcv1.IP{}
	}
	if neighborIP, ok := modelMap["neighbor_ip"].(string); ok {
		model.NeighborIP.Address = core.StringPtr(neighborIP)
	}

	if tunnelInterfaceIP, ok := modelMap["tunnel_interface_ip"].(string); ok {
		model.TunnelInterfaceIP.Address = core.StringPtr(tunnelInterfaceIP)
	}
	return model
}

func resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionIkeIdentityPrototype(modelMap map[string]interface{}) (vpcv1.VPNGatewayConnectionIkeIdentityPrototypeIntf, error) {
	model := &vpcv1.VPNGatewayConnectionIkeIdentityPrototype{}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	return model, nil
}

func resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionPolicyModePeerPrototype(modelMap map[string]interface{}) (vpcv1.VPNGatewayConnectionPolicyModePeerPrototypeIntf, error) {
	model := &vpcv1.VPNGatewayConnectionPolicyModePeerPrototype{}
	if modelMap["ike_identity"] != nil && len(modelMap["ike_identity"].([]interface{})) > 0 {
		IkeIdentityModel, err := resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionIkeIdentityPrototype(modelMap["ike_identity"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.IkeIdentity = IkeIdentityModel
	}
	if modelMap["address"] != nil && modelMap["address"].(string) != "" {
		model.Address = core.StringPtr(modelMap["address"].(string))
	}
	if modelMap["fqdn"] != nil && modelMap["fqdn"].(string) != "" {
		model.Fqdn = core.StringPtr(modelMap["fqdn"].(string))
	}
	if modelMap["cidrs"] != nil && modelMap["cidrs"].(*schema.Set).Len() > 0 {
		peerCidrs := flex.ExpandStringList((modelMap["cidrs"].(*schema.Set)).List())
		model.CIDRs = peerCidrs
	}
	return model, nil
}
func resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionStaticRouteModePeerPrototype(modelMap map[string]interface{}) (vpcv1.VPNGatewayConnectionStaticRouteModePeerPrototypeIntf, error) {
	model := &vpcv1.VPNGatewayConnectionStaticRouteModePeerPrototype{}
	if modelMap["ike_identity"] != nil && len(modelMap["ike_identity"].([]interface{})) > 0 {
		IkeIdentityModel, err := resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionIkeIdentityPrototype(modelMap["ike_identity"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.IkeIdentity = IkeIdentityModel
	}
	if modelMap["address"] != nil && modelMap["address"].(string) != "" {
		model.Address = core.StringPtr(modelMap["address"].(string))
	}
	if modelMap["fqdn"] != nil && modelMap["fqdn"].(string) != "" {
		model.Fqdn = core.StringPtr(modelMap["fqdn"].(string))
	}
	return model, nil
}

func resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionDynamicRouteModePeerPrototype(modelMap map[string]interface{}) (vpcv1.VPNGatewayConnectionDynamicRouteModePeerPrototypeIntf, error) {
	model := &vpcv1.VPNGatewayConnectionDynamicRouteModePeerPrototype{}
	if modelMap["asn"] != nil && modelMap["asn"].(int) > 0 {
		model.Asn = core.Int64Ptr(int64(modelMap["asn"].(int)))
	}
	if modelMap["ike_identity"] != nil && len(modelMap["ike_identity"].([]interface{})) > 0 {
		IkeIdentityModel, err := resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionIkeIdentityPrototype(modelMap["ike_identity"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.IkeIdentity = IkeIdentityModel
	}
	if modelMap["address"] != nil && modelMap["address"].(string) != "" {
		model.Address = core.StringPtr(modelMap["address"].(string))
	}
	if modelMap["fqdn"] != nil && modelMap["fqdn"].(string) != "" {
		model.Fqdn = core.StringPtr(modelMap["fqdn"].(string))
	}
	return model, nil
}

func resourceIBMIsVPNGatewayConnectionMapToVPNGatewayConnectionPeerPatch(d *schema.ResourceData, modelMap map[string]interface{}) (vpcv1.VPNGatewayConnectionPeerPatchIntf, error) {
	model := &vpcv1.VPNGatewayConnectionPeerPatch{}
	if d.HasChange("peer.0.address") && modelMap["address"] != nil && modelMap["address"].(string) != "" {
		model.Address = core.StringPtr(modelMap["address"].(string))
	}
	if d.HasChange("peer.0.fqdn") && modelMap["fqdn"] != nil && modelMap["fqdn"].(string) != "" {
		model.Fqdn = core.StringPtr(modelMap["fqdn"].(string))
	}
	if d.HasChange("peer.0.asn") && modelMap["asn"] != nil && modelMap["asn"].(int) > 0 {
		asn := core.Int64Ptr(int64(modelMap["asn"].(int)))
		model.Asn = asn
	}
	return model, nil
}
func resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionStaticRouteModeLocalToMap(model *vpcv1.VPNGatewayConnectionStaticRouteModeLocal) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	ikeIdentities := []map[string]interface{}{}
	for _, ikeIdentitiesItem := range model.IkeIdentities {
		ikeIdentitiesItemMap, err := resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionIkeIdentityToMap(ikeIdentitiesItem)
		if err != nil {
			return modelMap, err
		}
		ikeIdentities = append(ikeIdentities, ikeIdentitiesItemMap)
	}
	modelMap["ike_identities"] = ikeIdentities
	return modelMap, nil
}

func resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionIkeIdentityToMap(model vpcv1.VPNGatewayConnectionIkeIdentityIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.VPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityFqdn); ok {
		return resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityFqdnToMap(model.(*vpcv1.VPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityFqdn))
	} else if _, ok := model.(*vpcv1.VPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityHostname); ok {
		return resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityHostnameToMap(model.(*vpcv1.VPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityHostname))
	} else if _, ok := model.(*vpcv1.VPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityIPv4); ok {
		return resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityIPv4ToMap(model.(*vpcv1.VPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityIPv4))
	} else if _, ok := model.(*vpcv1.VPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityKeyID); ok {
		return resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityKeyIDToMap(model.(*vpcv1.VPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityKeyID))
	} else if _, ok := model.(*vpcv1.VPNGatewayConnectionIkeIdentity); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.VPNGatewayConnectionIkeIdentity)
		modelMap["type"] = model.Type
		if model.Value != nil {
			modelMap["value"] = model.Value
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.VPNGatewayConnectionIkeIdentityIntf subtype encountered")
	}
}

func resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityFqdnToMap(model *vpcv1.VPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityFqdn) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	modelMap["value"] = model.Value
	return modelMap, nil
}

func resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityHostnameToMap(model *vpcv1.VPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityHostname) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	modelMap["value"] = model.Value
	return modelMap, nil
}

func resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityIPv4ToMap(model *vpcv1.VPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityIPv4) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	modelMap["value"] = model.Value
	return modelMap, nil
}

func resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityKeyIDToMap(model *vpcv1.VPNGatewayConnectionIkeIdentityVPNGatewayConnectionIkeIdentityKeyID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	modelMap["value"] = string(*model.Value)
	return modelMap, nil
}

func resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionStaticRouteModePeerToMap(model vpcv1.VPNGatewayConnectionStaticRouteModePeerIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.VPNGatewayConnectionStaticRouteModePeerVPNGatewayConnectionPeerByAddress); ok {
		return resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionStaticRouteModePeerVPNGatewayConnectionPeerByAddressToMap(model.(*vpcv1.VPNGatewayConnectionStaticRouteModePeerVPNGatewayConnectionPeerByAddress))
	} else if _, ok := model.(*vpcv1.VPNGatewayConnectionStaticRouteModePeerVPNGatewayConnectionPeerByFqdn); ok {
		return resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionStaticRouteModePeerVPNGatewayConnectionPeerByFqdnToMap(model.(*vpcv1.VPNGatewayConnectionStaticRouteModePeerVPNGatewayConnectionPeerByFqdn))
	} else if _, ok := model.(*vpcv1.VPNGatewayConnectionStaticRouteModePeer); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.VPNGatewayConnectionStaticRouteModePeer)
		ikeIdentityMap, err := resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionIkeIdentityToMap(model.IkeIdentity)
		if err != nil {
			return modelMap, err
		}
		modelMap["ike_identity"] = []map[string]interface{}{ikeIdentityMap}
		modelMap["type"] = model.Type
		if model.Address != nil {
			modelMap["address"] = model.Address
		}
		if model.Fqdn != nil {
			modelMap["fqdn"] = model.Fqdn
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.VPNGatewayConnectionStaticRouteModePeerIntf subtype encountered")
	}
}

func resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionStaticRouteModePeerVPNGatewayConnectionPeerByAddressToMap(model *vpcv1.VPNGatewayConnectionStaticRouteModePeerVPNGatewayConnectionPeerByAddress) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	ikeIdentityMap, err := resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionIkeIdentityToMap(model.IkeIdentity)
	if err != nil {
		return modelMap, err
	}
	modelMap["ike_identity"] = []map[string]interface{}{ikeIdentityMap}
	modelMap["type"] = model.Type
	modelMap["address"] = model.Address
	return modelMap, nil
}

func resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionStaticRouteModePeerVPNGatewayConnectionPeerByFqdnToMap(model *vpcv1.VPNGatewayConnectionStaticRouteModePeerVPNGatewayConnectionPeerByFqdn) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	ikeIdentityMap, err := resourceIBMIsVPNGatewayConnectionVPNGatewayConnectionIkeIdentityToMap(model.IkeIdentity)
	if err != nil {
		return modelMap, err
	}
	modelMap["ike_identity"] = []map[string]interface{}{ikeIdentityMap}
	modelMap["type"] = model.Type
	modelMap["fqdn"] = model.Fqdn
	return modelMap, nil
}

func setvpnGatewayConnectionIntfResource(context context.Context, d *schema.ResourceData, vpn_gateway_id string, vpnGatewayConnectionIntf vpcv1.VPNGatewayConnectionIntf) diag.Diagnostics {
	var err error

	switch reflect.TypeOf(vpnGatewayConnectionIntf).String() {
	case "*vpcv1.VPNGatewayConnection":
		{
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnection)
			d.SetId(fmt.Sprintf("%s/%s", vpn_gateway_id, *vpnGatewayConnection.ID))
			if err = d.Set("admin_state_up", vpnGatewayConnection.AdminStateUp); err != nil {
				err = fmt.Errorf("Error setting admin_state_up: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-admin_state_up").GetDiag()
			}
			if err = d.Set("authentication_mode", vpnGatewayConnection.AuthenticationMode); err != nil {
				err = fmt.Errorf("Error setting authentication_mode: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-authentication_mode").GetDiag()
			}
			if err = d.Set("created_at", flex.DateTimeToString(vpnGatewayConnection.CreatedAt)); err != nil {
				err = fmt.Errorf("Error setting created_at: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-created_at").GetDiag()
			}

			if vpnGatewayConnection.DeadPeerDetection != nil {

				if err = d.Set(isVPNGatewayConnectionDeadPeerDetectionAction, vpnGatewayConnection.DeadPeerDetection.Action); err != nil {
					err = fmt.Errorf("Error setting action: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-action").GetDiag()
				}

				if err = d.Set(isVPNGatewayConnectionDeadPeerDetectionInterval, vpnGatewayConnection.DeadPeerDetection.Interval); err != nil {
					err = fmt.Errorf("Error setting interval: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-interval").GetDiag()
				}

				if err = d.Set(isVPNGatewayConnectionDeadPeerDetectionTimeout, vpnGatewayConnection.DeadPeerDetection.Timeout); err != nil {
					err = fmt.Errorf("Error setting timeout: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-timeout").GetDiag()
				}
			}
			if err = d.Set("href", vpnGatewayConnection.Href); err != nil {
				err = fmt.Errorf("Error setting href: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-href").GetDiag()
			}

			if vpnGatewayConnection.IkePolicy != nil {
				if err = d.Set("ike_policy", vpnGatewayConnection.IkePolicy.ID); err != nil {
					err = fmt.Errorf("Error setting ike_policy: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-ike_policy").GetDiag()
				}
			}

			if vpnGatewayConnection.IpsecPolicy != nil {
				if err = d.Set("ipsec_policy", vpnGatewayConnection.IpsecPolicy.ID); err != nil {
					err = fmt.Errorf("Error setting ipsec_policy: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-ipsec_policy").GetDiag()
				}
			}
			if err = d.Set("mode", vpnGatewayConnection.Mode); err != nil {
				err = fmt.Errorf("Error setting mode: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-mode").GetDiag()
			}
			if !core.IsNil(vpnGatewayConnection.DistributeTraffic) {
				if err = d.Set("distribute_traffic", vpnGatewayConnection.DistributeTraffic); err != nil {
					err = fmt.Errorf("Error setting distribute_traffic: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-distribute_traffic").GetDiag()
				}
			}
			if err = d.Set("name", vpnGatewayConnection.Name); err != nil {
				err = fmt.Errorf("Error setting name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-name").GetDiag()
			}

			// breaking changes
			if err = d.Set("establish_mode", vpnGatewayConnection.EstablishMode); err != nil {
				err = fmt.Errorf("Error setting establish_mode: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-establish_mode").GetDiag()
			}
			local := []map[string]interface{}{}
			if vpnGatewayConnection.Local != nil {
				modelMap, err := dataSourceIBMIsVPNGatewayConnectionVPNGatewayConnectionStaticRouteModeLocalToMap(vpnGatewayConnection.Local)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "local-to-map").GetDiag()
				}
				local = append(local, modelMap)
			}
			if err = d.Set("local", local); err != nil {
				err = fmt.Errorf("Error setting local: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-local").GetDiag()
			}

			peer := []map[string]interface{}{}
			if vpnGatewayConnection.Peer != nil {
				modelMap, err := dataSourceIBMIsVPNGatewayConnectionVPNGatewayConnectionStaticRouteModePeerToMap(vpnGatewayConnection.Peer)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "peer-to-map").GetDiag()
				}
				peer = append(peer, modelMap)
			}
			if err = d.Set("peer", peer); err != nil {
				err = fmt.Errorf("Error setting peer: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-peer").GetDiag()
			}
			// Deprecated
			if vpnGatewayConnection.Peer != nil {
				peer := vpnGatewayConnection.Peer.(*vpcv1.VPNGatewayConnectionStaticRouteModePeer)
				if err = d.Set("peer_address", peer.Address); err != nil {
					err = fmt.Errorf("Error setting peer_address: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-peer_address").GetDiag()
				}
			}
			if err = d.Set("preshared_key", vpnGatewayConnection.Psk); err != nil {
				err = fmt.Errorf("Error setting psk: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-psk").GetDiag()
			}
			if err = d.Set("resource_type", vpnGatewayConnection.ResourceType); err != nil {
				err = fmt.Errorf("Error setting resource_type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-resource_type").GetDiag()
			}
			if err = d.Set("status", vpnGatewayConnection.Status); err != nil {
				err = fmt.Errorf("Error setting status: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-status").GetDiag()
			}
			if err := d.Set("status_reasons", resourceVPNGatewayConnectionFlattenLifecycleReasons(vpnGatewayConnection.StatusReasons)); err != nil {
				err = fmt.Errorf("Error setting status_reasons: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-status_reasons").GetDiag()
			}
			if err = d.Set("routing_protocol", vpnGatewayConnection.RoutingProtocol); err != nil {
				err = fmt.Errorf("Error setting routing_protocol: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-routing_protocol").GetDiag()
			}

			if vpnGatewayConnection.Tunnels != nil {
				err = d.Set("tunnels", resourceVPNGatewayConnectionFlattenTunnels(vpnGatewayConnection.Tunnels))
				if err != nil {
					err = fmt.Errorf("Error setting tunnels: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-tunnels").GetDiag()
				}
			} else {
				d.Set("tunnels", []map[string]interface{}{})
			}
		}
	case "*vpcv1.VPNGatewayConnectionRouteMode":
		{
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteMode)
			d.SetId(fmt.Sprintf("%s/%s", vpn_gateway_id, *vpnGatewayConnection.ID))
			if err = d.Set("admin_state_up", vpnGatewayConnection.AdminStateUp); err != nil {
				err = fmt.Errorf("Error setting admin_state_up: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-admin_state_up").GetDiag()
			}
			if err = d.Set("authentication_mode", vpnGatewayConnection.AuthenticationMode); err != nil {
				err = fmt.Errorf("Error setting authentication_mode: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-authentication_mode").GetDiag()
			}
			if err = d.Set("created_at", flex.DateTimeToString(vpnGatewayConnection.CreatedAt)); err != nil {
				err = fmt.Errorf("Error setting created_at: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-created_at").GetDiag()
			}
			if !core.IsNil(vpnGatewayConnection.DistributeTraffic) {
				if err = d.Set("distribute_traffic", vpnGatewayConnection.DistributeTraffic); err != nil {
					err = fmt.Errorf("Error setting distribute_traffic: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-distribute_traffic").GetDiag()
				}
			}
			if vpnGatewayConnection.DeadPeerDetection != nil {
				if err = d.Set(isVPNGatewayConnectionDeadPeerDetectionAction, vpnGatewayConnection.DeadPeerDetection.Action); err != nil {
					err = fmt.Errorf("Error setting action: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-action").GetDiag()
				}

				if err = d.Set(isVPNGatewayConnectionDeadPeerDetectionInterval, vpnGatewayConnection.DeadPeerDetection.Interval); err != nil {
					err = fmt.Errorf("Error setting interval: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-interval").GetDiag()
				}

				if err = d.Set(isVPNGatewayConnectionDeadPeerDetectionTimeout, vpnGatewayConnection.DeadPeerDetection.Timeout); err != nil {
					err = fmt.Errorf("Error setting timeout: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-timeout").GetDiag()
				}
			}
			if err = d.Set("href", vpnGatewayConnection.Href); err != nil {
				err = fmt.Errorf("Error setting href: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-href").GetDiag()
			}

			if vpnGatewayConnection.IkePolicy != nil {
				if err = d.Set("ike_policy", vpnGatewayConnection.IkePolicy.ID); err != nil {
					err = fmt.Errorf("Error setting ike_policy: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-ike_policy").GetDiag()
				}
			}

			if vpnGatewayConnection.IpsecPolicy != nil {
				if err = d.Set("ipsec_policy", vpnGatewayConnection.IpsecPolicy.ID); err != nil {
					err = fmt.Errorf("Error setting ipsec_policy: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-ipsec_policy").GetDiag()
				}
			}
			if err = d.Set("mode", vpnGatewayConnection.Mode); err != nil {
				err = fmt.Errorf("Error setting mode: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-mode").GetDiag()
			}
			if err = d.Set("name", vpnGatewayConnection.Name); err != nil {
				err = fmt.Errorf("Error setting name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-name").GetDiag()
			}

			// breaking changes
			if err = d.Set("establish_mode", vpnGatewayConnection.EstablishMode); err != nil {
				err = fmt.Errorf("Error setting establish_mode: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-establish_mode").GetDiag()
			}
			local := []map[string]interface{}{}
			if vpnGatewayConnection.Local != nil {
				modelMap, err := dataSourceIBMIsVPNGatewayConnectionVPNGatewayConnectionStaticRouteModeLocalToMap(vpnGatewayConnection.Local)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "local-to-map").GetDiag()
				}
				local = append(local, modelMap)
			}
			if err = d.Set("local", local); err != nil {
				err = fmt.Errorf("Error setting status_reasons: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-local").GetDiag()
			}

			peer := []map[string]interface{}{}
			if vpnGatewayConnection.Peer != nil {
				modelMap, err := dataSourceIBMIsVPNGatewayConnectionVPNGatewayConnectionStaticRouteModePeerToMap(vpnGatewayConnection.Peer)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "peer-to-map").GetDiag()
				}
				peer = append(peer, modelMap)
			}
			if err = d.Set("peer", peer); err != nil {
				err = fmt.Errorf("Error setting peer: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-peer").GetDiag()
			}
			// Deprecated
			if vpnGatewayConnection.Peer != nil {
				peer := vpnGatewayConnection.Peer.(*vpcv1.VPNGatewayConnectionStaticRouteModePeer)
				if err = d.Set("peer_address", peer.Address); err != nil {
					err = fmt.Errorf("Error setting peer_address: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-peer_address").GetDiag()

				}
			}
			if err = d.Set("preshared_key", vpnGatewayConnection.Psk); err != nil {
				err = fmt.Errorf("Error setting psk: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-psk").GetDiag()
			}
			if err = d.Set("resource_type", vpnGatewayConnection.ResourceType); err != nil {
				err = fmt.Errorf("Error setting resource_type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-resource_type").GetDiag()
			}
			if err = d.Set("status", vpnGatewayConnection.Status); err != nil {
				err = fmt.Errorf("Error setting status: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-status").GetDiag()
			}
			if err := d.Set("status_reasons", resourceVPNGatewayConnectionFlattenLifecycleReasons(vpnGatewayConnection.StatusReasons)); err != nil {
				err = fmt.Errorf("Error setting status_reasons: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-status_reasons").GetDiag()

			}
			if err = d.Set("routing_protocol", vpnGatewayConnection.RoutingProtocol); err != nil {
				err = fmt.Errorf("Error setting routing_protocol: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-routing_protocol").GetDiag()
			}

			if vpnGatewayConnection.Tunnels != nil {
				err = d.Set("tunnels", resourceVPNGatewayConnectionFlattenTunnels(vpnGatewayConnection.Tunnels))
				if err != nil {
					err = fmt.Errorf("Error setting tunnels: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-tunnels").GetDiag()
				}
			} else {
				d.Set("tunnels", []map[string]interface{}{})
			}
		}
	case "*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionStaticRouteMode":
		{
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionStaticRouteMode)
			d.SetId(fmt.Sprintf("%s/%s", vpn_gateway_id, *vpnGatewayConnection.ID))
			if err = d.Set("admin_state_up", vpnGatewayConnection.AdminStateUp); err != nil {
				err = fmt.Errorf("Error setting admin_state_up: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-admin_state_up").GetDiag()
			}
			if err = d.Set("authentication_mode", vpnGatewayConnection.AuthenticationMode); err != nil {
				err = fmt.Errorf("Error setting authentication_mode: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-authentication_mode").GetDiag()
			}
			if err = d.Set("created_at", flex.DateTimeToString(vpnGatewayConnection.CreatedAt)); err != nil {
				err = fmt.Errorf("Error setting created_at: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-created_at").GetDiag()
			}
			if !core.IsNil(vpnGatewayConnection.DistributeTraffic) {
				if err = d.Set("distribute_traffic", vpnGatewayConnection.DistributeTraffic); err != nil {
					err = fmt.Errorf("Error setting distribute_traffic: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-distribute_traffic").GetDiag()
				}
			}
			if vpnGatewayConnection.DeadPeerDetection != nil {
				if err = d.Set(isVPNGatewayConnectionDeadPeerDetectionAction, vpnGatewayConnection.DeadPeerDetection.Action); err != nil {
					err = fmt.Errorf("Error setting action: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-action").GetDiag()
				}

				if err = d.Set(isVPNGatewayConnectionDeadPeerDetectionInterval, vpnGatewayConnection.DeadPeerDetection.Interval); err != nil {
					err = fmt.Errorf("Error setting interval: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-interval").GetDiag()
				}

				if err = d.Set(isVPNGatewayConnectionDeadPeerDetectionTimeout, vpnGatewayConnection.DeadPeerDetection.Timeout); err != nil {
					err = fmt.Errorf("Error setting timeout: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-timeout").GetDiag()
				}
			}
			if err = d.Set("href", vpnGatewayConnection.Href); err != nil {
				err = fmt.Errorf("Error setting href: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-href").GetDiag()
			}

			if vpnGatewayConnection.IkePolicy != nil {
				if err = d.Set("ike_policy", vpnGatewayConnection.IkePolicy.ID); err != nil {
					err = fmt.Errorf("Error setting ike_policy: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-ike_policy").GetDiag()
				}
			}

			if vpnGatewayConnection.IpsecPolicy != nil {
				if err = d.Set("ipsec_policy", vpnGatewayConnection.IpsecPolicy.ID); err != nil {
					err = fmt.Errorf("Error setting ipsec_policy: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-ipsec_policy").GetDiag()
				}
			}
			if err = d.Set("mode", vpnGatewayConnection.Mode); err != nil {
				err = fmt.Errorf("Error setting mode: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-mode").GetDiag()
			}
			if err = d.Set("name", vpnGatewayConnection.Name); err != nil {
				err = fmt.Errorf("Error setting name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-name").GetDiag()
			}

			// breaking changes
			if err = d.Set("establish_mode", vpnGatewayConnection.EstablishMode); err != nil {
				err = fmt.Errorf("Error setting establish_mode: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-establish_mode").GetDiag()
			}
			local := []map[string]interface{}{}
			if vpnGatewayConnection.Local != nil {
				modelMap, err := dataSourceIBMIsVPNGatewayConnectionVPNGatewayConnectionStaticRouteModeLocalToMap(vpnGatewayConnection.Local)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "local-to-map").GetDiag()
				}
				local = append(local, modelMap)
			}
			if err = d.Set("local", local); err != nil {
				err = fmt.Errorf("Error setting local: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-local").GetDiag()
			}

			peer := []map[string]interface{}{}
			if vpnGatewayConnection.Peer != nil {
				modelMap, err := dataSourceIBMIsVPNGatewayConnectionVPNGatewayConnectionStaticRouteModePeerToMap(vpnGatewayConnection.Peer)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "peer-to-map").GetDiag()
				}
				peer = append(peer, modelMap)
			}
			if err = d.Set("peer", peer); err != nil {
				err = fmt.Errorf("Error setting peer: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-peer").GetDiag()
			}
			// Deprecated
			if vpnGatewayConnection.Peer != nil {
				peer := vpnGatewayConnection.Peer.(*vpcv1.VPNGatewayConnectionStaticRouteModePeer)
				if err = d.Set("peer_address", peer.Address); err != nil {
					err = fmt.Errorf("Error setting peer_address: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-peer_address").GetDiag()

				}
			}
			if err = d.Set("preshared_key", vpnGatewayConnection.Psk); err != nil {
				err = fmt.Errorf("Error setting psk: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-psk").GetDiag()
			}
			if err = d.Set("resource_type", vpnGatewayConnection.ResourceType); err != nil {
				err = fmt.Errorf("Error setting resource_type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-resource_type").GetDiag()
			}
			if err = d.Set("status", vpnGatewayConnection.Status); err != nil {
				err = fmt.Errorf("Error setting status: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-status").GetDiag()
			}
			if err := d.Set("status_reasons", resourceVPNGatewayConnectionFlattenLifecycleReasons(vpnGatewayConnection.StatusReasons)); err != nil {
				err = fmt.Errorf("Error setting status_reasons: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-status_reasons").GetDiag()
			}
			if err = d.Set("routing_protocol", vpnGatewayConnection.RoutingProtocol); err != nil {
				err = fmt.Errorf("Error setting routing_protocol: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-routing_protocol").GetDiag()
			}

			if vpnGatewayConnection.Tunnels != nil {
				err = d.Set("tunnels", resourceVPNGatewayConnectionFlattenTunnels(vpnGatewayConnection.Tunnels))
				if err != nil {
					err = fmt.Errorf("Error setting tunnels: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-tunnels").GetDiag()
				}
			} else {
				d.Set("tunnels", []map[string]interface{}{})
			}
		}
	case "*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionDynamicRouteMode":
		{
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionDynamicRouteMode)
			d.SetId(fmt.Sprintf("%s/%s", vpn_gateway_id, *vpnGatewayConnection.ID))
			if err = d.Set("admin_state_up", vpnGatewayConnection.AdminStateUp); err != nil {
				err = fmt.Errorf("Error setting admin_state_up: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-admin_state_up").GetDiag()
			}
			if err = d.Set("authentication_mode", vpnGatewayConnection.AuthenticationMode); err != nil {
				err = fmt.Errorf("Error setting authentication_mode: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-authentication_mode").GetDiag()
			}
			if err = d.Set("created_at", flex.DateTimeToString(vpnGatewayConnection.CreatedAt)); err != nil {
				err = fmt.Errorf("Error setting created_at: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-created_at").GetDiag()
			}
			if !core.IsNil(vpnGatewayConnection.DistributeTraffic) {
				if err = d.Set("distribute_traffic", vpnGatewayConnection.DistributeTraffic); err != nil {
					err = fmt.Errorf("Error setting distribute_traffic: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-distribute_traffic").GetDiag()
				}
			}
			if vpnGatewayConnection.DeadPeerDetection != nil {
				if err = d.Set(isVPNGatewayConnectionDeadPeerDetectionAction, vpnGatewayConnection.DeadPeerDetection.Action); err != nil {
					err = fmt.Errorf("Error setting action: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-action").GetDiag()
				}

				if err = d.Set(isVPNGatewayConnectionDeadPeerDetectionInterval, vpnGatewayConnection.DeadPeerDetection.Interval); err != nil {
					err = fmt.Errorf("Error setting interval: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-interval").GetDiag()
				}

				if err = d.Set(isVPNGatewayConnectionDeadPeerDetectionTimeout, vpnGatewayConnection.DeadPeerDetection.Timeout); err != nil {
					err = fmt.Errorf("Error setting timeout: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-timeout").GetDiag()
				}
			}
			if err = d.Set("href", vpnGatewayConnection.Href); err != nil {
				err = fmt.Errorf("Error setting href: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-href").GetDiag()
			}

			if vpnGatewayConnection.IkePolicy != nil {
				if err = d.Set("ike_policy", vpnGatewayConnection.IkePolicy.ID); err != nil {
					err = fmt.Errorf("Error setting ike_policy: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-ike_policy").GetDiag()
				}
			}

			if vpnGatewayConnection.IpsecPolicy != nil {
				if err = d.Set("ipsec_policy", vpnGatewayConnection.IpsecPolicy.ID); err != nil {
					err = fmt.Errorf("Error setting ipsec_policy: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-ipsec_policy").GetDiag()
				}
			}
			if err = d.Set("mode", vpnGatewayConnection.Mode); err != nil {
				err = fmt.Errorf("Error setting mode: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-mode").GetDiag()
			}
			if err = d.Set("name", vpnGatewayConnection.Name); err != nil {
				err = fmt.Errorf("Error setting name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-name").GetDiag()
			}

			// breaking changes
			if err = d.Set("establish_mode", vpnGatewayConnection.EstablishMode); err != nil {
				err = fmt.Errorf("Error setting establish_mode: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-establish_mode").GetDiag()
			}
			local := []map[string]interface{}{}
			if vpnGatewayConnection.Local != nil {
				modelMap, err := dataSourceIBMIsVPNGatewayConnectionVPNGatewayConnectionDynamicRouteModeLocalToMap(vpnGatewayConnection.Local)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "local-to-map").GetDiag()

				}
				local = append(local, modelMap)
			}
			if err = d.Set("local", local); err != nil {
				err = fmt.Errorf("Error setting status_reasons: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-local").GetDiag()
			}

			peer := []map[string]interface{}{}
			if vpnGatewayConnection.Peer != nil {
				modelMap, err := dataSourceIBMIsVPNGatewayConnectionVPNGatewayConnectionDynamicRouteModePeerToMap(vpnGatewayConnection.Peer)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "peer-to-map").GetDiag()

				}
				peer = append(peer, modelMap)
			}
			if err = d.Set("peer", peer); err != nil {
				err = fmt.Errorf("Error setting peer: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-peer").GetDiag()
			}
			// Deprecated
			if vpnGatewayConnection.Peer != nil {
				peer := vpnGatewayConnection.Peer.(*vpcv1.VPNGatewayConnectionDynamicRouteModePeer)
				if err = d.Set("peer_address", peer.Address); err != nil {
					err = fmt.Errorf("Error setting peer_address: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-peer_address").GetDiag()
				}
			}
			if err = d.Set("preshared_key", vpnGatewayConnection.Psk); err != nil {
				err = fmt.Errorf("Error setting psk: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-psk").GetDiag()
			}
			if err = d.Set("resource_type", vpnGatewayConnection.ResourceType); err != nil {
				err = fmt.Errorf("Error setting resource_type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-resource_type").GetDiag()
			}
			if err = d.Set("status", vpnGatewayConnection.Status); err != nil {
				err = fmt.Errorf("Error setting status: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-status").GetDiag()
			}
			if err := d.Set("status_reasons", resourceVPNGatewayConnectionFlattenLifecycleReasons(vpnGatewayConnection.StatusReasons)); err != nil {
				err = fmt.Errorf("Error setting status_reasons: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-status_reasons").GetDiag()
			}
			if err = d.Set("routing_protocol", vpnGatewayConnection.RoutingProtocol); err != nil {
				err = fmt.Errorf("Error setting routing_protocol: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-routing_protocol").GetDiag()
			}

			if vpnGatewayConnection.Tunnels != nil {
				err = d.Set("tunnels", resourceVPNGatewayConnectionsFlattenDynamicTunnels(vpnGatewayConnection.Tunnels))
				if err != nil {
					err = fmt.Errorf("Error setting tunnels: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-tunnels").GetDiag()
				}
			} else {
				d.Set("tunnels", []map[string]interface{}{})
			}
		}
	case "*vpcv1.VPNGatewayConnectionPolicyMode":
		{
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionPolicyMode)
			d.SetId(fmt.Sprintf("%s/%s", vpn_gateway_id, *vpnGatewayConnection.ID))
			if err = d.Set("admin_state_up", vpnGatewayConnection.AdminStateUp); err != nil {
				err = fmt.Errorf("Error setting admin_state_up: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-admin_state_up").GetDiag()
			}
			if err = d.Set("authentication_mode", vpnGatewayConnection.AuthenticationMode); err != nil {
				err = fmt.Errorf("Error setting authentication_mode: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-authentication_mode").GetDiag()
			}
			if err = d.Set("created_at", flex.DateTimeToString(vpnGatewayConnection.CreatedAt)); err != nil {
				err = fmt.Errorf("Error setting created_at: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-created_at").GetDiag()
			}

			if vpnGatewayConnection.DeadPeerDetection != nil {
				if err = d.Set(isVPNGatewayConnectionDeadPeerDetectionAction, vpnGatewayConnection.DeadPeerDetection.Action); err != nil {
					err = fmt.Errorf("Error setting action: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-action").GetDiag()
				}

				if err = d.Set(isVPNGatewayConnectionDeadPeerDetectionInterval, vpnGatewayConnection.DeadPeerDetection.Interval); err != nil {
					err = fmt.Errorf("Error setting interval: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-interval").GetDiag()
				}

				if err = d.Set(isVPNGatewayConnectionDeadPeerDetectionTimeout, vpnGatewayConnection.DeadPeerDetection.Timeout); err != nil {
					err = fmt.Errorf("Error setting timeout: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-timeout").GetDiag()
				}
			}
			if err = d.Set("href", vpnGatewayConnection.Href); err != nil {
				err = fmt.Errorf("Error setting href: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-href").GetDiag()
			}

			if vpnGatewayConnection.IkePolicy != nil {
				if err = d.Set("ike_policy", vpnGatewayConnection.IkePolicy.ID); err != nil {
					err = fmt.Errorf("Error setting ike_policy: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-ike_policy").GetDiag()
				}

			}

			if vpnGatewayConnection.IpsecPolicy != nil {
				if err = d.Set("ipsec_policy", vpnGatewayConnection.IpsecPolicy.ID); err != nil {
					err = fmt.Errorf("Error setting ipsec_policy: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-ipsec_policy").GetDiag()
				}

			}
			if err = d.Set("mode", vpnGatewayConnection.Mode); err != nil {
				err = fmt.Errorf("Error setting mode: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-mode").GetDiag()
			}
			if err = d.Set("name", vpnGatewayConnection.Name); err != nil {
				err = fmt.Errorf("Error setting name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-name").GetDiag()
			}

			// breaking changes
			if err = d.Set("establish_mode", vpnGatewayConnection.EstablishMode); err != nil {
				err = fmt.Errorf("Error setting establish_mode: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-establish_mode").GetDiag()
			}
			local := []map[string]interface{}{}
			if vpnGatewayConnection.Local != nil {
				modelMap, err := dataSourceIBMIsVPNGatewayConnectionVPNGatewayConnectionPolicyModeLocalToMap(vpnGatewayConnection.Local)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "local-to-map").GetDiag()
				}
				local = append(local, modelMap)
			}
			if err = d.Set("local", local); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-local").GetDiag()
			}

			peer := []map[string]interface{}{}
			if vpnGatewayConnection.Peer != nil {
				modelMap, err := dataSourceIBMIsVPNGatewayConnectionVPNGatewayConnectionPolicyModePeerToMap(vpnGatewayConnection.Peer)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "peer-to-map").GetDiag()
				}
				peer = append(peer, modelMap)
			}
			if err = d.Set("peer", peer); err != nil {
				err = fmt.Errorf("Error setting peer: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-peer").GetDiag()
			}
			tunnels := []map[string]interface{}{}

			if err = d.Set("tunnels", tunnels); err != nil {
				err = fmt.Errorf("Error setting peer_address: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-tunnels").GetDiag()
			}

			// Deprecated
			if vpnGatewayConnection.Peer != nil {
				peer := vpnGatewayConnection.Peer.(*vpcv1.VPNGatewayConnectionPolicyModePeer)
				if err = d.Set("peer_address", peer.Address); err != nil {
					err = fmt.Errorf("Error setting peer_address: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-peer_address").GetDiag()
				}
				if len(peer.CIDRs) > 0 {
					err = d.Set("peer_cidrs", peer.CIDRs)
					if err != nil {
						err = fmt.Errorf("Error setting peer_cidrs: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-peer_cidrs").GetDiag()
					}
				}
			}
			if err = d.Set("preshared_key", vpnGatewayConnection.Psk); err != nil {
				err = fmt.Errorf("Error setting psk: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-psk").GetDiag()
			}
			if err = d.Set("resource_type", vpnGatewayConnection.ResourceType); err != nil {
				err = fmt.Errorf("Error setting resource_type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-resource_type").GetDiag()
			}
			if err = d.Set("status", vpnGatewayConnection.Status); err != nil {
				err = fmt.Errorf("Error setting status: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-status").GetDiag()
			}
			if err := d.Set("status_reasons", resourceVPNGatewayConnectionFlattenLifecycleReasons(vpnGatewayConnection.StatusReasons)); err != nil {
				err = fmt.Errorf("Error setting status_reasons: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-status_reasons").GetDiag()
			}
			// Deprecated
			if vpnGatewayConnection.Local != nil {
				local := vpnGatewayConnection.Local
				if len(local.CIDRs) > 0 {
					err = d.Set("local_cidrs", local.CIDRs)
					if err != nil {
						err = fmt.Errorf("Error setting local_cidrs: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_connection", "read", "set-local_cidrs").GetDiag()
					}
				}
			}

		}
	}

	return nil
}

func resourceVPNGatewayConnectionFlattenTunnels(result []vpcv1.VPNGatewayConnectionStaticRouteModeTunnel) (tunnels []map[string]interface{}) {
	for _, tunnelsItem := range result {
		tunnels = append(tunnels, resourceVPNGatewayConnectionTunnelsToMap(tunnelsItem))
	}

	return tunnels
}

func resourceVPNGatewayConnectionTunnelsToMap(tunnelsItem vpcv1.VPNGatewayConnectionStaticRouteModeTunnel) (tunnelsMap map[string]interface{}) {
	tunnelsMap = map[string]interface{}{}

	if tunnelsItem.PublicIP != nil {
		tunnelsMap["address"] = tunnelsItem.PublicIP.Address
	}
	if tunnelsItem.Status != nil {
		tunnelsMap["status"] = tunnelsItem.Status
	}

	return tunnelsMap
}

func resourceVPNGatewayConnectionsFlattenDynamicTunnels(result []vpcv1.VPNGatewayConnectionDynamicRouteModeTunnel) (tunnels []map[string]interface{}) {
	for _, tunnelsItem := range result {
		tunnels = append(tunnels, resourceVPNGatewayConnectionsDynamicTunnelsToMap(tunnelsItem))
	}

	return tunnels
}

func resourceVPNGatewayConnectionsDynamicTunnelsToMap(tunnelsItem vpcv1.VPNGatewayConnectionDynamicRouteModeTunnel) (tunnelsMap map[string]interface{}) {
	tunnelsMap = map[string]interface{}{}

	if tunnelsItem.NeighborIP != nil {
		tunnelsMap["neighbor_ip"] = tunnelsItem.NeighborIP.Address
	}
	if tunnelsItem.ProtocolState != nil {
		tunnelsMap["protocol_state"] = tunnelsItem.ProtocolState
	}
	if tunnelsItem.PublicIP != nil {
		tunnelsMap["address"] = tunnelsItem.PublicIP.Address
	}
	if tunnelsItem.Status != nil {
		tunnelsMap["status"] = tunnelsItem.Status
	}
	if tunnelsItem.TunnelInterfaceIP != nil {
		tunnelsMap["tunnel_interface_ip"] = tunnelsItem.TunnelInterfaceIP.Address
	}

	return tunnelsMap
}
