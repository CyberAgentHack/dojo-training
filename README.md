# CA Tech Dojo 事前準備
golang, Docker, Docker Composeをインストールして下記のコマンドを実行できるか確認してください

```
$ go run main.go
Hello World !
```

## 手順
### 1. golangをインストールする<br>
詳細なインストール方法についてはここでは省略します。<br>
versionは1.17以降としてください。<br>
MacやLinuxで開発を行う場合はanyenvやgoenvなどのバージョン管理ツールでインストールするのがおすすめです。
### 2. Docker, Docker Composeをインストールする<br>
Docker DesktopのPersonal Planをインストールしてください<br>
https://www.docker.com/products/personal <br><br>

※ 開発にWindowsを用いる場合、環境によってはDockerコンテナの起動が困難な可能性があります。自分では解決できなそうと感じた場合はSlack等でご相談ください。
### 3. このリポジトリをcloneする<br>
```
$ git clone https://github.com/CyberAgentHack/dojo-training.git
```
### 4. コンテナを起動する<br>
```
$ cd dojo-training
$ docker-compose up -d
Creating network "dojo-training_default" with the default driver
Creating volume "dojo-training_db-data" with local driver
Creating dojo-training_mysql_1 ... done
```
成功するとMySQL(データベース)のコンテナが立ち上がります。<br>
```
$ docker-compose ps
        Name                       Command               State                 Ports              
--------------------------------------------------------------------------------------------------
dojo-training_mysql_1   docker-entrypoint.sh mysql ...   Up      0.0.0.0:3306->3306/tcp, 33060/tcp
```
### 5. main.goを実行する<br>
go run コマンドでmain.goを実行し、`Hello World !` と表示されることを確認してください。<br>
表示されない、もしくはエラーが表示される場合はNGとします。
```
$ go run main.go
Hello World !
```
### 6. 実行できたことを確認できたらコンテナを停止させる
```
$ docker-compose down -v
Stopping dojo-training_mysql_1 ... done
Removing dojo-training_mysql_1 ... done
Removing network dojo-training_default
Removing volume dojo-training_db-data
```

### 7. golangci-lintをインストールする
```shell
$ go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1
```

### 8. lintエラーが出ないことを確認
```shell
$ golangci-lint run
```
