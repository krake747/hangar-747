#!/bin/bash

WORKFLOW_TEMPLATE=$(cat .github/go-workflow-template.yml)

mkdir -p .github/workflows

# iterate each go package directory in the golang directory
for dir in golang/*/; do
	package=$(basename "$dir")
	echo "generating workflow for golang/${package}"

	# replace template package placeholder with package name
	WORKFLOW=$(echo "${WORKFLOW_TEMPLATE}" | sed "s/{{package}}/${package}/g")

	# save workflow to .github/workflows/{package}
	echo "${WORKFLOW}" > .github/workflows/${package}.yml
done
