// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Replacement messages.
//
// ## Import
//
// # SystemReplacemsg Spam can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemreplacemsgSpam:SystemreplacemsgSpam labelname {{msg_type}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemreplacemsgSpam:SystemreplacemsgSpam labelname {{msg_type}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemreplacemsgSpam struct {
	pulumi.CustomResourceState

	// Message string.
	Buffer pulumi.StringPtrOutput `pulumi:"buffer"`
	// Format flag.
	Format pulumi.StringOutput `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringOutput `pulumi:"header"`
	// Message type.
	MsgType pulumi.StringOutput `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemreplacemsgSpam registers a new resource with the given unique name, arguments, and options.
func NewSystemreplacemsgSpam(ctx *pulumi.Context,
	name string, args *SystemreplacemsgSpamArgs, opts ...pulumi.ResourceOption) (*SystemreplacemsgSpam, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemreplacemsgSpam
	err := ctx.RegisterResource("fortios:index/systemreplacemsgSpam:SystemreplacemsgSpam", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemreplacemsgSpam gets an existing SystemreplacemsgSpam resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemreplacemsgSpam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemreplacemsgSpamState, opts ...pulumi.ResourceOption) (*SystemreplacemsgSpam, error) {
	var resource SystemreplacemsgSpam
	err := ctx.ReadResource("fortios:index/systemreplacemsgSpam:SystemreplacemsgSpam", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemreplacemsgSpam resources.
type systemreplacemsgSpamState struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType *string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemreplacemsgSpamState struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemreplacemsgSpamState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemreplacemsgSpamState)(nil)).Elem()
}

type systemreplacemsgSpamArgs struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemreplacemsgSpam resource.
type SystemreplacemsgSpamArgs struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemreplacemsgSpamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemreplacemsgSpamArgs)(nil)).Elem()
}

type SystemreplacemsgSpamInput interface {
	pulumi.Input

	ToSystemreplacemsgSpamOutput() SystemreplacemsgSpamOutput
	ToSystemreplacemsgSpamOutputWithContext(ctx context.Context) SystemreplacemsgSpamOutput
}

func (*SystemreplacemsgSpam) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemreplacemsgSpam)(nil)).Elem()
}

func (i *SystemreplacemsgSpam) ToSystemreplacemsgSpamOutput() SystemreplacemsgSpamOutput {
	return i.ToSystemreplacemsgSpamOutputWithContext(context.Background())
}

func (i *SystemreplacemsgSpam) ToSystemreplacemsgSpamOutputWithContext(ctx context.Context) SystemreplacemsgSpamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgSpamOutput)
}

// SystemreplacemsgSpamArrayInput is an input type that accepts SystemreplacemsgSpamArray and SystemreplacemsgSpamArrayOutput values.
// You can construct a concrete instance of `SystemreplacemsgSpamArrayInput` via:
//
//	SystemreplacemsgSpamArray{ SystemreplacemsgSpamArgs{...} }
type SystemreplacemsgSpamArrayInput interface {
	pulumi.Input

	ToSystemreplacemsgSpamArrayOutput() SystemreplacemsgSpamArrayOutput
	ToSystemreplacemsgSpamArrayOutputWithContext(context.Context) SystemreplacemsgSpamArrayOutput
}

type SystemreplacemsgSpamArray []SystemreplacemsgSpamInput

func (SystemreplacemsgSpamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemreplacemsgSpam)(nil)).Elem()
}

func (i SystemreplacemsgSpamArray) ToSystemreplacemsgSpamArrayOutput() SystemreplacemsgSpamArrayOutput {
	return i.ToSystemreplacemsgSpamArrayOutputWithContext(context.Background())
}

func (i SystemreplacemsgSpamArray) ToSystemreplacemsgSpamArrayOutputWithContext(ctx context.Context) SystemreplacemsgSpamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgSpamArrayOutput)
}

// SystemreplacemsgSpamMapInput is an input type that accepts SystemreplacemsgSpamMap and SystemreplacemsgSpamMapOutput values.
// You can construct a concrete instance of `SystemreplacemsgSpamMapInput` via:
//
//	SystemreplacemsgSpamMap{ "key": SystemreplacemsgSpamArgs{...} }
type SystemreplacemsgSpamMapInput interface {
	pulumi.Input

	ToSystemreplacemsgSpamMapOutput() SystemreplacemsgSpamMapOutput
	ToSystemreplacemsgSpamMapOutputWithContext(context.Context) SystemreplacemsgSpamMapOutput
}

type SystemreplacemsgSpamMap map[string]SystemreplacemsgSpamInput

func (SystemreplacemsgSpamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemreplacemsgSpam)(nil)).Elem()
}

func (i SystemreplacemsgSpamMap) ToSystemreplacemsgSpamMapOutput() SystemreplacemsgSpamMapOutput {
	return i.ToSystemreplacemsgSpamMapOutputWithContext(context.Background())
}

func (i SystemreplacemsgSpamMap) ToSystemreplacemsgSpamMapOutputWithContext(ctx context.Context) SystemreplacemsgSpamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgSpamMapOutput)
}

type SystemreplacemsgSpamOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgSpamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemreplacemsgSpam)(nil)).Elem()
}

func (o SystemreplacemsgSpamOutput) ToSystemreplacemsgSpamOutput() SystemreplacemsgSpamOutput {
	return o
}

func (o SystemreplacemsgSpamOutput) ToSystemreplacemsgSpamOutputWithContext(ctx context.Context) SystemreplacemsgSpamOutput {
	return o
}

// Message string.
func (o SystemreplacemsgSpamOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemreplacemsgSpam) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

// Format flag.
func (o SystemreplacemsgSpamOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgSpam) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Header flag. Valid values: `none`, `http`, `8bit`.
func (o SystemreplacemsgSpamOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgSpam) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

// Message type.
func (o SystemreplacemsgSpamOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgSpam) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemreplacemsgSpamOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemreplacemsgSpam) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemreplacemsgSpamArrayOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgSpamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemreplacemsgSpam)(nil)).Elem()
}

func (o SystemreplacemsgSpamArrayOutput) ToSystemreplacemsgSpamArrayOutput() SystemreplacemsgSpamArrayOutput {
	return o
}

func (o SystemreplacemsgSpamArrayOutput) ToSystemreplacemsgSpamArrayOutputWithContext(ctx context.Context) SystemreplacemsgSpamArrayOutput {
	return o
}

func (o SystemreplacemsgSpamArrayOutput) Index(i pulumi.IntInput) SystemreplacemsgSpamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemreplacemsgSpam {
		return vs[0].([]*SystemreplacemsgSpam)[vs[1].(int)]
	}).(SystemreplacemsgSpamOutput)
}

type SystemreplacemsgSpamMapOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgSpamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemreplacemsgSpam)(nil)).Elem()
}

func (o SystemreplacemsgSpamMapOutput) ToSystemreplacemsgSpamMapOutput() SystemreplacemsgSpamMapOutput {
	return o
}

func (o SystemreplacemsgSpamMapOutput) ToSystemreplacemsgSpamMapOutputWithContext(ctx context.Context) SystemreplacemsgSpamMapOutput {
	return o
}

func (o SystemreplacemsgSpamMapOutput) MapIndex(k pulumi.StringInput) SystemreplacemsgSpamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemreplacemsgSpam {
		return vs[0].(map[string]*SystemreplacemsgSpam)[vs[1].(string)]
	}).(SystemreplacemsgSpamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgSpamInput)(nil)).Elem(), &SystemreplacemsgSpam{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgSpamArrayInput)(nil)).Elem(), SystemreplacemsgSpamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgSpamMapInput)(nil)).Elem(), SystemreplacemsgSpamMap{})
	pulumi.RegisterOutputType(SystemreplacemsgSpamOutput{})
	pulumi.RegisterOutputType(SystemreplacemsgSpamArrayOutput{})
	pulumi.RegisterOutputType(SystemreplacemsgSpamMapOutput{})
}
