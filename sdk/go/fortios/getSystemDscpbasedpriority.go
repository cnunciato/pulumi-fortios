// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system dscpbasedpriority
func LookupSystemDscpbasedpriority(ctx *pulumi.Context, args *LookupSystemDscpbasedpriorityArgs, opts ...pulumi.InvokeOption) (*LookupSystemDscpbasedpriorityResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemDscpbasedpriorityResult
	err := ctx.Invoke("fortios:index/getSystemDscpbasedpriority:getSystemDscpbasedpriority", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemDscpbasedpriority.
type LookupSystemDscpbasedpriorityArgs struct {
	// Specify the fosid of the desired system dscpbasedpriority.
	Fosid int `pulumi:"fosid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemDscpbasedpriority.
type LookupSystemDscpbasedpriorityResult struct {
	// DSCP(DiffServ) DS value (0 - 63).
	Ds int `pulumi:"ds"`
	// Item ID.
	Fosid int `pulumi:"fosid"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// DSCP based priority level.
	Priority  string  `pulumi:"priority"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemDscpbasedpriorityOutput(ctx *pulumi.Context, args LookupSystemDscpbasedpriorityOutputArgs, opts ...pulumi.InvokeOption) LookupSystemDscpbasedpriorityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemDscpbasedpriorityResult, error) {
			args := v.(LookupSystemDscpbasedpriorityArgs)
			r, err := LookupSystemDscpbasedpriority(ctx, &args, opts...)
			var s LookupSystemDscpbasedpriorityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemDscpbasedpriorityResultOutput)
}

// A collection of arguments for invoking getSystemDscpbasedpriority.
type LookupSystemDscpbasedpriorityOutputArgs struct {
	// Specify the fosid of the desired system dscpbasedpriority.
	Fosid pulumi.IntInput `pulumi:"fosid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemDscpbasedpriorityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemDscpbasedpriorityArgs)(nil)).Elem()
}

// A collection of values returned by getSystemDscpbasedpriority.
type LookupSystemDscpbasedpriorityResultOutput struct{ *pulumi.OutputState }

func (LookupSystemDscpbasedpriorityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemDscpbasedpriorityResult)(nil)).Elem()
}

func (o LookupSystemDscpbasedpriorityResultOutput) ToLookupSystemDscpbasedpriorityResultOutput() LookupSystemDscpbasedpriorityResultOutput {
	return o
}

func (o LookupSystemDscpbasedpriorityResultOutput) ToLookupSystemDscpbasedpriorityResultOutputWithContext(ctx context.Context) LookupSystemDscpbasedpriorityResultOutput {
	return o
}

// DSCP(DiffServ) DS value (0 - 63).
func (o LookupSystemDscpbasedpriorityResultOutput) Ds() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemDscpbasedpriorityResult) int { return v.Ds }).(pulumi.IntOutput)
}

// Item ID.
func (o LookupSystemDscpbasedpriorityResultOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemDscpbasedpriorityResult) int { return v.Fosid }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemDscpbasedpriorityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDscpbasedpriorityResult) string { return v.Id }).(pulumi.StringOutput)
}

// DSCP based priority level.
func (o LookupSystemDscpbasedpriorityResultOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDscpbasedpriorityResult) string { return v.Priority }).(pulumi.StringOutput)
}

func (o LookupSystemDscpbasedpriorityResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemDscpbasedpriorityResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemDscpbasedpriorityResultOutput{})
}
