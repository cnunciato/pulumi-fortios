// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetPushupdateArgs, GetPushupdateResult, GetPushupdateOutputArgs } from "./getPushupdate";
export const getPushupdate: typeof import("./getPushupdate").getPushupdate = null as any;
export const getPushupdateOutput: typeof import("./getPushupdate").getPushupdateOutput = null as any;
utilities.lazyLoad(exports, ["getPushupdate","getPushupdateOutput"], () => require("./getPushupdate"));

export { GetScheduleArgs, GetScheduleResult, GetScheduleOutputArgs } from "./getSchedule";
export const getSchedule: typeof import("./getSchedule").getSchedule = null as any;
export const getScheduleOutput: typeof import("./getSchedule").getScheduleOutput = null as any;
utilities.lazyLoad(exports, ["getSchedule","getScheduleOutput"], () => require("./getSchedule"));

export { GetTunnelingArgs, GetTunnelingResult, GetTunnelingOutputArgs } from "./getTunneling";
export const getTunneling: typeof import("./getTunneling").getTunneling = null as any;
export const getTunnelingOutput: typeof import("./getTunneling").getTunnelingOutput = null as any;
utilities.lazyLoad(exports, ["getTunneling","getTunnelingOutput"], () => require("./getTunneling"));

export { PushupdateArgs, PushupdateState } from "./pushupdate";
export type Pushupdate = import("./pushupdate").Pushupdate;
export const Pushupdate: typeof import("./pushupdate").Pushupdate = null as any;
utilities.lazyLoad(exports, ["Pushupdate"], () => require("./pushupdate"));

export { ScheduleArgs, ScheduleState } from "./schedule";
export type Schedule = import("./schedule").Schedule;
export const Schedule: typeof import("./schedule").Schedule = null as any;
utilities.lazyLoad(exports, ["Schedule"], () => require("./schedule"));

export { TunnelingArgs, TunnelingState } from "./tunneling";
export type Tunneling = import("./tunneling").Tunneling;
export const Tunneling: typeof import("./tunneling").Tunneling = null as any;
utilities.lazyLoad(exports, ["Tunneling"], () => require("./tunneling"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "fortios:system/autoupdate/pushupdate:Pushupdate":
                return new Pushupdate(name, <any>undefined, { urn })
            case "fortios:system/autoupdate/schedule:Schedule":
                return new Schedule(name, <any>undefined, { urn })
            case "fortios:system/autoupdate/tunneling:Tunneling":
                return new Tunneling(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("fortios", "system/autoupdate/pushupdate", _module)
pulumi.runtime.registerResourceModule("fortios", "system/autoupdate/schedule", _module)
pulumi.runtime.registerResourceModule("fortios", "system/autoupdate/tunneling", _module)