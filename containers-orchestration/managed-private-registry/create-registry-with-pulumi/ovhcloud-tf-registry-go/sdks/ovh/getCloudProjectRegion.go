// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCloudProjectRegion(ctx *pulumi.Context, args *GetCloudProjectRegionArgs, opts ...pulumi.InvokeOption) (*GetCloudProjectRegionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetCloudProjectRegionResult
	err = ctx.InvokePackage("ovh:index/getCloudProjectRegion:getCloudProjectRegion", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectRegion.
type GetCloudProjectRegionArgs struct {
	Id          *string `pulumi:"id"`
	Name        string  `pulumi:"name"`
	ServiceName *string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectRegion.
type GetCloudProjectRegionResult struct {
	ContinentCode      string                         `pulumi:"continentCode"`
	DatacenterLocation string                         `pulumi:"datacenterLocation"`
	Id                 string                         `pulumi:"id"`
	Name               string                         `pulumi:"name"`
	ServiceName        *string                        `pulumi:"serviceName"`
	Services           []GetCloudProjectRegionService `pulumi:"services"`
}

func GetCloudProjectRegionOutput(ctx *pulumi.Context, args GetCloudProjectRegionOutputArgs, opts ...pulumi.InvokeOption) GetCloudProjectRegionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloudProjectRegionResult, error) {
			args := v.(GetCloudProjectRegionArgs)
			r, err := GetCloudProjectRegion(ctx, &args, opts...)
			var s GetCloudProjectRegionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloudProjectRegionResultOutput)
}

// A collection of arguments for invoking getCloudProjectRegion.
type GetCloudProjectRegionOutputArgs struct {
	Id          pulumi.StringPtrInput `pulumi:"id"`
	Name        pulumi.StringInput    `pulumi:"name"`
	ServiceName pulumi.StringPtrInput `pulumi:"serviceName"`
}

func (GetCloudProjectRegionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectRegionArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectRegion.
type GetCloudProjectRegionResultOutput struct{ *pulumi.OutputState }

func (GetCloudProjectRegionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectRegionResult)(nil)).Elem()
}

func (o GetCloudProjectRegionResultOutput) ToGetCloudProjectRegionResultOutput() GetCloudProjectRegionResultOutput {
	return o
}

func (o GetCloudProjectRegionResultOutput) ToGetCloudProjectRegionResultOutputWithContext(ctx context.Context) GetCloudProjectRegionResultOutput {
	return o
}

func (o GetCloudProjectRegionResultOutput) ContinentCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectRegionResult) string { return v.ContinentCode }).(pulumi.StringOutput)
}

func (o GetCloudProjectRegionResultOutput) DatacenterLocation() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectRegionResult) string { return v.DatacenterLocation }).(pulumi.StringOutput)
}

func (o GetCloudProjectRegionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectRegionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCloudProjectRegionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectRegionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetCloudProjectRegionResultOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCloudProjectRegionResult) *string { return v.ServiceName }).(pulumi.StringPtrOutput)
}

func (o GetCloudProjectRegionResultOutput) Services() GetCloudProjectRegionServiceArrayOutput {
	return o.ApplyT(func(v GetCloudProjectRegionResult) []GetCloudProjectRegionService { return v.Services }).(GetCloudProjectRegionServiceArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudProjectRegionResultOutput{})
}