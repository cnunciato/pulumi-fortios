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
    /// Global firewall settings. Applies to FortiOS Version `&gt;= 7.2.1`.
    /// 
    /// ## Import
    /// 
    /// Firewall Global can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/firewallGlobal:FirewallGlobal labelname FirewallGlobal
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/firewallGlobal:FirewallGlobal labelname FirewallGlobal
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/firewallGlobal:FirewallGlobal")]
    public partial class FirewallGlobal : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Persistency of banned IPs across power cycling. Valid values: `disabled`, `permanent-only`, `all`.
        /// </summary>
        [Output("bannedIpPersistency")]
        public Output<string> BannedIpPersistency { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallGlobal resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallGlobal(string name, FirewallGlobalArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/firewallGlobal:FirewallGlobal", name, args ?? new FirewallGlobalArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallGlobal(string name, Input<string> id, FirewallGlobalState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/firewallGlobal:FirewallGlobal", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallGlobal resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallGlobal Get(string name, Input<string> id, FirewallGlobalState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallGlobal(name, id, state, options);
        }
    }

    public sealed class FirewallGlobalArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Persistency of banned IPs across power cycling. Valid values: `disabled`, `permanent-only`, `all`.
        /// </summary>
        [Input("bannedIpPersistency")]
        public Input<string>? BannedIpPersistency { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FirewallGlobalArgs()
        {
        }
        public static new FirewallGlobalArgs Empty => new FirewallGlobalArgs();
    }

    public sealed class FirewallGlobalState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Persistency of banned IPs across power cycling. Valid values: `disabled`, `permanent-only`, `all`.
        /// </summary>
        [Input("bannedIpPersistency")]
        public Input<string>? BannedIpPersistency { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FirewallGlobalState()
        {
        }
        public static new FirewallGlobalState Empty => new FirewallGlobalState();
    }
}
