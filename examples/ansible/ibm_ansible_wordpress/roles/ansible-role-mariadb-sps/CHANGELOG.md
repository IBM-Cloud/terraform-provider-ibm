# Change log

This file contains al notable changes to the mariadb Ansible role.

This file adheres to the guidelines of <http://keepachangelog.com/>. Versioning follows [Semantic Versioning](http://semver.org/).

## SPS - Updated sever_my.cnf file to support DB replication 
## 2018-08-21


## 2.1.0 - 2018-06-01

### Added

- (GH-11) Allow default character set and collation to be configured (credit: Ripon Banik)
- (GH-12) Added several server configuration variables for better performance. See `templates/etc_my.cnf.d_server.cnf.j2` (credit: Tom Stechele)
- (GH-13) Allow the role to skip setting the swappiness of the mariadb process.
- (GH-16) Allow user privileges to be added instead of overwritten (credit: Cédric Delgehier)
- Allow custom configuration not specified in the template

### Changed

- Bump default MariaDB version to 10.3, the current stable
- Fix MariaDB service name in handler

## 2.0.2 - 2017-11-21

### Changed

- Removed Ansible 2.4 deprecation warnings (include -> include_tasks)

## 2.0.1 - 2017-09-06

### Changed

- (GH-7, GH-8) Fix for downloading from the MariaDB repository on RHEL systems (credit: [@raznikk](https://github.com/raznikk), [@piuma](https://github.com/piuma))

## 2.0.0 - 2017-07-13

### Added

- (GH-5) MariaDB is installed from the project repositories and the version can be chosen.
- (GH-3) Added template for server.cnf (credit: [Louis Tournayre](https://github.com/louiznk))
    - Made settings in server.cnf configurable and added default values.
    - (GH-1) Added performance related settings (credit: [Thomas Eylenbosch](https://github.com/EylenboschThomas))
- Added Docker based playbook and functional tests executed on Travis CI on each commit for all supported platforms (CentOS, Fedora)
- Swappiness can be configured, and is turned off by default. This means that swapping processes to disk is avoided at all cost.

### Changed

- (GH-2) Made database initialisation idempotent. **In order to implement this, a breaking change was introduced in the way `mariadb_databases` is defined.**
- (GH-4) Remove initialisation scripts from the server after executing them.
- (GH-6) Test code was moved to a separate branch
- The database root password is no longer stored on the server in `~/.my.cnf`
- Fix ignored `host:` option in `mariadb_users`

## 1.1.3 - 2016-10-01

### Added

- In variable `mysql_users`, it is now possible to set a `host` pattern that specifies the hosts from where a user is allowed to query the database.

## 1.1.2 - 2016-05-10

### Added

- Support for Fedora

### Changed

- Fixed Ansible 2.0 deprecation warnings
- Prepared the role for supporting multiple platforms

## 1.1.1 - 2015-04-28

### Changed

- Fixed a bug that makes the role fail when `mariadb_init_scripts` is undefined.

## 1.1.0 - 2015-04-26

### Added

- Databases can be initialised with a SQL script

## 1.0.0 - 2015-04-25

First release!

### Added

- Install MariaDB
- Secure installation (change root password, delete anonymous user, test table)
- Manage bind_port
- Create databases and users

