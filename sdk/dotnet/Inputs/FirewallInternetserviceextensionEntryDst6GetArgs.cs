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

    public sealed class FirewallInternetserviceextensionEntryDst6GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Select the destination address or address group object from available options.
        /// 
        /// The `dst6` block supports:
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallInternetserviceextensionEntryDst6GetArgs()
        {
        }
        public static new FirewallInternetserviceextensionEntryDst6GetArgs Empty => new FirewallInternetserviceextensionEntryDst6GetArgs();
    }
}
