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
    public sealed class FirewallInternetservicecustomEntry
    {
        /// <summary>
        /// Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`.
        /// </summary>
        public readonly string? AddrMode;
        /// <summary>
        /// Destination address6 or address6 group name. The structure of `dst6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.FirewallInternetservicecustomEntryDst6> Dst6s;
        /// <summary>
        /// Destination address or address group name. The structure of `dst` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.FirewallInternetservicecustomEntryDst> Dsts;
        /// <summary>
        /// Entry ID(1-255).
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Port ranges in the custom entry. The structure of `port_range` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.FirewallInternetservicecustomEntryPortRange> PortRanges;
        /// <summary>
        /// Integer value for the protocol type as defined by IANA (0 - 255).
        /// </summary>
        public readonly int? Protocol;

        [OutputConstructor]
        private FirewallInternetservicecustomEntry(
            string? addrMode,

            ImmutableArray<Outputs.FirewallInternetservicecustomEntryDst6> dst6s,

            ImmutableArray<Outputs.FirewallInternetservicecustomEntryDst> dsts,

            int? id,

            ImmutableArray<Outputs.FirewallInternetservicecustomEntryPortRange> portRanges,

            int? protocol)
        {
            AddrMode = addrMode;
            Dst6s = dst6s;
            Dsts = dsts;
            Id = id;
            PortRanges = portRanges;
            Protocol = protocol;
        }
    }
}
