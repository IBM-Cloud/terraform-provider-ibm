package ibm

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataIBMCISFirewallRecord() *schema.Resource {
	return &schema.Resource{
		Read: dataIBMCISFirewallRecordRead,
		Schema: map[string]*schema.Schema{
			"cis_id": {
				Type:        schema.TypeString,
				Description: "CIS object id",
				Required:    true,
			},
			"domain_id": {
				Type:        schema.TypeString,
				Description: "Associated CIS domain",
				Required:    true,
			},
			"firewall_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of firewall.Allowable values are access-rules,ua-rules,lockdowns",
			},
			"lockdown": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Lockdown json Data",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lockdown_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"paused": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"urls": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"configurations": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
func dataIBMCISFirewallRecordRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	cisID := d.Get("cis_id").(string)
	zoneID := d.Get("domain_id").(string)
	firewallType := d.Get("firewall_type").(string)

	recordList, err := cisClient.Firewall().ListFirewall(cisID, zoneID, firewallType)
	if err != nil {
		log.Printf("[WARN] Error getting zone during Firewall Record List %v\n", err)
		return err
	}

	record := make([]map[string]interface{}, 0, len(recordList))
	if firewallType == "lockdowns" {
		for _, l := range recordList {
			configurationList := make([]map[string]interface{}, 0, len(l.Configurations))
			for _, c := range l.Configurations {
				configuration := make(map[string]interface{})
				configuration["target"] = c.Target
				configuration["value"] = c.Value
				configurationList = append(configurationList, configuration)
			}
			lockdown := make(map[string]interface{})
			lockdown["lockdown_id"] = l.ID
			lockdown["paused"] = l.Paused
			lockdown["priority"] = l.Priority
			lockdown["configurations"] = configurationList
			lockdown["urls"] = l.Urls
			record = append(record, lockdown)
		}
		d.Set("lockdown", record)
	}
	d.SetId(firewallType + ":" + cisID)
	d.Set("cis_id", cisID)
	d.Set("domain_id", zoneID)
	d.Set("firewall_type", firewallType)

	return nil
}
