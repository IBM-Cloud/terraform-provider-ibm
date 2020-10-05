package ibm

import (
	"log"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/IBM/networking-go-sdk/zonessettingsv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	cisDomainSettingsDNSSEC                          = "dnssec"
	cisDomainSettingsWAF                             = "waf"
	cisDomainSettingsSSL                             = "ssl"
	cisDomainSettingsUniversalSSL                    = "ssl_universal"
	cisDomainSettingsCertificateStatus               = "certificate_status"
	cisDomainSettingsMinTLSVersion                   = "min_tls_version"
	cisDomainSettingsCNAMEFlattening                 = "cname_flattening"
	cisDomainSettingsOpportunisticEncryption         = "opportunistic_encryption"
	cisDomainSettingsAutomaticHTPSRewrites           = "automatic_https_rewrites"
	cisDomainSettingsAlwaysUseHTTPS                  = "always_use_https"
	cisDomainSettingsIPv6                            = "ipv6"
	cisDomainSettingsBrowserCheck                    = "browser_check"
	cisDomainSettingsHotlinkProtection               = "hotlink_protection"
	cisDomainSettingsHTTP2                           = "http2"
	cisDomainSettingsImageLoadOptimization           = "image_load_optimization"
	cisDomainSettingsImageSizeOptimization           = "image_size_optimization"
	cisDomainSettingsIPGeoLocation                   = "ip_geolocation"
	cisDomainSettingsOriginErrorPagePassThru         = "origin_error_page_pass_thru"
	cisDomainSettingsBrotli                          = "brotli"
	cisDomainSettingsPseudoIPv4                      = "pseudo_ipv4"
	cisDomainSettingsPrefetchPreload                 = "prefetch_preload"
	cisDomainSettingsResponseBuffering               = "response_buffering"
	cisDomainSettingsScriptLoadOptimisation          = "script_load_optimization"
	cisDomainSettingsServerSideExclude               = "server_side_exclude"
	cisDomainSettingsTLSClientAuth                   = "tls_client_auth"
	cisDomainSettingsTrueClientIPHeader              = "true_client_ip_header"
	cisDomainSettingsWebSockets                      = "websockets"
	cisDomainSettingsChallengeTTL                    = "challenge_ttl"
	cisDomainSettingsMinify                          = "minify"
	cisDomainSettingsMinifyCSS                       = "css"
	cisDomainSettingsMinifyHTML                      = "html"
	cisDomainSettingsMinifyJS                        = "js"
	cisDomainSettingsSecurityHeader                  = "security_header"
	cisDomainSettingsSecurityHeaderEnabled           = "enabled"
	cisDomainSettingsSecurityHeaderMaxAge            = "max_age"
	cisDomainSettingsSecurityHeaderIncludeSubdomains = "include_subdomains"
	cisDomainSettingsSecurityHeaderNoSniff           = "nosniff"
	cisDomainSettingsMobileRedirect                  = "mobile_redirect"
	cisDomainSettingsMobileRedirectStatus            = "status"
	cisDomainSettingsMobileRedirectMobileSubdomain   = "mobile_subdomain"
	cisDomainSettingsMobileRedirectStripURI          = "strip_uri"
	cisDomainSettingsMaxUpload                       = "max_upload"
	cisDomainSettingsCipher                          = "cipher"
)

func resourceIBMCISSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Description:      "Associated CIS domain",
				Required:         true,
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisDomainSettingsDNSSEC: {
				Type:         schema.TypeString,
				Description:  "DNS Sec setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"active", "disabled"}),
			},
			cisDomainSettingsWAF: {
				Type:         schema.TypeString,
				Description:  "WAF setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"off", "on"}),
			},
			cisDomainSettingsSSL: {
				Type:         schema.TypeString,
				Description:  "SSL/TLS setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"off", "flexible", "full", "strict", "origin_pull"}),
			},
			cisDomainSettingsCertificateStatus: {
				Type:        schema.TypeString,
				Description: "Certificate status",
				Computed:    true,
				Deprecated:  "This field is deprecated",
			},
			cisDomainSettingsMinTLSVersion: {
				Type:         schema.TypeString,
				Description:  "Minimum version of TLS required",
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"1.1", "1.2", "1.3", "1.4"}),
				Default:      "1.1",
			},
			cisDomainSettingsCNAMEFlattening: {
				Type:         schema.TypeString,
				Description:  "cname_flattening setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"flatten_at_root", "flatten_all", "flatten_none"}),
			},
			cisDomainSettingsOpportunisticEncryption: {
				Type:         schema.TypeString,
				Description:  "opportunistic_encryption setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsAutomaticHTPSRewrites: {
				Type:         schema.TypeString,
				Description:  "automatic_https_rewrites setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsAlwaysUseHTTPS: {
				Type:         schema.TypeString,
				Description:  "always_use_https setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsIPv6: {
				Type:         schema.TypeString,
				Description:  "ipv6 setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsBrowserCheck: {
				Type:         schema.TypeString,
				Description:  "browser_check setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsHotlinkProtection: {
				Type:         schema.TypeString,
				Description:  "hotlink_protection setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsHTTP2: {
				Type:         schema.TypeString,
				Description:  "http2 setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsImageLoadOptimization: {
				Type:         schema.TypeString,
				Description:  "image_load_optimization setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsImageSizeOptimization: {
				Type:         schema.TypeString,
				Description:  "image_size_optimization setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"lossless", "off", "lossy"}),
			},
			cisDomainSettingsIPGeoLocation: {
				Type:         schema.TypeString,
				Description:  "ip_geolocation setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsOriginErrorPagePassThru: {
				Type:         schema.TypeString,
				Description:  "origin_error_page_pass_thru setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsBrotli: {
				Type:         schema.TypeString,
				Description:  "brotli setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsPseudoIPv4: {
				Type:         schema.TypeString,
				Description:  "pseudo_ipv4 setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"overwrite_header", "off", "add_header"}),
			},
			cisDomainSettingsPrefetchPreload: {
				Type:         schema.TypeString,
				Description:  "prefetch_preload setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsResponseBuffering: {
				Type:         schema.TypeString,
				Description:  "response_buffering setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsScriptLoadOptimisation: {
				Type:         schema.TypeString,
				Description:  "script_load_optimization setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsServerSideExclude: {
				Type:         schema.TypeString,
				Description:  "server_side_exclude setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsTLSClientAuth: {
				Type:         schema.TypeString,
				Description:  "tls_client_auth setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsTrueClientIPHeader: {
				Type:         schema.TypeString,
				Description:  "true_client_ip_header setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsWebSockets: {
				Type:         schema.TypeString,
				Description:  "websockets setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			cisDomainSettingsChallengeTTL: {
				Type:        schema.TypeInt,
				Description: "Challenge TTL setting",
				Optional:    true,
				Computed:    true,
				ValidateFunc: validateAllowedIntValue([]int{300, 900, 1800, 2700,
					3600, 7200, 10800, 14400, 28800, 57600, 86400, 604800, 2592000, 31536000}),
			},
			cisDomainSettingsMaxUpload: {
				Type:        schema.TypeInt,
				Description: "Maximum upload",
				Optional:    true,
				Computed:    true,
				ValidateFunc: validateAllowedIntValue([]int{100, 125, 150, 175, 200, 225, 250, 275,
					300, 325, 350, 375, 400, 425, 450, 475, 500}),
			},
			cisDomainSettingsCipher: {
				Type:        schema.TypeSet,
				Description: "Cipher settings",
				Optional:    true,
				Computed:    true,
				Set:         schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validateAllowedStringValue([]string{
						zonessettingsv1.UpdateCiphersOptions_Value_Aes128GcmSha256,
						zonessettingsv1.UpdateCiphersOptions_Value_Aes128Sha,
						zonessettingsv1.UpdateCiphersOptions_Value_Aes128Sha256,
						zonessettingsv1.UpdateCiphersOptions_Value_Aes256GcmSha384,
						zonessettingsv1.UpdateCiphersOptions_Value_Aes256Sha,
						zonessettingsv1.UpdateCiphersOptions_Value_Aes256Sha256,
						zonessettingsv1.UpdateCiphersOptions_Value_DesCbc3Sha,
						zonessettingsv1.UpdateCiphersOptions_Value_EcdheEcdsaAes128GcmSha256,
						zonessettingsv1.UpdateCiphersOptions_Value_EcdheEcdsaAes128Sha,
						zonessettingsv1.UpdateCiphersOptions_Value_EcdheEcdsaAes128Sha256,
						zonessettingsv1.UpdateCiphersOptions_Value_EcdheEcdsaAes256GcmSha384,
						zonessettingsv1.UpdateCiphersOptions_Value_EcdheEcdsaAes256Sha384,
						zonessettingsv1.UpdateCiphersOptions_Value_EcdheEcdsaChacha20Poly1305,
						zonessettingsv1.UpdateCiphersOptions_Value_EcdheRsaAes128GcmSha256,
						zonessettingsv1.UpdateCiphersOptions_Value_EcdheRsaAes128Sha,
						zonessettingsv1.UpdateCiphersOptions_Value_EcdheRsaAes128Sha256,
						zonessettingsv1.UpdateCiphersOptions_Value_EcdheRsaAes256GcmSha384,
						zonessettingsv1.UpdateCiphersOptions_Value_EcdheRsaAes256Sha,
						zonessettingsv1.UpdateCiphersOptions_Value_EcdheRsaAes256Sha384,
						zonessettingsv1.UpdateCiphersOptions_Value_EcdheRsaChacha20Poly1305,
					}),
				},
			},
			cisDomainSettingsMinify: {
				Type:        schema.TypeList,
				Description: "Minify setting",
				Optional:    true,
				Computed:    true,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cisDomainSettingsMinifyCSS: {
							Type:         schema.TypeString,
							Description:  "Minify CSS setting",
							Required:     true,
							ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
						},
						cisDomainSettingsMinifyHTML: {
							Type:         schema.TypeString,
							Description:  "Minify HTML setting",
							Required:     true,
							ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
						},
						cisDomainSettingsMinifyJS: {
							Type:         schema.TypeString,
							Description:  "Minify JS setting",
							Required:     true,
							ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
						},
					},
				},
			},
			cisDomainSettingsSecurityHeader: {
				Type:        schema.TypeList,
				Description: "Security Header Setting",
				Optional:    true,
				Computed:    true,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cisDomainSettingsSecurityHeaderEnabled: {
							Type:        schema.TypeBool,
							Description: "security header enabled/disabled",
							Required:    true,
						},
						cisDomainSettingsSecurityHeaderIncludeSubdomains: {
							Type:        schema.TypeBool,
							Description: "security header subdomain included or not",
							Required:    true,
						},
						cisDomainSettingsSecurityHeaderMaxAge: {
							Type:        schema.TypeInt,
							Description: "security header max age",
							Required:    true,
						},
						cisDomainSettingsSecurityHeaderNoSniff: {
							Type:        schema.TypeBool,
							Description: "security header no sniff",
							Required:    true,
						},
					},
				},
			},
			cisDomainSettingsMobileRedirect: {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cisDomainSettingsMobileRedirectStatus: {
							Type:         schema.TypeString,
							Description:  "mobile redirect status",
							Required:     true,
							ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
						},
						cisDomainSettingsMobileRedirectMobileSubdomain: {
							Type:        schema.TypeString,
							Description: "Mobile redirect subdomain",
							Optional:    true,
							Computed:    true,
						},
						cisDomainSettingsMobileRedirectStripURI: {
							Type:        schema.TypeBool,
							Description: "mobile redirect strip URI",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},

		Create:   resourceCISSettingsCreate,
		Read:     resourceCISSettingsRead,
		Update:   resourceCISSettingsUpdate,
		Delete:   resourceCISSettingsDelete,
		Importer: &schema.ResourceImporter{},
	}
}

var settingsList = []string{
	cisDomainSettingsDNSSEC,
	cisDomainSettingsWAF,
	cisDomainSettingsSSL,
	cisDomainSettingsMinTLSVersion,
	cisDomainSettingsCNAMEFlattening,
	cisDomainSettingsOpportunisticEncryption,
	cisDomainSettingsAutomaticHTPSRewrites,
	cisDomainSettingsAlwaysUseHTTPS,
	cisDomainSettingsIPv6,
	cisDomainSettingsBrowserCheck,
	cisDomainSettingsHotlinkProtection,
	cisDomainSettingsHTTP2,
	cisDomainSettingsImageLoadOptimization,
	cisDomainSettingsImageSizeOptimization,
	cisDomainSettingsIPGeoLocation,
	cisDomainSettingsOriginErrorPagePassThru,
	cisDomainSettingsBrotli,
	cisDomainSettingsPseudoIPv4,
	cisDomainSettingsPrefetchPreload,
	cisDomainSettingsResponseBuffering,
	cisDomainSettingsScriptLoadOptimisation,
	cisDomainSettingsServerSideExclude,
	cisDomainSettingsTLSClientAuth,
	cisDomainSettingsTrueClientIPHeader,
	cisDomainSettingsWebSockets,
	cisDomainSettingsChallengeTTL,
	cisDomainSettingsMinify,
	cisDomainSettingsSecurityHeader,
	cisDomainSettingsMobileRedirect,
	cisDomainSettingsMaxUpload,
	cisDomainSettingsCipher,
}

func resourceCISSettingsCreate(d *schema.ResourceData, meta interface{}) error {
	cisID := d.Get(cisID).(string)
	zoneID := d.Get(cisDomainID).(string)

	d.SetId(convertCisToTfTwoVar(zoneID, cisID))
	return resourceCISSettingsUpdate(d, meta)
}

func resourceCISSettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisDomainSettingsClientSession()
	if err != nil {
		return err
	}

	zoneID, cisID, _ := convertTftoCisTwoVar(d.Id())
	cisClient.Crn = core.StringPtr(cisID)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	for _, item := range settingsList {
		var err error
		var resp interface{}

		switch item {
		case cisDomainSettingsDNSSEC:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateZoneDnssecOptions()
					opt.SetStatus(v.(string))
					_, resp, err = cisClient.UpdateZoneDnssec(opt)
				}
			}
		case cisDomainSettingsWAF:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateWebApplicationFirewallOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateWebApplicationFirewall(opt)
				}
			}
		case cisDomainSettingsSSL:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					cisClient, err := meta.(ClientSession).CisSSLClientSession()
					if err != nil {
						return err
					}
					cisClient.Crn = core.StringPtr(cisID)
					cisClient.ZoneIdentifier = core.StringPtr(zoneID)
					opt := cisClient.NewChangeSslSettingOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.ChangeSslSetting(opt)
				}
			}

		case cisDomainSettingsMinTLSVersion:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateMinTlsVersionOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateMinTlsVersion(opt)
				}
			}
		case cisDomainSettingsCNAMEFlattening:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateZoneCnameFlatteningOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateZoneCnameFlattening(opt)
				}
			}
		case cisDomainSettingsOpportunisticEncryption:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateOpportunisticEncryptionOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateOpportunisticEncryption(opt)
				}
			}
		case cisDomainSettingsAutomaticHTPSRewrites:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateAutomaticHttpsRewritesOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateAutomaticHttpsRewrites(opt)
				}
			}
		case cisDomainSettingsAlwaysUseHTTPS:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateAlwaysUseHttpsOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateAlwaysUseHttps(opt)
				}
			}
		case cisDomainSettingsIPv6:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateIpv6Options()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateIpv6(opt)
				}
			}
		case cisDomainSettingsBrowserCheck:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateBrowserCheckOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateBrowserCheck(opt)
				}
			}
		case cisDomainSettingsHotlinkProtection:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateHotlinkProtectionOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateHotlinkProtection(opt)
				}
			}
		case cisDomainSettingsHTTP2:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateHttp2Options()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateHttp2(opt)
				}
			}
		case cisDomainSettingsImageLoadOptimization:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateImageLoadOptimizationOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateImageLoadOptimization(opt)
				}
			}
		case cisDomainSettingsImageSizeOptimization:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateImageSizeOptimizationOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateImageSizeOptimization(opt)
				}
			}
		case cisDomainSettingsIPGeoLocation:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateIpGeolocationOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateIpGeolocation(opt)
				}
			}
		case cisDomainSettingsOriginErrorPagePassThru:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateEnableErrorPagesOnOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateEnableErrorPagesOn(opt)
				}
			}
		case cisDomainSettingsPseudoIPv4:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdatePseudoIpv4Options()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdatePseudoIpv4(opt)
				}
			}
		case cisDomainSettingsPrefetchPreload:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdatePrefetchPreloadOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdatePrefetchPreload(opt)
				}
			}
		case cisDomainSettingsResponseBuffering:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateResponseBufferingOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateResponseBuffering(opt)
				}
			}
		case cisDomainSettingsScriptLoadOptimisation:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateScriptLoadOptimizationOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateScriptLoadOptimization(opt)
				}
			}
		case cisDomainSettingsServerSideExclude:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateServerSideExcludeOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateServerSideExclude(opt)
				}
			}
		case cisDomainSettingsTLSClientAuth:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateTlsClientAuthOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateTlsClientAuth(opt)
				}
			}
		case cisDomainSettingsTrueClientIPHeader:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateTrueClientIpOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateTrueClientIp(opt)
				}
			}
		case cisDomainSettingsWebSockets:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateWebSocketsOptions()
					opt.SetValue(v.(string))
					_, resp, err = cisClient.UpdateWebSockets(opt)
				}
			}
		case cisDomainSettingsChallengeTTL:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateChallengeTtlOptions()
					opt.SetValue(int64(v.(int)))
					_, resp, err = cisClient.UpdateChallengeTTL(opt)
				}
			}
		case cisDomainSettingsMaxUpload:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					opt := cisClient.NewUpdateMaxUploadOptions()
					opt.SetValue(int64(v.(int)))
					_, resp, err = cisClient.UpdateMaxUpload(opt)
				}
			}
		case cisDomainSettingsCipher:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					cipherValue := expandStringList(v.(*schema.Set).List())
					opt := cisClient.NewUpdateCiphersOptions()
					opt.SetValue(cipherValue)
					_, resp, err = cisClient.UpdateCiphers(opt)
				}
			}
		case cisDomainSettingsMinify:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					dataMap := v.([]interface{})[0].(map[string]interface{})
					css := dataMap[cisDomainSettingsMinifyCSS].(string)
					html := dataMap[cisDomainSettingsMinifyHTML].(string)
					js := dataMap[cisDomainSettingsMinifyJS].(string)
					minifyVal, err := cisClient.NewMinifySettingValue(css, html, js)
					if err != nil {
						log.Println("Invalid minfiy setting values")
						return err
					}
					opt := cisClient.NewUpdateMinifyOptions()
					opt.SetValue(minifyVal)
					_, resp, err = cisClient.UpdateMinify(opt)
				}
			}
		case cisDomainSettingsSecurityHeader:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					dataMap := v.([]interface{})[0].(map[string]interface{})
					enabled := dataMap[cisDomainSettingsSecurityHeaderEnabled].(bool)
					nosniff := dataMap[cisDomainSettingsSecurityHeaderNoSniff].(bool)
					includeSubdomain := dataMap[cisDomainSettingsSecurityHeaderIncludeSubdomains].(bool)
					maxAge := int64(dataMap[cisDomainSettingsSecurityHeaderMaxAge].(int))
					securityVal, err := cisClient.NewSecurityHeaderSettingValueStrictTransportSecurity(
						enabled, maxAge, includeSubdomain, nosniff)
					if err != nil {
						log.Println("Invalid security header setting values")
						return err
					}
					securityOpt, err := cisClient.NewSecurityHeaderSettingValue(securityVal)
					if err != nil {
						log.Println("Invalid security header setting options")
						return err
					}
					opt := cisClient.NewUpdateSecurityHeaderOptions()
					opt.SetValue(securityOpt)
					_, resp, err = cisClient.UpdateSecurityHeader(opt)
				}
			}
		case cisDomainSettingsMobileRedirect:
			if d.HasChange(item) {
				if v, ok := d.GetOk(item); ok {
					dataMap := v.([]interface{})[0].(map[string]interface{})
					status := dataMap[cisDomainSettingsMobileRedirectStatus].(string)
					mobileSubdomain := dataMap[cisDomainSettingsMobileRedirectMobileSubdomain].(string)
					stripURI := dataMap[cisDomainSettingsMobileRedirectStripURI].(bool)
					mobileOpt, err := cisClient.NewMobileRedirecSettingValue(status, mobileSubdomain, stripURI)
					if err != nil {
						log.Println("Invalid mobile redirect options")
						return err
					}
					opt := cisClient.NewUpdateMobileRedirectOptions()
					opt.SetValue(mobileOpt)
					_, resp, err = cisClient.UpdateMobileRedirect(opt)
				}
			}
		}
		if err != nil {
			log.Printf("Update settings Failed on %s, %v\n", item, resp)
			return err
		}
	}

	return resourceCISSettingsRead(d, meta)
}

func resourceCISSettingsRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisDomainSettingsClientSession()
	if err != nil {
		return err
	}

	zoneID, crn, _ := convertTftoCisTwoVar(d.Id())
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	for _, item := range settingsList {

		switch item {
		case cisDomainSettingsDNSSEC:
			opt := cisClient.NewGetZoneDnssecOptions()
			result, resp, err := cisClient.GetZoneDnssec(opt)
			if err != nil {
				log.Printf("DNS SEC get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsDNSSEC, result.Result.Status)

		case cisDomainSettingsWAF:
			opt := cisClient.NewGetWebApplicationFirewallOptions()
			result, resp, err := cisClient.GetWebApplicationFirewall(opt)
			if err != nil {
				log.Printf("Waf get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsWAF, result.Result.Value)

		case cisDomainSettingsSSL:
			cisClient, err := meta.(ClientSession).CisSSLClientSession()
			if err != nil {
				return err
			}
			cisClient.Crn = core.StringPtr(crn)
			cisClient.ZoneIdentifier = core.StringPtr(zoneID)
			opt := cisClient.NewGetSslSettingOptions()
			result, resp, err := cisClient.GetSslSetting(opt)
			if err != nil {
				log.Printf("SSL setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsSSL, result.Result.Value)

		case cisDomainSettingsBrotli:
			cisClient, err := meta.(ClientSession).CisAPI()
			if err != nil {
				return err
			}
			settingsResult, err := cisClient.Settings().GetSetting(crn, zoneID, item)
			if err != nil {
				log.Printf("[WARN] Error getting %s zone during Domain read %v\n", item, err)
				return err
			}
			settingsObj := *settingsResult
			d.Set(item, settingsObj.Value)

		case cisDomainSettingsMinTLSVersion:
			opt := cisClient.NewGetMinTlsVersionOptions()
			result, resp, err := cisClient.GetMinTlsVersion(opt)
			if err != nil {
				log.Printf("Min TLS Version setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsMinTLSVersion, result.Result.Value)

		case cisDomainSettingsCNAMEFlattening:
			opt := cisClient.NewGetZoneCnameFlatteningOptions()
			result, resp, err := cisClient.GetZoneCnameFlattening(opt)
			if err != nil {
				log.Printf("CNAME Flattening setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsCNAMEFlattening, result.Result.Value)

		case cisDomainSettingsOpportunisticEncryption:
			opt := cisClient.NewGetOpportunisticEncryptionOptions()
			result, resp, err := cisClient.GetOpportunisticEncryption(opt)
			if err != nil {
				log.Printf("Opportunistic Encryption setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsOpportunisticEncryption, result.Result.Value)

		case cisDomainSettingsAutomaticHTPSRewrites:
			opt := cisClient.NewGetAutomaticHttpsRewritesOptions()
			result, resp, err := cisClient.GetAutomaticHttpsRewrites(opt)
			if err != nil {
				log.Printf("Automatic HTTPS Rewrites setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsAutomaticHTPSRewrites, result.Result.Value)

		case cisDomainSettingsAlwaysUseHTTPS:
			opt := cisClient.NewGetAlwaysUseHttpsOptions()
			result, resp, err := cisClient.GetAlwaysUseHttps(opt)
			if err != nil {
				log.Printf("Always use HTTPS setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsAlwaysUseHTTPS, result.Result.Value)

		case cisDomainSettingsIPv6:
			opt := cisClient.NewGetIpv6Options()
			result, resp, err := cisClient.GetIpv6(opt)
			if err != nil {
				log.Printf("IPv6 setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsIPv6, result.Result.Value)

		case cisDomainSettingsBrowserCheck:
			opt := cisClient.NewGetBrowserCheckOptions()
			result, resp, err := cisClient.GetBrowserCheck(opt)
			if err != nil {
				log.Printf("Browser Check setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsBrowserCheck, result.Result.Value)

		case cisDomainSettingsHotlinkProtection:
			opt := cisClient.NewGetHotlinkProtectionOptions()
			result, resp, err := cisClient.GetHotlinkProtection(opt)
			if err != nil {
				log.Printf("Hotlink protection setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsHotlinkProtection, result.Result.Value)

		case cisDomainSettingsHTTP2:
			opt := cisClient.NewGetHttp2Options()
			result, resp, err := cisClient.GetHttp2(opt)
			if err != nil {
				log.Printf("HTTP2 setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsHTTP2, result.Result.Value)

		case cisDomainSettingsImageLoadOptimization:
			opt := cisClient.NewGetImageLoadOptimizationOptions()
			result, resp, err := cisClient.GetImageLoadOptimization(opt)
			if err != nil {
				log.Printf("Image load optimization setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsImageLoadOptimization, result.Result.Value)

		case cisDomainSettingsImageSizeOptimization:
			opt := cisClient.NewGetImageSizeOptimizationOptions()
			result, resp, err := cisClient.GetImageSizeOptimization(opt)
			if err != nil {
				log.Printf("Image size optimization setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsImageSizeOptimization, result.Result.Value)

		case cisDomainSettingsIPGeoLocation:
			opt := cisClient.NewGetIpGeolocationOptions()
			result, resp, err := cisClient.GetIpGeolocation(opt)
			if err != nil {
				log.Printf("IP Geo location setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsIPGeoLocation, result.Result.Value)

		case cisDomainSettingsOriginErrorPagePassThru:
			opt := cisClient.NewGetEnableErrorPagesOnOptions()
			result, resp, err := cisClient.GetEnableErrorPagesOn(opt)
			if err != nil {
				log.Printf("Origin error page pass thru setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsOriginErrorPagePassThru, result.Result.Value)

		case cisDomainSettingsPseudoIPv4:
			opt := cisClient.NewGetPseudoIpv4Options()
			result, resp, err := cisClient.GetPseudoIpv4(opt)
			if err != nil {
				log.Printf("Pseudo IPv4 setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsPseudoIPv4, result.Result.Value)

		case cisDomainSettingsPrefetchPreload:
			opt := cisClient.NewGetPrefetchPreloadOptions()
			result, resp, err := cisClient.GetPrefetchPreload(opt)
			if err != nil {
				log.Printf("Prefetch preload setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsPrefetchPreload, result.Result.Value)

		case cisDomainSettingsResponseBuffering:
			opt := cisClient.NewGetResponseBufferingOptions()
			result, resp, err := cisClient.GetResponseBuffering(opt)
			if err != nil {
				log.Printf("Response buffering setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsResponseBuffering, result.Result.Value)

		case cisDomainSettingsScriptLoadOptimisation:
			opt := cisClient.NewGetScriptLoadOptimizationOptions()
			result, resp, err := cisClient.GetScriptLoadOptimization(opt)
			if err != nil {
				log.Printf("Script load optimisation setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsScriptLoadOptimisation, result.Result.Value)

		case cisDomainSettingsServerSideExclude:
			opt := cisClient.NewGetServerSideExcludeOptions()
			result, resp, err := cisClient.GetServerSideExclude(opt)
			if err != nil {
				log.Printf("Service side exclude setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsServerSideExclude, result.Result.Value)

		case cisDomainSettingsTLSClientAuth:
			opt := cisClient.NewGetTlsClientAuthOptions()
			result, resp, err := cisClient.GetTlsClientAuth(opt)
			if err != nil {
				log.Printf("TLS Client Auth setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsTLSClientAuth, result.Result.Value)

		case cisDomainSettingsTrueClientIPHeader:
			opt := cisClient.NewGetTrueClientIpOptions()
			result, resp, err := cisClient.GetTrueClientIp(opt)
			if err != nil {
				log.Printf("True Client IP Header setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsTrueClientIPHeader, result.Result.Value)

		case cisDomainSettingsWebSockets:
			opt := cisClient.NewGetWebSocketsOptions()
			result, resp, err := cisClient.GetWebSockets(opt)
			if err != nil {
				log.Printf("Get websockets setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsWebSockets, result.Result.Value)

		case cisDomainSettingsChallengeTTL:
			opt := cisClient.NewGetChallengeTtlOptions()
			result, resp, err := cisClient.GetChallengeTTL(opt)
			if err != nil {
				log.Printf("Challenge TTL setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsChallengeTTL, result.Result.Value)

		case cisDomainSettingsMaxUpload:
			opt := cisClient.NewGetMaxUploadOptions()
			result, resp, err := cisClient.GetMaxUpload(opt)
			if err != nil {
				log.Printf("Max upload setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsMaxUpload, result.Result.Value)

		case cisDomainSettingsCipher:
			opt := cisClient.NewGetCiphersOptions()
			result, resp, err := cisClient.GetCiphers(opt)
			if err != nil {
				log.Printf("Cipher setting get request failed : %v", resp)
				return err
			}
			d.Set(cisDomainSettingsCipher, result.Result.Value)

		case cisDomainSettingsMinify:
			opt := cisClient.NewGetMinifyOptions()
			result, resp, err := cisClient.GetMinify(opt)
			if err != nil {
				log.Printf("Minify setting get request failed : %v", resp)
				return err
			}
			minify := result.Result.Value
			value := map[string]string{
				cisDomainSettingsMinifyCSS:  *minify.Css,
				cisDomainSettingsMinifyHTML: *minify.HTML,
				cisDomainSettingsMinifyJS:   *minify.Js,
			}
			d.Set(cisDomainSettingsMinify, []interface{}{value})

		case cisDomainSettingsSecurityHeader:
			opt := cisClient.NewGetSecurityHeaderOptions()
			result, resp, err := cisClient.GetSecurityHeader(opt)
			if err != nil {
				log.Printf("Security header setting get request failed : %v", resp)
				return err
			}
			securityHeader := result.Result.Value.StrictTransportSecurity
			value := map[string]interface{}{
				cisDomainSettingsSecurityHeaderEnabled:           *securityHeader.Enabled,
				cisDomainSettingsSecurityHeaderNoSniff:           *securityHeader.Nosniff,
				cisDomainSettingsSecurityHeaderIncludeSubdomains: *securityHeader.IncludeSubdomains,
				cisDomainSettingsSecurityHeaderMaxAge:            *securityHeader.MaxAge,
			}
			d.Set(cisDomainSettingsSecurityHeader, []interface{}{value})

		case cisDomainSettingsMobileRedirect:
			opt := cisClient.NewGetMobileRedirectOptions()
			result, resp, err := cisClient.GetMobileRedirect(opt)
			if err != nil {
				log.Printf("Mobile redirect strip URI setting get request failed : %v", resp)
				return err
			}
			value := result.Result.Value
			uri := map[string]interface{}{
				cisDomainSettingsMobileRedirectMobileSubdomain: *value.MobileSubdomain,
				cisDomainSettingsMobileRedirectStatus:          *value.Status,
				cisDomainSettingsMobileRedirectStripURI:        *value.StripURI,
			}
			d.Set(cisDomainSettingsMobileRedirect, []interface{}{uri})
		}
	}
	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)
	return nil
}

func resourceCISSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	// Nothing to delete on CIS resource
	d.SetId("")
	return nil
}
