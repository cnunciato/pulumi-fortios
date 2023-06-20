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
    public static class GetSystemAccprofile
    {
        /// <summary>
        /// Use this data source to get information on an fortios system accprofile
        /// </summary>
        public static Task<GetSystemAccprofileResult> InvokeAsync(GetSystemAccprofileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemAccprofileResult>("fortios:index/getSystemAccprofile:getSystemAccprofile", args ?? new GetSystemAccprofileArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system accprofile
        /// </summary>
        public static Output<GetSystemAccprofileResult> Invoke(GetSystemAccprofileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemAccprofileResult>("fortios:index/getSystemAccprofile:getSystemAccprofile", args ?? new GetSystemAccprofileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemAccprofileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system accprofile.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemAccprofileArgs()
        {
        }
        public static new GetSystemAccprofileArgs Empty => new GetSystemAccprofileArgs();
    }

    public sealed class GetSystemAccprofileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system accprofile.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemAccprofileInvokeArgs()
        {
        }
        public static new GetSystemAccprofileInvokeArgs Empty => new GetSystemAccprofileInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemAccprofileResult
    {
        /// <summary>
        /// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
        /// </summary>
        public readonly int Admintimeout;
        /// <summary>
        /// Enable/disable overriding the global administrator idle timeout.
        /// </summary>
        public readonly string AdmintimeoutOverride;
        /// <summary>
        /// Administrator access to Users and Devices.
        /// </summary>
        public readonly string Authgrp;
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comments;
        /// <summary>
        /// FortiView.
        /// </summary>
        public readonly string Ftviewgrp;
        /// <summary>
        /// Administrator access to the Firewall configuration.
        /// </summary>
        public readonly string Fwgrp;
        /// <summary>
        /// Custom firewall permission. The structure of `fwgrp_permission` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemAccprofileFwgrpPermissionResult> FwgrpPermissions;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Administrator access to Logging and Reporting including viewing log messages.
        /// </summary>
        public readonly string Loggrp;
        /// <summary>
        /// Custom Log &amp; Report permission. The structure of `loggrp_permission` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemAccprofileLoggrpPermissionResult> LoggrpPermissions;
        /// <summary>
        /// Profile name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Network Configuration.
        /// </summary>
        public readonly string Netgrp;
        /// <summary>
        /// Custom network permission. The structure of `netgrp_permission` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemAccprofileNetgrpPermissionResult> NetgrpPermissions;
        /// <summary>
        /// Scope of admin access: global or specific VDOM(s).
        /// </summary>
        public readonly string Scope;
        /// <summary>
        /// Security Fabric.
        /// </summary>
        public readonly string Secfabgrp;
        /// <summary>
        /// System Configuration.
        /// </summary>
        public readonly string Sysgrp;
        /// <summary>
        /// Custom system permission. The structure of `sysgrp_permission` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemAccprofileSysgrpPermissionResult> SysgrpPermissions;
        /// <summary>
        /// Enable/disable permission to run system diagnostic commands.
        /// </summary>
        public readonly string SystemDiagnostics;
        /// <summary>
        /// Enable/disable permission to execute SSH commands.
        /// </summary>
        public readonly string SystemExecuteSsh;
        /// <summary>
        /// Enable/disable permission to execute TELNET commands.
        /// </summary>
        public readonly string SystemExecuteTelnet;
        /// <summary>
        /// Administrator access to Security Profiles.
        /// </summary>
        public readonly string Utmgrp;
        /// <summary>
        /// Custom Security Profile permissions. The structure of `utmgrp_permission` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemAccprofileUtmgrpPermissionResult> UtmgrpPermissions;
        public readonly string? Vdomparam;
        /// <summary>
        /// Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
        /// </summary>
        public readonly string Vpngrp;
        /// <summary>
        /// Administrator access to WAN Opt &amp; Cache.
        /// </summary>
        public readonly string Wanoptgrp;
        /// <summary>
        /// Administrator access to the WiFi controller and Switch controller.
        /// </summary>
        public readonly string Wifi;

        [OutputConstructor]
        private GetSystemAccprofileResult(
            int admintimeout,

            string admintimeoutOverride,

            string authgrp,

            string comments,

            string ftviewgrp,

            string fwgrp,

            ImmutableArray<Outputs.GetSystemAccprofileFwgrpPermissionResult> fwgrpPermissions,

            string id,

            string loggrp,

            ImmutableArray<Outputs.GetSystemAccprofileLoggrpPermissionResult> loggrpPermissions,

            string name,

            string netgrp,

            ImmutableArray<Outputs.GetSystemAccprofileNetgrpPermissionResult> netgrpPermissions,

            string scope,

            string secfabgrp,

            string sysgrp,

            ImmutableArray<Outputs.GetSystemAccprofileSysgrpPermissionResult> sysgrpPermissions,

            string systemDiagnostics,

            string systemExecuteSsh,

            string systemExecuteTelnet,

            string utmgrp,

            ImmutableArray<Outputs.GetSystemAccprofileUtmgrpPermissionResult> utmgrpPermissions,

            string? vdomparam,

            string vpngrp,

            string wanoptgrp,

            string wifi)
        {
            Admintimeout = admintimeout;
            AdmintimeoutOverride = admintimeoutOverride;
            Authgrp = authgrp;
            Comments = comments;
            Ftviewgrp = ftviewgrp;
            Fwgrp = fwgrp;
            FwgrpPermissions = fwgrpPermissions;
            Id = id;
            Loggrp = loggrp;
            LoggrpPermissions = loggrpPermissions;
            Name = name;
            Netgrp = netgrp;
            NetgrpPermissions = netgrpPermissions;
            Scope = scope;
            Secfabgrp = secfabgrp;
            Sysgrp = sysgrp;
            SysgrpPermissions = sysgrpPermissions;
            SystemDiagnostics = systemDiagnostics;
            SystemExecuteSsh = systemExecuteSsh;
            SystemExecuteTelnet = systemExecuteTelnet;
            Utmgrp = utmgrp;
            UtmgrpPermissions = utmgrpPermissions;
            Vdomparam = vdomparam;
            Vpngrp = vpngrp;
            Wanoptgrp = wanoptgrp;
            Wifi = wifi;
        }
    }
}
