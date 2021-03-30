package beater

import (
	"fmt"
	"context"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"

	"github.com/DmitryZ-outten/trivybeat/config"

	DockerClient "github.com/docker/engine-api/client"
	DockerTypes "github.com/docker/engine-api/types"

	image2 "github.com/aquasecurity/fanal/artifact/image"
	"github.com/aquasecurity/fanal/cache"
	"github.com/aquasecurity/fanal/image"
	"github.com/aquasecurity/trivy/pkg/log"
	"github.com/aquasecurity/trivy/pkg/report"
	"github.com/aquasecurity/trivy/pkg/rpc/client"
	"github.com/aquasecurity/trivy/pkg/scanner"
	"github.com/aquasecurity/trivy/pkg/types"
	"golang.org/x/xerrors"
)

// trivybeat configuration.
type trivybeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of trivybeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &trivybeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts trivybeat.
func (bt *trivybeat) Run(b *beat.Beat) error {
	logp.Info("trivybeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	// Create a Docker client
	ctx := context.Background()
	cli, err := DockerClient.NewClient("unix:///var/run/docker.sock", "v1.41", nil, nil)
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		containers, err := cli.ContainerList(ctx, DockerTypes.ContainerListOptions{})
		if err != nil {
			panic(err)
		}

		for _, container := range containers {
			logp.Info(container.Image)
			results := TrivyScan( string(container.Image), bt.config.Server )
			for _, vulnerability := range results[0].Vulnerabilities {
				logp.Info("%+v\n", vulnerability.VulnerabilityID)
				event := beat.Event{
					Timestamp: time.Now(),
					Fields: common.MapStr{
						"type":    b.Info.Name,
						"container.image.name": string(container.Image),
						"vulnerability.id": vulnerability.VulnerabilityID,
						"vulnerability.severity": vulnerability.Vulnerability.Severity,
						"vulnerability.description": vulnerability.Vulnerability.Description,
						"vulnerability.reference": vulnerability.Vulnerability.References,
						"vulnerability.pkgname": vulnerability.PkgName,
					},
				}
				bt.client.Publish(event)
			}
		}
	}
}

// Scan with Trivy
func TrivyScan(imageFlag string, url string) report.Results {

	if err := log.InitLogger(true, false); err != nil {
		log.Logger.Fatalf("error happened: %v", xerrors.Errorf("failed to initialize a logger: %w", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
	defer cancel()

	localCache, err := cache.NewFSCache("")
	if err != nil {
		log.Logger.Fatalf("could not initialize f: %v", err)
	}

	sc, cleanUp, err := initializeDockerScanner(ctx, imageFlag, localCache, client.CustomHeaders{}, client.RemoteURL(url), time.Second*5000)
	if err != nil {
		log.Logger.Fatalf("could not initialize scanner: %v", err)
	}

	defer cleanUp()

	results, err := sc.ScanArtifact(ctx, types.ScanOptions{
		VulnType:            []string{"os", "library"},
		ScanRemovedPackages: true,
		ListAllPackages:     true,
	})
	if err != nil {
		log.Logger.Fatalf("could not scan image: %v", err)
	}

	if len(results) > 0 {
		log.Logger.Infof("%d vulnerability/ies found", len(results[0].Vulnerabilities))
	} else {
		log.Logger.Infof("no vulnerabilities found for image %s", imageFlag)
	}

	return results
}

// Initialize Docker Scanner
func initializeDockerScanner(ctx context.Context, imageName string, artifactCache cache.ArtifactCache, customHeaders client.CustomHeaders, url client.RemoteURL, timeout time.Duration) (scanner.Scanner, func(), error) {
	scannerScanner := client.NewProtobufClient(url)
	clientScanner := client.NewScanner(customHeaders, scannerScanner)
	dockerOption, err := types.GetDockerOption(timeout)
	if err != nil {
		return scanner.Scanner{}, nil, err
	}
	imageImage, cleanup, err := image.NewDockerImage(ctx, imageName, dockerOption)
	if err != nil {
		return scanner.Scanner{}, nil, err
	}
	artifact := image2.NewArtifact(imageImage, artifactCache)
	scanner2 := scanner.NewScanner(clientScanner, artifact)
	return scanner2, func() {
		cleanup()
	}, nil
}

// Stop stops trivybeat.
func (bt *trivybeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
