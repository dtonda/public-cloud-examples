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

type CloudProjectContainerregistry struct {
	pulumi.CustomResourceState

	// Registry creation date
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Registry name
	Name pulumi.StringOutput `pulumi:"name"`
	// Plan ID of the registry.
	PlanId pulumi.StringOutput `pulumi:"planId"`
	// Plan of the registry
	Plans CloudProjectContainerregistryPlanArrayOutput `pulumi:"plans"`
	// Project ID of your registry
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Region of the registry.
	Region      pulumi.StringOutput    `pulumi:"region"`
	ServiceName pulumi.StringPtrOutput `pulumi:"serviceName"`
	// Current size of the registry (bytes)
	Size pulumi.Float64Output `pulumi:"size"`
	// Registry status
	Status pulumi.StringOutput `pulumi:"status"`
	// Registry last update date
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Access url of the registry
	Url pulumi.StringOutput `pulumi:"url"`
	// Version of your registry
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewCloudProjectContainerregistry registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectContainerregistry(ctx *pulumi.Context,
	name string, args *CloudProjectContainerregistryArgs, opts ...pulumi.ResourceOption) (*CloudProjectContainerregistry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource CloudProjectContainerregistry
	err = ctx.RegisterPackageResource("ovh:index/cloudProjectContainerregistry:CloudProjectContainerregistry", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectContainerregistry gets an existing CloudProjectContainerregistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectContainerregistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectContainerregistryState, opts ...pulumi.ResourceOption) (*CloudProjectContainerregistry, error) {
	var resource CloudProjectContainerregistry
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("ovh:index/cloudProjectContainerregistry:CloudProjectContainerregistry", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectContainerregistry resources.
type cloudProjectContainerregistryState struct {
	// Registry creation date
	CreatedAt *string `pulumi:"createdAt"`
	// Registry name
	Name *string `pulumi:"name"`
	// Plan ID of the registry.
	PlanId *string `pulumi:"planId"`
	// Plan of the registry
	Plans []CloudProjectContainerregistryPlan `pulumi:"plans"`
	// Project ID of your registry
	ProjectId *string `pulumi:"projectId"`
	// Region of the registry.
	Region      *string `pulumi:"region"`
	ServiceName *string `pulumi:"serviceName"`
	// Current size of the registry (bytes)
	Size *float64 `pulumi:"size"`
	// Registry status
	Status *string `pulumi:"status"`
	// Registry last update date
	UpdatedAt *string `pulumi:"updatedAt"`
	// Access url of the registry
	Url *string `pulumi:"url"`
	// Version of your registry
	Version *string `pulumi:"version"`
}

type CloudProjectContainerregistryState struct {
	// Registry creation date
	CreatedAt pulumi.StringPtrInput
	// Registry name
	Name pulumi.StringPtrInput
	// Plan ID of the registry.
	PlanId pulumi.StringPtrInput
	// Plan of the registry
	Plans CloudProjectContainerregistryPlanArrayInput
	// Project ID of your registry
	ProjectId pulumi.StringPtrInput
	// Region of the registry.
	Region      pulumi.StringPtrInput
	ServiceName pulumi.StringPtrInput
	// Current size of the registry (bytes)
	Size pulumi.Float64PtrInput
	// Registry status
	Status pulumi.StringPtrInput
	// Registry last update date
	UpdatedAt pulumi.StringPtrInput
	// Access url of the registry
	Url pulumi.StringPtrInput
	// Version of your registry
	Version pulumi.StringPtrInput
}

func (CloudProjectContainerregistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectContainerregistryState)(nil)).Elem()
}

type cloudProjectContainerregistryArgs struct {
	// Registry name
	Name *string `pulumi:"name"`
	// Plan ID of the registry.
	PlanId *string `pulumi:"planId"`
	// Region of the registry.
	Region      string  `pulumi:"region"`
	ServiceName *string `pulumi:"serviceName"`
}

// The set of arguments for constructing a CloudProjectContainerregistry resource.
type CloudProjectContainerregistryArgs struct {
	// Registry name
	Name pulumi.StringPtrInput
	// Plan ID of the registry.
	PlanId pulumi.StringPtrInput
	// Region of the registry.
	Region      pulumi.StringInput
	ServiceName pulumi.StringPtrInput
}

func (CloudProjectContainerregistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectContainerregistryArgs)(nil)).Elem()
}

type CloudProjectContainerregistryInput interface {
	pulumi.Input

	ToCloudProjectContainerregistryOutput() CloudProjectContainerregistryOutput
	ToCloudProjectContainerregistryOutputWithContext(ctx context.Context) CloudProjectContainerregistryOutput
}

func (*CloudProjectContainerregistry) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectContainerregistry)(nil)).Elem()
}

func (i *CloudProjectContainerregistry) ToCloudProjectContainerregistryOutput() CloudProjectContainerregistryOutput {
	return i.ToCloudProjectContainerregistryOutputWithContext(context.Background())
}

func (i *CloudProjectContainerregistry) ToCloudProjectContainerregistryOutputWithContext(ctx context.Context) CloudProjectContainerregistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectContainerregistryOutput)
}

type CloudProjectContainerregistryOutput struct{ *pulumi.OutputState }

func (CloudProjectContainerregistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectContainerregistry)(nil)).Elem()
}

func (o CloudProjectContainerregistryOutput) ToCloudProjectContainerregistryOutput() CloudProjectContainerregistryOutput {
	return o
}

func (o CloudProjectContainerregistryOutput) ToCloudProjectContainerregistryOutputWithContext(ctx context.Context) CloudProjectContainerregistryOutput {
	return o
}

// Registry creation date
func (o CloudProjectContainerregistryOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerregistry) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Registry name
func (o CloudProjectContainerregistryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerregistry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Plan ID of the registry.
func (o CloudProjectContainerregistryOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerregistry) pulumi.StringOutput { return v.PlanId }).(pulumi.StringOutput)
}

// Plan of the registry
func (o CloudProjectContainerregistryOutput) Plans() CloudProjectContainerregistryPlanArrayOutput {
	return o.ApplyT(func(v *CloudProjectContainerregistry) CloudProjectContainerregistryPlanArrayOutput { return v.Plans }).(CloudProjectContainerregistryPlanArrayOutput)
}

// Project ID of your registry
func (o CloudProjectContainerregistryOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerregistry) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Region of the registry.
func (o CloudProjectContainerregistryOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerregistry) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o CloudProjectContainerregistryOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudProjectContainerregistry) pulumi.StringPtrOutput { return v.ServiceName }).(pulumi.StringPtrOutput)
}

// Current size of the registry (bytes)
func (o CloudProjectContainerregistryOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v *CloudProjectContainerregistry) pulumi.Float64Output { return v.Size }).(pulumi.Float64Output)
}

// Registry status
func (o CloudProjectContainerregistryOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerregistry) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Registry last update date
func (o CloudProjectContainerregistryOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerregistry) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Access url of the registry
func (o CloudProjectContainerregistryOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerregistry) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Version of your registry
func (o CloudProjectContainerregistryOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerregistry) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectContainerregistryInput)(nil)).Elem(), &CloudProjectContainerregistry{})
	pulumi.RegisterOutputType(CloudProjectContainerregistryOutput{})
}
