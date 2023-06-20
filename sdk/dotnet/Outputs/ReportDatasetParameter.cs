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
    public sealed class ReportDatasetParameter
    {
        /// <summary>
        /// Data type. Valid values: `text`, `integer`, `double`, `long-integer`, `date-time`.
        /// </summary>
        public readonly string? DataType;
        /// <summary>
        /// Display name.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// SQL field name.
        /// </summary>
        public readonly string? Field;
        /// <summary>
        /// Parameter ID (1 to number of columns in SQL result).
        /// </summary>
        public readonly int? Id;

        [OutputConstructor]
        private ReportDatasetParameter(
            string? dataType,

            string? displayName,

            string? field,

            int? id)
        {
            DataType = dataType;
            DisplayName = displayName;
            Field = field;
            Id = id;
        }
    }
}
