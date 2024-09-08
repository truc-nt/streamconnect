package com.hcmut.gateway.model.service;

import com.hcmut.gateway.controller.AuthenticationController.LoginUserDto;
import com.hcmut.shared_lib.model.entity.User;
import com.hcmut.shared_lib.model.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;

@Service
public class AuthenticationServiceImpl implements AuthenticationService {

    private final UserRepository userRepository;
    private final PasswordEncoder passwordEncoder;
    private final AuthenticationManager authenticationManager;

    @Autowired
    public AuthenticationServiceImpl(
            UserRepository userRepository,
            AuthenticationManager authenticationManager,
            PasswordEncoder passwordEncoder
    ) {
        this.authenticationManager = authenticationManager;
        this.userRepository = userRepository;
        this.passwordEncoder = passwordEncoder;
    }

    public User authenticate(LoginUserDto input) {
        authenticationManager.authenticate(
                new UsernamePasswordAuthenticationToken(
                        input.getUsername(),
                        input.getPassword()
                )
        );

        return userRepository.findByUsername(input.getUsername())
                .orElseThrow(() -> new IllegalArgumentException("User not found"));
    }

    public User register(User account) {
        if (!userRepository.findByUsernameOrEmail(account.getUsername(), account.getEmail()).isEmpty()) {
            throw new IllegalArgumentException("Username already exists");
        }
        account.setHashedPassword(passwordEncoder.encode(account.getPassword()));
        account.setCreatedDateTime(LocalDateTime.now());
        account.setLastUpdatedDateTime(LocalDateTime.now());
        return userRepository.save(account);
    }
}
