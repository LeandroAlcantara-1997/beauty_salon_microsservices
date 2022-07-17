package com.github.LeandroAlcantara1997.user.entity;

import java.util.Date;

import lombok.AccessLevel;
import lombok.Data;
import lombok.Setter;

@Data
public class User {
    @Setter(value = AccessLevel.PRIVATE)
    private Long id;
    private String name;
    private String lastName;
    private String cpf;
    private String rg;
    private Date birthDate;
    private Contact contact;
    private Address address;
    private Login login;
    
}
