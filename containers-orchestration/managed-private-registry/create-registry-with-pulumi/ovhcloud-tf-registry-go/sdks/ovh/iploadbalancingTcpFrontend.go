// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IploadbalancingTcpFrontend struct {
	pulumi.CustomResourceState

	AllowedSources pulumi.StringArrayOutput `pulumi:"allowedSources"`
	DedicatedIpfos pulumi.StringArrayOutput `pulumi:"dedicatedIpfos"`
	DefaultFarmId  pulumi.Float64Output     `pulumi:"defaultFarmId"`
	DefaultSslId   pulumi.Float64Output     `pulumi:"defaultSslId"`
	DeniedSources  pulumi.StringArrayOutput `pulumi:"deniedSources"`
	Disabled       pulumi.BoolPtrOutput     `pulumi:"disabled"`
	DisplayName    pulumi.StringPtrOutput   `pulumi:"displayName"`
	Port           pulumi.StringOutput      `pulumi:"port"`
	ServiceName    pulumi.StringOutput      `pulumi:"serviceName"`
	Ssl            pulumi.BoolPtrOutput     `pulumi:"ssl"`
	Zone           pulumi.StringOutput      `pulumi:"zone"`
}

// NewIploadbalancingTcpFrontend registers a new resource with the given unique name, arguments, and options.
func NewIploadbalancingTcpFrontend(ctx *pulumi.Context,
	name string, args *IploadbalancingTcpFrontendArgs, opts ...pulumi.ResourceOption) (*IploadbalancingTcpFrontend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource IploadbalancingTcpFrontend
	err = ctx.RegisterPackageResource("ovh:index/iploadbalancingTcpFrontend:IploadbalancingTcpFrontend", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIploadbalancingTcpFrontend gets an existing IploadbalancingTcpFrontend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIploadbalancingTcpFrontend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IploadbalancingTcpFrontendState, opts ...pulumi.ResourceOption) (*IploadbalancingTcpFrontend, error) {
	var resource IploadbalancingTcpFrontend
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("ovh:index/iploadbalancingTcpFrontend:IploadbalancingTcpFrontend", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IploadbalancingTcpFrontend resources.
type iploadbalancingTcpFrontendState struct {
	AllowedSources []string `pulumi:"allowedSources"`
	DedicatedIpfos []string `pulumi:"dedicatedIpfos"`
	DefaultFarmId  *float64 `pulumi:"defaultFarmId"`
	DefaultSslId   *float64 `pulumi:"defaultSslId"`
	DeniedSources  []string `pulumi:"deniedSources"`
	Disabled       *bool    `pulumi:"disabled"`
	DisplayName    *string  `pulumi:"displayName"`
	Port           *string  `pulumi:"port"`
	ServiceName    *string  `pulumi:"serviceName"`
	Ssl            *bool    `pulumi:"ssl"`
	Zone           *string  `pulumi:"zone"`
}

type IploadbalancingTcpFrontendState struct {
	AllowedSources pulumi.StringArrayInput
	DedicatedIpfos pulumi.StringArrayInput
	DefaultFarmId  pulumi.Float64PtrInput
	DefaultSslId   pulumi.Float64PtrInput
	DeniedSources  pulumi.StringArrayInput
	Disabled       pulumi.BoolPtrInput
	DisplayName    pulumi.StringPtrInput
	Port           pulumi.StringPtrInput
	ServiceName    pulumi.StringPtrInput
	Ssl            pulumi.BoolPtrInput
	Zone           pulumi.StringPtrInput
}

func (IploadbalancingTcpFrontendState) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadbalancingTcpFrontendState)(nil)).Elem()
}

type iploadbalancingTcpFrontendArgs struct {
	AllowedSources []string `pulumi:"allowedSources"`
	DedicatedIpfos []string `pulumi:"dedicatedIpfos"`
	DefaultFarmId  *float64 `pulumi:"defaultFarmId"`
	DefaultSslId   *float64 `pulumi:"defaultSslId"`
	DeniedSources  []string `pulumi:"deniedSources"`
	Disabled       *bool    `pulumi:"disabled"`
	DisplayName    *string  `pulumi:"displayName"`
	Port           string   `pulumi:"port"`
	ServiceName    string   `pulumi:"serviceName"`
	Ssl            *bool    `pulumi:"ssl"`
	Zone           string   `pulumi:"zone"`
}

// The set of arguments for constructing a IploadbalancingTcpFrontend resource.
type IploadbalancingTcpFrontendArgs struct {
	AllowedSources pulumi.StringArrayInput
	DedicatedIpfos pulumi.StringArrayInput
	DefaultFarmId  pulumi.Float64PtrInput
	DefaultSslId   pulumi.Float64PtrInput
	DeniedSources  pulumi.StringArrayInput
	Disabled       pulumi.BoolPtrInput
	DisplayName    pulumi.StringPtrInput
	Port           pulumi.StringInput
	ServiceName    pulumi.StringInput
	Ssl            pulumi.BoolPtrInput
	Zone           pulumi.StringInput
}

func (IploadbalancingTcpFrontendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadbalancingTcpFrontendArgs)(nil)).Elem()
}

type IploadbalancingTcpFrontendInput interface {
	pulumi.Input

	ToIploadbalancingTcpFrontendOutput() IploadbalancingTcpFrontendOutput
	ToIploadbalancingTcpFrontendOutputWithContext(ctx context.Context) IploadbalancingTcpFrontendOutput
}

func (*IploadbalancingTcpFrontend) ElementType() reflect.Type {
	return reflect.TypeOf((**IploadbalancingTcpFrontend)(nil)).Elem()
}

func (i *IploadbalancingTcpFrontend) ToIploadbalancingTcpFrontendOutput() IploadbalancingTcpFrontendOutput {
	return i.ToIploadbalancingTcpFrontendOutputWithContext(context.Background())
}

func (i *IploadbalancingTcpFrontend) ToIploadbalancingTcpFrontendOutputWithContext(ctx context.Context) IploadbalancingTcpFrontendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingTcpFrontendOutput)
}

type IploadbalancingTcpFrontendOutput struct{ *pulumi.OutputState }

func (IploadbalancingTcpFrontendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IploadbalancingTcpFrontend)(nil)).Elem()
}

func (o IploadbalancingTcpFrontendOutput) ToIploadbalancingTcpFrontendOutput() IploadbalancingTcpFrontendOutput {
	return o
}

func (o IploadbalancingTcpFrontendOutput) ToIploadbalancingTcpFrontendOutputWithContext(ctx context.Context) IploadbalancingTcpFrontendOutput {
	return o
}

func (o IploadbalancingTcpFrontendOutput) AllowedSources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IploadbalancingTcpFrontend) pulumi.StringArrayOutput { return v.AllowedSources }).(pulumi.StringArrayOutput)
}

func (o IploadbalancingTcpFrontendOutput) DedicatedIpfos() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IploadbalancingTcpFrontend) pulumi.StringArrayOutput { return v.DedicatedIpfos }).(pulumi.StringArrayOutput)
}

func (o IploadbalancingTcpFrontendOutput) DefaultFarmId() pulumi.Float64Output {
	return o.ApplyT(func(v *IploadbalancingTcpFrontend) pulumi.Float64Output { return v.DefaultFarmId }).(pulumi.Float64Output)
}

func (o IploadbalancingTcpFrontendOutput) DefaultSslId() pulumi.Float64Output {
	return o.ApplyT(func(v *IploadbalancingTcpFrontend) pulumi.Float64Output { return v.DefaultSslId }).(pulumi.Float64Output)
}

func (o IploadbalancingTcpFrontendOutput) DeniedSources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IploadbalancingTcpFrontend) pulumi.StringArrayOutput { return v.DeniedSources }).(pulumi.StringArrayOutput)
}

func (o IploadbalancingTcpFrontendOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IploadbalancingTcpFrontend) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

func (o IploadbalancingTcpFrontendOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IploadbalancingTcpFrontend) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o IploadbalancingTcpFrontendOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *IploadbalancingTcpFrontend) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

func (o IploadbalancingTcpFrontendOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IploadbalancingTcpFrontend) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

func (o IploadbalancingTcpFrontendOutput) Ssl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IploadbalancingTcpFrontend) pulumi.BoolPtrOutput { return v.Ssl }).(pulumi.BoolPtrOutput)
}

func (o IploadbalancingTcpFrontendOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *IploadbalancingTcpFrontend) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IploadbalancingTcpFrontendInput)(nil)).Elem(), &IploadbalancingTcpFrontend{})
	pulumi.RegisterOutputType(IploadbalancingTcpFrontendOutput{})
}