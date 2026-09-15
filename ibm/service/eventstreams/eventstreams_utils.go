// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventstreams

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/version"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/sarama"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// clientPool maintains Kafka admin client for each instance.
// key is instance's CRN
var clientPool = map[string]sarama.ClusterAdmin{}

func createSaramaAdminClient(d *schema.ResourceData, meta interface{}) (sarama.ClusterAdmin, string, error) {
	bxSession, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		log.Printf("[DEBUG] createSaramaAdminClient BluemixSession err %s", err)
		return nil, "", err
	}
	instanceCRN := d.Get("resource_instance_id").(string)
	if len(instanceCRN) == 0 {
		topicID := d.Id()
		if len(topicID) == 0 || !strings.Contains(topicID, ":") {
			log.Printf("[DEBUG] createSaramaAdminClient resource_instance_id is missing")
			return nil, "", fmt.Errorf("resource_instance_id is required")
		}
		instanceCRN = getInstanceCRN(topicID)
	}
	instance, err := getInstanceDetails(instanceCRN, meta)
	if err != nil {
		return nil, "", err
	}
	adminURL := instance.Extensions["kafka_http_url"].(string)
	d.Set("kafka_http_url", adminURL)
	log.Printf("[INFO] createSaramaAdminClient kafka_http_url is set to %s", adminURL)
	brokerAddress := flex.ExpandStringList(instance.Extensions["kafka_brokers_sasl"].([]interface{}))
	slices.Sort(brokerAddress)
	d.Set("kafka_brokers_sasl", brokerAddress)
	log.Printf("[INFO] createSaramaAdminClient kafka_brokers_sasl is set to %s", brokerAddress)
	var adminClient sarama.ClusterAdmin
	var ok bool
	if adminClient, ok = clientPool[instanceCRN]; ok {
		log.Printf("[DEBUG] createSaramaAdminClient got client from pool for instance %s", instanceCRN)
		return adminClient, instanceCRN, nil
	}
	config := sarama.NewConfig()
	config.ClientID = fmt.Sprintf("terraform-provider-ibm/%s", version.Version)
	config.Net.SASL.Enable = true
	config.Net.TLS.Enable = true
	config.Version = sarama.MaxVersion
	tenantID := strings.TrimPrefix(strings.Split(adminURL, ".")[0], "https://")
	if tenantID != "" && tenantID != "admin" {
		config.Net.SASL.AuthIdentity = tenantID
	} else {
		config.Net.SASL.AuthIdentity = instanceCRN
	}
	config.Admin.Timeout = adminClientTimeout
	config.Net.SASL.Mechanism = sarama.SASLTypeOAuth
	config.Net.SASL.TokenProvider, err = newAccessTokenProvider(bxSession)
	if err != nil {
		return nil, "", err
	}
	adminClient, err = sarama.NewClusterAdmin(brokerAddress, config)
	if err != nil {
		log.Printf("[DEBUG] createSaramaAdminClient NewClusterAdmin err %s", err)
		return nil, "", err
	}
	clientPool[instanceCRN] = adminClient
	log.Printf("[INFO] createSaramaAdminClient instance %s 's client is initialized", instanceCRN)
	return adminClient, instanceCRN, nil
}

func topicDetail2Config(topicConfigEntries map[string]*string) map[string]*string {
	configs := map[string]*string{}
	for key, value := range topicConfigEntries {
		if flex.IndexOf(key, allowedTopicConfigs) != -1 {
			configs[key] = value
		}
	}
	return configs
}

func config2TopicDetail(config map[string]interface{}) map[string]*string {
	configEntries := make(map[string]*string)
	for key, value := range config {
		switch value := value.(type) {
		case string:
			configEntries[key] = &value
		}
	}
	return configEntries
}

func getTopicID(instanceCRN string, topicName string) string {
	crnSegments := strings.Split(instanceCRN, ":")
	crnSegments[8] = "topic"
	crnSegments[9] = topicName
	return strings.Join(crnSegments, ":")
}

func getTopicName(topicID string) string {
	return strings.Split(topicID, ":")[9]
}

func getInstanceCRN(topicID string) string {
	crnSegments := strings.Split(topicID, ":")
	crnSegments[8] = ""
	crnSegments[9] = ""
	return strings.Join(crnSegments, ":")
}

type accessTokenProvider struct {
	authenticator *core.IamAuthenticator
}

func newAccessTokenProvider(sess *session.Session) (*accessTokenProvider, error) {
	iamEndpoint, err := sess.Config.EndpointLocator.IAMEndpoint()
	if err != nil {
		log.Printf("[DEBUG] newAccessTokenProvider.IAMEndpoint() error:%s", err)
		return nil, err
	}
	authenticator, err := core.NewIamAuthenticatorBuilder().
		SetURL(iamEndpoint).
		SetApiKey(sess.Config.BluemixAPIKey).
		SetRefreshToken(sess.Config.IAMRefreshToken).
		SetClientIDSecret("bx", "bx").
		Build()
	if err != nil {
		log.Printf("[DEBUG] newAccessTokenProvider.NewIamAuthenticatorBuilder() error:%s", err)
		return nil, err
	}
	return &accessTokenProvider{authenticator}, nil
}

// Token() implements sarama.AccessTokenProvider interface for sasl.mechanism=OAUTHBEARER
func (tp *accessTokenProvider) Token() (*sarama.AccessToken, error) {
	token, err := tp.authenticator.GetToken()
	if err != nil {
		log.Printf("[DEBUG] accessTokenProvider.GetToken() error:%s", err)
		return nil, err
	}
	return &sarama.AccessToken{Token: token}, nil
}
