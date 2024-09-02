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

type VrackDedicatedServer struct {
	pulumi.CustomResourceState

	ServerId pulumi.StringOutput `pulumi:"serverId"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewVrackDedicatedServer registers a new resource with the given unique name, arguments, and options.
func NewVrackDedicatedServer(ctx *pulumi.Context,
	name string, args *VrackDedicatedServerArgs, opts ...pulumi.ResourceOption) (*VrackDedicatedServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServerId == nil {
		return nil, errors.New("invalid value for required argument 'ServerId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource VrackDedicatedServer
	err = ctx.RegisterPackageResource("ovh:index/vrackDedicatedServer:VrackDedicatedServer", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVrackDedicatedServer gets an existing VrackDedicatedServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVrackDedicatedServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VrackDedicatedServerState, opts ...pulumi.ResourceOption) (*VrackDedicatedServer, error) {
	var resource VrackDedicatedServer
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("ovh:index/vrackDedicatedServer:VrackDedicatedServer", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VrackDedicatedServer resources.
type vrackDedicatedServerState struct {
	ServerId *string `pulumi:"serverId"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName *string `pulumi:"serviceName"`
}

type VrackDedicatedServerState struct {
	ServerId pulumi.StringPtrInput
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringPtrInput
}

func (VrackDedicatedServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*vrackDedicatedServerState)(nil)).Elem()
}

type vrackDedicatedServerArgs struct {
	ServerId string `pulumi:"serverId"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a VrackDedicatedServer resource.
type VrackDedicatedServerArgs struct {
	ServerId pulumi.StringInput
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringInput
}

func (VrackDedicatedServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vrackDedicatedServerArgs)(nil)).Elem()
}

type VrackDedicatedServerInput interface {
	pulumi.Input

	ToVrackDedicatedServerOutput() VrackDedicatedServerOutput
	ToVrackDedicatedServerOutputWithContext(ctx context.Context) VrackDedicatedServerOutput
}

func (*VrackDedicatedServer) ElementType() reflect.Type {
	return reflect.TypeOf((**VrackDedicatedServer)(nil)).Elem()
}

func (i *VrackDedicatedServer) ToVrackDedicatedServerOutput() VrackDedicatedServerOutput {
	return i.ToVrackDedicatedServerOutputWithContext(context.Background())
}

func (i *VrackDedicatedServer) ToVrackDedicatedServerOutputWithContext(ctx context.Context) VrackDedicatedServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackDedicatedServerOutput)
}

type VrackDedicatedServerOutput struct{ *pulumi.OutputState }

func (VrackDedicatedServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VrackDedicatedServer)(nil)).Elem()
}

func (o VrackDedicatedServerOutput) ToVrackDedicatedServerOutput() VrackDedicatedServerOutput {
	return o
}

func (o VrackDedicatedServerOutput) ToVrackDedicatedServerOutputWithContext(ctx context.Context) VrackDedicatedServerOutput {
	return o
}

func (o VrackDedicatedServerOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *VrackDedicatedServer) pulumi.StringOutput { return v.ServerId }).(pulumi.StringOutput)
}

// Service name of the resource representing the id of the cloud project.
func (o VrackDedicatedServerOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VrackDedicatedServer) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VrackDedicatedServerInput)(nil)).Elem(), &VrackDedicatedServer{})
	pulumi.RegisterOutputType(VrackDedicatedServerOutput{})
}