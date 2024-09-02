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

type CloudProjectKubeNodepool struct {
	pulumi.CustomResourceState

	// Enable anti affinity groups for nodes in the pool
	AntiAffinity pulumi.BoolOutput `pulumi:"antiAffinity"`
	// Enable auto-scaling for the pool
	Autoscale pulumi.BoolOutput `pulumi:"autoscale"`
	// scaleDownUnneededTimeSeconds for autoscaling
	AutoscalingScaleDownUnneededTimeSeconds pulumi.Float64Output `pulumi:"autoscalingScaleDownUnneededTimeSeconds"`
	// scaleDownUnreadyTimeSeconds for autoscaling
	AutoscalingScaleDownUnreadyTimeSeconds pulumi.Float64Output `pulumi:"autoscalingScaleDownUnreadyTimeSeconds"`
	// scaleDownUtilizationThreshold for autoscaling
	AutoscalingScaleDownUtilizationThreshold pulumi.Float64Output `pulumi:"autoscalingScaleDownUtilizationThreshold"`
	// Number of nodes which are actually ready in the pool
	AvailableNodes pulumi.Float64Output `pulumi:"availableNodes"`
	// Creation date
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Number of nodes present in the pool
	CurrentNodes pulumi.Float64Output `pulumi:"currentNodes"`
	// Number of nodes you desire in the pool
	DesiredNodes pulumi.Float64Output `pulumi:"desiredNodes"`
	// Flavor name
	Flavor pulumi.StringOutput `pulumi:"flavor"`
	// Flavor name
	FlavorName pulumi.StringOutput `pulumi:"flavorName"`
	// Kube ID
	KubeId pulumi.StringOutput `pulumi:"kubeId"`
	// Number of nodes you desire in the pool
	MaxNodes pulumi.Float64Output `pulumi:"maxNodes"`
	// Number of nodes you desire in the pool
	MinNodes pulumi.Float64Output `pulumi:"minNodes"`
	// Enable monthly billing on all nodes in the pool
	MonthlyBilled pulumi.BoolOutput `pulumi:"monthlyBilled"`
	// NodePool resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Project id
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Service name
	ServiceName pulumi.StringPtrOutput `pulumi:"serviceName"`
	// Status describing the state between number of nodes wanted and available ones
	SizeStatus pulumi.StringOutput `pulumi:"sizeStatus"`
	// Current status
	Status pulumi.StringOutput `pulumi:"status"`
	// Node pool template
	Template CloudProjectKubeNodepoolTemplatePtrOutput `pulumi:"template"`
	Timeouts CloudProjectKubeNodepoolTimeoutsPtrOutput `pulumi:"timeouts"`
	// Number of nodes with latest version installed in the pool
	UpToDateNodes pulumi.Float64Output `pulumi:"upToDateNodes"`
	// Last update date
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewCloudProjectKubeNodepool registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectKubeNodepool(ctx *pulumi.Context,
	name string, args *CloudProjectKubeNodepoolArgs, opts ...pulumi.ResourceOption) (*CloudProjectKubeNodepool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FlavorName == nil {
		return nil, errors.New("invalid value for required argument 'FlavorName'")
	}
	if args.KubeId == nil {
		return nil, errors.New("invalid value for required argument 'KubeId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource CloudProjectKubeNodepool
	err = ctx.RegisterPackageResource("ovh:index/cloudProjectKubeNodepool:CloudProjectKubeNodepool", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectKubeNodepool gets an existing CloudProjectKubeNodepool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectKubeNodepool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectKubeNodepoolState, opts ...pulumi.ResourceOption) (*CloudProjectKubeNodepool, error) {
	var resource CloudProjectKubeNodepool
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("ovh:index/cloudProjectKubeNodepool:CloudProjectKubeNodepool", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectKubeNodepool resources.
type cloudProjectKubeNodepoolState struct {
	// Enable anti affinity groups for nodes in the pool
	AntiAffinity *bool `pulumi:"antiAffinity"`
	// Enable auto-scaling for the pool
	Autoscale *bool `pulumi:"autoscale"`
	// scaleDownUnneededTimeSeconds for autoscaling
	AutoscalingScaleDownUnneededTimeSeconds *float64 `pulumi:"autoscalingScaleDownUnneededTimeSeconds"`
	// scaleDownUnreadyTimeSeconds for autoscaling
	AutoscalingScaleDownUnreadyTimeSeconds *float64 `pulumi:"autoscalingScaleDownUnreadyTimeSeconds"`
	// scaleDownUtilizationThreshold for autoscaling
	AutoscalingScaleDownUtilizationThreshold *float64 `pulumi:"autoscalingScaleDownUtilizationThreshold"`
	// Number of nodes which are actually ready in the pool
	AvailableNodes *float64 `pulumi:"availableNodes"`
	// Creation date
	CreatedAt *string `pulumi:"createdAt"`
	// Number of nodes present in the pool
	CurrentNodes *float64 `pulumi:"currentNodes"`
	// Number of nodes you desire in the pool
	DesiredNodes *float64 `pulumi:"desiredNodes"`
	// Flavor name
	Flavor *string `pulumi:"flavor"`
	// Flavor name
	FlavorName *string `pulumi:"flavorName"`
	// Kube ID
	KubeId *string `pulumi:"kubeId"`
	// Number of nodes you desire in the pool
	MaxNodes *float64 `pulumi:"maxNodes"`
	// Number of nodes you desire in the pool
	MinNodes *float64 `pulumi:"minNodes"`
	// Enable monthly billing on all nodes in the pool
	MonthlyBilled *bool `pulumi:"monthlyBilled"`
	// NodePool resource name
	Name *string `pulumi:"name"`
	// Project id
	ProjectId *string `pulumi:"projectId"`
	// Service name
	ServiceName *string `pulumi:"serviceName"`
	// Status describing the state between number of nodes wanted and available ones
	SizeStatus *string `pulumi:"sizeStatus"`
	// Current status
	Status *string `pulumi:"status"`
	// Node pool template
	Template *CloudProjectKubeNodepoolTemplate `pulumi:"template"`
	Timeouts *CloudProjectKubeNodepoolTimeouts `pulumi:"timeouts"`
	// Number of nodes with latest version installed in the pool
	UpToDateNodes *float64 `pulumi:"upToDateNodes"`
	// Last update date
	UpdatedAt *string `pulumi:"updatedAt"`
}

type CloudProjectKubeNodepoolState struct {
	// Enable anti affinity groups for nodes in the pool
	AntiAffinity pulumi.BoolPtrInput
	// Enable auto-scaling for the pool
	Autoscale pulumi.BoolPtrInput
	// scaleDownUnneededTimeSeconds for autoscaling
	AutoscalingScaleDownUnneededTimeSeconds pulumi.Float64PtrInput
	// scaleDownUnreadyTimeSeconds for autoscaling
	AutoscalingScaleDownUnreadyTimeSeconds pulumi.Float64PtrInput
	// scaleDownUtilizationThreshold for autoscaling
	AutoscalingScaleDownUtilizationThreshold pulumi.Float64PtrInput
	// Number of nodes which are actually ready in the pool
	AvailableNodes pulumi.Float64PtrInput
	// Creation date
	CreatedAt pulumi.StringPtrInput
	// Number of nodes present in the pool
	CurrentNodes pulumi.Float64PtrInput
	// Number of nodes you desire in the pool
	DesiredNodes pulumi.Float64PtrInput
	// Flavor name
	Flavor pulumi.StringPtrInput
	// Flavor name
	FlavorName pulumi.StringPtrInput
	// Kube ID
	KubeId pulumi.StringPtrInput
	// Number of nodes you desire in the pool
	MaxNodes pulumi.Float64PtrInput
	// Number of nodes you desire in the pool
	MinNodes pulumi.Float64PtrInput
	// Enable monthly billing on all nodes in the pool
	MonthlyBilled pulumi.BoolPtrInput
	// NodePool resource name
	Name pulumi.StringPtrInput
	// Project id
	ProjectId pulumi.StringPtrInput
	// Service name
	ServiceName pulumi.StringPtrInput
	// Status describing the state between number of nodes wanted and available ones
	SizeStatus pulumi.StringPtrInput
	// Current status
	Status pulumi.StringPtrInput
	// Node pool template
	Template CloudProjectKubeNodepoolTemplatePtrInput
	Timeouts CloudProjectKubeNodepoolTimeoutsPtrInput
	// Number of nodes with latest version installed in the pool
	UpToDateNodes pulumi.Float64PtrInput
	// Last update date
	UpdatedAt pulumi.StringPtrInput
}

func (CloudProjectKubeNodepoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectKubeNodepoolState)(nil)).Elem()
}

type cloudProjectKubeNodepoolArgs struct {
	// Enable anti affinity groups for nodes in the pool
	AntiAffinity *bool `pulumi:"antiAffinity"`
	// Enable auto-scaling for the pool
	Autoscale *bool `pulumi:"autoscale"`
	// scaleDownUnneededTimeSeconds for autoscaling
	AutoscalingScaleDownUnneededTimeSeconds *float64 `pulumi:"autoscalingScaleDownUnneededTimeSeconds"`
	// scaleDownUnreadyTimeSeconds for autoscaling
	AutoscalingScaleDownUnreadyTimeSeconds *float64 `pulumi:"autoscalingScaleDownUnreadyTimeSeconds"`
	// scaleDownUtilizationThreshold for autoscaling
	AutoscalingScaleDownUtilizationThreshold *float64 `pulumi:"autoscalingScaleDownUtilizationThreshold"`
	// Number of nodes you desire in the pool
	DesiredNodes *float64 `pulumi:"desiredNodes"`
	// Flavor name
	FlavorName string `pulumi:"flavorName"`
	// Kube ID
	KubeId string `pulumi:"kubeId"`
	// Number of nodes you desire in the pool
	MaxNodes *float64 `pulumi:"maxNodes"`
	// Number of nodes you desire in the pool
	MinNodes *float64 `pulumi:"minNodes"`
	// Enable monthly billing on all nodes in the pool
	MonthlyBilled *bool `pulumi:"monthlyBilled"`
	// NodePool resource name
	Name *string `pulumi:"name"`
	// Service name
	ServiceName *string `pulumi:"serviceName"`
	// Node pool template
	Template *CloudProjectKubeNodepoolTemplate `pulumi:"template"`
	Timeouts *CloudProjectKubeNodepoolTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a CloudProjectKubeNodepool resource.
type CloudProjectKubeNodepoolArgs struct {
	// Enable anti affinity groups for nodes in the pool
	AntiAffinity pulumi.BoolPtrInput
	// Enable auto-scaling for the pool
	Autoscale pulumi.BoolPtrInput
	// scaleDownUnneededTimeSeconds for autoscaling
	AutoscalingScaleDownUnneededTimeSeconds pulumi.Float64PtrInput
	// scaleDownUnreadyTimeSeconds for autoscaling
	AutoscalingScaleDownUnreadyTimeSeconds pulumi.Float64PtrInput
	// scaleDownUtilizationThreshold for autoscaling
	AutoscalingScaleDownUtilizationThreshold pulumi.Float64PtrInput
	// Number of nodes you desire in the pool
	DesiredNodes pulumi.Float64PtrInput
	// Flavor name
	FlavorName pulumi.StringInput
	// Kube ID
	KubeId pulumi.StringInput
	// Number of nodes you desire in the pool
	MaxNodes pulumi.Float64PtrInput
	// Number of nodes you desire in the pool
	MinNodes pulumi.Float64PtrInput
	// Enable monthly billing on all nodes in the pool
	MonthlyBilled pulumi.BoolPtrInput
	// NodePool resource name
	Name pulumi.StringPtrInput
	// Service name
	ServiceName pulumi.StringPtrInput
	// Node pool template
	Template CloudProjectKubeNodepoolTemplatePtrInput
	Timeouts CloudProjectKubeNodepoolTimeoutsPtrInput
}

func (CloudProjectKubeNodepoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectKubeNodepoolArgs)(nil)).Elem()
}

type CloudProjectKubeNodepoolInput interface {
	pulumi.Input

	ToCloudProjectKubeNodepoolOutput() CloudProjectKubeNodepoolOutput
	ToCloudProjectKubeNodepoolOutputWithContext(ctx context.Context) CloudProjectKubeNodepoolOutput
}

func (*CloudProjectKubeNodepool) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectKubeNodepool)(nil)).Elem()
}

func (i *CloudProjectKubeNodepool) ToCloudProjectKubeNodepoolOutput() CloudProjectKubeNodepoolOutput {
	return i.ToCloudProjectKubeNodepoolOutputWithContext(context.Background())
}

func (i *CloudProjectKubeNodepool) ToCloudProjectKubeNodepoolOutputWithContext(ctx context.Context) CloudProjectKubeNodepoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectKubeNodepoolOutput)
}

type CloudProjectKubeNodepoolOutput struct{ *pulumi.OutputState }

func (CloudProjectKubeNodepoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectKubeNodepool)(nil)).Elem()
}

func (o CloudProjectKubeNodepoolOutput) ToCloudProjectKubeNodepoolOutput() CloudProjectKubeNodepoolOutput {
	return o
}

func (o CloudProjectKubeNodepoolOutput) ToCloudProjectKubeNodepoolOutputWithContext(ctx context.Context) CloudProjectKubeNodepoolOutput {
	return o
}

// Enable anti affinity groups for nodes in the pool
func (o CloudProjectKubeNodepoolOutput) AntiAffinity() pulumi.BoolOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.BoolOutput { return v.AntiAffinity }).(pulumi.BoolOutput)
}

// Enable auto-scaling for the pool
func (o CloudProjectKubeNodepoolOutput) Autoscale() pulumi.BoolOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.BoolOutput { return v.Autoscale }).(pulumi.BoolOutput)
}

// scaleDownUnneededTimeSeconds for autoscaling
func (o CloudProjectKubeNodepoolOutput) AutoscalingScaleDownUnneededTimeSeconds() pulumi.Float64Output {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.Float64Output {
		return v.AutoscalingScaleDownUnneededTimeSeconds
	}).(pulumi.Float64Output)
}

// scaleDownUnreadyTimeSeconds for autoscaling
func (o CloudProjectKubeNodepoolOutput) AutoscalingScaleDownUnreadyTimeSeconds() pulumi.Float64Output {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.Float64Output {
		return v.AutoscalingScaleDownUnreadyTimeSeconds
	}).(pulumi.Float64Output)
}

// scaleDownUtilizationThreshold for autoscaling
func (o CloudProjectKubeNodepoolOutput) AutoscalingScaleDownUtilizationThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.Float64Output {
		return v.AutoscalingScaleDownUtilizationThreshold
	}).(pulumi.Float64Output)
}

// Number of nodes which are actually ready in the pool
func (o CloudProjectKubeNodepoolOutput) AvailableNodes() pulumi.Float64Output {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.Float64Output { return v.AvailableNodes }).(pulumi.Float64Output)
}

// Creation date
func (o CloudProjectKubeNodepoolOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Number of nodes present in the pool
func (o CloudProjectKubeNodepoolOutput) CurrentNodes() pulumi.Float64Output {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.Float64Output { return v.CurrentNodes }).(pulumi.Float64Output)
}

// Number of nodes you desire in the pool
func (o CloudProjectKubeNodepoolOutput) DesiredNodes() pulumi.Float64Output {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.Float64Output { return v.DesiredNodes }).(pulumi.Float64Output)
}

// Flavor name
func (o CloudProjectKubeNodepoolOutput) Flavor() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.StringOutput { return v.Flavor }).(pulumi.StringOutput)
}

// Flavor name
func (o CloudProjectKubeNodepoolOutput) FlavorName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.StringOutput { return v.FlavorName }).(pulumi.StringOutput)
}

// Kube ID
func (o CloudProjectKubeNodepoolOutput) KubeId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.StringOutput { return v.KubeId }).(pulumi.StringOutput)
}

// Number of nodes you desire in the pool
func (o CloudProjectKubeNodepoolOutput) MaxNodes() pulumi.Float64Output {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.Float64Output { return v.MaxNodes }).(pulumi.Float64Output)
}

// Number of nodes you desire in the pool
func (o CloudProjectKubeNodepoolOutput) MinNodes() pulumi.Float64Output {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.Float64Output { return v.MinNodes }).(pulumi.Float64Output)
}

// Enable monthly billing on all nodes in the pool
func (o CloudProjectKubeNodepoolOutput) MonthlyBilled() pulumi.BoolOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.BoolOutput { return v.MonthlyBilled }).(pulumi.BoolOutput)
}

// NodePool resource name
func (o CloudProjectKubeNodepoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Project id
func (o CloudProjectKubeNodepoolOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Service name
func (o CloudProjectKubeNodepoolOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.StringPtrOutput { return v.ServiceName }).(pulumi.StringPtrOutput)
}

// Status describing the state between number of nodes wanted and available ones
func (o CloudProjectKubeNodepoolOutput) SizeStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.StringOutput { return v.SizeStatus }).(pulumi.StringOutput)
}

// Current status
func (o CloudProjectKubeNodepoolOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Node pool template
func (o CloudProjectKubeNodepoolOutput) Template() CloudProjectKubeNodepoolTemplatePtrOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) CloudProjectKubeNodepoolTemplatePtrOutput { return v.Template }).(CloudProjectKubeNodepoolTemplatePtrOutput)
}

func (o CloudProjectKubeNodepoolOutput) Timeouts() CloudProjectKubeNodepoolTimeoutsPtrOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) CloudProjectKubeNodepoolTimeoutsPtrOutput { return v.Timeouts }).(CloudProjectKubeNodepoolTimeoutsPtrOutput)
}

// Number of nodes with latest version installed in the pool
func (o CloudProjectKubeNodepoolOutput) UpToDateNodes() pulumi.Float64Output {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.Float64Output { return v.UpToDateNodes }).(pulumi.Float64Output)
}

// Last update date
func (o CloudProjectKubeNodepoolOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectKubeNodepoolInput)(nil)).Elem(), &CloudProjectKubeNodepool{})
	pulumi.RegisterOutputType(CloudProjectKubeNodepoolOutput{})
}
