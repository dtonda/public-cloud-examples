// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCloudProjectKubeNodepoolNodes(ctx *pulumi.Context, args *GetCloudProjectKubeNodepoolNodesArgs, opts ...pulumi.InvokeOption) (*GetCloudProjectKubeNodepoolNodesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetCloudProjectKubeNodepoolNodesResult
	err = ctx.InvokePackage("ovh:index/getCloudProjectKubeNodepoolNodes:getCloudProjectKubeNodepoolNodes", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectKubeNodepoolNodes.
type GetCloudProjectKubeNodepoolNodesArgs struct {
	Id          *string `pulumi:"id"`
	KubeId      string  `pulumi:"kubeId"`
	Name        string  `pulumi:"name"`
	ServiceName *string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectKubeNodepoolNodes.
type GetCloudProjectKubeNodepoolNodesResult struct {
	Id          string                                 `pulumi:"id"`
	KubeId      string                                 `pulumi:"kubeId"`
	Name        string                                 `pulumi:"name"`
	Nodes       []GetCloudProjectKubeNodepoolNodesNode `pulumi:"nodes"`
	ServiceName *string                                `pulumi:"serviceName"`
}

func GetCloudProjectKubeNodepoolNodesOutput(ctx *pulumi.Context, args GetCloudProjectKubeNodepoolNodesOutputArgs, opts ...pulumi.InvokeOption) GetCloudProjectKubeNodepoolNodesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCloudProjectKubeNodepoolNodesResult, error) {
			args := v.(GetCloudProjectKubeNodepoolNodesArgs)
			r, err := GetCloudProjectKubeNodepoolNodes(ctx, &args, opts...)
			var s GetCloudProjectKubeNodepoolNodesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCloudProjectKubeNodepoolNodesResultOutput)
}

// A collection of arguments for invoking getCloudProjectKubeNodepoolNodes.
type GetCloudProjectKubeNodepoolNodesOutputArgs struct {
	Id          pulumi.StringPtrInput `pulumi:"id"`
	KubeId      pulumi.StringInput    `pulumi:"kubeId"`
	Name        pulumi.StringInput    `pulumi:"name"`
	ServiceName pulumi.StringPtrInput `pulumi:"serviceName"`
}

func (GetCloudProjectKubeNodepoolNodesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectKubeNodepoolNodesArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectKubeNodepoolNodes.
type GetCloudProjectKubeNodepoolNodesResultOutput struct{ *pulumi.OutputState }

func (GetCloudProjectKubeNodepoolNodesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudProjectKubeNodepoolNodesResult)(nil)).Elem()
}

func (o GetCloudProjectKubeNodepoolNodesResultOutput) ToGetCloudProjectKubeNodepoolNodesResultOutput() GetCloudProjectKubeNodepoolNodesResultOutput {
	return o
}

func (o GetCloudProjectKubeNodepoolNodesResultOutput) ToGetCloudProjectKubeNodepoolNodesResultOutputWithContext(ctx context.Context) GetCloudProjectKubeNodepoolNodesResultOutput {
	return o
}

func (o GetCloudProjectKubeNodepoolNodesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectKubeNodepoolNodesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCloudProjectKubeNodepoolNodesResultOutput) KubeId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectKubeNodepoolNodesResult) string { return v.KubeId }).(pulumi.StringOutput)
}

func (o GetCloudProjectKubeNodepoolNodesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudProjectKubeNodepoolNodesResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetCloudProjectKubeNodepoolNodesResultOutput) Nodes() GetCloudProjectKubeNodepoolNodesNodeArrayOutput {
	return o.ApplyT(func(v GetCloudProjectKubeNodepoolNodesResult) []GetCloudProjectKubeNodepoolNodesNode { return v.Nodes }).(GetCloudProjectKubeNodepoolNodesNodeArrayOutput)
}

func (o GetCloudProjectKubeNodepoolNodesResultOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCloudProjectKubeNodepoolNodesResult) *string { return v.ServiceName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudProjectKubeNodepoolNodesResultOutput{})
}