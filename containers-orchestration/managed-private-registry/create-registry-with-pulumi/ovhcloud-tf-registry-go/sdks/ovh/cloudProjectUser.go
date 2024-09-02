// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudProjectUser struct {
	pulumi.CustomResourceState

	CreationDate pulumi.StringOutput             `pulumi:"creationDate"`
	Description  pulumi.StringPtrOutput          `pulumi:"description"`
	OpenstackRc  pulumi.StringMapOutput          `pulumi:"openstackRc"`
	Password     pulumi.StringOutput             `pulumi:"password"`
	RoleName     pulumi.StringPtrOutput          `pulumi:"roleName"`
	RoleNames    pulumi.StringArrayOutput        `pulumi:"roleNames"`
	Roles        CloudProjectUserRoleArrayOutput `pulumi:"roles"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringPtrOutput `pulumi:"serviceName"`
	Status      pulumi.StringOutput    `pulumi:"status"`
	Username    pulumi.StringOutput    `pulumi:"username"`
}

// NewCloudProjectUser registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectUser(ctx *pulumi.Context,
	name string, args *CloudProjectUserArgs, opts ...pulumi.ResourceOption) (*CloudProjectUser, error) {
	if args == nil {
		args = &CloudProjectUserArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource CloudProjectUser
	err = ctx.RegisterPackageResource("ovh:index/cloudProjectUser:CloudProjectUser", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectUser gets an existing CloudProjectUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectUserState, opts ...pulumi.ResourceOption) (*CloudProjectUser, error) {
	var resource CloudProjectUser
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("ovh:index/cloudProjectUser:CloudProjectUser", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectUser resources.
type cloudProjectUserState struct {
	CreationDate *string                `pulumi:"creationDate"`
	Description  *string                `pulumi:"description"`
	OpenstackRc  map[string]string      `pulumi:"openstackRc"`
	Password     *string                `pulumi:"password"`
	RoleName     *string                `pulumi:"roleName"`
	RoleNames    []string               `pulumi:"roleNames"`
	Roles        []CloudProjectUserRole `pulumi:"roles"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName *string `pulumi:"serviceName"`
	Status      *string `pulumi:"status"`
	Username    *string `pulumi:"username"`
}

type CloudProjectUserState struct {
	CreationDate pulumi.StringPtrInput
	Description  pulumi.StringPtrInput
	OpenstackRc  pulumi.StringMapInput
	Password     pulumi.StringPtrInput
	RoleName     pulumi.StringPtrInput
	RoleNames    pulumi.StringArrayInput
	Roles        CloudProjectUserRoleArrayInput
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringPtrInput
	Status      pulumi.StringPtrInput
	Username    pulumi.StringPtrInput
}

func (CloudProjectUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectUserState)(nil)).Elem()
}

type cloudProjectUserArgs struct {
	Description *string           `pulumi:"description"`
	OpenstackRc map[string]string `pulumi:"openstackRc"`
	RoleName    *string           `pulumi:"roleName"`
	RoleNames   []string          `pulumi:"roleNames"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName *string `pulumi:"serviceName"`
}

// The set of arguments for constructing a CloudProjectUser resource.
type CloudProjectUserArgs struct {
	Description pulumi.StringPtrInput
	OpenstackRc pulumi.StringMapInput
	RoleName    pulumi.StringPtrInput
	RoleNames   pulumi.StringArrayInput
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringPtrInput
}

func (CloudProjectUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectUserArgs)(nil)).Elem()
}

type CloudProjectUserInput interface {
	pulumi.Input

	ToCloudProjectUserOutput() CloudProjectUserOutput
	ToCloudProjectUserOutputWithContext(ctx context.Context) CloudProjectUserOutput
}

func (*CloudProjectUser) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectUser)(nil)).Elem()
}

func (i *CloudProjectUser) ToCloudProjectUserOutput() CloudProjectUserOutput {
	return i.ToCloudProjectUserOutputWithContext(context.Background())
}

func (i *CloudProjectUser) ToCloudProjectUserOutputWithContext(ctx context.Context) CloudProjectUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectUserOutput)
}

type CloudProjectUserOutput struct{ *pulumi.OutputState }

func (CloudProjectUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectUser)(nil)).Elem()
}

func (o CloudProjectUserOutput) ToCloudProjectUserOutput() CloudProjectUserOutput {
	return o
}

func (o CloudProjectUserOutput) ToCloudProjectUserOutputWithContext(ctx context.Context) CloudProjectUserOutput {
	return o
}

func (o CloudProjectUserOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o CloudProjectUserOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CloudProjectUserOutput) OpenstackRc() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.StringMapOutput { return v.OpenstackRc }).(pulumi.StringMapOutput)
}

func (o CloudProjectUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

func (o CloudProjectUserOutput) RoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.StringPtrOutput { return v.RoleName }).(pulumi.StringPtrOutput)
}

func (o CloudProjectUserOutput) RoleNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.StringArrayOutput { return v.RoleNames }).(pulumi.StringArrayOutput)
}

func (o CloudProjectUserOutput) Roles() CloudProjectUserRoleArrayOutput {
	return o.ApplyT(func(v *CloudProjectUser) CloudProjectUserRoleArrayOutput { return v.Roles }).(CloudProjectUserRoleArrayOutput)
}

// Service name of the resource representing the id of the cloud project.
func (o CloudProjectUserOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.StringPtrOutput { return v.ServiceName }).(pulumi.StringPtrOutput)
}

func (o CloudProjectUserOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o CloudProjectUserOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectUser) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectUserInput)(nil)).Elem(), &CloudProjectUser{})
	pulumi.RegisterOutputType(CloudProjectUserOutput{})
}