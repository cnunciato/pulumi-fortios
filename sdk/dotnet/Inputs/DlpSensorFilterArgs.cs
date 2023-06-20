// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class DlpSensorFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to take with content that this DLP sensor matches. Valid values: `allow`, `log-only`, `block`, `quarantine-ip`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Enable/disable DLP archiving. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("archive")]
        public Input<string>? Archive { get; set; }

        /// <summary>
        /// Enter a company identifier watermark to match. Only watermarks that your company has placed on the files are matched.
        /// </summary>
        [Input("companyIdentifier")]
        public Input<string>? CompanyIdentifier { get; set; }

        /// <summary>
        /// Quarantine duration in days, hours, minutes format (dddhhmm).
        /// </summary>
        [Input("expiry")]
        public Input<string>? Expiry { get; set; }

        /// <summary>
        /// Match files this size or larger (0 - 4294967295 kbytes).
        /// </summary>
        [Input("fileSize")]
        public Input<int>? FileSize { get; set; }

        /// <summary>
        /// Select the number of a DLP file pattern table to match.
        /// </summary>
        [Input("fileType")]
        public Input<int>? FileType { get; set; }

        /// <summary>
        /// Select the type of content to match.
        /// </summary>
        [Input("filterBy")]
        public Input<string>? FilterBy { get; set; }

        [Input("fpSensitivities")]
        private InputList<Inputs.DlpSensorFilterFpSensitivityArgs>? _fpSensitivities;

        /// <summary>
        /// Select a DLP file pattern sensitivity to match. The structure of `fp_sensitivity` block is documented below.
        /// </summary>
        public InputList<Inputs.DlpSensorFilterFpSensitivityArgs> FpSensitivities
        {
            get => _fpSensitivities ?? (_fpSensitivities = new InputList<Inputs.DlpSensorFilterFpSensitivityArgs>());
            set => _fpSensitivities = value;
        }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Percentage of fingerprints in the fingerprint databases designated with the selected fp-sensitivity to match.
        /// </summary>
        [Input("matchPercentage")]
        public Input<int>? MatchPercentage { get; set; }

        /// <summary>
        /// Filter name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Check messages or files over one or more of these protocols.
        /// </summary>
        [Input("proto")]
        public Input<string>? Proto { get; set; }

        /// <summary>
        /// Enter a regular expression to match (max. 255 characters).
        /// </summary>
        [Input("regexp")]
        public Input<string>? Regexp { get; set; }

        [Input("sensitivities")]
        private InputList<Inputs.DlpSensorFilterSensitivityArgs>? _sensitivities;

        /// <summary>
        /// Select a DLP file pattern sensitivity to match. The structure of `sensitivity` block is documented below.
        /// </summary>
        public InputList<Inputs.DlpSensorFilterSensitivityArgs> Sensitivities
        {
            get => _sensitivities ?? (_sensitivities = new InputList<Inputs.DlpSensorFilterSensitivityArgs>());
            set => _sensitivities = value;
        }

        /// <summary>
        /// Select the severity or threat level that matches this filter. Valid values: `info`, `low`, `medium`, `high`, `critical`.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// Select whether to check the content of messages (an email message) or files (downloaded files or email attachments).  Valid values: `file`, `message`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DlpSensorFilterArgs()
        {
        }
        public static new DlpSensorFilterArgs Empty => new DlpSensorFilterArgs();
    }
}
