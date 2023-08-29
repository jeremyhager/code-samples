provider "aws" {
  region = "us-east-1"
}

data "aws_iam_policy" "adminAccess" {
  name = "AdministratorAccess"
}

module "iam_user" {
  source = "../../../modules/iam-user"

  name = "iamadmin"
  password_reset_required = false
  policy_arns = [data.aws_iam_policy.adminAccess.arn]

}