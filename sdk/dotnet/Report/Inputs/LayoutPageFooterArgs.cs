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

    public sealed class LayoutPageFooterArgs : global::Pulumi.ResourceArgs
    {
        [Input("footerItems")]
        private InputList<Inputs.LayoutPageFooterFooterItemArgs>? _footerItems;

        /// <summary>
        /// Configure report footer item. The structure of `footer_item` block is documented below.
        /// </summary>
        public InputList<Inputs.LayoutPageFooterFooterItemArgs> FooterItems
        {
            get => _footerItems ?? (_footerItems = new InputList<Inputs.LayoutPageFooterFooterItemArgs>());
            set => _footerItems = value;
        }

        /// <summary>
        /// Report footer style.
        /// </summary>
        [Input("style")]
        public Input<string>? Style { get; set; }

        public LayoutPageFooterArgs()
        {
        }
        public static new LayoutPageFooterArgs Empty => new LayoutPageFooterArgs();
    }
}