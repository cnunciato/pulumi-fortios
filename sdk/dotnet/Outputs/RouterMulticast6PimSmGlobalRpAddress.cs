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
    public sealed class RouterMulticast6PimSmGlobalRpAddress
    {
        /// <summary>
        /// ID of the entry.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// RP router IPv6 address.
        /// </summary>
        public readonly string? Ip6Address;

        [OutputConstructor]
        private RouterMulticast6PimSmGlobalRpAddress(
            int? id,

            string? ip6Address)
        {
            Id = id;
            Ip6Address = ip6Address;
        }
    }
}
