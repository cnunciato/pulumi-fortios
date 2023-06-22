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
    /// Configure IPS rule setting.
    /// 
    /// ## Import
    /// 
    /// Ips RuleSettings can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/ipsRulesettings:IpsRulesettings labelname {{fosid}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/ipsRulesettings:IpsRulesettings labelname {{fosid}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/ipsRulesettings:IpsRulesettings")]
    public partial class IpsRulesettings : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Rule ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a IpsRulesettings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpsRulesettings(string name, IpsRulesettingsArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/ipsRulesettings:IpsRulesettings", name, args ?? new IpsRulesettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpsRulesettings(string name, Input<string> id, IpsRulesettingsState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/ipsRulesettings:IpsRulesettings", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IpsRulesettings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpsRulesettings Get(string name, Input<string> id, IpsRulesettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new IpsRulesettings(name, id, state, options);
        }
    }

    public sealed class IpsRulesettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Rule ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IpsRulesettingsArgs()
        {
        }
        public static new IpsRulesettingsArgs Empty => new IpsRulesettingsArgs();
    }

    public sealed class IpsRulesettingsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Rule ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IpsRulesettingsState()
        {
        }
        public static new IpsRulesettingsState Empty => new IpsRulesettingsState();
    }
}