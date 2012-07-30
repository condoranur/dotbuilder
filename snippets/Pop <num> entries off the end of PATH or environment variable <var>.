# title: ppop(): Pop <num> entries off the end of PATH or environment variable <var>.
# tags: function | path
# Pop <num> entries off the end of PATH or environment variable <var>.
# Usage: ppop [-n <num>] [<var>]
ppop () {
    local n=1 i=0
    [ "$1" = "-n" ] && { n=$2; shift 2; }
    while [ $i -lt $n ]
    do eval "${1:-PATH}='\${${1:-PATH}%:*}'"
       i=$(( i + 1 ))
    done
}
