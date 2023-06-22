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

    public sealed class Wirelesscontrollerhotspot20IconIconListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Icon file.
        /// </summary>
        [Input("file")]
        public Input<string>? File { get; set; }

        /// <summary>
        /// Icon height.
        /// </summary>
        [Input("height")]
        public Input<int>? Height { get; set; }

        /// <summary>
        /// Language code.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// Icon name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Icon type. Valid values: `bmp`, `gif`, `jpeg`, `png`, `tiff`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Icon width.
        /// </summary>
        [Input("width")]
        public Input<int>? Width { get; set; }

        public Wirelesscontrollerhotspot20IconIconListGetArgs()
        {
        }
        public static new Wirelesscontrollerhotspot20IconIconListGetArgs Empty => new Wirelesscontrollerhotspot20IconIconListGetArgs();
    }
}