// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system alias
func LookupSystemAlias(ctx *pulumi.Context, args *LookupSystemAliasArgs, opts ...pulumi.InvokeOption) (*LookupSystemAliasResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemAliasResult
	err := ctx.Invoke("fortios:index/getSystemAlias:getSystemAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemAlias.
type LookupSystemAliasArgs struct {
	// Specify the name of the desired system alias.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemAlias.
type LookupSystemAliasResult struct {
	// Command list to execute.
	Command string `pulumi:"command"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Alias command name.
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemAliasOutput(ctx *pulumi.Context, args LookupSystemAliasOutputArgs, opts ...pulumi.InvokeOption) LookupSystemAliasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemAliasResult, error) {
			args := v.(LookupSystemAliasArgs)
			r, err := LookupSystemAlias(ctx, &args, opts...)
			var s LookupSystemAliasResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemAliasResultOutput)
}

// A collection of arguments for invoking getSystemAlias.
type LookupSystemAliasOutputArgs struct {
	// Specify the name of the desired system alias.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemAliasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemAliasArgs)(nil)).Elem()
}

// A collection of values returned by getSystemAlias.
type LookupSystemAliasResultOutput struct{ *pulumi.OutputState }

func (LookupSystemAliasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemAliasResult)(nil)).Elem()
}

func (o LookupSystemAliasResultOutput) ToLookupSystemAliasResultOutput() LookupSystemAliasResultOutput {
	return o
}

func (o LookupSystemAliasResultOutput) ToLookupSystemAliasResultOutputWithContext(ctx context.Context) LookupSystemAliasResultOutput {
	return o
}

// Command list to execute.
func (o LookupSystemAliasResultOutput) Command() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAliasResult) string { return v.Command }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemAliasResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAliasResult) string { return v.Id }).(pulumi.StringOutput)
}

// Alias command name.
func (o LookupSystemAliasResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAliasResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSystemAliasResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemAliasResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemAliasResultOutput{})
}
