// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Coordinate federated upgrades within the Security Fabric. Applies to FortiOS Version `>= 7.0.0`.
//
// ## Import
//
// # System FederatedUpgrade can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemFederatedupgrade:SystemFederatedupgrade labelname SystemFederatedUpgrade
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemFederatedupgrade:SystemFederatedupgrade labelname SystemFederatedUpgrade
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemFederatedupgrade struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Serial number of the node to include.
	FailureDevice pulumi.StringOutput `pulumi:"failureDevice"`
	// Reason for upgrade failure. Valid values: `none`, `internal`, `timeout`, `device-type-unsupported`, `download-failed`, `device-missing`, `version-unavailable`, `staging-failed`, `reboot-failed`, `device-not-reconnected`, `node-not-ready`, `no-final-confirmation`, `no-confirmation-query`.
	FailureReason pulumi.StringOutput `pulumi:"failureReason"`
	// The index of the next image to upgrade to.
	NextPathIndex pulumi.IntOutput `pulumi:"nextPathIndex"`
	// Nodes which will be included in the upgrade. The structure of `nodeList` block is documented below.
	NodeLists SystemFederatedupgradeNodeListArrayOutput `pulumi:"nodeLists"`
	// Current status of the upgrade.
	Status pulumi.StringOutput `pulumi:"status"`
	// Unique identifier for this upgrade.
	UpgradeId pulumi.IntOutput `pulumi:"upgradeId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemFederatedupgrade registers a new resource with the given unique name, arguments, and options.
func NewSystemFederatedupgrade(ctx *pulumi.Context,
	name string, args *SystemFederatedupgradeArgs, opts ...pulumi.ResourceOption) (*SystemFederatedupgrade, error) {
	if args == nil {
		args = &SystemFederatedupgradeArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemFederatedupgrade
	err := ctx.RegisterResource("fortios:index/systemFederatedupgrade:SystemFederatedupgrade", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemFederatedupgrade gets an existing SystemFederatedupgrade resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemFederatedupgrade(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemFederatedupgradeState, opts ...pulumi.ResourceOption) (*SystemFederatedupgrade, error) {
	var resource SystemFederatedupgrade
	err := ctx.ReadResource("fortios:index/systemFederatedupgrade:SystemFederatedupgrade", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemFederatedupgrade resources.
type systemFederatedupgradeState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Serial number of the node to include.
	FailureDevice *string `pulumi:"failureDevice"`
	// Reason for upgrade failure. Valid values: `none`, `internal`, `timeout`, `device-type-unsupported`, `download-failed`, `device-missing`, `version-unavailable`, `staging-failed`, `reboot-failed`, `device-not-reconnected`, `node-not-ready`, `no-final-confirmation`, `no-confirmation-query`.
	FailureReason *string `pulumi:"failureReason"`
	// The index of the next image to upgrade to.
	NextPathIndex *int `pulumi:"nextPathIndex"`
	// Nodes which will be included in the upgrade. The structure of `nodeList` block is documented below.
	NodeLists []SystemFederatedupgradeNodeList `pulumi:"nodeLists"`
	// Current status of the upgrade.
	Status *string `pulumi:"status"`
	// Unique identifier for this upgrade.
	UpgradeId *int `pulumi:"upgradeId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemFederatedupgradeState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Serial number of the node to include.
	FailureDevice pulumi.StringPtrInput
	// Reason for upgrade failure. Valid values: `none`, `internal`, `timeout`, `device-type-unsupported`, `download-failed`, `device-missing`, `version-unavailable`, `staging-failed`, `reboot-failed`, `device-not-reconnected`, `node-not-ready`, `no-final-confirmation`, `no-confirmation-query`.
	FailureReason pulumi.StringPtrInput
	// The index of the next image to upgrade to.
	NextPathIndex pulumi.IntPtrInput
	// Nodes which will be included in the upgrade. The structure of `nodeList` block is documented below.
	NodeLists SystemFederatedupgradeNodeListArrayInput
	// Current status of the upgrade.
	Status pulumi.StringPtrInput
	// Unique identifier for this upgrade.
	UpgradeId pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemFederatedupgradeState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFederatedupgradeState)(nil)).Elem()
}

type systemFederatedupgradeArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Serial number of the node to include.
	FailureDevice *string `pulumi:"failureDevice"`
	// Reason for upgrade failure. Valid values: `none`, `internal`, `timeout`, `device-type-unsupported`, `download-failed`, `device-missing`, `version-unavailable`, `staging-failed`, `reboot-failed`, `device-not-reconnected`, `node-not-ready`, `no-final-confirmation`, `no-confirmation-query`.
	FailureReason *string `pulumi:"failureReason"`
	// The index of the next image to upgrade to.
	NextPathIndex *int `pulumi:"nextPathIndex"`
	// Nodes which will be included in the upgrade. The structure of `nodeList` block is documented below.
	NodeLists []SystemFederatedupgradeNodeList `pulumi:"nodeLists"`
	// Current status of the upgrade.
	Status *string `pulumi:"status"`
	// Unique identifier for this upgrade.
	UpgradeId *int `pulumi:"upgradeId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemFederatedupgrade resource.
type SystemFederatedupgradeArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Serial number of the node to include.
	FailureDevice pulumi.StringPtrInput
	// Reason for upgrade failure. Valid values: `none`, `internal`, `timeout`, `device-type-unsupported`, `download-failed`, `device-missing`, `version-unavailable`, `staging-failed`, `reboot-failed`, `device-not-reconnected`, `node-not-ready`, `no-final-confirmation`, `no-confirmation-query`.
	FailureReason pulumi.StringPtrInput
	// The index of the next image to upgrade to.
	NextPathIndex pulumi.IntPtrInput
	// Nodes which will be included in the upgrade. The structure of `nodeList` block is documented below.
	NodeLists SystemFederatedupgradeNodeListArrayInput
	// Current status of the upgrade.
	Status pulumi.StringPtrInput
	// Unique identifier for this upgrade.
	UpgradeId pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemFederatedupgradeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFederatedupgradeArgs)(nil)).Elem()
}

type SystemFederatedupgradeInput interface {
	pulumi.Input

	ToSystemFederatedupgradeOutput() SystemFederatedupgradeOutput
	ToSystemFederatedupgradeOutputWithContext(ctx context.Context) SystemFederatedupgradeOutput
}

func (*SystemFederatedupgrade) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFederatedupgrade)(nil)).Elem()
}

func (i *SystemFederatedupgrade) ToSystemFederatedupgradeOutput() SystemFederatedupgradeOutput {
	return i.ToSystemFederatedupgradeOutputWithContext(context.Background())
}

func (i *SystemFederatedupgrade) ToSystemFederatedupgradeOutputWithContext(ctx context.Context) SystemFederatedupgradeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFederatedupgradeOutput)
}

// SystemFederatedupgradeArrayInput is an input type that accepts SystemFederatedupgradeArray and SystemFederatedupgradeArrayOutput values.
// You can construct a concrete instance of `SystemFederatedupgradeArrayInput` via:
//
//	SystemFederatedupgradeArray{ SystemFederatedupgradeArgs{...} }
type SystemFederatedupgradeArrayInput interface {
	pulumi.Input

	ToSystemFederatedupgradeArrayOutput() SystemFederatedupgradeArrayOutput
	ToSystemFederatedupgradeArrayOutputWithContext(context.Context) SystemFederatedupgradeArrayOutput
}

type SystemFederatedupgradeArray []SystemFederatedupgradeInput

func (SystemFederatedupgradeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFederatedupgrade)(nil)).Elem()
}

func (i SystemFederatedupgradeArray) ToSystemFederatedupgradeArrayOutput() SystemFederatedupgradeArrayOutput {
	return i.ToSystemFederatedupgradeArrayOutputWithContext(context.Background())
}

func (i SystemFederatedupgradeArray) ToSystemFederatedupgradeArrayOutputWithContext(ctx context.Context) SystemFederatedupgradeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFederatedupgradeArrayOutput)
}

// SystemFederatedupgradeMapInput is an input type that accepts SystemFederatedupgradeMap and SystemFederatedupgradeMapOutput values.
// You can construct a concrete instance of `SystemFederatedupgradeMapInput` via:
//
//	SystemFederatedupgradeMap{ "key": SystemFederatedupgradeArgs{...} }
type SystemFederatedupgradeMapInput interface {
	pulumi.Input

	ToSystemFederatedupgradeMapOutput() SystemFederatedupgradeMapOutput
	ToSystemFederatedupgradeMapOutputWithContext(context.Context) SystemFederatedupgradeMapOutput
}

type SystemFederatedupgradeMap map[string]SystemFederatedupgradeInput

func (SystemFederatedupgradeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFederatedupgrade)(nil)).Elem()
}

func (i SystemFederatedupgradeMap) ToSystemFederatedupgradeMapOutput() SystemFederatedupgradeMapOutput {
	return i.ToSystemFederatedupgradeMapOutputWithContext(context.Background())
}

func (i SystemFederatedupgradeMap) ToSystemFederatedupgradeMapOutputWithContext(ctx context.Context) SystemFederatedupgradeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFederatedupgradeMapOutput)
}

type SystemFederatedupgradeOutput struct{ *pulumi.OutputState }

func (SystemFederatedupgradeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFederatedupgrade)(nil)).Elem()
}

func (o SystemFederatedupgradeOutput) ToSystemFederatedupgradeOutput() SystemFederatedupgradeOutput {
	return o
}

func (o SystemFederatedupgradeOutput) ToSystemFederatedupgradeOutputWithContext(ctx context.Context) SystemFederatedupgradeOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SystemFederatedupgradeOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemFederatedupgrade) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Serial number of the node to include.
func (o SystemFederatedupgradeOutput) FailureDevice() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFederatedupgrade) pulumi.StringOutput { return v.FailureDevice }).(pulumi.StringOutput)
}

// Reason for upgrade failure. Valid values: `none`, `internal`, `timeout`, `device-type-unsupported`, `download-failed`, `device-missing`, `version-unavailable`, `staging-failed`, `reboot-failed`, `device-not-reconnected`, `node-not-ready`, `no-final-confirmation`, `no-confirmation-query`.
func (o SystemFederatedupgradeOutput) FailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFederatedupgrade) pulumi.StringOutput { return v.FailureReason }).(pulumi.StringOutput)
}

// The index of the next image to upgrade to.
func (o SystemFederatedupgradeOutput) NextPathIndex() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemFederatedupgrade) pulumi.IntOutput { return v.NextPathIndex }).(pulumi.IntOutput)
}

// Nodes which will be included in the upgrade. The structure of `nodeList` block is documented below.
func (o SystemFederatedupgradeOutput) NodeLists() SystemFederatedupgradeNodeListArrayOutput {
	return o.ApplyT(func(v *SystemFederatedupgrade) SystemFederatedupgradeNodeListArrayOutput { return v.NodeLists }).(SystemFederatedupgradeNodeListArrayOutput)
}

// Current status of the upgrade.
func (o SystemFederatedupgradeOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFederatedupgrade) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Unique identifier for this upgrade.
func (o SystemFederatedupgradeOutput) UpgradeId() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemFederatedupgrade) pulumi.IntOutput { return v.UpgradeId }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemFederatedupgradeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemFederatedupgrade) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemFederatedupgradeArrayOutput struct{ *pulumi.OutputState }

func (SystemFederatedupgradeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFederatedupgrade)(nil)).Elem()
}

func (o SystemFederatedupgradeArrayOutput) ToSystemFederatedupgradeArrayOutput() SystemFederatedupgradeArrayOutput {
	return o
}

func (o SystemFederatedupgradeArrayOutput) ToSystemFederatedupgradeArrayOutputWithContext(ctx context.Context) SystemFederatedupgradeArrayOutput {
	return o
}

func (o SystemFederatedupgradeArrayOutput) Index(i pulumi.IntInput) SystemFederatedupgradeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemFederatedupgrade {
		return vs[0].([]*SystemFederatedupgrade)[vs[1].(int)]
	}).(SystemFederatedupgradeOutput)
}

type SystemFederatedupgradeMapOutput struct{ *pulumi.OutputState }

func (SystemFederatedupgradeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFederatedupgrade)(nil)).Elem()
}

func (o SystemFederatedupgradeMapOutput) ToSystemFederatedupgradeMapOutput() SystemFederatedupgradeMapOutput {
	return o
}

func (o SystemFederatedupgradeMapOutput) ToSystemFederatedupgradeMapOutputWithContext(ctx context.Context) SystemFederatedupgradeMapOutput {
	return o
}

func (o SystemFederatedupgradeMapOutput) MapIndex(k pulumi.StringInput) SystemFederatedupgradeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemFederatedupgrade {
		return vs[0].(map[string]*SystemFederatedupgrade)[vs[1].(string)]
	}).(SystemFederatedupgradeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFederatedupgradeInput)(nil)).Elem(), &SystemFederatedupgrade{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFederatedupgradeArrayInput)(nil)).Elem(), SystemFederatedupgradeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFederatedupgradeMapInput)(nil)).Elem(), SystemFederatedupgradeMap{})
	pulumi.RegisterOutputType(SystemFederatedupgradeOutput{})
	pulumi.RegisterOutputType(SystemFederatedupgradeArrayOutput{})
	pulumi.RegisterOutputType(SystemFederatedupgradeMapOutput{})
}
