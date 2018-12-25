package mediaserver

/*
#cgo CPPFLAGS: -I${SRCDIR}/external/libsrtp/include
#cgo CPPFLAGS: -I${SRCDIR}/external/openssl/include
#cgo CPPFLAGS: -I${SRCDIR}/external/mp4v2/include
#cgo CPPFLAGS: -I${SRCDIR}/media-server/include
#cgo LDFLAGS: -L${SRCDIR}/media-server/bin/release/ -lmediaserver-${GOOS}-${GOARCH}
#cgo LDFLAGS: -L${SRCDIR}/external/libsrtp -lsrtp2-${GOOS}-${GOARCH}
#cgo LDFLAGS: -L${SRCDIR}/external/openssl -lcrypto-${GOOS}-${GOARCH}
#cgo LDFLAGS: -L${SRCDIR}/external/openssl -lssl-${GOOS}-${GOARCH}
#cgo LDFLAGS: -L${SRCDIR}/external/mp4v2/.libs -lmp4v2-${GOOS}-${GOARCH}
*/
import "C"

func init() {
	MediaServerInitialize()
}

func EnableLog(flag bool) {
	MediaServerEnableLog(flag)
}

func EnableDebug(flag bool) {
	MediaServerEnableDebug(flag)
}

func SetPortRange(minPort, maxPort int) bool {
	return MediaServerSetPortRange(minPort, maxPort)
}

func EnableUltraDebug(flag bool) {
	MediaServerEnableUltraDebug(flag)
}
