#!/bin/sh
# フォーマット
echo "gofmt..."
## フォーマットしたかどうか
formatted=false
dirs=("api/" "config/" "controllers/" "db/" "entity/" "models/" "oauth/" "tools/" "utils/")
for dir in ${dirs[@]}
do
    # フォーマット違いがあるものだけフォーマットする
    formatDiffFileCount=`gofmt -l $dir | wc -l`
    if [ $formatDiffFileCount -gt 0 ] ; then
        gofmt -w $dir
        formatted=true
    fi
done

if [ "$formatted" = true ] ; then
    echo "フォーマットしました。commitしてください。 git commit -am 'gofmt' "
else
    echo "フォーマットする必要ありませんでした。"
fi
echo "done"

# ビルド
echo -n "go build..."
cd api
if [ -e ./api ] ; then
    rm -f ./api
fi
go build
echo "done"

# 実行
echo "run"
./api -http
