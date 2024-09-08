package com.hcmut.core.model.constant;

public enum LivestreamStatus {
    CREATED("CREATED"),
    STREAMING("STREAMING"),
    ENDED("ENDED");

    private final String value;

    LivestreamStatus(String value) {
        this.value = value;
    }

    public String getValue() {
        return value;
    }

    public static LivestreamStatus fromValue(String value) {
        for (LivestreamStatus status : LivestreamStatus.values()) {
            if (status.getValue().equals(value)) {
                return status;
            }
        }
        return null;
    }
}
