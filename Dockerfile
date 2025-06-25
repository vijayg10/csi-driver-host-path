# Copyright 2021 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Stage 1: Build the Go binary
FROM golang:1.24-alpine AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /hostpathplugin ./cmd/hostpathplugin

# Stage 2: Create the runtime image
FROM alpine
LABEL maintainers="Kubernetes Authors"
LABEL description="HostPath Driver"
RUN apk add --no-cache util-linux coreutils socat tar
COPY --from=builder /hostpathplugin /hostpathplugin
ENTRYPOINT ["/hostpathplugin"]
