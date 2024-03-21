package com.photomanager.photomanager.main.home

import androidx.compose.material3.Tab
import androidx.compose.material3.TabRow
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.res.stringResource
import com.photomanager.photomanager.R
import com.photomanager.photomanager.main.home.model.WorkflowStage

@Composable
 fun HomeTabs(
    state: HomeState,
    onTabSelected: (WorkflowStage) -> Unit
) {
    TabRow(selectedTabIndex = state.currentMode.ordinal) {
        Tab(
            selected = state.currentMode == WorkflowStage.FOOTAGE,
            onClick = { onTabSelected(WorkflowStage.FOOTAGE) }) {
            Text(text = stringResource(id = R.string.footage))
        }
        Tab(
            selected = state.currentMode == WorkflowStage.COLLECTION,
            onClick = { onTabSelected(WorkflowStage.COLLECTION) }) {
            Text(text = stringResource(id = R.string.collection))
        }
    }
}