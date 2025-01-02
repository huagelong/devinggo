> install
```shell
sudo chmod 0777 devinggo
cp devinggo.service /usr/lib/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable devinggo
sudo systemctl start devinggo
```

> dev
```shell
gf gen service
gf gen service -s=modules/system/logic -d=modules/system/service
gf gen service -s=modules/api/logic -d=modules/api/service
```
