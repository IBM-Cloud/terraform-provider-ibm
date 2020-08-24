package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	cisDNSRecords = "cis_dns_records"
)

func dataSourceIBMCISDNSRecord() *schema.Resource {
	return &schema.Resource{
		Read:     dataSourceIBMCISDNSRecordsRead,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "DNS Zone CRN",
			},
			cisDomainID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Zone Id",
			},

			cisDNSRecords: {
				Type:        schema.TypeList,
				Description: "Collection of dns resource records",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "DNS record id",
						},
						cisDNSRecordID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "DNS record id",
						},
						cisZoneName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "DNS Record Name",
						},
						cisDNSRecordName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "DNS Record Name",
						},
						cisDNSRecordCreatedOn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "DNS Record created on",
						},
						cisDNSRecordModifiedOn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "DNS Record modified on",
						},
						cisDNSRecordType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "DNS Record Type",
						},
						cisDNSRecordContent: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "DNS Record conent info",
						},
						cisDNSRecordPriority: {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "DNS Record MX priority",
						},
						cisDNSRecordProxiable: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "DNS Record proxiable",
						},
						cisDNSRecordProxied: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "DNS Record proxied",
						},
						cisDNSRecordTTL: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "DNS Record Time To Live",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMCISDNSRecordsRead(d *schema.ResourceData, meta interface{}) error {
	var (
		crn     string
		zoneID  string
		records []map[string]interface{}
	)
	sess, err := meta.(ClientSession).CisDNSRecordClientSession()
	if err != nil {
		return err
	}

	// session options
	crn = d.Get(cisID).(string)
	zoneID = d.Get(cisDomainID).(string)
	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	opt := sess.NewListAllDnsRecordsOptions()
	result, response, err := sess.ListAllDnsRecords(opt)
	if err != nil {
		log.Printf("Error reading dns records: %s", response)
		return err
	}

	records = make([]map[string]interface{}, 0)
	for _, instance := range result.Result {
		record := map[string]interface{}{}
		record["id"] = convertCisToTfThreeVar(*instance.ID, zoneID, crn)
		record[cisDNSRecordID] = *instance.ID
		record[cisZoneName] = *instance.ZoneName
		record[cisDNSRecordCreatedOn] = *instance.CreatedOn
		record[cisDNSRecordModifiedOn] = *instance.ModifiedOn
		record[cisDNSRecordName] = *instance.Name
		record[cisDNSRecordType] = *instance.Type
		record[cisDNSRecordContent] = *instance.Content
		record[cisDNSRecordProxiable] = *instance.Proxiable
		record[cisDNSRecordProxied] = *instance.Proxied
		record[cisDNSRecordTTL] = *instance.TTL

		switch *instance.Type {
		// for MX & SRV records ouptut
		case cisDNSRecordTypeMX, cisDNSRecordTypeSRV:
			record[cisDNSRecordPriority] = *instance.Priority
		}
		records = append(records, record)
	}
	d.SetId(dataSourceIBMCISDNSRecordID(d))
	d.Set(cisDNSRecords, records)
	return nil
}

// dataSourceIBMCISDNSRecordID returns a reasonable ID for dns zones list.
func dataSourceIBMCISDNSRecordID(d *schema.ResourceData) string {
	zoneID := d.Get(cisDomainID)
	crn := d.Get(cisID)
	return fmt.Sprintf("%s:%s", zoneID, crn)
}
