# README

Skeleton OAS server implementation for TAPI Latest, note this was built on 20th Nov 2018 .

## Generate

Read this [page](https://github.com/MEF-GIT/MEF-LSO-Presto-SDK/tree/master/published/swagger) to understand how to generate a single OAS specification from a set of YANG files.

```bash
java -cp ./yang2swagger-tapi-cli-1.4-cli.jar com.amartus.y2s.Generator -yang-dir TAPI/YANG -output tapi-2.2.0-RC2.yml
```

Followed by 

```bash
swagger generate model -f tapi-2.2.0-RC2.yml
```

## Install

```bash
go install github.com/damianoneill/tapi/...
```