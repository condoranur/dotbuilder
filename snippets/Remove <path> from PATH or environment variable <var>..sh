# title: prm(): Remove <path> from PATH or environment variable <var>.
# tags: function | path
# requires: pls(): List path entries of PATH or environment variable <var>.
# Usage: prm <path> [<var>]
prm () { eval "${2:-PATH}='$(pls $2 |grep -v "^$1\$" |tr '\n' :)'"; }
