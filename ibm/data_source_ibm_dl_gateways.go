package ibm

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/networking-go-sdk/directlinkapisv1"
)

const (
	dlGateways   = "gateways"
	dlGatewaysId = "id"
)

func dataSourceIBMDLGateways() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMDLGatewaysRead,
		Schema: map[string]*schema.Schema{
			dlGateways: {
				Type:        schema.TypeList,
				Description: "Collection of direct link gateways",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						dlGatewaysId: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Id of the data source gateways",
						},
						dlBgpAsn: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "BGP ASN",
						},
						dlBgpBaseCidr: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "BGP base CIDR",
						},
						dlBgpCerCidr: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "BGP customer edge router CIDR",
						},
						dlBgpIbmAsn: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "IBM BGP ASN",
						},
						dlBgpIbmCidr: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "BGP IBM CIDR",
						},
						dlBgpStatus: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway BGP status",
						},
						dlCompletionNoticeRejectReason: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Reason for completion notice rejection",
						},
						dlCreatedAt: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time resource was created",
						},
						dlCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN (Cloud Resource Name) of this gateway",
						},
						dlCrossConnectRouter: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cross connect router",
						},
						dlDedicatedHostingID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Dedicated host id",
						},
						dlGlobal: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Gateways with global routing (true) can connect to networks outside their associated region",
						},
						dlLinkStatus: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway link status",
						},
						dlLocationDisplayName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway location long name",
						},
						dlLocationName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway location",
						},
						dlMetered: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Metered billing option",
						},
						dlName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique user-defined name for this gateway",
						},
						dlOperationalStatus: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway operational status",
						},
						dlPort: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway port",
						},
						dlProviderAPIManaged: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether gateway was created through a provider portal",
						},
						dlResourceGroup: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway resource group",
						},
						dlSpeedMbps: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Gateway speed in megabits per second",
						},
						dlType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway type",
						},
						dlVlan: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "VLAN allocated for this gateway",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMDLGatewaysRead(d *schema.ResourceData, meta interface{}) error {
	directLink, err := directlinkClient(meta)
	if err != nil {
		return err
	}
	listGatewaysOptionsModel := &directlinkapisv1.ListGatewaysOptions{}
	listGateways, _, err := directLink.ListGateways(listGatewaysOptionsModel)
	if err != nil {
		return err
	}
	gateways := make([]map[string]interface{}, 0)
	for _, instance := range listGateways.Gateways {
		gateway := map[string]interface{}{}
		if instance.ID != nil {
			gateway["id"] = *instance.ID
		}
		if instance.Name != nil {
			gateway[dlName] = *instance.Name
		}
		if instance.Crn != nil {
			gateway[dlCrn] = *instance.Crn
		}
		if instance.BgpAsn != nil {
			gateway[dlBgpAsn] = *instance.BgpAsn
		}
		if instance.BgpIbmCidr != nil {
			gateway[dlBgpIbmCidr] = *instance.BgpIbmCidr
		}
		if instance.BgpIbmAsn != nil {
			gateway[dlBgpIbmAsn] = *instance.BgpIbmAsn
		}
		if instance.Metered != nil {
			gateway[dlMetered] = *instance.Metered
		}
		if instance.CrossConnectRouter != nil {
			gateway[dlCrossConnectRouter] = *instance.CrossConnectRouter
		}
		if instance.BgpBaseCidr != nil {
			gateway[dlBgpBaseCidr] = *instance.BgpBaseCidr
		}
		if instance.BgpCerCidr != nil {
			gateway[dlBgpCerCidr] = *instance.BgpCerCidr
		}
		if instance.DedicatedHostingID != nil {
			gateway[dlDedicatedHostingID] = *instance.DedicatedHostingID
		}
		if instance.ProviderApiManaged != nil {
			gateway[dlProviderAPIManaged] = *instance.ProviderApiManaged
		}
		if instance.Type != nil {
			gateway[dlType] = *instance.Type
		}
		if instance.SpeedMbps != nil {
			gateway[dlSpeedMbps] = *instance.SpeedMbps
		}
		if instance.OperationalStatus != nil {
			gateway[dlOperationalStatus] = *instance.OperationalStatus
		}
		if instance.BgpStatus != nil {
			gateway[dlBgpStatus] = *instance.BgpStatus
		}
		if instance.LocationName != nil {
			gateway[dlLocationName] = *instance.LocationName
		}
		if instance.LocationDisplayName != nil {
			gateway[dlLocationDisplayName] = *instance.LocationDisplayName
		}
		if instance.Vlan != nil {
			gateway[dlVlan] = *instance.Vlan
		}
		if instance.Global != nil {
			gateway[dlGlobal] = *instance.Global
		}
		if instance.Port != nil {
			gateway[dlPort] = *instance.Port.ID
		}
		if instance.LinkStatus != nil {
			gateway[dlLinkStatus] = *instance.LinkStatus
		}
		if instance.CreatedAt != nil {
			gateway[dlCreatedAt] = instance.CreatedAt.String()
		}
		if instance.ResourceGroup != nil {
			rg := instance.ResourceGroup
			gateway[dlResourceGroup] = *rg.ID
		}
		gateways = append(gateways, gateway)
	}
	d.SetId(dataSourceIBMDLGatewaysID(d))
	d.Set(dlGateways, gateways)
	return nil
}

// dataSourceIBMDLGatewaysID returns a reasonable ID for a direct link gateways list.
func dataSourceIBMDLGatewaysID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
