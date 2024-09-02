// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCloudProjectDatabaseOpensearchUser(ctx *pulumi.Context, args *LookupCloudProjectDatabaseOpensearchUserArgs, opts ...pulumi.InvokeOption) (*LookupCloudProjectDatabaseOpensearchUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupCloudProjectDatabaseOpensearchUserResult
	err = ctx.InvokePackage("ovh:index/getCloudProjectDatabaseOpensearchUser:getCloudProjectDatabaseOpensearchUser", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectDatabaseOpensearchUser.
type LookupCloudProjectDatabaseOpensearchUserArgs struct {
	ClusterId   string  `pulumi:"clusterId"`
	Id          *string `pulumi:"id"`
	Name        string  `pulumi:"name"`
	ServiceName *string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectDatabaseOpensearchUser.
type LookupCloudProjectDatabaseOpensearchUserResult struct {
	Acls        []GetCloudProjectDatabaseOpensearchUserAcl `pulumi:"acls"`
	ClusterId   string                                     `pulumi:"clusterId"`
	CreatedAt   string                                     `pulumi:"createdAt"`
	Id          string                                     `pulumi:"id"`
	Name        string                                     `pulumi:"name"`
	ServiceName *string                                    `pulumi:"serviceName"`
	Status      string                                     `pulumi:"status"`
}

func LookupCloudProjectDatabaseOpensearchUserOutput(ctx *pulumi.Context, args LookupCloudProjectDatabaseOpensearchUserOutputArgs, opts ...pulumi.InvokeOption) LookupCloudProjectDatabaseOpensearchUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudProjectDatabaseOpensearchUserResult, error) {
			args := v.(LookupCloudProjectDatabaseOpensearchUserArgs)
			r, err := LookupCloudProjectDatabaseOpensearchUser(ctx, &args, opts...)
			var s LookupCloudProjectDatabaseOpensearchUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudProjectDatabaseOpensearchUserResultOutput)
}

// A collection of arguments for invoking getCloudProjectDatabaseOpensearchUser.
type LookupCloudProjectDatabaseOpensearchUserOutputArgs struct {
	ClusterId   pulumi.StringInput    `pulumi:"clusterId"`
	Id          pulumi.StringPtrInput `pulumi:"id"`
	Name        pulumi.StringInput    `pulumi:"name"`
	ServiceName pulumi.StringPtrInput `pulumi:"serviceName"`
}

func (LookupCloudProjectDatabaseOpensearchUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProjectDatabaseOpensearchUserArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectDatabaseOpensearchUser.
type LookupCloudProjectDatabaseOpensearchUserResultOutput struct{ *pulumi.OutputState }

func (LookupCloudProjectDatabaseOpensearchUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProjectDatabaseOpensearchUserResult)(nil)).Elem()
}

func (o LookupCloudProjectDatabaseOpensearchUserResultOutput) ToLookupCloudProjectDatabaseOpensearchUserResultOutput() LookupCloudProjectDatabaseOpensearchUserResultOutput {
	return o
}

func (o LookupCloudProjectDatabaseOpensearchUserResultOutput) ToLookupCloudProjectDatabaseOpensearchUserResultOutputWithContext(ctx context.Context) LookupCloudProjectDatabaseOpensearchUserResultOutput {
	return o
}

func (o LookupCloudProjectDatabaseOpensearchUserResultOutput) Acls() GetCloudProjectDatabaseOpensearchUserAclArrayOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabaseOpensearchUserResult) []GetCloudProjectDatabaseOpensearchUserAcl {
		return v.Acls
	}).(GetCloudProjectDatabaseOpensearchUserAclArrayOutput)
}

func (o LookupCloudProjectDatabaseOpensearchUserResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabaseOpensearchUserResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o LookupCloudProjectDatabaseOpensearchUserResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabaseOpensearchUserResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupCloudProjectDatabaseOpensearchUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabaseOpensearchUserResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudProjectDatabaseOpensearchUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabaseOpensearchUserResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCloudProjectDatabaseOpensearchUserResultOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabaseOpensearchUserResult) *string { return v.ServiceName }).(pulumi.StringPtrOutput)
}

func (o LookupCloudProjectDatabaseOpensearchUserResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabaseOpensearchUserResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudProjectDatabaseOpensearchUserResultOutput{})
}
