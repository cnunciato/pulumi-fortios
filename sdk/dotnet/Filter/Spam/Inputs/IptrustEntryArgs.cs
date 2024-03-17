// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Spam.Inputs
{

    public sealed class IptrustEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of address. Valid values: `ipv4`, `ipv6`.
        /// </summary>
        [Input("addrType")]
        public Input<string>? AddrType { get; set; }

        /// <summary>
        /// Trusted IP entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// IPv4 network address or network address/subnet mask bits.
        /// </summary>
        [Input("ip4Subnet")]
        public Input<string>? Ip4Subnet { get; set; }

        /// <summary>
        /// IPv6 network address/subnet mask bits.
        /// </summary>
        [Input("ip6Subnet")]
        public Input<string>? Ip6Subnet { get; set; }

        /// <summary>
        /// Enable/disable status. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public IptrustEntryArgs()
        {
        }
        public static new IptrustEntryArgs Empty => new IptrustEntryArgs();
    }
}