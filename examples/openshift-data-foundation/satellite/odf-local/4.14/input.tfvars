## DEFAULT VALUES ARE SET ##
## Please change according to your configuratiom ##


# Common for both storage configuration and assignment 
ibmcloud_api_key = ""
location = "" #Location of your storage configuration and assignment
configName = "" #Name of your storage configuration
region = ""


#ODF Storage Configuration

storageTemplateName = "odf-local"
storageTemplateVersion = "4.14"

## User Parameters

autoDiscoverDevices = "true"
osdDevicePaths = ""
billingType = "advanced"
clusterEncryption = "false"
kmsBaseUrl = null
kmsEncryption = "false"
kmsInstanceId = null
kmsInstanceName = null
kmsTokenUrl = null
ibmCosEndpoint = null
ibmCosLocation = null
ignoreNoobaa = false
numOfOsd = "1"
ocsUpgrade = "false"
workerNodes = null
encryptionInTransit = false
disableNoobaaLB = false
performCleanup = false
taintNodes = false
addSingleReplicaPool = false
prepareForDisasterRecovery = false

## Secret Parameters
ibmCosAccessKey = null
ibmCosSecretKey = null
iamAPIKey = "" #Required
kmsApiKey = null
kmsRootKey = null

#ODF Storage Assignment
assignmentName = ""
cluster = ""
updateConfigRevision = false

## NOTE ##
# The following variables will cause issues to your storage assignment lifecycle, so please use only with a storage configuration resource.
deleteAssignments = false
updateAssignments = false