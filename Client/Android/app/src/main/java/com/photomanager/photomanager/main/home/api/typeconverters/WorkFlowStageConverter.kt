package com.photomanager.photomanager.main.home.api.typeconverters

import com.photomanager.photomanager.main.home.api.model.WorkflowStage
import kotlinx.serialization.KSerializer
import kotlinx.serialization.descriptors.PrimitiveKind
import kotlinx.serialization.descriptors.PrimitiveSerialDescriptor
import kotlinx.serialization.encoding.Decoder
import kotlinx.serialization.encoding.Encoder

object WorkflowStageSerializer : KSerializer<WorkflowStage> {
    override val descriptor = PrimitiveSerialDescriptor("WorkflowStage", PrimitiveKind.STRING)

    override fun serialize(encoder: Encoder, value: WorkflowStage) {
        encoder.encodeString(value.value)
    }

    override fun deserialize(decoder: Decoder): WorkflowStage {
        val value = decoder.decodeString()
        return WorkflowStage.entries.find { it.value == value } ?: throw IllegalArgumentException("Unknown value for WorkflowStage: $value")
    }
}