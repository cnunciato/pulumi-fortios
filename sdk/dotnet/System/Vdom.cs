// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System
{
    /// <summary>
    /// Configure virtual domain.
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
    ///     var trname = new Fortios.System.Vdom("trname", new()
    ///     {
    ///         ShortName = "testvdom",
    ///         Temporary = 0,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// System Vdom can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/vdom:Vdom labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/vdom:Vdom labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/vdom:Vdom")]
    public partial class Vdom : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Flag.
        /// </summary>
        [Output("flag")]
        public Output<int> Flag { get; private set; } = null!;

        /// <summary>
        /// VDOM name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// VDOM short name.
        /// </summary>
        [Output("shortName")]
        public Output<string> ShortName { get; private set; } = null!;

        /// <summary>
        /// Temporary.
        /// </summary>
        [Output("temporary")]
        public Output<int> Temporary { get; private set; } = null!;

        /// <summary>
        /// Virtual cluster ID (0 - 4294967295).
        /// </summary>
        [Output("vclusterId")]
        public Output<int> VclusterId { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Vdom resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vdom(string name, VdomArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/vdom:Vdom", name, args ?? new VdomArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Vdom(string name, Input<string> id, VdomState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/vdom:Vdom", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Vdom resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vdom Get(string name, Input<string> id, VdomState? state = null, CustomResourceOptions? options = null)
        {
            return new Vdom(name, id, state, options);
        }
    }

    public sealed class VdomArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Flag.
        /// </summary>
        [Input("flag")]
        public Input<int>? Flag { get; set; }

        /// <summary>
        /// VDOM name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// VDOM short name.
        /// </summary>
        [Input("shortName")]
        public Input<string>? ShortName { get; set; }

        /// <summary>
        /// Temporary.
        /// </summary>
        [Input("temporary")]
        public Input<int>? Temporary { get; set; }

        /// <summary>
        /// Virtual cluster ID (0 - 4294967295).
        /// </summary>
        [Input("vclusterId")]
        public Input<int>? VclusterId { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public VdomArgs()
        {
        }
        public static new VdomArgs Empty => new VdomArgs();
    }

    public sealed class VdomState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Flag.
        /// </summary>
        [Input("flag")]
        public Input<int>? Flag { get; set; }

        /// <summary>
        /// VDOM name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// VDOM short name.
        /// </summary>
        [Input("shortName")]
        public Input<string>? ShortName { get; set; }

        /// <summary>
        /// Temporary.
        /// </summary>
        [Input("temporary")]
        public Input<int>? Temporary { get; set; }

        /// <summary>
        /// Virtual cluster ID (0 - 4294967295).
        /// </summary>
        [Input("vclusterId")]
        public Input<int>? VclusterId { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public VdomState()
        {
        }
        public static new VdomState Empty => new VdomState();
    }
}