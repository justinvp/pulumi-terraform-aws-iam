package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Fixed types and added `pulumi:` tags to fields, with optional fields marked as such.
type IamAccountArgs struct {
	GetCallerIdentity           bool               `pulumi:"getCallerIdentity,optional"`
	AccountAlias                pulumi.StringInput `pulumi:"accountAlias"`
	CreateAccountPasswordPolicy bool               `pulumi:"createAccountPasswordPolicy,optional"`
	MaxPasswordAge              pulumi.IntInput    `pulumi:"maxPasswordAge,optional"`
	MinimumPasswordLength       pulumi.IntInput    `pulumi:"minimumPasswordLength,optional"`
	AllowUsersToChangePassword  pulumi.BoolInput   `pulumi:"allowUsersToChangePassword,optional"`
	HardExpiry                  pulumi.BoolInput   `pulumi:"hardExpiry,optional"`
	PasswordReusePrevention     pulumi.IntPtrInput `pulumi:"passwordReusePrevention,optional"`
	RequireLowercaseCharacters  pulumi.BoolInput   `pulumi:"requireLowercaseCharacters,optional"`
	RequireUppercaseCharacters  pulumi.BoolInput   `pulumi:"requireUppercaseCharacters,optional"`
	RequireNumbers              pulumi.BoolInput   `pulumi:"requireNumbers,optional"`
	RequireSymbols              pulumi.BoolInput   `pulumi:"requireSymbols,optional"`
}

// (Optional) Added Annotate for descriptions.
func (args *IamAccountArgs) Annotate(a infer.Annotator) {
	a.Describe(&args.GetCallerIdentity, "Whether to get AWS account ID, User ID, and ARN in which Terraform is authorized.")
	a.Describe(&args.AccountAlias, "AWS IAM account alias for this account.")
	a.Describe(&args.CreateAccountPasswordPolicy, "Whether to create AWS IAM account password policy.")
	a.Describe(&args.MaxPasswordAge, "The number of days that an user password is valid.")
	a.Describe(&args.MinimumPasswordLength, "Minimum length to require for user passwords.")
	a.Describe(&args.AllowUsersToChangePassword, "Whether to allow users to change their own password.")
	a.Describe(&args.HardExpiry, "Whether users are prevented from setting a new password after their password has expired (i.e. require administrator reset).")
	a.Describe(&args.PasswordReusePrevention, "The number of previous passwords that users are prevented from reusing.")
	a.Describe(&args.RequireLowercaseCharacters, "Whether to require lowercase characters for user passwords.")
	a.Describe(&args.RequireUppercaseCharacters, "Whether to require uppercase characters for user passwords.")
	a.Describe(&args.RequireNumbers, "Whether to require numbers for user passwords.")
	a.Describe(&args.RequireSymbols, "Whether to require symbols for user passwords.")
}

// Fixed types and added `pulumi:` tags to fields.
// Renamed to have "State" suffix.
type IamAccountState struct {
	pulumi.ResourceState
	CallerIdentityAccountId                 pulumi.StringOutput `pulumi:"callerIdentityAccountId"`
	CallerIdentityArn                       pulumi.StringOutput `pulumi:"callerIdentityArn"`
	CallerIdentityUserId                    pulumi.StringOutput `pulumi:"callerIdentityUserId"`
	IamAccountPasswordPolicyExpirePasswords pulumi.BoolOutput   `pulumi:"iamAccountPasswordPolicyExpirePasswords"`
}

// (Optional) Added Annotate for descriptions.
func (state *IamAccountState) Annotate(a infer.Annotator) {
	a.Describe(&state.CallerIdentityAccountId, "The AWS Account ID number of the account that owns or contains the calling entity.")
	a.Describe(&state.CallerIdentityArn, "The AWS ARN associated with the calling entity.")
	a.Describe(&state.CallerIdentityUserId, "The unique identifier of the calling entity.")
	a.Describe(&state.IamAccountPasswordPolicyExpirePasswords, "Indicates whether passwords in the account expire. Returns true if max_password_age contains a value greater than 0. Returns false if it is 0 or not present.")
}

// Lowercased function name (no need to export) and added `typ` parameter.
func newIamAccount(
	ctx *pulumi.Context,
	typ,
	name string,
	args *IamAccountArgs,
	opts ...pulumi.ResourceOption,
) (*IamAccountState, error) {
	var componentResource IamAccountState
	// Pass the typ rather than using the typ that the converter chose.
	err := ctx.RegisterComponentResource(typ, name, &componentResource, opts...)
	if err != nil {
		return nil, err
	}

	var accountId, arn, userId string
	if args.GetCallerIdentity {
		result, err := aws.GetCallerIdentity(ctx, pulumi.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		accountId, arn, userId = result.AccountId, result.Arn, result.UserId
	}

	_, err = iam.NewAccountAlias(ctx, fmt.Sprintf("%s-alias", name), &iam.AccountAliasArgs{
		AccountAlias: args.AccountAlias,
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}

	expirePasswords := pulumi.Bool(false).ToBoolOutput()
	if args.CreateAccountPasswordPolicy {
		policy, err := iam.NewAccountPasswordPolicy(ctx, fmt.Sprintf("%s-policy", name), &iam.AccountPasswordPolicyArgs{
			MaxPasswordAge:             args.MaxPasswordAge,
			MinimumPasswordLength:      args.MinimumPasswordLength,
			AllowUsersToChangePassword: args.AllowUsersToChangePassword,
			HardExpiry:                 args.HardExpiry,
			PasswordReusePrevention:    args.PasswordReusePrevention,
			RequireLowercaseCharacters: args.RequireLowercaseCharacters,
			RequireUppercaseCharacters: args.RequireUppercaseCharacters,
			RequireNumbers:             args.RequireNumbers,
			RequireSymbols:             args.RequireSymbols,
		}, pulumi.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		expirePasswords = policy.ExpirePasswords
	}

	// No need to call ctx.RegisterResourceOutputs. The Go Provider SDK will do it for us
	// using the `pulumi` tagged fields on the struct.
	componentResource.CallerIdentityAccountId = pulumi.String(accountId).ToStringOutput()
	componentResource.CallerIdentityArn = pulumi.String(arn).ToStringOutput()
	componentResource.CallerIdentityUserId = pulumi.String(userId).ToStringOutput()
	componentResource.IamAccountPasswordPolicyExpirePasswords = expirePasswords

	return &componentResource, nil
}
