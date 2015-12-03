```

nrcq(8)             System Administration Utilities            nrcq(8)

NAME
  nrcq - NagRestConf Query utility

SYNOPSIS
  nrcq [options] URL ENDPOINT

DESCRIPTION
  -c, --complete=false: Also show fields with empty values.
  -d, --data=[]: Set extra data to send, 'option:value'.
	The user should not urlencode data, nrcq will do it.
	May be used multiple times.
  -f, --filter="": A client side RE2 regex filter, 'option:regex[,option:regex]...'
  -F, --folder="local": The system folder to query.
  -j, --json=false: Output in JSON format.
  -l, --list="": List all options for the specified table. Required fields are
	 preceded by a star, '*'.
  -L, --listendpoints=false: List all endpoints/tables.
  -p, --pack=false: Remove spaces and lines from the Json output.
  -P, --password="": Password for Basic Auth.
  -U, --username="": Username for Basic Auth.
  -v, --version=false: Show the version of this program.

EXAMPLES
  Show all valid endpoints:
    nrcq -l servicesets

  List all nagios options for the servicesets table:
    nrcq -l servicesets

  Show all hosts:
    nrcq http://server/rest show/hosts

  Show a subset of hosts using a simple RE2 regular expression:
    nrcq http://server/rest show/hosts -f "name:host2"

  Show a subset of services using a complex RE2 regular expression:
  (See https://github.com/google/re2/wiki/Syntax)
    nrcq http://server/rest show/services -f "name:\bhost2\b|web,svcdesc:(?i)swap"

  Add a new host:
    nrcq http://server/rest add/hosts \
      -d name:server1 \
      -d alias:server1 \
      -d ipaddress:server1.there.gq \
      -d template:hsttmpl-local \
      -d servicesets:example-lin

  Delete a host and all of its services:
    nrcq http://server/rest delete/services \
      -d name:server1 \
      -d "svcdesc:.*"
    nrcq http://server/rest delete/hosts \
      -d name:server1 \

```

Compiled binaries are available for Linux, Mac and Windows at [sourceforge](https://sourceforge.net/projects/nagrestconf/files/nrcq).

To compile from source, install [Google Go](https://golang.org/dl/) and [Git](https://git-scm.com/downloads) then use:

    go get github.com/mclarkson/nrcq

There is a short [REST Tutorial](http://nagrestconf.smorg.co.uk/documentation/resttut.php).

