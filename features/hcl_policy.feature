Feature: HCL Policy

  Background:
    Given I have "go" command installed
    When I run `go build -o ../../bin/vaultpolicy2json-int-testing ../../main.go`
    Then the exit status should be 0

  Scenario: Valid Policy
    Given a file named "policy.hcl" with:
      """
      path "secret/*" {
        capabilities = ["create", "read", "update", "delete", "list"]
      }

      path "secret/super-secret" {
        capabilities = ["deny"]
      }
      """
    When I run `vaultpolicy2json-int-testing` interactively
    And I pipe in the file "policy.hcl"
    Then the exit status should be 0
    And the output should contain exactly:
      """

      path "secret/*" {   capabilities = ["create", "read", "update", "delete", "list"] }  path "secret/super-secret" {   capabilities = ["deny"] }
      """
    And the exit status should be 0

  Scenario: Invalid Policy
    Given a file named "invalid.hcl" with:
      """
      INVALID HCL
      """
    When I run `vaultpolicy2json-int-testing` interactively
    And I pipe in the file "invalid.hcl"
    Then the output should contain exactly:
      """

      unable to parse HCL: At 1:13: key 'INVALID HCL' expected start of object ('{') or assignment ('=')
      """
    And the exit status should be 1

