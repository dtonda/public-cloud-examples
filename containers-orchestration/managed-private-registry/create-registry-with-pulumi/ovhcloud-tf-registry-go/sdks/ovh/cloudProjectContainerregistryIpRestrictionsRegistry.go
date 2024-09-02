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

type CloudProjectContainerregistryIpRestrictionsRegistry struct {
	pulumi.CustomResourceState

	// List your IP restrictions applied on artifact manager component
	IpRestrictions pulumi.StringMapArrayOutput `pulumi:"ipRestrictions"`
	// RegistryID
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// Service name
	ServiceName pulumi.StringPtrOutput `pulumi:"serviceName"`
}

// NewCloudProjectContainerregistryIpRestrictionsRegistry registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectContainerregistryIpRestrictionsRegistry(ctx *pulumi.Context,
	name string, args *CloudProjectContainerregistryIpRestrictionsRegistryArgs, opts ...pulumi.ResourceOption) (*CloudProjectContainerregistryIpRestrictionsRegistry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpRestrictions == nil {
		return nil, errors.New("invalid value for required argument 'IpRestrictions'")
	}
	if args.RegistryId == nil {
		return nil, errors.New("invalid value for required argument 'RegistryId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource CloudProjectContainerregistryIpRestrictionsRegistry
	err = ctx.RegisterPackageResource("ovh:index/cloudProjectContainerregistryIpRestrictionsRegistry:CloudProjectContainerregistryIpRestrictionsRegistry", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectContainerregistryIpRestrictionsRegistry gets an existing CloudProjectContainerregistryIpRestrictionsRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectContainerregistryIpRestrictionsRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectContainerregistryIpRestrictionsRegistryState, opts ...pulumi.ResourceOption) (*CloudProjectContainerregistryIpRestrictionsRegistry, error) {
	var resource CloudProjectContainerregistryIpRestrictionsRegistry
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("ovh:index/cloudProjectContainerregistryIpRestrictionsRegistry:CloudProjectContainerregistryIpRestrictionsRegistry", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectContainerregistryIpRestrictionsRegistry resources.
type cloudProjectContainerregistryIpRestrictionsRegistryState struct {
	// List your IP restrictions applied on artifact manager component
	IpRestrictions []map[string]string `pulumi:"ipRestrictions"`
	// RegistryID
	RegistryId *string `pulumi:"registryId"`
	// Service name
	ServiceName *string `pulumi:"serviceName"`
}

type CloudProjectContainerregistryIpRestrictionsRegistryState struct {
	// List your IP restrictions applied on artifact manager component
	IpRestrictions pulumi.StringMapArrayInput
	// RegistryID
	RegistryId pulumi.StringPtrInput
	// Service name
	ServiceName pulumi.StringPtrInput
}

func (CloudProjectContainerregistryIpRestrictionsRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectContainerregistryIpRestrictionsRegistryState)(nil)).Elem()
}

type cloudProjectContainerregistryIpRestrictionsRegistryArgs struct {
	// List your IP restrictions applied on artifact manager component
	IpRestrictions []map[string]string `pulumi:"ipRestrictions"`
	// RegistryID
	RegistryId string `pulumi:"registryId"`
	// Service name
	ServiceName *string `pulumi:"serviceName"`
}

// The set of arguments for constructing a CloudProjectContainerregistryIpRestrictionsRegistry resource.
type CloudProjectContainerregistryIpRestrictionsRegistryArgs struct {
	// List your IP restrictions applied on artifact manager component
	IpRestrictions pulumi.StringMapArrayInput
	// RegistryID
	RegistryId pulumi.StringInput
	// Service name
	ServiceName pulumi.StringPtrInput
}

func (CloudProjectContainerregistryIpRestrictionsRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectContainerregistryIpRestrictionsRegistryArgs)(nil)).Elem()
}

type CloudProjectContainerregistryIpRestrictionsRegistryInput interface {
	pulumi.Input

	ToCloudProjectContainerregistryIpRestrictionsRegistryOutput() CloudProjectContainerregistryIpRestrictionsRegistryOutput
	ToCloudProjectContainerregistryIpRestrictionsRegistryOutputWithContext(ctx context.Context) CloudProjectContainerregistryIpRestrictionsRegistryOutput
}

func (*CloudProjectContainerregistryIpRestrictionsRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectContainerregistryIpRestrictionsRegistry)(nil)).Elem()
}

func (i *CloudProjectContainerregistryIpRestrictionsRegistry) ToCloudProjectContainerregistryIpRestrictionsRegistryOutput() CloudProjectContainerregistryIpRestrictionsRegistryOutput {
	return i.ToCloudProjectContainerregistryIpRestrictionsRegistryOutputWithContext(context.Background())
}

func (i *CloudProjectContainerregistryIpRestrictionsRegistry) ToCloudProjectContainerregistryIpRestrictionsRegistryOutputWithContext(ctx context.Context) CloudProjectContainerregistryIpRestrictionsRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectContainerregistryIpRestrictionsRegistryOutput)
}

type CloudProjectContainerregistryIpRestrictionsRegistryOutput struct{ *pulumi.OutputState }

func (CloudProjectContainerregistryIpRestrictionsRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectContainerregistryIpRestrictionsRegistry)(nil)).Elem()
}

func (o CloudProjectContainerregistryIpRestrictionsRegistryOutput) ToCloudProjectContainerregistryIpRestrictionsRegistryOutput() CloudProjectContainerregistryIpRestrictionsRegistryOutput {
	return o
}

func (o CloudProjectContainerregistryIpRestrictionsRegistryOutput) ToCloudProjectContainerregistryIpRestrictionsRegistryOutputWithContext(ctx context.Context) CloudProjectContainerregistryIpRestrictionsRegistryOutput {
	return o
}

// List your IP restrictions applied on artifact manager component
func (o CloudProjectContainerregistryIpRestrictionsRegistryOutput) IpRestrictions() pulumi.StringMapArrayOutput {
	return o.ApplyT(func(v *CloudProjectContainerregistryIpRestrictionsRegistry) pulumi.StringMapArrayOutput {
		return v.IpRestrictions
	}).(pulumi.StringMapArrayOutput)
}

// RegistryID
func (o CloudProjectContainerregistryIpRestrictionsRegistryOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerregistryIpRestrictionsRegistry) pulumi.StringOutput { return v.RegistryId }).(pulumi.StringOutput)
}

// Service name
func (o CloudProjectContainerregistryIpRestrictionsRegistryOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudProjectContainerregistryIpRestrictionsRegistry) pulumi.StringPtrOutput {
		return v.ServiceName
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectContainerregistryIpRestrictionsRegistryInput)(nil)).Elem(), &CloudProjectContainerregistryIpRestrictionsRegistry{})
	pulumi.RegisterOutputType(CloudProjectContainerregistryIpRestrictionsRegistryOutput{})
}