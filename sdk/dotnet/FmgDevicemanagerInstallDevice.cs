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
    /// This resource supports installing devicemanager script from FortiManager to the related device
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test1 = new Fortios.FmgDevicemanagerInstallDevice("test1", new()
    ///     {
    ///         TargetDevname = "FGVM64-test",
    ///         Timeout = 5,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FortiosResourceType("fortios:index/fmgDevicemanagerInstallDevice:FmgDevicemanagerInstallDevice")]
    public partial class FmgDevicemanagerInstallDevice : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Source ADOM name. default is 'root'
        /// </summary>
        [Output("adom")]
        public Output<string?> Adom { get; private set; } = null!;

        /// <summary>
        /// Target device name.
        /// </summary>
        [Output("targetDevname")]
        public Output<string> TargetDevname { get; private set; } = null!;

        /// <summary>
        /// Timeout for installing the script to the target, default: 3 minutes.
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;

        /// <summary>
        /// Vdom of managed device. default is 'root'
        /// </summary>
        [Output("vdom")]
        public Output<string?> Vdom { get; private set; } = null!;


        /// <summary>
        /// Create a FmgDevicemanagerInstallDevice resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FmgDevicemanagerInstallDevice(string name, FmgDevicemanagerInstallDeviceArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/fmgDevicemanagerInstallDevice:FmgDevicemanagerInstallDevice", name, args ?? new FmgDevicemanagerInstallDeviceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FmgDevicemanagerInstallDevice(string name, Input<string> id, FmgDevicemanagerInstallDeviceState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/fmgDevicemanagerInstallDevice:FmgDevicemanagerInstallDevice", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FmgDevicemanagerInstallDevice resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FmgDevicemanagerInstallDevice Get(string name, Input<string> id, FmgDevicemanagerInstallDeviceState? state = null, CustomResourceOptions? options = null)
        {
            return new FmgDevicemanagerInstallDevice(name, id, state, options);
        }
    }

    public sealed class FmgDevicemanagerInstallDeviceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Source ADOM name. default is 'root'
        /// </summary>
        [Input("adom")]
        public Input<string>? Adom { get; set; }

        /// <summary>
        /// Target device name.
        /// </summary>
        [Input("targetDevname", required: true)]
        public Input<string> TargetDevname { get; set; } = null!;

        /// <summary>
        /// Timeout for installing the script to the target, default: 3 minutes.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Vdom of managed device. default is 'root'
        /// </summary>
        [Input("vdom")]
        public Input<string>? Vdom { get; set; }

        public FmgDevicemanagerInstallDeviceArgs()
        {
        }
        public static new FmgDevicemanagerInstallDeviceArgs Empty => new FmgDevicemanagerInstallDeviceArgs();
    }

    public sealed class FmgDevicemanagerInstallDeviceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Source ADOM name. default is 'root'
        /// </summary>
        [Input("adom")]
        public Input<string>? Adom { get; set; }

        /// <summary>
        /// Target device name.
        /// </summary>
        [Input("targetDevname")]
        public Input<string>? TargetDevname { get; set; }

        /// <summary>
        /// Timeout for installing the script to the target, default: 3 minutes.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Vdom of managed device. default is 'root'
        /// </summary>
        [Input("vdom")]
        public Input<string>? Vdom { get; set; }

        public FmgDevicemanagerInstallDeviceState()
        {
        }
        public static new FmgDevicemanagerInstallDeviceState Empty => new FmgDevicemanagerInstallDeviceState();
    }
}