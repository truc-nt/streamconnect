package com.hcmut.streamconnect.account;

import java.util.List;

public class Role {
    public static final Role VIEWER = new Role("BUYER", "Viewer of livestream");
    public static final Role SELLER = new Role("SELLER", "Seller of livestream");
    public static final List<Role> ALL_ROLES = List.of(VIEWER, SELLER);
    private String id;
    private String name;

    // List of roles that can be created by this role
    private Role() {}

    private Role(String id, String name) {
        this.id = id;
        this.name = name;
    }

    public static Role getRoleById(String id) {
        for (Role role: ALL_ROLES) {
            if (role.id.equals(id)) {
                return role;
            }
        }
        return null;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
