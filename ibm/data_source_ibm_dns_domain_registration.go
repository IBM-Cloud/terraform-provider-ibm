package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/services"
	"log"
	//"reflect"
)

// The SoftLayer_Dns_Domain_Registration data type represents a domain registration record.
// type Dns_Domain_Registration struct {
// 	Entity

// 	// The SoftLayer customer account that the domain is registered to.
// 	Account *Account `json:"account,omitempty" xmlrpc:"account,omitempty"`

// 	// no documentation yet
// 	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate,omitempty"`

// 	// The domain registration status.
// 	DomainRegistrationStatus *Dns_Domain_Registration_Status `json:"domainRegistrationStatus,omitempty" xmlrpc:"domainRegistrationStatus,omitempty"`

// 	// no documentation yet
// 	DomainRegistrationStatusId *int `json:"domainRegistrationStatusId,omitempty" xmlrpc:"domainRegistrationStatusId,omitempty"`

// 	// The date that the domain registration will expire.
// 	ExpireDate *Time `json:"expireDate,omitempty" xmlrpc:"expireDate,omitempty"`

// 	// A domain record's internal identifier.
// 	Id *int `json:"id,omitempty" xmlrpc:"id,omitempty"`

// 	// Indicates whether a domain is locked or unlocked.
// 	LockedFlag *int `json:"lockedFlag,omitempty" xmlrpc:"lockedFlag,omitempty"`

// 	// no documentation yet
// 	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate,omitempty"`

// 	// A domain's name, for example "example.com".
// 	Name *string `json:"name,omitempty" xmlrpc:"name,omitempty"`

// 	// The registrant verification status.
// 	RegistrantVerificationStatus *Dns_Domain_Registration_Registrant_Verification_Status `json:"registrantVerificationStatus,omitempty" xmlrpc:"registrantVerificationStatus,omitempty"`

// 	// no documentation yet
// 	RegistrantVerificationStatusId *int `json:"registrantVerificationStatusId,omitempty" xmlrpc:"registrantVerificationStatusId,omitempty"`

// 	// no documentation yet
// 	ServiceProvider *Service_Provider `json:"serviceProvider,omitempty" xmlrpc:"serviceProvider,omitempty"`

// 	// no documentation yet
// 	ServiceProviderId *int `json:"serviceProviderId,omitempty" xmlrpc:"serviceProviderId,omitempty"`
// }

// SoftLayer_Dns_Domain_ResourceRecord_NsType is a SoftLayer_Dns_Domain_ResourceRecord object whose ''type'' property is set to "ns" and defines a DNS NS record on a SoftLayer hosted domain. An NS record defines the authoritative name server for a domain. All SoftLayer hosted domains contain NS records for "ns1.softlayer.com" and "ns2.softlayer.com" . For instance, if example.org is hosted on ns1.softlayer.com, then example.org contains an NS record whose ''host'' property equals "@" and whose ''data'' property equals "ns1.example.org".
//
// NS resource records pointing to ns1.softlayer.com or ns2.softlayer.com many not be removed from a SoftLayer hosted domain.
// type Dns_Domain_ResourceRecord_NsType struct {
// 	Dns_Domain_ResourceRecord
// }

func dataSourceIBMDNSDomainRegistration() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMDNSDomainRegistrationRead,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: "A domain registration record's internal identifier",
				Type:        schema.TypeInt,
				Computed:    true,
			},

			"name": &schema.Schema{
				Description: "The name of the domain registration",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name_servers": &schema.Schema{
				Description: "Custom name servers for the domain registration",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceIBMDNSDomainRegistrationRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetAccountService(sess)

	name := d.Get("name").(string)
	names, err := service.
		Filter(filter.Build(filter.Path("domainRegistrations.name").Eq(name))).
		Mask("id,name").
		GetDomainRegistrations()

	if err != nil {
		return fmt.Errorf("Error retrieving domain registration: %s", err)
	}

	if len(names) == 0 {
		return fmt.Errorf("No domain registration found with name [%s]", name)
	}

	log.Printf("names %v\n", names)
	dnsId := *names[0].Id
	log.Printf("Domain Registration Id %d\n", dnsId)
	// Get nameservers for domain

	nService := services.GetDnsDomainRegistrationService(sess)

	// dnsId, _ := strconv.Atoi(d.Id())

	// retrieve remote object state
	dns_domain_nameservers, err := nService.Id(dnsId).
		Mask("nameservers.name").
		GetDomainNameservers()

	log.Printf("list %v\n", dns_domain_nameservers)

	//ns := *dns_domain_nameservers[0].Nameservers[0].Name
	ns := make([]string, len(dns_domain_nameservers[0].Nameservers))
	//var ns []string
	for i, elem := range dns_domain_nameservers[0].Nameservers {
		ns[i] = *elem.Name
	}

	log.Printf("names %v\n", ns)

	if err != nil {
		return fmt.Errorf("Error retrieving domain registration nameservers: %s", err)
	}

	d.SetId(fmt.Sprintf("%d", dnsId))
	d.Set("name_servers", ns)
	return nil
}
