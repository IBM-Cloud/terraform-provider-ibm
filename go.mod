module github.com/IBM-Cloud/terraform-provider-ibm

go 1.12

require (
	github.com/Bowery/prompt v0.0.0-20190916142128-fa8279994f75 // indirect
	github.com/IBM-Cloud/bluemix-go v0.0.0-20200131055858-778531c8c923
	github.com/IBM-Cloud/power-go-client v0.0.0-00010101000000-000000000000
	github.com/IBM/ibm-cos-sdk-go v1.2.0
	github.com/IBM/keyprotect-go-client v0.3.2
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/apache/incubator-openwhisk-client-go v0.0.0-20171128215515-ad814bc98c32
	github.com/apparentlymart/go-cidr v1.0.1
	github.com/asaskevich/govalidator v0.0.0-20180720115003-f9ffefc3facf // indirect
	github.com/cloudfoundry/jibber_jabber v0.0.0-20151120183258-bcc4c8345a21 // indirect
	github.com/dchest/safefile v0.0.0-20151022103144-855e8d98f185 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-openapi/analysis v0.0.0-20180825180245-b006789cd277 // indirect
	github.com/go-openapi/errors v0.0.0-20180827163446-87bb65328877 // indirect
	github.com/go-openapi/jsonpointer v0.0.0-20180825180259-52eb3d4b47c6 // indirect
	github.com/go-openapi/jsonreference v0.0.0-20180825180305-1c6a3fa339f2 // indirect
	github.com/go-openapi/loads v0.0.0-20180825180312-fd899182a268 // indirect
	github.com/go-openapi/runtime v0.0.0-20180825180317-95364c1e5610 // indirect
	github.com/go-openapi/spec v0.0.0-20180825180323-f1468acb3b29 // indirect
	github.com/go-openapi/strfmt v0.0.0-20180825180330-6292e8b5de3b
	github.com/go-openapi/swag v0.0.0-20180825180336-0dc164b7900e // indirect
	github.com/go-openapi/validate v0.0.0-20180825180342-e0648ff40507 // indirect
	github.com/go-test/deep v1.0.4 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/google/shlex v0.0.0-20181106134648-c34317bd91bf // indirect
	github.com/hashicorp/go-uuid v1.0.1
	github.com/hashicorp/go-version v1.2.0
	github.com/hashicorp/terraform-plugin-sdk v1.0.0
	github.com/hashicorp/tf-sdk-migrator v1.0.0 // indirect
	github.com/hokaccha/go-prettyjson v0.0.0-20170213120834-e6b9231a2b1c // indirect
	github.com/kardianos/govendor v1.0.9 // indirect
	github.com/mailru/easyjson v0.0.0-20180823135443-60711f1a8329 // indirect
	github.com/minsikl/netscaler-nitro-go v0.0.0-20170827154432-5b14ce3643e3
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/gox v1.0.1 // indirect
	github.com/nicksnyder/go-i18n v0.0.0-20171206142411-aa0ce51472e0 // indirect
	github.com/onsi/ginkgo v1.10.3 // indirect
	github.com/onsi/gomega v1.7.1 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pelletier/go-toml v0.0.0-20171024211038-4e9e0ee19b60 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/renier/xmlrpc v0.0.0-20170708154548-ce4a1a486c03 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/softlayer/softlayer-go v0.0.0-20190814165317-b9062a914a22
	github.com/stretchr/testify v1.4.0 // indirect
	github.ibm.com/Bluemix/riaas-go-client v0.0.0-20191018070922-afd27ac04d4f
	golang.org/x/tools v0.0.0-20200128002243-345141a36859 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce // indirect
)

replace github.ibm.com/Bluemix/riaas-go-client v0.0.0-20191018070922-afd27ac04d4f => ./common/github.ibm.com/Bluemix/riaas-go-client

replace github.com/softlayer/softlayer-go v0.0.0-20190814165317-b9062a914a22 => ./common/github.com/softlayer/softlayer-go

replace github.com/IBM-Cloud/power-go-client => ./common/github.com/IBM-Cloud/power-go-client
