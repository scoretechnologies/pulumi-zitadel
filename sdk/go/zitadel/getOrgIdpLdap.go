// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource representing an LDAP IdP on the organization.
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
// 		_, err := zitadel.LookupOrgIdpLdap(ctx, &GetOrgIdpLdapArgs{
// 			Id: "177073614158299139",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupOrgIdpLdap(ctx *pulumi.Context, args *LookupOrgIdpLdapArgs, opts ...pulumi.InvokeOption) (*LookupOrgIdpLdapResult, error) {
	var rv LookupOrgIdpLdapResult
	err := ctx.Invoke("zitadel:index/getOrgIdpLdap:getOrgIdpLdap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrgIdpLdap.
type LookupOrgIdpLdapArgs struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
	// ID of the organization
	OrgId string `pulumi:"orgId"`
}

// A collection of values returned by getOrgIdpLdap.
type LookupOrgIdpLdapResult struct {
	// User attribute for the avatar url
	AvatarUrlAttribute string `pulumi:"avatarUrlAttribute"`
	// Base DN for LDAP connections
	BaseDn string `pulumi:"baseDn"`
	// Bind DN for LDAP connections
	BindDn string `pulumi:"bindDn"`
	// Bind password for LDAP connections
	BindPassword string `pulumi:"bindPassword"`
	// User attribute for the display name
	DisplayNameAttribute string `pulumi:"displayNameAttribute"`
	// User attribute for the email
	EmailAttribute string `pulumi:"emailAttribute"`
	// User attribute for the email verified state
	EmailVerifiedAttribute string `pulumi:"emailVerifiedAttribute"`
	// User attribute for the first name
	FirstNameAttribute string `pulumi:"firstNameAttribute"`
	// The ID of this resource.
	Id string `pulumi:"id"`
	// User attribute for the id
	IdAttribute string `pulumi:"idAttribute"`
	// enabled if a new account in ZITADEL are created automatically on login with an external account
	IsAutoCreation bool `pulumi:"isAutoCreation"`
	// enabled if a the ZITADEL account fields are updated automatically on each login
	IsAutoUpdate bool `pulumi:"isAutoUpdate"`
	// enabled if users are able to create a new account in ZITADEL when using an external account
	IsCreationAllowed bool `pulumi:"isCreationAllowed"`
	// enabled if users are able to link an existing ZITADEL user with an external account
	IsLinkingAllowed bool `pulumi:"isLinkingAllowed"`
	// User attribute for the last name
	LastNameAttribute string `pulumi:"lastNameAttribute"`
	// Name of the IDP
	Name string `pulumi:"name"`
	// User attribute for the nick name
	NickNameAttribute string `pulumi:"nickNameAttribute"`
	// ID of the organization
	OrgId string `pulumi:"orgId"`
	// User attribute for the phone
	PhoneAttribute string `pulumi:"phoneAttribute"`
	// User attribute for the phone verified state
	PhoneVerifiedAttribute string `pulumi:"phoneVerifiedAttribute"`
	// User attribute for the preferred language
	PreferredLanguageAttribute string `pulumi:"preferredLanguageAttribute"`
	// User attribute for the preferred username
	PreferredUsernameAttribute string `pulumi:"preferredUsernameAttribute"`
	// User attribute for the profile
	ProfileAttribute string `pulumi:"profileAttribute"`
	// Servers to try in order for establishing LDAP connections
	Servers []string `pulumi:"servers"`
	// Wether to use StartTLS for LDAP connections
	StartTls bool `pulumi:"startTls"`
	// Timeout for LDAP connections
	Timeout string `pulumi:"timeout"`
	// User base for LDAP connections
	UserBase string `pulumi:"userBase"`
	// User filters for LDAP connections
	UserFilters []string `pulumi:"userFilters"`
	// User object classes for LDAP connections
	UserObjectClasses []string `pulumi:"userObjectClasses"`
}

func LookupOrgIdpLdapOutput(ctx *pulumi.Context, args LookupOrgIdpLdapOutputArgs, opts ...pulumi.InvokeOption) LookupOrgIdpLdapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrgIdpLdapResult, error) {
			args := v.(LookupOrgIdpLdapArgs)
			r, err := LookupOrgIdpLdap(ctx, &args, opts...)
			var s LookupOrgIdpLdapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrgIdpLdapResultOutput)
}

// A collection of arguments for invoking getOrgIdpLdap.
type LookupOrgIdpLdapOutputArgs struct {
	// The ID of this resource.
	Id pulumi.StringInput `pulumi:"id"`
	// ID of the organization
	OrgId pulumi.StringInput `pulumi:"orgId"`
}

func (LookupOrgIdpLdapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgIdpLdapArgs)(nil)).Elem()
}

// A collection of values returned by getOrgIdpLdap.
type LookupOrgIdpLdapResultOutput struct{ *pulumi.OutputState }

func (LookupOrgIdpLdapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgIdpLdapResult)(nil)).Elem()
}

func (o LookupOrgIdpLdapResultOutput) ToLookupOrgIdpLdapResultOutput() LookupOrgIdpLdapResultOutput {
	return o
}

func (o LookupOrgIdpLdapResultOutput) ToLookupOrgIdpLdapResultOutputWithContext(ctx context.Context) LookupOrgIdpLdapResultOutput {
	return o
}

// User attribute for the avatar url
func (o LookupOrgIdpLdapResultOutput) AvatarUrlAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.AvatarUrlAttribute }).(pulumi.StringOutput)
}

// Base DN for LDAP connections
func (o LookupOrgIdpLdapResultOutput) BaseDn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.BaseDn }).(pulumi.StringOutput)
}

// Bind DN for LDAP connections
func (o LookupOrgIdpLdapResultOutput) BindDn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.BindDn }).(pulumi.StringOutput)
}

// Bind password for LDAP connections
func (o LookupOrgIdpLdapResultOutput) BindPassword() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.BindPassword }).(pulumi.StringOutput)
}

// User attribute for the display name
func (o LookupOrgIdpLdapResultOutput) DisplayNameAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.DisplayNameAttribute }).(pulumi.StringOutput)
}

// User attribute for the email
func (o LookupOrgIdpLdapResultOutput) EmailAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.EmailAttribute }).(pulumi.StringOutput)
}

// User attribute for the email verified state
func (o LookupOrgIdpLdapResultOutput) EmailVerifiedAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.EmailVerifiedAttribute }).(pulumi.StringOutput)
}

// User attribute for the first name
func (o LookupOrgIdpLdapResultOutput) FirstNameAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.FirstNameAttribute }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o LookupOrgIdpLdapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.Id }).(pulumi.StringOutput)
}

// User attribute for the id
func (o LookupOrgIdpLdapResultOutput) IdAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.IdAttribute }).(pulumi.StringOutput)
}

// enabled if a new account in ZITADEL are created automatically on login with an external account
func (o LookupOrgIdpLdapResultOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) bool { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enabled if a the ZITADEL account fields are updated automatically on each login
func (o LookupOrgIdpLdapResultOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) bool { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enabled if users are able to create a new account in ZITADEL when using an external account
func (o LookupOrgIdpLdapResultOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) bool { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enabled if users are able to link an existing ZITADEL user with an external account
func (o LookupOrgIdpLdapResultOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) bool { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// User attribute for the last name
func (o LookupOrgIdpLdapResultOutput) LastNameAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.LastNameAttribute }).(pulumi.StringOutput)
}

// Name of the IDP
func (o LookupOrgIdpLdapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.Name }).(pulumi.StringOutput)
}

// User attribute for the nick name
func (o LookupOrgIdpLdapResultOutput) NickNameAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.NickNameAttribute }).(pulumi.StringOutput)
}

// ID of the organization
func (o LookupOrgIdpLdapResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// User attribute for the phone
func (o LookupOrgIdpLdapResultOutput) PhoneAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.PhoneAttribute }).(pulumi.StringOutput)
}

// User attribute for the phone verified state
func (o LookupOrgIdpLdapResultOutput) PhoneVerifiedAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.PhoneVerifiedAttribute }).(pulumi.StringOutput)
}

// User attribute for the preferred language
func (o LookupOrgIdpLdapResultOutput) PreferredLanguageAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.PreferredLanguageAttribute }).(pulumi.StringOutput)
}

// User attribute for the preferred username
func (o LookupOrgIdpLdapResultOutput) PreferredUsernameAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.PreferredUsernameAttribute }).(pulumi.StringOutput)
}

// User attribute for the profile
func (o LookupOrgIdpLdapResultOutput) ProfileAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.ProfileAttribute }).(pulumi.StringOutput)
}

// Servers to try in order for establishing LDAP connections
func (o LookupOrgIdpLdapResultOutput) Servers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) []string { return v.Servers }).(pulumi.StringArrayOutput)
}

// Wether to use StartTLS for LDAP connections
func (o LookupOrgIdpLdapResultOutput) StartTls() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) bool { return v.StartTls }).(pulumi.BoolOutput)
}

// Timeout for LDAP connections
func (o LookupOrgIdpLdapResultOutput) Timeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.Timeout }).(pulumi.StringOutput)
}

// User base for LDAP connections
func (o LookupOrgIdpLdapResultOutput) UserBase() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) string { return v.UserBase }).(pulumi.StringOutput)
}

// User filters for LDAP connections
func (o LookupOrgIdpLdapResultOutput) UserFilters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) []string { return v.UserFilters }).(pulumi.StringArrayOutput)
}

// User object classes for LDAP connections
func (o LookupOrgIdpLdapResultOutput) UserObjectClasses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOrgIdpLdapResult) []string { return v.UserObjectClasses }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrgIdpLdapResultOutput{})
}
