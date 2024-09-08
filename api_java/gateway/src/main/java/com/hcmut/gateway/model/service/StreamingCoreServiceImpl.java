package com.hcmut.gateway.model.service;

import com.hcmut.gateway.account.CurrentUserDetails;
import com.hcmut.gateway.configuration.ExternalServiceClient;
import com.hcmut.gateway.util.AccountUtils;
import com.hcmut.shared_lib.model.DTO.LivestreamDTO;
import com.hcmut.shared_lib.model.entity.Livestream;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.core.ParameterizedTypeReference;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.util.StringUtils;

import java.util.List;

@Service
public class StreamingCoreServiceImpl implements StreamingCoreService {
    private final ExternalServiceClient coreServiceClient;

    @Autowired
    public StreamingCoreServiceImpl(@Qualifier("coreServiceClient") ExternalServiceClient coreServiceClient) {
        this.coreServiceClient = coreServiceClient;
    }

    public Livestream createLivestream(LivestreamDTO request) {
        CurrentUserDetails userDetails = AccountUtils.getCurrentUserDetails();
        request.setCreatorId(userDetails.getId());
        String uriString = coreServiceClient.getUriBuilder().path("/api/v1/livestream").toUriString();
        ResponseEntity<Livestream> response = coreServiceClient.getRestClient().post()
                .uri(uriString).contentType(MediaType.APPLICATION_JSON).body(request)
                .retrieve().toEntity(Livestream.class);
        return response.getBody();
    }
    public List<Livestream> fetchLivestreams(String status, boolean fetchAll) {
        var uriBuilder = coreServiceClient.getUriBuilder().path("/api/v1/livestream");
        if (StringUtils.hasText(status)) {
            uriBuilder.queryParam("status", status);
        }
        if (!fetchAll) {
            CurrentUserDetails userDetails = AccountUtils.getCurrentUserDetails();
            uriBuilder.queryParam("ownerId", userDetails.getId());
        }
        return coreServiceClient.getRestClient().get().uri(uriBuilder.toUriString()).retrieve()
                .body(new ParameterizedTypeReference<>() {});

    }

}
