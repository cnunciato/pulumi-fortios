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

    public sealed class ExtensioncontrollerExtenderprofileCellularSmsNotificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SMS alert list. The structure of `alert` block is documented below.
        /// </summary>
        [Input("alert")]
        public Input<Inputs.ExtensioncontrollerExtenderprofileCellularSmsNotificationAlertArgs>? Alert { get; set; }

        [Input("receivers")]
        private InputList<Inputs.ExtensioncontrollerExtenderprofileCellularSmsNotificationReceiverArgs>? _receivers;

        /// <summary>
        /// SMS notification receiver list. The structure of `receiver` block is documented below.
        /// </summary>
        public InputList<Inputs.ExtensioncontrollerExtenderprofileCellularSmsNotificationReceiverArgs> Receivers
        {
            get => _receivers ?? (_receivers = new InputList<Inputs.ExtensioncontrollerExtenderprofileCellularSmsNotificationReceiverArgs>());
            set => _receivers = value;
        }

        /// <summary>
        /// FortiExtender SMS notification status. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ExtensioncontrollerExtenderprofileCellularSmsNotificationArgs()
        {
        }
        public static new ExtensioncontrollerExtenderprofileCellularSmsNotificationArgs Empty => new ExtensioncontrollerExtenderprofileCellularSmsNotificationArgs();
    }
}
