// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    /// <summary>
    /// Configure/list NAC devices learned on the managed FortiSwitch ports which matches NAC policy. Applies to FortiOS Version `6.4.0,6.4.1,6.4.2,6.4.10,7.0.0`.
    /// 
    /// ## Import
    /// 
    /// SwitchController NacDevice can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerNacdevice:SwitchcontrollerNacdevice labelname {{fosid}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerNacdevice:SwitchcontrollerNacdevice labelname {{fosid}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/switchcontrollerNacdevice:SwitchcontrollerNacdevice")]
    public partial class SwitchcontrollerNacdevice : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description for the learned NAC device.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Device ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Managed FortiSwitch port where NAC device is last learned.
        /// </summary>
        [Output("lastKnownPort")]
        public Output<string> LastKnownPort { get; private set; } = null!;

        /// <summary>
        /// Managed FortiSwitch where NAC device is last learned.
        /// </summary>
        [Output("lastKnownSwitch")]
        public Output<string> LastKnownSwitch { get; private set; } = null!;

        /// <summary>
        /// Device last seen.
        /// </summary>
        [Output("lastSeen")]
        public Output<int> LastSeen { get; private set; } = null!;

        /// <summary>
        /// MAC address of the learned NAC device.
        /// </summary>
        [Output("mac")]
        public Output<string> Mac { get; private set; } = null!;

        /// <summary>
        /// MAC policy to be applied on this learned NAC device.
        /// </summary>
        [Output("macPolicy")]
        public Output<string> MacPolicy { get; private set; } = null!;

        /// <summary>
        /// Matched NAC policy for the learned NAC device.
        /// </summary>
        [Output("matchedNacPolicy")]
        public Output<string> MatchedNacPolicy { get; private set; } = null!;

        /// <summary>
        /// Port policy to be applied on this learned NAC device.
        /// </summary>
        [Output("portPolicy")]
        public Output<string> PortPolicy { get; private set; } = null!;

        /// <summary>
        /// Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SwitchcontrollerNacdevice resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SwitchcontrollerNacdevice(string name, SwitchcontrollerNacdeviceArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerNacdevice:SwitchcontrollerNacdevice", name, args ?? new SwitchcontrollerNacdeviceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SwitchcontrollerNacdevice(string name, Input<string> id, SwitchcontrollerNacdeviceState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerNacdevice:SwitchcontrollerNacdevice", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SwitchcontrollerNacdevice resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SwitchcontrollerNacdevice Get(string name, Input<string> id, SwitchcontrollerNacdeviceState? state = null, CustomResourceOptions? options = null)
        {
            return new SwitchcontrollerNacdevice(name, id, state, options);
        }
    }

    public sealed class SwitchcontrollerNacdeviceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description for the learned NAC device.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Device ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Managed FortiSwitch port where NAC device is last learned.
        /// </summary>
        [Input("lastKnownPort")]
        public Input<string>? LastKnownPort { get; set; }

        /// <summary>
        /// Managed FortiSwitch where NAC device is last learned.
        /// </summary>
        [Input("lastKnownSwitch")]
        public Input<string>? LastKnownSwitch { get; set; }

        /// <summary>
        /// Device last seen.
        /// </summary>
        [Input("lastSeen")]
        public Input<int>? LastSeen { get; set; }

        /// <summary>
        /// MAC address of the learned NAC device.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        /// <summary>
        /// MAC policy to be applied on this learned NAC device.
        /// </summary>
        [Input("macPolicy")]
        public Input<string>? MacPolicy { get; set; }

        /// <summary>
        /// Matched NAC policy for the learned NAC device.
        /// </summary>
        [Input("matchedNacPolicy")]
        public Input<string>? MatchedNacPolicy { get; set; }

        /// <summary>
        /// Port policy to be applied on this learned NAC device.
        /// </summary>
        [Input("portPolicy")]
        public Input<string>? PortPolicy { get; set; }

        /// <summary>
        /// Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchcontrollerNacdeviceArgs()
        {
        }
        public static new SwitchcontrollerNacdeviceArgs Empty => new SwitchcontrollerNacdeviceArgs();
    }

    public sealed class SwitchcontrollerNacdeviceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description for the learned NAC device.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Device ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Managed FortiSwitch port where NAC device is last learned.
        /// </summary>
        [Input("lastKnownPort")]
        public Input<string>? LastKnownPort { get; set; }

        /// <summary>
        /// Managed FortiSwitch where NAC device is last learned.
        /// </summary>
        [Input("lastKnownSwitch")]
        public Input<string>? LastKnownSwitch { get; set; }

        /// <summary>
        /// Device last seen.
        /// </summary>
        [Input("lastSeen")]
        public Input<int>? LastSeen { get; set; }

        /// <summary>
        /// MAC address of the learned NAC device.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        /// <summary>
        /// MAC policy to be applied on this learned NAC device.
        /// </summary>
        [Input("macPolicy")]
        public Input<string>? MacPolicy { get; set; }

        /// <summary>
        /// Matched NAC policy for the learned NAC device.
        /// </summary>
        [Input("matchedNacPolicy")]
        public Input<string>? MatchedNacPolicy { get; set; }

        /// <summary>
        /// Port policy to be applied on this learned NAC device.
        /// </summary>
        [Input("portPolicy")]
        public Input<string>? PortPolicy { get; set; }

        /// <summary>
        /// Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchcontrollerNacdeviceState()
        {
        }
        public static new SwitchcontrollerNacdeviceState Empty => new SwitchcontrollerNacdeviceState();
    }
}
