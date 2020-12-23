package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the simple Terraform module in examples/terraform-basic-example using Terratest.
func TestTerraformVPC(t *testing.T) {
	t.Parallel()

	expectedCIDRBlock := "10.0.1.0/24"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/terraform-aws-vpc",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"cidr_block": expectedCIDRBlock,
		},

		// Variables to pass to our Terraform code using -var-file options
		// VarFiles: []string{"varfile.tfvars"},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	})

	// website::tag::4::Clean up resources with "terraform destroy". Using "defer" runs the command at the end of the test, whether the test succeeds or fails.
	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// website::tag::2::Run "terraform init" and "terraform apply".
	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	actualCIDRBlock := terraform.Output(t, terraformOptions, "cidr_block")
	fmt.Println(actualCIDRBlock)

	// website::tag::3::Check the output against expected values.
	// Verify we're getting back the outputs we expect
	assert.Equal(t, expectedCIDRBlock, actualCIDRBlock)
}

func TestTerraformVPCWithTags(t *testing.T) {
	t.Parallel()

	expectedCIDRBlock := "10.0.2.0/24"
	VPCTags := map[string]string{"environment": "development", "location": "eu-west-2"}

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/terraform-aws-vpc",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"cidr_block": expectedCIDRBlock,
			"tags":       VPCTags,
		},

		// Variables to pass to our Terraform code using -var-file options
		// VarFiles: []string{"varfile.tfvars"},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	})

	// website::tag::4::Clean up resources with "terraform destroy". Using "defer" runs the command at the end of the test, whether the test succeeds or fails.
	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// website::tag::2::Run "terraform init" and "terraform apply".
	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	actualCIDRBlock := terraform.Output(t, terraformOptions, "cidr_block")
	actualVPCTags := terraform.Output(t, terraformOptions, "tags")

	// website::tag::3::Check the output against expected values.
	// Verify we're getting back the outputs we expect
	assert.Equal(t, actualCIDRBlock, expectedCIDRBlock)
	assert.Contains(t, actualVPCTags, "environment")
	assert.Contains(t, actualVPCTags, "location")
}

func TestTerraformVPCWithTagsDNS(t *testing.T) {
	t.Parallel()

	expectedCIDRBlock := "10.0.2.0/24"
	VPCTags := map[string]string{"environment": "development", "location": "eu-west-2"}
	enableDNSSsupport := "false"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/terraform-aws-vpc",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"cidr_block":         expectedCIDRBlock,
			"tags":               VPCTags,
			"enable_dns_support": enableDNSSsupport,
		},

		// Variables to pass to our Terraform code using -var-file options
		// VarFiles: []string{"varfile.tfvars"},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	})

	// website::tag::4::Clean up resources with "terraform destroy". Using "defer" runs the command at the end of the test, whether the test succeeds or fails.
	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// website::tag::2::Run "terraform init" and "terraform apply".
	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	actualCIDRBlock := terraform.Output(t, terraformOptions, "cidr_block")
	actualVPCTags := terraform.Output(t, terraformOptions, "tags")
	actualenablednssupport := terraform.Output(t, terraformOptions, "enable_dns_support")

	// website::tag::3::Check the output against expected values.
	// Verify we're getting back the outputs we expect
	assert.Equal(t, actualCIDRBlock, expectedCIDRBlock)
	assert.Contains(t, actualVPCTags, "environment")
	assert.Contains(t, actualVPCTags, "location")
	assert.Equal(t, "false", actualenablednssupport)
}
