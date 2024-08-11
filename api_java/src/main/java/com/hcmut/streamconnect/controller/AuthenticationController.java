package com.hcmut.streamconnect.controller;

import com.hcmut.streamconnect.configuration.JwtService;
import com.hcmut.streamconnect.model.entity.Account;
import com.hcmut.streamconnect.model.service.AuthenticationService;
import com.hcmut.streamconnect.model.service.AuthenticationServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1/auth")
public class AuthenticationController {
    private final JwtService jwtService;

    private final AuthenticationService authenticationService;

    @Autowired
    public AuthenticationController(JwtService jwtService, AuthenticationService authenticationService) {
        this.jwtService = jwtService;
        this.authenticationService = authenticationService;
    }

    @PostMapping("/login")
    public ResponseEntity<LoginResponse> authenticate(@RequestBody LoginUserDto loginUserDto) {
        Account authenticatedUser = authenticationService.authenticate(loginUserDto);
        String jwtToken = jwtService.generateToken(authenticatedUser);
        LoginResponse loginResponse = new LoginResponse().setToken(jwtToken).setExpiresIn(jwtService.getExpirationTime());
        return ResponseEntity.ok(loginResponse);
    }

    @PostMapping("/register")
    public ResponseEntity<Account> register(@RequestBody Account account) {
        Account registeredAccount = authenticationService.register(account);
        return ResponseEntity.ok(registeredAccount);
    }

    public static class LoginUserDto {
        private String username;
        private String password;

        public String getUsername() {
            return username;
        }

        public String getPassword() {
            return password;
        }
    }

    public static class LoginResponse {
        private String token;

        private long expiresIn;

        public String getToken() {
            return token;
        }

        public long getExpiresIn() {
            return expiresIn;
        }

        public LoginResponse setToken(String token) {
            this.token = token;
            return this;
        }

        public LoginResponse setExpiresIn(long expiresIn) {
            this.expiresIn = expiresIn;
            return this;
        }
    }
}
