#!/bin/bash
rm gin-example
go build -o gin-example
nohup gin-example
