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
    public sealed class DnsfilterProfileFtgdDnsFilter
    {
        /// <summary>
        /// Action to take for DNS requests matching the category. Valid values: `block`, `monitor`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Category number.
        /// </summary>
        public readonly int? Category;
        /// <summary>
        /// ID number.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Enable/disable DNS filter logging for this DNS profile. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Log;

        [OutputConstructor]
        private DnsfilterProfileFtgdDnsFilter(
            string? action,

            int? category,

            int? id,

            string? log)
        {
            Action = action;
            Category = category;
            Id = id;
            Log = log;
        }
    }
}
