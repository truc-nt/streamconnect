package com.hcmut.shared_lib.model.constant;

public enum NotificationType implements KeyEnum<NotificationType> {
    MESSAGE("message"), ORDER("order"), PRODUCT("product");
    private final String key;

    private NotificationType(String key) {
        this.key = key;
    }

    public String getKey() {
        return key;
    }

    public static NotificationType get(String key) {
        return KeyEnumUtils.get(NotificationType.class, key);
    }

}