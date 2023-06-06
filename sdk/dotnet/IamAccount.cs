// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.TerraformAwsIam
{
    [TerraformAwsIamResourceType("terraform-aws-iam:index:IamAccount")]
    public partial class IamAccount : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// The AWS Account ID number of the account that owns or contains the calling entity.
        /// </summary>
        [Output("callerIdentityAccountId")]
        public Output<string> CallerIdentityAccountId { get; private set; } = null!;

        /// <summary>
        /// The AWS ARN associated with the calling entity.
        /// </summary>
        [Output("callerIdentityArn")]
        public Output<string> CallerIdentityArn { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the calling entity.
        /// </summary>
        [Output("callerIdentityUserId")]
        public Output<string> CallerIdentityUserId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether passwords in the account expire. Returns true if max_password_age contains a value greater than 0. Returns false if it is 0 or not present.
        /// </summary>
        [Output("iamAccountPasswordPolicyExpirePasswords")]
        public Output<bool> IamAccountPasswordPolicyExpirePasswords { get; private set; } = null!;


        /// <summary>
        /// Create a IamAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamAccount(string name, IamAccountArgs args, ComponentResourceOptions? options = null)
            : base("terraform-aws-iam:index:IamAccount", name, args ?? new IamAccountArgs(), MakeResourceOptions(options, ""), remote: true)
        {
        }

        private static ComponentResourceOptions MakeResourceOptions(ComponentResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ComponentResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = ComponentResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class IamAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS IAM account alias for this account.
        /// </summary>
        [Input("accountAlias", required: true)]
        public Input<string> AccountAlias { get; set; } = null!;

        /// <summary>
        /// Whether to allow users to change their own password.
        /// </summary>
        [Input("allowUsersToChangePassword")]
        public Input<bool>? AllowUsersToChangePassword { get; set; }

        /// <summary>
        /// Whether to create AWS IAM account password policy.
        /// </summary>
        [Input("createAccountPasswordPolicy")]
        public bool? CreateAccountPasswordPolicy { get; set; }

        /// <summary>
        /// Whether to get AWS account ID, User ID, and ARN in which Terraform is authorized.
        /// </summary>
        [Input("getCallerIdentity")]
        public bool? GetCallerIdentity { get; set; }

        /// <summary>
        /// Whether users are prevented from setting a new password after their password has expired (i.e. require administrator reset).
        /// </summary>
        [Input("hardExpiry")]
        public Input<bool>? HardExpiry { get; set; }

        /// <summary>
        /// The number of days that an user password is valid.
        /// </summary>
        [Input("maxPasswordAge")]
        public Input<int>? MaxPasswordAge { get; set; }

        /// <summary>
        /// Minimum length to require for user passwords.
        /// </summary>
        [Input("minimumPasswordLength")]
        public Input<int>? MinimumPasswordLength { get; set; }

        /// <summary>
        /// The number of previous passwords that users are prevented from reusing.
        /// </summary>
        [Input("passwordReusePrevention")]
        public Input<int>? PasswordReusePrevention { get; set; }

        /// <summary>
        /// Whether to require lowercase characters for user passwords.
        /// </summary>
        [Input("requireLowercaseCharacters")]
        public Input<bool>? RequireLowercaseCharacters { get; set; }

        /// <summary>
        /// Whether to require numbers for user passwords.
        /// </summary>
        [Input("requireNumbers")]
        public Input<bool>? RequireNumbers { get; set; }

        /// <summary>
        /// Whether to require symbols for user passwords.
        /// </summary>
        [Input("requireSymbols")]
        public Input<bool>? RequireSymbols { get; set; }

        /// <summary>
        /// Whether to require uppercase characters for user passwords.
        /// </summary>
        [Input("requireUppercaseCharacters")]
        public Input<bool>? RequireUppercaseCharacters { get; set; }

        public IamAccountArgs()
        {
        }
        public static new IamAccountArgs Empty => new IamAccountArgs();
    }
}
