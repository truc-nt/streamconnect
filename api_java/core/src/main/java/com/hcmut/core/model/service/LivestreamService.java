package com.hcmut.core.model.service;


import com.hcmut.shared_lib.model.DTO.LivestreamDTO;
import com.hcmut.shared_lib.model.entity.Livestream;

import java.util.List;

public interface LivestreamService {

    Livestream createLivestream(LivestreamDTO livestream);

//    Livestream startLivestream(Long livestreamId);

    List<Livestream> fetchLivestreams(String status, Long ownerId);
}
