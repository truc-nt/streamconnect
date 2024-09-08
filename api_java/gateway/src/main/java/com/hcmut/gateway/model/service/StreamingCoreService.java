package com.hcmut.gateway.model.service;

import com.hcmut.shared_lib.model.DTO.LivestreamDTO;
import com.hcmut.shared_lib.model.entity.Livestream;

import java.util.List;

public interface StreamingCoreService {
    Livestream createLivestream(LivestreamDTO request);
    List<Livestream> fetchLivestreams(String status, boolean fetchAll);
}
