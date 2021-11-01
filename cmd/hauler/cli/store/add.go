package store

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/rancherfederal/hauler/pkg/apis/hauler.cattle.io/v1alpha1"
	"github.com/rancherfederal/hauler/pkg/content/chart"
	"github.com/rancherfederal/hauler/pkg/content/file"
	"github.com/rancherfederal/hauler/pkg/content/image"
	"github.com/rancherfederal/hauler/pkg/log"
	"github.com/rancherfederal/hauler/pkg/store"
)

type AddFileOpts struct {
	Name string
}

func (o *AddFileOpts) AddFlags(cmd *cobra.Command) {
	f := cmd.Flags()
	f.StringVarP(&o.Name, "name", "n", "", "(Optional) Name to assign to file in store")
}

func AddFileCmd(ctx context.Context, o *AddFileOpts, s *store.Store, ref string) error {
	l := log.FromContext(ctx)
	l.Debugf("running cli command `hauler store add`")

	s.Open()
	defer s.Close()

	cfg := v1alpha1.File{
		Ref:  ref,
		Name: o.Name,
	}

	f := file.NewFile(cfg)
	return s.Add(ctx, f)
}

type AddImageOpts struct {
	Name string
}

func (o *AddImageOpts) AddFlags(cmd *cobra.Command) {
	f := cmd.Flags()
	_ = f
}

func AddImageCmd(ctx context.Context, o *AddImageOpts, s *store.Store, ref string) error {
	l := log.FromContext(ctx)
	l.Debugf("running cli command `hauler store add image`")

	s.Open()
	defer s.Close()

	cfg := v1alpha1.Image{
		Ref: ref,
	}

	i := image.NewImage(cfg)
	return s.Add(ctx, i)
}

type AddChartOpts struct {
	Name    string
	Version string
	RepoURL string

	// TODO: Support helm auth
	Username              string
	Password              string
	PassCredentialsAll    bool
	CertFile              string
	KeyFile               string
	CaFile                string
	InsecureSkipTLSverify bool
	RepositoryConfig      string
	RepositoryCache       string
}

func (o *AddChartOpts) AddFlags(cmd *cobra.Command) {
	f := cmd.Flags()

	f.StringVarP(&o.RepoURL, "repo", "r", "", "Chart repository URL")
	f.StringVar(&o.Version, "version", "", "(Optional) Version of the chart to download, defaults to latest if not specified")
}

func AddChartCmd(ctx context.Context, o *AddChartOpts, s *store.Store, name string) error {
	l := log.FromContext(ctx)
	l.Debugf("running cli command `hauler store add chart`")

	s.Open()
	defer s.Close()

	cfg := v1alpha1.Chart{
		Name:    name,
		RepoURL: o.RepoURL,
		Version: o.Version,
	}

	c := chart.NewChart(cfg)
	return s.Add(ctx, c)
}