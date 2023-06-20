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

    public sealed class SystemSdwanZoneGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Zone name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Method of selecting member if more than one meets the SLA.
        /// </summary>
        [Input("serviceSlaTieBreak")]
        public Input<string>? ServiceSlaTieBreak { get; set; }

        public SystemSdwanZoneGetArgs()
        {
        }
        public static new SystemSdwanZoneGetArgs Empty => new SystemSdwanZoneGetArgs();
    }
}
