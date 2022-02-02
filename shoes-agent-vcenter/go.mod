module github.com/whywaita/shoes-agent-vcenter/shoes-agent-vcenter

go 1.17

replace (
	github.com/whywaita/shoes-agent-vcenter/pkg/vmware => ../pkg/vmware
	github.com/whywaita/shoes-agent/agent => ../../shoes-agent/agent
	github.com/whywaita/shoes-agent/proto.go => ../../shoes-agent/proto.go
	github.com/whywaita/shoes-agent/shoes-agent => ../../shoes-agent/shoes-agent
)

require (
	github.com/google/uuid v1.2.0
	github.com/hashicorp/go-plugin v1.4.3
	github.com/vmware/govmomi v0.27.2
	github.com/whywaita/myshoes v1.10.4
	github.com/whywaita/shoes-agent-vcenter/pkg/vmware v0.0.0-00010101000000-000000000000
	github.com/whywaita/shoes-agent/agent v0.0.0-00010101000000-000000000000
	github.com/whywaita/shoes-agent/proto.go v0.0.0-00010101000000-000000000000
	github.com/whywaita/shoes-agent/shoes-agent v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.43.0
)

require (
	github.com/bradleyfalzon/ghinstallation/v2 v2.0.3 // indirect
	github.com/fatih/color v1.7.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/golang-jwt/jwt/v4 v4.0.0 // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/google/go-github/v35 v35.2.0 // indirect
	github.com/google/go-github/v39 v39.0.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/hashicorp/go-hclog v0.14.1 // indirect
	github.com/hashicorp/go-version v1.3.0 // indirect
	github.com/hashicorp/yamux v0.0.0-20180604194846-3520598351bb // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/mitchellh/go-testing-interface v0.0.0-20171004221916-a61a99592b77 // indirect
	github.com/oklog/run v1.0.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	github.com/tklauser/numcpus v0.3.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/sys v0.0.0-20210816074244-15123e1e1f71 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20200825200019-8632dd797987 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
)
