package com.hcmut.gateway.controller;

import com.hcmut.gateway.model.service.SocketService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1/test")
public class TestController {

    private final SocketService socketService;

    @Autowired
    public TestController(SocketService socketService) {
        this.socketService = socketService;
    }

    @GetMapping("/public/hello")
    public String hello() {
        return "Hello World!";
    }

    @GetMapping("/secret")
    public String secret() {
        return "Secret";
    }
//
//    @PostMapping("/public/ping-user")
//    public void pingUser(@RequestBody PingMessage message) {
//        socketService.notifyUser(message.getUsername(), message.getMessage());
//    }

    public static class PingMessage {
        public String username;
        public String message;

        public PingMessage() {
        }

        public String getUsername() {
            return username;
        }

        public void setUsername(String username) {
            this.username = username;
        }

        public String getMessage() {
            return message;
        }

        public void setMessage(String message) {
            this.message = message;
        }
    }

}
