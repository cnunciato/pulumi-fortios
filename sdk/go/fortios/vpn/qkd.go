// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure Quantum Key Distribution servers Applies to FortiOS Version `>= 7.4.2`.
//
// ## Import
//
// Vpn Qkd can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:vpn/qkd:Qkd labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:vpn/qkd:Qkd labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Qkd struct {
	pulumi.CustomResourceState

	// Names of up to 4 certificates to offer to the KME. The structure of `certificate` block is documented below.
	Certificates QkdCertificateArrayOutput `pulumi:"certificates"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Quantum Key Distribution ID assigned by the KME.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	// Quantum Key Distribution configuration name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Authenticate Quantum Key Device's certificate with the peer/peergrp.
	Peer pulumi.StringOutput `pulumi:"peer"`
	// Port to connect to on the KME.
	Port pulumi.IntOutput `pulumi:"port"`
	// IPv4, IPv6 or DNS address of the KME.
	Server pulumi.StringOutput `pulumi:"server"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewQkd registers a new resource with the given unique name, arguments, and options.
func NewQkd(ctx *pulumi.Context,
	name string, args *QkdArgs, opts ...pulumi.ResourceOption) (*Qkd, error) {
	if args == nil {
		args = &QkdArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Qkd
	err := ctx.RegisterResource("fortios:vpn/qkd:Qkd", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQkd gets an existing Qkd resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQkd(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QkdState, opts ...pulumi.ResourceOption) (*Qkd, error) {
	var resource Qkd
	err := ctx.ReadResource("fortios:vpn/qkd:Qkd", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Qkd resources.
type qkdState struct {
	// Names of up to 4 certificates to offer to the KME. The structure of `certificate` block is documented below.
	Certificates []QkdCertificate `pulumi:"certificates"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Quantum Key Distribution ID assigned by the KME.
	Fosid *string `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Quantum Key Distribution configuration name.
	Name *string `pulumi:"name"`
	// Authenticate Quantum Key Device's certificate with the peer/peergrp.
	Peer *string `pulumi:"peer"`
	// Port to connect to on the KME.
	Port *int `pulumi:"port"`
	// IPv4, IPv6 or DNS address of the KME.
	Server *string `pulumi:"server"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type QkdState struct {
	// Names of up to 4 certificates to offer to the KME. The structure of `certificate` block is documented below.
	Certificates QkdCertificateArrayInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Quantum Key Distribution ID assigned by the KME.
	Fosid pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Quantum Key Distribution configuration name.
	Name pulumi.StringPtrInput
	// Authenticate Quantum Key Device's certificate with the peer/peergrp.
	Peer pulumi.StringPtrInput
	// Port to connect to on the KME.
	Port pulumi.IntPtrInput
	// IPv4, IPv6 or DNS address of the KME.
	Server pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (QkdState) ElementType() reflect.Type {
	return reflect.TypeOf((*qkdState)(nil)).Elem()
}

type qkdArgs struct {
	// Names of up to 4 certificates to offer to the KME. The structure of `certificate` block is documented below.
	Certificates []QkdCertificate `pulumi:"certificates"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Quantum Key Distribution ID assigned by the KME.
	Fosid *string `pulumi:"fosid"`
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables *string `pulumi:"getAllTables"`
	// Quantum Key Distribution configuration name.
	Name *string `pulumi:"name"`
	// Authenticate Quantum Key Device's certificate with the peer/peergrp.
	Peer *string `pulumi:"peer"`
	// Port to connect to on the KME.
	Port *int `pulumi:"port"`
	// IPv4, IPv6 or DNS address of the KME.
	Server *string `pulumi:"server"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Qkd resource.
type QkdArgs struct {
	// Names of up to 4 certificates to offer to the KME. The structure of `certificate` block is documented below.
	Certificates QkdCertificateArrayInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Quantum Key Distribution ID assigned by the KME.
	Fosid pulumi.StringPtrInput
	// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
	GetAllTables pulumi.StringPtrInput
	// Quantum Key Distribution configuration name.
	Name pulumi.StringPtrInput
	// Authenticate Quantum Key Device's certificate with the peer/peergrp.
	Peer pulumi.StringPtrInput
	// Port to connect to on the KME.
	Port pulumi.IntPtrInput
	// IPv4, IPv6 or DNS address of the KME.
	Server pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (QkdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*qkdArgs)(nil)).Elem()
}

type QkdInput interface {
	pulumi.Input

	ToQkdOutput() QkdOutput
	ToQkdOutputWithContext(ctx context.Context) QkdOutput
}

func (*Qkd) ElementType() reflect.Type {
	return reflect.TypeOf((**Qkd)(nil)).Elem()
}

func (i *Qkd) ToQkdOutput() QkdOutput {
	return i.ToQkdOutputWithContext(context.Background())
}

func (i *Qkd) ToQkdOutputWithContext(ctx context.Context) QkdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QkdOutput)
}

// QkdArrayInput is an input type that accepts QkdArray and QkdArrayOutput values.
// You can construct a concrete instance of `QkdArrayInput` via:
//
//	QkdArray{ QkdArgs{...} }
type QkdArrayInput interface {
	pulumi.Input

	ToQkdArrayOutput() QkdArrayOutput
	ToQkdArrayOutputWithContext(context.Context) QkdArrayOutput
}

type QkdArray []QkdInput

func (QkdArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Qkd)(nil)).Elem()
}

func (i QkdArray) ToQkdArrayOutput() QkdArrayOutput {
	return i.ToQkdArrayOutputWithContext(context.Background())
}

func (i QkdArray) ToQkdArrayOutputWithContext(ctx context.Context) QkdArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QkdArrayOutput)
}

// QkdMapInput is an input type that accepts QkdMap and QkdMapOutput values.
// You can construct a concrete instance of `QkdMapInput` via:
//
//	QkdMap{ "key": QkdArgs{...} }
type QkdMapInput interface {
	pulumi.Input

	ToQkdMapOutput() QkdMapOutput
	ToQkdMapOutputWithContext(context.Context) QkdMapOutput
}

type QkdMap map[string]QkdInput

func (QkdMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Qkd)(nil)).Elem()
}

func (i QkdMap) ToQkdMapOutput() QkdMapOutput {
	return i.ToQkdMapOutputWithContext(context.Background())
}

func (i QkdMap) ToQkdMapOutputWithContext(ctx context.Context) QkdMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QkdMapOutput)
}

type QkdOutput struct{ *pulumi.OutputState }

func (QkdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Qkd)(nil)).Elem()
}

func (o QkdOutput) ToQkdOutput() QkdOutput {
	return o
}

func (o QkdOutput) ToQkdOutputWithContext(ctx context.Context) QkdOutput {
	return o
}

// Names of up to 4 certificates to offer to the KME. The structure of `certificate` block is documented below.
func (o QkdOutput) Certificates() QkdCertificateArrayOutput {
	return o.ApplyT(func(v *Qkd) QkdCertificateArrayOutput { return v.Certificates }).(QkdCertificateArrayOutput)
}

// Comment.
func (o QkdOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Qkd) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o QkdOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Qkd) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Quantum Key Distribution ID assigned by the KME.
func (o QkdOutput) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v *Qkd) pulumi.StringOutput { return v.Fosid }).(pulumi.StringOutput)
}

// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
func (o QkdOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Qkd) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

// Quantum Key Distribution configuration name.
func (o QkdOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Qkd) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Authenticate Quantum Key Device's certificate with the peer/peergrp.
func (o QkdOutput) Peer() pulumi.StringOutput {
	return o.ApplyT(func(v *Qkd) pulumi.StringOutput { return v.Peer }).(pulumi.StringOutput)
}

// Port to connect to on the KME.
func (o QkdOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Qkd) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// IPv4, IPv6 or DNS address of the KME.
func (o QkdOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *Qkd) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o QkdOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Qkd) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type QkdArrayOutput struct{ *pulumi.OutputState }

func (QkdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Qkd)(nil)).Elem()
}

func (o QkdArrayOutput) ToQkdArrayOutput() QkdArrayOutput {
	return o
}

func (o QkdArrayOutput) ToQkdArrayOutputWithContext(ctx context.Context) QkdArrayOutput {
	return o
}

func (o QkdArrayOutput) Index(i pulumi.IntInput) QkdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Qkd {
		return vs[0].([]*Qkd)[vs[1].(int)]
	}).(QkdOutput)
}

type QkdMapOutput struct{ *pulumi.OutputState }

func (QkdMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Qkd)(nil)).Elem()
}

func (o QkdMapOutput) ToQkdMapOutput() QkdMapOutput {
	return o
}

func (o QkdMapOutput) ToQkdMapOutputWithContext(ctx context.Context) QkdMapOutput {
	return o
}

func (o QkdMapOutput) MapIndex(k pulumi.StringInput) QkdOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Qkd {
		return vs[0].(map[string]*Qkd)[vs[1].(string)]
	}).(QkdOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QkdInput)(nil)).Elem(), &Qkd{})
	pulumi.RegisterInputType(reflect.TypeOf((*QkdArrayInput)(nil)).Elem(), QkdArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QkdMapInput)(nil)).Elem(), QkdMap{})
	pulumi.RegisterOutputType(QkdOutput{})
	pulumi.RegisterOutputType(QkdArrayOutput{})
	pulumi.RegisterOutputType(QkdMapOutput{})
}