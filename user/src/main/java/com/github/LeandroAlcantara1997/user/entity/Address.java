package com.github.LeandroAlcantara1997.user.entity;

import lombok.AccessLevel;
import lombok.Data;
import lombok.Setter;

@Data
public class Address {
    @Setter(value = AccessLevel.PRIVATE)
    private Long id;
    private String country;
    private String state;
    private String city;
    private String district;
    private String street;
    private String number;
    private String complement;
}
