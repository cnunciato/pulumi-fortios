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
    public sealed class ReportDatasetField
    {
        /// <summary>
        /// Display name.
        /// </summary>
        public readonly string? Displayname;
        /// <summary>
        /// Field ID (1 to number of columns in SQL result).
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Field type. Valid values: `text`, `integer`, `double`.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private ReportDatasetField(
            string? displayname,

            int? id,

            string? name,

            string? type)
        {
            Displayname = displayname;
            Id = id;
            Name = name;
            Type = type;
        }
    }
}