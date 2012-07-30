# title: put ports on the paths if /opt/local exists
# tags: osx | path
# put ports on the paths if /opt/local exists
test -x /opt/local -a ! -L /opt/local && {
    PORTS=/opt/local

    # setup the PATH and MANPATH
    PATH="$PORTS/bin:$PORTS/sbin:$PATH"
    MANPATH="$PORTS/share/man:$MANPATH"

    # nice little port alias
    alias port="sudo nice -n +18 $PORTS/bin/port"
}
