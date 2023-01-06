ARM `ror` (rotate right) operations are not supported. According to the [docs],
the syntax used for `ror` is `@>`, which does not compile. A logical right
shift `lsr` does work using the `>>` operator.

```
go version go1.19.4 darwin/arm64
```

```
GO111MODULE=""
GOARCH="arm64"
GOBIN=""
GOCACHE="/Users/dadrian/Library/Caches/go-build"
GOENV="/Users/dadrian/Library/Application Support/go/env"
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="arm64"
GOHOSTOS="darwin"
GOINSECURE=""
GOMODCACHE="/Users/dadrian/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="darwin"
GOPATH="/Users/dadrian/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/opt/homebrew/Cellar/go/1.19.4/libexec"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/opt/homebrew/Cellar/go/1.19.4/libexec/pkg/tool/darwin_arm64"
GOVCS=""
GOVERSION="go1.19.4"
GCCGO="gccgo"
AR="ar"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD="/Users/dadrian/src/github.com/dadrian/go-bugs/go.mod"
GOWORK=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -arch arm64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/7_/yc9s7tp94lq83r7fmlncg58c0000gn/T/go-build2852821403=/tmp/go-build -gno-record-gcc-switches -fno-common"
```
