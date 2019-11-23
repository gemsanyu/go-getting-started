#!/bin/bash
case $1 in
    start)
        echo "Starting Application"
        ./sgapp > app.log &
        ;;
    stop)
        echo "Stoping Application"
        sudo kill $(sudo lsof -t -i:8024)
        ;;
    restart)
        echo "Stoping Application"
        sudo kill $(sudo lsof -t -i:8024)
        echo "Starting Application"
        ./sgapp > app.log &
        ;;
    *)
        echo "Golang app service"
        echo $"Usage $0 {start|stop}"
        exit 1
esac
exit 0
