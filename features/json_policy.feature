Feature: JSON Policy

  Background:
    Given I have "go" command installed
    When I run `go build -o ../../bin/vaultpolicy2json-int-testing ../../main.go`
    Then the exit status should be 0

  Scenario:
    Given a file named "policy.hcl" with:
      """
      path "secret/*" {
        capabilities = ["create", "read", "update", "delete", "list"]
      }

      path "secret/super-secret" {
        capabilities = ["deny"]
      }
      """
    When I run `vaultpolicy2json-int-testing --json-policy` interactively
    And I pipe in the file "policy.hcl"
    Then the output should contain exactly:
      """

      {"path":[{"secret/*":{"capabilities":["create","read","update","delete","list"]}},{"secret/super-secret":{"capabilities":["deny"]}}]}
      """
    And the exit status should be 0
