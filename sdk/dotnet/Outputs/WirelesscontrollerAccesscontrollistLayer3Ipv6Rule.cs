// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class WirelesscontrollerAccesscontrollistLayer3Ipv6Rule
    {
        /// <summary>
        /// Policy action (allow | deny). Valid values: `allow`, `deny`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Description.
        /// </summary>
        public readonly string? Comment;
        /// <summary>
        /// Destination IPv6 address (any | local-LAN | IPv6 address[/prefix length]), default = any.
        /// </summary>
        public readonly string? Dstaddr;
        /// <summary>
        /// Destination port (0 - 65535, default = 0, meaning any).
        /// </summary>
        public readonly int? Dstport;
        /// <summary>
        /// Protocol type as defined by IANA (0 - 255, default = 255, meaning any).
        /// </summary>
        public readonly int? Protocol;
        /// <summary>
        /// Rule ID (1 - 65535).
        /// </summary>
        public readonly int? RuleId;
        /// <summary>
        /// Source IPv6 address (any | local-LAN | IPv6 address[/prefix length]), default = any.
        /// </summary>
        public readonly string? Srcaddr;
        /// <summary>
        /// Source port (0 - 65535, default = 0, meaning any).
        /// </summary>
        public readonly int? Srcport;

        [OutputConstructor]
        private WirelesscontrollerAccesscontrollistLayer3Ipv6Rule(
            string? action,

            string? comment,

            string? dstaddr,

            int? dstport,

            int? protocol,

            int? ruleId,

            string? srcaddr,

            int? srcport)
        {
            Action = action;
            Comment = comment;
            Dstaddr = dstaddr;
            Dstport = dstport;
            Protocol = protocol;
            RuleId = ruleId;
            Srcaddr = srcaddr;
            Srcport = srcport;
        }
    }
}