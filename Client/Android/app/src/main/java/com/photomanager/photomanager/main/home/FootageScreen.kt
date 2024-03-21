package com.photomanager.photomanager.main.home

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.ElevatedButton
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import com.photomanager.photomanager.R
import com.photomanager.photomanager.main.components.ImagePicker

@Composable
fun FootageScreen(
    items: LazyBulk<ImageUIDescriptor>,
    dispatchEvent: (HomeEvent) -> Unit,
) {
    Column(modifier = Modifier.fillMaxSize()) {
        var pickerShown by remember { mutableStateOf(false) }
        ElevatedButton(onClick = { pickerShown = !pickerShown }) {
            Text(text = stringResource(id = R.string.open_picker))
        }

        if (pickerShown) {
            ImagePicker { uris ->
                if (!uris.isNullOrEmpty()) {
                    dispatchEvent(HomeEvent.OnImagesPicked(uris))
                }
                pickerShown = false
            }
        } else {
            PhotoGrid(items, onApproachingWindowEnd = { index ->
                dispatchEvent(HomeEvent.OnFootageAboutToMiss(index))
            }) { dispatchEvent(HomeEvent.OnImageClicked(it)) }
        }
    }
}