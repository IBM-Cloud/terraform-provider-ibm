module github.com/IBM-Cloud/terraform-provider-ibm

go 1.16

require (
	github.com/IBM-Cloud/bluemix-go v0.0.0-20220407050707-b4cd0d4da813
	github.com/IBM-Cloud/container-services-go-sdk v0.0.0-20210705152127-41ca00fc9a62
	github.com/IBM-Cloud/power-go-client v1.1.4
	github.com/IBM/apigateway-go-sdk v0.0.0-20210714141226-a5d5d49caaca
	github.com/IBM/appconfiguration-go-admin-sdk v0.1.0
	github.com/IBM/appid-management-go-sdk v0.0.0-20210908164609-dd0e0eaf732f
	github.com/IBM/cloud-databases-go-sdk v0.2.0
	github.com/IBM/cloudant-go-sdk v0.0.36
	github.com/IBM/container-registry-go-sdk v0.0.15
	github.com/IBM/event-notifications-go-admin-sdk v0.0.7
	github.com/IBM/eventstreams-go-sdk v1.2.0
	github.com/IBM/go-sdk-core/v5 v5.9.2
	github.com/IBM/ibm-cos-sdk-go v1.8.0
	github.com/IBM/ibm-cos-sdk-go-config v1.2.0
	github.com/IBM/ibm-hpcs-tke-sdk v0.0.0-20211109141421-a4b61b05f7d1
	github.com/IBM/keyprotect-go-client v0.7.0
	github.com/IBM/networking-go-sdk v0.28.0
	github.com/IBM/platform-services-go-sdk v0.22.6
	github.com/IBM/push-notifications-go-sdk v0.0.0-20210310100607-5790b96c47f5
	github.com/IBM/scc-go-sdk v1.3.4
	github.com/IBM/scc-go-sdk/v3 v3.1.6
	github.com/IBM/schematics-go-sdk v0.1.3
	github.com/IBM/secrets-manager-go-sdk v0.1.19
	github.com/IBM/vpc-go-sdk v0.19.0
	github.com/PromonLogicalis/asn1 v0.0.0-20190312173541-d60463189a56 // indirect
	github.com/ScaleFT/sshkeys v0.0.0-20200327173127-6142f742bca5
	github.com/Shopify/sarama v1.29.1
	github.com/apache/openwhisk-client-go v0.0.0-20200201143223-a804fb82d105
	github.com/apparentlymart/go-cidr v1.1.0
	github.com/aws/aws-sdk-go v1.37.0 // indirect
	github.com/cloudfoundry/jibber_jabber v0.0.0-20151120183258-bcc4c8345a21 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/strfmt v0.21.2
	github.com/go-test/deep v1.0.4 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/go-cmp v0.5.6
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/uuid v1.3.0
	github.com/hashicorp/go-uuid v1.0.2
	github.com/hashicorp/go-version v1.4.0
	github.com/hashicorp/hcl/v2 v2.8.2 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.10.1
	github.com/hokaccha/go-prettyjson v0.0.0-20170213120834-e6b9231a2b1c // indirect
	github.com/jinzhu/copier v0.3.2
	github.com/minsikl/netscaler-nitro-go v0.0.0-20170827154432-5b14ce3643e3
	github.com/mitchellh/go-homedir v1.1.0
	github.com/softlayer/softlayer-go v1.0.3
	github.ibm.com/org-ids/tekton-pipeline-go-sdk v0.0.1
	github.ibm.com/org-ids/toolchain-go-sdk v0.0.2
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/sys v0.0.0-20220330033206-e17cdc41300f // indirect
	golang.org/x/tools v0.1.10 // indirect
	google.golang.org/api v0.34.0 // indirect
	gotest.tools v2.2.0+incompatible
)

replace github.com/softlayer/softlayer-go v1.0.3 => github.com/IBM-Cloud/softlayer-go v1.0.5-tf

replace github.com/dgrijalva/jwt-go v3.2.0+incompatible => github.com/golang-jwt/jwt v3.2.1+incompatible

replace github.ibm.com/org-ids/toolchain-go-sdk => ../toolchain-go-sdk

replace github.ibm.com/org-ids/tekton-pipeline-go-sdk => ../tekton-pipeline-go-sdk
