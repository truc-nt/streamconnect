package com.hcmut.gateway.controller;

import com.hcmut.gateway.model.DTO.RegisterLivestreamProductFollowerRequest;
import com.hcmut.gateway.model.service.EcommerceService;
import com.hcmut.gateway.model.service.NotificationService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/v1/ecommerce")
//same controller for all external call to ecommerce service
public class EcommerceController {
    private final EcommerceService ecommerceService;
    private final NotificationService notificationService;

    @Autowired
    public EcommerceController(EcommerceService ecommerceService, NotificationService notificationService) {
        this.ecommerceService = ecommerceService;
        this.notificationService = notificationService;
    }

    @PostMapping(value = "/register-livestream-product-follower")
    public ResponseEntity<Void> registerLivestreamProductFollower(
            @RequestBody RegisterLivestreamProductFollowerRequest request) {
        ecommerceService.registerLivestreamProductFollower(request);
        return ResponseEntity.ok().build();
    }

    @GetMapping(value = "/notify-livestream-product-follower")
    public ResponseEntity<Void> notifyLivestreamProductFollower(@RequestParam("productId") Long productId) {
        ecommerceService.notifyLivestreamProductFollower(productId);
        return ResponseEntity.ok().build();
    }
}
