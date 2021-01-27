# Zone Printer demo application

This web application prints the Google Cloud datacenter it’s running on with
information about where the datacenter is located (city, country and flag).

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

```sh
gcloud run deploy --platform=managed --allow-unauthenticated \
        --image gcr.io/ahmetb-public/zoneprinter
```

Follow the [Serving traffic from multiple
regions](https://cloud.google.com/run/docs/multiple-regions) tutorial from
Cloud Run documentation to see how you can deploy this application to all
regions and set up a load balancer on top.

If you want to automate further, check out [this Terraform
guide](https://github.com/ahmetb/cloud-run-multi-region-terraform/) for
achieving the same task.

-----

This is not an official Google project.
