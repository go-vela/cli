module github.com/go-vela/cli

go 1.13

replace github.com/go-vela/pkg-executor => ../pkg-executor

require (
	github.com/Masterminds/semver/v3 v3.1.1
	github.com/buildkite/yaml v0.0.0-20181016232759-0caa5f0796e3
	github.com/cli/browser v1.0.0
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dustin/go-humanize v1.0.0
	github.com/go-vela/compiler v0.6.1-0.20210111222408-a8b5012a5943
	github.com/go-vela/mock v0.6.1-0.20210111203055-2629f560e9b7
	github.com/go-vela/pkg-executor v0.0.0-00010101000000-000000000000
	github.com/go-vela/pkg-runtime v0.6.1-0.20210111230700-42483a9ea3c1
	github.com/go-vela/sdk-go v0.6.1-0.20210111215046-d4ccd4904b94
	github.com/go-vela/types v0.6.1-0.20210111181528-d3bb371e9ec6
	github.com/gosuri/uitable v0.0.4
	github.com/lunixbochs/vtclean v1.0.0 // indirect
	github.com/manifoldco/promptui v0.8.0
	github.com/mattn/go-runewidth v0.0.6 // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/afero v1.5.1
	github.com/urfave/cli/v2 v2.3.0
	gopkg.in/yaml.v2 v2.4.0
)
