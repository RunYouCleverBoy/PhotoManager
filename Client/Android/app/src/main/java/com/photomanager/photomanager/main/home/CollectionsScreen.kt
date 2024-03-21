package com.photomanager.photomanager.main.home

import androidx.compose.runtime.Composable

@Composable
fun CollectionsScreen(
    collection: LazyBulk<ImageUIDescriptor>,
    dispatchEvent: (HomeEvent) -> Unit
) {
    PhotoGrid(collection, onApproachingWindowEnd = { index ->
        dispatchEvent(HomeEvent.OnCollectionAboutToMiss(index))
    }) { dispatchEvent(HomeEvent.OnImageClicked(it)) }
}