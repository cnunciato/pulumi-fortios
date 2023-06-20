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

    public sealed class SwitchcontrollerManagedswitchSwitchLogGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable to configure local logging settings that override global logging settings. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("localOverride")]
        public Input<string>? LocalOverride { get; set; }

        /// <summary>
        /// Severity of FortiSwitch logs that are added to the FortiGate event log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// Enable/disable adding FortiSwitch logs to the FortiGate event log. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public SwitchcontrollerManagedswitchSwitchLogGetArgs()
        {
        }
        public static new SwitchcontrollerManagedswitchSwitchLogGetArgs Empty => new SwitchcontrollerManagedswitchSwitchLogGetArgs();
    }
}
