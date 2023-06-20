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

    public sealed class ReportLayoutPageHeaderArgs : global::Pulumi.ResourceArgs
    {
        [Input("headerItems")]
        private InputList<Inputs.ReportLayoutPageHeaderHeaderItemArgs>? _headerItems;

        /// <summary>
        /// Configure report header item. The structure of `header_item` block is documented below.
        /// </summary>
        public InputList<Inputs.ReportLayoutPageHeaderHeaderItemArgs> HeaderItems
        {
            get => _headerItems ?? (_headerItems = new InputList<Inputs.ReportLayoutPageHeaderHeaderItemArgs>());
            set => _headerItems = value;
        }

        /// <summary>
        /// Report header style.
        /// </summary>
        [Input("style")]
        public Input<string>? Style { get; set; }

        public ReportLayoutPageHeaderArgs()
        {
        }
        public static new ReportLayoutPageHeaderArgs Empty => new ReportLayoutPageHeaderArgs();
    }
}
