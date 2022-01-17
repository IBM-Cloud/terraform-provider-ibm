// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIBMSatconClusterGroup() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMSatconClusterGroupCreate,
		Read:     resourceIBMSatconClusterGroupRead,
		Update:   resourceIBMSatconClusterGroupUpdate,
		Delete:   resourceIBMSatconClusterGroupDelete,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"uuid": {
				Description: "ID of the clustergroup",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name or id of the clustergroup",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"created": {
				Description: "Creation time of the clustergroup",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"clusters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": {
							Description: "ID of the cluster",
							Type:        schema.TypeString,
							Required:    true,
						},
						"name": {
							Description: "Name of the cluster",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceIBMSatconClusterGroupCreate(d *schema.ResourceData, meta interface{}) error {
	satconClient, err := meta.(ClientSession).SatellitConfigClientSession()
	if err != nil {
		return err
	}

	satconGroupAPI := satconClient.Groups

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	name := d.Get("name").(string)
	log.Printf("[DEBUG] create clustergroup with name: %s , userid: %v\n", name, userDetails.userAccount)
	addDetails, err := satconGroupAPI.AddGroup(userDetails.userAccount, name)
	if err != nil {
		log.Printf("[DEBUG] resourceIBMSatconClusterGroupCreate AddGroup failed with: %v\n", err)
		return fmt.Errorf("error creating satellite clustergroup: %s", err)
	}

	d.SetId(name)
	d.Set("uuid", addDetails.UUID)

	//TODO Wait for clusters to be able to attach to group

	return resourceIBMSatconClusterGroupRead(d, meta)
}

func resourceIBMSatconClusterGroupRead(d *schema.ResourceData, meta interface{}) error {
	groupName := d.Id()

	if groupName == "" {
		return fmt.Errorf("satellite clustergroup name is empty")
	}

	satconClient, err := meta.(ClientSession).SatellitConfigClientSession()
	if err != nil {
		return err
	}

	satconGroupAPI := satconClient.Groups

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] get clustergroup with name: %s , userid: %v\n", name, userDetails.userAccount)

	group, err := satconGroupAPI.GroupByName(userDetails.userAccount, groupName)
	if err != nil {
		return fmt.Errorf("error retrieving satellite clustergroup: %s", err)
	}

	clusters := make([]map[string]interface{}, 0)
	for _, c := range group.Clusters {
		cluster := map[string]interface{}{
			"id":   c.ID,
			"name": c.Name,
		}
		clusters = append(clusters, cluster)
	}

	d.Set("uuid", group.UUID)
	d.Set("name", group.Name)
	d.Set("created", group.Created)
	d.Set("clusters", clusters)

	return nil
}

func resourceIBMSatconClusterGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	//TODO
	return nil
}

func resourceIBMSatconClusterGroupDelete(d *schema.ResourceData, meta interface{}) error {
	groupName := d.Id()

	if groupName == "" {
		return fmt.Errorf("satellite clustergroup name is empty")
	}

	satconClient, err := meta.(ClientSession).SatellitConfigClientSession()
	if err != nil {
		return err
	}

	satconGroupAPI := satconClient.Groups

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] remove clustergroup with name: %s , userid: %v\n", name, userDetails.userAccount)

	removeDetails, err := satconGroupAPI.RemoveGroupByName(userDetails.userAccount, groupName)
	if err != nil {
		return fmt.Errorf("failed deleting satellite clustergroup: %s", err)
	}

	log.Printf("[INFO] Removed satellite clustergroup with name: %s, uuid: %s", groupName, removeDetails.UUID)

	d.SetId("")
	return nil
}
