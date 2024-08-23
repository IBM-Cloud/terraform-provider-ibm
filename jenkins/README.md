# Generate a new Terraform provider with Jenkins


## Overview

This documentation will guide you through how to generate new `terraform-provider-ibm` for `linux` OS by a jenkins job and also upload it to artifactory.

## How To


- Jenkins job location: https://wcp-lox-team-jenkins.swg-devops.com/view/Product-Lifecycle/job/product-lifecycle-generate-terraform-provider/

- Artifactory provider folder: https://eu.artifactory.swg-devops.com/artifactory/wcp-lox-team-private-provider-terraform-local/

- Terraform provider repo: https://github.ibm.com/LOX/terraform-provider-ibm


Steps:

- Login to Jenkins ( see url above )
- Click on the `product-lifecycle-generate-terraform-provider` job under `Product-Lifecycle` view.
- On the left panel click on `Build with Parameters` 
- Fill out the provider version (check the latest version under `Releases` and add a new version here)
- Click on `Build`
- Verify that provider zip is uploaded to artifactory
- Create a release under `Releases` . Use the version you provided as a parameter above. Also included in release notes the changes and the current publi `terraform-provider-ibm` version.


