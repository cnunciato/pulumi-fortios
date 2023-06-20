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
    public sealed class ReportChartColumn
    {
        /// <summary>
        /// Detail unit of column.
        /// </summary>
        public readonly string? DetailUnit;
        /// <summary>
        /// Detail value of column.
        /// </summary>
        public readonly string? DetailValue;
        /// <summary>
        /// Footer unit of column.
        /// </summary>
        public readonly string? FooterUnit;
        /// <summary>
        /// Footer value of column.
        /// </summary>
        public readonly string? FooterValue;
        /// <summary>
        /// Display name of table header.
        /// </summary>
        public readonly string? HeaderValue;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Show detail in certain display value for certain condition. The structure of `mapping` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ReportChartColumnMapping> Mappings;

        [OutputConstructor]
        private ReportChartColumn(
            string? detailUnit,

            string? detailValue,

            string? footerUnit,

            string? footerValue,

            string? headerValue,

            int? id,

            ImmutableArray<Outputs.ReportChartColumnMapping> mappings)
        {
            DetailUnit = detailUnit;
            DetailValue = detailValue;
            FooterUnit = footerUnit;
            FooterValue = footerValue;
            HeaderValue = headerValue;
            Id = id;
            Mappings = mappings;
        }
    }
}
