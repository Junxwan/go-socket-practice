# Golang Socket

    go run service.og 

    go run client.og 

client:

    Connect success ....
    >Please input data
    client   # enter
    >success
    >Please input data

service:

    Start listen
    Connect : 127.0.0.1:60295
    Receive 'client' from '127.0.0.1:60295'

## net.Listen()

監聽某個tcp,utp,id prot

## net.Dial()

連接到某個tcp,utp,id的prot

## Conn
透過net.Dial建立跟service連線後會得到一個Conn，該interface實作了兩個func

    // 讀取service傳遞過來的資料
    Read(b []byte) (n int, err error)

    // 發送資料到service
    Write(b []byte) (n int, err error)

## listener.Accept()
等待同意Client的連線

## net.ResolveTCPAddr
將ip:prot解析成Tcp Addr，該interface實作了兩個func

    // 是tcp or utp
    Network() string
    
    // ip
    String() string  