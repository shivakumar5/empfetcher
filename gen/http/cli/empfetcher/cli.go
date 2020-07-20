// Code generated by goa v3.2.0, DO NOT EDIT.
//
// empfetcher HTTP client CLI support package
//
// Command:
// $ goa gen github.com/flexera/empfetcher/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	empfetcherc "github.com/flexera/empfetcher/gen/http/empfetcher/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `empfetcher (add|update|list|show|delete|restore|viewdeleted|search)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` empfetcher add --body '{
      "address": "Bangalore",
      "department": "development",
      "id": "fgfhjsddctybnjgjh",
      "name": "shiva",
      "skills": "golang, docker"
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		empfetcherFlags = flag.NewFlagSet("empfetcher", flag.ContinueOnError)

		empfetcherAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		empfetcherAddBodyFlag = empfetcherAddFlags.String("body", "REQUIRED", "")

		empfetcherUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		empfetcherUpdateBodyFlag = empfetcherUpdateFlags.String("body", "REQUIRED", "")
		empfetcherUpdateIDFlag   = empfetcherUpdateFlags.String("id", "REQUIRED", "Unique ID of an Employee")

		empfetcherListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		empfetcherShowFlags  = flag.NewFlagSet("show", flag.ExitOnError)
		empfetcherShowIDFlag = empfetcherShowFlags.String("id", "REQUIRED", "ID is the unique id of an employee")

		empfetcherDeleteFlags          = flag.NewFlagSet("delete", flag.ExitOnError)
		empfetcherDeleteIDFlag         = empfetcherDeleteFlags.String("id", "REQUIRED", "ID is the unique id of an employee")
		empfetcherDeletePermdeleteFlag = empfetcherDeleteFlags.String("permdelete", "", "")

		empfetcherRestoreFlags  = flag.NewFlagSet("restore", flag.ExitOnError)
		empfetcherRestoreIDFlag = empfetcherRestoreFlags.String("id", "REQUIRED", "ID is the unique id of an employee")

		empfetcherViewdeletedFlags = flag.NewFlagSet("viewdeleted", flag.ExitOnError)

		empfetcherSearchFlags    = flag.NewFlagSet("search", flag.ExitOnError)
		empfetcherSearchBodyFlag = empfetcherSearchFlags.String("body", "REQUIRED", "")
	)
	empfetcherFlags.Usage = empfetcherUsage
	empfetcherAddFlags.Usage = empfetcherAddUsage
	empfetcherUpdateFlags.Usage = empfetcherUpdateUsage
	empfetcherListFlags.Usage = empfetcherListUsage
	empfetcherShowFlags.Usage = empfetcherShowUsage
	empfetcherDeleteFlags.Usage = empfetcherDeleteUsage
	empfetcherRestoreFlags.Usage = empfetcherRestoreUsage
	empfetcherViewdeletedFlags.Usage = empfetcherViewdeletedUsage
	empfetcherSearchFlags.Usage = empfetcherSearchUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "empfetcher":
			svcf = empfetcherFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "empfetcher":
			switch epn {
			case "add":
				epf = empfetcherAddFlags

			case "update":
				epf = empfetcherUpdateFlags

			case "list":
				epf = empfetcherListFlags

			case "show":
				epf = empfetcherShowFlags

			case "delete":
				epf = empfetcherDeleteFlags

			case "restore":
				epf = empfetcherRestoreFlags

			case "viewdeleted":
				epf = empfetcherViewdeletedFlags

			case "search":
				epf = empfetcherSearchFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "empfetcher":
			c := empfetcherc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = empfetcherc.BuildAddPayload(*empfetcherAddBodyFlag)
			case "update":
				endpoint = c.Update()
				data, err = empfetcherc.BuildUpdatePayload(*empfetcherUpdateBodyFlag, *empfetcherUpdateIDFlag)
			case "list":
				endpoint = c.List()
				data = nil
			case "show":
				endpoint = c.Show()
				data, err = empfetcherc.BuildShowPayload(*empfetcherShowIDFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = empfetcherc.BuildDeletePayload(*empfetcherDeleteIDFlag, *empfetcherDeletePermdeleteFlag)
			case "restore":
				endpoint = c.Restore()
				data, err = empfetcherc.BuildRestorePayload(*empfetcherRestoreIDFlag)
			case "viewdeleted":
				endpoint = c.Viewdeleted()
				data = nil
			case "search":
				endpoint = c.Search()
				data, err = empfetcherc.BuildSearchPayload(*empfetcherSearchBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// empfetcherUsage displays the usage of the empfetcher command and its
// subcommands.
func empfetcherUsage() {
	fmt.Fprintf(os.Stderr, `Service is the empfetcher service interface.
Usage:
    %s [globalflags] empfetcher COMMAND [flags]

COMMAND:
    add: Adds an Employee Details
    update: Updates an Employee Details
    list: List All Employee Details
    show: Show Employee Details based on ID
    delete: Delete Employee Details
    restore: Restores an Employee Details
    viewdeleted: View All deactivated Employee Details
    search: Search employees by name

Additional help:
    %s empfetcher COMMAND --help
`, os.Args[0], os.Args[0])
}
func empfetcherAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] empfetcher add -body JSON

Adds an Employee Details
    -body JSON: 

Example:
    `+os.Args[0]+` empfetcher add --body '{
      "address": "Bangalore",
      "department": "development",
      "id": "fgfhjsddctybnjgjh",
      "name": "shiva",
      "skills": "golang, docker"
   }'
`, os.Args[0])
}

func empfetcherUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] empfetcher update -body JSON -id STRING

Updates an Employee Details
    -body JSON: 
    -id STRING: Unique ID of an Employee

Example:
    `+os.Args[0]+` empfetcher update --body '{
      "address": "Bangalore",
      "department": "development",
      "name": "shiva",
      "skills": "golang, docker"
   }' --id "fgfhjsddctybnjgjh"
`, os.Args[0])
}

func empfetcherListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] empfetcher list

List All Employee Details

Example:
    `+os.Args[0]+` empfetcher list
`, os.Args[0])
}

func empfetcherShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] empfetcher show -id STRING

Show Employee Details based on ID
    -id STRING: ID is the unique id of an employee

Example:
    `+os.Args[0]+` empfetcher show --id "Necessitatibus delectus veniam et officia veniam."
`, os.Args[0])
}

func empfetcherDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] empfetcher delete -id STRING -permdelete BOOL

Delete Employee Details
    -id STRING: ID is the unique id of an employee
    -permdelete BOOL: 

Example:
    `+os.Args[0]+` empfetcher delete --id "Doloremque maiores qui." --permdelete false
`, os.Args[0])
}

func empfetcherRestoreUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] empfetcher restore -id STRING

Restores an Employee Details
    -id STRING: ID is the unique id of an employee

Example:
    `+os.Args[0]+` empfetcher restore --id "Sed architecto architecto."
`, os.Args[0])
}

func empfetcherViewdeletedUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] empfetcher viewdeleted

View All deactivated Employee Details

Example:
    `+os.Args[0]+` empfetcher viewdeleted
`, os.Args[0])
}

func empfetcherSearchUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] empfetcher search -body JSON

Search employees by name
    -body JSON: 

Example:
    `+os.Args[0]+` empfetcher search --body '{
      "searchString": "Non illum et aspernatur."
   }'
`, os.Args[0])
}
