// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isNetworkACLName              = "name"
	isNetworkACLRules             = "rules"
	isNetworkACLSubnets           = "subnets"
	isNetworkACLRuleID            = "id"
	isNetworkACLRuleName          = "name"
	isNetworkACLRuleAction        = "action"
	isNetworkACLRuleIPVersion     = "ip_version"
	isNetworkACLRuleSource        = "source"
	isNetworkACLRuleDestination   = "destination"
	isNetworkACLRuleDirection     = "direction"
	isNetworkACLRuleProtocol      = "protocol"
	isNetworkACLRuleICMP          = "icmp"
	isNetworkACLRuleICMPCode      = "code"
	isNetworkACLRuleICMPType      = "type"
	isNetworkACLRuleTCP           = "tcp"
	isNetworkACLRuleUDP           = "udp"
	isNetworkACLRulePortMax       = "port_max"
	isNetworkACLRulePortMin       = "port_min"
	isNetworkACLRuleSourcePortMax = "source_port_max"
	isNetworkACLRuleSourcePortMin = "source_port_min"
	isNetworkACLVPC               = "vpc"
	isNetworkACLResourceGroup     = "resource_group"
	isNetworkACLTags              = "tags"
	isNetworkACLAccessTags        = "access_tags"
	isNetworkACLCRN               = "crn"
	isNetworkACLRuleUpdateMode    = "surgical_rule_update"
)

func ResourceIBMISNetworkACL() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISNetworkACLCreate,
		ReadContext:   resourceIBMISNetworkACLRead,
		UpdateContext: resourceIBMISNetworkACLUpdate,
		DeleteContext: resourceIBMISNetworkACLDelete,
		Exists:        resourceIBMISNetworkACLExists,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceTagsCustomizeDiff(diff)
				},
			),
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceValidateAccessTags(diff, v)
				}),
		),

		Schema: map[string]*schema.Schema{
			isNetworkACLName: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLName),
				Description:  "Network ACL name",
			},
			isNetworkACLVPC: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "Network ACL VPC name",
			},
			isNetworkACLResourceGroup: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "Resource group ID for the network ACL",
			},
			isNetworkACLTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of tags",
			},

			isNetworkACLAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", "accesstag")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
			},

			isNetworkACLCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},
			flex.ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},
			flex.ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			flex.ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
			isNetworkACLRuleUpdateMode: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "When set to true, enables surgical inline rule updates (add, remove, reorder, patch, recreate only changed rules). When false (default), any change to inline rules deletes all existing rules and recreates them from the configuration.",
			},
			isNetworkACLRules: {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isNetworkACLRuleID: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isNetworkACLRuleName: {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     false,
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRuleName),
						},
						isNetworkACLRuleAction: {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     false,
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRuleAction),
						},
						isNetworkACLRuleIPVersion: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isNetworkACLRuleSource: {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     false,
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRuleSource),
						},
						isNetworkACLRuleDestination: {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     false,
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRuleDestination),
						},
						isNetworkACLRuleDirection: {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     false,
							Description:  "Direction of traffic to enforce, either inbound or outbound",
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRuleDirection),
						},
						isNetworkACLSubnets: {
							Type:     schema.TypeInt,
							Computed: true,
						},
						isNetworkACLRuleProtocol: {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							Description:  "The name of the network protocol",
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRuleProtocol),
						},
						isNetworkACLRuleICMPCode: {
							Type:         schema.TypeInt,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRuleICMPCode),
							Description:  "The ICMP traffic code to allow. Valid values from 0 to 255.",
						},
						isNetworkACLRuleICMPType: {
							Type:         schema.TypeInt,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRuleICMPType),
							Description:  "The ICMP traffic type to allow. Valid values from 0 to 254.",
						},
						isNetworkACLRulePortMax: {
							Type:             schema.TypeInt,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: suppressNullValues,
							ValidateFunc:     validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRulePortMax),
							Description:      "The highest port in the range of ports to be matched",
						},
						isNetworkACLRulePortMin: {
							Type:             schema.TypeInt,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: suppressNullValues,
							ValidateFunc:     validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRulePortMin),
							Description:      "The lowest port in the range of ports to be matched",
						},
						isNetworkACLRuleSourcePortMax: {
							Type:             schema.TypeInt,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: suppressNullValues,
							ValidateFunc:     validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRuleSourcePortMax),
							Description:      "The highest port in the range of ports to be matched",
						},
						isNetworkACLRuleSourcePortMin: {
							Type:             schema.TypeInt,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: suppressNullValues,
							ValidateFunc:     validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRuleSourcePortMin),
							Description:      "The lowest port in the range of ports to be matched",
						},
						isNetworkACLRuleICMP: {
							Type:       schema.TypeList,
							MinItems:   0,
							MaxItems:   1,
							Optional:   true,
							Computed:   true,
							Deprecated: "icmp is deprecated, use 'protocol', 'code', and 'type' instead.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isNetworkACLRuleICMPCode: {
										Type:         schema.TypeInt,
										Optional:     true,
										ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRuleICMPCode),
									},
									isNetworkACLRuleICMPType: {
										Type:         schema.TypeInt,
										Optional:     true,
										ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRuleICMPType),
									},
								},
							},
						},

						isNetworkACLRuleTCP: {
							Type:       schema.TypeList,
							MinItems:   0,
							MaxItems:   1,
							Optional:   true,
							Computed:   true,
							Deprecated: "tcp is deprecated, use 'protocol', 'port_min', 'port_max', 'source_port_min', and 'source_port_max' instead.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isNetworkACLRulePortMax: {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      65535,
										ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRulePortMax),
									},
									isNetworkACLRulePortMin: {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      1,
										ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRulePortMin),
									},
									isNetworkACLRuleSourcePortMax: {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      65535,
										ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRuleSourcePortMax),
									},
									isNetworkACLRuleSourcePortMin: {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      1,
										ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRuleSourcePortMin),
									},
								},
							},
						},

						isNetworkACLRuleUDP: {
							Type:       schema.TypeList,
							MinItems:   0,
							MaxItems:   1,
							Optional:   true,
							Computed:   true,
							Deprecated: "udp is deprecated, use 'protocol', 'port_min', 'port_max', 'source_port_min', and 'source_port_max' instead.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isNetworkACLRulePortMax: {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      65535,
										ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRulePortMax),
									},
									isNetworkACLRulePortMin: {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      1,
										ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRulePortMin),
									},
									isNetworkACLRuleSourcePortMax: {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      65535,
										ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRuleSourcePortMax),
									},
									isNetworkACLRuleSourcePortMin: {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      1,
										ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRuleSourcePortMin),
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

func suppressNullValues(k, old, new string, d *schema.ResourceData) bool {
	parts := strings.Split(k, ".")
	if len(parts) < 3 {
		return false
	}

	// Build the path to the protocol field
	ruleIndex := parts[1]
	protocolKey := fmt.Sprintf("rules.%s.protocol", ruleIndex)

	// Get the protocol value
	protocol, ok := d.GetOk(protocolKey)
	if !ok {
		return false
	}

	protocolStr := protocol.(string)

	// Only suppress for TCP or UDP protocols
	if protocolStr != "tcp" && protocolStr != "udp" {
		return false
	}

	// When TypeInt field is null, it comes as "0"
	// Suppress if new is "0" and old was a positive value
	if new == "0" && old != "" && old != "0" && d.Id() != "" {
		return true
	}
	return false
}

func ResourceIBMISNetworkACLValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	direction := "inbound, outbound"
	action := "allow, deny"
	protocol := "tcp, udp, icmp, ah, any, esp, gre, icmp_tcp_udp, ip_in_ip, l2tp, number_10, number_100, number_101, number_102, number_103, number_104, number_105, number_106, number_107, number_108, number_109, number_11, number_110, number_111, number_113, number_114, number_116, number_117, number_118, number_119, number_12, number_120, number_121, number_122, number_123, number_124, number_125, number_126, number_127, number_128, number_129, number_13, number_130, number_131, number_133, number_134, number_136, number_137, number_138, number_139, number_14, number_140, number_141, number_142, number_143, number_144, number_145, number_146, number_147, number_148, number_149, number_15, number_150, number_151, number_152, number_153, number_154, number_155, number_156, number_157, number_158, number_159, number_16, number_160, number_161, number_162, number_163, number_164, number_165, number_166, number_167, number_168, number_169, number_170, number_171, number_172, number_173, number_174, number_175, number_176, number_177, number_178, number_179, number_18, number_180, number_181, number_182, number_183, number_184, number_185, number_186, number_187, number_188, number_189, number_19, number_190, number_191, number_192, number_193, number_194, number_195, number_196, number_197, number_198, number_199, number_2, number_20, number_200, number_201, number_202, number_203, number_204, number_205, number_206, number_207, number_208, number_209, number_21, number_210, number_211, number_212, number_213, number_214, number_215, number_216, number_217, number_218, number_219, number_22, number_220, number_221, number_222, number_223, number_224, number_225, number_226, number_227, number_228, number_229, number_23, number_230, number_231, number_232, number_233, number_234, number_235, number_236, number_237, number_238, number_239, number_24, number_240, number_241, number_242, number_243, number_244, number_245, number_246, number_247, number_248, number_249, number_25, number_250, number_251, number_252, number_253, number_254, number_255, number_26, number_27, number_28, number_29, number_3, number_30, number_31, number_32, number_33, number_34, number_35, number_36, number_37, number_38, number_39, number_40, number_41, number_42, number_45, number_48, number_49, number_5, number_52, number_53, number_54, number_55, number_56, number_57, number_61, number_62, number_63, number_64, number_65, number_66, number_67, number_68, number_69, number_7, number_70, number_71, number_72, number_73, number_74, number_75, number_76, number_77, number_78, number_79, number_8, number_80, number_81, number_82, number_83, number_84, number_85, number_86, number_87, number_88, number_89, number_9, number_90, number_91, number_92, number_93, number_94, number_95, number_96, number_97, number_98, number_99, rsvp, sctp, vrrp"
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isNetworkACLRuleAction,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              action})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isNetworkACLRuleDirection,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              direction})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isNetworkACLName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   false,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isNetworkACLRuleName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isNetworkACLRuleDestination,
			ValidateFunctionIdentifier: validate.ValidateIPorCIDR,
			Type:                       validate.TypeString,
			Required:                   true})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isNetworkACLRuleSource,
			ValidateFunctionIdentifier: validate.ValidateIPorCIDR,
			Type:                       validate.TypeString,
			Required:                   true})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isNetworkACLRuleICMPType,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "0",
			MaxValue:                   "254"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isNetworkACLRuleICMPCode,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "0",
			MaxValue:                   "255"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isNetworkACLRulePortMin,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "1",
			MaxValue:                   "65535"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isNetworkACLRulePortMax,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "1",
			MaxValue:                   "65535"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isNetworkACLRuleSourcePortMin,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "1",
			MaxValue:                   "65535"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isNetworkACLRuleSourcePortMax,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "1",
			MaxValue:                   "65535"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "accesstag",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$`,
			MinValueLength:             1,
			MaxValueLength:             128})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isNetworkACLRuleProtocol,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			AllowedValues:              protocol})

	ibmISNetworkACLResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_network_acl", Schema: validateSchema}
	return &ibmISNetworkACLResourceValidator
}

func resourceIBMISNetworkACLCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	name := d.Get(isNetworkACLName).(string)
	err := nwaclCreate(context, d, meta, name)
	if err != nil {
		return err
	}
	return resourceIBMISNetworkACLRead(context, d, meta)

}

func nwaclCreate(context context.Context, d *schema.ResourceData, meta interface{}, name string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	var vpc, rg string
	if vpcID, ok := d.GetOk(isNetworkACLVPC); ok {
		vpc = vpcID.(string)
	} else {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl", "create", "parse-vpc").GetDiag()
	}

	nwaclTemplate := &vpcv1.NetworkACLPrototype{
		VPC: &vpcv1.VPCIdentity{
			ID: &vpc,
		},
	}
	if name != "" {
		nwaclTemplate.Name = &name
	}

	if grp, ok := d.GetOk(isNetworkACLResourceGroup); ok {
		rg = grp.(string)
		nwaclTemplate.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}
	// validate each rule before attempting to create the ACL
	var rules []interface{}
	if rls, ok := d.GetOk(isNetworkACLRules); ok {
		rules = rls.([]interface{})
	}
	err = validateInlineRules(d, rules)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl", "create", "validate-inline-rules").GetDiag()
	}

	options := &vpcv1.CreateNetworkACLOptions{
		NetworkACLPrototype: nwaclTemplate,
	}

	nwacl, _, err := sess.CreateNetworkACLWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateNetworkACLWithContext failed: %s", err.Error()), "ibm_is_network_acl", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*nwacl.ID)
	log.Printf("[INFO] Network ACL : %s", *nwacl.ID)
	nwaclid := *nwacl.ID

	//Remove default rules
	err = clearRules(sess, nwaclid)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("clearRules failed: %s", err.Error()), "ibm_is_network_acl", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	err = createInlineRules(d, sess, nwaclid, rules)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("createInlineRules failed: %s", err.Error()), "ibm_is_network_acl", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isNetworkACLTags); ok || v != "" {
		oldList, newList := d.GetChange(isNetworkACLTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *nwacl.CRN, "", isUserTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource network acl (%s) tags: %s", d.Id(), err)
		}
	}
	if _, ok := d.GetOk(isNetworkACLAccessTags); ok {
		oldList, newList := d.GetChange(isNetworkACLAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *nwacl.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource network acl (%s) access tags: %s", d.Id(), err)
		}
	}
	return nil
}

func resourceIBMISNetworkACLRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	err := nwaclGet(context, d, meta, id)
	if err != nil {
		return err
	}
	return nil
}

func nwaclGet(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getNetworkAclOptions := &vpcv1.GetNetworkACLOptions{
		ID: &id,
	}
	nwacl, response, err := sess.GetNetworkACLWithContext(context, getNetworkAclOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetNetworkACLWithContext failed: %s", err.Error()), "ibm_is_network_acl", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if !core.IsNil(nwacl.Name) {
		if err = d.Set("name", nwacl.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl", "read", "set-name").GetDiag()
		}
	}

	if !core.IsNil(nwacl.VPC) {
		if err = d.Set(isNetworkACLVPC, *nwacl.VPC.ID); err != nil {
			err = fmt.Errorf("Error setting vpc: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl", "read", "set-vpc").GetDiag()
		}
	}
	if nwacl.ResourceGroup != nil {
		if err = d.Set(isNetworkACLResourceGroup, *nwacl.ResourceGroup.ID); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl", "read", "set-resource_group").GetDiag()
		}
		if err = d.Set(flex.ResourceGroupName, *nwacl.ResourceGroup.Name); err != nil {
			err = fmt.Errorf("Error setting resource_group_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl", "read", "set-resource_group_name").GetDiag()
		}
	}
	tags, err := flex.GetGlobalTagsUsingCRN(meta, *nwacl.CRN, "", isUserTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource network acl (%s) tags: %s", d.Id(), err)
	}

	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *nwacl.CRN, "", isAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource network acl (%s) access tags: %s", d.Id(), err)
	}

	if err = d.Set(isNetworkACLTags, tags); err != nil {
		err = fmt.Errorf("Error setting tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl", "read", "set-tags").GetDiag()
	}
	if err = d.Set(isNetworkACLAccessTags, accesstags); err != nil {
		err = fmt.Errorf("Error setting access_tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl", "read", "set-access_tags").GetDiag()
	}
	if err = d.Set("crn", nwacl.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl", "read", "set-crn").GetDiag()
	}
	rules := make([]interface{}, 0)
	if len(nwacl.Rules) > 0 {
		for index, rulex := range nwacl.Rules {
			log.Println("[DEBUG] Type of the Rule", reflect.TypeOf(rulex))
			rule := make(map[string]interface{})
			rule[isNetworkACLSubnets] = len(nwacl.Subnets)
			switch reflect.TypeOf(rulex).String() {
			case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIcmp":
				{
					rulex := rulex.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIcmp)
					rule[isNetworkACLRuleID] = *rulex.ID
					rule[isNetworkACLRuleName] = *rulex.Name
					rule[isNetworkACLRuleAction] = *rulex.Action
					rule[isNetworkACLRuleIPVersion] = *rulex.IPVersion
					rule[isNetworkACLRuleProtocol] = *rulex.Protocol
					rule[isNetworkACLRuleSource] = *rulex.Source
					rule[isNetworkACLRuleDestination] = *rulex.Destination
					rule[isNetworkACLRuleDirection] = *rulex.Direction

					// Always populate new design fields
					if rulex.Code != nil {
						rule[isNetworkACLRuleICMPCode] = int(*rulex.Code)
					}
					if rulex.Type != nil {
						rule[isNetworkACLRuleICMPType] = int(*rulex.Type)
					}

					// Only populate deprecated icmp block if user was using old-style
					icmpPath := fmt.Sprintf("rules.%d.icmp", index)
					usingDeprecatedIcmp := false
					if _, ok := d.GetOk(icmpPath); ok {
						usingDeprecatedIcmp = true
					}
					if usingDeprecatedIcmp {
						icmpProtocol := map[string]int{}
						if rulex.Code != nil {
							icmpProtocol[isNetworkACLRuleICMPCode] = int(*rulex.Code)
						}
						if rulex.Type != nil {
							icmpProtocol[isNetworkACLRuleICMPType] = int(*rulex.Type)
						}
						protocolList := make([]map[string]int, 1, 1)
						if len(icmpProtocol) > 0 {
							protocolList[0] = icmpProtocol
						}
						rule[isNetworkACLRuleICMP] = protocolList
					} else {
						rule[isNetworkACLRuleICMP] = make([]map[string]int, 0, 0)
					}
					rule[isNetworkACLRuleTCP] = make([]map[string]int, 0, 0)
					rule[isNetworkACLRuleUDP] = make([]map[string]int, 0, 0)
				}
			case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTcpudp":
				{
					rulex := rulex.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTcpudp)
					rule[isNetworkACLRuleID] = *rulex.ID
					rule[isNetworkACLRuleName] = *rulex.Name
					rule[isNetworkACLRuleAction] = *rulex.Action
					rule[isNetworkACLRuleIPVersion] = *rulex.IPVersion
					rule[isNetworkACLRuleProtocol] = *rulex.Protocol
					rule[isNetworkACLRuleSource] = *rulex.Source
					rule[isNetworkACLRuleDestination] = *rulex.Destination
					rule[isNetworkACLRuleDirection] = *rulex.Direction

					// Always populate new design fields
					rule[isNetworkACLRuleSourcePortMax] = checkNetworkACLNil(rulex.SourcePortMax)
					rule[isNetworkACLRuleSourcePortMin] = checkNetworkACLNil(rulex.SourcePortMin)
					rule[isNetworkACLRulePortMax] = checkNetworkACLNil(rulex.DestinationPortMax)
					rule[isNetworkACLRulePortMin] = checkNetworkACLNil(rulex.DestinationPortMin)

					// Only populate deprecated tcp/udp blocks if user was using old-style
					tcpPath := fmt.Sprintf("rules.%d.tcp", index)
					udpPath := fmt.Sprintf("rules.%d.udp", index)
					usingDeprecatedBlock := false
					if v, ok := d.GetOk(tcpPath); ok {
						if tcpList, ok := v.([]interface{}); ok && len(tcpList) > 0 {
							usingDeprecatedBlock = true
						}
					}
					if v, ok := d.GetOk(udpPath); ok {
						if udpList, ok := v.([]interface{}); ok && len(udpList) > 0 {
							usingDeprecatedBlock = true
						}
					}

					if usingDeprecatedBlock {
						tcpudpProtocol := map[string]int{}
						if rulex.SourcePortMax != nil {
							tcpudpProtocol[isNetworkACLRuleSourcePortMax] = checkNetworkACLNil(rulex.SourcePortMax)
						}
						if rulex.SourcePortMin != nil {
							tcpudpProtocol[isNetworkACLRuleSourcePortMin] = checkNetworkACLNil(rulex.SourcePortMin)
						}
						if rulex.DestinationPortMax != nil {
							tcpudpProtocol[isNetworkACLRulePortMax] = checkNetworkACLNil(rulex.DestinationPortMax)
						}
						if rulex.DestinationPortMin != nil {
							tcpudpProtocol[isNetworkACLRulePortMin] = checkNetworkACLNil(rulex.DestinationPortMin)
						}
						protocolList := make([]map[string]int, 0)
						if len(tcpudpProtocol) > 0 {
							protocolList = append(protocolList, tcpudpProtocol)
						}
						if *rulex.Protocol == "tcp" {
							rule[isNetworkACLRuleICMP] = make([]map[string]int, 0, 0)
							rule[isNetworkACLRuleUDP] = make([]map[string]int, 0, 0)
							rule[isNetworkACLRuleTCP] = protocolList
						} else if *rulex.Protocol == "udp" {
							rule[isNetworkACLRuleICMP] = make([]map[string]int, 0, 0)
							rule[isNetworkACLRuleTCP] = make([]map[string]int, 0, 0)
							rule[isNetworkACLRuleUDP] = protocolList
						}
					} else {
						rule[isNetworkACLRuleICMP] = make([]map[string]int, 0, 0)
						rule[isNetworkACLRuleTCP] = make([]map[string]int, 0, 0)
						rule[isNetworkACLRuleUDP] = make([]map[string]int, 0, 0)
					}
				}
			case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolAny":
				{
					rulex := rulex.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolAny)
					rule[isNetworkACLRuleID] = *rulex.ID
					rule[isNetworkACLRuleName] = *rulex.Name
					rule[isNetworkACLRuleAction] = *rulex.Action
					rule[isNetworkACLRuleIPVersion] = *rulex.IPVersion
					rule[isNetworkACLRuleProtocol] = *rulex.Protocol
					rule[isNetworkACLRuleSource] = *rulex.Source
					rule[isNetworkACLRuleDestination] = *rulex.Destination
					rule[isNetworkACLRuleDirection] = *rulex.Direction
					rule[isNetworkACLRuleICMP] = make([]map[string]int, 0, 0)
					rule[isNetworkACLRuleTCP] = make([]map[string]int, 0, 0)
					rule[isNetworkACLRuleUDP] = make([]map[string]int, 0, 0)
				}
			case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIndividual":
				{
					rulex := rulex.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIndividual)
					rule[isNetworkACLRuleID] = *rulex.ID
					rule[isNetworkACLRuleName] = *rulex.Name
					rule[isNetworkACLRuleAction] = *rulex.Action
					rule[isNetworkACLRuleIPVersion] = *rulex.IPVersion
					rule[isNetworkACLRuleProtocol] = *rulex.Protocol
					rule[isNetworkACLRuleSource] = *rulex.Source
					rule[isNetworkACLRuleDestination] = *rulex.Destination
					rule[isNetworkACLRuleDirection] = *rulex.Direction
					rule[isNetworkACLRuleICMP] = make([]map[string]int, 0, 0)
					rule[isNetworkACLRuleTCP] = make([]map[string]int, 0, 0)
					rule[isNetworkACLRuleUDP] = make([]map[string]int, 0, 0)
				}
			case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIcmptcpudp":
				{
					rulex := rulex.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIcmptcpudp)
					rule[isNetworkACLRuleID] = *rulex.ID
					rule[isNetworkACLRuleName] = *rulex.Name
					rule[isNetworkACLRuleAction] = *rulex.Action
					rule[isNetworkACLRuleIPVersion] = *rulex.IPVersion
					rule[isNetworkACLRuleProtocol] = *rulex.Protocol
					rule[isNetworkACLRuleSource] = *rulex.Source
					rule[isNetworkACLRuleDestination] = *rulex.Destination
					rule[isNetworkACLRuleDirection] = *rulex.Direction
					rule[isNetworkACLRuleICMP] = make([]map[string]int, 0, 0)
					rule[isNetworkACLRuleTCP] = make([]map[string]int, 0, 0)
					rule[isNetworkACLRuleUDP] = make([]map[string]int, 0, 0)
				}
			case "*vpcv1.NetworkACLRuleItem":
				{
					rulex := rulex.(*vpcv1.NetworkACLRuleItem)
					rule[isNetworkACLRuleID] = *rulex.ID
					rule[isNetworkACLRuleName] = *rulex.Name
					rule[isNetworkACLRuleAction] = *rulex.Action
					rule[isNetworkACLRuleIPVersion] = *rulex.IPVersion
					rule[isNetworkACLRuleSource] = *rulex.Source
					rule[isNetworkACLRuleDestination] = *rulex.Destination
					rule[isNetworkACLRuleDirection] = *rulex.Direction
					rule[isNetworkACLRuleICMP] = make([]map[string]int, 0, 0)
					rule[isNetworkACLRuleTCP] = make([]map[string]int, 0, 0)
					rule[isNetworkACLRuleUDP] = make([]map[string]int, 0, 0)
				}
			}
			rules = append(rules, rule)
		}
	}
	if err = d.Set(isNetworkACLRules, rules); err != nil {
		err = fmt.Errorf("Error setting rules: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl", "read", "set-rules").GetDiag()
	}
	controller, err := flex.GetBaseController(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBaseController failed: %s", err.Error()), "ibm_is_network_acl", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/network/acl"); err != nil {
		err = fmt.Errorf("Error setting resource_controller_url: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl", "read", "set-resource_controller_url").GetDiag()
	}
	if err = d.Set(flex.ResourceName, *nwacl.Name); err != nil {
		err = fmt.Errorf("Error setting resource_name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl", "read", "set-resource_name").GetDiag()
	}
	// d.Set(flex.ResourceCRN, *nwacl.Crn)
	return nil
}

func resourceIBMISNetworkACLUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()

	name := ""
	hasChanged := false

	if d.HasChange(isNetworkACLName) {
		name = d.Get(isNetworkACLName).(string)
		hasChanged = true
	}

	err := nwaclUpdate(context, d, meta, id, name, hasChanged)
	if err != nil {
		return err
	}
	return resourceIBMISNetworkACLRead(context, d, meta)
}

func nwaclUpdate(context context.Context, d *schema.ResourceData, meta interface{}, id, name string, hasChanged bool) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if hasChanged {
		updateNetworkACLOptions := &vpcv1.UpdateNetworkACLOptions{
			ID: &id,
		}
		networkACLPatchModel := &vpcv1.NetworkACLPatch{
			Name: &name,
		}
		networkACLPatch, err := networkACLPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("networkACLPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_network_acl", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateNetworkACLOptions.NetworkACLPatch = networkACLPatch
		_, _, err = sess.UpdateNetworkACLWithContext(context, updateNetworkACLOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateNetworkACLWithContext failed: %s", err.Error()), "ibm_is_network_acl", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	if d.HasChange(isNetworkACLTags) {
		oldList, newList := d.GetChange(isNetworkACLTags)
		err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, d.Get(isNetworkACLCRN).(string), "", isUserTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource network acl (%s) tags: %s", d.Id(), err)
		}
	}
	if d.HasChange(isNetworkACLAccessTags) {
		oldList, newList := d.GetChange(isNetworkACLAccessTags)
		err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, d.Get(isNetworkACLCRN).(string), "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource network acl (%s) access tags: %s", d.Id(), err)
		}
	}
	if d.HasChange(isNetworkACLRules) && !d.IsNewResource() {
		// ── Branch: legacy clear+recreate vs surgical update ───────────────────
		if !d.Get(isNetworkACLRuleUpdateMode).(bool) {
			// Legacy path: delete all rules, recreate from current config.
			rules := d.Get(isNetworkACLRules).([]interface{})
			if err := validateInlineRulesForUpdate(d, make([]interface{}, len(rules))); err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("validateInlineRulesForUpdate failed: %s", err.Error()), "ibm_is_network_acl", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			if err := clearRules(sess, id); err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("clearRules failed: %s", err.Error()), "ibm_is_network_acl", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			for i, r := range rules {
				if err := createSingleNwaclRuleForUpdateLegacy(d, sess, id, r.(map[string]interface{}), i); err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("createSingleNwaclRuleForUpdateLegacy failed: %s", err.Error()), "ibm_is_network_acl", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
			return nil
		}

		// ── Surgical update path ───────────────────────────────────────────────
		ots, _ := d.GetChange(isNetworkACLRules)
		otsIntf := ots.([]interface{})

		// Build old-state map: name → {id, protocol, index}
		type oldRuleInfo struct {
			id       string
			protocol string
			index    int
			data     map[string]interface{}
		}
		oldStateMap := make(map[string]oldRuleInfo, len(otsIntf))
		oldStateOrder := make([]string, len(otsIntf)) // names in old state order
		for i, r := range otsIntf {
			rm := r.(map[string]interface{})
			name, _ := rm[isNetworkACLRuleName].(string)
			ruleId, _ := rm[isNetworkACLRuleID].(string)
			proto, _ := rm[isNetworkACLRuleProtocol].(string)
			oldStateMap[name] = oldRuleInfo{id: ruleId, protocol: proto, index: i, data: rm}
			oldStateOrder[i] = name
		}

		// Build raw-config map: name → {rawVal, index}
		type rawRuleInfo struct {
			val      cty.Value
			index    int
			protocol string // protocol implied by raw config
		}
		rawCfg := d.GetRawConfig()
		rawCfgRules := rawCfg.GetAttr("rules")

		rawConfigMap := make(map[string]rawRuleInfo)
		rawConfigOrder := []string{} // names in desired (raw config) order

		if !rawCfgRules.IsNull() && rawCfgRules.IsKnown() {
			for i := 0; i < rawCfgRules.LengthInt(); i++ {
				rv := rawCfgRules.Index(cty.NumberIntVal(int64(i)))
				if rv.IsNull() {
					continue
				}
				name := ""
				if nameAttr := rv.GetAttr("name"); !nameAttr.IsNull() && nameAttr.IsKnown() {
					name = nameAttr.AsString()
				}
				if name == "" {
					continue
				}
				rawConfigMap[name] = rawRuleInfo{
					val:      rv,
					index:    i,
					protocol: nwaclProtocolFromRawVal(rv),
				}
				rawConfigOrder = append(rawConfigOrder, name)
			}
		}

		// Validate cross-field consistency (only one protocol block, port/icmp vs protocol).
		// action/direction are already enforced by the schema-level validator at plan time.
		// The slice just needs the correct length so the validator can index into GetRawConfig.
		if err := validateInlineRulesForUpdate(d, make([]interface{}, len(rawConfigOrder))); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("validateInlineRulesForUpdate failed: %s", err.Error()), "ibm_is_network_acl", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		// Step 1: Delete rules removed from config
		for _, oldName := range oldStateOrder {
			if _, existsInConfig := rawConfigMap[oldName]; !existsInConfig {
				oldInfo := oldStateMap[oldName]
				log.Printf("[DEBUG] nwaclUpdate: deleting removed rule %q (id=%s)", oldName, oldInfo.id)
				deleteOpts := &vpcv1.DeleteNetworkACLRuleOptions{
					NetworkACLID: &id,
					ID:           &oldInfo.id,
				}
				if resp, err := sess.DeleteNetworkACLRuleWithContext(context, deleteOpts); err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteNetworkACLRuleWithContext (remove rule %q) failed: %s\n%s", oldName, err.Error(), resp), "ibm_is_network_acl", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		}

		// Step 2: Process desired rules in raw-config order
		// We maintain a live slice that tracks the current rule ordering by name so
		// we can compute the correct `before` ID for inserts and repositions.
		// Start from the post-deletion old state order.
		liveOrder := make([]string, 0, len(rawConfigOrder))
		for _, name := range oldStateOrder {
			if _, existsInConfig := rawConfigMap[name]; existsInConfig {
				liveOrder = append(liveOrder, name)
			}
		}

		// liveIDs maps name → current rule ID (updated when rules are created).
		liveIDs := make(map[string]string, len(oldStateMap))
		for name, info := range oldStateMap {
			liveIDs[name] = info.id
		}

		// beforeIDForIndex returns the rule ID that the rule at desiredIndex should
		// be placed before in the API linked-list, based on the current liveOrder.
		// It looks at what rule should immediately follow desiredIndex in rawConfigOrder.
		beforeIDForIndex := func(desiredIndex int) string {
			if desiredIndex+1 >= len(rawConfigOrder) {
				return "" // last rule — no `before`
			}
			// Find the first rule after desiredIndex that already exists in liveIDs.
			for j := desiredIndex + 1; j < len(rawConfigOrder); j++ {
				nextName := rawConfigOrder[j]
				if rid, ok := liveIDs[nextName]; ok && rid != "" {
					return rid
				}
			}
			return ""
		}

		// insertIntoLiveOrder inserts name at position idx, shifting others right.
		// If idx >= len(liveOrder) the name is appended at the end.
		insertIntoLiveOrder := func(name string, idx int) {
			if idx >= len(liveOrder) {
				liveOrder = append(liveOrder, name)
				return
			}
			liveOrder = append(liveOrder, "")
			copy(liveOrder[idx+1:], liveOrder[idx:])
			liveOrder[idx] = name
		}

		// removeFromLiveOrder removes name from liveOrder.
		removeFromLiveOrder := func(name string) {
			for i, n := range liveOrder {
				if n == name {
					liveOrder = append(liveOrder[:i], liveOrder[i+1:]...)
					return
				}
			}
		}

		// currentLiveIndex returns the current position of name in liveOrder, -1 if not found.
		currentLiveIndex := func(name string) int {
			for i, n := range liveOrder {
				if n == name {
					return i
				}
			}
			return -1
		}

		for desiredIdx, ruleName := range rawConfigOrder {
			rawInfo := rawConfigMap[ruleName]
			oldInfo, existsInOld := oldStateMap[ruleName]

			// Scenario: New rule (addition)
			if !existsInOld {
				beforeID := beforeIDForIndex(desiredIdx)
				log.Printf("[DEBUG] nwaclUpdate: adding new rule %q at index %d (before=%s)", ruleName, desiredIdx, beforeID)

				// Build a minimal state map for createSingleNwaclRuleForUpdate.
				stateMap := nwaclRawValToStateMap(rawInfo.val)
				newID, err := createSingleNwaclRuleForUpdate(d, sess, id, stateMap, rawInfo.index, beforeID)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("createSingleNwaclRuleForUpdate (add %q) failed: %s", ruleName, err.Error()), "ibm_is_network_acl", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				liveIDs[ruleName] = newID
				insertIntoLiveOrder(ruleName, desiredIdx)
				continue
			}

			// Scenario: Immutable field changed (name or protocol)
			// Name change: the rule name in rawConfig differs from oldState key.
			// Protocol change: rawConfig protocol differs from old state protocol.
			oldProtocol := oldInfo.protocol
			newProtocol := rawInfo.protocol
			protocolChanged := oldProtocol != "" && newProtocol != "" && oldProtocol != newProtocol

			if protocolChanged {
				log.Printf("[DEBUG] nwaclUpdate: protocol changed for rule %q (%s→%s), delete+recreate", ruleName, oldProtocol, newProtocol)
				beforeID := beforeIDForIndex(desiredIdx)

				deleteOpts := &vpcv1.DeleteNetworkACLRuleOptions{
					NetworkACLID: &id,
					ID:           &oldInfo.id,
				}
				if resp, err := sess.DeleteNetworkACLRuleWithContext(context, deleteOpts); err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteNetworkACLRuleWithContext (protocol change %q) failed: %s\n%s", ruleName, err.Error(), resp), "ibm_is_network_acl", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				removeFromLiveOrder(ruleName)
				delete(liveIDs, ruleName)

				stateMap := nwaclRawValToStateMap(rawInfo.val)
				newID, err := createSingleNwaclRuleForUpdate(d, sess, id, stateMap, rawInfo.index, beforeID)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("createSingleNwaclRuleForUpdate (protocol change %q) failed: %s", ruleName, err.Error()), "ibm_is_network_acl", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				liveIDs[ruleName] = newID
				insertIntoLiveOrder(ruleName, desiredIdx)
				continue
			}

			// Scenario: Reorder only
			curLiveIdx := currentLiveIndex(ruleName)
			needsReorder := curLiveIdx != desiredIdx

			// Scenario: Mutable attribute update
			// Gather patch fields by comparing raw config values against old state.
			rulePatch := &vpcv1.NetworkACLRulePatch{}
			hasRulePatch := false

			// action
			if newAction := nwaclStringAttr(rawInfo.val, "action"); newAction != oldInfo.data[isNetworkACLRuleAction] {
				if newAction != "" {
					rulePatch.Action = &newAction
					hasRulePatch = true
				}
			}
			// source
			if newSrc := nwaclStringAttr(rawInfo.val, "source"); newSrc != oldInfo.data[isNetworkACLRuleSource] {
				if newSrc != "" {
					rulePatch.Source = &newSrc
					hasRulePatch = true
				}
			}
			// destination
			if newDest := nwaclStringAttr(rawInfo.val, "destination"); newDest != oldInfo.data[isNetworkACLRuleDestination] {
				if newDest != "" {
					rulePatch.Destination = &newDest
					hasRulePatch = true
				}
			}
			// direction
			if newDir := nwaclStringAttr(rawInfo.val, "direction"); newDir != oldInfo.data[isNetworkACLRuleDirection] {
				if newDir != "" {
					rulePatch.Direction = &newDir
					hasRulePatch = true
				}
			}
			// icmp type/code — resolve from deprecated icmp{} block if present,
			// otherwise from top-level flat fields. Old state always stores both
			// at the top level (nwaclGet populates rule["type"] and rule["code"]
			// unconditionally), so the comparison target is always oldInfo.data.
			if rawInfo.protocol == "icmp" {
				icmpSrc := rawInfo.val // default: top-level flat fields
				icmpAttr := rawInfo.val.GetAttr("icmp")
				if !icmpAttr.IsNull() && icmpAttr.LengthInt() > 0 {
					if elem := icmpAttr.Index(cty.NumberIntVal(0)); !elem.IsNull() {
						icmpSrc = elem // deprecated block: read type/code from inside icmp{}
					}
				}
				if newType := nwaclInt64AttrFromRaw(icmpSrc, "type"); newType != nil {
					oldTypeVal, _ := oldInfo.data[isNetworkACLRuleICMPType].(int)
					if int64(oldTypeVal) != *newType {
						rulePatch.Type = newType
						hasRulePatch = true
					}
				}
				if newCode := nwaclInt64AttrFromRaw(icmpSrc, "code"); newCode != nil {
					oldCodeVal, _ := oldInfo.data[isNetworkACLRuleICMPCode].(int)
					if int64(oldCodeVal) != *newCode {
						rulePatch.Code = newCode
						hasRulePatch = true
					}
				}
			}
			// tcp/udp port fields — resolve from deprecated tcp{}/udp{} block if present,
			// otherwise from top-level flat fields. Old state always stores port values
			// at the top level (nwaclGet populates them unconditionally).
			if rawInfo.protocol == "tcp" || rawInfo.protocol == "udp" {
				portSrc := rawInfo.val // default: top-level flat fields
				tcpAttr := rawInfo.val.GetAttr("tcp")
				udpAttr := rawInfo.val.GetAttr("udp")
				if !tcpAttr.IsNull() && tcpAttr.LengthInt() > 0 {
					if elem := tcpAttr.Index(cty.NumberIntVal(0)); !elem.IsNull() {
						portSrc = elem // deprecated tcp{} block
					}
				} else if !udpAttr.IsNull() && udpAttr.LengthInt() > 0 {
					if elem := udpAttr.Index(cty.NumberIntVal(0)); !elem.IsNull() {
						portSrc = elem // deprecated udp{} block
					}
				}

				if pm := nwaclInt64AttrFromRaw(portSrc, "port_min"); pm != nil {
					oldVal, _ := oldInfo.data[isNetworkACLRulePortMin].(int)
					if int64(oldVal) != *pm {
						rulePatch.DestinationPortMin = pm
						hasRulePatch = true
					}
				}
				if pm := nwaclInt64AttrFromRaw(portSrc, "port_max"); pm != nil {
					oldVal, _ := oldInfo.data[isNetworkACLRulePortMax].(int)
					if int64(oldVal) != *pm {
						rulePatch.DestinationPortMax = pm
						hasRulePatch = true
					}
				}
				if pm := nwaclInt64AttrFromRaw(portSrc, "source_port_min"); pm != nil {
					oldVal, _ := oldInfo.data[isNetworkACLRuleSourcePortMin].(int)
					if int64(oldVal) != *pm {
						rulePatch.SourcePortMin = pm
						hasRulePatch = true
					}
				}
				if pm := nwaclInt64AttrFromRaw(portSrc, "source_port_max"); pm != nil {
					oldVal, _ := oldInfo.data[isNetworkACLRuleSourcePortMax].(int)
					if int64(oldVal) != *pm {
						rulePatch.SourcePortMax = pm
						hasRulePatch = true
					}
				}
			}

			// Apply attribute patch if needed.
			if hasRulePatch {
				log.Printf("[DEBUG] nwaclUpdate: patching mutable attrs on rule %q (id=%s)", ruleName, oldInfo.id)
				patchMap, err := rulePatch.AsPatch()
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("NetworkACLRulePatch.AsPatch() failed for %q: %s", ruleName, err.Error()), "ibm_is_network_acl", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				updateOpts := &vpcv1.UpdateNetworkACLRuleOptions{
					NetworkACLID:        &id,
					ID:                  &oldInfo.id,
					NetworkACLRulePatch: patchMap,
				}
				if _, _, err := sess.UpdateNetworkACLRuleWithContext(context, updateOpts); err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateNetworkACLRuleWithContext (mutable update %q) failed: %s", ruleName, err.Error()), "ibm_is_network_acl", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}

			// Apply reorder patch if needed.
			if needsReorder {
				beforeID := beforeIDForIndex(desiredIdx)
				log.Printf("[DEBUG] nwaclUpdate: reordering rule %q from live pos %d to %d (before=%s)", ruleName, curLiveIdx, desiredIdx, beforeID)
				reorderPatch := &vpcv1.NetworkACLRulePatch{}
				if beforeID != "" {
					reorderPatch.Before = &vpcv1.NetworkACLRuleBeforePatchNetworkACLRuleIdentityByID{
						ID: &beforeID,
					}
				}
				reorderPatchMap, err := reorderPatch.AsPatch()
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("NetworkACLRulePatch.AsPatch() (reorder %q) failed: %s", ruleName, err.Error()), "ibm_is_network_acl", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				reorderOpts := &vpcv1.UpdateNetworkACLRuleOptions{
					NetworkACLID:        &id,
					ID:                  &oldInfo.id,
					NetworkACLRulePatch: reorderPatchMap,
				}
				if _, _, err := sess.UpdateNetworkACLRuleWithContext(context, reorderOpts); err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateNetworkACLRuleWithContext (reorder %q) failed: %s", ruleName, err.Error()), "ibm_is_network_acl", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				// Update liveOrder to reflect the new position.
				removeFromLiveOrder(ruleName)
				insertIntoLiveOrder(ruleName, desiredIdx)
			}
		}
	}
	return nil
}

// nwaclProtocolFromRawVal determines the effective protocol from a raw cty rule value.
// It prefers explicit deprecated blocks (icmp/tcp/udp) over the flat protocol attribute.
func nwaclProtocolFromRawVal(rv cty.Value) string {
	if rv.IsNull() || !rv.IsKnown() {
		return ""
	}
	if icmpAttr := rv.GetAttr("icmp"); !icmpAttr.IsNull() && icmpAttr.LengthInt() > 0 {
		return "icmp"
	}
	if tcpAttr := rv.GetAttr("tcp"); !tcpAttr.IsNull() && tcpAttr.LengthInt() > 0 {
		return "tcp"
	}
	if udpAttr := rv.GetAttr("udp"); !udpAttr.IsNull() && udpAttr.LengthInt() > 0 {
		return "udp"
	}
	if p := rv.GetAttr("protocol"); !p.IsNull() && p.IsKnown() {
		return p.AsString()
	}
	return ""
}

// nwaclStringAttr safely extracts a string attribute from a cty.Value.
func nwaclStringAttr(rv cty.Value, attr string) string {
	if rv.IsNull() || !rv.IsKnown() {
		return ""
	}
	v := rv.GetAttr(attr)
	if v.IsNull() || !v.IsKnown() {
		return ""
	}
	return v.AsString()
}

// nwaclInt64AttrFromRaw extracts an int64 pointer from a cty.Value attribute.
// Returns nil if the attribute is null or unknown.
func nwaclInt64AttrFromRaw(rv cty.Value, attr string) *int64 {
	if rv.IsNull() || !rv.IsKnown() {
		return nil
	}
	v := rv.GetAttr(attr)
	if v.IsNull() || !v.IsKnown() {
		return nil
	}
	n, _ := v.AsBigFloat().Int64()
	return &n
}

// nwaclRawValToStateMap converts a raw cty rule value to the map[string]interface{}
// format expected by createSingleNwaclRuleForUpdate. Only the fields needed for the
// create call are populated; the function reads protocol/port detail directly from
// GetRawConfig via the index parameter passed alongside.
func nwaclRawValToStateMap(rv cty.Value) map[string]interface{} {
	m := map[string]interface{}{
		isNetworkACLRuleName:        nwaclStringAttr(rv, "name"),
		isNetworkACLRuleAction:      nwaclStringAttr(rv, "action"),
		isNetworkACLRuleSource:      nwaclStringAttr(rv, "source"),
		isNetworkACLRuleDestination: nwaclStringAttr(rv, "destination"),
		isNetworkACLRuleDirection:   nwaclStringAttr(rv, "direction"),
		isNetworkACLRuleProtocol:    nwaclStringAttr(rv, "protocol"),
		isNetworkACLRuleICMP:        []interface{}{},
		isNetworkACLRuleTCP:         []interface{}{},
		isNetworkACLRuleUDP:         []interface{}{},
	}
	return m
}

func resourceIBMISNetworkACLDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	err := nwaclDelete(context, d, meta, id)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func nwaclDelete(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getNetworkAclOptions := &vpcv1.GetNetworkACLOptions{
		ID: &id,
	}
	_, response, err := sess.GetNetworkACLWithContext(context, getNetworkAclOptions)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetNetworkACLWithContext failed: %s", err.Error()), "ibm_is_network_acl", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteNetworkAclOptions := &vpcv1.DeleteNetworkACLOptions{
		ID: &id,
	}
	response, err = sess.DeleteNetworkACLWithContext(context, deleteNetworkAclOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteNetworkACLWithContext failed: %s", err.Error()), "ibm_is_network_acl", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}

func resourceIBMISNetworkACLExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id := d.Id()
	exists, err := nwaclExists(d, meta, id)
	return exists, err
}

func nwaclExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	getNetworkAclOptions := &vpcv1.GetNetworkACLOptions{
		ID: &id,
	}
	_, response, err := sess.GetNetworkACL(getNetworkAclOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetNetworkACL failed: %s", err.Error()), "ibm_is_network_acl", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}

func checkNetworkACLNil(ptr *int64) int {
	if ptr == nil {
		return 0
	}
	return int(*ptr)
}

func clearRules(nwaclC *vpcv1.VpcV1, nwaclid string) error {
	start := ""
	allrecs := []vpcv1.NetworkACLRuleItemIntf{}
	for {
		listNetworkAclRulesOptions := &vpcv1.ListNetworkACLRulesOptions{
			NetworkACLID: &nwaclid,
		}
		if start != "" {
			listNetworkAclRulesOptions.Start = &start
		}
		rawrules, response, err := nwaclC.ListNetworkACLRules(listNetworkAclRulesOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error Listing network ACL rules : %s\n%s", err, response)
		}
		start = flex.GetNext(rawrules.Next)
		allrecs = append(allrecs, rawrules.Rules...)
		if start == "" {
			break
		}
	}

	for _, rule := range allrecs {
		deleteNetworkAclRuleOptions := &vpcv1.DeleteNetworkACLRuleOptions{
			NetworkACLID: &nwaclid,
		}
		switch reflect.TypeOf(rule).String() {
		case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIcmp":
			rule := rule.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIcmp)
			deleteNetworkAclRuleOptions.ID = rule.ID
		case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTcpudp":
			rule := rule.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTcpudp)
			deleteNetworkAclRuleOptions.ID = rule.ID
		case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolAny":
			rule := rule.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolAny)
			deleteNetworkAclRuleOptions.ID = rule.ID
		case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIcmptcpudp":
			rule := rule.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIcmptcpudp)
			deleteNetworkAclRuleOptions.ID = rule.ID
		case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIndividual":
			rule := rule.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIndividual)
			deleteNetworkAclRuleOptions.ID = rule.ID
		case "*vpcv1.NetworkACLRuleItem":
			rule := rule.(*vpcv1.NetworkACLRuleItem)
			deleteNetworkAclRuleOptions.ID = rule.ID
		}

		response, err := nwaclC.DeleteNetworkACLRule(deleteNetworkAclRuleOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error Deleting network ACL rule : %s\n%s", err, response)
		}
	}
	return nil
}

func validateInlineRules(d *schema.ResourceData, rules []interface{}) error {
	for i, rule := range rules {
		rulex := rule.(map[string]interface{})
		action := rulex[isNetworkACLRuleAction].(string)
		if (action != "allow") && (action != "deny") {
			return fmt.Errorf("[ERROR] Invalid action. valid values are allow|deny")
		}

		direction := rulex[isNetworkACLRuleDirection].(string)
		direction = strings.ToLower(direction)

		// Use GetRawConfig to get the actual HCL configuration without state merging
		// This correctly detects which protocol blocks are defined in the user's config
		rawConfig := d.GetRawConfig()
		rulesAttr := rawConfig.GetAttr("rules")

		icmp := false
		tcp := false
		udp := false

		if !rulesAttr.IsNull() && rulesAttr.LengthInt() > i {
			ruleVal := rulesAttr.Index(cty.NumberIntVal(int64(i)))
			if !ruleVal.IsNull() {
				icmpAttr := ruleVal.GetAttr("icmp")
				tcpAttr := ruleVal.GetAttr("tcp")
				udpAttr := ruleVal.GetAttr("udp")

				icmp = !icmpAttr.IsNull() && icmpAttr.LengthInt() > 0
				tcp = !tcpAttr.IsNull() && tcpAttr.LengthInt() > 0
				udp = !udpAttr.IsNull() && udpAttr.LengthInt() > 0
			}
		}

		log.Printf("[DEBUG] validateInlineRules rule[%d] from RawConfig: icmp=%t, tcp=%t, udp=%t", i, icmp, tcp, udp)

		if (icmp && tcp) || (icmp && udp) || (tcp && udp) {
			return fmt.Errorf("Only one of icmp|tcp|udp can be defined per rule")
		}

		protocol := rulex[isNetworkACLRuleProtocol]
		icmpType := fmt.Sprintf("rules.%d.type", i)
		icmpCode := fmt.Sprintf("rules.%d.code", i)
		portMin := fmt.Sprintf("rules.%d.port_min", i)
		portMax := fmt.Sprintf("rules.%d.port_max", i)
		srcPortMin := fmt.Sprintf("rules.%d.source_port_min", i)
		srcPortMax := fmt.Sprintf("rules.%d.source_port_max", i)
		if protocol != "icmp" && protocol != "" {
			if _, ok := d.GetOk(icmpType); ok {
				return fmt.Errorf("attribute 'type' conflicts with protocol %q; 'type' is only valid for icmp protocol", protocol)
			}
			if _, ok := d.GetOk(icmpCode); ok {
				return fmt.Errorf("attribute 'code' conflicts with protocol %q; 'code' is only valid for icmp protocol", protocol)
			}
		}

		if protocol != "tcp" && protocol != "udp" && protocol != "" {
			if _, ok := d.GetOk(portMin); ok {
				return fmt.Errorf("attribute 'port_min' conflicts with protocol %s; ports apply only to tcp/udp protocol", protocol)
			}
			if _, ok := d.GetOk(portMax); ok {
				return fmt.Errorf("attribute 'port_max' conflicts with protocol %s; ports apply only to tcp/udp protocol", protocol)
			}

			if _, ok := d.GetOk(srcPortMin); ok {
				return fmt.Errorf("attribute 'source_port_min' conflicts with protocol %s; ports apply only to tcp/udp protocol", protocol)
			}
			if _, ok := d.GetOk(srcPortMax); ok {
				return fmt.Errorf("attribute 'source_port_max' conflicts with protocol %s; ports apply only to tcp/udp protocol", protocol)
			}
		}

	}
	return nil
}

func validateInlineRulesForUpdate(d *schema.ResourceData, rules []interface{}) error {
	// Read everything from GetRawConfig so state-merged values don't cause
	// false positives on update. On create GetRawConfig is fully populated;
	// on update it is also fully populated (only destroy returns null).
	rawConfig := d.GetRawConfig()
	var rawRulesAttr cty.Value
	if rawConfig.IsKnown() && !rawConfig.IsNull() {
		rawRulesAttr = rawConfig.GetAttr("rules")
	} else {
		rawRulesAttr = cty.NullVal(cty.DynamicPseudoType)
	}

	for i := range rules {
		hasIcmpBlock, hasTcpBlock, hasUdpBlock := false, false, false
		protocol := ""
		hasIcmpType, hasIcmpCode := false, false
		hasPortMin, hasPortMax, hasSrcPortMin, hasSrcPortMax := false, false, false, false

		if !rawRulesAttr.IsNull() && rawRulesAttr.IsKnown() && rawRulesAttr.LengthInt() > i {
			ruleVal := rawRulesAttr.Index(cty.NumberIntVal(int64(i)))
			if !ruleVal.IsNull() {
				icmpAttr := ruleVal.GetAttr("icmp")
				tcpAttr := ruleVal.GetAttr("tcp")
				udpAttr := ruleVal.GetAttr("udp")
				hasIcmpBlock = !icmpAttr.IsNull() && icmpAttr.LengthInt() > 0
				hasTcpBlock = !tcpAttr.IsNull() && tcpAttr.LengthInt() > 0
				hasUdpBlock = !udpAttr.IsNull() && udpAttr.LengthInt() > 0

				if p := ruleVal.GetAttr("protocol"); !p.IsNull() && p.IsKnown() {
					protocol = p.AsString()
				}
				if v := ruleVal.GetAttr("type"); !v.IsNull() && v.IsKnown() {
					hasIcmpType = true
				}
				if v := ruleVal.GetAttr("code"); !v.IsNull() && v.IsKnown() {
					hasIcmpCode = true
				}
				if v := ruleVal.GetAttr("port_min"); !v.IsNull() && v.IsKnown() {
					hasPortMin = true
				}
				if v := ruleVal.GetAttr("port_max"); !v.IsNull() && v.IsKnown() {
					hasPortMax = true
				}
				if v := ruleVal.GetAttr("source_port_min"); !v.IsNull() && v.IsKnown() {
					hasSrcPortMin = true
				}
				if v := ruleVal.GetAttr("source_port_max"); !v.IsNull() && v.IsKnown() {
					hasSrcPortMax = true
				}
			}
		}

		log.Printf("[DEBUG] validateInlineRules rule[%d] from RawConfig: icmp=%t, tcp=%t, udp=%t, protocol=%q", i, hasIcmpBlock, hasTcpBlock, hasUdpBlock, protocol)

		if (hasIcmpBlock && hasTcpBlock) || (hasIcmpBlock && hasUdpBlock) || (hasTcpBlock && hasUdpBlock) {
			return fmt.Errorf("Only one of icmp|tcp|udp can be defined per rule")
		}

		if protocol != "icmp" && protocol != "" {
			if hasIcmpType {
				return fmt.Errorf("attribute 'type' conflicts with protocol %q; 'type' is only valid for icmp protocol", protocol)
			}
			if hasIcmpCode {
				return fmt.Errorf("attribute 'code' conflicts with protocol %q; 'code' is only valid for icmp protocol", protocol)
			}
		}

		if protocol != "tcp" && protocol != "udp" && protocol != "" {
			if hasPortMin {
				return fmt.Errorf("attribute 'port_min' conflicts with protocol %s; ports apply only to tcp/udp protocol", protocol)
			}
			if hasPortMax {
				return fmt.Errorf("attribute 'port_max' conflicts with protocol %s; ports apply only to tcp/udp protocol", protocol)
			}
			if hasSrcPortMin {
				return fmt.Errorf("attribute 'source_port_min' conflicts with protocol %s; ports apply only to tcp/udp protocol", protocol)
			}
			if hasSrcPortMax {
				return fmt.Errorf("attribute 'source_port_max' conflicts with protocol %s; ports apply only to tcp/udp protocol", protocol)
			}
		}
	}
	return nil
}

func createInlineRules(d *schema.ResourceData, nwaclC *vpcv1.VpcV1, nwaclid string, rules []interface{}) error {
	before := ""

	for i := 0; i <= len(rules)-1; i++ {
		rulex := rules[i].(map[string]interface{})

		name := rulex[isNetworkACLRuleName].(string)
		source := rulex[isNetworkACLRuleSource].(string)
		destination := rulex[isNetworkACLRuleDestination].(string)
		action := rulex[isNetworkACLRuleAction].(string)
		direction := rulex[isNetworkACLRuleDirection].(string)
		icmp := rulex[isNetworkACLRuleICMP].([]interface{})
		tcp := rulex[isNetworkACLRuleTCP].([]interface{})
		udp := rulex[isNetworkACLRuleUDP].([]interface{})
		icmptype := int64(-1)
		icmpcode := int64(-1)
		minport := int64(-1)
		maxport := int64(-1)
		sourceminport := int64(-1)
		sourcemaxport := int64(-1)
		protocol := "icmp_tcp_udp"
		if action == "deny" {
			protocol = "any"
		}
		if protocolVal, ok := rulex[isNetworkACLRuleProtocol]; ok {
			if str, ok := protocolVal.(string); ok && str != "" {
				protocol = str
			}
		}
		ruleTemplate := &vpcv1.NetworkACLRulePrototype{
			Action:      &action,
			Destination: &destination,
			Direction:   &direction,
			Source:      &source,
			Name:        &name,
		}

		if before != "" {
			ruleTemplate.Before = &vpcv1.NetworkACLRuleBeforePrototype{
				ID: &before,
			}
		}

		// Detect if user is using new-style top-level fields vs deprecated blocks
		// by checking which set of fields has actually changed
		useTopLevelPorts := false
		if protocol == "tcp" || protocol == "udp" {
			portMinPath := fmt.Sprintf("rules.%d.port_min", i)
			portMaxPath := fmt.Sprintf("rules.%d.port_max", i)
			srcPortMinPath := fmt.Sprintf("rules.%d.source_port_min", i)
			srcPortMaxPath := fmt.Sprintf("rules.%d.source_port_max", i)
			if d.HasChange(portMinPath) || d.HasChange(portMaxPath) ||
				d.HasChange(srcPortMinPath) || d.HasChange(srcPortMaxPath) {
				useTopLevelPorts = true
			}
		}
		useTopLevelIcmp := false
		if protocol == "icmp" {
			icmpTypePath := fmt.Sprintf("rules.%d.type", i)
			icmpCodePath := fmt.Sprintf("rules.%d.code", i)
			if d.HasChange(icmpTypePath) || d.HasChange(icmpCodePath) {
				useTopLevelIcmp = true
			}
		}

		if len(icmp) > 0 && !useTopLevelIcmp {
			protocol = "icmp"
			ruleTemplate.Protocol = &protocol
			if !isNil(icmp[0]) {
				icmpTypePath := fmt.Sprintf("rules.%d.icmp.0.%s", i, isNetworkACLRuleICMPType)
				icmpCodePath := fmt.Sprintf("rules.%d.icmp.0.%s", i, isNetworkACLRuleICMPCode)
				if val, ok := d.GetOkExists(icmpTypePath); ok {
					icmptype = int64(val.(int))
					ruleTemplate.Type = &icmptype
				}
				if val, ok := d.GetOkExists(icmpCodePath); ok {
					icmpcode = int64(val.(int))
					ruleTemplate.Code = &icmpcode
				}
				if ruleTemplate.Type != nil && ruleTemplate.Code == nil {
					v := int64(0)
					ruleTemplate.Code = &v
				}
				if ruleTemplate.Code != nil && ruleTemplate.Type == nil {
					v := int64(0)
					ruleTemplate.Type = &v
				}
			}
		} else if protocol == "icmp" {
			icmpType := fmt.Sprintf("rules.%d.type", i)
			icmpCode := fmt.Sprintf("rules.%d.code", i)
			ruleTemplate.Protocol = &protocol
			if val, ok := d.GetOkExists(icmpType); ok {
				icmptype = int64(val.(int))
				ruleTemplate.Type = &icmptype
			}
			if val, ok := d.GetOkExists(icmpCode); ok {
				icmpcode = int64(val.(int))
				ruleTemplate.Code = &icmpcode
			}
		}

		if len(tcp) > 0 && !useTopLevelPorts {
			protocol = "tcp"
			ruleTemplate.Protocol = &protocol
			if !isNil(tcp[0]) {
				tcpval := tcp[0].(map[string]interface{})
				if val, ok := tcpval[isNetworkACLRulePortMin]; ok {
					minport = int64(val.(int))
					ruleTemplate.DestinationPortMin = &minport
				}
				if val, ok := tcpval[isNetworkACLRulePortMax]; ok {
					maxport = int64(val.(int))
					ruleTemplate.DestinationPortMax = &maxport
				}
				if val, ok := tcpval[isNetworkACLRuleSourcePortMin]; ok {
					sourceminport = int64(val.(int))
					ruleTemplate.SourcePortMin = &sourceminport
				}
				if val, ok := tcpval[isNetworkACLRuleSourcePortMax]; ok {
					sourcemaxport = int64(val.(int))
					ruleTemplate.SourcePortMax = &sourcemaxport
				}
			}
		} else if protocol == "tcp" {
			ruleTemplate.Protocol = &protocol
			if val, ok := rulex[isNetworkACLRulePortMin]; ok {
				minport = int64(val.(int))
				ruleTemplate.DestinationPortMin = &minport
			}
			if val, ok := rulex[isNetworkACLRulePortMax]; ok {
				maxport = int64(val.(int))
				ruleTemplate.DestinationPortMax = &maxport
			}
			if val, ok := rulex[isNetworkACLRuleSourcePortMin]; ok {
				sourceminport = int64(val.(int))
				ruleTemplate.SourcePortMin = &sourceminport
			}
			if val, ok := rulex[isNetworkACLRuleSourcePortMax]; ok {
				sourcemaxport = int64(val.(int))
				ruleTemplate.SourcePortMax = &sourcemaxport
			}
			if minport == 0 {
				ruleTemplate.DestinationPortMin = nil
			}
			if maxport == 0 {
				ruleTemplate.DestinationPortMax = nil
			}
			if sourceminport == 0 {
				ruleTemplate.SourcePortMin = nil
			}
			if sourcemaxport == 0 {
				ruleTemplate.SourcePortMax = nil
			}
		}

		if len(udp) > 0 && !useTopLevelPorts {
			protocol = "udp"
			ruleTemplate.Protocol = &protocol
			if !isNil(udp[0]) {
				udpval := udp[0].(map[string]interface{})
				if val, ok := udpval[isNetworkACLRulePortMin]; ok {
					minport = int64(val.(int))
					ruleTemplate.DestinationPortMin = &minport
				}
				if val, ok := udpval[isNetworkACLRulePortMax]; ok {
					maxport = int64(val.(int))
					ruleTemplate.DestinationPortMax = &maxport
				}
				if val, ok := udpval[isNetworkACLRuleSourcePortMin]; ok {
					sourceminport = int64(val.(int))
					ruleTemplate.SourcePortMin = &sourceminport
				}
				if val, ok := udpval[isNetworkACLRuleSourcePortMax]; ok {
					sourcemaxport = int64(val.(int))
					ruleTemplate.SourcePortMax = &sourcemaxport
				}
			}
		} else if protocol == "udp" {
			ruleTemplate.Protocol = &protocol
			if val, ok := rulex[isNetworkACLRulePortMin]; ok {
				minport = int64(val.(int))
				ruleTemplate.DestinationPortMin = &minport
			}
			if val, ok := rulex[isNetworkACLRulePortMax]; ok {
				maxport = int64(val.(int))
				ruleTemplate.DestinationPortMax = &maxport
			}
			if val, ok := rulex[isNetworkACLRuleSourcePortMin]; ok {
				sourceminport = int64(val.(int))
				ruleTemplate.SourcePortMin = &sourceminport
			}
			if val, ok := rulex[isNetworkACLRuleSourcePortMax]; ok {
				sourcemaxport = int64(val.(int))
				ruleTemplate.SourcePortMax = &sourcemaxport
			}
			if minport == 0 {
				ruleTemplate.DestinationPortMin = nil
			}
			if maxport == 0 {
				ruleTemplate.DestinationPortMax = nil
			}
			if sourceminport == 0 {
				ruleTemplate.SourcePortMin = nil
			}
			if sourcemaxport == 0 {
				ruleTemplate.SourcePortMax = nil
			}
		}
		ruleTemplate.Protocol = &protocol

		createNetworkAclRuleOptions := &vpcv1.CreateNetworkACLRuleOptions{
			NetworkACLID:            &nwaclid,
			NetworkACLRulePrototype: ruleTemplate,
		}
		_, response, err := nwaclC.CreateNetworkACLRule(createNetworkAclRuleOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error Creating network ACL rule : %s\n%s", err, response)
		}
	}
	return nil
}

func isNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}

// createSingleNwaclRuleForUpdateLegacy is the legacy variant used by the clear+recreate
// update path (rule_update_mode = false). It does not accept a `before` parameter and
// does not return a rule ID — rules are simply appended in order after a full clear.
func createSingleNwaclRuleForUpdateLegacy(d *schema.ResourceData, nwaclC *vpcv1.VpcV1, nwaclid string, rulex map[string]interface{}, i int) error {
	name := rulex[isNetworkACLRuleName].(string)
	source := rulex[isNetworkACLRuleSource].(string)
	destination := rulex[isNetworkACLRuleDestination].(string)
	action := rulex[isNetworkACLRuleAction].(string)
	direction := rulex[isNetworkACLRuleDirection].(string)

	hasIcmpBlock, hasTcpBlock, hasUdpBlock := false, false, false
	protocol := "icmp_tcp_udp"
	if action == "deny" {
		protocol = "any"
	}

	var rawIcmpType, rawIcmpCode *int64
	var rawPortMin, rawPortMax, rawSrcPortMin, rawSrcPortMax *int64

	rawConfig := d.GetRawConfig()
	var rulesAttr cty.Value
	if rawConfig.IsKnown() && !rawConfig.IsNull() {
		rulesAttr = rawConfig.GetAttr("rules")
	} else {
		rulesAttr = cty.NullVal(cty.DynamicPseudoType)
	}

	if rulesAttr.IsKnown() && !rulesAttr.IsNull() && rulesAttr.LengthInt() > i {
		ruleVal := rulesAttr.Index(cty.NumberIntVal(int64(i)))
		if !ruleVal.IsNull() {
			protocolAttr := ruleVal.GetAttr("protocol")
			if !protocolAttr.IsNull() && protocolAttr.IsKnown() {
				if p := protocolAttr.AsString(); p != "" {
					protocol = p
				}
			}

			icmpAttr := ruleVal.GetAttr("icmp")
			tcpAttr := ruleVal.GetAttr("tcp")
			udpAttr := ruleVal.GetAttr("udp")
			hasIcmpBlock = !icmpAttr.IsNull() && icmpAttr.IsKnown() && icmpAttr.LengthInt() > 0
			hasTcpBlock = !tcpAttr.IsNull() && tcpAttr.IsKnown() && tcpAttr.LengthInt() > 0
			hasUdpBlock = !udpAttr.IsNull() && udpAttr.IsKnown() && udpAttr.LengthInt() > 0

			if hasIcmpBlock {
				elem := icmpAttr.Index(cty.NumberIntVal(0))
				if !elem.IsNull() {
					if t := elem.GetAttr("type"); !t.IsNull() && t.IsKnown() {
						v, _ := t.AsBigFloat().Int64()
						rawIcmpType = &v
					}
					if c := elem.GetAttr("code"); !c.IsNull() && c.IsKnown() {
						v, _ := c.AsBigFloat().Int64()
						rawIcmpCode = &v
					}
				}
			}
			if !hasIcmpBlock {
				if t := ruleVal.GetAttr("type"); !t.IsNull() && t.IsKnown() {
					v, _ := t.AsBigFloat().Int64()
					rawIcmpType = &v
				}
				if c := ruleVal.GetAttr("code"); !c.IsNull() && c.IsKnown() {
					v, _ := c.AsBigFloat().Int64()
					rawIcmpCode = &v
				}
			}
			if hasTcpBlock {
				elem := tcpAttr.Index(cty.NumberIntVal(0))
				if !elem.IsNull() {
					if v := elem.GetAttr("port_min"); !v.IsNull() && v.IsKnown() {
						n, _ := v.AsBigFloat().Int64()
						rawPortMin = &n
					}
					if v := elem.GetAttr("port_max"); !v.IsNull() && v.IsKnown() {
						n, _ := v.AsBigFloat().Int64()
						rawPortMax = &n
					}
					if v := elem.GetAttr("source_port_min"); !v.IsNull() && v.IsKnown() {
						n, _ := v.AsBigFloat().Int64()
						rawSrcPortMin = &n
					}
					if v := elem.GetAttr("source_port_max"); !v.IsNull() && v.IsKnown() {
						n, _ := v.AsBigFloat().Int64()
						rawSrcPortMax = &n
					}
				}
			}
			if hasUdpBlock {
				elem := udpAttr.Index(cty.NumberIntVal(0))
				if !elem.IsNull() {
					if v := elem.GetAttr("port_min"); !v.IsNull() && v.IsKnown() {
						n, _ := v.AsBigFloat().Int64()
						rawPortMin = &n
					}
					if v := elem.GetAttr("port_max"); !v.IsNull() && v.IsKnown() {
						n, _ := v.AsBigFloat().Int64()
						rawPortMax = &n
					}
					if v := elem.GetAttr("source_port_min"); !v.IsNull() && v.IsKnown() {
						n, _ := v.AsBigFloat().Int64()
						rawSrcPortMin = &n
					}
					if v := elem.GetAttr("source_port_max"); !v.IsNull() && v.IsKnown() {
						n, _ := v.AsBigFloat().Int64()
						rawSrcPortMax = &n
					}
				}
			}
			if !hasTcpBlock && !hasUdpBlock {
				if v := ruleVal.GetAttr("port_min"); !v.IsNull() && v.IsKnown() {
					n, _ := v.AsBigFloat().Int64()
					rawPortMin = &n
				}
				if v := ruleVal.GetAttr("port_max"); !v.IsNull() && v.IsKnown() {
					n, _ := v.AsBigFloat().Int64()
					rawPortMax = &n
				}
				if v := ruleVal.GetAttr("source_port_min"); !v.IsNull() && v.IsKnown() {
					n, _ := v.AsBigFloat().Int64()
					rawSrcPortMin = &n
				}
				if v := ruleVal.GetAttr("source_port_max"); !v.IsNull() && v.IsKnown() {
					n, _ := v.AsBigFloat().Int64()
					rawSrcPortMax = &n
				}
			}
		}
	}

	if hasIcmpBlock {
		protocol = "icmp"
	} else if hasTcpBlock {
		protocol = "tcp"
	} else if hasUdpBlock {
		protocol = "udp"
	}

	ruleTemplate := &vpcv1.NetworkACLRulePrototype{
		Action:      &action,
		Destination: &destination,
		Direction:   &direction,
		Source:      &source,
		Name:        &name,
		Protocol:    &protocol,
	}

	switch protocol {
	case "icmp":
		ruleTemplate.Type = rawIcmpType
		ruleTemplate.Code = rawIcmpCode
		if hasIcmpBlock {
			if ruleTemplate.Type != nil && ruleTemplate.Code == nil {
				v := int64(0)
				ruleTemplate.Code = &v
			}
			if ruleTemplate.Code != nil && ruleTemplate.Type == nil {
				v := int64(0)
				ruleTemplate.Type = &v
			}
		}
	case "tcp", "udp":
		ruleTemplate.DestinationPortMin = rawPortMin
		ruleTemplate.DestinationPortMax = rawPortMax
		ruleTemplate.SourcePortMin = rawSrcPortMin
		ruleTemplate.SourcePortMax = rawSrcPortMax
	}

	createNetworkAclRuleOptions := &vpcv1.CreateNetworkACLRuleOptions{
		NetworkACLID:            &nwaclid,
		NetworkACLRulePrototype: ruleTemplate,
	}
	_, response, err := nwaclC.CreateNetworkACLRule(createNetworkAclRuleOptions)
	if err != nil {
		return fmt.Errorf("[ERROR] Error Creating network ACL rule : %s\n%s", err, response)
	}
	return nil
}

func createSingleNwaclRuleForUpdate(d *schema.ResourceData, nwaclC *vpcv1.VpcV1, nwaclid string, rulex map[string]interface{}, i int, before string) (string, error) {
	name := rulex[isNetworkACLRuleName].(string)
	source := rulex[isNetworkACLRuleSource].(string)
	destination := rulex[isNetworkACLRuleDestination].(string)
	action := rulex[isNetworkACLRuleAction].(string)
	direction := rulex[isNetworkACLRuleDirection].(string)

	// Use GetRawConfig exclusively for all field values — never read from the
	// merged state map (rulex/d.Get) for protocol or port/icmp fields, as
	// Computed fields bleed old state values through on update.
	hasIcmpBlock, hasTcpBlock, hasUdpBlock := false, false, false
	protocol := "icmp_tcp_udp"
	if action == "deny" {
		protocol = "any"
	}

	// rawIcmpType/rawIcmpCode: nil means not set in config.
	var rawIcmpType, rawIcmpCode *int64
	var rawPortMin, rawPortMax, rawSrcPortMin, rawSrcPortMax *int64

	rawConfig := d.GetRawConfig()
	rulesAttr := rawConfig.GetAttr("rules")
	if !rulesAttr.IsNull() && rulesAttr.LengthInt() > i {
		ruleVal := rulesAttr.Index(cty.NumberIntVal(int64(i)))
		if !ruleVal.IsNull() {
			// Protocol
			protocolAttr := ruleVal.GetAttr("protocol")
			if !protocolAttr.IsNull() && protocolAttr.IsKnown() {
				if p := protocolAttr.AsString(); p != "" {
					protocol = p
				}
			}

			// Deprecated blocks
			icmpAttr := ruleVal.GetAttr("icmp")
			tcpAttr := ruleVal.GetAttr("tcp")
			udpAttr := ruleVal.GetAttr("udp")
			hasIcmpBlock = !icmpAttr.IsNull() && icmpAttr.LengthInt() > 0
			hasTcpBlock = !tcpAttr.IsNull() && tcpAttr.LengthInt() > 0
			hasUdpBlock = !udpAttr.IsNull() && udpAttr.LengthInt() > 0

			// icmp block fields
			if hasIcmpBlock {
				elem := icmpAttr.Index(cty.NumberIntVal(0))
				if !elem.IsNull() {
					if t := elem.GetAttr("type"); !t.IsNull() && t.IsKnown() {
						v, _ := t.AsBigFloat().Int64()
						rawIcmpType = &v
					}
					if c := elem.GetAttr("code"); !c.IsNull() && c.IsKnown() {
						v, _ := c.AsBigFloat().Int64()
						rawIcmpCode = &v
					}
				}
			}

			// top-level icmp fields (new style: protocol="icmp" + type/code flat)
			if !hasIcmpBlock {
				if t := ruleVal.GetAttr("type"); !t.IsNull() && t.IsKnown() {
					v, _ := t.AsBigFloat().Int64()
					rawIcmpType = &v
				}
				if c := ruleVal.GetAttr("code"); !c.IsNull() && c.IsKnown() {
					v, _ := c.AsBigFloat().Int64()
					rawIcmpCode = &v
				}
			}

			// tcp block fields
			if hasTcpBlock {
				elem := tcpAttr.Index(cty.NumberIntVal(0))
				if !elem.IsNull() {
					if v := elem.GetAttr("port_min"); !v.IsNull() && v.IsKnown() {
						n, _ := v.AsBigFloat().Int64()
						rawPortMin = &n
					}
					if v := elem.GetAttr("port_max"); !v.IsNull() && v.IsKnown() {
						n, _ := v.AsBigFloat().Int64()
						rawPortMax = &n
					}
					if v := elem.GetAttr("source_port_min"); !v.IsNull() && v.IsKnown() {
						n, _ := v.AsBigFloat().Int64()
						rawSrcPortMin = &n
					}
					if v := elem.GetAttr("source_port_max"); !v.IsNull() && v.IsKnown() {
						n, _ := v.AsBigFloat().Int64()
						rawSrcPortMax = &n
					}
				}
			}

			// udp block fields
			if hasUdpBlock {
				elem := udpAttr.Index(cty.NumberIntVal(0))
				if !elem.IsNull() {
					if v := elem.GetAttr("port_min"); !v.IsNull() && v.IsKnown() {
						n, _ := v.AsBigFloat().Int64()
						rawPortMin = &n
					}
					if v := elem.GetAttr("port_max"); !v.IsNull() && v.IsKnown() {
						n, _ := v.AsBigFloat().Int64()
						rawPortMax = &n
					}
					if v := elem.GetAttr("source_port_min"); !v.IsNull() && v.IsKnown() {
						n, _ := v.AsBigFloat().Int64()
						rawSrcPortMin = &n
					}
					if v := elem.GetAttr("source_port_max"); !v.IsNull() && v.IsKnown() {
						n, _ := v.AsBigFloat().Int64()
						rawSrcPortMax = &n
					}
				}
			}

			// top-level port fields (new style: protocol="tcp"/"udp" + flat ports)
			if !hasTcpBlock && !hasUdpBlock {
				if v := ruleVal.GetAttr("port_min"); !v.IsNull() && v.IsKnown() {
					n, _ := v.AsBigFloat().Int64()
					rawPortMin = &n
				}
				if v := ruleVal.GetAttr("port_max"); !v.IsNull() && v.IsKnown() {
					n, _ := v.AsBigFloat().Int64()
					rawPortMax = &n
				}
				if v := ruleVal.GetAttr("source_port_min"); !v.IsNull() && v.IsKnown() {
					n, _ := v.AsBigFloat().Int64()
					rawSrcPortMin = &n
				}
				if v := ruleVal.GetAttr("source_port_max"); !v.IsNull() && v.IsKnown() {
					n, _ := v.AsBigFloat().Int64()
					rawSrcPortMax = &n
				}
			}
		}
	}

	// Deprecated blocks override whatever the flat protocol attribute says.
	if hasIcmpBlock {
		protocol = "icmp"
	} else if hasTcpBlock {
		protocol = "tcp"
	} else if hasUdpBlock {
		protocol = "udp"
	}
	// action=deny with no explicit protocol means "any"
	if action == "deny" && protocol == "icmp_tcp_udp" {
		protocol = "any"
	}

	ruleTemplate := &vpcv1.NetworkACLRulePrototype{
		Action:      &action,
		Destination: &destination,
		Direction:   &direction,
		Source:      &source,
		Name:        &name,
	}

	if before != "" {
		ruleTemplate.Before = &vpcv1.NetworkACLRuleBeforePrototype{
			ID: &before,
		}
	}

	ruleTemplate.Protocol = &protocol

	switch protocol {
	case "icmp":
		ruleTemplate.Type = rawIcmpType
		ruleTemplate.Code = rawIcmpCode
		if hasIcmpBlock {
			if ruleTemplate.Type != nil && ruleTemplate.Code == nil {
				v := int64(0)
				ruleTemplate.Code = &v
			}
			if ruleTemplate.Code != nil && ruleTemplate.Type == nil {
				v := int64(0)
				ruleTemplate.Type = &v
			}
		}
	case "tcp", "udp":
		ruleTemplate.DestinationPortMin = rawPortMin
		ruleTemplate.DestinationPortMax = rawPortMax
		ruleTemplate.SourcePortMin = rawSrcPortMin
		ruleTemplate.SourcePortMax = rawSrcPortMax
	}

	createNetworkAclRuleOptions := &vpcv1.CreateNetworkACLRuleOptions{
		NetworkACLID:            &nwaclid,
		NetworkACLRulePrototype: ruleTemplate,
	}
	rule, response, err := nwaclC.CreateNetworkACLRule(createNetworkAclRuleOptions)
	if err != nil {
		return "", fmt.Errorf("[ERROR] Error Creating network ACL rule : %s\n%s", err, response)
	}
	newID := ""
	if rule != nil {
		switch r := rule.(type) {
		case *vpcv1.NetworkACLRuleNetworkACLRuleProtocolIcmp:
			newID = *r.ID
		case *vpcv1.NetworkACLRuleNetworkACLRuleProtocolTcpudp:
			newID = *r.ID
		case *vpcv1.NetworkACLRuleNetworkACLRuleProtocolAny:
			newID = *r.ID
		case *vpcv1.NetworkACLRuleNetworkACLRuleProtocolIndividual:
			newID = *r.ID
		case *vpcv1.NetworkACLRuleNetworkACLRuleProtocolIcmptcpudp:
			newID = *r.ID
		case *vpcv1.NetworkACLRule:
			newID = *r.ID
		}
	}
	return newID, nil
}
