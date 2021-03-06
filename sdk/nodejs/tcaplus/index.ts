// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./cluster";
export * from "./clusters";
export * from "./idl";
export * from "./idls";
export * from "./table";
export * from "./tableGroup";
export * from "./tableGroups";
export * from "./tables";

// Import resources to register:
import { Cluster } from "./cluster";
import { Idl } from "./idl";
import { Table } from "./table";
import { TableGroup } from "./tableGroup";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tctest:Tcaplus/cluster:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "tctest:Tcaplus/idl:Idl":
                return new Idl(name, <any>undefined, { urn })
            case "tctest:Tcaplus/table:Table":
                return new Table(name, <any>undefined, { urn })
            case "tctest:Tcaplus/tableGroup:TableGroup":
                return new TableGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tctest", "Tcaplus/cluster", _module)
pulumi.runtime.registerResourceModule("tctest", "Tcaplus/idl", _module)
pulumi.runtime.registerResourceModule("tctest", "Tcaplus/table", _module)
pulumi.runtime.registerResourceModule("tctest", "Tcaplus/tableGroup", _module)
