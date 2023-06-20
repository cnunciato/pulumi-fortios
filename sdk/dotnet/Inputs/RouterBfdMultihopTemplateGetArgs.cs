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

    public sealed class RouterBfdMultihopTemplateGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication mode. Valid values: `none`, `md5`.
        /// </summary>
        [Input("authMode")]
        public Input<string>? AuthMode { get; set; }

        /// <summary>
        /// BFD desired minimal transmit interval (milliseconds).
        /// </summary>
        [Input("bfdDesiredMinTx")]
        public Input<int>? BfdDesiredMinTx { get; set; }

        /// <summary>
        /// BFD detection multiplier.
        /// </summary>
        [Input("bfdDetectMult")]
        public Input<int>? BfdDetectMult { get; set; }

        /// <summary>
        /// BFD required minimal receive interval (milliseconds).
        /// </summary>
        [Input("bfdRequiredMinRx")]
        public Input<int>? BfdRequiredMinRx { get; set; }

        /// <summary>
        /// Destination prefix.
        /// </summary>
        [Input("dst")]
        public Input<string>? Dst { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("md5Key")]
        private Input<string>? _md5Key;

        /// <summary>
        /// MD5 key of key ID 1.
        /// </summary>
        public Input<string>? Md5Key
        {
            get => _md5Key;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _md5Key = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Source prefix.
        /// </summary>
        [Input("src")]
        public Input<string>? Src { get; set; }

        public RouterBfdMultihopTemplateGetArgs()
        {
        }
        public static new RouterBfdMultihopTemplateGetArgs Empty => new RouterBfdMultihopTemplateGetArgs();
    }
}
