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

    public sealed class WirelesscontrollerVapVlanPoolGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// WTP group name.
        /// </summary>
        [Input("wtpGroup")]
        public Input<string>? WtpGroup { get; set; }

        public WirelesscontrollerVapVlanPoolGetArgs()
        {
        }
        public static new WirelesscontrollerVapVlanPoolGetArgs Empty => new WirelesscontrollerVapVlanPoolGetArgs();
    }
}
