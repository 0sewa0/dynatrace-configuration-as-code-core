#!/bin/sh

export GO_POST_PROCESS_FILE="gofmt -w"
openapi-generator generate -i https://api.dynatrace.com/spec-json -g go -t template -o account_management --remove-operation-id-prefix --enable-post-process-file --additional-properties=packageName=accountmanagement
openapi-generator generate -i tenant_specs/environment_api.json -g go -t template -o environment --remove-operation-id-prefix --enable-post-process-file --additional-properties=packageName=environment,withGoMod=false,generateInterfaces=true

#/components/schemas/EntityShortRepresentation
#/components/schemas/EventDto
