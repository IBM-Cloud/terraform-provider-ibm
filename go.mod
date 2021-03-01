module github.com/IBM-Cloud/terraform-provider-ibm

go 1.16

require (
	github.com/Bowery/prompt v0.0.0-20190916142128-fa8279994f75 // indirect
	github.com/IBM-Cloud/bluemix-go v0.0.0-20210903091337-e7648c203ed5
	github.com/IBM-Cloud/container-services-go-sdk v0.0.0-20210705152127-41ca00fc9a62
	github.com/IBM-Cloud/power-go-client v1.0.55
	github.com/IBM/apigateway-go-sdk v0.0.0-20210714141226-a5d5d49caaca
	github.com/IBM/appconfiguration-go-admin-sdk v0.1.0
	github.com/IBM/appid-management-go-sdk v0.0.0-20210727091553-7e0e5823e707
	github.com/IBM/cloudant-go-sdk v0.0.36
	github.com/IBM/container-registry-go-sdk v0.0.13
	github.com/IBM/go-sdk-core/v4 v4.10.0
	github.com/IBM/go-sdk-core/v5 v5.6.3
	github.com/IBM/ibm-cos-sdk-go v1.7.0
	github.com/IBM/ibm-cos-sdk-go-config v1.2.0
	github.com/IBM/ibm-hpcs-tke-sdk v0.0.0-20210723145459-a232c3f3ac91
	github.com/IBM/keyprotect-go-client v0.7.0
	github.com/IBM/networking-go-sdk v0.19.0
	github.com/IBM/platform-services-go-sdk v0.19.3
	github.com/IBM/push-notifications-go-sdk v0.0.0-20210310100607-5790b96c47f5
	github.com/IBM/schematics-go-sdk v0.0.2
	github.com/IBM/secrets-manager-go-sdk v0.1.19
	github.com/IBM/vpc-go-sdk v0.9.0
	github.com/PromonLogicalis/asn1 v0.0.0-20190312173541-d60463189a56 // indirect
	github.com/ScaleFT/sshkeys v0.0.0-20200327173127-6142f742bca5
	github.com/Shopify/sarama v1.29.1
	github.com/apache/openwhisk-client-go v0.0.0-20200201143223-a804fb82d105
	github.com/apparentlymart/go-cidr v1.1.0
	github.com/cloudfoundry/jibber_jabber v0.0.0-20151120183258-bcc4c8345a21 // indirect
	github.com/dchest/safefile v0.0.0-20151022103144-855e8d98f185 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/strfmt v0.20.2
	github.com/go-openapi/validate v0.20.1 // indirect
	github.com/go-test/deep v1.0.4 // indirect
	github.com/golang-jwt/jwt v3.2.1+incompatible
	github.com/google/go-cmp v0.5.6
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/google/uuid v1.2.0
	github.com/hashicorp/go-uuid v1.0.2
	github.com/hashicorp/go-version v1.3.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/hokaccha/go-prettyjson v0.0.0-20170213120834-e6b9231a2b1c // indirect
	github.com/jinzhu/copier v0.3.2
	github.com/kardianos/govendor v1.0.9 // indirect
	github.com/minsikl/netscaler-nitro-go v0.0.0-20170827154432-5b14ce3643e3
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/gox v1.0.1 // indirect
	github.com/softlayer/softlayer-go v1.0.3
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	golang.org/x/sys v0.0.0-20210903071746-97244b99971b // indirect
	golang.org/x/tools v0.1.5 // indirect
	gotest.tools v2.2.0+incompatible
)

replace github.com/softlayer/softlayer-go v1.0.3 => github.com/IBM-Cloud/softlayer-go v1.0.3-tf
