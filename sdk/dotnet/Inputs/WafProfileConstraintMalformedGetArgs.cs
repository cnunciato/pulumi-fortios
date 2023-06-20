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

    public sealed class WafProfileConstraintMalformedGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action. Valid values: `allow`, `block`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Enable/disable logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Severity. Valid values: `high`, `medium`, `low`.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// Enable/disable the constraint. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public WafProfileConstraintMalformedGetArgs()
        {
        }
        public static new WafProfileConstraintMalformedGetArgs Empty => new WafProfileConstraintMalformedGetArgs();
    }
}
