package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/hashicorp/hcl"

	removesingletonarrays "github.com/petems/remove-singleton-arrays"
)

// Version is what is returned by the `-v` flag
const Version = "0.1.0"

// gitCommit is the gitcommit its built from
var gitCommit = "development"

// SingletonVaultExceptions are the parameters that are allowed to have singular arrays
var SingletonVaultExceptions = []string{"capabilities", "allowed_parameters", "required_parameters"}

func main() {
	version := flag.Bool("version", false, "Prints current app version")
	jsonPolicy := flag.Bool("json-policy", false, "Convert the HCL policy to JSON also")
	apiJSON := flag.Bool("api-json", true, "Output to JSON to supply to the API")

	flag.Parse()

	if *version {
		fmt.Printf("%s %s\n", Version, gitCommit)
		return
	}

	err := toJSON(*jsonPolicy, *apiJSON)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func toJSON(jsonPolicy bool, apiJSON bool) error {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("unable to read from stdin: %s", err)
	}

	var hclPolicy interface{}
	err = hcl.Unmarshal(input, &hclPolicy)
	if err != nil {
		return fmt.Errorf("unable to parse HCL: %s", err)
	}

	payload := make(map[string]string)

	policyString := ""

	if jsonPolicy {

		jsonPolicy, err := json.Marshal(hclPolicy)
		if err != nil {
			return fmt.Errorf("unable to marshal json: %s", err)
		}

		deSingletonPolicy, err := removesingletonarrays.WithIgnores(string(jsonPolicy), SingletonVaultExceptions)

		if err != nil {
			return fmt.Errorf("unable to remove singletons from json: %s", err)
		}

		policyString = string(deSingletonPolicy)
	} else {
		policyString = string(input)
	}

	re := regexp.MustCompile(`\r?\n`)
	policyString = re.ReplaceAllString(policyString, " ")

	if apiJSON {
		payload["policy"] = policyString
		out, err := json.Marshal(payload)

		if err != nil {
			return fmt.Errorf("unable to marshal JSON: %s", err)
		}

		fmt.Println(string(out))
	} else {

		fmt.Println(string(policyString))
	}

	return nil
}
