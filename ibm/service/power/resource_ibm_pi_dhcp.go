// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_service_d_h_c_p"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

// Attributes and Arguments defined in data_source_ibm_pi_dhcp.go
func ResourceIBMPIDhcp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIDhcpCreate,
		ReadContext:   resourceIBMPIDhcpRead,
		DeleteContext: resourceIBMPIDhcpDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			// Required Attributes
			PICloudInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PI cloud instance ID",
				ForceNew:    true,
			},
			PIDhcpCloudConnection: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The cloud connection uuid to connect with DHCP private network",
				ForceNew:    true,
			},

			//Computed Attributes
			DhcpID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the DHCP Server",
			},
			DhcpStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the DHCP Server",
			},
			DhcpNetwork: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The DHCP Server private network",
			},
			DhcpLeases: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of DHCP Server PVM Instance leases",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						DhcpInstanceIP: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP of the PVM Instance",
						},
						DhcpInstanceMAC: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The MAC Address of the PVM Instance",
						},
					},
				},
			},
		},
	}
}

func resourceIBMPIDhcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)

	body := &models.DHCPServerCreate{}
	if c, ok := d.GetOk(PIDhcpCloudConnection); ok {
		body.CloudConnectionID = c.(string)
	}

	client := st.NewIBMPIDhcpClient(ctx, sess, cloudInstanceID)
	dhcpServer, err := client.Create(body)
	if err != nil {
		log.Printf("[DEBUG] create DHCP failed %v", err)
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, *dhcpServer.ID))

	_, err = waitForIBMDHCPStatus(ctx, client, *dhcpServer.ID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		diag.FromErr(err)
	}

	return resourceIBMPIDhcpRead(ctx, d, meta)
}

func resourceIBMPIDhcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, dhcpID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := st.NewIBMPIDhcpClient(ctx, sess, cloudInstanceID)
	dhcpServer, err := client.Get(dhcpID)
	if err != nil {
		uErr := errors.Unwrap(err)
		switch uErr.(type) {
		case *p_cloud_service_d_h_c_p.PcloudDhcpGetNotFound:
			log.Printf("[DEBUG] dhcp does not exist %v", err)
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] get DHCP failed %v", err)
		return diag.FromErr(err)
	}

	d.Set(DhcpID, *dhcpServer.ID)
	d.Set(DhcpStatus, *dhcpServer.Status)

	if dhcpServer.Network != nil {
		d.Set(DhcpNetwork, dhcpServer.Network.ID)
	}
	if dhcpServer.Leases != nil {
		leaseList := make([]map[string]string, len(dhcpServer.Leases))
		for i, lease := range dhcpServer.Leases {
			leaseList[i] = map[string]string{
				DhcpInstanceIP:  *lease.InstanceIP,
				DhcpInstanceMAC: *lease.InstanceMacAddress,
			}
		}
		d.Set(DhcpLeases, leaseList)
	}

	return nil
}
func resourceIBMPIDhcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, dhcpID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := st.NewIBMPIDhcpClient(ctx, sess, cloudInstanceID)
	err = client.Delete(dhcpID)
	if err != nil {
		uErr := errors.Unwrap(err)
		switch uErr.(type) {
		case *p_cloud_service_d_h_c_p.PcloudDhcpDeleteNotFound:
			log.Printf("[DEBUG] dhcp does not exist %v", err)
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] delete DHCP failed %v", err)
		return diag.FromErr(err)
	}
	_, err = waitForIBMDHCPStatusDeleted(ctx, client, dhcpID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func waitForIBMDHCPStatus(ctx context.Context, client *st.IBMPIDhcpClient, dhcpID string, timeout time.Duration) (interface{}, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"Building"},
		Target:  []string{"ACTIVE"},
		Refresh: func() (interface{}, string, error) {
			dhcpServer, err := client.Get(dhcpID)
			if err != nil {
				log.Printf("[DEBUG] get DHCP failed %v", err)
				return nil, "", err
			}
			if *dhcpServer.Status != "ACTIVE" {
				return dhcpServer, "Building", nil
			}
			return dhcpServer, *dhcpServer.Status, nil
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForStateContext(ctx)
}

func waitForIBMDHCPStatusDeleted(ctx context.Context, client *st.IBMPIDhcpClient, dhcpID string, timeout time.Duration) (interface{}, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"Deleting"},
		Target:  []string{"Deleted"},
		Refresh: func() (interface{}, string, error) {
			dhcpServer, err := client.Get(dhcpID)
			if err != nil {
				log.Printf("[DEBUG] dhcp does not exist %v", err)
				return dhcpServer, "Deleted", nil
			}
			return dhcpServer, "Deleting", nil
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForStateContext(ctx)
}
