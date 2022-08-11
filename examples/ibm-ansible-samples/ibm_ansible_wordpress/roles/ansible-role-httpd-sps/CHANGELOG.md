# Change log

This file contains al notable changes to the bertvv.httpd Ansible role.

This file adheres to the guidelines of [http://keepachangelog.com/](http://keepachangelog.com/). Versioning follows [Semantic Versioning](http://semver.org/).

## 1.2.1 - 2015-05-10

### Added

- tests for supported platforms

### Changed

- Installing scripting support is better generalized.
- Moved platform specific values to `vars/RedHat.yml`.

## 1.2.0 - 2015-03-15

### Added

- Fedora support (credit [Richard Marko](https://github.com/sorki))
- more SSL configuration options (credit [Richard Marko](https://github.com/sorki)
- `httpd_ServerTokens` variable (credit [Richard Marko](https://github.com/sorki)

## 1.1.0 - 2015-03-06

### Added

- Optional support for PHP

## 1.0.0 - 2015-03-06

First release!

### Added

- Install `httpd` and `mod_ssl` packages
- Configure Apache with template for `httpd.conf`
- Configure `mode_ssl` with template for `ssl.conf`
- Basic settings like port number, log file locations, and certificates can be customized


