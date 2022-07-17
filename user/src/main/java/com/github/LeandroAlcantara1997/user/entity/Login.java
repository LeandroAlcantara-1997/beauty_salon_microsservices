package com.github.LeandroAlcantara1997.user.entity;

import lombok.AccessLevel;
import lombok.Data;
import lombok.Setter;

@Data
public class Login {
    @Setter(value = AccessLevel.PRIVATE)
    private Long id;
    private String password;
}