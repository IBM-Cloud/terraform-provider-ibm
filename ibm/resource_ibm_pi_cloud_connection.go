// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_cloud_connections"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

const (
	PICloudConnectionId = "cloud_connection_id"
)

func resourceIBMPICloudConnection() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPICloudConnectionCreate,
		ReadContext:   resourceIBMPICloudConnectionRead,
		UpdateContext: resourceIBMPICloudConnectionUpdate,
		DeleteContext: resourceIBMPICloudConnectionDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			// Required Attributes
			helpers.PICloudInstanceId: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PI cloud instance ID",
			},
			helpers.PICloudConnectionName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the cloud connection",
			},
			helpers.PICloudConnectionSpeed: {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validateAllowedIntValue([]int{50, 100, 200, 500, 1000, 2000, 5000, 10000}),
				Description:  "Speed of the cloud connection (speed in megabits per second)",
			},

			// Optional Attributes
			helpers.PICloudConnectionGlobalRouting: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable global routing for this cloud connection",
			},
			helpers.PICloudConnectionMetered: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable metered for this cloud connection",
			},
			helpers.PICloudConnectionNetworks: {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Set of Networks to attach to this cloud connection",
			},
			helpers.PICloudConnectionClassicEnabled: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable classic endpoint destination",
			},
			helpers.PICloudConnectionClassicGreCidr: {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{helpers.PICloudConnectionClassicEnabled, helpers.PICloudConnectionClassicGreDest},
				Description:  "GRE network in CIDR notation",
			},
			helpers.PICloudConnectionClassicGreDest: {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{helpers.PICloudConnectionClassicEnabled, helpers.PICloudConnectionClassicGreCidr},
				Description:  "GRE destination IP address",
			},
			helpers.PICloudConnectionVPCEnabled: {
				Type:         schema.TypeBool,
				Optional:     true,
				Default:      false,
				RequiredWith: []string{helpers.PICloudConnectionVPCCRNs},
				Description:  "Enable VPC for this cloud connection",
			},
			helpers.PICloudConnectionVPCCRNs: {
				Type:         schema.TypeSet,
				Optional:     true,
				Elem:         &schema.Schema{Type: schema.TypeString},
				RequiredWith: []string{helpers.PICloudConnectionVPCEnabled},
				Description:  "Set of VPCs to attach to this cloud connection",
			},

			//Computed Attributes
			PICloudConnectionId: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloud connection ID",
			},
			PICloudConnectionStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Link status",
			},
			PICloudConnectionIBMIPAddress: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IBM IP address",
			},
			PICloudConnectionUserIPAddress: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User IP address",
			},
			PICloudConnectionPort: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Port",
			},
			PICloudConnectionClassicGreSource: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GRE auto-assigned source IP address",
			},
		},
	}
}

func resourceIBMPICloudConnectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(helpers.PICloudInstanceId).(string)
	name := d.Get(helpers.PICloudConnectionName).(string)
	speed := int64(d.Get(helpers.PICloudConnectionSpeed).(int))

	body := &models.CloudConnectionCreate{
		Name:  &name,
		Speed: &speed,
	}
	if v, ok := d.GetOk(helpers.PICloudConnectionGlobalRouting); ok {
		body.GlobalRouting = v.(bool)
	}
	if v, ok := d.GetOk(helpers.PICloudConnectionMetered); ok {
		body.Metered = v.(bool)
	}
	// networks
	if v, ok := d.GetOk(helpers.PICloudConnectionNetworks); ok && v.(*schema.Set).Len() > 0 {
		body.Subnets = expandStringList(v.(*schema.Set).List())
	}
	// classic
	if v, ok := d.GetOk(helpers.PICloudConnectionClassicEnabled); ok {
		classicEnabled := v.(bool)
		classic := &models.CloudConnectionEndpointClassicUpdate{
			Enabled: classicEnabled,
		}
		if v, ok := d.GetOk(helpers.PICloudConnectionClassicGreCidr); ok {
			greCIDR := v.(string)
			classic.Gre.Cidr = &greCIDR
		}
		if v, ok := d.GetOk(helpers.PICloudConnectionClassicGreDest); ok {
			greDest := v.(string)
			classic.Gre.DestIPAddress = &greDest
		}
		body.Classic = classic
	}

	// VPC
	if v, ok := d.GetOk(helpers.PICloudConnectionVPCEnabled); ok {
		vpcEnabled := v.(bool)
		vpc := &models.CloudConnectionEndpointVPC{
			Enabled: vpcEnabled,
		}
		if v, ok := d.GetOk(helpers.PICloudConnectionVPCCRNs); ok && v.(*schema.Set).Len() > 0 {
			vpcIds := expandStringList(v.(*schema.Set).List())
			vpcs := make([]*models.CloudConnectionVPC, len(vpcIds))
			for i, vpcId := range vpcIds {
				vpcs[i] = &models.CloudConnectionVPC{
					VpcID: &vpcId,
				}
			}
			vpc.Vpcs = vpcs
		}
		body.Vpc = vpc
	}

	client := st.NewIBMPICloudConnectionClient(sess, cloudInstanceID)
	cloudConnection, cloudConnectionJob, err := client.CreateWithContext(ctx, body, cloudInstanceID)
	if err != nil {
		log.Printf("[DEBUG] create cloud connection failed %v", err)
		return diag.Errorf(errors.CreateCloudConnectionOperationFailed, cloudInstanceID, err)
	}
	var cloudConnectionID string
	if cloudConnection != nil {
		cloudConnectionID = *cloudConnection.CloudConnectionID
	} else if cloudConnectionJob != nil {
		cloudConnectionID = *cloudConnectionJob.CloudConnectionID
		jobID := *cloudConnectionJob.JobRef.ID

		client := st.NewIBMPIJobClient(sess, cloudInstanceID)
		_, err = waitForIBMPIJobCompleted(ctx, client, jobID, cloudInstanceID, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return diag.Errorf(errors.CreateCloudConnectionOperationFailed, cloudInstanceID, err)
		}
	}

	d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, cloudConnectionID))

	return resourceIBMPICloudConnectionRead(ctx, d, meta)
}

func resourceIBMPICloudConnectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := parts[0]
	cloudConnectionID := parts[1]

	ccName := d.Get(helpers.PICloudConnectionName).(string)
	ccSpeed := int64(d.Get(helpers.PICloudConnectionSpeed).(int))

	client := st.NewIBMPICloudConnectionClient(sess, cloudInstanceID)
	jobClient := st.NewIBMPIJobClient(sess, cloudInstanceID)

	if d.HasChangesExcept(helpers.PICloudConnectionNetworks) {

		body := &models.CloudConnectionUpdate{
			Name:  &ccName,
			Speed: &ccSpeed,
		}
		if v, ok := d.GetOk(helpers.PICloudConnectionGlobalRouting); ok {
			globalRouting := v.(bool)
			body.GlobalRouting = &globalRouting
		}
		if v, ok := d.GetOk(helpers.PICloudConnectionMetered); ok {
			metered := v.(bool)
			body.Metered = &metered
		}
		// classic
		if v, ok := d.GetOk(helpers.PICloudConnectionClassicEnabled); ok {
			classicEnabled := v.(bool)
			classic := &models.CloudConnectionEndpointClassicUpdate{
				Enabled: classicEnabled,
			}
			if v, ok := d.GetOk(helpers.PICloudConnectionClassicGreCidr); ok {
				greCIDR := v.(string)
				classic.Gre.Cidr = &greCIDR
			}
			if v, ok := d.GetOk(helpers.PICloudConnectionClassicGreDest); ok {
				greDest := v.(string)
				classic.Gre.DestIPAddress = &greDest
			}
			body.Classic = classic
		} else {
			// need to disable classic if not provided
			classic := &models.CloudConnectionEndpointClassicUpdate{
				Enabled: false,
			}
			body.Classic = classic
		}
		// vpc
		if v, ok := d.GetOk(helpers.PICloudConnectionVPCEnabled); ok {
			vpcEnabled := v.(bool)
			vpc := &models.CloudConnectionEndpointVPC{
				Enabled: vpcEnabled,
			}
			if v, ok := d.GetOk(helpers.PICloudConnectionVPCCRNs); ok && v.(*schema.Set).Len() > 0 {
				vpcIds := expandStringList(v.(*schema.Set).List())
				vpcs := make([]*models.CloudConnectionVPC, len(vpcIds))
				for i, vpcId := range vpcIds {
					vpcs[i] = &models.CloudConnectionVPC{
						VpcID: &vpcId,
					}
				}
				vpc.Vpcs = vpcs
			}
			body.Vpc = vpc
		} else {
			// need to disable VPC if not provided
			vpc := &models.CloudConnectionEndpointVPC{
				Enabled: false,
			}
			body.Vpc = vpc
		}

		_, cloudConnectionJob, err := client.UpdateWithContext(ctx, cloudConnectionID, cloudInstanceID, body)
		if err != nil {
			return diag.Errorf(errors.UpdateCloudConnectionOperationFailed, cloudConnectionID, err)
		}
		if cloudConnectionJob != nil {
			_, err = waitForIBMPIJobCompleted(ctx, jobClient, *cloudConnectionJob.ID, cloudInstanceID, d.Timeout(schema.TimeoutCreate))
			if err != nil {
				return diag.Errorf(errors.UpdateCloudConnectionOperationFailed, cloudConnectionID, err)
			}
		}
	}
	if d.HasChange(helpers.PICloudConnectionNetworks) {
		oldRaw, newRaw := d.GetChange(helpers.PICloudConnectionNetworks)
		old := oldRaw.(*schema.Set)
		new := newRaw.(*schema.Set)

		toAdd := new.Difference(old)
		toRemove := old.Difference(new)

		// call network add api for each toAdd
		for _, n := range expandStringList(toAdd.List()) {
			jobReference, err := client.AddNetworkWithContext(ctx, n, cloudConnectionID, cloudInstanceID)
			if err != nil {
				return diag.Errorf("failed to perform the network add operation... %v", err)
			}
			_, err = waitForIBMPIJobCompleted(ctx, jobClient, *jobReference.ID, cloudInstanceID, d.Timeout(schema.TimeoutUpdate))
			if err != nil {
				return diag.Errorf(errors.UpdateCloudConnectionOperationFailed, cloudConnectionID, err)
			}
		}

		// call network delete api for each toRemove
		for _, n := range expandStringList(toRemove.List()) {
			jobReference, err := client.DeleteNetworkWithContext(ctx, n, cloudConnectionID, cloudInstanceID)
			if err != nil {
				return diag.Errorf("failed to perform the network delete operation... %v", err)
			}
			_, err = waitForIBMPIJobCompleted(ctx, jobClient, *jobReference.ID, cloudInstanceID, d.Timeout(schema.TimeoutUpdate))
			if err != nil {
				return diag.Errorf(errors.UpdateCloudConnectionOperationFailed, cloudConnectionID, err)
			}
		}
	}

	return resourceIBMPICloudConnectionRead(ctx, d, meta)
}

func resourceIBMPICloudConnectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := parts[0]
	cloudConnectionID := parts[1]

	client := st.NewIBMPICloudConnectionClient(sess, cloudInstanceID)
	cloudConnection, err := client.GetWithContext(ctx, cloudConnectionID, cloudInstanceID)
	if err != nil {
		switch err.(type) {
		case *p_cloud_cloud_connections.PcloudCloudconnectionsGetNotFound:
			log.Printf("[DEBUG] cloud connection does not exist %v", err)
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] get cloud connection failed %v", err)
		return diag.Errorf(errors.GetCloudConnectionOperationFailed, cloudConnectionID, err)
	}

	d.Set(PICloudConnectionId, cloudConnection.CloudConnectionID)
	d.Set(helpers.PICloudConnectionName, cloudConnection.Name)
	d.Set(helpers.PICloudConnectionGlobalRouting, cloudConnection.GlobalRouting)
	d.Set(helpers.PICloudConnectionMetered, cloudConnection.Metered)
	d.Set(PICloudConnectionIBMIPAddress, cloudConnection.IbmIPAddress)
	d.Set(PICloudConnectionUserIPAddress, cloudConnection.UserIPAddress)
	d.Set(PICloudConnectionStatus, cloudConnection.LinkStatus)
	d.Set(PICloudConnectionPort, cloudConnection.Port)
	d.Set(helpers.PICloudConnectionSpeed, cloudConnection.Speed)
	d.Set(helpers.PICloudInstanceId, cloudInstanceID)
	if cloudConnection.Networks != nil {
		networks := make([]string, 0)
		for _, ccNetwork := range cloudConnection.Networks {
			if ccNetwork != nil {
				networks = append(networks, *ccNetwork.NetworkID)
			}
		}
		d.Set(helpers.PICloudConnectionNetworks, networks)
	}
	if cloudConnection.Classic != nil {
		d.Set(helpers.PICloudConnectionClassicEnabled, cloudConnection.Classic.Enabled)
		if cloudConnection.Classic.Gre != nil {
			d.Set(helpers.PICloudConnectionClassicGreDest, cloudConnection.Classic.Gre.DestIPAddress)
			d.Set(PICloudConnectionClassicGreSource, cloudConnection.Classic.Gre.SourceIPAddress)
		}
	}
	if cloudConnection.Vpc != nil {
		d.Set(helpers.PICloudConnectionVPCEnabled, cloudConnection.Vpc.Enabled)
		if cloudConnection.Vpc.Vpcs != nil && len(cloudConnection.Vpc.Vpcs) > 0 {
			vpcCRNs := make([]string, len(cloudConnection.Vpc.Vpcs))
			for i, vpc := range cloudConnection.Vpc.Vpcs {
				vpcCRNs[i] = *vpc.VpcID
			}
			d.Set(helpers.PICloudConnectionVPCCRNs, vpcCRNs)
		}
	}

	return nil
}
func resourceIBMPICloudConnectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := parts[0]
	cloudConnectionID := parts[1]

	client := st.NewIBMPICloudConnectionClient(sess, cloudInstanceID)
	_, err = client.GetWithContext(ctx, cloudConnectionID, cloudInstanceID)
	if err != nil {
		switch err.(type) {
		case *p_cloud_cloud_connections.PcloudCloudconnectionsGetNotFound:
			log.Printf("[DEBUG] cloud connection does not exist %v", err)
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] get cloud connection failed %v", err)
		return diag.Errorf(errors.GetCloudConnectionOperationFailed, cloudConnectionID, err)
	}
	log.Printf("[INFO] Found cloud connection with id %s", cloudConnectionID)

	_, deleteJob, err := client.DeleteWithContext(ctx, cloudConnectionID, cloudInstanceID)
	if err != nil {
		log.Printf("[DEBUG] delete cloud connection failed %v", err)
		return diag.Errorf(errors.DeleteCloudConnectionOperationFailed, cloudConnectionID, err)
	}
	if deleteJob != nil {
		jobID := *deleteJob.ID

		client := st.NewIBMPIJobClient(sess, cloudInstanceID)
		_, err = waitForIBMPIJobCompleted(ctx, client, jobID, cloudInstanceID, d.Timeout(schema.TimeoutDelete))
		if err != nil {
			return diag.Errorf(errors.DeleteCloudConnectionOperationFailed, cloudConnectionID, err)
		}
	}

	d.SetId("")
	return nil
}
