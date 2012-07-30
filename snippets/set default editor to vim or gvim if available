# title: set default editor to vim or gvim if available
# tags: pager & editor
# set default editor to vim or gvim if available
HAVE_VIM=$(command -v vim)
HAVE_GVIM=$(command -v gvim)

test -n "$HAVE_VIM" &&
EDITOR=vim ||
EDITOR=vi
export EDITOR
