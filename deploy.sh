#!/usr/bin/env sh

# 确保脚本抛出遇到的错误
set -e


push_addr=`git remote get-url --push origin`
commit_info=`git describe --all --always --long`
dist_path=docs/.vuepress/dist 
push_branch=gh-pages

# 生成静态文件
yarn build:win

# 进入生成的文件夹
cp -r $dist_path ./

git add -A
git commit -m "deploy, $commit_info"
git push -f $push_addr

rm -rf $dist_path