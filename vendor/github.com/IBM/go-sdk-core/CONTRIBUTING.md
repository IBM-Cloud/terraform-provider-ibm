# Issues

If you encounter an issue with the SDK, you are welcome to submit a [bug report](https://github.com/IBM/go-sdk-core/issues).
Before that, please search for similar issues. It's possible somebody has encountered this issue already.

# Code
## Commit Messages
Commit messages should follow the [Angular Commit Message Guidelines](https://github.com/angular/angular/blob/master/CONTRIBUTING.md#-commit-message-guidelines).
This is because our release tool - [semantic-release](https://github.com/semantic-release/semantic-release) -
uses this format for determining release versions and generating changelogs.
Tools such as [commitizen](https://github.com/commitizen/cz-cli) or [commitlint](https://github.com/conventional-changelog/commitlint)
can be used to help contributors and enforce commit messages.
Here are some examples of acceptable commit messages, along with the release type that would be done based on the commit message:

| Commit message                                                                                                                                                              | Release type               |
|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------|
| `fix(IAM Authentication) propagate token request errors back to request invocation thread`                                                                                  | Patch Release              |
| `feat(JSON Serialization): add custom deserializer for dynamic models`                                                                                                      | ~~Minor~~ Feature Release  |
| `feat(BaseService): added baseURL as param to BaseService ctor`<br><br>`BREAKING CHANGE: The global-search service has been updated to reflect version 3 of the API.`       | ~~Major~~ Breaking Release |

# Pull Requests

If you want to contribute to the repository, here's a quick guide:
  1. Fork the repository  
  
  2. The `go-sdk-core` project uses Go modules for dependency management, so do NOT set the `GOPATH` environment 
  variable to include your local `go-sdk-core` project directory.  
  
  3. Clone the respository into a local directory.  
  
  4. Install the `golangci-lint` tool:
     ```sh
     curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(go env GOPATH)/bin v1.21.0
     ```  
  * Note: As of this writing, the 1.21.0 version of `golangci-lint` is being used by this project.
  Please check the `curl` command found in the `.travis.yml` file to see the version of this tool that is currently 
  being used at the time you are planning to commit changes. This will ensure that you are using the same version 
  of the linter as the Travis build automation, which will ensure that you are using the same set of linter checks
  that the automated build uses.  
  
  5. Make your code changes as needed.  Be sure to add new tests for any new or modified functionality.  
  
  6. Test your changes:
     ```sh
     go test ./...
     ```  
  
  7. Check your code for lint issues
     ```sh
     golangci-lint run
     ```  
  
  8. Commit your changes:
  * Commit messages should follow the Angular commit message guidelines as mentioned above.
  
  9. Push your commit(s) to your fork and submit a pull request to the **master** branch.

# Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
   have the right to submit it under the open source license
   indicated in the file; or

(b) The contribution is based upon previous work that, to the best
   of my knowledge, is covered under an appropriate open source
   license and I have the right under that license to submit that
   work with modifications, whether created in whole or in part
   by me, under the same open source license (unless I am
   permitted to submit under a different license), as indicated
   in the file; or

(c) The contribution was provided directly to me by some other
   person who certified (a), (b) or (c) and I have not modified
   it.

(d) I understand and agree that this project and the contribution
   are public and that a record of the contribution (including all
   personal information I submit with it, including my sign-off) is
   maintained indefinitely and may be redistributed consistent with
   this project or the open source license(s) involved.

# Additional Resources
+ [General GitHub documentation](https://help.github.com/)
+ [GitHub pull request documentation](https://help.github.com/send-pull-requests/)

[dw]: https://developer.ibm.com/answers/questions/ask.html
[stackoverflow]: http://stackoverflow.com/questions/ask?tags=ibm
[dep]: https://github.com/golang/dep
