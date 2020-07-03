package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/networking-go-sdk/transitgatewayapisv1"
)

func dataSourceIBMTransitGateway() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMTransitGatewayRead,

		Schema: map[string]*schema.Schema{
			tgName: {
				Type:     schema.TypeString,
				Required: true,
			},
			tgCrn: {
				Type:     schema.TypeString,
				Computed: true,
			},
			tgLocation: {
				Type:     schema.TypeString,
				Computed: true,
			},
			tgCreatedAt: {
				Type:     schema.TypeString,
				Computed: true,
			},
			tgGlobal: {
				Type:     schema.TypeBool,
				Computed: true,
			},
			tgStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
			tgUpdatedAt: {
				Type:     schema.TypeString,
				Computed: true,
			},
			tgResourceGroup: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMTransitGatewayRead(d *schema.ResourceData, meta interface{}) error {

	client, err := transitgatewayClient(meta)
	if err != nil {
		return err
	}

	listTransitGatewaysOptionsModel := &transitgatewayapisv1.ListTransitGatewaysOptions{}
	listTransitGateways, response, err := client.ListTransitGateways(listTransitGatewaysOptionsModel)
	if err != nil {
		return fmt.Errorf("Error while listing transit gateways %s\n%s", err, response)
	}

	gwName := d.Get(tgName).(string)
	var foundGateway bool
	for _, tgw := range listTransitGateways.TransitGateways {

		if *tgw.Name == gwName {
			d.SetId(*tgw.ID)
			d.Set(tgCrn, tgw.Crn)
			d.Set(tgName, tgw.Name)
			d.Set(tgLocation, tgw.Location)
			d.Set(tgCreatedAt, tgw.CreatedAt.String())

			if tgw.UpdatedAt != nil {
				d.Set(tgUpdatedAt, tgw.UpdatedAt.String())
			}
			d.Set(tgGlobal, tgw.Global)
			d.Set(tgStatus, tgw.Status)

			if tgw.ResourceGroup != nil {
				rg := tgw.ResourceGroup
				d.Set(tgResourceGroup, *rg.ID)
			}
			foundGateway = true
		}
	}

	if !foundGateway {
		return fmt.Errorf(
			"Couldn't find any gateway with the specified name: (%s)", gwName)
	}

	return nil

}
