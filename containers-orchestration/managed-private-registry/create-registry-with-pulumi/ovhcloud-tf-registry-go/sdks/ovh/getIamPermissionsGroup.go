// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIamPermissionsGroup(ctx *pulumi.Context, args *LookupIamPermissionsGroupArgs, opts ...pulumi.InvokeOption) (*LookupIamPermissionsGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupIamPermissionsGroupResult
	err = ctx.InvokePackage("ovh:index/getIamPermissionsGroup:getIamPermissionsGroup", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIamPermissionsGroup.
type LookupIamPermissionsGroupArgs struct {
	Allows      []string `pulumi:"allows"`
	Denies      []string `pulumi:"denies"`
	Description *string  `pulumi:"description"`
	Excepts     []string `pulumi:"excepts"`
	UpdatedAt   *string  `pulumi:"updatedAt"`
	Urn         string   `pulumi:"urn"`
}

// A collection of values returned by getIamPermissionsGroup.
type LookupIamPermissionsGroupResult struct {
	Allows      []string `pulumi:"allows"`
	CreatedAt   string   `pulumi:"createdAt"`
	Denies      []string `pulumi:"denies"`
	Description *string  `pulumi:"description"`
	Excepts     []string `pulumi:"excepts"`
	Id          string   `pulumi:"id"`
	Name        string   `pulumi:"name"`
	Owner       string   `pulumi:"owner"`
	UpdatedAt   string   `pulumi:"updatedAt"`
	Urn         string   `pulumi:"urn"`
}

func LookupIamPermissionsGroupOutput(ctx *pulumi.Context, args LookupIamPermissionsGroupOutputArgs, opts ...pulumi.InvokeOption) LookupIamPermissionsGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIamPermissionsGroupResult, error) {
			args := v.(LookupIamPermissionsGroupArgs)
			r, err := LookupIamPermissionsGroup(ctx, &args, opts...)
			var s LookupIamPermissionsGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIamPermissionsGroupResultOutput)
}

// A collection of arguments for invoking getIamPermissionsGroup.
type LookupIamPermissionsGroupOutputArgs struct {
	Allows      pulumi.StringArrayInput `pulumi:"allows"`
	Denies      pulumi.StringArrayInput `pulumi:"denies"`
	Description pulumi.StringPtrInput   `pulumi:"description"`
	Excepts     pulumi.StringArrayInput `pulumi:"excepts"`
	UpdatedAt   pulumi.StringPtrInput   `pulumi:"updatedAt"`
	Urn         pulumi.StringInput      `pulumi:"urn"`
}

func (LookupIamPermissionsGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIamPermissionsGroupArgs)(nil)).Elem()
}

// A collection of values returned by getIamPermissionsGroup.
type LookupIamPermissionsGroupResultOutput struct{ *pulumi.OutputState }

func (LookupIamPermissionsGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIamPermissionsGroupResult)(nil)).Elem()
}

func (o LookupIamPermissionsGroupResultOutput) ToLookupIamPermissionsGroupResultOutput() LookupIamPermissionsGroupResultOutput {
	return o
}

func (o LookupIamPermissionsGroupResultOutput) ToLookupIamPermissionsGroupResultOutputWithContext(ctx context.Context) LookupIamPermissionsGroupResultOutput {
	return o
}

func (o LookupIamPermissionsGroupResultOutput) Allows() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIamPermissionsGroupResult) []string { return v.Allows }).(pulumi.StringArrayOutput)
}

func (o LookupIamPermissionsGroupResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamPermissionsGroupResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupIamPermissionsGroupResultOutput) Denies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIamPermissionsGroupResult) []string { return v.Denies }).(pulumi.StringArrayOutput)
}

func (o LookupIamPermissionsGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIamPermissionsGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupIamPermissionsGroupResultOutput) Excepts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIamPermissionsGroupResult) []string { return v.Excepts }).(pulumi.StringArrayOutput)
}

func (o LookupIamPermissionsGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamPermissionsGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIamPermissionsGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamPermissionsGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIamPermissionsGroupResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamPermissionsGroupResult) string { return v.Owner }).(pulumi.StringOutput)
}

func (o LookupIamPermissionsGroupResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamPermissionsGroupResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupIamPermissionsGroupResultOutput) Urn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamPermissionsGroupResult) string { return v.Urn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIamPermissionsGroupResultOutput{})
}