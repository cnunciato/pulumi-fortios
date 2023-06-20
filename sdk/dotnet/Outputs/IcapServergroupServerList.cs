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
    public sealed class IcapServergroupServerList
    {
        /// <summary>
        /// ICAP server name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Optionally assign a weight of the ICAP server for weighted load balancing (1 - 100, default = 10)
        /// </summary>
        public readonly int? Weight;

        [OutputConstructor]
        private IcapServergroupServerList(
            string? name,

            int? weight)
        {
            Name = name;
            Weight = weight;
        }
    }
}
