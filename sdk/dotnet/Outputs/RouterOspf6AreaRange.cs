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
    public sealed class RouterOspf6AreaRange
    {
        /// <summary>
        /// Enable/disable advertise status. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Advertise;
        /// <summary>
        /// Range entry ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// IPv6 prefix.
        /// </summary>
        public readonly string? Prefix6;

        [OutputConstructor]
        private RouterOspf6AreaRange(
            string? advertise,

            int? id,

            string? prefix6)
        {
            Advertise = advertise;
            Id = id;
            Prefix6 = prefix6;
        }
    }
}
