package com.hcmut.gateway.controller;

import com.hcmut.gateway.account.CurrentUserDetails;
import com.hcmut.gateway.model.service.NotificationService;
import com.hcmut.gateway.util.AccountUtils;
import com.hcmut.shared_lib.model.entity.Notification;
import com.hcmut.shared_lib.model.repository.NotificationRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/v1/notification")
public class NotificationController {

    private final NotificationRepository notificationRepository;
    private final NotificationService notificationService;

    @Autowired
    public NotificationController(NotificationRepository notificationRepository, NotificationService notificationService) {
        this.notificationRepository = notificationRepository;
        this.notificationService = notificationService;
    }

    //write one endpoint to fetch all, and one to fetch by id
    @GetMapping(value = "")
    public ResponseEntity<List<Notification>> getAllNotifications(
            @RequestParam(value = "userId", required = false) Long userId,
            @RequestParam(value = "status", required = false) String status
    ) {
        CurrentUserDetails userDetails = AccountUtils.getCurrentUserDetails();
        if (userId == null) {
            userId = userDetails.getId();
        }
        List<Notification> notifications = notificationRepository.findAllByUserId(userId, status);
        return ResponseEntity.ok(notifications);
    }

    @GetMapping(value = "/{notification_id}")
    public ResponseEntity<Notification> getNotificationById(@PathVariable("notification_id") Long notificationId) {
        Notification notification = notificationRepository.findById(notificationId)
                .orElseThrow(() -> new IllegalArgumentException("Notification not found"));
        return ResponseEntity.ok(notification);
    }

    @PostMapping(value = "/batch-status-update")
    public ResponseEntity<Void> batchUpdateStatus(@RequestBody NotificationStatusUpdateRequest request) {
        notificationService.batchUpdateStatus(request);
        return ResponseEntity.ok().build();
    }

    public static class NotificationStatusUpdateRequest {
        private List<Long> notificationIds;
        private String status;

        public List<Long> getNotificationIds() {
            return notificationIds;
        }

        public String getStatus() {
            return status;
        }
    }
}
