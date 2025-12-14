#!/bin/bash

WORKFLOW_TEMPLATE=$(cat .github/dotnet-workflow-template.yml)

mkdir -p .github/workflows

# iterate each dotnet project directory in the dotnet directory
for dir in dotnet/*/; do
	project=$(basename "$dir")
	name=$(echo "$project" | sed 's/\.Api$//' | tr '[:upper:]' '[:lower:]')
	echo "generating workflow for dotnet/${project}"

	# replace template placeholders with project name and workflow name
	WORKFLOW=$(echo "${WORKFLOW_TEMPLATE}" | sed "s/{{name}}/${name}/g" | sed "s/{{project}}/${project}/g")

	# save workflow to .github/workflows/dotnet-{name}
	echo "${WORKFLOW}" > .github/workflows/dotnet-${name}.yml
done