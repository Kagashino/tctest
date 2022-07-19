// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./provider";

// Export sub-modules:
import * as address from "./address";
import * as api from "./api";
import * as apigateway from "./apigateway";
import * as as from "./as";
import * as audit from "./audit";
import * as audits from "./audits";
import * as availability from "./availability";
import * as cam from "./cam";
import * as cbs from "./cbs";
import * as ccn from "./ccn";
import * as cdh from "./cdh";
import * as cdn from "./cdn";
import * as cfs from "./cfs";
import * as ckafka from "./ckafka";
import * as clb from "./clb";
import * as cloud from "./cloud";
import * as cls from "./cls";
import * as config from "./config";
import * as container from "./container";
import * as cos from "./cos";
import * as cvm from "./cvm";
import * as cynosdb from "./cynosdb";
import * as dayu from "./dayu";
import * as dc from "./dc";
import * as dcx from "./dcx";
import * as dnat from "./dnat";
import * as dnats from "./dnats";
import * as dnspod from "./dnspod";
import * as eip from "./eip";
import * as eips from "./eips";
import * as eks from "./eks";
import * as elasticsearch from "./elasticsearch";
import * as emr from "./emr";
import * as eni from "./eni";
import * as enis from "./enis";
import * as gaap from "./gaap";
import * as ha from "./ha";
import * as image from "./image";
import * as instances from "./instances";
import * as key from "./key";
import * as kms from "./kms";
import * as mongodb from "./mongodb";
import * as monitor from "./monitor";
import * as mysql from "./mysql";
import * as nat from "./nat";
import * as nats from "./nats";
import * as placement from "./placement";
import * as postgresql from "./postgresql";
import * as privatedns from "./privatedns";
import * as protocol from "./protocol";
import * as redis from "./redis";
import * as reserved from "./reserved";
import * as route from "./route";
import * as scf from "./scf";
import * as security from "./security";
import * as sqlserver from "./sqlserver";
import * as ssl from "./ssl";
import * as ssm from "./ssm";
import * as subnet from "./subnet";
import * as tcaplus from "./tcaplus";
import * as tcr from "./tcr";
import * as tdmq from "./tdmq";
import * as tke from "./tke";
import * as types from "./types";
import * as user from "./user";
import * as vod from "./vod";
import * as vpc from "./vpc";
import * as vpn from "./vpn";

export {
    address,
    api,
    apigateway,
    as,
    audit,
    audits,
    availability,
    cam,
    cbs,
    ccn,
    cdh,
    cdn,
    cfs,
    ckafka,
    clb,
    cloud,
    cls,
    config,
    container,
    cos,
    cvm,
    cynosdb,
    dayu,
    dc,
    dcx,
    dnat,
    dnats,
    dnspod,
    eip,
    eips,
    eks,
    elasticsearch,
    emr,
    eni,
    enis,
    gaap,
    ha,
    image,
    instances,
    key,
    kms,
    mongodb,
    monitor,
    mysql,
    nat,
    nats,
    placement,
    postgresql,
    privatedns,
    protocol,
    redis,
    reserved,
    route,
    scf,
    security,
    sqlserver,
    ssl,
    ssm,
    subnet,
    tcaplus,
    tcr,
    tdmq,
    tke,
    types,
    user,
    vod,
    vpc,
    vpn,
};

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("tctest", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:tctest") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
