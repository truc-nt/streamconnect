package com.hcmut.streamconnect.model.service;

import com.hcmut.streamconnect.controller.AuthenticationController.LoginUserDto;
import com.hcmut.streamconnect.model.entity.Account;
import com.hcmut.streamconnect.model.repository.AccountRepository;
import java.time.LocalDateTime;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

@Service
public class AuthenticationServiceImpl implements AuthenticationService {

    private final AccountRepository accountRepository;
    private final PasswordEncoder passwordEncoder;
    private final AuthenticationManager authenticationManager;

    @Autowired
    public AuthenticationServiceImpl(
            AccountRepository accountRepository,
            AuthenticationManager authenticationManager,
            PasswordEncoder passwordEncoder
    ) {
        this.authenticationManager = authenticationManager;
        this.accountRepository = accountRepository;
        this.passwordEncoder = passwordEncoder;
    }

    public Account authenticate(LoginUserDto input) {
        authenticationManager.authenticate(
                new UsernamePasswordAuthenticationToken(
                        input.getUsername(),
                        input.getPassword()
                )
        );

        return accountRepository.findByUsername(input.getUsername())
                .orElseThrow(() -> new IllegalArgumentException("User not found"));
    }

    public Account register(Account account) {
        if (!accountRepository.findByUsernameOrEmail(account.getUsername(), account.getEmail()).isEmpty()) {
            throw new IllegalArgumentException("Username already exists");
        }
        account.setHashedPassword(passwordEncoder.encode(account.getPassword()));
        account.setCreatedDateTime(LocalDateTime.now());
        account.setLastUpdatedDateTime(LocalDateTime.now());
        return accountRepository.save(account);
    }
}
