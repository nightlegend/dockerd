# dockerd
Docker daemon api


## Configure in your docker daemon

```
cd /etc/systemd/system/docker.service.d/
touch tcp.conf
echo
"[Service]
ExecStart=
ExecStart=/usr/bin/docker daemon -H tcp://0.0.0.0:2375 -H unix:///var/run/docker.sock"
>> tcp.conf
systemctl daemon-reload
systemctl restart docker

```

## Clone source
```
git clone https://github.com/nightlegend/dockerd.git into your workdir

```

## Get package and build

```
go get
go build

```

## Start

```
go run main.go

```
