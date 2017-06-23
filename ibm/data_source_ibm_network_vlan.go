package ibm

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/services"
)

func dataSourceIBMNetworkVlan() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMNetworkVlanRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"number": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"router_hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"subnets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceIBMNetworkVlanRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetAccountService(sess)

	name := d.Get("name").(string)
	number := d.Get("number").(int)
	routerHostname := d.Get("router_hostname").(string)
	var vlan *datatypes.Network_Vlan
	var err error

	if number != 0 && routerHostname != "" {
		// Got vlan number and router, get vlan, and compute name
		vlan, err = getVlan(number, routerHostname, meta)
		if err != nil {
			return err
		}

		d.SetId(fmt.Sprintf("%d", *vlan.Id))
		if vlan.Name != nil {
			d.Set("name", *vlan.Name)
		}
	} else if name != "" {
		// Got name, get vlan, and compute router hostname and vlan number
		networkVlans, err := service.
			Mask("id,vlanNumber,name,primaryRouter[hostname],primarySubnets[networkIdentifier,cidr]").
			Filter(filter.Path("networkVlans.name").Eq(name).Build()).
			GetNetworkVlans()
		if err != nil {
			return fmt.Errorf("Error obtaining VLAN id: %s", err)
		} else if len(networkVlans) == 0 {
			return fmt.Errorf("No VLAN was found with the name '%s'", name)
		}

		vlan = &networkVlans[0]
		d.SetId(fmt.Sprintf("%d", *vlan.Id))
		d.Set("number", *vlan.VlanNumber)

		if vlan.PrimaryRouter != nil && vlan.PrimaryRouter.Hostname != nil {
			d.Set("router_hostname", *vlan.PrimaryRouter.Hostname)
		}
	} else {
		return errors.New("Missing required properties. Need a VLAN name, or the VLAN's number and router hostname.")
	}

	// Get subnets in cidr format for display
	if len(vlan.PrimarySubnets) > 0 {
		subnets := make([]string, len(vlan.PrimarySubnets))
		for i, subnet := range vlan.PrimarySubnets {
			subnets[i] = fmt.Sprintf("%s/%d", *subnet.NetworkIdentifier, *subnet.Cidr)
		}

		d.Set("subnets", subnets)
	}

	return nil
}

func getVlan(vlanNumber int, primaryRouterHostname string, meta interface{}) (*datatypes.Network_Vlan, error) {
	service := services.GetAccountService(meta.(ClientSession).SoftLayerSession())

	networkVlans, err := service.
		Mask("id,name,primarySubnets[networkIdentifier,cidr]").
		Filter(
			filter.Build(
				filter.Path("networkVlans.primaryRouter.hostname").Eq(primaryRouterHostname),
				filter.Path("networkVlans.vlanNumber").Eq(vlanNumber),
			),
		).
		GetNetworkVlans()

	if err != nil {
		return &datatypes.Network_Vlan{}, fmt.Errorf("Error looking up Vlan: %s", err)
	}

	if len(networkVlans) < 1 {
		return &datatypes.Network_Vlan{}, fmt.Errorf(
			"Unable to locate a vlan matching the provided router hostname and vlan number: %s/%d",
			primaryRouterHostname,
			vlanNumber)
	}

	return &networkVlans[0], nil
}
