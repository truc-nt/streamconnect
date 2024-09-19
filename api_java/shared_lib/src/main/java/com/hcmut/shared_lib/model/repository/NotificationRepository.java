package com.hcmut.shared_lib.model.repository;

import com.hcmut.shared_lib.model.entity.Notification;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import java.util.List;
import java.util.Optional;

public interface NotificationRepository extends JpaRepository<Notification, Long> {

    @Query("SELECT n FROM Notification n WHERE n.userId = :userId AND (:status IS NULL OR n.status = :status)")
    List<Notification> findAllByUserId(@Param("userId") long userId, @Param("status") String status);
    Optional<Notification> findById(long id);
}
