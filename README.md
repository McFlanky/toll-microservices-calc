# toll-calculator

```
docker compose up -d
```

### Install protobuf compiler (protoc compiler)
For linux users or WSL2
```
sudo apt install -y protobuf-compiler
```

For Mac users you can use Brew for this
```
brew install protobuf
```

### Install GRPC and Protobuffer plugins for Golang
1. Protobuffers
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```
2. GRPC
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
3. NOTE that you need to set /go/bin directory in your path or wherever your Go directory lives
```
PATH="${PATH}:${HOME}/go/bin"
```

4. Install the package dependencies 
4.1 protobuffer package
```
go get google.golang.org/protobuf/runtime/protoimpl
```
4.2 grpc package
```
go get google.golang.org/grpc
```

### Installing Prometheus
Install Prometheus in a Docker Container
```
docker run -p 9090:9090 -v ./.config/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus
```

Installing Prometheus golang client
```
go get github.com/prometheus/client_golang/prometheus
```

Installing Prometheus natively in your system
1. Clone the repository
```
git clone https://github.com/prometheus/prometheus.git
cd prometheus
```

2. Install it
```
make build
```

3. Run the Prometheus Deamon
```
./prometheus --config.file=<your_config_file>yml
```

4. In this project's case, that would be (running from inside the project directory)
```
./prometheus/prometheus --config.file=.config/prometheus.yml
```

### Installing Grafana
Install Grafana locally in Debian/Ubuntu for WSL2 users
Install the prerequisite packages
```
sudo apt-get install -y apt-transport-https software-properties-common wget
```
Import the GPG key:
```
sudo mkdir -p /etc/apt/keyrings/
wget -q -O - https://apt.grafana.com/gpg.key | gpg --dearmor | sudo tee /etc/apt/keyrings/grafana.gpg > /dev/null
```
To add a repository for stable releases, run the following command:
```
echo "deb [signed-by=/etc/apt/keyrings/grafana.gpg] https://apt.grafana.com stable main" | sudo tee -a /etc/apt/sources.list.d/grafana.list
```
Run the following command to update the list of available packages:
```
# Updates the list of available packages
sudo apt-get update
```
Then:
```
# Installs the latest OSS release:
sudo apt-get install grafana
```

Start:
```
sudo systemctl start grafana-server
```