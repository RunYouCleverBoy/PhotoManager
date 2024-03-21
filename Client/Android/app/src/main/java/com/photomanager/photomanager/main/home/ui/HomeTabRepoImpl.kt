package com.photomanager.photomanager.main.home.ui

import com.photomanager.photomanager.R
import com.photomanager.photomanager.main.home.model.WorkflowStage
import javax.inject.Inject

class HomeTabRepoImpl @Inject constructor(): HomeTabRepo {
    override fun getTabs(): List<HomeTabRepo.TabDescriptor> {
        return listOf(
            HomeTabRepo.TabDescriptor(R.string.footage, WorkflowStage.FOOTAGE),
            HomeTabRepo.TabDescriptor(R.string.collection, WorkflowStage.COLLECTION)
        )
    }
}