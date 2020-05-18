# Issues

If you encounter an issue with the project, you are welcome to submit a [bug report](<github-repo-url>/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

# Pull Requests

If you want to contribute to the repository, here's a quick guide:
  1. Fork the repository

  2. The `vpc-go-sdk` project uses Go modules for dependency management, so do NOT set the `GOPATH` environment
  variable to include your local `vpc-go-sdk` project directory.

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
  * Commits should follow the [Angular commit message guidelines](https://github.com/angular/angular/blob/master/CONTRIBUTING.md#-commit-message-guidelines). This is because our release tool uses this format for determining release versions and generating changelogs. To make this easier, we recommend using the [Commitizen CLI](https://github.com/commitizen/cz-cli) with the `cz-conventional-changelog` adapter.

  9. Push to your fork and submit a pull request to the **master** branch

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

## Additional Resources
+ [General GitHub documentation](https://help.github.com/)
+ [GitHub pull request documentation](https://help.github.com/send-pull-requests/)

[Maven]: https://maven.apache.org/guides/getting-started/maven-in-five-minutes.html
