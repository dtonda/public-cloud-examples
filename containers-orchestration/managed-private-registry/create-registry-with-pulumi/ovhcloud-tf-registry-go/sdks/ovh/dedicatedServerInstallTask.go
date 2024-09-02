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

type DedicatedServerInstallTask struct {
	pulumi.CustomResourceState

	// If set, reboot the server on the specified boot id during destroy phase
	BootidOnDestroy pulumi.Float64PtrOutput `pulumi:"bootidOnDestroy"`
	// Details of this task
	Comment pulumi.StringOutput                        `pulumi:"comment"`
	Details DedicatedServerInstallTaskDetailsPtrOutput `pulumi:"details"`
	// Completion date
	DoneDate pulumi.StringOutput `pulumi:"doneDate"`
	// Function name
	Function pulumi.StringOutput `pulumi:"function"`
	// Last update
	LastUpdate pulumi.StringOutput `pulumi:"lastUpdate"`
	// Partition scheme name.
	PartitionSchemeName pulumi.StringPtrOutput `pulumi:"partitionSchemeName"`
	// The internal name of your dedicated server.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Task Creation date
	StartDate pulumi.StringOutput `pulumi:"startDate"`
	// Task status
	Status pulumi.StringOutput `pulumi:"status"`
	// Template name
	TemplateName  pulumi.StringOutput                               `pulumi:"templateName"`
	Timeouts      DedicatedServerInstallTaskTimeoutsPtrOutput       `pulumi:"timeouts"`
	UserMetadatas DedicatedServerInstallTaskUserMetadataArrayOutput `pulumi:"userMetadatas"`
}

// NewDedicatedServerInstallTask registers a new resource with the given unique name, arguments, and options.
func NewDedicatedServerInstallTask(ctx *pulumi.Context,
	name string, args *DedicatedServerInstallTaskArgs, opts ...pulumi.ResourceOption) (*DedicatedServerInstallTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DedicatedServerInstallTask
	err = ctx.RegisterPackageResource("ovh:index/dedicatedServerInstallTask:DedicatedServerInstallTask", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedServerInstallTask gets an existing DedicatedServerInstallTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedServerInstallTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedServerInstallTaskState, opts ...pulumi.ResourceOption) (*DedicatedServerInstallTask, error) {
	var resource DedicatedServerInstallTask
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("ovh:index/dedicatedServerInstallTask:DedicatedServerInstallTask", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedServerInstallTask resources.
type dedicatedServerInstallTaskState struct {
	// If set, reboot the server on the specified boot id during destroy phase
	BootidOnDestroy *float64 `pulumi:"bootidOnDestroy"`
	// Details of this task
	Comment *string                            `pulumi:"comment"`
	Details *DedicatedServerInstallTaskDetails `pulumi:"details"`
	// Completion date
	DoneDate *string `pulumi:"doneDate"`
	// Function name
	Function *string `pulumi:"function"`
	// Last update
	LastUpdate *string `pulumi:"lastUpdate"`
	// Partition scheme name.
	PartitionSchemeName *string `pulumi:"partitionSchemeName"`
	// The internal name of your dedicated server.
	ServiceName *string `pulumi:"serviceName"`
	// Task Creation date
	StartDate *string `pulumi:"startDate"`
	// Task status
	Status *string `pulumi:"status"`
	// Template name
	TemplateName  *string                                  `pulumi:"templateName"`
	Timeouts      *DedicatedServerInstallTaskTimeouts      `pulumi:"timeouts"`
	UserMetadatas []DedicatedServerInstallTaskUserMetadata `pulumi:"userMetadatas"`
}

type DedicatedServerInstallTaskState struct {
	// If set, reboot the server on the specified boot id during destroy phase
	BootidOnDestroy pulumi.Float64PtrInput
	// Details of this task
	Comment pulumi.StringPtrInput
	Details DedicatedServerInstallTaskDetailsPtrInput
	// Completion date
	DoneDate pulumi.StringPtrInput
	// Function name
	Function pulumi.StringPtrInput
	// Last update
	LastUpdate pulumi.StringPtrInput
	// Partition scheme name.
	PartitionSchemeName pulumi.StringPtrInput
	// The internal name of your dedicated server.
	ServiceName pulumi.StringPtrInput
	// Task Creation date
	StartDate pulumi.StringPtrInput
	// Task status
	Status pulumi.StringPtrInput
	// Template name
	TemplateName  pulumi.StringPtrInput
	Timeouts      DedicatedServerInstallTaskTimeoutsPtrInput
	UserMetadatas DedicatedServerInstallTaskUserMetadataArrayInput
}

func (DedicatedServerInstallTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedServerInstallTaskState)(nil)).Elem()
}

type dedicatedServerInstallTaskArgs struct {
	// If set, reboot the server on the specified boot id during destroy phase
	BootidOnDestroy *float64                           `pulumi:"bootidOnDestroy"`
	Details         *DedicatedServerInstallTaskDetails `pulumi:"details"`
	// Partition scheme name.
	PartitionSchemeName *string `pulumi:"partitionSchemeName"`
	// The internal name of your dedicated server.
	ServiceName string `pulumi:"serviceName"`
	// Template name
	TemplateName  string                                   `pulumi:"templateName"`
	Timeouts      *DedicatedServerInstallTaskTimeouts      `pulumi:"timeouts"`
	UserMetadatas []DedicatedServerInstallTaskUserMetadata `pulumi:"userMetadatas"`
}

// The set of arguments for constructing a DedicatedServerInstallTask resource.
type DedicatedServerInstallTaskArgs struct {
	// If set, reboot the server on the specified boot id during destroy phase
	BootidOnDestroy pulumi.Float64PtrInput
	Details         DedicatedServerInstallTaskDetailsPtrInput
	// Partition scheme name.
	PartitionSchemeName pulumi.StringPtrInput
	// The internal name of your dedicated server.
	ServiceName pulumi.StringInput
	// Template name
	TemplateName  pulumi.StringInput
	Timeouts      DedicatedServerInstallTaskTimeoutsPtrInput
	UserMetadatas DedicatedServerInstallTaskUserMetadataArrayInput
}

func (DedicatedServerInstallTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedServerInstallTaskArgs)(nil)).Elem()
}

type DedicatedServerInstallTaskInput interface {
	pulumi.Input

	ToDedicatedServerInstallTaskOutput() DedicatedServerInstallTaskOutput
	ToDedicatedServerInstallTaskOutputWithContext(ctx context.Context) DedicatedServerInstallTaskOutput
}

func (*DedicatedServerInstallTask) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedServerInstallTask)(nil)).Elem()
}

func (i *DedicatedServerInstallTask) ToDedicatedServerInstallTaskOutput() DedicatedServerInstallTaskOutput {
	return i.ToDedicatedServerInstallTaskOutputWithContext(context.Background())
}

func (i *DedicatedServerInstallTask) ToDedicatedServerInstallTaskOutputWithContext(ctx context.Context) DedicatedServerInstallTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedServerInstallTaskOutput)
}

type DedicatedServerInstallTaskOutput struct{ *pulumi.OutputState }

func (DedicatedServerInstallTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedServerInstallTask)(nil)).Elem()
}

func (o DedicatedServerInstallTaskOutput) ToDedicatedServerInstallTaskOutput() DedicatedServerInstallTaskOutput {
	return o
}

func (o DedicatedServerInstallTaskOutput) ToDedicatedServerInstallTaskOutputWithContext(ctx context.Context) DedicatedServerInstallTaskOutput {
	return o
}

// If set, reboot the server on the specified boot id during destroy phase
func (o DedicatedServerInstallTaskOutput) BootidOnDestroy() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DedicatedServerInstallTask) pulumi.Float64PtrOutput { return v.BootidOnDestroy }).(pulumi.Float64PtrOutput)
}

// Details of this task
func (o DedicatedServerInstallTaskOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServerInstallTask) pulumi.StringOutput { return v.Comment }).(pulumi.StringOutput)
}

func (o DedicatedServerInstallTaskOutput) Details() DedicatedServerInstallTaskDetailsPtrOutput {
	return o.ApplyT(func(v *DedicatedServerInstallTask) DedicatedServerInstallTaskDetailsPtrOutput { return v.Details }).(DedicatedServerInstallTaskDetailsPtrOutput)
}

// Completion date
func (o DedicatedServerInstallTaskOutput) DoneDate() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServerInstallTask) pulumi.StringOutput { return v.DoneDate }).(pulumi.StringOutput)
}

// Function name
func (o DedicatedServerInstallTaskOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServerInstallTask) pulumi.StringOutput { return v.Function }).(pulumi.StringOutput)
}

// Last update
func (o DedicatedServerInstallTaskOutput) LastUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServerInstallTask) pulumi.StringOutput { return v.LastUpdate }).(pulumi.StringOutput)
}

// Partition scheme name.
func (o DedicatedServerInstallTaskOutput) PartitionSchemeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DedicatedServerInstallTask) pulumi.StringPtrOutput { return v.PartitionSchemeName }).(pulumi.StringPtrOutput)
}

// The internal name of your dedicated server.
func (o DedicatedServerInstallTaskOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServerInstallTask) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Task Creation date
func (o DedicatedServerInstallTaskOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServerInstallTask) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

// Task status
func (o DedicatedServerInstallTaskOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServerInstallTask) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Template name
func (o DedicatedServerInstallTaskOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServerInstallTask) pulumi.StringOutput { return v.TemplateName }).(pulumi.StringOutput)
}

func (o DedicatedServerInstallTaskOutput) Timeouts() DedicatedServerInstallTaskTimeoutsPtrOutput {
	return o.ApplyT(func(v *DedicatedServerInstallTask) DedicatedServerInstallTaskTimeoutsPtrOutput { return v.Timeouts }).(DedicatedServerInstallTaskTimeoutsPtrOutput)
}

func (o DedicatedServerInstallTaskOutput) UserMetadatas() DedicatedServerInstallTaskUserMetadataArrayOutput {
	return o.ApplyT(func(v *DedicatedServerInstallTask) DedicatedServerInstallTaskUserMetadataArrayOutput {
		return v.UserMetadatas
	}).(DedicatedServerInstallTaskUserMetadataArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedServerInstallTaskInput)(nil)).Elem(), &DedicatedServerInstallTask{})
	pulumi.RegisterOutputType(DedicatedServerInstallTaskOutput{})
}