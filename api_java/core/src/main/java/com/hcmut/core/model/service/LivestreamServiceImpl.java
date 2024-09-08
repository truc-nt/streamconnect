package com.hcmut.core.model.service;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.hcmut.core.model.constant.LivestreamStatus;
import com.hcmut.shared_lib.model.DTO.LivestreamDTO;
import com.hcmut.shared_lib.model.entity.Livestream;
import com.hcmut.shared_lib.model.repository.LivestreamRepository;
import com.hcmut.shared_lib.model.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Service;
import org.springframework.util.StringUtils;
import org.springframework.web.client.RestClient;

import java.util.List;

@Service
public class LivestreamServiceImpl implements LivestreamService {
    private static final String VIDEO_SDK_BASE_URL = "https://api.videosdk.live/v2";

    @Value("${videosdk.token}")
    private String videoSdkToken;

    private final UserRepository userRepository;

    private final LivestreamRepository livestreamRepository;

    private final RestClient restClient = RestClient.create();

    private final ObjectMapper objectMapper = new ObjectMapper();

    @Autowired
    public LivestreamServiceImpl(UserRepository userRepository, LivestreamRepository livestreamRepository) {
        this.userRepository = userRepository;
        this.livestreamRepository = livestreamRepository;
    }

    @Override
    public Livestream createLivestream(LivestreamDTO livestream) {
        Livestream newLiveStream = new Livestream();
        copyForUpdate(livestream, newLiveStream);
        newLiveStream.setOwner(userRepository.findById(livestream.getCreatorId())
                .orElseThrow(() -> new IllegalArgumentException("Owner not found")));
        newLiveStream.setStatus(LivestreamStatus.CREATED.getValue());
        newLiveStream.setMeetingId(createVideoSdkRoom());
        return livestreamRepository.save(newLiveStream);
    }

//    @Override
//    public Livestream startLivestream(Long livestreamId) {
//        Livestream livestream = livestreamRepository.findById(livestreamId)
//                .orElseThrow(() -> new IllegalArgumentException("Livestream not found"));
//        if (!LivestreamStatus.CREATED.getValue().equals(livestream.getStatus())) {
//            throw new IllegalArgumentException("Livestream is not in created state");
//        }
//        livestream.setStatus(LivestreamStatus.STREAMING.getValue());
//        livestream.setStartTime(LocalDateTime.now());
//        livestream.setHlsUrl(startRoomHsl(livestream.getMeetingId()));
//        return livestreamRepository.save(livestream);
//    }

    @Override
    public List<Livestream> fetchLivestreams(String status, Long ownerId) {
        if (StringUtils.hasText(status) && LivestreamStatus.fromValue(status) == null) {
            throw new IllegalArgumentException("Invalid status");
        }
        if (ownerId == null) {
            return livestreamRepository.findAllByStatus(status);
        }
        return livestreamRepository.findAllByStatusAndOwnerId(status, ownerId);
    }

    private String createVideoSdkRoom() {
        String response = restClient.post().uri(VIDEO_SDK_BASE_URL + "/rooms")
                .contentType(MediaType.APPLICATION_JSON)
                .header("Authorization", videoSdkToken)
                .body("{}").retrieve().body(String.class);
        try {
            JsonNode parent = objectMapper.readTree(response);
            return parent.path("roomId").asText();
        } catch (JsonProcessingException e) {
            throw new RuntimeException("Failed to create room");
        }
    }

//    private String startRoomHsl(String meetingId) {
//        String response = restClient.post().uri(VIDEO_SDK_BASE_URL + "/hls/start")
//                .contentType(MediaType.APPLICATION_JSON)
//                .header("Authorization", videoSdkToken)
//                .body(String.format("{\"roomId\": \"%s\"}", meetingId)).retrieve().body(String.class);
//        try {
//            JsonNode parent = objectMapper.readTree(response);
//            return parent.path("livestreamUrl").asText();
//        } catch (JsonProcessingException e) {
//            throw new RuntimeException("Failed to start hls");
//        }
//    }

    private void copyForUpdate(LivestreamDTO livestream, Livestream newLiveStream) {
        newLiveStream.setTitle(livestream.getTitle());
        newLiveStream.setDescription(livestream.getDescription());
    }
//
//    private void validateLivestream(LivestreamDTO livestream) {
//        if (!StringUtils.hasText(livestream.getTitle())) {
//            throw new IllegalArgumentException("Title is required");
//        }
//    }


}
