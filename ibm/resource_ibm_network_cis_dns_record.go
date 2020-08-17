package ibm

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	cisCRN                 = "crn"
	cisZoneID              = "zone_id"
	cisZoneName            = "zone_name"
	cisDNSRecordCreatedOn  = "created_on"
	cisDNSRecordModifiedOn = "modified_on"
	cisDNSRecordName       = "name"
	cisDNSRecordType       = "type"
	cisDNSRecordContent    = "content"
	cisDNSRecordZoneID     = "zone_id"
	cisDNSRecordZoneName   = "zone_name"
	cisDNSRecordProxiable  = "proxiable"
	cisDNSRecordProxied    = "proxied"
	cisDNSRecordTTL        = "ttl"
	cisDNSRecordPriority   = "priority"
	cisDNSRecordData       = "data"
)

// Constants associated with the DNS Record Type property.
// dns record type.
const (
	cisDNSRecordTypeA     = "A"
	cisDNSRecordTypeAAAA  = "AAAA"
	cisDNSRecordTypeCAA   = "CAA"
	cisDNSRecordTypeCNAME = "CNAME"
	cisDNSRecordTypeLOC   = "LOC"
	cisDNSRecordTypeMX    = "MX"
	cisDNSRecordTypeNS    = "NS"
	cisDNSRecordTypeSPF   = "SPF"
	cisDNSRecordTypeSRV   = "SRV"
	cisDNSRecordTypeTXT   = "TXT"
)

var allowedDNSRecordTypes = []string{
	cisDNSRecordTypeA,
	cisDNSRecordTypeAAAA,
	cisDNSRecordTypeCAA,
	cisDNSRecordTypeCNAME,
	cisDNSRecordTypeLOC,
	cisDNSRecordTypeMX,
	cisDNSRecordTypeNS,
	cisDNSRecordTypeSPF,
	cisDNSRecordTypeSRV,
	cisDNSRecordTypeTXT,
}

func resourceIBMNetworkCISDNSRecords() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMNetworkCISDNSRecordCreate,
		Update:   resourceIBMNetworkCISDNSRecordUpdate,
		Read:     resourceIBMNetworkCISDNSRecordRead,
		Delete:   resourceIBMNetworkCISDNSRecordDelete,
		Exists:   resourceIBMNetworkCISDNSRecordExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			cisCRN: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "CRN",
			},
			cisZoneID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Zone ID",
			},
			cisZoneName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DNS Record Zone Name",
			},
			cisDNSRecordName: {
				Type:     schema.TypeString,
				Optional: true,
				// ForceNew:    true,
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
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: func(val interface{}, field string) (warnings []string, errors []error) {
					value := val.(string)
					for _, DNSRecordType := range allowedDNSRecordTypes {
						if value == DNSRecordType {
							return
						}
					}

					errors = append(
						errors,
						fmt.Errorf("%s is not one of the valid domain record types: %s",
							value, strings.Join(allowedPrivateDomainRecordTypes, ", "),
						),
					)
					return
				},
				Description: "DNS Record Type",
			},
			cisDNSRecordContent: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "DNS Record conent info",
			},
			cisDNSRecordData: {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "DNS Record data",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// for LOC record
						"altitude": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"lat_degrees": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"lat_direction": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"lat_minutes": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"lat_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"long_degrees": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"long_direction": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"long_minutes": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"long_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"precision_horz": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"precision_vert": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"size": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						// for CAA record
						"tag": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						// for SRV record
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"proto": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"service": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"target": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"weight": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			cisDNSRecordPriority: {
				Type:        schema.TypeInt,
				Optional:    true,
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
				Optional:    true,
				Description: "DNS Record Time To Live",
			},
		},
	}
}

func resourceIBMNetworkCISDNSRecordCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).CisDNSRecordClientSession()
	if err != nil {
		log.Printf("Error: %s", err)
		return err
	}
	var (
		crn            string
		zoneID         string
		recordName     string
		recordType     string
		recordContent  string
		recordPriority int
		ttl            int
		ok             bool
		data           interface{}
		recordData     map[string]interface{}
	)
	// session options
	crn = d.Get(cisCRN).(string)
	zoneID = d.Get(cisZoneID).(string)
	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	// Input options
	opt := sess.NewCreateDnsRecordOptions()

	// set record type
	recordType = d.Get(cisDNSRecordType).(string)
	opt.SetType(recordType)
	// set ttl value
	ttl = d.Get(cisDNSRecordTTL).(int)
	opt.SetTTL(int64(ttl))

	switch recordType {
	// A, AAAA, CNAME, SPF, TXT & NS records inputs
	case cisDNSRecordTypeA,
		cisDNSRecordTypeAAAA,
		cisDNSRecordTypeCNAME,
		cisDNSRecordTypeSPF,
		cisDNSRecordTypeTXT,
		cisDNSRecordTypeNS:
		// set record name & content
		recordName = d.Get(cisDNSRecordName).(string)
		opt.SetName(recordName)
		recordContent = d.Get(cisDNSRecordContent).(string)
		opt.SetContent(recordContent)

	// MX Record inputs
	case cisDNSRecordTypeMX:

		// set record name
		recordName = d.Get(cisDNSRecordName).(string)
		opt.SetName(recordName)
		recordContent = d.Get(cisDNSRecordContent).(string)
		opt.SetContent(recordContent)
		recordPriority = d.Get(cisDNSRecordPriority).(int)
		opt.SetPriority(int64(recordPriority))

	// LOC Record inputs
	case cisDNSRecordTypeLOC:

		// set record name
		recordName = d.Get(cisDNSRecordName).(string)
		opt.SetName(recordName)
		data, ok = d.GetOk(cisDNSRecordData)
		if ok == false {
			log.Printf("Error in getting data")
			return err
		}
		recordData = make(map[string]interface{}, 0)
		var dataMap map[string]interface{} = data.(map[string]interface{})

		// altitude
		v, ok := strconv.Atoi(dataMap["altitude"].(string))
		if ok != nil {
			return ok
		}
		recordData["altitude"] = v

		// lat_degrees
		v, ok = strconv.Atoi(dataMap["lat_degrees"].(string))
		if ok != nil {
			return ok
		}
		recordData["lat_degrees"] = v

		// lat_direction
		recordData["lat_direction"] = dataMap["lat_direction"].(string)

		// long_direction
		recordData["long_direction"] = dataMap["long_direction"].(string)

		// lat_minutes
		v, ok = strconv.Atoi(dataMap["lat_minutes"].(string))
		if ok != nil {
			return ok
		}
		recordData["lat_minutes"] = v

		// lat_seconds
		v, ok = strconv.Atoi(dataMap["lat_seconds"].(string))
		if ok != nil {
			return ok
		}
		recordData["lat_seconds"] = v

		// long_degrees
		v, ok = strconv.Atoi(dataMap["long_degrees"].(string))
		if ok != nil {
			return ok
		}
		recordData["long_degrees"] = v

		// long_minutes
		v, ok = strconv.Atoi(dataMap["long_minutes"].(string))
		if ok != nil {
			return ok
		}
		recordData["long_minutes"] = v

		// long_seconds
		v, ok = strconv.Atoi(dataMap["long_seconds"].(string))
		if ok != nil {
			return ok
		}
		recordData["long_seconds"] = v

		// percision_horz
		v, ok = strconv.Atoi(dataMap["precision_horz"].(string))
		if ok != nil {
			return ok
		}
		recordData["precision_horz"] = v

		// precision_vert
		v, ok = strconv.Atoi(dataMap["precision_vert"].(string))
		if ok != nil {
			return ok
		}
		recordData["precision_vert"] = v

		// size
		v, ok = strconv.Atoi(dataMap["size"].(string))
		if ok != nil {
			return ok
		}
		recordData["size"] = v

		opt.SetData(recordData)

	// CAA Record inputs
	case cisDNSRecordTypeCAA:

		// set record name
		recordName = d.Get(cisDNSRecordName).(string)
		opt.SetName(recordName)
		data, ok = d.GetOk(cisDNSRecordData)
		if ok == false {
			log.Printf("Error in getting data")
			return err
		}
		recordData = make(map[string]interface{}, 0)
		var dataMap map[string]interface{} = data.(map[string]interface{})

		// tag
		v := dataMap["tag"].(string)
		recordData["tag"] = v

		// value
		v = dataMap["value"].(string)
		recordData["value"] = v

		opt.SetData(recordData)

	// SRV record input
	case cisDNSRecordTypeSRV:
		data, ok = d.GetOk(cisDNSRecordData)
		if ok == false {
			log.Printf("Error in getting data")
			return err
		}
		recordData = make(map[string]interface{}, 0)
		var dataMap map[string]interface{} = data.(map[string]interface{})

		// name
		v := dataMap["name"].(string)
		recordData["name"] = v

		// target
		v = dataMap["target"].(string)
		recordData["target"] = v

		// proto
		v = dataMap["proto"].(string)
		recordData["proto"] = v

		// service
		v = dataMap["service"].(string)
		recordData["service"] = v
		opt.SetData(recordData)

		// port
		s, ok := strconv.Atoi(dataMap["port"].(string))
		if ok != nil {
			return ok
		}
		recordData["port"] = s

		// priority
		s, ok = strconv.Atoi(dataMap["priority"].(string))
		if ok != nil {
			return ok
		}
		recordData["priority"] = s

		// weight
		s, ok = strconv.Atoi(dataMap["weight"].(string))
		if ok != nil {
			return ok
		}
		recordData["weight"] = s
		opt.SetData(recordData)
	}

	result, response, err := sess.CreateDnsRecord(opt)
	if err != nil {
		log.Printf("Error creating dns record: %s, error %s", response, err)
		return err
	}
	d.SetId(*result.Result.ID)
	d.Set(cisZoneName, *result.Result.ZoneName)
	d.Set(cisDNSRecordCreatedOn, *result.Result.CreatedOn)
	d.Set(cisDNSRecordModifiedOn, *result.Result.ModifiedOn)
	d.Set(cisDNSRecordName, *result.Result.Name)
	d.Set(cisDNSRecordType, *result.Result.Type)
	d.Set(cisDNSRecordContent, *result.Result.Content)
	d.Set(cisDNSRecordZoneID, *result.Result.ZoneID)
	d.Set(cisDNSRecordZoneName, *result.Result.ZoneName)
	d.Set(cisDNSRecordProxiable, *result.Result.Proxiable)
	d.Set(cisDNSRecordProxied, *result.Result.Proxied)
	d.Set(cisDNSRecordTTL, *result.Result.TTL)

	switch recordType {
	// for MX & SRV records ouptut
	case cisDNSRecordTypeMX, cisDNSRecordTypeSRV:
		d.Set(cisDNSRecordPriority, *result.Result.Priority)
	// for LOC & CAA records output
	case cisDNSRecordTypeLOC, cisDNSRecordTypeCAA:
		d.Set(cisDNSRecordData, result.Result.Data)
	}

	return err
}

func resourceIBMNetworkCISDNSRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).CisDNSRecordClientSession()
	if err != nil {
		log.Printf("Error: %s", err)
		return err
	}
	var (
		ID             string
		crn            string
		zoneID         string
		recordName     string
		recordType     string
		recordContent  string
		recordPriority int
		ttl            int
		ok             bool
		data           interface{}
		recordData     map[string]interface{}
	)
	// session options
	ID = d.Id()
	crn = d.Get(cisCRN).(string)
	zoneID = d.Get(cisZoneID).(string)
	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	// Input options
	opt := sess.NewUpdateDnsRecordOptions(ID)

	// set record type
	if d.HasChange(cisDNSRecordType) {
		recordType = d.Get(cisDNSRecordType).(string)
		opt.SetType(recordType)
	}
	// set ttl value
	if d.HasChange(cisDNSRecordTTL) {
		ttl = d.Get(cisDNSRecordTTL).(int)
		opt.SetTTL(int64(ttl))
	}

	switch recordType {
	// A, AAAA, CNAME, SPF, TXT & NS records inputs
	case cisDNSRecordTypeA,
		cisDNSRecordTypeAAAA,
		cisDNSRecordTypeCNAME,
		cisDNSRecordTypeSPF,
		cisDNSRecordTypeTXT,
		cisDNSRecordTypeNS:
		// set record name & content
		if d.HasChange(cisDNSRecordName) {
			recordName = d.Get(cisDNSRecordName).(string)
			opt.SetName(recordName)
		}
		if d.HasChange(cisDNSRecordContent) {
			recordContent = d.Get(cisDNSRecordContent).(string)
			opt.SetContent(recordContent)
		}

	// MX Record inputs
	case cisDNSRecordTypeMX:

		// set record name
		if d.HasChange(cisDNSRecordName) {
			recordName = d.Get(cisDNSRecordName).(string)
			opt.SetName(recordName)
		}

		// set content
		if d.HasChange(cisDNSRecordContent) {
			recordContent = d.Get(cisDNSRecordContent).(string)
			opt.SetContent(recordContent)
		}

		// set priority
		if d.HasChange(cisDNSRecordPriority) {
			recordPriority = d.Get(cisDNSRecordPriority).(int)
			opt.SetPriority(int64(recordPriority))
		}

	// LOC Record inputs
	case cisDNSRecordTypeLOC:

		// set record name
		if d.HasChange(cisDNSRecordName) {
			recordName = d.Get(cisDNSRecordName).(string)
			opt.SetName(recordName)
		}

		if d.HasChange(cisDNSRecordData) {
			data, ok = d.GetOk(cisDNSRecordData)
			if ok == false {
				log.Printf("Error in getting data")
				return err
			}
			recordData = make(map[string]interface{}, 0)
			var dataMap map[string]interface{} = data.(map[string]interface{})

			// altitude
			v, ok := strconv.Atoi(dataMap["altitude"].(string))
			if ok != nil {
				return ok
			}
			recordData["altitude"] = v

			// lat_degrees
			v, ok = strconv.Atoi(dataMap["lat_degrees"].(string))
			if ok != nil {
				return ok
			}
			recordData["lat_degrees"] = v

			// lat_direction
			recordData["lat_direction"] = dataMap["lat_direction"].(string)

			// long_direction
			recordData["long_direction"] = dataMap["long_direction"].(string)

			// lat_minutes
			v, ok = strconv.Atoi(dataMap["lat_minutes"].(string))
			if ok != nil {
				return ok
			}
			recordData["lat_minutes"] = v

			// lat_seconds
			v, ok = strconv.Atoi(dataMap["lat_seconds"].(string))
			if ok != nil {
				return ok
			}
			recordData["lat_seconds"] = v

			// long_degrees
			v, ok = strconv.Atoi(dataMap["long_degrees"].(string))
			if ok != nil {
				return ok
			}
			recordData["long_degrees"] = v

			// long_minutes
			v, ok = strconv.Atoi(dataMap["long_minutes"].(string))
			if ok != nil {
				return ok
			}
			recordData["long_minutes"] = v

			// long_seconds
			v, ok = strconv.Atoi(dataMap["long_seconds"].(string))
			if ok != nil {
				return ok
			}
			recordData["long_seconds"] = v

			// percision_horz
			v, ok = strconv.Atoi(dataMap["precision_horz"].(string))
			if ok != nil {
				return ok
			}
			recordData["precision_horz"] = v

			// precision_vert
			v, ok = strconv.Atoi(dataMap["precision_vert"].(string))
			if ok != nil {
				return ok
			}
			recordData["precision_vert"] = v

			// size
			v, ok = strconv.Atoi(dataMap["size"].(string))
			if ok != nil {
				return ok
			}
			recordData["size"] = v

			opt.SetData(recordData)
		}

	// CAA Record inputs
	case cisDNSRecordTypeCAA:

		// set record name
		if d.HasChange(cisDNSRecordName) {
			recordName = d.Get(cisDNSRecordName).(string)
			opt.SetName(recordName)
		}
		if d.HasChange(cisDNSRecordData) {
			data, ok = d.GetOk(cisDNSRecordData)
			if ok == false {
				log.Printf("Error in getting data")
				return err
			}
			recordData = make(map[string]interface{}, 0)
			var dataMap map[string]interface{} = data.(map[string]interface{})

			// tag
			v := dataMap["tag"].(string)
			recordData["tag"] = v

			// value
			v = dataMap["value"].(string)
			recordData["value"] = v

			opt.SetData(recordData)
		}

	// SRV record input
	case cisDNSRecordTypeSRV:
		if d.HasChange(cisDNSRecordData) {
			data, ok = d.GetOk(cisDNSRecordData)
			if ok == false {
				log.Printf("Error in getting data")
				return err
			}
			recordData = make(map[string]interface{}, 0)
			var dataMap map[string]interface{} = data.(map[string]interface{})

			// name
			v := dataMap["name"].(string)
			recordData["name"] = v

			// target
			v = dataMap["target"].(string)
			recordData["target"] = v

			// proto
			v = dataMap["proto"].(string)
			recordData["proto"] = v

			// service
			v = dataMap["service"].(string)
			recordData["service"] = v
			opt.SetData(recordData)

			// port
			s, ok := strconv.Atoi(dataMap["port"].(string))
			if ok != nil {
				return ok
			}
			recordData["port"] = s

			// priority
			s, ok = strconv.Atoi(dataMap["priority"].(string))
			if ok != nil {
				return ok
			}
			recordData["priority"] = s

			// weight
			s, ok = strconv.Atoi(dataMap["weight"].(string))
			if ok != nil {
				return ok
			}
			recordData["weight"] = s
			opt.SetData(recordData)
		}
	}

	result, response, err := sess.UpdateDnsRecord(opt)
	if err != nil {
		log.Printf("Error creating dns record: %s, error %s", response, err)
		return err
	}
	d.SetId(*result.Result.ID)
	d.Set(cisZoneName, *result.Result.ZoneName)
	d.Set(cisDNSRecordCreatedOn, *result.Result.CreatedOn)
	d.Set(cisDNSRecordModifiedOn, *result.Result.ModifiedOn)
	d.Set(cisDNSRecordName, *result.Result.Name)
	d.Set(cisDNSRecordType, *result.Result.Type)
	d.Set(cisDNSRecordContent, *result.Result.Content)
	d.Set(cisDNSRecordZoneID, *result.Result.ZoneID)
	d.Set(cisDNSRecordZoneName, *result.Result.ZoneName)
	d.Set(cisDNSRecordProxiable, *result.Result.Proxiable)
	d.Set(cisDNSRecordProxied, *result.Result.Proxied)
	d.Set(cisDNSRecordTTL, *result.Result.TTL)

	switch recordType {
	// for MX & SRV records ouptut
	case cisDNSRecordTypeMX, cisDNSRecordTypeSRV:
		d.Set(cisDNSRecordPriority, *result.Result.Priority)
	// for LOC & CAA records output
	case cisDNSRecordTypeLOC, cisDNSRecordTypeCAA:
		d.Set(cisDNSRecordData, result.Result.Data)
	}

	return err
}

func resourceIBMNetworkCISDNSRecordRead(d *schema.ResourceData, meta interface{}) error {
	var (
		crn      string
		zoneID   string
		recordID string
	)
	sess, err := meta.(ClientSession).CisDNSRecordClientSession()
	if err != nil {
		return err
	}
	// session options
	crn = d.Get(cisCRN).(string)
	zoneID = d.Get(cisZoneID).(string)
	recordID = d.Id()

	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	opt := sess.NewGetDnsRecordOptions(recordID)
	result, response, err := sess.GetDnsRecord(opt)
	if err != nil {
		log.Printf("Error reading dns record: %s", response)
		return err
	}

	d.SetId(*result.Result.ID)
	d.Set(cisZoneName, *result.Result.ZoneName)
	d.Set(cisDNSRecordCreatedOn, *result.Result.CreatedOn)
	d.Set(cisDNSRecordModifiedOn, *result.Result.ModifiedOn)
	d.Set(cisDNSRecordName, *result.Result.Name)
	d.Set(cisDNSRecordType, *result.Result.Type)
	d.Set(cisDNSRecordContent, *result.Result.Content)
	d.Set(cisDNSRecordZoneID, *result.Result.ZoneID)
	d.Set(cisDNSRecordZoneName, *result.Result.ZoneName)
	d.Set(cisDNSRecordProxiable, *result.Result.Proxiable)
	d.Set(cisDNSRecordProxied, *result.Result.Proxied)
	d.Set(cisDNSRecordTTL, *result.Result.TTL)

	switch *result.Result.Type {
	// for MX & SRV records ouptut
	case cisDNSRecordTypeMX, cisDNSRecordTypeSRV:
		d.Set(cisDNSRecordPriority, *result.Result.Priority)
	// for LOC & CAA records output
	case cisDNSRecordTypeLOC, cisDNSRecordTypeCAA:
		d.Set(cisDNSRecordData, result.Result.Data)
	}
	return err
}

func resourceIBMNetworkCISDNSRecordDelete(d *schema.ResourceData, meta interface{}) error {
	var (
		crn      string
		zoneID   string
		recordID string
	)
	sess, err := meta.(ClientSession).CisDNSRecordClientSession()
	if err != nil {
		return err
	}

	// session options
	crn = d.Get(cisCRN).(string)
	zoneID = d.Get(cisZoneID).(string)
	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)
	recordID = d.Id()

	opt := sess.NewDeleteDnsRecordOptions(recordID)
	result, response, err := sess.DeleteDnsRecord(opt)
	if err != nil {
		log.Printf("Error deleting dns record: %s", response)
		return err
	}
	log.Printf("record id: %s", *result.Result.ID)
	return err
}

func resourceIBMNetworkCISDNSRecordExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).CisDNSRecordClientSession()
	if err != nil {
		return false, err
	}

	// session options
	crn := d.Get(cisCRN).(string)
	zoneID := d.Get(cisZoneID).(string)
	recordID := d.Id()

	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	opt := sess.NewGetDnsRecordOptions(recordID)
	_, response, err := sess.GetDnsRecord(opt)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
