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

    public sealed class WirelesscontrollerWtpprofileRadio2ChannelGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Channel number.
        /// </summary>
        [Input("chan")]
        public Input<string>? Chan { get; set; }

        public WirelesscontrollerWtpprofileRadio2ChannelGetArgs()
        {
        }
        public static new WirelesscontrollerWtpprofileRadio2ChannelGetArgs Empty => new WirelesscontrollerWtpprofileRadio2ChannelGetArgs();
    }
}
