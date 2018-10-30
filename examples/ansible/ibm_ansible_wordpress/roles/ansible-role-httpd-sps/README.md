# Ansible role `httpd`

A simple Ansible role for installing and configuring the Apache web server for RHEL/CentOS 7 and Fedora 21+. Specifically, the responsibilities of this role are to:

- Install the necessary packages;
- Maintain the main configuration file;
- Maintain the configuration file for `mod_ssl`.
- Enable and maintain the configuration file for `mod_status`.
- Install support for scripting language (currently only PHP)

HTTPS/TLS is enabled, by default using the standard self-signed certificate. You can provide your own certificate by setting the appropriate role variables.

Currently, no virtual hosts or other features are provided.

## Requirements

- The firewall settings are not managed by this role.
- If you want to use custom certificates, you have to make sure that they are installed on the system before applying this role.

## Role Variables

If no variables are set, applying this role will result in a configuration equivalent to the default install. Consequently, no variables are required.

| Variable                        | Default                                    | Comments (type)                                                                       |
| :---                            | :---                                       | :---                                                                                  |
| `httpd_AccessLog_ssl`           | logs/ssl_access_log                        |                                                                                       |
| `httpd_DocumentRoot`            | '/var/www/html'                            |                                                                                       |
| `httpd_ErrorLog_ssl`            | logs/ssl_error_log                         |                                                                                       |
| `httpd_ErrorLog`                | logs/error_log                             |                                                                                       |
| `httpd_Listen_ssl`              | 443                                        |                                                                                       |
| `httpd_Listen`                  | 80                                         |                                                                                       |
| `httpd_LogLevel_ssl`            | warn                                       |                                                                                       |
| `httpd_LogLevel`                | warn                                       |                                                                                       |
| `httpd_SSLCACertificateFile`    | -                                          |                                                                                       |
| `httpd_SSLCertificateChainFile` | -                                          |                                                                                       |
| `httpd_SSLCertificateFile`      | /etc/pki/tls/certs/localhost.crt           |                                                                                       |
| `httpd_SSLCertificateKeyFile`   | /etc/pki/tls/private/localhost.key         |                                                                                       |
| `httpd_SSLCipherSuite`          | See [default variables](defaults/main.yml) |                                                                                       |
| `httpd_SSLHonorCipherOrder`     | 'on'                                       |                                                                                       |
| `httpd_SSLProtocol`             | 'all -SSLv3 -TLSv1'                        |                                                                                       |
| `httpd_ServerAdmin`             | root@localhost                             |                                                                                       |
| `httpd_ServerRoot`              | '/etc/httpd'                               |                                                                                       |
| `httpd_ServerTokens`            | Prod                                       | See [documentation](https://httpd.apache.org/docs/current/mod/core.html#servertokens) |
| `httpd_scripting`               | 'none'                                     | Allowed values: `php`                                                                 |
| `httpd_StatusEnable`            | false                                      | Enable mod_status                                                                     |
| `httpd_StatusLocation`          | '/server-status'                           | Location for mod_status status page                                                   |
| `httpd_StatusRequire`           | 'host localhost'                           | Access control for mod_status                                                         |
| `httpd_ExtendedStatus`          | on                                         | Enable ExtendedStatus                                                                 |

## Dependencies

No dependencies.

## Example Playbook

See the test playbooks in either the [Vagrant](https://github.com/bertvv/ansible-role-httpd/blob/vagrant-tests/test.yml) or [Docker](https://github.com/bertvv/ansible-role-httpd/blob/docker-tests/test.yml) test environment. See the section Testing for details.

## Testing

There are two types of test environments available. One powered by Vagrant, another by Docker. The latter is suitable for running automated tests on Travis-CI. Test code is kept in separate orphan branches. For details of how to set up these test environments on your own machine, see the README files in the respective branches:

- Vagrant: [vagrant-tests](https://github.com/bertvv/ansible-role-httpd/tree/vagrant-tests)
- Docker: [docker-tests](https://github.com/bertvv/ansible-role-httpd/tree/docker-tests)

## Contributing

Issues, feature requests, ideas are appreciated and can be posted in the Issues section.

Pull requests are also very welcome. The best way to submit a PR is by first creating a fork of this Github project, then creating a topic branch for the suggested change and pushing that branch to your own fork. Github can then easily create a PR based on that branch.

## License

2-clause BSD license, see [LICENSE.md](LICENSE.md)

## Contributors

- [Bert Van Vreckem](https://github.com/bertvv/) (maintainer)
- [Richard Marko](https://github.com/sorki)
- [Lander Van den Bulcke](https://github.com/landervdb/)
