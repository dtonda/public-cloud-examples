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

type IpFirewallRule struct {
	pulumi.CustomResourceState

	// Possible values for action
	Action       pulumi.StringOutput `pulumi:"action"`
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// Destination ip for your rule
	Destination pulumi.StringOutput `pulumi:"destination"`
	// Destination port for your rule. Only with TCP/UDP protocol
	DestinationPort pulumi.Float64Output `pulumi:"destinationPort"`
	// Destination port range for your rule. Only with TCP/UDP protocol
	DestinationPortDesc pulumi.StringOutput `pulumi:"destinationPortDesc"`
	// Fragments option
	Fragments pulumi.BoolOutput `pulumi:"fragments"`
	// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
	Ip pulumi.StringOutput `pulumi:"ip"`
	// IPv4 address (e.g., 192.0.2.0)
	IpOnFirewall pulumi.StringOutput `pulumi:"ipOnFirewall"`
	// Possible values for protocol
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	Rule     pulumi.StringOutput `pulumi:"rule"`
	// Possible values for action
	Sequence pulumi.Float64Output `pulumi:"sequence"`
	// IPv4 CIDR notation (e.g., 192.0.2.0/24)
	Source pulumi.StringOutput `pulumi:"source"`
	// Source port for your rule. Only with TCP/UDP protocol
	SourcePort pulumi.Float64Output `pulumi:"sourcePort"`
	// Source port for your rule. Only with TCP/UDP protocol
	SourcePortDesc pulumi.StringOutput `pulumi:"sourcePortDesc"`
	// Current state of your rule
	State pulumi.StringOutput `pulumi:"state"`
	// TCP option on your rule
	TcpOption pulumi.StringOutput `pulumi:"tcpOption"`
}

// NewIpFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewIpFirewallRule(ctx *pulumi.Context,
	name string, args *IpFirewallRuleArgs, opts ...pulumi.ResourceOption) (*IpFirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	if args.IpOnFirewall == nil {
		return nil, errors.New("invalid value for required argument 'IpOnFirewall'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.Sequence == nil {
		return nil, errors.New("invalid value for required argument 'Sequence'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource IpFirewallRule
	err = ctx.RegisterPackageResource("ovh:index/ipFirewallRule:IpFirewallRule", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpFirewallRule gets an existing IpFirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpFirewallRuleState, opts ...pulumi.ResourceOption) (*IpFirewallRule, error) {
	var resource IpFirewallRule
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("ovh:index/ipFirewallRule:IpFirewallRule", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpFirewallRule resources.
type ipFirewallRuleState struct {
	// Possible values for action
	Action       *string `pulumi:"action"`
	CreationDate *string `pulumi:"creationDate"`
	// Destination ip for your rule
	Destination *string `pulumi:"destination"`
	// Destination port for your rule. Only with TCP/UDP protocol
	DestinationPort *float64 `pulumi:"destinationPort"`
	// Destination port range for your rule. Only with TCP/UDP protocol
	DestinationPortDesc *string `pulumi:"destinationPortDesc"`
	// Fragments option
	Fragments *bool `pulumi:"fragments"`
	// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
	Ip *string `pulumi:"ip"`
	// IPv4 address (e.g., 192.0.2.0)
	IpOnFirewall *string `pulumi:"ipOnFirewall"`
	// Possible values for protocol
	Protocol *string `pulumi:"protocol"`
	Rule     *string `pulumi:"rule"`
	// Possible values for action
	Sequence *float64 `pulumi:"sequence"`
	// IPv4 CIDR notation (e.g., 192.0.2.0/24)
	Source *string `pulumi:"source"`
	// Source port for your rule. Only with TCP/UDP protocol
	SourcePort *float64 `pulumi:"sourcePort"`
	// Source port for your rule. Only with TCP/UDP protocol
	SourcePortDesc *string `pulumi:"sourcePortDesc"`
	// Current state of your rule
	State *string `pulumi:"state"`
	// TCP option on your rule
	TcpOption *string `pulumi:"tcpOption"`
}

type IpFirewallRuleState struct {
	// Possible values for action
	Action       pulumi.StringPtrInput
	CreationDate pulumi.StringPtrInput
	// Destination ip for your rule
	Destination pulumi.StringPtrInput
	// Destination port for your rule. Only with TCP/UDP protocol
	DestinationPort pulumi.Float64PtrInput
	// Destination port range for your rule. Only with TCP/UDP protocol
	DestinationPortDesc pulumi.StringPtrInput
	// Fragments option
	Fragments pulumi.BoolPtrInput
	// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
	Ip pulumi.StringPtrInput
	// IPv4 address (e.g., 192.0.2.0)
	IpOnFirewall pulumi.StringPtrInput
	// Possible values for protocol
	Protocol pulumi.StringPtrInput
	Rule     pulumi.StringPtrInput
	// Possible values for action
	Sequence pulumi.Float64PtrInput
	// IPv4 CIDR notation (e.g., 192.0.2.0/24)
	Source pulumi.StringPtrInput
	// Source port for your rule. Only with TCP/UDP protocol
	SourcePort pulumi.Float64PtrInput
	// Source port for your rule. Only with TCP/UDP protocol
	SourcePortDesc pulumi.StringPtrInput
	// Current state of your rule
	State pulumi.StringPtrInput
	// TCP option on your rule
	TcpOption pulumi.StringPtrInput
}

func (IpFirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipFirewallRuleState)(nil)).Elem()
}

type ipFirewallRuleArgs struct {
	// Possible values for action
	Action string `pulumi:"action"`
	// Destination port for your rule. Only with TCP/UDP protocol
	DestinationPort *float64 `pulumi:"destinationPort"`
	// Fragments option
	Fragments *bool `pulumi:"fragments"`
	// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
	Ip string `pulumi:"ip"`
	// IPv4 address (e.g., 192.0.2.0)
	IpOnFirewall string `pulumi:"ipOnFirewall"`
	// Possible values for protocol
	Protocol string `pulumi:"protocol"`
	// Possible values for action
	Sequence float64 `pulumi:"sequence"`
	// IPv4 CIDR notation (e.g., 192.0.2.0/24)
	Source *string `pulumi:"source"`
	// Source port for your rule. Only with TCP/UDP protocol
	SourcePort *float64 `pulumi:"sourcePort"`
	// TCP option on your rule
	TcpOption *string `pulumi:"tcpOption"`
}

// The set of arguments for constructing a IpFirewallRule resource.
type IpFirewallRuleArgs struct {
	// Possible values for action
	Action pulumi.StringInput
	// Destination port for your rule. Only with TCP/UDP protocol
	DestinationPort pulumi.Float64PtrInput
	// Fragments option
	Fragments pulumi.BoolPtrInput
	// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
	Ip pulumi.StringInput
	// IPv4 address (e.g., 192.0.2.0)
	IpOnFirewall pulumi.StringInput
	// Possible values for protocol
	Protocol pulumi.StringInput
	// Possible values for action
	Sequence pulumi.Float64Input
	// IPv4 CIDR notation (e.g., 192.0.2.0/24)
	Source pulumi.StringPtrInput
	// Source port for your rule. Only with TCP/UDP protocol
	SourcePort pulumi.Float64PtrInput
	// TCP option on your rule
	TcpOption pulumi.StringPtrInput
}

func (IpFirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipFirewallRuleArgs)(nil)).Elem()
}

type IpFirewallRuleInput interface {
	pulumi.Input

	ToIpFirewallRuleOutput() IpFirewallRuleOutput
	ToIpFirewallRuleOutputWithContext(ctx context.Context) IpFirewallRuleOutput
}

func (*IpFirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**IpFirewallRule)(nil)).Elem()
}

func (i *IpFirewallRule) ToIpFirewallRuleOutput() IpFirewallRuleOutput {
	return i.ToIpFirewallRuleOutputWithContext(context.Background())
}

func (i *IpFirewallRule) ToIpFirewallRuleOutputWithContext(ctx context.Context) IpFirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpFirewallRuleOutput)
}

type IpFirewallRuleOutput struct{ *pulumi.OutputState }

func (IpFirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpFirewallRule)(nil)).Elem()
}

func (o IpFirewallRuleOutput) ToIpFirewallRuleOutput() IpFirewallRuleOutput {
	return o
}

func (o IpFirewallRuleOutput) ToIpFirewallRuleOutputWithContext(ctx context.Context) IpFirewallRuleOutput {
	return o
}

// Possible values for action
func (o IpFirewallRuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

func (o IpFirewallRuleOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// Destination ip for your rule
func (o IpFirewallRuleOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringOutput { return v.Destination }).(pulumi.StringOutput)
}

// Destination port for your rule. Only with TCP/UDP protocol
func (o IpFirewallRuleOutput) DestinationPort() pulumi.Float64Output {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.Float64Output { return v.DestinationPort }).(pulumi.Float64Output)
}

// Destination port range for your rule. Only with TCP/UDP protocol
func (o IpFirewallRuleOutput) DestinationPortDesc() pulumi.StringOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringOutput { return v.DestinationPortDesc }).(pulumi.StringOutput)
}

// Fragments option
func (o IpFirewallRuleOutput) Fragments() pulumi.BoolOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.BoolOutput { return v.Fragments }).(pulumi.BoolOutput)
}

// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
func (o IpFirewallRuleOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// IPv4 address (e.g., 192.0.2.0)
func (o IpFirewallRuleOutput) IpOnFirewall() pulumi.StringOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringOutput { return v.IpOnFirewall }).(pulumi.StringOutput)
}

// Possible values for protocol
func (o IpFirewallRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

func (o IpFirewallRuleOutput) Rule() pulumi.StringOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringOutput { return v.Rule }).(pulumi.StringOutput)
}

// Possible values for action
func (o IpFirewallRuleOutput) Sequence() pulumi.Float64Output {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.Float64Output { return v.Sequence }).(pulumi.Float64Output)
}

// IPv4 CIDR notation (e.g., 192.0.2.0/24)
func (o IpFirewallRuleOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// Source port for your rule. Only with TCP/UDP protocol
func (o IpFirewallRuleOutput) SourcePort() pulumi.Float64Output {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.Float64Output { return v.SourcePort }).(pulumi.Float64Output)
}

// Source port for your rule. Only with TCP/UDP protocol
func (o IpFirewallRuleOutput) SourcePortDesc() pulumi.StringOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringOutput { return v.SourcePortDesc }).(pulumi.StringOutput)
}

// Current state of your rule
func (o IpFirewallRuleOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// TCP option on your rule
func (o IpFirewallRuleOutput) TcpOption() pulumi.StringOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringOutput { return v.TcpOption }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpFirewallRuleInput)(nil)).Elem(), &IpFirewallRule{})
	pulumi.RegisterOutputType(IpFirewallRuleOutput{})
}