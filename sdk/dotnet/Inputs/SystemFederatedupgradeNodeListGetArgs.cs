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

    public sealed class SystemFederatedupgradeNodeListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The serial of the FortiGate that controls this device
        /// </summary>
        [Input("coordinatingFortigate")]
        public Input<string>? CoordinatingFortigate { get; set; }

        /// <summary>
        /// What type of device this node represents.
        /// </summary>
        [Input("deviceType")]
        public Input<string>? DeviceType { get; set; }

        /// <summary>
        /// Serial number of the node to include.
        /// </summary>
        [Input("serial")]
        public Input<string>? Serial { get; set; }

        /// <summary>
        /// When the upgrade was configured. Format hh:mm yyyy/mm/dd UTC.
        /// </summary>
        [Input("setupTime")]
        public Input<string>? SetupTime { get; set; }

        /// <summary>
        /// Scheduled time for the upgrade. Format hh:mm yyyy/mm/dd UTC.
        /// </summary>
        [Input("time")]
        public Input<string>? Time { get; set; }

        /// <summary>
        /// Whether the upgrade should be run immediately, or at a scheduled time. Valid values: `immediate`, `scheduled`.
        /// </summary>
        [Input("timing")]
        public Input<string>? Timing { get; set; }

        /// <summary>
        /// Image IDs to upgrade through.
        /// </summary>
        [Input("upgradePath")]
        public Input<string>? UpgradePath { get; set; }

        public SystemFederatedupgradeNodeListGetArgs()
        {
        }
        public static new SystemFederatedupgradeNodeListGetArgs Empty => new SystemFederatedupgradeNodeListGetArgs();
    }
}
