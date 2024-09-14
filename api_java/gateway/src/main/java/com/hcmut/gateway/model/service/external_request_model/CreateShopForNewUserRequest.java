package com.hcmut.gateway.model.service.external_request_model;

import com.fasterxml.jackson.annotation.JsonProperty;

public class CreateShopForNewUserRequest {
    @JsonProperty("user_id")
    private Long userId;

    @JsonProperty("shop_name")
    private String shopName;

    public CreateShopForNewUserRequest() {}

    public CreateShopForNewUserRequest(Long userId, String shopName) {
        this.userId = userId;
        this.shopName = shopName;
    }

    public Long getUserId() {
        return userId;
    }

    public void setUserId(Long userId) {
        this.userId = userId;
    }

    public String getShopName() {
        return shopName;
    }

    public void setShopName(String shopName) {
        this.shopName = shopName;
    }
}
