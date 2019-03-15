Feature: Version Command

  Background:
    Given I have "go" command installed
    When I run `go build -o ../../bin/vaultpolicy2json-int-testing ../../main.go`
    Then the exit status should be 0

  Scenario: Version Command
    Given a build of 'vaultpolicy2json'
    When I run `bin/vaultpolicy2json-int-testing --version`
    Then the output should match /^\d\.\d\.\d/
    And the exit status should be 0
