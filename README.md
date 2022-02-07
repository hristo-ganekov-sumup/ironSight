# ironSight

Its a tool for detecting drift on state to AWS infra basis.

Its hardcoded for use with autosg pattern.

Its planned to be used together with [stateGetter](https://github.com/TFArmada/stateGetter)


## Usage
```shell
TFE_TOKEN=<SUPER_SECRET_TFE_TOKEN> stateGetter -organization yourorg -workspace autosg-workspace -filename live.tfstate

./ironSight_linux_amd64 --help
Usage of ./ironSight_linux_amd64:
  -region string
    	AWS Region (default "eu-west-1")
  -state string
    	State file (default "live.tfstate")

```