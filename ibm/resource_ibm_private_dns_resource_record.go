package ibm

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var allowedPrivateDomainRecordTypes = []string{
	"A", "AAAA", "CNAME", "MX", "PTR", "SRV", "TXT",
}

const (
	pdnsResourceRecordID = "resource_record_id"
	pdnsRecordType       = "type"
	pdnsRecordTTL        = "ttl"
	pdnsRecordName       = "name"
	pdnsRdata            = "rdata"
	pdnsMxPreference     = "preference"
	pdnsSrvPort          = "port"
	pdnsSrvPriority      = "priority"
	pdnsSrvWeight        = "weight"
	pdnsSrvProtocol      = "protocol"
	pdnsSrvService       = "service"
	pdnsRecordCreatedOn  = "created_on"
	pdnsRecordModifiedOn = "modified_on"
)

func caseDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
	return strings.ToUpper(old) == strings.ToUpper(new)
}

func resourceIBMPrivateDNSResourceRecord() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMPrivateDNSResourceRecordCreate,
		Read:     resourceIBMPrivateDNSResourceRecordRead,
		Update:   resourceIBMPrivateDNSResourceRecordUpdate,
		Delete:   resourceIBMPrivateDNSResourceRecordDelete,
		Exists:   resourceIBMPrivateDNSResourceRecordExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			pdnsResourceRecordID: {
				Type:     schema.TypeString,
				Computed: true,
			},

			pdnsInstanceID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			pdnsZoneID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			pdnsRecordName: {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: caseDiffSuppress,
			},

			pdnsRecordType: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: func(val interface{}, field string) (warnings []string, errors []error) {
					value := val.(string)
					for _, rtype := range allowedPrivateDomainRecordTypes {
						if value == rtype {
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
			},

			pdnsRdata: {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(val interface{}, field string) (warnings []string, errors []error) {
					value := val.(string)
					if ipv6Regexp.MatchString(value) && upcaseRegexp.MatchString(value) {
						errors = append(
							errors,
							fmt.Errorf(
								"IPv6 addresses in the data property cannot have upper case letters: %s",
								value,
							),
						)
					}
					return
				},
			},

			pdnsRecordTTL: {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  900,
				DefaultFunc: func() (interface{}, error) {
					return 900, nil
				},
			},

			pdnsMxPreference: {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},

			pdnsSrvPort: {
				Type:     schema.TypeInt,
				Optional: true,
			},

			pdnsSrvPriority: {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},

			pdnsSrvWeight: {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},

			pdnsSrvService: {
				Type:     schema.TypeString,
				Optional: true,
			},

			pdnsSrvProtocol: {
				Type:     schema.TypeString,
				Optional: true,
			},

			pdnsRecordCreatedOn: {
				Type:     schema.TypeString,
				Computed: true,
			},

			pdnsRecordModifiedOn: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMPrivateDNSResourceRecordCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	var (
		instanceID string
		zoneID     string
		recordType string
		name       string
		rdata      string
		service    string
		protocol   string
		ttl        int
		preference int
		port       int
		priority   int
		weight     int
	)

	instanceID = d.Get(pdnsInstanceID).(string)
	zoneID = d.Get(pdnsZoneID).(string)
	recordType = d.Get(pdnsRecordType).(string)
	name = d.Get(pdnsRecordName).(string)
	rdata = d.Get(pdnsRdata).(string)

	if v, ok := d.GetOk(pdnsRecordTTL); ok {
		ttl = v.(int)
	}

	createResourceRecordOptions := sess.NewCreateResourceRecordOptions(instanceID, zoneID)
	createResourceRecordOptions.SetName(name)
	createResourceRecordOptions.SetType(recordType)
	createResourceRecordOptions.SetTTL(int64(ttl))

	switch recordType {
	case "A":
		resourceRecordAData, err := sess.NewResourceRecordInputRdataRdataARecord(rdata)
		if err != nil {
			return err
		}
		createResourceRecordOptions.SetRdata(resourceRecordAData)
	case "AAAA":
		resourceRecordAaaaData, err := sess.NewResourceRecordInputRdataRdataAaaaRecord(rdata)
		if err != nil {
			return err
		}
		createResourceRecordOptions.SetRdata(resourceRecordAaaaData)
	case "CNAME":
		resourceRecordCnameData, err := sess.NewResourceRecordInputRdataRdataCnameRecord(rdata)
		if err != nil {
			return err
		}
		createResourceRecordOptions.SetRdata(resourceRecordCnameData)
	case "PTR":
		resourceRecordPtrData, err := sess.NewResourceRecordInputRdataRdataPtrRecord(rdata)
		if err != nil {
			return err
		}
		createResourceRecordOptions.SetRdata(resourceRecordPtrData)
	case "TXT":
		resourceRecordTxtData, err := sess.NewResourceRecordInputRdataRdataTxtRecord(rdata)
		if err != nil {
			return err
		}
		createResourceRecordOptions.SetRdata(resourceRecordTxtData)
	case "MX":
		if v, ok := d.GetOk(pdnsMxPreference); ok {
			preference = v.(int)
		}
		resourceRecordMxData, err := sess.NewResourceRecordInputRdataRdataMxRecord(rdata, int64(preference))
		if err != nil {
			return err
		}
		createResourceRecordOptions.SetRdata(resourceRecordMxData)
	case "SRV":
		if v, ok := d.GetOk(pdnsSrvPort); ok {
			port = v.(int)
		}
		if v, ok := d.GetOk(pdnsSrvPriority); ok {
			priority = v.(int)
		}
		if v, ok := d.GetOk(pdnsSrvWeight); ok {
			weight = v.(int)
		}
		resourceRecordSrvData, err := sess.NewResourceRecordInputRdataRdataSrvRecord(int64(port), int64(priority), rdata, int64(weight))
		if err != nil {
			return err
		}
		if v, ok := d.GetOk(pdnsSrvService); ok {
			service = v.(string)
		}
		if v, ok := d.GetOk(pdnsSrvProtocol); ok {
			protocol = v.(string)
		}
		createResourceRecordOptions.SetRdata(resourceRecordSrvData)
		createResourceRecordOptions.SetService(service)
		createResourceRecordOptions.SetProtocol(protocol)
	}

	response, _, err := sess.CreateResourceRecord(createResourceRecordOptions)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", instanceID, zoneID, *response.ID))
	d.Set(pdnsResourceRecordID, *response.ID)

	return resourceIBMPrivateDNSResourceRecordRead(d, meta)
}

func resourceIBMPrivateDNSResourceRecordRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")
	getResourceRecordOptions := sess.NewGetResourceRecordOptions(id_set[0], id_set[1], id_set[2])
	response, _, err := sess.GetResourceRecord(getResourceRecordOptions)
	if err != nil {
		return err
	}

	// extract the record name by removing zone details
	var recordName string
	zone := strings.Split(id_set[1], ":")
	name := strings.Split(*response.Name, zone[0])
	name[0] = strings.Trim(name[0], ".")
	recordName = name[0]

	if *response.Type == "SRV" {
		// "_sip._udp.testsrv"
		temp := strings.Split(name[0], ".")
		recordName = temp[2]
	}

	d.Set("id", response.ID)
	d.Set(pdnsResourceRecordID, response.ID)
	d.Set(pdnsInstanceID, id_set[0])
	d.Set(pdnsZoneID, id_set[1])
	d.Set(pdnsRecordName, recordName)
	d.Set(pdnsRdata, response.Rdata)
	d.Set(pdnsRecordType, response.Type)
	d.Set(pdnsRecordTTL, response.TTL)
	d.Set(pdnsRecordCreatedOn, response.CreatedOn)
	d.Set(pdnsRecordModifiedOn, response.ModifiedOn)

	if *response.Type == "SRV" {
		data := response.Rdata.(map[string]interface{})
		d.Set(pdnsSrvPort, data["port"])
		d.Set(pdnsSrvPriority, data["priority"])
		d.Set(pdnsSrvWeight, data["weight"])
		d.Set(pdnsRdata, data["target"])
		d.Set(pdnsSrvService, response.Service)
		d.Set(pdnsSrvProtocol, response.Protocol)
	}

	if *response.Type == "MX" {
		data := response.Rdata.(map[string]interface{})
		d.Set(pdnsMxPreference, data["preference"])
		d.Set(pdnsRdata, data["exchange"])
	}

	return nil
}

func resourceIBMPrivateDNSResourceRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")
	getResourceRecordOptions := sess.NewGetResourceRecordOptions(id_set[0], id_set[1], id_set[2])
	response, _, err := sess.GetResourceRecord(getResourceRecordOptions)
	if err != nil {
		return err
	}

	updateResourceRecordOptions := sess.NewUpdateResourceRecordOptions(id_set[0], id_set[1], id_set[2])
	updateResourceRecordOptions.SetName(*response.Type)

	//
	var ttl int64
	var rdata string

	temp := d.Get(pdnsRecordTTL).(int)
	ttl = int64(temp)

	recordType := *response.Type
	switch recordType {
	case "A":
		if d.HasChange(pdnsRecordTTL) || d.HasChange(pdnsRdata) {
			updateResourceRecordOptions.SetTTL(ttl)
			rdata = d.Get(pdnsRdata).(string)
			resourceRecordAData, err := sess.NewResourceRecordUpdateInputRdataRdataARecord(rdata)
			if err != nil {
				return err
			}
			updateResourceRecordOptions.SetRdata(resourceRecordAData)
		}
	case "AAAA":
		if d.HasChange(pdnsRecordTTL) || d.HasChange(pdnsRdata) {
			updateResourceRecordOptions.SetTTL(ttl)
			rdata = d.Get(pdnsRdata).(string)
			resourceRecordAaaaData, err := sess.NewResourceRecordUpdateInputRdataRdataAaaaRecord(rdata)
			if err != nil {
				return err
			}
			updateResourceRecordOptions.SetRdata(resourceRecordAaaaData)
		}
	case "CNAME":
		if d.HasChange(pdnsRecordTTL) || d.HasChange(pdnsRdata) {
			updateResourceRecordOptions.SetTTL(ttl)
			rdata = d.Get(pdnsRdata).(string)
			resourceRecordCnameData, err := sess.NewResourceRecordUpdateInputRdataRdataCnameRecord(rdata)
			if err != nil {
				return err
			}
			updateResourceRecordOptions.SetRdata(resourceRecordCnameData)
		}
	case "PTR":
		if d.HasChange(pdnsRecordTTL) || d.HasChange(pdnsRdata) {
			updateResourceRecordOptions.SetTTL(ttl)
			rdata = d.Get(pdnsRdata).(string)
			resourceRecordPtrData, err := sess.NewResourceRecordUpdateInputRdataRdataPtrRecord(rdata)
			if err != nil {
				return err
			}
			updateResourceRecordOptions.SetRdata(resourceRecordPtrData)
		}
	case "TXT":
		if d.HasChange(pdnsRecordTTL) || d.HasChange(pdnsRdata) {
			updateResourceRecordOptions.SetTTL(ttl)
			rdata = d.Get(pdnsRdata).(string)
			resourceRecordTxtData, err := sess.NewResourceRecordUpdateInputRdataRdataTxtRecord(rdata)
			if err != nil {
				return err
			}
			updateResourceRecordOptions.SetRdata(resourceRecordTxtData)
		}
	case "MX":
		if d.HasChange(pdnsRecordTTL) || d.HasChange(pdnsRdata) ||
			d.HasChange(pdnsMxPreference) {

			updateResourceRecordOptions.SetTTL(ttl)
			rdata = d.Get(pdnsRdata).(string)
			preference := d.Get(pdnsMxPreference).(int64)

			resourceRecordMxData, err := sess.NewResourceRecordUpdateInputRdataRdataMxRecord(rdata, preference)
			if err != nil {
				return err
			}
			updateResourceRecordOptions.SetRdata(resourceRecordMxData)
		}
	case "SRV":
		if d.HasChange(pdnsRecordTTL) || d.HasChange(pdnsRdata) ||
			d.HasChange(pdnsSrvPort) || d.HasChange(pdnsSrvPriority) ||
			d.HasChange(pdnsSrvWeight) || d.HasChange(pdnsSrvService) ||
			d.HasChange(pdnsSrvProtocol) {

			updateResourceRecordOptions.SetTTL(ttl)
			rdata = d.Get(pdnsRdata).(string)
			port := d.Get(pdnsSrvPort).(int64)
			priority := d.Get(pdnsSrvPriority).(int64)
			weight := d.Get(pdnsSrvWeight).(int64)

			resourceRecordSrvData, err := sess.NewResourceRecordUpdateInputRdataRdataSrvRecord(port, priority, rdata, weight)
			if err != nil {
				return err
			}
			updateResourceRecordOptions.SetRdata(resourceRecordSrvData)

			service := d.Get(pdnsSrvService).(string)
			protocol := d.Get(pdnsSrvProtocol).(string)
			updateResourceRecordOptions.SetService(service)
			updateResourceRecordOptions.SetProtocol(protocol)
		}
	}

	//
	_, _, err = sess.UpdateResourceRecord(updateResourceRecordOptions)
	if err != nil {
		return err
	}

	return resourceIBMPrivateDNSResourceRecordRead(d, meta)
}

func resourceIBMPrivateDNSResourceRecordDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")
	deleteResourceRecordOptions := sess.NewDeleteResourceRecordOptions(id_set[0], id_set[1], id_set[2])
	response, err := sess.DeleteResourceRecord(deleteResourceRecordOptions)
	if err != nil && response.StatusCode != 404 {
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMPrivateDNSResourceRecordExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return false, err
	}

	id_set := strings.Split(d.Id(), "/")
	getResourceRecordOptions := sess.NewGetResourceRecordOptions(id_set[0], id_set[1], id_set[2])
	_, response, err := sess.GetResourceRecord(getResourceRecordOptions)

	if err != nil && response.StatusCode != 404 {
		if response.StatusCode == 404 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
