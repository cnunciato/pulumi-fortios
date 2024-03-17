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

    public sealed class OspfAreaVirtualLinkMd5KeyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Area entry IP address.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("keyString")]
        private Input<string>? _keyString;

        /// <summary>
        /// Password for the key.
        /// </summary>
        public Input<string>? KeyString
        {
            get => _keyString;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _keyString = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public OspfAreaVirtualLinkMd5KeyGetArgs()
        {
        }
        public static new OspfAreaVirtualLinkMd5KeyGetArgs Empty => new OspfAreaVirtualLinkMd5KeyGetArgs();
    }
}