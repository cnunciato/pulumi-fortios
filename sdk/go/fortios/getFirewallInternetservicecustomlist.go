// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `FirewallInternetservicecustom`.
func GetFirewallInternetservicecustomlist(ctx *pulumi.Context, args *GetFirewallInternetservicecustomlistArgs, opts ...pulumi.InvokeOption) (*GetFirewallInternetservicecustomlistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallInternetservicecustomlistResult
	err := ctx.Invoke("fortios:index/getFirewallInternetservicecustomlist:getFirewallInternetservicecustomlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallInternetservicecustomlist.
type GetFirewallInternetservicecustomlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallInternetservicecustomlist.
type GetFirewallInternetservicecustomlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `FirewallInternetservicecustom`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetFirewallInternetservicecustomlistOutput(ctx *pulumi.Context, args GetFirewallInternetservicecustomlistOutputArgs, opts ...pulumi.InvokeOption) GetFirewallInternetservicecustomlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallInternetservicecustomlistResult, error) {
			args := v.(GetFirewallInternetservicecustomlistArgs)
			r, err := GetFirewallInternetservicecustomlist(ctx, &args, opts...)
			var s GetFirewallInternetservicecustomlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallInternetservicecustomlistResultOutput)
}

// A collection of arguments for invoking getFirewallInternetservicecustomlist.
type GetFirewallInternetservicecustomlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallInternetservicecustomlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallInternetservicecustomlistArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallInternetservicecustomlist.
type GetFirewallInternetservicecustomlistResultOutput struct{ *pulumi.OutputState }

func (GetFirewallInternetservicecustomlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallInternetservicecustomlistResult)(nil)).Elem()
}

func (o GetFirewallInternetservicecustomlistResultOutput) ToGetFirewallInternetservicecustomlistResultOutput() GetFirewallInternetservicecustomlistResultOutput {
	return o
}

func (o GetFirewallInternetservicecustomlistResultOutput) ToGetFirewallInternetservicecustomlistResultOutputWithContext(ctx context.Context) GetFirewallInternetservicecustomlistResultOutput {
	return o
}

func (o GetFirewallInternetservicecustomlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallInternetservicecustomlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallInternetservicecustomlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallInternetservicecustomlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `FirewallInternetservicecustom`.
func (o GetFirewallInternetservicecustomlistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFirewallInternetservicecustomlistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetFirewallInternetservicecustomlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallInternetservicecustomlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallInternetservicecustomlistResultOutput{})
}