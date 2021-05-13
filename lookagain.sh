#!/bin/bash 
find . -type f -name '*.sh' |rev| cut -c4- |cut -f1 -d'/'|rev 
