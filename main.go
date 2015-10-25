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
	encode  bool   // Encode output
}

var Args = args{}

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
		"An RE2 regex filter in the form 'property:regex[,property:regex]...'")
	flag.BoolVarP(&Args.newline, "pack", "p", false,
		"Remove spaces and lines from the Json output.")
	flag.BoolVarP(&Args.brief, "complete", "c", false,
		"Show fields with empty values in Json output.")
	flag.BoolVarP(&Args.encode, "encode", "e", false,
		"Encode output so it can be piped to another tool.")
	flag.Parse()
}

func main() {

	// Args left after flag finishes
	url := flag.Arg(0) // Base URL, eg. "http://1.2.3.4/rest"
	ep := flag.Arg(1)  // end point, eg. "show/hosts"

	nrc.SetEncode(Args.encode)

	if url == "" || ep == "" {
		fmt.Fprintf(os.Stderr, "ERROR: 2 non-option arguments expected.\n")
		flag.Usage()
	}

	if strings.HasPrefix(ep, "show/") {

		// GET REQUESTS

		switch {

		case strings.HasSuffix(ep, "/hosts"):
			jh := nrc.NewNrcHosts()
			err := jh.GetHosts(url, ep+`?json={"folder":"`+Args.folder+`"}`)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowHostsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/services"):
			jh := nrc.NewNrcServices()
			err := jh.GetServices(url, ep+`?json={"folder":"`+Args.folder+`"}`)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowServicesJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/servicesets"):
			jh := nrc.NewNrcServicesets()
			err := jh.GetServicesets(url, ep+`?json={"folder":"`+Args.folder+`"}`)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowServicesetsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/hosttemplates"):
			jh := nrc.NewNrcHosttemplates()
			err := jh.GetHosttemplates(url, ep+`?json={"folder":"`+Args.folder+`"}`)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowHosttemplatesJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/servicetemplates"):
			jh := nrc.NewNrcServicetemplates()
			err := jh.GetServicetemplates(url, ep+`?json={"folder":"`+Args.folder+`"}`)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowServicetemplatesJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/hostgroups"):
			jh := nrc.NewNrcHostgroups()
			err := jh.GetHostgroups(url, ep+`?json={"folder":"`+Args.folder+`"}`)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowHostgroupsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/servicegroups"):
			jh := nrc.NewNrcServicegroups()
			err := jh.GetServicegroups(url, ep+`?json={"folder":"`+Args.folder+`"}`)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowServicegroupsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/contacts"):
			jh := nrc.NewNrcContacts()
			err := jh.GetContacts(url, ep+`?json={"folder":"`+Args.folder+`"}`)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowContactsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/contactgroups"):
			jh := nrc.NewNrcContactgroups()
			err := jh.GetContactgroups(url, ep+`?json={"folder":"`+Args.folder+`"}`)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowContactgroupsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/timeperiods"):
			jh := nrc.NewNrcTimeperiods()
			err := jh.GetTimeperiods(url, ep+`?json={"folder":"`+Args.folder+`"}`)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowTimeperiodsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/commands"):
			jh := nrc.NewNrcCommands()
			err := jh.GetCommands(url, ep+`?json={"folder":"`+Args.folder+`"}`)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowCommandsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/servicedeps"):
			jh := nrc.NewNrcServicedeps()
			err := jh.GetServicedeps(url, ep+`?json={"folder":"`+Args.folder+`"}`)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowServicedepsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/hostdeps"):
			jh := nrc.NewNrcHostdeps()
			err := jh.GetHostdeps(url, ep+`?json={"folder":"`+Args.folder+`"}`)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowHostdepsJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/serviceesc"):
			jh := nrc.NewNrcServiceesc()
			err := jh.GetServiceesc(url, ep+`?json={"folder":"`+Args.folder+`"}`)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowServiceescJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/hostesc"):
			jh := nrc.NewNrcHostesc()
			err := jh.GetHostesc(url, ep+`?json={"folder":"`+Args.folder+`"}`)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowHostescJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/serviceextinfo"):
			jh := nrc.NewNrcServiceextinfo()
			err := jh.GetServiceextinfo(url, ep+`?json={"folder":"`+Args.folder+`"}`)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowServiceextinfoJson(Args.newline, Args.brief, Args.filter)

		case strings.HasSuffix(ep, "/hostextinfo"):
			jh := nrc.NewNrcHostextinfo()
			err := jh.GetHostextinfo(url, ep+`?json={"folder":"`+Args.folder+`"}`)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
			jh.ShowHostextinfoJson(Args.newline, Args.brief, Args.filter)

		default:
			fmt.Printf("ERROR: Invalid endpoint.\n")
			os.Exit(1)
		}
	}

}
