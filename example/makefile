user			:=	$(shell whoami)
rev				:= 	$(shell git rev-parse --short HEAD)
lastDate		:=	$(shell git log -1 --date=format:"%Y %m %d %T" --format="%ad")
os 				:=	$(shell uname -s)

basePath 		:=	$(shell pwd)
project			:= 	$(shell sed -n 1p go.mod | cut -d ' ' -f2)
output 			:=	$(shell echo "${project}_"`echo ${os} | tr '[A-Z]' '[a-z]'`)
outputDir		:=	"./bin"

ldflags			:=	-ldflags "-s -w -X main.version=${rev} -X \"main.lastDate=${lastDate}\""

remoteHost		:= ""
remoteWorkDir	:= ""
xlsxWorkDir		:= ""

.PHONY: help build_office test

help:
	@echo "help:"
	@echo "  office_build"

office_build:
	@echo "office build"
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${project}_linux ${project}.go
	@scp ${project}_linux ${remoteHost}:${remoteWorkDir}/${project}
	@rm ${project}_linux

test:
	@echo "test"

_xlsx_run:
	@set -e \
		&& echo "${XLSX_WORK_NAME} start" \
		&& $(MAKE) _xlsx_cmd XLSX_CMD="${XLSX_CMD}" \
		&& $(MAKE) _xlsx_work XLSX_WORK_NAME="${XLSX_WORK_NAME}"

_xlsx_cmd:
	@ssh -t ${remoteHost} "cd ${remoteWorkDir}; ./${project} ${XLSX_CMD}"

_xlsx_work:
	@set -e \
		&& scp ${remoteHost}:${remoteWorkDir}/log/xlsx_${XLSX_WORK_NAME} ./output/logs/ \
		&& _date=$(shell date +"%Y-%m-%d"); _xlsxDir=${xlsxWorkDir}/$${_date}; mkdir -p $${_xlsxDir}; echo $${_xlsxDir} \
		&& cat ./output/logs/xlsx_${XLSX_WORK_NAME} | tail -n 1 | xargs -I {} sh -c 'scp "${remoteHost}":"{}" '$${_xlsxDir} \
		&& cat ./output/logs/xlsx_${XLSX_WORK_NAME} | head -n 1 | xargs -I {} open "$${_xlsxDir}/{}" \
		&& cat ./output/logs/xlsx_${XLSX_WORK_NAME} | head -n 1 | xargs -I {} echo "$${_xlsxDir}/{}"

_example:
	@XLSX_WORK_NAME="_example"; \
		$(MAKE) _xlsx_run XLSX_WORK_NAME="$${XLSX_WORK_NAME}" XLSX_CMD="cmd $${XLSX_WORK_NAME} ${ARG}"

