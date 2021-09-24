#!/bin/sh
### BEGIN INIT INFO
# Provides:          vcloud-inet
# Required-Start:    $local_fs $network $named $time $syslog
# Required-Stop:    $local_fs $network $named $time $syslog
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Description:       vcloud-inet
### END INIT INFO

SCRIPT=/usr/bin/vcloud-inet
RUNAS=root

LOGFILE=/var/log/vcloud-inet.log

start() {
  echo 'Starting service…' >&2
  local CMD="$SCRIPT &> \"$LOGFILE\" & echo \$!"
  su -c "$CMD" $RUNAS
  echo 'Service started' >&2
}

stop() {
  echo 'Stopping service…' >&2
}

case "$1" in
  start)
    start
    ;;
  stop)
    stop
    ;;
  *)
esac
