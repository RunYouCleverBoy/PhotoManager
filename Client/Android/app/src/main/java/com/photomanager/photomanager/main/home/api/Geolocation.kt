package com.photomanager.photomanager.main.home.api

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class Geolocation(
    @SerialName("latitude") val latitude: Double,
    @SerialName("longitude") val longitude: Double
)

@Serializable
data class Place(
    @SerialName("name") val name: String,
    @SerialName("aliases") val aliases: List<String>,
    @SerialName("city") val city: String,
    @SerialName("country") val country: String
)