CC=gcc
CXX=g++
CLibPath=./server/clib
CXXLibPath=./server/C++lib

buildLib:
	#buildCLib
	$(CC) -g -fPIC -c -o $(CLibPath)/lib/lib.o $(CLibPath)/lib/lib.c
	$(CC) -g -fPIC -shared -o $(CLibPath)/lib/liblib.so $(CLibPath)/lib/lib.o
	rm $(CLibPath)/lib/lib.o
	#LD_LIBRARY_PATH=. go run main.go
	cp $(CLibPath)/lib/liblib.so $(CLibPath)/wrapper/
	cp $(CLibPath)/lib/liblib.so ./lib/

	#buildCPlusPlusLib
	$(CXX) -g -fPIC -c $(CXXLibPath)/lib/wrapper.cpp $(CXXLibPath)/lib/cplusplus.cpp
	$(CXX) -g -fPIC -shared -o $(CXXLibPath)/lib/libwrapper.so wrapper.o cplusplus.o
	cp $(CXXLibPath)/lib/libwrapper.so $(CXXLibPath)/wrapper/
	cp $(CXXLibPath)/lib/libwrapper.so ./lib/
	rm wrapper.o
	rm cplusplus.o

buildBin:
	cd ./server && go build && mv ./server ../bin/
	cd ./cmd && go build && mv ./cmd ../bin/

TestCgoWrapper:
	cd $(CLibPath) && LD_LIBRARY_PATH=./lib go run main2.go

clean:
	rm $(CLibPath)/lib/liblib.so
	rm $(CLibPath)/wrapper/liblib.so
	rm ./bin/*
	rm ./lib/*

package:
	export GO111MODULE=on && go mod tidy && go mod vendor



