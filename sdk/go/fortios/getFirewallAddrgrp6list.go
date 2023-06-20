// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `FirewallAddrgrp6`.
func GetFirewallAddrgrp6list(ctx *pulumi.Context, args *GetFirewallAddrgrp6listArgs, opts ...pulumi.InvokeOption) (*GetFirewallAddrgrp6listResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallAddrgrp6listResult
	err := ctx.Invoke("fortios:index/getFirewallAddrgrp6list:getFirewallAddrgrp6list", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallAddrgrp6list.
type GetFirewallAddrgrp6listArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallAddrgrp6list.
type GetFirewallAddrgrp6listResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `FirewallAddrgrp6`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetFirewallAddrgrp6listOutput(ctx *pulumi.Context, args GetFirewallAddrgrp6listOutputArgs, opts ...pulumi.InvokeOption) GetFirewallAddrgrp6listResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallAddrgrp6listResult, error) {
			args := v.(GetFirewallAddrgrp6listArgs)
			r, err := GetFirewallAddrgrp6list(ctx, &args, opts...)
			var s GetFirewallAddrgrp6listResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallAddrgrp6listResultOutput)
}

// A collection of arguments for invoking getFirewallAddrgrp6list.
type GetFirewallAddrgrp6listOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallAddrgrp6listOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallAddrgrp6listArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallAddrgrp6list.
type GetFirewallAddrgrp6listResultOutput struct{ *pulumi.OutputState }

func (GetFirewallAddrgrp6listResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallAddrgrp6listResult)(nil)).Elem()
}

func (o GetFirewallAddrgrp6listResultOutput) ToGetFirewallAddrgrp6listResultOutput() GetFirewallAddrgrp6listResultOutput {
	return o
}

func (o GetFirewallAddrgrp6listResultOutput) ToGetFirewallAddrgrp6listResultOutputWithContext(ctx context.Context) GetFirewallAddrgrp6listResultOutput {
	return o
}

func (o GetFirewallAddrgrp6listResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallAddrgrp6listResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallAddrgrp6listResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallAddrgrp6listResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `FirewallAddrgrp6`.
func (o GetFirewallAddrgrp6listResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFirewallAddrgrp6listResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetFirewallAddrgrp6listResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallAddrgrp6listResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallAddrgrp6listResultOutput{})
}
