Please follow below steps to release a new Go SDK on public repository.

1. Create a branch name `release-YYYY-MM-DD` with the date as API version date, from the below repository:  
`https://github.com/IBM/vpc-go-sdk`
  
2. Update the generated code from the API spec `https://pages.github.ibm.com/riaas/api-spec/spec_genesis_production_redirect`. Verify the generator version and `core-sdk` compatibility and update the `go.mod` file accordingly. Compatibility Chart - `https://github.ibm.com/CloudEngineering/openapi-sdkgen/wiki/Compatibility-Chart`

3. Must verify the integration test and examples has been added for the new feature going in the api version.

4. Update the `README.md` file with next release version. Commit the code (Semantic versionoing is diabled as of now) 

5. push the branch and cretae a PR `https://github.com/IBM/vpc-go-sdk` to master branch

6. Post the new PR link to slack channel `#vpc-client-sdk-terraform-internal`

7. Once the PR will approved and merge, go to `https://github.com/IBM/vpc-go-sdk/releases` and create a manual release by adding a new tag and update the release notes. 