package com.photomanager.photomanager.main.home.ui

import androidx.compose.runtime.Composable
import com.photomanager.photomanager.main.home.LazyBulk
import com.photomanager.photomanager.main.home.PhotoGrid

@Composable
fun CollectionsScreen(
    collection: LazyBulk<ImageUIDescriptor>,
    dispatchEvent: (HomeEvent) -> Unit
) {
    PhotoGrid(collection, onApproachingWindowEnd = { index ->
        dispatchEvent(HomeEvent.OnCollectionAboutToMiss(index))
    }) { dispatchEvent(HomeEvent.OnImageClicked(it)) }
}