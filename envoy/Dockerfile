FROM envoyproxy/envoy:v1.14.3
COPY ./envoy.yaml /etc/envoy/envoy.yaml
EXPOSE 5002
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml