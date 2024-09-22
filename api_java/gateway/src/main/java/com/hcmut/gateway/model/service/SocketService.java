package com.hcmut.gateway.model.service;

import com.hcmut.shared_lib.model.entity.Notification;
import com.hcmut.shared_lib.model.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.messaging.simp.SimpMessagingTemplate;
import org.springframework.stereotype.Service;

@Service
public class SocketService {
    private final SimpMessagingTemplate messagingTemplate;
    private final UserRepository userRepository;
    private final NotificationService notificationService;

    @Autowired
    public SocketService(SimpMessagingTemplate messagingTemplate, UserRepository userRepository, NotificationService notificationService) {
        this.messagingTemplate = messagingTemplate;
        this.userRepository = userRepository;
        this.notificationService = notificationService;
    }

//    @Transactional
//    public void notifyUser(final String userName, final String message) {
//        User user = userRepository.findByUsername(userName).orElseThrow(() -> new RuntimeException("User not found"));
//        Notification newNoti = new Notification();
//        newNoti.setUserId(user.getId());
//        newNoti.setTitle("Test title");
//        newNoti.setMessage(message);
//        newNoti.setType(NotificationType.MESSAGE.getKey());
//        newNoti = notificationService.createNewNotification(newNoti);
//        messagingTemplate.convertAndSendToUser(user.getUsername(), "/topic/notification", newNoti);
//    }

    public void notifyUser(long userId, Notification notification) {
        messagingTemplate.convertAndSendToUser(String.valueOf(userId), "/topic/notification", notification);
    }
}
