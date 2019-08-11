module goTestProject

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190804053845-51ab0e2deafa
)

require github.com/sirupsen/logrus v1.4.2
