# title: source ~/.shenv now if it exists
# tags: shell
# source ~/.shenv now if it exists
test -r ~/.shenv &&
. ~/.shenv
