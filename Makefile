all:	
	mkdir mgl64
	cp -R mgl32/* mgl64
	gofmt -w -r "float32 -> float64" mgl64/*.go
	gofmt -w -r "a.Float32 -> a.Float64" mgl64/*.go
	gofmt -w -r "mgl32 -> mgl64" mgl64/*.go
	gofmt -w -r "math.MaxFloat32 -> math.MaxFloat64" mgl64/*.go
	go fmt ./...
	cp mgl32/matstack/* mgl64/matstack
	gofmt -w -r "float32 -> float64" mgl64/matstack/*.go
	gofmt -w -r "mgl32 -> mgl64" mgl64/matstack/*.go
	goimports -w=true mgl64/matstack/*.go
