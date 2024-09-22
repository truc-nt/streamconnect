package com.hcmut.gateway.model.service;

import com.hcmut.gateway.model.DTO.RegisterLivestreamProductFollowerRequest;
import com.hcmut.shared_lib.model.entity.User;

public interface EcommerceService {
    void createShopForNewUser(User user);

    void registerLivestreamProductFollower(RegisterLivestreamProductFollowerRequest request);

    void notifyLivestreamProductFollower(Long productId);
}
