// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCloudProjectDatabaseKafkaUserAccess(ctx *pulumi.Context, args *GetCloudProjectDatabaseKafkaUserAccessArgs, opts ...pulumi.InvokeOption) (*GetCloudProjectDatabaseKafkaUserAccessResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetCloudProjectDatabaseKafkaUserAccessResult
	err = ctx.InvokePackage("ovh:index/getCloudProjectDatabaseKafkaUserAccess:getCloudProjectDatabaseKafkaUserAccess", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectDatabaseKafkaUserAccess.
type GetCloudProjectDatabaseKafkaUserAccessArgs struct {
	ClusterId   string  `pulumi:"clusterId"`
	Id          *string `pulumi:"id"`
	ServiceName *string `pulumi:"serviceName"`
	UserId      string  `pulumi:"userId"`
}

// A collection of values returned by getCloudProjectDatabaseKafkaUserAccess.
type GetCloudProjectDatabaseKafkaUserAccessResult struct {
	Cert        string  `pulumi:"cert"`
	ClusterId   string  `pulumi:"clusterId"`
	Id          string  `pulumi:"id"`
	Key         string  `pulumi:"key"`
	ServiceName *string `pulumi:"serviceName"`
	UserId      string  `pulumi:"userId"`
}

func GetCloudProjectDatabaseKafkaUserAccessOutput(ctx *pulumi.Context, args GetCloudProjectDatabaseKafkaUserAccessOutputArgs, opts ...pulumi.InvokeOption) GetCloudProjectDatabaseKafkaUserAccessResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloudProjectDatabaseKafkaUserAccessResult, error) {
			args := v.(GetCloudProjectDatabaseKafkaUserAccessArgs)
			r, err := GetCloudProjectDatabaseKafkaUserAccess(ctx, &args, opts...)
			var s GetCloudProjectDatabaseKafkaUserAccessResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloudProjectDatabaseKafkaUserAccessResultOutput)
}

// A collection of arguments for invoking getCloudProjectDatabaseKafkaUserAccess.
type GetCloudProjectDatabaseKafkaUserAccessOutputArgs struct {
	ClusterId   pulumi.StringInput    `pulumi:"clusterId"`
	Id          pulumi.StringPtrInput `pulumi:"id"`
	ServiceName pulumi.StringPtrInput `pulumi:"serviceName"`
	UserId      pulumi.StringInput    `pulumi:"userId"`
}

func (GetCloudProjectDatabaseKafkaUserAccessOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectDatabaseKafkaUserAccessArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectDatabaseKafkaUserAccess.
type GetCloudProjectDatabaseKafkaUserAccessResultOutput struct{ *pulumi.OutputState }

func (GetCloudProjectDatabaseKafkaUserAccessResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectDatabaseKafkaUserAccessResult)(nil)).Elem()
}

func (o GetCloudProjectDatabaseKafkaUserAccessResultOutput) ToGetCloudProjectDatabaseKafkaUserAccessResultOutput() GetCloudProjectDatabaseKafkaUserAccessResultOutput {
	return o
}

func (o GetCloudProjectDatabaseKafkaUserAccessResultOutput) ToGetCloudProjectDatabaseKafkaUserAccessResultOutputWithContext(ctx context.Context) GetCloudProjectDatabaseKafkaUserAccessResultOutput {
	return o
}

func (o GetCloudProjectDatabaseKafkaUserAccessResultOutput) Cert() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectDatabaseKafkaUserAccessResult) string { return v.Cert }).(pulumi.StringOutput)
}

func (o GetCloudProjectDatabaseKafkaUserAccessResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectDatabaseKafkaUserAccessResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o GetCloudProjectDatabaseKafkaUserAccessResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectDatabaseKafkaUserAccessResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCloudProjectDatabaseKafkaUserAccessResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectDatabaseKafkaUserAccessResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o GetCloudProjectDatabaseKafkaUserAccessResultOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCloudProjectDatabaseKafkaUserAccessResult) *string { return v.ServiceName }).(pulumi.StringPtrOutput)
}

func (o GetCloudProjectDatabaseKafkaUserAccessResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectDatabaseKafkaUserAccessResult) string { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudProjectDatabaseKafkaUserAccessResultOutput{})
}
