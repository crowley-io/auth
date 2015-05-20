GO_PACKAGE=github.com/crowley-io/auth

setup:
	go get -d -v ./...

setup-test: setup
	go get -d -t -v ./...

test:	setup-test
	go test $(GO_PACKAGE)/ssh

coverage: ssh/cover.out mfa/cover.out

	@echo "mode: set" > $@ && cat $^ 2>/dev/null | grep -v mode: | sort -r | \
		awk '{if($$1 != last) {print $$0;last=$$1}}' >> $@
	go tool cover -html=$@ -o $@.html
	@rm $^ 2>/dev/null || true

# cover.out: setup-test
# 	go test -coverprofile=$@ -coverpkg=./... $(GO_PACKAGE) 2>/dev/null

ssh/cover.out: setup-test
	go test -coverprofile=$@ $(GO_PACKAGE)/ssh

mfa/cover.out: setup-test
	go test -coverprofile=$@ $(GO_PACKAGE)/mfa