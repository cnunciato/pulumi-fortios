// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Web.Inputs
{

    public sealed class UrlfilterEntryGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to take for URL filter matches. Valid values: `exempt`, `block`, `allow`, `monitor`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Action to take for AntiPhishing matches. Valid values: `block`, `log`.
        /// </summary>
        [Input("antiphishAction")]
        public Input<string>? AntiphishAction { get; set; }

        /// <summary>
        /// Resolve IPv4 address, IPv6 address, or both from DNS server. Valid values: `ipv4`, `ipv6`, `both`.
        /// </summary>
        [Input("dnsAddressFamily")]
        public Input<string>? DnsAddressFamily { get; set; }

        /// <summary>
        /// If action is set to exempt, select the security profile operations that exempt URLs skip. Separate multiple options with a space.
        /// </summary>
        [Input("exempt")]
        public Input<string>? Exempt { get; set; }

        /// <summary>
        /// Id.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Referrer host name.
        /// </summary>
        [Input("referrerHost")]
        public Input<string>? ReferrerHost { get; set; }

        /// <summary>
        /// Enable/disable this URL filter. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Filter type (simple, regex, or wildcard). Valid values: `simple`, `regex`, `wildcard`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// URL to be filtered.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Web proxy profile.
        /// </summary>
        [Input("webProxyProfile")]
        public Input<string>? WebProxyProfile { get; set; }

        public UrlfilterEntryGetArgs()
        {
        }
        public static new UrlfilterEntryGetArgs Empty => new UrlfilterEntryGetArgs();
    }
}
