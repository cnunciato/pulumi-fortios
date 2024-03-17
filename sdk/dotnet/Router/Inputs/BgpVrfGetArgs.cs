// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Inputs
{

    public sealed class BgpVrfGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("exportRts")]
        private InputList<Inputs.BgpVrfExportRtGetArgs>? _exportRts;

        /// <summary>
        /// List of export route target. The structure of `export_rt` block is documented below.
        /// </summary>
        public InputList<Inputs.BgpVrfExportRtGetArgs> ExportRts
        {
            get => _exportRts ?? (_exportRts = new InputList<Inputs.BgpVrfExportRtGetArgs>());
            set => _exportRts = value;
        }

        /// <summary>
        /// Import route map.
        /// </summary>
        [Input("importRouteMap")]
        public Input<string>? ImportRouteMap { get; set; }

        [Input("importRts")]
        private InputList<Inputs.BgpVrfImportRtGetArgs>? _importRts;

        /// <summary>
        /// List of import route target. The structure of `import_rt` block is documented below.
        /// </summary>
        public InputList<Inputs.BgpVrfImportRtGetArgs> ImportRts
        {
            get => _importRts ?? (_importRts = new InputList<Inputs.BgpVrfImportRtGetArgs>());
            set => _importRts = value;
        }

        [Input("leakTargets")]
        private InputList<Inputs.BgpVrfLeakTargetGetArgs>? _leakTargets;

        /// <summary>
        /// Target VRF table. The structure of `leak_target` block is documented below.
        /// </summary>
        public InputList<Inputs.BgpVrfLeakTargetGetArgs> LeakTargets
        {
            get => _leakTargets ?? (_leakTargets = new InputList<Inputs.BgpVrfLeakTargetGetArgs>());
            set => _leakTargets = value;
        }

        /// <summary>
        /// Route Distinguisher: AA|AA:NN.
        /// </summary>
        [Input("rd")]
        public Input<string>? Rd { get; set; }

        /// <summary>
        /// VRF role. Valid values: `standalone`, `ce`, `pe`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Origin VRF ID (0 - 63).
        /// </summary>
        [Input("vrf")]
        public Input<string>? Vrf { get; set; }

        public BgpVrfGetArgs()
        {
        }
        public static new BgpVrfGetArgs Empty => new BgpVrfGetArgs();
    }
}