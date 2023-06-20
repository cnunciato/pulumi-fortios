// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemTosbasedpriority`.
func GetSystemTosbasedprioritylist(ctx *pulumi.Context, args *GetSystemTosbasedprioritylistArgs, opts ...pulumi.InvokeOption) (*GetSystemTosbasedprioritylistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemTosbasedprioritylistResult
	err := ctx.Invoke("fortios:index/getSystemTosbasedprioritylist:getSystemTosbasedprioritylist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemTosbasedprioritylist.
type GetSystemTosbasedprioritylistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemTosbasedprioritylist.
type GetSystemTosbasedprioritylistResult struct {
	Filter *string `pulumi:"filter"`
	// A list of the `SystemTosbasedpriority`.
	Fosidlists []int `pulumi:"fosidlists"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func GetSystemTosbasedprioritylistOutput(ctx *pulumi.Context, args GetSystemTosbasedprioritylistOutputArgs, opts ...pulumi.InvokeOption) GetSystemTosbasedprioritylistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemTosbasedprioritylistResult, error) {
			args := v.(GetSystemTosbasedprioritylistArgs)
			r, err := GetSystemTosbasedprioritylist(ctx, &args, opts...)
			var s GetSystemTosbasedprioritylistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemTosbasedprioritylistResultOutput)
}

// A collection of arguments for invoking getSystemTosbasedprioritylist.
type GetSystemTosbasedprioritylistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemTosbasedprioritylistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemTosbasedprioritylistArgs)(nil)).Elem()
}

// A collection of values returned by getSystemTosbasedprioritylist.
type GetSystemTosbasedprioritylistResultOutput struct{ *pulumi.OutputState }

func (GetSystemTosbasedprioritylistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemTosbasedprioritylistResult)(nil)).Elem()
}

func (o GetSystemTosbasedprioritylistResultOutput) ToGetSystemTosbasedprioritylistResultOutput() GetSystemTosbasedprioritylistResultOutput {
	return o
}

func (o GetSystemTosbasedprioritylistResultOutput) ToGetSystemTosbasedprioritylistResultOutputWithContext(ctx context.Context) GetSystemTosbasedprioritylistResultOutput {
	return o
}

func (o GetSystemTosbasedprioritylistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemTosbasedprioritylistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// A list of the `SystemTosbasedpriority`.
func (o GetSystemTosbasedprioritylistResultOutput) Fosidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetSystemTosbasedprioritylistResult) []int { return v.Fosidlists }).(pulumi.IntArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemTosbasedprioritylistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemTosbasedprioritylistResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemTosbasedprioritylistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemTosbasedprioritylistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemTosbasedprioritylistResultOutput{})
}
