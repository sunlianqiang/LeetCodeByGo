# to replace your project protofile
PROTO_FILE = ucs.201000.203000.proto 
# to replace your proto packagename
PROTO_PACKAGE_NAME = ucs
UGC_MESSAGE_PATH = ../ugc-message/proto

PROTO_GO_FILE = $(patsubst %.proto,%.pb.go,$(PROTO_FILE)) 
PROTO_PATH = proto/$(PROTO_PACKAGE_NAME)

SRCDIRS_EXCLUDE = log test utils proto scripts vendor Godeps
SRCDIRS_ALL = $(sort $(subst /,,$(dir $(wildcard */*))))
SRCDIRS = $(filter-out $(SRCDIRS_EXCLUDE), $(SRCDIRS_ALL))

PKGDIRS_EXCLUDE=$(GOROOT)/pkg
PKGDIRS_ALL = $(addsuffix /pkg, $(subst :, ,$(GOPATH)))
PKGDIRS = $(filter-out $(PKGDIRS_EXCLUDE), $(PKGDIRS_ALL))

GOLDFLAG ?=

VERSION_GIT_COMMIT = $(shell git rev-parse --short HEAD)

VERFLAG := \
	-X ucs-api-gateway/version.gitCommit=$(VERSION_GIT_COMMIT)

GOLDFLAG := $(GOLDFLAG) $(VERFLAG)

all: build_main 
	@for subdir in $(SRCDIRS);do \
		cd $$subdir; go install; cd ..; \
	done 

parse_proto:
	@if [ ! -d $(UGC_MESSAGE_PATH) ]; \
	then \
		echo "ugc-message does not exsit, please clone ugc-message repository at your upper dir"; \
		exit 1; \
	fi 
	@mkdir -p $(PROTO_PATH); \
	cp $(UGC_MESSAGE_PATH)/$(PROTO_FILE) $(PROTO_PATH); \
	frame_work_path=`echo $(CURDIR)|sed "s/src.*/src\/uframework/"`; \
	cd $(PROTO_PATH); \
	protoc --proto_path=$$frame_work_path/message/proto/ucloud --proto_path=. --go_out=. $(PROTO_FILE); \
	sed -i "s/^.*package ucloud.*$$/package $(PROTO_PACKAGE_NAME)/" $(PROTO_GO_FILE); \
	sed -i "s/^.*import ucloud.*$$/import ucloud \"uframework\/message\/proto\/ucloud\"/" $(PROTO_GO_FILE);\
	go install; \
	cd $(CURDIR); 



build_main: parse_proto
	@CGO_ENABLED=0 GO15VENDOREXPERIMENT=1 godep go build -ldflags "$(GOLDFLAG)" ucs-api-gateway.go;

show:
	@echo "==================src====================="
	@echo SRCDIRS_ALL: $(SRCDIRS_ALL)
	@echo SRCDIRS_EXCLUDE: $(SRCDIRS_EXCLUDE)
	@echo SRCDIRS: $(SRCDIRS)
	@echo "==================pkg====================="
	@echo PKGDIRS_EXCLUDE: $(PKGDIRS_EXCLUDE)
	@echo PKGDIRS_ALL: $(PKGDIRS_ALL)
	@echo PKGDIRS: $(PKGDIRS)
	@echo "================clean====================="
	@for subdir in $(PKGDIRS); do \
		cd $$subdir;\
		module_name=`echo $(CURDIR)|awk -F"/" '{print $$(NF)}'`;\
		result=`find . |grep $$module_name |head -n1|awk -F"." '{print $$2}'`; \
		if [ -n "$$result" ];then \
			echo clean_dirs:$$subdir$$result; \
		fi \
	done

clean:
	@for subdir in $(PKGDIRS); do \
		cd $$subdir;\
		module_name=`echo $(CURDIR)|awk -F"/" '{print $$(NF)}'`;\
		result=`find . |grep $$module_name|head -n1|awk -F"." '{print $$2}'`; \
		if [ -n "$$result" ];then \
			cd $$subdir$$result; rm -rf *; \
		fi \
	done

clean_go:
	@for subdir in $(PKGDIRS); do \
		cd $$subdir;\
		go clean;\
	done
