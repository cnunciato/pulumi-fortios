// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure authentication based routing.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fortios.NewRouterAuthpath(ctx, "trname", &fortios.RouterAuthpathArgs{
//				Device:  pulumi.String("port3"),
//				Gateway: pulumi.String("1.1.1.1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # Router AuthPath can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/routerAuthpath:RouterAuthpath labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/routerAuthpath:RouterAuthpath labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type RouterAuthpath struct {
	pulumi.CustomResourceState

	// Outgoing interface.
	Device pulumi.StringOutput `pulumi:"device"`
	// Gateway IP address.
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// Name of the entry.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterAuthpath registers a new resource with the given unique name, arguments, and options.
func NewRouterAuthpath(ctx *pulumi.Context,
	name string, args *RouterAuthpathArgs, opts ...pulumi.ResourceOption) (*RouterAuthpath, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Device == nil {
		return nil, errors.New("invalid value for required argument 'Device'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RouterAuthpath
	err := ctx.RegisterResource("fortios:index/routerAuthpath:RouterAuthpath", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterAuthpath gets an existing RouterAuthpath resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterAuthpath(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterAuthpathState, opts ...pulumi.ResourceOption) (*RouterAuthpath, error) {
	var resource RouterAuthpath
	err := ctx.ReadResource("fortios:index/routerAuthpath:RouterAuthpath", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterAuthpath resources.
type routerAuthpathState struct {
	// Outgoing interface.
	Device *string `pulumi:"device"`
	// Gateway IP address.
	Gateway *string `pulumi:"gateway"`
	// Name of the entry.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RouterAuthpathState struct {
	// Outgoing interface.
	Device pulumi.StringPtrInput
	// Gateway IP address.
	Gateway pulumi.StringPtrInput
	// Name of the entry.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterAuthpathState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerAuthpathState)(nil)).Elem()
}

type routerAuthpathArgs struct {
	// Outgoing interface.
	Device string `pulumi:"device"`
	// Gateway IP address.
	Gateway *string `pulumi:"gateway"`
	// Name of the entry.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterAuthpath resource.
type RouterAuthpathArgs struct {
	// Outgoing interface.
	Device pulumi.StringInput
	// Gateway IP address.
	Gateway pulumi.StringPtrInput
	// Name of the entry.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterAuthpathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerAuthpathArgs)(nil)).Elem()
}

type RouterAuthpathInput interface {
	pulumi.Input

	ToRouterAuthpathOutput() RouterAuthpathOutput
	ToRouterAuthpathOutputWithContext(ctx context.Context) RouterAuthpathOutput
}

func (*RouterAuthpath) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterAuthpath)(nil)).Elem()
}

func (i *RouterAuthpath) ToRouterAuthpathOutput() RouterAuthpathOutput {
	return i.ToRouterAuthpathOutputWithContext(context.Background())
}

func (i *RouterAuthpath) ToRouterAuthpathOutputWithContext(ctx context.Context) RouterAuthpathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAuthpathOutput)
}

// RouterAuthpathArrayInput is an input type that accepts RouterAuthpathArray and RouterAuthpathArrayOutput values.
// You can construct a concrete instance of `RouterAuthpathArrayInput` via:
//
//	RouterAuthpathArray{ RouterAuthpathArgs{...} }
type RouterAuthpathArrayInput interface {
	pulumi.Input

	ToRouterAuthpathArrayOutput() RouterAuthpathArrayOutput
	ToRouterAuthpathArrayOutputWithContext(context.Context) RouterAuthpathArrayOutput
}

type RouterAuthpathArray []RouterAuthpathInput

func (RouterAuthpathArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterAuthpath)(nil)).Elem()
}

func (i RouterAuthpathArray) ToRouterAuthpathArrayOutput() RouterAuthpathArrayOutput {
	return i.ToRouterAuthpathArrayOutputWithContext(context.Background())
}

func (i RouterAuthpathArray) ToRouterAuthpathArrayOutputWithContext(ctx context.Context) RouterAuthpathArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAuthpathArrayOutput)
}

// RouterAuthpathMapInput is an input type that accepts RouterAuthpathMap and RouterAuthpathMapOutput values.
// You can construct a concrete instance of `RouterAuthpathMapInput` via:
//
//	RouterAuthpathMap{ "key": RouterAuthpathArgs{...} }
type RouterAuthpathMapInput interface {
	pulumi.Input

	ToRouterAuthpathMapOutput() RouterAuthpathMapOutput
	ToRouterAuthpathMapOutputWithContext(context.Context) RouterAuthpathMapOutput
}

type RouterAuthpathMap map[string]RouterAuthpathInput

func (RouterAuthpathMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterAuthpath)(nil)).Elem()
}

func (i RouterAuthpathMap) ToRouterAuthpathMapOutput() RouterAuthpathMapOutput {
	return i.ToRouterAuthpathMapOutputWithContext(context.Background())
}

func (i RouterAuthpathMap) ToRouterAuthpathMapOutputWithContext(ctx context.Context) RouterAuthpathMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAuthpathMapOutput)
}

type RouterAuthpathOutput struct{ *pulumi.OutputState }

func (RouterAuthpathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterAuthpath)(nil)).Elem()
}

func (o RouterAuthpathOutput) ToRouterAuthpathOutput() RouterAuthpathOutput {
	return o
}

func (o RouterAuthpathOutput) ToRouterAuthpathOutputWithContext(ctx context.Context) RouterAuthpathOutput {
	return o
}

// Outgoing interface.
func (o RouterAuthpathOutput) Device() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterAuthpath) pulumi.StringOutput { return v.Device }).(pulumi.StringOutput)
}

// Gateway IP address.
func (o RouterAuthpathOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterAuthpath) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// Name of the entry.
func (o RouterAuthpathOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterAuthpath) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o RouterAuthpathOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterAuthpath) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RouterAuthpathArrayOutput struct{ *pulumi.OutputState }

func (RouterAuthpathArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterAuthpath)(nil)).Elem()
}

func (o RouterAuthpathArrayOutput) ToRouterAuthpathArrayOutput() RouterAuthpathArrayOutput {
	return o
}

func (o RouterAuthpathArrayOutput) ToRouterAuthpathArrayOutputWithContext(ctx context.Context) RouterAuthpathArrayOutput {
	return o
}

func (o RouterAuthpathArrayOutput) Index(i pulumi.IntInput) RouterAuthpathOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterAuthpath {
		return vs[0].([]*RouterAuthpath)[vs[1].(int)]
	}).(RouterAuthpathOutput)
}

type RouterAuthpathMapOutput struct{ *pulumi.OutputState }

func (RouterAuthpathMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterAuthpath)(nil)).Elem()
}

func (o RouterAuthpathMapOutput) ToRouterAuthpathMapOutput() RouterAuthpathMapOutput {
	return o
}

func (o RouterAuthpathMapOutput) ToRouterAuthpathMapOutputWithContext(ctx context.Context) RouterAuthpathMapOutput {
	return o
}

func (o RouterAuthpathMapOutput) MapIndex(k pulumi.StringInput) RouterAuthpathOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterAuthpath {
		return vs[0].(map[string]*RouterAuthpath)[vs[1].(string)]
	}).(RouterAuthpathOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterAuthpathInput)(nil)).Elem(), &RouterAuthpath{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterAuthpathArrayInput)(nil)).Elem(), RouterAuthpathArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterAuthpathMapInput)(nil)).Elem(), RouterAuthpathMap{})
	pulumi.RegisterOutputType(RouterAuthpathOutput{})
	pulumi.RegisterOutputType(RouterAuthpathArrayOutput{})
	pulumi.RegisterOutputType(RouterAuthpathMapOutput{})
}
