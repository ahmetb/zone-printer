# Cloud Run - Zone Printer demo application

This web application prints the Google Cloud datacenter it’s running on with
country and city name.

It's used to demonstrate global load balancing capabilities of Google Cloud
HTTPS Load Balancer, as it routes the request to the compute region closest
to the visitor.

Example:

> ```text
> Welcome from Google Cloud datacenters at:
> The Dalles, Oregon, USA
>
> You are now connected to "us-west1"
> ```

## Deploy to Cloud Run

Build the container image yourself, or use pre-built image:

    gcloud run deploy --platform=managed --allow-unauthenticated \
        --image gcr.io/ahmetb-public/zoneprinter

If you want to automate deploying this to all available
Cloud Run regions, check out [this Terraform guide](https://github.com/ahmetb/cloud-run-multi-region-terraform/).
