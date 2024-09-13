package com.hcmut.gateway.configuration;

import com.hcmut.gateway.account.CurrentUserDetails;
import com.hcmut.gateway.util.AccountUtils;
import org.springframework.http.HttpMethod;
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

    public RestClient.RequestBodySpec getBodySpec(HttpMethod method, String uri) {
        var bodySpec = restClient.method(method).uri(uri);
        CurrentUserDetails currentUserDetails = AccountUtils.getCurrentUserDetails();
        if (currentUserDetails.getId() != null) {
            bodySpec = bodySpec.header("user_id", String.valueOf(currentUserDetails.getId()));
        }
        return bodySpec;
    }

    public UriComponentsBuilder getUriBuilder() {
        return UriComponentsBuilder.newInstance().scheme("http").host(host).port(port);
    }
}