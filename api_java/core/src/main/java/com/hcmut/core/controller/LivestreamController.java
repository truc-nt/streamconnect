package com.hcmut.core.controller;


import com.hcmut.core.model.service.LivestreamService;
import com.hcmut.shared_lib.model.DTO.LivestreamDTO;
import com.hcmut.shared_lib.model.entity.Livestream;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/v1/livestream")
public class LivestreamController {
    private final LivestreamService livestreamService;

    @Autowired
    public LivestreamController(LivestreamService livestreamService) {
        this.livestreamService = livestreamService;
    }

    @PostMapping("")
    public ResponseEntity<Livestream> createLivestream(@RequestBody LivestreamDTO livestreamDTO) {
        Livestream createdLivestream = livestreamService.createLivestream(livestreamDTO);
        return ResponseEntity.ok(createdLivestream);
    }

    @GetMapping("")
    public ResponseEntity<List<Livestream>> fetchLivestreams(
            @RequestParam(value = "status", required = false) String status,
            @RequestParam(value = "ownerId", required = false) Long ownerId
    ) {
        List<Livestream> livestreams = livestreamService.fetchLivestreams(status, ownerId);
        return ResponseEntity.ok(livestreams);
    }
//
//    @GetMapping("/start/{livestreamId}")
//    public ResponseEntity<Livestream> startLivestream(@PathVariable("livestreamId") Long livestreamId) {
//        Livestream livestream = livestreamService.startLivestream(livestreamId);
//        return ResponseEntity.ok(livestream);
//    }
}
