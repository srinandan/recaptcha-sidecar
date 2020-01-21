# recaptcha-sidecar

This service is meant to run as a sidecar to the Apigee hybrid API gateway (also known as Message Processor). The service takes a Google Cloud Service Account as a parameter and is used to create assessments using [Enterprise reCAPTCHA](https://cloud.google.com/recaptcha-enterprise/docs/).

## Use Case

Apigee's API Management platform can expose APIs securely (protected via API keys, OAuth etc.). There are times when such methods of protection are not possible (ex: registration API, login API etc.). Such APIs may become tagets for BOTs. Enterprise reCAPTCHA is a BOT detection tool from Google Cloud and can be integrated with Apigee. User experiences include reCAPTCHA [instrumentation](https://cloud.google.com/recaptcha-enterprise/docs/instrument-web-pages). This enables the UX to obtain a reCAPTCHA token. The UX presents Apigee with the reCAPTCHA token. Apigee's API gateway invokes the reCAPTCHA assessment API and determine the risk score for the client. If the risk threshold is within the boundaries set by the enterprise, the API call can be let through.

## Prerequisites

* Apigee hybrid runtime installed on GKE or GKE on-premises (v1.13.x)
* A GCP Project with Enterprise reCAPTCHA enabled

## Prerequisites to build

* kubectl 1.13 or higher
* docker 19.x or higher (if not using skaffold)
* skaffold 1.1.0 or higher (optional)

## Installation

### Installation via kubectl

1. Build the [docker image](./Dockerfile) `docker build -t gcr.io/{project-id}/recaptcha-sidecar`
2. Push to a container registry `docker push gcr.io/{project-id}/recaptcha-sidecar`
3. Modify the kubernetes [manifest](./recaptcha-sidecar.yaml)

```bash

kubectl apply -n {namespace} -f recaptcha-sidecar.yaml
```

### Installation via Skaffold

This application can also be installed via [skaffold](https://skaffold.dev/). Modify the [skaffold.yaml](./skaffold.yaml) to set the appropriate project name.

```bash

skaffold run
```

## Supported Operations

### Environment Variables

The following environment variables are mandatory:

* `GOOGLE_APPLICATION_CREDENTIALS` - Path to service account json
* `PROJECT_NUMBER` - GCP Project number where reCAPTCHA is enabled
* `SITE_KEY` - reCAPTCHA site key

### Store data (Populate Cache)

Path: `/assessment/{token}`
Method: `GET`

```bash

curl 0.0.0.0:8080/token/xxxx 
```

## Access patterns from Apigee hyrid

A typical pattern/example would be to use a [Service Callout policy](https://docs.apigee.com/api-platform/reference/policies/service-callout-policy) to access operations supported by the service. 

This sample [sharedflow](./sharedflowbundle) within any API proxy. The shredflow parses the reCAPTCHA response and sets `recaptchaDecision` and `riskScore` message context variables. That API proxy can make decisions based the message context variables.