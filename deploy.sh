#! /bin/sh
rm -rf ./backend_datainfra
#rm -rf ./logs/*
git pull
# go get -u github.com/ndcinfra/backend_datainfra
killall -9 ./backend_datainfra
go build
nohup ./backend_datainfra &
cd logs


