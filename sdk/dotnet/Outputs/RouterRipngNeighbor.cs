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
    public sealed class RouterRipngNeighbor
    {
        /// <summary>
        /// Neighbor entry ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Interface name.
        /// </summary>
        public readonly string? Interface;
        /// <summary>
        /// IPv6 link-local address.
        /// </summary>
        public readonly string? Ip6;

        [OutputConstructor]
        private RouterRipngNeighbor(
            int? id,

            string? @interface,

            string? ip6)
        {
            Id = id;
            Interface = @interface;
            Ip6 = ip6;
        }
    }
}
