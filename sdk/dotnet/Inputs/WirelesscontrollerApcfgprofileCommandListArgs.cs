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

    public sealed class WirelesscontrollerApcfgprofileCommandListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Command ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// AP local configuration command name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("passwdValue")]
        private Input<string>? _passwdValue;

        /// <summary>
        /// AP local configuration command password value.
        /// </summary>
        public Input<string>? PasswdValue
        {
            get => _passwdValue;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passwdValue = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The command type (default = non-password). Valid values: `non-password`, `password`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// AP local configuration command value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public WirelesscontrollerApcfgprofileCommandListArgs()
        {
        }
        public static new WirelesscontrollerApcfgprofileCommandListArgs Empty => new WirelesscontrollerApcfgprofileCommandListArgs();
    }
}
