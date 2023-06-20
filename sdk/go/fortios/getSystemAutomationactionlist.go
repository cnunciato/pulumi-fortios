// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemAutomationaction`.
func GetSystemAutomationactionlist(ctx *pulumi.Context, args *GetSystemAutomationactionlistArgs, opts ...pulumi.InvokeOption) (*GetSystemAutomationactionlistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemAutomationactionlistResult
	err := ctx.Invoke("fortios:index/getSystemAutomationactionlist:getSystemAutomationactionlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemAutomationactionlist.
type GetSystemAutomationactionlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemAutomationactionlist.
type GetSystemAutomationactionlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `SystemAutomationaction`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemAutomationactionlistOutput(ctx *pulumi.Context, args GetSystemAutomationactionlistOutputArgs, opts ...pulumi.InvokeOption) GetSystemAutomationactionlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemAutomationactionlistResult, error) {
			args := v.(GetSystemAutomationactionlistArgs)
			r, err := GetSystemAutomationactionlist(ctx, &args, opts...)
			var s GetSystemAutomationactionlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemAutomationactionlistResultOutput)
}

// A collection of arguments for invoking getSystemAutomationactionlist.
type GetSystemAutomationactionlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemAutomationactionlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemAutomationactionlistArgs)(nil)).Elem()
}

// A collection of values returned by getSystemAutomationactionlist.
type GetSystemAutomationactionlistResultOutput struct{ *pulumi.OutputState }

func (GetSystemAutomationactionlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemAutomationactionlistResult)(nil)).Elem()
}

func (o GetSystemAutomationactionlistResultOutput) ToGetSystemAutomationactionlistResultOutput() GetSystemAutomationactionlistResultOutput {
	return o
}

func (o GetSystemAutomationactionlistResultOutput) ToGetSystemAutomationactionlistResultOutputWithContext(ctx context.Context) GetSystemAutomationactionlistResultOutput {
	return o
}

func (o GetSystemAutomationactionlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemAutomationactionlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemAutomationactionlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemAutomationactionlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `SystemAutomationaction`.
func (o GetSystemAutomationactionlistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemAutomationactionlistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemAutomationactionlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemAutomationactionlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemAutomationactionlistResultOutput{})
}
