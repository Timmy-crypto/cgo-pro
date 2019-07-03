server and cmd

run server
    
    cd ./bin
    //默认监听ip未本地环回地址，无法通过真实ip访问
    LD_LIBRARY_PATH=../lib ./server
    
    // 设置sever ws_host ip 为0.0.0.0则，可通过真实ip访问
    // 否则只监听的　127.0.0.1本地环回地址，无法通过真实ip访问
    LD_LIBRARY_PATH=../lib/ ./server --ws_host 0.0.0.0
    
    // 设置本地监听地址为实际ip ,本机无法通过环回地址访问
    LD_LIBRARY_PATH=../lib/ ./server --ws_host $(actual_ip)
    
run cmd

    cd ./bin
    
    //默认server_ip 为127.0.0.1
    ./cmd
    
    //指定server_ip为服务端实际ip
    ./cmd --server_ip $(server-ip)
    
cmd example

    rpc -m CgoTestF5
    rpc -m CgoTestCxxSum -p 5,1,2,3,4,5