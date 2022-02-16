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
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
)

// Attributes and Arguments defined in data_source_ibm_pi_network_port.go
func ResourceIBMPINetworkPort() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPINetworkPortCreate,
		ReadContext:   resourceIBMPINetworkPortRead,
		UpdateContext: resourceIBMPINetworkPortUpdate,
		DeleteContext: resourceIBMPINetworkPortDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			PINetworkPortName: {
				Type:     schema.TypeString,
				Required: true,
			},
			PICloudInstanceID: {
				Type:     schema.TypeString,
				Required: true,
			},
			PINetworkPortDescription: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			PINetworkPortIP: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			//Computed Attributes
			NetworkPortMAC: {
				Type:     schema.TypeString,
				Computed: true,
			},
			NetworkPortID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			NetworkPortStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
			NetworkPortPublicIP: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMPINetworkPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	networkname := d.Get(PINetworkPortName).(string)
	description := d.Get(PINetworkPortDescription).(string)

	ipaddress := d.Get(PINetworkPortIP).(string)

	nwportBody := &models.NetworkPortCreate{Description: description}

	if ipaddress != "" {
		log.Printf("IP address provided. ")
		nwportBody.IPAddress = ipaddress
	}

	client := st.NewIBMPINetworkClient(ctx, sess, cloudInstanceID)

	networkPortResponse, err := client.CreatePort(networkname, nwportBody)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("Printing the networkresponse %+v", &networkPortResponse)

	IBMPINetworkPortID := *networkPortResponse.PortID

	d.SetId(fmt.Sprintf("%s/%s/%s", cloudInstanceID, networkname, IBMPINetworkPortID))

	_, err = isWaitForIBMPINetworkPortAvailable(ctx, client, IBMPINetworkPortID, networkname, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIBMPINetworkPortRead(ctx, d, meta)
}

func resourceIBMPINetworkPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	cloudInstanceID := parts[0]
	networkname := parts[1]
	portID := parts[2]

	networkC := st.NewIBMPINetworkClient(ctx, sess, cloudInstanceID)
	networkdata, err := networkC.GetPort(networkname, portID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set(PINetworkPortIP, networkdata.IPAddress)
	d.Set(PINetworkPortDescription, networkdata.Description)
	d.Set(NetworkPortMAC, networkdata.MacAddress)
	d.Set(NetworkPortStatus, networkdata.Status)
	d.Set(NetworkPortID, networkdata.PortID)
	d.Set(NetworkPortPublicIP, networkdata.ExternalIP)

	return nil
}

func resourceIBMPINetworkPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceIBMPINetworkPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	log.Printf("Calling the network delete functions. ")
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	cloudInstanceID := parts[0]
	networkname := parts[1]
	portID := parts[2]

	client := st.NewIBMPINetworkClient(ctx, sess, cloudInstanceID)

	log.Printf("Calling the delete with the following params delete with cloud instance (%s) and networkid (%s) and portid (%s) ", cloudInstanceID, networkname, portID)
	err = client.DeletePort(networkname, portID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func isWaitForIBMPINetworkPortAvailable(ctx context.Context, client *st.IBMPINetworkClient, id string, networkname string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Power Network (%s) that was created for Network Zone (%s) to be available.", id, networkname)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", "build"},
		Target:     []string{"DOWN"},
		Refresh:    isIBMPINetworkPortRefreshFunc(client, id, networkname),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Minute,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isIBMPINetworkPortRefreshFunc(client *st.IBMPINetworkClient, id, networkname string) resource.StateRefreshFunc {

	log.Printf("Calling the IsIBMPINetwork Refresh Function....with the following id (%s) for network port and following id (%s) for network name and waiting for network to be READY", id, networkname)
	return func() (interface{}, string, error) {
		network, err := client.GetPort(networkname, id)
		if err != nil {
			return nil, "", err
		}

		if &network.PortID != nil {
			//if network.State == "available" {
			log.Printf(" The port has been created with the following ip address and attached to an instance ")
			return network, "DOWN", nil
		}

		return network, "build", nil
	}
}
