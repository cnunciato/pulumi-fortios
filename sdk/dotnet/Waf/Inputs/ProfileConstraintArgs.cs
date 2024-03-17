// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Waf.Inputs
{

    public sealed class ProfileConstraintArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// HTTP content length in request. The structure of `content_length` block is documented below.
        /// </summary>
        [Input("contentLength")]
        public Input<Inputs.ProfileConstraintContentLengthArgs>? ContentLength { get; set; }

        [Input("exceptions")]
        private InputList<Inputs.ProfileConstraintExceptionArgs>? _exceptions;

        /// <summary>
        /// HTTP constraint exception. The structure of `exception` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileConstraintExceptionArgs> Exceptions
        {
            get => _exceptions ?? (_exceptions = new InputList<Inputs.ProfileConstraintExceptionArgs>());
            set => _exceptions = value;
        }

        /// <summary>
        /// HTTP header length in request. The structure of `header_length` block is documented below.
        /// </summary>
        [Input("headerLength")]
        public Input<Inputs.ProfileConstraintHeaderLengthArgs>? HeaderLength { get; set; }

        /// <summary>
        /// Enable/disable hostname check. The structure of `hostname` block is documented below.
        /// </summary>
        [Input("hostname")]
        public Input<Inputs.ProfileConstraintHostnameArgs>? Hostname { get; set; }

        /// <summary>
        /// HTTP line length in request. The structure of `line_length` block is documented below.
        /// </summary>
        [Input("lineLength")]
        public Input<Inputs.ProfileConstraintLineLengthArgs>? LineLength { get; set; }

        /// <summary>
        /// Enable/disable malformed HTTP request check. The structure of `malformed` block is documented below.
        /// </summary>
        [Input("malformed")]
        public Input<Inputs.ProfileConstraintMalformedArgs>? Malformed { get; set; }

        /// <summary>
        /// Maximum number of cookies in HTTP request. The structure of `max_cookie` block is documented below.
        /// </summary>
        [Input("maxCookie")]
        public Input<Inputs.ProfileConstraintMaxCookieArgs>? MaxCookie { get; set; }

        /// <summary>
        /// Maximum number of HTTP header line. The structure of `max_header_line` block is documented below.
        /// </summary>
        [Input("maxHeaderLine")]
        public Input<Inputs.ProfileConstraintMaxHeaderLineArgs>? MaxHeaderLine { get; set; }

        /// <summary>
        /// Maximum number of range segments in HTTP range line. The structure of `max_range_segment` block is documented below.
        /// </summary>
        [Input("maxRangeSegment")]
        public Input<Inputs.ProfileConstraintMaxRangeSegmentArgs>? MaxRangeSegment { get; set; }

        /// <summary>
        /// Maximum number of parameters in URL. The structure of `max_url_param` block is documented below.
        /// </summary>
        [Input("maxUrlParam")]
        public Input<Inputs.ProfileConstraintMaxUrlParamArgs>? MaxUrlParam { get; set; }

        /// <summary>
        /// Enable/disable HTTP method check. The structure of `method` block is documented below.
        /// </summary>
        [Input("method")]
        public Input<Inputs.ProfileConstraintMethodArgs>? Method { get; set; }

        /// <summary>
        /// Maximum length of parameter in URL, HTTP POST request or HTTP body. The structure of `param_length` block is documented below.
        /// </summary>
        [Input("paramLength")]
        public Input<Inputs.ProfileConstraintParamLengthArgs>? ParamLength { get; set; }

        /// <summary>
        /// Maximum length of parameter in URL. The structure of `url_param_length` block is documented below.
        /// </summary>
        [Input("urlParamLength")]
        public Input<Inputs.ProfileConstraintUrlParamLengthArgs>? UrlParamLength { get; set; }

        /// <summary>
        /// Enable/disable HTTP version check. The structure of `version` block is documented below.
        /// </summary>
        [Input("version")]
        public Input<Inputs.ProfileConstraintVersionArgs>? Version { get; set; }

        public ProfileConstraintArgs()
        {
        }
        public static new ProfileConstraintArgs Empty => new ProfileConstraintArgs();
    }
}