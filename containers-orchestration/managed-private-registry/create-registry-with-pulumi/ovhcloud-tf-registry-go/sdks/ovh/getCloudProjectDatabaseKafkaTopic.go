// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCloudProjectDatabaseKafkaTopic(ctx *pulumi.Context, args *LookupCloudProjectDatabaseKafkaTopicArgs, opts ...pulumi.InvokeOption) (*LookupCloudProjectDatabaseKafkaTopicResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupCloudProjectDatabaseKafkaTopicResult
	err = ctx.InvokePackage("ovh:index/getCloudProjectDatabaseKafkaTopic:getCloudProjectDatabaseKafkaTopic", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectDatabaseKafkaTopic.
type LookupCloudProjectDatabaseKafkaTopicArgs struct {
	ClusterId   string  `pulumi:"clusterId"`
	Id          string  `pulumi:"id"`
	ServiceName *string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectDatabaseKafkaTopic.
type LookupCloudProjectDatabaseKafkaTopicResult struct {
	ClusterId         string  `pulumi:"clusterId"`
	Id                string  `pulumi:"id"`
	MinInsyncReplicas float64 `pulumi:"minInsyncReplicas"`
	Name              string  `pulumi:"name"`
	Partitions        float64 `pulumi:"partitions"`
	Replication       float64 `pulumi:"replication"`
	RetentionBytes    float64 `pulumi:"retentionBytes"`
	RetentionHours    float64 `pulumi:"retentionHours"`
	ServiceName       *string `pulumi:"serviceName"`
}

func LookupCloudProjectDatabaseKafkaTopicOutput(ctx *pulumi.Context, args LookupCloudProjectDatabaseKafkaTopicOutputArgs, opts ...pulumi.InvokeOption) LookupCloudProjectDatabaseKafkaTopicResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudProjectDatabaseKafkaTopicResult, error) {
			args := v.(LookupCloudProjectDatabaseKafkaTopicArgs)
			r, err := LookupCloudProjectDatabaseKafkaTopic(ctx, &args, opts...)
			var s LookupCloudProjectDatabaseKafkaTopicResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudProjectDatabaseKafkaTopicResultOutput)
}

// A collection of arguments for invoking getCloudProjectDatabaseKafkaTopic.
type LookupCloudProjectDatabaseKafkaTopicOutputArgs struct {
	ClusterId   pulumi.StringInput    `pulumi:"clusterId"`
	Id          pulumi.StringInput    `pulumi:"id"`
	ServiceName pulumi.StringPtrInput `pulumi:"serviceName"`
}

func (LookupCloudProjectDatabaseKafkaTopicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProjectDatabaseKafkaTopicArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectDatabaseKafkaTopic.
type LookupCloudProjectDatabaseKafkaTopicResultOutput struct{ *pulumi.OutputState }

func (LookupCloudProjectDatabaseKafkaTopicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProjectDatabaseKafkaTopicResult)(nil)).Elem()
}

func (o LookupCloudProjectDatabaseKafkaTopicResultOutput) ToLookupCloudProjectDatabaseKafkaTopicResultOutput() LookupCloudProjectDatabaseKafkaTopicResultOutput {
	return o
}

func (o LookupCloudProjectDatabaseKafkaTopicResultOutput) ToLookupCloudProjectDatabaseKafkaTopicResultOutputWithContext(ctx context.Context) LookupCloudProjectDatabaseKafkaTopicResultOutput {
	return o
}

func (o LookupCloudProjectDatabaseKafkaTopicResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabaseKafkaTopicResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o LookupCloudProjectDatabaseKafkaTopicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabaseKafkaTopicResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudProjectDatabaseKafkaTopicResultOutput) MinInsyncReplicas() pulumi.Float64Output {
	return o.ApplyT(func(v LookupCloudProjectDatabaseKafkaTopicResult) float64 { return v.MinInsyncReplicas }).(pulumi.Float64Output)
}

func (o LookupCloudProjectDatabaseKafkaTopicResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabaseKafkaTopicResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCloudProjectDatabaseKafkaTopicResultOutput) Partitions() pulumi.Float64Output {
	return o.ApplyT(func(v LookupCloudProjectDatabaseKafkaTopicResult) float64 { return v.Partitions }).(pulumi.Float64Output)
}

func (o LookupCloudProjectDatabaseKafkaTopicResultOutput) Replication() pulumi.Float64Output {
	return o.ApplyT(func(v LookupCloudProjectDatabaseKafkaTopicResult) float64 { return v.Replication }).(pulumi.Float64Output)
}

func (o LookupCloudProjectDatabaseKafkaTopicResultOutput) RetentionBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupCloudProjectDatabaseKafkaTopicResult) float64 { return v.RetentionBytes }).(pulumi.Float64Output)
}

func (o LookupCloudProjectDatabaseKafkaTopicResultOutput) RetentionHours() pulumi.Float64Output {
	return o.ApplyT(func(v LookupCloudProjectDatabaseKafkaTopicResult) float64 { return v.RetentionHours }).(pulumi.Float64Output)
}

func (o LookupCloudProjectDatabaseKafkaTopicResultOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabaseKafkaTopicResult) *string { return v.ServiceName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudProjectDatabaseKafkaTopicResultOutput{})
}
