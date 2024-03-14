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

// Datasource representing the project, which can then be granted to different organizations or users directly, containing different applications.
func GetProjects(ctx *pulumi.Context, args *GetProjectsArgs, opts ...pulumi.InvokeOption) (*GetProjectsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectsResult
	err := ctx.Invoke("zitadel:index/getProjects:getProjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjects.
type GetProjectsArgs struct {
	// Name of the project
	Name string `pulumi:"name"`
	// Method for querying projects by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
	NameMethod *string `pulumi:"nameMethod"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
}

// A collection of values returned by getProjects.
type GetProjectsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the project
	Name string `pulumi:"name"`
	// Method for querying projects by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
	NameMethod *string `pulumi:"nameMethod"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// A set of all project IDs.
	ProjectIds []string `pulumi:"projectIds"`
}

func GetProjectsOutput(ctx *pulumi.Context, args GetProjectsOutputArgs, opts ...pulumi.InvokeOption) GetProjectsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectsResult, error) {
			args := v.(GetProjectsArgs)
			r, err := GetProjects(ctx, &args, opts...)
			var s GetProjectsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectsResultOutput)
}

// A collection of arguments for invoking getProjects.
type GetProjectsOutputArgs struct {
	// Name of the project
	Name pulumi.StringInput `pulumi:"name"`
	// Method for querying projects by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
	NameMethod pulumi.StringPtrInput `pulumi:"nameMethod"`
	// ID of the organization
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
}

func (GetProjectsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsArgs)(nil)).Elem()
}

// A collection of values returned by getProjects.
type GetProjectsResultOutput struct{ *pulumi.OutputState }

func (GetProjectsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsResult)(nil)).Elem()
}

func (o GetProjectsResultOutput) ToGetProjectsResultOutput() GetProjectsResultOutput {
	return o
}

func (o GetProjectsResultOutput) ToGetProjectsResultOutputWithContext(ctx context.Context) GetProjectsResultOutput {
	return o
}

func (o GetProjectsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetProjectsResult] {
	return pulumix.Output[GetProjectsResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the project
func (o GetProjectsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Method for querying projects by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
func (o GetProjectsResultOutput) NameMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.NameMethod }).(pulumi.StringPtrOutput)
}

// ID of the organization
func (o GetProjectsResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// A set of all project IDs.
func (o GetProjectsResultOutput) ProjectIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProjectsResult) []string { return v.ProjectIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectsResultOutput{})
}