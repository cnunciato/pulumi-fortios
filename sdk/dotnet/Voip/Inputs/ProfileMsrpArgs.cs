// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Voip.Inputs
{

    public sealed class ProfileMsrpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable logging of MSRP violations. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("logViolations")]
        public Input<string>? LogViolations { get; set; }

        /// <summary>
        /// Maximum allowable MSRP message size (1-65535).
        /// </summary>
        [Input("maxMsgSize")]
        public Input<int>? MaxMsgSize { get; set; }

        /// <summary>
        /// Action for violation of max-msg-size. Valid values: `pass`, `block`, `reset`, `monitor`.
        /// </summary>
        [Input("maxMsgSizeAction")]
        public Input<string>? MaxMsgSizeAction { get; set; }

        /// <summary>
        /// Enable/disable MSRP. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ProfileMsrpArgs()
        {
        }
        public static new ProfileMsrpArgs Empty => new ProfileMsrpArgs();
    }
}