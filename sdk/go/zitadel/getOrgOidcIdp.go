// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource representing a generic OIDC IdP on the organization.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-zitadel/sdk/go/zitadel"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		orgOidcIdpOrgOidcIdp, err := zitadel.GetOrgOidcIdp(ctx, &GetOrgOidcIdpArgs{
// 			OrgId: data.Zitadel_org.Org.Id,
// 			IdpId: "177073612581240835",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("orgOidcIdp", orgOidcIdpOrgOidcIdp)
// 		return nil
// 	})
// }
// ```
func GetOrgOidcIdp(ctx *pulumi.Context, args *GetOrgOidcIdpArgs, opts ...pulumi.InvokeOption) (*GetOrgOidcIdpResult, error) {
	var rv GetOrgOidcIdpResult
	err := ctx.Invoke("zitadel:index/getOrgOidcIdp:getOrgOidcIdp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrgOidcIdp.
type GetOrgOidcIdpArgs struct {
	// The ID of this resource.
	IdpId string `pulumi:"idpId"`
	// ID of the organization
	OrgId string `pulumi:"orgId"`
}

// A collection of values returned by getOrgOidcIdp.
type GetOrgOidcIdpResult struct {
	// auto register for users from this idp
	AutoRegister bool `pulumi:"autoRegister"`
	// client id generated by the identity provider
	ClientId string `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret string `pulumi:"clientSecret"`
	// definition which field is mapped to the display name of the user
	DisplayNameMapping string `pulumi:"displayNameMapping"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ID of this resource.
	IdpId string `pulumi:"idpId"`
	// the oidc issuer of the identity provider
	Issuer string `pulumi:"issuer"`
	// Name of the IDP
	Name string `pulumi:"name"`
	// ID of the organization
	OrgId string `pulumi:"orgId"`
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
	// Some identity providers specify the styling of the button to their login
	StylingType string `pulumi:"stylingType"`
	// definition which field is mapped to the email of the user
	UsernameMapping string `pulumi:"usernameMapping"`
}

func GetOrgOidcIdpOutput(ctx *pulumi.Context, args GetOrgOidcIdpOutputArgs, opts ...pulumi.InvokeOption) GetOrgOidcIdpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOrgOidcIdpResult, error) {
			args := v.(GetOrgOidcIdpArgs)
			r, err := GetOrgOidcIdp(ctx, &args, opts...)
			var s GetOrgOidcIdpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOrgOidcIdpResultOutput)
}

// A collection of arguments for invoking getOrgOidcIdp.
type GetOrgOidcIdpOutputArgs struct {
	// The ID of this resource.
	IdpId pulumi.StringInput `pulumi:"idpId"`
	// ID of the organization
	OrgId pulumi.StringInput `pulumi:"orgId"`
}

func (GetOrgOidcIdpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrgOidcIdpArgs)(nil)).Elem()
}

// A collection of values returned by getOrgOidcIdp.
type GetOrgOidcIdpResultOutput struct{ *pulumi.OutputState }

func (GetOrgOidcIdpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrgOidcIdpResult)(nil)).Elem()
}

func (o GetOrgOidcIdpResultOutput) ToGetOrgOidcIdpResultOutput() GetOrgOidcIdpResultOutput {
	return o
}

func (o GetOrgOidcIdpResultOutput) ToGetOrgOidcIdpResultOutputWithContext(ctx context.Context) GetOrgOidcIdpResultOutput {
	return o
}

// auto register for users from this idp
func (o GetOrgOidcIdpResultOutput) AutoRegister() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) bool { return v.AutoRegister }).(pulumi.BoolOutput)
}

// client id generated by the identity provider
func (o GetOrgOidcIdpResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o GetOrgOidcIdpResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

// definition which field is mapped to the display name of the user
func (o GetOrgOidcIdpResultOutput) DisplayNameMapping() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) string { return v.DisplayNameMapping }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOrgOidcIdpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o GetOrgOidcIdpResultOutput) IdpId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) string { return v.IdpId }).(pulumi.StringOutput)
}

// the oidc issuer of the identity provider
func (o GetOrgOidcIdpResultOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) string { return v.Issuer }).(pulumi.StringOutput)
}

// Name of the IDP
func (o GetOrgOidcIdpResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o GetOrgOidcIdpResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o GetOrgOidcIdpResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

// Some identity providers specify the styling of the button to their login
func (o GetOrgOidcIdpResultOutput) StylingType() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) string { return v.StylingType }).(pulumi.StringOutput)
}

// definition which field is mapped to the email of the user
func (o GetOrgOidcIdpResultOutput) UsernameMapping() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrgOidcIdpResult) string { return v.UsernameMapping }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrgOidcIdpResultOutput{})
}
