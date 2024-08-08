package com.hcmut.streamconnect.model.repository;

import com.hcmut.streamconnect.model.entity.Account;
import java.util.List;
import java.util.Optional;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.util.Assert;

public interface AccountRepository extends JpaRepository<Account, Long>{

    Optional<Account> findByUsername(String username);
    List<Account> findByUsernameOrEmail(String username, String email);

    default Optional<Account> findByUsernameOrEmail(String usernameOrEmail) {
        usernameOrEmail = usernameOrEmail.trim();
        List<Account> accounts = findByUsernameOrEmail(usernameOrEmail, usernameOrEmail.toLowerCase());
        Assert.isTrue(accounts.size() <= 1, "Multiple accounts found for username or email: " + usernameOrEmail);
        return accounts.stream().findFirst();
    }
}
