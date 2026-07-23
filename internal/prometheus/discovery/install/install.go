// Package install registers the Prometheus service discovery types Alloy keeps
// in the default engine build. This intentionally omits moby-based discovery
// so the shipped binary does not pull Docker client dependencies through the
// blanket Prometheus installer package.
package install

import (
	_ "github.com/prometheus/prometheus/discovery/aws"
	_ "github.com/prometheus/prometheus/discovery/azure"
	_ "github.com/prometheus/prometheus/discovery/consul"
	_ "github.com/prometheus/prometheus/discovery/digitalocean"
	_ "github.com/prometheus/prometheus/discovery/dns"
	_ "github.com/prometheus/prometheus/discovery/eureka"
	_ "github.com/prometheus/prometheus/discovery/file"
	_ "github.com/prometheus/prometheus/discovery/gce"
	_ "github.com/prometheus/prometheus/discovery/hetzner"
	_ "github.com/prometheus/prometheus/discovery/http"
	_ "github.com/prometheus/prometheus/discovery/ionos"
	_ "github.com/prometheus/prometheus/discovery/kubernetes"
	_ "github.com/prometheus/prometheus/discovery/linode"
	_ "github.com/prometheus/prometheus/discovery/marathon"
	_ "github.com/prometheus/prometheus/discovery/nomad"
	_ "github.com/prometheus/prometheus/discovery/openstack"
	_ "github.com/prometheus/prometheus/discovery/ovhcloud"
	_ "github.com/prometheus/prometheus/discovery/puppetdb"
	_ "github.com/prometheus/prometheus/discovery/scaleway"
	_ "github.com/prometheus/prometheus/discovery/stackit"
	_ "github.com/prometheus/prometheus/discovery/triton"
	_ "github.com/prometheus/prometheus/discovery/uyuni"
	_ "github.com/prometheus/prometheus/discovery/vultr"
	_ "github.com/prometheus/prometheus/discovery/xds"
	_ "github.com/prometheus/prometheus/discovery/zookeeper"
)
