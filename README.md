# Natsy
### Simple [NATS](https://nats.io/) client that allows to publish and make requests

## Usage
Natsy supports both command line flags and configuration file (flags override file configuration).

##### Flags
```bash
            --message string     nats message
            --request            nats request
            --subject string     nats subject
            --timeout duration   nats timeout (request only) (default 1s)
            --url string         nats url
```

###### E.g
 ```bash
 ./natsy --url "demo.nats.io" --subject "foo"  --message "Hello, world." --timeout 500ms
 ```
 ```bash
 demo.nats.io - foo > Hello, world.
 demo.nats.io - foo < published
 ```
  
 ###### _Request_
 `natsy` can be used to make request also if the `--request` flag is provided:
 ```bash
 ./natsy --url "demo.nats.io" --subject "foo"  --message "Hello, world." --request
 ```
 _Example output_:
 ```bash
 demo.nats.io - foo > Hello, world.
 demo.nats.io - foo < demo.nats.io - foo < err: nats: timeout
 ```


##### Configuration file
Natsy looks for a configuration file called `config.yaml` in the same directory it is executed.

###### E.g
_config.yaml_
```yaml
url: demo.nats.io
subject: foo
message: "Hello, world."
timeout: 500ms
request: false
```

###### E.g
 ```bash
 ./natsy
 ```
 ```bash
 demo.nats.io - foo > Hello, world.
 demo.nats.io - foo < published
 ```