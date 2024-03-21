package com.photomanager.photomanager.main.home.ui

import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Tab
import androidx.compose.material3.TabRow
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.photomanager.photomanager.main.home.model.WorkflowStage

@Composable
fun HomeTabs(
    tabs: List<HomeTabRepo.TabDescriptor>,
    workflowStage: WorkflowStage,
    onTabSelected: (WorkflowStage) -> Unit
) {
    TabRow(selectedTabIndex = tabs.indexOfFirst { it.stage == workflowStage }) {
        tabs.forEach { tab ->
            Tab(
                modifier = Modifier.padding(8.dp),
                selected = workflowStage == tab.stage,
                onClick = { onTabSelected(tab.stage) }) {
                Text(text = stringResource(id = tab.title))
            }
        }
    }
}