package com.hcmut.core;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.domain.EntityScan;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;

@SpringBootApplication
@ComponentScan({"com.hcmut.core", "com.hcmut.shared_lib.model.repository"})
@EnableJpaRepositories("com.hcmut.shared_lib.model.repository")
@EntityScan("com.hcmut.shared_lib.model.entity")
public class LivestreamServiceMain {
    public static void main(String[] args) {
        SpringApplication.run(LivestreamServiceMain.class, args);
    }
}