#!/usr/bin/env bash

repoPath="repo"
ReNameCode="ShubNig"
VERSION_MAJOR=0
VERSION_MINOR=0
VERSION_PATCH=1
VERSION_BUILD=0

VersionCode=$[$[VERSION_MAJOR * 100000000] + $[VERSION_MINOR * 100000] + $[VERSION_PATCH * 100] + $[VERSION_BUILD]]
VersionName="${VERSION_MAJOR}.${VERSION_MINOR}.${VERSION_PATCH}.${VERSION_BUILD}"
packageReName="${ReNameCode}_${VersionName}"

shell_running_path=$(cd `dirname $0`; pwd)

echo -e "============\nPrint build info start"
go version
which go
echo -e "Your settings is
\tVersion Name -> ${ReNameCode}
\tVersion code -> ${VersionCode}
\tVersion name -> ${VersionName}
\tPackage rename -> ${packageReName}
\tOut Path -> ${shell_running_path}
"
echo -e "Print build info end\n============"

echo "start build OSX 64"
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
mv main "${shell_running_path}/${packageReName}_osx_64"
echo "build OSX 64 finish"