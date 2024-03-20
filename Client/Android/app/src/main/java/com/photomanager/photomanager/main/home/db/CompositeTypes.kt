package com.photomanager.photomanager.main.home.db

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class Comment(
    val id: String,
    val text: String,
    val date: Long,
    val userId: String,
    val userName: String
)

@Serializable
data class Comments(val comments: List<Comment>)

@Serializable
data class UserVisibility(@SerialName("userIds") val userIds: List<String>)
