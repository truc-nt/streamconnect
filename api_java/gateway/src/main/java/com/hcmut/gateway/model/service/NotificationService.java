package com.hcmut.gateway.model.service;

import com.hcmut.gateway.controller.NotificationController;
import com.hcmut.shared_lib.model.entity.Notification;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

public interface NotificationService {
    @Transactional
    void batchUpdateStatus(NotificationController.NotificationStatusUpdateRequest request);

    Notification createNewNotification(Notification notification);

    List<Notification> batchCreateNewNotification(List<Notification> notifications);
}
