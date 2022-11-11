#!/bin/bash
shellDir=$(cd "$(dirname $0)";pwd)

protoc -I=$shellDir/proto --go_out=$shellDir $shellDir/proto/*.proto