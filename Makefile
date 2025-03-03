# Use either DEBUG or RELEASE
BUILD_TYPE=DEBUG

BINARY_NAME=chibi
BUILD_DIR=build
DEB_PATH=chibi_debian

ifeq (${BUILD_TYPE}, RELEASE)
	LATEST_TAG:=$(shell git describe --tags --abbrev=0)
	LDFLAGS:="-X 'main.VERSION=${LATEST_TAG}' -s -w"
else 
	LATEST_TAG:=$(shell git describe --tags)
	LDFLAGS:="-X 'main.VERSION=${LATEST_TAG}'"
endif

LINUX_BIN=${BINARY_NAME}_${LATEST_TAG}_linux_amd64
WIN_BIN=${BINARY_NAME}_${LATEST_TAG}_win_amd64.exe
WIN_ARM_BIN=${BINARY_NAME}_${LATEST_TAG}_win_arm64.exe
APPLE_INTEL_BIN=${BINARY_NAME}_${LATEST_TAG}_darwin_amd64
APPLE_SILICON_BIN=${BINARY_NAME}_${LATEST_TAG}_darwin_arm64

.PHONY: all

all: clean compile pack_deb

clean:
	if [ -d ${BUILD_DIR} ]; then rm -rf ${BUILD_DIR}; fi
	mkdir ${BUILD_DIR}

compile:
	go mod tidy
	GOARCH=amd64 GOOS=darwin go build -ldflags=${LDFLAGS} -o ${BUILD_DIR}/${APPLE_INTEL_BIN}
	GOARCH=arm64 GOOS=darwin go build -ldflags=${LDFLAGS} -o ${BUILD_DIR}/${APPLE_SILICON_BIN}
	GOARCH=amd64 GOOS=windows go build -ldflags=${LDFLAGS} -o ${BUILD_DIR}/${WIN_BIN}
	GOARCH=arm64 GOOS=windows go build -ldflags=${LDFLAGS} -o ${BUILD_DIR}/${WIN_ARM_BIN}
	GOARCH=amd64 GOOS=linux go build -ldflags=${LDFLAGS} -o ${BUILD_DIR}/${LINUX_BIN}

#TODO: Implement dynamic versioning
pack_deb:
	echo "Packing debian archive"
	mkdir -p ${DEB_PATH}/DEBIAN ${DEB_PATH}/usr/local/bin
	touch ${DEB_PATH}/DEBIAN/control
	cp ${BUILD_DIR}/${LINUX_BIN} ${DEB_PATH}/usr/local/bin/chibi

	@echo "Package: chibi" > ${DEB_PATH}/DEBIAN/control
	@echo "Version: ${LATEST_TAG}" >> ${DEB_PATH}/DEBIAN/control
	@echo "Section: base" >> ${DEB_PATH}/DEBIAN/control
	@echo "Priority: optional" >> ${DEB_PATH}/DEBIAN/control
	@echo "Architecture: amd64" >> ${DEB_PATH}/DEBIAN/control
	@echo "Maintainer: Cosmic Predator" >> ${DEB_PATH}/DEBIAN/control
	@echo "Description: Chibi for AniList" >> ${DEB_PATH}/DEBIAN/control
	@echo "    A lightweight anime & manga tracker CLI app powered by AniList, written in Go." >> ${DEB_PATH}/DEBIAN/control
	@echo "Build-Using: go-1.23 (= 1.23.4)" >> chibi_debian/DEBIAN/control
	@echo "Homepage: https://github.com/CosmicPredator/chibi-cli" >> ${DEB_PATH}/DEBIAN/control

	dpkg-deb --build ${DEB_PATH}
	mv chibi_debian.deb ${BUILD_DIR}/chibi_${LATEST_TAG}_amd64.deb

	rm -rf ${DEB_PATH}/