DIR=$(dirname $0)
PWD=$(pwd)
CURRPATH=$PWD/$DIR
cd $CURRPATH

./iam-server --host=0.0.0.0 --port=8089
