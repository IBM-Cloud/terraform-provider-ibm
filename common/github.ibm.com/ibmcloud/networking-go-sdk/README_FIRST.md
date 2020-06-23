# IBM Cloud Go SDK Template Usage Instructions

This repository serves as a template for Go SDKs that are produced with the
[IBM OpenAPI SDK Generator](https://github.ibm.com/CloudEngineering/openapi-sdkgen).

You can use the contents of this repository to create your own Go SDK repository.

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      markdown-toc -i --maxdepth 4 README_FIRST.md
  -->

<!-- toc -->

- [How to use this repository](#how-to-use-this-repository)
  * [1. Create your new github repository from this template](#1-create-your-new-github-repository-from-this-template)
  * [2. Sanity-check your new repository](#2-sanity-check-your-new-repository)
  * [3. Modify selected files](#3-modify-selected-files)
  * [4. Add one or more services to the project](#4-add-one-or-more-services-to-the-project)
  * [5. Build and test the project](#5-build-and-test-the-project)
- [Integration tests](#integration-tests)
- [Continuous Integration](#continuous-integration)
  * [Release management with semantic-release](#release-management-with-semantic-release)
  * [Encrypting secrets](#encrypting-secrets)
- [Setting the ``User-Agent`` Header In Preparation for SDK Metrics Gathering](#setting-the-user-agent-header-in-preparation-for-sdk-metrics-gathering)

<!-- tocstop -->

## How to use this repository

### 1. Create your new github repository from this template
This SDK template repository is implemented as a
[github template](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template),
which makes it easy to create new projects from it.

To create a new SDK repository from this template, follow these instructions:  
1. In your browser, open the link for this
[template repository](https://github.ibm.com/CloudEngineering/go-sdk-template).

2. Click on the `Use this template` button that appears next to the `Clone or download` button.

3. In the next window:  
- Select the `Owner`. This is the github id or organization where the new repository should be created
- Enter the respository name (e.g. `platform-services-go-sdk`):  
  - Recommendation: use a name of the form `<service-category>-<language>-sdk`, where:  
    - `<service-category>` refers to the IBM Cloud service category associated with the services that
	  will be included in the project (e.g. `platform-services`)
    - `<language>` is the language associated with the SDK project (e.g. `go`)
	
4. Click the `Create repository from template` button to create the new repository  

If your goal is to create the new SDK repository on the `Github Enterprise` server (github.ibm.com),
then you are finished creating the new repository and you can proceed to section 2.

On the other hand, if your goal is to create the new SDK repository on the `Public Github` server (github.com),
then perform these additional steps:

5. Create a new **EMPTY** repository on the Public Github server:  
- Select "No template" for the "Repository template" option
- Select the `Owner` (your personal id or an organization)
- Enter the same respository name that you used when creating the new repository above (e.g. my-go-sdk)
- Do NOT select the `Initialize this repository with a README` option
- Select `None` for the `Add .gitignore` and `Add a license` options
- Click the `Create repository` button.
- After the new empty repository has been created, you will be at the main page
of your new repository, which will include this text:
```
...or push an existing repository from the command line

git remote add origin git@github.com:padamstx/my-go-sdk.git
git push -u origin master
```
- Take note of the two git commands listed above for your new repository, as we'll execute these later

6. Clone your new `Github Enterprise` repository (created in steps 1-3 above)
to your local development environment:  

```sh
[/work/demos]
$ git clone git@github.ibm.com:phil-adams/my-go-sdk.git
Cloning into 'my-go-sdk'...
remote: Enumerating objects: 36, done.
remote: Counting objects: 100% (36/36), done.
remote: Compressing objects: 100% (32/32), done.
remote: Total 36 (delta 1), reused 0 (delta 0), pack-reused 0
Receiving objects: 100% (36/36), 28.74 KiB | 577.00 KiB/s, done.
Resolving deltas: 100% (1/1), done.
```

7. "cd" into your project's root directory:

```sh
[/work/demos]
$ cd my-go-sdk
[/work/demos/my-go-sdk]
$ 
```

8. Remove the existing remote:  
```sh
[/work/demos/my-go-sdk]
$ git remote remove origin
```

9. Add a new remote which reflects your new `Public Github` repository:

```sh
[/work/demos/my-go-sdk]
$ git remote add origin git@github.com:padamstx/my-go-sdk.git
```

10. Push your local repository to the new remote (Public Github):  

```sh
[/work/demos/my-go-sdk]
$ git push -u origin master
Enumerating objects: 36, done.
Counting objects: 100% (36/36), done.
Delta compression using up to 12 threads
Compressing objects: 100% (31/31), done.
Writing objects: 100% (36/36), 28.74 KiB | 28.74 MiB/s, done.
Total 36 (delta 1), reused 36 (delta 1)
remote: Resolving deltas: 100% (1/1), done.
To github.com:padamstx/my-go-sdk.git
 * [new branch]      master -> master
Branch 'master' set up to track remote branch 'master' from 'origin'.
```

You have now created your new SDK repository on the `Public Github` server.

You may want to now delete the new SDK repository that you created on the `Github Enterprise`
server since it will no longer be used now that you have created your repository on `Public Github`.


### 2. Sanity-check your new repository

After creating your new SDK repository from the template repository, and cloning it
into your local development environment, you can do a quick sanity check by
running this command in the project root directory:
```
go test ./...
```
You should see output like this:
```
$ go test ./...
go: finding github.com/IBM/go-sdk-core/v3 v3.2.4
go: finding github.com/go-playground/locales v0.12.1
go: finding github.com/stretchr/testify v1.4.0
.
.
.
ok  	github.ibm.com/CloudEngineering/go-sdk-template/common	0.002s
ok  	github.ibm.com/CloudEngineering/go-sdk-template/exampleservicev1	0.006s
```

Note: the first time you build and test the project, you'll see output showing
that the Go engine is downloading the dependencies needed by the project since
they're not yet cached in your environment.

Note: this project uses go "modules" for dependency management.
For this reason, make sure the `GOPATH` environment variable is not set in
your shell when executing the `go` commands above.


### 3. Modify selected files

- In this section, you'll modify various files within your new SDK repository to reflect
the proper names and settings for your specific project.

- The template repository comes with an example service included, but this should be removed
from your project.  Remove the following directory and its contents:
  - exampleservicev1

- Next, here is a list of the various files within the project with comments
that will guide you in the required modifications:

  - `common/headers.go`:
    - modify the `sdkName` constant to reflect your project name (e.g. `platform-services-go-sdk`)
    - read the comments in the `GetSdkHeaders()` function and follow as appropriate

  - `common/version.go`:
    - make sure the `Version` constant is set to "0.0.1", as this will be the starting version
      number (release) of the project.

  - `go.mod`/`go.sum`:
    - Remove the `go.mod` and `go.sum` files
    - Run this command to create a new `go.mod` file which will contain your project's
      github url as the module import path:
      ```sh
         go mod init <module-import-path>
      ```
      where `<module-import-path>` should be the correct module import path for your project.
      This will be the github repository URL without the `https` scheme
      (e.g. `github.ibm.com/ibmcloud/platform-services-go-sdk`).

  - `.travis.yml`:
    - Remove the `jobs:` section, as this is applicable only to the template repository's build.

  - `README.md`:
    - Change the title to reflect your project; leave the version in the title as `0.0.1`
    - Change the `cloud.ibm.com/apidocs` link to reflect the correct service category
      (e.g. `platform-services`)
    - In the Overview section, modify `IBM Cloud MySDK Go SDK` to reflect your project
      (e.g. `IBM Cloud Platform Services Go SDK`)
    - In the table of services, remove the entry for the example service; later you'll list each
      service contained in your SDK project in this table, along with a link to the online reference docs
      and the name of the generated service struct.
    - In the Installation section, update the examples to reflect your new
      project's module import path (e.g. `github.ibm.com/ibmcloud/platform-services-go-sdk`).
    - In the "Issues" section, modify `<github-repo-url>` to reflect the Github URL for your project.
    - Note that the README.md file contains a link to a common README document where general
      SDK usage information can be found.
    - When finished, read through the document and make any other changes that might be necessary.

  - `CONTRIBUTING.md`:
    - In the "Issues" section, modify `<github-repo-url>` to reflect the Github URL for your project.

At this point, it's probably a good idea to commit the changes that you have made so far.
Be sure to use proper commit messages when committing changes (follow the link in `CONTRIBUTING.md`
to the common CONTRIBUTING document).  
Example:
```sh
cd <project-root>
git add .
git commit -m "chore: initial SDK project setup"
```


### 4. Add one or more services to the project
For each service that you'd like to add to your SDK project, follow
[these instructions](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/CONTRIBUTING_go.md#adding-a-new-service).

### 5. Build and test the project
If you made it this far, congratulate yourself!

After preparing your new Go SDK project and then generating the Go
code for your service(s), it's time to build and test your project.

To build and test all of the code within your project, you can run these commands in the project
root directory:
```
go test ./...
```
If everything builds and tests cleanly, you should see output like this:
```
$ go test ./...
ok  	github.ibm.com/CloudEngineering/go-sdk-template/common	0.002s
ok  	github.ibm.com/CloudEngineering/go-sdk-template/exampleservicev1	0.006s
```
Note: The above output reflects the module import path for the `go-sdk-template` repository and the
example service that is shipped with it.  Your output should reflect your Go SDK project's
module prefix and your project's set of packages.

If you encounter compile issues with the service or unit test code generated by the SDK Generator,
please let us know by posting on the `#ibm-sdk-generation` slack channel or by opening an issue
in the [`github.ibm.com/arf/arf-planning-sdk`](https://github.ibm.com/arf/planning-sdk-squad/)
issue repository.

Our goal is to generate the SDK service and unit test code that can be built and tested without
manual intervention.  If we fall short of that goal, we'd love to hear about it.


## Integration tests
Integration tests must be developed by hand.
For integration tests to run properly with an actual running instance of the service,
credentials (e.g. IAM api key, etc.) must be provided as external configuration properties.
Details about this can be found
[here](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/README.md#using-external-configuration).

An example integration test is located at `exampleservicev1/example_service_v1_integration_test.go`.
In order to run the "example service" integration test,
you'll need an actual running instance of the example service.
To run this service, clone the [Example Service repo](https://github.ibm.com/CloudEngineering/example-service)
and follow the instructions there for how to start up an instance of the example service.


## Continuous Integration
This repository is set up to use [Travis](https://travis-ci.com/)
or [Travis Enterprise](https://travis.ibm.com) for continuous integration.

The `.travis.yml` file contains all the instructions necessary to run the build.

For details related to the `travis.yml` file, see
[this](https://docs.travis-ci.com/user/customizing-the-build/)

### Release management with semantic-release
The `.travis.yml` file included in this template repository is configured to
perform automated release management with
[semantic-release](https://semantic-release.gitbook.io/semantic-release/).

When you configure your SDK project in Travis, be sure to set this environment variable in your
Travis build settings:  
- `GH_TOKEN`: set this to the Github oauth token for a user having "push" access to your repository

If you are using `Travis Enterprise` (travis.ibm.com), you'll need to add these environment variables
as well:  
- `GH_URL`: set this to the string `https://github.ibm.com`
- `GH_PREFIX`: set this to the string `/api/v3`

### Encrypting secrets
To run integration tests within a Travis build, you'll need to encrypt the file containing the
required external configuration properties.
For details on how to do this, please see
[this](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/EncryptingSecrets.md)


## Setting the ``User-Agent`` Header In Preparation for SDK Metrics Gathering

If you plan to gather metrics for your SDK, the `User-Agent` header value must be
a string similar to the following:
   `my-go-sdk/0.0.1 (lang=go; arch=x86_64; os=Linux; go.version=1.12.9)`

The key parts are the sdk name (`my-go-sdk`), version (`0.0.1`) and the
language name (`lang=go`).
This is required because the analytics data collector uses the User-Agent header included
with each request to gather usage data for IBM Cloud services.

The default implementation of the `common.GetSDKHeaders()` method provided in this SDK template
repository will need to be modified slightly for your SDK.
Replace the `my-go-sdk/0.0.1` part with the name and version of your
Go SDK. The rest of the system information should remain as-is.

For example, suppose your Go SDK project is called `platform-services-go-sdk` and its
version is `2.3.1`.
The `User-Agent` header value should be:
   `platform-services-go-sdk/2.3.1 (lang=go; arch=x86_64; os=Linux; go.version=1.12.9)`

__Note__: It is very important that the sdk name ends with the string `-sdk`,
as the analytics data collector uses this to gather usage data.

More information about the analytics tool, and other steps you should take to start gathering
metrics for your SDK can be found [here](https://github.ibm.com/CloudEngineering/sdk-analytics).
