// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GenericApiArgs, GenericApiState } from "./genericApi";
export type GenericApi = import("./genericApi").GenericApi;
export const GenericApi: typeof import("./genericApi").GenericApi = null as any;
utilities.lazyLoad(exports, ["GenericApi"], () => require("./genericApi"));

export { GetGenericApiArgs, GetGenericApiResult, GetGenericApiOutputArgs } from "./getGenericApi";
export const getGenericApi: typeof import("./getGenericApi").getGenericApi = null as any;
export const getGenericApiOutput: typeof import("./getGenericApi").getGenericApiOutput = null as any;
utilities.lazyLoad(exports, ["getGenericApi","getGenericApiOutput"], () => require("./getGenericApi"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "fortios:json/genericApi:GenericApi":
                return new GenericApi(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("fortios", "json/genericApi", _module)