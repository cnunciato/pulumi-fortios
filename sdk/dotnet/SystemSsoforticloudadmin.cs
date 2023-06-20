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
    /// Configure FortiCloud SSO admin users. Applies to FortiOS Version `&gt;= 7.0.1`.
    /// 
    /// ## Import
    /// 
    /// System SsoForticloudAdmin can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemSsoforticloudadmin:SystemSsoforticloudadmin labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemSsoforticloudadmin:SystemSsoforticloudadmin labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/systemSsoforticloudadmin:SystemSsoforticloudadmin")]
    public partial class SystemSsoforticloudadmin : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// FortiCloud SSO admin name.
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
        public Output<ImmutableArray<Outputs.SystemSsoforticloudadminVdom>> Vdoms { get; private set; } = null!;


        /// <summary>
        /// Create a SystemSsoforticloudadmin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemSsoforticloudadmin(string name, SystemSsoforticloudadminArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemSsoforticloudadmin:SystemSsoforticloudadmin", name, args ?? new SystemSsoforticloudadminArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemSsoforticloudadmin(string name, Input<string> id, SystemSsoforticloudadminState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemSsoforticloudadmin:SystemSsoforticloudadmin", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemSsoforticloudadmin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemSsoforticloudadmin Get(string name, Input<string> id, SystemSsoforticloudadminState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemSsoforticloudadmin(name, id, state, options);
        }
    }

    public sealed class SystemSsoforticloudadminArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// FortiCloud SSO admin name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        [Input("vdoms")]
        private InputList<Inputs.SystemSsoforticloudadminVdomArgs>? _vdoms;

        /// <summary>
        /// Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSsoforticloudadminVdomArgs> Vdoms
        {
            get => _vdoms ?? (_vdoms = new InputList<Inputs.SystemSsoforticloudadminVdomArgs>());
            set => _vdoms = value;
        }

        public SystemSsoforticloudadminArgs()
        {
        }
        public static new SystemSsoforticloudadminArgs Empty => new SystemSsoforticloudadminArgs();
    }

    public sealed class SystemSsoforticloudadminState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// FortiCloud SSO admin name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        [Input("vdoms")]
        private InputList<Inputs.SystemSsoforticloudadminVdomGetArgs>? _vdoms;

        /// <summary>
        /// Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSsoforticloudadminVdomGetArgs> Vdoms
        {
            get => _vdoms ?? (_vdoms = new InputList<Inputs.SystemSsoforticloudadminVdomGetArgs>());
            set => _vdoms = value;
        }

        public SystemSsoforticloudadminState()
        {
        }
        public static new SystemSsoforticloudadminState Empty => new SystemSsoforticloudadminState();
    }
}
