package ibm

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

var UnSupportedFields4Update = []string{
	"service",
	"customer",
	"region",
	"tcp_range",
	"udp_range",
	"max_speed",
	"dedicated",
	"multi_tenant",
}

func resourceCSEInstance() *schema.Resource {

	return &schema.Resource{
		Create:   resourceCSEInstanceCreate,
		Read:     resourceCSEInstanceRead,
		Update:   resourceCSEInstanceUpdate,
		Delete:   resourceCSEInstanceDelete,
		Exists:   resourceCSEInstanceExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"service": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The service name for the CSE.",
				Required:    true,
			},
			"customer": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The customer name for the CSE.",
				Required:    true,
			},
			"service_addresses": &schema.Schema{
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The service private addresses for the CSE.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
			},
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The region to create CSE.",
				Required:    true,
			},
			"data_centers": &schema.Schema{
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The data centers to create CSE.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
			},
			"tcp_ports": &schema.Schema{
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The exposed tcp ports for the CSE.",
				Elem:        &schema.Schema{Type: schema.TypeInt},
				//Set:         schema.HashInt,
				Set: HashInt,
			},
			"udp_ports": &schema.Schema{
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The exposed udp ports for the CSE.",
				Elem:        &schema.Schema{Type: schema.TypeInt},
				//Set:         schema.HashInt,
				Set: HashInt,
			},
			"tcp_range": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The exposed tcp port range for the CSE.",
				Optional:    true,
			},
			"udp_range": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The exposed udp port range for the CSE.",
				Optional:    true,
			},
			"max_speed": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The maxium network speed for the CSE.",
				Optional:    true,
			},
			"estado_proto": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The estado protocol for the CSE, value could be tcp, http or https.",
				Optional:    true,
			},
			"estado_port": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The estado port for the CSE",
				Optional:    true,
			},
			"estado_path": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The estado path for the CSE",
				Optional:    true,
			},
			"dedicated": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The value is 1 will create the CSE in dedicated device, is 0 will create in shared device.",
				Optional:    true,
			},
			"multi_tenant": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The dedicated device will only run this CSE when the value is 1.",
				Optional:    true,
			},
			"acl": &schema.Schema{
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The acl setting for this CSE.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"static_addresses": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func resourceCSEInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	cseClient, err := meta.(ClientSession).CseAPI()
	if err != nil {
		return err
	}

	seAPI := cseClient.ServiceEndpoints()

	payload := genCreatePayload(d)

	log.Printf("resourceCSEInstanceCreate: payload=%v\n", payload)

	srvId, err := seAPI.CreateServiceEndpoint(payload)

	if err != nil {
		return err
	}

	d.SetId(srvId)
	d.Set("service_id", srvId)

	return nil
}

func resourceCSEInstanceRead(d *schema.ResourceData, meta interface{}) error {
	cseClient, err := meta.(ClientSession).CseAPI()
	if err != nil {
		return err
	}

	seAPI := cseClient.ServiceEndpoints()

	srvId := d.Get("service_id").(string)

	srvObj, err := seAPI.GetServiceEndpoint(srvId)

	if err != nil {
		return err
	}

	d.Set("url", srvObj.Service.URL)
	addresses := []string{}
	for _, se := range srvObj.Endpoints {
		addresses = append(addresses, se.StaticAddress)
	}

	d.Set("static_addresses", addresses)

	return nil
}

func resourceCSEInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	payload, err := genUpdatePayload(d)
	if err != nil {
		return err
	}

	if len(payload) == 0 {
		return errors.New("No things to change")
	}

	log.Printf("resourceCSEInstanceUpdate: payload=%v", payload)

	cseClient, err := meta.(ClientSession).CseAPI()
	if err != nil {
		return err
	}

	seAPI := cseClient.ServiceEndpoints()

	srvId := d.Get("service_id").(string)

	err = seAPI.UpdateServiceEndpoint(srvId, payload)

	if err != nil {
		return err
	}

	return resourceCSEInstanceRead(d, meta)
}

func resourceCSEInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	cseClient, err := meta.(ClientSession).CseAPI()
	if err != nil {
		return err
	}

	seAPI := cseClient.ServiceEndpoints()

	srvId := d.Get("service_id").(string)

	err = seAPI.DeleteServiceEndpoint(srvId)

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceCSEInstanceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return true, nil
}

func genCreatePayload(d *schema.ResourceData) map[string]interface{} {
	payload := map[string]interface{}{}
	payload["service"] = d.Get("service").(string)
	payload["customer"] = d.Get("customer").(string)
	payload["serviceAddresses"] = expandStringList(d.Get("service_addresses").(*schema.Set).List())
	payload["region"] = d.Get("region").(string)
	payload["dataCenters"] = expandStringList(d.Get("data_centers").(*schema.Set).List())

	if tcpPorts, ok := d.GetOk("tcp_ports"); ok {
		payload["tcpports"] = expandIntList(tcpPorts.(*schema.Set).List())
	}

	if udpPorts, ok := d.GetOk("udp_ports"); ok {
		payload["udpports"] = expandIntList(udpPorts.(*schema.Set).List())
	}

	if tcpRange, ok := d.GetOk("tcp_range"); ok {
		payload["tcpportrange"] = tcpRange.(string)
	}

	if udpRange, ok := d.GetOk("udp_range"); ok {
		payload["udpportrange"] = udpRange.(string)
	}

	if speed, ok := d.GetOk("max_speed"); ok {
		payload["maxSpeed"] = speed.(string)
	}

	if estadoProto, ok := d.GetOk("estado_proto"); ok {
		payload["estadoProto"] = estadoProto.(string)
	}

	if estadoPort, ok := d.GetOk("estado_port"); ok {
		payload["estadoPort"] = estadoPort.(int)
	}

	if estadoPath, ok := d.GetOk("estado_path"); ok {
		payload["estadoPath"] = estadoPath.(string)
	}

	if dedicated, ok := d.GetOkExists("dedicated"); ok {
		payload["dedicated"] = dedicated.(int)
	}

	if multiTenant, ok := d.GetOkExists("multi_tenant"); ok {
		payload["multitenant"] = multiTenant.(int)
	}

	if acl, ok := d.GetOk("acl"); ok {
		payload["acl"] = expandStringList(acl.(*schema.Set).List())
	}

	return payload
}

func genUpdatePayload(d *schema.ResourceData) (map[string]interface{}, error) {
	ret := map[string]interface{}{}

	if d.HasChange("service_addresses") {
		_, newv := d.GetChange("service_addresses")
		ret["serviceAddresses"] = expandStringList(newv.(*schema.Set).List())
	} else {
		if v, ok := d.GetOk("service_addresses"); ok {
			ret["serviceAddresses"] = expandStringList(v.(*schema.Set).List())
		}
	}

	if d.HasChange("estado_proto") {
		_, newv := d.GetChange("estado_proto")
		ret["estadoProto"] = newv.(string)
	} else {
		if v, ok := d.GetOk("estado_proto"); ok {
			ret["estadoProto"] = v.(string)
		}
	}

	if d.HasChange("estado_port") {
		_, newv := d.GetChange("estado_port")
		ret["estadoPort"] = newv.(int)
	} else {
		if v, ok := d.GetOk("estado_port"); ok {
			ret["estadoPort"] = v.(int)
		}
	}

	if d.HasChange("estado_path") {
		_, newv := d.GetChange("estado_path")
		ret["estadoPath"] = newv.(string)
	} else {
		if v, ok := d.GetOk("estado_path"); ok {
			ret["estadoPath"] = v.(string)
		}
	}

	if d.HasChange("tcp_ports") {
		_, newv := d.GetChange("tcp_ports")
		ret["tcpports"] = expandIntList(newv.(*schema.Set).List())
	} else {
		if v, ok := d.GetOk("tcp_ports"); ok {
			ret["tcpports"] = expandIntList(v.(*schema.Set).List())
		}
	}

	if d.HasChange("udp_ports") {
		_, newv := d.GetChange("udp_ports")
		ret["udpports"] = expandIntList(newv.(*schema.Set).List())
	} else {
		if v, ok := d.GetOk("udp_ports"); ok {
			ret["udpports"] = expandIntList(v.(*schema.Set).List())
		}
	}

	if d.HasChange("data_centers") {
		_, newv := d.GetChange("data_centers")
		ret["dataCenters"] = expandStringList(newv.(*schema.Set).List())
	} else {
		if v, ok := d.GetOk("data_centers"); ok {
			ret["dataCenters"] = expandStringList(v.(*schema.Set).List())
		}
	}

	if d.HasChange("acl") {
		_, newv := d.GetChange("acl")
		ret["acl"] = expandStringList(newv.(*schema.Set).List())
	} else {
		if v, ok := d.GetOk("acl"); ok {
			ret["acl"] = expandStringList(v.(*schema.Set).List())
		}
	}

	for _, v := range UnSupportedFields4Update {
		if d.HasChange(v) {
			return ret, errors.New("Unsupported update field:" + v)
		}
	}

	return ret, nil
}
