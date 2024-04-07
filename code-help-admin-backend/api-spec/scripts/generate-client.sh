#!/bin/sh

directory=$(dirname "$0")

mkdir -p $directory/../generated/code-help-forum
mkdir -p $directory/../generated/code-help-admin
mkdir -p $directory/../generated/code-help-admin-core

$directory/../pkg/oapi-codegen.sh -generate types,client,spec -package codeHelpAdminCoreGen $directory/../code-help-admin-core-api.yaml > $directory/../generated/code-help-admin-core/code-help-admin-core.go
$directory/../pkg/oapi-codegen.sh -generate types,client,spec -package codeHelpForumGen $directory/../code-help-forum-api.yaml > $directory/../generated/code-help-forum/code-help-forum.go
$directory/../pkg/oapi-codegen.sh -generate types,std-http,spec -package codeHelpAdminGen $directory/../code-help-admin-api.yaml > $directory/../generated/code-help-admin/code-help-admin.go