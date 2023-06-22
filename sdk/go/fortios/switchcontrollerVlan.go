// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure VLANs for switch controller. Applies to FortiOS Version `<= 6.2.0`.
//
// ## Import
//
// # SwitchController Vlan can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerVlan:SwitchcontrollerVlan labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerVlan:SwitchcontrollerVlan labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SwitchcontrollerVlan struct {
	pulumi.CustomResourceState

	// Authentication. Valid values: `radius`, `usergroup`.
	Auth pulumi.StringOutput `pulumi:"auth"`
	// Color of icon on the GUI.
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comments pulumi.StringOutput `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Switch VLAN name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specify captive portal replacement message override group.
	PortalMessageOverrideGroup pulumi.StringOutput `pulumi:"portalMessageOverrideGroup"`
	// Individual message overrides. The structure of `portalMessageOverrides` block is documented below.
	PortalMessageOverrides SwitchcontrollerVlanPortalMessageOverridesOutput `pulumi:"portalMessageOverrides"`
	// Authentication radius server.
	RadiusServer pulumi.StringOutput `pulumi:"radiusServer"`
	// Security. Valid values: `open`, `captive-portal`, `8021x`.
	Security pulumi.StringOutput `pulumi:"security"`
	// Selected user group. The structure of `selectedUsergroups` block is documented below.
	SelectedUsergroups SwitchcontrollerVlanSelectedUsergroupArrayOutput `pulumi:"selectedUsergroups"`
	// Authentication usergroup.
	Usergroup pulumi.StringOutput `pulumi:"usergroup"`
	// Virtual domain,
	Vdom pulumi.StringOutput `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// VLAN ID.
	Vlanid pulumi.IntOutput `pulumi:"vlanid"`
}

// NewSwitchcontrollerVlan registers a new resource with the given unique name, arguments, and options.
func NewSwitchcontrollerVlan(ctx *pulumi.Context,
	name string, args *SwitchcontrollerVlanArgs, opts ...pulumi.ResourceOption) (*SwitchcontrollerVlan, error) {
	if args == nil {
		args = &SwitchcontrollerVlanArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchcontrollerVlan
	err := ctx.RegisterResource("fortios:index/switchcontrollerVlan:SwitchcontrollerVlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchcontrollerVlan gets an existing SwitchcontrollerVlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchcontrollerVlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchcontrollerVlanState, opts ...pulumi.ResourceOption) (*SwitchcontrollerVlan, error) {
	var resource SwitchcontrollerVlan
	err := ctx.ReadResource("fortios:index/switchcontrollerVlan:SwitchcontrollerVlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchcontrollerVlan resources.
type switchcontrollerVlanState struct {
	// Authentication. Valid values: `radius`, `usergroup`.
	Auth *string `pulumi:"auth"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Switch VLAN name.
	Name *string `pulumi:"name"`
	// Specify captive portal replacement message override group.
	PortalMessageOverrideGroup *string `pulumi:"portalMessageOverrideGroup"`
	// Individual message overrides. The structure of `portalMessageOverrides` block is documented below.
	PortalMessageOverrides *SwitchcontrollerVlanPortalMessageOverrides `pulumi:"portalMessageOverrides"`
	// Authentication radius server.
	RadiusServer *string `pulumi:"radiusServer"`
	// Security. Valid values: `open`, `captive-portal`, `8021x`.
	Security *string `pulumi:"security"`
	// Selected user group. The structure of `selectedUsergroups` block is documented below.
	SelectedUsergroups []SwitchcontrollerVlanSelectedUsergroup `pulumi:"selectedUsergroups"`
	// Authentication usergroup.
	Usergroup *string `pulumi:"usergroup"`
	// Virtual domain,
	Vdom *string `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// VLAN ID.
	Vlanid *int `pulumi:"vlanid"`
}

type SwitchcontrollerVlanState struct {
	// Authentication. Valid values: `radius`, `usergroup`.
	Auth pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Switch VLAN name.
	Name pulumi.StringPtrInput
	// Specify captive portal replacement message override group.
	PortalMessageOverrideGroup pulumi.StringPtrInput
	// Individual message overrides. The structure of `portalMessageOverrides` block is documented below.
	PortalMessageOverrides SwitchcontrollerVlanPortalMessageOverridesPtrInput
	// Authentication radius server.
	RadiusServer pulumi.StringPtrInput
	// Security. Valid values: `open`, `captive-portal`, `8021x`.
	Security pulumi.StringPtrInput
	// Selected user group. The structure of `selectedUsergroups` block is documented below.
	SelectedUsergroups SwitchcontrollerVlanSelectedUsergroupArrayInput
	// Authentication usergroup.
	Usergroup pulumi.StringPtrInput
	// Virtual domain,
	Vdom pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// VLAN ID.
	Vlanid pulumi.IntPtrInput
}

func (SwitchcontrollerVlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerVlanState)(nil)).Elem()
}

type switchcontrollerVlanArgs struct {
	// Authentication. Valid values: `radius`, `usergroup`.
	Auth *string `pulumi:"auth"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Switch VLAN name.
	Name *string `pulumi:"name"`
	// Specify captive portal replacement message override group.
	PortalMessageOverrideGroup *string `pulumi:"portalMessageOverrideGroup"`
	// Individual message overrides. The structure of `portalMessageOverrides` block is documented below.
	PortalMessageOverrides *SwitchcontrollerVlanPortalMessageOverrides `pulumi:"portalMessageOverrides"`
	// Authentication radius server.
	RadiusServer *string `pulumi:"radiusServer"`
	// Security. Valid values: `open`, `captive-portal`, `8021x`.
	Security *string `pulumi:"security"`
	// Selected user group. The structure of `selectedUsergroups` block is documented below.
	SelectedUsergroups []SwitchcontrollerVlanSelectedUsergroup `pulumi:"selectedUsergroups"`
	// Authentication usergroup.
	Usergroup *string `pulumi:"usergroup"`
	// Virtual domain,
	Vdom *string `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// VLAN ID.
	Vlanid *int `pulumi:"vlanid"`
}

// The set of arguments for constructing a SwitchcontrollerVlan resource.
type SwitchcontrollerVlanArgs struct {
	// Authentication. Valid values: `radius`, `usergroup`.
	Auth pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Switch VLAN name.
	Name pulumi.StringPtrInput
	// Specify captive portal replacement message override group.
	PortalMessageOverrideGroup pulumi.StringPtrInput
	// Individual message overrides. The structure of `portalMessageOverrides` block is documented below.
	PortalMessageOverrides SwitchcontrollerVlanPortalMessageOverridesPtrInput
	// Authentication radius server.
	RadiusServer pulumi.StringPtrInput
	// Security. Valid values: `open`, `captive-portal`, `8021x`.
	Security pulumi.StringPtrInput
	// Selected user group. The structure of `selectedUsergroups` block is documented below.
	SelectedUsergroups SwitchcontrollerVlanSelectedUsergroupArrayInput
	// Authentication usergroup.
	Usergroup pulumi.StringPtrInput
	// Virtual domain,
	Vdom pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// VLAN ID.
	Vlanid pulumi.IntPtrInput
}

func (SwitchcontrollerVlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerVlanArgs)(nil)).Elem()
}

type SwitchcontrollerVlanInput interface {
	pulumi.Input

	ToSwitchcontrollerVlanOutput() SwitchcontrollerVlanOutput
	ToSwitchcontrollerVlanOutputWithContext(ctx context.Context) SwitchcontrollerVlanOutput
}

func (*SwitchcontrollerVlan) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerVlan)(nil)).Elem()
}

func (i *SwitchcontrollerVlan) ToSwitchcontrollerVlanOutput() SwitchcontrollerVlanOutput {
	return i.ToSwitchcontrollerVlanOutputWithContext(context.Background())
}

func (i *SwitchcontrollerVlan) ToSwitchcontrollerVlanOutputWithContext(ctx context.Context) SwitchcontrollerVlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerVlanOutput)
}

// SwitchcontrollerVlanArrayInput is an input type that accepts SwitchcontrollerVlanArray and SwitchcontrollerVlanArrayOutput values.
// You can construct a concrete instance of `SwitchcontrollerVlanArrayInput` via:
//
//	SwitchcontrollerVlanArray{ SwitchcontrollerVlanArgs{...} }
type SwitchcontrollerVlanArrayInput interface {
	pulumi.Input

	ToSwitchcontrollerVlanArrayOutput() SwitchcontrollerVlanArrayOutput
	ToSwitchcontrollerVlanArrayOutputWithContext(context.Context) SwitchcontrollerVlanArrayOutput
}

type SwitchcontrollerVlanArray []SwitchcontrollerVlanInput

func (SwitchcontrollerVlanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerVlan)(nil)).Elem()
}

func (i SwitchcontrollerVlanArray) ToSwitchcontrollerVlanArrayOutput() SwitchcontrollerVlanArrayOutput {
	return i.ToSwitchcontrollerVlanArrayOutputWithContext(context.Background())
}

func (i SwitchcontrollerVlanArray) ToSwitchcontrollerVlanArrayOutputWithContext(ctx context.Context) SwitchcontrollerVlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerVlanArrayOutput)
}

// SwitchcontrollerVlanMapInput is an input type that accepts SwitchcontrollerVlanMap and SwitchcontrollerVlanMapOutput values.
// You can construct a concrete instance of `SwitchcontrollerVlanMapInput` via:
//
//	SwitchcontrollerVlanMap{ "key": SwitchcontrollerVlanArgs{...} }
type SwitchcontrollerVlanMapInput interface {
	pulumi.Input

	ToSwitchcontrollerVlanMapOutput() SwitchcontrollerVlanMapOutput
	ToSwitchcontrollerVlanMapOutputWithContext(context.Context) SwitchcontrollerVlanMapOutput
}

type SwitchcontrollerVlanMap map[string]SwitchcontrollerVlanInput

func (SwitchcontrollerVlanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerVlan)(nil)).Elem()
}

func (i SwitchcontrollerVlanMap) ToSwitchcontrollerVlanMapOutput() SwitchcontrollerVlanMapOutput {
	return i.ToSwitchcontrollerVlanMapOutputWithContext(context.Background())
}

func (i SwitchcontrollerVlanMap) ToSwitchcontrollerVlanMapOutputWithContext(ctx context.Context) SwitchcontrollerVlanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerVlanMapOutput)
}

type SwitchcontrollerVlanOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerVlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerVlan)(nil)).Elem()
}

func (o SwitchcontrollerVlanOutput) ToSwitchcontrollerVlanOutput() SwitchcontrollerVlanOutput {
	return o
}

func (o SwitchcontrollerVlanOutput) ToSwitchcontrollerVlanOutputWithContext(ctx context.Context) SwitchcontrollerVlanOutput {
	return o
}

// Authentication. Valid values: `radius`, `usergroup`.
func (o SwitchcontrollerVlanOutput) Auth() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerVlan) pulumi.StringOutput { return v.Auth }).(pulumi.StringOutput)
}

// Color of icon on the GUI.
func (o SwitchcontrollerVlanOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerVlan) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o SwitchcontrollerVlanOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerVlan) pulumi.StringOutput { return v.Comments }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SwitchcontrollerVlanOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerVlan) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Switch VLAN name.
func (o SwitchcontrollerVlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerVlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specify captive portal replacement message override group.
func (o SwitchcontrollerVlanOutput) PortalMessageOverrideGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerVlan) pulumi.StringOutput { return v.PortalMessageOverrideGroup }).(pulumi.StringOutput)
}

// Individual message overrides. The structure of `portalMessageOverrides` block is documented below.
func (o SwitchcontrollerVlanOutput) PortalMessageOverrides() SwitchcontrollerVlanPortalMessageOverridesOutput {
	return o.ApplyT(func(v *SwitchcontrollerVlan) SwitchcontrollerVlanPortalMessageOverridesOutput {
		return v.PortalMessageOverrides
	}).(SwitchcontrollerVlanPortalMessageOverridesOutput)
}

// Authentication radius server.
func (o SwitchcontrollerVlanOutput) RadiusServer() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerVlan) pulumi.StringOutput { return v.RadiusServer }).(pulumi.StringOutput)
}

// Security. Valid values: `open`, `captive-portal`, `8021x`.
func (o SwitchcontrollerVlanOutput) Security() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerVlan) pulumi.StringOutput { return v.Security }).(pulumi.StringOutput)
}

// Selected user group. The structure of `selectedUsergroups` block is documented below.
func (o SwitchcontrollerVlanOutput) SelectedUsergroups() SwitchcontrollerVlanSelectedUsergroupArrayOutput {
	return o.ApplyT(func(v *SwitchcontrollerVlan) SwitchcontrollerVlanSelectedUsergroupArrayOutput {
		return v.SelectedUsergroups
	}).(SwitchcontrollerVlanSelectedUsergroupArrayOutput)
}

// Authentication usergroup.
func (o SwitchcontrollerVlanOutput) Usergroup() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerVlan) pulumi.StringOutput { return v.Usergroup }).(pulumi.StringOutput)
}

// Virtual domain,
func (o SwitchcontrollerVlanOutput) Vdom() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerVlan) pulumi.StringOutput { return v.Vdom }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SwitchcontrollerVlanOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerVlan) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// VLAN ID.
func (o SwitchcontrollerVlanOutput) Vlanid() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerVlan) pulumi.IntOutput { return v.Vlanid }).(pulumi.IntOutput)
}

type SwitchcontrollerVlanArrayOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerVlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerVlan)(nil)).Elem()
}

func (o SwitchcontrollerVlanArrayOutput) ToSwitchcontrollerVlanArrayOutput() SwitchcontrollerVlanArrayOutput {
	return o
}

func (o SwitchcontrollerVlanArrayOutput) ToSwitchcontrollerVlanArrayOutputWithContext(ctx context.Context) SwitchcontrollerVlanArrayOutput {
	return o
}

func (o SwitchcontrollerVlanArrayOutput) Index(i pulumi.IntInput) SwitchcontrollerVlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchcontrollerVlan {
		return vs[0].([]*SwitchcontrollerVlan)[vs[1].(int)]
	}).(SwitchcontrollerVlanOutput)
}

type SwitchcontrollerVlanMapOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerVlanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerVlan)(nil)).Elem()
}

func (o SwitchcontrollerVlanMapOutput) ToSwitchcontrollerVlanMapOutput() SwitchcontrollerVlanMapOutput {
	return o
}

func (o SwitchcontrollerVlanMapOutput) ToSwitchcontrollerVlanMapOutputWithContext(ctx context.Context) SwitchcontrollerVlanMapOutput {
	return o
}

func (o SwitchcontrollerVlanMapOutput) MapIndex(k pulumi.StringInput) SwitchcontrollerVlanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchcontrollerVlan {
		return vs[0].(map[string]*SwitchcontrollerVlan)[vs[1].(string)]
	}).(SwitchcontrollerVlanOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerVlanInput)(nil)).Elem(), &SwitchcontrollerVlan{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerVlanArrayInput)(nil)).Elem(), SwitchcontrollerVlanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerVlanMapInput)(nil)).Elem(), SwitchcontrollerVlanMap{})
	pulumi.RegisterOutputType(SwitchcontrollerVlanOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerVlanArrayOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerVlanMapOutput{})
}