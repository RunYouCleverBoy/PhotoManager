package com.photomanager.photomanager.main.home.db

import androidx.room.TypeConverter
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json

class CompositeTypeConverters {
    @TypeConverter
    fun fromComments(value: Comments): String {
        return Json.encodeToString(value)
    }

    @TypeConverter
    fun toComments(value: String): Comments {
        return Json.decodeFromString(value)
    }

    @TypeConverter
    fun fromUserVisibility(value: UserVisibility): String {
        return Json.encodeToString(value)
    }

    @TypeConverter
    fun toUserVisibility(value: String): UserVisibility {
        return Json.decodeFromString(value)
    }

    @TypeConverter
    fun fromWorkFlowStage(value: WorkFlowStage): String {
        return value.value
    }

    @TypeConverter
    fun toWorkFlowStage(value: String): WorkFlowStage {
        return WorkFlowStage.entries.first { it.value == value }
    }
}