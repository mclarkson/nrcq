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
		fmt.Fprintf(os.Stderr, "  Show all hosts:\n")
		fmt.Fprintf(os.Stderr, "    nrcq http://server/rest show/hosts\n")
		fmt.Fprintf(os.Stderr, "\n  Show a subset of services:\n")
		fmt.Fprintf(os.Stderr, "    nrcq http://server/rest show/services")
		fmt.Fprintf(os.Stderr, " -f \"name:\\bhost2\\b|web,svcdesc:(?i)swap\"\n")
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
		"List all options for the specified table.")
	flag.BoolVarP(&Args.json, "json", "j", false,
		"Output the table list (-l) in JSON.")
	//flag.StringVarP(&Args.data, "data", "d", "",
	//	"Set extra data to send, 'option:value[,option:value]...'")
	flag.VarP(&dataFlag, "data", "d",
		"Set extra data to send, 'option:value[,option:value]...'\n\tMay be used multiple times.")
}

func DisplayArray(a []string) {
	l := 0
	fmt.Println()
	for _, j := range a {
		l += len(j) + 2
		if l > 79 {
			l = 0
			fmt.Printf("\n")
		}
		fmt.Printf("  %s", j)
	}
	fmt.Printf("\n\n")
}

func main() {

	flag.Parse()

	fmt.Printf("%s\n", dataFlag)
	Args.data = string(dataFlag)

	// Args left after flag finishes
	url := flag.Arg(0) // Base URL, eg. "http://1.2.3.4/rest"
	ep := flag.Arg(1)  // end point, eg. "show/hosts"

	nrc.SetEncode(Args.encode)

	if url == "" || ep == "" {
		fmt.Fprintf(os.Stderr, "ERROR: 2 non-option arguments expected.\n")
		flag.Usage()
	}

	if Args.list != "" {
		switch Args.list {
		case "hosts":
			if Args.json == true {
				fmt.Printf("%s\n", []byte(nrc.HostsFieldsJson()))
			} else {
				DisplayArray(nrc.HostsFields())
			}
		case "services":
			if Args.json == true {
				fmt.Printf("%s\n", []byte(nrc.ServicesFieldsJson()))
			} else {
				DisplayArray(nrc.ServicesFields())
			}
		case "servicesets":
			if Args.json == true {
				fmt.Printf("%s\n", []byte(nrc.ServicesetsFieldsJson()))
			} else {
				DisplayArray(nrc.ServicesetsFields())
			}
		case "hosttemplates":
			if Args.json == true {
				fmt.Printf("%s\n", []byte(nrc.HosttemplatesFieldsJson()))
			} else {
				DisplayArray(nrc.HosttemplatesFields())
			}
		case "servicetemplates":
			if Args.json == true {
				fmt.Printf("%s\n", []byte(nrc.ServicetemplatesFieldsJson()))
			} else {
				DisplayArray(nrc.ServicetemplatesFields())
			}
		case "hostgroups":
			if Args.json == true {
				fmt.Printf("%s\n", []byte(nrc.HostgroupsFieldsJson()))
			} else {
				DisplayArray(nrc.HostgroupsFields())
			}
		case "servicegroups":
			if Args.json == true {
				fmt.Printf("%s\n", []byte(nrc.ServicegroupsFieldsJson()))
			} else {
				DisplayArray(nrc.ServicegroupsFields())
			}
		case "contacts":
			if Args.json == true {
				fmt.Printf("%s\n", []byte(nrc.ContactsFieldsJson()))
			} else {
				DisplayArray(nrc.ContactsFields())
			}
		case "contactgroups":
			if Args.json == true {
				fmt.Printf("%s\n", []byte(nrc.ContactgroupsFieldsJson()))
			} else {
				DisplayArray(nrc.ContactgroupsFields())
			}
		case "timeperiods":
			if Args.json == true {
				fmt.Printf("%s\n", []byte(nrc.TimeperiodsFieldsJson()))
			} else {
				DisplayArray(nrc.TimeperiodsFields())
			}
		case "commands":
			if Args.json == true {
				fmt.Printf("%s\n", []byte(nrc.CommandsFieldsJson()))
			} else {
				DisplayArray(nrc.CommandsFields())
			}
		case "servicedeps":
			if Args.json == true {
				fmt.Printf("%s\n", []byte(nrc.ServicedepsFieldsJson()))
			} else {
				DisplayArray(nrc.ServicedepsFields())
			}
		case "hostdeps":
			if Args.json == true {
				fmt.Printf("%s\n", []byte(nrc.HostdepsFieldsJson()))
			} else {
				DisplayArray(nrc.HostdepsFields())
			}
		case "serviceesc":
			if Args.json == true {
				fmt.Printf("%s\n", []byte(nrc.ServiceescFieldsJson()))
			} else {
				DisplayArray(nrc.ServiceescFields())
			}
		case "hostesc":
			if Args.json == true {
				fmt.Printf("%s\n", []byte(nrc.HostescFieldsJson()))
			} else {
				DisplayArray(nrc.HostescFields())
			}
		case "serviceextinfo":
			if Args.json == true {
				fmt.Printf("%s\n", []byte(nrc.ServiceextinfoFieldsJson()))
			} else {
				DisplayArray(nrc.ServiceextinfoFields())
			}
		case "hostextinfo":
			if Args.json == true {
				fmt.Printf("%s\n", []byte(nrc.HostextinfoFieldsJson()))
			} else {
				DisplayArray(nrc.HostextinfoFields())
			}
		default:
			fmt.Printf("ERROR: Unknown table\n")
		}
		os.Exit(0)
	}

	if strings.HasPrefix(ep, "show/") {

		// GET REQUESTS

		switch {

		case strings.HasSuffix(ep, "/hosts"):
			jh := nrc.NewNrcHosts()
			err := jh.GetHosts(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowHostsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/services"):
			jh := nrc.NewNrcServices()
			err := jh.GetServices(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowServicesJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/servicesets"):
			jh := nrc.NewNrcServicesets()
			err := jh.GetServicesets(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowServicesetsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/hosttemplates"):
			jh := nrc.NewNrcHosttemplates()
			err := jh.GetHosttemplates(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowHosttemplatesJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/servicetemplates"):
			jh := nrc.NewNrcServicetemplates()
			err := jh.GetServicetemplates(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowServicetemplatesJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/hostgroups"):
			jh := nrc.NewNrcHostgroups()
			err := jh.GetHostgroups(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowHostgroupsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/servicegroups"):
			jh := nrc.NewNrcServicegroups()
			err := jh.GetServicegroups(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowServicegroupsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/contacts"):
			jh := nrc.NewNrcContacts()
			err := jh.GetContacts(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowContactsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/contactgroups"):
			jh := nrc.NewNrcContactgroups()
			err := jh.GetContactgroups(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowContactgroupsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/timeperiods"):
			jh := nrc.NewNrcTimeperiods()
			err := jh.GetTimeperiods(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowTimeperiodsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/commands"):
			jh := nrc.NewNrcCommands()
			err := jh.GetCommands(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowCommandsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/servicedeps"):
			jh := nrc.NewNrcServicedeps()
			err := jh.GetServicedeps(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowServicedepsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/hostdeps"):
			jh := nrc.NewNrcHostdeps()
			err := jh.GetHostdeps(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowHostdepsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/serviceesc"):
			jh := nrc.NewNrcServiceesc()
			err := jh.GetServiceesc(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowServiceescJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/hostesc"):
			jh := nrc.NewNrcHostesc()
			err := jh.GetHostesc(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowHostescJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/serviceextinfo"):
			jh := nrc.NewNrcServiceextinfo()
			err := jh.GetServiceextinfo(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowServiceextinfoJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/hostextinfo"):
			jh := nrc.NewNrcHostextinfo()
			err := jh.GetHostextinfo(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowHostextinfoJson(Args.newline, Args.brief, Args.filter)

		default:
			fmt.Printf("ERROR: Invalid endpoint.\n")
			os.Exit(1)
		}

	} else if strings.HasPrefix(ep, "add/") || strings.HasPrefix(ep, "modify/") ||
		strings.HasPrefix(ep, "delete/") {

		// GET REQUESTS

		switch {

		case strings.HasSuffix(ep, "/hosts"):
			jh := nrc.NewNrcHosts()
			err := jh.PostHosts(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("SUCCESS\n")

		case strings.HasSuffix(ep, "/services"):
			jh := nrc.NewNrcServices()
			err := jh.PostServices(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("SUCCESS\n")

		case strings.HasSuffix(ep, "/servicesets"):
			jh := nrc.NewNrcServicesets()
			err := jh.PostServicesets(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("SUCCESS\n")

		case strings.HasSuffix(ep, "/hosttemplates"):
			jh := nrc.NewNrcHosttemplates()
			err := jh.PostHosttemplates(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("SUCCESS\n")

		case strings.HasSuffix(ep, "/servicetemplates"):
			jh := nrc.NewNrcServicetemplates()
			err := jh.PostServicetemplates(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("SUCCESS\n")

		case strings.HasSuffix(ep, "/hostgroups"):
			jh := nrc.NewNrcHostgroups()
			err := jh.PostHostgroups(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("SUCCESS\n")

		case strings.HasSuffix(ep, "/servicegroups"):
			jh := nrc.NewNrcServicegroups()
			err := jh.PostServicegroups(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("SUCCESS\n")

		case strings.HasSuffix(ep, "/contacts"):
			jh := nrc.NewNrcContacts()
			err := jh.PostContacts(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("SUCCESS\n")

		case strings.HasSuffix(ep, "/contactgroups"):
			jh := nrc.NewNrcContactgroups()
			err := jh.PostContactgroups(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("SUCCESS\n")

		case strings.HasSuffix(ep, "/timeperiods"):
			jh := nrc.NewNrcTimeperiods()
			err := jh.PostTimeperiods(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("SUCCESS\n")

		case strings.HasSuffix(ep, "/commands"):
			jh := nrc.NewNrcCommands()
			err := jh.PostCommands(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("SUCCESS\n")

		case strings.HasSuffix(ep, "/servicedeps"):
			jh := nrc.NewNrcServicedeps()
			err := jh.PostServicedeps(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("SUCCESS\n")

		case strings.HasSuffix(ep, "/hostdeps"):
			jh := nrc.NewNrcHostdeps()
			err := jh.PostHostdeps(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("SUCCESS\n")

		case strings.HasSuffix(ep, "/serviceesc"):
			jh := nrc.NewNrcServiceesc()
			err := jh.PostServiceesc(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("SUCCESS\n")

		case strings.HasSuffix(ep, "/hostesc"):
			jh := nrc.NewNrcHostesc()
			err := jh.PostHostesc(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("SUCCESS\n")

		case strings.HasSuffix(ep, "/serviceextinfo"):
			jh := nrc.NewNrcServiceextinfo()
			err := jh.PostServiceextinfo(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("SUCCESS\n")

		case strings.HasSuffix(ep, "/hostextinfo"):
			jh := nrc.NewNrcHostextinfo()
			err := jh.PostHostextinfo(url, ep, Args.folder, Args.data)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("SUCCESS\n")

		default:
			fmt.Printf("ERROR: Invalid endpoint.\n")
			os.Exit(1)
		}
	}

}
