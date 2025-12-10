# Manage password daatbase and online users for ocserv

This solution assume deploy on same server with ocserv
Use basic auth to access web interface. Password for authorization store in passwd file in clear text (TODO rewrite to use bcrypt)

## usage

```
% ./ocserv-admin -h
Usage of ./ocserv-admin:
  -ocpasswd string
	ocserv password file (default "/etc/ocserv/ocpasswd")
  -passwd string
	local auth file (default "passwd")
```

## build
```
buildah build -t ocserv-admin:`cat .tag`
```

## run
```
podman run -d --restart always -p 80:80 -v /etc/ocserv/ocpasswd:/app/ocpasswd --name ocserv-admin localhost/ocserv-admin:`cat .tag`
```

## online users
```
root@vpn:~# occtl show users
      id     user    vhost             ip         vpn-ip device   since    dtls-cipher    status
  171536 user1  default  1.1.1.1 192.168.254.18   vpns10  1h:28m      (no-dtls) connected
  171516 user2  default   1.1.1.2 192.168.254.24  vpns7   2h:01m      (no-dtls) connected
  171369 user3  default  1.1.1.3 192.168.254.220  vpns3   3h:45m      (no-dtls) connected
  171316 user4  default   1.1.1.4 192.168.254.173 vpns4   4h:50m      (no-dtls) connected
```

```
# occtl show id 171516
        ID: 171516
        Username: user4   Groupname: admin
        State: connected
        vhost: default
        Device: vpns4   MTU: 1434
        Remote IP: 1.1.1.4   Location: unknown
        Local Device IP: 192.168.40.5
        IPv4: 192.168.254.24   P-t-P IPv4: 192.168.254.1
        User-Agent: Open AnyConnect VPN Agent v8.20-1
        RX: 1651643 (1.7 MB)   TX: 25270855 (25.3 MB)
        Average bandwidth RX: 164 bytes/s  TX: 2.5 kB/s
        DPD: 90   KeepAlive: 32400
        Hostname: user-laptop
        Connected at: 2025-12-04 13:44 ( 2h:47m)
        Session: 0UXOAK
        TLS ciphersuite: (TLS1.2)-(ECDHE-SECP256R1)-(RSA-SHA256)-(AES-256-GCM)

        DNS: 192.168.41.41
        Split-DNS-Domains: ritm.lan
        Routes: 192.168.9.0/255.255.255.0
                192.168.20.0/255.255.255.0
                192.168.26.0/255.255.255.0
                192.168.100.20/255.255.255.0
                192.168.41.41/255.255.255.255
        Restricted to routes: False
```
