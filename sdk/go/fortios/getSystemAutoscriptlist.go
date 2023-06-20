// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemAutoscript`.
func GetSystemAutoscriptlist(ctx *pulumi.Context, args *GetSystemAutoscriptlistArgs, opts ...pulumi.InvokeOption) (*GetSystemAutoscriptlistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemAutoscriptlistResult
	err := ctx.Invoke("fortios:index/getSystemAutoscriptlist:getSystemAutoscriptlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemAutoscriptlist.
type GetSystemAutoscriptlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemAutoscriptlist.
type GetSystemAutoscriptlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `SystemAutoscript`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemAutoscriptlistOutput(ctx *pulumi.Context, args GetSystemAutoscriptlistOutputArgs, opts ...pulumi.InvokeOption) GetSystemAutoscriptlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemAutoscriptlistResult, error) {
			args := v.(GetSystemAutoscriptlistArgs)
			r, err := GetSystemAutoscriptlist(ctx, &args, opts...)
			var s GetSystemAutoscriptlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemAutoscriptlistResultOutput)
}

// A collection of arguments for invoking getSystemAutoscriptlist.
type GetSystemAutoscriptlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemAutoscriptlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemAutoscriptlistArgs)(nil)).Elem()
}

// A collection of values returned by getSystemAutoscriptlist.
type GetSystemAutoscriptlistResultOutput struct{ *pulumi.OutputState }

func (GetSystemAutoscriptlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemAutoscriptlistResult)(nil)).Elem()
}

func (o GetSystemAutoscriptlistResultOutput) ToGetSystemAutoscriptlistResultOutput() GetSystemAutoscriptlistResultOutput {
	return o
}

func (o GetSystemAutoscriptlistResultOutput) ToGetSystemAutoscriptlistResultOutputWithContext(ctx context.Context) GetSystemAutoscriptlistResultOutput {
	return o
}

func (o GetSystemAutoscriptlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemAutoscriptlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemAutoscriptlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemAutoscriptlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `SystemAutoscript`.
func (o GetSystemAutoscriptlistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemAutoscriptlistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemAutoscriptlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemAutoscriptlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemAutoscriptlistResultOutput{})
}
