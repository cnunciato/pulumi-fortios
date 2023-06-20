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
    public sealed class SwitchcontrollerQuarantineTarget
    {
        /// <summary>
        /// Description for the quarantine MAC.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// FSW entry id for the quarantine MAC.
        /// </summary>
        public readonly int? EntryId;
        /// <summary>
        /// Quarantine MAC.
        /// </summary>
        public readonly string? Mac;
        /// <summary>
        /// Tags for the quarantine MAC. The structure of `tag` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.SwitchcontrollerQuarantineTargetTag> Tags;

        [OutputConstructor]
        private SwitchcontrollerQuarantineTarget(
            string? description,

            int? entryId,

            string? mac,

            ImmutableArray<Outputs.SwitchcontrollerQuarantineTargetTag> tags)
        {
            Description = description;
            EntryId = entryId;
            Mac = mac;
            Tags = tags;
        }
    }
}
