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
    public sealed class CifsProfileFileFilter
    {
        /// <summary>
        /// File filter entries. The structure of `entries` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.CifsProfileFileFilterEntry> Entries;
        /// <summary>
        /// Enable/disable file filter logging. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Log;
        /// <summary>
        /// Enable/disable file filter. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private CifsProfileFileFilter(
            ImmutableArray<Outputs.CifsProfileFileFilterEntry> entries,

            string? log,

            string? status)
        {
            Entries = entries;
            Log = log;
            Status = status;
        }
    }
}
