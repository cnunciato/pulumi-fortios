// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Inputs
{

    public sealed class GlobalDisableDiscoveryGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Managed device ID.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GlobalDisableDiscoveryGetArgs()
        {
        }
        public static new GlobalDisableDiscoveryGetArgs Empty => new GlobalDisableDiscoveryGetArgs();
    }
}