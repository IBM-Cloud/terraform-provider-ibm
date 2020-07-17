module github.com/IBM-Cloud/terraform-provider-ibm

go 1.12

require (
	github.com/Bowery/prompt v0.0.0-20190916142128-fa8279994f75 // indirect
	github.com/IBM-Cloud/bluemix-go v0.0.0-20200714095434-4dd1df3322eb
	github.com/IBM-Cloud/power-go-client v1.0.37
	github.com/IBM/apigateway-go-sdk v0.0.0-20200414212859-416e5948678a
	github.com/IBM/dns-svcs-go-sdk v0.0.3
	github.com/IBM/go-sdk-core/v3 v3.3.1
	github.com/IBM/ibm-cos-sdk-go v1.3.1
	github.com/IBM/ibm-cos-sdk-go-config v1.0.0
	github.com/IBM/keyprotect-go-client v0.3.5-0.20200325142150-b63163832e26
	github.com/IBM/vpc-go-sdk v0.0.1
	github.com/ScaleFT/sshkeys v0.0.0-20200327173127-6142f742bca5
	github.com/apache/incubator-openwhisk-client-go v0.0.0-20171128215515-ad814bc98c32
	github.com/apparentlymart/go-cidr v1.0.1
	github.com/cloudfoundry/jibber_jabber v0.0.0-20151120183258-bcc4c8345a21 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/runtime v0.19.15 // indirect
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-openapi/validate v0.19.8 // indirect
	github.com/go-test/deep v1.0.4 // indirect
	github.com/hashicorp/go-uuid v1.0.1
	github.com/hashicorp/go-version v1.2.0
	github.com/hashicorp/hil v0.0.0-20200423225030-a18a1cd20038 // indirect
	github.com/hashicorp/terraform v0.12.28 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.6.0
	github.com/hokaccha/go-prettyjson v0.0.0-20170213120834-e6b9231a2b1c // indirect
	github.com/minsikl/netscaler-nitro-go v0.0.0-20170827154432-5b14ce3643e3
	github.com/mitchellh/go-homedir v1.1.0
	github.com/renier/xmlrpc v0.0.0-20170708154548-ce4a1a486c03 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/softlayer/softlayer-go v0.0.0-20190814165317-b9062a914a22
	github.ibm.com/ibmcloud/namespace-go-sdk v0.0.0-00010101000000-000000000000 // indirect
	github.ibm.com/ibmcloud/networking-go-sdk v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
)

replace github.com/softlayer/softlayer-go v0.0.0-20190814165317-b9062a914a22 => ./common/github.com/softlayer/softlayer-go

replace github.ibm.com/ibmcloud/networking-go-sdk => ./common/github.ibm.com/ibmcloud/networking-go-sdk

replace github.ibm.com/ibmcloud/namespace-go-sdk => ./common/github.ibm.com/ibmcloud/namespace-go-sdk
