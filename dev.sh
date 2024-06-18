#!/bin/bash

air -c .air.toml && \
templ generate -watch && \
tailwind -i ./cmd/web/assets/css/input.css -o ./cmd/web/assets/css/output.css --watch
