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

type IploadbalancingTcpRouteRule struct {
	pulumi.CustomResourceState

	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	Field       pulumi.StringOutput    `pulumi:"field"`
	Match       pulumi.StringOutput    `pulumi:"match"`
	Negate      pulumi.BoolOutput      `pulumi:"negate"`
	Pattern     pulumi.StringPtrOutput `pulumi:"pattern"`
	RouteId     pulumi.StringOutput    `pulumi:"routeId"`
	ServiceName pulumi.StringOutput    `pulumi:"serviceName"`
	SubField    pulumi.StringPtrOutput `pulumi:"subField"`
}

// NewIploadbalancingTcpRouteRule registers a new resource with the given unique name, arguments, and options.
func NewIploadbalancingTcpRouteRule(ctx *pulumi.Context,
	name string, args *IploadbalancingTcpRouteRuleArgs, opts ...pulumi.ResourceOption) (*IploadbalancingTcpRouteRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Field == nil {
		return nil, errors.New("invalid value for required argument 'Field'")
	}
	if args.Match == nil {
		return nil, errors.New("invalid value for required argument 'Match'")
	}
	if args.RouteId == nil {
		return nil, errors.New("invalid value for required argument 'RouteId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource IploadbalancingTcpRouteRule
	err = ctx.RegisterPackageResource("ovh:index/iploadbalancingTcpRouteRule:IploadbalancingTcpRouteRule", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIploadbalancingTcpRouteRule gets an existing IploadbalancingTcpRouteRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIploadbalancingTcpRouteRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IploadbalancingTcpRouteRuleState, opts ...pulumi.ResourceOption) (*IploadbalancingTcpRouteRule, error) {
	var resource IploadbalancingTcpRouteRule
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("ovh:index/iploadbalancingTcpRouteRule:IploadbalancingTcpRouteRule", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IploadbalancingTcpRouteRule resources.
type iploadbalancingTcpRouteRuleState struct {
	DisplayName *string `pulumi:"displayName"`
	Field       *string `pulumi:"field"`
	Match       *string `pulumi:"match"`
	Negate      *bool   `pulumi:"negate"`
	Pattern     *string `pulumi:"pattern"`
	RouteId     *string `pulumi:"routeId"`
	ServiceName *string `pulumi:"serviceName"`
	SubField    *string `pulumi:"subField"`
}

type IploadbalancingTcpRouteRuleState struct {
	DisplayName pulumi.StringPtrInput
	Field       pulumi.StringPtrInput
	Match       pulumi.StringPtrInput
	Negate      pulumi.BoolPtrInput
	Pattern     pulumi.StringPtrInput
	RouteId     pulumi.StringPtrInput
	ServiceName pulumi.StringPtrInput
	SubField    pulumi.StringPtrInput
}

func (IploadbalancingTcpRouteRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadbalancingTcpRouteRuleState)(nil)).Elem()
}

type iploadbalancingTcpRouteRuleArgs struct {
	DisplayName *string `pulumi:"displayName"`
	Field       string  `pulumi:"field"`
	Match       string  `pulumi:"match"`
	Negate      *bool   `pulumi:"negate"`
	Pattern     *string `pulumi:"pattern"`
	RouteId     string  `pulumi:"routeId"`
	ServiceName string  `pulumi:"serviceName"`
	SubField    *string `pulumi:"subField"`
}

// The set of arguments for constructing a IploadbalancingTcpRouteRule resource.
type IploadbalancingTcpRouteRuleArgs struct {
	DisplayName pulumi.StringPtrInput
	Field       pulumi.StringInput
	Match       pulumi.StringInput
	Negate      pulumi.BoolPtrInput
	Pattern     pulumi.StringPtrInput
	RouteId     pulumi.StringInput
	ServiceName pulumi.StringInput
	SubField    pulumi.StringPtrInput
}

func (IploadbalancingTcpRouteRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadbalancingTcpRouteRuleArgs)(nil)).Elem()
}

type IploadbalancingTcpRouteRuleInput interface {
	pulumi.Input

	ToIploadbalancingTcpRouteRuleOutput() IploadbalancingTcpRouteRuleOutput
	ToIploadbalancingTcpRouteRuleOutputWithContext(ctx context.Context) IploadbalancingTcpRouteRuleOutput
}

func (*IploadbalancingTcpRouteRule) ElementType() reflect.Type {
	return reflect.TypeOf((**IploadbalancingTcpRouteRule)(nil)).Elem()
}

func (i *IploadbalancingTcpRouteRule) ToIploadbalancingTcpRouteRuleOutput() IploadbalancingTcpRouteRuleOutput {
	return i.ToIploadbalancingTcpRouteRuleOutputWithContext(context.Background())
}

func (i *IploadbalancingTcpRouteRule) ToIploadbalancingTcpRouteRuleOutputWithContext(ctx context.Context) IploadbalancingTcpRouteRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingTcpRouteRuleOutput)
}

type IploadbalancingTcpRouteRuleOutput struct{ *pulumi.OutputState }

func (IploadbalancingTcpRouteRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IploadbalancingTcpRouteRule)(nil)).Elem()
}

func (o IploadbalancingTcpRouteRuleOutput) ToIploadbalancingTcpRouteRuleOutput() IploadbalancingTcpRouteRuleOutput {
	return o
}

func (o IploadbalancingTcpRouteRuleOutput) ToIploadbalancingTcpRouteRuleOutputWithContext(ctx context.Context) IploadbalancingTcpRouteRuleOutput {
	return o
}

func (o IploadbalancingTcpRouteRuleOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IploadbalancingTcpRouteRule) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o IploadbalancingTcpRouteRuleOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v *IploadbalancingTcpRouteRule) pulumi.StringOutput { return v.Field }).(pulumi.StringOutput)
}

func (o IploadbalancingTcpRouteRuleOutput) Match() pulumi.StringOutput {
	return o.ApplyT(func(v *IploadbalancingTcpRouteRule) pulumi.StringOutput { return v.Match }).(pulumi.StringOutput)
}

func (o IploadbalancingTcpRouteRuleOutput) Negate() pulumi.BoolOutput {
	return o.ApplyT(func(v *IploadbalancingTcpRouteRule) pulumi.BoolOutput { return v.Negate }).(pulumi.BoolOutput)
}

func (o IploadbalancingTcpRouteRuleOutput) Pattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IploadbalancingTcpRouteRule) pulumi.StringPtrOutput { return v.Pattern }).(pulumi.StringPtrOutput)
}

func (o IploadbalancingTcpRouteRuleOutput) RouteId() pulumi.StringOutput {
	return o.ApplyT(func(v *IploadbalancingTcpRouteRule) pulumi.StringOutput { return v.RouteId }).(pulumi.StringOutput)
}

func (o IploadbalancingTcpRouteRuleOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IploadbalancingTcpRouteRule) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

func (o IploadbalancingTcpRouteRuleOutput) SubField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IploadbalancingTcpRouteRule) pulumi.StringPtrOutput { return v.SubField }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IploadbalancingTcpRouteRuleInput)(nil)).Elem(), &IploadbalancingTcpRouteRule{})
	pulumi.RegisterOutputType(IploadbalancingTcpRouteRuleOutput{})
}