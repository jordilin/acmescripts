#!/bin/sh

gitop="git diff"

if test "$2" = "staged";
then
	gitop="git diff --staged"
fi

$gitop $1 | 9p write acme/new/body
last=$(9p ls acme | sort -g | tail -n 1)
echo "name diff-$1" | 9p write acme/$last/ctl
echo -n "clean" | 9p write acme/$last/ctl
echo -n "0,0" | 9p write acme/$last/addr
echo -n "dot=addr" | 9p write acme/$last/ctl
echo -n "show" | 9p write acme/$last/ctl

