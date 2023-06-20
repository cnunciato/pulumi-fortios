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

    public sealed class EndpointcontrolProfileForticlientWinmacSettingsForticlientRunningAppArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application name.
        /// </summary>
        [Input("appName")]
        public Input<string>? AppName { get; set; }

        /// <summary>
        /// App's SHA256 signature.
        /// </summary>
        [Input("appSha256Signature")]
        public Input<string>? AppSha256Signature { get; set; }

        /// <summary>
        /// App's SHA256 Signature.
        /// </summary>
        [Input("appSha256Signature2")]
        public Input<string>? AppSha256Signature2 { get; set; }

        /// <summary>
        /// App's SHA256 Signature.
        /// </summary>
        [Input("appSha256Signature3")]
        public Input<string>? AppSha256Signature3 { get; set; }

        /// <summary>
        /// App's SHA256 Signature.
        /// </summary>
        [Input("appSha256Signature4")]
        public Input<string>? AppSha256Signature4 { get; set; }

        /// <summary>
        /// Application check rule. Valid values: `present`, `absent`.
        /// </summary>
        [Input("applicationCheckRule")]
        public Input<string>? ApplicationCheckRule { get; set; }

        /// <summary>
        /// Application ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Process name.
        /// </summary>
        [Input("processName")]
        public Input<string>? ProcessName { get; set; }

        /// <summary>
        /// Process name.
        /// </summary>
        [Input("processName2")]
        public Input<string>? ProcessName2 { get; set; }

        /// <summary>
        /// Process name.
        /// </summary>
        [Input("processName3")]
        public Input<string>? ProcessName3 { get; set; }

        /// <summary>
        /// Process name.
        /// </summary>
        [Input("processName4")]
        public Input<string>? ProcessName4 { get; set; }

        public EndpointcontrolProfileForticlientWinmacSettingsForticlientRunningAppArgs()
        {
        }
        public static new EndpointcontrolProfileForticlientWinmacSettingsForticlientRunningAppArgs Empty => new EndpointcontrolProfileForticlientWinmacSettingsForticlientRunningAppArgs();
    }
}
