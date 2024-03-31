#!/bin/sh

directory=$(dirname "$0")

mkdir -p $directory/../generated/code-help-core
mkdir -p $directory/../generated/code-help-forum
mkdir -p $directory/../generated/code-help-admin

$directory/../pkg/oapi-codegen.sh -generate types,client,spec -package codeHelpCoreGen $directory/../code-help-core-api.yaml > $directory/../generated/code-help-core/code-help-core.go
$directory/../pkg/oapi-codegen.sh -generate types,client,spec -package codeHelpForumGen $directory/../code-help-forum-api.yaml > $directory/../generated/code-help-forum/code-help-forum.go
$directory/../pkg/oapi-codegen.sh -generate types,std-http,spec -package codeHelpAdminGen $directory/../code-help-admin-api.yaml > $directory/../generated/code-help-admin/code-help-admin.go