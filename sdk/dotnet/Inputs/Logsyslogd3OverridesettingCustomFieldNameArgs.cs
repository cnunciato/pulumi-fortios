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

    public sealed class Logsyslogd3OverridesettingCustomFieldNameArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Field custom name.
        /// </summary>
        [Input("custom")]
        public Input<string>? Custom { get; set; }

        /// <summary>
        /// Entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Field name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public Logsyslogd3OverridesettingCustomFieldNameArgs()
        {
        }
        public static new Logsyslogd3OverridesettingCustomFieldNameArgs Empty => new Logsyslogd3OverridesettingCustomFieldNameArgs();
    }
}