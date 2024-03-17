// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure redundant internet connections using SD-WAN (formerly virtual WAN link). Applies to FortiOS Version `<= 6.4.0`.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/system"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := system.NewVirtualwanlink(ctx, "trname", &system.VirtualwanlinkArgs{
//				FailDetect:      pulumi.String("disable"),
//				LoadBalanceMode: pulumi.String("source-ip-based"),
//				Status:          pulumi.String("disable"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// System VirtualWanLink can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:system/virtualwanlink:Virtualwanlink labelname SystemVirtualWanLink
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:system/virtualwanlink:Virtualwanlink labelname SystemVirtualWanLink
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Virtualwanlink struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
	FailAlertInterfaces VirtualwanlinkFailAlertInterfaceArrayOutput `pulumi:"failAlertInterfaces"`
	// Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
	FailDetect pulumi.StringOutput `pulumi:"failDetect"`
	// SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `healthCheck` block is documented below.
	HealthChecks VirtualwanlinkHealthCheckArrayOutput `pulumi:"healthChecks"`
	// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
	LoadBalanceMode pulumi.StringOutput `pulumi:"loadBalanceMode"`
	// Physical FortiGate interfaces added to the virtual-wan-link. The structure of `members` block is documented below.
	Members VirtualwanlinkMemberArrayOutput `pulumi:"members"`
	// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
	NeighborHoldBootTime pulumi.IntOutput `pulumi:"neighborHoldBootTime"`
	// Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
	NeighborHoldDown pulumi.StringOutput `pulumi:"neighborHoldDown"`
	// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
	NeighborHoldDownTime pulumi.IntOutput `pulumi:"neighborHoldDownTime"`
	// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
	Neighbors VirtualwanlinkNeighborArrayOutput `pulumi:"neighbors"`
	// Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
	Services VirtualwanlinkServiceArrayOutput `pulumi:"services"`
	// Enable/disable SD-WAN. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Configure SD-WAN zones. The structure of `zone` block is documented below.
	Zones VirtualwanlinkZoneArrayOutput `pulumi:"zones"`
}

// NewVirtualwanlink registers a new resource with the given unique name, arguments, and options.
func NewVirtualwanlink(ctx *pulumi.Context,
	name string, args *VirtualwanlinkArgs, opts ...pulumi.ResourceOption) (*Virtualwanlink, error) {
	if args == nil {
		args = &VirtualwanlinkArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Virtualwanlink
	err := ctx.RegisterResource("fortios:system/virtualwanlink:Virtualwanlink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualwanlink gets an existing Virtualwanlink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualwanlink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualwanlinkState, opts ...pulumi.ResourceOption) (*Virtualwanlink, error) {
	var resource Virtualwanlink
	err := ctx.ReadResource("fortios:system/virtualwanlink:Virtualwanlink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Virtualwanlink resources.
type virtualwanlinkState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
	FailAlertInterfaces []VirtualwanlinkFailAlertInterface `pulumi:"failAlertInterfaces"`
	// Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
	FailDetect *string `pulumi:"failDetect"`
	// SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `healthCheck` block is documented below.
	HealthChecks []VirtualwanlinkHealthCheck `pulumi:"healthChecks"`
	// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
	LoadBalanceMode *string `pulumi:"loadBalanceMode"`
	// Physical FortiGate interfaces added to the virtual-wan-link. The structure of `members` block is documented below.
	Members []VirtualwanlinkMember `pulumi:"members"`
	// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
	NeighborHoldBootTime *int `pulumi:"neighborHoldBootTime"`
	// Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
	NeighborHoldDown *string `pulumi:"neighborHoldDown"`
	// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
	NeighborHoldDownTime *int `pulumi:"neighborHoldDownTime"`
	// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
	Neighbors []VirtualwanlinkNeighbor `pulumi:"neighbors"`
	// Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
	Services []VirtualwanlinkService `pulumi:"services"`
	// Enable/disable SD-WAN. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Configure SD-WAN zones. The structure of `zone` block is documented below.
	Zones []VirtualwanlinkZone `pulumi:"zones"`
}

type VirtualwanlinkState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
	FailAlertInterfaces VirtualwanlinkFailAlertInterfaceArrayInput
	// Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
	FailDetect pulumi.StringPtrInput
	// SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `healthCheck` block is documented below.
	HealthChecks VirtualwanlinkHealthCheckArrayInput
	// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
	LoadBalanceMode pulumi.StringPtrInput
	// Physical FortiGate interfaces added to the virtual-wan-link. The structure of `members` block is documented below.
	Members VirtualwanlinkMemberArrayInput
	// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
	NeighborHoldBootTime pulumi.IntPtrInput
	// Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
	NeighborHoldDown pulumi.StringPtrInput
	// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
	NeighborHoldDownTime pulumi.IntPtrInput
	// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
	Neighbors VirtualwanlinkNeighborArrayInput
	// Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
	Services VirtualwanlinkServiceArrayInput
	// Enable/disable SD-WAN. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Configure SD-WAN zones. The structure of `zone` block is documented below.
	Zones VirtualwanlinkZoneArrayInput
}

func (VirtualwanlinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualwanlinkState)(nil)).Elem()
}

type virtualwanlinkArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
	FailAlertInterfaces []VirtualwanlinkFailAlertInterface `pulumi:"failAlertInterfaces"`
	// Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
	FailDetect *string `pulumi:"failDetect"`
	// SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `healthCheck` block is documented below.
	HealthChecks []VirtualwanlinkHealthCheck `pulumi:"healthChecks"`
	// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
	LoadBalanceMode *string `pulumi:"loadBalanceMode"`
	// Physical FortiGate interfaces added to the virtual-wan-link. The structure of `members` block is documented below.
	Members []VirtualwanlinkMember `pulumi:"members"`
	// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
	NeighborHoldBootTime *int `pulumi:"neighborHoldBootTime"`
	// Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
	NeighborHoldDown *string `pulumi:"neighborHoldDown"`
	// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
	NeighborHoldDownTime *int `pulumi:"neighborHoldDownTime"`
	// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
	Neighbors []VirtualwanlinkNeighbor `pulumi:"neighbors"`
	// Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
	Services []VirtualwanlinkService `pulumi:"services"`
	// Enable/disable SD-WAN. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Configure SD-WAN zones. The structure of `zone` block is documented below.
	Zones []VirtualwanlinkZone `pulumi:"zones"`
}

// The set of arguments for constructing a Virtualwanlink resource.
type VirtualwanlinkArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
	FailAlertInterfaces VirtualwanlinkFailAlertInterfaceArrayInput
	// Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
	FailDetect pulumi.StringPtrInput
	// SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `healthCheck` block is documented below.
	HealthChecks VirtualwanlinkHealthCheckArrayInput
	// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
	LoadBalanceMode pulumi.StringPtrInput
	// Physical FortiGate interfaces added to the virtual-wan-link. The structure of `members` block is documented below.
	Members VirtualwanlinkMemberArrayInput
	// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
	NeighborHoldBootTime pulumi.IntPtrInput
	// Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
	NeighborHoldDown pulumi.StringPtrInput
	// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
	NeighborHoldDownTime pulumi.IntPtrInput
	// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
	Neighbors VirtualwanlinkNeighborArrayInput
	// Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
	Services VirtualwanlinkServiceArrayInput
	// Enable/disable SD-WAN. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Configure SD-WAN zones. The structure of `zone` block is documented below.
	Zones VirtualwanlinkZoneArrayInput
}

func (VirtualwanlinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualwanlinkArgs)(nil)).Elem()
}

type VirtualwanlinkInput interface {
	pulumi.Input

	ToVirtualwanlinkOutput() VirtualwanlinkOutput
	ToVirtualwanlinkOutputWithContext(ctx context.Context) VirtualwanlinkOutput
}

func (*Virtualwanlink) ElementType() reflect.Type {
	return reflect.TypeOf((**Virtualwanlink)(nil)).Elem()
}

func (i *Virtualwanlink) ToVirtualwanlinkOutput() VirtualwanlinkOutput {
	return i.ToVirtualwanlinkOutputWithContext(context.Background())
}

func (i *Virtualwanlink) ToVirtualwanlinkOutputWithContext(ctx context.Context) VirtualwanlinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualwanlinkOutput)
}

// VirtualwanlinkArrayInput is an input type that accepts VirtualwanlinkArray and VirtualwanlinkArrayOutput values.
// You can construct a concrete instance of `VirtualwanlinkArrayInput` via:
//
//	VirtualwanlinkArray{ VirtualwanlinkArgs{...} }
type VirtualwanlinkArrayInput interface {
	pulumi.Input

	ToVirtualwanlinkArrayOutput() VirtualwanlinkArrayOutput
	ToVirtualwanlinkArrayOutputWithContext(context.Context) VirtualwanlinkArrayOutput
}

type VirtualwanlinkArray []VirtualwanlinkInput

func (VirtualwanlinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Virtualwanlink)(nil)).Elem()
}

func (i VirtualwanlinkArray) ToVirtualwanlinkArrayOutput() VirtualwanlinkArrayOutput {
	return i.ToVirtualwanlinkArrayOutputWithContext(context.Background())
}

func (i VirtualwanlinkArray) ToVirtualwanlinkArrayOutputWithContext(ctx context.Context) VirtualwanlinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualwanlinkArrayOutput)
}

// VirtualwanlinkMapInput is an input type that accepts VirtualwanlinkMap and VirtualwanlinkMapOutput values.
// You can construct a concrete instance of `VirtualwanlinkMapInput` via:
//
//	VirtualwanlinkMap{ "key": VirtualwanlinkArgs{...} }
type VirtualwanlinkMapInput interface {
	pulumi.Input

	ToVirtualwanlinkMapOutput() VirtualwanlinkMapOutput
	ToVirtualwanlinkMapOutputWithContext(context.Context) VirtualwanlinkMapOutput
}

type VirtualwanlinkMap map[string]VirtualwanlinkInput

func (VirtualwanlinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Virtualwanlink)(nil)).Elem()
}

func (i VirtualwanlinkMap) ToVirtualwanlinkMapOutput() VirtualwanlinkMapOutput {
	return i.ToVirtualwanlinkMapOutputWithContext(context.Background())
}

func (i VirtualwanlinkMap) ToVirtualwanlinkMapOutputWithContext(ctx context.Context) VirtualwanlinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualwanlinkMapOutput)
}

type VirtualwanlinkOutput struct{ *pulumi.OutputState }

func (VirtualwanlinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Virtualwanlink)(nil)).Elem()
}

func (o VirtualwanlinkOutput) ToVirtualwanlinkOutput() VirtualwanlinkOutput {
	return o
}

func (o VirtualwanlinkOutput) ToVirtualwanlinkOutputWithContext(ctx context.Context) VirtualwanlinkOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o VirtualwanlinkOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Virtualwanlink) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
func (o VirtualwanlinkOutput) FailAlertInterfaces() VirtualwanlinkFailAlertInterfaceArrayOutput {
	return o.ApplyT(func(v *Virtualwanlink) VirtualwanlinkFailAlertInterfaceArrayOutput { return v.FailAlertInterfaces }).(VirtualwanlinkFailAlertInterfaceArrayOutput)
}

// Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
func (o VirtualwanlinkOutput) FailDetect() pulumi.StringOutput {
	return o.ApplyT(func(v *Virtualwanlink) pulumi.StringOutput { return v.FailDetect }).(pulumi.StringOutput)
}

// SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `healthCheck` block is documented below.
func (o VirtualwanlinkOutput) HealthChecks() VirtualwanlinkHealthCheckArrayOutput {
	return o.ApplyT(func(v *Virtualwanlink) VirtualwanlinkHealthCheckArrayOutput { return v.HealthChecks }).(VirtualwanlinkHealthCheckArrayOutput)
}

// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
func (o VirtualwanlinkOutput) LoadBalanceMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Virtualwanlink) pulumi.StringOutput { return v.LoadBalanceMode }).(pulumi.StringOutput)
}

// Physical FortiGate interfaces added to the virtual-wan-link. The structure of `members` block is documented below.
func (o VirtualwanlinkOutput) Members() VirtualwanlinkMemberArrayOutput {
	return o.ApplyT(func(v *Virtualwanlink) VirtualwanlinkMemberArrayOutput { return v.Members }).(VirtualwanlinkMemberArrayOutput)
}

// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
func (o VirtualwanlinkOutput) NeighborHoldBootTime() pulumi.IntOutput {
	return o.ApplyT(func(v *Virtualwanlink) pulumi.IntOutput { return v.NeighborHoldBootTime }).(pulumi.IntOutput)
}

// Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
func (o VirtualwanlinkOutput) NeighborHoldDown() pulumi.StringOutput {
	return o.ApplyT(func(v *Virtualwanlink) pulumi.StringOutput { return v.NeighborHoldDown }).(pulumi.StringOutput)
}

// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
func (o VirtualwanlinkOutput) NeighborHoldDownTime() pulumi.IntOutput {
	return o.ApplyT(func(v *Virtualwanlink) pulumi.IntOutput { return v.NeighborHoldDownTime }).(pulumi.IntOutput)
}

// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
func (o VirtualwanlinkOutput) Neighbors() VirtualwanlinkNeighborArrayOutput {
	return o.ApplyT(func(v *Virtualwanlink) VirtualwanlinkNeighborArrayOutput { return v.Neighbors }).(VirtualwanlinkNeighborArrayOutput)
}

// Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
func (o VirtualwanlinkOutput) Services() VirtualwanlinkServiceArrayOutput {
	return o.ApplyT(func(v *Virtualwanlink) VirtualwanlinkServiceArrayOutput { return v.Services }).(VirtualwanlinkServiceArrayOutput)
}

// Enable/disable SD-WAN. Valid values: `disable`, `enable`.
func (o VirtualwanlinkOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Virtualwanlink) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o VirtualwanlinkOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Virtualwanlink) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Configure SD-WAN zones. The structure of `zone` block is documented below.
func (o VirtualwanlinkOutput) Zones() VirtualwanlinkZoneArrayOutput {
	return o.ApplyT(func(v *Virtualwanlink) VirtualwanlinkZoneArrayOutput { return v.Zones }).(VirtualwanlinkZoneArrayOutput)
}

type VirtualwanlinkArrayOutput struct{ *pulumi.OutputState }

func (VirtualwanlinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Virtualwanlink)(nil)).Elem()
}

func (o VirtualwanlinkArrayOutput) ToVirtualwanlinkArrayOutput() VirtualwanlinkArrayOutput {
	return o
}

func (o VirtualwanlinkArrayOutput) ToVirtualwanlinkArrayOutputWithContext(ctx context.Context) VirtualwanlinkArrayOutput {
	return o
}

func (o VirtualwanlinkArrayOutput) Index(i pulumi.IntInput) VirtualwanlinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Virtualwanlink {
		return vs[0].([]*Virtualwanlink)[vs[1].(int)]
	}).(VirtualwanlinkOutput)
}

type VirtualwanlinkMapOutput struct{ *pulumi.OutputState }

func (VirtualwanlinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Virtualwanlink)(nil)).Elem()
}

func (o VirtualwanlinkMapOutput) ToVirtualwanlinkMapOutput() VirtualwanlinkMapOutput {
	return o
}

func (o VirtualwanlinkMapOutput) ToVirtualwanlinkMapOutputWithContext(ctx context.Context) VirtualwanlinkMapOutput {
	return o
}

func (o VirtualwanlinkMapOutput) MapIndex(k pulumi.StringInput) VirtualwanlinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Virtualwanlink {
		return vs[0].(map[string]*Virtualwanlink)[vs[1].(string)]
	}).(VirtualwanlinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualwanlinkInput)(nil)).Elem(), &Virtualwanlink{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualwanlinkArrayInput)(nil)).Elem(), VirtualwanlinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualwanlinkMapInput)(nil)).Elem(), VirtualwanlinkMap{})
	pulumi.RegisterOutputType(VirtualwanlinkOutput{})
	pulumi.RegisterOutputType(VirtualwanlinkArrayOutput{})
	pulumi.RegisterOutputType(VirtualwanlinkMapOutput{})
}