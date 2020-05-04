#!/bin/bash

source "$(dirname $0)/base.sh"
info "[$0]"
(
    set -e
    cd $ROOT_DIR/go/
    go run main.go
)
if [ $? == 0 ] ; then
    info "[$0 OK]"
else
    error "[$0 FAILED]"
fi