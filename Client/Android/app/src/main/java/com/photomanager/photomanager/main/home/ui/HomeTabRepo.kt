package com.photomanager.photomanager.main.home.ui

import androidx.annotation.StringRes
import com.photomanager.photomanager.main.home.model.WorkflowStage

interface HomeTabRepo {
    data class TabDescriptor(@StringRes val title: Int, val stage: WorkflowStage)
    fun getTabs(): List<TabDescriptor>
}