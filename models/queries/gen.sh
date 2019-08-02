(cd gen && rm *.go)
cp pkg.go.template gen/pkg.go
cp verboseBlock.go.template gen/verboseBlock.go
go test -v && (cd gen && go build)
