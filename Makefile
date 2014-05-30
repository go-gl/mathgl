all:	
	mkdir -p mgl64
	cp mgl32/* mgl64
	gofmt -w -r "float32 -> float64" mgl64/*.go
	gofmt -w -r "a.Float32 -> a.Float64" mgl64/*.go
	gofmt -w -r "mgl32 -> mgl64" mgl64/*.go
	go fmt ./...
