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

    public sealed class Wirelesscontrollerhotspot20HsprofileOsuProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// OSU provider name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public Wirelesscontrollerhotspot20HsprofileOsuProviderArgs()
        {
        }
        public static new Wirelesscontrollerhotspot20HsprofileOsuProviderArgs Empty => new Wirelesscontrollerhotspot20HsprofileOsuProviderArgs();
    }
}
