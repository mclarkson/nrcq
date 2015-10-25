
nrcq(8)             System Administration Utilities            nrcq(8)

NAME
  nrcq - NagRestConf Query utility

SYNOPSIS
  nrcq [options] URL ENDPOINT

DESCRIPTION
  -c, --complete=false: Show fields with empty values in Json output.
  -d, --data="": Set extra data to send, 'option:value[,option:value]...'
  -e, --encode=false: Encode output so it can be piped to another tool.
  -f, --filter="": A client side RE2 regex filter, 'option:regex[,option:regex]...'
  -F, --folder="local": The system folder to query.
  -p, --pack=false: Remove spaces and lines from the Json output.

EXAMPLES
  Show all hosts:
    nrcq http://server/rest show/hosts

  Show a subset of services:
    nrcq http://server/rest show/services -f "name:\bhost2\b|web,svcdesc:(?i)swap"

