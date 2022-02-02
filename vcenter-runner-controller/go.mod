module github.com/whywaita/shoes-agent-vcenter/vcenter-runner-controller

go 1.17

replace github.com/whywaita/shoes-agent-vcenter/pkg/vmware => ../pkg/vmware

require (
	github.com/google/uuid v1.2.0
	github.com/vmware/govmomi v0.27.2
	github.com/whywaita/shoes-agent-vcenter/pkg/vmware v0.0.0-00010101000000-000000000000
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
)
