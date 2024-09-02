// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetIamPolicies(ctx *pulumi.Context, args *GetIamPoliciesArgs, opts ...pulumi.InvokeOption) (*GetIamPoliciesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetIamPoliciesResult
	err = ctx.InvokePackage("ovh:index/getIamPolicies:getIamPolicies", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIamPolicies.
type GetIamPoliciesArgs struct {
	Id *string `pulumi:"id"`
}

// A collection of values returned by getIamPolicies.
type GetIamPoliciesResult struct {
	Id       string   `pulumi:"id"`
	Policies []string `pulumi:"policies"`
}

func GetIamPoliciesOutput(ctx *pulumi.Context, args GetIamPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetIamPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIamPoliciesResult, error) {
			args := v.(GetIamPoliciesArgs)
			r, err := GetIamPolicies(ctx, &args, opts...)
			var s GetIamPoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIamPoliciesResultOutput)
}

// A collection of arguments for invoking getIamPolicies.
type GetIamPoliciesOutputArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (GetIamPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIamPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getIamPolicies.
type GetIamPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetIamPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIamPoliciesResult)(nil)).Elem()
}

func (o GetIamPoliciesResultOutput) ToGetIamPoliciesResultOutput() GetIamPoliciesResultOutput {
	return o
}

func (o GetIamPoliciesResultOutput) ToGetIamPoliciesResultOutputWithContext(ctx context.Context) GetIamPoliciesResultOutput {
	return o
}

func (o GetIamPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIamPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIamPoliciesResultOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIamPoliciesResult) []string { return v.Policies }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIamPoliciesResultOutput{})
}