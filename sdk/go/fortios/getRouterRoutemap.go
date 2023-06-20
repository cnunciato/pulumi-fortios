// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios router routemap
func LookupRouterRoutemap(ctx *pulumi.Context, args *LookupRouterRoutemapArgs, opts ...pulumi.InvokeOption) (*LookupRouterRoutemapResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterRoutemapResult
	err := ctx.Invoke("fortios:index/getRouterRoutemap:getRouterRoutemap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouterRoutemap.
type LookupRouterRoutemapArgs struct {
	// Specify the name of the desired router routemap.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getRouterRoutemap.
type LookupRouterRoutemapResult struct {
	// Optional comments.
	Comments string `pulumi:"comments"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name.
	Name string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules     []GetRouterRoutemapRule `pulumi:"rules"`
	Vdomparam *string                 `pulumi:"vdomparam"`
}

func LookupRouterRoutemapOutput(ctx *pulumi.Context, args LookupRouterRoutemapOutputArgs, opts ...pulumi.InvokeOption) LookupRouterRoutemapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterRoutemapResult, error) {
			args := v.(LookupRouterRoutemapArgs)
			r, err := LookupRouterRoutemap(ctx, &args, opts...)
			var s LookupRouterRoutemapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterRoutemapResultOutput)
}

// A collection of arguments for invoking getRouterRoutemap.
type LookupRouterRoutemapOutputArgs struct {
	// Specify the name of the desired router routemap.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterRoutemapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterRoutemapArgs)(nil)).Elem()
}

// A collection of values returned by getRouterRoutemap.
type LookupRouterRoutemapResultOutput struct{ *pulumi.OutputState }

func (LookupRouterRoutemapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterRoutemapResult)(nil)).Elem()
}

func (o LookupRouterRoutemapResultOutput) ToLookupRouterRoutemapResultOutput() LookupRouterRoutemapResultOutput {
	return o
}

func (o LookupRouterRoutemapResultOutput) ToLookupRouterRoutemapResultOutputWithContext(ctx context.Context) LookupRouterRoutemapResultOutput {
	return o
}

// Optional comments.
func (o LookupRouterRoutemapResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterRoutemapResult) string { return v.Comments }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterRoutemapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterRoutemapResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name.
func (o LookupRouterRoutemapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterRoutemapResult) string { return v.Name }).(pulumi.StringOutput)
}

// Rule. The structure of `rule` block is documented below.
func (o LookupRouterRoutemapResultOutput) Rules() GetRouterRoutemapRuleArrayOutput {
	return o.ApplyT(func(v LookupRouterRoutemapResult) []GetRouterRoutemapRule { return v.Rules }).(GetRouterRoutemapRuleArrayOutput)
}

func (o LookupRouterRoutemapResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterRoutemapResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterRoutemapResultOutput{})
}
