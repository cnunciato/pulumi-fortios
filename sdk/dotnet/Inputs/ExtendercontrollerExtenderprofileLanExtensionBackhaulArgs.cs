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

    public sealed class ExtendercontrollerExtenderprofileLanExtensionBackhaulArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// FortiExtender LAN extension backhaul name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// FortiExtender uplink port. Valid values: `wan`, `lte1`, `lte2`, `port1`, `port2`, `port3`, `port4`, `port5`, `sfp`.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// FortiExtender uplink port. Valid values: `primary`, `secondary`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// WRR weight parameter
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ExtendercontrollerExtenderprofileLanExtensionBackhaulArgs()
        {
        }
        public static new ExtendercontrollerExtenderprofileLanExtensionBackhaulArgs Empty => new ExtendercontrollerExtenderprofileLanExtensionBackhaulArgs();
    }
}
