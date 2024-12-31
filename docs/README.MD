> install
```shell
sudo chmod 0777 devinggo
cp devinggo.service /usr/lib/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable devinggo
sudo systemctl start devinggo
```