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
    /// This resource supports Create/Read/Update/Delete devicemanager script for FortiManager.
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
    ///     var test1 = new Fortios.FmgDevicemanagerScript("test1", new()
    ///     {
    ///         Content = @"config system interface 
    ///  edit port3 
    /// 	 set vdom ""root""
    /// 	 set ip 10.7.0.200 255.255.0.0 
    /// 	 set allowaccess ping http https
    /// 	 next 
    ///  end
    /// ",
    ///         Description = "description",
    ///         Target = "remote_device",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FortiosResourceType("fortios:index/fmgDevicemanagerScript:FmgDevicemanagerScript")]
    public partial class FmgDevicemanagerScript : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ADOM name. default is 'root'.
        /// </summary>
        [Output("adom")]
        public Output<string?> Adom { get; private set; } = null!;

        /// <summary>
        /// Script content, only cli script is supported now
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// Description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Script name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Script target, Enum: ["device_database", "remote_device", "adom_database"]
        /// </summary>
        [Output("target")]
        public Output<string?> Target { get; private set; } = null!;


        /// <summary>
        /// Create a FmgDevicemanagerScript resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FmgDevicemanagerScript(string name, FmgDevicemanagerScriptArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/fmgDevicemanagerScript:FmgDevicemanagerScript", name, args ?? new FmgDevicemanagerScriptArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FmgDevicemanagerScript(string name, Input<string> id, FmgDevicemanagerScriptState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/fmgDevicemanagerScript:FmgDevicemanagerScript", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FmgDevicemanagerScript resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FmgDevicemanagerScript Get(string name, Input<string> id, FmgDevicemanagerScriptState? state = null, CustomResourceOptions? options = null)
        {
            return new FmgDevicemanagerScript(name, id, state, options);
        }
    }

    public sealed class FmgDevicemanagerScriptArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ADOM name. default is 'root'.
        /// </summary>
        [Input("adom")]
        public Input<string>? Adom { get; set; }

        /// <summary>
        /// Script content, only cli script is supported now
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Script name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Script target, Enum: ["device_database", "remote_device", "adom_database"]
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        public FmgDevicemanagerScriptArgs()
        {
        }
        public static new FmgDevicemanagerScriptArgs Empty => new FmgDevicemanagerScriptArgs();
    }

    public sealed class FmgDevicemanagerScriptState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ADOM name. default is 'root'.
        /// </summary>
        [Input("adom")]
        public Input<string>? Adom { get; set; }

        /// <summary>
        /// Script content, only cli script is supported now
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Script name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Script target, Enum: ["device_database", "remote_device", "adom_database"]
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        public FmgDevicemanagerScriptState()
        {
        }
        public static new FmgDevicemanagerScriptState Empty => new FmgDevicemanagerScriptState();
    }
}
