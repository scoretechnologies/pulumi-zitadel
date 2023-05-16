// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing an action belonging to an organization.
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
// 		_, err := zitadel.NewAction(ctx, "action", &zitadel.ActionArgs{
// 			OrgId:         pulumi.Any(zitadel_org.Org.Id),
// 			Script:        pulumi.String("testscript"),
// 			Timeout:       pulumi.String("10s"),
// 			AllowedToFail: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Action struct {
	pulumi.CustomResourceState

	// when true, the next action will be called even if this action fails
	AllowedToFail pulumi.BoolOutput   `pulumi:"allowedToFail"`
	Name          pulumi.StringOutput `pulumi:"name"`
	// ID of the organization
	OrgId  pulumi.StringOutput `pulumi:"orgId"`
	Script pulumi.StringOutput `pulumi:"script"`
	// the state of the action
	State pulumi.IntOutput `pulumi:"state"`
	// after which time the action will be terminated if not finished
	Timeout pulumi.StringOutput `pulumi:"timeout"`
}

// NewAction registers a new resource with the given unique name, arguments, and options.
func NewAction(ctx *pulumi.Context,
	name string, args *ActionArgs, opts ...pulumi.ResourceOption) (*Action, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowedToFail == nil {
		return nil, errors.New("invalid value for required argument 'AllowedToFail'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.Script == nil {
		return nil, errors.New("invalid value for required argument 'Script'")
	}
	if args.Timeout == nil {
		return nil, errors.New("invalid value for required argument 'Timeout'")
	}
	var resource Action
	err := ctx.RegisterResource("zitadel:index/action:Action", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAction gets an existing Action resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionState, opts ...pulumi.ResourceOption) (*Action, error) {
	var resource Action
	err := ctx.ReadResource("zitadel:index/action:Action", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Action resources.
type actionState struct {
	// when true, the next action will be called even if this action fails
	AllowedToFail *bool   `pulumi:"allowedToFail"`
	Name          *string `pulumi:"name"`
	// ID of the organization
	OrgId  *string `pulumi:"orgId"`
	Script *string `pulumi:"script"`
	// the state of the action
	State *int `pulumi:"state"`
	// after which time the action will be terminated if not finished
	Timeout *string `pulumi:"timeout"`
}

type ActionState struct {
	// when true, the next action will be called even if this action fails
	AllowedToFail pulumi.BoolPtrInput
	Name          pulumi.StringPtrInput
	// ID of the organization
	OrgId  pulumi.StringPtrInput
	Script pulumi.StringPtrInput
	// the state of the action
	State pulumi.IntPtrInput
	// after which time the action will be terminated if not finished
	Timeout pulumi.StringPtrInput
}

func (ActionState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionState)(nil)).Elem()
}

type actionArgs struct {
	// when true, the next action will be called even if this action fails
	AllowedToFail bool    `pulumi:"allowedToFail"`
	Name          *string `pulumi:"name"`
	// ID of the organization
	OrgId  string `pulumi:"orgId"`
	Script string `pulumi:"script"`
	// after which time the action will be terminated if not finished
	Timeout string `pulumi:"timeout"`
}

// The set of arguments for constructing a Action resource.
type ActionArgs struct {
	// when true, the next action will be called even if this action fails
	AllowedToFail pulumi.BoolInput
	Name          pulumi.StringPtrInput
	// ID of the organization
	OrgId  pulumi.StringInput
	Script pulumi.StringInput
	// after which time the action will be terminated if not finished
	Timeout pulumi.StringInput
}

func (ActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionArgs)(nil)).Elem()
}

type ActionInput interface {
	pulumi.Input

	ToActionOutput() ActionOutput
	ToActionOutputWithContext(ctx context.Context) ActionOutput
}

func (*Action) ElementType() reflect.Type {
	return reflect.TypeOf((**Action)(nil)).Elem()
}

func (i *Action) ToActionOutput() ActionOutput {
	return i.ToActionOutputWithContext(context.Background())
}

func (i *Action) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionOutput)
}

// ActionArrayInput is an input type that accepts ActionArray and ActionArrayOutput values.
// You can construct a concrete instance of `ActionArrayInput` via:
//
//          ActionArray{ ActionArgs{...} }
type ActionArrayInput interface {
	pulumi.Input

	ToActionArrayOutput() ActionArrayOutput
	ToActionArrayOutputWithContext(context.Context) ActionArrayOutput
}

type ActionArray []ActionInput

func (ActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Action)(nil)).Elem()
}

func (i ActionArray) ToActionArrayOutput() ActionArrayOutput {
	return i.ToActionArrayOutputWithContext(context.Background())
}

func (i ActionArray) ToActionArrayOutputWithContext(ctx context.Context) ActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionArrayOutput)
}

// ActionMapInput is an input type that accepts ActionMap and ActionMapOutput values.
// You can construct a concrete instance of `ActionMapInput` via:
//
//          ActionMap{ "key": ActionArgs{...} }
type ActionMapInput interface {
	pulumi.Input

	ToActionMapOutput() ActionMapOutput
	ToActionMapOutputWithContext(context.Context) ActionMapOutput
}

type ActionMap map[string]ActionInput

func (ActionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Action)(nil)).Elem()
}

func (i ActionMap) ToActionMapOutput() ActionMapOutput {
	return i.ToActionMapOutputWithContext(context.Background())
}

func (i ActionMap) ToActionMapOutputWithContext(ctx context.Context) ActionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionMapOutput)
}

type ActionOutput struct{ *pulumi.OutputState }

func (ActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Action)(nil)).Elem()
}

func (o ActionOutput) ToActionOutput() ActionOutput {
	return o
}

func (o ActionOutput) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return o
}

// when true, the next action will be called even if this action fails
func (o ActionOutput) AllowedToFail() pulumi.BoolOutput {
	return o.ApplyT(func(v *Action) pulumi.BoolOutput { return v.AllowedToFail }).(pulumi.BoolOutput)
}

func (o ActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Action) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o ActionOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *Action) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

func (o ActionOutput) Script() pulumi.StringOutput {
	return o.ApplyT(func(v *Action) pulumi.StringOutput { return v.Script }).(pulumi.StringOutput)
}

// the state of the action
func (o ActionOutput) State() pulumi.IntOutput {
	return o.ApplyT(func(v *Action) pulumi.IntOutput { return v.State }).(pulumi.IntOutput)
}

// after which time the action will be terminated if not finished
func (o ActionOutput) Timeout() pulumi.StringOutput {
	return o.ApplyT(func(v *Action) pulumi.StringOutput { return v.Timeout }).(pulumi.StringOutput)
}

type ActionArrayOutput struct{ *pulumi.OutputState }

func (ActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Action)(nil)).Elem()
}

func (o ActionArrayOutput) ToActionArrayOutput() ActionArrayOutput {
	return o
}

func (o ActionArrayOutput) ToActionArrayOutputWithContext(ctx context.Context) ActionArrayOutput {
	return o
}

func (o ActionArrayOutput) Index(i pulumi.IntInput) ActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Action {
		return vs[0].([]*Action)[vs[1].(int)]
	}).(ActionOutput)
}

type ActionMapOutput struct{ *pulumi.OutputState }

func (ActionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Action)(nil)).Elem()
}

func (o ActionMapOutput) ToActionMapOutput() ActionMapOutput {
	return o
}

func (o ActionMapOutput) ToActionMapOutputWithContext(ctx context.Context) ActionMapOutput {
	return o
}

func (o ActionMapOutput) MapIndex(k pulumi.StringInput) ActionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Action {
		return vs[0].(map[string]*Action)[vs[1].(string)]
	}).(ActionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActionInput)(nil)).Elem(), &Action{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionArrayInput)(nil)).Elem(), ActionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionMapInput)(nil)).Elem(), ActionMap{})
	pulumi.RegisterOutputType(ActionOutput{})
	pulumi.RegisterOutputType(ActionArrayOutput{})
	pulumi.RegisterOutputType(ActionMapOutput{})
}
