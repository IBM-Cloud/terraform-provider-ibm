Please follow below steps to release a new Go SDK on enterprise and public repository.

**Note:**  Check the generator version, compatible go-sdk-core version from `https://github.ibm.com/CloudEngineering/openapi-sdkgen/wiki/Compatibility-Chart`. If there is any change in the generated sdk then proceed.

1. Create a branch name `release-YYYY-MM-DD` with the date as API release  date, on the below repository:  
`https://github.ibm.com/ibmcloud/vpc-go-sdk`
  
2. Update the generated code from the API spec `https://pages.github.ibm.com/riaas/api-spec/spec_genesis_production_redirect`. Verify the generator version and `core-sdk` compatibility and update the `go.mod` file accordingly. Compatibility Chart - `https://github.ibm.com/CloudEngineering/openapi-sdkgen/wiki/Compatibility-Chart` and make the handedits :

        2.1.1 replace `CidRs` -> `CIDRs` and `Cidrs` -> `CIDRs` using case sensitive replacements
        2.1.2 remove required comment from `VpcV1Options` struct.
        2.1.3 add the below code to set the default version of the api in `NewVpcV1` function
           
                	if options.Version == nil {
		                	options.Version = core.StringPtr("2022-03-29")
	                	}
             

3. Must verify the integration test and examples has been added for the new feature going in the api spec release date. Run the test and examples on enterprise repo. Then make a commit with an appropriate commit message and with this description `Signed-off-by: Your Name <your-email@ibm.com>`

4. Update the `README.md` file with next release version. Commit the code (Semantic versioning is disabled as of now). Replace `github.ibm.com/ibmcloud` with `github.com/IBM` when moving from enterprise to public. Verify the test and examples on public repo. Then make a commit with an appropriate commit message and with this description `Signed-off-by: Your Name <your-email@ibm.com>`

5. Push the branch and create a PR `https://github.com/IBM/vpc-go-sdk` to master branch

6. Post the new PR link to slack channel `#vpc-client-sdk-terraform-internal`

7. Once the PR will be approved and merge, go to `https://github.com/IBM/vpc-go-sdk/releases` and create a manual release by adding a new tag and update the release notes in the following order
    * Breaking changes
    * New features
    * Changes
    * Bug fixes