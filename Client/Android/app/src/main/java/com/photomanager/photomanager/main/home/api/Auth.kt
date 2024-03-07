package com.photomanager.photomanager.main.home.api

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class LoginRequest(
    @SerialName("email") val email: String,
    @SerialName("password") val password: String
)

@Serializable
data class LoginWithTokenRequest(
    @SerialName("oldToken") val oldToken: String,
    @SerialName("refreshToken") val refreshToken: String
)