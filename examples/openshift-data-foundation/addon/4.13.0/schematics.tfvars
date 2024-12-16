## DEFAULT VALUES ARE SET ##
## Please change according to your configuratiom ##


# To enable ODF AddOn on your cluster
ibmcloud_api_key = ""
cluster = ""
region = ""
odfVersion = ""


# To create the Ocscluster Custom Resource Definition, with the following specs
autoDiscoverDevices = "false"
billingType = "advanced"
clusterEncryption = "false"
hpcsBaseUrl = null
hpcsEncryption = "false"
hpcsInstanceId = null
hpcsSecretName = null
hpcsServiceName = null
hpcsTokenUrl = null
ignoreNoobaa = false
numOfOsd = "1"
ocsUpgrade = "false"
osdDevicePaths = null
osdSize = "512Gi"
osdStorageClassName = "ibmc-vpc-block-metro-10iops-tier"
workerNodes = null
encryptionInTransit = false
disableNoobaaLB = false
taintNodes = false