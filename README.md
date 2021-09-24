# vcloud-inet

## Build

```
make go-build
```

## Run

```
make go-run
```

## How to use with VM template

Build executable binary
```
git clone git@gitlab.e-magnum.kz:devops/scripts/golang/vcloud-inet.git
make go-build
mv vcloud-inet /usr/bin/
```

Create init script (execute on boot)
```
mv vcloud-inet.sh /etc/init.d/
chown root:root /etc/init.d/vcloud-inet.sh
chmod +x /etc/init.d/vcloud-inet.sh
update-rc.d vcloud-inet.sh defaults
```
