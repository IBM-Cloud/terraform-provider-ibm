package ibm

import (
	fmt "fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/helpers/product"
	services "github.com/softlayer/softlayer-go/services"
	session1 "github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

const (
	AdditionalSSLServicesPackageType            = "ADDITIONAL_SERVICES"
	AdditionalServicesSSLCertificatePackageType = "ADDITIONAL_SERVICES_SSL_CERTIFICATE"

	SSLMask = "id"
)

func resourceIBMSSLCertificate() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMSSLCertificateCreate,
		Read:     resourceIBMSSLCertificateRead,
		Update:   resourceIBMSSLCertificateUpdate,
		Delete:   resourceIBMSSLCertificateDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"serverCount": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"serverType": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"validityMonths": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"sslType": {
				Type:     schema.TypeString,
				Required: true,
			},

			"certificateSigningRequest": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"renewalFlag": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"orderApproverEmailAddress": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"technicalContactSameAsOrgAddressFlag": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"administrativeContactSameAsTechnicalFlag": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"billingContactSameAsTechnicalFlag": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"organizationInformation": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"org_address": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"org_addressLine1": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"org_addressLine2": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},

									"org_city": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"org_countryCode": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"org_postalCode": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"org_state": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"org_organizationName": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"org_phoneNumber": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"org_faxNumber": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"technicalContact": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"tech_address": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tech_addressLine1": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"tech_addressLine2": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},

									"tech_city": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"tech_countryCode": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"tech_postalCode": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"tech_state": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"tech_organizationName": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"tech_firstName": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"tech_lastName": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"tech_emailAddress": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"tech_phoneNumber": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"tech_faxNumber": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"tech_title": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"billingContact": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"billing_address": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"billing_addressLine1": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"billing_addressLine2": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},

									"billing_city": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"billing_countryCode": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"billing_postalCode": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"billing_state": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"billing_organizationName": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"billing_firstName": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"billing_lastName": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"billing_emailAddress": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"billing_phoneNumber": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"billing_faxNumber": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"billing_title": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"administrativeContact": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"admin_address": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"admin_addressLine1": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"admin_addressLine2": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},

									"admin_city": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"admin_countryCode": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"admin_postalCode": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"admin_state": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"admin_organizationName": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"admin_firstName": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"admin_lastName": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"admin_emailAddress": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"admin_phoneNumber": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"admin_faxNumber": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"admin_title": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}
func resourceIBMSSLCertificateCreate(d *schema.ResourceData, m interface{}) error {
	sess := m.(ClientSession).SoftLayerSession()
	service := services.GetSecurityCertificateRequestService(sess.SetRetries(0))
	sslKeyName := sl.String(d.Get("sslType").(string))
	pkg, err := product.GetPackageByType(sess, AdditionalServicesSSLCertificatePackageType)
	if err != nil {
		return err
	}
	productItems, err := product.GetPackageProducts(sess, *pkg.Id)
	if err != nil {
		return err
	}
	var itemId *int
	for _, item := range productItems {
		if *item.KeyName == *sslKeyName {
			itemId = item.Id
		}
	}
	validCSR, err := service.ValidateCsr(sl.String(d.Get("certificateSigningRequest").(string)), sl.Int(d.Get("validityMonths").(int)), itemId, sl.String(d.Get("serverType").(string)))
	if err != nil {
		return fmt.Errorf("Error during validation of CSR: %s", err)
	}
	if validCSR == true {
		productOrderContainer, err := buildSSLProductOrderContainer(d, sess, AdditionalServicesSSLCertificatePackageType)
		if err != nil {
			// Find price items with AdditionalServices
			productOrderContainer, err = buildSSLProductOrderContainer(d, sess, AdditionalSSLServicesPackageType)
			if err != nil {
				return fmt.Errorf("Error creating SSL certificate: %s", err)
			}
		}
		log.Printf("[INFO] Creating SSL Certificate")
		verifiedOrderContainer, err := services.GetProductOrderService(sess).VerifyOrder(productOrderContainer)
		if err != nil {
			return fmt.Errorf("Order verification failed: %s", err)
		}

		servercorecount := verifiedOrderContainer.ServerCoreCount
		log.Println(verifiedOrderContainer)
		log.Printf("ServerCoreCount: %d", servercorecount)
		receipt, err := services.GetProductOrderService(sess).PlaceOrder(productOrderContainer, sl.Bool(false))

		if err != nil {
			return fmt.Errorf("Error during creation of ssl: %s", err)
		}

		ssl, err := findSSLByOrderId(sess, *receipt.OrderId)
		d.SetId(fmt.Sprintf("%d", *ssl.Id))
		return resourceIBMSSLCertificateRead(d, m)
	} else {
		log.Println("Provided CSR is not valid.")
		return nil
	}
}

func resourceIBMSSLCertificateRead(d *schema.ResourceData, m interface{}) error {
	sess := m.(ClientSession).SoftLayerSession()
	service := services.GetSecurityCertificateRequestService(sess)
	sslId, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Not a valid SSL ID, must be an integer: %s", err)
	}

	ssl, err := service.Id(sslId).Mask(SSLMask).GetObject()
	if err != nil {
		return fmt.Errorf("Error retrieving SSL: %s", err)
	}
	d.Set("id", *ssl.Id)
	d.Set("CertificateSigningRequest", ssl.CertificateSigningRequest)
	return nil
}

func resourceIBMSSLCertificateUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIBMSSLCertificateDelete(d *schema.ResourceData, m interface{}) error {
	sess := m.(ClientSession).SoftLayerSession()
	service := services.GetSecurityCertificateService(sess)
	service1 := services.GetSecurityCertificateRequestService(sess)
	sslId, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Not a valid SSL ID, must be an integer: %s", err)
	}

	value, err := service1.Id(sslId).GetObject()
	if err != nil {
		return fmt.Errorf("Not a valid Object ID: %s", err)
	}
	sslReqId := value.StatusId

	if *sslReqId == 49 || *sslReqId == 43 {
		deleteObject, err := service.Id(sslId).DeleteObject()
		if deleteObject == false {
			return fmt.Errorf("Error deleting SSL: %s", err)
		} else {
			d.SetId("")
			return nil
		}
	} else if *sslReqId == 50 {
		cancelObject, err := service1.Id(sslId).CancelSslOrder()
		if cancelObject == false {
			return fmt.Errorf("Error deleting SSL: %s", err)
		} else {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
}

func normalizedCert(cert interface{}) string {
	if cert == nil || cert == (*string)(nil) {
		return ""
	}

	switch cert.(type) {
	case string:
		return strings.TrimSpace(cert.(string))
	default:
		return ""
	}
}

func buildSSLProductOrderContainer(d *schema.ResourceData, sess *session1.Session, packageType string) (*datatypes.Container_Product_Order_Security_Certificate, error) {
	certificateSigningRequest := sl.String(d.Get("certificateSigningRequest").(string))
	orderApproverEmailAddress := sl.String(d.Get("orderApproverEmailAddress").(string))
	renewalFlag := sl.Bool(d.Get("renewalFlag").(bool))
	serverCount := sl.Int(d.Get("serverCount").(int))
	validityMonths := sl.Int(d.Get("validityMonths").(int))
	serverType := sl.String(d.Get("serverType").(string))
	sslType := sl.String(d.Get("sslType").(string))
	orgnizationInfoList := d.Get("organizationInformation").(*schema.Set).List()
	var addressline1, addressline2, city, countryCode, state, postalCode, organizationName, phoneNumber, faxNumber string
	for _, orgnizationInfo := range orgnizationInfoList {
		org_info := orgnizationInfo.(map[string]interface{})
		org_addressList := org_info["org_address"].(*schema.Set).List()
		for _, org_address := range org_addressList {
			org_addr := org_address.(map[string]interface{})
			addressline1 = org_addr["org_addressLine1"].(string)
			addressline2 = org_addr["org_addressLine2"].(string)
			city = org_addr["org_city"].(string)
			countryCode = org_addr["org_countryCode"].(string)
			state = org_addr["org_state"].(string)
			postalCode = org_addr["org_postalCode"].(string)
		}
		organizationName = org_info["org_organizationName"].(string)
		phoneNumber = org_info["org_phoneNumber"].(string)
		faxNumber = org_info["org_faxNumber"].(string)
	}
	org_address_information := datatypes.Container_Product_Order_Attribute_Address{
		AddressLine1: &addressline1,
		AddressLine2: &addressline2,
		City:         &city,
		CountryCode:  &countryCode,
		PostalCode:   &postalCode,
		State:        &state,
	}
	org_information := datatypes.Container_Product_Order_Attribute_Organization{
		Address:          &org_address_information,
		OrganizationName: &organizationName,
		PhoneNumber:      &phoneNumber,
		FaxNumber:        &faxNumber,
	}
	TechInfoList := d.Get("technicalContact").(*schema.Set).List()
	var tech_addressline1, tech_addressline2, tech_city, tech_countryCode, tech_state, tech_postalCode, tech_organizationName, tech_phoneNumber, tech_faxNumber, tech_emailAddress, tech_firstName, tech_lastName, tech_title string
	for _, technicalcont := range TechInfoList {
		tech_contact := technicalcont.(map[string]interface{})
		tect_addressList := tech_contact["tech_address"].(*schema.Set).List()
		for _, tech_address := range tect_addressList {
			tech_addr := tech_address.(map[string]interface{})
			tech_addressline1 = tech_addr["tech_addressLine1"].(string)
			tech_addressline2 = tech_addr["tech_addressLine2"].(string)
			tech_city = tech_addr["tech_city"].(string)
			tech_countryCode = tech_addr["tech_countryCode"].(string)
			tech_state = tech_addr["tech_state"].(string)
			tech_postalCode = tech_addr["tech_postalCode"].(string)
		}
		tech_organizationName = tech_contact["tech_organizationName"].(string)
		tech_phoneNumber = tech_contact["tech_phoneNumber"].(string)
		tech_faxNumber = tech_contact["tech_faxNumber"].(string)
		tech_emailAddress = tech_contact["tech_emailAddress"].(string)
		tech_firstName = tech_contact["tech_firstName"].(string)
		tech_lastName = tech_contact["tech_lastName"].(string)
		tech_title = tech_contact["tech_title"].(string)
	}
	tech_address_information := datatypes.Container_Product_Order_Attribute_Address{
		AddressLine1: &tech_addressline1,
		AddressLine2: &tech_addressline2,
		City:         &tech_city,
		CountryCode:  &tech_countryCode,
		PostalCode:   &tech_postalCode,
		State:        &tech_state,
	}
	techAddressFlag := d.Get("technicalContactSameAsOrgAddressFlag").(bool)
	var technical_contact_attr datatypes.Container_Product_Order_Attribute_Contact
	if techAddressFlag {
		technical_contact_attr = datatypes.Container_Product_Order_Attribute_Contact{
			Address:          &org_address_information,
			EmailAddress:     &tech_emailAddress,
			FirstName:        &tech_firstName,
			LastName:         &tech_lastName,
			OrganizationName: &tech_organizationName,
			PhoneNumber:      &tech_phoneNumber,
			FaxNumber:        &tech_faxNumber,
			Title:            &tech_title,
		}
	} else {
		technical_contact_attr = datatypes.Container_Product_Order_Attribute_Contact{
			Address:          &tech_address_information,
			EmailAddress:     &tech_emailAddress,
			FirstName:        &tech_firstName,
			LastName:         &tech_lastName,
			OrganizationName: &tech_organizationName,
			PhoneNumber:      &tech_phoneNumber,
			FaxNumber:        &tech_faxNumber,
			Title:            &tech_title,
		}
	}

	administrativeContactList := d.Get("administrativeContact").(*schema.Set).List()
	var admin_addressline1, admin_addressline2, admin_city, admin_countryCode, admin_state, admin_postalCode, admin_organizationName, admin_phoneNumber, admin_faxNumber, admin_emailAddress, admin_firstName, admin_lastName, admin_title string
	for _, administrativecont := range administrativeContactList {
		administrative_contact := administrativecont.(map[string]interface{})
		administrative_addressList := administrative_contact["admin_address"].(*schema.Set).List()
		for _, admin_address := range administrative_addressList {
			admin_addr := admin_address.(map[string]interface{})
			admin_addressline1 = admin_addr["admin_addressLine1"].(string)
			admin_addressline2 = admin_addr["admin_addressLine2"].(string)
			admin_city = admin_addr["admin_city"].(string)
			admin_countryCode = admin_addr["admin_countryCode"].(string)
			admin_state = admin_addr["admin_state"].(string)
			admin_postalCode = admin_addr["admin_postalCode"].(string)
		}
		admin_organizationName = administrative_contact["admin_organizationName"].(string)
		admin_phoneNumber = administrative_contact["admin_phoneNumber"].(string)
		admin_faxNumber = administrative_contact["admin_faxNumber"].(string)
		admin_emailAddress = administrative_contact["admin_emailAddress"].(string)
		admin_firstName = administrative_contact["admin_firstName"].(string)
		admin_lastName = administrative_contact["admin_lastName"].(string)
		admin_title = administrative_contact["admin_title"].(string)
	}
	administrative_address_information := datatypes.Container_Product_Order_Attribute_Address{
		AddressLine1: &admin_addressline1,
		AddressLine2: &admin_addressline2,
		City:         &admin_city,
		CountryCode:  &admin_countryCode,
		PostalCode:   &admin_postalCode,
		State:        &admin_state,
	}

	administrative_contact_attr := datatypes.Container_Product_Order_Attribute_Contact{
		Address:          &administrative_address_information,
		EmailAddress:     &admin_emailAddress,
		FirstName:        &admin_firstName,
		LastName:         &admin_lastName,
		OrganizationName: &admin_organizationName,
		PhoneNumber:      &admin_phoneNumber,
		FaxNumber:        &admin_faxNumber,
		Title:            &admin_title,
	}

	billingContactList := d.Get("billingContact").(*schema.Set).List()
	var bill_addressline1, bill_addressline2, bill_city, bill_countryCode, bill_state, bill_postalCode, bill_organizationName, bill_phoneNumber, bill_faxNumber, bill_emailAddress, bill_firstName, bill_lastName, bill_title string
	for _, billingcont := range billingContactList {
		billing_contact := billingcont.(map[string]interface{})
		billing_addressList := billing_contact["billing_address"].(*schema.Set).List()
		for _, billing_address := range billing_addressList {
			billing_addr := billing_address.(map[string]interface{})
			bill_addressline1 = billing_addr["billing_addressLine1"].(string)
			bill_addressline2 = billing_addr["billing_addressLine2"].(string)
			bill_city = billing_addr["billing_city"].(string)
			bill_countryCode = billing_addr["billing_countryCode"].(string)
			bill_state = billing_addr["billing_state"].(string)
			bill_postalCode = billing_addr["billing_postalCode"].(string)
		}
		bill_organizationName = billing_contact["billing_organizationName"].(string)
		bill_phoneNumber = billing_contact["billing_phoneNumber"].(string)
		bill_faxNumber = billing_contact["billing_faxNumber"].(string)
		bill_emailAddress = billing_contact["billing_emailAddress"].(string)
		bill_firstName = billing_contact["billing_firstName"].(string)
		bill_lastName = billing_contact["billing_lastName"].(string)
		bill_title = billing_contact["billing_title"].(string)
	}
	billing_address_information := datatypes.Container_Product_Order_Attribute_Address{
		AddressLine1: &bill_addressline1,
		AddressLine2: &bill_addressline2,
		City:         &bill_city,
		CountryCode:  &bill_countryCode,
		PostalCode:   &bill_postalCode,
		State:        &bill_state,
	}

	billing_contact_attr := datatypes.Container_Product_Order_Attribute_Contact{
		Address:          &billing_address_information,
		EmailAddress:     &bill_emailAddress,
		FirstName:        &bill_firstName,
		LastName:         &bill_lastName,
		OrganizationName: &bill_organizationName,
		PhoneNumber:      &bill_phoneNumber,
		FaxNumber:        &bill_faxNumber,
		Title:            &bill_title,
	}

	administrativeContactSameAsTechnical := d.Get("administrativeContactSameAsTechnicalFlag").(bool)
	billingContactSameAsTechnical := d.Get("billingContactSameAsTechnicalFlag").(bool)
	if administrativeContactSameAsTechnical {
		administrative_contact_attr = technical_contact_attr
	}
	if billingContactSameAsTechnical {
		billing_contact_attr = technical_contact_attr
	}
	pkg, err := product.GetPackageByType(sess, packageType)
	if err != nil {
		return &datatypes.Container_Product_Order_Security_Certificate{}, err
	}

	productItems, err := product.GetPackageProducts(sess, *pkg.Id)
	if err != nil {
		return &datatypes.Container_Product_Order_Security_Certificate{}, err
	}
	sslKeyName := sslType

	sslItems := []datatypes.Product_Item{}
	for _, item := range productItems {
		if *item.KeyName == *sslKeyName {
			sslItems = append(sslItems, item)
		}
	}

	if len(sslItems) == 0 {
		return &datatypes.Container_Product_Order_Security_Certificate{},
			fmt.Errorf("No product items matching %p could be found", sslKeyName)
	}
	sslContainer := datatypes.Container_Product_Order_Security_Certificate{
		Container_Product_Order: datatypes.Container_Product_Order{
			PackageId: pkg.Id,
			Prices: []datatypes.Product_Item_Price{
				{
					Id: sslItems[0].Prices[0].Id,
				},
			},
			Quantity: sl.Int(1),
		},
		AdministrativeContact:     &administrative_contact_attr,
		BillingContact:            &billing_contact_attr,
		CertificateSigningRequest: certificateSigningRequest,
		OrderApproverEmailAddress: orderApproverEmailAddress,
		OrganizationInformation:   &org_information,
		RenewalFlag:               renewalFlag,
		ServerCount:               serverCount,
		ServerType:                serverType,
		TechnicalContact:          &technical_contact_attr,
		ValidityMonths:            validityMonths,
	}

	return &sslContainer, nil
}

func findSSLByOrderId(sess *session1.Session, orderId int) (datatypes.Security_Certificate_Request, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"pending"},
		Target:  []string{"complete"},
		Refresh: func() (interface{}, string, error) {
			acc := services.GetAccountService(sess)
			acc_attr, err := acc.GetAttributes()
			acc_id := acc_attr[0].AccountId
			ssls, err := services.GetSecurityCertificateRequestService(sess).Filter(filter.Path("securityCertificateRequest.order.id").Eq(strconv.Itoa(orderId)).Build()).Mask("id").GetSslCertificateRequests(acc_id)
			if err != nil {
				return datatypes.Security_Certificate_Request{}, "", err
			}

			if len(ssls) >= 1 {
				return ssls[0], "complete", nil
			} else {
				return nil, "pending", nil
			}
		},
		Timeout:    10 * time.Minute,
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}
	pendingResult, err := stateConf.WaitForState()

	if err != nil {
		return datatypes.Security_Certificate_Request{}, err
	}

	var result, ok = pendingResult.(datatypes.Security_Certificate_Request)

	if ok {
		return result, nil
	}

	return datatypes.Security_Certificate_Request{},
		fmt.Errorf("Cannot find SSl with order id '%d'", orderId)
}
