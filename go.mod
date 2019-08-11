module goTestProject

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3 => /src/golang.org/x/net
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c => /src/golang.org/x/net
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859 => /src/golang.org/x/net
	golang.org/x/sync v0.0.0-20190423024810-112230192c58 => /src/golang.org/x/sync
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190804053845-51ab0e2deafa
	golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223 => /src/golang.org/x/sys
	golang.org/x/text v0.3.0 => /src/golang.org/x/text
	golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e => /src/golang.org/x/tools
	golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7 => /src/golang.org/x/xerrors
)

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/sirupsen/logrus v1.4.2
)
