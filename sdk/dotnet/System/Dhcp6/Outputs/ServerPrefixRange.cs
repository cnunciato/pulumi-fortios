// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Dhcp6.Outputs
{

    [OutputType]
    public sealed class ServerPrefixRange
    {
        /// <summary>
        /// End of prefix range.
        /// </summary>
        public readonly string? EndPrefix;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Prefix length.
        /// </summary>
        public readonly int? PrefixLength;
        /// <summary>
        /// Start of prefix range.
        /// </summary>
        public readonly string? StartPrefix;

        [OutputConstructor]
        private ServerPrefixRange(
            string? endPrefix,

            int? id,

            int? prefixLength,

            string? startPrefix)
        {
            EndPrefix = endPrefix;
            Id = id;
            PrefixLength = prefixLength;
            StartPrefix = startPrefix;
        }
    }
}