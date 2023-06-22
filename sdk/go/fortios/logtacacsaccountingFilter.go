// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Settings for TACACS+ accounting events filter. Applies to FortiOS Version `>= 7.0.2`.
//
// ## Import
//
// # LogTacacsAccounting Filter can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/logtacacsaccountingFilter:LogtacacsaccountingFilter labelname LogTacacsAccountingFilter
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/logtacacsaccountingFilter:LogtacacsaccountingFilter labelname LogTacacsAccountingFilter
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type LogtacacsaccountingFilter struct {
	pulumi.CustomResourceState

	// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
	CliCmdAudit pulumi.StringOutput `pulumi:"cliCmdAudit"`
	// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
	ConfigChangeAudit pulumi.StringOutput `pulumi:"configChangeAudit"`
	// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
	LoginAudit pulumi.StringOutput `pulumi:"loginAudit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewLogtacacsaccountingFilter registers a new resource with the given unique name, arguments, and options.
func NewLogtacacsaccountingFilter(ctx *pulumi.Context,
	name string, args *LogtacacsaccountingFilterArgs, opts ...pulumi.ResourceOption) (*LogtacacsaccountingFilter, error) {
	if args == nil {
		args = &LogtacacsaccountingFilterArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogtacacsaccountingFilter
	err := ctx.RegisterResource("fortios:index/logtacacsaccountingFilter:LogtacacsaccountingFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogtacacsaccountingFilter gets an existing LogtacacsaccountingFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogtacacsaccountingFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogtacacsaccountingFilterState, opts ...pulumi.ResourceOption) (*LogtacacsaccountingFilter, error) {
	var resource LogtacacsaccountingFilter
	err := ctx.ReadResource("fortios:index/logtacacsaccountingFilter:LogtacacsaccountingFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogtacacsaccountingFilter resources.
type logtacacsaccountingFilterState struct {
	// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
	CliCmdAudit *string `pulumi:"cliCmdAudit"`
	// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
	ConfigChangeAudit *string `pulumi:"configChangeAudit"`
	// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
	LoginAudit *string `pulumi:"loginAudit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type LogtacacsaccountingFilterState struct {
	// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
	CliCmdAudit pulumi.StringPtrInput
	// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
	ConfigChangeAudit pulumi.StringPtrInput
	// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
	LoginAudit pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogtacacsaccountingFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logtacacsaccountingFilterState)(nil)).Elem()
}

type logtacacsaccountingFilterArgs struct {
	// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
	CliCmdAudit *string `pulumi:"cliCmdAudit"`
	// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
	ConfigChangeAudit *string `pulumi:"configChangeAudit"`
	// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
	LoginAudit *string `pulumi:"loginAudit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogtacacsaccountingFilter resource.
type LogtacacsaccountingFilterArgs struct {
	// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
	CliCmdAudit pulumi.StringPtrInput
	// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
	ConfigChangeAudit pulumi.StringPtrInput
	// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
	LoginAudit pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogtacacsaccountingFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logtacacsaccountingFilterArgs)(nil)).Elem()
}

type LogtacacsaccountingFilterInput interface {
	pulumi.Input

	ToLogtacacsaccountingFilterOutput() LogtacacsaccountingFilterOutput
	ToLogtacacsaccountingFilterOutputWithContext(ctx context.Context) LogtacacsaccountingFilterOutput
}

func (*LogtacacsaccountingFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogtacacsaccountingFilter)(nil)).Elem()
}

func (i *LogtacacsaccountingFilter) ToLogtacacsaccountingFilterOutput() LogtacacsaccountingFilterOutput {
	return i.ToLogtacacsaccountingFilterOutputWithContext(context.Background())
}

func (i *LogtacacsaccountingFilter) ToLogtacacsaccountingFilterOutputWithContext(ctx context.Context) LogtacacsaccountingFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogtacacsaccountingFilterOutput)
}

// LogtacacsaccountingFilterArrayInput is an input type that accepts LogtacacsaccountingFilterArray and LogtacacsaccountingFilterArrayOutput values.
// You can construct a concrete instance of `LogtacacsaccountingFilterArrayInput` via:
//
//	LogtacacsaccountingFilterArray{ LogtacacsaccountingFilterArgs{...} }
type LogtacacsaccountingFilterArrayInput interface {
	pulumi.Input

	ToLogtacacsaccountingFilterArrayOutput() LogtacacsaccountingFilterArrayOutput
	ToLogtacacsaccountingFilterArrayOutputWithContext(context.Context) LogtacacsaccountingFilterArrayOutput
}

type LogtacacsaccountingFilterArray []LogtacacsaccountingFilterInput

func (LogtacacsaccountingFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogtacacsaccountingFilter)(nil)).Elem()
}

func (i LogtacacsaccountingFilterArray) ToLogtacacsaccountingFilterArrayOutput() LogtacacsaccountingFilterArrayOutput {
	return i.ToLogtacacsaccountingFilterArrayOutputWithContext(context.Background())
}

func (i LogtacacsaccountingFilterArray) ToLogtacacsaccountingFilterArrayOutputWithContext(ctx context.Context) LogtacacsaccountingFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogtacacsaccountingFilterArrayOutput)
}

// LogtacacsaccountingFilterMapInput is an input type that accepts LogtacacsaccountingFilterMap and LogtacacsaccountingFilterMapOutput values.
// You can construct a concrete instance of `LogtacacsaccountingFilterMapInput` via:
//
//	LogtacacsaccountingFilterMap{ "key": LogtacacsaccountingFilterArgs{...} }
type LogtacacsaccountingFilterMapInput interface {
	pulumi.Input

	ToLogtacacsaccountingFilterMapOutput() LogtacacsaccountingFilterMapOutput
	ToLogtacacsaccountingFilterMapOutputWithContext(context.Context) LogtacacsaccountingFilterMapOutput
}

type LogtacacsaccountingFilterMap map[string]LogtacacsaccountingFilterInput

func (LogtacacsaccountingFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogtacacsaccountingFilter)(nil)).Elem()
}

func (i LogtacacsaccountingFilterMap) ToLogtacacsaccountingFilterMapOutput() LogtacacsaccountingFilterMapOutput {
	return i.ToLogtacacsaccountingFilterMapOutputWithContext(context.Background())
}

func (i LogtacacsaccountingFilterMap) ToLogtacacsaccountingFilterMapOutputWithContext(ctx context.Context) LogtacacsaccountingFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogtacacsaccountingFilterMapOutput)
}

type LogtacacsaccountingFilterOutput struct{ *pulumi.OutputState }

func (LogtacacsaccountingFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogtacacsaccountingFilter)(nil)).Elem()
}

func (o LogtacacsaccountingFilterOutput) ToLogtacacsaccountingFilterOutput() LogtacacsaccountingFilterOutput {
	return o
}

func (o LogtacacsaccountingFilterOutput) ToLogtacacsaccountingFilterOutputWithContext(ctx context.Context) LogtacacsaccountingFilterOutput {
	return o
}

// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
func (o LogtacacsaccountingFilterOutput) CliCmdAudit() pulumi.StringOutput {
	return o.ApplyT(func(v *LogtacacsaccountingFilter) pulumi.StringOutput { return v.CliCmdAudit }).(pulumi.StringOutput)
}

// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
func (o LogtacacsaccountingFilterOutput) ConfigChangeAudit() pulumi.StringOutput {
	return o.ApplyT(func(v *LogtacacsaccountingFilter) pulumi.StringOutput { return v.ConfigChangeAudit }).(pulumi.StringOutput)
}

// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
func (o LogtacacsaccountingFilterOutput) LoginAudit() pulumi.StringOutput {
	return o.ApplyT(func(v *LogtacacsaccountingFilter) pulumi.StringOutput { return v.LoginAudit }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o LogtacacsaccountingFilterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogtacacsaccountingFilter) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type LogtacacsaccountingFilterArrayOutput struct{ *pulumi.OutputState }

func (LogtacacsaccountingFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogtacacsaccountingFilter)(nil)).Elem()
}

func (o LogtacacsaccountingFilterArrayOutput) ToLogtacacsaccountingFilterArrayOutput() LogtacacsaccountingFilterArrayOutput {
	return o
}

func (o LogtacacsaccountingFilterArrayOutput) ToLogtacacsaccountingFilterArrayOutputWithContext(ctx context.Context) LogtacacsaccountingFilterArrayOutput {
	return o
}

func (o LogtacacsaccountingFilterArrayOutput) Index(i pulumi.IntInput) LogtacacsaccountingFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogtacacsaccountingFilter {
		return vs[0].([]*LogtacacsaccountingFilter)[vs[1].(int)]
	}).(LogtacacsaccountingFilterOutput)
}

type LogtacacsaccountingFilterMapOutput struct{ *pulumi.OutputState }

func (LogtacacsaccountingFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogtacacsaccountingFilter)(nil)).Elem()
}

func (o LogtacacsaccountingFilterMapOutput) ToLogtacacsaccountingFilterMapOutput() LogtacacsaccountingFilterMapOutput {
	return o
}

func (o LogtacacsaccountingFilterMapOutput) ToLogtacacsaccountingFilterMapOutputWithContext(ctx context.Context) LogtacacsaccountingFilterMapOutput {
	return o
}

func (o LogtacacsaccountingFilterMapOutput) MapIndex(k pulumi.StringInput) LogtacacsaccountingFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogtacacsaccountingFilter {
		return vs[0].(map[string]*LogtacacsaccountingFilter)[vs[1].(string)]
	}).(LogtacacsaccountingFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogtacacsaccountingFilterInput)(nil)).Elem(), &LogtacacsaccountingFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogtacacsaccountingFilterArrayInput)(nil)).Elem(), LogtacacsaccountingFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogtacacsaccountingFilterMapInput)(nil)).Elem(), LogtacacsaccountingFilterMap{})
	pulumi.RegisterOutputType(LogtacacsaccountingFilterOutput{})
	pulumi.RegisterOutputType(LogtacacsaccountingFilterArrayOutput{})
	pulumi.RegisterOutputType(LogtacacsaccountingFilterMapOutput{})
}