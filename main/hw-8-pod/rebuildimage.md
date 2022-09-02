### build a docker image, push to docker hub

oliver@ASUS-Desktop-Win11:~/hw3$ sudo su -
Welcome to Ubuntu 20.04.4 LTS (GNU/Linux 5.10.16.3-microsoft-standard-WSL2 x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

  System information as of Fri Sep  2 11:38:03 PDT 2022

  System load:  0.0                Processes:                17
  Usage of /:   0.7% of 250.98GB   Users logged in:          0
  Memory usage: 1%                 IPv4 address for docker0: 172.17.0.1
  Swap usage:   0%                 IPv4 address for eth0:    172.25.216.167

116 updates can be applied immediately.
69 of these updates are standard security updates.
To see these additional updates run: apt list --upgradable

root@ASUS-Desktop-Win11:~# cd /home/oliver/hw3/

root@ASUS-Desktop-Win11:/home/oliver/hw3# ls
Dockerfile  main.go

root@ASUS-Desktop-Win11:/home/oliver/hw3# docker build -t httpserver:v1.0 -f Dockerfile .
Sending build context to Docker daemon   5.12kB
Step 1/9 : FROM golang:1.17 AS builder
1.17: Pulling from library/golang
d836772a1c1f: Pull complete
66a9e63c657a: Pull complete
d1989b6e74cf: Pull complete
c28818711e1e: Pull complete
9d6246ba248c: Pull complete
21d43f0d73c2: Pull complete
d8a1c5873f40: Pull complete
Digest: sha256:87262e4a4c7db56158a80a18fefdc4fee5accc41b59cde821e691d05541bbb18
Status: Downloaded newer image for golang:1.17
 ---> 742df529b073
Step 2/9 : ENV GO111MODULE=off  CGO_ENABLED=0   GOOS=linux      GOARCH=amd64
 ---> Running in 12559cb45124
Removing intermediate container 12559cb45124
 ---> 4e9d391beafd
Step 3/9 : WORKDIR /build
 ---> Running in 3b46088e7169
Removing intermediate container 3b46088e7169
 ---> 2f45ac4d24cd
Step 4/9 : COPY . .
 ---> fecbdf4e0cb2
Step 5/9 : RUN go build -o httpserver .
 ---> Running in c5101e5d8700
Removing intermediate container c5101e5d8700
 ---> 537ba83158d8
Step 6/9 : FROM scratch
 --->
Step 7/9 : COPY --from=builder /build/httpserver /
 ---> 97912efd0932
Step 8/9 : EXPOSE 80
 ---> Running in 9910e19c491d
Removing intermediate container 9910e19c491d
 ---> 9b635b3869f7
Step 9/9 : ENTRYPOINT ["/httpserver"]
 ---> Running in d3d53eae5dd3
Removing intermediate container d3d53eae5dd3
 ---> 58203a4ad45c
Successfully built 58203a4ad45c
Successfully tagged httpserver:v1.0
root@ASUS-Desktop-Win11:/home/oliver/hw3# docker images
REPOSITORY   TAG       IMAGE ID       CREATED          SIZE
httpserver   v1.0      58203a4ad45c   40 seconds ago   7.04MB
<none>       <none>    537ba83158d8   41 seconds ago   964MB
golang       1.17      742df529b073   4 weeks ago      942MB

root@ASUS-Desktop-Win11:/home/oliver/hw3# docker login
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: oliverzhang17
Password:
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
root@ASUS-Desktop-Win11:/home/oliver/hw3# docker tag httpserver:v1.0 oliverzhang17/httpserver:0.2

root@ASUS-Desktop-Win11:/home/oliver/hw3# docker images
REPOSITORY                 TAG       IMAGE ID       CREATED         SIZE
httpserver                 v1.0      58203a4ad45c   3 minutes ago   7.04MB
oliverzhang17/httpserver   0.2       58203a4ad45c   3 minutes ago   7.04MB
<none>                     <none>    537ba83158d8   3 minutes ago   964MB
golang                     1.17      742df529b073   4 weeks ago     942MB

root@ASUS-Desktop-Win11:/home/oliver/hw3# docker push oliverzhang17/httpserver:0.2
The push refers to repository [docker.io/oliverzhang17/httpserver]
a8407fac9ecf: Pushed
0.2: digest: sha256:8bdf67780d7affd76ed721313e80d8124b741b2546d52ea2f909db5ce8d09fc5 size: 528

### create container 

root@ASUS-Desktop-Win11:/home/oliver/hw3# docker run -d httpserver:v1.0
ad0f04747b7d522042f1508217703bcc26183b0e529fdedca9d5333c8fec34db
root@ASUS-Desktop-Win11:/home/oliver/hw3# docker ps
CONTAINER ID   IMAGE             COMMAND         CREATED          STATUS          PORTS     NAMES
ad0f04747b7d   httpserver:v1.0   "/httpserver"   17 seconds ago   Up 16 seconds   80/tcp    recursing_germain

### get container ip

root@ASUS-Desktop-Win11:/home/oliver/hw3# PID=$(docker inspect --format "{{ .State.Pid }}" recursing_germain)
root@ASUS-Desktop-Win11:/home/oliver/hw3# echo $PID
2918
root@ASUS-Desktop-Win11:/home/oliver/hw3# nsenter -t $PID -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: sit0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/sit 0.0.0.0 brd 0.0.0.0
10: eth0@if11: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever

### test

root@ASUS-Desktop-Win11:/home/oliver/hw3# curl 172.17.0.2
Success! Client IP:  172.17.0.1
root@ASUS-Desktop-Win11:/home/oliver/hw3# curl 172.17.0.2/healthz
wokring, response code is 200
root@ASUS-Desktop-Win11:/home/oliver/hw3#
