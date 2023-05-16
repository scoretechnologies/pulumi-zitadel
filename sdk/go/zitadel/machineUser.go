// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing a serviceaccount situated under an organization, which then can be authorized through memberships or direct grants on other resources.
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
// 		_, err := zitadel.NewMachineUser(ctx, "machineUser", &zitadel.MachineUserArgs{
// 			OrgId:       pulumi.Any(zitadel_org.Org.Id),
// 			UserName:    pulumi.String("machine@localhost.com"),
// 			Description: pulumi.String("description"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type MachineUser struct {
	pulumi.CustomResourceState

	// Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
	AccessTokenType pulumi.StringPtrOutput `pulumi:"accessTokenType"`
	// Description of the user
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Loginnames
	LoginNames pulumi.StringArrayOutput `pulumi:"loginNames"`
	// Name of the machine user
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the organization
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Preferred login name
	PreferredLoginName pulumi.StringOutput `pulumi:"preferredLoginName"`
	// State of the user
	State pulumi.StringOutput `pulumi:"state"`
	// Username
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewMachineUser registers a new resource with the given unique name, arguments, and options.
func NewMachineUser(ctx *pulumi.Context,
	name string, args *MachineUserArgs, opts ...pulumi.ResourceOption) (*MachineUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	var resource MachineUser
	err := ctx.RegisterResource("zitadel:index/machineUser:MachineUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineUser gets an existing MachineUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineUserState, opts ...pulumi.ResourceOption) (*MachineUser, error) {
	var resource MachineUser
	err := ctx.ReadResource("zitadel:index/machineUser:MachineUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineUser resources.
type machineUserState struct {
	// Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
	AccessTokenType *string `pulumi:"accessTokenType"`
	// Description of the user
	Description *string `pulumi:"description"`
	// Loginnames
	LoginNames []string `pulumi:"loginNames"`
	// Name of the machine user
	Name *string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// Preferred login name
	PreferredLoginName *string `pulumi:"preferredLoginName"`
	// State of the user
	State *string `pulumi:"state"`
	// Username
	UserName *string `pulumi:"userName"`
}

type MachineUserState struct {
	// Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
	AccessTokenType pulumi.StringPtrInput
	// Description of the user
	Description pulumi.StringPtrInput
	// Loginnames
	LoginNames pulumi.StringArrayInput
	// Name of the machine user
	Name pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// Preferred login name
	PreferredLoginName pulumi.StringPtrInput
	// State of the user
	State pulumi.StringPtrInput
	// Username
	UserName pulumi.StringPtrInput
}

func (MachineUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineUserState)(nil)).Elem()
}

type machineUserArgs struct {
	// Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
	AccessTokenType *string `pulumi:"accessTokenType"`
	// Description of the user
	Description *string `pulumi:"description"`
	// Name of the machine user
	Name *string `pulumi:"name"`
	// ID of the organization
	OrgId string `pulumi:"orgId"`
	// Username
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a MachineUser resource.
type MachineUserArgs struct {
	// Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
	AccessTokenType pulumi.StringPtrInput
	// Description of the user
	Description pulumi.StringPtrInput
	// Name of the machine user
	Name pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringInput
	// Username
	UserName pulumi.StringInput
}

func (MachineUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineUserArgs)(nil)).Elem()
}

type MachineUserInput interface {
	pulumi.Input

	ToMachineUserOutput() MachineUserOutput
	ToMachineUserOutputWithContext(ctx context.Context) MachineUserOutput
}

func (*MachineUser) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineUser)(nil)).Elem()
}

func (i *MachineUser) ToMachineUserOutput() MachineUserOutput {
	return i.ToMachineUserOutputWithContext(context.Background())
}

func (i *MachineUser) ToMachineUserOutputWithContext(ctx context.Context) MachineUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineUserOutput)
}

// MachineUserArrayInput is an input type that accepts MachineUserArray and MachineUserArrayOutput values.
// You can construct a concrete instance of `MachineUserArrayInput` via:
//
//          MachineUserArray{ MachineUserArgs{...} }
type MachineUserArrayInput interface {
	pulumi.Input

	ToMachineUserArrayOutput() MachineUserArrayOutput
	ToMachineUserArrayOutputWithContext(context.Context) MachineUserArrayOutput
}

type MachineUserArray []MachineUserInput

func (MachineUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MachineUser)(nil)).Elem()
}

func (i MachineUserArray) ToMachineUserArrayOutput() MachineUserArrayOutput {
	return i.ToMachineUserArrayOutputWithContext(context.Background())
}

func (i MachineUserArray) ToMachineUserArrayOutputWithContext(ctx context.Context) MachineUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineUserArrayOutput)
}

// MachineUserMapInput is an input type that accepts MachineUserMap and MachineUserMapOutput values.
// You can construct a concrete instance of `MachineUserMapInput` via:
//
//          MachineUserMap{ "key": MachineUserArgs{...} }
type MachineUserMapInput interface {
	pulumi.Input

	ToMachineUserMapOutput() MachineUserMapOutput
	ToMachineUserMapOutputWithContext(context.Context) MachineUserMapOutput
}

type MachineUserMap map[string]MachineUserInput

func (MachineUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MachineUser)(nil)).Elem()
}

func (i MachineUserMap) ToMachineUserMapOutput() MachineUserMapOutput {
	return i.ToMachineUserMapOutputWithContext(context.Background())
}

func (i MachineUserMap) ToMachineUserMapOutputWithContext(ctx context.Context) MachineUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineUserMapOutput)
}

type MachineUserOutput struct{ *pulumi.OutputState }

func (MachineUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineUser)(nil)).Elem()
}

func (o MachineUserOutput) ToMachineUserOutput() MachineUserOutput {
	return o
}

func (o MachineUserOutput) ToMachineUserOutputWithContext(ctx context.Context) MachineUserOutput {
	return o
}

// Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
func (o MachineUserOutput) AccessTokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineUser) pulumi.StringPtrOutput { return v.AccessTokenType }).(pulumi.StringPtrOutput)
}

// Description of the user
func (o MachineUserOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineUser) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Loginnames
func (o MachineUserOutput) LoginNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MachineUser) pulumi.StringArrayOutput { return v.LoginNames }).(pulumi.StringArrayOutput)
}

// Name of the machine user
func (o MachineUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o MachineUserOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineUser) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Preferred login name
func (o MachineUserOutput) PreferredLoginName() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineUser) pulumi.StringOutput { return v.PreferredLoginName }).(pulumi.StringOutput)
}

// State of the user
func (o MachineUserOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineUser) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Username
func (o MachineUserOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineUser) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

type MachineUserArrayOutput struct{ *pulumi.OutputState }

func (MachineUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MachineUser)(nil)).Elem()
}

func (o MachineUserArrayOutput) ToMachineUserArrayOutput() MachineUserArrayOutput {
	return o
}

func (o MachineUserArrayOutput) ToMachineUserArrayOutputWithContext(ctx context.Context) MachineUserArrayOutput {
	return o
}

func (o MachineUserArrayOutput) Index(i pulumi.IntInput) MachineUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MachineUser {
		return vs[0].([]*MachineUser)[vs[1].(int)]
	}).(MachineUserOutput)
}

type MachineUserMapOutput struct{ *pulumi.OutputState }

func (MachineUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MachineUser)(nil)).Elem()
}

func (o MachineUserMapOutput) ToMachineUserMapOutput() MachineUserMapOutput {
	return o
}

func (o MachineUserMapOutput) ToMachineUserMapOutputWithContext(ctx context.Context) MachineUserMapOutput {
	return o
}

func (o MachineUserMapOutput) MapIndex(k pulumi.StringInput) MachineUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MachineUser {
		return vs[0].(map[string]*MachineUser)[vs[1].(string)]
	}).(MachineUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MachineUserInput)(nil)).Elem(), &MachineUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineUserArrayInput)(nil)).Elem(), MachineUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineUserMapInput)(nil)).Elem(), MachineUserMap{})
	pulumi.RegisterOutputType(MachineUserOutput{})
	pulumi.RegisterOutputType(MachineUserArrayOutput{})
	pulumi.RegisterOutputType(MachineUserMapOutput{})
}