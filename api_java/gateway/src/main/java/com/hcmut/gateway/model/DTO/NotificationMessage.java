package com.hcmut.gateway.model.DTO;

public class NotificationMessage {
    private String message;

    public NotificationMessage() {
    }

    public NotificationMessage(String message) {
        this.message = message;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }
}
