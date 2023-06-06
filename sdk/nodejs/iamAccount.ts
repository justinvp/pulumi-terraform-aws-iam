// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class IamAccount extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'terraform-aws-iam:index:IamAccount';

    /**
     * Returns true if the given object is an instance of IamAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IamAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IamAccount.__pulumiType;
    }

    /**
     * The AWS Account ID number of the account that owns or contains the calling entity.
     */
    public /*out*/ readonly callerIdentityAccountId!: pulumi.Output<string>;
    /**
     * The AWS ARN associated with the calling entity.
     */
    public /*out*/ readonly callerIdentityArn!: pulumi.Output<string>;
    /**
     * The unique identifier of the calling entity.
     */
    public /*out*/ readonly callerIdentityUserId!: pulumi.Output<string>;
    /**
     * Indicates whether passwords in the account expire. Returns true if max_password_age contains a value greater than 0. Returns false if it is 0 or not present.
     */
    public /*out*/ readonly iamAccountPasswordPolicyExpirePasswords!: pulumi.Output<boolean>;

    /**
     * Create a IamAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IamAccountArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.accountAlias === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountAlias'");
            }
            resourceInputs["accountAlias"] = args ? args.accountAlias : undefined;
            resourceInputs["allowUsersToChangePassword"] = args ? args.allowUsersToChangePassword : undefined;
            resourceInputs["createAccountPasswordPolicy"] = args ? args.createAccountPasswordPolicy : undefined;
            resourceInputs["getCallerIdentity"] = args ? args.getCallerIdentity : undefined;
            resourceInputs["hardExpiry"] = args ? args.hardExpiry : undefined;
            resourceInputs["maxPasswordAge"] = args ? args.maxPasswordAge : undefined;
            resourceInputs["minimumPasswordLength"] = args ? args.minimumPasswordLength : undefined;
            resourceInputs["passwordReusePrevention"] = args ? args.passwordReusePrevention : undefined;
            resourceInputs["requireLowercaseCharacters"] = args ? args.requireLowercaseCharacters : undefined;
            resourceInputs["requireNumbers"] = args ? args.requireNumbers : undefined;
            resourceInputs["requireSymbols"] = args ? args.requireSymbols : undefined;
            resourceInputs["requireUppercaseCharacters"] = args ? args.requireUppercaseCharacters : undefined;
            resourceInputs["callerIdentityAccountId"] = undefined /*out*/;
            resourceInputs["callerIdentityArn"] = undefined /*out*/;
            resourceInputs["callerIdentityUserId"] = undefined /*out*/;
            resourceInputs["iamAccountPasswordPolicyExpirePasswords"] = undefined /*out*/;
        } else {
            resourceInputs["callerIdentityAccountId"] = undefined /*out*/;
            resourceInputs["callerIdentityArn"] = undefined /*out*/;
            resourceInputs["callerIdentityUserId"] = undefined /*out*/;
            resourceInputs["iamAccountPasswordPolicyExpirePasswords"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IamAccount.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a IamAccount resource.
 */
export interface IamAccountArgs {
    /**
     * AWS IAM account alias for this account.
     */
    accountAlias: pulumi.Input<string>;
    /**
     * Whether to allow users to change their own password.
     */
    allowUsersToChangePassword?: pulumi.Input<boolean>;
    /**
     * Whether to create AWS IAM account password policy.
     */
    createAccountPasswordPolicy?: pulumi.Input<boolean>;
    /**
     * Whether to get AWS account ID, User ID, and ARN in which Terraform is authorized.
     */
    getCallerIdentity?: pulumi.Input<boolean>;
    /**
     * Whether users are prevented from setting a new password after their password has expired (i.e. require administrator reset).
     */
    hardExpiry?: pulumi.Input<boolean>;
    /**
     * The number of days that an user password is valid.
     */
    maxPasswordAge?: pulumi.Input<number>;
    /**
     * Minimum length to require for user passwords.
     */
    minimumPasswordLength?: pulumi.Input<number>;
    /**
     * The number of previous passwords that users are prevented from reusing.
     */
    passwordReusePrevention?: pulumi.Input<number>;
    /**
     * Whether to require lowercase characters for user passwords.
     */
    requireLowercaseCharacters?: pulumi.Input<boolean>;
    /**
     * Whether to require numbers for user passwords.
     */
    requireNumbers?: pulumi.Input<boolean>;
    /**
     * Whether to require symbols for user passwords.
     */
    requireSymbols?: pulumi.Input<boolean>;
    /**
     * Whether to require uppercase characters for user passwords.
     */
    requireUppercaseCharacters?: pulumi.Input<boolean>;
}
