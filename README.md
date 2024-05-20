# Toll Calculator
Microservice-based architecture application containing services:
- OBU
- Data Receiver
- Distance Calculator
- Invoice Aggregator </br>

All are used effectively together to calculate a toll fee displayed from the invoice aggregator based on simulated OBU data of </br>
latitude/longitude converted to distance traveled, then multipled by base price to get the toll fee for specific obu ID, among </br>
other processes as described below...

```
This application was made to demonstrate the effectiveness of building microservices in Golang. While also testing myself
to not only create my custom transport layers and metrics system with Prometheus/Grafana, but also a custom microservice
architecture without microservice frameworks.

As an added layer of complexity I decided to rewrite my entire aggregator microservice using go-kit to build a more
enterprise-like architecture to learn from. This has really cemented my skills in Golang coming from web servers/JSON apis.
```
##### Below is a graphic representing how the services interact with each other:
![image](https://github.com/McFlanky/toll-microservices-calc/assets/153543951/5e233021-2a41-402c-ad82-43e7ae8e7ec9)

### ⚙️ Tech Stack
- Golang
- Apache Kafka
- Gorilla (Golang web toolkit)
- gRPC & Protobuffers
- Prometheus
- Grafana


## Getting Started & How to Make
### Run docker desktop then paste this in your terminal to boot up container
```
docker compose up -d
```

### Make .env file & setup enviroment variables
Check repo http.go/main.go in aggregator directory for further details
```
AGG_HTTP_ENDPOINT=<your chosen port number>
AGG_GRPC_ENDPOINT=<your chosen port number>
AGG_STORE_TYPE=memory
```
You can run the base application from here by running the servers as according to their MakeFile names, in order of: </br>
- Make aggregator
- Make calculator
- Make receiver
- Make obu
  
HTTP custom Gateway layer can be called with:
- make gate 

To send a GET request in http://localhost:6000/obu?={obuID} to get back how long it "took" & "uri"

Now you can send GET requests to:</br>
http://localhost:4000/invoice?obu={obuID_from_make_aggregator_terminal} to get individual information from each invoice</br>
And POST requests to:</br>
http://localhost:4000/aggregate to post your own "invoice" in the form of JSON with a value(int), obuID(int), unix(int)</br>

All information for each Microservice, in each seperate terminal, is shown as such:</br>
### Invoice Aggregator: 🏁 Finish Line</br>
![image](https://github.com/McFlanky/toll-microservices-calc/assets/153543951/991c31ab-96c5-46fd-bf64-1b87eb1cf43b)</br>
### Distance Calculator: ⬆️</br>
![image](https://github.com/McFlanky/toll-microservices-calc/assets/153543951/47020253-a881-439e-ad36-77b52198de9b)</br>
### Data Receiver: ⬆️</br>
![image](https://github.com/McFlanky/toll-microservices-calc/assets/153543951/79dc3231-d686-4c0c-9b39-fb54f7084b13)</br>
### OBU: ⬆️</br>
![image](https://github.com/McFlanky/toll-microservices-calc/assets/153543951/1005af63-bcf7-42af-b99b-482778fe9ba7)</br>
**MakeFile warning affects nothing**

## Further Installation

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

Now at http://localhost:4000/metrics there will be a log of metrics showing if errors occured, if aggregation was called or if calculation was called and how many times and will be graphed in prometheus 

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
