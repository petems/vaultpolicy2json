# `vaultpolicy2json`

[![Build Status](https://travis-ci.org/petems/vaultpolicy2json.svg?branch=master)](https://travis-ci.org/petems/vaultpolicy2json)

A CLI tool for converting HCL policies into JSON to use with the Vault policy API

## Background

Lets say I have the following HCL policy for reading secrets:

```bash
cat all_secrets.hcl
# List, create, update, and delete key/value secrets
path "secret/*"
{
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}
```

When using the Vault CLI, you can provide it a policy as an HCL file and it works fine:

```bash
vault policy write all_secrets all_secrets.hcl
Success! Uploaded policy: all_secrets
```

However, you cant just provide this HCL file to the API:

```bash
$ curl --header "X-Vault-Token: $VAULT_TOKEN"   --request PUT   --data @all_secrets.hcl  $VAULT_ADDR/v1/sys/policy/secret
{"errors":["failed to parse JSON input: invalid character '#' looking for beginning of value"]}
```

Instead, you will have to provide the information as JSON file, with a key of "policy" with the value is the HCL policy with valid escapes to work in JSON.

Like this:

```json
{
  "policy":"# List, create, update, and delete key/value secrets\npath \"secret/*\"\n{\n  capabilities = [\"create\", \"read\", \"update\", \"delete\", \"list\", \"sudo\"]\n}"
}
```

This isn't too hard to script...

```bash
$ tee to_json_policy.sh <<EOF
#!/usr/bin/env bash
printf "{\"policy\": \""
cat \$1 | sed '/^[[:blank:]]*#/d;s/#.*//' | sed 's/\"/\\\"/g' | tr -d '\n'
printf "\"}"
EOF
$ bash to_json_policy.sh all_secrets.hcl
{"policy": "path \"secret/*\"{  capabilities = [\"create\", \"read\", \"update\", \"delete\", \"list\", \"sudo\"]}"}
```

But wouldn't it be nice to have this all done for you, and to check the HCL and JSON are valid?

## Usage

The tool waits for input from STDIN, so you can do something like this:

```bash
$ cat all_secrets.hcl
# List, create, update, and delete key/value secrets
path "secret/*"
{
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}
$ vaultpolicy2json < all_secrets.hcl
{"policy":"# List, create, update, and delete key/value secrets\npath \"secret/*\"\n{\n  capabilities = [\"create\", \"read\", \"update\", \"delete\", \"list\", \"sudo\"]\n}"}
$ curl --header "X-Vault-Token: $VAULT_TOKEN"   --request PUT   --data @all_secrets.json  $VAULT_ADDR/v1/sys/policy/secret
$ curl --header "X-Vault-Token: $VAULT_TOKEN" $VAULT_ADDR/v1/sys/policy/all_secrets
{"name":"all_secrets","rules":"# List, create, update, and delete key/value secrets\npath \"secret/*\"\n{\n  capabilities = [\"create\", \"read\", \"update\", \"delete\", \"list\", \"sudo\"]\n}","request_id":"fd2d45f1-0ba4-7270-2c3f-44c016c6555b","lease_id":"","renewable":false,"lease_duration":0,"data":{"name":"all_secrets","rules":"# List, create, update, and delete key/value secrets\npath \"secret/*\"\n{\n  capabilities = [\"create\", \"read\", \"update\", \"delete\", \"list\", \"sudo\"]\n}"},"wrap_info":null,"warnings":null,"auth":null}
```

## Installation

```bash
go get -v github.com/petems/vaultpolicy2json
```

Eventually I'll configure Travis to build binaries and setup a `brew tap` for OSX and Linux.

## Testing

The test suite consists of aruba/cucumber tests

###Aruba tests

* `bundle install`
* `bundle exec cucumber`
