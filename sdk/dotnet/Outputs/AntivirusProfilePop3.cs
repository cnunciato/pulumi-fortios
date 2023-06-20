// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class AntivirusProfilePop3
    {
        /// <summary>
        /// Select the archive types to block.
        /// </summary>
        public readonly string? ArchiveBlock;
        /// <summary>
        /// Select the archive types to log.
        /// </summary>
        public readonly string? ArchiveLog;
        /// <summary>
        /// Enable AntiVirus scan service. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        public readonly string? AvScan;
        /// <summary>
        /// AV Content Disarm and Reconstruction settings. The structure of `content_disarm` block is documented below.
        /// </summary>
        public readonly string? ContentDisarm;
        /// <summary>
        /// Enable/disable the virus emulator. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Emulator;
        /// <summary>
        /// Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default`, `virus`.
        /// </summary>
        public readonly string? Executables;
        /// <summary>
        /// One or more external malware block lists. The structure of `external_blocklist` block is documented below.
        /// </summary>
        public readonly string? ExternalBlocklist;
        /// <summary>
        /// Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        public readonly string? Fortiai;
        /// <summary>
        /// Enable scanning of files by FortiNDR. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        public readonly string? Fortindr;
        /// <summary>
        /// Enable scanning of files by FortiSandbox. Valid values: `disable`, `block`, `monitor`.
        /// </summary>
        public readonly string? Fortisandbox;
        /// <summary>
        /// Enable/disable HTTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `avmonitor`, `quarantine`.
        /// </summary>
        public readonly string? Options;
        /// <summary>
        /// Configure Virus Outbreak Prevention settings. The structure of `outbreak_prevention` block is documented below.
        /// </summary>
        public readonly string? OutbreakPrevention;
        /// <summary>
        /// Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Quarantine;

        [OutputConstructor]
        private AntivirusProfilePop3(
            string? archiveBlock,

            string? archiveLog,

            string? avScan,

            string? contentDisarm,

            string? emulator,

            string? executables,

            string? externalBlocklist,

            string? fortiai,

            string? fortindr,

            string? fortisandbox,

            string? options,

            string? outbreakPrevention,

            string? quarantine)
        {
            ArchiveBlock = archiveBlock;
            ArchiveLog = archiveLog;
            AvScan = avScan;
            ContentDisarm = contentDisarm;
            Emulator = emulator;
            Executables = executables;
            ExternalBlocklist = externalBlocklist;
            Fortiai = fortiai;
            Fortindr = fortindr;
            Fortisandbox = fortisandbox;
            Options = options;
            OutbreakPrevention = outbreakPrevention;
            Quarantine = quarantine;
        }
    }
}
