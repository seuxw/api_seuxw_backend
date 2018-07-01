#!/bin/bash -e

# set some environment variables
readonly SEUXW_ROOT=$(cd $(dirname "${BASH_SOURCE}")/.. && pwd -P)
readonly SEUXW_OUTPUT="${SEUXW_ROOT}/_output/local"
readonly SEUXW_OUTPUT_SRCPATH="${SEUXW_OUTPUT}/src"
readonly SEUXW_OUTPUT_BINPATH="${SEUXW_OUTPUT}/bin"

readonly SEUXW_TARGETS=(
	filter/test
  )

eval $(go env)

# enable/disable failpoints
toggle_failpoints() {
	FAILPKGS="seuxwserver/ seuxwserver/auth/"

	mode="disable"
	if [ ! -z "$FAILPOINTS" ]; then mode="enable"; fi
	if [ ! -z "$1" ]; then mode="$1"; fi

	if which gofail >/dev/null 2>&1; then
		gofail "$mode" $FAILPKGS
	elif [ "$mode" != "disable" ]; then
		echo "FAILPOINTS set but gofail not found"
		exit 1
	fi
}

seuxw_setup_gopath() {
	# preserve old gopath to support building with unvendored tooling deps (e.g., gofail)
	if [ -n "$GOPATH" ]; then
		GOPATH=":$GOPATH"
	fi
	export GOPATH=${SEUXW_OUTPUT}

	rm -rf ${SEUXW_OUTPUT_SRCPATH}
	mkdir -p ${SEUXW_OUTPUT_SRCPATH}

	ln -s ${SEUXW_ROOT} ${SEUXW_OUTPUT_SRCPATH}/seuxw
}

seuxw_build_target() {
	toggle_failpoints

	for arg; do
		# echo "target: ${arg}, ${arg##*/}"
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $GO_BUILD_FLAGS \
		-installsuffix cgo -ldflags "$GO_LDFLAGS" \
		-o ${SEUXW_OUTPUT_BINPATH}/${arg##*/}.x seuxw/${arg} || return
	done
}

seuxw_make_ldflag() {
  local key=${1}
  local val=${2}

  echo "-X seuxw/filter/version.${key}=${val}"
}

# Prints the value that needs to be passed to the -ldflags parameter of go build
# in order to set the project on the git tree status.
seuxw_version_ldflags() {
	local -a ldflags=($(seuxw_make_ldflag "buildDate" "$(date -u +'%Y-%m-%dT%H:%M:%SZ')"))

	local git_sha=`git rev-parse --short HEAD || echo "GitNotFound"`
	if [ ! -z "$FAILPOINTS" ]; then
		git_sha="$git_sha"-FAILPOINTS
	fi

	ldflags+=($(seuxw_make_ldflag "gitSHA" "${git_sha}"))

	echo "${ldflags[*]-}"
}

toggle_failpoints

# only build when called directly, not sourced
if echo "$0" | grep "build.sh$" >/dev/null; then
	# force new gopath so builds outside of gopath work
	seuxw_setup_gopath
	seuxw_version_ldflags
	#seuxw_build
	seuxw_build_target "${SEUXW_TARGETS[@]}"
fi
