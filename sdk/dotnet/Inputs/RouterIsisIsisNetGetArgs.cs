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

    public sealed class RouterIsisIsisNetGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// isis-net ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// IS-IS net xx.xxxx. ... .xxxx.xx.
        /// </summary>
        [Input("net")]
        public Input<string>? Net { get; set; }

        public RouterIsisIsisNetGetArgs()
        {
        }
        public static new RouterIsisIsisNetGetArgs Empty => new RouterIsisIsisNetGetArgs();
    }
}
