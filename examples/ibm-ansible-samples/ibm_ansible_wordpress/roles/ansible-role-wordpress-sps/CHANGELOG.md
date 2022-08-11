# Change log

## SPS - force update of Salts file if already existing to make execution idempotent
## 2018-08-21

This file contains al notable changes to the wordpress Ansible role.

This file adheres to the guidelines of [http://keepachangelog.com/](http://keepachangelog.com/). Versioning follows [Semantic Versioning](http://semver.org/).

## 1.2.0 - 2017-01-25

### Added

- (GH-7) Added variable `mariadb_database_host` (credit [Kwinten Guillaume](https://github.com/kwinteng))

### Changed

- Removed hard-coded paths to config files
- Set SELinux boolean `httpd_can_network_connect_db` when necessary
- Check whether Apache is already installed

### Removed

- (GH-7) Removed dependency on `bertvv.mariadb` role

## 1.1.4 - 2016-05-10

### Added

- Explicit support for Fedora and CentOS 7, tests for these platforms.

### Changes

- Removed Ansible 2.0 deprecation warnings

## 1.1.3 - 2015-10-30

This is a bugfix release

### Changes

- Fixed #5 (attempting to log in to the admin page redirects to the login page without error message)

## 1.1.2 - 2015-10-11

This is a bugfix release

### Changes

- Fixed #2 (downloading plugins/themes without a version number). Credit to [Jordi Stevens](https://github.com/Xplendit)
- Fixed #4 (Playbook sometimes crashes when getting new salts). As a consequence of the changes, the playbook will no longer fetch new salts every time it is run. When you want to get new salts, delete /usr/share/wordpress/wp-salts.php and re-run the playbook.
- Replace hard-coded values of Wordpress installation directory with a variable

## 1.1.1 - 2015-10-07

### Changes

- Fixed missing value of `wordpress_themes`

## 1.1.0 - 2015-10-07

### Added

- Install plugins with role variable `wordpress_plugins`
- Install themes with role variable `wordpress_themes`

## 1.0.0 - 2015-04-28

First release!

### Added

- Installs Wordpress and generates `wp-config.php` with safe secret keys and salts

