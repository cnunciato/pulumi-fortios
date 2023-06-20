// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `FirewallCentralsnatmap`.
func GetFirewallCentralsnatmaplist(ctx *pulumi.Context, args *GetFirewallCentralsnatmaplistArgs, opts ...pulumi.InvokeOption) (*GetFirewallCentralsnatmaplistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallCentralsnatmaplistResult
	err := ctx.Invoke("fortios:index/getFirewallCentralsnatmaplist:getFirewallCentralsnatmaplist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallCentralsnatmaplist.
type GetFirewallCentralsnatmaplistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallCentralsnatmaplist.
type GetFirewallCentralsnatmaplistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `FirewallCentralsnatmap`.
	Policyidlists []int   `pulumi:"policyidlists"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

func GetFirewallCentralsnatmaplistOutput(ctx *pulumi.Context, args GetFirewallCentralsnatmaplistOutputArgs, opts ...pulumi.InvokeOption) GetFirewallCentralsnatmaplistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallCentralsnatmaplistResult, error) {
			args := v.(GetFirewallCentralsnatmaplistArgs)
			r, err := GetFirewallCentralsnatmaplist(ctx, &args, opts...)
			var s GetFirewallCentralsnatmaplistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallCentralsnatmaplistResultOutput)
}

// A collection of arguments for invoking getFirewallCentralsnatmaplist.
type GetFirewallCentralsnatmaplistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallCentralsnatmaplistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallCentralsnatmaplistArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallCentralsnatmaplist.
type GetFirewallCentralsnatmaplistResultOutput struct{ *pulumi.OutputState }

func (GetFirewallCentralsnatmaplistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallCentralsnatmaplistResult)(nil)).Elem()
}

func (o GetFirewallCentralsnatmaplistResultOutput) ToGetFirewallCentralsnatmaplistResultOutput() GetFirewallCentralsnatmaplistResultOutput {
	return o
}

func (o GetFirewallCentralsnatmaplistResultOutput) ToGetFirewallCentralsnatmaplistResultOutputWithContext(ctx context.Context) GetFirewallCentralsnatmaplistResultOutput {
	return o
}

func (o GetFirewallCentralsnatmaplistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallCentralsnatmaplistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallCentralsnatmaplistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallCentralsnatmaplistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `FirewallCentralsnatmap`.
func (o GetFirewallCentralsnatmaplistResultOutput) Policyidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetFirewallCentralsnatmaplistResult) []int { return v.Policyidlists }).(pulumi.IntArrayOutput)
}

func (o GetFirewallCentralsnatmaplistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallCentralsnatmaplistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallCentralsnatmaplistResultOutput{})
}
