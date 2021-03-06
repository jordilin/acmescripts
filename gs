#!/bin/bash

# Middle click gs on a repository and it will tell you if there's any file
# that has been modified or is untracked. It will prepend the "git add"
# command before the filename, so adding it is just a one click operation.

output=$(git status)
modified=$(echo "$output" | awk '/Changes not staged/{f=1; line=$0} f && /modified/ { print $1 " (gd " $2 ") (ga " $2 ") (gc) (gac " $2 ") (gp) (git checkout -f " $2 ")"} /no changes added/{f=0} ')
untracked=$(echo "$output" | awk '/Untracked files:/{f=1; line=$0} f && /\t[a-zA-Z]/ { print "untracked: (ga " $1 ") (gc) (gac " $1 ") (gp) (rm " $1 ")"} /untracked files present/{f=0} ')

if test "$modified" && test "$untracked";
then
	echo "$modified"
	echo "$untracked"
	exit 0
elif test "$modified";
then
	echo "$modified"
	exit 0
elif test "$untracked";
then
	echo "$untracked"
	exit 0
else
	echo "Nothing to add."
fi
