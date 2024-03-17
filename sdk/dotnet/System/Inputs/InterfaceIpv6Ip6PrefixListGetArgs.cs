// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Inputs
{

    public sealed class InterfaceIpv6Ip6PrefixListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable the autonomous flag. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("autonomousFlag")]
        public Input<string>? AutonomousFlag { get; set; }

        [Input("dnssls")]
        private InputList<Inputs.InterfaceIpv6Ip6PrefixListDnsslGetArgs>? _dnssls;

        /// <summary>
        /// DNS search list option. The structure of `dnssl` block is documented below.
        /// </summary>
        public InputList<Inputs.InterfaceIpv6Ip6PrefixListDnsslGetArgs> Dnssls
        {
            get => _dnssls ?? (_dnssls = new InputList<Inputs.InterfaceIpv6Ip6PrefixListDnsslGetArgs>());
            set => _dnssls = value;
        }

        /// <summary>
        /// Enable/disable the onlink flag. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("onlinkFlag")]
        public Input<string>? OnlinkFlag { get; set; }

        /// <summary>
        /// Preferred life time (sec).
        /// </summary>
        [Input("preferredLifeTime")]
        public Input<int>? PreferredLifeTime { get; set; }

        /// <summary>
        /// IPv6 prefix.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// Recursive DNS server option.
        /// 
        /// The `dhcp6_iapd_list` block supports:
        /// </summary>
        [Input("rdnss")]
        public Input<string>? Rdnss { get; set; }

        /// <summary>
        /// Valid life time (sec).
        /// </summary>
        [Input("validLifeTime")]
        public Input<int>? ValidLifeTime { get; set; }

        public InterfaceIpv6Ip6PrefixListGetArgs()
        {
        }
        public static new InterfaceIpv6Ip6PrefixListGetArgs Empty => new InterfaceIpv6Ip6PrefixListGetArgs();
    }
}