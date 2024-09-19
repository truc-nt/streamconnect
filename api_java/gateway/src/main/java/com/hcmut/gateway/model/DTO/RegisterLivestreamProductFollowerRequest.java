package com.hcmut.gateway.model.DTO;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public class RegisterLivestreamProductFollowerRequest {

    @JsonProperty("id_livestream_products")
    private List<Long> idLivestreamProducts;

    @JsonProperty("id_livestream")
    private Long idLivestream;

    @JsonProperty("id_user")
    private Long idUser;

    public List<Long> getIdLivestreamProducts() {
        return idLivestreamProducts;
    }

    public void setIdLivestreamProducts(List<Long> idLivestreamProducts) {
        this.idLivestreamProducts = idLivestreamProducts;
    }

    public Long getIdLivestream() {
        return idLivestream;
    }

    public void setIdLivestream(Long idLivestream) {
        this.idLivestream = idLivestream;
    }

    public Long getIdUser() {
        return idUser;
    }
}
