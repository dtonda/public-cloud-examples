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

type DedicatedServerRebootTask struct {
	pulumi.CustomResourceState

	// Details of this task
	Comment pulumi.StringOutput `pulumi:"comment"`
	// Completion date
	DoneDate pulumi.StringOutput `pulumi:"doneDate"`
	// Function name
	Function pulumi.StringOutput `pulumi:"function"`
	// Change this value to recreate a reboot task.
	Keepers pulumi.StringArrayOutput `pulumi:"keepers"`
	// Last update
	LastUpdate pulumi.StringOutput `pulumi:"lastUpdate"`
	// The internal name of your dedicated server.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Task Creation date
	StartDate pulumi.StringOutput `pulumi:"startDate"`
	// Task status
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewDedicatedServerRebootTask registers a new resource with the given unique name, arguments, and options.
func NewDedicatedServerRebootTask(ctx *pulumi.Context,
	name string, args *DedicatedServerRebootTaskArgs, opts ...pulumi.ResourceOption) (*DedicatedServerRebootTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Keepers == nil {
		return nil, errors.New("invalid value for required argument 'Keepers'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DedicatedServerRebootTask
	err = ctx.RegisterPackageResource("ovh:index/dedicatedServerRebootTask:DedicatedServerRebootTask", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedServerRebootTask gets an existing DedicatedServerRebootTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedServerRebootTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedServerRebootTaskState, opts ...pulumi.ResourceOption) (*DedicatedServerRebootTask, error) {
	var resource DedicatedServerRebootTask
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("ovh:index/dedicatedServerRebootTask:DedicatedServerRebootTask", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedServerRebootTask resources.
type dedicatedServerRebootTaskState struct {
	// Details of this task
	Comment *string `pulumi:"comment"`
	// Completion date
	DoneDate *string `pulumi:"doneDate"`
	// Function name
	Function *string `pulumi:"function"`
	// Change this value to recreate a reboot task.
	Keepers []string `pulumi:"keepers"`
	// Last update
	LastUpdate *string `pulumi:"lastUpdate"`
	// The internal name of your dedicated server.
	ServiceName *string `pulumi:"serviceName"`
	// Task Creation date
	StartDate *string `pulumi:"startDate"`
	// Task status
	Status *string `pulumi:"status"`
}

type DedicatedServerRebootTaskState struct {
	// Details of this task
	Comment pulumi.StringPtrInput
	// Completion date
	DoneDate pulumi.StringPtrInput
	// Function name
	Function pulumi.StringPtrInput
	// Change this value to recreate a reboot task.
	Keepers pulumi.StringArrayInput
	// Last update
	LastUpdate pulumi.StringPtrInput
	// The internal name of your dedicated server.
	ServiceName pulumi.StringPtrInput
	// Task Creation date
	StartDate pulumi.StringPtrInput
	// Task status
	Status pulumi.StringPtrInput
}

func (DedicatedServerRebootTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedServerRebootTaskState)(nil)).Elem()
}

type dedicatedServerRebootTaskArgs struct {
	// Change this value to recreate a reboot task.
	Keepers []string `pulumi:"keepers"`
	// The internal name of your dedicated server.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a DedicatedServerRebootTask resource.
type DedicatedServerRebootTaskArgs struct {
	// Change this value to recreate a reboot task.
	Keepers pulumi.StringArrayInput
	// The internal name of your dedicated server.
	ServiceName pulumi.StringInput
}

func (DedicatedServerRebootTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedServerRebootTaskArgs)(nil)).Elem()
}

type DedicatedServerRebootTaskInput interface {
	pulumi.Input

	ToDedicatedServerRebootTaskOutput() DedicatedServerRebootTaskOutput
	ToDedicatedServerRebootTaskOutputWithContext(ctx context.Context) DedicatedServerRebootTaskOutput
}

func (*DedicatedServerRebootTask) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedServerRebootTask)(nil)).Elem()
}

func (i *DedicatedServerRebootTask) ToDedicatedServerRebootTaskOutput() DedicatedServerRebootTaskOutput {
	return i.ToDedicatedServerRebootTaskOutputWithContext(context.Background())
}

func (i *DedicatedServerRebootTask) ToDedicatedServerRebootTaskOutputWithContext(ctx context.Context) DedicatedServerRebootTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedServerRebootTaskOutput)
}

type DedicatedServerRebootTaskOutput struct{ *pulumi.OutputState }

func (DedicatedServerRebootTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedServerRebootTask)(nil)).Elem()
}

func (o DedicatedServerRebootTaskOutput) ToDedicatedServerRebootTaskOutput() DedicatedServerRebootTaskOutput {
	return o
}

func (o DedicatedServerRebootTaskOutput) ToDedicatedServerRebootTaskOutputWithContext(ctx context.Context) DedicatedServerRebootTaskOutput {
	return o
}

// Details of this task
func (o DedicatedServerRebootTaskOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServerRebootTask) pulumi.StringOutput { return v.Comment }).(pulumi.StringOutput)
}

// Completion date
func (o DedicatedServerRebootTaskOutput) DoneDate() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServerRebootTask) pulumi.StringOutput { return v.DoneDate }).(pulumi.StringOutput)
}

// Function name
func (o DedicatedServerRebootTaskOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServerRebootTask) pulumi.StringOutput { return v.Function }).(pulumi.StringOutput)
}

// Change this value to recreate a reboot task.
func (o DedicatedServerRebootTaskOutput) Keepers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DedicatedServerRebootTask) pulumi.StringArrayOutput { return v.Keepers }).(pulumi.StringArrayOutput)
}

// Last update
func (o DedicatedServerRebootTaskOutput) LastUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServerRebootTask) pulumi.StringOutput { return v.LastUpdate }).(pulumi.StringOutput)
}

// The internal name of your dedicated server.
func (o DedicatedServerRebootTaskOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServerRebootTask) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Task Creation date
func (o DedicatedServerRebootTaskOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServerRebootTask) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

// Task status
func (o DedicatedServerRebootTaskOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServerRebootTask) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedServerRebootTaskInput)(nil)).Elem(), &DedicatedServerRebootTask{})
	pulumi.RegisterOutputType(DedicatedServerRebootTaskOutput{})
}