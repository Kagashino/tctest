// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./cchttpPolicies";
export * from "./cchttpPolicy";
export * from "./cchttpsPolicies";
export * from "./cchttpsPolicy";
export * from "./ccpolicyV2";
export * from "./dayuEipEip";
export * from "./ddosPolicies";
export * from "./ddosPolicy";
export * from "./ddosPolicyAttachment";
export * from "./ddosPolicyAttachments";
export * from "./ddosPolicyCase";
export * from "./ddosPolicyCases";
export * from "./ddosPolicyV2";
export * from "./l4rule";
export * from "./l4ruleV2";
export * from "./l4rules";
export * from "./l4rulesV2";
export * from "./l7rule";
export * from "./l7ruleV2";
export * from "./l7rules";
export * from "./l7rulesV2";

// Import resources to register:
import { CCHttpPolicy } from "./cchttpPolicy";
import { CCHttpsPolicy } from "./cchttpsPolicy";
import { CCPolicyV2 } from "./ccpolicyV2";
import { DayuEipEip } from "./dayuEipEip";
import { DdosPolicy } from "./ddosPolicy";
import { DdosPolicyAttachment } from "./ddosPolicyAttachment";
import { DdosPolicyCase } from "./ddosPolicyCase";
import { DdosPolicyV2 } from "./ddosPolicyV2";
import { L4Rule } from "./l4rule";
import { L4RuleV2 } from "./l4ruleV2";
import { L7Rule } from "./l7rule";
import { L7RuleV2 } from "./l7ruleV2";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tctest:Dayu/cCHttpPolicy:CCHttpPolicy":
                return new CCHttpPolicy(name, <any>undefined, { urn })
            case "tctest:Dayu/cCHttpsPolicy:CCHttpsPolicy":
                return new CCHttpsPolicy(name, <any>undefined, { urn })
            case "tctest:Dayu/cCPolicyV2:CCPolicyV2":
                return new CCPolicyV2(name, <any>undefined, { urn })
            case "tctest:Dayu/dayuEipEip:DayuEipEip":
                return new DayuEipEip(name, <any>undefined, { urn })
            case "tctest:Dayu/ddosPolicy:DdosPolicy":
                return new DdosPolicy(name, <any>undefined, { urn })
            case "tctest:Dayu/ddosPolicyAttachment:DdosPolicyAttachment":
                return new DdosPolicyAttachment(name, <any>undefined, { urn })
            case "tctest:Dayu/ddosPolicyCase:DdosPolicyCase":
                return new DdosPolicyCase(name, <any>undefined, { urn })
            case "tctest:Dayu/ddosPolicyV2:DdosPolicyV2":
                return new DdosPolicyV2(name, <any>undefined, { urn })
            case "tctest:Dayu/l4Rule:L4Rule":
                return new L4Rule(name, <any>undefined, { urn })
            case "tctest:Dayu/l4RuleV2:L4RuleV2":
                return new L4RuleV2(name, <any>undefined, { urn })
            case "tctest:Dayu/l7Rule:L7Rule":
                return new L7Rule(name, <any>undefined, { urn })
            case "tctest:Dayu/l7RuleV2:L7RuleV2":
                return new L7RuleV2(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tctest", "Dayu/cCHttpPolicy", _module)
pulumi.runtime.registerResourceModule("tctest", "Dayu/cCHttpsPolicy", _module)
pulumi.runtime.registerResourceModule("tctest", "Dayu/cCPolicyV2", _module)
pulumi.runtime.registerResourceModule("tctest", "Dayu/dayuEipEip", _module)
pulumi.runtime.registerResourceModule("tctest", "Dayu/ddosPolicy", _module)
pulumi.runtime.registerResourceModule("tctest", "Dayu/ddosPolicyAttachment", _module)
pulumi.runtime.registerResourceModule("tctest", "Dayu/ddosPolicyCase", _module)
pulumi.runtime.registerResourceModule("tctest", "Dayu/ddosPolicyV2", _module)
pulumi.runtime.registerResourceModule("tctest", "Dayu/l4Rule", _module)
pulumi.runtime.registerResourceModule("tctest", "Dayu/l4RuleV2", _module)
pulumi.runtime.registerResourceModule("tctest", "Dayu/l7Rule", _module)
pulumi.runtime.registerResourceModule("tctest", "Dayu/l7RuleV2", _module)
