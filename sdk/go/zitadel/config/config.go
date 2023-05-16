// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// Domain used to connect to the ZITADEL instance
func GetDomain(ctx *pulumi.Context) string {
	return config.Get(ctx, "zitadel:domain")
}

// Use insecure connection
func GetInsecure(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "zitadel:insecure")
}

// Path to the file containing credentials to connect to ZITADEL. Either 'jwt_profile_file' or 'jwt_profile_json' is
// required
func GetJwtProfileFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "zitadel:jwtProfileFile")
}

// JSON value of credentials to connect to ZITADEL. Either 'jwt_profile_file' or 'jwt_profile_json' is required
func GetJwtProfileJson(ctx *pulumi.Context) string {
	return config.Get(ctx, "zitadel:jwtProfileJson")
}

// Used port if not the default ports 80 or 443 are configured
func GetPort(ctx *pulumi.Context) string {
	return config.Get(ctx, "zitadel:port")
}

// Path to the file containing credentials to connect to ZITADEL
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "zitadel:token")
}