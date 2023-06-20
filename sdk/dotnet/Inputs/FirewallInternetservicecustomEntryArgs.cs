// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class FirewallInternetservicecustomEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`.
        /// </summary>
        [Input("addrMode")]
        public Input<string>? AddrMode { get; set; }

        [Input("dst6s")]
        private InputList<Inputs.FirewallInternetservicecustomEntryDst6Args>? _dst6s;

        /// <summary>
        /// Destination address6 or address6 group name. The structure of `dst6` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallInternetservicecustomEntryDst6Args> Dst6s
        {
            get => _dst6s ?? (_dst6s = new InputList<Inputs.FirewallInternetservicecustomEntryDst6Args>());
            set => _dst6s = value;
        }

        [Input("dsts")]
        private InputList<Inputs.FirewallInternetservicecustomEntryDstArgs>? _dsts;

        /// <summary>
        /// Destination address or address group name. The structure of `dst` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallInternetservicecustomEntryDstArgs> Dsts
        {
            get => _dsts ?? (_dsts = new InputList<Inputs.FirewallInternetservicecustomEntryDstArgs>());
            set => _dsts = value;
        }

        /// <summary>
        /// Entry ID(1-255).
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("portRanges")]
        private InputList<Inputs.FirewallInternetservicecustomEntryPortRangeArgs>? _portRanges;

        /// <summary>
        /// Port ranges in the custom entry. The structure of `port_range` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallInternetservicecustomEntryPortRangeArgs> PortRanges
        {
            get => _portRanges ?? (_portRanges = new InputList<Inputs.FirewallInternetservicecustomEntryPortRangeArgs>());
            set => _portRanges = value;
        }

        /// <summary>
        /// Integer value for the protocol type as defined by IANA (0 - 255).
        /// </summary>
        [Input("protocol")]
        public Input<int>? Protocol { get; set; }

        public FirewallInternetservicecustomEntryArgs()
        {
        }
        public static new FirewallInternetservicecustomEntryArgs Empty => new FirewallInternetservicecustomEntryArgs();
    }
}
