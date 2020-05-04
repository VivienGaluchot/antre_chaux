#!/bin/bash

source "$(dirname $0)/base.sh"
info "[$0]"
(
    set -e
    cd $ROOT_DIR/go/

    PROMOTE="--no-promote"
    if [ "$1" == "--promote" ] ; then
        PROMOTE="--promote"
    fi

    info "Deploy with $PROMOTE"
    gcloud --project=$PROJECT_NAME app deploy $PROMOTE
    echo ""

    info "Done, go to https://console.cloud.google.com/ to try, promote the deployed version and delete old one."
)
if [ $? == 0 ] ; then
    info "[$0 OK]"
else
    error "[$0 FAILED]"
fi