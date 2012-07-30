# title: punshift(): Shift <path> onto the beginning of PATH or environment variable <var>.
# tags: function | path
# Shift <path> onto the beginning of PATH or environment variable <var>.
# Usage: punshift <path> [<var>]
punshift () { eval "${2:-PATH}='$1:$(eval echo \$${2:-PATH})'"; }
