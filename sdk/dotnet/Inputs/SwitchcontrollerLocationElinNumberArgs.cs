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

    public sealed class SwitchcontrollerLocationElinNumberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configure ELIN callback number.
        /// </summary>
        [Input("elinNum")]
        public Input<string>? ElinNum { get; set; }

        /// <summary>
        /// Parent key name.
        /// </summary>
        [Input("parentKey")]
        public Input<string>? ParentKey { get; set; }

        public SwitchcontrollerLocationElinNumberArgs()
        {
        }
        public static new SwitchcontrollerLocationElinNumberArgs Empty => new SwitchcontrollerLocationElinNumberArgs();
    }
}
