# alcubierre

alcubierre是一套分流传输系统。客户端通过N种不同途径连接到服务器端，形成N个可靠通道（tcp通道），并且将这N个通道互相绑定在一起。此后客户端服务器端各自起一个udp通道，将收到的所有报文从N个通道中的一个发送出去。将N个通道中收到的数据从udp通道中发送出去。以此达到增加传输带宽，减少传输特征的目的。

alcubierre最佳的合作方式，是和openvpn结合，作为openvpn之间的通道。