# Сreated by Catborisovv (c) 2020-2024
# MakeFile fro my PetAPI

# build for mac m1 (arm-silicon)
build-mac-arm64:
	go env GOARCH=arm64
	go env GOOS=darwin
	go build cmd/app/main.go

# build for mac amd64
build-mac-amd64: 
	go env GOARCH=amd64
	go env GOOS=darwin
	go build cmd/app/main.go	

# build for linux-arm
build-linux-arm64:
	go env GOARCH=arm64
	go env GOOS=linux
	go build cmd/app/main.go	

# build for linux-amd64 
build-linux-amd64:
	go env GOARCH=amd64
	go env GOOS=linux
	go build cmd/app/main.go