// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDbaasLogsOutputGraylogStream(ctx *pulumi.Context, args *LookupDbaasLogsOutputGraylogStreamArgs, opts ...pulumi.InvokeOption) (*LookupDbaasLogsOutputGraylogStreamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupDbaasLogsOutputGraylogStreamResult
	err = ctx.InvokePackage("ovh:index/getDbaasLogsOutputGraylogStream:getDbaasLogsOutputGraylogStream", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDbaasLogsOutputGraylogStream.
type LookupDbaasLogsOutputGraylogStreamArgs struct {
	Id          *string `pulumi:"id"`
	ServiceName string  `pulumi:"serviceName"`
	Title       string  `pulumi:"title"`
}

// A collection of values returned by getDbaasLogsOutputGraylogStream.
type LookupDbaasLogsOutputGraylogStreamResult struct {
	CanAlert                 bool    `pulumi:"canAlert"`
	ColdStorageCompression   string  `pulumi:"coldStorageCompression"`
	ColdStorageContent       string  `pulumi:"coldStorageContent"`
	ColdStorageEnabled       bool    `pulumi:"coldStorageEnabled"`
	ColdStorageNotifyEnabled bool    `pulumi:"coldStorageNotifyEnabled"`
	ColdStorageRetention     float64 `pulumi:"coldStorageRetention"`
	ColdStorageTarget        string  `pulumi:"coldStorageTarget"`
	CreatedAt                string  `pulumi:"createdAt"`
	Description              string  `pulumi:"description"`
	Id                       string  `pulumi:"id"`
	IndexingEnabled          bool    `pulumi:"indexingEnabled"`
	IndexingMaxSize          float64 `pulumi:"indexingMaxSize"`
	IndexingNotifyEnabled    bool    `pulumi:"indexingNotifyEnabled"`
	IsEditable               bool    `pulumi:"isEditable"`
	IsShareable              bool    `pulumi:"isShareable"`
	NbAlertCondition         float64 `pulumi:"nbAlertCondition"`
	NbArchive                float64 `pulumi:"nbArchive"`
	ParentStreamId           string  `pulumi:"parentStreamId"`
	PauseIndexingOnMaxSize   bool    `pulumi:"pauseIndexingOnMaxSize"`
	RetentionId              string  `pulumi:"retentionId"`
	ServiceName              string  `pulumi:"serviceName"`
	StreamId                 string  `pulumi:"streamId"`
	Title                    string  `pulumi:"title"`
	UpdatedAt                string  `pulumi:"updatedAt"`
	WebSocketEnabled         bool    `pulumi:"webSocketEnabled"`
	WriteToken               string  `pulumi:"writeToken"`
}

func LookupDbaasLogsOutputGraylogStreamOutput(ctx *pulumi.Context, args LookupDbaasLogsOutputGraylogStreamOutputArgs, opts ...pulumi.InvokeOption) LookupDbaasLogsOutputGraylogStreamResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDbaasLogsOutputGraylogStreamResult, error) {
			args := v.(LookupDbaasLogsOutputGraylogStreamArgs)
			r, err := LookupDbaasLogsOutputGraylogStream(ctx, &args, opts...)
			var s LookupDbaasLogsOutputGraylogStreamResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDbaasLogsOutputGraylogStreamResultOutput)
}

// A collection of arguments for invoking getDbaasLogsOutputGraylogStream.
type LookupDbaasLogsOutputGraylogStreamOutputArgs struct {
	Id          pulumi.StringPtrInput `pulumi:"id"`
	ServiceName pulumi.StringInput    `pulumi:"serviceName"`
	Title       pulumi.StringInput    `pulumi:"title"`
}

func (LookupDbaasLogsOutputGraylogStreamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDbaasLogsOutputGraylogStreamArgs)(nil)).Elem()
}

// A collection of values returned by getDbaasLogsOutputGraylogStream.
type LookupDbaasLogsOutputGraylogStreamResultOutput struct{ *pulumi.OutputState }

func (LookupDbaasLogsOutputGraylogStreamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDbaasLogsOutputGraylogStreamResult)(nil)).Elem()
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) ToLookupDbaasLogsOutputGraylogStreamResultOutput() LookupDbaasLogsOutputGraylogStreamResultOutput {
	return o
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) ToLookupDbaasLogsOutputGraylogStreamResultOutputWithContext(ctx context.Context) LookupDbaasLogsOutputGraylogStreamResultOutput {
	return o
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) CanAlert() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) bool { return v.CanAlert }).(pulumi.BoolOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) ColdStorageCompression() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) string { return v.ColdStorageCompression }).(pulumi.StringOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) ColdStorageContent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) string { return v.ColdStorageContent }).(pulumi.StringOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) ColdStorageEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) bool { return v.ColdStorageEnabled }).(pulumi.BoolOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) ColdStorageNotifyEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) bool { return v.ColdStorageNotifyEnabled }).(pulumi.BoolOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) ColdStorageRetention() pulumi.Float64Output {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) float64 { return v.ColdStorageRetention }).(pulumi.Float64Output)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) ColdStorageTarget() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) string { return v.ColdStorageTarget }).(pulumi.StringOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) IndexingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) bool { return v.IndexingEnabled }).(pulumi.BoolOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) IndexingMaxSize() pulumi.Float64Output {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) float64 { return v.IndexingMaxSize }).(pulumi.Float64Output)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) IndexingNotifyEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) bool { return v.IndexingNotifyEnabled }).(pulumi.BoolOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) IsEditable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) bool { return v.IsEditable }).(pulumi.BoolOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) IsShareable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) bool { return v.IsShareable }).(pulumi.BoolOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) NbAlertCondition() pulumi.Float64Output {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) float64 { return v.NbAlertCondition }).(pulumi.Float64Output)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) NbArchive() pulumi.Float64Output {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) float64 { return v.NbArchive }).(pulumi.Float64Output)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) ParentStreamId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) string { return v.ParentStreamId }).(pulumi.StringOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) PauseIndexingOnMaxSize() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) bool { return v.PauseIndexingOnMaxSize }).(pulumi.BoolOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) RetentionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) string { return v.RetentionId }).(pulumi.StringOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) StreamId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) string { return v.StreamId }).(pulumi.StringOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) string { return v.Title }).(pulumi.StringOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) WebSocketEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) bool { return v.WebSocketEnabled }).(pulumi.BoolOutput)
}

func (o LookupDbaasLogsOutputGraylogStreamResultOutput) WriteToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDbaasLogsOutputGraylogStreamResult) string { return v.WriteToken }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDbaasLogsOutputGraylogStreamResultOutput{})
}
