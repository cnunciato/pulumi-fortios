// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Remote certificate as a PEM file.
//
// ## Import
//
// # VpnCertificate Remote can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/vpncertificateRemote:VpncertificateRemote labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/vpncertificateRemote:VpncertificateRemote labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type VpncertificateRemote struct {
	pulumi.CustomResourceState

	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringOutput `pulumi:"range"`
	// Remote certificate.
	Remote pulumi.StringOutput `pulumi:"remote"`
	// Remote certificate source type.
	Source pulumi.StringOutput `pulumi:"source"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewVpncertificateRemote registers a new resource with the given unique name, arguments, and options.
func NewVpncertificateRemote(ctx *pulumi.Context,
	name string, args *VpncertificateRemoteArgs, opts ...pulumi.ResourceOption) (*VpncertificateRemote, error) {
	if args == nil {
		args = &VpncertificateRemoteArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource VpncertificateRemote
	err := ctx.RegisterResource("fortios:index/vpncertificateRemote:VpncertificateRemote", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpncertificateRemote gets an existing VpncertificateRemote resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpncertificateRemote(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpncertificateRemoteState, opts ...pulumi.ResourceOption) (*VpncertificateRemote, error) {
	var resource VpncertificateRemote
	err := ctx.ReadResource("fortios:index/vpncertificateRemote:VpncertificateRemote", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpncertificateRemote resources.
type vpncertificateRemoteState struct {
	// Name.
	Name *string `pulumi:"name"`
	// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
	Range *string `pulumi:"range"`
	// Remote certificate.
	Remote *string `pulumi:"remote"`
	// Remote certificate source type.
	Source *string `pulumi:"source"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type VpncertificateRemoteState struct {
	// Name.
	Name pulumi.StringPtrInput
	// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringPtrInput
	// Remote certificate.
	Remote pulumi.StringPtrInput
	// Remote certificate source type.
	Source pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VpncertificateRemoteState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpncertificateRemoteState)(nil)).Elem()
}

type vpncertificateRemoteArgs struct {
	// Name.
	Name *string `pulumi:"name"`
	// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
	Range *string `pulumi:"range"`
	// Remote certificate.
	Remote *string `pulumi:"remote"`
	// Remote certificate source type.
	Source *string `pulumi:"source"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a VpncertificateRemote resource.
type VpncertificateRemoteArgs struct {
	// Name.
	Name pulumi.StringPtrInput
	// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringPtrInput
	// Remote certificate.
	Remote pulumi.StringPtrInput
	// Remote certificate source type.
	Source pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VpncertificateRemoteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpncertificateRemoteArgs)(nil)).Elem()
}

type VpncertificateRemoteInput interface {
	pulumi.Input

	ToVpncertificateRemoteOutput() VpncertificateRemoteOutput
	ToVpncertificateRemoteOutputWithContext(ctx context.Context) VpncertificateRemoteOutput
}

func (*VpncertificateRemote) ElementType() reflect.Type {
	return reflect.TypeOf((**VpncertificateRemote)(nil)).Elem()
}

func (i *VpncertificateRemote) ToVpncertificateRemoteOutput() VpncertificateRemoteOutput {
	return i.ToVpncertificateRemoteOutputWithContext(context.Background())
}

func (i *VpncertificateRemote) ToVpncertificateRemoteOutputWithContext(ctx context.Context) VpncertificateRemoteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpncertificateRemoteOutput)
}

// VpncertificateRemoteArrayInput is an input type that accepts VpncertificateRemoteArray and VpncertificateRemoteArrayOutput values.
// You can construct a concrete instance of `VpncertificateRemoteArrayInput` via:
//
//	VpncertificateRemoteArray{ VpncertificateRemoteArgs{...} }
type VpncertificateRemoteArrayInput interface {
	pulumi.Input

	ToVpncertificateRemoteArrayOutput() VpncertificateRemoteArrayOutput
	ToVpncertificateRemoteArrayOutputWithContext(context.Context) VpncertificateRemoteArrayOutput
}

type VpncertificateRemoteArray []VpncertificateRemoteInput

func (VpncertificateRemoteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpncertificateRemote)(nil)).Elem()
}

func (i VpncertificateRemoteArray) ToVpncertificateRemoteArrayOutput() VpncertificateRemoteArrayOutput {
	return i.ToVpncertificateRemoteArrayOutputWithContext(context.Background())
}

func (i VpncertificateRemoteArray) ToVpncertificateRemoteArrayOutputWithContext(ctx context.Context) VpncertificateRemoteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpncertificateRemoteArrayOutput)
}

// VpncertificateRemoteMapInput is an input type that accepts VpncertificateRemoteMap and VpncertificateRemoteMapOutput values.
// You can construct a concrete instance of `VpncertificateRemoteMapInput` via:
//
//	VpncertificateRemoteMap{ "key": VpncertificateRemoteArgs{...} }
type VpncertificateRemoteMapInput interface {
	pulumi.Input

	ToVpncertificateRemoteMapOutput() VpncertificateRemoteMapOutput
	ToVpncertificateRemoteMapOutputWithContext(context.Context) VpncertificateRemoteMapOutput
}

type VpncertificateRemoteMap map[string]VpncertificateRemoteInput

func (VpncertificateRemoteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpncertificateRemote)(nil)).Elem()
}

func (i VpncertificateRemoteMap) ToVpncertificateRemoteMapOutput() VpncertificateRemoteMapOutput {
	return i.ToVpncertificateRemoteMapOutputWithContext(context.Background())
}

func (i VpncertificateRemoteMap) ToVpncertificateRemoteMapOutputWithContext(ctx context.Context) VpncertificateRemoteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpncertificateRemoteMapOutput)
}

type VpncertificateRemoteOutput struct{ *pulumi.OutputState }

func (VpncertificateRemoteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpncertificateRemote)(nil)).Elem()
}

func (o VpncertificateRemoteOutput) ToVpncertificateRemoteOutput() VpncertificateRemoteOutput {
	return o
}

func (o VpncertificateRemoteOutput) ToVpncertificateRemoteOutputWithContext(ctx context.Context) VpncertificateRemoteOutput {
	return o
}

// Name.
func (o VpncertificateRemoteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateRemote) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
func (o VpncertificateRemoteOutput) Range() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateRemote) pulumi.StringOutput { return v.Range }).(pulumi.StringOutput)
}

// Remote certificate.
func (o VpncertificateRemoteOutput) Remote() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateRemote) pulumi.StringOutput { return v.Remote }).(pulumi.StringOutput)
}

// Remote certificate source type.
func (o VpncertificateRemoteOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateRemote) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o VpncertificateRemoteOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpncertificateRemote) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type VpncertificateRemoteArrayOutput struct{ *pulumi.OutputState }

func (VpncertificateRemoteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpncertificateRemote)(nil)).Elem()
}

func (o VpncertificateRemoteArrayOutput) ToVpncertificateRemoteArrayOutput() VpncertificateRemoteArrayOutput {
	return o
}

func (o VpncertificateRemoteArrayOutput) ToVpncertificateRemoteArrayOutputWithContext(ctx context.Context) VpncertificateRemoteArrayOutput {
	return o
}

func (o VpncertificateRemoteArrayOutput) Index(i pulumi.IntInput) VpncertificateRemoteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpncertificateRemote {
		return vs[0].([]*VpncertificateRemote)[vs[1].(int)]
	}).(VpncertificateRemoteOutput)
}

type VpncertificateRemoteMapOutput struct{ *pulumi.OutputState }

func (VpncertificateRemoteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpncertificateRemote)(nil)).Elem()
}

func (o VpncertificateRemoteMapOutput) ToVpncertificateRemoteMapOutput() VpncertificateRemoteMapOutput {
	return o
}

func (o VpncertificateRemoteMapOutput) ToVpncertificateRemoteMapOutputWithContext(ctx context.Context) VpncertificateRemoteMapOutput {
	return o
}

func (o VpncertificateRemoteMapOutput) MapIndex(k pulumi.StringInput) VpncertificateRemoteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpncertificateRemote {
		return vs[0].(map[string]*VpncertificateRemote)[vs[1].(string)]
	}).(VpncertificateRemoteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpncertificateRemoteInput)(nil)).Elem(), &VpncertificateRemote{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpncertificateRemoteArrayInput)(nil)).Elem(), VpncertificateRemoteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpncertificateRemoteMapInput)(nil)).Elem(), VpncertificateRemoteMap{})
	pulumi.RegisterOutputType(VpncertificateRemoteOutput{})
	pulumi.RegisterOutputType(VpncertificateRemoteArrayOutput{})
	pulumi.RegisterOutputType(VpncertificateRemoteMapOutput{})
}
