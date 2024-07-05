// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel/internal"
)

// Resource representing the custom privacy policy of an organization.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := zitadel.NewPrivacyPolicy(ctx, "default", &zitadel.PrivacyPolicyArgs{
//				OrgId:        pulumi.Any(data.Zitadel_org.Default.Id),
//				TosLink:      pulumi.String("https://example.com/tos"),
//				PrivacyLink:  pulumi.String("https://example.com/privacy"),
//				HelpLink:     pulumi.String("https://example.com/help"),
//				SupportEmail: pulumi.String("support@example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// bash The resource can be imported using the ID format `<[org_id]>`, e.g.
//
// ```sh
//
//	$ pulumi import zitadel:index/privacyPolicy:PrivacyPolicy imported '123456789012345678'
//
// ```
type PrivacyPolicy struct {
	pulumi.CustomResourceState

	HelpLink pulumi.StringPtrOutput `pulumi:"helpLink"`
	// ID of the organization
	OrgId        pulumi.StringPtrOutput `pulumi:"orgId"`
	PrivacyLink  pulumi.StringPtrOutput `pulumi:"privacyLink"`
	SupportEmail pulumi.StringPtrOutput `pulumi:"supportEmail"`
	TosLink      pulumi.StringPtrOutput `pulumi:"tosLink"`
}

// NewPrivacyPolicy registers a new resource with the given unique name, arguments, and options.
func NewPrivacyPolicy(ctx *pulumi.Context,
	name string, args *PrivacyPolicyArgs, opts ...pulumi.ResourceOption) (*PrivacyPolicy, error) {
	if args == nil {
		args = &PrivacyPolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PrivacyPolicy
	err := ctx.RegisterResource("zitadel:index/privacyPolicy:PrivacyPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivacyPolicy gets an existing PrivacyPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivacyPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivacyPolicyState, opts ...pulumi.ResourceOption) (*PrivacyPolicy, error) {
	var resource PrivacyPolicy
	err := ctx.ReadResource("zitadel:index/privacyPolicy:PrivacyPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivacyPolicy resources.
type privacyPolicyState struct {
	HelpLink *string `pulumi:"helpLink"`
	// ID of the organization
	OrgId        *string `pulumi:"orgId"`
	PrivacyLink  *string `pulumi:"privacyLink"`
	SupportEmail *string `pulumi:"supportEmail"`
	TosLink      *string `pulumi:"tosLink"`
}

type PrivacyPolicyState struct {
	HelpLink pulumi.StringPtrInput
	// ID of the organization
	OrgId        pulumi.StringPtrInput
	PrivacyLink  pulumi.StringPtrInput
	SupportEmail pulumi.StringPtrInput
	TosLink      pulumi.StringPtrInput
}

func (PrivacyPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*privacyPolicyState)(nil)).Elem()
}

type privacyPolicyArgs struct {
	HelpLink *string `pulumi:"helpLink"`
	// ID of the organization
	OrgId        *string `pulumi:"orgId"`
	PrivacyLink  *string `pulumi:"privacyLink"`
	SupportEmail *string `pulumi:"supportEmail"`
	TosLink      *string `pulumi:"tosLink"`
}

// The set of arguments for constructing a PrivacyPolicy resource.
type PrivacyPolicyArgs struct {
	HelpLink pulumi.StringPtrInput
	// ID of the organization
	OrgId        pulumi.StringPtrInput
	PrivacyLink  pulumi.StringPtrInput
	SupportEmail pulumi.StringPtrInput
	TosLink      pulumi.StringPtrInput
}

func (PrivacyPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privacyPolicyArgs)(nil)).Elem()
}

type PrivacyPolicyInput interface {
	pulumi.Input

	ToPrivacyPolicyOutput() PrivacyPolicyOutput
	ToPrivacyPolicyOutputWithContext(ctx context.Context) PrivacyPolicyOutput
}

func (*PrivacyPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivacyPolicy)(nil)).Elem()
}

func (i *PrivacyPolicy) ToPrivacyPolicyOutput() PrivacyPolicyOutput {
	return i.ToPrivacyPolicyOutputWithContext(context.Background())
}

func (i *PrivacyPolicy) ToPrivacyPolicyOutputWithContext(ctx context.Context) PrivacyPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivacyPolicyOutput)
}

func (i *PrivacyPolicy) ToOutput(ctx context.Context) pulumix.Output[*PrivacyPolicy] {
	return pulumix.Output[*PrivacyPolicy]{
		OutputState: i.ToPrivacyPolicyOutputWithContext(ctx).OutputState,
	}
}

// PrivacyPolicyArrayInput is an input type that accepts PrivacyPolicyArray and PrivacyPolicyArrayOutput values.
// You can construct a concrete instance of `PrivacyPolicyArrayInput` via:
//
//	PrivacyPolicyArray{ PrivacyPolicyArgs{...} }
type PrivacyPolicyArrayInput interface {
	pulumi.Input

	ToPrivacyPolicyArrayOutput() PrivacyPolicyArrayOutput
	ToPrivacyPolicyArrayOutputWithContext(context.Context) PrivacyPolicyArrayOutput
}

type PrivacyPolicyArray []PrivacyPolicyInput

func (PrivacyPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivacyPolicy)(nil)).Elem()
}

func (i PrivacyPolicyArray) ToPrivacyPolicyArrayOutput() PrivacyPolicyArrayOutput {
	return i.ToPrivacyPolicyArrayOutputWithContext(context.Background())
}

func (i PrivacyPolicyArray) ToPrivacyPolicyArrayOutputWithContext(ctx context.Context) PrivacyPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivacyPolicyArrayOutput)
}

func (i PrivacyPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*PrivacyPolicy] {
	return pulumix.Output[[]*PrivacyPolicy]{
		OutputState: i.ToPrivacyPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// PrivacyPolicyMapInput is an input type that accepts PrivacyPolicyMap and PrivacyPolicyMapOutput values.
// You can construct a concrete instance of `PrivacyPolicyMapInput` via:
//
//	PrivacyPolicyMap{ "key": PrivacyPolicyArgs{...} }
type PrivacyPolicyMapInput interface {
	pulumi.Input

	ToPrivacyPolicyMapOutput() PrivacyPolicyMapOutput
	ToPrivacyPolicyMapOutputWithContext(context.Context) PrivacyPolicyMapOutput
}

type PrivacyPolicyMap map[string]PrivacyPolicyInput

func (PrivacyPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivacyPolicy)(nil)).Elem()
}

func (i PrivacyPolicyMap) ToPrivacyPolicyMapOutput() PrivacyPolicyMapOutput {
	return i.ToPrivacyPolicyMapOutputWithContext(context.Background())
}

func (i PrivacyPolicyMap) ToPrivacyPolicyMapOutputWithContext(ctx context.Context) PrivacyPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivacyPolicyMapOutput)
}

func (i PrivacyPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PrivacyPolicy] {
	return pulumix.Output[map[string]*PrivacyPolicy]{
		OutputState: i.ToPrivacyPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type PrivacyPolicyOutput struct{ *pulumi.OutputState }

func (PrivacyPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivacyPolicy)(nil)).Elem()
}

func (o PrivacyPolicyOutput) ToPrivacyPolicyOutput() PrivacyPolicyOutput {
	return o
}

func (o PrivacyPolicyOutput) ToPrivacyPolicyOutputWithContext(ctx context.Context) PrivacyPolicyOutput {
	return o
}

func (o PrivacyPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*PrivacyPolicy] {
	return pulumix.Output[*PrivacyPolicy]{
		OutputState: o.OutputState,
	}
}

func (o PrivacyPolicyOutput) HelpLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivacyPolicy) pulumi.StringPtrOutput { return v.HelpLink }).(pulumi.StringPtrOutput)
}

// ID of the organization
func (o PrivacyPolicyOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivacyPolicy) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

func (o PrivacyPolicyOutput) PrivacyLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivacyPolicy) pulumi.StringPtrOutput { return v.PrivacyLink }).(pulumi.StringPtrOutput)
}

func (o PrivacyPolicyOutput) SupportEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivacyPolicy) pulumi.StringPtrOutput { return v.SupportEmail }).(pulumi.StringPtrOutput)
}

func (o PrivacyPolicyOutput) TosLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivacyPolicy) pulumi.StringPtrOutput { return v.TosLink }).(pulumi.StringPtrOutput)
}

type PrivacyPolicyArrayOutput struct{ *pulumi.OutputState }

func (PrivacyPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivacyPolicy)(nil)).Elem()
}

func (o PrivacyPolicyArrayOutput) ToPrivacyPolicyArrayOutput() PrivacyPolicyArrayOutput {
	return o
}

func (o PrivacyPolicyArrayOutput) ToPrivacyPolicyArrayOutputWithContext(ctx context.Context) PrivacyPolicyArrayOutput {
	return o
}

func (o PrivacyPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PrivacyPolicy] {
	return pulumix.Output[[]*PrivacyPolicy]{
		OutputState: o.OutputState,
	}
}

func (o PrivacyPolicyArrayOutput) Index(i pulumi.IntInput) PrivacyPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrivacyPolicy {
		return vs[0].([]*PrivacyPolicy)[vs[1].(int)]
	}).(PrivacyPolicyOutput)
}

type PrivacyPolicyMapOutput struct{ *pulumi.OutputState }

func (PrivacyPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivacyPolicy)(nil)).Elem()
}

func (o PrivacyPolicyMapOutput) ToPrivacyPolicyMapOutput() PrivacyPolicyMapOutput {
	return o
}

func (o PrivacyPolicyMapOutput) ToPrivacyPolicyMapOutputWithContext(ctx context.Context) PrivacyPolicyMapOutput {
	return o
}

func (o PrivacyPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PrivacyPolicy] {
	return pulumix.Output[map[string]*PrivacyPolicy]{
		OutputState: o.OutputState,
	}
}

func (o PrivacyPolicyMapOutput) MapIndex(k pulumi.StringInput) PrivacyPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrivacyPolicy {
		return vs[0].(map[string]*PrivacyPolicy)[vs[1].(string)]
	}).(PrivacyPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivacyPolicyInput)(nil)).Elem(), &PrivacyPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivacyPolicyArrayInput)(nil)).Elem(), PrivacyPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivacyPolicyMapInput)(nil)).Elem(), PrivacyPolicyMap{})
	pulumi.RegisterOutputType(PrivacyPolicyOutput{})
	pulumi.RegisterOutputType(PrivacyPolicyArrayOutput{})
	pulumi.RegisterOutputType(PrivacyPolicyMapOutput{})
}
