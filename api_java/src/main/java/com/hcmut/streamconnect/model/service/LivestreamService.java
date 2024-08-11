package com.hcmut.streamconnect.model.service;

import com.hcmut.streamconnect.model.DTO.LivestreamDTO;
import com.hcmut.streamconnect.model.entity.Livestream;
import java.util.List;

public interface LivestreamService {

    Livestream createLivestream(LivestreamDTO livestream);

    Livestream startLivestream(Long livestreamId);

    List<Livestream> fetchLiveStream(String status, boolean fetchAll);
}
