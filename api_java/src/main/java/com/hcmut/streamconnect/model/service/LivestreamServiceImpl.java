package com.hcmut.streamconnect.model.service;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.hcmut.streamconnect.account.CurrentUserDetails;
import com.hcmut.streamconnect.model.DTO.LivestreamDTO;
import com.hcmut.streamconnect.model.constant.LivestreamStatus;
import com.hcmut.streamconnect.model.entity.Livestream;
import com.hcmut.streamconnect.model.repository.AccountRepository;
import com.hcmut.streamconnect.model.repository.LivestreamRepository;
import com.hcmut.streamconnect.util.AccountUtils;
import java.time.LocalDateTime;
import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Service;
import org.springframework.util.StringUtils;
import org.springframework.web.client.RestClient;

@Service
public class LivestreamServiceImpl implements LivestreamService{
    private static final String VIDEO_SDK_BASE_URL = "https://api.videosdk.live/v2";

    @Value("${videosdk.token}")
    private String videoSdkToken;

    private final AccountRepository accountRepository;

    private final LivestreamRepository livestreamRepository;

    private final RestClient restClient = RestClient.create();

    private final ObjectMapper objectMapper = new ObjectMapper();

    @Autowired
    public LivestreamServiceImpl(AccountRepository accountRepository, LivestreamRepository livestreamRepository) {
        this.accountRepository = accountRepository;
        this.livestreamRepository = livestreamRepository;
    }

    @Override
    public Livestream createLivestream(LivestreamDTO livestream) {
        Livestream newLiveStream = new Livestream();
        copyForUpdate(livestream, newLiveStream);
        CurrentUserDetails currentUserDetails = AccountUtils.getCurrentUserDetails();
        newLiveStream.setOwner(accountRepository.findById(currentUserDetails.getId())
                .orElseThrow(() -> new IllegalArgumentException("Owner not found")));
        newLiveStream.setStatus(LivestreamStatus.CREATED.getValue());
        newLiveStream.setMeetingId(createVideoSdkRoom());
        return livestreamRepository.save(newLiveStream);
    }

    @Override
    public Livestream startLivestream(Long livestreamId) {
        Livestream livestream = livestreamRepository.findById(livestreamId)
                .orElseThrow(() -> new IllegalArgumentException("Livestream not found"));
        if (!LivestreamStatus.CREATED.getValue().equals(livestream.getStatus())) {
            throw new IllegalArgumentException("Livestream is not in created state");
        }
        livestream.setStatus(LivestreamStatus.STREAMING.getValue());
        livestream.setStartTime(LocalDateTime.now());
        livestream.setHlsUrl(startRoomHsl(livestream.getMeetingId()));
        return livestreamRepository.save(livestream);
    }

    @Override
    public List<Livestream> fetchLiveStream(String status, boolean fetchAll) {
        if (StringUtils.hasText(status) && LivestreamStatus.fromValue(status) == null) {
            throw new IllegalArgumentException("Invalid status");
        }
        status = StringUtils.hasText(status) ? status : null;
        CurrentUserDetails currentUserDetails = AccountUtils.getCurrentUserDetails();
        if (fetchAll) {
            return livestreamRepository.findAllByStatus(status);
        }
        return livestreamRepository.findAllByStatusAndOwnerId(status, currentUserDetails.getId());
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

    private String startRoomHsl(String meetingId) {
        String response = restClient.post().uri(VIDEO_SDK_BASE_URL + "/hls/start")
                .contentType(MediaType.APPLICATION_JSON)
                .header("Authorization", videoSdkToken)
                .body(String.format("{\"roomId\": \"%s\"}", meetingId)).retrieve().body(String.class);
        try {
            JsonNode parent = objectMapper.readTree(response);
            return parent.path("livestreamUrl").asText();
        } catch (JsonProcessingException e) {
            throw new RuntimeException("Failed to start hls");
        }
    }

    private void copyForUpdate(LivestreamDTO livestream, Livestream newLiveStream) {
        newLiveStream.setTitle(livestream.getTitle());
        newLiveStream.setDescription(livestream.getDescription());
    }

    private void validateLivestream(LivestreamDTO livestream) {
        if (!StringUtils.hasText(livestream.getTitle())) {
            throw new IllegalArgumentException("Title is required");
        }
    }


}
