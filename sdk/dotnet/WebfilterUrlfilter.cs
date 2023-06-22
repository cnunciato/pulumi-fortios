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
    /// Configure URL filter lists.
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
    ///     var trname = new Fortios.WebfilterUrlfilter("trname", new()
    ///     {
    ///         Fosid = 1,
    ///         IpAddrBlock = "enable",
    ///         OneArmIpsUrlfilter = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Webfilter Urlfilter can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/webfilterUrlfilter:WebfilterUrlfilter labelname {{fosid}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/webfilterUrlfilter:WebfilterUrlfilter labelname {{fosid}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/webfilterUrlfilter:WebfilterUrlfilter")]
    public partial class WebfilterUrlfilter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optional comments.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// URL filter entries. The structure of `entries` block is documented below.
        /// </summary>
        [Output("entries")]
        public Output<ImmutableArray<Outputs.WebfilterUrlfilterEntry>> Entries { get; private set; } = null!;

        /// <summary>
        /// ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Enable/disable blocking URLs when the hostname appears as an IP address. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ipAddrBlock")]
        public Output<string> IpAddrBlock { get; private set; } = null!;

        /// <summary>
        /// Name of URL filter list.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable DNS resolver for one-arm IPS URL filter operation. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("oneArmIpsUrlfilter")]
        public Output<string> OneArmIpsUrlfilter { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a WebfilterUrlfilter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebfilterUrlfilter(string name, WebfilterUrlfilterArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/webfilterUrlfilter:WebfilterUrlfilter", name, args ?? new WebfilterUrlfilterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebfilterUrlfilter(string name, Input<string> id, WebfilterUrlfilterState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/webfilterUrlfilter:WebfilterUrlfilter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WebfilterUrlfilter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebfilterUrlfilter Get(string name, Input<string> id, WebfilterUrlfilterState? state = null, CustomResourceOptions? options = null)
        {
            return new WebfilterUrlfilter(name, id, state, options);
        }
    }

    public sealed class WebfilterUrlfilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("entries")]
        private InputList<Inputs.WebfilterUrlfilterEntryArgs>? _entries;

        /// <summary>
        /// URL filter entries. The structure of `entries` block is documented below.
        /// </summary>
        public InputList<Inputs.WebfilterUrlfilterEntryArgs> Entries
        {
            get => _entries ?? (_entries = new InputList<Inputs.WebfilterUrlfilterEntryArgs>());
            set => _entries = value;
        }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("fosid", required: true)]
        public Input<int> Fosid { get; set; } = null!;

        /// <summary>
        /// Enable/disable blocking URLs when the hostname appears as an IP address. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipAddrBlock")]
        public Input<string>? IpAddrBlock { get; set; }

        /// <summary>
        /// Name of URL filter list.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable DNS resolver for one-arm IPS URL filter operation. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("oneArmIpsUrlfilter")]
        public Input<string>? OneArmIpsUrlfilter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public WebfilterUrlfilterArgs()
        {
        }
        public static new WebfilterUrlfilterArgs Empty => new WebfilterUrlfilterArgs();
    }

    public sealed class WebfilterUrlfilterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("entries")]
        private InputList<Inputs.WebfilterUrlfilterEntryGetArgs>? _entries;

        /// <summary>
        /// URL filter entries. The structure of `entries` block is documented below.
        /// </summary>
        public InputList<Inputs.WebfilterUrlfilterEntryGetArgs> Entries
        {
            get => _entries ?? (_entries = new InputList<Inputs.WebfilterUrlfilterEntryGetArgs>());
            set => _entries = value;
        }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Enable/disable blocking URLs when the hostname appears as an IP address. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipAddrBlock")]
        public Input<string>? IpAddrBlock { get; set; }

        /// <summary>
        /// Name of URL filter list.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable DNS resolver for one-arm IPS URL filter operation. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("oneArmIpsUrlfilter")]
        public Input<string>? OneArmIpsUrlfilter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public WebfilterUrlfilterState()
        {
        }
        public static new WebfilterUrlfilterState Empty => new WebfilterUrlfilterState();
    }
}