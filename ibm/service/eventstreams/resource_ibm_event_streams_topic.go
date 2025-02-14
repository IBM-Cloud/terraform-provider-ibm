// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventstreams

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"
	"time"

	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/version"
	"github.com/IBM/go-sdk-core/v5/core"
	iamidentity "github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/IBM/sarama"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	defaultReplicationFactor = 3
	defaultCleanupPolicy     = "delete"
	defaultRetentionBytes    = 1073741824 // 100 MB
	defaultRetentionMs       = 86400000   // 24 hours
	defaultSegmentBytes      = 536870912  // 512 MB
)

var (
	adminClientTimeout  = 30 * time.Second
	allowedTopicConfigs = []string{
		"cleanup.policy",
		"retention.ms",
		"retention.bytes",
		"segment.ms",
		"segment.bytes",
		"segment.index.bytes",
		"message.audit.enable", // enterprise only
	}
)

func ResourceIBMEventStreamsTopic() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMEventStreamsTopicCreate,
		ReadContext:   resourceIBMEventStreamsTopicRead,
		UpdateContext: resourceIBMEventStreamsTopicUpdate,
		DeleteContext: resourceIBMEventStreamsTopicDelete,
		Importer:      &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			"resource_instance_id": {
				Type:        schema.TypeString,
				Description: "The CRN of the Event Streams instance",
				Required:    true,
			},
			"kafka_http_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "API endpoint for interacting with Event Streams REST API",
			},
			"kafka_brokers_sasl": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Kafka brokers addresses for interacting with Kafka native API",
			},
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the topic",
				Required:    true,
			},
			"partitions": {
				Type:        schema.TypeInt,
				Description: "The number of partitions",
				Optional:    true,
				Default:     1,
			},
			"config": {
				Type:        schema.TypeMap,
				Description: "The configuration parameters of a topic",
				Optional:    true,
			},
		},
	}
}

// clientPool maintains Kafka admin client for each instance.
// key is instance's CRN
var clientPool = map[string]sarama.ClusterAdmin{}

func resourceIBMEventStreamsTopicExists(context context.Context, d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Printf("[DEBUG] resourceIBMEventStreamsTopicExists")
	adminClient, _, err := createSaramaAdminClient(d, meta)
	if err != nil {
		log.Printf("[DEBUG] resourceIBMEventStreamsTopicExists createSaramaAdminClient err %s", err)
		return false, err
	}
	topicName := d.Get("name").(string)
	topicsMetadata, err := adminClient.DescribeTopics([]string{topicName})
	if err != nil {
		descErr := fmt.Errorf("[ERROR] Error describing topic %s : %v", topicName, err)
		log.Printf("[DEBUG] resourceIBMEventStreamsTopicExists DescribeTopics: %s, err %v", topicName, descErr)
		return false, descErr
	}
	if len(topicsMetadata) == 0 {
		descErr := fmt.Errorf("no metadata was returned for topic %s", topicName)
		log.Printf("[DEBUG] resourceIBMEventStreamsTopicExists DescribeTopics: %s, err %v", topicName, descErr)
		return false, descErr
	}
	if topicsMetadata[0].Err != sarama.ErrNoError {
		metadataErr := topicsMetadata[0].Err
		log.Printf("[DEBUG] resourceIBMEventStreamsTopicExists DescribeTopics: %s, err %v", topicName, metadataErr)
		return false, metadataErr
	}
	log.Printf("[INFO] resourceIBMEventStreamsTopicExists topic %s exists", topicName)
	return true, nil
}

func resourceIBMEventStreamsTopicCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] resourceIBMEventStreamsTopicCreate")
	adminClient, instanceCRN, err := createSaramaAdminClient(d, meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("resourceIBMEventStreamsTopicCreate createSaramaAdminClient: %s", err), "ibm_event_streams_topic", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	topicName := d.Get("name").(string)
	partitions := d.Get("partitions").(int)
	config := d.Get("config").(map[string]interface{})
	topicDetail := sarama.TopicDetail{
		NumPartitions:     int32(partitions),
		ReplicationFactor: int16(defaultReplicationFactor),
		ConfigEntries:     config2TopicDetail(config),
	}
	err = adminClient.CreateTopic(topicName, &topicDetail, false)
	if err != nil {
		if kafkaErr, ok := err.(*sarama.TopicError); ok {
			if kafkaErr != nil && kafkaErr.Err == sarama.ErrTopicAlreadyExists {
				exists, err := resourceIBMEventStreamsTopicExists(context, d, meta)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("resourceIBMEventStreamsTopicCreate resourceIBMEventStreamsTopicExists %s: %s", topicName, err), "ibm_event_streams_topic", "create")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				if exists {
					d.SetId(getTopicID(instanceCRN, topicName))
					return resourceIBMEventStreamsTopicRead(context, d, meta)
				}
			}
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("resourceIBMEventStreamsTopicCreate CreateTopic %s: %s", topicName, err), "ibm_event_streams_topic", "create")
		log.Printf("[ERROR]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	log.Printf("[INFO] resourceIBMEventStreamsTopicCreate CreateTopic: topic is %s, detail is %v", topicName, topicDetail)
	d.SetId(getTopicID(instanceCRN, topicName))
	return resourceIBMEventStreamsTopicRead(context, d, meta)
}

func resourceIBMEventStreamsTopicRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] resourceIBMEventStreamsTopicRead")
	adminClient, instanceCRN, err := createSaramaAdminClient(d, meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("resourceIBMEventStreamsTopicRead createSaramaAdminClient: %s", err), "ibm_event_streams_topic", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	topicID := d.Id()
	topicName := getTopicName(topicID)
	topics, err := adminClient.ListTopics()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("resourceIBMEventStreamsTopicRead ListTopics: %s", err), "ibm_event_streams_topic", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	for name, detail := range topics {
		if name == topicName {
			d.Set("resource_instance_id", instanceCRN)
			d.Set("name", name)
			d.Set("partitions", detail.NumPartitions)
			if config := d.Get("config"); config != nil {
				savedConfig := map[string]*string{}
				for k := range config.(map[string]interface{}) {
					if value, ok := detail.ConfigEntries[k]; ok {
						savedConfig[k] = value
					}
				}
				d.Set("config", topicDetail2Config(savedConfig))
			}
			return nil
		}
	}
	log.Printf("[INFO] resourceIBMEventStreamsTopicRead topic %s does not exist", topicName)
	d.SetId("")
	return nil
}

func resourceIBMEventStreamsTopicUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] resourceIBMEventStreamsTopicUpdate")
	adminClient, _, err := createSaramaAdminClient(d, meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("resourceIBMEventStreamsTopicUpdate createSaramaAdminClient: %s", err), "ibm_event_streams_topic", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	topicName := d.Get("name").(string)
	if d.HasChange("partitions") {
		oi, ni := d.GetChange("partitions")
		oldPartitions := oi.(int)
		newPartitions := ni.(int)
		log.Printf("[INFO]resourceIBMEventStreamsTopicUpdate Updating partitions from %d to %d", oldPartitions, newPartitions)
		err = adminClient.CreatePartitions(topicName, int32(newPartitions), nil, false)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("resourceIBMEventStreamsTopicUpdate CreatePartitions: %s", err), "ibm_event_streams_topic", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		d.Set("partitions", int32(newPartitions))
		log.Printf("[INFO]resourceIBMEventStreamsTopicUpdate partitions is set to %d", newPartitions)
	}
	if d.HasChange("config") {
		config := d.Get("config").(map[string]interface{})
		configEntries := config2TopicDetail(config)
		err = adminClient.AlterConfig(sarama.TopicResource, topicName, configEntries, false)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("resourceIBMEventStreamsTopicUpdate AlterConfig: %s", err), "ibm_event_streams_topic", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		d.Set("config", topicDetail2Config(configEntries))
		log.Printf("[INFO]resourceIBMEventStreamsTopicUpdate config is set to %v", topicDetail2Config(configEntries))
	}
	return resourceIBMEventStreamsTopicRead(context, d, meta)
}

func resourceIBMEventStreamsTopicDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] resourceIBMEventStreamsTopicDelete")
	adminClient, _, err := createSaramaAdminClient(d, meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("resourceIBMEventStreamsTopicDelete createSaramaAdminClient: %s", err), "ibm_event_streams_topic", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	topicName := d.Get("name").(string)
	err = adminClient.DeleteTopic(topicName)
	if err != nil {
		if kerr, ok := err.(sarama.KError); ok {
			if kerr == sarama.ErrUnknownTopicOrPartition {
				d.SetId("")
				log.Printf("[INFO]resourceIBMEventStreamsTopicDelete topic %s does not exist", topicName)
				return nil
			}
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("resourceIBMEventStreamsTopicDelete DeleteClient: %s", err), "ibm_event_streams_topic", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	log.Printf("[INFO]resourceIBMEventStreamsTopicDelete topic %s deleted", topicName)
	return nil
}

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
	var adminClient sarama.ClusterAdmin
	var ok bool
	if adminClient, ok = clientPool[instanceCRN]; ok {
		log.Printf("[DEBUG] createSaramaAdminClient got client from pool for instance %s", instanceCRN)
		return adminClient, instanceCRN, nil
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
	if bxSession.Config.BluemixAPIKey != "" {
		config.Net.SASL.User = "token"
		config.Net.SASL.Password = bxSession.Config.BluemixAPIKey
		config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
		log.Printf("[DEBUG] createSaramaAdminClient configured SASL mechanism=PLAIN")
	} else if _, err = validateToken(bxSession.Config.IAMAccessToken); err == nil {
		config.Net.SASL.Mechanism = sarama.SASLTypeOAuth
		config.Net.SASL.TokenProvider = accessTokenProvider{clientSession: bxSession}
		log.Printf("[DEBUG] createSaramaAdminClient configured SASL mechanism=OAUTHBEARER")
	} else {
		return nil, "", errors.New("either IBMCLOUD_API_KEY or IAM_TOKEN needs to be configured")
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
	clientSession *session.Session
}

// Token() implements sarama.AccessTokenProvider interface for sasl.mechanism=OAUTHBEARER
func (tp accessTokenProvider) Token() (*sarama.AccessToken, error) {
	token, err := validateToken(tp.clientSession.Config.IAMAccessToken)
	if err != nil {
		log.Printf("[DEBUG] accessTokenProvider.Token() error:%s", err)
		return nil, err
	}
	if expired(token) {
		authenticator := &core.IamAuthenticator{
			RefreshToken: tp.clientSession.Config.IAMRefreshToken,
			ClientId:     "bx",
			ClientSecret: "bx",
			URL:          conns.EnvFallBack([]string{"IBMCLOUD_IAM_API_ENDPOINT"}, iamidentity.DefaultServiceURL),
		}
		token, err = authenticator.GetToken()
		if err != nil {
			log.Printf("[DEBUG] accessTokenProvider.authenticator.GetToken() error:%s", err)
			return nil, err
		}
	}
	return &sarama.AccessToken{Token: token}, nil
}

func validateToken(token string) (string, error) {
	if len(token) == 0 {
		return "", errors.New("IAMAccessToken is required")
	}
	token = strings.TrimPrefix(token, "Bearer")
	token = strings.Trim(token, " ")
	if len(strings.Split(token, ".")) != 3 {
		return "", errors.New("IAMAccessToken is malformed")
	}
	return token, nil
}

func expired(token string) (expired bool) {
	expired = true
	tokenString, err := base64.RawURLEncoding.DecodeString(strings.Split(token, ".")[1])
	if err != nil {
		log.Printf("[DEBUG] expired.DecodeString() error:%s", err)
		return
	}
	claims := jwt.RegisteredClaims{}
	err = json.Unmarshal(tokenString, &claims)
	if err != nil {
		log.Printf("[DEBUG] expired.Unmarshal() error:%s", err)
		return
	}
	if claims.ID == "" {
		log.Printf("[DEBUG] expired.jit is empty")
		return
	}
	if claims.ExpiresAt.IsZero() {
		log.Printf("[DEBUG] expired.exp is zero")
		return
	}
	return claims.ExpiresAt.Before(time.Now().Add(10 * time.Second))
}
