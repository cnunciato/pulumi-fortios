// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall proxyaddress
func LookupFirewallProxyaddress(ctx *pulumi.Context, args *LookupFirewallProxyaddressArgs, opts ...pulumi.InvokeOption) (*LookupFirewallProxyaddressResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallProxyaddressResult
	err := ctx.Invoke("fortios:index/getFirewallProxyaddress:getFirewallProxyaddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallProxyaddress.
type LookupFirewallProxyaddressArgs struct {
	// Specify the name of the desired firewall proxyaddress.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getFirewallProxyaddress.
type LookupFirewallProxyaddressResult struct {
	// SaaS application. The structure of `application` block is documented below.
	Applications []GetFirewallProxyaddressApplication `pulumi:"applications"`
	// Case sensitivity in pattern.
	CaseSensitivity string `pulumi:"caseSensitivity"`
	// Tag category.
	Categories []GetFirewallProxyaddressCategory `pulumi:"categories"`
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color int `pulumi:"color"`
	// Optional comments.
	Comment string `pulumi:"comment"`
	// HTTP header regular expression.
	Header string `pulumi:"header"`
	// HTTP header group. The structure of `headerGroup` block is documented below.
	HeaderGroups []GetFirewallProxyaddressHeaderGroup `pulumi:"headerGroups"`
	// HTTP header.
	HeaderName string `pulumi:"headerName"`
	// Address object for the host.
	Host string `pulumi:"host"`
	// Host name as a regular expression.
	HostRegex string `pulumi:"hostRegex"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// HTTP request methods to be used.
	Method string `pulumi:"method"`
	// SaaS applicaton name.
	Name string `pulumi:"name"`
	// URL path as a regular expression.
	Path string `pulumi:"path"`
	// Match the query part of the URL as a regular expression.
	Query string `pulumi:"query"`
	// Enable/disable use of referrer field in the HTTP header to match the address.
	Referrer string `pulumi:"referrer"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []GetFirewallProxyaddressTagging `pulumi:"taggings"`
	// Proxy address type.
	Type string `pulumi:"type"`
	// Names of browsers to be used as user agent.
	Ua string `pulumi:"ua"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid      string  `pulumi:"uuid"`
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable visibility of the object in the GUI.
	Visibility string `pulumi:"visibility"`
}

func LookupFirewallProxyaddressOutput(ctx *pulumi.Context, args LookupFirewallProxyaddressOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallProxyaddressResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallProxyaddressResult, error) {
			args := v.(LookupFirewallProxyaddressArgs)
			r, err := LookupFirewallProxyaddress(ctx, &args, opts...)
			var s LookupFirewallProxyaddressResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallProxyaddressResultOutput)
}

// A collection of arguments for invoking getFirewallProxyaddress.
type LookupFirewallProxyaddressOutputArgs struct {
	// Specify the name of the desired firewall proxyaddress.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallProxyaddressOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallProxyaddressArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallProxyaddress.
type LookupFirewallProxyaddressResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallProxyaddressResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallProxyaddressResult)(nil)).Elem()
}

func (o LookupFirewallProxyaddressResultOutput) ToLookupFirewallProxyaddressResultOutput() LookupFirewallProxyaddressResultOutput {
	return o
}

func (o LookupFirewallProxyaddressResultOutput) ToLookupFirewallProxyaddressResultOutputWithContext(ctx context.Context) LookupFirewallProxyaddressResultOutput {
	return o
}

// SaaS application. The structure of `application` block is documented below.
func (o LookupFirewallProxyaddressResultOutput) Applications() GetFirewallProxyaddressApplicationArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) []GetFirewallProxyaddressApplication { return v.Applications }).(GetFirewallProxyaddressApplicationArrayOutput)
}

// Case sensitivity in pattern.
func (o LookupFirewallProxyaddressResultOutput) CaseSensitivity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) string { return v.CaseSensitivity }).(pulumi.StringOutput)
}

// Tag category.
func (o LookupFirewallProxyaddressResultOutput) Categories() GetFirewallProxyaddressCategoryArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) []GetFirewallProxyaddressCategory { return v.Categories }).(GetFirewallProxyaddressCategoryArrayOutput)
}

// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
func (o LookupFirewallProxyaddressResultOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) int { return v.Color }).(pulumi.IntOutput)
}

// Optional comments.
func (o LookupFirewallProxyaddressResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) string { return v.Comment }).(pulumi.StringOutput)
}

// HTTP header regular expression.
func (o LookupFirewallProxyaddressResultOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) string { return v.Header }).(pulumi.StringOutput)
}

// HTTP header group. The structure of `headerGroup` block is documented below.
func (o LookupFirewallProxyaddressResultOutput) HeaderGroups() GetFirewallProxyaddressHeaderGroupArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) []GetFirewallProxyaddressHeaderGroup { return v.HeaderGroups }).(GetFirewallProxyaddressHeaderGroupArrayOutput)
}

// HTTP header.
func (o LookupFirewallProxyaddressResultOutput) HeaderName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) string { return v.HeaderName }).(pulumi.StringOutput)
}

// Address object for the host.
func (o LookupFirewallProxyaddressResultOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) string { return v.Host }).(pulumi.StringOutput)
}

// Host name as a regular expression.
func (o LookupFirewallProxyaddressResultOutput) HostRegex() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) string { return v.HostRegex }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallProxyaddressResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) string { return v.Id }).(pulumi.StringOutput)
}

// HTTP request methods to be used.
func (o LookupFirewallProxyaddressResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) string { return v.Method }).(pulumi.StringOutput)
}

// SaaS applicaton name.
func (o LookupFirewallProxyaddressResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) string { return v.Name }).(pulumi.StringOutput)
}

// URL path as a regular expression.
func (o LookupFirewallProxyaddressResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) string { return v.Path }).(pulumi.StringOutput)
}

// Match the query part of the URL as a regular expression.
func (o LookupFirewallProxyaddressResultOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) string { return v.Query }).(pulumi.StringOutput)
}

// Enable/disable use of referrer field in the HTTP header to match the address.
func (o LookupFirewallProxyaddressResultOutput) Referrer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) string { return v.Referrer }).(pulumi.StringOutput)
}

// Config object tagging. The structure of `tagging` block is documented below.
func (o LookupFirewallProxyaddressResultOutput) Taggings() GetFirewallProxyaddressTaggingArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) []GetFirewallProxyaddressTagging { return v.Taggings }).(GetFirewallProxyaddressTaggingArrayOutput)
}

// Proxy address type.
func (o LookupFirewallProxyaddressResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) string { return v.Type }).(pulumi.StringOutput)
}

// Names of browsers to be used as user agent.
func (o LookupFirewallProxyaddressResultOutput) Ua() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) string { return v.Ua }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o LookupFirewallProxyaddressResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupFirewallProxyaddressResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable visibility of the object in the GUI.
func (o LookupFirewallProxyaddressResultOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyaddressResult) string { return v.Visibility }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallProxyaddressResultOutput{})
}
