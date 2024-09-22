package com.hcmut.gateway.model.DTO;

import com.fasterxml.jackson.annotation.JsonProperty;

public class LivestreamDTO {
    @JsonProperty("id_livestream")
    private long idLivestream;

    @JsonProperty("title")
    private String title;

    @JsonProperty("description")
    private String description;

    @JsonProperty("status")
    private String status;

    // Getters and setters
    public long getIdLivestream() {
        return idLivestream;
    }

    public void setIdLivestream(long idLivestream) {
        this.idLivestream = idLivestream;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }
}