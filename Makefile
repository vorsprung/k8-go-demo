# Project-specific variables
CONVEY_PORT ?=	9042


# Common variables
SOURCES :=	$(shell find . -type f -name "*.go")
COMMANDS :=	$(shell go list ./... | grep -v /vendor/ | grep /cmd/)
PACKAGES :=	$(shell go list ./... | grep -v /vendor/ | grep -v /cmd/)
GO ?=		$(GOENV) go
GODEP ?=	$(GOENV) godep
DOCKERUSER =		"vorsprung"



ifndef XKEY
$(error XKEY is not set, should have a apikey)
endif

.PHONY: test
test:
	$(GO) get -t .
	$(GO) test -v .

.PHONY: convey
convey:
	$(GO) get github.com/smartystreets/goconvey
	goconvey -cover -port=$(CONVEY_PORT) -workDir="$(realpath .)" -depth=1


.PHONY:	cover
cover:	profile.out


profile.out:	$(SOURCES)
	rm -f $@
	$(GO) test -covermode=count -coverpkg=. -coverprofile=$@ .

.PHONY: browsercover
browsercover: test cover
	$(GO) tool cover -html=profile.out

.PHONY: image
image:
	docker build --build-arg apikey=$(XKEY) -t $(DOCKERUSER)/stock .
	docker push $(DOCKERUSER)/stock

.PHONY: secret
secret:
	python secret_file.py | kubectl apply

k8: secret
	kubectl apply -f stock.yml
	-kubectl expose deployment stockapp --type=NodePort --name=stockapp-service

k8test: k8
	$(eval URL = $(shell minikube service stockapp-service --url))
	curl $(URL)?symbol=ORCL&ndays=3
	@echo

k8delete:
	kubectl delete stockapp-service
	rm -f secret.yml s2.txt
