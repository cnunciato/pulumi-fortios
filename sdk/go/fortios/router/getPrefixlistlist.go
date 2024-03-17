// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `router.Prefixlist`.
func GetPrefixlistlist(ctx *pulumi.Context, args *GetPrefixlistlistArgs, opts ...pulumi.InvokeOption) (*GetPrefixlistlistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPrefixlistlistResult
	err := ctx.Invoke("fortios:router/getPrefixlistlist:getPrefixlistlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrefixlistlist.
type GetPrefixlistlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getPrefixlistlist.
type GetPrefixlistlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `router.Prefixlist`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetPrefixlistlistOutput(ctx *pulumi.Context, args GetPrefixlistlistOutputArgs, opts ...pulumi.InvokeOption) GetPrefixlistlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPrefixlistlistResult, error) {
			args := v.(GetPrefixlistlistArgs)
			r, err := GetPrefixlistlist(ctx, &args, opts...)
			var s GetPrefixlistlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPrefixlistlistResultOutput)
}

// A collection of arguments for invoking getPrefixlistlist.
type GetPrefixlistlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetPrefixlistlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPrefixlistlistArgs)(nil)).Elem()
}

// A collection of values returned by getPrefixlistlist.
type GetPrefixlistlistResultOutput struct{ *pulumi.OutputState }

func (GetPrefixlistlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPrefixlistlistResult)(nil)).Elem()
}

func (o GetPrefixlistlistResultOutput) ToGetPrefixlistlistResultOutput() GetPrefixlistlistResultOutput {
	return o
}

func (o GetPrefixlistlistResultOutput) ToGetPrefixlistlistResultOutputWithContext(ctx context.Context) GetPrefixlistlistResultOutput {
	return o
}

func (o GetPrefixlistlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPrefixlistlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPrefixlistlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPrefixlistlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `router.Prefixlist`.
func (o GetPrefixlistlistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPrefixlistlistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetPrefixlistlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPrefixlistlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPrefixlistlistResultOutput{})
}