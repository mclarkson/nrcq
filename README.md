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

EXAMPLES
  List all nagios options for the servicesets table:
    nrcq -l servicesets

  Show all hosts:
    nrcq http://server/rest show/hosts

  Show a subset of services using an RE2 regular expression:
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

Compiled binaries are available at [sourceforge](https://sourceforge.net/projects/nagrestconf/files/ncrq) for Linux, Mac and Windows.

To compile from source install [Google Go](https://golang.org/dl/) and [Git](https://git-scm.com/downloads) then use:

    go get github.com/mclarkson/nrcq

