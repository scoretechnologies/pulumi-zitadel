// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel/internal"
)

// Resource representing the membership of a user on an project, defined with the given role.
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
//			_, err := zitadel.NewProjectMember(ctx, "default", &zitadel.ProjectMemberArgs{
//				OrgId:     pulumi.Any(data.Zitadel_org.Default.Id),
//				ProjectId: pulumi.Any(data.Zitadel_project.Default.Id),
//				UserId:    pulumi.Any(data.Zitadel_human_user.Default.Id),
//				Roles: pulumi.StringArray{
//					pulumi.String("PROJECT_OWNER"),
//				},
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
// bash The resource can be imported using the ID format `<project_id:user_id[:org_id]>`, e.g.
//
// ```sh
//
//	$ pulumi import zitadel:index/projectMember:ProjectMember imported '123456789012345678:123456789012345678:123456789012345678'
//
// ```
type ProjectMember struct {
	pulumi.CustomResourceState

	// ID of the organization
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// ID of the project
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// List of roles granted
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// ID of the user
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewProjectMember registers a new resource with the given unique name, arguments, and options.
func NewProjectMember(ctx *pulumi.Context,
	name string, args *ProjectMemberArgs, opts ...pulumi.ResourceOption) (*ProjectMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectMember
	err := ctx.RegisterResource("zitadel:index/projectMember:ProjectMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectMember gets an existing ProjectMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectMemberState, opts ...pulumi.ResourceOption) (*ProjectMember, error) {
	var resource ProjectMember
	err := ctx.ReadResource("zitadel:index/projectMember:ProjectMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectMember resources.
type projectMemberState struct {
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// ID of the project
	ProjectId *string `pulumi:"projectId"`
	// List of roles granted
	Roles []string `pulumi:"roles"`
	// ID of the user
	UserId *string `pulumi:"userId"`
}

type ProjectMemberState struct {
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// ID of the project
	ProjectId pulumi.StringPtrInput
	// List of roles granted
	Roles pulumi.StringArrayInput
	// ID of the user
	UserId pulumi.StringPtrInput
}

func (ProjectMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMemberState)(nil)).Elem()
}

type projectMemberArgs struct {
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// ID of the project
	ProjectId string `pulumi:"projectId"`
	// List of roles granted
	Roles []string `pulumi:"roles"`
	// ID of the user
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a ProjectMember resource.
type ProjectMemberArgs struct {
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// ID of the project
	ProjectId pulumi.StringInput
	// List of roles granted
	Roles pulumi.StringArrayInput
	// ID of the user
	UserId pulumi.StringInput
}

func (ProjectMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMemberArgs)(nil)).Elem()
}

type ProjectMemberInput interface {
	pulumi.Input

	ToProjectMemberOutput() ProjectMemberOutput
	ToProjectMemberOutputWithContext(ctx context.Context) ProjectMemberOutput
}

func (*ProjectMember) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectMember)(nil)).Elem()
}

func (i *ProjectMember) ToProjectMemberOutput() ProjectMemberOutput {
	return i.ToProjectMemberOutputWithContext(context.Background())
}

func (i *ProjectMember) ToProjectMemberOutputWithContext(ctx context.Context) ProjectMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMemberOutput)
}

func (i *ProjectMember) ToOutput(ctx context.Context) pulumix.Output[*ProjectMember] {
	return pulumix.Output[*ProjectMember]{
		OutputState: i.ToProjectMemberOutputWithContext(ctx).OutputState,
	}
}

// ProjectMemberArrayInput is an input type that accepts ProjectMemberArray and ProjectMemberArrayOutput values.
// You can construct a concrete instance of `ProjectMemberArrayInput` via:
//
//	ProjectMemberArray{ ProjectMemberArgs{...} }
type ProjectMemberArrayInput interface {
	pulumi.Input

	ToProjectMemberArrayOutput() ProjectMemberArrayOutput
	ToProjectMemberArrayOutputWithContext(context.Context) ProjectMemberArrayOutput
}

type ProjectMemberArray []ProjectMemberInput

func (ProjectMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectMember)(nil)).Elem()
}

func (i ProjectMemberArray) ToProjectMemberArrayOutput() ProjectMemberArrayOutput {
	return i.ToProjectMemberArrayOutputWithContext(context.Background())
}

func (i ProjectMemberArray) ToProjectMemberArrayOutputWithContext(ctx context.Context) ProjectMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMemberArrayOutput)
}

func (i ProjectMemberArray) ToOutput(ctx context.Context) pulumix.Output[[]*ProjectMember] {
	return pulumix.Output[[]*ProjectMember]{
		OutputState: i.ToProjectMemberArrayOutputWithContext(ctx).OutputState,
	}
}

// ProjectMemberMapInput is an input type that accepts ProjectMemberMap and ProjectMemberMapOutput values.
// You can construct a concrete instance of `ProjectMemberMapInput` via:
//
//	ProjectMemberMap{ "key": ProjectMemberArgs{...} }
type ProjectMemberMapInput interface {
	pulumi.Input

	ToProjectMemberMapOutput() ProjectMemberMapOutput
	ToProjectMemberMapOutputWithContext(context.Context) ProjectMemberMapOutput
}

type ProjectMemberMap map[string]ProjectMemberInput

func (ProjectMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectMember)(nil)).Elem()
}

func (i ProjectMemberMap) ToProjectMemberMapOutput() ProjectMemberMapOutput {
	return i.ToProjectMemberMapOutputWithContext(context.Background())
}

func (i ProjectMemberMap) ToProjectMemberMapOutputWithContext(ctx context.Context) ProjectMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMemberMapOutput)
}

func (i ProjectMemberMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ProjectMember] {
	return pulumix.Output[map[string]*ProjectMember]{
		OutputState: i.ToProjectMemberMapOutputWithContext(ctx).OutputState,
	}
}

type ProjectMemberOutput struct{ *pulumi.OutputState }

func (ProjectMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectMember)(nil)).Elem()
}

func (o ProjectMemberOutput) ToProjectMemberOutput() ProjectMemberOutput {
	return o
}

func (o ProjectMemberOutput) ToProjectMemberOutputWithContext(ctx context.Context) ProjectMemberOutput {
	return o
}

func (o ProjectMemberOutput) ToOutput(ctx context.Context) pulumix.Output[*ProjectMember] {
	return pulumix.Output[*ProjectMember]{
		OutputState: o.OutputState,
	}
}

// ID of the organization
func (o ProjectMemberOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectMember) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// ID of the project
func (o ProjectMemberOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectMember) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// List of roles granted
func (o ProjectMemberOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProjectMember) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// ID of the user
func (o ProjectMemberOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectMember) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type ProjectMemberArrayOutput struct{ *pulumi.OutputState }

func (ProjectMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectMember)(nil)).Elem()
}

func (o ProjectMemberArrayOutput) ToProjectMemberArrayOutput() ProjectMemberArrayOutput {
	return o
}

func (o ProjectMemberArrayOutput) ToProjectMemberArrayOutputWithContext(ctx context.Context) ProjectMemberArrayOutput {
	return o
}

func (o ProjectMemberArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ProjectMember] {
	return pulumix.Output[[]*ProjectMember]{
		OutputState: o.OutputState,
	}
}

func (o ProjectMemberArrayOutput) Index(i pulumi.IntInput) ProjectMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectMember {
		return vs[0].([]*ProjectMember)[vs[1].(int)]
	}).(ProjectMemberOutput)
}

type ProjectMemberMapOutput struct{ *pulumi.OutputState }

func (ProjectMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectMember)(nil)).Elem()
}

func (o ProjectMemberMapOutput) ToProjectMemberMapOutput() ProjectMemberMapOutput {
	return o
}

func (o ProjectMemberMapOutput) ToProjectMemberMapOutputWithContext(ctx context.Context) ProjectMemberMapOutput {
	return o
}

func (o ProjectMemberMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ProjectMember] {
	return pulumix.Output[map[string]*ProjectMember]{
		OutputState: o.OutputState,
	}
}

func (o ProjectMemberMapOutput) MapIndex(k pulumi.StringInput) ProjectMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectMember {
		return vs[0].(map[string]*ProjectMember)[vs[1].(string)]
	}).(ProjectMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMemberInput)(nil)).Elem(), &ProjectMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMemberArrayInput)(nil)).Elem(), ProjectMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMemberMapInput)(nil)).Elem(), ProjectMemberMap{})
	pulumi.RegisterOutputType(ProjectMemberOutput{})
	pulumi.RegisterOutputType(ProjectMemberArrayOutput{})
	pulumi.RegisterOutputType(ProjectMemberMapOutput{})
}
