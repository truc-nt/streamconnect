package com.hcmut.gateway.model.service;

import com.hcmut.gateway.configuration.ExternalServiceClient;
import com.hcmut.gateway.model.DTO.LivestreamProductFollowerDTO;
import com.hcmut.gateway.model.DTO.RegisterLivestreamProductFollowerRequest;
import com.hcmut.gateway.model.external_request_model.CreateShopForNewUserRequest;
import com.hcmut.shared_lib.common_util.ExternalRequestUtils;
import com.hcmut.shared_lib.model.constant.NotificationType;
import com.hcmut.shared_lib.model.entity.Notification;
import com.hcmut.shared_lib.model.entity.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.HttpMethod;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

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
        LivestreamProductFollowerDTO livestreamProductFollower = getLivestreamProductFollower(productId);
        List<Notification> notifications = livestreamProductFollower.getUserIds().stream()
                .map(userId -> buildProductFollowerNotification(userId, livestreamProductFollower))
                .toList();
        notifications = notificationService.batchCreateNewNotification(notifications);
        for (Notification notification : notifications) {
            socketService.notifyUser(notification.getUserId(), notification);
        }
    }

    private LivestreamProductFollowerDTO getLivestreamProductFollower(Long productId) {
        final String path = String.format("/api/livestreams/-/livestream_products/%d/followers", productId);
        String uriString = ecommerceServiceClient.getUriBuilder().path(path).toUriString();
        ResponseEntity<LivestreamProductFollowerDTO> response = ecommerceServiceClient.getBodySpec(HttpMethod.GET, uriString).retrieve()
                .onStatus(status -> status.is4xxClientError() || status.is5xxServerError(),
                        ExternalRequestUtils::handleErrorRequest
                ).toEntity(LivestreamProductFollowerDTO.class);
        return response.getBody();
    }

    private Notification buildProductFollowerNotification(Long userId, LivestreamProductFollowerDTO livestreamProductFollower) {
        Notification notification = new Notification();
        notification.setUserId(userId);
        notification.setTitle("Thông báo sản phẩm");
        notification.setMessage(String.format("Sản phẩm %s đang được bán trực tiếp trên livestream %s",
                livestreamProductFollower.getLivestreamProduct().getName(),
                livestreamProductFollower.getLivestream().getTitle()));
        notification.setType(NotificationType.PRODUCT.getKey());
        return notification;
    }
}
