
BUILD		:= release
VERSION		:= v1.0.4
REVISION	:= $(shell git rev-parse --short HEAD)

SRCS		:= $(shell find ./src -type f -name '*.go')
VSRCS		:= $(shell find ./vendor -type f -name '*.go')
VSRC_EXIST	:= $(shell find ./vendor -name src)
LDFLAGS		:= -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\""

define vendor_restore
	@gb vendor restore 
endef

BuildDeb = mkdir -p work/opt/ir-http; \
		   mkdir -p work/DEBIAN ; \
		   cp ./bin/$(1) ./work/opt/ir-http/ir-http; \
		   cp -Rp ./cmd ./work/opt/ir-http; \
		   cp ./config.json ./work/opt/ir-http; \
		   cp ./deb/post* ./work/DEBIAN/ ; \
		   sed -e 's/%%DBG%%/$(2)/' -e 's/%%ARCH%%/$(3)/' ./deb/conf > ./work/DEBIAN/control ; \
		   fakeroot dpkg-deb --build ./work . ; \
		   rm -rf ./work

BuildCommand = GOARCH=$(1) GOOS=linux gb build -tags '$(2)' $(3) all 

ifeq ($(BUILD),debug)
TAGS	:= debug
IS_DBG	:= -debug
else
TAGS	:= 
IS_DBG	:= 
endif

.PHONY: all
all: amd64

.PHONY: clean
clean:
	@rm -rf bin
	@rm -rf pkg
	@rm -rf *.deb

all-clean:
	@rm -rf bin
	@rm -rf pkg
	@rm -rf vendor/src
	@rm -rf *.deb

setup :
	$(if $(VSRC_EXIST) ,,$(vendor_restore))
386: $(SRCS)
	$(if $(VSRC_EXIST) ,,$(vendor_restore))
	$(call BuildCommand,386,$(TAGS),$(LDFLAGS))

arm: $(SRCS)
	$(if $(VSRC_EXIST) ,,$(vendor_restore))
	$(call BuildCommand,arm,$(TAGS),$(LDFLAGS))

arm64: $(SRCS)
	$(if $(VSRC_EXIST) ,,$(vendor_restore))
	$(call BuildCommand,arm64,$(TAGS),$(LDFLAGS))

amd64: $(SRCS)
	$(if $(VSRC_EXIST) ,,$(vendor_restore))
	$(call BuildCommand,amd64,$(TAGS),$(LDFLAGS))
#	@mv bin/main$(IS_DBG) bin/main-linux-amd64$(IS_DBG)

deb: deb-amd64

deb-386: ./bin/main-linux-386$(IS_DBG)
	$(call BuildDeb,main-linux-386$(IS_DBG),$(IS_DBG),i386)

deb-amd64: ./bin/main-linux-amd64$(IS_DBG)
	$(call BuildDeb,main-linux-amd64$(IS_DBG),$(IS_DBG),amd64)

deb-armhf: ./bin/main-linux-arm$(IS_DBG)
	$(call BuildDeb,main-linux-arm$(IS_DBG),$(IS_DBG),armhf)

deb-arm64: ./bin/main-linux-arm64$(IS_DBG)
	$(call BuildDeb,main-linux-arm64$(IS_DBG),$(IS_DBG),arm64)

archive:
	@git archive HEAD --output=src.zip
