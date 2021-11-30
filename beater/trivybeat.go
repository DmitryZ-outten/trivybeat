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
	"github.com/aquasecurity/fanal/image"
	"github.com/aquasecurity/trivy/pkg/log"
	"github.com/aquasecurity/trivy/pkg/report"
	"github.com/aquasecurity/trivy/pkg/rpc/client"
	"github.com/aquasecurity/trivy/pkg/scanner"
	"github.com/aquasecurity/trivy/pkg/types"
	"github.com/aquasecurity/trivy/pkg/cache"
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

// Define a map for CVE severity
	severity_id := make(map[string]int)
	severity_id["CRITICAL"] = 1
	severity_id["HIGH"] = 2
	severity_id["MEDIUM"] = 3
	severity_id["LOW"] = 4

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		containers := GetContainers( )
		results := TrivyScan( containers, bt.config.Server )

		for _, container := range results {
			fmt.Printf("\n==================\n")
			fmt.Printf("%+v\n", container[0].Target)
			for _, vulnerability := range container[0].Vulnerabilities {
					event := beat.Event{
						Timestamp: time.Now(),
						Fields: common.MapStr{
							"type":    b.Info.Name,
							"container": common.MapStr{ 
								"image": common.MapStr{ 
									"name": string(container[0].Target),
								},
							},
							"vulnerability": common.MapStr{
								"id": vulnerability.VulnerabilityID,
								"severity": vulnerability.Vulnerability.Severity,
								"severity_id": severity_id[vulnerability.Vulnerability.Severity],
								"description": vulnerability.Vulnerability.Description,
								"reference": vulnerability.Vulnerability.References,
								"pkgname": vulnerability.PkgName,
							},
						},
					}
					fmt.Printf("event created\n")
					bt.client.Publish(event)
			}

		}
		fmt.Printf("\n+++++++++++++++++++++\n")
		
	}
}

// Get containers
func GetContainers() []DockerTypes.Container {

	// Create a Docker client
	ctx := context.Background()
	cli, err := DockerClient.NewClient("unix:///var/run/docker.sock", "v1.41", nil, nil)
	if err != nil {
		panic(err)
	}

	// Get running containers
	containers, err := cli.ContainerList(ctx, DockerTypes.ContainerListOptions{})
	if err != nil {
		logp.Info("could not get running containers: %v", err)
	}

	return containers
}

// Scan with Trivy
func TrivyScan(containers []DockerTypes.Container, url string) []report.Results {

	if err := log.InitLogger(true, false); err != nil {
		log.Logger.Fatalf("error happened: %v", xerrors.Errorf("failed to initialize a logger: %w", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
	defer cancel()

	var vuln []report.Results
	for _, container := range containers {
		fmt.Printf("%+v\n", container)
    	sc, cleanUp, err := initializeDockerScanner(ctx, container.Image, client.CustomHeaders{}, client.RemoteURL(url), time.Second*5000)
		if err != nil {
			log.Logger.Fatalf("could not initialize scanner: %v", err)
		}

		defer cleanUp()

		results, err := sc.ScanArtifact(ctx, types.ScanOptions{
			VulnType:            []string{"os"},
			ScanRemovedPackages: false,
		})
	    if err != nil {
	    	fmt.Printf("%+v\n", err)
	    	//log.Logger.Fatalf("error in image scan: %v", err)
	    } else {
			vuln = append(vuln, results)
		}
	}

	return vuln
}

// Initialize Docker Scanner
func initializeDockerScanner(ctx context.Context, imageName string, customHeaders client.CustomHeaders, url client.RemoteURL, timeout time.Duration) (scanner.Scanner, func(), error) {
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
	artifactCache := cache.NewRemoteCache(cache.RemoteURL(url), nil)
	artifact := image2.NewArtifact(imageImage, artifactCache, nil)
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
