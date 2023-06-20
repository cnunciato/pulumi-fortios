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

    public sealed class ExtendercontrollerExtenderprofileCellularGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// FortiExtender controller report configuration. The structure of `controller_report` block is documented below.
        /// </summary>
        [Input("controllerReport")]
        public Input<Inputs.ExtendercontrollerExtenderprofileCellularControllerReportGetArgs>? ControllerReport { get; set; }

        [Input("dataplans")]
        private InputList<Inputs.ExtendercontrollerExtenderprofileCellularDataplanGetArgs>? _dataplans;

        /// <summary>
        /// Dataplan names. The structure of `dataplan` block is documented below.
        /// </summary>
        public InputList<Inputs.ExtendercontrollerExtenderprofileCellularDataplanGetArgs> Dataplans
        {
            get => _dataplans ?? (_dataplans = new InputList<Inputs.ExtendercontrollerExtenderprofileCellularDataplanGetArgs>());
            set => _dataplans = value;
        }

        /// <summary>
        /// Configuration options for modem 1. The structure of `modem1` block is documented below.
        /// </summary>
        [Input("modem1")]
        public Input<Inputs.ExtendercontrollerExtenderprofileCellularModem1GetArgs>? Modem1 { get; set; }

        /// <summary>
        /// Configuration options for modem 2. The structure of `modem2` block is documented below.
        /// </summary>
        [Input("modem2")]
        public Input<Inputs.ExtendercontrollerExtenderprofileCellularModem2GetArgs>? Modem2 { get; set; }

        /// <summary>
        /// FortiExtender cellular SMS notification configuration. The structure of `sms_notification` block is documented below.
        /// </summary>
        [Input("smsNotification")]
        public Input<Inputs.ExtendercontrollerExtenderprofileCellularSmsNotificationGetArgs>? SmsNotification { get; set; }

        public ExtendercontrollerExtenderprofileCellularGetArgs()
        {
        }
        public static new ExtendercontrollerExtenderprofileCellularGetArgs Empty => new ExtendercontrollerExtenderprofileCellularGetArgs();
    }
}
