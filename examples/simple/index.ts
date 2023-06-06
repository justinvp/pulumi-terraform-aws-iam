import * as iam from "@pulumi/terraform-aws-iam";

const account = new iam.IamAccount("account", {
    accountAlias: "new-test-account-awesome-company",
    minimumPasswordLength: 6,
    requireNumbers: false,
});
