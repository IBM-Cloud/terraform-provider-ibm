package ibm

import (
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	pdnsPermittedNetworks = "permitted_networks"
)

func dataSourceIBMPrivateDNSPermittedNetworks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMPrivateDNSPermittedNetworksRead,

		Schema: map[string]*schema.Schema{

			pdnsInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Instance ID",
			},

			pdnsZoneID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Zone Id",
			},

			pdnsPermittedNetworks: {

				Type:        schema.TypeList,
				Description: "Collection of permitted networks",
				Computed:    true,
				Elem: &schema.Resource{

					Schema: map[string]*schema.Schema{

						pdnsPermittedNetworkID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network Id",
						},

						pdnsInstanceID: {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Instance Id",
						},

						pdnsZoneID: {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Zone Id",
						},

						pdnsNetworkType: {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							Default:      "vpc",
							ValidateFunc: validateAllowedStringValue([]string{"vpc"}),
							Description:  "Network Type",
						},

						pdnsPermittedNetwork: {
							Type:        schema.TypeMap,
							Description: "permitted network",
							Computed:    true,
							Elem: &schema.Resource{

								Schema: map[string]*schema.Schema{

									pdnsVpcCRN: {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: "VPC CRN id",
									},
								},
							},
						},

						pdnsPermittedNetworkCreatedOn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network creation date",
						},

						pdnsPermittedNetworkModifiedOn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network Modification date",
						},

						pdnsPermittedNetworkState: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network status",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPrivateDNSPermittedNetworksRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	instanceID := d.Get(pdnsInstanceID).(string)
	dnsZoneID := d.Get(pdnsZoneID).(string)
	listPermittedNetworkOptions := sess.NewListPermittedNetworksOptions(instanceID, dnsZoneID)
	availablePermittedNetworks, detail, err := sess.ListPermittedNetworks(listPermittedNetworkOptions)
	if err != nil {
		log.Printf("Error reading list of permitted networks:%s", detail)
		return err
	}

	permittedNetworks := make([]map[string]interface{}, 0)

	for _, instance := range availablePermittedNetworks.PermittedNetworks {
		permittedNetwork := map[string]interface{}{}
		permittedNetworkVpcCrn := map[string]interface{}{}
		permittedNetwork[pdnsInstanceID] = instanceID
		permittedNetwork[pdnsPermittedNetworkID] = instance.ID
		permittedNetwork[pdnsPermittedNetworkCreatedOn] = instance.CreatedOn
		permittedNetwork[pdnsPermittedNetworkModifiedOn] = instance.ModifiedOn
		permittedNetwork[pdnsPermittedNetworkState] = instance.State
		permittedNetwork[pdnsNetworkType] = instance.Type
		permittedNetworkVpcCrn[pdnsVpcCRN] = instance.PermittedNetwork.VpcCrn
		permittedNetwork[pdnsPermittedNetwork] = permittedNetworkVpcCrn
		permittedNetwork[pdnsZoneID] = dnsZoneID

		permittedNetworks = append(permittedNetworks, permittedNetwork)
		log.Println("inside array : ", *instance.ID, *instance.CreatedOn, *instance.ModifiedOn)
	}
	d.SetId(dataSourceIBMPrivateDNSPermittedNetworkID(d))
	d.Set(pdnsPermittedNetworks, permittedNetworks)
	return nil
}

// dataSourceIBMPrivateDnsPermittedNetworkID returns a reasonable ID for dns permitted network list.
func dataSourceIBMPrivateDNSPermittedNetworkID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
