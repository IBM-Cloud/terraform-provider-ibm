// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isSecurityGroupRuleCode             = "code"
	isSecurityGroupRuleDirection        = "direction"
	isSecurityGroupRuleIPVersion        = "ip_version"
	isSecurityGroupRuleIPVersionDefault = "ipv4"
	isSecurityGroupRulePortMax          = "port_max"
	isSecurityGroupRulePortMin          = "port_min"
	isSecurityGroupRuleProtocolICMP     = "icmp"
	isSecurityGroupRuleProtocolTCP      = "tcp"
	isSecurityGroupRuleProtocolUDP      = "udp"
	isSecurityGroupRuleProtocol         = "protocol"
	isSecurityGroupRuleRemote           = "remote"
	isSecurityGroupRuleLocal            = "local"
	isSecurityGroupRuleType             = "type"
	isSecurityGroupID                   = "group"
	isSecurityGroupRuleID               = "rule_id"
)

func ResourceIBMISSecurityGroupRule() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceIBMISSecurityGroupRuleCreate,
		ReadContext:   resourceIBMISSecurityGroupRuleRead,
		UpdateContext: resourceIBMISSecurityGroupRuleUpdate,
		DeleteContext: resourceIBMISSecurityGroupRuleDelete,
		Exists:        resourceIBMISSecurityGroupRuleExists,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{

			isSecurityGroupID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Security group id",
				ForceNew:    true,
			},

			isSecurityGroupRuleID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Rule id",
			},

			isSecurityGroupRuleDirection: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Direction of traffic to enforce, either inbound or outbound",
				ValidateFunc: validate.InvokeValidator("ibm_is_security_group_rule", isSecurityGroupRuleDirection),
			},

			isSecurityGroupRuleIPVersion: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "IP version: ipv4",
				Default:      isSecurityGroupRuleIPVersionDefault,
				ValidateFunc: validate.InvokeValidator("ibm_is_security_group_rule", isSecurityGroupRuleIPVersion),
			},

			isSecurityGroupRuleRemote: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Security group local ip: an IP address, a CIDR block",
			},

			isSecurityGroupRuleLocal: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Security group id: an IP address, a CIDR block, or a single security group identifier",
			},

			isSecurityGroupRuleProtocolICMP: {
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				ForceNew:      true,
				MinItems:      1,
				ConflictsWith: []string{isSecurityGroupRuleProtocolTCP, isSecurityGroupRuleProtocolUDP, isSecurityGroupRuleProtocol},
				Deprecated:    "icmp is deprecated, use 'protocol', 'code', and 'type' instead.",
				Description:   "protocol=icmp",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isSecurityGroupRuleType: {
							Type:         schema.TypeInt,
							Optional:     true,
							ForceNew:     false,
							ValidateFunc: validate.InvokeValidator("ibm_is_security_group_rule", isSecurityGroupRuleType),
						},
						isSecurityGroupRuleCode: {
							Type:         schema.TypeInt,
							Optional:     true,
							ForceNew:     false,
							ValidateFunc: validate.InvokeValidator("ibm_is_security_group_rule", isSecurityGroupRuleCode),
						},
					},
				},
			},

			isSecurityGroupRuleProtocolTCP: {
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				MinItems:      1,
				ForceNew:      true,
				Description:   "protocol=tcp",
				Deprecated:    "tcp is deprecated, use 'protocol', 'code', and 'type' instead.",
				ConflictsWith: []string{isSecurityGroupRuleProtocolUDP, isSecurityGroupRuleProtocolICMP, isSecurityGroupRuleProtocol, isSecurityGroupRuleType, isSecurityGroupRuleCode},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isSecurityGroupRulePortMin: {
							Type:         schema.TypeInt,
							Optional:     true,
							ForceNew:     false,
							Default:      1,
							ValidateFunc: validate.InvokeValidator("ibm_is_security_group_rule", isSecurityGroupRulePortMin),
						},
						isSecurityGroupRulePortMax: {
							Type:         schema.TypeInt,
							Optional:     true,
							ForceNew:     false,
							Default:      65535,
							ValidateFunc: validate.InvokeValidator("ibm_is_security_group_rule", isSecurityGroupRulePortMax),
						},
					},
				},
			},

			isSecurityGroupRuleProtocolUDP: {
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				ForceNew:      true,
				MinItems:      1,
				Description:   "protocol=udp",
				Deprecated:    "udp is deprecated, use 'protocol', 'port_min', and 'port_max' instead.",
				ConflictsWith: []string{isSecurityGroupRuleProtocolTCP, isSecurityGroupRuleProtocolICMP, isSecurityGroupRuleProtocol},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isSecurityGroupRulePortMin: {
							Type:         schema.TypeInt,
							Optional:     true,
							ForceNew:     false,
							Default:      1,
							ValidateFunc: validate.InvokeValidator("ibm_is_security_group_rule", isSecurityGroupRulePortMin),
						},
						isSecurityGroupRulePortMax: {
							Type:         schema.TypeInt,
							Optional:     true,
							ForceNew:     false,
							Default:      65535,
							ValidateFunc: validate.InvokeValidator("ibm_is_security_group_rule", isSecurityGroupRulePortMax),
						},
					},
				},
			},

			flex.RelatedCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the Security Group",
			},
			isSecurityGroupRuleProtocol: {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				Description:   "The name of the network protocol",
				ConflictsWith: []string{isSecurityGroupRuleProtocolTCP, isSecurityGroupRuleProtocolICMP, isSecurityGroupRuleProtocolUDP},
				ValidateFunc:  validate.InvokeValidator("ibm_is_security_group_rule", isSecurityGroupRuleProtocol),
			},
			isSecurityGroupRuleType: {
				Type:          schema.TypeInt,
				Optional:      true,
				RequiredWith:  []string{isSecurityGroupRuleProtocol},
				ConflictsWith: []string{isSecurityGroupRuleProtocolICMP, isSecurityGroupRuleProtocolUDP, isSecurityGroupRuleProtocolTCP, isSecurityGroupRulePortMin, isSecurityGroupRulePortMax},
				ValidateFunc:  validate.InvokeValidator("ibm_is_security_group_rule", isSecurityGroupRuleType),
			},
			isSecurityGroupRuleCode: {
				Type:          schema.TypeInt,
				Optional:      true,
				RequiredWith:  []string{isSecurityGroupRuleProtocol},
				ConflictsWith: []string{isSecurityGroupRuleProtocolICMP, isSecurityGroupRuleProtocolUDP, isSecurityGroupRuleProtocolTCP, isSecurityGroupRulePortMin, isSecurityGroupRulePortMax},
				ValidateFunc:  validate.InvokeValidator("ibm_is_security_group_rule", isSecurityGroupRuleCode),
			},
			isSecurityGroupRulePortMin: {
				Type:          schema.TypeInt,
				Optional:      true,
				RequiredWith:  []string{isSecurityGroupRuleProtocol, isSecurityGroupRulePortMax},
				ConflictsWith: []string{isSecurityGroupRuleProtocolICMP, isSecurityGroupRuleProtocolTCP, isSecurityGroupRuleProtocolUDP, isSecurityGroupRuleType, isSecurityGroupRuleCode},
				ValidateFunc:  validate.InvokeValidator("ibm_is_security_group_rule", isSecurityGroupRulePortMin),
			},
			isSecurityGroupRulePortMax: {
				Type:          schema.TypeInt,
				Optional:      true,
				RequiredWith:  []string{isSecurityGroupRuleProtocol, isSecurityGroupRulePortMin},
				ConflictsWith: []string{isSecurityGroupRuleProtocolICMP, isSecurityGroupRuleProtocolTCP, isSecurityGroupRuleProtocolUDP, isSecurityGroupRuleType, isSecurityGroupRuleCode},
				ValidateFunc:  validate.InvokeValidator("ibm_is_security_group_rule", isSecurityGroupRulePortMax),
			},
		},
	}
}

func ResourceIBMISSecurityGroupRuleValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	direction := "inbound, outbound"
	protocol := "tcp, udp, icmp, ah, any, esp, gre, icmp_tcp_udp, ip_in_ip, l2tp, number_0, number_10, number_100, number_101, number_102, number_103, number_104, number_105, number_106, number_107, number_108, number_109, number_11, number_110, number_111, number_113, number_114, number_116, number_117, number_118, number_119, number_12, number_120, number_121, number_122, number_123, number_124, number_125, number_126, number_127, number_128, number_129, number_13, number_130, number_131, number_133, number_134, number_135, number_136, number_137, number_138, number_139, number_14, number_140, number_141, number_142, number_143, number_144, number_145, number_146, number_147, number_148, number_149, number_15, number_150, number_151, number_152, number_153, number_154, number_155, number_156, number_157, number_158, number_159, number_16, number_160, number_161, number_162, number_163, number_164, number_165, number_166, number_167, number_168, number_169, number_170, number_171, number_172, number_173, number_174, number_175, number_176, number_177, number_178, number_179, number_18, number_180, number_181, number_182, number_183, number_184, number_185, number_186, number_187, number_188, number_189, number_19, number_190, number_191, number_192, number_193, number_194, number_195, number_196, number_197, number_198, number_199, number_2, number_20, number_200, number_201, number_202, number_203, number_204, number_205, number_206, number_207, number_208, number_209, number_21, number_210, number_211, number_212, number_213, number_214, number_215, number_216, number_217, number_218, number_219, number_22, number_220, number_221, number_222, number_223, number_224, number_225, number_226, number_227, number_228, number_229, number_23, number_230, number_231, number_232, number_233, number_234, number_235, number_236, number_237, number_238, number_239, number_24, number_240, number_241, number_242, number_243, number_244, number_245, number_246, number_247, number_248, number_249, number_25, number_250, number_251, number_252, number_253, number_254, number_255, number_26, number_27, number_28, number_29, number_3, number_30, number_31, number_32, number_33, number_34, number_35, number_36, number_37, number_38, number_39, number_40, number_41, number_42, number_43, number_44, number_45, number_48, number_49, number_5, number_52, number_53, number_54, number_55, number_56, number_57, number_58, number_59, number_60, number_61, number_62, number_63, number_64, number_65, number_66, number_67, number_68, number_69, number_7, number_70, number_71, number_72, number_73, number_74, number_75, number_76, number_77, number_78, number_79, number_8, number_80, number_81, number_82, number_83, number_84, number_85, number_86, number_87, number_88, number_89, number_9, number_90, number_91, number_92, number_93, number_94, number_95, number_96, number_97, number_98, number_99, rsvp, sctp, vrrp"
	ip_version := "ipv4"

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isSecurityGroupRuleDirection,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              direction})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isSecurityGroupRuleIPVersion,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              ip_version})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isSecurityGroupRuleType,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "0",
			MaxValue:                   "254"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isSecurityGroupRuleCode,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "0",
			MaxValue:                   "255"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isSecurityGroupRulePortMin,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "1",
			MaxValue:                   "65535"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isSecurityGroupRulePortMax,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "1",
			MaxValue:                   "65535"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isSecurityGroupRuleProtocol,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			AllowedValues:              protocol})

	ibmISSecurityGroupRuleResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_security_group_rule", Schema: validateSchema}
	return &ibmISSecurityGroupRuleResourceValidator
}

func resourceIBMISSecurityGroupRuleCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parsed, sgTemplate, _, err := parseIBMISSecurityGroupRuleDictionary(d, "create", sess)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "create", "parse-request-body").GetDiag()
	}
	isSecurityGroupRuleKey := "security_group_rule_key_" + parsed.secgrpID
	conns.IbmMutexKV.Lock(isSecurityGroupRuleKey)
	defer conns.IbmMutexKV.Unlock(isSecurityGroupRuleKey)

	options := &vpcv1.CreateSecurityGroupRuleOptions{
		SecurityGroupID:            &parsed.secgrpID,
		SecurityGroupRulePrototype: sgTemplate,
	}

	rule, _, err := sess.CreateSecurityGroupRuleWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateSecurityGroupRuleWithContext failed: %s", err.Error()), "ibm_is_security_group_rule", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	switch reflect.TypeOf(rule).String() {
	case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolIcmp":
		{
			sgrule := rule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolIcmp)
			d.Set(isSecurityGroupRuleID, *sgrule.ID)
			tfID := makeTerraformRuleID(parsed.secgrpID, *sgrule.ID)
			d.SetId(tfID)
		}
	case "*vpcv1.SecurityGroupRuleProtocolAny":
		{
			sgrule := rule.(*vpcv1.SecurityGroupRuleProtocolAny)
			d.Set(isSecurityGroupRuleID, *sgrule.ID)
			tfID := makeTerraformRuleID(parsed.secgrpID, *sgrule.ID)
			d.SetId(tfID)
		}
	case "*vpcv1.SecurityGroupRuleProtocolIndividual":
		{
			sgrule := rule.(*vpcv1.SecurityGroupRuleProtocolIndividual)
			d.Set(isSecurityGroupRuleID, *sgrule.ID)
			tfID := makeTerraformRuleID(parsed.secgrpID, *sgrule.ID)
			d.SetId(tfID)
		}
	case "*vpcv1.SecurityGroupRuleProtocolIcmptcpudp":
		{
			sgrule := rule.(*vpcv1.SecurityGroupRuleProtocolIcmptcpudp)
			d.Set(isSecurityGroupRuleID, *sgrule.ID)
			tfID := makeTerraformRuleID(parsed.secgrpID, *sgrule.ID)
			d.SetId(tfID)
		}
	case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp":
		{
			sgrule := rule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp)
			d.Set(isSecurityGroupRuleID, *sgrule.ID)
			tfID := makeTerraformRuleID(parsed.secgrpID, *sgrule.ID)
			d.SetId(tfID)
		}
	}
	return resourceIBMISSecurityGroupRuleRead(context, d, meta)
}

func resourceIBMISSecurityGroupRuleRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	secgrpID, ruleID, err := parseISTerraformID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "sep-id-parts").GetDiag()
	}

	getSecurityGroupRuleOptions := &vpcv1.GetSecurityGroupRuleOptions{
		SecurityGroupID: &secgrpID,
		ID:              &ruleID,
	}
	sgrule, response, err := sess.GetSecurityGroupRuleWithContext(context, getSecurityGroupRuleOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSecurityGroupRuleWithContext failed: %s", err.Error()), "ibm_is_security_group_rule", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.Set(isSecurityGroupID, secgrpID)
	getSecurityGroupOptions := &vpcv1.GetSecurityGroupOptions{
		ID: &secgrpID,
	}
	sg, response, err := sess.GetSecurityGroupWithContext(context, getSecurityGroupOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSecurityGroupWithContext failed: %s", err.Error()), "ibm_is_security_group_rule", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(flex.RelatedCRN, *sg.CRN); err != nil {
		err = fmt.Errorf("Error setting related_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-related_crn").GetDiag()
	}
	switch reflect.TypeOf(sgrule).String() {
	case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolIcmp":
		{
			securityGroupRule := sgrule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolIcmp)
			d.Set(isSecurityGroupRuleID, *securityGroupRule.ID)
			tfID := makeTerraformRuleID(secgrpID, *securityGroupRule.ID)
			d.SetId(tfID)
			if err = d.Set("direction", securityGroupRule.Direction); err != nil {
				err = fmt.Errorf("Error setting direction: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-direction").GetDiag()
			}
			if !core.IsNil(securityGroupRule.IPVersion) {
				if err = d.Set("ip_version", securityGroupRule.IPVersion); err != nil {
					err = fmt.Errorf("Error setting ip_version: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-ip_version").GetDiag()
				}
			}
			if err = d.Set("protocol", securityGroupRule.Protocol); err != nil {
				err = fmt.Errorf("Error setting protocol: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-protocol").GetDiag()
			}
			icmpList := d.Get("icmp").([]interface{})
			if len(icmpList) > 0 {
				icmpProtocol := map[string]interface{}{}

				if securityGroupRule.Type != nil {
					icmpProtocol["type"] = *securityGroupRule.Type
				}
				if securityGroupRule.Code != nil {
					icmpProtocol["code"] = *securityGroupRule.Code
				}
				protocolList := make([]map[string]interface{}, 0)
				protocolList = append(protocolList, icmpProtocol)
				if err = d.Set(isSecurityGroupRuleProtocolICMP, protocolList); err != nil {
					err = fmt.Errorf("Error setting icmp: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-icmp").GetDiag()
				}
			}
			remote, ok := securityGroupRule.Remote.(*vpcv1.SecurityGroupRuleRemote)
			if ok {
				if remote != nil && reflect.ValueOf(remote).IsNil() == false {
					if remote.ID != nil {
						if err = d.Set(isSecurityGroupRuleRemote, remote.ID); err != nil {
							err = fmt.Errorf("Error setting remote: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-remote").GetDiag()
						}
					} else if remote.Address != nil {
						if err = d.Set(isSecurityGroupRuleRemote, remote.Address); err != nil {
							err = fmt.Errorf("Error setting remote: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-remote").GetDiag()
						}
					} else if remote.CIDRBlock != nil {
						if err = d.Set(isSecurityGroupRuleRemote, remote.CIDRBlock); err != nil {
							err = fmt.Errorf("Error setting remote: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-remote").GetDiag()
						}
					}
				}
			}
			local, ok := securityGroupRule.Local.(*vpcv1.SecurityGroupRuleLocal)
			if ok {
				if local != nil && reflect.ValueOf(local).IsNil() == false {
					if local.Address != nil {
						if err = d.Set(isSecurityGroupRuleLocal, local.Address); err != nil {
							err = fmt.Errorf("Error setting local: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-local").GetDiag()
						}
					} else if local.CIDRBlock != nil {
						if err = d.Set(isSecurityGroupRuleLocal, local.CIDRBlock); err != nil {
							err = fmt.Errorf("Error setting local: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-local").GetDiag()
						}
					}
				}
			}
		}
	case "*vpcv1.SecurityGroupRuleProtocolAny":
		{
			securityGroupRule := sgrule.(*vpcv1.SecurityGroupRuleProtocolAny)
			d.Set(isSecurityGroupRuleID, *securityGroupRule.ID)
			tfID := makeTerraformRuleID(secgrpID, *securityGroupRule.ID)
			d.SetId(tfID)
			if err = d.Set("direction", securityGroupRule.Direction); err != nil {
				err = fmt.Errorf("Error setting direction: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-direction").GetDiag()
			}
			if !core.IsNil(securityGroupRule.IPVersion) {
				if err = d.Set("ip_version", securityGroupRule.IPVersion); err != nil {
					err = fmt.Errorf("Error setting ip_version: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-ip_version").GetDiag()
				}
			}
			if err = d.Set("protocol", securityGroupRule.Protocol); err != nil {
				err = fmt.Errorf("Error setting protocol: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-protocol").GetDiag()
			}
			remote, ok := securityGroupRule.Remote.(*vpcv1.SecurityGroupRuleRemote)
			if ok {
				if remote != nil && reflect.ValueOf(remote).IsNil() == false {
					if remote.ID != nil {
						if err = d.Set(isSecurityGroupRuleRemote, remote.ID); err != nil {
							err = fmt.Errorf("Error setting remote: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-remote").GetDiag()
						}
					} else if remote.Address != nil {
						if err = d.Set(isSecurityGroupRuleRemote, remote.Address); err != nil {
							err = fmt.Errorf("Error setting remote: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-remote").GetDiag()
						}
					} else if remote.CIDRBlock != nil {
						if err = d.Set(isSecurityGroupRuleRemote, remote.CIDRBlock); err != nil {
							err = fmt.Errorf("Error setting remote: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-remote").GetDiag()
						}
					}
				}
			}
			local, ok := securityGroupRule.Local.(*vpcv1.SecurityGroupRuleLocal)
			if ok {
				if local != nil && reflect.ValueOf(local).IsNil() == false {
					if local.Address != nil {
						if err = d.Set(isSecurityGroupRuleLocal, local.Address); err != nil {
							err = fmt.Errorf("Error setting local: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-local").GetDiag()
						}
					} else if local.CIDRBlock != nil {
						if err = d.Set(isSecurityGroupRuleLocal, local.CIDRBlock); err != nil {
							err = fmt.Errorf("Error setting local: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-local").GetDiag()
						}
					}
				}
			}
		}
	case "*vpcv1.SecurityGroupRuleProtocolIndividual":
		{
			securityGroupRule := sgrule.(*vpcv1.SecurityGroupRuleProtocolIndividual)
			d.Set(isSecurityGroupRuleID, *securityGroupRule.ID)
			tfID := makeTerraformRuleID(secgrpID, *securityGroupRule.ID)
			d.SetId(tfID)
			if err = d.Set("direction", securityGroupRule.Direction); err != nil {
				err = fmt.Errorf("Error setting direction: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-direction").GetDiag()
			}
			if !core.IsNil(securityGroupRule.IPVersion) {
				if err = d.Set("ip_version", securityGroupRule.IPVersion); err != nil {
					err = fmt.Errorf("Error setting ip_version: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-ip_version").GetDiag()
				}
			}
			if err = d.Set("protocol", securityGroupRule.Protocol); err != nil {
				err = fmt.Errorf("Error setting protocol: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-protocol").GetDiag()
			}
			remote, ok := securityGroupRule.Remote.(*vpcv1.SecurityGroupRuleRemote)
			if ok {
				if remote != nil && reflect.ValueOf(remote).IsNil() == false {
					if remote.ID != nil {
						if err = d.Set(isSecurityGroupRuleRemote, remote.ID); err != nil {
							err = fmt.Errorf("Error setting remote: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-remote").GetDiag()
						}
					} else if remote.Address != nil {
						if err = d.Set(isSecurityGroupRuleRemote, remote.Address); err != nil {
							err = fmt.Errorf("Error setting remote: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-remote").GetDiag()
						}
					} else if remote.CIDRBlock != nil {
						if err = d.Set(isSecurityGroupRuleRemote, remote.CIDRBlock); err != nil {
							err = fmt.Errorf("Error setting remote: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-remote").GetDiag()
						}
					}
				}
			}
			local, ok := securityGroupRule.Local.(*vpcv1.SecurityGroupRuleLocal)
			if ok {
				if local != nil && reflect.ValueOf(local).IsNil() == false {
					if local.Address != nil {
						if err = d.Set(isSecurityGroupRuleLocal, local.Address); err != nil {
							err = fmt.Errorf("Error setting local: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-local").GetDiag()
						}
					} else if local.CIDRBlock != nil {
						if err = d.Set(isSecurityGroupRuleLocal, local.CIDRBlock); err != nil {
							err = fmt.Errorf("Error setting local: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-local").GetDiag()
						}
					}
				}
			}
		}

	case "*vpcv1.SecurityGroupRuleProtocolIcmptcpudp":
		{
			securityGroupRule := sgrule.(*vpcv1.SecurityGroupRuleProtocolIcmptcpudp)
			d.Set(isSecurityGroupRuleID, *securityGroupRule.ID)
			tfID := makeTerraformRuleID(secgrpID, *securityGroupRule.ID)
			d.SetId(tfID)
			if err = d.Set("direction", securityGroupRule.Direction); err != nil {
				err = fmt.Errorf("Error setting direction: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-direction").GetDiag()
			}
			if !core.IsNil(securityGroupRule.IPVersion) {
				if err = d.Set("ip_version", securityGroupRule.IPVersion); err != nil {
					err = fmt.Errorf("Error setting ip_version: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-ip_version").GetDiag()
				}
			}
			if err = d.Set("protocol", securityGroupRule.Protocol); err != nil {
				err = fmt.Errorf("Error setting protocol: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-protocol").GetDiag()
			}
			remote, ok := securityGroupRule.Remote.(*vpcv1.SecurityGroupRuleRemote)
			if ok {
				if remote != nil && reflect.ValueOf(remote).IsNil() == false {
					if remote.ID != nil {
						if err = d.Set(isSecurityGroupRuleRemote, remote.ID); err != nil {
							err = fmt.Errorf("Error setting remote: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-remote").GetDiag()
						}
					} else if remote.Address != nil {
						if err = d.Set(isSecurityGroupRuleRemote, remote.Address); err != nil {
							err = fmt.Errorf("Error setting remote: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-remote").GetDiag()
						}
					} else if remote.CIDRBlock != nil {
						if err = d.Set(isSecurityGroupRuleRemote, remote.CIDRBlock); err != nil {
							err = fmt.Errorf("Error setting remote: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-remote").GetDiag()
						}
					}
				}
			}
			local, ok := securityGroupRule.Local.(*vpcv1.SecurityGroupRuleLocal)
			if ok {
				if local != nil && reflect.ValueOf(local).IsNil() == false {
					if local.Address != nil {
						if err = d.Set(isSecurityGroupRuleLocal, local.Address); err != nil {
							err = fmt.Errorf("Error setting local: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-local").GetDiag()
						}
					} else if local.CIDRBlock != nil {
						if err = d.Set(isSecurityGroupRuleLocal, local.CIDRBlock); err != nil {
							err = fmt.Errorf("Error setting local: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-local").GetDiag()
						}
					}
				}
			}
		}
	case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp":
		{
			securityGroupRule := sgrule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp)
			d.Set(isSecurityGroupRuleID, *securityGroupRule.ID)
			tfID := makeTerraformRuleID(secgrpID, *securityGroupRule.ID)
			d.SetId(tfID)
			if err = d.Set("direction", securityGroupRule.Direction); err != nil {
				err = fmt.Errorf("Error setting direction: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-direction").GetDiag()
			}
			if !core.IsNil(securityGroupRule.IPVersion) {
				if err = d.Set("ip_version", securityGroupRule.IPVersion); err != nil {
					err = fmt.Errorf("Error setting ip_version: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-ip_version").GetDiag()
				}
			}
			if err = d.Set("protocol", securityGroupRule.Protocol); err != nil {
				err = fmt.Errorf("Error setting protocol: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-protocol").GetDiag()
			}
			tcpList := d.Get("tcp").([]interface{})
			udpList := d.Get("udp").([]interface{})
			if len(tcpList) > 0 || len(udpList) > 0 {
				tcpProtocol := map[string]interface{}{}

				if securityGroupRule.PortMin != nil {
					tcpProtocol["port_min"] = *securityGroupRule.PortMin
				}
				if securityGroupRule.PortMax != nil {
					tcpProtocol["port_max"] = *securityGroupRule.PortMax
				}
				protocolList := make([]map[string]interface{}, 0)
				protocolList = append(protocolList, tcpProtocol)
				if *securityGroupRule.Protocol == isSecurityGroupRuleProtocolTCP {
					if err = d.Set(isSecurityGroupRuleProtocolTCP, protocolList); err != nil {
						err = fmt.Errorf("Error setting tcp: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-tcp").GetDiag()
					}
				} else {
					if err = d.Set(isSecurityGroupRuleProtocolUDP, protocolList); err != nil {
						err = fmt.Errorf("Error setting udp: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-udp").GetDiag()
					}
				}
			}
			remote, ok := securityGroupRule.Remote.(*vpcv1.SecurityGroupRuleRemote)
			if ok {
				if remote != nil && reflect.ValueOf(remote).IsNil() == false {
					if remote.ID != nil {
						if err = d.Set(isSecurityGroupRuleRemote, remote.ID); err != nil {
							err = fmt.Errorf("Error setting remote: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-remote").GetDiag()
						}
					} else if remote.Address != nil {
						if err = d.Set(isSecurityGroupRuleRemote, remote.Address); err != nil {
							err = fmt.Errorf("Error setting remote: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-remote").GetDiag()
						}
					} else if remote.CIDRBlock != nil {
						if err = d.Set(isSecurityGroupRuleRemote, remote.CIDRBlock); err != nil {
							err = fmt.Errorf("Error setting remote: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-remote").GetDiag()
						}
					}
				}
			}
			local, ok := securityGroupRule.Local.(*vpcv1.SecurityGroupRuleLocal)
			if ok {
				if local != nil && reflect.ValueOf(local).IsNil() == false {
					if local.Address != nil {
						if err = d.Set(isSecurityGroupRuleLocal, local.Address); err != nil {
							err = fmt.Errorf("Error setting local: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-local").GetDiag()
						}
					} else if local.CIDRBlock != nil {
						if err = d.Set(isSecurityGroupRuleLocal, local.CIDRBlock); err != nil {
							err = fmt.Errorf("Error setting local: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "read", "set-local").GetDiag()
						}
					}
				}
			}
		}
	}
	return nil
}

func resourceIBMISSecurityGroupRuleUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parsed, _, sgTemplate, err := parseIBMISSecurityGroupRuleDictionary(d, "update", sess)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "update", "sep-id-parts").GetDiag()
	}
	isSecurityGroupRuleKey := "security_group_rule_key_" + parsed.secgrpID
	conns.IbmMutexKV.Lock(isSecurityGroupRuleKey)
	defer conns.IbmMutexKV.Unlock(isSecurityGroupRuleKey)

	updateSecurityGroupRuleOptions := sgTemplate
	_, _, err = sess.UpdateSecurityGroupRuleWithContext(context, updateSecurityGroupRuleOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSecurityGroupRuleWithContext failed: %s", err.Error()), "ibm_is_security_group_rule", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	return resourceIBMISSecurityGroupRuleRead(context, d, meta)
}

func resourceIBMISSecurityGroupRuleDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	secgrpID, ruleID, err := parseISTerraformID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "delete", "sep-id-parts").GetDiag()
	}

	isSecurityGroupRuleKey := "security_group_rule_key_" + secgrpID
	conns.IbmMutexKV.Lock(isSecurityGroupRuleKey)
	defer conns.IbmMutexKV.Unlock(isSecurityGroupRuleKey)

	getSecurityGroupRuleOptions := &vpcv1.GetSecurityGroupRuleOptions{
		SecurityGroupID: &secgrpID,
		ID:              &ruleID,
	}
	_, response, err := sess.GetSecurityGroupRuleWithContext(context, getSecurityGroupRuleOptions)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSecurityGroupRuleWithContext failed: %s", err.Error()), "ibm_is_security_group_rule", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteSecurityGroupRuleOptions := &vpcv1.DeleteSecurityGroupRuleOptions{
		SecurityGroupID: &secgrpID,
		ID:              &ruleID,
	}
	response, err = sess.DeleteSecurityGroupRuleWithContext(context, deleteSecurityGroupRuleOptions)
	if err != nil && response.StatusCode != 404 {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteSecurityGroupRuleWithContext failed: %s", err.Error()), "ibm_is_security_group_rule", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}

func resourceIBMISSecurityGroupRuleExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_rule", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	secgrpID, ruleID, err := parseISTerraformID(d.Id())
	if err != nil {
		return false, err
	}

	getSecurityGroupRuleOptions := &vpcv1.GetSecurityGroupRuleOptions{
		SecurityGroupID: &secgrpID,
		ID:              &ruleID,
	}
	_, response, err := sess.GetSecurityGroupRule(getSecurityGroupRuleOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSecurityGroupRule failed: %s", err.Error()), "ibm_is_security_group_rule", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}

func parseISTerraformID(s string) (string, string, error) {
	segments := strings.Split(s, ".")
	if len(segments) != 2 {
		return "", "", fmt.Errorf("invalid terraform Id %s (incorrect number of segments)", s)
	}
	if segments[0] == "" || segments[1] == "" {
		return "", "", fmt.Errorf("invalid terraform Id %s (one or more empty segments)", s)
	}
	return segments[0], segments[1], nil
}

type parsedIBMISSecurityGroupRuleDictionary struct {
	// After parsing, unused string fields are set to
	// "" and unused int64 fields will be set to -1.
	// This ("" for unused strings and -1 for unused int64s)
	// is expected by our riaas API client.
	secgrpID       string
	ruleID         string
	direction      string
	ipversion      string
	remote         string
	remoteAddress  string
	remoteCIDR     string
	remoteSecGrpID string
	local          string
	localAddress   string
	localCIDR      string
	protocol       string
	icmpType       int64
	icmpCode       int64
	portMin        int64
	portMax        int64
}

func inferRemoteSecurityGroup(s string) (address, cidr, id string, err error) {
	if validate.IsSecurityGroupAddress(s) {
		address = s
		return
	} else if validate.IsSecurityGroupCIDR(s) {
		cidr = s
		return
	} else {
		id = s
		return
	}
}

func inferLocalSecurityGroup(s string) (address, cidr string, err error) {
	if validate.IsSecurityGroupAddress(s) {
		address = s
		return
	} else if validate.IsSecurityGroupCIDR(s) {
		cidr = s
		return
	}
	return
}

func parseIBMISSecurityGroupRuleDictionary(d *schema.ResourceData, tag string, sess *vpcv1.VpcV1) (*parsedIBMISSecurityGroupRuleDictionary, *vpcv1.SecurityGroupRulePrototype, *vpcv1.UpdateSecurityGroupRuleOptions, error) {
	parsed := &parsedIBMISSecurityGroupRuleDictionary{}
	sgTemplate := &vpcv1.SecurityGroupRulePrototype{}
	sgTemplateUpdate := &vpcv1.UpdateSecurityGroupRuleOptions{}
	var err error
	parsed.icmpType = -1
	parsed.icmpCode = -1
	parsed.portMin = -1
	parsed.portMax = -1

	parsed.secgrpID, parsed.ruleID, err = parseISTerraformID(d.Id())
	if err != nil {
		parsed.secgrpID = d.Get(isSecurityGroupID).(string)
	} else {
		sgTemplateUpdate.SecurityGroupID = &parsed.secgrpID
		sgTemplateUpdate.ID = &parsed.ruleID
	}

	securityGroupRulePatchModel := &vpcv1.SecurityGroupRulePatch{}

	parsed.direction = d.Get(isSecurityGroupRuleDirection).(string)
	sgTemplate.Direction = &parsed.direction
	securityGroupRulePatchModel.Direction = &parsed.direction

	if version, ok := d.GetOk(isSecurityGroupRuleIPVersion); ok {
		parsed.ipversion = version.(string)
		sgTemplate.IPVersion = &parsed.ipversion
		securityGroupRulePatchModel.IPVersion = &parsed.ipversion
	} else {
		parsed.ipversion = "IPv4"
		sgTemplate.IPVersion = &parsed.ipversion
		securityGroupRulePatchModel.IPVersion = &parsed.ipversion
	}

	parsed.remote = ""
	if pr, ok := d.GetOk(isSecurityGroupRuleRemote); ok {
		parsed.remote = pr.(string)
	}
	parsed.remoteAddress = ""
	parsed.remoteCIDR = ""
	parsed.remoteSecGrpID = ""
	err = nil
	if parsed.remote != "" {
		parsed.remoteAddress, parsed.remoteCIDR, parsed.remoteSecGrpID, err = inferRemoteSecurityGroup(parsed.remote)
		remoteTemplate := &vpcv1.SecurityGroupRuleRemotePrototype{}
		remoteTemplateUpdate := &vpcv1.SecurityGroupRuleRemotePatch{}
		if parsed.remoteAddress != "" {
			remoteTemplate.Address = &parsed.remoteAddress
			remoteTemplateUpdate.Address = &parsed.remoteAddress
		} else if parsed.remoteCIDR != "" {
			remoteTemplate.CIDRBlock = &parsed.remoteCIDR
			remoteTemplateUpdate.CIDRBlock = &parsed.remoteCIDR
		} else if parsed.remoteSecGrpID != "" {
			remoteTemplate.ID = &parsed.remoteSecGrpID
			remoteTemplateUpdate.ID = &parsed.remoteSecGrpID

			// check if remote is actually a SG identifier
			getSecurityGroupOptions := &vpcv1.GetSecurityGroupOptions{
				ID: &parsed.remoteSecGrpID,
			}
			sg, res, err := sess.GetSecurityGroup(getSecurityGroupOptions)
			if err != nil || sg == nil {
				if res != nil && res.StatusCode == 404 {
					return nil, nil, nil, fmt.Errorf("[ERROR] Invalid remote provided (%s): %s\n%s", parsed.remoteSecGrpID, err, res)
				}
				return nil, nil, nil, fmt.Errorf("[ERROR] Invalid remote provided (%s): %s", parsed.remoteSecGrpID, err)
			}
		}
		sgTemplate.Remote = remoteTemplate
		securityGroupRulePatchModel.Remote = remoteTemplateUpdate
	}

	if err != nil {
		return nil, nil, nil, err
	}

	parsed.protocol = "icmp_tcp_udp"
	if protocol, ok := d.GetOk(isSecurityGroupRuleProtocol); ok {
		parsed.protocol = protocol.(string)
	}
	sgTemplate.Protocol = &parsed.protocol

	if parsed.protocol != "icmp" {
		if _, ok := d.GetOk("type"); ok {
			return nil, nil, nil, fmt.Errorf("attribute 'type' conflicts with protocol %s; 'type' is only valid for icmp protocol", parsed.protocol)
		}
		if _, ok := d.GetOk("code"); ok {
			return nil, nil, nil, fmt.Errorf("attribute 'code' conflicts with protocol %q; 'code' is only valid for icmp protocol", parsed.protocol)
		}
	}

	if parsed.protocol != "tcp" && parsed.protocol != "udp" {
		if _, ok := d.GetOk("port_min"); ok {
			return nil, nil, nil, fmt.Errorf("attribute 'port_min' conflicts with protocol %s; ports apply only to tcp/udp protocol", parsed.protocol)
		}
		if _, ok := d.GetOk("port_max"); ok {
			return nil, nil, nil, fmt.Errorf("attribute 'port_max' conflicts with protocol %s; ports apply only to tcp/udp protocol", parsed.protocol)
		}
	}

	//Local
	parsed.local = ""
	if pl, ok := d.GetOk(isSecurityGroupRuleLocal); ok {
		parsed.local = pl.(string)
	}
	parsed.localAddress = ""
	parsed.localCIDR = ""
	err = nil
	if parsed.local != "" {
		parsed.localAddress, parsed.localCIDR, err = inferLocalSecurityGroup(parsed.local)
		localTemplate := &vpcv1.SecurityGroupRuleLocalPrototype{}
		localTemplateUpdate := &vpcv1.SecurityGroupRuleLocalPatch{}
		if parsed.localAddress != "" {
			localTemplate.Address = &parsed.localAddress
			localTemplateUpdate.Address = &parsed.localAddress
		} else if parsed.localCIDR != "" {
			localTemplate.CIDRBlock = &parsed.localCIDR
			localTemplateUpdate.CIDRBlock = &parsed.localCIDR
		}
		sgTemplate.Local = localTemplate
		securityGroupRulePatchModel.Local = localTemplateUpdate
	}
	if err != nil {
		return nil, nil, nil, err
	}

	if icmpInterface, ok := d.GetOk("icmp"); ok {
		if icmpInterface.([]interface{})[0] != nil {
			haveType := false
			if value, ok := d.GetOk("icmp.0.type"); ok {
				parsed.icmpType = int64(value.(int))
				haveType = true
				sgTemplate.Type = &parsed.icmpType
				securityGroupRulePatchModel.Type = &parsed.icmpType
			}
			if value, ok := d.GetOk("icmp.0.code"); ok {
				if !haveType {
					return nil, nil, nil, fmt.Errorf("icmp code requires icmp type")
				}
				parsed.icmpCode = int64(value.(int))
				sgTemplate.Code = &parsed.icmpCode
				securityGroupRulePatchModel.Code = &parsed.icmpCode
			}
		}
		parsed.protocol = "icmp"
		sgTemplate.Protocol = &parsed.protocol
	} else {
		if parsed.protocol == "icmp" {
			haveType := false
			if value, ok := d.GetOk("type"); ok {
				parsed.icmpType = int64(value.(int))
				haveType = true
				sgTemplate.Type = &parsed.icmpType
				securityGroupRulePatchModel.Type = &parsed.icmpType
			}
			if value, ok := d.GetOk("code"); ok {
				if !haveType {
					return nil, nil, nil, fmt.Errorf("icmp code requires icmp type")
				}
				parsed.icmpCode = int64(value.(int))
				sgTemplate.Code = &parsed.icmpCode
				securityGroupRulePatchModel.Code = &parsed.icmpCode
			}
		}
	}
	for _, prot := range []string{"tcp", "udp"} {
		if tcpInterface, ok := d.GetOk(prot); ok {
			if tcpInterface.([]interface{})[0] != nil {
				haveMin := false
				haveMax := false
				ports := tcpInterface.([]interface{})[0].(map[string]interface{})
				if value, ok := ports["port_min"]; ok {
					parsed.portMin = int64(value.(int))
					haveMin = true
				}
				if value, ok := ports["port_max"]; ok {
					parsed.portMax = int64(value.(int))
					haveMax = true
				}

				// If only min or max is set, ensure that both min and max are set to the same value
				if haveMin && !haveMax {
					parsed.portMax = parsed.portMin
				}
				if haveMax && !haveMin {
					parsed.portMin = parsed.portMax
				}
			}
			parsed.protocol = prot
			sgTemplate.Protocol = &parsed.protocol
			if tcpInterface.([]interface{})[0] == nil {
				parsed.portMax = 65535
				parsed.portMin = 1
			}
			sgTemplate.PortMax = &parsed.portMax
			sgTemplate.PortMin = &parsed.portMin
			securityGroupRulePatchModel.PortMax = &parsed.portMax
			securityGroupRulePatchModel.PortMin = &parsed.portMin
		} else {
			if parsed.protocol == prot {
				if value, ok := d.GetOk("port_min"); ok {
					parsed.portMin = int64(value.(int))
				}
				if value, ok := d.GetOk("port_max"); ok {
					parsed.portMax = int64(value.(int))
				}
				parsed.protocol = prot
				sgTemplate.Protocol = &parsed.protocol
				if parsed.portMin == -1 && parsed.portMax == -1 {
					sgTemplate.PortMax = nil
					sgTemplate.PortMin = nil
					securityGroupRulePatchModel.PortMax = nil
					securityGroupRulePatchModel.PortMin = nil
				} else {
					sgTemplate.PortMax = &parsed.portMax
					sgTemplate.PortMin = &parsed.portMin
					securityGroupRulePatchModel.PortMax = &parsed.portMax
					securityGroupRulePatchModel.PortMin = &parsed.portMin
				}
			}
		}
	}

	sgTemplate.Protocol = &parsed.protocol
	securityGroupRulePatch, err := securityGroupRulePatchModel.AsPatch()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("[ERROR] Error calling asPatch for SecurityGroupRulePatch: %s", err)
	}
	if _, ok := d.GetOk("icmp"); ok {

		if parsed.icmpType == -1 {
			securityGroupRulePatch["type"] = nil
		}
		if parsed.icmpCode == -1 {
			securityGroupRulePatch["code"] = nil
		}
	}
	sgTemplateUpdate.SecurityGroupRulePatch = securityGroupRulePatch
	//	log.Printf("[DEBUG] parse tag=%s\n\t%v  \n\t%v  \n\t%v  \n\t%v  \n\t%v \n\t%v \n\t%v \n\t%v  \n\t%v  \n\t%v  \n\t%v  \n\t%v ",
	//		tag, parsed.secgrpID, parsed.ruleID, parsed.direction, parsed.ipversion, parsed.protocol, parsed.remoteAddress,
	//		parsed.remoteCIDR, parsed.remoteSecGrpID, parsed.icmpType, parsed.icmpCode, parsed.portMin, parsed.portMax)
	return parsed, sgTemplate, sgTemplateUpdate, nil
}

func makeTerraformRuleID(id1, id2 string) string {
	// Include both group and rule id to create a unique Terraform id.  As a bonus,
	// we can extract the group id as needed for API calls such as READ.
	return id1 + "." + id2
}
