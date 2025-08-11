---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_githubconsolidated"
description: |-
  Manages cd_toolchain_tool_githubconsolidated.
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain_tool_githubconsolidated

Create, update, and delete cd_toolchain_tool_githubconsolidateds with this resource.

See the [tool integration](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-github) page for more information.

## Example Usage

```hcl
resource "ibm_cd_toolchain_tool_githubconsolidated" "cd_toolchain_tool_githubconsolidated_instance" {
  initialization {
		git_id = "github"
		owner_id = "<github-user-id>"
		repo_name = "myrepo"
		source_repo_url = "https://github.com/source-repo-owner/source-repo"
		type = "clone"
		private_repo = true
  }
  parameters {
		integration_owner = "my-userid"
		auth_type = "pat"
		api_token = "<api_token>"
		toolchain_issues_enabled = true
  }
  toolchain_id = ibm_cd_toolchain.cd_toolchain.id
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `initialization` - (Required, List) 
Nested schema for **initialization**:
	* `auto_init` - (Optional, Forces new resource, Boolean) Setting this value to true will initialize this repository with a README.  This parameter is only used when creating a new repository.
	  * Constraints: The default value is `false`.
	* `blind_connection` - (Optional, Forces new resource, Boolean) Setting this value to true means the server is not addressable on the public internet. IBM Cloud will not be able to validate the connection details you provide. Certain functionality that requires API access to the git server will be disabled. Delivery pipeline will only work using a private worker that has network access to the git server.
	  * Constraints: The default value is `false`.
	* `git_id` - (Optional, Forces new resource, String) Set this value to 'github' for github.com, or 'githubcustom' for a custom GitHub Enterprise server.
	* `owner_id` - (Optional, Forces new resource, String) The GitHub user or organization that owns the repository.  This parameter is required when creating a new repository, cloning, or forking a repository.  The value will be computed when linking to an existing repository.
	* `private_repo` - (Optional, Forces new resource, Boolean) Set this value to 'true' to make the repository private when creating a new repository or when cloning or forking a repository.  This parameter is not used when linking to an existing repository.
	  * Constraints: The default value is `false`.
	* `repo_name` - (Optional, Forces new resource, String) The name of the new GitHub repository to create.  This parameter is required when creating a new repository, cloning, or forking a repository.  The value will be computed when linking to an existing repository.
	* `repo_url` - (Optional, Forces new resource, String) The URL of the GitHub repository for this tool integration.  This parameter is required when linking to an existing repository.  The value will be computed when creating a new repository, cloning, or forking a repository.
	* `root_url` - (Optional, Forces new resource, String) The Root URL of the server. e.g. https://github.example.com.
	* `source_repo_url` - (Optional, Forces new resource, String) The URL of the repository that you are forking or cloning.  This parameter is required when forking or cloning a repository.  It is not used when creating a new repository or linking to an existing repository.
	* `title` - (Optional, Forces new resource, String) The title of the server. e.g. My GitHub Enterprise Server.
	* `type` - (Required, Forces new resource, String) The operation that should be performed to initialize the new tool integration. Use 'new' or 'new_if_not_exists' to create a new git repository, 'clone' or 'clone_if_not_exists' to clone an existing repository into a new git repository, 'fork' or 'fork_if_not_exists' to fork an existing git repository, or 'link' to link to an existing git repository. If you attempt to apply a resource with type 'new', 'clone', or 'fork' when the target repo already exists, the attempt will fail. If you apply a resource with type 'new_if_not_exists`, 'clone_if_not_exists', or 'fork_if_not_exists' when the target repo already exists, the existing repo will be used as-is.
	  * Constraints: Allowable values are: `new`, `fork`, `clone`, `link`, `new_if_not_exists`, `clone_if_not_exists`, `fork_if_not_exists`.
* `name` - (Optional, String) Name of the tool.
  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$/`.
* `parameters` - (Required, List) Unique key-value pairs representing parameters to be used to create the tool. A list of parameters for each tool integration can be found in the <a href="https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-integrations">Configuring tool integrations page</a>.
Nested schema for **parameters**:
	* `api_root_url` - (Computed, String) The API root URL for the GitHub server.
	* `api_token` - (Optional, String) Personal Access Token. Required if ‘auth_type’ is set to ‘pat’, ignored otherwise.
	* `auth_type` - (Optional, String) Select the method of authentication that will be used to access the git provider. The default value is 'oauth'.
	  * Constraints: Allowable values are: `oauth`, `pat`.
	* `auto_init` - (Computed, Boolean) Setting this value to true will initialize this repository with a README.  This parameter is only used when creating a new repository.
	  * Constraints: The default value is `false`.
	* `blind_connection` - (Computed, Boolean) Setting this value to true means the server is not addressable on the public internet. IBM Cloud will not be able to validate the connection details you provide. Certain functionality that requires API access to the git server will be disabled. Delivery pipeline will only work using a private worker that has network access to the git server.
	  * Constraints: The default value is `false`.
	* `default_branch` - (Computed, String) The default branch of the git repository.
	* `enable_traceability` - (Optional, Boolean) Set this value to 'true' to track the deployment of code changes by creating tags, labels and comments on commits, pull requests and referenced issues.
	  * Constraints: The default value is `false`.
	* `git_id` - (Computed, String) Set this value to 'github' for github.com, or 'githubcustom' for a custom GitHub Enterprise server.
	* `integration_owner` - (Optional, String) Select the user which git operations will be performed as.
	* `owner_id` - (Computed, String) The GitHub user or organization that owns the repository.  This parameter is required when creating a new repository, cloning, or forking a repository.  The value will be computed when linking to an existing repository.
	* `private_repo` - (Computed, Boolean) Set this value to 'true' to make the repository private when creating a new repository or when cloning or forking a repository.  This parameter is not used when linking to an existing repository.
	  * Constraints: The default value is `false`.
	* `repo_id` - (Computed, String) The ID of the GitHub repository.
	* `repo_name` - (Computed, String) The name of the new GitHub repository to create.  This parameter is required when creating a new repository, cloning, or forking a repository.  The value will be computed when linking to an existing repository.
	* `repo_url` - (Computed, String) The URL of the GitHub repository for this tool integration.  This parameter is required when linking to an existing repository.  The value will be computed when creating a new repository, cloning, or forking a repository.
	* `root_url` - (Computed, String) The Root URL of the server. e.g. https://github.example.com.
	* `source_repo_url` - (Computed, String) The URL of the repository that you are forking or cloning.  This parameter is required when forking or cloning a repository.  It is not used when creating a new repository or linking to an existing repository.
	* `title` - (Computed, String) The title of the server. e.g. My GitHub Enterprise Server.
	* `token_url` - (Computed, String) The token URL used for authorizing with the GitHub server.
	* `toolchain_issues_enabled` - (Optional, Boolean) Setting this value to true will enable issues on the GitHub repository and add an issues tool card to the toolchain.  Setting the value to false will remove the tool card from the toolchain, but will not impact whether or not issues are enabled on the GitHub repository itself.
	  * Constraints: The default value is `true`.
	* `type` - (Computed, String) The operation that should be performed to initialize the new tool integration. Use 'new' or 'new_if_not_exists' to create a new git repository, 'clone' or 'clone_if_not_exists' to clone an existing repository into a new git repository, 'fork' or 'fork_if_not_exists' to fork an existing git repository, or 'link' to link to an existing git repository. If you attempt to apply a resource with type 'new', 'clone', or 'fork' when the target repo already exists, the attempt will fail. If you apply a resource with type 'new_if_not_exists`, 'clone_if_not_exists', or 'fork_if_not_exists' when the target repo already exists, the existing repo will be used as-is.
	  * Constraints: Allowable values are: `new`, `fork`, `clone`, `link`, `new_if_not_exists`, `clone_if_not_exists`, `fork_if_not_exists`.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind the tool to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the cd_toolchain_tool_githubconsolidated.
* `crn` - (String) Tool CRN.
* `href` - (String) URI representing the tool.
* `referent` - (List) Information on URIs to access this resource through the UI or API.
Nested schema for **referent**:
	* `api_href` - (String) URI representing this resource through an API.
	* `ui_href` - (String) URI representing this resource through the UI.
* `resource_group_id` - (String) Resource group where the tool is located.
* `state` - (String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.
* `tool_id` - (String) Tool ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.
* `toolchain_crn` - (String) CRN of toolchain which the tool is bound to.
* `updated_at` - (String) Latest tool update timestamp.


## Import

You can import the `ibm_cd_toolchain_tool_githubconsolidated` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `tool_id` in the following format:

<pre>
&lt;toolchain_id&gt;/&lt;tool_id&gt;
</pre>
* `toolchain_id`: A string. ID of the toolchain to bind the tool to.
* `tool_id`: A string. Tool ID.

# Syntax
<pre>
$ terraform import ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated &lt;toolchain_id&gt;/&lt;tool_id&gt;
</pre>
