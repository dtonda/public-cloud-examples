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

type CloudProjectAlerting struct {
	pulumi.CustomResourceState

	// Alerting creation date
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// Possible values for delay between two alerts in seconds
	Delay pulumi.Float64Output `pulumi:"delay"`
	// Email to contact
	Email pulumi.StringOutput `pulumi:"email"`
	// Formatted monthly threshold for this alerting
	FormattedMonthlyThreshold CloudProjectAlertingFormattedMonthlyThresholdOutput `pulumi:"formattedMonthlyThreshold"`
	// Monthly threshold for this alerting in currency
	MonthlyThreshold pulumi.Float64Output `pulumi:"monthlyThreshold"`
	// The project id
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewCloudProjectAlerting registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectAlerting(ctx *pulumi.Context,
	name string, args *CloudProjectAlertingArgs, opts ...pulumi.ResourceOption) (*CloudProjectAlerting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Delay == nil {
		return nil, errors.New("invalid value for required argument 'Delay'")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.MonthlyThreshold == nil {
		return nil, errors.New("invalid value for required argument 'MonthlyThreshold'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource CloudProjectAlerting
	err = ctx.RegisterPackageResource("ovh:index/cloudProjectAlerting:CloudProjectAlerting", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectAlerting gets an existing CloudProjectAlerting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectAlerting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectAlertingState, opts ...pulumi.ResourceOption) (*CloudProjectAlerting, error) {
	var resource CloudProjectAlerting
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("ovh:index/cloudProjectAlerting:CloudProjectAlerting", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectAlerting resources.
type cloudProjectAlertingState struct {
	// Alerting creation date
	CreationDate *string `pulumi:"creationDate"`
	// Possible values for delay between two alerts in seconds
	Delay *float64 `pulumi:"delay"`
	// Email to contact
	Email *string `pulumi:"email"`
	// Formatted monthly threshold for this alerting
	FormattedMonthlyThreshold *CloudProjectAlertingFormattedMonthlyThreshold `pulumi:"formattedMonthlyThreshold"`
	// Monthly threshold for this alerting in currency
	MonthlyThreshold *float64 `pulumi:"monthlyThreshold"`
	// The project id
	ServiceName *string `pulumi:"serviceName"`
}

type CloudProjectAlertingState struct {
	// Alerting creation date
	CreationDate pulumi.StringPtrInput
	// Possible values for delay between two alerts in seconds
	Delay pulumi.Float64PtrInput
	// Email to contact
	Email pulumi.StringPtrInput
	// Formatted monthly threshold for this alerting
	FormattedMonthlyThreshold CloudProjectAlertingFormattedMonthlyThresholdPtrInput
	// Monthly threshold for this alerting in currency
	MonthlyThreshold pulumi.Float64PtrInput
	// The project id
	ServiceName pulumi.StringPtrInput
}

func (CloudProjectAlertingState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectAlertingState)(nil)).Elem()
}

type cloudProjectAlertingArgs struct {
	// Possible values for delay between two alerts in seconds
	Delay float64 `pulumi:"delay"`
	// Email to contact
	Email string `pulumi:"email"`
	// Monthly threshold for this alerting in currency
	MonthlyThreshold float64 `pulumi:"monthlyThreshold"`
	// The project id
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a CloudProjectAlerting resource.
type CloudProjectAlertingArgs struct {
	// Possible values for delay between two alerts in seconds
	Delay pulumi.Float64Input
	// Email to contact
	Email pulumi.StringInput
	// Monthly threshold for this alerting in currency
	MonthlyThreshold pulumi.Float64Input
	// The project id
	ServiceName pulumi.StringInput
}

func (CloudProjectAlertingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectAlertingArgs)(nil)).Elem()
}

type CloudProjectAlertingInput interface {
	pulumi.Input

	ToCloudProjectAlertingOutput() CloudProjectAlertingOutput
	ToCloudProjectAlertingOutputWithContext(ctx context.Context) CloudProjectAlertingOutput
}

func (*CloudProjectAlerting) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectAlerting)(nil)).Elem()
}

func (i *CloudProjectAlerting) ToCloudProjectAlertingOutput() CloudProjectAlertingOutput {
	return i.ToCloudProjectAlertingOutputWithContext(context.Background())
}

func (i *CloudProjectAlerting) ToCloudProjectAlertingOutputWithContext(ctx context.Context) CloudProjectAlertingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectAlertingOutput)
}

type CloudProjectAlertingOutput struct{ *pulumi.OutputState }

func (CloudProjectAlertingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectAlerting)(nil)).Elem()
}

func (o CloudProjectAlertingOutput) ToCloudProjectAlertingOutput() CloudProjectAlertingOutput {
	return o
}

func (o CloudProjectAlertingOutput) ToCloudProjectAlertingOutputWithContext(ctx context.Context) CloudProjectAlertingOutput {
	return o
}

// Alerting creation date
func (o CloudProjectAlertingOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectAlerting) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// Possible values for delay between two alerts in seconds
func (o CloudProjectAlertingOutput) Delay() pulumi.Float64Output {
	return o.ApplyT(func(v *CloudProjectAlerting) pulumi.Float64Output { return v.Delay }).(pulumi.Float64Output)
}

// Email to contact
func (o CloudProjectAlertingOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectAlerting) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// Formatted monthly threshold for this alerting
func (o CloudProjectAlertingOutput) FormattedMonthlyThreshold() CloudProjectAlertingFormattedMonthlyThresholdOutput {
	return o.ApplyT(func(v *CloudProjectAlerting) CloudProjectAlertingFormattedMonthlyThresholdOutput {
		return v.FormattedMonthlyThreshold
	}).(CloudProjectAlertingFormattedMonthlyThresholdOutput)
}

// Monthly threshold for this alerting in currency
func (o CloudProjectAlertingOutput) MonthlyThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v *CloudProjectAlerting) pulumi.Float64Output { return v.MonthlyThreshold }).(pulumi.Float64Output)
}

// The project id
func (o CloudProjectAlertingOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectAlerting) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectAlertingInput)(nil)).Elem(), &CloudProjectAlerting{})
	pulumi.RegisterOutputType(CloudProjectAlertingOutput{})
}
