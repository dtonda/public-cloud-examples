// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCloudProjectLoadbalancer(ctx *pulumi.Context, args *GetCloudProjectLoadbalancerArgs, opts ...pulumi.InvokeOption) (*GetCloudProjectLoadbalancerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetCloudProjectLoadbalancerResult
	err = ctx.InvokePackage("ovh:index/getCloudProjectLoadbalancer:getCloudProjectLoadbalancer", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectLoadbalancer.
type GetCloudProjectLoadbalancerArgs struct {
	Id          string `pulumi:"id"`
	RegionName  string `pulumi:"regionName"`
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectLoadbalancer.
type GetCloudProjectLoadbalancerResult struct {
	CreatedAt          string                                `pulumi:"createdAt"`
	FlavorId           string                                `pulumi:"flavorId"`
	FloatingIp         GetCloudProjectLoadbalancerFloatingIp `pulumi:"floatingIp"`
	Id                 string                                `pulumi:"id"`
	Name               string                                `pulumi:"name"`
	OperatingStatus    string                                `pulumi:"operatingStatus"`
	ProvisioningStatus string                                `pulumi:"provisioningStatus"`
	RegionName         string                                `pulumi:"regionName"`
	ServiceName        string                                `pulumi:"serviceName"`
	UpdatedAt          string                                `pulumi:"updatedAt"`
	VipAddress         string                                `pulumi:"vipAddress"`
	VipNetworkId       string                                `pulumi:"vipNetworkId"`
	VipSubnetId        string                                `pulumi:"vipSubnetId"`
}

func GetCloudProjectLoadbalancerOutput(ctx *pulumi.Context, args GetCloudProjectLoadbalancerOutputArgs, opts ...pulumi.InvokeOption) GetCloudProjectLoadbalancerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloudProjectLoadbalancerResult, error) {
			args := v.(GetCloudProjectLoadbalancerArgs)
			r, err := GetCloudProjectLoadbalancer(ctx, &args, opts...)
			var s GetCloudProjectLoadbalancerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloudProjectLoadbalancerResultOutput)
}

// A collection of arguments for invoking getCloudProjectLoadbalancer.
type GetCloudProjectLoadbalancerOutputArgs struct {
	Id          pulumi.StringInput `pulumi:"id"`
	RegionName  pulumi.StringInput `pulumi:"regionName"`
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetCloudProjectLoadbalancerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectLoadbalancerArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectLoadbalancer.
type GetCloudProjectLoadbalancerResultOutput struct{ *pulumi.OutputState }

func (GetCloudProjectLoadbalancerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectLoadbalancerResult)(nil)).Elem()
}

func (o GetCloudProjectLoadbalancerResultOutput) ToGetCloudProjectLoadbalancerResultOutput() GetCloudProjectLoadbalancerResultOutput {
	return o
}

func (o GetCloudProjectLoadbalancerResultOutput) ToGetCloudProjectLoadbalancerResultOutputWithContext(ctx context.Context) GetCloudProjectLoadbalancerResultOutput {
	return o
}

func (o GetCloudProjectLoadbalancerResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectLoadbalancerResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o GetCloudProjectLoadbalancerResultOutput) FlavorId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectLoadbalancerResult) string { return v.FlavorId }).(pulumi.StringOutput)
}

func (o GetCloudProjectLoadbalancerResultOutput) FloatingIp() GetCloudProjectLoadbalancerFloatingIpOutput {
	return o.ApplyT(func(v GetCloudProjectLoadbalancerResult) GetCloudProjectLoadbalancerFloatingIp { return v.FloatingIp }).(GetCloudProjectLoadbalancerFloatingIpOutput)
}

func (o GetCloudProjectLoadbalancerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectLoadbalancerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCloudProjectLoadbalancerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectLoadbalancerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetCloudProjectLoadbalancerResultOutput) OperatingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectLoadbalancerResult) string { return v.OperatingStatus }).(pulumi.StringOutput)
}

func (o GetCloudProjectLoadbalancerResultOutput) ProvisioningStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectLoadbalancerResult) string { return v.ProvisioningStatus }).(pulumi.StringOutput)
}

func (o GetCloudProjectLoadbalancerResultOutput) RegionName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectLoadbalancerResult) string { return v.RegionName }).(pulumi.StringOutput)
}

func (o GetCloudProjectLoadbalancerResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectLoadbalancerResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o GetCloudProjectLoadbalancerResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectLoadbalancerResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o GetCloudProjectLoadbalancerResultOutput) VipAddress() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectLoadbalancerResult) string { return v.VipAddress }).(pulumi.StringOutput)
}

func (o GetCloudProjectLoadbalancerResultOutput) VipNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectLoadbalancerResult) string { return v.VipNetworkId }).(pulumi.StringOutput)
}

func (o GetCloudProjectLoadbalancerResultOutput) VipSubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectLoadbalancerResult) string { return v.VipSubnetId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudProjectLoadbalancerResultOutput{})
}
