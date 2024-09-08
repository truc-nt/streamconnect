package com.hcmut.gateway.configuration;

import org.springframework.web.client.RestClient;
import org.springframework.web.util.UriComponentsBuilder;

public class ExternalServiceClient {

    private final RestClient restClient;
    private final String host;
    private final int port;

    public ExternalServiceClient(String host, int port) {
        this.restClient = RestClient.create();
        this.host = host;
        this.port = port;
    }

    public RestClient getRestClient() {
        return restClient;
    }

    public UriComponentsBuilder getUriBuilder() {
        return UriComponentsBuilder.newInstance().scheme("http").host(host).port(port);
    }
}