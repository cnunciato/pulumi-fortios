// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router
{
    /// <summary>
    /// Configure key-chain.
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
    ///     var trname = new Fortios.Router.Keychain("trname", new()
    ///     {
    ///         Keys = new[]
    ///         {
    ///             new Fortios.Router.Inputs.KeychainKeyArgs
    ///             {
    ///                 AcceptLifetime = "04:00:00 01 01 2008 04:00:00 01 01 2022",
    ///                 KeyString = "ewiwn3i23232s212",
    ///                 SendLifetime = "04:00:00 01 01 2008 04:00:00 01 01 2022",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Router KeyChain can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:router/keychain:Keychain labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:router/keychain:Keychain labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:router/keychain:Keychain")]
    public partial class Keychain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Configuration method to edit key settings. The structure of `key` block is documented below.
        /// </summary>
        [Output("keys")]
        public Output<ImmutableArray<Outputs.KeychainKey>> Keys { get; private set; } = null!;

        /// <summary>
        /// Key-chain name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Keychain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Keychain(string name, KeychainArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:router/keychain:Keychain", name, args ?? new KeychainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Keychain(string name, Input<string> id, KeychainState? state = null, CustomResourceOptions? options = null)
            : base("fortios:router/keychain:Keychain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Keychain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Keychain Get(string name, Input<string> id, KeychainState? state = null, CustomResourceOptions? options = null)
        {
            return new Keychain(name, id, state, options);
        }
    }

    public sealed class KeychainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("keys")]
        private InputList<Inputs.KeychainKeyArgs>? _keys;

        /// <summary>
        /// Configuration method to edit key settings. The structure of `key` block is documented below.
        /// </summary>
        public InputList<Inputs.KeychainKeyArgs> Keys
        {
            get => _keys ?? (_keys = new InputList<Inputs.KeychainKeyArgs>());
            set => _keys = value;
        }

        /// <summary>
        /// Key-chain name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public KeychainArgs()
        {
        }
        public static new KeychainArgs Empty => new KeychainArgs();
    }

    public sealed class KeychainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("keys")]
        private InputList<Inputs.KeychainKeyGetArgs>? _keys;

        /// <summary>
        /// Configuration method to edit key settings. The structure of `key` block is documented below.
        /// </summary>
        public InputList<Inputs.KeychainKeyGetArgs> Keys
        {
            get => _keys ?? (_keys = new InputList<Inputs.KeychainKeyGetArgs>());
            set => _keys = value;
        }

        /// <summary>
        /// Key-chain name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public KeychainState()
        {
        }
        public static new KeychainState Empty => new KeychainState();
    }
}