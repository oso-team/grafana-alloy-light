package flowcmd

import (
	"github.com/grafana/alloy"
	"github.com/grafana/alloy/internal/alloycli"
	"github.com/grafana/alloy/internal/build"
	_ "github.com/grafana/alloy/internal/prometheus/discovery/install"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"

	// Register Prometheus SD components
	_ "github.com/grafana/alloy/internal/loki/promtail/discovery/consulagent"

	// Register integrations
	_ "github.com/grafana/alloy/internal/static/integrations/install"

	// Embed application manifest for Windows builds
	_ "github.com/grafana/alloy/internal/winmanifest"
)

func init() {
	// If the build version wasn't set by the build process, we'll set it based
	// on the version in .release-please-manifest.json.
	if build.Version == "" || build.Version == "v0.0.0" {
		build.Version = alloy.FallbackVersion()
	}

	prometheus.MustRegister(build.NewCollector("alloy"))
}

// RootCommand exposes the root Cobra command constructed by the internal alloy CLI.
func RootCommand() *cobra.Command {
	return alloycli.Command()
}

func RunCommand() *cobra.Command {
	return alloycli.RunCommand()
}
