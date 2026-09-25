// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventstreams

import (
	"errors"
	"fmt"
	"log"
	"maps"
	"slices"
	"strings"

	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/version"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/IBM/sarama"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// clientPool maintains Kafka admin client for each instance.
// key is instance's CRN
var clientPool = map[string]sarama.ClusterAdmin{}

type extensions struct {
	platformGeneration int
	adminURL           string
	bootstrapServers   []string
}

// formatObject flattens the map representation of a JSON object into a string.
// Property names are preserved, but values are replaced by their Go type.
// Intent is to provide debug information for service extensions without logging
// values.
func formatObject(m map[string]interface{}) string {
	sb := &strings.Builder{}
	for idx, propertyName := range slices.Sorted(maps.Keys(m)) {
		if idx != 0 {
			sb.WriteString(",")
		}
		value := m[propertyName]
		if mapValue, ok := value.(map[string]interface{}); ok {
			sb.WriteString(fmt.Sprintf("%q:{%s}", propertyName, formatObject(mapValue)))
		} else {
			sb.WriteString(fmt.Sprintf("%q:%T", propertyName, value))
		}
	}
	return sb.String()
}

func parseInstanceExtensions(instance *resourcecontrollerv2.ResourceInstance, meta interface{}) (*extensions, error) {
	if kafkaHTTP, ok := instance.Extensions["kafka_http_url"].(string); ok {
		if brokersSASL, ok := instance.Extensions["kafka_brokers_sasl"].([]interface{}); ok {
			bootstrapServers := flex.ExpandStringList(brokersSASL)
			slices.Sort(bootstrapServers)
			return &extensions{
				adminURL:           kafkaHTTP,
				bootstrapServers:   bootstrapServers,
				platformGeneration: 1,
			}, nil
		}
	}
	if dataservices, ok := instance.Extensions["dataservices"].(map[string]any); ok {
		if connection, ok := dataservices["connection"].(map[string]any); ok {
			if bootstrapServers, ok := connection["bootstrap_servers"].(string); ok {
				if restURL, ok := connection["rest_url"].(string); ok {
					return &extensions{
						adminURL:           restURL,
						bootstrapServers:   strings.Split(bootstrapServers, ","),
						platformGeneration: 2,
					}, nil
				}
			}
		}
	}
	log.Printf("[DEBUG] unexpected format of instance extensions: %s", formatObject(instance.Extensions))
	return nil, errors.New("unexpected format of instance extensions")
}

func createSaramaAdminClient(d *schema.ResourceData, meta interface{}) (sarama.ClusterAdmin, *extensions, string, error) {
	bxSession, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		log.Printf("[DEBUG] createSaramaAdminClient BluemixSession err %s", err)
		return nil, nil, "", err
	}
	instanceCRN := d.Get("resource_instance_id").(string)
	if len(instanceCRN) == 0 {
		id := d.Id()
		if len(id) == 0 || !strings.Contains(id, ":") {
			log.Printf("[DEBUG] createSaramaAdminClient resource_instance_id is missing")
			return nil, nil, "", fmt.Errorf("resource_instance_id is required")
		}
		instanceCRN = getInstanceCRN(id)
	}
	instance, err := getInstanceDetails(instanceCRN, meta)
	if err != nil {
		log.Printf("[DEBUG] createSaramaAdminClient err %s", err)
		return nil, nil, "", err
	}
	ext, err := parseInstanceExtensions(instance, meta)
	if err != nil {
		return nil, nil, "", err
	}
	log.Printf("[INFO] createSaramaAdminClient kafka_http_url is set to %s", ext.adminURL)
	log.Printf("[INFO] createSaramaAdminClient kafka_brokers_sasl is set to %s", strings.Join(ext.bootstrapServers, ","))
	var adminClient sarama.ClusterAdmin
	var ok bool
	if adminClient, ok = clientPool[instanceCRN]; ok {
		log.Printf("[DEBUG] createSaramaAdminClient got client from pool for instance %s", instanceCRN)
		return adminClient, ext, instanceCRN, nil
	}
	config := sarama.NewConfig()
	config.ClientID = fmt.Sprintf("terraform-provider-ibm/%s", version.Version)
	config.Net.SASL.Enable = true
	config.Net.TLS.Enable = true
	config.Version = sarama.V3_8_1_0
	if ext.platformGeneration >= 2 {
		config.Version = sarama.V4_1_0_0
	}
	tenantID := strings.TrimPrefix(strings.Split(ext.adminURL, ".")[0], "https://")
	if ext.platformGeneration == 1 && tenantID != "" && tenantID != "admin" {
		config.Net.SASL.AuthIdentity = tenantID
	} else {
		config.Net.SASL.AuthIdentity = instanceCRN
	}
	config.Admin.Timeout = adminClientTimeout
	config.Net.SASL.Mechanism = sarama.SASLTypeOAuth
	config.Net.SASL.TokenProvider, err = newAccessTokenProvider(bxSession)
	if err != nil {
		return nil, nil, "", err
	}
	adminClient, err = sarama.NewClusterAdmin(ext.bootstrapServers, config)
	if err != nil {
		log.Printf("[DEBUG] createSaramaAdminClient NewClusterAdmin err %s", err)
		return nil, nil, "", err
	}
	clientPool[instanceCRN] = adminClient
	log.Printf("[INFO] createSaramaAdminClient instance %s 's client is initialized", instanceCRN)
	return adminClient, ext, instanceCRN, nil
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
