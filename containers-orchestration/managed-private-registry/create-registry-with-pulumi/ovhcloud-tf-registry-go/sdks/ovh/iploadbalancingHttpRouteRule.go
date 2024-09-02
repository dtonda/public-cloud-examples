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

type IploadbalancingHttpRouteRule struct {
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

// NewIploadbalancingHttpRouteRule registers a new resource with the given unique name, arguments, and options.
func NewIploadbalancingHttpRouteRule(ctx *pulumi.Context,
	name string, args *IploadbalancingHttpRouteRuleArgs, opts ...pulumi.ResourceOption) (*IploadbalancingHttpRouteRule, error) {
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
	var resource IploadbalancingHttpRouteRule
	err = ctx.RegisterPackageResource("ovh:index/iploadbalancingHttpRouteRule:IploadbalancingHttpRouteRule", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIploadbalancingHttpRouteRule gets an existing IploadbalancingHttpRouteRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIploadbalancingHttpRouteRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IploadbalancingHttpRouteRuleState, opts ...pulumi.ResourceOption) (*IploadbalancingHttpRouteRule, error) {
	var resource IploadbalancingHttpRouteRule
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("ovh:index/iploadbalancingHttpRouteRule:IploadbalancingHttpRouteRule", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IploadbalancingHttpRouteRule resources.
type iploadbalancingHttpRouteRuleState struct {
	DisplayName *string `pulumi:"displayName"`
	Field       *string `pulumi:"field"`
	Match       *string `pulumi:"match"`
	Negate      *bool   `pulumi:"negate"`
	Pattern     *string `pulumi:"pattern"`
	RouteId     *string `pulumi:"routeId"`
	ServiceName *string `pulumi:"serviceName"`
	SubField    *string `pulumi:"subField"`
}

type IploadbalancingHttpRouteRuleState struct {
	DisplayName pulumi.StringPtrInput
	Field       pulumi.StringPtrInput
	Match       pulumi.StringPtrInput
	Negate      pulumi.BoolPtrInput
	Pattern     pulumi.StringPtrInput
	RouteId     pulumi.StringPtrInput
	ServiceName pulumi.StringPtrInput
	SubField    pulumi.StringPtrInput
}

func (IploadbalancingHttpRouteRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadbalancingHttpRouteRuleState)(nil)).Elem()
}

type iploadbalancingHttpRouteRuleArgs struct {
	DisplayName *string `pulumi:"displayName"`
	Field       string  `pulumi:"field"`
	Match       string  `pulumi:"match"`
	Negate      *bool   `pulumi:"negate"`
	Pattern     *string `pulumi:"pattern"`
	RouteId     string  `pulumi:"routeId"`
	ServiceName string  `pulumi:"serviceName"`
	SubField    *string `pulumi:"subField"`
}

// The set of arguments for constructing a IploadbalancingHttpRouteRule resource.
type IploadbalancingHttpRouteRuleArgs struct {
	DisplayName pulumi.StringPtrInput
	Field       pulumi.StringInput
	Match       pulumi.StringInput
	Negate      pulumi.BoolPtrInput
	Pattern     pulumi.StringPtrInput
	RouteId     pulumi.StringInput
	ServiceName pulumi.StringInput
	SubField    pulumi.StringPtrInput
}

func (IploadbalancingHttpRouteRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadbalancingHttpRouteRuleArgs)(nil)).Elem()
}

type IploadbalancingHttpRouteRuleInput interface {
	pulumi.Input

	ToIploadbalancingHttpRouteRuleOutput() IploadbalancingHttpRouteRuleOutput
	ToIploadbalancingHttpRouteRuleOutputWithContext(ctx context.Context) IploadbalancingHttpRouteRuleOutput
}

func (*IploadbalancingHttpRouteRule) ElementType() reflect.Type {
	return reflect.TypeOf((**IploadbalancingHttpRouteRule)(nil)).Elem()
}

func (i *IploadbalancingHttpRouteRule) ToIploadbalancingHttpRouteRuleOutput() IploadbalancingHttpRouteRuleOutput {
	return i.ToIploadbalancingHttpRouteRuleOutputWithContext(context.Background())
}

func (i *IploadbalancingHttpRouteRule) ToIploadbalancingHttpRouteRuleOutputWithContext(ctx context.Context) IploadbalancingHttpRouteRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingHttpRouteRuleOutput)
}

type IploadbalancingHttpRouteRuleOutput struct{ *pulumi.OutputState }

func (IploadbalancingHttpRouteRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IploadbalancingHttpRouteRule)(nil)).Elem()
}

func (o IploadbalancingHttpRouteRuleOutput) ToIploadbalancingHttpRouteRuleOutput() IploadbalancingHttpRouteRuleOutput {
	return o
}

func (o IploadbalancingHttpRouteRuleOutput) ToIploadbalancingHttpRouteRuleOutputWithContext(ctx context.Context) IploadbalancingHttpRouteRuleOutput {
	return o
}

func (o IploadbalancingHttpRouteRuleOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IploadbalancingHttpRouteRule) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o IploadbalancingHttpRouteRuleOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v *IploadbalancingHttpRouteRule) pulumi.StringOutput { return v.Field }).(pulumi.StringOutput)
}

func (o IploadbalancingHttpRouteRuleOutput) Match() pulumi.StringOutput {
	return o.ApplyT(func(v *IploadbalancingHttpRouteRule) pulumi.StringOutput { return v.Match }).(pulumi.StringOutput)
}

func (o IploadbalancingHttpRouteRuleOutput) Negate() pulumi.BoolOutput {
	return o.ApplyT(func(v *IploadbalancingHttpRouteRule) pulumi.BoolOutput { return v.Negate }).(pulumi.BoolOutput)
}

func (o IploadbalancingHttpRouteRuleOutput) Pattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IploadbalancingHttpRouteRule) pulumi.StringPtrOutput { return v.Pattern }).(pulumi.StringPtrOutput)
}

func (o IploadbalancingHttpRouteRuleOutput) RouteId() pulumi.StringOutput {
	return o.ApplyT(func(v *IploadbalancingHttpRouteRule) pulumi.StringOutput { return v.RouteId }).(pulumi.StringOutput)
}

func (o IploadbalancingHttpRouteRuleOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IploadbalancingHttpRouteRule) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

func (o IploadbalancingHttpRouteRuleOutput) SubField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IploadbalancingHttpRouteRule) pulumi.StringPtrOutput { return v.SubField }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IploadbalancingHttpRouteRuleInput)(nil)).Elem(), &IploadbalancingHttpRouteRule{})
	pulumi.RegisterOutputType(IploadbalancingHttpRouteRuleOutput{})
}