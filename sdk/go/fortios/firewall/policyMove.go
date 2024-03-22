// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

type PolicyMove struct {
	pulumi.CustomResourceState

	Comment              pulumi.StringPtrOutput `pulumi:"comment"`
	Move                 pulumi.StringOutput    `pulumi:"move"`
	PolicyidDst          pulumi.IntOutput       `pulumi:"policyidDst"`
	PolicyidSrc          pulumi.IntOutput       `pulumi:"policyidSrc"`
	StatePolicySrcdstPos pulumi.StringPtrOutput `pulumi:"statePolicySrcdstPos"`
	Vdomparam            pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewPolicyMove registers a new resource with the given unique name, arguments, and options.
func NewPolicyMove(ctx *pulumi.Context,
	name string, args *PolicyMoveArgs, opts ...pulumi.ResourceOption) (*PolicyMove, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Move == nil {
		return nil, errors.New("invalid value for required argument 'Move'")
	}
	if args.PolicyidDst == nil {
		return nil, errors.New("invalid value for required argument 'PolicyidDst'")
	}
	if args.PolicyidSrc == nil {
		return nil, errors.New("invalid value for required argument 'PolicyidSrc'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PolicyMove
	err := ctx.RegisterResource("fortios:firewall/policyMove:PolicyMove", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyMove gets an existing PolicyMove resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyMove(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyMoveState, opts ...pulumi.ResourceOption) (*PolicyMove, error) {
	var resource PolicyMove
	err := ctx.ReadResource("fortios:firewall/policyMove:PolicyMove", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyMove resources.
type policyMoveState struct {
	Comment              *string `pulumi:"comment"`
	Move                 *string `pulumi:"move"`
	PolicyidDst          *int    `pulumi:"policyidDst"`
	PolicyidSrc          *int    `pulumi:"policyidSrc"`
	StatePolicySrcdstPos *string `pulumi:"statePolicySrcdstPos"`
	Vdomparam            *string `pulumi:"vdomparam"`
}

type PolicyMoveState struct {
	Comment              pulumi.StringPtrInput
	Move                 pulumi.StringPtrInput
	PolicyidDst          pulumi.IntPtrInput
	PolicyidSrc          pulumi.IntPtrInput
	StatePolicySrcdstPos pulumi.StringPtrInput
	Vdomparam            pulumi.StringPtrInput
}

func (PolicyMoveState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyMoveState)(nil)).Elem()
}

type policyMoveArgs struct {
	Comment              *string `pulumi:"comment"`
	Move                 string  `pulumi:"move"`
	PolicyidDst          int     `pulumi:"policyidDst"`
	PolicyidSrc          int     `pulumi:"policyidSrc"`
	StatePolicySrcdstPos *string `pulumi:"statePolicySrcdstPos"`
	Vdomparam            *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a PolicyMove resource.
type PolicyMoveArgs struct {
	Comment              pulumi.StringPtrInput
	Move                 pulumi.StringInput
	PolicyidDst          pulumi.IntInput
	PolicyidSrc          pulumi.IntInput
	StatePolicySrcdstPos pulumi.StringPtrInput
	Vdomparam            pulumi.StringPtrInput
}

func (PolicyMoveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyMoveArgs)(nil)).Elem()
}

type PolicyMoveInput interface {
	pulumi.Input

	ToPolicyMoveOutput() PolicyMoveOutput
	ToPolicyMoveOutputWithContext(ctx context.Context) PolicyMoveOutput
}

func (*PolicyMove) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyMove)(nil)).Elem()
}

func (i *PolicyMove) ToPolicyMoveOutput() PolicyMoveOutput {
	return i.ToPolicyMoveOutputWithContext(context.Background())
}

func (i *PolicyMove) ToPolicyMoveOutputWithContext(ctx context.Context) PolicyMoveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyMoveOutput)
}

// PolicyMoveArrayInput is an input type that accepts PolicyMoveArray and PolicyMoveArrayOutput values.
// You can construct a concrete instance of `PolicyMoveArrayInput` via:
//
//	PolicyMoveArray{ PolicyMoveArgs{...} }
type PolicyMoveArrayInput interface {
	pulumi.Input

	ToPolicyMoveArrayOutput() PolicyMoveArrayOutput
	ToPolicyMoveArrayOutputWithContext(context.Context) PolicyMoveArrayOutput
}

type PolicyMoveArray []PolicyMoveInput

func (PolicyMoveArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyMove)(nil)).Elem()
}

func (i PolicyMoveArray) ToPolicyMoveArrayOutput() PolicyMoveArrayOutput {
	return i.ToPolicyMoveArrayOutputWithContext(context.Background())
}

func (i PolicyMoveArray) ToPolicyMoveArrayOutputWithContext(ctx context.Context) PolicyMoveArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyMoveArrayOutput)
}

// PolicyMoveMapInput is an input type that accepts PolicyMoveMap and PolicyMoveMapOutput values.
// You can construct a concrete instance of `PolicyMoveMapInput` via:
//
//	PolicyMoveMap{ "key": PolicyMoveArgs{...} }
type PolicyMoveMapInput interface {
	pulumi.Input

	ToPolicyMoveMapOutput() PolicyMoveMapOutput
	ToPolicyMoveMapOutputWithContext(context.Context) PolicyMoveMapOutput
}

type PolicyMoveMap map[string]PolicyMoveInput

func (PolicyMoveMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyMove)(nil)).Elem()
}

func (i PolicyMoveMap) ToPolicyMoveMapOutput() PolicyMoveMapOutput {
	return i.ToPolicyMoveMapOutputWithContext(context.Background())
}

func (i PolicyMoveMap) ToPolicyMoveMapOutputWithContext(ctx context.Context) PolicyMoveMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyMoveMapOutput)
}

type PolicyMoveOutput struct{ *pulumi.OutputState }

func (PolicyMoveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyMove)(nil)).Elem()
}

func (o PolicyMoveOutput) ToPolicyMoveOutput() PolicyMoveOutput {
	return o
}

func (o PolicyMoveOutput) ToPolicyMoveOutputWithContext(ctx context.Context) PolicyMoveOutput {
	return o
}

func (o PolicyMoveOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyMove) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o PolicyMoveOutput) Move() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyMove) pulumi.StringOutput { return v.Move }).(pulumi.StringOutput)
}

func (o PolicyMoveOutput) PolicyidDst() pulumi.IntOutput {
	return o.ApplyT(func(v *PolicyMove) pulumi.IntOutput { return v.PolicyidDst }).(pulumi.IntOutput)
}

func (o PolicyMoveOutput) PolicyidSrc() pulumi.IntOutput {
	return o.ApplyT(func(v *PolicyMove) pulumi.IntOutput { return v.PolicyidSrc }).(pulumi.IntOutput)
}

func (o PolicyMoveOutput) StatePolicySrcdstPos() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyMove) pulumi.StringPtrOutput { return v.StatePolicySrcdstPos }).(pulumi.StringPtrOutput)
}

func (o PolicyMoveOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyMove) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type PolicyMoveArrayOutput struct{ *pulumi.OutputState }

func (PolicyMoveArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyMove)(nil)).Elem()
}

func (o PolicyMoveArrayOutput) ToPolicyMoveArrayOutput() PolicyMoveArrayOutput {
	return o
}

func (o PolicyMoveArrayOutput) ToPolicyMoveArrayOutputWithContext(ctx context.Context) PolicyMoveArrayOutput {
	return o
}

func (o PolicyMoveArrayOutput) Index(i pulumi.IntInput) PolicyMoveOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicyMove {
		return vs[0].([]*PolicyMove)[vs[1].(int)]
	}).(PolicyMoveOutput)
}

type PolicyMoveMapOutput struct{ *pulumi.OutputState }

func (PolicyMoveMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyMove)(nil)).Elem()
}

func (o PolicyMoveMapOutput) ToPolicyMoveMapOutput() PolicyMoveMapOutput {
	return o
}

func (o PolicyMoveMapOutput) ToPolicyMoveMapOutputWithContext(ctx context.Context) PolicyMoveMapOutput {
	return o
}

func (o PolicyMoveMapOutput) MapIndex(k pulumi.StringInput) PolicyMoveOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicyMove {
		return vs[0].(map[string]*PolicyMove)[vs[1].(string)]
	}).(PolicyMoveOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyMoveInput)(nil)).Elem(), &PolicyMove{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyMoveArrayInput)(nil)).Elem(), PolicyMoveArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyMoveMapInput)(nil)).Elem(), PolicyMoveMap{})
	pulumi.RegisterOutputType(PolicyMoveOutput{})
	pulumi.RegisterOutputType(PolicyMoveArrayOutput{})
	pulumi.RegisterOutputType(PolicyMoveMapOutput{})
}