OUT        = ../../pkg/client
SWAGGER    = ../api/swagger
INCLUDEDIR = $(CURDIR)/includes
PROTOS     = $(shell find . -type f -name '*.proto')
APP_NAME   = srv-users

# Proto Dependencies
#PROTODEPS += github.com/gogo/protobuf
#PROTODEPS += github.com/grpc-ecosystem/grpc-gateway

all: codegen

.PHONY: codegen
codegen: $(PROTODEPS) | $(PROTOS)

# Links dependencies into includes directory
.PHONY: $(PROTODEPS)
$(PROTODEPS): %:
ifeq ($(wildcard $(INCLUDEDIR)/$*),)
	$(info linking $*)
	mkdir -p $(INCLUDEDIR)/$*
	ln -s  $(shell go list -m -f '{{ .Dir }}' $*)/* $(INCLUDEDIR)/$*
endif

# Generate protocol buffer files
.PHONY: $(PROTOS)
$(PROTOS): %.proto:
	@mkdir -p $(OUT)
	@protoc \
  	-I . \
	--go_out=plugins=grpc:$(OUT) \
	--govalidators_out=gogoimport=true:$(OUT) \
	$*.proto
	cd $(OUT) && $(MAKE) all

