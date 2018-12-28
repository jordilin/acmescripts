#!/bin/bash

if test "$1" = "-f";
then
	git push origin +master
else 
	git push origin master
fi
