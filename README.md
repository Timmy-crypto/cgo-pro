provide cmd and server to test cgo

generate lib:

    make buildLib


generate cmd and server
    
    make buildBin
    
run server
    
    cd ./bin
    LD_LIBRARY_PATH=../lib ./server
    
run cmd

    cd ./bin
    ./cmd
    
cmd example
    rpc -m CgoTestF5
    rpc -m CgoTestCxxSum -p 5,1,2,3,4,5
    