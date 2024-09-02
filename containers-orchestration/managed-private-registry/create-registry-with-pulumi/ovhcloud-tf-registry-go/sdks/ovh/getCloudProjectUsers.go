// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCloudProjectUsers(ctx *pulumi.Context, args *GetCloudProjectUsersArgs, opts ...pulumi.InvokeOption) (*GetCloudProjectUsersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetCloudProjectUsersResult
	err = ctx.InvokePackage("ovh:index/getCloudProjectUsers:getCloudProjectUsers", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectUsers.
type GetCloudProjectUsersArgs struct {
	Id          *string `pulumi:"id"`
	ServiceName *string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectUsers.
type GetCloudProjectUsersResult struct {
	Id          string                     `pulumi:"id"`
	ServiceName *string                    `pulumi:"serviceName"`
	Users       []GetCloudProjectUsersUser `pulumi:"users"`
}

func GetCloudProjectUsersOutput(ctx *pulumi.Context, args GetCloudProjectUsersOutputArgs, opts ...pulumi.InvokeOption) GetCloudProjectUsersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloudProjectUsersResult, error) {
			args := v.(GetCloudProjectUsersArgs)
			r, err := GetCloudProjectUsers(ctx, &args, opts...)
			var s GetCloudProjectUsersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloudProjectUsersResultOutput)
}

// A collection of arguments for invoking getCloudProjectUsers.
type GetCloudProjectUsersOutputArgs struct {
	Id          pulumi.StringPtrInput `pulumi:"id"`
	ServiceName pulumi.StringPtrInput `pulumi:"serviceName"`
}

func (GetCloudProjectUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectUsersArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectUsers.
type GetCloudProjectUsersResultOutput struct{ *pulumi.OutputState }

func (GetCloudProjectUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectUsersResult)(nil)).Elem()
}

func (o GetCloudProjectUsersResultOutput) ToGetCloudProjectUsersResultOutput() GetCloudProjectUsersResultOutput {
	return o
}

func (o GetCloudProjectUsersResultOutput) ToGetCloudProjectUsersResultOutputWithContext(ctx context.Context) GetCloudProjectUsersResultOutput {
	return o
}

func (o GetCloudProjectUsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectUsersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCloudProjectUsersResultOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCloudProjectUsersResult) *string { return v.ServiceName }).(pulumi.StringPtrOutput)
}

func (o GetCloudProjectUsersResultOutput) Users() GetCloudProjectUsersUserArrayOutput {
	return o.ApplyT(func(v GetCloudProjectUsersResult) []GetCloudProjectUsersUser { return v.Users }).(GetCloudProjectUsersUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudProjectUsersResultOutput{})
}