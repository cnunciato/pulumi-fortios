// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a FortiAPI Generic Interface data source.
func LookupJsonGenericApi(ctx *pulumi.Context, args *LookupJsonGenericApiArgs, opts ...pulumi.InvokeOption) (*LookupJsonGenericApiResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupJsonGenericApiResult
	err := ctx.Invoke("fortios:index/getJsonGenericApi:getJsonGenericApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getJsonGenericApi.
type LookupJsonGenericApiArgs struct {
	// FortiAPI URL path.
	Path string `pulumi:"path"`
	// URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..
	Specialparams *string `pulumi:"specialparams"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

// A collection of values returned by getJsonGenericApi.
type LookupJsonGenericApiResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// FortiAPI URL path.
	Path string `pulumi:"path"`
	// FortiAPI returns results.
	Response string `pulumi:"response"`
	// URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..
	Specialparams *string `pulumi:"specialparams"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

func LookupJsonGenericApiOutput(ctx *pulumi.Context, args LookupJsonGenericApiOutputArgs, opts ...pulumi.InvokeOption) LookupJsonGenericApiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJsonGenericApiResult, error) {
			args := v.(LookupJsonGenericApiArgs)
			r, err := LookupJsonGenericApi(ctx, &args, opts...)
			var s LookupJsonGenericApiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJsonGenericApiResultOutput)
}

// A collection of arguments for invoking getJsonGenericApi.
type LookupJsonGenericApiOutputArgs struct {
	// FortiAPI URL path.
	Path pulumi.StringInput `pulumi:"path"`
	// URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..
	Specialparams pulumi.StringPtrInput `pulumi:"specialparams"`
	Vdomparam     pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupJsonGenericApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJsonGenericApiArgs)(nil)).Elem()
}

// A collection of values returned by getJsonGenericApi.
type LookupJsonGenericApiResultOutput struct{ *pulumi.OutputState }

func (LookupJsonGenericApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJsonGenericApiResult)(nil)).Elem()
}

func (o LookupJsonGenericApiResultOutput) ToLookupJsonGenericApiResultOutput() LookupJsonGenericApiResultOutput {
	return o
}

func (o LookupJsonGenericApiResultOutput) ToLookupJsonGenericApiResultOutputWithContext(ctx context.Context) LookupJsonGenericApiResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupJsonGenericApiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJsonGenericApiResult) string { return v.Id }).(pulumi.StringOutput)
}

// FortiAPI URL path.
func (o LookupJsonGenericApiResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJsonGenericApiResult) string { return v.Path }).(pulumi.StringOutput)
}

// FortiAPI returns results.
func (o LookupJsonGenericApiResultOutput) Response() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJsonGenericApiResult) string { return v.Response }).(pulumi.StringOutput)
}

// URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..
func (o LookupJsonGenericApiResultOutput) Specialparams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJsonGenericApiResult) *string { return v.Specialparams }).(pulumi.StringPtrOutput)
}

func (o LookupJsonGenericApiResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJsonGenericApiResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJsonGenericApiResultOutput{})
}