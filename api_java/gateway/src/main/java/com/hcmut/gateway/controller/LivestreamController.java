package com.hcmut.gateway.controller;


import com.hcmut.gateway.model.service.StreamingCoreService;
import com.hcmut.shared_lib.model.DTO.LivestreamDTO;
import com.hcmut.shared_lib.model.entity.Livestream;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/v1/livestream")
public class LivestreamController {

    private final StreamingCoreService streamingCoreService;

    @Autowired
    public LivestreamController(StreamingCoreService streamingCoreService) {
        this.streamingCoreService = streamingCoreService;
    }

//    @PreAuthorize("hasAnyAuthority('SELLER')")
    @PostMapping("")
    public ResponseEntity<Livestream> createLivestream(@RequestBody LivestreamDTO livestreamDTO) {
        Livestream createdLivestream = streamingCoreService.createLivestream(livestreamDTO);
        return ResponseEntity.ok(createdLivestream);
    }

    @GetMapping("")
    public ResponseEntity<List<Livestream>> fetchLivestreams(
            @RequestParam(value = "status", required = false) String status,
            @RequestParam(value = "fetchAll", required = false, defaultValue = "false") boolean fetchAll
    ) {
        List<Livestream> livestreams = streamingCoreService.fetchLivestreams(status, fetchAll);
        return ResponseEntity.ok(livestreams);
    }
//
//    @GetMapping("/start/{livestreamId}")
//    public ResponseEntity<Livestream> startLivestream(@PathVariable("livestreamId") Long livestreamId) {
//        Livestream livestream = livestreamService.startLivestream(livestreamId);
//        return ResponseEntity.ok(livestream);
//    }
}
