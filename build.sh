DIR=$(dirname $0)
PWD=$(pwd)
CURRPATH=$PWD/$DIR
cd $CURRPATH

export GOPATH=$CURRPATH:$GOPATH
go build iam/cmd/iam-server
