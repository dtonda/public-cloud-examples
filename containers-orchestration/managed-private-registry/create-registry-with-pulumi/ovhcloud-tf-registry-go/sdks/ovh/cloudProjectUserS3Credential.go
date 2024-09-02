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

type CloudProjectUserS3Credential struct {
	pulumi.CustomResourceState

	AccessKeyId     pulumi.StringOutput `pulumi:"accessKeyId"`
	InternalUserId  pulumi.StringOutput `pulumi:"internalUserId"`
	SecretAccessKey pulumi.StringOutput `pulumi:"secretAccessKey"`
	// Service name of the resource representing the ID of the cloud project.
	ServiceName pulumi.StringPtrOutput `pulumi:"serviceName"`
	// The user ID
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewCloudProjectUserS3Credential registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectUserS3Credential(ctx *pulumi.Context,
	name string, args *CloudProjectUserS3CredentialArgs, opts ...pulumi.ResourceOption) (*CloudProjectUserS3Credential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secretAccessKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource CloudProjectUserS3Credential
	err = ctx.RegisterPackageResource("ovh:index/cloudProjectUserS3Credential:CloudProjectUserS3Credential", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectUserS3Credential gets an existing CloudProjectUserS3Credential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectUserS3Credential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectUserS3CredentialState, opts ...pulumi.ResourceOption) (*CloudProjectUserS3Credential, error) {
	var resource CloudProjectUserS3Credential
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("ovh:index/cloudProjectUserS3Credential:CloudProjectUserS3Credential", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectUserS3Credential resources.
type cloudProjectUserS3CredentialState struct {
	AccessKeyId     *string `pulumi:"accessKeyId"`
	InternalUserId  *string `pulumi:"internalUserId"`
	SecretAccessKey *string `pulumi:"secretAccessKey"`
	// Service name of the resource representing the ID of the cloud project.
	ServiceName *string `pulumi:"serviceName"`
	// The user ID
	UserId *string `pulumi:"userId"`
}

type CloudProjectUserS3CredentialState struct {
	AccessKeyId     pulumi.StringPtrInput
	InternalUserId  pulumi.StringPtrInput
	SecretAccessKey pulumi.StringPtrInput
	// Service name of the resource representing the ID of the cloud project.
	ServiceName pulumi.StringPtrInput
	// The user ID
	UserId pulumi.StringPtrInput
}

func (CloudProjectUserS3CredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectUserS3CredentialState)(nil)).Elem()
}

type cloudProjectUserS3CredentialArgs struct {
	// Service name of the resource representing the ID of the cloud project.
	ServiceName *string `pulumi:"serviceName"`
	// The user ID
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a CloudProjectUserS3Credential resource.
type CloudProjectUserS3CredentialArgs struct {
	// Service name of the resource representing the ID of the cloud project.
	ServiceName pulumi.StringPtrInput
	// The user ID
	UserId pulumi.StringInput
}

func (CloudProjectUserS3CredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectUserS3CredentialArgs)(nil)).Elem()
}

type CloudProjectUserS3CredentialInput interface {
	pulumi.Input

	ToCloudProjectUserS3CredentialOutput() CloudProjectUserS3CredentialOutput
	ToCloudProjectUserS3CredentialOutputWithContext(ctx context.Context) CloudProjectUserS3CredentialOutput
}

func (*CloudProjectUserS3Credential) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectUserS3Credential)(nil)).Elem()
}

func (i *CloudProjectUserS3Credential) ToCloudProjectUserS3CredentialOutput() CloudProjectUserS3CredentialOutput {
	return i.ToCloudProjectUserS3CredentialOutputWithContext(context.Background())
}

func (i *CloudProjectUserS3Credential) ToCloudProjectUserS3CredentialOutputWithContext(ctx context.Context) CloudProjectUserS3CredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectUserS3CredentialOutput)
}

type CloudProjectUserS3CredentialOutput struct{ *pulumi.OutputState }

func (CloudProjectUserS3CredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectUserS3Credential)(nil)).Elem()
}

func (o CloudProjectUserS3CredentialOutput) ToCloudProjectUserS3CredentialOutput() CloudProjectUserS3CredentialOutput {
	return o
}

func (o CloudProjectUserS3CredentialOutput) ToCloudProjectUserS3CredentialOutputWithContext(ctx context.Context) CloudProjectUserS3CredentialOutput {
	return o
}

func (o CloudProjectUserS3CredentialOutput) AccessKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectUserS3Credential) pulumi.StringOutput { return v.AccessKeyId }).(pulumi.StringOutput)
}

func (o CloudProjectUserS3CredentialOutput) InternalUserId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectUserS3Credential) pulumi.StringOutput { return v.InternalUserId }).(pulumi.StringOutput)
}

func (o CloudProjectUserS3CredentialOutput) SecretAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectUserS3Credential) pulumi.StringOutput { return v.SecretAccessKey }).(pulumi.StringOutput)
}

// Service name of the resource representing the ID of the cloud project.
func (o CloudProjectUserS3CredentialOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudProjectUserS3Credential) pulumi.StringPtrOutput { return v.ServiceName }).(pulumi.StringPtrOutput)
}

// The user ID
func (o CloudProjectUserS3CredentialOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectUserS3Credential) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectUserS3CredentialInput)(nil)).Elem(), &CloudProjectUserS3Credential{})
	pulumi.RegisterOutputType(CloudProjectUserS3CredentialOutput{})
}
