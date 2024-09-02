// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHostingPrivatedatabase(ctx *pulumi.Context, args *LookupHostingPrivatedatabaseArgs, opts ...pulumi.InvokeOption) (*LookupHostingPrivatedatabaseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupHostingPrivatedatabaseResult
	err = ctx.InvokePackage("ovh:index/getHostingPrivatedatabase:getHostingPrivatedatabase", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHostingPrivatedatabase.
type LookupHostingPrivatedatabaseArgs struct {
	Id          *string `pulumi:"id"`
	ServiceName string  `pulumi:"serviceName"`
}

// A collection of values returned by getHostingPrivatedatabase.
type LookupHostingPrivatedatabaseResult struct {
	Cpu            float64 `pulumi:"cpu"`
	Datacenter     string  `pulumi:"datacenter"`
	DisplayName    string  `pulumi:"displayName"`
	Hostname       string  `pulumi:"hostname"`
	HostnameFtp    string  `pulumi:"hostnameFtp"`
	Id             string  `pulumi:"id"`
	Infrastructure string  `pulumi:"infrastructure"`
	Offer          string  `pulumi:"offer"`
	Port           float64 `pulumi:"port"`
	PortFtp        float64 `pulumi:"portFtp"`
	QuotaSize      float64 `pulumi:"quotaSize"`
	QuotaUsed      float64 `pulumi:"quotaUsed"`
	Ram            float64 `pulumi:"ram"`
	Server         string  `pulumi:"server"`
	ServiceName    string  `pulumi:"serviceName"`
	State          string  `pulumi:"state"`
	Type           string  `pulumi:"type"`
	Urn            string  `pulumi:"urn"`
	Version        string  `pulumi:"version"`
	VersionLabel   string  `pulumi:"versionLabel"`
	VersionNumber  float64 `pulumi:"versionNumber"`
}

func LookupHostingPrivatedatabaseOutput(ctx *pulumi.Context, args LookupHostingPrivatedatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupHostingPrivatedatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHostingPrivatedatabaseResult, error) {
			args := v.(LookupHostingPrivatedatabaseArgs)
			r, err := LookupHostingPrivatedatabase(ctx, &args, opts...)
			var s LookupHostingPrivatedatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHostingPrivatedatabaseResultOutput)
}

// A collection of arguments for invoking getHostingPrivatedatabase.
type LookupHostingPrivatedatabaseOutputArgs struct {
	Id          pulumi.StringPtrInput `pulumi:"id"`
	ServiceName pulumi.StringInput    `pulumi:"serviceName"`
}

func (LookupHostingPrivatedatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostingPrivatedatabaseArgs)(nil)).Elem()
}

// A collection of values returned by getHostingPrivatedatabase.
type LookupHostingPrivatedatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupHostingPrivatedatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostingPrivatedatabaseResult)(nil)).Elem()
}

func (o LookupHostingPrivatedatabaseResultOutput) ToLookupHostingPrivatedatabaseResultOutput() LookupHostingPrivatedatabaseResultOutput {
	return o
}

func (o LookupHostingPrivatedatabaseResultOutput) ToLookupHostingPrivatedatabaseResultOutputWithContext(ctx context.Context) LookupHostingPrivatedatabaseResultOutput {
	return o
}

func (o LookupHostingPrivatedatabaseResultOutput) Cpu() pulumi.Float64Output {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) float64 { return v.Cpu }).(pulumi.Float64Output)
}

func (o LookupHostingPrivatedatabaseResultOutput) Datacenter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) string { return v.Datacenter }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) string { return v.Hostname }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseResultOutput) HostnameFtp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) string { return v.HostnameFtp }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseResultOutput) Infrastructure() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) string { return v.Infrastructure }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseResultOutput) Offer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) string { return v.Offer }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseResultOutput) Port() pulumi.Float64Output {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) float64 { return v.Port }).(pulumi.Float64Output)
}

func (o LookupHostingPrivatedatabaseResultOutput) PortFtp() pulumi.Float64Output {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) float64 { return v.PortFtp }).(pulumi.Float64Output)
}

func (o LookupHostingPrivatedatabaseResultOutput) QuotaSize() pulumi.Float64Output {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) float64 { return v.QuotaSize }).(pulumi.Float64Output)
}

func (o LookupHostingPrivatedatabaseResultOutput) QuotaUsed() pulumi.Float64Output {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) float64 { return v.QuotaUsed }).(pulumi.Float64Output)
}

func (o LookupHostingPrivatedatabaseResultOutput) Ram() pulumi.Float64Output {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) float64 { return v.Ram }).(pulumi.Float64Output)
}

func (o LookupHostingPrivatedatabaseResultOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) string { return v.Server }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseResultOutput) Urn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) string { return v.Urn }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) string { return v.Version }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseResultOutput) VersionLabel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) string { return v.VersionLabel }).(pulumi.StringOutput)
}

func (o LookupHostingPrivatedatabaseResultOutput) VersionNumber() pulumi.Float64Output {
	return o.ApplyT(func(v LookupHostingPrivatedatabaseResult) float64 { return v.VersionNumber }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(LookupHostingPrivatedatabaseResultOutput{})
}