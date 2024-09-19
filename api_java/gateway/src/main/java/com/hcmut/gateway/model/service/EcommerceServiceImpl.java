package com.hcmut.gateway.model.service;

import com.hcmut.gateway.configuration.ExternalServiceClient;
import com.hcmut.gateway.model.DTO.LivestreamProductFollower;
import com.hcmut.gateway.model.DTO.RegisterLivestreamProductFollowerRequest;
import com.hcmut.gateway.model.external_request_model.CreateShopForNewUserRequest;
import com.hcmut.shared_lib.common_util.ExternalRequestUtils;
import com.hcmut.shared_lib.model.constant.NotificationType;
import com.hcmut.shared_lib.model.entity.Notification;
import com.hcmut.shared_lib.model.entity.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.core.ParameterizedTypeReference;
import org.springframework.http.HttpMethod;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

import java.util.Collections;
import java.util.List;

@Service
public class EcommerceServiceImpl implements EcommerceService {
    private final ExternalServiceClient ecommerceServiceClient;
    private final NotificationService notificationService;
    private final SocketService socketService;

    @Autowired
    public EcommerceServiceImpl(
            @Qualifier("ecommerceServiceClient") ExternalServiceClient ecommerceServiceClient,
            NotificationService notificationService,
            SocketService socketService
    ) {
        this.ecommerceServiceClient = ecommerceServiceClient;
        this.notificationService = notificationService;
        this.socketService = socketService;
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

    @Override
    public void registerLivestreamProductFollower(RegisterLivestreamProductFollowerRequest request) {
        final String path = String.format("/api/livestreams/%d/register_product_follower", request.getIdLivestream());
        String uriString = ecommerceServiceClient.getUriBuilder().path(path).toUriString();
        ecommerceServiceClient.getBodySpec(HttpMethod.POST, uriString).contentType(MediaType.APPLICATION_JSON)
                .body(request).retrieve()
                .onStatus(status -> status.is4xxClientError() || status.is5xxServerError(),
                        ExternalRequestUtils::handleErrorRequest
                ).toBodilessEntity();
    }

    @Override
    public void notifyLivestreamProductFollower(Long productId) {
        List<Long> followerIds = getLivestreamProductFollowerIds(productId);
        List<Notification> notifications = followerIds.stream()
                .map(userId -> buildProductFollowerNotification(userId, productId))
                .toList();
        notifications = notificationService.batchCreateNewNotification(notifications);
        for (Notification notification : notifications) {
            socketService.notifyUser(notification.getUserId(), notification);
        }
    }

    private List<Long> getLivestreamProductFollowerIds(Long productId) {
        final String path = String.format("/api/livestreams/-/livestream_products/%d/followers", productId);
        String uriString = ecommerceServiceClient.getUriBuilder().path(path).toUriString();
        ResponseEntity<List<LivestreamProductFollower>> response = ecommerceServiceClient.getBodySpec(HttpMethod.GET, uriString).retrieve()
                .onStatus(status -> status.is4xxClientError() || status.is5xxServerError(),
                        ExternalRequestUtils::handleErrorRequest
                ).toEntity(new ParameterizedTypeReference<>() {});
        if (response.getBody() == null) {
            return Collections.emptyList();
        }
        return response.getBody().stream().map(LivestreamProductFollower::getIdUser).toList();
    }

    private Notification buildProductFollowerNotification(Long userId, Long productId) {
        Notification notification = new Notification();
        notification.setUserId(userId);
        notification.setTitle("Thông báo sản phẩm");
        notification.setMessage("Sản phẩm bạn theo dõi đang được bày bán " + productId);
        notification.setType(NotificationType.ORDER.getKey());
        return notification;
    }
}
