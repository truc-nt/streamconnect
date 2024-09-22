package com.hcmut.gateway.configuration;

import com.hcmut.gateway.account.CurrentUserDetails;
import com.hcmut.gateway.util.AccountUtils;
import org.springframework.http.server.ServerHttpRequest;
import org.springframework.stereotype.Service;
import org.springframework.web.socket.WebSocketHandler;
import org.springframework.web.socket.server.support.DefaultHandshakeHandler;

import java.security.Principal;
import java.util.Map;

@Service
public class UserHandshakeHandler extends DefaultHandshakeHandler {
    @Override
    protected Principal determineUser(ServerHttpRequest request, WebSocketHandler wsHandler, Map<String, Object> attributes) {
        CurrentUserDetails currentUserDetails = AccountUtils.getCurrentUserDetails();
        return () -> currentUserDetails.getId().toString();
    }
}
