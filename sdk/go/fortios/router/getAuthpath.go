// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Use this data source to get information on an fortios router authpath
func LookupAuthpath(ctx *pulumi.Context, args *LookupAuthpathArgs, opts ...pulumi.InvokeOption) (*LookupAuthpathResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAuthpathResult
	err := ctx.Invoke("fortios:router/getAuthpath:getAuthpath", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAuthpath.
type LookupAuthpathArgs struct {
	// Specify the name of the desired router authpath.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getAuthpath.
type LookupAuthpathResult struct {
	// Outgoing interface.
	Device string `pulumi:"device"`
	// Gateway IP address.
	Gateway string `pulumi:"gateway"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the entry.
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupAuthpathOutput(ctx *pulumi.Context, args LookupAuthpathOutputArgs, opts ...pulumi.InvokeOption) LookupAuthpathResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAuthpathResult, error) {
			args := v.(LookupAuthpathArgs)
			r, err := LookupAuthpath(ctx, &args, opts...)
			var s LookupAuthpathResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAuthpathResultOutput)
}

// A collection of arguments for invoking getAuthpath.
type LookupAuthpathOutputArgs struct {
	// Specify the name of the desired router authpath.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupAuthpathOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthpathArgs)(nil)).Elem()
}

// A collection of values returned by getAuthpath.
type LookupAuthpathResultOutput struct{ *pulumi.OutputState }

func (LookupAuthpathResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthpathResult)(nil)).Elem()
}

func (o LookupAuthpathResultOutput) ToLookupAuthpathResultOutput() LookupAuthpathResultOutput {
	return o
}

func (o LookupAuthpathResultOutput) ToLookupAuthpathResultOutputWithContext(ctx context.Context) LookupAuthpathResultOutput {
	return o
}

// Outgoing interface.
func (o LookupAuthpathResultOutput) Device() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthpathResult) string { return v.Device }).(pulumi.StringOutput)
}

// Gateway IP address.
func (o LookupAuthpathResultOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthpathResult) string { return v.Gateway }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAuthpathResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthpathResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the entry.
func (o LookupAuthpathResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthpathResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAuthpathResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthpathResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthpathResultOutput{})
}