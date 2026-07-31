# Zitadel Resource Provider

The Zitadel Resource Provider lets you manage [Zitadel](https://zitadel.com/) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @scoretechnologies/zitadel
```

or `yarn`:

```bash
yarn add @scoretechnologies/zitadel
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_zitadel
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/scoretechnologies/pulumi-zitadel/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package scoretechnologies.Zitadel
```

## Configuration

The following configuration points are available for the `zitadel` provider:

- `zitadel:domain` - domain used to connect to the ZITADEL instance (required)
- `zitadel:port` - used port if not the default ports 80 or 443 are configured
- `zitadel:insecure` - use insecure connection
- `zitadel:insecureSkipVerifyTls` - disable TLS certificate verification. Only use in development or testing environments with self-signed certificates
- `zitadel:transportHeaders` - custom headers added to both the HTTP authentication and gRPC API requests, for example a `Proxy-Authorization` header when fronted by GCP IAP

Exactly one of the following credentials is required:

- `zitadel:accessToken` - Personal Access Token to connect to ZITADEL
- `zitadel:jwtProfileFile` - path to the file containing JWT Profile credentials
- `zitadel:jwtProfileJson` - JSON value of the JWT Profile credentials
- `zitadel:jwtFile` - path to a file containing a presigned JWT
- `zitadel:systemApi` - configuration block for authenticating against the ZITADEL System API with a PEM encoded key
- `zitadel:token` - **deprecated**, use `accessToken` for Personal Access Tokens or `jwtProfileFile` for JWT Profile credentials instead

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/zitadel/api-docs/).
