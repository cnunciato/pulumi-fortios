// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class ReportChartXSeriesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// X-series caption.
        /// </summary>
        [Input("caption")]
        public Input<string>? Caption { get; set; }

        /// <summary>
        /// X-series caption font size.
        /// </summary>
        [Input("captionFontSize")]
        public Input<int>? CaptionFontSize { get; set; }

        /// <summary>
        /// X-series value expression.
        /// </summary>
        [Input("databind")]
        public Input<string>? Databind { get; set; }

        /// <summary>
        /// X-series label font size.
        /// </summary>
        [Input("fontSize")]
        public Input<int>? FontSize { get; set; }

        /// <summary>
        /// X-series represent category or not. Valid values: `yes`, `no`.
        /// </summary>
        [Input("isCategory")]
        public Input<string>? IsCategory { get; set; }

        /// <summary>
        /// X-series label angle. Valid values: `45-degree`, `vertical`, `horizontal`.
        /// </summary>
        [Input("labelAngle")]
        public Input<string>? LabelAngle { get; set; }

        /// <summary>
        /// Scale increase or decrease. Valid values: `decrease`, `increase`.
        /// </summary>
        [Input("scaleDirection")]
        public Input<string>? ScaleDirection { get; set; }

        /// <summary>
        /// Date/time format. Valid values: `YYYY-MM-DD-HH-MM`, `YYYY-MM-DD HH`, `YYYY-MM-DD`, `YYYY-MM`, `YYYY`, `HH-MM`, `MM-DD`.
        /// </summary>
        [Input("scaleFormat")]
        public Input<string>? ScaleFormat { get; set; }

        /// <summary>
        /// Scale step.
        /// </summary>
        [Input("scaleStep")]
        public Input<int>? ScaleStep { get; set; }

        /// <summary>
        /// Scale unit. Valid values: `minute`, `hour`, `day`, `month`, `year`.
        /// </summary>
        [Input("scaleUnit")]
        public Input<string>? ScaleUnit { get; set; }

        /// <summary>
        /// X-series unit.
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        public ReportChartXSeriesArgs()
        {
        }
        public static new ReportChartXSeriesArgs Empty => new ReportChartXSeriesArgs();
    }
}
