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

    public sealed class WirelesscontrollerQosprofileDscpWmmBkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DSCP WMM mapping numbers (0 - 63).
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public WirelesscontrollerQosprofileDscpWmmBkArgs()
        {
        }
        public static new WirelesscontrollerQosprofileDscpWmmBkArgs Empty => new WirelesscontrollerQosprofileDscpWmmBkArgs();
    }
}