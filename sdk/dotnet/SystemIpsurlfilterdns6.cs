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
    /// Configure IPS URL filter IPv6 DNS servers.
    /// 
    /// ## Import
    /// 
    /// System IpsUrlfilterDns6 can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemIpsurlfilterdns6:SystemIpsurlfilterdns6 labelname {{address6}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemIpsurlfilterdns6:SystemIpsurlfilterdns6 labelname {{address6}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/systemIpsurlfilterdns6:SystemIpsurlfilterdns6")]
    public partial class SystemIpsurlfilterdns6 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IPv6 address of DNS server.
        /// </summary>
        [Output("address6")]
        public Output<string> Address6 { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SystemIpsurlfilterdns6 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemIpsurlfilterdns6(string name, SystemIpsurlfilterdns6Args? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemIpsurlfilterdns6:SystemIpsurlfilterdns6", name, args ?? new SystemIpsurlfilterdns6Args(), MakeResourceOptions(options, ""))
        {
        }

        private SystemIpsurlfilterdns6(string name, Input<string> id, SystemIpsurlfilterdns6State? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemIpsurlfilterdns6:SystemIpsurlfilterdns6", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemIpsurlfilterdns6 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemIpsurlfilterdns6 Get(string name, Input<string> id, SystemIpsurlfilterdns6State? state = null, CustomResourceOptions? options = null)
        {
            return new SystemIpsurlfilterdns6(name, id, state, options);
        }
    }

    public sealed class SystemIpsurlfilterdns6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IPv6 address of DNS server.
        /// </summary>
        [Input("address6")]
        public Input<string>? Address6 { get; set; }

        /// <summary>
        /// Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemIpsurlfilterdns6Args()
        {
        }
        public static new SystemIpsurlfilterdns6Args Empty => new SystemIpsurlfilterdns6Args();
    }

    public sealed class SystemIpsurlfilterdns6State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IPv6 address of DNS server.
        /// </summary>
        [Input("address6")]
        public Input<string>? Address6 { get; set; }

        /// <summary>
        /// Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemIpsurlfilterdns6State()
        {
        }
        public static new SystemIpsurlfilterdns6State Empty => new SystemIpsurlfilterdns6State();
    }
}
