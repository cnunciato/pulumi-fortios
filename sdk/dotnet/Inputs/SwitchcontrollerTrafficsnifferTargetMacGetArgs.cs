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

    public sealed class SwitchcontrollerTrafficsnifferTargetMacGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description for the sniffer MAC.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Sniffer MAC.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        public SwitchcontrollerTrafficsnifferTargetMacGetArgs()
        {
        }
        public static new SwitchcontrollerTrafficsnifferTargetMacGetArgs Empty => new SwitchcontrollerTrafficsnifferTargetMacGetArgs();
    }
}
