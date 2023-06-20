// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class ExtendercontrollerExtenderModem1AutoSwitch
    {
        /// <summary>
        /// Automatically switch based on data usage. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Dataplan;
        /// <summary>
        /// Auto switch by disconnect. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Disconnect;
        /// <summary>
        /// Automatically switch based on disconnect period.
        /// </summary>
        public readonly int? DisconnectPeriod;
        /// <summary>
        /// Automatically switch based on disconnect threshold.
        /// </summary>
        public readonly int? DisconnectThreshold;
        /// <summary>
        /// Automatically switch based on signal strength. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Signal;
        /// <summary>
        /// Auto switch with switch back multi-options. Valid values: `time`, `timer`.
        /// </summary>
        public readonly string? SwitchBack;
        /// <summary>
        /// Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).
        /// </summary>
        public readonly string? SwitchBackTime;
        /// <summary>
        /// Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).
        /// </summary>
        public readonly int? SwitchBackTimer;

        [OutputConstructor]
        private ExtendercontrollerExtenderModem1AutoSwitch(
            string? dataplan,

            string? disconnect,

            int? disconnectPeriod,

            int? disconnectThreshold,

            string? signal,

            string? switchBack,

            string? switchBackTime,

            int? switchBackTimer)
        {
            Dataplan = dataplan;
            Disconnect = disconnect;
            DisconnectPeriod = disconnectPeriod;
            DisconnectThreshold = disconnectThreshold;
            Signal = signal;
            SwitchBack = switchBack;
            SwitchBackTime = switchBackTime;
            SwitchBackTimer = switchBackTimer;
        }
    }
}
