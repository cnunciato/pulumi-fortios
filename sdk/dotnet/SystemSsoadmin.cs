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
    /// Configure SSO admin users. Applies to FortiOS Version `&gt;= 6.2.4`.
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
    ///     var trname = new Fortios.SystemSsoadmin("trname", new()
    ///     {
    ///         Accprofile = "super_admin",
    ///         Vdoms = new[]
    ///         {
    ///             new Fortios.Inputs.SystemSsoadminVdomArgs
    ///             {
    ///                 Name = "root",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System SsoAdmin can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemSsoadmin:SystemSsoadmin labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemSsoadmin:SystemSsoadmin labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/systemSsoadmin:SystemSsoadmin")]
    public partial class SystemSsoadmin : global::Pulumi.CustomResource
    {
        /// <summary>
        /// SSO admin user access profile.
        /// </summary>
        [Output("accprofile")]
        public Output<string> Accprofile { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// The FortiOS version to ignore release overview prompt for.
        /// </summary>
        [Output("guiIgnoreReleaseOverviewVersion")]
        public Output<string> GuiIgnoreReleaseOverviewVersion { get; private set; } = null!;

        /// <summary>
        /// SSO admin name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
        /// </summary>
        [Output("vdoms")]
        public Output<ImmutableArray<Outputs.SystemSsoadminVdom>> Vdoms { get; private set; } = null!;


        /// <summary>
        /// Create a SystemSsoadmin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemSsoadmin(string name, SystemSsoadminArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/systemSsoadmin:SystemSsoadmin", name, args ?? new SystemSsoadminArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemSsoadmin(string name, Input<string> id, SystemSsoadminState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemSsoadmin:SystemSsoadmin", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemSsoadmin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemSsoadmin Get(string name, Input<string> id, SystemSsoadminState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemSsoadmin(name, id, state, options);
        }
    }

    public sealed class SystemSsoadminArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SSO admin user access profile.
        /// </summary>
        [Input("accprofile", required: true)]
        public Input<string> Accprofile { get; set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// The FortiOS version to ignore release overview prompt for.
        /// </summary>
        [Input("guiIgnoreReleaseOverviewVersion")]
        public Input<string>? GuiIgnoreReleaseOverviewVersion { get; set; }

        /// <summary>
        /// SSO admin name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        [Input("vdoms")]
        private InputList<Inputs.SystemSsoadminVdomArgs>? _vdoms;

        /// <summary>
        /// Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSsoadminVdomArgs> Vdoms
        {
            get => _vdoms ?? (_vdoms = new InputList<Inputs.SystemSsoadminVdomArgs>());
            set => _vdoms = value;
        }

        public SystemSsoadminArgs()
        {
        }
        public static new SystemSsoadminArgs Empty => new SystemSsoadminArgs();
    }

    public sealed class SystemSsoadminState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SSO admin user access profile.
        /// </summary>
        [Input("accprofile")]
        public Input<string>? Accprofile { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// The FortiOS version to ignore release overview prompt for.
        /// </summary>
        [Input("guiIgnoreReleaseOverviewVersion")]
        public Input<string>? GuiIgnoreReleaseOverviewVersion { get; set; }

        /// <summary>
        /// SSO admin name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        [Input("vdoms")]
        private InputList<Inputs.SystemSsoadminVdomGetArgs>? _vdoms;

        /// <summary>
        /// Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSsoadminVdomGetArgs> Vdoms
        {
            get => _vdoms ?? (_vdoms = new InputList<Inputs.SystemSsoadminVdomGetArgs>());
            set => _vdoms = value;
        }

        public SystemSsoadminState()
        {
        }
        public static new SystemSsoadminState Empty => new SystemSsoadminState();
    }
}
