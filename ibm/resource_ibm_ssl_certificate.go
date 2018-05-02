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
			"csr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"address1": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"address2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"city": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"country_code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"postal_code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"state": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"org_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"phone_no": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"first_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"last_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"renewal": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"server_count": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"validity_month": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"ssl_type": {
				Type:     schema.TypeString,
				Required: true,
			},

			/*"common_name": {
				Type:     schema.TypeString,
				Required: true,
			},*/

			"title": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIBMSSLCertificateCreate(d *schema.ResourceData, m interface{}) error {
	sess := m.(ClientSession).SoftLayerSession()
	service := services.GetSecurityCertificateRequestService(sess.SetRetries(0))
	sslKeyName := sl.String(d.Get("ssl_type").(string))
	/*commonName := sl.String(d.Get("common_name").(string))

	common_name, err := service.GetAdministratorEmailDomains(commonName)
	if err != nil {
		return err
	}

	log.Println(common_name)*/
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
	log.Printf("///////////////////===%d===////////////////////", *itemId)
	validCSR, err := service.ValidateCsr(sl.String(d.Get("csr").(string)), sl.Int(d.Get("validity_month").(int)), itemId, sl.String(d.Get("server_type").(string)))
	if err != nil {
		return fmt.Errorf("Error during validation of CSR: %s", err)
	}
	log.Println(validCSR)
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

		server_cnt := verifiedOrderContainer.ServerCoreCount
		log.Println(verifiedOrderContainer)
		log.Printf("Server_count: %d", server_cnt)
		receipt, err := services.GetProductOrderService(sess).PlaceOrder(productOrderContainer, sl.Bool(false))

		if err != nil {
			return fmt.Errorf("Error during creation of ssl: %s", err)
		}

		ssl, err := findSSLByOrderId(sess, *receipt.OrderId)

		d.SetId(fmt.Sprintf("%d", *ssl.Id))

		log.Println("//////////////////////////////SSL_ID///////////////////////////////////////")
		log.Printf("%d", *ssl.Id)
		log.Printf("ssl created successfully with OrderID: %d", receipt.OrderId)
		log.Println("\\\\\\\\\\\\\\\\\\\\\\\\\\\\SSL_ID//////////////////////////////////////////")
		return resourceIBMSSLCertificateRead(d, m)
	} else {
		log.Println("Please enter valid CSR")
		return nil
	}
}

func resourceIBMSSLCertificateRead(d *schema.ResourceData, m interface{}) error {
	sess := m.(ClientSession).SoftLayerSession()
	service := services.GetSecurityCertificateRequestService(sess)
	sslservice := services.GetSecurityCertificateRequestServerTypeService(sess)
	sslId, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Not a valid SSL ID, must be an integer: %s", err)
	}

	ssl, err := service.Id(sslId).Mask(SSLMask).GetObject()
	if err != nil {
		return fmt.Errorf("Error retrieving SSL: %s", err)
	}

	log.Println(ssl)
	sslserver, err := sslservice.GetAllObjects()
	if err != nil {
		return fmt.Errorf("Error retrieving SSL servers: %s", err)
	}
	log.Println(sslserver)

	d.Set("id", *ssl.Id)
	d.Set("csr", ssl.CertificateSigningRequest)

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
			id := sslId
			log.Printf("ID: %d", id)
			d.SetId("")
			return nil
		}
	} else if *sslReqId == 50 {
		cancelObject, err := service1.Id(sslId).CancelSslOrder()
		if cancelObject == false {
			return fmt.Errorf("Error deleting SSL: %s", err)
		} else {
			id := sslId
			log.Printf("ID: %d", id)
			d.SetId("")
			return nil
		}
	} else {
		id := sslId
		log.Printf("ID: %d", id)
		d.SetId("")
		return nil
	}
	/*if value == false {
		return fmt.Errorf("Error deleting SSL: %s", err)
	} else {
		id := sslId
		log.Printf("ID: %d", id)
		d.SetId("")
		return nil
	}*/

	//_, err = services.GetProductOrderService(sess).Id(*billingItem.Id).CheckItemAvailability()

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
	address_attr := datatypes.Container_Product_Order_Attribute_Address{
		AddressLine1: sl.String(d.Get("address1").(string)),
		AddressLine2: sl.String(d.Get("address2").(string)),
		City:         sl.String(d.Get("city").(string)),
		CountryCode:  sl.String(d.Get("country_code").(string)),
		PostalCode:   sl.String(d.Get("postal_code").(string)),
		State:        sl.String(d.Get("state").(string)),
	}

	contact_attr := datatypes.Container_Product_Order_Attribute_Contact{
		Address:          &address_attr,
		EmailAddress:     sl.String(d.Get("email").(string)),
		FirstName:        sl.String(d.Get("first_name").(string)),
		LastName:         sl.String(d.Get("last_name").(string)),
		OrganizationName: sl.String(d.Get("org_name").(string)),
		PhoneNumber:      sl.String(d.Get("phone_no").(string)),
		Title:            sl.String(d.Get("title").(string)),
	}
	csr := sl.String(d.Get("csr").(string))
	email := sl.String(d.Get("email").(string))

	org_attr := datatypes.Container_Product_Order_Attribute_Organization{
		Address:          &address_attr,
		OrganizationName: sl.String(d.Get("org_name").(string)),
		PhoneNumber:      sl.String(d.Get("phone_no").(string)),
	}

	renewalflag := sl.Bool(d.Get("renewal").(bool))
	server_count := sl.Int(d.Get("server_count").(int))
	validity_month := sl.Int(d.Get("validity_month").(int))

	pkg, err := product.GetPackageByType(sess, packageType)
	if err != nil {
		return &datatypes.Container_Product_Order_Security_Certificate{}, err
	}

	productItems, err := product.GetPackageProducts(sess, *pkg.Id)
	if err != nil {
		return &datatypes.Container_Product_Order_Security_Certificate{}, err
	}
	sslKeyName := sl.String(d.Get("ssl_type").(string))

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
		AdministrativeContact:     &contact_attr,
		BillingContact:            &contact_attr,
		CertificateSigningRequest: csr,
		OrderApproverEmailAddress: email,
		OrganizationInformation:   &org_attr,
		RenewalFlag:               renewalflag,
		ServerCount:               server_count,
		ServerType:                sl.String(d.Get("server_type").(string)),
		TechnicalContact:          &contact_attr,
		ValidityMonths:            validity_month,
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
