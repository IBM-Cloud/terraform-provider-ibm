// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isNwACLID         = "network_acl"
	isNwACLRuleId     = "rule_id"
	isNwACLRuleBefore = "before"
)

func ResourceIBMISNetworkACLRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISNetworkACLRuleCreate,
		ReadContext:   resourceIBMISNetworkACLRuleRead,
		UpdateContext: resourceIBMISNetworkACLRuleUpdate,
		DeleteContext: resourceIBMISNetworkACLRuleDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				return flex.ResourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{
			isNwACLID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Network ACL id",
			},
			isNwACLRuleId: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The network acl rule id.",
			},
			isNwACLRuleBefore: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The rule that this rule is immediately before. If absent, this is the last rule.",
			},
			isNetworkACLRuleProtocol: {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{isNetworkACLRuleTCP, isNetworkACLRuleUDP, isNetworkACLRuleICMP},
				Description:   "The name of the network protocol",
				ValidateFunc:  validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRuleProtocol),
			},
			isNetworkACLRuleICMPCode: {
				Type:          schema.TypeInt,
				Optional:      true,
				RequiredWith:  []string{isNetworkACLRuleProtocol},
				ConflictsWith: []string{isNetworkACLRuleICMP, isNetworkACLRuleUDP, isNetworkACLRuleTCP, isNetworkACLRulePortMin, isNetworkACLRulePortMax, isNetworkACLRuleSourcePortMax, isNetworkACLRuleSourcePortMin},
				ValidateFunc:  validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRuleICMPCode),
				Description:   "The ICMP traffic code to allow. Valid values from 0 to 255.",
			},
			isNetworkACLRuleICMPType: {
				Type:          schema.TypeInt,
				Optional:      true,
				RequiredWith:  []string{isNetworkACLRuleProtocol},
				ConflictsWith: []string{isNetworkACLRuleICMP, isNetworkACLRuleUDP, isNetworkACLRuleTCP, isNetworkACLRulePortMin, isNetworkACLRulePortMax, isNetworkACLRuleSourcePortMax, isNetworkACLRuleSourcePortMin},
				ValidateFunc:  validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRuleICMPType),
				Description:   "The ICMP traffic type to allow. Valid values from 0 to 254.",
			},
			isNetworkACLRulePortMax: {
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				ValidateFunc:  validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRulePortMax),
				RequiredWith:  []string{isNetworkACLRuleProtocol, isNetworkACLRulePortMin, isNetworkACLRuleSourcePortMax, isNetworkACLRuleSourcePortMin},
				ConflictsWith: []string{isNetworkACLRuleICMP, isNetworkACLRuleUDP, isNetworkACLRuleTCP, isNetworkACLRuleICMPCode, isNetworkACLRuleICMPType},
				Description:   "The highest port in the range of ports to be matched",
			},
			isNetworkACLRulePortMin: {
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				ValidateFunc:  validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRulePortMin),
				RequiredWith:  []string{isNetworkACLRuleProtocol, isNetworkACLRulePortMax, isNetworkACLRuleSourcePortMax, isNetworkACLRuleSourcePortMin},
				ConflictsWith: []string{isNetworkACLRuleICMP, isNetworkACLRuleUDP, isNetworkACLRuleTCP, isNetworkACLRuleICMPCode, isNetworkACLRuleICMPType},
				Description:   "The lowest port in the range of ports to be matched",
			},
			isNetworkACLRuleSourcePortMax: {
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				ValidateFunc:  validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRuleSourcePortMax),
				RequiredWith:  []string{isNetworkACLRuleProtocol, isNetworkACLRulePortMax, isNetworkACLRulePortMin, isNetworkACLRuleSourcePortMin},
				ConflictsWith: []string{isNetworkACLRuleICMP, isNetworkACLRuleUDP, isNetworkACLRuleTCP, isNetworkACLRuleICMPCode, isNetworkACLRuleICMPType},
				Description:   "The highest port in the range of ports to be matched",
			},
			isNetworkACLRuleSourcePortMin: {
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				ValidateFunc:  validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRuleSourcePortMin),
				RequiredWith:  []string{isNetworkACLRuleProtocol, isNetworkACLRulePortMax, isNetworkACLRulePortMin, isNetworkACLRuleSourcePortMax},
				ConflictsWith: []string{isNetworkACLRuleICMP, isNetworkACLRuleUDP, isNetworkACLRuleTCP, isNetworkACLRuleICMPCode, isNetworkACLRuleICMPType},
				Description:   "The lowest port in the range of ports to be matched",
			},
			isNetworkACLRuleHref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The url of the rule.",
			},
			isNetworkACLRuleName: {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     false,
				Description:  "The user-defined name for this rule. Names must be unique within the network ACL the rule resides in. If unspecified, the name will be a hyphenated list of randomly-selected words.",
				ValidateFunc: validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRuleName),
			},
			isNetworkACLRuleAction: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				Description:  "Whether to allow or deny matching traffic",
				ValidateFunc: validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRuleAction),
			},
			isNetworkACLRuleIPVersion: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP version for this rule.",
			},
			isNetworkACLRuleSource: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				Description:  "The source CIDR block. The CIDR block 0.0.0.0/0 applies to all addresses.",
				ValidateFunc: validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRuleSource),
			},
			isNetworkACLRuleDestination: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRuleDestination),
				Description:  "The destination CIDR block. The CIDR block 0.0.0.0/0 applies to all addresses.",
			},
			isNetworkACLRuleDirection: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				Description:  "Direction of traffic to enforce, either inbound or outbound",
				ValidateFunc: validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRuleDirection),
			},
			isNetworkACLRuleICMP: {
				Type:          schema.TypeList,
				MinItems:      0,
				MaxItems:      1,
				Optional:      true,
				ConflictsWith: []string{isNetworkACLRuleTCP, isNetworkACLRuleUDP, isNetworkACLRuleProtocol},
				ForceNew:      true,
				Deprecated:    "icmp is deprecated, use 'protocol', 'code', and 'type' instead.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isNetworkACLRuleICMPCode: {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRuleICMPCode),
							Description:  "The ICMP traffic code to allow. Valid values from 0 to 255.",
						},
						isNetworkACLRuleICMPType: {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRuleICMPType),
							Description:  "The ICMP traffic type to allow. Valid values from 0 to 254.",
						},
					},
				},
			},

			isNetworkACLRuleTCP: {
				Type:          schema.TypeList,
				MinItems:      0,
				MaxItems:      1,
				Optional:      true,
				ConflictsWith: []string{isNetworkACLRuleICMP, isNetworkACLRuleUDP, isNetworkACLRuleProtocol},
				ForceNew:      true,
				Deprecated:    "tcp is deprecated, use 'protocol', 'port_min', 'port_max', 'source_port_min', and 'source_port_max' instead.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isNetworkACLRulePortMax: {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      65535,
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRulePortMax),
							Description:  "The highest port in the range of ports to be matched",
						},
						isNetworkACLRulePortMin: {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      1,
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRulePortMin),
							Description:  "The lowest port in the range of ports to be matched",
						},
						isNetworkACLRuleSourcePortMax: {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      65535,
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRuleSourcePortMax),
							Description:  "The highest port in the range of ports to be matched",
						},
						isNetworkACLRuleSourcePortMin: {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      1,
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRuleSourcePortMin),
							Description:  "The lowest port in the range of ports to be matched",
						},
					},
				},
			},

			isNetworkACLRuleUDP: {
				Type:          schema.TypeList,
				MinItems:      0,
				MaxItems:      1,
				Optional:      true,
				ConflictsWith: []string{isNetworkACLRuleICMP, isNetworkACLRuleTCP, isNetworkACLRuleProtocol},
				ForceNew:      true,
				Deprecated:    "udp is deprecated, use 'protocol', 'port_min', 'port_max', 'source_port_min', and 'source_port_max' instead.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isNetworkACLRulePortMax: {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      65535,
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRulePortMax),
							Description:  "The highest port in the range of ports to be matched",
						},
						isNetworkACLRulePortMin: {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      1,
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRulePortMin),
							Description:  "The lowest port in the range of ports to be matched",
						},
						isNetworkACLRuleSourcePortMax: {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      65535,
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRuleSourcePortMax),
							Description:  "The highest port in the range of ports to be matched",
						},
						isNetworkACLRuleSourcePortMin: {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      1,
							ValidateFunc: validate.InvokeValidator("ibm_is_network_acl_rule", isNetworkACLRuleSourcePortMin),
							Description:  "The lowest port in the range of ports to be matched",
						},
					},
				},
			},
		},
	}
}

func ResourceIBMISNetworkACLRuleValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	direction := "inbound, outbound"
	protocol := "tcp, udp, icmp, ah, any, esp, gre, icmp_tcp_udp, ip_in_ip, l2tp, number_0, number_10, number_100, number_101, number_102, number_103, number_104, number_105, number_106, number_107, number_108, number_109, number_11, number_110, number_111, number_113, number_114, number_116, number_117, number_118, number_119, number_12, number_120, number_121, number_122, number_123, number_124, number_125, number_126, number_127, number_128, number_129, number_13, number_130, number_131, number_133, number_134, number_135, number_136, number_137, number_138, number_139, number_14, number_140, number_141, number_142, number_143, number_144, number_145, number_146, number_147, number_148, number_149, number_15, number_150, number_151, number_152, number_153, number_154, number_155, number_156, number_157, number_158, number_159, number_16, number_160, number_161, number_162, number_163, number_164, number_165, number_166, number_167, number_168, number_169, number_170, number_171, number_172, number_173, number_174, number_175, number_176, number_177, number_178, number_179, number_18, number_180, number_181, number_182, number_183, number_184, number_185, number_186, number_187, number_188, number_189, number_19, number_190, number_191, number_192, number_193, number_194, number_195, number_196, number_197, number_198, number_199, number_2, number_20, number_200, number_201, number_202, number_203, number_204, number_205, number_206, number_207, number_208, number_209, number_21, number_210, number_211, number_212, number_213, number_214, number_215, number_216, number_217, number_218, number_219, number_22, number_220, number_221, number_222, number_223, number_224, number_225, number_226, number_227, number_228, number_229, number_23, number_230, number_231, number_232, number_233, number_234, number_235, number_236, number_237, number_238, number_239, number_24, number_240, number_241, number_242, number_243, number_244, number_245, number_246, number_247, number_248, number_249, number_25, number_250, number_251, number_252, number_253, number_254, number_255, number_26, number_27, number_28, number_29, number_3, number_30, number_31, number_32, number_33, number_34, number_35, number_36, number_37, number_38, number_39, number_40, number_41, number_42, number_43, number_44, number_45, number_48, number_49, number_5, number_52, number_53, number_54, number_55, number_56, number_57, number_58, number_59, number_60, number_61, number_62, number_63, number_64, number_65, number_66, number_67, number_68, number_69, number_7, number_70, number_71, number_72, number_73, number_74, number_75, number_76, number_77, number_78, number_79, number_8, number_80, number_81, number_82, number_83, number_84, number_85, number_86, number_87, number_88, number_89, number_9, number_90, number_91, number_92, number_93, number_94, number_95, number_96, number_97, number_98, number_99, rsvp, sctp, vrrp"
	action := "allow, deny"

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
			Identifier:                 isNwACLID,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
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
			Identifier:                 isSecurityGroupRuleProtocol,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			AllowedValues:              protocol})

	ibmISNetworkACLRuleResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_network_acl_rule", Schema: validateSchema}
	return &ibmISNetworkACLRuleResourceValidator
}

func resourceIBMISNetworkACLRuleCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	nwACLID := d.Get(isNwACLID).(string)

	err := nwaclRuleCreate(context, d, meta, nwACLID)
	if err != nil {
		return err
	}
	return resourceIBMISNetworkACLRuleRead(context, d, meta)

}

func nwaclRuleCreate(context context.Context, d *schema.ResourceData, meta interface{}, nwACLID string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	action := d.Get(isNetworkACLRuleAction).(string)

	direction := d.Get(isNetworkACLRuleDirection).(string)
	// creating rule
	name := d.Get(isNetworkACLRuleName).(string)
	source := d.Get(isNetworkACLRuleSource).(string)
	destination := d.Get(isNetworkACLRuleDestination).(string)
	icmp := d.Get(isNetworkACLRuleICMP).([]interface{})
	tcp := d.Get(isNetworkACLRuleTCP).([]interface{})
	udp := d.Get(isNetworkACLRuleUDP).([]interface{})
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
	if protocolVal, ok := d.GetOk(isNetworkACLRuleProtocol); ok {
		protocol = protocolVal.(string)
	}

	if protocol != "icmp" {
		if _, ok := d.GetOk(isNetworkACLRuleICMPType); ok {
			err = fmt.Errorf("attribute 'type' conflicts with protocol %s; 'type' is only valid for icmp protocol", protocol)
			tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "create", "initialize-client")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if _, ok := d.GetOk(isNetworkACLRuleICMPCode); ok {
			err = fmt.Errorf("attribute 'code' conflicts with protocol %q; 'code' is only valid for icmp protocol", protocol)
			tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "create", "initialize-client")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	if protocol != "tcp" && protocol != "udp" {
		if _, ok := d.GetOk(isNetworkACLRulePortMin); ok {
			err = fmt.Errorf("attribute 'port_min' conflicts with protocol %s; ports apply only to tcp/udp protocol", protocol)
			tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "create", "")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if _, ok := d.GetOk(isNetworkACLRulePortMax); ok {
			err = fmt.Errorf("attribute 'port_max' conflicts with protocol %s; ports apply only to tcp/udp protocol", protocol)
			tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "create", "")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if _, ok := d.GetOk(isNetworkACLRuleSourcePortMax); ok {
			err = fmt.Errorf("attribute 'source_port_max' conflicts with protocol %s; ports apply only to tcp/udp protocol", protocol)
			tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "create", "")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if _, ok := d.GetOk(isNetworkACLRuleSourcePortMin); ok {
			err = fmt.Errorf("attribute 'source_port_min' conflicts with protocol %s; ports apply only to tcp/udp protocol", protocol)
			tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "create", "")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	ruleTemplate := &vpcv1.NetworkACLRulePrototype{
		Action:      &action,
		Destination: &destination,
		Direction:   &direction,
		Source:      &source,
		Name:        &name,
	}

	if before, ok := d.GetOk(isNwACLRuleBefore); ok {
		beforeStr := before.(string)
		ruleTemplate.Before = &vpcv1.NetworkACLRuleBeforePrototype{
			ID: &beforeStr,
		}
	}

	if len(icmp) > 0 {
		protocol = "icmp"
		ruleTemplate.Protocol = &protocol
		if !isNil(icmp[0]) {
			icmpval := icmp[0].(map[string]interface{})
			if val, ok := icmpval[isNetworkACLRuleICMPType]; ok {
				icmptype = int64(val.(int))
				ruleTemplate.Type = &icmptype
			}
			if val, ok := icmpval[isNetworkACLRuleICMPCode]; ok {
				icmpcode = int64(val.(int))
				ruleTemplate.Code = &icmpcode
			}
		}
	} else if protocol == "icmp" {
		ruleTemplate.Protocol = &protocol
		if val, ok := d.GetOk(isNetworkACLRuleICMPType); ok {
			icmptype = int64(val.(int))
			ruleTemplate.Type = &icmptype
		}
		if val, ok := d.GetOk(isNetworkACLRuleICMPCode); ok {
			icmpcode = int64(val.(int))
			ruleTemplate.Code = &icmpcode
		}
	}

	if len(tcp) > 0 {
		protocol = "tcp"
		ruleTemplate.Protocol = &protocol
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
	} else if protocol == "tcp" {
		ruleTemplate.Protocol = &protocol
		if val, ok := d.GetOk(isNetworkACLRulePortMin); ok {
			minport = int64(val.(int))
			ruleTemplate.DestinationPortMin = &minport
		}
		if val, ok := d.GetOk(isNetworkACLRulePortMax); ok {
			maxport = int64(val.(int))
			ruleTemplate.DestinationPortMax = &maxport
		}
		if val, ok := d.GetOk(isNetworkACLRuleSourcePortMin); ok {
			sourceminport = int64(val.(int))
			ruleTemplate.SourcePortMin = &sourceminport
		}
		if val, ok := d.GetOk(isNetworkACLRuleSourcePortMax); ok {
			sourcemaxport = int64(val.(int))
			ruleTemplate.SourcePortMax = &sourcemaxport
		}
		// Setting to nil if no value passed since there is no defaults
		if minport == -1 {
			ruleTemplate.DestinationPortMin = nil
		}
		if maxport == -1 {
			ruleTemplate.DestinationPortMax = nil
		}
		if sourceminport == -1 {
			ruleTemplate.SourcePortMin = nil
		}
		if sourcemaxport == -1 {
			ruleTemplate.SourcePortMax = nil
		}
	}

	if len(udp) > 0 {
		protocol = "udp"
		ruleTemplate.Protocol = &protocol
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
	} else if protocol == "udp" {
		ruleTemplate.Protocol = &protocol
		if val, ok := d.GetOk(isNetworkACLRulePortMin); ok {
			minport = int64(val.(int))
			ruleTemplate.DestinationPortMin = &minport
		}
		if val, ok := d.GetOk(isNetworkACLRulePortMax); ok {
			maxport = int64(val.(int))
			ruleTemplate.DestinationPortMax = &maxport
		}
		if val, ok := d.GetOk(isNetworkACLRuleSourcePortMin); ok {
			sourceminport = int64(val.(int))
			ruleTemplate.SourcePortMin = &sourceminport
		}
		if val, ok := d.GetOk(isNetworkACLRuleSourcePortMax); ok {
			sourcemaxport = int64(val.(int))
			ruleTemplate.SourcePortMax = &sourcemaxport
		}
		// Setting to nil if no value passed since there is no defaults
		if minport == -1 {
			ruleTemplate.DestinationPortMin = nil
		}
		if maxport == -1 {
			ruleTemplate.DestinationPortMax = nil
		}
		if sourceminport == -1 {
			ruleTemplate.SourcePortMin = nil
		}
		if sourcemaxport == -1 {
			ruleTemplate.SourcePortMax = nil
		}
	}

	ruleTemplate.Protocol = &protocol

	createNetworkAclRuleOptions := &vpcv1.CreateNetworkACLRuleOptions{
		NetworkACLID:            &nwACLID,
		NetworkACLRulePrototype: ruleTemplate,
	}
	nwaclRule, _, err := sess.CreateNetworkACLRuleWithContext(context, createNetworkAclRuleOptions)
	if err != nil || nwaclRule == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateNetworkACLRuleWithContext failed: %s", err.Error()), "ibm_is_network_acl_rule", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	diagErr := nwaclRuleGet(context, d, meta, nwACLID, nwaclRule)
	if diagErr != nil {
		return diagErr
	}
	return nil
}

func resourceIBMISNetworkACLRuleRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	nwACLID, ruleId, err := parseNwACLTerraformID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "sep-id-parts").GetDiag()
	}
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getNetworkAclRuleOptions := &vpcv1.GetNetworkACLRuleOptions{
		NetworkACLID: &nwACLID,
		ID:           &ruleId,
	}
	nwaclRule, response, err := sess.GetNetworkACLRuleWithContext(context, getNetworkAclRuleOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetNetworkACLRuleWithContext failed: %s", err.Error()), "ibm_is_network_acl_rule", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	diagErr := nwaclRuleGet(context, d, meta, nwACLID, nwaclRule)
	if diagErr != nil {
		return diagErr
	}
	return nil
}

func nwaclRuleGet(context context.Context, d *schema.ResourceData, meta interface{}, nwACLID string, nwaclRule interface{}) diag.Diagnostics {

	log.Println("[DEBUG] Type of the Rule", reflect.TypeOf(nwaclRule))

	if err := d.Set(isNwACLID, nwACLID); err != nil {
		err = fmt.Errorf("Error setting network_acl: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-network_acl").GetDiag()
	}
	switch reflect.TypeOf(nwaclRule).String() {
	case "*vpcv1.NetworkACLRuleNetworkACLRuleProtocolIcmp":
		{
			var err error
			networkACLRule := nwaclRule.(*vpcv1.NetworkACLRuleNetworkACLRuleProtocolIcmp)
			d.SetId(makeTerraformACLRuleID(nwACLID, *networkACLRule.ID))
			d.Set(isNwACLRuleId, *networkACLRule.ID)
			if networkACLRule.Before != nil {
				if err = d.Set(isNwACLRuleBefore, *networkACLRule.Before.ID); err != nil {
					err = fmt.Errorf("Error setting before: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-before").GetDiag()
				}
			}

			if err = d.Set("name", networkACLRule.Name); err != nil {
				err = fmt.Errorf("Error setting name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-name").GetDiag()
			}
			if err = d.Set("protocol", networkACLRule.Protocol); err != nil {
				err = fmt.Errorf("Error setting protocol: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-protocol").GetDiag()
			}
			if err = d.Set("action", networkACLRule.Action); err != nil {
				err = fmt.Errorf("Error setting action: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-action").GetDiag()
			}
			if !core.IsNil(networkACLRule.IPVersion) {
				if err = d.Set("ip_version", networkACLRule.IPVersion); err != nil {
					err = fmt.Errorf("Error setting ip_version: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-ip_version").GetDiag()
				}
			}
			if err = d.Set("source", networkACLRule.Source); err != nil {
				err = fmt.Errorf("Error setting source: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-source").GetDiag()
			}
			if err = d.Set("destination", networkACLRule.Destination); err != nil {
				err = fmt.Errorf("Error setting destination: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-destination").GetDiag()
			}
			if err = d.Set("direction", networkACLRule.Direction); err != nil {
				err = fmt.Errorf("Error setting direction: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-direction").GetDiag()
			}
			d.Set(isNetworkACLRuleTCP, make([]map[string]int, 0, 0))
			d.Set(isNetworkACLRuleUDP, make([]map[string]int, 0, 0))
			// To determine to set values only for old config and skip for new config as it sets '0' for nil values
			icmpList := d.Get("icmp").([]interface{})
			if len(icmpList) > 0 {
				icmp := make([]map[string]int, 1, 1)
				if networkACLRule.Code != nil && networkACLRule.Type != nil {
					icmp[0] = map[string]int{
						isNetworkACLRuleICMPCode: int(*networkACLRule.Code),
						isNetworkACLRuleICMPType: int(*networkACLRule.Type),
					}
				}
				if err = d.Set(isNetworkACLRuleICMP, icmp); err != nil {
					err = fmt.Errorf("Error setting icmp: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-icmp").GetDiag()
				}
			}
		}
	case "*vpcv1.NetworkACLRuleNetworkACLRuleProtocolTcpudp":
		{
			var err error
			networkACLRule := nwaclRule.(*vpcv1.NetworkACLRuleNetworkACLRuleProtocolTcpudp)
			d.SetId(makeTerraformACLRuleID(nwACLID, *networkACLRule.ID))
			if err = d.Set(isNwACLRuleId, *networkACLRule.ID); err != nil {
				err = fmt.Errorf("Error setting rule_id: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-rule_id").GetDiag()
			}
			if networkACLRule.Before != nil {
				if err = d.Set(isNwACLRuleBefore, *networkACLRule.Before.ID); err != nil {
					err = fmt.Errorf("Error setting before: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-before").GetDiag()
				}
			}
			if err = d.Set("href", networkACLRule.Href); err != nil {
				err = fmt.Errorf("Error setting href: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-href").GetDiag()
			}
			if err = d.Set("protocol", networkACLRule.Protocol); err != nil {
				err = fmt.Errorf("Error setting protocol: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-protocol").GetDiag()
			}
			if err = d.Set("action", networkACLRule.Action); err != nil {
				err = fmt.Errorf("Error setting action: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-action").GetDiag()
			}
			if !core.IsNil(networkACLRule.IPVersion) {
				if err = d.Set("ip_version", networkACLRule.IPVersion); err != nil {
					err = fmt.Errorf("Error setting ip_version: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-ip_version").GetDiag()
				}
			}
			if err = d.Set("source", networkACLRule.Source); err != nil {
				err = fmt.Errorf("Error setting source: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-source").GetDiag()
			}
			if err = d.Set("destination", networkACLRule.Destination); err != nil {
				err = fmt.Errorf("Error setting destination: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-destination").GetDiag()
			}
			if err = d.Set("direction", networkACLRule.Direction); err != nil {
				err = fmt.Errorf("Error setting direction: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-direction").GetDiag()
			}
			// To determine to set values for old or new config. New config is set as it has default values
			tcpList := d.Get("tcp").([]interface{})
			udpList := d.Get("udp").([]interface{})
			if len(tcpList) > 0 || len(udpList) > 0 {
				if *networkACLRule.Protocol == "tcp" {
					d.Set(isNetworkACLRuleICMP, make([]map[string]int, 0, 0))
					d.Set(isNetworkACLRuleUDP, make([]map[string]int, 0, 0))
					tcp := make([]map[string]int, 1, 1)
					tcp[0] = map[string]int{
						isNetworkACLRuleSourcePortMax: checkNetworkACLNil(networkACLRule.SourcePortMax),
						isNetworkACLRuleSourcePortMin: checkNetworkACLNil(networkACLRule.SourcePortMin),
					}
					tcp[0][isNetworkACLRulePortMax] = checkNetworkACLNil(networkACLRule.DestinationPortMax)
					tcp[0][isNetworkACLRulePortMin] = checkNetworkACLNil(networkACLRule.DestinationPortMin)
					if err = d.Set(isNetworkACLRuleTCP, tcp); err != nil {
						err = fmt.Errorf("Error setting tcp: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-tcp").GetDiag()
					}
				} else if *networkACLRule.Protocol == "udp" {
					d.Set(isNetworkACLRuleICMP, make([]map[string]int, 0, 0))
					d.Set(isNetworkACLRuleTCP, make([]map[string]int, 0, 0))
					udp := make([]map[string]int, 1, 1)
					udp[0] = map[string]int{
						isNetworkACLRuleSourcePortMax: checkNetworkACLNil(networkACLRule.SourcePortMax),
						isNetworkACLRuleSourcePortMin: checkNetworkACLNil(networkACLRule.SourcePortMin),
					}
					udp[0][isNetworkACLRulePortMax] = checkNetworkACLNil(networkACLRule.DestinationPortMax)
					udp[0][isNetworkACLRulePortMin] = checkNetworkACLNil(networkACLRule.DestinationPortMin)
					if err = d.Set(isNetworkACLRuleUDP, udp); err != nil {
						err = fmt.Errorf("Error setting udp: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-udp").GetDiag()
					}
				}
			} else {
				d.Set(isNetworkACLRuleSourcePortMax, checkNetworkACLNil(networkACLRule.SourcePortMax))
				d.Set(isNetworkACLRuleSourcePortMin, checkNetworkACLNil(networkACLRule.SourcePortMin))
				d.Set(isNetworkACLRulePortMax, checkNetworkACLNil(networkACLRule.DestinationPortMax))
				d.Set(isNetworkACLRulePortMin, checkNetworkACLNil(networkACLRule.DestinationPortMin))
			}
		}
	case "*vpcv1.NetworkACLRuleNetworkACLRuleProtocolAny":
		{
			var err error
			networkACLRule := nwaclRule.(*vpcv1.NetworkACLRuleNetworkACLRuleProtocolAny)
			d.SetId(makeTerraformACLRuleID(nwACLID, *networkACLRule.ID))
			d.Set(isNwACLRuleId, *networkACLRule.ID)
			if networkACLRule.Before != nil {
				if err = d.Set(isNwACLRuleBefore, *networkACLRule.Before.ID); err != nil {
					err = fmt.Errorf("Error setting before: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-before").GetDiag()
				}
			}
			if err = d.Set("href", networkACLRule.Href); err != nil {
				err = fmt.Errorf("Error setting href: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-href").GetDiag()
			}
			if err = d.Set("protocol", networkACLRule.Protocol); err != nil {
				err = fmt.Errorf("Error setting protocol: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-protocol").GetDiag()
			}
			if err = d.Set("action", networkACLRule.Action); err != nil {
				err = fmt.Errorf("Error setting action: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-action").GetDiag()
			}
			if !core.IsNil(networkACLRule.IPVersion) {
				if err = d.Set("ip_version", networkACLRule.IPVersion); err != nil {
					err = fmt.Errorf("Error setting ip_version: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-ip_version").GetDiag()
				}
			}
			if err = d.Set("source", networkACLRule.Source); err != nil {
				err = fmt.Errorf("Error setting source: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-source").GetDiag()
			}
			if err = d.Set("destination", networkACLRule.Destination); err != nil {
				err = fmt.Errorf("Error setting destination: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-destination").GetDiag()
			}
			if err = d.Set("direction", networkACLRule.Direction); err != nil {
				err = fmt.Errorf("Error setting direction: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-direction").GetDiag()
			}
			d.Set(isNetworkACLRuleICMP, make([]map[string]int, 0, 0))
			d.Set(isNetworkACLRuleTCP, make([]map[string]int, 0, 0))
			d.Set(isNetworkACLRuleUDP, make([]map[string]int, 0, 0))
		}
	case "*vpcv1.NetworkACLRuleNetworkACLRuleProtocolIndividual":
		{
			var err error
			networkACLRule := nwaclRule.(*vpcv1.NetworkACLRuleNetworkACLRuleProtocolIndividual)
			d.SetId(makeTerraformACLRuleID(nwACLID, *networkACLRule.ID))
			d.Set(isNwACLRuleId, *networkACLRule.ID)
			if networkACLRule.Before != nil {
				if err = d.Set(isNwACLRuleBefore, *networkACLRule.Before.ID); err != nil {
					err = fmt.Errorf("Error setting before: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-before").GetDiag()
				}
			}
			if err = d.Set("href", networkACLRule.Href); err != nil {
				err = fmt.Errorf("Error setting href: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-href").GetDiag()
			}
			if err = d.Set("protocol", networkACLRule.Protocol); err != nil {
				err = fmt.Errorf("Error setting protocol: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-protocol").GetDiag()
			}
			if err = d.Set("action", networkACLRule.Action); err != nil {
				err = fmt.Errorf("Error setting action: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-action").GetDiag()
			}
			if !core.IsNil(networkACLRule.IPVersion) {
				if err = d.Set("ip_version", networkACLRule.IPVersion); err != nil {
					err = fmt.Errorf("Error setting ip_version: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-ip_version").GetDiag()
				}
			}
			if err = d.Set("source", networkACLRule.Source); err != nil {
				err = fmt.Errorf("Error setting source: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-source").GetDiag()
			}
			if err = d.Set("destination", networkACLRule.Destination); err != nil {
				err = fmt.Errorf("Error setting destination: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-destination").GetDiag()
			}
			if err = d.Set("direction", networkACLRule.Direction); err != nil {
				err = fmt.Errorf("Error setting direction: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-direction").GetDiag()
			}
			d.Set(isNetworkACLRuleICMP, make([]map[string]int, 0, 0))
			d.Set(isNetworkACLRuleTCP, make([]map[string]int, 0, 0))
			d.Set(isNetworkACLRuleUDP, make([]map[string]int, 0, 0))
		}
	case "*vpcv1.NetworkACLRuleNetworkACLRuleProtocolIcmptcpudp":
		{
			var err error
			networkACLRule := nwaclRule.(*vpcv1.NetworkACLRuleNetworkACLRuleProtocolIcmptcpudp)
			d.SetId(makeTerraformACLRuleID(nwACLID, *networkACLRule.ID))
			d.Set(isNwACLRuleId, *networkACLRule.ID)
			if networkACLRule.Before != nil {
				if err = d.Set(isNwACLRuleBefore, *networkACLRule.Before.ID); err != nil {
					err = fmt.Errorf("Error setting before: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-before").GetDiag()
				}
			}
			if err = d.Set("href", networkACLRule.Href); err != nil {
				err = fmt.Errorf("Error setting href: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-href").GetDiag()
			}
			if err = d.Set("protocol", networkACLRule.Protocol); err != nil {
				err = fmt.Errorf("Error setting protocol: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-protocol").GetDiag()
			}
			if err = d.Set("action", networkACLRule.Action); err != nil {
				err = fmt.Errorf("Error setting action: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-action").GetDiag()
			}
			if !core.IsNil(networkACLRule.IPVersion) {
				if err = d.Set("ip_version", networkACLRule.IPVersion); err != nil {
					err = fmt.Errorf("Error setting ip_version: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-ip_version").GetDiag()
				}
			}
			if err = d.Set("source", networkACLRule.Source); err != nil {
				err = fmt.Errorf("Error setting source: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-source").GetDiag()
			}
			if err = d.Set("destination", networkACLRule.Destination); err != nil {
				err = fmt.Errorf("Error setting destination: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-destination").GetDiag()
			}
			if err = d.Set("direction", networkACLRule.Direction); err != nil {
				err = fmt.Errorf("Error setting direction: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "read", "set-direction").GetDiag()
			}
			d.Set(isNetworkACLRuleICMP, make([]map[string]int, 0, 0))
			d.Set(isNetworkACLRuleTCP, make([]map[string]int, 0, 0))
			d.Set(isNetworkACLRuleUDP, make([]map[string]int, 0, 0))
		}

	}

	return nil
}

func resourceIBMISNetworkACLRuleUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	nwACLId, ruleId, err := parseNwACLTerraformID(id)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "update", "sep-id-parts").GetDiag()
	}

	diagErr := nwaclRuleUpdate(context, d, meta, ruleId, nwACLId)
	if diagErr != nil {
		return diagErr
	}
	return resourceIBMISNetworkACLRuleRead(context, d, meta)
}

func nwaclRuleUpdate(context context.Context, d *schema.ResourceData, meta interface{}, id, nwACLId string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	updateNetworkACLRuleOptions := &vpcv1.UpdateNetworkACLRuleOptions{
		NetworkACLID: &nwACLId,
		ID:           &id,
	}
	updateNetworkACLOptionsPatchModel := &vpcv1.NetworkACLRulePatch{}

	hasChanged := false

	if d.HasChange(isNetworkACLRuleAction) {
		hasChanged = true
		if actionVar, ok := d.GetOk(isNetworkACLRuleAction); ok {
			action := actionVar.(string)
			updateNetworkACLOptionsPatchModel.Action = &action
		}
	}
	aclRuleBeforeNull := false
	if d.HasChange(isNwACLRuleBefore) {
		hasChanged = true
		beforeVar := d.Get(isNwACLRuleBefore).(string)
		if beforeVar == "null" {
			aclRuleBeforeNull = true
		} else if beforeVar != "" {
			updateNetworkACLOptionsPatchModel.Before = &vpcv1.NetworkACLRuleBeforePatchNetworkACLRuleIdentityByID{
				ID: &beforeVar,
			}
		}
	}

	if d.HasChange(isNetworkACLRuleName) {
		hasChanged = true
		if nameVar, ok := d.GetOk(isNetworkACLRuleName); ok {
			nameStr := nameVar.(string)
			updateNetworkACLOptionsPatchModel.Name = &nameStr
		}
	}
	if d.HasChange(isNetworkACLRuleDirection) {
		hasChanged = true
		if directionVar, ok := d.GetOk(isNetworkACLRuleDirection); ok {
			directionStr := directionVar.(string)
			updateNetworkACLOptionsPatchModel.Direction = &directionStr
		}
	}
	if d.HasChange(isNetworkACLRuleDestination) {
		hasChanged = true
		if destinationVar, ok := d.GetOk(isNetworkACLRuleDestination); ok {
			destination := destinationVar.(string)
			updateNetworkACLOptionsPatchModel.Destination = &destination
		}
	}
	if d.HasChange(isNetworkACLRuleICMP) {
		icmpCode := fmt.Sprint(isNetworkACLRuleICMP, ".0.", isNetworkACLRuleICMPCode)
		icmpType := fmt.Sprint(isNetworkACLRuleICMP, ".0.", isNetworkACLRuleICMPType)
		if d.HasChange(icmpCode) {
			hasChanged = true
			if codeVar, ok := d.GetOk(icmpCode); ok {
				code := int64(codeVar.(int))
				updateNetworkACLOptionsPatchModel.Code = &code
			}
		}
		if d.HasChange(icmpType) {
			hasChanged = true
			if typeVar, ok := d.GetOk(icmpType); ok {
				typeInt := int64(typeVar.(int))
				updateNetworkACLOptionsPatchModel.Type = &typeInt
			}
		}
	}

	if d.HasChange(isNetworkACLRuleICMPCode) {
		hasChanged = true
		if codeVar, ok := d.GetOk(isNetworkACLRuleICMPCode); ok {
			code := int64(codeVar.(int))
			updateNetworkACLOptionsPatchModel.Code = &code
		}
	}
	if d.HasChange(isNetworkACLRuleICMPType) {
		hasChanged = true
		if typeVar, ok := d.GetOk(isNetworkACLRuleICMPType); ok {
			typeInt := int64(typeVar.(int))
			updateNetworkACLOptionsPatchModel.Type = &typeInt
		}
	}

	if d.HasChange(isNetworkACLRuleTCP) {
		tcp := d.Get(isNetworkACLRuleTCP).([]interface{})
		tcpval := tcp[0].(map[string]interface{})
		max := fmt.Sprint(isNetworkACLRuleTCP, ".0.", isNetworkACLRulePortMax)
		min := fmt.Sprint(isNetworkACLRuleTCP, ".0.", isNetworkACLRulePortMin)
		maxSource := fmt.Sprint(isNetworkACLRuleTCP, ".0.", isNetworkACLRuleSourcePortMax)
		minSource := fmt.Sprint(isNetworkACLRuleTCP, ".0.", isNetworkACLRuleSourcePortMin)
		if d.HasChange(max) {
			hasChanged = true
			if destinationVar, ok := tcpval[isNetworkACLRulePortMax]; ok {
				destination := int64(destinationVar.(int))
				updateNetworkACLOptionsPatchModel.DestinationPortMax = &destination
			}
		}
		if d.HasChange(min) {
			hasChanged = true
			if destinationVar, ok := tcpval[isNetworkACLRulePortMin]; ok {
				destination := int64(destinationVar.(int))
				updateNetworkACLOptionsPatchModel.DestinationPortMin = &destination
			}
		}
		if d.HasChange(maxSource) {
			hasChanged = true
			if sourceVar, ok := tcpval[isNetworkACLRuleSourcePortMax]; ok {
				source := int64(sourceVar.(int))
				updateNetworkACLOptionsPatchModel.SourcePortMax = &source
			}
		}
		if d.HasChange(minSource) {
			hasChanged = true
			if sourceVar, ok := tcpval[isNetworkACLRuleSourcePortMin]; ok {
				source := int64(sourceVar.(int))
				updateNetworkACLOptionsPatchModel.SourcePortMin = &source
			}
		}
	}
	if d.HasChange(isNetworkACLRulePortMax) {
		hasChanged = true
		if destinationVar, ok := d.GetOk(isNetworkACLRulePortMax); ok {
			destination := int64(destinationVar.(int))
			updateNetworkACLOptionsPatchModel.DestinationPortMax = &destination
		}
	}
	if d.HasChange(isNetworkACLRulePortMin) {
		hasChanged = true
		if destinationVar, ok := d.GetOk(isNetworkACLRulePortMin); ok {
			destination := int64(destinationVar.(int))
			updateNetworkACLOptionsPatchModel.DestinationPortMin = &destination
		}
	}
	if d.HasChange(isNetworkACLRuleSourcePortMax) {
		hasChanged = true
		if sourceVar, ok := d.GetOk(isNetworkACLRuleSourcePortMax); ok {
			source := int64(sourceVar.(int))
			updateNetworkACLOptionsPatchModel.SourcePortMax = &source
		}
	}
	if d.HasChange(isNetworkACLRuleSourcePortMin) {
		hasChanged = true
		if sourceVar, ok := d.GetOk(isNetworkACLRuleSourcePortMin); ok {
			source := int64(sourceVar.(int))
			updateNetworkACLOptionsPatchModel.SourcePortMin = &source
		}
	}
	if d.HasChange(isNetworkACLRuleUDP) {
		udp := d.Get(isNetworkACLRuleUDP).([]interface{})
		udpval := udp[0].(map[string]interface{})
		max := fmt.Sprint(isNetworkACLRuleUDP, ".0.", isNetworkACLRulePortMax)
		min := fmt.Sprint(isNetworkACLRuleUDP, ".0.", isNetworkACLRulePortMin)
		maxSource := fmt.Sprint(isNetworkACLRuleUDP, ".0.", isNetworkACLRuleSourcePortMax)
		minSource := fmt.Sprint(isNetworkACLRuleUDP, ".0.", isNetworkACLRuleSourcePortMin)

		if d.HasChange(max) {
			hasChanged = true
			if destinationVar, ok := udpval[isNetworkACLRulePortMax]; ok {
				destination := int64(destinationVar.(int))
				updateNetworkACLOptionsPatchModel.DestinationPortMax = &destination
			}
		}
		if d.HasChange(min) {
			hasChanged = true
			if destinationVar, ok := udpval[isNetworkACLRulePortMin]; ok {
				destination := int64(destinationVar.(int))
				updateNetworkACLOptionsPatchModel.DestinationPortMin = &destination
			}
		}
		if d.HasChange(maxSource) {
			hasChanged = true
			if sourceVar, ok := udpval[isNetworkACLRuleSourcePortMax]; ok {
				source := int64(sourceVar.(int))
				updateNetworkACLOptionsPatchModel.SourcePortMax = &source
			}
		}
		if d.HasChange(minSource) {
			hasChanged = true
			if sourceVar, ok := udpval[isNetworkACLRuleSourcePortMin]; ok {
				source := int64(sourceVar.(int))
				updateNetworkACLOptionsPatchModel.SourcePortMin = &source
			}
		}
	}

	if d.HasChange(isNetworkACLRuleSource) {
		hasChanged = true
		if sourceVar, ok := d.GetOk(isNetworkACLRuleSource); ok {
			source := sourceVar.(string)
			updateNetworkACLOptionsPatchModel.Source = &source
		}
	}

	if hasChanged {
		updateNetworkACLOptionsPatch, err := updateNetworkACLOptionsPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("updateNetworkACLOptionsPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_network_acl_rule", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if aclRuleBeforeNull {
			updateNetworkACLOptionsPatch["before"] = nil
		}
		updateNetworkACLRuleOptions.NetworkACLRulePatch = updateNetworkACLOptionsPatch
		_, _, err = sess.UpdateNetworkACLRuleWithContext(context, updateNetworkACLRuleOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateNetworkACLRuleWithContext failed: %s", err.Error()), "ibm_is_network_acl_rule", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return nil
}

func resourceIBMISNetworkACLRuleDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	nwACLID, ruleId, err := parseNwACLTerraformID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "delete", "sep-id-parts").GetDiag()
	}

	diagErr := nwaclRuleDelete(context, d, meta, ruleId, nwACLID)
	if diagErr != nil {
		return diagErr
	}
	d.SetId("")
	return nil
}

func nwaclRuleDelete(context context.Context, d *schema.ResourceData, meta interface{}, id, nwACLId string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getNetworkAclRuleOptions := &vpcv1.GetNetworkACLRuleOptions{
		NetworkACLID: &nwACLId,
		ID:           &id,
	}
	_, response, err := sess.GetNetworkACLRuleWithContext(context, getNetworkAclRuleOptions)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetNetworkACLRuleWithContext failed: %s", err.Error()), "ibm_is_network_acl_rule", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteNetworkAclRuleOptions := &vpcv1.DeleteNetworkACLRuleOptions{
		NetworkACLID: &nwACLId,
		ID:           &id,
	}
	response, err = sess.DeleteNetworkACLRuleWithContext(context, deleteNetworkAclRuleOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteNetworkACLRuleWithContext failed: %s", err.Error()), "ibm_is_network_acl_rule", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}

func nwaclRuleExists(d *schema.ResourceData, meta interface{}, id, nwACLId string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_network_acl_rule", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	getNetworkAclRuleOptions := &vpcv1.GetNetworkACLRuleOptions{
		NetworkACLID: &nwACLId,
		ID:           &id,
	}
	_, response, err := sess.GetNetworkACLRule(getNetworkAclRuleOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetNetworkACLRule failed: %s", err.Error()), "ibm_is_network_acl_rule", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}

func makeTerraformACLRuleID(id1, id2 string) string {
	// Include both network acl id and rule id to create a unique Terraform id.  As a bonus,
	// we can extract the network acl id as needed for API calls such as READ.
	return fmt.Sprintf("%s/%s", id1, id2)
}

func parseNwACLTerraformID(s string) (string, string, error) {
	segments := strings.Split(s, "/")
	if len(segments) != 2 {
		return "", "", fmt.Errorf("invalid terraform Id %s (incorrect number of segments)", s)
	}
	if segments[0] == "" || segments[1] == "" {
		return "", "", fmt.Errorf("invalid terraform Id %s (one or more empty segments)", s)
	}
	return segments[0], segments[1], nil
}
