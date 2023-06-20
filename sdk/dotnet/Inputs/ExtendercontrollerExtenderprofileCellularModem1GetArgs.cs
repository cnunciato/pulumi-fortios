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

    public sealed class ExtendercontrollerExtenderprofileCellularModem1GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// FortiExtender auto switch configuration. The structure of `auto_switch` block is documented below.
        /// </summary>
        [Input("autoSwitch")]
        public Input<Inputs.ExtendercontrollerExtenderprofileCellularModem1AutoSwitchGetArgs>? AutoSwitch { get; set; }

        /// <summary>
        /// Connection status.
        /// </summary>
        [Input("connStatus")]
        public Input<int>? ConnStatus { get; set; }

        /// <summary>
        /// Default SIM selection. Valid values: `sim1`, `sim2`, `carrier`, `cost`.
        /// </summary>
        [Input("defaultSim")]
        public Input<string>? DefaultSim { get; set; }

        /// <summary>
        /// FortiExtender GPS enable/disable. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("gps")]
        public Input<string>? Gps { get; set; }

        /// <summary>
        /// Preferred carrier.
        /// </summary>
        [Input("preferredCarrier")]
        public Input<string>? PreferredCarrier { get; set; }

        /// <summary>
        /// Redundant interface.
        /// </summary>
        [Input("redundantIntf")]
        public Input<string>? RedundantIntf { get; set; }

        /// <summary>
        /// FortiExtender mode. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("redundantMode")]
        public Input<string>? RedundantMode { get; set; }

        /// <summary>
        /// SIM #1 PIN status. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sim1Pin")]
        public Input<string>? Sim1Pin { get; set; }

        /// <summary>
        /// SIM #1 PIN password.
        /// </summary>
        [Input("sim1PinCode")]
        public Input<string>? Sim1PinCode { get; set; }

        /// <summary>
        /// SIM #2 PIN status. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sim2Pin")]
        public Input<string>? Sim2Pin { get; set; }

        /// <summary>
        /// SIM #2 PIN password.
        /// </summary>
        [Input("sim2PinCode")]
        public Input<string>? Sim2PinCode { get; set; }

        public ExtendercontrollerExtenderprofileCellularModem1GetArgs()
        {
        }
        public static new ExtendercontrollerExtenderprofileCellularModem1GetArgs Empty => new ExtendercontrollerExtenderprofileCellularModem1GetArgs();
    }
}
