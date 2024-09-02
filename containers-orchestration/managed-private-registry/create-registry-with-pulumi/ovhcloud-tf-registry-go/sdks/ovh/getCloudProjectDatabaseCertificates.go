// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCloudProjectDatabaseCertificates(ctx *pulumi.Context, args *GetCloudProjectDatabaseCertificatesArgs, opts ...pulumi.InvokeOption) (*GetCloudProjectDatabaseCertificatesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetCloudProjectDatabaseCertificatesResult
	err = ctx.InvokePackage("ovh:index/getCloudProjectDatabaseCertificates:getCloudProjectDatabaseCertificates", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectDatabaseCertificates.
type GetCloudProjectDatabaseCertificatesArgs struct {
	ClusterId   string  `pulumi:"clusterId"`
	Engine      string  `pulumi:"engine"`
	Id          *string `pulumi:"id"`
	ServiceName *string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectDatabaseCertificates.
type GetCloudProjectDatabaseCertificatesResult struct {
	Ca          string  `pulumi:"ca"`
	ClusterId   string  `pulumi:"clusterId"`
	Engine      string  `pulumi:"engine"`
	Id          string  `pulumi:"id"`
	ServiceName *string `pulumi:"serviceName"`
}

func GetCloudProjectDatabaseCertificatesOutput(ctx *pulumi.Context, args GetCloudProjectDatabaseCertificatesOutputArgs, opts ...pulumi.InvokeOption) GetCloudProjectDatabaseCertificatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloudProjectDatabaseCertificatesResult, error) {
			args := v.(GetCloudProjectDatabaseCertificatesArgs)
			r, err := GetCloudProjectDatabaseCertificates(ctx, &args, opts...)
			var s GetCloudProjectDatabaseCertificatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloudProjectDatabaseCertificatesResultOutput)
}

// A collection of arguments for invoking getCloudProjectDatabaseCertificates.
type GetCloudProjectDatabaseCertificatesOutputArgs struct {
	ClusterId   pulumi.StringInput    `pulumi:"clusterId"`
	Engine      pulumi.StringInput    `pulumi:"engine"`
	Id          pulumi.StringPtrInput `pulumi:"id"`
	ServiceName pulumi.StringPtrInput `pulumi:"serviceName"`
}

func (GetCloudProjectDatabaseCertificatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectDatabaseCertificatesArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectDatabaseCertificates.
type GetCloudProjectDatabaseCertificatesResultOutput struct{ *pulumi.OutputState }

func (GetCloudProjectDatabaseCertificatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectDatabaseCertificatesResult)(nil)).Elem()
}

func (o GetCloudProjectDatabaseCertificatesResultOutput) ToGetCloudProjectDatabaseCertificatesResultOutput() GetCloudProjectDatabaseCertificatesResultOutput {
	return o
}

func (o GetCloudProjectDatabaseCertificatesResultOutput) ToGetCloudProjectDatabaseCertificatesResultOutputWithContext(ctx context.Context) GetCloudProjectDatabaseCertificatesResultOutput {
	return o
}

func (o GetCloudProjectDatabaseCertificatesResultOutput) Ca() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectDatabaseCertificatesResult) string { return v.Ca }).(pulumi.StringOutput)
}

func (o GetCloudProjectDatabaseCertificatesResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectDatabaseCertificatesResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o GetCloudProjectDatabaseCertificatesResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectDatabaseCertificatesResult) string { return v.Engine }).(pulumi.StringOutput)
}

func (o GetCloudProjectDatabaseCertificatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectDatabaseCertificatesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCloudProjectDatabaseCertificatesResultOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCloudProjectDatabaseCertificatesResult) *string { return v.ServiceName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudProjectDatabaseCertificatesResultOutput{})
}
