// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing a GitLab Self Hosted IDP on the instance.
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
// 		_, err := zitadel.NewIdpGitlabSelfHosted(ctx, "gitlabSelfHosted", &zitadel.IdpGitlabSelfHostedArgs{
// 			ClientId:          pulumi.String("15765e..."),
// 			ClientSecret:      pulumi.String("*****abcxyz"),
// 			IsAutoCreation:    pulumi.Bool(false),
// 			IsAutoUpdate:      pulumi.Bool(true),
// 			IsCreationAllowed: pulumi.Bool(true),
// 			IsLinkingAllowed:  pulumi.Bool(false),
// 			Issuer:            pulumi.String("https://my.issuer"),
// 			Scopes: pulumi.StringArray{
// 				pulumi.String("openid"),
// 				pulumi.String("profile"),
// 				pulumi.String("email"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type IdpGitlabSelfHosted struct {
	pulumi.CustomResourceState

	// client id generated by the identity provider
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolOutput `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolOutput `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolOutput `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolOutput `pulumi:"isLinkingAllowed"`
	// the providers issuer
	Issuer pulumi.StringOutput `pulumi:"issuer"`
	// Name of the IDP
	Name pulumi.StringOutput `pulumi:"name"`
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
}

// NewIdpGitlabSelfHosted registers a new resource with the given unique name, arguments, and options.
func NewIdpGitlabSelfHosted(ctx *pulumi.Context,
	name string, args *IdpGitlabSelfHostedArgs, opts ...pulumi.ResourceOption) (*IdpGitlabSelfHosted, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'ClientSecret'")
	}
	if args.IsAutoCreation == nil {
		return nil, errors.New("invalid value for required argument 'IsAutoCreation'")
	}
	if args.IsAutoUpdate == nil {
		return nil, errors.New("invalid value for required argument 'IsAutoUpdate'")
	}
	if args.IsCreationAllowed == nil {
		return nil, errors.New("invalid value for required argument 'IsCreationAllowed'")
	}
	if args.IsLinkingAllowed == nil {
		return nil, errors.New("invalid value for required argument 'IsLinkingAllowed'")
	}
	if args.Issuer == nil {
		return nil, errors.New("invalid value for required argument 'Issuer'")
	}
	var resource IdpGitlabSelfHosted
	err := ctx.RegisterResource("zitadel:index/idpGitlabSelfHosted:IdpGitlabSelfHosted", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdpGitlabSelfHosted gets an existing IdpGitlabSelfHosted resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdpGitlabSelfHosted(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdpGitlabSelfHostedState, opts ...pulumi.ResourceOption) (*IdpGitlabSelfHosted, error) {
	var resource IdpGitlabSelfHosted
	err := ctx.ReadResource("zitadel:index/idpGitlabSelfHosted:IdpGitlabSelfHosted", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdpGitlabSelfHosted resources.
type idpGitlabSelfHostedState struct {
	// client id generated by the identity provider
	ClientId *string `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret *string `pulumi:"clientSecret"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation *bool `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate *bool `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed *bool `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed *bool `pulumi:"isLinkingAllowed"`
	// the providers issuer
	Issuer *string `pulumi:"issuer"`
	// Name of the IDP
	Name *string `pulumi:"name"`
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
}

type IdpGitlabSelfHostedState struct {
	// client id generated by the identity provider
	ClientId pulumi.StringPtrInput
	// client secret generated by the identity provider
	ClientSecret pulumi.StringPtrInput
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolPtrInput
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolPtrInput
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolPtrInput
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolPtrInput
	// the providers issuer
	Issuer pulumi.StringPtrInput
	// Name of the IDP
	Name pulumi.StringPtrInput
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayInput
}

func (IdpGitlabSelfHostedState) ElementType() reflect.Type {
	return reflect.TypeOf((*idpGitlabSelfHostedState)(nil)).Elem()
}

type idpGitlabSelfHostedArgs struct {
	// client id generated by the identity provider
	ClientId string `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret string `pulumi:"clientSecret"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation bool `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate bool `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed bool `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed bool `pulumi:"isLinkingAllowed"`
	// the providers issuer
	Issuer string `pulumi:"issuer"`
	// Name of the IDP
	Name *string `pulumi:"name"`
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
}

// The set of arguments for constructing a IdpGitlabSelfHosted resource.
type IdpGitlabSelfHostedArgs struct {
	// client id generated by the identity provider
	ClientId pulumi.StringInput
	// client secret generated by the identity provider
	ClientSecret pulumi.StringInput
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolInput
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolInput
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolInput
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolInput
	// the providers issuer
	Issuer pulumi.StringInput
	// Name of the IDP
	Name pulumi.StringPtrInput
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayInput
}

func (IdpGitlabSelfHostedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*idpGitlabSelfHostedArgs)(nil)).Elem()
}

type IdpGitlabSelfHostedInput interface {
	pulumi.Input

	ToIdpGitlabSelfHostedOutput() IdpGitlabSelfHostedOutput
	ToIdpGitlabSelfHostedOutputWithContext(ctx context.Context) IdpGitlabSelfHostedOutput
}

func (*IdpGitlabSelfHosted) ElementType() reflect.Type {
	return reflect.TypeOf((**IdpGitlabSelfHosted)(nil)).Elem()
}

func (i *IdpGitlabSelfHosted) ToIdpGitlabSelfHostedOutput() IdpGitlabSelfHostedOutput {
	return i.ToIdpGitlabSelfHostedOutputWithContext(context.Background())
}

func (i *IdpGitlabSelfHosted) ToIdpGitlabSelfHostedOutputWithContext(ctx context.Context) IdpGitlabSelfHostedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpGitlabSelfHostedOutput)
}

// IdpGitlabSelfHostedArrayInput is an input type that accepts IdpGitlabSelfHostedArray and IdpGitlabSelfHostedArrayOutput values.
// You can construct a concrete instance of `IdpGitlabSelfHostedArrayInput` via:
//
//          IdpGitlabSelfHostedArray{ IdpGitlabSelfHostedArgs{...} }
type IdpGitlabSelfHostedArrayInput interface {
	pulumi.Input

	ToIdpGitlabSelfHostedArrayOutput() IdpGitlabSelfHostedArrayOutput
	ToIdpGitlabSelfHostedArrayOutputWithContext(context.Context) IdpGitlabSelfHostedArrayOutput
}

type IdpGitlabSelfHostedArray []IdpGitlabSelfHostedInput

func (IdpGitlabSelfHostedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdpGitlabSelfHosted)(nil)).Elem()
}

func (i IdpGitlabSelfHostedArray) ToIdpGitlabSelfHostedArrayOutput() IdpGitlabSelfHostedArrayOutput {
	return i.ToIdpGitlabSelfHostedArrayOutputWithContext(context.Background())
}

func (i IdpGitlabSelfHostedArray) ToIdpGitlabSelfHostedArrayOutputWithContext(ctx context.Context) IdpGitlabSelfHostedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpGitlabSelfHostedArrayOutput)
}

// IdpGitlabSelfHostedMapInput is an input type that accepts IdpGitlabSelfHostedMap and IdpGitlabSelfHostedMapOutput values.
// You can construct a concrete instance of `IdpGitlabSelfHostedMapInput` via:
//
//          IdpGitlabSelfHostedMap{ "key": IdpGitlabSelfHostedArgs{...} }
type IdpGitlabSelfHostedMapInput interface {
	pulumi.Input

	ToIdpGitlabSelfHostedMapOutput() IdpGitlabSelfHostedMapOutput
	ToIdpGitlabSelfHostedMapOutputWithContext(context.Context) IdpGitlabSelfHostedMapOutput
}

type IdpGitlabSelfHostedMap map[string]IdpGitlabSelfHostedInput

func (IdpGitlabSelfHostedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdpGitlabSelfHosted)(nil)).Elem()
}

func (i IdpGitlabSelfHostedMap) ToIdpGitlabSelfHostedMapOutput() IdpGitlabSelfHostedMapOutput {
	return i.ToIdpGitlabSelfHostedMapOutputWithContext(context.Background())
}

func (i IdpGitlabSelfHostedMap) ToIdpGitlabSelfHostedMapOutputWithContext(ctx context.Context) IdpGitlabSelfHostedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpGitlabSelfHostedMapOutput)
}

type IdpGitlabSelfHostedOutput struct{ *pulumi.OutputState }

func (IdpGitlabSelfHostedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdpGitlabSelfHosted)(nil)).Elem()
}

func (o IdpGitlabSelfHostedOutput) ToIdpGitlabSelfHostedOutput() IdpGitlabSelfHostedOutput {
	return o
}

func (o IdpGitlabSelfHostedOutput) ToIdpGitlabSelfHostedOutputWithContext(ctx context.Context) IdpGitlabSelfHostedOutput {
	return o
}

// client id generated by the identity provider
func (o IdpGitlabSelfHostedOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGitlabSelfHosted) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o IdpGitlabSelfHostedOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGitlabSelfHosted) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// enable if a new account in ZITADEL should be created automatically on login with an external account
func (o IdpGitlabSelfHostedOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGitlabSelfHosted) pulumi.BoolOutput { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enable if a the ZITADEL account fields should be updated automatically on each login
func (o IdpGitlabSelfHostedOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGitlabSelfHosted) pulumi.BoolOutput { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enable if users should be able to create a new account in ZITADEL when using an external account
func (o IdpGitlabSelfHostedOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGitlabSelfHosted) pulumi.BoolOutput { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enable if users should be able to link an existing ZITADEL user with an external account
func (o IdpGitlabSelfHostedOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGitlabSelfHosted) pulumi.BoolOutput { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// the providers issuer
func (o IdpGitlabSelfHostedOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGitlabSelfHosted) pulumi.StringOutput { return v.Issuer }).(pulumi.StringOutput)
}

// Name of the IDP
func (o IdpGitlabSelfHostedOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGitlabSelfHosted) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o IdpGitlabSelfHostedOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IdpGitlabSelfHosted) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

type IdpGitlabSelfHostedArrayOutput struct{ *pulumi.OutputState }

func (IdpGitlabSelfHostedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdpGitlabSelfHosted)(nil)).Elem()
}

func (o IdpGitlabSelfHostedArrayOutput) ToIdpGitlabSelfHostedArrayOutput() IdpGitlabSelfHostedArrayOutput {
	return o
}

func (o IdpGitlabSelfHostedArrayOutput) ToIdpGitlabSelfHostedArrayOutputWithContext(ctx context.Context) IdpGitlabSelfHostedArrayOutput {
	return o
}

func (o IdpGitlabSelfHostedArrayOutput) Index(i pulumi.IntInput) IdpGitlabSelfHostedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IdpGitlabSelfHosted {
		return vs[0].([]*IdpGitlabSelfHosted)[vs[1].(int)]
	}).(IdpGitlabSelfHostedOutput)
}

type IdpGitlabSelfHostedMapOutput struct{ *pulumi.OutputState }

func (IdpGitlabSelfHostedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdpGitlabSelfHosted)(nil)).Elem()
}

func (o IdpGitlabSelfHostedMapOutput) ToIdpGitlabSelfHostedMapOutput() IdpGitlabSelfHostedMapOutput {
	return o
}

func (o IdpGitlabSelfHostedMapOutput) ToIdpGitlabSelfHostedMapOutputWithContext(ctx context.Context) IdpGitlabSelfHostedMapOutput {
	return o
}

func (o IdpGitlabSelfHostedMapOutput) MapIndex(k pulumi.StringInput) IdpGitlabSelfHostedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IdpGitlabSelfHosted {
		return vs[0].(map[string]*IdpGitlabSelfHosted)[vs[1].(string)]
	}).(IdpGitlabSelfHostedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdpGitlabSelfHostedInput)(nil)).Elem(), &IdpGitlabSelfHosted{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdpGitlabSelfHostedArrayInput)(nil)).Elem(), IdpGitlabSelfHostedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdpGitlabSelfHostedMapInput)(nil)).Elem(), IdpGitlabSelfHostedMap{})
	pulumi.RegisterOutputType(IdpGitlabSelfHostedOutput{})
	pulumi.RegisterOutputType(IdpGitlabSelfHostedArrayOutput{})
	pulumi.RegisterOutputType(IdpGitlabSelfHostedMapOutput{})
}
