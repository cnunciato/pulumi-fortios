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

    public sealed class ReportChartYSeriesGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Y-series caption.
        /// </summary>
        [Input("caption")]
        public Input<string>? Caption { get; set; }

        /// <summary>
        /// Y-series caption font size.
        /// </summary>
        [Input("captionFontSize")]
        public Input<int>? CaptionFontSize { get; set; }

        /// <summary>
        /// Y-series value expression.
        /// </summary>
        [Input("databind")]
        public Input<string>? Databind { get; set; }

        /// <summary>
        /// Extra Y-series value.
        /// </summary>
        [Input("extraDatabind")]
        public Input<string>? ExtraDatabind { get; set; }

        /// <summary>
        /// Allow another Y-series value Valid values: `enable`, `disable`.
        /// </summary>
        [Input("extraY")]
        public Input<string>? ExtraY { get; set; }

        /// <summary>
        /// Extra Y-series legend type/name.
        /// </summary>
        [Input("extraYLegend")]
        public Input<string>? ExtraYLegend { get; set; }

        /// <summary>
        /// Y-series label font size.
        /// </summary>
        [Input("fontSize")]
        public Input<int>? FontSize { get; set; }

        /// <summary>
        /// Y-series group option.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// Y-series label angle. Valid values: `45-degree`, `vertical`, `horizontal`.
        /// </summary>
        [Input("labelAngle")]
        public Input<string>? LabelAngle { get; set; }

        /// <summary>
        /// Y-series unit.
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        /// <summary>
        /// First Y-series legend type/name.
        /// </summary>
        [Input("yLegend")]
        public Input<string>? YLegend { get; set; }

        public ReportChartYSeriesGetArgs()
        {
        }
        public static new ReportChartYSeriesGetArgs Empty => new ReportChartYSeriesGetArgs();
    }
}
