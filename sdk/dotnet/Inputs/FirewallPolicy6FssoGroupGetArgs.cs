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

    public sealed class FirewallPolicy6FssoGroupGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Names of FSSO groups.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallPolicy6FssoGroupGetArgs()
        {
        }
        public static new FirewallPolicy6FssoGroupGetArgs Empty => new FirewallPolicy6FssoGroupGetArgs();
    }
}
