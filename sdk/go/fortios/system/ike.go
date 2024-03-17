// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure IKE global attributes. Applies to FortiOS Version `>= 7.0.1`.
//
// ## Import
//
// System Ike can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/ike:Ike labelname SystemIke
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/ike:Ike labelname SystemIke
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Ike struct {
	pulumi.CustomResourceState

	// Diffie-Hellman group 1 (MODP-768). The structure of `dhGroup1` block is documented below.
	DhGroup1 IkeDhGroup1Output `pulumi:"dhGroup1"`
	// Diffie-Hellman group 14 (MODP-2048). The structure of `dhGroup14` block is documented below.
	DhGroup14 IkeDhGroup14Output `pulumi:"dhGroup14"`
	// Diffie-Hellman group 15 (MODP-3072). The structure of `dhGroup15` block is documented below.
	DhGroup15 IkeDhGroup15Output `pulumi:"dhGroup15"`
	// Diffie-Hellman group 16 (MODP-4096). The structure of `dhGroup16` block is documented below.
	DhGroup16 IkeDhGroup16Output `pulumi:"dhGroup16"`
	// Diffie-Hellman group 17 (MODP-6144). The structure of `dhGroup17` block is documented below.
	DhGroup17 IkeDhGroup17Output `pulumi:"dhGroup17"`
	// Diffie-Hellman group 18 (MODP-8192). The structure of `dhGroup18` block is documented below.
	DhGroup18 IkeDhGroup18Output `pulumi:"dhGroup18"`
	// Diffie-Hellman group 19 (EC-P256). The structure of `dhGroup19` block is documented below.
	DhGroup19 IkeDhGroup19Output `pulumi:"dhGroup19"`
	// Diffie-Hellman group 2 (MODP-1024). The structure of `dhGroup2` block is documented below.
	DhGroup2 IkeDhGroup2Output `pulumi:"dhGroup2"`
	// Diffie-Hellman group 20 (EC-P384). The structure of `dhGroup20` block is documented below.
	DhGroup20 IkeDhGroup20Output `pulumi:"dhGroup20"`
	// Diffie-Hellman group 21 (EC-P521). The structure of `dhGroup21` block is documented below.
	DhGroup21 IkeDhGroup21Output `pulumi:"dhGroup21"`
	// Diffie-Hellman group 27 (EC-P224BP). The structure of `dhGroup27` block is documented below.
	DhGroup27 IkeDhGroup27Output `pulumi:"dhGroup27"`
	// Diffie-Hellman group 28 (EC-P256BP). The structure of `dhGroup28` block is documented below.
	DhGroup28 IkeDhGroup28Output `pulumi:"dhGroup28"`
	// Diffie-Hellman group 29 (EC-P384BP). The structure of `dhGroup29` block is documented below.
	DhGroup29 IkeDhGroup29Output `pulumi:"dhGroup29"`
	// Diffie-Hellman group 30 (EC-P512BP). The structure of `dhGroup30` block is documented below.
	DhGroup30 IkeDhGroup30Output `pulumi:"dhGroup30"`
	// Diffie-Hellman group 31 (EC-X25519). The structure of `dhGroup31` block is documented below.
	DhGroup31 IkeDhGroup31Output `pulumi:"dhGroup31"`
	// Diffie-Hellman group 32 (EC-X448). The structure of `dhGroup32` block is documented below.
	DhGroup32 IkeDhGroup32Output `pulumi:"dhGroup32"`
	// Diffie-Hellman group 5 (MODP-1536). The structure of `dhGroup5` block is documented below.
	DhGroup5 IkeDhGroup5Output `pulumi:"dhGroup5"`
	// Enable/disable Diffie-Hellman key pair cache. Valid values: `enable`, `disable`.
	DhKeypairCache pulumi.StringOutput `pulumi:"dhKeypairCache"`
	// Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).
	DhKeypairCount pulumi.IntOutput `pulumi:"dhKeypairCount"`
	// Enable/disable Diffie-Hellman key pair cache CPU throttling. Valid values: `enable`, `disable`.
	DhKeypairThrottle pulumi.StringOutput `pulumi:"dhKeypairThrottle"`
	// Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations. Valid values: `software`, `hardware`.
	DhMode pulumi.StringOutput `pulumi:"dhMode"`
	// Enable/disable multiprocess Diffie-Hellman daemon for IKE. Valid values: `enable`, `disable`.
	DhMultiprocess pulumi.StringOutput `pulumi:"dhMultiprocess"`
	// Number of Diffie-Hellman workers to start.
	DhWorkerCount pulumi.IntOutput `pulumi:"dhWorkerCount"`
	// Maximum number of IPsec tunnels to negotiate simultaneously.
	EmbryonicLimit pulumi.IntOutput `pulumi:"embryonicLimit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	//
	// The `dhGroup1` block supports:
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewIke registers a new resource with the given unique name, arguments, and options.
func NewIke(ctx *pulumi.Context,
	name string, args *IkeArgs, opts ...pulumi.ResourceOption) (*Ike, error) {
	if args == nil {
		args = &IkeArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ike
	err := ctx.RegisterResource("fortios:system/ike:Ike", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIke gets an existing Ike resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIke(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IkeState, opts ...pulumi.ResourceOption) (*Ike, error) {
	var resource Ike
	err := ctx.ReadResource("fortios:system/ike:Ike", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ike resources.
type ikeState struct {
	// Diffie-Hellman group 1 (MODP-768). The structure of `dhGroup1` block is documented below.
	DhGroup1 *IkeDhGroup1 `pulumi:"dhGroup1"`
	// Diffie-Hellman group 14 (MODP-2048). The structure of `dhGroup14` block is documented below.
	DhGroup14 *IkeDhGroup14 `pulumi:"dhGroup14"`
	// Diffie-Hellman group 15 (MODP-3072). The structure of `dhGroup15` block is documented below.
	DhGroup15 *IkeDhGroup15 `pulumi:"dhGroup15"`
	// Diffie-Hellman group 16 (MODP-4096). The structure of `dhGroup16` block is documented below.
	DhGroup16 *IkeDhGroup16 `pulumi:"dhGroup16"`
	// Diffie-Hellman group 17 (MODP-6144). The structure of `dhGroup17` block is documented below.
	DhGroup17 *IkeDhGroup17 `pulumi:"dhGroup17"`
	// Diffie-Hellman group 18 (MODP-8192). The structure of `dhGroup18` block is documented below.
	DhGroup18 *IkeDhGroup18 `pulumi:"dhGroup18"`
	// Diffie-Hellman group 19 (EC-P256). The structure of `dhGroup19` block is documented below.
	DhGroup19 *IkeDhGroup19 `pulumi:"dhGroup19"`
	// Diffie-Hellman group 2 (MODP-1024). The structure of `dhGroup2` block is documented below.
	DhGroup2 *IkeDhGroup2 `pulumi:"dhGroup2"`
	// Diffie-Hellman group 20 (EC-P384). The structure of `dhGroup20` block is documented below.
	DhGroup20 *IkeDhGroup20 `pulumi:"dhGroup20"`
	// Diffie-Hellman group 21 (EC-P521). The structure of `dhGroup21` block is documented below.
	DhGroup21 *IkeDhGroup21 `pulumi:"dhGroup21"`
	// Diffie-Hellman group 27 (EC-P224BP). The structure of `dhGroup27` block is documented below.
	DhGroup27 *IkeDhGroup27 `pulumi:"dhGroup27"`
	// Diffie-Hellman group 28 (EC-P256BP). The structure of `dhGroup28` block is documented below.
	DhGroup28 *IkeDhGroup28 `pulumi:"dhGroup28"`
	// Diffie-Hellman group 29 (EC-P384BP). The structure of `dhGroup29` block is documented below.
	DhGroup29 *IkeDhGroup29 `pulumi:"dhGroup29"`
	// Diffie-Hellman group 30 (EC-P512BP). The structure of `dhGroup30` block is documented below.
	DhGroup30 *IkeDhGroup30 `pulumi:"dhGroup30"`
	// Diffie-Hellman group 31 (EC-X25519). The structure of `dhGroup31` block is documented below.
	DhGroup31 *IkeDhGroup31 `pulumi:"dhGroup31"`
	// Diffie-Hellman group 32 (EC-X448). The structure of `dhGroup32` block is documented below.
	DhGroup32 *IkeDhGroup32 `pulumi:"dhGroup32"`
	// Diffie-Hellman group 5 (MODP-1536). The structure of `dhGroup5` block is documented below.
	DhGroup5 *IkeDhGroup5 `pulumi:"dhGroup5"`
	// Enable/disable Diffie-Hellman key pair cache. Valid values: `enable`, `disable`.
	DhKeypairCache *string `pulumi:"dhKeypairCache"`
	// Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).
	DhKeypairCount *int `pulumi:"dhKeypairCount"`
	// Enable/disable Diffie-Hellman key pair cache CPU throttling. Valid values: `enable`, `disable`.
	DhKeypairThrottle *string `pulumi:"dhKeypairThrottle"`
	// Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations. Valid values: `software`, `hardware`.
	DhMode *string `pulumi:"dhMode"`
	// Enable/disable multiprocess Diffie-Hellman daemon for IKE. Valid values: `enable`, `disable`.
	DhMultiprocess *string `pulumi:"dhMultiprocess"`
	// Number of Diffie-Hellman workers to start.
	DhWorkerCount *int `pulumi:"dhWorkerCount"`
	// Maximum number of IPsec tunnels to negotiate simultaneously.
	EmbryonicLimit *int `pulumi:"embryonicLimit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	//
	// The `dhGroup1` block supports:
	Vdomparam *string `pulumi:"vdomparam"`
}

type IkeState struct {
	// Diffie-Hellman group 1 (MODP-768). The structure of `dhGroup1` block is documented below.
	DhGroup1 IkeDhGroup1PtrInput
	// Diffie-Hellman group 14 (MODP-2048). The structure of `dhGroup14` block is documented below.
	DhGroup14 IkeDhGroup14PtrInput
	// Diffie-Hellman group 15 (MODP-3072). The structure of `dhGroup15` block is documented below.
	DhGroup15 IkeDhGroup15PtrInput
	// Diffie-Hellman group 16 (MODP-4096). The structure of `dhGroup16` block is documented below.
	DhGroup16 IkeDhGroup16PtrInput
	// Diffie-Hellman group 17 (MODP-6144). The structure of `dhGroup17` block is documented below.
	DhGroup17 IkeDhGroup17PtrInput
	// Diffie-Hellman group 18 (MODP-8192). The structure of `dhGroup18` block is documented below.
	DhGroup18 IkeDhGroup18PtrInput
	// Diffie-Hellman group 19 (EC-P256). The structure of `dhGroup19` block is documented below.
	DhGroup19 IkeDhGroup19PtrInput
	// Diffie-Hellman group 2 (MODP-1024). The structure of `dhGroup2` block is documented below.
	DhGroup2 IkeDhGroup2PtrInput
	// Diffie-Hellman group 20 (EC-P384). The structure of `dhGroup20` block is documented below.
	DhGroup20 IkeDhGroup20PtrInput
	// Diffie-Hellman group 21 (EC-P521). The structure of `dhGroup21` block is documented below.
	DhGroup21 IkeDhGroup21PtrInput
	// Diffie-Hellman group 27 (EC-P224BP). The structure of `dhGroup27` block is documented below.
	DhGroup27 IkeDhGroup27PtrInput
	// Diffie-Hellman group 28 (EC-P256BP). The structure of `dhGroup28` block is documented below.
	DhGroup28 IkeDhGroup28PtrInput
	// Diffie-Hellman group 29 (EC-P384BP). The structure of `dhGroup29` block is documented below.
	DhGroup29 IkeDhGroup29PtrInput
	// Diffie-Hellman group 30 (EC-P512BP). The structure of `dhGroup30` block is documented below.
	DhGroup30 IkeDhGroup30PtrInput
	// Diffie-Hellman group 31 (EC-X25519). The structure of `dhGroup31` block is documented below.
	DhGroup31 IkeDhGroup31PtrInput
	// Diffie-Hellman group 32 (EC-X448). The structure of `dhGroup32` block is documented below.
	DhGroup32 IkeDhGroup32PtrInput
	// Diffie-Hellman group 5 (MODP-1536). The structure of `dhGroup5` block is documented below.
	DhGroup5 IkeDhGroup5PtrInput
	// Enable/disable Diffie-Hellman key pair cache. Valid values: `enable`, `disable`.
	DhKeypairCache pulumi.StringPtrInput
	// Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).
	DhKeypairCount pulumi.IntPtrInput
	// Enable/disable Diffie-Hellman key pair cache CPU throttling. Valid values: `enable`, `disable`.
	DhKeypairThrottle pulumi.StringPtrInput
	// Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations. Valid values: `software`, `hardware`.
	DhMode pulumi.StringPtrInput
	// Enable/disable multiprocess Diffie-Hellman daemon for IKE. Valid values: `enable`, `disable`.
	DhMultiprocess pulumi.StringPtrInput
	// Number of Diffie-Hellman workers to start.
	DhWorkerCount pulumi.IntPtrInput
	// Maximum number of IPsec tunnels to negotiate simultaneously.
	EmbryonicLimit pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	//
	// The `dhGroup1` block supports:
	Vdomparam pulumi.StringPtrInput
}

func (IkeState) ElementType() reflect.Type {
	return reflect.TypeOf((*ikeState)(nil)).Elem()
}

type ikeArgs struct {
	// Diffie-Hellman group 1 (MODP-768). The structure of `dhGroup1` block is documented below.
	DhGroup1 *IkeDhGroup1 `pulumi:"dhGroup1"`
	// Diffie-Hellman group 14 (MODP-2048). The structure of `dhGroup14` block is documented below.
	DhGroup14 *IkeDhGroup14 `pulumi:"dhGroup14"`
	// Diffie-Hellman group 15 (MODP-3072). The structure of `dhGroup15` block is documented below.
	DhGroup15 *IkeDhGroup15 `pulumi:"dhGroup15"`
	// Diffie-Hellman group 16 (MODP-4096). The structure of `dhGroup16` block is documented below.
	DhGroup16 *IkeDhGroup16 `pulumi:"dhGroup16"`
	// Diffie-Hellman group 17 (MODP-6144). The structure of `dhGroup17` block is documented below.
	DhGroup17 *IkeDhGroup17 `pulumi:"dhGroup17"`
	// Diffie-Hellman group 18 (MODP-8192). The structure of `dhGroup18` block is documented below.
	DhGroup18 *IkeDhGroup18 `pulumi:"dhGroup18"`
	// Diffie-Hellman group 19 (EC-P256). The structure of `dhGroup19` block is documented below.
	DhGroup19 *IkeDhGroup19 `pulumi:"dhGroup19"`
	// Diffie-Hellman group 2 (MODP-1024). The structure of `dhGroup2` block is documented below.
	DhGroup2 *IkeDhGroup2 `pulumi:"dhGroup2"`
	// Diffie-Hellman group 20 (EC-P384). The structure of `dhGroup20` block is documented below.
	DhGroup20 *IkeDhGroup20 `pulumi:"dhGroup20"`
	// Diffie-Hellman group 21 (EC-P521). The structure of `dhGroup21` block is documented below.
	DhGroup21 *IkeDhGroup21 `pulumi:"dhGroup21"`
	// Diffie-Hellman group 27 (EC-P224BP). The structure of `dhGroup27` block is documented below.
	DhGroup27 *IkeDhGroup27 `pulumi:"dhGroup27"`
	// Diffie-Hellman group 28 (EC-P256BP). The structure of `dhGroup28` block is documented below.
	DhGroup28 *IkeDhGroup28 `pulumi:"dhGroup28"`
	// Diffie-Hellman group 29 (EC-P384BP). The structure of `dhGroup29` block is documented below.
	DhGroup29 *IkeDhGroup29 `pulumi:"dhGroup29"`
	// Diffie-Hellman group 30 (EC-P512BP). The structure of `dhGroup30` block is documented below.
	DhGroup30 *IkeDhGroup30 `pulumi:"dhGroup30"`
	// Diffie-Hellman group 31 (EC-X25519). The structure of `dhGroup31` block is documented below.
	DhGroup31 *IkeDhGroup31 `pulumi:"dhGroup31"`
	// Diffie-Hellman group 32 (EC-X448). The structure of `dhGroup32` block is documented below.
	DhGroup32 *IkeDhGroup32 `pulumi:"dhGroup32"`
	// Diffie-Hellman group 5 (MODP-1536). The structure of `dhGroup5` block is documented below.
	DhGroup5 *IkeDhGroup5 `pulumi:"dhGroup5"`
	// Enable/disable Diffie-Hellman key pair cache. Valid values: `enable`, `disable`.
	DhKeypairCache *string `pulumi:"dhKeypairCache"`
	// Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).
	DhKeypairCount *int `pulumi:"dhKeypairCount"`
	// Enable/disable Diffie-Hellman key pair cache CPU throttling. Valid values: `enable`, `disable`.
	DhKeypairThrottle *string `pulumi:"dhKeypairThrottle"`
	// Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations. Valid values: `software`, `hardware`.
	DhMode *string `pulumi:"dhMode"`
	// Enable/disable multiprocess Diffie-Hellman daemon for IKE. Valid values: `enable`, `disable`.
	DhMultiprocess *string `pulumi:"dhMultiprocess"`
	// Number of Diffie-Hellman workers to start.
	DhWorkerCount *int `pulumi:"dhWorkerCount"`
	// Maximum number of IPsec tunnels to negotiate simultaneously.
	EmbryonicLimit *int `pulumi:"embryonicLimit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	//
	// The `dhGroup1` block supports:
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Ike resource.
type IkeArgs struct {
	// Diffie-Hellman group 1 (MODP-768). The structure of `dhGroup1` block is documented below.
	DhGroup1 IkeDhGroup1PtrInput
	// Diffie-Hellman group 14 (MODP-2048). The structure of `dhGroup14` block is documented below.
	DhGroup14 IkeDhGroup14PtrInput
	// Diffie-Hellman group 15 (MODP-3072). The structure of `dhGroup15` block is documented below.
	DhGroup15 IkeDhGroup15PtrInput
	// Diffie-Hellman group 16 (MODP-4096). The structure of `dhGroup16` block is documented below.
	DhGroup16 IkeDhGroup16PtrInput
	// Diffie-Hellman group 17 (MODP-6144). The structure of `dhGroup17` block is documented below.
	DhGroup17 IkeDhGroup17PtrInput
	// Diffie-Hellman group 18 (MODP-8192). The structure of `dhGroup18` block is documented below.
	DhGroup18 IkeDhGroup18PtrInput
	// Diffie-Hellman group 19 (EC-P256). The structure of `dhGroup19` block is documented below.
	DhGroup19 IkeDhGroup19PtrInput
	// Diffie-Hellman group 2 (MODP-1024). The structure of `dhGroup2` block is documented below.
	DhGroup2 IkeDhGroup2PtrInput
	// Diffie-Hellman group 20 (EC-P384). The structure of `dhGroup20` block is documented below.
	DhGroup20 IkeDhGroup20PtrInput
	// Diffie-Hellman group 21 (EC-P521). The structure of `dhGroup21` block is documented below.
	DhGroup21 IkeDhGroup21PtrInput
	// Diffie-Hellman group 27 (EC-P224BP). The structure of `dhGroup27` block is documented below.
	DhGroup27 IkeDhGroup27PtrInput
	// Diffie-Hellman group 28 (EC-P256BP). The structure of `dhGroup28` block is documented below.
	DhGroup28 IkeDhGroup28PtrInput
	// Diffie-Hellman group 29 (EC-P384BP). The structure of `dhGroup29` block is documented below.
	DhGroup29 IkeDhGroup29PtrInput
	// Diffie-Hellman group 30 (EC-P512BP). The structure of `dhGroup30` block is documented below.
	DhGroup30 IkeDhGroup30PtrInput
	// Diffie-Hellman group 31 (EC-X25519). The structure of `dhGroup31` block is documented below.
	DhGroup31 IkeDhGroup31PtrInput
	// Diffie-Hellman group 32 (EC-X448). The structure of `dhGroup32` block is documented below.
	DhGroup32 IkeDhGroup32PtrInput
	// Diffie-Hellman group 5 (MODP-1536). The structure of `dhGroup5` block is documented below.
	DhGroup5 IkeDhGroup5PtrInput
	// Enable/disable Diffie-Hellman key pair cache. Valid values: `enable`, `disable`.
	DhKeypairCache pulumi.StringPtrInput
	// Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).
	DhKeypairCount pulumi.IntPtrInput
	// Enable/disable Diffie-Hellman key pair cache CPU throttling. Valid values: `enable`, `disable`.
	DhKeypairThrottle pulumi.StringPtrInput
	// Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations. Valid values: `software`, `hardware`.
	DhMode pulumi.StringPtrInput
	// Enable/disable multiprocess Diffie-Hellman daemon for IKE. Valid values: `enable`, `disable`.
	DhMultiprocess pulumi.StringPtrInput
	// Number of Diffie-Hellman workers to start.
	DhWorkerCount pulumi.IntPtrInput
	// Maximum number of IPsec tunnels to negotiate simultaneously.
	EmbryonicLimit pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	//
	// The `dhGroup1` block supports:
	Vdomparam pulumi.StringPtrInput
}

func (IkeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ikeArgs)(nil)).Elem()
}

type IkeInput interface {
	pulumi.Input

	ToIkeOutput() IkeOutput
	ToIkeOutputWithContext(ctx context.Context) IkeOutput
}

func (*Ike) ElementType() reflect.Type {
	return reflect.TypeOf((**Ike)(nil)).Elem()
}

func (i *Ike) ToIkeOutput() IkeOutput {
	return i.ToIkeOutputWithContext(context.Background())
}

func (i *Ike) ToIkeOutputWithContext(ctx context.Context) IkeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IkeOutput)
}

// IkeArrayInput is an input type that accepts IkeArray and IkeArrayOutput values.
// You can construct a concrete instance of `IkeArrayInput` via:
//
//	IkeArray{ IkeArgs{...} }
type IkeArrayInput interface {
	pulumi.Input

	ToIkeArrayOutput() IkeArrayOutput
	ToIkeArrayOutputWithContext(context.Context) IkeArrayOutput
}

type IkeArray []IkeInput

func (IkeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ike)(nil)).Elem()
}

func (i IkeArray) ToIkeArrayOutput() IkeArrayOutput {
	return i.ToIkeArrayOutputWithContext(context.Background())
}

func (i IkeArray) ToIkeArrayOutputWithContext(ctx context.Context) IkeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IkeArrayOutput)
}

// IkeMapInput is an input type that accepts IkeMap and IkeMapOutput values.
// You can construct a concrete instance of `IkeMapInput` via:
//
//	IkeMap{ "key": IkeArgs{...} }
type IkeMapInput interface {
	pulumi.Input

	ToIkeMapOutput() IkeMapOutput
	ToIkeMapOutputWithContext(context.Context) IkeMapOutput
}

type IkeMap map[string]IkeInput

func (IkeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ike)(nil)).Elem()
}

func (i IkeMap) ToIkeMapOutput() IkeMapOutput {
	return i.ToIkeMapOutputWithContext(context.Background())
}

func (i IkeMap) ToIkeMapOutputWithContext(ctx context.Context) IkeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IkeMapOutput)
}

type IkeOutput struct{ *pulumi.OutputState }

func (IkeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ike)(nil)).Elem()
}

func (o IkeOutput) ToIkeOutput() IkeOutput {
	return o
}

func (o IkeOutput) ToIkeOutputWithContext(ctx context.Context) IkeOutput {
	return o
}

// Diffie-Hellman group 1 (MODP-768). The structure of `dhGroup1` block is documented below.
func (o IkeOutput) DhGroup1() IkeDhGroup1Output {
	return o.ApplyT(func(v *Ike) IkeDhGroup1Output { return v.DhGroup1 }).(IkeDhGroup1Output)
}

// Diffie-Hellman group 14 (MODP-2048). The structure of `dhGroup14` block is documented below.
func (o IkeOutput) DhGroup14() IkeDhGroup14Output {
	return o.ApplyT(func(v *Ike) IkeDhGroup14Output { return v.DhGroup14 }).(IkeDhGroup14Output)
}

// Diffie-Hellman group 15 (MODP-3072). The structure of `dhGroup15` block is documented below.
func (o IkeOutput) DhGroup15() IkeDhGroup15Output {
	return o.ApplyT(func(v *Ike) IkeDhGroup15Output { return v.DhGroup15 }).(IkeDhGroup15Output)
}

// Diffie-Hellman group 16 (MODP-4096). The structure of `dhGroup16` block is documented below.
func (o IkeOutput) DhGroup16() IkeDhGroup16Output {
	return o.ApplyT(func(v *Ike) IkeDhGroup16Output { return v.DhGroup16 }).(IkeDhGroup16Output)
}

// Diffie-Hellman group 17 (MODP-6144). The structure of `dhGroup17` block is documented below.
func (o IkeOutput) DhGroup17() IkeDhGroup17Output {
	return o.ApplyT(func(v *Ike) IkeDhGroup17Output { return v.DhGroup17 }).(IkeDhGroup17Output)
}

// Diffie-Hellman group 18 (MODP-8192). The structure of `dhGroup18` block is documented below.
func (o IkeOutput) DhGroup18() IkeDhGroup18Output {
	return o.ApplyT(func(v *Ike) IkeDhGroup18Output { return v.DhGroup18 }).(IkeDhGroup18Output)
}

// Diffie-Hellman group 19 (EC-P256). The structure of `dhGroup19` block is documented below.
func (o IkeOutput) DhGroup19() IkeDhGroup19Output {
	return o.ApplyT(func(v *Ike) IkeDhGroup19Output { return v.DhGroup19 }).(IkeDhGroup19Output)
}

// Diffie-Hellman group 2 (MODP-1024). The structure of `dhGroup2` block is documented below.
func (o IkeOutput) DhGroup2() IkeDhGroup2Output {
	return o.ApplyT(func(v *Ike) IkeDhGroup2Output { return v.DhGroup2 }).(IkeDhGroup2Output)
}

// Diffie-Hellman group 20 (EC-P384). The structure of `dhGroup20` block is documented below.
func (o IkeOutput) DhGroup20() IkeDhGroup20Output {
	return o.ApplyT(func(v *Ike) IkeDhGroup20Output { return v.DhGroup20 }).(IkeDhGroup20Output)
}

// Diffie-Hellman group 21 (EC-P521). The structure of `dhGroup21` block is documented below.
func (o IkeOutput) DhGroup21() IkeDhGroup21Output {
	return o.ApplyT(func(v *Ike) IkeDhGroup21Output { return v.DhGroup21 }).(IkeDhGroup21Output)
}

// Diffie-Hellman group 27 (EC-P224BP). The structure of `dhGroup27` block is documented below.
func (o IkeOutput) DhGroup27() IkeDhGroup27Output {
	return o.ApplyT(func(v *Ike) IkeDhGroup27Output { return v.DhGroup27 }).(IkeDhGroup27Output)
}

// Diffie-Hellman group 28 (EC-P256BP). The structure of `dhGroup28` block is documented below.
func (o IkeOutput) DhGroup28() IkeDhGroup28Output {
	return o.ApplyT(func(v *Ike) IkeDhGroup28Output { return v.DhGroup28 }).(IkeDhGroup28Output)
}

// Diffie-Hellman group 29 (EC-P384BP). The structure of `dhGroup29` block is documented below.
func (o IkeOutput) DhGroup29() IkeDhGroup29Output {
	return o.ApplyT(func(v *Ike) IkeDhGroup29Output { return v.DhGroup29 }).(IkeDhGroup29Output)
}

// Diffie-Hellman group 30 (EC-P512BP). The structure of `dhGroup30` block is documented below.
func (o IkeOutput) DhGroup30() IkeDhGroup30Output {
	return o.ApplyT(func(v *Ike) IkeDhGroup30Output { return v.DhGroup30 }).(IkeDhGroup30Output)
}

// Diffie-Hellman group 31 (EC-X25519). The structure of `dhGroup31` block is documented below.
func (o IkeOutput) DhGroup31() IkeDhGroup31Output {
	return o.ApplyT(func(v *Ike) IkeDhGroup31Output { return v.DhGroup31 }).(IkeDhGroup31Output)
}

// Diffie-Hellman group 32 (EC-X448). The structure of `dhGroup32` block is documented below.
func (o IkeOutput) DhGroup32() IkeDhGroup32Output {
	return o.ApplyT(func(v *Ike) IkeDhGroup32Output { return v.DhGroup32 }).(IkeDhGroup32Output)
}

// Diffie-Hellman group 5 (MODP-1536). The structure of `dhGroup5` block is documented below.
func (o IkeOutput) DhGroup5() IkeDhGroup5Output {
	return o.ApplyT(func(v *Ike) IkeDhGroup5Output { return v.DhGroup5 }).(IkeDhGroup5Output)
}

// Enable/disable Diffie-Hellman key pair cache. Valid values: `enable`, `disable`.
func (o IkeOutput) DhKeypairCache() pulumi.StringOutput {
	return o.ApplyT(func(v *Ike) pulumi.StringOutput { return v.DhKeypairCache }).(pulumi.StringOutput)
}

// Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).
func (o IkeOutput) DhKeypairCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Ike) pulumi.IntOutput { return v.DhKeypairCount }).(pulumi.IntOutput)
}

// Enable/disable Diffie-Hellman key pair cache CPU throttling. Valid values: `enable`, `disable`.
func (o IkeOutput) DhKeypairThrottle() pulumi.StringOutput {
	return o.ApplyT(func(v *Ike) pulumi.StringOutput { return v.DhKeypairThrottle }).(pulumi.StringOutput)
}

// Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations. Valid values: `software`, `hardware`.
func (o IkeOutput) DhMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Ike) pulumi.StringOutput { return v.DhMode }).(pulumi.StringOutput)
}

// Enable/disable multiprocess Diffie-Hellman daemon for IKE. Valid values: `enable`, `disable`.
func (o IkeOutput) DhMultiprocess() pulumi.StringOutput {
	return o.ApplyT(func(v *Ike) pulumi.StringOutput { return v.DhMultiprocess }).(pulumi.StringOutput)
}

// Number of Diffie-Hellman workers to start.
func (o IkeOutput) DhWorkerCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Ike) pulumi.IntOutput { return v.DhWorkerCount }).(pulumi.IntOutput)
}

// Maximum number of IPsec tunnels to negotiate simultaneously.
func (o IkeOutput) EmbryonicLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *Ike) pulumi.IntOutput { return v.EmbryonicLimit }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
//
// The `dhGroup1` block supports:
func (o IkeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ike) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type IkeArrayOutput struct{ *pulumi.OutputState }

func (IkeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ike)(nil)).Elem()
}

func (o IkeArrayOutput) ToIkeArrayOutput() IkeArrayOutput {
	return o
}

func (o IkeArrayOutput) ToIkeArrayOutputWithContext(ctx context.Context) IkeArrayOutput {
	return o
}

func (o IkeArrayOutput) Index(i pulumi.IntInput) IkeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ike {
		return vs[0].([]*Ike)[vs[1].(int)]
	}).(IkeOutput)
}

type IkeMapOutput struct{ *pulumi.OutputState }

func (IkeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ike)(nil)).Elem()
}

func (o IkeMapOutput) ToIkeMapOutput() IkeMapOutput {
	return o
}

func (o IkeMapOutput) ToIkeMapOutputWithContext(ctx context.Context) IkeMapOutput {
	return o
}

func (o IkeMapOutput) MapIndex(k pulumi.StringInput) IkeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ike {
		return vs[0].(map[string]*Ike)[vs[1].(string)]
	}).(IkeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IkeInput)(nil)).Elem(), &Ike{})
	pulumi.RegisterInputType(reflect.TypeOf((*IkeArrayInput)(nil)).Elem(), IkeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IkeMapInput)(nil)).Elem(), IkeMap{})
	pulumi.RegisterOutputType(IkeOutput{})
	pulumi.RegisterOutputType(IkeArrayOutput{})
	pulumi.RegisterOutputType(IkeMapOutput{})
}