// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package terraformawsiam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IamAccount struct {
	pulumi.ResourceState

	// The AWS Account ID number of the account that owns or contains the calling entity.
	CallerIdentityAccountId pulumi.StringOutput `pulumi:"callerIdentityAccountId"`
	// The AWS ARN associated with the calling entity.
	CallerIdentityArn pulumi.StringOutput `pulumi:"callerIdentityArn"`
	// The unique identifier of the calling entity.
	CallerIdentityUserId pulumi.StringOutput `pulumi:"callerIdentityUserId"`
	// Indicates whether passwords in the account expire. Returns true if max_password_age contains a value greater than 0. Returns false if it is 0 or not present.
	IamAccountPasswordPolicyExpirePasswords pulumi.BoolOutput `pulumi:"iamAccountPasswordPolicyExpirePasswords"`
}

// NewIamAccount registers a new resource with the given unique name, arguments, and options.
func NewIamAccount(ctx *pulumi.Context,
	name string, args *IamAccountArgs, opts ...pulumi.ResourceOption) (*IamAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountAlias == nil {
		return nil, errors.New("invalid value for required argument 'AccountAlias'")
	}
	var resource IamAccount
	err := ctx.RegisterRemoteComponentResource("terraform-aws-iam:index:IamAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type iamAccountArgs struct {
	// AWS IAM account alias for this account.
	AccountAlias string `pulumi:"accountAlias"`
	// Whether to allow users to change their own password.
	AllowUsersToChangePassword *bool `pulumi:"allowUsersToChangePassword"`
	// Whether to create AWS IAM account password policy.
	CreateAccountPasswordPolicy *bool `pulumi:"createAccountPasswordPolicy"`
	// Whether to get AWS account ID, User ID, and ARN in which Terraform is authorized.
	GetCallerIdentity *bool `pulumi:"getCallerIdentity"`
	// Whether users are prevented from setting a new password after their password has expired (i.e. require administrator reset).
	HardExpiry *bool `pulumi:"hardExpiry"`
	// The number of days that an user password is valid.
	MaxPasswordAge *int `pulumi:"maxPasswordAge"`
	// Minimum length to require for user passwords.
	MinimumPasswordLength *int `pulumi:"minimumPasswordLength"`
	// The number of previous passwords that users are prevented from reusing.
	PasswordReusePrevention *int `pulumi:"passwordReusePrevention"`
	// Whether to require lowercase characters for user passwords.
	RequireLowercaseCharacters *bool `pulumi:"requireLowercaseCharacters"`
	// Whether to require numbers for user passwords.
	RequireNumbers *bool `pulumi:"requireNumbers"`
	// Whether to require symbols for user passwords.
	RequireSymbols *bool `pulumi:"requireSymbols"`
	// Whether to require uppercase characters for user passwords.
	RequireUppercaseCharacters *bool `pulumi:"requireUppercaseCharacters"`
}

// The set of arguments for constructing a IamAccount resource.
type IamAccountArgs struct {
	// AWS IAM account alias for this account.
	AccountAlias pulumi.StringInput
	// Whether to allow users to change their own password.
	AllowUsersToChangePassword pulumi.BoolPtrInput
	// Whether to create AWS IAM account password policy.
	CreateAccountPasswordPolicy pulumi.BoolPtrInput
	// Whether to get AWS account ID, User ID, and ARN in which Terraform is authorized.
	GetCallerIdentity pulumi.BoolPtrInput
	// Whether users are prevented from setting a new password after their password has expired (i.e. require administrator reset).
	HardExpiry pulumi.BoolPtrInput
	// The number of days that an user password is valid.
	MaxPasswordAge pulumi.IntPtrInput
	// Minimum length to require for user passwords.
	MinimumPasswordLength pulumi.IntPtrInput
	// The number of previous passwords that users are prevented from reusing.
	PasswordReusePrevention pulumi.IntPtrInput
	// Whether to require lowercase characters for user passwords.
	RequireLowercaseCharacters pulumi.BoolPtrInput
	// Whether to require numbers for user passwords.
	RequireNumbers pulumi.BoolPtrInput
	// Whether to require symbols for user passwords.
	RequireSymbols pulumi.BoolPtrInput
	// Whether to require uppercase characters for user passwords.
	RequireUppercaseCharacters pulumi.BoolPtrInput
}

func (IamAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamAccountArgs)(nil)).Elem()
}

type IamAccountInput interface {
	pulumi.Input

	ToIamAccountOutput() IamAccountOutput
	ToIamAccountOutputWithContext(ctx context.Context) IamAccountOutput
}

func (*IamAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**IamAccount)(nil)).Elem()
}

func (i *IamAccount) ToIamAccountOutput() IamAccountOutput {
	return i.ToIamAccountOutputWithContext(context.Background())
}

func (i *IamAccount) ToIamAccountOutputWithContext(ctx context.Context) IamAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamAccountOutput)
}

type IamAccountOutput struct{ *pulumi.OutputState }

func (IamAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamAccount)(nil)).Elem()
}

func (o IamAccountOutput) ToIamAccountOutput() IamAccountOutput {
	return o
}

func (o IamAccountOutput) ToIamAccountOutputWithContext(ctx context.Context) IamAccountOutput {
	return o
}

// The AWS Account ID number of the account that owns or contains the calling entity.
func (o IamAccountOutput) CallerIdentityAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamAccount) pulumi.StringOutput { return v.CallerIdentityAccountId }).(pulumi.StringOutput)
}

// The AWS ARN associated with the calling entity.
func (o IamAccountOutput) CallerIdentityArn() pulumi.StringOutput {
	return o.ApplyT(func(v *IamAccount) pulumi.StringOutput { return v.CallerIdentityArn }).(pulumi.StringOutput)
}

// The unique identifier of the calling entity.
func (o IamAccountOutput) CallerIdentityUserId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamAccount) pulumi.StringOutput { return v.CallerIdentityUserId }).(pulumi.StringOutput)
}

// Indicates whether passwords in the account expire. Returns true if max_password_age contains a value greater than 0. Returns false if it is 0 or not present.
func (o IamAccountOutput) IamAccountPasswordPolicyExpirePasswords() pulumi.BoolOutput {
	return o.ApplyT(func(v *IamAccount) pulumi.BoolOutput { return v.IamAccountPasswordPolicyExpirePasswords }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamAccountInput)(nil)).Elem(), &IamAccount{})
	pulumi.RegisterOutputType(IamAccountOutput{})
}
