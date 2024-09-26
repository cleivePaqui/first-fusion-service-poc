# fusion-api-template

https://portal.n-able.dev/catalog/rm/system/fusion-api-template/docs

## Project Setup

Every Fusion Golang project has some basic tooling requirements that require installation and
sometimes version pinning to satisfy their build pipeline.

To set up your developer environment, follow the
[common tools](https://portal.n-able.dev/docs/rm/system/fusion-common-developer-docs/howto/getting-started/)
documentation for more information.

## Skaffold
1. Run this first to download the localstack docker image(avoids skaffold timing out trying to pull it)
    ```
    docker pull localstack/localstack:latest
    ```
2. Provide cerberus credentials to generate k8s Secrets
    ```
    CLIENT_ID=$APP_CERBERUS_CLIENT_ID CLIENT_SECRET=$APP_CERBERUS_CLIENT_SECRET make generate-local-cerberus-secrets
    ```
3. Ensure istioctl is installed via our brewfile in common repo and then run these commands.

   **Important:** make sure you're pointing to your local context!
   ```
   kubectl config use-context docker-desktop
   istioctl install -y
   kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.2/cert-manager.yaml
   kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.20/samples/addons/kiali.yaml
   kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.20/samples/addons/jaeger.yaml
   kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.20/samples/addons/prometheus.yaml
   kubectl apply -f deployments/k8s/overlays/local/resources/issuer.yaml
   ```

4. Add entry for ingest to hosts file: /etc/hosts
   ```
   127.0.0.1 api-template.local
   ```

5. Run skaffold!
    ```
    skaffold dev -p local
    ```

6. Ping api-template home page
   ``go to https://api-template.local/welcome``
