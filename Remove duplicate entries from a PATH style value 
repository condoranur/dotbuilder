# title: puniq(): Remove duplicate entries from a PATH style value
# tags: function | path
# Remove duplicate entries from a PATH style value 
# while retaining the original order. Use PATH if no <path> is given.
# Usage: puniq [<path>]
#
# Example:
#   $ puniq /usr/bin:/usr/local/bin:/usr/bin
#   /usr/bin:/usr/local/bin
puniq () {
    echo "$1" |tr : '\n' |nl |sort -u -k 2,2 |sort -n |
    cut -f 2- |tr '\n' : |sed -e 's/:$//' -e 's/^://'
}
