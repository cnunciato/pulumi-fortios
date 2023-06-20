// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemWccp`.
func GetSystemWccplist(ctx *pulumi.Context, args *GetSystemWccplistArgs, opts ...pulumi.InvokeOption) (*GetSystemWccplistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemWccplistResult
	err := ctx.Invoke("fortios:index/getSystemWccplist:getSystemWccplist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemWccplist.
type GetSystemWccplistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemWccplist.
type GetSystemWccplistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `SystemWccp`.
	ServiceIdlists []string `pulumi:"serviceIdlists"`
	Vdomparam      *string  `pulumi:"vdomparam"`
}

func GetSystemWccplistOutput(ctx *pulumi.Context, args GetSystemWccplistOutputArgs, opts ...pulumi.InvokeOption) GetSystemWccplistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemWccplistResult, error) {
			args := v.(GetSystemWccplistArgs)
			r, err := GetSystemWccplist(ctx, &args, opts...)
			var s GetSystemWccplistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemWccplistResultOutput)
}

// A collection of arguments for invoking getSystemWccplist.
type GetSystemWccplistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemWccplistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemWccplistArgs)(nil)).Elem()
}

// A collection of values returned by getSystemWccplist.
type GetSystemWccplistResultOutput struct{ *pulumi.OutputState }

func (GetSystemWccplistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemWccplistResult)(nil)).Elem()
}

func (o GetSystemWccplistResultOutput) ToGetSystemWccplistResultOutput() GetSystemWccplistResultOutput {
	return o
}

func (o GetSystemWccplistResultOutput) ToGetSystemWccplistResultOutputWithContext(ctx context.Context) GetSystemWccplistResultOutput {
	return o
}

func (o GetSystemWccplistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemWccplistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemWccplistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemWccplistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `SystemWccp`.
func (o GetSystemWccplistResultOutput) ServiceIdlists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemWccplistResult) []string { return v.ServiceIdlists }).(pulumi.StringArrayOutput)
}

func (o GetSystemWccplistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemWccplistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemWccplistResultOutput{})
}
