# Ansible role `wordpress`

An Ansible role for installing Wordpress. Specifically, the responsibilities of this role are to:

- install the EPEL repository and Wordpress dependencies
- install Wordpress
- set up the database and configure Apache
- fetch security keys and salts
- generate `wp-config.php`

## Requirements

no specific requirements

## Role Variables

| Variable                  | Default     | Comments (type)                                   |
| :---                      | :---        | :---                                              |
| `wordpress_database`      | 'wordpress' | The name of the database for Wordpress.           |
| `wordpress_database_host` | 'localhost' | The database server.                              |
| `wordpress_user`          | 'wordpress' | The name of the database user.                    |
| `wordpress_password`      | 'wordpress' | The password of the database user.                |
| `wordpress_plugins`       | []          | Plugins to be installed. See below. (since 1.1.0) |
| `wordpress_themes`        | []          | Themes to be installed. See below. (since 1.1.0)  |

**Remark:** it is **very strongly** suggested to change the default password.

To install plugins and themes (from the Wordpress Plugin and Theme Directory), you need to specify at least the name. Most plugins and themes also have a version, in which case you need to provide it as well. The version number should not be given if the plugins does't have one. An example:

```yaml
wordpress_plugins:
  - name: wp-super-cache
    version: 1.4.5
  - name: jetpack
    version: 3.7.2
  - name: lipsum  # Plugin without a version
wordpress_themes:
  - name: xcel
    version: 1.0.9
```

### Configuring Apache and Mariadb

The variables for this role are not mandatory, but in the dependent roles (`bertvv.httpd` and `bertvv.mariadb`), some variables have to be set:

```Yaml
httpd_scripting: 'php'
mariadb_databases:
  - wordpress_db
mariadb_users:
  - name: wordpress_usr
    password: ywIapecJalg6
    priv: wordpress_db.*:ALL
```

* PHP scripting should be enabled
* A database should be created. Variable `wordpress_database` should have the same value as `mariadb_databases`
* A database user with access to the database should be created. Variables `wordpress_user` and `wordpress_password` should have the same values as the respective settings here.

## Dependencies

- [bertvv.httpd](https://galaxy.ansible.com/list#/roles/3047)
- [bertvv.mariadb](https://galaxy.ansible.com/list#/roles/3518)

## Example Playbook

See the [test playbook](https://github.com/bertvv/ansible-role-wordpress/blob/vagrant-tests/test.yml).

## Testing

Test code is kept in a separate branch. See the associated [README](https://github.com/bertvv/ansible-role-wordpress/blob/vagrant-tests/README.md) for more information on how to set this up.

## Contributing

Issues, feature requests, ideas are appreciated and can be posted in the Issues section.

Pull requests are also very welcome. The best way to submit a PR is by first creating a fork of this Github project, then creating a topic branch for the suggested change and pushing that branch to your own fork. Github can then easily create a PR based on that branch.

## License

2-clause BSD license, see [LICENSE.md](LICENSE.md)

## Contributors

- [Bert Van Vreckem](https://github.com/bertvv/) (maintainer)
- [Jordi Stevens](https://github.com/Xplendit)
- [Kwinten Guillaume](https://github.com/kwinteng)
