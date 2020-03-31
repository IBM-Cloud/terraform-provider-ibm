# [3.3.0](https://github.com/IBM/go-sdk-core/compare/v3.2.4...v3.3.0) (2020-03-29)


### Features

* add unmarshal methods for maps of primitive types ([0afd3f7](https://github.com/IBM/go-sdk-core/commit/0afd3f7cc650ca9fdf868d6a2276c940cdb52651))

## [3.2.4](https://github.com/IBM/go-sdk-core/compare/v3.2.3...v3.2.4) (2020-02-24)


### Bug Fixes

* tolerate explicit JSON null values in UnmarshalXXX() methods ([3967601](https://github.com/IBM/go-sdk-core/commit/39676013711af6cb685c8c5ec7c631e226b266df))

## [3.2.3](https://github.com/IBM/go-sdk-core/compare/v3.2.2...v3.2.3) (2020-02-19)


### Bug Fixes

* Fix token caching ([880b0be](https://github.com/IBM/go-sdk-core/commit/880b0bed51187332f26ba140d01b47e079f8df0c))

## [3.2.2](https://github.com/IBM/go-sdk-core/compare/v3.2.1...v3.2.2) (2020-02-13)


### Bug Fixes

* correct go.mod ([64ff92d](https://github.com/IBM/go-sdk-core/commit/64ff92decff6e1595f3f1f7764b5839864bcca20))

## [3.2.1](https://github.com/IBM/go-sdk-core/compare/v3.2.0...v3.2.1) (2020-02-13)


### Bug Fixes

* tolerate non-compliant error response bodies ([f0e3a13](https://github.com/IBM/go-sdk-core/commit/f0e3a1301c028df05ddd315cda687fd6295e39ab))

# [3.2.0](https://github.com/IBM/go-sdk-core/compare/v3.1.1...v3.2.0) (2020-02-07)


### Features

* add unmarshal functions for 'any' ([55c1eee](https://github.com/IBM/go-sdk-core/commit/55c1eee879932086061c9d5849b972caf5d31094))

## [3.1.1](https://github.com/IBM/go-sdk-core/compare/v3.1.0...v3.1.1) (2020-01-09)


### Bug Fixes

* ensure version # is updated in go.mod ([8fdc596](https://github.com/IBM/go-sdk-core/commit/8fdc5961b6951cc8f2769fbad241f749cc983d9c))
* fixup version #'s to 3.1.0 ([ecdafe1](https://github.com/IBM/go-sdk-core/commit/ecdafe11762d060ff08fb56ab5bd3b37ca870bbc))

# [3.1.0](https://github.com/IBM/go-sdk-core/compare/v3.0.0...v3.1.0) (2020-01-06)


### Features

* add unmarshal utility functions for primitive types ([3f7299b](https://github.com/IBM/go-sdk-core/commit/3f7299b0203f0fec5f6a6ede6bd23f63568388c3))

# [3.0.0](https://github.com/IBM/go-sdk-core/compare/v2.1.0...v3.0.0) (2019-12-09)

### Features

* created new major version 3.0.0 in v3 directory ([1595df4](https://github.com/IBM/go-sdk-core/commit/1595df486aba57dd5b965354376f5590d435ecfb))

### BREAKING CHANGES

* created new major version 3.0.0 in v3 directory

# [2.1.0](https://github.com/IBM/go-sdk-core/compare/v2.0.1...v2.1.0) (2019-12-04)


### Features

* allow JSON response body to be streamed ([d1345d7](https://github.com/IBM/go-sdk-core/commit/d1345d7d5d7dc91959eafc0d8c1ddd79a6f31450))

## [2.0.1](https://github.com/IBM/go-sdk-core/compare/v2.0.0...v2.0.1) (2019-11-21)


### Bug Fixes

* add HEAD operation constant ([#41](https://github.com/IBM/go-sdk-core/issues/41)) ([47b5dc9](https://github.com/IBM/go-sdk-core/commit/47b5dc9e46c4fa25b3e93e2b1ff15136c16e1877))

# [2.0.0](https://github.com/IBM/go-sdk-core/compare/v1.1.0...v2.0.0) (2019-11-06)


### Features

* **loadFromVCAPServices:** Service configuration factory. ([87ac493](https://github.com/IBM/go-sdk-core/commit/87ac49304e600a4bac9e52f2a0a0b529e26f0db1))


### BREAKING CHANGES

* **loadFromVCAPServices:** NewBaseService constructor changes. `displayname`, and `serviceName` removed from construction function signature, since they are no longer used.

# [1.1.0](https://github.com/IBM/go-sdk-core/compare/v1.0.1...v1.1.0) (2019-11-06)


### Features

* **BaseService:** add new method ConfigureService() to BaseService struct ([27192a7](https://github.com/IBM/go-sdk-core/commit/27192a7a796038d172af5a579a7535f91973990f))

# [1.0.1](https://github.com/IBM/go-sdk-core/compare/v1.0.0...v1.0.1) (2019-10-18)
    
### Bug Fixes
    
* fixed ConstructHTTPURL to avoid '//' when path segment is empty string ([e618205](https://github.com/IBM/go-sdk-core/commit/e61820596fbab3d475f4c2ba1d4417d755b78557))
* use go module instead of dep for dependency management; use golangci-lint for linting ([58a9cf6](https://github.com/IBM/go-sdk-core/commit/58a9cf666216ab4a420b686347f5e050e78ef975))

# [1.0.0](https://github.com/IBM/go-sdk-core/compare/v0.8.0...v1.0.0) (2019-10-04)


### Bug Fixes

* use correct error message for SSL verification errors ([a7fe39e](https://github.com/IBM/go-sdk-core/commit/a7fe39e))


### Documentation

* Updated README with information about the authentticators ([7ef7f24](https://github.com/IBM/go-sdk-core/commit/7ef7f24))


### Features

* **creds:** Search creds in working dir before home ([bf4a377](https://github.com/IBM/go-sdk-core/commit/bf4a377))
* **creds:** Search creds in working dir before home ([fbb900b](https://github.com/IBM/go-sdk-core/commit/fbb900b))
* **disable ssl:** Add description on disabling ssl verification ([298ec08](https://github.com/IBM/go-sdk-core/commit/298ec08))


### BREAKING CHANGES

* this release introduces a new authentication scheme

# [0.9.0](https://github.com/IBM/go-sdk-core/compare/v0.8.0...v0.9.0) (2019-09-24)


### Features

* **creds:** Search creds in working dir before home ([bf4a377](https://github.com/IBM/go-sdk-core/commit/bf4a377))
* **creds:** Search creds in working dir before home ([fbb900b](https://github.com/IBM/go-sdk-core/commit/fbb900b))
* **disable ssl:** Add description on disabling ssl verification ([298ec08](https://github.com/IBM/go-sdk-core/commit/298ec08))

# [0.8.0](https://github.com/IBM/go-sdk-core/compare/v0.7.0...v0.8.0) (2019-09-23)


### Features

* Defer service URL validation ([6f51c35](https://github.com/IBM/go-sdk-core/commit/6f51c35)), closes [arf/planning-sdk-squad#1011](https://github.com/arf/planning-sdk-squad/issues/1011)
* **creds:** Search creds in working dir before home ([bf4a377](https://github.com/IBM/go-sdk-core/commit/bf4a377))
* **creds:** Search creds in working dir before home ([fbb900b](https://github.com/IBM/go-sdk-core/commit/fbb900b))
* **disable ssl:** Add description on disabling ssl verification ([298ec08](https://github.com/IBM/go-sdk-core/commit/298ec08))
