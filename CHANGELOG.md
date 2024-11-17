# Changelog

## 24.11.2

## Improvements

- [x] using ingress certificate for TLS Ingress Endpoint

## Breaking changes

- [x] revome entrypoint from config, shell will only on the first container of the setup

## 24.11.1

## Improvements

- [x] IPv6 support in setup definition
- [x] customizable egress proxy in setup definition
- [x] introduce examples
- [x] install uninstall verb and config to have multiple hab on same host
- [x] add Run verbs to run as daemon

## 24.11.0

## Improvements

- [x] A little bit more docs
- [x] Back to LXD Snap distribution for better package managements

## Refacotring

- [x] Runtime Controller reference by constant rather than string
- [x] Refactoring Controller getters on Controllers consumer
- [x] Grouping commands name and prefix in config
- [x] Agnosticity for plateform runtime

## 24.9.0

### Improvment

- [x] Ubuntu 24.04.1 base testing
- [x] Using Incus from APT rather than LXD from SNAP
- [x] implement unprovision network for plateform
- [x] ensure eveything is clean after unprovision of incus

### Breaking changes

- [x] default config massively changed
