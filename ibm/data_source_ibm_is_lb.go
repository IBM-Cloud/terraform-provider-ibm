package ibm

import (
	"fmt"
	"log"

	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMISLB() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISLBRead,

		Schema: map[string]*schema.Schema{
			isLBName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Load Balancer name",
			},

			isLBType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load Balancer type",
			},

			isLBStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load Balancer status",
			},

			isLBOperatingStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load Balancer operating status",
			},

			isLBPublicIPs: {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Load Balancer Public IPs",
			},

			isLBPrivateIPs: {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Load Balancer private IPs",
			},

			isLBSubnets: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "Load Balancer subnets list",
			},

			isLBTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         resourceIBMVPCHash,
				Description: "Tags associated to Load Balancer",
			},

			isLBResourceGroup: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load Balancer Resource group",
			},

			isLBHostName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load Balancer Host Name",
			},

			isLBListeners: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "Load Balancer Listeners list",
			},

			isLBPools: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "Load Balancer Pools list",
			},

			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},

			ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func dataSourceIBMISLBRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	name := d.Get(isLBName).(string)
	if userDetails.generation == 1 {
		err := classiclbGetbyName(d, meta, name)
		if err != nil {
			return err
		}
	} else {
		err := lbGetByName(d, meta, name)
		if err != nil {
			return err
		}
	}
	return nil
}

func classiclbGetbyName(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	listLoadBalancersOptions := &vpcclassicv1.ListLoadBalancersOptions{}
	lbs, response, err := sess.ListLoadBalancers(listLoadBalancersOptions)
	if err != nil {
		return fmt.Errorf("Error Fetching Load Balancers %s\n%s", err, response)
	}
	for _, lb := range lbs.LoadBalancers {
		if *lb.Name == name {
			d.SetId(*lb.ID)
			d.Set("id", *lb.ID)
			d.Set(isLBName, *lb.Name)
			if *lb.IsPublic {
				d.Set(isLBType, "public")
			} else {
				d.Set(isLBType, "private")
			}
			d.Set(isLBStatus, *lb.ProvisioningStatus)
			d.Set(isLBOperatingStatus, *lb.OperatingStatus)
			publicIpList := make([]string, 0)
			if lb.PublicIps != nil {
				for _, ip := range lb.PublicIps {
					if ip.Address != nil {
						pubip := *ip.Address
						publicIpList = append(publicIpList, pubip)
					}
				}
			}
			d.Set(isLBPublicIPs, publicIpList)
			privateIpList := make([]string, 0)
			if lb.PrivateIps != nil {
				for _, ip := range lb.PrivateIps {
					if ip.Address != nil {
						prip := *ip.Address
						privateIpList = append(privateIpList, prip)
					}
				}
			}
			d.Set(isLBPrivateIPs, privateIpList)
			if lb.Subnets != nil {
				subnetList := make([]string, 0)
				for _, subnet := range lb.Subnets {
					if subnet.ID != nil {
						sub := *subnet.ID
						subnetList = append(subnetList, sub)
					}
				}
				d.Set(isLBSubnets, subnetList)
			}
			if lb.Listeners != nil {
				listenerList := make([]string, 0)
				for _, listener := range lb.Listeners {
					if listener.ID != nil {
						lis := *listener.ID
						listenerList = append(listenerList, lis)
					}
				}
				d.Set(isLBListeners, listenerList)
			}
			if lb.Pools != nil {
				poolList := make([]string, 0)
				for _, pool := range lb.Pools {
					if pool.ID != nil {
						p := *pool.ID
						poolList = append(poolList, p)
					}
				}
				d.Set(isLBPools, poolList)
			}
			d.Set(isLBResourceGroup, *lb.ResourceGroup.ID)
			d.Set(isLBHostName, *lb.Hostname)
			tags, err := GetTagsUsingCRN(meta, *lb.CRN)
			if err != nil {
				log.Printf(
					"Error on get of resource vpc Load Balancer (%s) tags: %s", d.Id(), err)
			}
			d.Set(isLBTags, tags)
			controller, err := getBaseController(meta)
			if err != nil {
				return err
			}
			d.Set(ResourceControllerURL, controller+"/vpc/network/loadBalancers")
			d.Set(ResourceName, *lb.Name)
			if lb.ResourceGroup != nil {
				d.Set(ResourceGroupName, *lb.ResourceGroup.ID)
			}
			return nil
		}
	}
	return fmt.Errorf("No Load balancer found with name %s", name)
}

func lbGetByName(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	listLoadBalancersOptions := &vpcv1.ListLoadBalancersOptions{}
	lbs, response, err := sess.ListLoadBalancers(listLoadBalancersOptions)
	if err != nil {
		return fmt.Errorf("Error Fetching Load Balancers %s\n%s", err, response)
	}
	for _, lb := range lbs.LoadBalancers {
		if *lb.Name == name {
			d.SetId(*lb.ID)
			d.Set("id", *lb.ID)
			d.Set(isLBName, *lb.Name)
			if *lb.IsPublic {
				d.Set(isLBType, "public")
			} else {
				d.Set(isLBType, "private")
			}
			d.Set(isLBStatus, *lb.ProvisioningStatus)
			d.Set(isLBOperatingStatus, *lb.OperatingStatus)
			publicIpList := make([]string, 0)
			if lb.PublicIps != nil {
				for _, ip := range lb.PublicIps {
					if ip.Address != nil {
						pubip := *ip.Address
						publicIpList = append(publicIpList, pubip)
					}
				}
			}
			d.Set(isLBPublicIPs, publicIpList)
			privateIpList := make([]string, 0)
			if lb.PrivateIps != nil {
				for _, ip := range lb.PrivateIps {
					if ip.Address != nil {
						prip := *ip.Address
						privateIpList = append(privateIpList, prip)
					}
				}
			}
			d.Set(isLBPrivateIPs, privateIpList)
			if lb.Subnets != nil {
				subnetList := make([]string, 0)
				for _, subnet := range lb.Subnets {
					if subnet.ID != nil {
						sub := *subnet.ID
						subnetList = append(subnetList, sub)
					}
				}
				d.Set(isLBSubnets, subnetList)
			}
			if lb.Listeners != nil {
				listenerList := make([]string, 0)
				for _, listener := range lb.Listeners {
					if listener.ID != nil {
						lis := *listener.ID
						listenerList = append(listenerList, lis)
					}
				}
				d.Set(isLBListeners, listenerList)
			}
			if lb.Pools != nil {
				poolList := make([]string, 0)
				for _, pool := range lb.Pools {
					if pool.ID != nil {
						p := *pool.ID
						poolList = append(poolList, p)
					}
				}
				d.Set(isLBPools, poolList)
			}
			d.Set(isLBResourceGroup, *lb.ResourceGroup.ID)
			d.Set(isLBHostName, *lb.Hostname)
			tags, err := GetTagsUsingCRN(meta, *lb.CRN)
			if err != nil {
				log.Printf(
					"Error on get of resource vpc Load Balancer (%s) tags: %s", d.Id(), err)
			}
			d.Set(isLBTags, tags)
			controller, err := getBaseController(meta)
			if err != nil {
				return err
			}
			d.Set(ResourceControllerURL, controller+"/vpc-ext/network/loadBalancers")
			d.Set(ResourceName, *lb.Name)
			if lb.ResourceGroup != nil {
				d.Set(ResourceGroupName, *lb.ResourceGroup.ID)
			}
			return nil
		}
	}
	return fmt.Errorf("No Load balancer found with name %s", name)
}
