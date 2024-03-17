// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package user

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// SAML server entry configuration. Applies to FortiOS Version `>= 6.2.4`.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/user"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := user.NewSaml(ctx, "tr3", &user.SamlArgs{
//				Cert:               pulumi.String("Fortinet_Factory"),
//				EntityId:           pulumi.String("https://1.1.1.1"),
//				IdpCert:            pulumi.String("cer11"),
//				IdpEntityId:        pulumi.String("https://1.1.1.1/acc"),
//				IdpSingleLogoutUrl: pulumi.String("https://1.1.1.1/lo"),
//				IdpSingleSignOnUrl: pulumi.String("https://1.1.1.1/sou"),
//				SingleLogoutUrl:    pulumi.String("https://1.1.1.1/logout"),
//				SingleSignOnUrl:    pulumi.String("https://1.1.1.1/sign"),
//				UserName:           pulumi.String("ad111"),
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
// User Saml can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:user/saml:Saml labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:user/saml:Saml labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Saml struct {
	pulumi.CustomResourceState

	// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
	AdfsClaim pulumi.StringOutput `pulumi:"adfsClaim"`
	// URL to verify authentication.
	AuthUrl pulumi.StringPtrOutput `pulumi:"authUrl"`
	// Certificate to sign SAML messages.
	Cert pulumi.StringOutput `pulumi:"cert"`
	// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
	ClockTolerance pulumi.IntOutput `pulumi:"clockTolerance"`
	// Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
	DigestMethod pulumi.StringOutput `pulumi:"digestMethod"`
	// SP entity ID.
	EntityId pulumi.StringOutput `pulumi:"entityId"`
	// Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	GroupClaimType pulumi.StringOutput `pulumi:"groupClaimType"`
	// Group name in assertion statement.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// IDP Certificate name.
	IdpCert pulumi.StringOutput `pulumi:"idpCert"`
	// IDP entity ID.
	IdpEntityId pulumi.StringOutput `pulumi:"idpEntityId"`
	// IDP single logout url.
	IdpSingleLogoutUrl pulumi.StringOutput `pulumi:"idpSingleLogoutUrl"`
	// IDP single sign-on URL.
	IdpSingleSignOnUrl pulumi.StringOutput `pulumi:"idpSingleSignOnUrl"`
	// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
	LimitRelaystate pulumi.StringOutput `pulumi:"limitRelaystate"`
	// SAML server entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// SP single logout URL.
	SingleLogoutUrl pulumi.StringOutput `pulumi:"singleLogoutUrl"`
	// SP single sign-on URL.
	SingleSignOnUrl pulumi.StringOutput `pulumi:"singleSignOnUrl"`
	// User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	UserClaimType pulumi.StringOutput `pulumi:"userClaimType"`
	// User name in assertion statement.
	UserName pulumi.StringOutput `pulumi:"userName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSaml registers a new resource with the given unique name, arguments, and options.
func NewSaml(ctx *pulumi.Context,
	name string, args *SamlArgs, opts ...pulumi.ResourceOption) (*Saml, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EntityId == nil {
		return nil, errors.New("invalid value for required argument 'EntityId'")
	}
	if args.IdpCert == nil {
		return nil, errors.New("invalid value for required argument 'IdpCert'")
	}
	if args.IdpEntityId == nil {
		return nil, errors.New("invalid value for required argument 'IdpEntityId'")
	}
	if args.IdpSingleSignOnUrl == nil {
		return nil, errors.New("invalid value for required argument 'IdpSingleSignOnUrl'")
	}
	if args.SingleSignOnUrl == nil {
		return nil, errors.New("invalid value for required argument 'SingleSignOnUrl'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Saml
	err := ctx.RegisterResource("fortios:user/saml:Saml", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSaml gets an existing Saml resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSaml(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SamlState, opts ...pulumi.ResourceOption) (*Saml, error) {
	var resource Saml
	err := ctx.ReadResource("fortios:user/saml:Saml", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Saml resources.
type samlState struct {
	// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
	AdfsClaim *string `pulumi:"adfsClaim"`
	// URL to verify authentication.
	AuthUrl *string `pulumi:"authUrl"`
	// Certificate to sign SAML messages.
	Cert *string `pulumi:"cert"`
	// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
	ClockTolerance *int `pulumi:"clockTolerance"`
	// Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
	DigestMethod *string `pulumi:"digestMethod"`
	// SP entity ID.
	EntityId *string `pulumi:"entityId"`
	// Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	GroupClaimType *string `pulumi:"groupClaimType"`
	// Group name in assertion statement.
	GroupName *string `pulumi:"groupName"`
	// IDP Certificate name.
	IdpCert *string `pulumi:"idpCert"`
	// IDP entity ID.
	IdpEntityId *string `pulumi:"idpEntityId"`
	// IDP single logout url.
	IdpSingleLogoutUrl *string `pulumi:"idpSingleLogoutUrl"`
	// IDP single sign-on URL.
	IdpSingleSignOnUrl *string `pulumi:"idpSingleSignOnUrl"`
	// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
	LimitRelaystate *string `pulumi:"limitRelaystate"`
	// SAML server entry name.
	Name *string `pulumi:"name"`
	// SP single logout URL.
	SingleLogoutUrl *string `pulumi:"singleLogoutUrl"`
	// SP single sign-on URL.
	SingleSignOnUrl *string `pulumi:"singleSignOnUrl"`
	// User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	UserClaimType *string `pulumi:"userClaimType"`
	// User name in assertion statement.
	UserName *string `pulumi:"userName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SamlState struct {
	// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
	AdfsClaim pulumi.StringPtrInput
	// URL to verify authentication.
	AuthUrl pulumi.StringPtrInput
	// Certificate to sign SAML messages.
	Cert pulumi.StringPtrInput
	// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
	ClockTolerance pulumi.IntPtrInput
	// Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
	DigestMethod pulumi.StringPtrInput
	// SP entity ID.
	EntityId pulumi.StringPtrInput
	// Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	GroupClaimType pulumi.StringPtrInput
	// Group name in assertion statement.
	GroupName pulumi.StringPtrInput
	// IDP Certificate name.
	IdpCert pulumi.StringPtrInput
	// IDP entity ID.
	IdpEntityId pulumi.StringPtrInput
	// IDP single logout url.
	IdpSingleLogoutUrl pulumi.StringPtrInput
	// IDP single sign-on URL.
	IdpSingleSignOnUrl pulumi.StringPtrInput
	// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
	LimitRelaystate pulumi.StringPtrInput
	// SAML server entry name.
	Name pulumi.StringPtrInput
	// SP single logout URL.
	SingleLogoutUrl pulumi.StringPtrInput
	// SP single sign-on URL.
	SingleSignOnUrl pulumi.StringPtrInput
	// User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	UserClaimType pulumi.StringPtrInput
	// User name in assertion statement.
	UserName pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SamlState) ElementType() reflect.Type {
	return reflect.TypeOf((*samlState)(nil)).Elem()
}

type samlArgs struct {
	// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
	AdfsClaim *string `pulumi:"adfsClaim"`
	// URL to verify authentication.
	AuthUrl *string `pulumi:"authUrl"`
	// Certificate to sign SAML messages.
	Cert *string `pulumi:"cert"`
	// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
	ClockTolerance *int `pulumi:"clockTolerance"`
	// Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
	DigestMethod *string `pulumi:"digestMethod"`
	// SP entity ID.
	EntityId string `pulumi:"entityId"`
	// Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	GroupClaimType *string `pulumi:"groupClaimType"`
	// Group name in assertion statement.
	GroupName *string `pulumi:"groupName"`
	// IDP Certificate name.
	IdpCert string `pulumi:"idpCert"`
	// IDP entity ID.
	IdpEntityId string `pulumi:"idpEntityId"`
	// IDP single logout url.
	IdpSingleLogoutUrl *string `pulumi:"idpSingleLogoutUrl"`
	// IDP single sign-on URL.
	IdpSingleSignOnUrl string `pulumi:"idpSingleSignOnUrl"`
	// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
	LimitRelaystate *string `pulumi:"limitRelaystate"`
	// SAML server entry name.
	Name *string `pulumi:"name"`
	// SP single logout URL.
	SingleLogoutUrl *string `pulumi:"singleLogoutUrl"`
	// SP single sign-on URL.
	SingleSignOnUrl string `pulumi:"singleSignOnUrl"`
	// User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	UserClaimType *string `pulumi:"userClaimType"`
	// User name in assertion statement.
	UserName *string `pulumi:"userName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Saml resource.
type SamlArgs struct {
	// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
	AdfsClaim pulumi.StringPtrInput
	// URL to verify authentication.
	AuthUrl pulumi.StringPtrInput
	// Certificate to sign SAML messages.
	Cert pulumi.StringPtrInput
	// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
	ClockTolerance pulumi.IntPtrInput
	// Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
	DigestMethod pulumi.StringPtrInput
	// SP entity ID.
	EntityId pulumi.StringInput
	// Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	GroupClaimType pulumi.StringPtrInput
	// Group name in assertion statement.
	GroupName pulumi.StringPtrInput
	// IDP Certificate name.
	IdpCert pulumi.StringInput
	// IDP entity ID.
	IdpEntityId pulumi.StringInput
	// IDP single logout url.
	IdpSingleLogoutUrl pulumi.StringPtrInput
	// IDP single sign-on URL.
	IdpSingleSignOnUrl pulumi.StringInput
	// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
	LimitRelaystate pulumi.StringPtrInput
	// SAML server entry name.
	Name pulumi.StringPtrInput
	// SP single logout URL.
	SingleLogoutUrl pulumi.StringPtrInput
	// SP single sign-on URL.
	SingleSignOnUrl pulumi.StringInput
	// User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	UserClaimType pulumi.StringPtrInput
	// User name in assertion statement.
	UserName pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SamlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*samlArgs)(nil)).Elem()
}

type SamlInput interface {
	pulumi.Input

	ToSamlOutput() SamlOutput
	ToSamlOutputWithContext(ctx context.Context) SamlOutput
}

func (*Saml) ElementType() reflect.Type {
	return reflect.TypeOf((**Saml)(nil)).Elem()
}

func (i *Saml) ToSamlOutput() SamlOutput {
	return i.ToSamlOutputWithContext(context.Background())
}

func (i *Saml) ToSamlOutputWithContext(ctx context.Context) SamlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlOutput)
}

// SamlArrayInput is an input type that accepts SamlArray and SamlArrayOutput values.
// You can construct a concrete instance of `SamlArrayInput` via:
//
//	SamlArray{ SamlArgs{...} }
type SamlArrayInput interface {
	pulumi.Input

	ToSamlArrayOutput() SamlArrayOutput
	ToSamlArrayOutputWithContext(context.Context) SamlArrayOutput
}

type SamlArray []SamlInput

func (SamlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Saml)(nil)).Elem()
}

func (i SamlArray) ToSamlArrayOutput() SamlArrayOutput {
	return i.ToSamlArrayOutputWithContext(context.Background())
}

func (i SamlArray) ToSamlArrayOutputWithContext(ctx context.Context) SamlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlArrayOutput)
}

// SamlMapInput is an input type that accepts SamlMap and SamlMapOutput values.
// You can construct a concrete instance of `SamlMapInput` via:
//
//	SamlMap{ "key": SamlArgs{...} }
type SamlMapInput interface {
	pulumi.Input

	ToSamlMapOutput() SamlMapOutput
	ToSamlMapOutputWithContext(context.Context) SamlMapOutput
}

type SamlMap map[string]SamlInput

func (SamlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Saml)(nil)).Elem()
}

func (i SamlMap) ToSamlMapOutput() SamlMapOutput {
	return i.ToSamlMapOutputWithContext(context.Background())
}

func (i SamlMap) ToSamlMapOutputWithContext(ctx context.Context) SamlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlMapOutput)
}

type SamlOutput struct{ *pulumi.OutputState }

func (SamlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Saml)(nil)).Elem()
}

func (o SamlOutput) ToSamlOutput() SamlOutput {
	return o
}

func (o SamlOutput) ToSamlOutputWithContext(ctx context.Context) SamlOutput {
	return o
}

// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
func (o SamlOutput) AdfsClaim() pulumi.StringOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringOutput { return v.AdfsClaim }).(pulumi.StringOutput)
}

// URL to verify authentication.
func (o SamlOutput) AuthUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringPtrOutput { return v.AuthUrl }).(pulumi.StringPtrOutput)
}

// Certificate to sign SAML messages.
func (o SamlOutput) Cert() pulumi.StringOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringOutput { return v.Cert }).(pulumi.StringOutput)
}

// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
func (o SamlOutput) ClockTolerance() pulumi.IntOutput {
	return o.ApplyT(func(v *Saml) pulumi.IntOutput { return v.ClockTolerance }).(pulumi.IntOutput)
}

// Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
func (o SamlOutput) DigestMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringOutput { return v.DigestMethod }).(pulumi.StringOutput)
}

// SP entity ID.
func (o SamlOutput) EntityId() pulumi.StringOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringOutput { return v.EntityId }).(pulumi.StringOutput)
}

// Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
func (o SamlOutput) GroupClaimType() pulumi.StringOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringOutput { return v.GroupClaimType }).(pulumi.StringOutput)
}

// Group name in assertion statement.
func (o SamlOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// IDP Certificate name.
func (o SamlOutput) IdpCert() pulumi.StringOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringOutput { return v.IdpCert }).(pulumi.StringOutput)
}

// IDP entity ID.
func (o SamlOutput) IdpEntityId() pulumi.StringOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringOutput { return v.IdpEntityId }).(pulumi.StringOutput)
}

// IDP single logout url.
func (o SamlOutput) IdpSingleLogoutUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringOutput { return v.IdpSingleLogoutUrl }).(pulumi.StringOutput)
}

// IDP single sign-on URL.
func (o SamlOutput) IdpSingleSignOnUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringOutput { return v.IdpSingleSignOnUrl }).(pulumi.StringOutput)
}

// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
func (o SamlOutput) LimitRelaystate() pulumi.StringOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringOutput { return v.LimitRelaystate }).(pulumi.StringOutput)
}

// SAML server entry name.
func (o SamlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// SP single logout URL.
func (o SamlOutput) SingleLogoutUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringOutput { return v.SingleLogoutUrl }).(pulumi.StringOutput)
}

// SP single sign-on URL.
func (o SamlOutput) SingleSignOnUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringOutput { return v.SingleSignOnUrl }).(pulumi.StringOutput)
}

// User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
func (o SamlOutput) UserClaimType() pulumi.StringOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringOutput { return v.UserClaimType }).(pulumi.StringOutput)
}

// User name in assertion statement.
func (o SamlOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SamlOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Saml) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SamlArrayOutput struct{ *pulumi.OutputState }

func (SamlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Saml)(nil)).Elem()
}

func (o SamlArrayOutput) ToSamlArrayOutput() SamlArrayOutput {
	return o
}

func (o SamlArrayOutput) ToSamlArrayOutputWithContext(ctx context.Context) SamlArrayOutput {
	return o
}

func (o SamlArrayOutput) Index(i pulumi.IntInput) SamlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Saml {
		return vs[0].([]*Saml)[vs[1].(int)]
	}).(SamlOutput)
}

type SamlMapOutput struct{ *pulumi.OutputState }

func (SamlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Saml)(nil)).Elem()
}

func (o SamlMapOutput) ToSamlMapOutput() SamlMapOutput {
	return o
}

func (o SamlMapOutput) ToSamlMapOutputWithContext(ctx context.Context) SamlMapOutput {
	return o
}

func (o SamlMapOutput) MapIndex(k pulumi.StringInput) SamlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Saml {
		return vs[0].(map[string]*Saml)[vs[1].(string)]
	}).(SamlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SamlInput)(nil)).Elem(), &Saml{})
	pulumi.RegisterInputType(reflect.TypeOf((*SamlArrayInput)(nil)).Elem(), SamlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SamlMapInput)(nil)).Elem(), SamlMap{})
	pulumi.RegisterOutputType(SamlOutput{})
	pulumi.RegisterOutputType(SamlArrayOutput{})
	pulumi.RegisterOutputType(SamlMapOutput{})
}