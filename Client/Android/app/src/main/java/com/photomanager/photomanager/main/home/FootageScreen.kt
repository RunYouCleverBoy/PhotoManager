package com.photomanager.photomanager.main.home

import androidx.compose.material3.ElevatedButton
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.res.stringResource
import com.photomanager.photomanager.R
import com.photomanager.photomanager.main.components.ImagePicker

@Composable
fun FootageScreen(
    items: List<ImageUIDescriptor>,
    dispatchEvent: (HomeEvent) -> Unit,
) {
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
    }

    PhotoGrid(items, onApproachingWindowEnd = { index ->
        dispatchEvent(HomeEvent.OnApproachingFootageWindowEnd(index))
    }) { dispatchEvent(HomeEvent.OnImageClicked(it)) }
}