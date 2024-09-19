package com.hcmut.gateway.model.DTO;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.time.LocalDateTime;

public class LivestreamProductFollower {

    @JsonProperty("fk_user")
    private Long idUser;

    @JsonProperty("fk_livestream_product")
    private Long idLivestreamProduct;

    @JsonProperty("created_at")
    private LocalDateTime createdAt;

    public Long getIdUser() {
        return idUser;
    }

    public Long getIdLivestreamProduct() {
        return idLivestreamProduct;
    }

    public LocalDateTime getCreatedAt() {
        return createdAt;
    }

    public void setIdUser(Long idUser) {
        this.idUser = idUser;
    }

    public void setIdLivestreamProduct(Long idLivestreamProduct) {
        this.idLivestreamProduct = idLivestreamProduct;
    }

    public void setCreatedAt(LocalDateTime createdAt) {
        this.createdAt = createdAt;
    }
}
