package com.hcmut.shared_lib.model.repository;

import com.hcmut.shared_lib.model.entity.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.util.Assert;

import java.util.List;
import java.util.Optional;

public interface UserRepository extends JpaRepository<User, Long>{

    Optional<User> findByUsername(String username);

    @Query("SELECT a FROM User a WHERE a.username = ?1 OR a.email = ?2")
    List<User> findByUsernameOrEmail(String username, String email);

    default Optional<User> findByUsernameOrEmail(String usernameOrEmail) {
        usernameOrEmail = usernameOrEmail.trim();
        List<User> users = findByUsernameOrEmail(usernameOrEmail, usernameOrEmail.toLowerCase());
        Assert.isTrue(users.size() <= 1, "Multiple accounts found for username or email: " + usernameOrEmail);
        return users.stream().findFirst();
    }
}
