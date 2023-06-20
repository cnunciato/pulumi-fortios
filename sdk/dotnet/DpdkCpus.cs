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
    /// Configure CPUs enabled to run engines in each DPDK stage. Applies to FortiOS Version `&gt;= 6.2.4`.
    /// 
    /// ## Import
    /// 
    /// Dpdk Cpus can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/dpdkCpus:DpdkCpus labelname DpdkCpus
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/dpdkCpus:DpdkCpus labelname DpdkCpus
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/dpdkCpus:DpdkCpus")]
    public partial class DpdkCpus : global::Pulumi.CustomResource
    {
        /// <summary>
        /// CPUs enabled to run DPDK IPS engines.
        /// </summary>
        [Output("ipsCpus")]
        public Output<string> IpsCpus { get; private set; } = null!;

        /// <summary>
        /// CPUs isolated to run only the DPDK engines with the exception of processes that have affinity explicitly set by either a user configuration or by their implementation.
        /// </summary>
        [Output("isolatedCpus")]
        public Output<string> IsolatedCpus { get; private set; } = null!;

        /// <summary>
        /// CPUs enabled to run DPDK RX engines.
        /// </summary>
        [Output("rxCpus")]
        public Output<string> RxCpus { get; private set; } = null!;

        /// <summary>
        /// CPUs enabled to run DPDK TX engines.
        /// </summary>
        [Output("txCpus")]
        public Output<string> TxCpus { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// CPUs enabled to run DPDK VNP engines.
        /// </summary>
        [Output("vnpCpus")]
        public Output<string> VnpCpus { get; private set; } = null!;


        /// <summary>
        /// Create a DpdkCpus resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DpdkCpus(string name, DpdkCpusArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/dpdkCpus:DpdkCpus", name, args ?? new DpdkCpusArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DpdkCpus(string name, Input<string> id, DpdkCpusState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/dpdkCpus:DpdkCpus", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DpdkCpus resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DpdkCpus Get(string name, Input<string> id, DpdkCpusState? state = null, CustomResourceOptions? options = null)
        {
            return new DpdkCpus(name, id, state, options);
        }
    }

    public sealed class DpdkCpusArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CPUs enabled to run DPDK IPS engines.
        /// </summary>
        [Input("ipsCpus")]
        public Input<string>? IpsCpus { get; set; }

        /// <summary>
        /// CPUs isolated to run only the DPDK engines with the exception of processes that have affinity explicitly set by either a user configuration or by their implementation.
        /// </summary>
        [Input("isolatedCpus")]
        public Input<string>? IsolatedCpus { get; set; }

        /// <summary>
        /// CPUs enabled to run DPDK RX engines.
        /// </summary>
        [Input("rxCpus")]
        public Input<string>? RxCpus { get; set; }

        /// <summary>
        /// CPUs enabled to run DPDK TX engines.
        /// </summary>
        [Input("txCpus")]
        public Input<string>? TxCpus { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// CPUs enabled to run DPDK VNP engines.
        /// </summary>
        [Input("vnpCpus")]
        public Input<string>? VnpCpus { get; set; }

        public DpdkCpusArgs()
        {
        }
        public static new DpdkCpusArgs Empty => new DpdkCpusArgs();
    }

    public sealed class DpdkCpusState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CPUs enabled to run DPDK IPS engines.
        /// </summary>
        [Input("ipsCpus")]
        public Input<string>? IpsCpus { get; set; }

        /// <summary>
        /// CPUs isolated to run only the DPDK engines with the exception of processes that have affinity explicitly set by either a user configuration or by their implementation.
        /// </summary>
        [Input("isolatedCpus")]
        public Input<string>? IsolatedCpus { get; set; }

        /// <summary>
        /// CPUs enabled to run DPDK RX engines.
        /// </summary>
        [Input("rxCpus")]
        public Input<string>? RxCpus { get; set; }

        /// <summary>
        /// CPUs enabled to run DPDK TX engines.
        /// </summary>
        [Input("txCpus")]
        public Input<string>? TxCpus { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// CPUs enabled to run DPDK VNP engines.
        /// </summary>
        [Input("vnpCpus")]
        public Input<string>? VnpCpus { get; set; }

        public DpdkCpusState()
        {
        }
        public static new DpdkCpusState Empty => new DpdkCpusState();
    }
}
