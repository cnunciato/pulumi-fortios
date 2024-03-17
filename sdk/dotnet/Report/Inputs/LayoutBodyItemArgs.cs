// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Report.Inputs
{

    public sealed class LayoutBodyItemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Report item chart name.
        /// </summary>
        [Input("chart")]
        public Input<string>? Chart { get; set; }

        /// <summary>
        /// Report chart options. Valid values: `include-no-data`, `hide-title`, `show-caption`.
        /// </summary>
        [Input("chartOptions")]
        public Input<string>? ChartOptions { get; set; }

        /// <summary>
        /// Report section column number.
        /// </summary>
        [Input("column")]
        public Input<int>? Column { get; set; }

        /// <summary>
        /// Report item text content.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Control how drill down charts are shown.
        /// </summary>
        [Input("drillDownItems")]
        public Input<string>? DrillDownItems { get; set; }

        /// <summary>
        /// Control whether keys from the parent being combined or not.
        /// </summary>
        [Input("drillDownTypes")]
        public Input<string>? DrillDownTypes { get; set; }

        /// <summary>
        /// Enable/disable hide item in report. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("hide")]
        public Input<string>? Hide { get; set; }

        /// <summary>
        /// Report item ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Report item image file name.
        /// </summary>
        [Input("imgSrc")]
        public Input<string>? ImgSrc { get; set; }

        /// <summary>
        /// Report item list component. Valid values: `bullet`, `numbered`.
        /// </summary>
        [Input("listComponent")]
        public Input<string>? ListComponent { get; set; }

        [Input("lists")]
        private InputList<Inputs.LayoutBodyItemListArgs>? _lists;

        /// <summary>
        /// Configure report list item. The structure of `list` block is documented below.
        /// </summary>
        public InputList<Inputs.LayoutBodyItemListArgs> Lists
        {
            get => _lists ?? (_lists = new InputList<Inputs.LayoutBodyItemListArgs>());
            set => _lists = value;
        }

        /// <summary>
        /// Report item miscellaneous component. Valid values: `hline`, `page-break`, `column-break`, `section-start`.
        /// </summary>
        [Input("miscComponent")]
        public Input<string>? MiscComponent { get; set; }

        [Input("parameters")]
        private InputList<Inputs.LayoutBodyItemParameterArgs>? _parameters;

        /// <summary>
        /// Parameters. The structure of `parameters` block is documented below.
        /// </summary>
        public InputList<Inputs.LayoutBodyItemParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.LayoutBodyItemParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// Report item style.
        /// </summary>
        [Input("style")]
        public Input<string>? Style { get; set; }

        /// <summary>
        /// Table chart caption style.
        /// </summary>
        [Input("tableCaptionStyle")]
        public Input<string>? TableCaptionStyle { get; set; }

        /// <summary>
        /// Report item table column widths.
        /// </summary>
        [Input("tableColumnWidths")]
        public Input<string>? TableColumnWidths { get; set; }

        /// <summary>
        /// Table chart even row style.
        /// </summary>
        [Input("tableEvenRowStyle")]
        public Input<string>? TableEvenRowStyle { get; set; }

        /// <summary>
        /// Table chart head style.
        /// </summary>
        [Input("tableHeadStyle")]
        public Input<string>? TableHeadStyle { get; set; }

        /// <summary>
        /// Table chart odd row style.
        /// </summary>
        [Input("tableOddRowStyle")]
        public Input<string>? TableOddRowStyle { get; set; }

        /// <summary>
        /// Report item text component. Valid values: `text`, `heading1`, `heading2`, `heading3`.
        /// </summary>
        [Input("textComponent")]
        public Input<string>? TextComponent { get; set; }

        /// <summary>
        /// Report section title.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// Value of top.
        /// </summary>
        [Input("topN")]
        public Input<int>? TopN { get; set; }

        /// <summary>
        /// Report item type. Valid values: `text`, `image`, `chart`, `misc`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public LayoutBodyItemArgs()
        {
        }
        public static new LayoutBodyItemArgs Empty => new LayoutBodyItemArgs();
    }
}