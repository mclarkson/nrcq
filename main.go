// nrcq - NagRestConf Query utility
// Copyright (C) 2014  Mark Clarkson
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"github.com/mclarkson/nagrestconf-golib"
	flag "github.com/ogier/pflag" // Drop-in replacement for flag
	"os"
	"strings"
)

var Version = "0.1.2"

type args struct {
	folder        string   // The system folder to query
	newline       bool     // True - Output newlines
	brief         bool     // True - Omit empty fields
	filter        string   // Regex filter
	data          []string // Data to send
	encode        bool     // Encode output
	list          string   // Encode output
	listendpoints bool     // Encode output
	json          bool     // Encode output
	username      string   // Encode output
	password      string   // Encode output
	version       bool     // Whether to display the version
}

var Args = args{}

type data []string

var dataFlag data

func (a *data) Set(value string) error {
	*a = append(*a, (value))
	return nil
}
func (a *data) String() string {
	return fmt.Sprint(*a)
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "nrcq(8)             System Administration Utilities            nrcq(8)\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "NAME\n")
		fmt.Fprintf(os.Stderr, "  nrcq - NagRestConf Query utility\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "SYNOPSIS\n")
		fmt.Fprintf(os.Stderr, "  nrcq [options] URL ENDPOINT\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "DESCRIPTION\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "EXAMPLES\n")
		fmt.Fprintf(os.Stderr, "  Show all valid endpoints:\n")
		fmt.Fprintf(os.Stderr, "    nrcq -L\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "  List all nagios options for the servicesets table:\n")
		fmt.Fprintf(os.Stderr, "    nrcq -l servicesets\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "  Show all hosts:\n")
		fmt.Fprintf(os.Stderr, "    nrcq http://server/rest show/hosts\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "  Show a subset of hosts using a simple RE2 regular expression:\n")
		fmt.Fprintf(os.Stderr, "    nrcq http://server/rest show/hosts")
		fmt.Fprintf(os.Stderr, " -f \"name:host2\"\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "  Show a subset of services using a complex RE2 regular expression:\n")
		fmt.Fprintf(os.Stderr, "  (See https://github.com/google/re2/wiki/Syntax)\n")
		fmt.Fprintf(os.Stderr, "    nrcq http://server/rest show/services")
		fmt.Fprintf(os.Stderr, " -f \"name:\\bhost2\\b|web,svcdesc:(?i)swap\"\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "  Add a new host:\n")
		fmt.Fprintf(os.Stderr, "    nrcq http://server/rest add/hosts \\\n")
		fmt.Fprintf(os.Stderr, "      -d name:server1 \\\n")
		fmt.Fprintf(os.Stderr, "      -d alias:server1 \\\n")
		fmt.Fprintf(os.Stderr, "      -d ipaddress:server1.there.gq \\\n")
		fmt.Fprintf(os.Stderr, "      -d template:hsttmpl-local \\\n")
		fmt.Fprintf(os.Stderr, "      -d servicesets:example-lin\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "  Delete a host and all of its services:\n")
		fmt.Fprintf(os.Stderr, "    nrcq http://server/rest delete/services \\\n")
		fmt.Fprintf(os.Stderr, "      -d name:server1 \\\n")
		fmt.Fprintf(os.Stderr, "      -d \"svcdesc:.*\"\n")
		fmt.Fprintf(os.Stderr, "    nrcq http://server/rest delete/hosts \\\n")
		fmt.Fprintf(os.Stderr, "      -d name:server1 \\\n")
		fmt.Fprintf(os.Stderr, "\n")
	}

	flag.StringVarP(&Args.folder, "folder", "F", "local",
		"The system folder to query.")
	flag.StringVarP(&Args.filter, "filter", "f", "",
		"A client side RE2 regex filter, 'option:regex[,option:regex]...'")
	flag.BoolVarP(&Args.version, "version", "v", false,
		"Show the version of this program.")
	flag.BoolVarP(&Args.newline, "pack", "p", false,
		"Remove spaces and lines from the Json output.")
	flag.BoolVarP(&Args.brief, "complete", "c", false,
		"Also show fields with empty values.")
	// Done automatically now:
	//flag.BoolVarP(&Args.encode, "encode", "e", false,
	//	"URL Encode output where necessary so it can be piped to another tool.")
	flag.BoolVarP(&Args.listendpoints, "listendpoints", "L", false,
		"List all endpoints/tables.")
	flag.StringVarP(&Args.username, "username", "U", "",
		"Username for Basic Auth.")
	flag.StringVarP(&Args.password, "password", "P", "",
		"Password for Basic Auth.")
	flag.StringVarP(&Args.list, "list", "l", "",
		"List all options for the specified table. Required fields are\n\t preceded by a star, '*'.")
	flag.BoolVarP(&Args.json, "json", "j", false,
		"Output in JSON format.")
	flag.VarP(&dataFlag, "data", "d",
		"Set extra data to send, 'option:value'.\n\tThe user should not urlencode data, nrcq will do it.\n\tMay be used multiple times.")
}

func DisplayArray(a, r []string) {
	l := 0
	fmt.Println()
	for _, j := range a {
		s := ""
		sl := 0
		for _, k := range r {
			if j == k {
				sl = 1
				s = "*"
				break
			}
		}
		l += len(j) + 2 + sl
		if l > 79 {
			l = len(j) + 2
			fmt.Printf("\n")
		}
		fmt.Printf("  %s%s", s, j)
	}
	fmt.Printf("\n\n")
}

func endpointarr() []string {
	a := []string{
		"check/nagiosconfig",
		"apply/nagiosconfig",
		"apply/nagioslastgoodconfig",
		"restart/nagios",
		"show|add|modify|delete/hosts",
		"show|add|modify|delete/services",
		"show|add|modify|delete/servicesets",
		"show|add|modify|delete/hosttemplates",
		"show|add|modify|delete/servicetemplates",
		"show|add|modify|delete/hostgroups",
		"show|add|modify|delete/servicegroups",
		"show|add|modify|delete/contacts",
		"show|add|modify|delete/contactgroups",
		"show|add|modify|delete/timeperiods",
		"show|add|modify|delete/commands",
		"show|add|modify|delete/servicedeps",
		"show|add|modify|delete/hostdeps",
		"show|add|modify|delete/serviceesc",
		"show|add|modify|delete/hostesc",
		"show|add|modify|delete/serviceextinfo",
		"show|add|modify|delete/hostextinfo",
	}

	return a
}

func createObject(object string) (n nrc.NrcQuery) {
	switch object {
	case "applynagiosconfig":
		n = nrc.NewNrcApplyConfig(Args.username, Args.password)
	case "nagioslastgoodconfig":
		n = nrc.NewNrcLastGood(Args.username, Args.password)
	case "restart":
		n = nrc.NewNrcRestart(Args.username, Args.password)
	case "nagios":
		n = nrc.NewNrcRestart(Args.username, Args.password)
	case "check":
		n = nrc.NewNrcCheck(Args.username, Args.password)
	case "nagiosconfig": //show
		n = nrc.NewNrcCheck(Args.username, Args.password)
	case "hosts":
		n = nrc.NewNrcHosts(Args.username, Args.password)
	case "services":
		n = nrc.NewNrcServices(Args.username, Args.password)
	case "servicesets":
		n = nrc.NewNrcServicesets(Args.username, Args.password)
	case "hosttemplates":
		n = nrc.NewNrcHosttemplates(Args.username, Args.password)
	case "servicetemplates":
		n = nrc.NewNrcServicetemplates(Args.username, Args.password)
	case "hostgroups":
		n = nrc.NewNrcHostgroups(Args.username, Args.password)
	case "servicegroups":
		n = nrc.NewNrcServicegroups(Args.username, Args.password)
	case "contacts":
		n = nrc.NewNrcContacts(Args.username, Args.password)
	case "contactgroups":
		n = nrc.NewNrcContactgroups(Args.username, Args.password)
	case "timeperiods":
		n = nrc.NewNrcTimeperiods(Args.username, Args.password)
	case "commands":
		n = nrc.NewNrcCommands(Args.username, Args.password)
	case "servicedeps":
		n = nrc.NewNrcServicedeps(Args.username, Args.password)
	case "hostdeps":
		n = nrc.NewNrcHostdeps(Args.username, Args.password)
	case "serviceesc":
		n = nrc.NewNrcServiceesc(Args.username, Args.password)
	case "hostesc":
		n = nrc.NewNrcHostesc(Args.username, Args.password)
	case "serviceextinfo":
		n = nrc.NewNrcServiceextinfo(Args.username, Args.password)
	case "hostextinfo":
		n = nrc.NewNrcHostextinfo(Args.username, Args.password)
	default:
		fmt.Printf("ERROR: Unknown endpoint\n")
		os.Exit(1)
	}
	return n
}

func main() {

	flag.Parse()

	Args.data = []string(dataFlag)

	if Args.version {
		fmt.Printf("Nrcq version is %s\n", Version)
		os.Exit(0)
	}

	if Args.json {
		Args.encode = true
	} else {
		Args.encode = false
	}

	// Xfer the encode setting to the library
	nrc.SetEncode(Args.encode)

	// Args left after flag finishes
	url := flag.Arg(0) // Base URL, eg. "http://1.2.3.4/rest"
	ep := flag.Arg(1)  // end point, eg. "show/hosts"

	if Args.list != "" {

		n := createObject(Args.list)

		if Args.json == true {
			fmt.Printf("%s\n", []byte(n.OptionsJson()))
		} else {
			DisplayArray(n.Options(), n.RequiredOptions())
		}

		os.Exit(0)
	}

	if Args.listendpoints {

		n := endpointarr()

		if Args.json == true {
			fmt.Printf("%s\n", n)
		} else {
			DisplayArray(n, []string{})
		}

		os.Exit(0)
	}

	if url == "" || ep == "" {
		fmt.Fprintf(os.Stderr, "ERROR: 2 non-option arguments expected.\n")
		flag.Usage()
	}

	if strings.HasPrefix(ep, "check/") {

		// GET REQUESTS

		n := createObject("check")

		err := n.Get(url, ep, Args.folder, Args.data)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			os.Exit(1)
		}
		if Args.json == true {
			n.ShowJson(Args.newline, false, "")
		} else {
			n.Show(false, "")
		}

	} else if strings.HasPrefix(ep, "show/") {

		// GET REQUESTS

		cmd := strings.Split(ep, "/")
		n := createObject(cmd[1])

		err := n.Get(url, ep, Args.folder, Args.data)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			os.Exit(1)
		}
		if Args.json == true {
			n.ShowJson(Args.newline, Args.brief, Args.filter)
		} else {
			n.Show(Args.brief, Args.filter)
		}

	} else if strings.HasPrefix(ep, "add/") || strings.HasPrefix(ep, "modify/") ||
		strings.HasPrefix(ep, "delete/") || strings.HasPrefix(ep, "restart/") ||
		ep == "apply/nagioslastgoodconfig" {

		// POST REQUESTS

		cmd := strings.Split(ep, "/")
		n := createObject(cmd[1])

		err := n.Post(url, ep, Args.folder, Args.data)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("SUCCESS\n")

	} else if ep == "apply/nagiosconfig" {

		// This is the only Post request that produces output

		n := createObject("applynagiosconfig")

		err := n.Post(url, ep, Args.folder, Args.data)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			os.Exit(1)
		}
		if Args.json == true {
			n.ShowJson(Args.newline, Args.brief, Args.filter)
		} else {
			n.Show(Args.brief, Args.filter)
		}
	}
}
