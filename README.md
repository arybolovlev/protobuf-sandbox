# Protobuf Sandbox

## Server

### Run

```console
$ go build -o bin/misha
$ bin/misha -server
2023/03/06 10:09:15 Running Server
2023/03/06 10:09:15 server listening at [::]:8090
```

## Client

### Run

```console
$ bin/misha -client -file _examples/job.hcl
2023/03/06 10:09:36 Running Client
2023/03/06 10:09:36 New job sucessfully created: id:"white" name:"4e3fd8a0-6fab-4658-a136-0f60021f8d3d"
2023/03/06 10:09:36 New job sucessfully created: id:"blue" name:"9d9dd5bc-01e6-492e-9e17-db640adfd860"
2023/03/06 10:09:36 New job sucessfully created: id:"red" name:"b8d09a87-d1e4-47cc-8c09-7e79a519843b"
```
