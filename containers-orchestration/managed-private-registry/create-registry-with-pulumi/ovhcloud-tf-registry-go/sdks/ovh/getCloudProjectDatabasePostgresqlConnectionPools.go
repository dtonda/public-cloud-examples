// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCloudProjectDatabasePostgresqlConnectionPools(ctx *pulumi.Context, args *GetCloudProjectDatabasePostgresqlConnectionPoolsArgs, opts ...pulumi.InvokeOption) (*GetCloudProjectDatabasePostgresqlConnectionPoolsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetCloudProjectDatabasePostgresqlConnectionPoolsResult
	err = ctx.InvokePackage("ovh:index/getCloudProjectDatabasePostgresqlConnectionPools:getCloudProjectDatabasePostgresqlConnectionPools", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectDatabasePostgresqlConnectionPools.
type GetCloudProjectDatabasePostgresqlConnectionPoolsArgs struct {
	ClusterId   string  `pulumi:"clusterId"`
	Id          *string `pulumi:"id"`
	ServiceName *string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectDatabasePostgresqlConnectionPools.
type GetCloudProjectDatabasePostgresqlConnectionPoolsResult struct {
	ClusterId         string   `pulumi:"clusterId"`
	ConnectionPoolIds []string `pulumi:"connectionPoolIds"`
	Id                string   `pulumi:"id"`
	ServiceName       *string  `pulumi:"serviceName"`
}

func GetCloudProjectDatabasePostgresqlConnectionPoolsOutput(ctx *pulumi.Context, args GetCloudProjectDatabasePostgresqlConnectionPoolsOutputArgs, opts ...pulumi.InvokeOption) GetCloudProjectDatabasePostgresqlConnectionPoolsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloudProjectDatabasePostgresqlConnectionPoolsResult, error) {
			args := v.(GetCloudProjectDatabasePostgresqlConnectionPoolsArgs)
			r, err := GetCloudProjectDatabasePostgresqlConnectionPools(ctx, &args, opts...)
			var s GetCloudProjectDatabasePostgresqlConnectionPoolsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloudProjectDatabasePostgresqlConnectionPoolsResultOutput)
}

// A collection of arguments for invoking getCloudProjectDatabasePostgresqlConnectionPools.
type GetCloudProjectDatabasePostgresqlConnectionPoolsOutputArgs struct {
	ClusterId   pulumi.StringInput    `pulumi:"clusterId"`
	Id          pulumi.StringPtrInput `pulumi:"id"`
	ServiceName pulumi.StringPtrInput `pulumi:"serviceName"`
}

func (GetCloudProjectDatabasePostgresqlConnectionPoolsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectDatabasePostgresqlConnectionPoolsArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectDatabasePostgresqlConnectionPools.
type GetCloudProjectDatabasePostgresqlConnectionPoolsResultOutput struct{ *pulumi.OutputState }

func (GetCloudProjectDatabasePostgresqlConnectionPoolsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectDatabasePostgresqlConnectionPoolsResult)(nil)).Elem()
}

func (o GetCloudProjectDatabasePostgresqlConnectionPoolsResultOutput) ToGetCloudProjectDatabasePostgresqlConnectionPoolsResultOutput() GetCloudProjectDatabasePostgresqlConnectionPoolsResultOutput {
	return o
}

func (o GetCloudProjectDatabasePostgresqlConnectionPoolsResultOutput) ToGetCloudProjectDatabasePostgresqlConnectionPoolsResultOutputWithContext(ctx context.Context) GetCloudProjectDatabasePostgresqlConnectionPoolsResultOutput {
	return o
}

func (o GetCloudProjectDatabasePostgresqlConnectionPoolsResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectDatabasePostgresqlConnectionPoolsResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o GetCloudProjectDatabasePostgresqlConnectionPoolsResultOutput) ConnectionPoolIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCloudProjectDatabasePostgresqlConnectionPoolsResult) []string { return v.ConnectionPoolIds }).(pulumi.StringArrayOutput)
}

func (o GetCloudProjectDatabasePostgresqlConnectionPoolsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectDatabasePostgresqlConnectionPoolsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCloudProjectDatabasePostgresqlConnectionPoolsResultOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCloudProjectDatabasePostgresqlConnectionPoolsResult) *string { return v.ServiceName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudProjectDatabasePostgresqlConnectionPoolsResultOutput{})
}