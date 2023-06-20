// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system networkvisibility
func LookupSystemNetworkvisibility(ctx *pulumi.Context, args *LookupSystemNetworkvisibilityArgs, opts ...pulumi.InvokeOption) (*LookupSystemNetworkvisibilityResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemNetworkvisibilityResult
	err := ctx.Invoke("fortios:index/getSystemNetworkvisibility:getSystemNetworkvisibility", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemNetworkvisibility.
type LookupSystemNetworkvisibilityArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemNetworkvisibility.
type LookupSystemNetworkvisibilityResult struct {
	// Enable/disable logging of destination hostname visibility.
	DestinationHostnameVisibility string `pulumi:"destinationHostnameVisibility"`
	// Enable/disable logging of destination geographical location visibility.
	DestinationLocation string `pulumi:"destinationLocation"`
	// Enable/disable logging of destination visibility.
	DestinationVisibility string `pulumi:"destinationVisibility"`
	// Limit of the number of hostname table entries (0 - 50000).
	HostnameLimit int `pulumi:"hostnameLimit"`
	// TTL of hostname table entries (60 - 86400).
	HostnameTtl int `pulumi:"hostnameTtl"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable/disable logging of source geographical location visibility.
	SourceLocation string  `pulumi:"sourceLocation"`
	Vdomparam      *string `pulumi:"vdomparam"`
}

func LookupSystemNetworkvisibilityOutput(ctx *pulumi.Context, args LookupSystemNetworkvisibilityOutputArgs, opts ...pulumi.InvokeOption) LookupSystemNetworkvisibilityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemNetworkvisibilityResult, error) {
			args := v.(LookupSystemNetworkvisibilityArgs)
			r, err := LookupSystemNetworkvisibility(ctx, &args, opts...)
			var s LookupSystemNetworkvisibilityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemNetworkvisibilityResultOutput)
}

// A collection of arguments for invoking getSystemNetworkvisibility.
type LookupSystemNetworkvisibilityOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemNetworkvisibilityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemNetworkvisibilityArgs)(nil)).Elem()
}

// A collection of values returned by getSystemNetworkvisibility.
type LookupSystemNetworkvisibilityResultOutput struct{ *pulumi.OutputState }

func (LookupSystemNetworkvisibilityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemNetworkvisibilityResult)(nil)).Elem()
}

func (o LookupSystemNetworkvisibilityResultOutput) ToLookupSystemNetworkvisibilityResultOutput() LookupSystemNetworkvisibilityResultOutput {
	return o
}

func (o LookupSystemNetworkvisibilityResultOutput) ToLookupSystemNetworkvisibilityResultOutputWithContext(ctx context.Context) LookupSystemNetworkvisibilityResultOutput {
	return o
}

// Enable/disable logging of destination hostname visibility.
func (o LookupSystemNetworkvisibilityResultOutput) DestinationHostnameVisibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNetworkvisibilityResult) string { return v.DestinationHostnameVisibility }).(pulumi.StringOutput)
}

// Enable/disable logging of destination geographical location visibility.
func (o LookupSystemNetworkvisibilityResultOutput) DestinationLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNetworkvisibilityResult) string { return v.DestinationLocation }).(pulumi.StringOutput)
}

// Enable/disable logging of destination visibility.
func (o LookupSystemNetworkvisibilityResultOutput) DestinationVisibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNetworkvisibilityResult) string { return v.DestinationVisibility }).(pulumi.StringOutput)
}

// Limit of the number of hostname table entries (0 - 50000).
func (o LookupSystemNetworkvisibilityResultOutput) HostnameLimit() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemNetworkvisibilityResult) int { return v.HostnameLimit }).(pulumi.IntOutput)
}

// TTL of hostname table entries (60 - 86400).
func (o LookupSystemNetworkvisibilityResultOutput) HostnameTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemNetworkvisibilityResult) int { return v.HostnameTtl }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemNetworkvisibilityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNetworkvisibilityResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable/disable logging of source geographical location visibility.
func (o LookupSystemNetworkvisibilityResultOutput) SourceLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNetworkvisibilityResult) string { return v.SourceLocation }).(pulumi.StringOutput)
}

func (o LookupSystemNetworkvisibilityResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemNetworkvisibilityResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemNetworkvisibilityResultOutput{})
}
