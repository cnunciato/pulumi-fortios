// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Inputs
{

    public sealed class BfdNeighborGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interface name.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// IPv4 address of the BFD neighbor.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        public BfdNeighborGetArgs()
        {
        }
        public static new BfdNeighborGetArgs Empty => new BfdNeighborGetArgs();
    }
}