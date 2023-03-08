#!/bin/bash
# @file tests.sh
# @brief Start and stop the docker stack containing the Jira runtime.
#
# @description The script starts and stops the docker compose stack used for the local jira
# instance. The ``local/tron:test`` image is build before starting Jira.
#
# === Script Arguments
#
# * *$1* (string): Command (``start``, ``logs``, ``stop``)
#
# === Script Example
#
# [source, bash]
# ```
# ./tests.sh start
# ./tests.sh logs
# ./tests.sh stop
# ```


readonly START="start"
readonly LOGS="logs"
readonly STOP="stop"

ARG="$1"
readonly ARG

# Include local bash modules
source "../bash-modules/log.sh"


if [ -z "$ARG" ]; then
  LOG_ERROR "Param missing: command (start | logs | stop)"
  LOG_ERROR "exit" && exit 0
fi


set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace


case "$ARG" in
  "$START" )
    LOG_HEADER "Start stack"
    docker-compose up --build -d
  ;;
  "$LOGS" )
    LOG_HEADER "Show logs"
    docker-compose logs
  ;;
  "$STOP" )
    LOG_HEADER "Stop stack"
    docker-compose down -v --rmi all --remove-orphans
  ;;
esac
