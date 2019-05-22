package ibm

import (
	"fmt"
	"log"
	"strconv"

	"github.com/softlayer/softlayer-go/sl"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
)

const str string = ".cdnedge.bluemix.net"

func resourceIBMCDN() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMCDNCreate,
		Read:   resourceIBMCDNRead,
		Update: resourceIBMCDNUpdate,
		Delete: resourceIBMCDNDelete,
		Exists: resourceIBMCDNExists,

		Schema: map[string]*schema.Schema{
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vendor_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "akamai",
				ForceNew: true,
			},

			"origin_type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "HOST_SERVER",
				ForceNew:     true,
				ValidateFunc: validateAllowedStringValue([]string{"HOST_SERVER", "OBJECT_STORAGE"}),
			},
			"origin_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"bucket_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "HTTP",
				ForceNew:     true,
				ValidateFunc: validateAllowedStringValue([]string{"HTTP", "HTTPS", "HTTP_AND_HTTPS"}),
			},
			"http_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  80,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"https_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  443,
			},
			"cname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"respect_headers": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"file_extension": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"certificate_type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"SHARED_SAN_CERT", "WILDCARD_CERT"}),
				ForceNew:     true,
			},
			"cache_key_query_rule": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"include-all", "ignore-all", "ignore: space separated query-args", "include: space separated query-args"}),
				Default:      "include-all",
			},
			"performance_configuration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "General web delivery",
				ForceNew: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/*",
				ForceNew: true,
			},
		},
	}
}

func resourceIBMCDNCreate(d *schema.ResourceData, meta interface{}) error {
	///create  session
	sess := meta.(ClientSession).SoftLayerSession()
	///get the value of all the parameters
	domain := d.Get("host_name").(string)
	vendorname := d.Get("vendor_name").(string)
	origintype := d.Get("origin_type").(string)
	originaddress := d.Get("origin_address").(string)
	protocol := d.Get("protocol").(string)
	httpport := d.Get("http_port").(int)
	httpsport := d.Get("https_port").(int)
	bucketname := d.Get("bucket_name").(string)
	path := d.Get("path").(string)
	header := d.Get("header").(string)
	cachekeyqueryrule := d.Get("cache_key_query_rule").(string)
	performanceconfiguration := d.Get("performance_configuration").(string)
	respectheaders := d.Get("respect_headers").(bool)
	var rHeader = "0"
	if respectheaders {
		rHeader = "1"
	}
	cname := d.Get("cname").(string)
	certificateType := d.Get("certificate_type").(string)
	if name, ok := d.GetOk("cname"); ok {
		cname = name.(string) + str
	}

	///creat an object of CDN service
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)
	//////pass the parameters to create domain mapping
	if origintype == "OBJECT_STORAGE" && protocol == "HTTP" {
		receipt1, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Cname:                    sl.String(cname),
			Protocol:                 sl.String(protocol),
			HttpPort:                 sl.Int(httpport),
			OriginType:               sl.String(origintype),
			BucketName:               sl.String(bucketname),
			Header:                   sl.String(header),
			RespectHeaders:           sl.String(rHeader),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		if err != nil {
			return fmt.Errorf("Error creating CDN: %s", err)
		}

		d.SetId(*receipt1[0].UniqueId)
		id, err := strconv.Atoi((d.Id()))
		result1, err := service.VerifyDomainMapping(&id)
		log.Print("The status of domain mapping ", result1)
		return resourceIBMCDNRead(d, meta)

	}
	if origintype == "OBJECT_STORAGE" && protocol == "HTTPS" {
		receipt2, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Cname:                    sl.String(cname),
			Protocol:                 sl.String(protocol),
			HttpsPort:                sl.Int(httpsport),
			OriginType:               sl.String(origintype),
			BucketName:               sl.String(bucketname),
			Header:                   sl.String(header),
			RespectHeaders:           sl.String(rHeader),
			CertificateType:          sl.String(certificateType),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		if err != nil {
			return fmt.Errorf("Error creating CDN: %s", err)
		}

		d.SetId(*receipt2[0].UniqueId)
		id, err := strconv.Atoi((d.Id()))
		result2, err := service.VerifyDomainMapping(&id)
		log.Print("The status of domain mapping ", result2)
		return resourceIBMCDNRead(d, meta)
	}
	if origintype == "OBJECT_STORAGE" && protocol == "HTTP_AND_HTTPS" {
		receipt3, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Cname:                    sl.String(cname),
			Protocol:                 sl.String(protocol),
			HttpPort:                 sl.Int(httpport),
			HttpsPort:                sl.Int(httpsport),
			OriginType:               sl.String(origintype),
			BucketName:               sl.String(bucketname),
			Header:                   sl.String(header),
			RespectHeaders:           sl.String(rHeader),
			CertificateType:          sl.String(certificateType),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		if err != nil {
			return fmt.Errorf("Error creating CDN: %s", err)
		}

		d.SetId(*receipt3[0].UniqueId)
		id, err := strconv.Atoi((d.Id()))
		result3, err := service.VerifyDomainMapping(&id)
		log.Print("The status of domain mapping ", result3)
		return resourceIBMCDNRead(d, meta)
	}
	if origintype == "HOST_SERVER" && protocol == "HTTP" {
		receipt4, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Cname:                    sl.String(cname),
			Protocol:                 sl.String(protocol),
			HttpPort:                 sl.Int(httpport),
			OriginType:               sl.String(origintype),
			Header:                   sl.String(header),
			RespectHeaders:           sl.String(rHeader),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		if err != nil {
			return fmt.Errorf("Error creating CDN: %s", err)
		}

		d.SetId(*receipt4[0].UniqueId)
		id, err := strconv.Atoi((d.Id()))
		result4, err := service.VerifyDomainMapping(&id)
		log.Print("The status of domain mapping ", result4)
		return resourceIBMCDNRead(d, meta)
	}
	if origintype == "HOST_SERVER" && protocol == "HTTPS" {
		receipt5, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Cname:                    sl.String(cname),
			Protocol:                 sl.String(protocol),
			HttpsPort:                sl.Int(httpsport),
			OriginType:               sl.String(origintype),
			Header:                   sl.String(header),
			RespectHeaders:           sl.String(rHeader),
			CertificateType:          sl.String(certificateType),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		if err != nil {
			return fmt.Errorf("Error creating CDN: %s", err)
		}

		d.SetId(*receipt5[0].UniqueId)
		id, err := strconv.Atoi((d.Id()))
		result5, err := service.VerifyDomainMapping(&id)
		log.Print("The status of domain mapping ", result5)
		return resourceIBMCDNRead(d, meta)
	}
	if origintype == "HOST_SERVER" && protocol == "HTTP_AND_HTTPS" {
		receipt6, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Cname:                    sl.String(cname),
			Protocol:                 sl.String(protocol),
			HttpPort:                 sl.Int(httpport),
			HttpsPort:                sl.Int(httpsport),
			OriginType:               sl.String(origintype),
			Header:                   sl.String(header),
			RespectHeaders:           sl.String(rHeader),
			CertificateType:          sl.String(certificateType),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		if err != nil {
			return fmt.Errorf("Error creating CDN: %s", err)
		}

		d.SetId(*receipt6[0].UniqueId)
		id, err := strconv.Atoi((d.Id()))
		result6, err := service.VerifyDomainMapping(&id)
		log.Print("The status of domain mapping ", result6)
		return resourceIBMCDNRead(d, meta)
	}

	return nil
}

func resourceIBMCDNRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)
	cdnId := sl.String(d.Id())
	///read the changes in the remote resource and update in the local resource.
	read, err := service.ListDomainMappingByUniqueId(cdnId)
	///Print the response of the requested the service.
	d.Set("id", *read[0].UniqueId)
	d.Set("originaddress", *read[0].OriginHost)
	d.Set("vendorname", *read[0].VendorName)
	d.Set("domain", *read[0].Domain)
	d.Set("header", *read[0].Header)
	d.Set("cname", *read[0].Cname)
	d.Set("origin_type", *read[0].OriginType)
	d.Set("status", *read[0].Status)
	if *read[0].OriginType == "OBJECT_STORAGE" {
		d.Set("bucketname", *read[0].BucketName)
	}
	if *read[0].Protocol == "HTTP" || *read[0].Protocol == "HTTP_AND_HTTPS" {
		d.Set("httpport", *read[0].HttpPort)
	}
	if *read[0].Protocol == "HTTPS" || *read[0].Protocol == "HTTP_AND_HTTPS" {
		d.Set("httpsport", *read[0].HttpsPort)
	}
	d.Set("protocol", *read[0].Protocol)
	d.Set("respectheaders", *read[0].RespectHeaders)
	d.Set("certificationtype", *read[0].CertificateType)
	d.Set("cachekeyqueryrule", *read[0].CacheKeyQueryRule)
	d.Set("path", *read[0].Path)
	d.Set("performanceconfiguration", *read[0].PerformanceConfiguration)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func resourceIBMCDNUpdate(d *schema.ResourceData, meta interface{}) error {
	/// Nothing to update for now. Not supported.
	sess := meta.(ClientSession).SoftLayerSession()
	domain := d.Get("host_name").(string)
	vendorname := d.Get("vendor_name").(string)
	origintype := d.Get("origin_type").(string)
	originaddress := d.Get("origin_address").(string)
	protocol := d.Get("protocol").(string)
	httpport := d.Get("http_port").(int)
	httpsport := d.Get("https_port").(int)
	path := d.Get("path").(string)
	cname := d.Get("cname").(string)
	header := d.Get("header").(string)
	bucketname := d.Get("bucket_name").(string)
	fileextension := d.Get("file_extension").(string)
	respectheaders := d.Get("respect_headers").(bool)
	var rHeader = "0"
	if respectheaders {
		rHeader = "1"
	}
	certificateType := d.Get("certificate_type").(string)
	cachekeyqueryrule := d.Get("cache_key_query_rule").(string)
	performanceconfiguration := d.Get("performance_configuration").(string)
	uniqueId := d.Id()
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)
	///pass the changed as well as unchanged parameters to update the resource.

	if origintype == "HOST_SERVER" && protocol == "HTTP_AND_HTTPS" {
		update1, err := service.UpdateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Protocol:                 sl.String(protocol),
			Cname:                    sl.String(cname),
			HttpPort:                 sl.Int(httpport),
			HttpsPort:                sl.Int(httpsport),
			OriginType:               sl.String(origintype),
			RespectHeaders:           sl.String(rHeader),
			Header:                   sl.String(header),
			UniqueId:                 sl.String(uniqueId),
			CertificateType:          sl.String(certificateType),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		///Print the response of the requested service.
		log.Print("Response for cdn update: ", update1)

		if err != nil {
			log.Println(err)
		}
		return resourceIBMCDNRead(d, meta)
	}

	if origintype == "HOST_SERVER" && protocol == "HTTPS" {
		update2, err := service.UpdateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Protocol:                 sl.String(protocol),
			Cname:                    sl.String(cname),
			HttpsPort:                sl.Int(httpsport),
			OriginType:               sl.String(origintype),
			RespectHeaders:           sl.String(rHeader),
			Header:                   sl.String(header),
			UniqueId:                 sl.String(uniqueId),
			CertificateType:          sl.String(certificateType),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		///Print the response of the requested service.
		log.Print("Response for cdn update: ", update2)
		if err != nil {
			log.Println(err)
		}
		return resourceIBMCDNRead(d, meta)

	}

	if origintype == "HOST_SERVER" && protocol == "HTTP" {
		update3, err := service.UpdateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Protocol:                 sl.String(protocol),
			Cname:                    sl.String(cname),
			HttpPort:                 sl.Int(httpport),
			OriginType:               sl.String(origintype),
			RespectHeaders:           sl.String(rHeader),
			Header:                   sl.String(header),
			UniqueId:                 sl.String(uniqueId),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		///Print the response of the requested service.
		log.Print("Response for cdn update: ", update3)
		if err != nil {
			log.Println(err)
		}
		return resourceIBMCDNRead(d, meta)

	}

	if origintype == "OBJECT_STORAGE" && protocol == "HTTP_AND_HTTPS" {
		update4, err := service.UpdateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Protocol:                 sl.String(protocol),
			Cname:                    sl.String(cname),
			HttpPort:                 sl.Int(httpport),
			HttpsPort:                sl.Int(httpsport),
			OriginType:               sl.String(origintype),
			RespectHeaders:           sl.String(rHeader),
			BucketName:               sl.String(bucketname),
			Header:                   sl.String(header),
			FileExtension:            sl.String(fileextension),
			UniqueId:                 sl.String(uniqueId),
			CertificateType:          sl.String(certificateType),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		///Print the response of the requested service.
		log.Print("Response for cdn update: ", update4)
		if err != nil {
			log.Println(err)
		}
		return resourceIBMCDNRead(d, meta)
	}

	if origintype == "OBJECT_STORAGE" && protocol == "HTTPS" {
		update5, err := service.UpdateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Protocol:                 sl.String(protocol),
			Cname:                    sl.String(cname),
			HttpsPort:                sl.Int(httpsport),
			OriginType:               sl.String(origintype),
			RespectHeaders:           sl.String(rHeader),
			BucketName:               sl.String(bucketname),
			Header:                   sl.String(header),
			FileExtension:            sl.String(fileextension),
			UniqueId:                 sl.String(uniqueId),
			CertificateType:          sl.String(certificateType),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		///Print the response of the requested service.
		log.Print("Response for cdn update: ", update5)
		if err != nil {
			log.Println(err)
		}
		return resourceIBMCDNRead(d, meta)
	}

	if origintype == "OBJECT_STORAGE" && protocol == "HTTP" {
		update6, err := service.UpdateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Protocol:                 sl.String(protocol),
			Cname:                    sl.String(cname),
			HttpPort:                 sl.Int(httpport),
			OriginType:               sl.String(origintype),
			RespectHeaders:           sl.String(rHeader),
			BucketName:               sl.String(bucketname),
			Header:                   sl.String(header),
			FileExtension:            sl.String(fileextension),
			UniqueId:                 sl.String(uniqueId),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		///Print the response of the requested service.
		log.Print("Response for cdn update: ", update6)
		if err != nil {
			log.Println(err)
		}
		return resourceIBMCDNRead(d, meta)
	}

	return nil
}

func resourceIBMCDNDelete(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)

	cdnId := sl.String(d.Id())
	///pass the id to delete the resource.
	delete, err := service.DeleteDomainMapping(cdnId)
	if err != nil {
		log.Println(err)
	}
	///print the delete response
	log.Print("Delete response is : ", delete)
	d.SetId("")
	return nil
}

func resourceIBMCDNExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)
	cdnId := sl.String(d.Id())
	///check if the resource exists with the given id.
	exists, err := service.ListDomainMappingByUniqueId(cdnId)
	///Print the response for exist request.
	log.Print("Exists response is : ", exists)
	if err != nil {
		if apiErr, ok := err.(sl.Error); ok {
			if apiErr.StatusCode == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error retrieving CDN mapping info: %s", err)
	}
	return true, nil
}
