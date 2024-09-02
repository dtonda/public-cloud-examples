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

type CloudProjectKubeIprestrictions struct {
	pulumi.CustomResourceState

	// List of IP restrictions for the cluster
	Ips pulumi.StringArrayOutput `pulumi:"ips"`
	// Kube ID
	KubeId pulumi.StringOutput `pulumi:"kubeId"`
	// Service name
	ServiceName pulumi.StringPtrOutput                          `pulumi:"serviceName"`
	Timeouts    CloudProjectKubeIprestrictionsTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewCloudProjectKubeIprestrictions registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectKubeIprestrictions(ctx *pulumi.Context,
	name string, args *CloudProjectKubeIprestrictionsArgs, opts ...pulumi.ResourceOption) (*CloudProjectKubeIprestrictions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ips == nil {
		return nil, errors.New("invalid value for required argument 'Ips'")
	}
	if args.KubeId == nil {
		return nil, errors.New("invalid value for required argument 'KubeId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource CloudProjectKubeIprestrictions
	err = ctx.RegisterPackageResource("ovh:index/cloudProjectKubeIprestrictions:CloudProjectKubeIprestrictions", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectKubeIprestrictions gets an existing CloudProjectKubeIprestrictions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectKubeIprestrictions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectKubeIprestrictionsState, opts ...pulumi.ResourceOption) (*CloudProjectKubeIprestrictions, error) {
	var resource CloudProjectKubeIprestrictions
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("ovh:index/cloudProjectKubeIprestrictions:CloudProjectKubeIprestrictions", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectKubeIprestrictions resources.
type cloudProjectKubeIprestrictionsState struct {
	// List of IP restrictions for the cluster
	Ips []string `pulumi:"ips"`
	// Kube ID
	KubeId *string `pulumi:"kubeId"`
	// Service name
	ServiceName *string                                 `pulumi:"serviceName"`
	Timeouts    *CloudProjectKubeIprestrictionsTimeouts `pulumi:"timeouts"`
}

type CloudProjectKubeIprestrictionsState struct {
	// List of IP restrictions for the cluster
	Ips pulumi.StringArrayInput
	// Kube ID
	KubeId pulumi.StringPtrInput
	// Service name
	ServiceName pulumi.StringPtrInput
	Timeouts    CloudProjectKubeIprestrictionsTimeoutsPtrInput
}

func (CloudProjectKubeIprestrictionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectKubeIprestrictionsState)(nil)).Elem()
}

type cloudProjectKubeIprestrictionsArgs struct {
	// List of IP restrictions for the cluster
	Ips []string `pulumi:"ips"`
	// Kube ID
	KubeId string `pulumi:"kubeId"`
	// Service name
	ServiceName *string                                 `pulumi:"serviceName"`
	Timeouts    *CloudProjectKubeIprestrictionsTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a CloudProjectKubeIprestrictions resource.
type CloudProjectKubeIprestrictionsArgs struct {
	// List of IP restrictions for the cluster
	Ips pulumi.StringArrayInput
	// Kube ID
	KubeId pulumi.StringInput
	// Service name
	ServiceName pulumi.StringPtrInput
	Timeouts    CloudProjectKubeIprestrictionsTimeoutsPtrInput
}

func (CloudProjectKubeIprestrictionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectKubeIprestrictionsArgs)(nil)).Elem()
}

type CloudProjectKubeIprestrictionsInput interface {
	pulumi.Input

	ToCloudProjectKubeIprestrictionsOutput() CloudProjectKubeIprestrictionsOutput
	ToCloudProjectKubeIprestrictionsOutputWithContext(ctx context.Context) CloudProjectKubeIprestrictionsOutput
}

func (*CloudProjectKubeIprestrictions) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectKubeIprestrictions)(nil)).Elem()
}

func (i *CloudProjectKubeIprestrictions) ToCloudProjectKubeIprestrictionsOutput() CloudProjectKubeIprestrictionsOutput {
	return i.ToCloudProjectKubeIprestrictionsOutputWithContext(context.Background())
}

func (i *CloudProjectKubeIprestrictions) ToCloudProjectKubeIprestrictionsOutputWithContext(ctx context.Context) CloudProjectKubeIprestrictionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectKubeIprestrictionsOutput)
}

type CloudProjectKubeIprestrictionsOutput struct{ *pulumi.OutputState }

func (CloudProjectKubeIprestrictionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectKubeIprestrictions)(nil)).Elem()
}

func (o CloudProjectKubeIprestrictionsOutput) ToCloudProjectKubeIprestrictionsOutput() CloudProjectKubeIprestrictionsOutput {
	return o
}

func (o CloudProjectKubeIprestrictionsOutput) ToCloudProjectKubeIprestrictionsOutputWithContext(ctx context.Context) CloudProjectKubeIprestrictionsOutput {
	return o
}

// List of IP restrictions for the cluster
func (o CloudProjectKubeIprestrictionsOutput) Ips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudProjectKubeIprestrictions) pulumi.StringArrayOutput { return v.Ips }).(pulumi.StringArrayOutput)
}

// Kube ID
func (o CloudProjectKubeIprestrictionsOutput) KubeId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectKubeIprestrictions) pulumi.StringOutput { return v.KubeId }).(pulumi.StringOutput)
}

// Service name
func (o CloudProjectKubeIprestrictionsOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudProjectKubeIprestrictions) pulumi.StringPtrOutput { return v.ServiceName }).(pulumi.StringPtrOutput)
}

func (o CloudProjectKubeIprestrictionsOutput) Timeouts() CloudProjectKubeIprestrictionsTimeoutsPtrOutput {
	return o.ApplyT(func(v *CloudProjectKubeIprestrictions) CloudProjectKubeIprestrictionsTimeoutsPtrOutput {
		return v.Timeouts
	}).(CloudProjectKubeIprestrictionsTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectKubeIprestrictionsInput)(nil)).Elem(), &CloudProjectKubeIprestrictions{})
	pulumi.RegisterOutputType(CloudProjectKubeIprestrictionsOutput{})
}