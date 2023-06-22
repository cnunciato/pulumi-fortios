// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Filters for FortiAnalyzer.
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
//			_, err := fortios.NewLogfortianalyzer2Filter(ctx, "trname", &fortios.Logfortianalyzer2FilterArgs{
//				Anomaly:          pulumi.String("enable"),
//				DlpArchive:       pulumi.String("enable"),
//				Dns:              pulumi.String("enable"),
//				FilterType:       pulumi.String("include"),
//				ForwardTraffic:   pulumi.String("enable"),
//				Gtp:              pulumi.String("enable"),
//				LocalTraffic:     pulumi.String("enable"),
//				MulticastTraffic: pulumi.String("enable"),
//				Severity:         pulumi.String("information"),
//				SnifferTraffic:   pulumi.String("enable"),
//				Ssh:              pulumi.String("enable"),
//				Voip:             pulumi.String("enable"),
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
// # LogFortianalyzer2 Filter can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/logfortianalyzer2Filter:Logfortianalyzer2Filter labelname LogFortianalyzer2Filter
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/logfortianalyzer2Filter:Logfortianalyzer2Filter labelname LogFortianalyzer2Filter
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Logfortianalyzer2Filter struct {
	pulumi.CustomResourceState

	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringOutput `pulumi:"anomaly"`
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive pulumi.StringOutput `pulumi:"dlpArchive"`
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns pulumi.StringOutput `pulumi:"dns"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// FortiAnalyzer 2 log filter.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringOutput `pulumi:"filterType"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringOutput `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles Logfortianalyzer2FilterFreeStyleArrayOutput `pulumi:"freeStyles"`
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp pulumi.StringOutput `pulumi:"gtp"`
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic pulumi.StringOutput `pulumi:"localTraffic"`
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic pulumi.StringOutput `pulumi:"multicastTraffic"`
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery pulumi.StringOutput `pulumi:"netscanDiscovery"`
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability pulumi.StringOutput `pulumi:"netscanVulnerability"`
	// Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic pulumi.StringOutput `pulumi:"snifferTraffic"`
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh pulumi.StringOutput `pulumi:"ssh"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip pulumi.StringOutput `pulumi:"voip"`
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic pulumi.StringOutput `pulumi:"ztnaTraffic"`
}

// NewLogfortianalyzer2Filter registers a new resource with the given unique name, arguments, and options.
func NewLogfortianalyzer2Filter(ctx *pulumi.Context,
	name string, args *Logfortianalyzer2FilterArgs, opts ...pulumi.ResourceOption) (*Logfortianalyzer2Filter, error) {
	if args == nil {
		args = &Logfortianalyzer2FilterArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Logfortianalyzer2Filter
	err := ctx.RegisterResource("fortios:index/logfortianalyzer2Filter:Logfortianalyzer2Filter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogfortianalyzer2Filter gets an existing Logfortianalyzer2Filter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogfortianalyzer2Filter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Logfortianalyzer2FilterState, opts ...pulumi.ResourceOption) (*Logfortianalyzer2Filter, error) {
	var resource Logfortianalyzer2Filter
	err := ctx.ReadResource("fortios:index/logfortianalyzer2Filter:Logfortianalyzer2Filter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Logfortianalyzer2Filter resources.
type logfortianalyzer2FilterState struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly *string `pulumi:"anomaly"`
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive *string `pulumi:"dlpArchive"`
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns *string `pulumi:"dns"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// FortiAnalyzer 2 log filter.
	Filter *string `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType *string `pulumi:"filterType"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic *string `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles []Logfortianalyzer2FilterFreeStyle `pulumi:"freeStyles"`
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp *string `pulumi:"gtp"`
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic *string `pulumi:"localTraffic"`
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic *string `pulumi:"multicastTraffic"`
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery *string `pulumi:"netscanDiscovery"`
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability *string `pulumi:"netscanVulnerability"`
	// Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity *string `pulumi:"severity"`
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic *string `pulumi:"snifferTraffic"`
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh *string `pulumi:"ssh"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip *string `pulumi:"voip"`
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic *string `pulumi:"ztnaTraffic"`
}

type Logfortianalyzer2FilterState struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringPtrInput
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive pulumi.StringPtrInput
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// FortiAnalyzer 2 log filter.
	Filter pulumi.StringPtrInput
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringPtrInput
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringPtrInput
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles Logfortianalyzer2FilterFreeStyleArrayInput
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp pulumi.StringPtrInput
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic pulumi.StringPtrInput
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic pulumi.StringPtrInput
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery pulumi.StringPtrInput
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability pulumi.StringPtrInput
	// Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringPtrInput
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic pulumi.StringPtrInput
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip pulumi.StringPtrInput
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic pulumi.StringPtrInput
}

func (Logfortianalyzer2FilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logfortianalyzer2FilterState)(nil)).Elem()
}

type logfortianalyzer2FilterArgs struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly *string `pulumi:"anomaly"`
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive *string `pulumi:"dlpArchive"`
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns *string `pulumi:"dns"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// FortiAnalyzer 2 log filter.
	Filter *string `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType *string `pulumi:"filterType"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic *string `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles []Logfortianalyzer2FilterFreeStyle `pulumi:"freeStyles"`
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp *string `pulumi:"gtp"`
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic *string `pulumi:"localTraffic"`
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic *string `pulumi:"multicastTraffic"`
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery *string `pulumi:"netscanDiscovery"`
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability *string `pulumi:"netscanVulnerability"`
	// Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity *string `pulumi:"severity"`
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic *string `pulumi:"snifferTraffic"`
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh *string `pulumi:"ssh"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip *string `pulumi:"voip"`
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic *string `pulumi:"ztnaTraffic"`
}

// The set of arguments for constructing a Logfortianalyzer2Filter resource.
type Logfortianalyzer2FilterArgs struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringPtrInput
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive pulumi.StringPtrInput
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// FortiAnalyzer 2 log filter.
	Filter pulumi.StringPtrInput
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringPtrInput
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringPtrInput
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles Logfortianalyzer2FilterFreeStyleArrayInput
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp pulumi.StringPtrInput
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic pulumi.StringPtrInput
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic pulumi.StringPtrInput
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery pulumi.StringPtrInput
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability pulumi.StringPtrInput
	// Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringPtrInput
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic pulumi.StringPtrInput
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip pulumi.StringPtrInput
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic pulumi.StringPtrInput
}

func (Logfortianalyzer2FilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logfortianalyzer2FilterArgs)(nil)).Elem()
}

type Logfortianalyzer2FilterInput interface {
	pulumi.Input

	ToLogfortianalyzer2FilterOutput() Logfortianalyzer2FilterOutput
	ToLogfortianalyzer2FilterOutputWithContext(ctx context.Context) Logfortianalyzer2FilterOutput
}

func (*Logfortianalyzer2Filter) ElementType() reflect.Type {
	return reflect.TypeOf((**Logfortianalyzer2Filter)(nil)).Elem()
}

func (i *Logfortianalyzer2Filter) ToLogfortianalyzer2FilterOutput() Logfortianalyzer2FilterOutput {
	return i.ToLogfortianalyzer2FilterOutputWithContext(context.Background())
}

func (i *Logfortianalyzer2Filter) ToLogfortianalyzer2FilterOutputWithContext(ctx context.Context) Logfortianalyzer2FilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Logfortianalyzer2FilterOutput)
}

// Logfortianalyzer2FilterArrayInput is an input type that accepts Logfortianalyzer2FilterArray and Logfortianalyzer2FilterArrayOutput values.
// You can construct a concrete instance of `Logfortianalyzer2FilterArrayInput` via:
//
//	Logfortianalyzer2FilterArray{ Logfortianalyzer2FilterArgs{...} }
type Logfortianalyzer2FilterArrayInput interface {
	pulumi.Input

	ToLogfortianalyzer2FilterArrayOutput() Logfortianalyzer2FilterArrayOutput
	ToLogfortianalyzer2FilterArrayOutputWithContext(context.Context) Logfortianalyzer2FilterArrayOutput
}

type Logfortianalyzer2FilterArray []Logfortianalyzer2FilterInput

func (Logfortianalyzer2FilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Logfortianalyzer2Filter)(nil)).Elem()
}

func (i Logfortianalyzer2FilterArray) ToLogfortianalyzer2FilterArrayOutput() Logfortianalyzer2FilterArrayOutput {
	return i.ToLogfortianalyzer2FilterArrayOutputWithContext(context.Background())
}

func (i Logfortianalyzer2FilterArray) ToLogfortianalyzer2FilterArrayOutputWithContext(ctx context.Context) Logfortianalyzer2FilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Logfortianalyzer2FilterArrayOutput)
}

// Logfortianalyzer2FilterMapInput is an input type that accepts Logfortianalyzer2FilterMap and Logfortianalyzer2FilterMapOutput values.
// You can construct a concrete instance of `Logfortianalyzer2FilterMapInput` via:
//
//	Logfortianalyzer2FilterMap{ "key": Logfortianalyzer2FilterArgs{...} }
type Logfortianalyzer2FilterMapInput interface {
	pulumi.Input

	ToLogfortianalyzer2FilterMapOutput() Logfortianalyzer2FilterMapOutput
	ToLogfortianalyzer2FilterMapOutputWithContext(context.Context) Logfortianalyzer2FilterMapOutput
}

type Logfortianalyzer2FilterMap map[string]Logfortianalyzer2FilterInput

func (Logfortianalyzer2FilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Logfortianalyzer2Filter)(nil)).Elem()
}

func (i Logfortianalyzer2FilterMap) ToLogfortianalyzer2FilterMapOutput() Logfortianalyzer2FilterMapOutput {
	return i.ToLogfortianalyzer2FilterMapOutputWithContext(context.Background())
}

func (i Logfortianalyzer2FilterMap) ToLogfortianalyzer2FilterMapOutputWithContext(ctx context.Context) Logfortianalyzer2FilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Logfortianalyzer2FilterMapOutput)
}

type Logfortianalyzer2FilterOutput struct{ *pulumi.OutputState }

func (Logfortianalyzer2FilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Logfortianalyzer2Filter)(nil)).Elem()
}

func (o Logfortianalyzer2FilterOutput) ToLogfortianalyzer2FilterOutput() Logfortianalyzer2FilterOutput {
	return o
}

func (o Logfortianalyzer2FilterOutput) ToLogfortianalyzer2FilterOutputWithContext(ctx context.Context) Logfortianalyzer2FilterOutput {
	return o
}

// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
func (o Logfortianalyzer2FilterOutput) Anomaly() pulumi.StringOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringOutput { return v.Anomaly }).(pulumi.StringOutput)
}

// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
func (o Logfortianalyzer2FilterOutput) DlpArchive() pulumi.StringOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringOutput { return v.DlpArchive }).(pulumi.StringOutput)
}

// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
func (o Logfortianalyzer2FilterOutput) Dns() pulumi.StringOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringOutput { return v.Dns }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o Logfortianalyzer2FilterOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// FortiAnalyzer 2 log filter.
func (o Logfortianalyzer2FilterOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
func (o Logfortianalyzer2FilterOutput) FilterType() pulumi.StringOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringOutput { return v.FilterType }).(pulumi.StringOutput)
}

// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
func (o Logfortianalyzer2FilterOutput) ForwardTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringOutput { return v.ForwardTraffic }).(pulumi.StringOutput)
}

// Free Style Filters The structure of `freeStyle` block is documented below.
func (o Logfortianalyzer2FilterOutput) FreeStyles() Logfortianalyzer2FilterFreeStyleArrayOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) Logfortianalyzer2FilterFreeStyleArrayOutput { return v.FreeStyles }).(Logfortianalyzer2FilterFreeStyleArrayOutput)
}

// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
func (o Logfortianalyzer2FilterOutput) Gtp() pulumi.StringOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringOutput { return v.Gtp }).(pulumi.StringOutput)
}

// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
func (o Logfortianalyzer2FilterOutput) LocalTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringOutput { return v.LocalTraffic }).(pulumi.StringOutput)
}

// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
func (o Logfortianalyzer2FilterOutput) MulticastTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringOutput { return v.MulticastTraffic }).(pulumi.StringOutput)
}

// Enable/disable netscan discovery event logging.
func (o Logfortianalyzer2FilterOutput) NetscanDiscovery() pulumi.StringOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringOutput { return v.NetscanDiscovery }).(pulumi.StringOutput)
}

// Enable/disable netscan vulnerability event logging.
func (o Logfortianalyzer2FilterOutput) NetscanVulnerability() pulumi.StringOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringOutput { return v.NetscanVulnerability }).(pulumi.StringOutput)
}

// Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
func (o Logfortianalyzer2FilterOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
func (o Logfortianalyzer2FilterOutput) SnifferTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringOutput { return v.SnifferTraffic }).(pulumi.StringOutput)
}

// Enable/disable SSH logging. Valid values: `enable`, `disable`.
func (o Logfortianalyzer2FilterOutput) Ssh() pulumi.StringOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringOutput { return v.Ssh }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Logfortianalyzer2FilterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
func (o Logfortianalyzer2FilterOutput) Voip() pulumi.StringOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringOutput { return v.Voip }).(pulumi.StringOutput)
}

// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
func (o Logfortianalyzer2FilterOutput) ZtnaTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Logfortianalyzer2Filter) pulumi.StringOutput { return v.ZtnaTraffic }).(pulumi.StringOutput)
}

type Logfortianalyzer2FilterArrayOutput struct{ *pulumi.OutputState }

func (Logfortianalyzer2FilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Logfortianalyzer2Filter)(nil)).Elem()
}

func (o Logfortianalyzer2FilterArrayOutput) ToLogfortianalyzer2FilterArrayOutput() Logfortianalyzer2FilterArrayOutput {
	return o
}

func (o Logfortianalyzer2FilterArrayOutput) ToLogfortianalyzer2FilterArrayOutputWithContext(ctx context.Context) Logfortianalyzer2FilterArrayOutput {
	return o
}

func (o Logfortianalyzer2FilterArrayOutput) Index(i pulumi.IntInput) Logfortianalyzer2FilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Logfortianalyzer2Filter {
		return vs[0].([]*Logfortianalyzer2Filter)[vs[1].(int)]
	}).(Logfortianalyzer2FilterOutput)
}

type Logfortianalyzer2FilterMapOutput struct{ *pulumi.OutputState }

func (Logfortianalyzer2FilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Logfortianalyzer2Filter)(nil)).Elem()
}

func (o Logfortianalyzer2FilterMapOutput) ToLogfortianalyzer2FilterMapOutput() Logfortianalyzer2FilterMapOutput {
	return o
}

func (o Logfortianalyzer2FilterMapOutput) ToLogfortianalyzer2FilterMapOutputWithContext(ctx context.Context) Logfortianalyzer2FilterMapOutput {
	return o
}

func (o Logfortianalyzer2FilterMapOutput) MapIndex(k pulumi.StringInput) Logfortianalyzer2FilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Logfortianalyzer2Filter {
		return vs[0].(map[string]*Logfortianalyzer2Filter)[vs[1].(string)]
	}).(Logfortianalyzer2FilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Logfortianalyzer2FilterInput)(nil)).Elem(), &Logfortianalyzer2Filter{})
	pulumi.RegisterInputType(reflect.TypeOf((*Logfortianalyzer2FilterArrayInput)(nil)).Elem(), Logfortianalyzer2FilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Logfortianalyzer2FilterMapInput)(nil)).Elem(), Logfortianalyzer2FilterMap{})
	pulumi.RegisterOutputType(Logfortianalyzer2FilterOutput{})
	pulumi.RegisterOutputType(Logfortianalyzer2FilterArrayOutput{})
	pulumi.RegisterOutputType(Logfortianalyzer2FilterMapOutput{})
}