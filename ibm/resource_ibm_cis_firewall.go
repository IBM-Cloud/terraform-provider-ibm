package ibm

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
)

func resourceIBMCISFirewallRecord() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMCISFirewallRecordCreate,
		Read:     resourceIBMCISFirewallRecordRead,
		Update:   resourceIBMCISFirewallRecordUpdate,
		Delete:   resourceIBMCISFirewallRecordDelete,
		Exists:   resourceIBMCISFirewallRecordExists,
		Importer: &schema.ResourceImporter{},

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
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Type of firewall.Allowable values are access-rules,ua-rules,lockdowns",
				ValidateFunc: validateAllowedStringValue([]string{"lockdowns", "access_rules", "ua_rules"}),
			},

			"lockdown": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Lockdown json Data",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lockdown_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"paused": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"urls": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"configurations": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
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

func resourceIBMCISFirewallRecordCreate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	cisID := d.Get("cis_id").(string)
	zoneID := d.Get("domain_id").(string)
	firewallType := d.Get("firewall_type").(string)

	newRecord := v1.FirewallBody{}

	if lockdown, ok := d.GetOk("lockdown"); ok {
		lockdownlist := lockdown.([]interface{})
		for _, l := range lockdownlist {
			lockdownMap, _ := l.(map[string]interface{})

			if paused := lockdownMap["paused"]; paused != nil {
				newRecord.Paused = paused.(bool)
			}
			if description := lockdownMap["description"]; description != nil {
				newRecord.Description = description.(string)
			}
			if priority := lockdownMap["priority"]; priority != nil {
				newRecord.Priority = priority.(int)
			}
			var urls = make([]string, 0)
			if u, ok := lockdownMap["urls"]; ok && u != nil {
				for _, url := range u.([]interface{}) {
					urls = append(urls, fmt.Sprintf("%v", url))
				}
				newRecord.Urls = urls
			}
			var configurationList = make([]v1.Configuration, 0)
			if res, ok := lockdownMap["configurations"]; ok {
				configurations := res.([]interface{})
				for _, c := range configurations {
					r, _ := c.(map[string]interface{})
					if target, value := r["target"], r["value"]; target != nil && value != nil {
						configurationRecord := v1.Configuration{}
						configurationRecord.Target, configurationRecord.Value = target.(string), value.(string)
						configurationList = append(configurationList, configurationRecord)
					}
				}
				newRecord.Configurations = configurationList
			}
		}
	}

	var recordPtr *v1.FirewallRecord
	recordPtr, err = cisClient.Firewall().CreateFirewall(cisID, zoneID, firewallType, newRecord)
	if err != nil {
		return fmt.Errorf("Failed to create Firewall: %s", err)
	}

	// In the Event that the API returns an empty Firewall, we verify that the
	// ID returned is not the default ""

	record := *recordPtr

	if record.ID == "" {
		return fmt.Errorf("Failed to find record in Create response; Record was empty")
	}

	d.SetId(convertCisToTfFourVar(firewallType, record.ID, zoneID, cisID))

	return resourceIBMCISFirewallRecordRead(d, meta)
}

func resourceIBMCISFirewallRecordRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	firewallType, recordID, zoneID, cisID, _ := convertTfToCisFourVar(d.Id())
	if err != nil {
		return err
	}

	var recordPtr *v1.FirewallRecord
	recordPtr, err = cisClient.Firewall().GetFirewall(cisID, zoneID, firewallType, recordID)
	if err != nil {
		if strings.Contains(err.Error(), "Request failed with status code: 404") {
			d.SetId("")
			return nil
		}
		log.Printf("[WARN] Error getting zone during Firewall Read %v\n", err)
		return err
	}

	record := *recordPtr
	d.Set("cis_id", cisID)
	d.Set("firewall_type", firewallType)
	d.Set("domain_id", zoneID)

	if &record.Paused != nil || record.Urls != nil || record.Configurations != nil {
		configuration := make([]map[string]interface{}, 0, len(record.Configurations))
		for _, c := range record.Configurations {
			r := map[string]interface{}{
				"target": c.Target,
				"value":  c.Value,
			}
			configuration = append(configuration, r)
		}
		lockdown := map[string]interface{}{
			"lockdown_id":    recordID,
			"paused":         record.Paused,
			"urls":           record.Urls,
			"configurations": configuration,
		}
		if &record.Description != nil {
			lockdown["description"] = record.Description
		}
		if &record.Priority != nil {
			lockdown["priority"] = record.Priority
		}
		lockdowns := make([]map[string]interface{}, 0, 1)
		lockdowns = append(lockdowns, lockdown)

		d.Set("lockdown", lockdowns)

	}
	return nil
}

func resourceIBMCISFirewallRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	firewallType, recordID, zoneID, cisID, _ := convertTfToCisFourVar(d.Id())
	if err != nil {
		return err
	}
	updateRecord := v1.FirewallBody{}
	if d.HasChange("lockdown") {
		if lockdown, ok := d.GetOk("lockdown"); ok {
			lockdownlist := lockdown.([]interface{})
			for _, l := range lockdownlist {
				lockdownMap, _ := l.(map[string]interface{})
				if paused := lockdownMap["paused"]; paused != nil {
					updateRecord.Paused = paused.(bool)
				}
				if description := lockdownMap["description"]; description != nil {
					updateRecord.Description = description.(string)
				}
				if priority := lockdownMap["priority"]; priority != nil {
					updateRecord.Priority = priority.(int)
				}
				var urls = make([]string, 0)
				if u, ok := lockdownMap["urls"]; ok && u != nil {
					for _, url := range u.([]interface{}) {
						urls = append(urls, fmt.Sprintf("%v", url))
					}
					updateRecord.Urls = urls
				}
				var configurationList = make([]v1.Configuration, 0)
				if res, ok := lockdownMap["configurations"]; ok {
					configurations := res.([]interface{})
					for _, c := range configurations {
						r, _ := c.(map[string]interface{})
						if target, value := r["target"], r["value"]; target != nil && value != nil {
							configurationRecord := v1.Configuration{}
							configurationRecord.Target, configurationRecord.Value = target.(string), value.(string)
							configurationList = append(configurationList, configurationRecord)
						}
					}
					updateRecord.Configurations = configurationList
				}
			}
		}
	}
	_, err = cisClient.Firewall().UpdateFirewall(cisID, zoneID, firewallType, recordID, updateRecord)
	if err != nil {
		log.Printf("[WARN] Error getting zone during Firewall Update %v\n", err)
		return err
	}

	return resourceIBMCISFirewallRecordRead(d, meta)
}

func resourceIBMCISFirewallRecordDelete(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	firewallType, recordID, zoneID, cisID, _ := convertTfToCisFourVar(d.Id())
	if err != nil {
		return err
	}
	err = cisClient.Firewall().DeleteFirewall(cisID, zoneID, firewallType, recordID)
	if err != nil && !strings.Contains(err.Error(), "Request failed with status code: 404") {
		return fmt.Errorf("Error deleting IBMCISFirewall: %s", err)
	}
	return nil
}
func resourceIBMCISFirewallRecordExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return false, err
	}
	firewallType, recordID, zoneID, cisID, _ := convertTfToCisFourVar(d.Id())
	if err != nil {
		return false, err
	}
	_, err = cisClient.Firewall().GetFirewall(cisID, zoneID, firewallType, recordID)
	if err != nil {
		if strings.Contains(err.Error(), "Request failed with status code: 404") {
			return false, nil
		}
		return false, fmt.Errorf("Error getting IBMCISFirewall: %s", err)
	}

	return true, nil
}
