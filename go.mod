module github.com/DmitryZ-outten/trivybeat

go 1.16

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.2.0+incompatible
	github.com/Microsoft/go-winio => github.com/bi-zone/go-winio v0.4.15
	github.com/Shopify/sarama => github.com/elastic/sarama v1.19.1-0.20210120173147-5c8cb347d877
	github.com/cucumber/godog => github.com/cucumber/godog v0.8.1
	github.com/docker/docker => github.com/docker/engine v0.0.0-20191113042239-ea84732a7725
	github.com/docker/go-plugins-helpers => github.com/elastic/go-plugins-helpers v0.0.0-20200207104224-bdf17607b79f
	github.com/dop251/goja => github.com/andrewkroh/goja v0.0.0-20190128172624-dd2ac4456e20
	github.com/dop251/goja_nodejs => github.com/dop251/goja_nodejs v0.0.0-20171011081505-adff31b136e6
	github.com/fsnotify/fsevents => github.com/elastic/fsevents v0.0.0-20181029231046-e1d381a4d270
	github.com/fsnotify/fsnotify => github.com/adriansr/fsnotify v0.0.0-20180417234312-c9bbe1f46f1d
	github.com/google/gopacket => github.com/adriansr/gopacket v1.1.18-0.20200327165309-dd62abfa8a41
	github.com/insomniacslk/dhcp => github.com/elastic/dhcp v0.0.0-20200227161230-57ec251c7eb3 // indirect
	github.com/kardianos/service => github.com/blakerouse/service v1.1.1-0.20200924160513-057808572ffa
	github.com/tonistiigi/fifo => github.com/containerd/fifo v0.0.0-20190816180239-bda0ff6ed73c
	golang.org/x/tools => golang.org/x/tools v0.0.0-20200602230032-c00d67ef29d0 // release 1.14
)

require (
	github.com/Shopify/sarama v1.27.0 // indirect
	github.com/akavel/rsrc v0.10.2 // indirect
	github.com/aquasecurity/fanal v0.0.0-20210501093021-8aaac3e8dea7
	github.com/aquasecurity/trivy v0.17.2
	github.com/aquasecurity/trivy-db v0.0.0-20210429114658-ae22941a55d0
	github.com/dgrijalva/jwt-go v3.2.1-0.20190620180102-5e25c22bd5d6+incompatible // indirect
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/docker/engine-api v0.4.0
	github.com/dop251/goja v0.0.0-20211022113120-dc8c55024d06 // indirect
	github.com/dop251/goja_nodejs v0.0.0-20211022123610-8dd9abb0616d // indirect
	github.com/elastic/beats/v7 v7.15.1
	github.com/elastic/elastic-agent-client/v7 v7.0.0-20210922110810-e6f1f402a9ed // indirect
	github.com/elastic/go-sysinfo v1.7.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/gin-gonic/gin v1.7.0 // indirect
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/gofrs/uuid v4.1.0+incompatible // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/nomad/api v0.0.0-20201203164818-6318a8ac7bf8 // indirect
	github.com/jcchavezs/porto v0.4.0 // indirect
	github.com/josephspurrier/goversioninfo v1.3.0 // indirect
	github.com/magefile/mage v1.11.0
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mitchellh/gox v1.0.1
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/pierrre/gotestcover v0.0.0-20160517101806-924dca7d15f0
	github.com/poy/eachers v0.0.0-20181020210610-23942921fe77 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/tsg/go-daemon v0.0.0-20200207173439-e704b93fd89b
	go.elastic.co/apm v1.14.0 // indirect
	go.elastic.co/ecszap v1.0.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/net v0.0.0-20211101193420-4a448f8816b3 // indirect
	golang.org/x/sys v0.0.0-20211102061401-a2f17f7b995c // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.7
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/genproto v0.0.0-20211101144312-62acf1d99145 // indirect
	google.golang.org/grpc v1.41.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	honnef.co/go/tools v0.1.3 // indirect
	howett.net/plist v0.0.0-20201203080718-1454fab16a06 // indirect
)
