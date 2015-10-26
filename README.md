```
nrcq(8)             System Administration Utilities            nrcq(8)

NAME
  nrcq - NagRestConf Query utility

SYNOPSIS
  nrcq [options] URL ENDPOINT

DESCRIPTION
  -c, --complete=false: Show fields with empty values in Json output.
  -d, --data=: Set extra data to send, 'option:value[,option:value]...'
        May be used multiple times.
  -e, --encode=false: Encode output so it can be piped to another tool.
  -f, --filter="": A client side RE2 regex filter, 'option:regex[,option:regex]...'
  -F, --folder="local": The system folder to query.
  -j, --json=false: Output the table list (-l) in JSON.
  -l, --list="": List all options for the specified table. Required fields are
         preceded by a star, '*'.
  -p, --pack=false: Remove spaces and lines from the Json output.

EXAMPLES
  List all nagios options in the servicesets table:
    nrcq -l servicesets

  Show all hosts:
    nrcq http://server/rest show/hosts

  Show a subset of services using an RE2 regular expression:
    nrcq http://server/rest show/services -f "name:\bhost2\b|web,svcdesc:(?i)swap"

  Add a new host:
    nrcq http://server/rest add/hosts \
      -d name:server1,alias:server1 \
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

