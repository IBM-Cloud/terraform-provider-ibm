# Terraform IBM Provider Activity Tracker
<!-- markdownlint-disable MD026 -->
This area is primarily for IBM provider contributors and maintainers. For information on _using_ Terraform and the IBM provider, see the links below.


## Environment variables to set for the tests
* COS_API_KEY   : API Key used for creating COS targets
* INGESTION_KEY : Deprecated. LogDNA Ingestion Key used for creating Logdna targets. This variable is no longer in use and will be removed in the next major version of the provider.
* IES_API_KEY   : Event streams password used for creating Event streams targets

## Handy Links
* [Find out about contributing](../../../CONTRIBUTING.md) to the IBM provider!
* IBM Provider Docs: [Home](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs)
* IBM Provider Docs: [One of the Activity Tracker resources](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/atracker_route)
* IBM API Docs: [IBM API Docs for Activity Tracker](https://cloud.ibm.com/apidocs/atracker)
* IBM Activity Tracker SDK: [IBM SDK for Activity Tracker](https://github.com/IBM/platform-services-go-sdk/tree/main/atrackerv2)
