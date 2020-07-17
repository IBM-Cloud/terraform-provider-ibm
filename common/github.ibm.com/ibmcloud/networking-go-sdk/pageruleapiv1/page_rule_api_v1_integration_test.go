/*
 * (C) Copyright IBM Corp. 2020.
 */

package pageruleapiv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/pageruleapiv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`pageruleapiv1`, func() {
	if _, err := os.Stat(configFile); err != nil {
		configLoaded = false
	}

	err := godotenv.Load(configFile)
	if err != nil {
		configLoaded = false
	}

	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("CIS_SERVICES_APIKEY"),
		URL:    os.Getenv("CIS_SERVICES_AUTH_URL"),
	}
	serviceURL := os.Getenv("API_ENDPOINT")
	crn := os.Getenv("CRN")
	zone_id := os.Getenv("ZONE_ID")
	url := os.Getenv("URL_MATCH")
	url_change := os.Getenv("CHANGE_URL_MATCH")
	globalOptions := &PageRuleApiV1Options{
		ServiceName:   "cis_services",
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
		ZoneID:        &zone_id,
	}

	service, serviceErr := NewPageRuleApiV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`create/update/change/delete/get page rule`, func() {
		Context(`create/update/change/delete/get page`, func() {
			BeforeEach(func() {
				shouldSkipTest()
				listOpt := service.NewListPageRulesOptions()
				listResult, listResp, listErr := service.ListPageRules(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				for _, rule := range listResult.Result {
					delOpt := service.NewDeletePageRuleOptions(*rule.ID)
					delResult, delResp, delErr := service.DeletePageRule(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				listOpt := service.NewListPageRulesOptions()
				listResult, listResp, listErr := service.ListPageRules(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				for _, rule := range listResult.Result {
					delOpt := service.NewDeletePageRuleOptions(*rule.ID)
					delResult, delResp, delErr := service.DeletePageRule(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
			It(`create/update/change/delete/get page rule url value`, func() {
				shouldSkipTest()

				// create page rule
				targetConstraintOpt := TargetsItemConstraint{
					Operator: core.StringPtr("matches"),
					Value:    &url,
				}
				targetOpt := []TargetsItem{
					{
						Target:     core.StringPtr("url"),
						Constraint: &targetConstraintOpt,
					}}

				pageRuleItem1 := PageRulesBodyActionsItem{
					ID: core.StringPtr(PageRulesBodyActionsItem_ID_DisableSecurity),
				}
				pageRuleItem2 := PageRulesBodyActionsItem{
					ID:    core.StringPtr(PageRulesBodyActionsItem_ID_BrowserCheck),
					Value: PageRulesBodyActionsItemActionsSecurityOptions_Value_On,
				}
				createOpt := service.NewCreatePageRuleOptions()
				createOpt.SetTargets(targetOpt)
				createOpt.SetActions([]PageRulesBodyActionsItemIntf{&pageRuleItem1, &pageRuleItem2})
				createOpt.SetPriority(1)
				createOpt.SetStatus("active")

				createResult, createResp, createErr := service.CreatePageRule(createOpt)
				Expect(createErr).To(BeNil())
				Expect(createResp).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).Should(BeTrue())

				// get page rule by id
				getOpt := service.NewGetPageRuleOptions(*createResult.Result.ID)
				getResult, getResp, getErr := service.GetPageRule(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				// update page rule
				pageRuleItem2 = PageRulesBodyActionsItem{
					ID:    core.StringPtr(PageRulesBodyActionsItem_ID_BrowserCheck),
					Value: PageRulesBodyActionsItemActionsSecurityOptions_Value_Off,
				}

				updateOpt := service.NewUpdatePageRuleOptions(*createResult.Result.ID)
				updateOpt.SetTargets(targetOpt)
				updateOpt.SetActions([]PageRulesBodyActionsItemIntf{&pageRuleItem1, &pageRuleItem2})
				updateOpt.SetPriority(5)
				updateOpt.SetStatus("active")

				targetConstraintOpt = TargetsItemConstraint{
					Operator: core.StringPtr("matches"),
					Value:    &url_change,
				}
				targetOpt = []TargetsItem{
					{
						Target:     core.StringPtr("url"),
						Constraint: &targetConstraintOpt,
					},
				}

				updateResult, updateResp, updateErr := service.UpdatePageRule(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())

				// change page rule
				pageRuleItem2 = PageRulesBodyActionsItem{
					ID:    core.StringPtr(PageRulesBodyActionsItem_ID_BrowserCheck),
					Value: PageRulesBodyActionsItemActionsSecurityOptions_Value_On,
				}

				changeOpt := service.NewChangePageRuleOptions(*createResult.Result.ID)
				changeOpt.SetTargets(targetOpt)
				changeOpt.SetActions([]PageRulesBodyActionsItemIntf{&pageRuleItem1, &pageRuleItem2})
				changeOpt.SetPriority(5)
				changeOpt.SetStatus("active")

				changeResult, changeResp, changeErr := service.ChangePageRule(changeOpt)
				Expect(changeErr).To(BeNil())
				Expect(changeResp).ToNot(BeNil())
				Expect(changeResult).ToNot(BeNil())
				Expect(*changeResult.Success).Should(BeTrue())

				delOpt := service.NewDeletePageRuleOptions(*createResult.Result.ID)
				delResult, delResp, delErr := service.DeletePageRule(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*delResult.Success).Should(BeTrue())

			})

			It(`create/update/change/delete/get page rule forwarding url`, func() {
				shouldSkipTest()
				targetConstraintOpt := TargetsItemConstraint{
					Operator: core.StringPtr("matches"),
					Value:    &url,
				}
				targetOpt := []TargetsItem{
					{
						Target:     core.StringPtr("url"),
						Constraint: &targetConstraintOpt,
					}}

				urlChange := fmt.Sprintf("https://%s", url_change)
				ActionsForwardingUrlValueOpt, err := service.NewPageRulesBodyActionsItemActionsForwardingURL(PageRulesBodyActionsItem_ID_ForwardingURL)
				Expect(err).To(BeNil())
				ActionsForwardingUrlValueOpt.Value = &ActionsForwardingUrlValue{
					URL:        &urlChange,
					StatusCode: core.Int64Ptr(302),
				}
				pageRuleItem1 := PageRulesBodyActionsItem{
					ID:    ActionsForwardingUrlValueOpt.ID,
					Value: ActionsForwardingUrlValueOpt.Value,
				}
				createOpt := service.NewCreatePageRuleOptions()
				createOpt.SetTargets(targetOpt)
				createOpt.SetActions([]PageRulesBodyActionsItemIntf{&pageRuleItem1})
				createOpt.SetPriority(1)
				createOpt.SetStatus("active")

				createResult, createResp, createErr := service.CreatePageRule(createOpt)
				Expect(createErr).To(BeNil())
				Expect(createResp).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).Should(BeTrue())

				// update page rule
				targetConstraintOpt = TargetsItemConstraint{
					Operator: core.StringPtr("matches"),
					Value:    &url_change,
				}
				targetOpt = []TargetsItem{
					{
						Target:     core.StringPtr("url"),
						Constraint: &targetConstraintOpt,
					}}

				urlChange = fmt.Sprintf("https://%s", url)
				ActionsForwardingUrlValueOpt, err = service.NewPageRulesBodyActionsItemActionsForwardingURL(PageRulesBodyActionsItem_ID_ForwardingURL)
				Expect(err).To(BeNil())
				ActionsForwardingUrlValueOpt.Value = &ActionsForwardingUrlValue{
					URL:        &urlChange,
					StatusCode: core.Int64Ptr(301),
				}
				updateOpt := service.NewUpdatePageRuleOptions(*createResult.Result.ID)
				updateOpt.SetTargets(targetOpt)
				updateOpt.SetActions([]PageRulesBodyActionsItemIntf{&pageRuleItem1})
				updateOpt.SetPriority(1)
				updateOpt.SetStatus("active")

				updateResult, updateResp, updateErr := service.UpdatePageRule(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())

				// change page rule
				targetConstraintOpt = TargetsItemConstraint{
					Operator: core.StringPtr("matches"),
					Value:    &url,
				}
				targetOpt = []TargetsItem{
					{
						Target:     core.StringPtr("url"),
						Constraint: &targetConstraintOpt,
					}}
				urlChange = fmt.Sprintf("https://%s", url_change)
				ActionsForwardingUrlValueOpt, err = service.NewPageRulesBodyActionsItemActionsForwardingURL(PageRulesBodyActionsItem_ID_ForwardingURL)
				Expect(err).To(BeNil())
				ActionsForwardingUrlValueOpt.Value = &ActionsForwardingUrlValue{
					URL:        &urlChange,
					StatusCode: core.Int64Ptr(302),
				}
				changeOpt := service.NewChangePageRuleOptions(*createResult.Result.ID)
				changeOpt.SetTargets(targetOpt)
				changeOpt.SetActions([]PageRulesBodyActionsItemIntf{&pageRuleItem1})
				changeOpt.SetPriority(1)
				changeOpt.SetStatus("active")

				changeResult, changeResp, changeErr := service.ChangePageRule(changeOpt)
				Expect(changeErr).To(BeNil())
				Expect(changeResp).ToNot(BeNil())
				Expect(changeResult).ToNot(BeNil())
				Expect(*changeResult.Success).Should(BeTrue())

				getOpt := service.NewGetPageRuleOptions(*createResult.Result.ID)
				getResult, getResp, getErr := service.GetPageRule(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				delOpt := service.NewDeletePageRuleOptions(*createResult.Result.ID)
				delResult, delResp, delErr := service.DeletePageRule(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*delResult.Success).Should(BeTrue())
			})
			It(`create/update/change/delete/get page rule security level`, func() {
				shouldSkipTest()

				// create page rule
				targetConstraintOpt := TargetsItemConstraint{
					Operator: core.StringPtr("matches"),
					Value:    &url,
				}
				targetOpt := []TargetsItem{
					{
						Target:     core.StringPtr("url"),
						Constraint: &targetConstraintOpt,
					}}

				security_levels := []string{
					PageRulesBodyActionsItemActionsSecurityLevel_Value_EssentiallyOff,
					PageRulesBodyActionsItemActionsSecurityLevel_Value_High,
					PageRulesBodyActionsItemActionsSecurityLevel_Value_Low,
					PageRulesBodyActionsItemActionsSecurityLevel_Value_Medium,
					PageRulesBodyActionsItemActionsSecurityLevel_Value_UnderAttack,
				}
				pageRuleItem1 := PageRulesBodyActionsItem{
					ID:    core.StringPtr(PageRulesBodyActionsItem_ID_SecurityLevel),
					Value: core.StringPtr(PageRulesBodyActionsItemActionsSecurityLevel_Value_Off),
				}
				createOpt := service.NewCreatePageRuleOptions()
				createOpt.SetTargets(targetOpt)
				createOpt.SetActions([]PageRulesBodyActionsItemIntf{
					&pageRuleItem1,
				})
				createOpt.SetPriority(1)
				createOpt.SetStatus("active")

				createResult, createResp, createErr := service.CreatePageRule(createOpt)
				Expect(createErr).To(BeNil())
				Expect(createResp).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).Should(BeTrue())

				for _, level := range security_levels {

					// update page rule
					pageRuleItem1 = PageRulesBodyActionsItem{
						ID:    core.StringPtr(PageRulesBodyActionsItem_ID_SecurityLevel),
						Value: core.StringPtr(level),
					}
					updateOpt := service.NewUpdatePageRuleOptions(*createResult.Result.ID)
					updateOpt.SetTargets(targetOpt)
					updateOpt.SetActions([]PageRulesBodyActionsItemIntf{
						&pageRuleItem1,
					})
					updateOpt.SetPriority(1)
					updateOpt.SetStatus("active")

					updateResult, updateResp, updateErr := service.UpdatePageRule(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())
				}

				// change page rule
				for _, level := range security_levels {
					pageRuleItem1 = PageRulesBodyActionsItem{
						ID:    core.StringPtr(PageRulesBodyActionsItem_ID_SecurityLevel),
						Value: core.StringPtr(level),
					}
					changeOpt := service.NewChangePageRuleOptions(*createResult.Result.ID)
					changeOpt.SetTargets(targetOpt)
					changeOpt.SetActions([]PageRulesBodyActionsItemIntf{
						&pageRuleItem1,
					})
					changeOpt.SetPriority(1)
					changeOpt.SetStatus("active")

					changeResult, changeResp, changeErr := service.ChangePageRule(changeOpt)
					Expect(changeErr).To(BeNil())
					Expect(changeResp).ToNot(BeNil())
					Expect(changeResult).ToNot(BeNil())
					Expect(*changeResult.Success).Should(BeTrue())
				}

				// get page rule
				getOpt := service.NewGetPageRuleOptions(*createResult.Result.ID)
				getResult, getResp, getErr := service.GetPageRule(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				// delete page rule
				delOpt := service.NewDeletePageRuleOptions(*createResult.Result.ID)
				delResult, delResp, delErr := service.DeletePageRule(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*delResult.Success).Should(BeTrue())
			})
			It(`create/update/change/delete/get page rule browser cache ttl`, func() {
				shouldSkipTest()

				// create page rule
				targetConstraintOpt := TargetsItemConstraint{
					Operator: core.StringPtr("matches"),
					Value:    &url,
				}
				targetOpt := []TargetsItem{
					{
						Target:     core.StringPtr("url"),
						Constraint: &targetConstraintOpt,
					}}

				cacheTTL := []int64{
					3600, 7200, 10800, 14400, 18000, 28800, 43200, 57600, 72000, 86400, 172800, 259200, 345600, 432000, 691200, 1382400, 2073600, 2678400, 5356800, 16070400, 31536000,
				}
				pageRuleItem1 := PageRulesBodyActionsItem{
					ID:    core.StringPtr(PageRulesBodyActionsItem_ID_BrowserCacheTTL),
					Value: core.Int64Ptr(1800),
				}
				createOpt := service.NewCreatePageRuleOptions()
				createOpt.SetTargets(targetOpt)
				createOpt.SetActions([]PageRulesBodyActionsItemIntf{
					&pageRuleItem1,
				})
				createOpt.SetPriority(1)
				createOpt.SetStatus("active")

				createResult, createResp, createErr := service.CreatePageRule(createOpt)
				Expect(createErr).To(BeNil())
				Expect(createResp).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).Should(BeTrue())

				for _, ttl := range cacheTTL {

					// update page rule
					pageRuleItem1 = PageRulesBodyActionsItem{
						ID:    core.StringPtr(PageRulesBodyActionsItem_ID_BrowserCacheTTL),
						Value: core.Int64Ptr(ttl),
					}
					updateOpt := service.NewUpdatePageRuleOptions(*createResult.Result.ID)
					updateOpt.SetTargets(targetOpt)
					updateOpt.SetActions([]PageRulesBodyActionsItemIntf{
						&pageRuleItem1,
					})
					updateOpt.SetPriority(1)
					updateOpt.SetStatus("active")

					updateResult, updateResp, updateErr := service.UpdatePageRule(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())
				}

				// change page rule
				for _, ttl := range cacheTTL {
					pageRuleItem1 = PageRulesBodyActionsItem{
						ID:    core.StringPtr(PageRulesBodyActionsItem_ID_BrowserCacheTTL),
						Value: core.Int64Ptr(ttl),
					}
					changeOpt := service.NewChangePageRuleOptions(*createResult.Result.ID)
					changeOpt.SetTargets(targetOpt)
					changeOpt.SetActions([]PageRulesBodyActionsItemIntf{
						&pageRuleItem1,
					})
					changeOpt.SetPriority(1)
					changeOpt.SetStatus("active")

					changeResult, changeResp, changeErr := service.ChangePageRule(changeOpt)
					Expect(changeErr).To(BeNil())
					Expect(changeResp).ToNot(BeNil())
					Expect(changeResult).ToNot(BeNil())
					Expect(*changeResult.Success).Should(BeTrue())
				}

				// get page rule
				getOpt := service.NewGetPageRuleOptions(*createResult.Result.ID)
				getResult, getResp, getErr := service.GetPageRule(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				// delete page rule
				delOpt := service.NewDeletePageRuleOptions(*createResult.Result.ID)
				delResult, delResp, delErr := service.DeletePageRule(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*delResult.Success).Should(BeTrue())
			})
			It(`create/update/change/delete/get page rule edge cache ttl`, func() {
				shouldSkipTest()

				// create page rule
				targetConstraintOpt := TargetsItemConstraint{
					Operator: core.StringPtr("matches"),
					Value:    &url,
				}
				targetOpt := []TargetsItem{
					{
						Target:     core.StringPtr("url"),
						Constraint: &targetConstraintOpt,
					}}

				cacheTTL := []int64{30, 60, 300, 600, 1200, 1800, 3600, 7200, 10800, 14400, 18000, 28800, 43200, 57600, 72000, 86400, 172800, 259200, 345600, 432000, 518400, 604800, 1209600, 2419200}
				pageRuleItem1 := PageRulesBodyActionsItem{
					ID:    core.StringPtr(PageRulesBodyActionsItem_ID_EdgeCacheTTL),
					Value: core.Int64Ptr(30),
				}
				createOpt := service.NewCreatePageRuleOptions()
				createOpt.SetTargets(targetOpt)
				createOpt.SetActions([]PageRulesBodyActionsItemIntf{
					&pageRuleItem1,
				})
				createOpt.SetPriority(1)
				createOpt.SetStatus("active")

				createResult, createResp, createErr := service.CreatePageRule(createOpt)
				Expect(createErr).To(BeNil())
				Expect(createResp).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).Should(BeTrue())

				for _, ttl := range cacheTTL {

					// update page rule
					pageRuleItem1 = PageRulesBodyActionsItem{
						ID:    core.StringPtr(PageRulesBodyActionsItem_ID_EdgeCacheTTL),
						Value: core.Int64Ptr(ttl),
					}
					updateOpt := service.NewUpdatePageRuleOptions(*createResult.Result.ID)
					updateOpt.SetTargets(targetOpt)
					updateOpt.SetActions([]PageRulesBodyActionsItemIntf{
						&pageRuleItem1,
					})
					updateOpt.SetPriority(1)
					updateOpt.SetStatus("active")

					updateResult, updateResp, updateErr := service.UpdatePageRule(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())
				}

				// change page rule
				for _, ttl := range cacheTTL {
					pageRuleItem1 = PageRulesBodyActionsItem{
						ID:    core.StringPtr(PageRulesBodyActionsItem_ID_EdgeCacheTTL),
						Value: core.Int64Ptr(ttl),
					}
					changeOpt := service.NewChangePageRuleOptions(*createResult.Result.ID)
					changeOpt.SetTargets(targetOpt)
					changeOpt.SetActions([]PageRulesBodyActionsItemIntf{
						&pageRuleItem1,
					})
					changeOpt.SetPriority(1)
					changeOpt.SetStatus("active")

					changeResult, changeResp, changeErr := service.ChangePageRule(changeOpt)
					Expect(changeErr).To(BeNil())
					Expect(changeResp).ToNot(BeNil())
					Expect(changeResult).ToNot(BeNil())
					Expect(*changeResult.Success).Should(BeTrue())
				}

				// get page rule
				getOpt := service.NewGetPageRuleOptions(*createResult.Result.ID)
				getResult, getResp, getErr := service.GetPageRule(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				// delete page rule
				delOpt := service.NewDeletePageRuleOptions(*createResult.Result.ID)
				delResult, delResp, delErr := service.DeletePageRule(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*delResult.Success).Should(BeTrue())
			})
			It(`create/update/change/delete/get page rule ssl`, func() {
				shouldSkipTest()

				// create page rule
				targetConstraintOpt := TargetsItemConstraint{
					Operator: core.StringPtr("matches"),
					Value:    &url,
				}
				targetOpt := []TargetsItem{
					{
						Target:     core.StringPtr("url"),
						Constraint: &targetConstraintOpt,
					}}

				ssl := []string{PageRulesBodyActionsItemActionsSsl_Value_Flexible,
					PageRulesBodyActionsItemActionsSsl_Value_Full,
					PageRulesBodyActionsItemActionsSsl_Value_OriginPull,
					PageRulesBodyActionsItemActionsSsl_Value_Strict,
				}
				pageRuleItem1 := PageRulesBodyActionsItem{
					ID:    core.StringPtr(PageRulesBodyActionsItem_ID_Ssl),
					Value: PageRulesBodyActionsItemActionsSsl_Value_Off,
				}
				createOpt := service.NewCreatePageRuleOptions()
				createOpt.SetTargets(targetOpt)
				createOpt.SetActions([]PageRulesBodyActionsItemIntf{
					&pageRuleItem1,
				})
				createOpt.SetPriority(1)
				createOpt.SetStatus("active")

				createResult, createResp, createErr := service.CreatePageRule(createOpt)
				Expect(createErr).To(BeNil())
				Expect(createResp).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).Should(BeTrue())

				for _, item := range ssl {

					// update page rule
					pageRuleItem1 = PageRulesBodyActionsItem{
						ID:    core.StringPtr(PageRulesBodyActionsItem_ID_Ssl),
						Value: item,
					}
					updateOpt := service.NewUpdatePageRuleOptions(*createResult.Result.ID)
					updateOpt.SetTargets(targetOpt)
					updateOpt.SetActions([]PageRulesBodyActionsItemIntf{
						&pageRuleItem1,
					})
					updateOpt.SetPriority(1)
					updateOpt.SetStatus("active")

					updateResult, updateResp, updateErr := service.UpdatePageRule(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())
				}

				// change page rule
				for _, item := range ssl {
					pageRuleItem1 = PageRulesBodyActionsItem{
						ID:    core.StringPtr(PageRulesBodyActionsItem_ID_Ssl),
						Value: item,
					}
					changeOpt := service.NewChangePageRuleOptions(*createResult.Result.ID)
					changeOpt.SetTargets(targetOpt)
					changeOpt.SetActions([]PageRulesBodyActionsItemIntf{
						&pageRuleItem1,
					})
					changeOpt.SetPriority(1)
					changeOpt.SetStatus("active")

					changeResult, changeResp, changeErr := service.ChangePageRule(changeOpt)
					Expect(changeErr).To(BeNil())
					Expect(changeResp).ToNot(BeNil())
					Expect(changeResult).ToNot(BeNil())
					Expect(*changeResult.Success).Should(BeTrue())
				}

				// get page rule
				getOpt := service.NewGetPageRuleOptions(*createResult.Result.ID)
				getResult, getResp, getErr := service.GetPageRule(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				// delete page rule
				delOpt := service.NewDeletePageRuleOptions(*createResult.Result.ID)
				delResult, delResp, delErr := service.DeletePageRule(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*delResult.Success).Should(BeTrue())
			})
			It(`create/update/change/delete/get page rule all security options`, func() {
				shouldSkipTest()

				// create page rule
				targetConstraintOpt := TargetsItemConstraint{
					Operator: core.StringPtr("matches"),
					Value:    &url,
				}
				targetOpt := []TargetsItem{
					{
						Target:     core.StringPtr("url"),
						Constraint: &targetConstraintOpt,
					}}

				securityOptions := []string{
					PageRulesBodyActionsItemActionsSecurityOptions_ID_CacheDeceptionArmor,
					PageRulesBodyActionsItemActionsSecurityOptions_ID_EmailObfuscation,
					PageRulesBodyActionsItemActionsSecurityOptions_ID_ExplicitCacheControl,
					PageRulesBodyActionsItemActionsSecurityOptions_ID_IpGeolocation,
					PageRulesBodyActionsItemActionsSecurityOptions_ID_ServerSideExclude,
					PageRulesBodyActionsItemActionsSecurityOptions_ID_Waf,
				}
				pageRuleItem1 := PageRulesBodyActionsItem{
					ID:    core.StringPtr(PageRulesBodyActionsItemActionsSecurityOptions_ID_BrowserCheck),
					Value: PageRulesBodyActionsItemActionsSsl_Value_Off,
				}
				createOpt := service.NewCreatePageRuleOptions()
				createOpt.SetTargets(targetOpt)
				createOpt.SetActions([]PageRulesBodyActionsItemIntf{
					&pageRuleItem1,
				})
				createOpt.SetPriority(1)
				createOpt.SetStatus("active")

				createResult, createResp, createErr := service.CreatePageRule(createOpt)
				Expect(createErr).To(BeNil())
				Expect(createResp).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).Should(BeTrue())

				for _, item := range securityOptions {

					// update page rule
					pageRuleItem1 = PageRulesBodyActionsItem{
						ID:    core.StringPtr(item),
						Value: PageRulesBodyActionsItemActionsSecurityOptions_Value_On,
					}
					updateOpt := service.NewUpdatePageRuleOptions(*createResult.Result.ID)
					updateOpt.SetTargets(targetOpt)
					updateOpt.SetActions([]PageRulesBodyActionsItemIntf{
						&pageRuleItem1,
					})
					updateOpt.SetPriority(1)
					updateOpt.SetStatus("active")

					updateResult, updateResp, updateErr := service.UpdatePageRule(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())
				}

				// change page rule
				for _, item := range securityOptions {
					pageRuleItem1 = PageRulesBodyActionsItem{
						ID:    core.StringPtr(item),
						Value: PageRulesBodyActionsItemActionsSecurityOptions_Value_Off,
					}
					changeOpt := service.NewChangePageRuleOptions(*createResult.Result.ID)
					changeOpt.SetTargets(targetOpt)
					changeOpt.SetActions([]PageRulesBodyActionsItemIntf{
						&pageRuleItem1,
					})
					changeOpt.SetPriority(1)
					changeOpt.SetStatus("active")

					changeResult, changeResp, changeErr := service.ChangePageRule(changeOpt)
					Expect(changeErr).To(BeNil())
					Expect(changeResp).ToNot(BeNil())
					Expect(changeResult).ToNot(BeNil())
					Expect(*changeResult.Success).Should(BeTrue())
				}

				// get page rule
				getOpt := service.NewGetPageRuleOptions(*createResult.Result.ID)
				getResult, getResp, getErr := service.GetPageRule(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				// delete page rule
				delOpt := service.NewDeletePageRuleOptions(*createResult.Result.ID)
				delResult, delResp, delErr := service.DeletePageRule(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*delResult.Success).Should(BeTrue())
			})
			It(`create/update/change/delete/get page rule bypass cache on cookie`, func() {
				shouldSkipTest()

				// create page rule
				targetConstraintOpt := TargetsItemConstraint{
					Operator: core.StringPtr("matches"),
					Value:    &url,
				}
				targetOpt := []TargetsItem{
					{
						Target:     core.StringPtr("url"),
						Constraint: &targetConstraintOpt,
					}}

				cacheLevelActions := []string{
					PageRulesBodyActionsItemActionsCacheLevel_Value_Aggressive,
					PageRulesBodyActionsItemActionsCacheLevel_Value_Bypass,
					PageRulesBodyActionsItemActionsCacheLevel_Value_CacheEverything,
					PageRulesBodyActionsItemActionsCacheLevel_Value_Simplified,
					PageRulesBodyActionsItemActionsCacheLevel_Value_Basic,
				}
				pageRuleItem1 := PageRulesBodyActionsItem{
					ID:    core.StringPtr(PageRulesBodyActionsItem_ID_BypassCacheOnCookie),
					Value: PageRulesBodyActionsItemActionsCacheLevel_Value_Basic,
				}
				createOpt := service.NewCreatePageRuleOptions()
				createOpt.SetTargets(targetOpt)
				createOpt.SetActions([]PageRulesBodyActionsItemIntf{
					&pageRuleItem1,
				})
				createOpt.SetPriority(1)
				createOpt.SetStatus("active")

				createResult, createResp, createErr := service.CreatePageRule(createOpt)
				Expect(createErr).To(BeNil())
				Expect(createResp).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).Should(BeTrue())

				for _, item := range cacheLevelActions {

					// update page rule
					pageRuleItem1 = PageRulesBodyActionsItem{
						ID:    core.StringPtr(PageRulesBodyActionsItem_ID_BypassCacheOnCookie),
						Value: item,
					}
					updateOpt := service.NewUpdatePageRuleOptions(*createResult.Result.ID)
					updateOpt.SetTargets(targetOpt)
					updateOpt.SetActions([]PageRulesBodyActionsItemIntf{
						&pageRuleItem1,
					})
					updateOpt.SetPriority(1)
					updateOpt.SetStatus("active")

					updateResult, updateResp, updateErr := service.UpdatePageRule(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())
				}

				// change page rule
				for _, item := range cacheLevelActions {
					pageRuleItem1 = PageRulesBodyActionsItem{
						ID:    core.StringPtr(PageRulesBodyActionsItem_ID_BypassCacheOnCookie),
						Value: item,
					}
					changeOpt := service.NewChangePageRuleOptions(*createResult.Result.ID)
					changeOpt.SetTargets(targetOpt)
					changeOpt.SetActions([]PageRulesBodyActionsItemIntf{
						&pageRuleItem1,
					})
					changeOpt.SetPriority(1)
					changeOpt.SetStatus("active")

					changeResult, changeResp, changeErr := service.ChangePageRule(changeOpt)
					Expect(changeErr).To(BeNil())
					Expect(changeResp).ToNot(BeNil())
					Expect(changeResult).ToNot(BeNil())
					Expect(*changeResult.Success).Should(BeTrue())
				}

				// get page rule
				getOpt := service.NewGetPageRuleOptions(*createResult.Result.ID)
				getResult, getResp, getErr := service.GetPageRule(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				// delete page rule
				delOpt := service.NewDeletePageRuleOptions(*createResult.Result.ID)
				delResult, delResp, delErr := service.DeletePageRule(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*delResult.Success).Should(BeTrue())
			})

		})
	})

})
