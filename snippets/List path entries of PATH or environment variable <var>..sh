# title: pls(): List path entries of PATH or environment variable <var>.
# tags: function | path
# List path entries of PATH or environment variable <var>.
# Usage: pls [<var>]
pls () { eval echo \$${1:-PATH} |tr : '\n'; }
