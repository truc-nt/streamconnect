package com.hcmut.gateway.model.service;

import com.hcmut.gateway.controller.NotificationController.NotificationStatusUpdateRequest;
import com.hcmut.shared_lib.model.constant.NotificationStatus;
import com.hcmut.shared_lib.model.entity.Notification;
import com.hcmut.shared_lib.model.repository.NotificationRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class NotificationServiceImpl implements NotificationService {

    private final NotificationRepository notificationRepository;

    @Autowired
    public NotificationServiceImpl(NotificationRepository notificationRepository) {
        this.notificationRepository = notificationRepository;
    }

    @Override
    @Transactional
    public void batchUpdateStatus(NotificationStatusUpdateRequest request) {
        if (request.getNotificationIds().isEmpty()) {
            return;
        }
        NotificationStatus status = NotificationStatus.get(request.getStatus());
        if (status == null) {
            throw new IllegalArgumentException("Unknown status: " + request.getStatus());
        }
        List<Notification> notifications = notificationRepository.findAllById(request.getNotificationIds());
        for (Notification notification : notifications) {
            updateNotificationStatus(notification, status);
        }
        notificationRepository.saveAll(notifications);
    }

    @Override
    public Notification createNewNotification(Notification notification) {
        notification.setStatus(NotificationStatus.NEW.name());
        return notificationRepository.save(notification);
    }

    @Override
    public List<Notification> batchCreateNewNotification(List<Notification> notifications) {
        for (Notification notification : notifications) {
            notification.setStatus(NotificationStatus.NEW.name());
        }
        return notificationRepository.saveAll(notifications);
    }

    private void updateNotificationStatus(Notification notification, NotificationStatus newStatus) {
        switch(newStatus) {
            case NEW:
                throw new IllegalArgumentException("Cannot update to NEW status");
            case SEND:
                if (notification.getStatus().equals(NotificationStatus.NEW.name())) {
                    notification.setStatus(NotificationStatus.SEND.name());
                }
                break;
            case READ:
                if (notification.getStatus().equals(NotificationStatus.SEND.name())) {
                    notification.setStatus(NotificationStatus.READ.name());
                }
                break;
            default:
                throw new IllegalArgumentException("Unknown status: " +  newStatus);
        }
    }
}
