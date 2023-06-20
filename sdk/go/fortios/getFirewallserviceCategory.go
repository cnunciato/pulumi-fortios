// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewallservice category
func LookupFirewallserviceCategory(ctx *pulumi.Context, args *LookupFirewallserviceCategoryArgs, opts ...pulumi.InvokeOption) (*LookupFirewallserviceCategoryResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallserviceCategoryResult
	err := ctx.Invoke("fortios:index/getFirewallserviceCategory:getFirewallserviceCategory", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallserviceCategory.
type LookupFirewallserviceCategoryArgs struct {
	// Specify the name of the desired firewallservice category.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallserviceCategory.
type LookupFirewallserviceCategoryResult struct {
	// Comment.
	Comment string `pulumi:"comment"`
	// Security Fabric global object setting.
	FabricObject string `pulumi:"fabricObject"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Service category name.
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupFirewallserviceCategoryOutput(ctx *pulumi.Context, args LookupFirewallserviceCategoryOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallserviceCategoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallserviceCategoryResult, error) {
			args := v.(LookupFirewallserviceCategoryArgs)
			r, err := LookupFirewallserviceCategory(ctx, &args, opts...)
			var s LookupFirewallserviceCategoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallserviceCategoryResultOutput)
}

// A collection of arguments for invoking getFirewallserviceCategory.
type LookupFirewallserviceCategoryOutputArgs struct {
	// Specify the name of the desired firewallservice category.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallserviceCategoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallserviceCategoryArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallserviceCategory.
type LookupFirewallserviceCategoryResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallserviceCategoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallserviceCategoryResult)(nil)).Elem()
}

func (o LookupFirewallserviceCategoryResultOutput) ToLookupFirewallserviceCategoryResultOutput() LookupFirewallserviceCategoryResultOutput {
	return o
}

func (o LookupFirewallserviceCategoryResultOutput) ToLookupFirewallserviceCategoryResultOutputWithContext(ctx context.Context) LookupFirewallserviceCategoryResultOutput {
	return o
}

// Comment.
func (o LookupFirewallserviceCategoryResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallserviceCategoryResult) string { return v.Comment }).(pulumi.StringOutput)
}

// Security Fabric global object setting.
func (o LookupFirewallserviceCategoryResultOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallserviceCategoryResult) string { return v.FabricObject }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallserviceCategoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallserviceCategoryResult) string { return v.Id }).(pulumi.StringOutput)
}

// Service category name.
func (o LookupFirewallserviceCategoryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallserviceCategoryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFirewallserviceCategoryResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallserviceCategoryResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallserviceCategoryResultOutput{})
}
