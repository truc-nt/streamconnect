package com.hcmut.streamconnect.controller;

import com.hcmut.streamconnect.model.DTO.LivestreamDTO;
import com.hcmut.streamconnect.model.entity.Livestream;
import com.hcmut.streamconnect.model.service.LivestreamService;
import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1/livestream")
public class LivestreamController {
    private final LivestreamService livestreamService;

    @Autowired
    public LivestreamController(LivestreamService livestreamService) {
        this.livestreamService = livestreamService;
    }

//    @PreAuthorize("hasAnyAuthority('SELLER')")
    @PostMapping("")
    public ResponseEntity<Livestream> createLivestream(@RequestBody LivestreamDTO livestreamDTO) {
        Livestream createdLivestream = livestreamService.createLivestream(livestreamDTO);
        return ResponseEntity.ok(createdLivestream);
    }

    @GetMapping("")
    public ResponseEntity<List<Livestream>> fetchLivestreams(
            @RequestParam(value = "status", required = false) String status,
            @RequestParam(value = "fetchAll", required = false, defaultValue = "false") boolean fetchAll
    ) {
        List<Livestream> livestreams = livestreamService.fetchLiveStream(status, fetchAll);
        return ResponseEntity.ok(livestreams);
    }

    @GetMapping("/start/{livestreamId}")
    public ResponseEntity<Livestream> startLivestream(@PathVariable("livestreamId") Long livestreamId) {
        Livestream livestream = livestreamService.startLivestream(livestreamId);
        return ResponseEntity.ok(livestream);
    }
}
