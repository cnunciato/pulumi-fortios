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

    public sealed class WebfilterProfileFileFilterEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action taken for matched file. Valid values: `log`, `block`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Match files transmitted in the session's originating or reply direction. Valid values: `incoming`, `outgoing`, `any`.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        [Input("fileTypes")]
        private InputList<Inputs.WebfilterProfileFileFilterEntryFileTypeArgs>? _fileTypes;

        /// <summary>
        /// Select file type. The structure of `file_type` block is documented below.
        /// </summary>
        public InputList<Inputs.WebfilterProfileFileFilterEntryFileTypeArgs> FileTypes
        {
            get => _fileTypes ?? (_fileTypes = new InputList<Inputs.WebfilterProfileFileFilterEntryFileTypeArgs>());
            set => _fileTypes = value;
        }

        /// <summary>
        /// Add a file filter.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// Match password-protected files. Valid values: `yes`, `any`.
        /// </summary>
        [Input("passwordProtected")]
        public Input<string>? PasswordProtected { get; set; }

        /// <summary>
        /// Protocols to apply with. Valid values: `http`, `ftp`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public WebfilterProfileFileFilterEntryArgs()
        {
        }
        public static new WebfilterProfileFileFilterEntryArgs Empty => new WebfilterProfileFileFilterEntryArgs();
    }
}
