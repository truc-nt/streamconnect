package com.hcmut.shared_lib.model.repository;

import com.hcmut.shared_lib.model.entity.Livestream;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import java.util.List;

public interface LivestreamRepository extends JpaRepository<Livestream, Long> {

    @Query("SELECT l FROM Livestream l WHERE :status IS NULL OR l.status = :status")
    List<Livestream> findAllByStatus(@Param("status") String status);

    @Query("SELECT l FROM Livestream l WHERE (:status is NULL OR l.status = :status) AND l.owner.id = :ownerId")
    List<Livestream> findAllByStatusAndOwnerId(@Param("status") String status, @Param("ownerId") Long ownerId);
}
