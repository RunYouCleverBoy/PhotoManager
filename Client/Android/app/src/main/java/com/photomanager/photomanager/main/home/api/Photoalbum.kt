package com.photomanager.photomanager.main.home.api

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

typealias ObjectId = String

@Serializable
data class PhotoAlbum(
    @SerialName("id") val id: ObjectId,
    @SerialName("cover_image_url") val coverImageUrl: String,
    @SerialName("name") val name: String,
    @SerialName("description") val description: String,
    @SerialName("owner") val owner: ObjectId,
    @SerialName("visible_to") val visibleTo: List<ObjectId>
)