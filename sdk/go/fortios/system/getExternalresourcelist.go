// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Provides a list of `system.Externalresource`.
func GetExternalresourcelist(ctx *pulumi.Context, args *GetExternalresourcelistArgs, opts ...pulumi.InvokeOption) (*GetExternalresourcelistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetExternalresourcelistResult
	err := ctx.Invoke("fortios:system/getExternalresourcelist:getExternalresourcelist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getExternalresourcelist.
type GetExternalresourcelistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getExternalresourcelist.
type GetExternalresourcelistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `system.Externalresource`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetExternalresourcelistOutput(ctx *pulumi.Context, args GetExternalresourcelistOutputArgs, opts ...pulumi.InvokeOption) GetExternalresourcelistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetExternalresourcelistResult, error) {
			args := v.(GetExternalresourcelistArgs)
			r, err := GetExternalresourcelist(ctx, &args, opts...)
			var s GetExternalresourcelistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetExternalresourcelistResultOutput)
}

// A collection of arguments for invoking getExternalresourcelist.
type GetExternalresourcelistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetExternalresourcelistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExternalresourcelistArgs)(nil)).Elem()
}

// A collection of values returned by getExternalresourcelist.
type GetExternalresourcelistResultOutput struct{ *pulumi.OutputState }

func (GetExternalresourcelistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExternalresourcelistResult)(nil)).Elem()
}

func (o GetExternalresourcelistResultOutput) ToGetExternalresourcelistResultOutput() GetExternalresourcelistResultOutput {
	return o
}

func (o GetExternalresourcelistResultOutput) ToGetExternalresourcelistResultOutputWithContext(ctx context.Context) GetExternalresourcelistResultOutput {
	return o
}

func (o GetExternalresourcelistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetExternalresourcelistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetExternalresourcelistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetExternalresourcelistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `system.Externalresource`.
func (o GetExternalresourcelistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetExternalresourcelistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetExternalresourcelistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetExternalresourcelistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetExternalresourcelistResultOutput{})
}
