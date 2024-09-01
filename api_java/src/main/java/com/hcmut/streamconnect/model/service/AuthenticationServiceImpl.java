package com.hcmut.streamconnect.model.service;

import com.hcmut.streamconnect.controller.AuthenticationController.LoginUserDto;
import com.hcmut.streamconnect.model.entity.User;
import com.hcmut.streamconnect.model.repository.UserRepository;
import java.time.LocalDateTime;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

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

    public User register(User user) {
        if (!userRepository.findByUsernameOrEmail(user.getUsername(), user.getEmail()).isEmpty()) {
            throw new IllegalArgumentException("Username already exists");
        }
        validateRegisterUser(user);
        user.setHashedPassword(passwordEncoder.encode(user.getPassword()));
        user.setCreatedDateTime(LocalDateTime.now());
        user.setLastUpdatedDateTime(LocalDateTime.now());
        return userRepository.save(user);
    }

    private void validateRegisterUser(User user) {
        if (user.getUsername() == null || user.getUsername().isEmpty()) {
            throw new IllegalArgumentException("Username cannot be empty");
        }

        if (user.getEmail() == null || user.getEmail().isEmpty() || !user.getEmail().matches("^[A-Za-z0-9+_.-]+@(.+)$")) {
            throw new IllegalArgumentException("Invalid email address");
        }

        if (user.getPassword() == null || user.getPassword().isEmpty() || user.getPassword().length() < 8) {
            throw new IllegalArgumentException("Password must be at least 8 characters long");
        }
    }
}
