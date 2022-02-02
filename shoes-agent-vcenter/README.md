# shoes-agent-vcenter

shoes-provider for `shoes-agent-vcenter`.

`sheoes-agent-vcenter` will create an instance using [Instant Clone](https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-853B1E2B-76CE-4240-A654-3806912820EB.html).

## Setup

### Auth (required)

- `SHOES_VCENTER_URL`
- `SHOES_VCENTER_USERNAME`
- `SHOES_VCENTER_PASSWORD`

### Configure (required)

- `SHOES_VCENTER_DATACENTER_NAME`
- `SHOES_VCENTER_DATASTORE_NAME`
- `SHOES_VCENTER_FOLDER_NAME`
- `SHOES_VCENTER_SOURCE_VM_NAME`
- `SHOES_VCENTER_SOURCE_SNAPSHOT_NAME`

### Optional

- `SHOES_VCENTER_NUMBER_OF_INSTANCE_PER_HOST`
  - Set number of an instance on same host
- `SHOES_VCENTER_IGNORE_VM_NAMES`
  - Set ignore virtual machine for `SHOES_VCENTER_NUMBER_OF_INSTANCE_PER_HOST`