// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHostingPrivatedatabaseUserGrant(ctx *pulumi.Context, args *LookupHostingPrivatedatabaseUserGrantArgs, opts ...pulumi.InvokeOption) (*LookupHostingPrivatedatabaseUserGrantResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupHostingPrivatedatabaseUserGrantResult
	err = ctx.InvokePackage("ovh:index/getHostingPrivatedatabaseUserGrant:getHostingPrivatedatabaseUserGrant", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHostingPrivatedatabaseUserGrant.
type LookupHostingPrivatedatabaseUserGrantArgs struct {
	DatabaseName string  `pulumi:"databaseName"`
	Id           *string `pulumi:"id"`
	ServiceName  string  `pulumi:"serviceName"`
	UserName     string  `pulumi:"userName"`
}

// A collection of values returned by getHostingPrivatedatabaseUserGrant.
type LookupHostingPrivatedatabaseUserGrantResult struct {
	CreationDate string `pulumi:"creationDate"`
	DatabaseName string `pulumi:"databaseName"`
	Grant        string `pulumi:"grant"`
	Id           string `pulumi:"id"`
	ServiceName  string `pulumi:"serviceName"`
	UserName     string `pulumi:"userName"`
}

func LookupHostingPrivatedatabaseUserGrantOutput(ctx *pulumi.Context, args LookupHostingPrivatedatabaseUserGrantOutputArgs, opts ...pulumi.InvokeOption) LookupHostingPrivatedatabaseUserGrantResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHostingPrivatedatabaseUserGrantResult, error) {
			args := v.(LookupHostingPrivatedatabaseUserGrantArgs)
			r, err := LookupHostingPrivatedatabaseUserGrant(ctx, &args, opts...)
			var s LookupHostingPrivatedatabaseUserGrantResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHostingPrivatedatabaseUserGrantResultOutput)
}

// A collection of arguments for invoking getHostingPrivatedatabaseUserGrant.
type LookupHostingPrivatedatabaseUserGrantOutputArgs struct {
	DatabaseName pulumi.StringInput    `pulumi:"databaseName"`
	Id           pulumi.StringPtrInput `pulumi:"id"`
	ServiceName  pulumi.StringInput    `pulumi:"serviceName"`
	UserName     pulumi.StringInput    `pulumi:"userName"`
}

func (LookupHostingPrivatedatabaseUserGrantOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostingPrivatedatabaseUserGrantArgs)(nil)).Elem()
}

// A collection of values returned by getHostingPrivatedatabaseUserGrant.
type LookupHostingPrivatedatabaseUserGrantResultOutput struct{ *pulumi.OutputState }

func (LookupHostingPrivatedatabaseUserGrantResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostingPrivatedatabaseUserGrantResult)(nil)).Elem()
}

func (o LookupHostingPrivatedatabaseUserGrantResultOutput) ToLookupHostingPrivatedatabaseUserGrantResultOutput() LookupHostingPrivatedatabaseUserGrantResultOutput {
	return o
}

func (o LookupHostingPrivatedatabaseUserGrantResultOutput) ToLookupHostingPrivatedatabaseUserGrantResultOutputWithContext(ctx context.Context) LookupHostingPrivatedatabaseUserGrantResultOutput {
	return o
}

func (o LookupHostingPrivatedatabaseUserGrantResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseUserGrantResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseUserGrantResultOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseUserGrantResult) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseUserGrantResultOutput) Grant() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseUserGrantResult) string { return v.Grant }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseUserGrantResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseUserGrantResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseUserGrantResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseUserGrantResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseUserGrantResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseUserGrantResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHostingPrivatedatabaseUserGrantResultOutput{})
}