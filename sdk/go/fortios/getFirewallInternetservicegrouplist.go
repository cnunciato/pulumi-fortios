// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `FirewallInternetservicegroup`.
func GetFirewallInternetservicegrouplist(ctx *pulumi.Context, args *GetFirewallInternetservicegrouplistArgs, opts ...pulumi.InvokeOption) (*GetFirewallInternetservicegrouplistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallInternetservicegrouplistResult
	err := ctx.Invoke("fortios:index/getFirewallInternetservicegrouplist:getFirewallInternetservicegrouplist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallInternetservicegrouplist.
type GetFirewallInternetservicegrouplistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallInternetservicegrouplist.
type GetFirewallInternetservicegrouplistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `FirewallInternetservicegroup`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetFirewallInternetservicegrouplistOutput(ctx *pulumi.Context, args GetFirewallInternetservicegrouplistOutputArgs, opts ...pulumi.InvokeOption) GetFirewallInternetservicegrouplistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallInternetservicegrouplistResult, error) {
			args := v.(GetFirewallInternetservicegrouplistArgs)
			r, err := GetFirewallInternetservicegrouplist(ctx, &args, opts...)
			var s GetFirewallInternetservicegrouplistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallInternetservicegrouplistResultOutput)
}

// A collection of arguments for invoking getFirewallInternetservicegrouplist.
type GetFirewallInternetservicegrouplistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallInternetservicegrouplistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallInternetservicegrouplistArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallInternetservicegrouplist.
type GetFirewallInternetservicegrouplistResultOutput struct{ *pulumi.OutputState }

func (GetFirewallInternetservicegrouplistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallInternetservicegrouplistResult)(nil)).Elem()
}

func (o GetFirewallInternetservicegrouplistResultOutput) ToGetFirewallInternetservicegrouplistResultOutput() GetFirewallInternetservicegrouplistResultOutput {
	return o
}

func (o GetFirewallInternetservicegrouplistResultOutput) ToGetFirewallInternetservicegrouplistResultOutputWithContext(ctx context.Context) GetFirewallInternetservicegrouplistResultOutput {
	return o
}

func (o GetFirewallInternetservicegrouplistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallInternetservicegrouplistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallInternetservicegrouplistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallInternetservicegrouplistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `FirewallInternetservicegroup`.
func (o GetFirewallInternetservicegrouplistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFirewallInternetservicegrouplistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetFirewallInternetservicegrouplistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallInternetservicegrouplistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallInternetservicegrouplistResultOutput{})
}