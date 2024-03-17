// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Dhcp6.Inputs
{

    public sealed class ServerPrefixRangeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// End of prefix range.
        /// </summary>
        [Input("endPrefix")]
        public Input<string>? EndPrefix { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Prefix length.
        /// </summary>
        [Input("prefixLength")]
        public Input<int>? PrefixLength { get; set; }

        /// <summary>
        /// Start of prefix range.
        /// </summary>
        [Input("startPrefix")]
        public Input<string>? StartPrefix { get; set; }

        public ServerPrefixRangeGetArgs()
        {
        }
        public static new ServerPrefixRangeGetArgs Empty => new ServerPrefixRangeGetArgs();
    }
}