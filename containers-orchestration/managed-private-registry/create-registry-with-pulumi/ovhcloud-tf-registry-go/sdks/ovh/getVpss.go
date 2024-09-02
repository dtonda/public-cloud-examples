// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVpss(ctx *pulumi.Context, args *GetVpssArgs, opts ...pulumi.InvokeOption) (*GetVpssResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetVpssResult
	err = ctx.InvokePackage("ovh:index/getVpss:getVpss", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpss.
type GetVpssArgs struct {
	Id *string `pulumi:"id"`
}

// A collection of values returned by getVpss.
type GetVpssResult struct {
	Id      string   `pulumi:"id"`
	Results []string `pulumi:"results"`
}

func GetVpssOutput(ctx *pulumi.Context, args GetVpssOutputArgs, opts ...pulumi.InvokeOption) GetVpssResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVpssResult, error) {
			args := v.(GetVpssArgs)
			r, err := GetVpss(ctx, &args, opts...)
			var s GetVpssResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVpssResultOutput)
}

// A collection of arguments for invoking getVpss.
type GetVpssOutputArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (GetVpssOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpssArgs)(nil)).Elem()
}

// A collection of values returned by getVpss.
type GetVpssResultOutput struct{ *pulumi.OutputState }

func (GetVpssResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpssResult)(nil)).Elem()
}

func (o GetVpssResultOutput) ToGetVpssResultOutput() GetVpssResultOutput {
	return o
}

func (o GetVpssResultOutput) ToGetVpssResultOutputWithContext(ctx context.Context) GetVpssResultOutput {
	return o
}

func (o GetVpssResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpssResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVpssResultOutput) Results() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVpssResult) []string { return v.Results }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpssResultOutput{})
}