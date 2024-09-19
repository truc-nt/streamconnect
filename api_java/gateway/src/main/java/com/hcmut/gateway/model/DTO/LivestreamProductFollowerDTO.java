package com.hcmut.gateway.model.DTO;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public class LivestreamProductFollowerDTO {

    @JsonProperty("user_ids")
    private List<Long> userIds;

    @JsonProperty("livestream")
    private LivestreamDTO livestream;

    @JsonProperty("livestream_product")
    private ProductDTO livestreamProduct;

    //setter and getter
    public List<Long> getUserIds() {
        return userIds;
    }

    public void setUserIds(List<Long> userIds) {
        this.userIds = userIds;
    }

    public LivestreamDTO getLivestream() {
        return livestream;
    }

    public void setLivestream(LivestreamDTO livestreamDTO) {
        this.livestream = livestreamDTO;
    }

    public ProductDTO getLivestreamProduct() {
        return livestreamProduct;
    }

    public void setLivestreamProduct(ProductDTO productDTO) {
        this.livestreamProduct = productDTO;
    }
}
