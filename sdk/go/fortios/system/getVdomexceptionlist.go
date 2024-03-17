// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `system.Vdomexception`.
func GetVdomexceptionlist(ctx *pulumi.Context, args *GetVdomexceptionlistArgs, opts ...pulumi.InvokeOption) (*GetVdomexceptionlistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVdomexceptionlistResult
	err := ctx.Invoke("fortios:system/getVdomexceptionlist:getVdomexceptionlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVdomexceptionlist.
type GetVdomexceptionlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getVdomexceptionlist.
type GetVdomexceptionlistResult struct {
	Filter *string `pulumi:"filter"`
	// A list of the `system.Vdomexception`.
	Fosidlists []int `pulumi:"fosidlists"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func GetVdomexceptionlistOutput(ctx *pulumi.Context, args GetVdomexceptionlistOutputArgs, opts ...pulumi.InvokeOption) GetVdomexceptionlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVdomexceptionlistResult, error) {
			args := v.(GetVdomexceptionlistArgs)
			r, err := GetVdomexceptionlist(ctx, &args, opts...)
			var s GetVdomexceptionlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVdomexceptionlistResultOutput)
}

// A collection of arguments for invoking getVdomexceptionlist.
type GetVdomexceptionlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetVdomexceptionlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVdomexceptionlistArgs)(nil)).Elem()
}

// A collection of values returned by getVdomexceptionlist.
type GetVdomexceptionlistResultOutput struct{ *pulumi.OutputState }

func (GetVdomexceptionlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVdomexceptionlistResult)(nil)).Elem()
}

func (o GetVdomexceptionlistResultOutput) ToGetVdomexceptionlistResultOutput() GetVdomexceptionlistResultOutput {
	return o
}

func (o GetVdomexceptionlistResultOutput) ToGetVdomexceptionlistResultOutputWithContext(ctx context.Context) GetVdomexceptionlistResultOutput {
	return o
}

func (o GetVdomexceptionlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVdomexceptionlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// A list of the `system.Vdomexception`.
func (o GetVdomexceptionlistResultOutput) Fosidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetVdomexceptionlistResult) []int { return v.Fosidlists }).(pulumi.IntArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVdomexceptionlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVdomexceptionlistResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVdomexceptionlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVdomexceptionlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVdomexceptionlistResultOutput{})
}