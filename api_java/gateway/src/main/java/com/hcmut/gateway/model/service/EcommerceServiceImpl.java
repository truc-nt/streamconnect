package com.hcmut.gateway.model.service;

import com.hcmut.gateway.configuration.ExternalServiceClient;
import com.hcmut.gateway.model.service.external_request_model.CreateShopForNewUserRequest;
import com.hcmut.shared_lib.common_util.ExternalRequestUtils;
import com.hcmut.shared_lib.model.entity.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.HttpMethod;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Service;

@Service
public class EcommerceServiceImpl implements EcommerceService {
    private final ExternalServiceClient ecommerceServiceClient;

    @Autowired
    public EcommerceServiceImpl(@Qualifier("ecommerceServiceClient") ExternalServiceClient ecommerceServiceClient) {
        this.ecommerceServiceClient = ecommerceServiceClient;
    }

    @Override
    public void createShopForNewUser(User user) {
        String uriString = ecommerceServiceClient.getUriBuilder().path("/api/shops/forNewUser").toUriString();
        CreateShopForNewUserRequest request = new CreateShopForNewUserRequest();
        request.setUserId(user.getId());
        ecommerceServiceClient.getBodySpec(HttpMethod.POST, uriString).header("user_id", String.valueOf(user.getId()))
                .contentType(MediaType.APPLICATION_JSON).body(request).retrieve()
                .onStatus(status -> status.is4xxClientError() || status.is5xxServerError(),
                        ExternalRequestUtils::handleErrorRequest
                ).toBodilessEntity();
    }
}
