package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IamAccount struct{}

func (r *IamAccount) Construct(ctx *pulumi.Context, name, typ string, args IamAccountArgs, opts pulumi.ResourceOption) (*IamAccountState, error) {
	comp, err := newIamAccount(ctx, typ, name, &args, opts)
	if err != nil {
		return nil, err
	}
	return comp, nil
}
