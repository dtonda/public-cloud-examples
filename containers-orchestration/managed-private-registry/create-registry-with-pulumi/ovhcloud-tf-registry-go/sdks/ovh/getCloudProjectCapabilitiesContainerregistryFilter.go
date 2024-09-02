// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCloudProjectCapabilitiesContainerregistryFilter(ctx *pulumi.Context, args *GetCloudProjectCapabilitiesContainerregistryFilterArgs, opts ...pulumi.InvokeOption) (*GetCloudProjectCapabilitiesContainerregistryFilterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetCloudProjectCapabilitiesContainerregistryFilterResult
	err = ctx.InvokePackage("ovh:index/getCloudProjectCapabilitiesContainerregistryFilter:getCloudProjectCapabilitiesContainerregistryFilter", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectCapabilitiesContainerregistryFilter.
type GetCloudProjectCapabilitiesContainerregistryFilterArgs struct {
	Id          *string `pulumi:"id"`
	PlanName    string  `pulumi:"planName"`
	Region      string  `pulumi:"region"`
	ServiceName *string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectCapabilitiesContainerregistryFilter.
type GetCloudProjectCapabilitiesContainerregistryFilterResult struct {
	Code           string                                                            `pulumi:"code"`
	CreatedAt      string                                                            `pulumi:"createdAt"`
	Features       []GetCloudProjectCapabilitiesContainerregistryFilterFeature       `pulumi:"features"`
	Id             string                                                            `pulumi:"id"`
	Name           string                                                            `pulumi:"name"`
	PlanName       string                                                            `pulumi:"planName"`
	Region         string                                                            `pulumi:"region"`
	RegistryLimits []GetCloudProjectCapabilitiesContainerregistryFilterRegistryLimit `pulumi:"registryLimits"`
	ServiceName    *string                                                           `pulumi:"serviceName"`
	UpdatedAt      string                                                            `pulumi:"updatedAt"`
}

func GetCloudProjectCapabilitiesContainerregistryFilterOutput(ctx *pulumi.Context, args GetCloudProjectCapabilitiesContainerregistryFilterOutputArgs, opts ...pulumi.InvokeOption) GetCloudProjectCapabilitiesContainerregistryFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloudProjectCapabilitiesContainerregistryFilterResult, error) {
			args := v.(GetCloudProjectCapabilitiesContainerregistryFilterArgs)
			r, err := GetCloudProjectCapabilitiesContainerregistryFilter(ctx, &args, opts...)
			var s GetCloudProjectCapabilitiesContainerregistryFilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloudProjectCapabilitiesContainerregistryFilterResultOutput)
}

// A collection of arguments for invoking getCloudProjectCapabilitiesContainerregistryFilter.
type GetCloudProjectCapabilitiesContainerregistryFilterOutputArgs struct {
	Id          pulumi.StringPtrInput `pulumi:"id"`
	PlanName    pulumi.StringInput    `pulumi:"planName"`
	Region      pulumi.StringInput    `pulumi:"region"`
	ServiceName pulumi.StringPtrInput `pulumi:"serviceName"`
}

func (GetCloudProjectCapabilitiesContainerregistryFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectCapabilitiesContainerregistryFilterArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectCapabilitiesContainerregistryFilter.
type GetCloudProjectCapabilitiesContainerregistryFilterResultOutput struct{ *pulumi.OutputState }

func (GetCloudProjectCapabilitiesContainerregistryFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectCapabilitiesContainerregistryFilterResult)(nil)).Elem()
}

func (o GetCloudProjectCapabilitiesContainerregistryFilterResultOutput) ToGetCloudProjectCapabilitiesContainerregistryFilterResultOutput() GetCloudProjectCapabilitiesContainerregistryFilterResultOutput {
	return o
}

func (o GetCloudProjectCapabilitiesContainerregistryFilterResultOutput) ToGetCloudProjectCapabilitiesContainerregistryFilterResultOutputWithContext(ctx context.Context) GetCloudProjectCapabilitiesContainerregistryFilterResultOutput {
	return o
}

func (o GetCloudProjectCapabilitiesContainerregistryFilterResultOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectCapabilitiesContainerregistryFilterResult) string { return v.Code }).(pulumi.StringOutput)
}

func (o GetCloudProjectCapabilitiesContainerregistryFilterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectCapabilitiesContainerregistryFilterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o GetCloudProjectCapabilitiesContainerregistryFilterResultOutput) Features() GetCloudProjectCapabilitiesContainerregistryFilterFeatureArrayOutput {
	return o.ApplyT(func(v GetCloudProjectCapabilitiesContainerregistryFilterResult) []GetCloudProjectCapabilitiesContainerregistryFilterFeature {
		return v.Features
	}).(GetCloudProjectCapabilitiesContainerregistryFilterFeatureArrayOutput)
}

func (o GetCloudProjectCapabilitiesContainerregistryFilterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectCapabilitiesContainerregistryFilterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCloudProjectCapabilitiesContainerregistryFilterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectCapabilitiesContainerregistryFilterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetCloudProjectCapabilitiesContainerregistryFilterResultOutput) PlanName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectCapabilitiesContainerregistryFilterResult) string { return v.PlanName }).(pulumi.StringOutput)
}

func (o GetCloudProjectCapabilitiesContainerregistryFilterResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectCapabilitiesContainerregistryFilterResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetCloudProjectCapabilitiesContainerregistryFilterResultOutput) RegistryLimits() GetCloudProjectCapabilitiesContainerregistryFilterRegistryLimitArrayOutput {
	return o.ApplyT(func(v GetCloudProjectCapabilitiesContainerregistryFilterResult) []GetCloudProjectCapabilitiesContainerregistryFilterRegistryLimit {
		return v.RegistryLimits
	}).(GetCloudProjectCapabilitiesContainerregistryFilterRegistryLimitArrayOutput)
}

func (o GetCloudProjectCapabilitiesContainerregistryFilterResultOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCloudProjectCapabilitiesContainerregistryFilterResult) *string { return v.ServiceName }).(pulumi.StringPtrOutput)
}

func (o GetCloudProjectCapabilitiesContainerregistryFilterResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectCapabilitiesContainerregistryFilterResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudProjectCapabilitiesContainerregistryFilterResultOutput{})
}