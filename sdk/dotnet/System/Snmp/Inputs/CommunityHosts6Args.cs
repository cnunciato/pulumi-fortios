// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Snmp.Inputs
{

    public sealed class CommunityHosts6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("haDirect")]
        public Input<string>? HaDirect { get; set; }

        /// <summary>
        /// Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both. Valid values: `any`, `query`, `trap`.
        /// </summary>
        [Input("hostType")]
        public Input<string>? HostType { get; set; }

        /// <summary>
        /// Host6 entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// SNMP manager IPv6 address prefix.
        /// </summary>
        [Input("ipv6")]
        public Input<string>? Ipv6 { get; set; }

        /// <summary>
        /// Source IPv6 address for SNMP traps.
        /// </summary>
        [Input("sourceIpv6")]
        public Input<string>? SourceIpv6 { get; set; }

        public CommunityHosts6Args()
        {
        }
        public static new CommunityHosts6Args Empty => new CommunityHosts6Args();
    }
}
