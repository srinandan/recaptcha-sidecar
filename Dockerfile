# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#docker build -t gcr.io/srinandans-apigee/recaptcha-sidecar .
#docker run -d -p 8080:8080 --name recaptcha-sidecar -v ~/sa.json:/sa.json -e GOOGLE_APPLICATION_CREDENTIALS="./sa.json" -e DEBUG="true" gcr.io/nandanks-151422/recaptcha-sidecar
FROM golang:latest as builder
RUN useradd -U app
ADD . /go/src/recaptcha-sidecar 
WORKDIR /go/src/recaptcha-sidecar
COPY . /go/src/recaptcha-sidecar
ENV GO111MODULE=on
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' -o /go/bin/recaptcha-sidecar

#without these certificates, we cannot verify the JWT token
FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
WORKDIR /
COPY --from=builder /go/bin/recaptcha-sidecar .
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
USER app
CMD ["./recaptcha-sidecar"]