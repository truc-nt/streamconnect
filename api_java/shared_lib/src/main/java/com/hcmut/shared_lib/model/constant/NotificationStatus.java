package com.hcmut.shared_lib.model.constant;

public enum NotificationStatus implements KeyEnum<NotificationStatus> {
    NEW("NEW"), SEND("SEND"), READ("READ");
    private final String key;

    private NotificationStatus(String key) {
        this.key = key;
    }

    public String getKey() {
        return key;
    }

    public static NotificationStatus get(String key) {
        return KeyEnumUtils.get(NotificationStatus.class, key);
    }

}
