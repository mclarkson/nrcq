package main

import (
	"fmt"
	"github.com/mclarkson/nagrestconf-golib"
	flag "github.com/ogier/pflag" // Drop-in replacement for flag
	"os"
	"strings"
)

type args struct {
	folder  string // The system folder to query
	newline bool   // True - Output newlines
	brief   bool   // True - Omit empty fields
	filter  string // Regex filter
	data    string // Data to send
	encode  bool   // Encode output
	list    string // Encode output
	json    bool   // Encode output
}

var Args = args{}

type data string

var dataFlag data

func (a *data) Set(value string) error {
	comma := ","
	if len(*a) == 0 {
		comma = ""
	}
	*a += data(comma + value)
	return nil
}
func (a *data) String() string {
	return fmt.Sprint(*a)
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nnrcq(8)             System Administration Utilities            nrcq(8)\n")
		fmt.Fprintf(os.Stderr, "\nNAME\n")
		fmt.Fprintf(os.Stderr, "  nrcq - NagRestConf Query utility\n")
		fmt.Fprintf(os.Stderr, "\nSYNOPSIS\n")
		fmt.Fprintf(os.Stderr, "  nrcq [options] URL ENDPOINT\n")
		fmt.Fprintf(os.Stderr, "\nDESCRIPTION\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nEXAMPLES\n")
		fmt.Fprintf(os.Stderr, "  List all nagios options for the servicesets table:\n")
		fmt.Fprintf(os.Stderr, "    nrcq -l servicesets\n")
		fmt.Fprintf(os.Stderr, "\n  Show all hosts:\n")
		fmt.Fprintf(os.Stderr, "    nrcq http://server/rest show/hosts\n")
		fmt.Fprintf(os.Stderr, "\n  Show a subset of services using an RE2 regular expression:\n")
		fmt.Fprintf(os.Stderr, "    nrcq http://server/rest show/services")
		fmt.Fprintf(os.Stderr, " -f \"name:\\bhost2\\b|web,svcdesc:(?i)swap\"\n")
		fmt.Fprintf(os.Stderr, "\n  Add a new host:\n")
		fmt.Fprintf(os.Stderr, "    nrcq http://server/rest add/hosts \\\n")
		fmt.Fprintf(os.Stderr, "      -d name:server1,alias:server1 \\\n")
		fmt.Fprintf(os.Stderr, "      -d ipaddress:server1.there.gq \\\n")
		fmt.Fprintf(os.Stderr, "      -d template:hsttmpl-local \\\n")
		fmt.Fprintf(os.Stderr, "      -d servicesets:example-lin\n")
		fmt.Fprintf(os.Stderr, "\n  Delete a host and all of its services:\n")
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
	flag.BoolVarP(&Args.newline, "pack", "p", false,
		"Remove spaces and lines from the Json output.")
	flag.BoolVarP(&Args.brief, "complete", "c", false,
		"Show fields with empty values in Json output.")
	flag.BoolVarP(&Args.encode, "encode", "e", false,
		"Encode output so it can be piped to another tool.")
	flag.StringVarP(&Args.list, "list", "l", "",
		"List all options for the specified table. Required fields are\n\t preceded by a star, '*'.")
	flag.BoolVarP(&Args.json, "json", "j", false,
		"Output in JSON format.")
	//flag.StringVarP(&Args.data, "data", "d", "",
	//	"Set extra data to send, 'option:value[,option:value]...'")
	flag.VarP(&dataFlag, "data", "d",
		"Set extra data to send, 'option:value[,option:value]...'\n\tMay be used multiple times.")
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

func createObject(object string) (n nrc.NrcQuery) {
	switch object {
	case "restart":
		n = nrc.NewNrcRestart()
	case "nagios":
		n = nrc.NewNrcRestart()
	case "check":
		n = nrc.NewNrcCheck()
	case "nagiosconfig":
		n = nrc.NewNrcCheck()
	case "hosts":
		n = nrc.NewNrcHosts()
	case "services":
		n = nrc.NewNrcServices()
	case "servicesets":
		n = nrc.NewNrcServicesets()
	case "hosttemplates":
		n = nrc.NewNrcHosttemplates()
	case "servicetemplates":
		n = nrc.NewNrcServicetemplates()
	case "hostgroups":
		n = nrc.NewNrcHostgroups()
	case "servicegroups":
		n = nrc.NewNrcServicegroups()
	case "contacts":
		n = nrc.NewNrcContacts()
	case "contactgroups":
		n = nrc.NewNrcContactgroups()
	case "timeperiods":
		n = nrc.NewNrcTimeperiods()
	case "commands":
		n = nrc.NewNrcCommands()
	case "servicedeps":
		n = nrc.NewNrcServicedeps()
	case "hostdeps":
		n = nrc.NewNrcHostdeps()
	case "serviceesc":
		n = nrc.NewNrcServiceesc()
	case "hostesc":
		n = nrc.NewNrcHostesc()
	case "serviceextinfo":
		n = nrc.NewNrcServiceextinfo()
	case "hostextinfo":
		n = nrc.NewNrcHostextinfo()
	default:
		fmt.Printf("ERROR: Unknown endpoint\n")
		os.Exit(1)
	}
	return n
}

func main() {

	flag.Parse()

	Args.data = string(dataFlag)

	// Args left after flag finishes
	url := flag.Arg(0) // Base URL, eg. "http://1.2.3.4/rest"
	ep := flag.Arg(1)  // end point, eg. "show/hosts"

	// Xfer the encode setting to the library
	nrc.SetEncode(Args.encode)

	if Args.list != "" {

		n := createObject(Args.list)

		if Args.json == true {
			fmt.Printf("%s\n", []byte(n.OptionsJson()))
		} else {
			DisplayArray(n.Options(), n.RequiredOptions())
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
		strings.HasPrefix(ep, "delete/") || strings.HasPrefix(ep, "restart/") {

		// POST REQUESTS

		cmd := strings.Split(ep, "/")
		n := createObject(cmd[1])

		err := n.Post(url, ep, Args.folder, Args.data)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("SUCCESS\n")
	}
}
