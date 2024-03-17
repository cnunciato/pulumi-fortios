// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.User
{
    /// <summary>
    /// Configure devices. Applies to FortiOS Version `&lt;= 6.2.0`.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.User.Device("trname", new()
    ///     {
    ///         Alias = "1",
    ///         Category = "amazon-device",
    ///         Mac = "08:00:20:0a:8c:6d",
    ///         Type = "unknown",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// User Device can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/device:Device labelname {{alias}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/device:Device labelname {{alias}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:user/device:Device")]
    public partial class Device : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Device alias.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// Image file for avatar (maximum 4K base64 encoded).
        /// </summary>
        [Output("avatar")]
        public Output<string?> Avatar { get; private set; } = null!;

        /// <summary>
        /// Device category. Valid values: `none`, `amazon-device`, `android-device`, `blackberry-device`, `fortinet-device`, `ios-device`, `windows-device`.
        /// </summary>
        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Device MAC address.
        /// </summary>
        [Output("mac")]
        public Output<string> Mac { get; private set; } = null!;

        /// <summary>
        /// Master device (optional).
        /// </summary>
        [Output("masterDevice")]
        public Output<string> MasterDevice { get; private set; } = null!;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        [Output("taggings")]
        public Output<ImmutableArray<Outputs.DeviceTagging>> Taggings { get; private set; } = null!;

        /// <summary>
        /// Device type. Valid values: `unknown`, `android-phone`, `android-tablet`, `blackberry-phone`, `blackberry-playbook`, `forticam`, `fortifone`, `fortinet-device`, `gaming-console`, `ip-phone`, `ipad`, `iphone`, `linux-pc`, `mac`, `media-streaming`, `printer`, `router-nat-device`, `windows-pc`, `windows-phone`, `windows-tablet`, `other-network-device`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// User name.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Device resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Device(string name, DeviceArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:user/device:Device", name, args ?? new DeviceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Device(string name, Input<string> id, DeviceState? state = null, CustomResourceOptions? options = null)
            : base("fortios:user/device:Device", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Device resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Device Get(string name, Input<string> id, DeviceState? state = null, CustomResourceOptions? options = null)
        {
            return new Device(name, id, state, options);
        }
    }

    public sealed class DeviceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Device alias.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Image file for avatar (maximum 4K base64 encoded).
        /// </summary>
        [Input("avatar")]
        public Input<string>? Avatar { get; set; }

        /// <summary>
        /// Device category. Valid values: `none`, `amazon-device`, `android-device`, `blackberry-device`, `fortinet-device`, `ios-device`, `windows-device`.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Device MAC address.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        /// <summary>
        /// Master device (optional).
        /// </summary>
        [Input("masterDevice")]
        public Input<string>? MasterDevice { get; set; }

        [Input("taggings")]
        private InputList<Inputs.DeviceTaggingArgs>? _taggings;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.DeviceTaggingArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.DeviceTaggingArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// Device type. Valid values: `unknown`, `android-phone`, `android-tablet`, `blackberry-phone`, `blackberry-playbook`, `forticam`, `fortifone`, `fortinet-device`, `gaming-console`, `ip-phone`, `ipad`, `iphone`, `linux-pc`, `mac`, `media-streaming`, `printer`, `router-nat-device`, `windows-pc`, `windows-phone`, `windows-tablet`, `other-network-device`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// User name.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DeviceArgs()
        {
        }
        public static new DeviceArgs Empty => new DeviceArgs();
    }

    public sealed class DeviceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Device alias.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Image file for avatar (maximum 4K base64 encoded).
        /// </summary>
        [Input("avatar")]
        public Input<string>? Avatar { get; set; }

        /// <summary>
        /// Device category. Valid values: `none`, `amazon-device`, `android-device`, `blackberry-device`, `fortinet-device`, `ios-device`, `windows-device`.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Device MAC address.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        /// <summary>
        /// Master device (optional).
        /// </summary>
        [Input("masterDevice")]
        public Input<string>? MasterDevice { get; set; }

        [Input("taggings")]
        private InputList<Inputs.DeviceTaggingGetArgs>? _taggings;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.DeviceTaggingGetArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.DeviceTaggingGetArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// Device type. Valid values: `unknown`, `android-phone`, `android-tablet`, `blackberry-phone`, `blackberry-playbook`, `forticam`, `fortifone`, `fortinet-device`, `gaming-console`, `ip-phone`, `ipad`, `iphone`, `linux-pc`, `mac`, `media-streaming`, `printer`, `router-nat-device`, `windows-pc`, `windows-phone`, `windows-tablet`, `other-network-device`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// User name.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DeviceState()
        {
        }
        public static new DeviceState Empty => new DeviceState();
    }
}