# title: pshift(): Shift <num> entries off the front of PATH or environment var <var>.
# tags: function | path
# requires: pls(): List path entries of PATH or environment variable <var>.
# Shift <num> entries off the front of PATH or environment var <var>.
# Usage: pshift [-n <num>] [<var>]
# with the <var> option. Useful: pshift $(pwd)
pshift () {
    local n=1
    [ "$1" = "-n" ] && { n=$(( $2 + 1 )); shift 2; }
    eval "${1:-PATH}='$(pls |tail -n +$n |tr '\n' :)'"
}
