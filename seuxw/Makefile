all: build

vet:
	@ for f in $$(find . -type f -iregex '.*\.go' | grep -v "vendor"); do \
		go fmt $$f; \
	done 

	@ for f in $$(find . -type f -iregex '.*\.go'  | grep -v "vendor"); do \
		go tool vet $$f; \
	done 

build:
	bash/build.sh $$func $$os

gen:
	@ if ! which thrift > /dev/null; then \
		echo "error: thrift not installed" >&2; \
		exit 1; \
	fi

	@ for f in $$(find proto -type f -iregex '.*\.thrift'); do \
		thrift -out proto -r --gen go $$f; \
	done

	@ for d in $$(find proto/* -type 'd' | grep "_sevice-remote"); do \
		rm -rf $$d; \
	done 
	
	@ for f in $$(find . -type f -i
		echo "error: thrift not installed" >&2; \
		exit 1; \
	fi

	@ for f in $$(find proto -type f -iregex '.*\.thrift'); do \
		thrift -out proto -r --gen go $$f; \
	done

	@ for d in $$(find proto/* -type 'd' | grep "_sevice-remote"); do \
		rm -rf $$d; \
	done 
	
	@ for f in $$(find . -type f -iregex '.*\.go'  | grep 'consts.go\|GoUnusedProtection__.go'); do \
		rm -f $$f; \
	done

clean:
	@ #go clean -i seuxw.cc/...
	@ rm -rf _output
	
	@ for d in $$(find proto/* -type 'd' | grep -v "thrift"); do \
		rm -rf $$d; \
	done 