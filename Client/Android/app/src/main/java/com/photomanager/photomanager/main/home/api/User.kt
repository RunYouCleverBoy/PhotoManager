package com.photomanager.photomanager.main.home.api

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

enum class Role(val value: String) {
    RoleAdmin("admin"),
    RoleUser("user")
}

@Serializable
data class User(
    @SerialName("id") val id: String,
    @SerialName("name") val name: String,
    @SerialName("email") val email: String,
    @SerialName("password") val password: String,
    @SerialName("token") val token: String?,
    @SerialName("refreshToken") val refreshToken: String?,
    @SerialName("role") val role: Role?
)