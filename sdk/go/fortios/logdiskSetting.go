// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Settings for local disk logging.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fortios.NewLogdiskSetting(ctx, "trname", &fortios.LogdiskSettingArgs{
//				Diskfull:                   pulumi.String("overwrite"),
//				DlpArchiveQuota:            pulumi.Int(0),
//				FullFinalWarningThreshold:  pulumi.Int(95),
//				FullFirstWarningThreshold:  pulumi.Int(75),
//				FullSecondWarningThreshold: pulumi.Int(90),
//				IpsArchive:                 pulumi.String("enable"),
//				LogQuota:                   pulumi.Int(0),
//				MaxLogFileSize:             pulumi.Int(20),
//				MaxPolicyPacketCaptureSize: pulumi.Int(100),
//				MaximumLogAge:              pulumi.Int(7),
//				ReportQuota:                pulumi.Int(0),
//				RollDay:                    pulumi.String("sunday"),
//				RollSchedule:               pulumi.String("daily"),
//				RollTime:                   pulumi.String("00:00"),
//				SourceIp:                   pulumi.String("0.0.0.0"),
//				Status:                     pulumi.String("enable"),
//				Upload:                     pulumi.String("disable"),
//				UploadDeleteFiles:          pulumi.String("enable"),
//				UploadDestination:          pulumi.String("ftp-server"),
//				UploadSslConn:              pulumi.String("default"),
//				Uploadip:                   pulumi.String("0.0.0.0"),
//				Uploadport:                 pulumi.Int(21),
//				Uploadsched:                pulumi.String("disable"),
//				Uploadtime:                 pulumi.String("00:00"),
//				Uploadtype:                 pulumi.String("traffic event virus webfilter IPS spamfilter dlp-archive anomaly voip dlp app-ctrl waf netscan gtp dns"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # LogDisk Setting can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/logdiskSetting:LogdiskSetting labelname LogDiskSetting
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/logdiskSetting:LogdiskSetting labelname LogDiskSetting
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type LogdiskSetting struct {
	pulumi.CustomResourceState

	// Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
	Diskfull pulumi.StringOutput `pulumi:"diskfull"`
	// DLP archive quota (MB).
	DlpArchiveQuota pulumi.IntOutput `pulumi:"dlpArchiveQuota"`
	// Log full final warning threshold as a percent (3 - 100, default = 95).
	FullFinalWarningThreshold pulumi.IntOutput `pulumi:"fullFinalWarningThreshold"`
	// Log full first warning threshold as a percent (1 - 98, default = 75).
	FullFirstWarningThreshold pulumi.IntOutput `pulumi:"fullFirstWarningThreshold"`
	// Log full second warning threshold as a percent (2 - 99, default = 90).
	FullSecondWarningThreshold pulumi.IntOutput `pulumi:"fullSecondWarningThreshold"`
	// Specify outgoing interface to reach server.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringOutput `pulumi:"interfaceSelectMethod"`
	// Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
	IpsArchive pulumi.StringOutput `pulumi:"ipsArchive"`
	// Disk log quota (MB).
	LogQuota pulumi.IntOutput `pulumi:"logQuota"`
	// Maximum log file size before rolling (1 - 100 Mbytes).
	MaxLogFileSize pulumi.IntOutput `pulumi:"maxLogFileSize"`
	// Maximum size of policy sniffer in MB (0 means unlimited).
	MaxPolicyPacketCaptureSize pulumi.IntOutput `pulumi:"maxPolicyPacketCaptureSize"`
	// Delete log files older than (days).
	MaximumLogAge pulumi.IntOutput `pulumi:"maximumLogAge"`
	// Report quota (MB).
	ReportQuota pulumi.IntOutput `pulumi:"reportQuota"`
	// Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	RollDay pulumi.StringOutput `pulumi:"rollDay"`
	// Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
	RollSchedule pulumi.StringOutput `pulumi:"rollSchedule"`
	// Time of day to roll the log file (hh:mm).
	RollTime pulumi.StringOutput `pulumi:"rollTime"`
	// Source IP address to use for uploading disk log files.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Enable/disable local disk logging. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
	Upload pulumi.StringOutput `pulumi:"upload"`
	// Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
	UploadDeleteFiles pulumi.StringOutput `pulumi:"uploadDeleteFiles"`
	// The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
	UploadDestination pulumi.StringOutput `pulumi:"uploadDestination"`
	// Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
	UploadSslConn pulumi.StringOutput `pulumi:"uploadSslConn"`
	// The remote directory on the FTP server to upload log files to.
	Uploaddir pulumi.StringOutput `pulumi:"uploaddir"`
	// IP address of the FTP server to upload log files to.
	Uploadip pulumi.StringOutput `pulumi:"uploadip"`
	// Password required to log into the FTP server to upload disk log files.
	Uploadpass pulumi.StringPtrOutput `pulumi:"uploadpass"`
	// TCP port to use for communicating with the FTP server (default = 21).
	Uploadport pulumi.IntOutput `pulumi:"uploadport"`
	// Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
	Uploadsched pulumi.StringOutput `pulumi:"uploadsched"`
	// Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
	Uploadtime pulumi.StringOutput `pulumi:"uploadtime"`
	// Types of log files to upload. Separate multiple entries with a space.
	Uploadtype pulumi.StringOutput `pulumi:"uploadtype"`
	// Username required to log into the FTP server to upload disk log files.
	Uploaduser pulumi.StringOutput `pulumi:"uploaduser"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewLogdiskSetting registers a new resource with the given unique name, arguments, and options.
func NewLogdiskSetting(ctx *pulumi.Context,
	name string, args *LogdiskSettingArgs, opts ...pulumi.ResourceOption) (*LogdiskSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.Uploadpass != nil {
		args.Uploadpass = pulumi.ToSecret(args.Uploadpass).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"uploadpass",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource LogdiskSetting
	err := ctx.RegisterResource("fortios:index/logdiskSetting:LogdiskSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogdiskSetting gets an existing LogdiskSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogdiskSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogdiskSettingState, opts ...pulumi.ResourceOption) (*LogdiskSetting, error) {
	var resource LogdiskSetting
	err := ctx.ReadResource("fortios:index/logdiskSetting:LogdiskSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogdiskSetting resources.
type logdiskSettingState struct {
	// Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
	Diskfull *string `pulumi:"diskfull"`
	// DLP archive quota (MB).
	DlpArchiveQuota *int `pulumi:"dlpArchiveQuota"`
	// Log full final warning threshold as a percent (3 - 100, default = 95).
	FullFinalWarningThreshold *int `pulumi:"fullFinalWarningThreshold"`
	// Log full first warning threshold as a percent (1 - 98, default = 75).
	FullFirstWarningThreshold *int `pulumi:"fullFirstWarningThreshold"`
	// Log full second warning threshold as a percent (2 - 99, default = 90).
	FullSecondWarningThreshold *int `pulumi:"fullSecondWarningThreshold"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
	IpsArchive *string `pulumi:"ipsArchive"`
	// Disk log quota (MB).
	LogQuota *int `pulumi:"logQuota"`
	// Maximum log file size before rolling (1 - 100 Mbytes).
	MaxLogFileSize *int `pulumi:"maxLogFileSize"`
	// Maximum size of policy sniffer in MB (0 means unlimited).
	MaxPolicyPacketCaptureSize *int `pulumi:"maxPolicyPacketCaptureSize"`
	// Delete log files older than (days).
	MaximumLogAge *int `pulumi:"maximumLogAge"`
	// Report quota (MB).
	ReportQuota *int `pulumi:"reportQuota"`
	// Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	RollDay *string `pulumi:"rollDay"`
	// Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
	RollSchedule *string `pulumi:"rollSchedule"`
	// Time of day to roll the log file (hh:mm).
	RollTime *string `pulumi:"rollTime"`
	// Source IP address to use for uploading disk log files.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable local disk logging. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
	Upload *string `pulumi:"upload"`
	// Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
	UploadDeleteFiles *string `pulumi:"uploadDeleteFiles"`
	// The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
	UploadDestination *string `pulumi:"uploadDestination"`
	// Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
	UploadSslConn *string `pulumi:"uploadSslConn"`
	// The remote directory on the FTP server to upload log files to.
	Uploaddir *string `pulumi:"uploaddir"`
	// IP address of the FTP server to upload log files to.
	Uploadip *string `pulumi:"uploadip"`
	// Password required to log into the FTP server to upload disk log files.
	Uploadpass *string `pulumi:"uploadpass"`
	// TCP port to use for communicating with the FTP server (default = 21).
	Uploadport *int `pulumi:"uploadport"`
	// Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
	Uploadsched *string `pulumi:"uploadsched"`
	// Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
	Uploadtime *string `pulumi:"uploadtime"`
	// Types of log files to upload. Separate multiple entries with a space.
	Uploadtype *string `pulumi:"uploadtype"`
	// Username required to log into the FTP server to upload disk log files.
	Uploaduser *string `pulumi:"uploaduser"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type LogdiskSettingState struct {
	// Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
	Diskfull pulumi.StringPtrInput
	// DLP archive quota (MB).
	DlpArchiveQuota pulumi.IntPtrInput
	// Log full final warning threshold as a percent (3 - 100, default = 95).
	FullFinalWarningThreshold pulumi.IntPtrInput
	// Log full first warning threshold as a percent (1 - 98, default = 75).
	FullFirstWarningThreshold pulumi.IntPtrInput
	// Log full second warning threshold as a percent (2 - 99, default = 90).
	FullSecondWarningThreshold pulumi.IntPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
	IpsArchive pulumi.StringPtrInput
	// Disk log quota (MB).
	LogQuota pulumi.IntPtrInput
	// Maximum log file size before rolling (1 - 100 Mbytes).
	MaxLogFileSize pulumi.IntPtrInput
	// Maximum size of policy sniffer in MB (0 means unlimited).
	MaxPolicyPacketCaptureSize pulumi.IntPtrInput
	// Delete log files older than (days).
	MaximumLogAge pulumi.IntPtrInput
	// Report quota (MB).
	ReportQuota pulumi.IntPtrInput
	// Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	RollDay pulumi.StringPtrInput
	// Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
	RollSchedule pulumi.StringPtrInput
	// Time of day to roll the log file (hh:mm).
	RollTime pulumi.StringPtrInput
	// Source IP address to use for uploading disk log files.
	SourceIp pulumi.StringPtrInput
	// Enable/disable local disk logging. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
	Upload pulumi.StringPtrInput
	// Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
	UploadDeleteFiles pulumi.StringPtrInput
	// The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
	UploadDestination pulumi.StringPtrInput
	// Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
	UploadSslConn pulumi.StringPtrInput
	// The remote directory on the FTP server to upload log files to.
	Uploaddir pulumi.StringPtrInput
	// IP address of the FTP server to upload log files to.
	Uploadip pulumi.StringPtrInput
	// Password required to log into the FTP server to upload disk log files.
	Uploadpass pulumi.StringPtrInput
	// TCP port to use for communicating with the FTP server (default = 21).
	Uploadport pulumi.IntPtrInput
	// Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
	Uploadsched pulumi.StringPtrInput
	// Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
	Uploadtime pulumi.StringPtrInput
	// Types of log files to upload. Separate multiple entries with a space.
	Uploadtype pulumi.StringPtrInput
	// Username required to log into the FTP server to upload disk log files.
	Uploaduser pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogdiskSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logdiskSettingState)(nil)).Elem()
}

type logdiskSettingArgs struct {
	// Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
	Diskfull *string `pulumi:"diskfull"`
	// DLP archive quota (MB).
	DlpArchiveQuota *int `pulumi:"dlpArchiveQuota"`
	// Log full final warning threshold as a percent (3 - 100, default = 95).
	FullFinalWarningThreshold *int `pulumi:"fullFinalWarningThreshold"`
	// Log full first warning threshold as a percent (1 - 98, default = 75).
	FullFirstWarningThreshold *int `pulumi:"fullFirstWarningThreshold"`
	// Log full second warning threshold as a percent (2 - 99, default = 90).
	FullSecondWarningThreshold *int `pulumi:"fullSecondWarningThreshold"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
	IpsArchive *string `pulumi:"ipsArchive"`
	// Disk log quota (MB).
	LogQuota *int `pulumi:"logQuota"`
	// Maximum log file size before rolling (1 - 100 Mbytes).
	MaxLogFileSize *int `pulumi:"maxLogFileSize"`
	// Maximum size of policy sniffer in MB (0 means unlimited).
	MaxPolicyPacketCaptureSize *int `pulumi:"maxPolicyPacketCaptureSize"`
	// Delete log files older than (days).
	MaximumLogAge *int `pulumi:"maximumLogAge"`
	// Report quota (MB).
	ReportQuota *int `pulumi:"reportQuota"`
	// Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	RollDay *string `pulumi:"rollDay"`
	// Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
	RollSchedule *string `pulumi:"rollSchedule"`
	// Time of day to roll the log file (hh:mm).
	RollTime *string `pulumi:"rollTime"`
	// Source IP address to use for uploading disk log files.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable local disk logging. Valid values: `enable`, `disable`.
	Status string `pulumi:"status"`
	// Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
	Upload *string `pulumi:"upload"`
	// Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
	UploadDeleteFiles *string `pulumi:"uploadDeleteFiles"`
	// The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
	UploadDestination *string `pulumi:"uploadDestination"`
	// Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
	UploadSslConn *string `pulumi:"uploadSslConn"`
	// The remote directory on the FTP server to upload log files to.
	Uploaddir *string `pulumi:"uploaddir"`
	// IP address of the FTP server to upload log files to.
	Uploadip *string `pulumi:"uploadip"`
	// Password required to log into the FTP server to upload disk log files.
	Uploadpass *string `pulumi:"uploadpass"`
	// TCP port to use for communicating with the FTP server (default = 21).
	Uploadport *int `pulumi:"uploadport"`
	// Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
	Uploadsched *string `pulumi:"uploadsched"`
	// Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
	Uploadtime *string `pulumi:"uploadtime"`
	// Types of log files to upload. Separate multiple entries with a space.
	Uploadtype *string `pulumi:"uploadtype"`
	// Username required to log into the FTP server to upload disk log files.
	Uploaduser *string `pulumi:"uploaduser"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogdiskSetting resource.
type LogdiskSettingArgs struct {
	// Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
	Diskfull pulumi.StringPtrInput
	// DLP archive quota (MB).
	DlpArchiveQuota pulumi.IntPtrInput
	// Log full final warning threshold as a percent (3 - 100, default = 95).
	FullFinalWarningThreshold pulumi.IntPtrInput
	// Log full first warning threshold as a percent (1 - 98, default = 75).
	FullFirstWarningThreshold pulumi.IntPtrInput
	// Log full second warning threshold as a percent (2 - 99, default = 90).
	FullSecondWarningThreshold pulumi.IntPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
	IpsArchive pulumi.StringPtrInput
	// Disk log quota (MB).
	LogQuota pulumi.IntPtrInput
	// Maximum log file size before rolling (1 - 100 Mbytes).
	MaxLogFileSize pulumi.IntPtrInput
	// Maximum size of policy sniffer in MB (0 means unlimited).
	MaxPolicyPacketCaptureSize pulumi.IntPtrInput
	// Delete log files older than (days).
	MaximumLogAge pulumi.IntPtrInput
	// Report quota (MB).
	ReportQuota pulumi.IntPtrInput
	// Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	RollDay pulumi.StringPtrInput
	// Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
	RollSchedule pulumi.StringPtrInput
	// Time of day to roll the log file (hh:mm).
	RollTime pulumi.StringPtrInput
	// Source IP address to use for uploading disk log files.
	SourceIp pulumi.StringPtrInput
	// Enable/disable local disk logging. Valid values: `enable`, `disable`.
	Status pulumi.StringInput
	// Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
	Upload pulumi.StringPtrInput
	// Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
	UploadDeleteFiles pulumi.StringPtrInput
	// The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
	UploadDestination pulumi.StringPtrInput
	// Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
	UploadSslConn pulumi.StringPtrInput
	// The remote directory on the FTP server to upload log files to.
	Uploaddir pulumi.StringPtrInput
	// IP address of the FTP server to upload log files to.
	Uploadip pulumi.StringPtrInput
	// Password required to log into the FTP server to upload disk log files.
	Uploadpass pulumi.StringPtrInput
	// TCP port to use for communicating with the FTP server (default = 21).
	Uploadport pulumi.IntPtrInput
	// Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
	Uploadsched pulumi.StringPtrInput
	// Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
	Uploadtime pulumi.StringPtrInput
	// Types of log files to upload. Separate multiple entries with a space.
	Uploadtype pulumi.StringPtrInput
	// Username required to log into the FTP server to upload disk log files.
	Uploaduser pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogdiskSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logdiskSettingArgs)(nil)).Elem()
}

type LogdiskSettingInput interface {
	pulumi.Input

	ToLogdiskSettingOutput() LogdiskSettingOutput
	ToLogdiskSettingOutputWithContext(ctx context.Context) LogdiskSettingOutput
}

func (*LogdiskSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**LogdiskSetting)(nil)).Elem()
}

func (i *LogdiskSetting) ToLogdiskSettingOutput() LogdiskSettingOutput {
	return i.ToLogdiskSettingOutputWithContext(context.Background())
}

func (i *LogdiskSetting) ToLogdiskSettingOutputWithContext(ctx context.Context) LogdiskSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogdiskSettingOutput)
}

// LogdiskSettingArrayInput is an input type that accepts LogdiskSettingArray and LogdiskSettingArrayOutput values.
// You can construct a concrete instance of `LogdiskSettingArrayInput` via:
//
//	LogdiskSettingArray{ LogdiskSettingArgs{...} }
type LogdiskSettingArrayInput interface {
	pulumi.Input

	ToLogdiskSettingArrayOutput() LogdiskSettingArrayOutput
	ToLogdiskSettingArrayOutputWithContext(context.Context) LogdiskSettingArrayOutput
}

type LogdiskSettingArray []LogdiskSettingInput

func (LogdiskSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogdiskSetting)(nil)).Elem()
}

func (i LogdiskSettingArray) ToLogdiskSettingArrayOutput() LogdiskSettingArrayOutput {
	return i.ToLogdiskSettingArrayOutputWithContext(context.Background())
}

func (i LogdiskSettingArray) ToLogdiskSettingArrayOutputWithContext(ctx context.Context) LogdiskSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogdiskSettingArrayOutput)
}

// LogdiskSettingMapInput is an input type that accepts LogdiskSettingMap and LogdiskSettingMapOutput values.
// You can construct a concrete instance of `LogdiskSettingMapInput` via:
//
//	LogdiskSettingMap{ "key": LogdiskSettingArgs{...} }
type LogdiskSettingMapInput interface {
	pulumi.Input

	ToLogdiskSettingMapOutput() LogdiskSettingMapOutput
	ToLogdiskSettingMapOutputWithContext(context.Context) LogdiskSettingMapOutput
}

type LogdiskSettingMap map[string]LogdiskSettingInput

func (LogdiskSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogdiskSetting)(nil)).Elem()
}

func (i LogdiskSettingMap) ToLogdiskSettingMapOutput() LogdiskSettingMapOutput {
	return i.ToLogdiskSettingMapOutputWithContext(context.Background())
}

func (i LogdiskSettingMap) ToLogdiskSettingMapOutputWithContext(ctx context.Context) LogdiskSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogdiskSettingMapOutput)
}

type LogdiskSettingOutput struct{ *pulumi.OutputState }

func (LogdiskSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogdiskSetting)(nil)).Elem()
}

func (o LogdiskSettingOutput) ToLogdiskSettingOutput() LogdiskSettingOutput {
	return o
}

func (o LogdiskSettingOutput) ToLogdiskSettingOutputWithContext(ctx context.Context) LogdiskSettingOutput {
	return o
}

// Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
func (o LogdiskSettingOutput) Diskfull() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.Diskfull }).(pulumi.StringOutput)
}

// DLP archive quota (MB).
func (o LogdiskSettingOutput) DlpArchiveQuota() pulumi.IntOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.IntOutput { return v.DlpArchiveQuota }).(pulumi.IntOutput)
}

// Log full final warning threshold as a percent (3 - 100, default = 95).
func (o LogdiskSettingOutput) FullFinalWarningThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.IntOutput { return v.FullFinalWarningThreshold }).(pulumi.IntOutput)
}

// Log full first warning threshold as a percent (1 - 98, default = 75).
func (o LogdiskSettingOutput) FullFirstWarningThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.IntOutput { return v.FullFirstWarningThreshold }).(pulumi.IntOutput)
}

// Log full second warning threshold as a percent (2 - 99, default = 90).
func (o LogdiskSettingOutput) FullSecondWarningThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.IntOutput { return v.FullSecondWarningThreshold }).(pulumi.IntOutput)
}

// Specify outgoing interface to reach server.
func (o LogdiskSettingOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
func (o LogdiskSettingOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

// Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
func (o LogdiskSettingOutput) IpsArchive() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.IpsArchive }).(pulumi.StringOutput)
}

// Disk log quota (MB).
func (o LogdiskSettingOutput) LogQuota() pulumi.IntOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.IntOutput { return v.LogQuota }).(pulumi.IntOutput)
}

// Maximum log file size before rolling (1 - 100 Mbytes).
func (o LogdiskSettingOutput) MaxLogFileSize() pulumi.IntOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.IntOutput { return v.MaxLogFileSize }).(pulumi.IntOutput)
}

// Maximum size of policy sniffer in MB (0 means unlimited).
func (o LogdiskSettingOutput) MaxPolicyPacketCaptureSize() pulumi.IntOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.IntOutput { return v.MaxPolicyPacketCaptureSize }).(pulumi.IntOutput)
}

// Delete log files older than (days).
func (o LogdiskSettingOutput) MaximumLogAge() pulumi.IntOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.IntOutput { return v.MaximumLogAge }).(pulumi.IntOutput)
}

// Report quota (MB).
func (o LogdiskSettingOutput) ReportQuota() pulumi.IntOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.IntOutput { return v.ReportQuota }).(pulumi.IntOutput)
}

// Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
func (o LogdiskSettingOutput) RollDay() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.RollDay }).(pulumi.StringOutput)
}

// Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
func (o LogdiskSettingOutput) RollSchedule() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.RollSchedule }).(pulumi.StringOutput)
}

// Time of day to roll the log file (hh:mm).
func (o LogdiskSettingOutput) RollTime() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.RollTime }).(pulumi.StringOutput)
}

// Source IP address to use for uploading disk log files.
func (o LogdiskSettingOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Enable/disable local disk logging. Valid values: `enable`, `disable`.
func (o LogdiskSettingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
func (o LogdiskSettingOutput) Upload() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.Upload }).(pulumi.StringOutput)
}

// Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
func (o LogdiskSettingOutput) UploadDeleteFiles() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.UploadDeleteFiles }).(pulumi.StringOutput)
}

// The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
func (o LogdiskSettingOutput) UploadDestination() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.UploadDestination }).(pulumi.StringOutput)
}

// Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
func (o LogdiskSettingOutput) UploadSslConn() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.UploadSslConn }).(pulumi.StringOutput)
}

// The remote directory on the FTP server to upload log files to.
func (o LogdiskSettingOutput) Uploaddir() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.Uploaddir }).(pulumi.StringOutput)
}

// IP address of the FTP server to upload log files to.
func (o LogdiskSettingOutput) Uploadip() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.Uploadip }).(pulumi.StringOutput)
}

// Password required to log into the FTP server to upload disk log files.
func (o LogdiskSettingOutput) Uploadpass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringPtrOutput { return v.Uploadpass }).(pulumi.StringPtrOutput)
}

// TCP port to use for communicating with the FTP server (default = 21).
func (o LogdiskSettingOutput) Uploadport() pulumi.IntOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.IntOutput { return v.Uploadport }).(pulumi.IntOutput)
}

// Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
func (o LogdiskSettingOutput) Uploadsched() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.Uploadsched }).(pulumi.StringOutput)
}

// Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
func (o LogdiskSettingOutput) Uploadtime() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.Uploadtime }).(pulumi.StringOutput)
}

// Types of log files to upload. Separate multiple entries with a space.
func (o LogdiskSettingOutput) Uploadtype() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.Uploadtype }).(pulumi.StringOutput)
}

// Username required to log into the FTP server to upload disk log files.
func (o LogdiskSettingOutput) Uploaduser() pulumi.StringOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringOutput { return v.Uploaduser }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o LogdiskSettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogdiskSetting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type LogdiskSettingArrayOutput struct{ *pulumi.OutputState }

func (LogdiskSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogdiskSetting)(nil)).Elem()
}

func (o LogdiskSettingArrayOutput) ToLogdiskSettingArrayOutput() LogdiskSettingArrayOutput {
	return o
}

func (o LogdiskSettingArrayOutput) ToLogdiskSettingArrayOutputWithContext(ctx context.Context) LogdiskSettingArrayOutput {
	return o
}

func (o LogdiskSettingArrayOutput) Index(i pulumi.IntInput) LogdiskSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogdiskSetting {
		return vs[0].([]*LogdiskSetting)[vs[1].(int)]
	}).(LogdiskSettingOutput)
}

type LogdiskSettingMapOutput struct{ *pulumi.OutputState }

func (LogdiskSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogdiskSetting)(nil)).Elem()
}

func (o LogdiskSettingMapOutput) ToLogdiskSettingMapOutput() LogdiskSettingMapOutput {
	return o
}

func (o LogdiskSettingMapOutput) ToLogdiskSettingMapOutputWithContext(ctx context.Context) LogdiskSettingMapOutput {
	return o
}

func (o LogdiskSettingMapOutput) MapIndex(k pulumi.StringInput) LogdiskSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogdiskSetting {
		return vs[0].(map[string]*LogdiskSetting)[vs[1].(string)]
	}).(LogdiskSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogdiskSettingInput)(nil)).Elem(), &LogdiskSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogdiskSettingArrayInput)(nil)).Elem(), LogdiskSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogdiskSettingMapInput)(nil)).Elem(), LogdiskSettingMap{})
	pulumi.RegisterOutputType(LogdiskSettingOutput{})
	pulumi.RegisterOutputType(LogdiskSettingArrayOutput{})
	pulumi.RegisterOutputType(LogdiskSettingMapOutput{})
}
