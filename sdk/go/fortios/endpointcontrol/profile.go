// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package endpointcontrol

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure FortiClient endpoint control profiles. Applies to FortiOS Version `<= 6.2.0`.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/endpointcontrol"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := endpointcontrol.NewProfile(ctx, "trname", &endpointcontrol.ProfileArgs{
//				DeviceGroups: endpointcontrol.ProfileDeviceGroupArray{
//					&endpointcontrol.ProfileDeviceGroupArgs{
//						Name: pulumi.String("Mobile Devices"),
//					},
//				},
//				ForticlientAndroidSettings: &endpointcontrol.ProfileForticlientAndroidSettingsArgs{
//					DisableWfWhenProtected:     pulumi.String("enable"),
//					ForticlientAdvancedVpn:     pulumi.String("disable"),
//					ForticlientVpnProvisioning: pulumi.String("disable"),
//					ForticlientWf:              pulumi.String("disable"),
//				},
//				ForticlientIosSettings: &endpointcontrol.ProfileForticlientIosSettingsArgs{
//					ClientVpnProvisioning:          pulumi.String("disable"),
//					DisableWfWhenProtected:         pulumi.String("enable"),
//					DistributeConfigurationProfile: pulumi.String("disable"),
//					ForticlientWf:                  pulumi.String("disable"),
//				},
//				ForticlientWinmacSettings: &endpointcontrol.ProfileForticlientWinmacSettingsArgs{
//					AvRealtimeProtection:                       pulumi.String("disable"),
//					AvSignatureUpToDate:                        pulumi.String("disable"),
//					ForticlientApplicationFirewall:             pulumi.String("disable"),
//					ForticlientAv:                              pulumi.String("disable"),
//					ForticlientEmsCompliance:                   pulumi.String("disable"),
//					ForticlientEmsComplianceAction:             pulumi.String("warning"),
//					ForticlientLinuxVer:                        pulumi.String("5.4.1"),
//					ForticlientLogUpload:                       pulumi.String("enable"),
//					ForticlientLogUploadLevel:                  pulumi.String("traffic vulnerability event"),
//					ForticlientMacVer:                          pulumi.String("5.4.1"),
//					ForticlientMinimumSoftwareVersion:          pulumi.String("disable"),
//					ForticlientRegistrationComplianceAction:    pulumi.String("warning"),
//					ForticlientSecurityPosture:                 pulumi.String("disable"),
//					ForticlientSecurityPostureComplianceAction: pulumi.String("warning"),
//					ForticlientSystemCompliance:                pulumi.String("enable"),
//					ForticlientSystemComplianceAction:          pulumi.String("warning"),
//					ForticlientVulnScan:                        pulumi.String("enable"),
//					ForticlientVulnScanComplianceAction:        pulumi.String("warning"),
//					ForticlientVulnScanEnforce:                 pulumi.String("high"),
//					ForticlientVulnScanEnforceGrace:            pulumi.Int(1),
//					ForticlientVulnScanExempt:                  pulumi.String("disable"),
//					ForticlientWf:                              pulumi.String("disable"),
//					ForticlientWinVer:                          pulumi.String("5.4.1"),
//					OsAvSoftwareInstalled:                      pulumi.String("disable"),
//					SandboxAnalysis:                            pulumi.String("disable"),
//				},
//				OnNetAddrs: endpointcontrol.ProfileOnNetAddrArray{
//					&endpointcontrol.ProfileOnNetAddrArgs{
//						Name: pulumi.String("all"),
//					},
//				},
//				ProfileName: pulumi.String("1"),
//				Users: endpointcontrol.ProfileUserArray{
//					&endpointcontrol.ProfileUserArgs{
//						Name: pulumi.String("guest"),
//					},
//				},
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
// EndpointControl Profile can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:endpointcontrol/profile:Profile labelname {{profile_name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:endpointcontrol/profile:Profile labelname {{profile_name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Profile struct {
	pulumi.CustomResourceState

	// Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Device groups. The structure of `deviceGroups` block is documented below.
	DeviceGroups ProfileDeviceGroupArrayOutput `pulumi:"deviceGroups"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// FortiClient settings for Android platform. The structure of `forticlientAndroidSettings` block is documented below.
	ForticlientAndroidSettings ProfileForticlientAndroidSettingsOutput `pulumi:"forticlientAndroidSettings"`
	// FortiClient settings for iOS platform. The structure of `forticlientIosSettings` block is documented below.
	ForticlientIosSettings ProfileForticlientIosSettingsOutput `pulumi:"forticlientIosSettings"`
	// FortiClient settings for Windows/Mac platform. The structure of `forticlientWinmacSettings` block is documented below.
	ForticlientWinmacSettings ProfileForticlientWinmacSettingsOutput `pulumi:"forticlientWinmacSettings"`
	// Addresses for on-net detection. The structure of `onNetAddr` block is documented below.
	OnNetAddrs ProfileOnNetAddrArrayOutput `pulumi:"onNetAddrs"`
	// Profile name.
	ProfileName pulumi.StringOutput `pulumi:"profileName"`
	// Select an endpoint control replacement message override group from available options.
	ReplacemsgOverrideGroup pulumi.StringOutput `pulumi:"replacemsgOverrideGroup"`
	// Source addresses. The structure of `srcAddr` block is documented below.
	SrcAddrs ProfileSrcAddrArrayOutput `pulumi:"srcAddrs"`
	// User groups. The structure of `userGroups` block is documented below.
	UserGroups ProfileUserGroupArrayOutput `pulumi:"userGroups"`
	// Users. The structure of `users` block is documented below.
	Users ProfileUserArrayOutput `pulumi:"users"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewProfile registers a new resource with the given unique name, arguments, and options.
func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil {
		args = &ProfileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Profile
	err := ctx.RegisterResource("fortios:endpointcontrol/profile:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfile gets an existing Profile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("fortios:endpointcontrol/profile:Profile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Profile resources.
type profileState struct {
	// Description.
	Description *string `pulumi:"description"`
	// Device groups. The structure of `deviceGroups` block is documented below.
	DeviceGroups []ProfileDeviceGroup `pulumi:"deviceGroups"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// FortiClient settings for Android platform. The structure of `forticlientAndroidSettings` block is documented below.
	ForticlientAndroidSettings *ProfileForticlientAndroidSettings `pulumi:"forticlientAndroidSettings"`
	// FortiClient settings for iOS platform. The structure of `forticlientIosSettings` block is documented below.
	ForticlientIosSettings *ProfileForticlientIosSettings `pulumi:"forticlientIosSettings"`
	// FortiClient settings for Windows/Mac platform. The structure of `forticlientWinmacSettings` block is documented below.
	ForticlientWinmacSettings *ProfileForticlientWinmacSettings `pulumi:"forticlientWinmacSettings"`
	// Addresses for on-net detection. The structure of `onNetAddr` block is documented below.
	OnNetAddrs []ProfileOnNetAddr `pulumi:"onNetAddrs"`
	// Profile name.
	ProfileName *string `pulumi:"profileName"`
	// Select an endpoint control replacement message override group from available options.
	ReplacemsgOverrideGroup *string `pulumi:"replacemsgOverrideGroup"`
	// Source addresses. The structure of `srcAddr` block is documented below.
	SrcAddrs []ProfileSrcAddr `pulumi:"srcAddrs"`
	// User groups. The structure of `userGroups` block is documented below.
	UserGroups []ProfileUserGroup `pulumi:"userGroups"`
	// Users. The structure of `users` block is documented below.
	Users []ProfileUser `pulumi:"users"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ProfileState struct {
	// Description.
	Description pulumi.StringPtrInput
	// Device groups. The structure of `deviceGroups` block is documented below.
	DeviceGroups ProfileDeviceGroupArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// FortiClient settings for Android platform. The structure of `forticlientAndroidSettings` block is documented below.
	ForticlientAndroidSettings ProfileForticlientAndroidSettingsPtrInput
	// FortiClient settings for iOS platform. The structure of `forticlientIosSettings` block is documented below.
	ForticlientIosSettings ProfileForticlientIosSettingsPtrInput
	// FortiClient settings for Windows/Mac platform. The structure of `forticlientWinmacSettings` block is documented below.
	ForticlientWinmacSettings ProfileForticlientWinmacSettingsPtrInput
	// Addresses for on-net detection. The structure of `onNetAddr` block is documented below.
	OnNetAddrs ProfileOnNetAddrArrayInput
	// Profile name.
	ProfileName pulumi.StringPtrInput
	// Select an endpoint control replacement message override group from available options.
	ReplacemsgOverrideGroup pulumi.StringPtrInput
	// Source addresses. The structure of `srcAddr` block is documented below.
	SrcAddrs ProfileSrcAddrArrayInput
	// User groups. The structure of `userGroups` block is documented below.
	UserGroups ProfileUserGroupArrayInput
	// Users. The structure of `users` block is documented below.
	Users ProfileUserArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileState)(nil)).Elem()
}

type profileArgs struct {
	// Description.
	Description *string `pulumi:"description"`
	// Device groups. The structure of `deviceGroups` block is documented below.
	DeviceGroups []ProfileDeviceGroup `pulumi:"deviceGroups"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// FortiClient settings for Android platform. The structure of `forticlientAndroidSettings` block is documented below.
	ForticlientAndroidSettings *ProfileForticlientAndroidSettings `pulumi:"forticlientAndroidSettings"`
	// FortiClient settings for iOS platform. The structure of `forticlientIosSettings` block is documented below.
	ForticlientIosSettings *ProfileForticlientIosSettings `pulumi:"forticlientIosSettings"`
	// FortiClient settings for Windows/Mac platform. The structure of `forticlientWinmacSettings` block is documented below.
	ForticlientWinmacSettings *ProfileForticlientWinmacSettings `pulumi:"forticlientWinmacSettings"`
	// Addresses for on-net detection. The structure of `onNetAddr` block is documented below.
	OnNetAddrs []ProfileOnNetAddr `pulumi:"onNetAddrs"`
	// Profile name.
	ProfileName *string `pulumi:"profileName"`
	// Select an endpoint control replacement message override group from available options.
	ReplacemsgOverrideGroup *string `pulumi:"replacemsgOverrideGroup"`
	// Source addresses. The structure of `srcAddr` block is documented below.
	SrcAddrs []ProfileSrcAddr `pulumi:"srcAddrs"`
	// User groups. The structure of `userGroups` block is documented below.
	UserGroups []ProfileUserGroup `pulumi:"userGroups"`
	// Users. The structure of `users` block is documented below.
	Users []ProfileUser `pulumi:"users"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Profile resource.
type ProfileArgs struct {
	// Description.
	Description pulumi.StringPtrInput
	// Device groups. The structure of `deviceGroups` block is documented below.
	DeviceGroups ProfileDeviceGroupArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// FortiClient settings for Android platform. The structure of `forticlientAndroidSettings` block is documented below.
	ForticlientAndroidSettings ProfileForticlientAndroidSettingsPtrInput
	// FortiClient settings for iOS platform. The structure of `forticlientIosSettings` block is documented below.
	ForticlientIosSettings ProfileForticlientIosSettingsPtrInput
	// FortiClient settings for Windows/Mac platform. The structure of `forticlientWinmacSettings` block is documented below.
	ForticlientWinmacSettings ProfileForticlientWinmacSettingsPtrInput
	// Addresses for on-net detection. The structure of `onNetAddr` block is documented below.
	OnNetAddrs ProfileOnNetAddrArrayInput
	// Profile name.
	ProfileName pulumi.StringPtrInput
	// Select an endpoint control replacement message override group from available options.
	ReplacemsgOverrideGroup pulumi.StringPtrInput
	// Source addresses. The structure of `srcAddr` block is documented below.
	SrcAddrs ProfileSrcAddrArrayInput
	// User groups. The structure of `userGroups` block is documented below.
	UserGroups ProfileUserGroupArrayInput
	// Users. The structure of `users` block is documented below.
	Users ProfileUserArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileArgs)(nil)).Elem()
}

type ProfileInput interface {
	pulumi.Input

	ToProfileOutput() ProfileOutput
	ToProfileOutputWithContext(ctx context.Context) ProfileOutput
}

func (*Profile) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (i *Profile) ToProfileOutput() ProfileOutput {
	return i.ToProfileOutputWithContext(context.Background())
}

func (i *Profile) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileOutput)
}

// ProfileArrayInput is an input type that accepts ProfileArray and ProfileArrayOutput values.
// You can construct a concrete instance of `ProfileArrayInput` via:
//
//	ProfileArray{ ProfileArgs{...} }
type ProfileArrayInput interface {
	pulumi.Input

	ToProfileArrayOutput() ProfileArrayOutput
	ToProfileArrayOutputWithContext(context.Context) ProfileArrayOutput
}

type ProfileArray []ProfileInput

func (ProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Profile)(nil)).Elem()
}

func (i ProfileArray) ToProfileArrayOutput() ProfileArrayOutput {
	return i.ToProfileArrayOutputWithContext(context.Background())
}

func (i ProfileArray) ToProfileArrayOutputWithContext(ctx context.Context) ProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileArrayOutput)
}

// ProfileMapInput is an input type that accepts ProfileMap and ProfileMapOutput values.
// You can construct a concrete instance of `ProfileMapInput` via:
//
//	ProfileMap{ "key": ProfileArgs{...} }
type ProfileMapInput interface {
	pulumi.Input

	ToProfileMapOutput() ProfileMapOutput
	ToProfileMapOutputWithContext(context.Context) ProfileMapOutput
}

type ProfileMap map[string]ProfileInput

func (ProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Profile)(nil)).Elem()
}

func (i ProfileMap) ToProfileMapOutput() ProfileMapOutput {
	return i.ToProfileMapOutputWithContext(context.Background())
}

func (i ProfileMap) ToProfileMapOutputWithContext(ctx context.Context) ProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileMapOutput)
}

type ProfileOutput struct{ *pulumi.OutputState }

func (ProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (o ProfileOutput) ToProfileOutput() ProfileOutput {
	return o
}

func (o ProfileOutput) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return o
}

// Description.
func (o ProfileOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Device groups. The structure of `deviceGroups` block is documented below.
func (o ProfileOutput) DeviceGroups() ProfileDeviceGroupArrayOutput {
	return o.ApplyT(func(v *Profile) ProfileDeviceGroupArrayOutput { return v.DeviceGroups }).(ProfileDeviceGroupArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ProfileOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// FortiClient settings for Android platform. The structure of `forticlientAndroidSettings` block is documented below.
func (o ProfileOutput) ForticlientAndroidSettings() ProfileForticlientAndroidSettingsOutput {
	return o.ApplyT(func(v *Profile) ProfileForticlientAndroidSettingsOutput { return v.ForticlientAndroidSettings }).(ProfileForticlientAndroidSettingsOutput)
}

// FortiClient settings for iOS platform. The structure of `forticlientIosSettings` block is documented below.
func (o ProfileOutput) ForticlientIosSettings() ProfileForticlientIosSettingsOutput {
	return o.ApplyT(func(v *Profile) ProfileForticlientIosSettingsOutput { return v.ForticlientIosSettings }).(ProfileForticlientIosSettingsOutput)
}

// FortiClient settings for Windows/Mac platform. The structure of `forticlientWinmacSettings` block is documented below.
func (o ProfileOutput) ForticlientWinmacSettings() ProfileForticlientWinmacSettingsOutput {
	return o.ApplyT(func(v *Profile) ProfileForticlientWinmacSettingsOutput { return v.ForticlientWinmacSettings }).(ProfileForticlientWinmacSettingsOutput)
}

// Addresses for on-net detection. The structure of `onNetAddr` block is documented below.
func (o ProfileOutput) OnNetAddrs() ProfileOnNetAddrArrayOutput {
	return o.ApplyT(func(v *Profile) ProfileOnNetAddrArrayOutput { return v.OnNetAddrs }).(ProfileOnNetAddrArrayOutput)
}

// Profile name.
func (o ProfileOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.ProfileName }).(pulumi.StringOutput)
}

// Select an endpoint control replacement message override group from available options.
func (o ProfileOutput) ReplacemsgOverrideGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.ReplacemsgOverrideGroup }).(pulumi.StringOutput)
}

// Source addresses. The structure of `srcAddr` block is documented below.
func (o ProfileOutput) SrcAddrs() ProfileSrcAddrArrayOutput {
	return o.ApplyT(func(v *Profile) ProfileSrcAddrArrayOutput { return v.SrcAddrs }).(ProfileSrcAddrArrayOutput)
}

// User groups. The structure of `userGroups` block is documented below.
func (o ProfileOutput) UserGroups() ProfileUserGroupArrayOutput {
	return o.ApplyT(func(v *Profile) ProfileUserGroupArrayOutput { return v.UserGroups }).(ProfileUserGroupArrayOutput)
}

// Users. The structure of `users` block is documented below.
func (o ProfileOutput) Users() ProfileUserArrayOutput {
	return o.ApplyT(func(v *Profile) ProfileUserArrayOutput { return v.Users }).(ProfileUserArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ProfileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ProfileArrayOutput struct{ *pulumi.OutputState }

func (ProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Profile)(nil)).Elem()
}

func (o ProfileArrayOutput) ToProfileArrayOutput() ProfileArrayOutput {
	return o
}

func (o ProfileArrayOutput) ToProfileArrayOutputWithContext(ctx context.Context) ProfileArrayOutput {
	return o
}

func (o ProfileArrayOutput) Index(i pulumi.IntInput) ProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Profile {
		return vs[0].([]*Profile)[vs[1].(int)]
	}).(ProfileOutput)
}

type ProfileMapOutput struct{ *pulumi.OutputState }

func (ProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Profile)(nil)).Elem()
}

func (o ProfileMapOutput) ToProfileMapOutput() ProfileMapOutput {
	return o
}

func (o ProfileMapOutput) ToProfileMapOutputWithContext(ctx context.Context) ProfileMapOutput {
	return o
}

func (o ProfileMapOutput) MapIndex(k pulumi.StringInput) ProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Profile {
		return vs[0].(map[string]*Profile)[vs[1].(string)]
	}).(ProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileInput)(nil)).Elem(), &Profile{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileArrayInput)(nil)).Elem(), ProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileMapInput)(nil)).Elem(), ProfileMap{})
	pulumi.RegisterOutputType(ProfileOutput{})
	pulumi.RegisterOutputType(ProfileArrayOutput{})
	pulumi.RegisterOutputType(ProfileMapOutput{})
}