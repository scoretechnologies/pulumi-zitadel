CHANGELOG
=========

## HEAD (Unreleased)

* Upgrade the upstream `terraform-provider-zitadel` from v2.2.0 to v3.3.0, adding 45
  resources and 34 data sources. The upgrade is additive: no existing resource, data
  source, or configuration option was renamed or removed.
* Bridge the upstream Plugin Framework provider alongside the SDKv2 one, so the
  message text and login text resources are now available.
* New provider configuration options: `accessToken`, `jwtFile`, `systemApi`,
  `insecureSkipVerifyTls` and `transportHeaders`. `token` is deprecated upstream in
  favour of `accessToken`.
* Upgrade `pulumi-terraform-bridge` to v3.135.0 and derive resource tokens
  automatically, so resources added upstream no longer need to be listed by hand.
  Building now requires Go 1.25.
* fix dotnetversion typo in release template #94
* fix `/usr/share/dotnet/host/xfr does not exist` issue when running `make build_dotnet`

---
