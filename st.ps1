 docker run --rm -it `
      -p '9902:9902' `
      -p '10000:10000' `
      'envoyproxy/envoy:dev-cacd32bc9e3285f6eee7efbc3fcb3e7291975055' `
         -c '/etc/envoy/envoy.yaml' `
         --config-yaml "$(Get-Content -Raw envoy-override.yaml)"
