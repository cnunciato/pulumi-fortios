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

    public sealed class SystemSdwanServicePriorityZoneGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Priority zone name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SystemSdwanServicePriorityZoneGetArgs()
        {
        }
        public static new SystemSdwanServicePriorityZoneGetArgs Empty => new SystemSdwanServicePriorityZoneGetArgs();
    }
}
